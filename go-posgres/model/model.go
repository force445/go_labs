package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID  `gorm:"type:uuid;"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Phone    string     `json:"phone"`
	Location []Location `gorm:"foreignKey:UserID"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	user.ID = uuid
	return
}

type Location struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;"`
	UserID  uuid.UUID `gorm:"type:uuid;"`
	Address string    `json:"address"`
	City    string    `json:"city"`
	State   string    `json:"state"`
	Zip     string    `json:"zip"`
}

type Locations struct {
	Locations []Location `json:"locations"`
}

func (location *Location) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	location.ID = uuid
	return
}

// One station has many drones
type Drone struct {
	gorm.Model
	ID              uuid.UUID    `gorm:"type:uuid;"`
	DroneID         string       `json:"drone_id"`
	SerialNo        string       `json:"serial_no"`
	LimitWeight     float64      `json:"limit_weight"`
	MaintenanceDate time.Weekday `json:"maintenance_date"`
	StationID       uuid.UUID    `gorm:"type:uuid;"`
}

type Station struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	StationID string    `json:"station_id"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	Drone     []Drone   `gorm:"foreignKey:StationID"`
}
