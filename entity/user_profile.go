package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	Id           uuid.UUID
	FullName     string
	DateBirth    time.Time
	Gender       rune
	Bio          string
	ImageDirPath string
	LocationLat  float32
	LocationLng  float32
}
