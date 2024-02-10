package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	IDStr     uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Model) BeforeCreate(tx *gorm.DB) (err error) {
	base.IDStr = uuid.NewV4()
	return
}

type User struct {
	Model
	Username string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	NIK      string `gorm:"not null"`
	Prefix   string `gorm:"not null;default:'+62'"`
	Password string `gorm:"not null"`
	Photo    *string

	MerchantID *uint
	Merchant   Merchant

	Role uint `gorm:"not null;"`
}

type Merchant struct {
	Model
	Name string `gorm:"not null"`

	RunningText *string
	Address     string `gorm:"not null"`
	Phone       string `gorm:"not null"`
	PICName     string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Photo       *string

	MerchantCategoryID uint `gorm:"not null"`
	MerchantCategory   MerchantCategory
}

type MerchantCategory struct {
	Model
	Name string `gorm:"not null"`
}

type Location struct {
	Model
	Name string `gorm:"not null"`
}

type Queue struct {
	Model
	LocationID   uint
	Location     Location
	LocationName string

	QueueNo       string `gorm:"not null"`
	MedicalRecord string `gorm:"not null"`
	Type          uint   `gorm:"not null"` // 1 = non racikan, 2 = racikan

	MerchantID   uint `gorm:"not null"`
	Merchant     Merchant
	MerchantName string

	// ----- who create this queue
	UserID   uint `gorm:"not null"`
	User     User
	UserName string

	IsFollowUp    bool `gorm:"not null;default:0"`
	FollowUpPhone *string

	Histories []QueueHistory
}

type QueueHistory struct {
	gorm.Model

	Status     uint `gorm:"not null"` // 1 = validasi, 2 = proses, 3 = siap diserahkan, 4 = diserahkan
	StartQueue *time.Time
	Duration   *float64
	EndQueue   *time.Time // sum of start + duration

	Type uint `gorm:"not null;default:1"` // 1 = default, 2 = extend

	QueueID uint `gorm:"not null"`
	Queue   Queue

	UserID   uint `gorm:"not null"`
	User     User
	UserName string
}

type AdvertisementCategory struct {
	Model
	Name        string `gorm:"not null"`
	Description *string
}

type Advertisement struct {
	Model
	Name    string `gorm:"not null"`
	Company string `gorm:"not null"`

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant

	AdvertisementCategoryID uint `gorm:"not null"`
	AdvertisementCategory   AdvertisementCategory

	DocumentPath string         `gorm:"not null"`
	DateStart    datatypes.Date `gorm:"not null"`
	DateEnd      datatypes.Date `gorm:"not null"`

	Description *string
}
