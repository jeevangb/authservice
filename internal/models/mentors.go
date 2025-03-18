package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title           string         `gorm:"type:varchar(255);not null"`
	Description     string         `gorm:"type:text;not null"`
	TechnologyStack pq.StringArray `gorm:"type:text[]"`
	MentorName      string         `gorm:"type:varchar(50);not null;default:'N/A'"`
	Status          string         `gorm:"type:varchar(50);not null"`
}
