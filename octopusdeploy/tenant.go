package octopusdeploy

import "github.com/go-playground/validator/v10"

type Tenants struct {
	Items []*Tenant `json:"Items"`
	PagedResults
}

type Tenant struct {
	ClonedFromTenantID  string              `json:"ClonedFromTenantId"`
	Description         string              `json:"Description"`
	Name                string              `json:"Name" validate:"required"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	SpaceID             string              `json:"SpaceId"`
	TenantTags          []string            `json:"TenantTags,omitempty"`

	resource
}

// NewTenant initializes a Tenant with a name and a description.
func NewTenant(name string, description string) *Tenant {
	return &Tenant{
		Description: description,
		Name:        name,
		resource:    *newResource(),
	}
}

// Validate checks the state of the tenant and returns an error if invalid.
func (t Tenant) Validate() error {
	return validator.New().Struct(t)
}
