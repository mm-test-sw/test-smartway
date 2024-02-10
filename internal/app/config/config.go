package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"test-smartway/internal/entity"
	"time"
)

type Config struct {
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout" validate:"required"`

	Server `yaml:"server" validate:"required"`

	DataBase `yaml:"dataBase" validate:"required"`

	DemoAccountAirlines []entity.Airline `yaml:"demoAccountAirlines"`
}

type Server struct {
	Port           string        `yaml:"port" validate:"required"`
	WriteTimeout   time.Duration `yaml:"writeTimeout" validate:"required"`
	ReadTimeout    time.Duration `yaml:"readTimeout" validate:"required"`
	IdleTimeout    time.Duration `yaml:"idleTimeout" validate:"required"`
	RequestTimeout time.Duration `yaml:"requestTimeout"  validate:"required"`
}

type DataBase struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"dBName" validate:"required"`
	Params   string `yaml:"params"`
}

func NewConfig() *Config {
	cfg := new(Config)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	if err := validator.New().Struct(cfg); err != nil {
		panic(err.Error())
	}

	return cfg
}
