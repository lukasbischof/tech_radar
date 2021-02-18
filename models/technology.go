package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"time"
)

// Technology is used by pop to map your technologies database table to your go code.
type Technology struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `json:"name" db:"name"`
	Category  string    `json:"category" db:"category"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Technology) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Technologies is not required by pop and may be deleted
type Technologies []Technology

func (t Technologies) GroupTechnologies() (groups map[string]Technologies) {
	groups = map[string]Technologies{}

	for _, technology := range t {
		if list, ok := groups[technology.Category]; ok {
			groups[technology.Category] = append(list, technology)
		} else {
			groups[technology.Category] = []Technology{technology}
		}
	}

	return groups
}

// String is not required by pop and may be deleted
func (t Technologies) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Technology) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Technology) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Technology) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
