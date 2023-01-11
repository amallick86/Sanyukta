package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type SubscriptionTypes struct {
	ID   int    `json:"ID"`
	Type string `json:"Type"`
}
type Subscription struct {
	ID                int               `json:"ID"`
	SubscriptionTypes SubscriptionTypes `json:"SubscriptionTypes"`
	Price             int               `json:"Price"`
	Duration          string            `json:"Duration"`
}
type User struct {
	UID            uuid.UUID `json:"UID" `
	FirstName      string    `json:"FirstName" `
	LastName       string    `json:"LastName" `
	Email          string    `json:"Email" `
	UserName       string    `json:"UserName" `
	Password       string    `json:"Password" `
	IsVerified     bool      `json:"IsVerified" `
	GoogleID       string    `json:"GoogleID" `
	SubscriptionID int       `json:"SubscriptionID" `
	CreatedAt      time.Time `json:"CreatedAt" `
}
type EmailOTP struct {
	ID         sql.NullInt64 `json:"ID"`
	UID        uuid.UUID     `json:"UID" `
	FullName   string        `json:"FullName" `
	OTP        sql.NullInt64 `json:"OTP"`
	CountOfDay sql.NullInt16 `json:"CountOfDay"`
	CreatedAt  sql.NullTime  `json:"CreatedAt" `
	ExpiredAt  sql.NullTime  `json:"ExpiredAt" `
}
type UserDetails struct {
	UDID            int       `json:"UDID" `
	UID             uuid.UUID `json:"UID" `
	Bio             string    `json:"Bio" `
	DateOfBirth     time.Time `json:"DateOfBirth" `
	Gender          int       `json:"Gender" `
	Child           int       `json:"Child" `
	Siblings        int       `json:"Siblings" `
	Height          string    `json:"Height" `
	Religion        int       `json:"Religion" `
	Caste           int       `json:"Caste" `
	Zodiac          int       `json:"Zodiac" `
	MaritalStatus   int       `json:"MaritalStatus" `
	HigherEducation int       `json:"HigherEducation" `
	Profession      int       `json:"Profession"`
	HomeCountry     int       `json:"HomeCountry" `
	HomeState       int       `json:"HomeState" `
	ResidingCountry int       `json:"ResidingCountry" `
	ResidingState   int       `json:"ResidingState" `
	ResidingStatus  int       `json:"ResidingStatus" `
	CreatedAt       time.Time `json:"CreatedAt" `
	UpdatedAt       time.Time `json:"UpdatedAt" `
}

type Sessions struct {
	ID           uuid.UUID `json:"ID" `
	UID          uuid.UUID `json:"UID" `
	RefreshToken string    `json:"RefreshToken"`
	UserAgent    string    `json:"UserAgent" `
	ClientIP     string    `json:"ClientIP" `
	IsBlocked    bool      `json:"IsBlocked" `
	ExpiresAt    time.Time `json:"ExpiresAt" `
	DeviceID     string    `json:"DeviceID"`
	CreatedAt    time.Time `json:"CreatedAt" `
}
