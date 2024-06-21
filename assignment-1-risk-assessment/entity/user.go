package entity

import (
	"time"
)

type User struct {
	ID             int                 `gorm:"primaryKey" json:"id"`
	Name           string              `gorm:"type:varchar;not null" json:"name"`
	Email          string              `gorm:"type:varchar;uniqueIndex;not null" json:"email"`
	RiskScore      int                 `json:"risk_score"`
	RiskCategory   ProfileRiskCategory `json:"risk_category"`
	RiskDefinition string              `json:"risk_definition"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}
