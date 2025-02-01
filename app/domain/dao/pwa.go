package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"landingConstructor/app/domain/enums"
	"time"
)

type PwaCreateRequest struct {
	Name         string             `json:"name"`
	TypeCampaign enums.TypeCampaign `gorm:"column:type" json:"type"`
	Icon         string             `json:"iconUrl"`
	BaseModel
}

type Pwa struct {
	ID           uuid.UUID          `gorm:"type:uuid;primaryKey;not null;" json:"id"`
	Name         string             `gorm:"unique;not null" json:"name"`
	TypeCampaign enums.TypeCampaign `gorm:"column:type" json:"type"`
	Icon         string             `json:"iconUrl"`
	PreLandingID uuid.UUID          `gorm:"foreignKey:PwaId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"preLandingID"`
	BaseModel
}

func (p *Pwa) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}

type Header struct {
	Name            string `gorm:"column:name" json:"name"`
	IconStore       string `gorm:"column:iconStore" json:"iconStore,omitempty"`
	IconApp         string `gorm:"column:iconApp" json:"iconApp,omitempty"`
	Developer       string `gorm:"column:developer" json:"developer"`
	Subtitle        string `gorm:"column:subtitle" json:"subtitle"`
	Rating          int    `gorm:"column:rating" json:"rating"`
	NumberOfReviews int    `gorm:"column:numberOfReviews" json:"numberOfReviews"`
}

type PreLanding struct {
	ID     uuid.UUID    `gorm:"type:uuid;primaryKey;not null;" json:"id"`
	PwaId  uuid.UUID    `gorm:"type:uuid;not null" json:"pwaId"`
	Design enums.Design `gorm:"column:design" json:"design"`
	Header
	BaseModel
}

func (p *PreLanding) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}

type Description struct {
	AboutThisGame string    `json:"aboutThisGame"`
	UpdatedOn     time.Time `json:"updatedOn"`
	DataSafety    string    `json:"dataSafety"`
}

type Ratings struct {
	One   int `json:"1"`
	Two   int `json:"2"`
	Three int `json:"3"`
	Four  int `json:"4"`
	Five  int `json:"5"`
}

type Comment struct {
	Name   string    `json:"name"`
	Avatar string    `json:"avatar"`
	Rating int       `json:"rating"`
	Date   time.Time `json:"date"`
	Text   string    `json:"text"`
	People int       `json:"people"`
}

type PreLandingCreateRequest struct {
	PwaId       uuid.UUID    `json:"pwaId"`
	Design      enums.Design `json:"design"`
	Header      Header       `json:"header"`
	Screenshots []string     `json:"screenshots,omitempty"`
	Description Description  `json:"description"`
	Ratings     Ratings      `json:"ratings"`
	Comments    []Comment    `json:"comments"`
}
