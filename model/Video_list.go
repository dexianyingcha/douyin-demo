package model

import "time"

type Video_list struct {
	//gorm.Model
	ID              int    `gorm:"not null;unique;primary_key"`
	User_id         string `gorm:"not null;unique"`
	Video_id        string `gorm:"not null;unique"`
	Video_desc      string `gorm:"not null;unique"`
	Play_url        string `gorm:"not null;unique"`
	Cover_url       string `gorm:"not null"`
	Favorite_counts int
	Comment_counts  int
	Creat_time      time.Time `gorm:"autoCreateTime"`
}
