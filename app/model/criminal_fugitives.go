package model

import (
	"time"

	"gorm.io/gorm"
)

type CriminalFugitives struct {
	ID                  *string        `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Full_Name           *string        `json:"full_name"`
	Crime_Description   *string        `json:"crime_description"`
	Danger_Level        *int           `json:"danger_level"`
	Last_Known_Location *string        `json:"last_known_location"`
	Is_Captured         *bool          `json:"is_captured"`
	Reward_Amount       *float64       `json:"reward_amount"`
	CreatedAt           time.Time      `gorm:"column:created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
