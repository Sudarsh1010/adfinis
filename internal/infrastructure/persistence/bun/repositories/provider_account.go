package repositories

import (
	"context"
	"fmt"

	"github.com/sudarsh1010/adfinis/internal/domain/provider_account"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/models"
	"github.com/uptrace/bun"
)

// ProviderAccountRepository implements persistence for ProviderAccount domain entity.
// This repository maps between domain.ProviderAccount entities and models.ProviderAccount database models
// using the Bun ORM.
type ProviderAccountRepository struct {
	db *bun.DB
}

// NewProviderAccountRepository creates a new ProviderAccountRepository with the given database connection.
func NewProviderAccountRepository(db *bun.DB) *ProviderAccountRepository {
	return &ProviderAccountRepository{db: db}
}

// toProviderAccountDomainEntity converts a models.ProviderAccount to a domain.ProviderAccount entity.
func toProviderAccountDomainEntity(m *models.ProviderAccount) *provider_account.ProviderAccount {
		pa := &provider_account.ProviderAccount{}
		pa.SetID(provider_account.ID(m.ID))
		pa.SetWorkspaceID(string(m.WorkspaceID))
		pa.SetProvider(m.Provider)
		pa.SetName(m.Name)
		pa.SetAPIKeyEncrypted(m.APIKeyEncrypted)
		pa.SetExtraConfigJSON(m.ExtraConfigJSON)
		pa.SetCreatedAt(m.CreatedAt)
		return pa
}

// toModel converts a domain.ProviderAccount entity to a database model.
func toModel(pa *provider_account.ProviderAccount) *models.ProviderAccount {
	return &models.ProviderAccount{
		ID:              string(pa.ID()),
		WorkspaceID:     string(pa.WorkspaceID()),
		Provider:        pa.Provider(),
		Name:            pa.Name(),
		APIKeyEncrypted: pa.APIKeyEncrypted(),
		ExtraConfigJSON: pa.ExtraConfigJSON(),
		CreatedAt:       pa.CreatedAt(),
	}
}

// Save persists a provider account entity to the database.
// Converts the domain entity to a database model and inserts it.
func (r *ProviderAccountRepository) Save(ctx context.Context, pa *provider_account.ProviderAccount) error {
	model := toModel(pa)

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

// FindByID retrieves a provider account by its ID from the database.
// Returns ErrProviderAccountNotFound if the provider account does not exist.
func (r *ProviderAccountRepository) FindByID(ctx context.Context, id provider_account.ID) (*provider_account.ProviderAccount, error) {
	var model models.ProviderAccount
	err := r.db.NewSelect().Model(&model).Where("id = ?", string(id)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("provider account not found: %w", err)
	}

			return toProviderAccountDomainEntity(&model), nil
}

// FindAll retrieves all provider accounts from the database.
func (r *ProviderAccountRepository) FindAll(ctx context.Context) ([]*provider_account.ProviderAccount, error) {
	var models []models.ProviderAccount
	err := r.db.NewSelect().Model(&models).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find all provider accounts: %w", err)
	}

	providerAccounts := make([]*provider_account.ProviderAccount, len(models))
	for i := range models {
				providerAccounts[i] = toProviderAccountDomainEntity(&models[i])
	}
	return providerAccounts, nil
}

// Update updates a provider account in the database.
// Returns ErrProviderAccountNotFound if the provider account does not exist.
func (r *ProviderAccountRepository) Update(ctx context.Context, pa *provider_account.ProviderAccount) error {
	model := toModel(pa)

	result, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", string(pa.ID())).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provider account: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return provider_account.ErrProviderAccountNotFound
	}

	return nil
}

// Delete removes a provider account from the database by its ID.
// Returns ErrProviderAccountNotFound if the provider account does not exist.
func (r *ProviderAccountRepository) Delete(ctx context.Context, id provider_account.ID) error {
	result, err := r.db.NewDelete().
		Model(&models.ProviderAccount{}).
		Where("id = ?", string(id)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete provider account: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return provider_account.ErrProviderAccountNotFound
	}

	return nil
}
