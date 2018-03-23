package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Weibo struct {
	gorm.Model
	Name    string `gorm:"size:255;index:name"`
	Content string `gorm:"not null;size:255"`
	//Images  []string
	Time   time.Time
	Client string
}

type WeiboForward struct {
	gorm.Model
	Name    string `gorm:"size:255;index:name"`
	Comment string `gorm:"size:255"`
	Content string `gorm:"not null;size:255"`
	//Images  []string
	Time   time.Time
	Client string
}
