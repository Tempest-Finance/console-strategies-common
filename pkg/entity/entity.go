package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseID struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseID) BeforeCreate(tx *gorm.DB) error {
	if base.ID != uuid.Nil {
		return nil
	}
	base.ID = uuid.New()
	return nil
}

type BaseCreatedUpdated struct {
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updatedAt" gorm:"autoUpdateTime"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
