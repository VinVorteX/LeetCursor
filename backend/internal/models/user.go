package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	Name         string
	Picture      string
	GoogleID     string `gorm:"uniqueIndex"`
	LeetcodeID   string
	CodeforcesID string
	CodechefID   string
	TotalSolved  int
	EasySolved   int
	MediumSolved int
	HardSolved   int
	LastUpdated  time.Time
	Profiles     []CodingProfile `gorm:"foreignKey:UserID"`
}

type CodingProfile struct {
	gorm.Model
	UserID      uint
	Platform    string // "leetcode", "codeforces", "codechef"
	Username    string
	TotalSolved int
	Rating      int
	LastUpdated time.Time
}
