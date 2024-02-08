package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nanmu42/gzip"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"test-smartway/internal/api/handler"
	"test-smartway/internal/api/middleware"
	"test-smartway/internal/app/config"
	"test-smartway/internal/entity"
	"test-smartway/internal/repository"
	"test-smartway/internal/service"
)

func Run(ctx context.Context, cfg *config.Config) {
	closer := newCloser()
	logger := newLogger()
	router := mux.NewRouter()
	db := newDataBase(ctx, logger, cfg)

	defer db.Close()

	// Repository
	airlineRepository := repository.NewAirlineRepository(db)
	providerRepository := repository.NewProviderRepository(db)
	schemaRepository := repository.NewSchemaRepository(db)
	accountRepository := repository.NewAccountRepository(db)

	// Service
	airlineService := service.NewAirlineService(airlineRepository)
	providerService := service.NewProviderService(providerRepository)
	schemaService := service.NewSchemaService(schemaRepository)
	accountService := service.NewAccountService(cfg, accountRepository)

	// API
	mw := middleware.NewMiddleware(logger)
	handler.RegisterAirlineHandlers(router, airlineService, logger, mw)
	handler.RegisterProviderHandlers(router, providerService, logger, mw)
	handler.RegisterSchemaHandlers(router, schemaService, logger, mw)
	handler.RegisterAccountHandlers(router, accountService, logger, mw)

	go func() {
		logger.DPanic("ListenAndServe упал", zap.Any("Error", configServer(cfg, router).ListenAndServe()))
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		logger.Error("Close err", zap.Error(err))
	}
}

func configServer(cfg *config.Config, router http.Handler) *http.Server {
	return &http.Server{
		Handler:        gzip.DefaultHandler().WrapHandler(router),
		Addr:           ":" + cfg.Port,
		WriteTimeout:   cfg.WriteTimeout,
		ReadTimeout:    cfg.ReadTimeout,
		IdleTimeout:    cfg.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
}

func newDataBase(ctx context.Context, logger *zap.Logger, cfg *config.Config) *pgxpool.Pool {

	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?%s",
		cfg.Username, cfg.Password, cfg.Address, cfg.DBName, cfg.Params)

	logger.Info("start establishing a connection to the database", zap.String("Connect", conn))

	db, err := pgxpool.New(ctx, conn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping(ctx)
	if err != nil {
		panic(err.Error())
	}

	logger.Info("successful connection to the database")

	return db
}

func newLogger() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = zap.New(logger.Core(), zap.AddCaller(), zap.AddCallerSkip(1))

	logger.Info("Zap Logger", zap.String("Level", logger.Level().String()))

	return logger
}

func newCloser() *entity.Closer {
	return &entity.Closer{}
}
