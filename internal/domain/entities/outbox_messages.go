package entities

import (
	"time"
)

type OutboxMessage struct {
	Id        int64     `gorm:"column:id;primaryKey"`
	Message   string    `gorm:"column:message"`
	CreatedOn time.Time `gorm:"column:created_on"`
}
