package service

import (
	"context"
	"errors"
	"test-smartway/internal/entity"
)

type schemaService struct {
	schemaRepo entity.ISchemaRepository
}

func NewSchemaService(schemaRepo entity.ISchemaRepository) entity.ISchemaService {
	return &schemaService{schemaRepo: schemaRepo}
}

func (s *schemaService) AddSchema(ctx context.Context, schema *entity.Schema) (*entity.Schema, error) {
	return s.schemaRepo.InsertSchema(ctx, schema)
}

func (s *schemaService) GetSchema(ctx context.Context, name string) (*entity.Schema, error) {
	return s.schemaRepo.SelectSchemaByName(ctx, name)
}

func (s *schemaService) PatchSchema(ctx context.Context, schema *entity.Schema) (*entity.Schema, error) {
	if len(schema.Name) == 0 && schema.Providers == nil {
		return schema, nil
	}

	tx, err := s.schemaRepo.GetTx(ctx)
	if err != nil {
		return nil, err
	}

	if len(schema.Name) != 0 {
		err = s.schemaRepo.TxUpdateSchemaName(ctx, tx, schema.Id, schema.Name)
		if err != nil {
			return nil, err
		}
	}

	if schema.Providers != nil {
		err = s.schemaRepo.TxReplaceSchemaProviders(ctx, tx, schema.Id, schema.Providers)
		if err != nil {
			return nil, err
		}
	}

	schema, err = s.schemaRepo.TxSelectSchema(ctx, tx, schema.Id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func (s *schemaService) DeleteSchema(ctx context.Context, id string) error {
	isSchemeAssignedToAccount, err := s.schemaRepo.IsSchemeAssignedToAccount(ctx, id)
	if err != nil {
		return err
	}

	if isSchemeAssignedToAccount {
		return entity.NewLogicError(errors.New("scheme assigned to account"), "scheme assigned to account", 400)
	}

	err = s.schemaRepo.DeleteSchema(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
