package provider_account

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// ID is a ULID wrapper (value object).
type ID string

func NewProviderAccountID() ID {
	return ID(ulid.Make().String())
}

// ProviderAccount is a pure domain entity (NO Bun dependencies).
type ProviderAccount struct {
	id              ID
	workspaceID     string
	provider        string
	name            string
	apiKeyEncrypted string
	extraConfigJSON string
	createdAt       time.Time
	updatedAt       time.Time
}

func NewProviderAccount(workspaceID, provider, name, apiKeyEncrypted string) *ProviderAccount {
	return &ProviderAccount{
		id:              NewProviderAccountID(),
		workspaceID:     workspaceID,
		provider:        provider,
		name:            name,
		apiKeyEncrypted: apiKeyEncrypted,
		extraConfigJSON: "",
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}
}

func (p *ProviderAccount) ID() ID {
	return p.id
}

func (p *ProviderAccount) WorkspaceID() string {
	return p.workspaceID
}

func (p *ProviderAccount) Provider() string {
	return p.provider
}

func (p *ProviderAccount) Name() string {
	return p.name
}

func (p *ProviderAccount) APIKeyEncrypted() string {
	return p.apiKeyEncrypted
}

func (p *ProviderAccount) ExtraConfigJSON() string {
	return p.extraConfigJSON
}

func (p *ProviderAccount) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProviderAccount) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *ProviderAccount) SetID(id ID) {
	p.id = id
}

func (p *ProviderAccount) SetWorkspaceID(workspaceID string) {
	p.workspaceID = workspaceID
}

func (p *ProviderAccount) SetProvider(provider string) {
	p.provider = provider
}

func (p *ProviderAccount) SetName(name string) {
	p.name = name
}

func (p *ProviderAccount) SetAPIKeyEncrypted(apiKeyEncrypted string) {
	p.apiKeyEncrypted = apiKeyEncrypted
}

func (p *ProviderAccount) SetExtraConfigJSON(extraConfigJSON string) {
	p.extraConfigJSON = extraConfigJSON
}

func (p *ProviderAccount) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *ProviderAccount) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}
