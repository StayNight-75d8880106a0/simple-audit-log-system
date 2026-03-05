package model

import (
	"time"

	"gorm.io/datatypes"
)

type AuditLogs struct {
	ID          *string        `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Entity_Name *string        `json:"entity_name"`
	Entity_ID   *string        `json:"entity_id"`
	Action      *string        `json:"action"`
	Old_Data    datatypes.JSON `json:"old_data"`
	New_Data    datatypes.JSON `json:"new_data"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
}
