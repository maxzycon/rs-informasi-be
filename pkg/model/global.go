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

	Role uint `gorm:"not null"`
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

type Floor struct {
	Model
	Name string `gorm:"not null"`

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type Facility struct {
	Model
	Name  string `gorm:"not null"`
	Desc  *string
	Photo *string

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type Services struct {
	Model
	Name  string `gorm:"not null"`
	Desc  *string
	Photo *string

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type ProductCategory struct {
	Model
	Name string `gorm:"not null"`

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type Information struct {
	Model
	Name                  string `gorm:"not null"`
	Desc                  *string
	Photo                 *string
	InformationCategoryID uint
	InformationCategory   InformationCategory

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type InformationCategory struct {
	Model
	Name string `gorm:"not null"`

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type Product struct {
	Model
	Name              string `gorm:"not null"`
	ProductCategoryID uint
	ProductCategory   ProductCategory

	Price             float64 `gorm:"not null"`
	IsDiscount        bool    `gorm:"not null;default:0"`
	AmountDiscount    *float64
	DiscountStartDate *datatypes.Date
	DiscountEndDate   *datatypes.Date

	Photo  *string
	Detail []DetailProduct

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type DetailProduct struct {
	Model
	Name      string `gorm:"not null"`
	ProductID uint
	Product   Product
}

type Organ struct {
	Model
	Name string `gorm:"not null"`
}

type Doctor struct {
	Model
	Name             string `gorm:"not null"`
	SpecializationID uint   `gorm:"not null"`
	Specialization   Specialization

	Skill     []DoctorSkill
	Education []DoctorEducation
	Slot      []DoctorSlot

	Photo *string

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}

type DoctorSkill struct {
	Model
	DoctorID uint `gorm:"not null"`
	Doctor   Doctor
	Name     string `gorm:"not null"`
}

type DoctorSlot struct {
	Model
	DoctorID  uint `gorm:"not null"`
	Doctor    Doctor
	Day       uint   `gorm:"not null"`
	StartTime string `gorm:"not null"`
	EndTime   string `gorm:"not null"`
}

type DoctorEducation struct {
	Model
	DoctorID uint `gorm:"not null"`
	Doctor   Doctor
	Grade    string `gorm:"not null"`
	Major    string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

type Specialization struct {
	Model
	Name string `gorm:"not null"`

	OrganID uint `gorm:"not null"`
	Organ   Organ

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
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

type LogsPage struct {
	Model
	Url string `gorm:"not null"`
}

type Room struct {
	Model
	Name    string `gorm:"not null"`
	FloorID uint   `gorm:"not null"`
	Floor   Floor

	MerchantID uint `gorm:"not null"`
	Merchant   Merchant
}
