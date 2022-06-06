package model

import (
	"fmt"
	"strconv"
)

type User_info struct {
	//gorm.Model
	ID            int    `gorm:"not null;unique;primaryKey;autoIncrement:true"`
	User_id       string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	Username      string `gorm:"not null"`
	Fans_counts   int
	Follow_counts int
}

type User_info_dto struct {
	//gorm.Model
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func Get_User_Info(User_id string) User_info_dto {
	var user User_info
	result := db.Select("user_id,username,fans_counts,follow_counts").Where("user_id = ?", User_id).First(&user)
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	id, _ := strconv.ParseInt(user.User_id, 10, 64)
	return User_info_dto{Id: id, Name: user.Username, FollowCount: int64(user.Follow_counts), FollowerCount: int64(user.Fans_counts), IsFollow: false}
}

func Get_User_Info_By_Username(Username string) User_info_dto {
	var user User_info
	result := db.Select("user_id,username,fans_counts,follow_counts").Where("username = ?", Username).First(&user)
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	id, _ := strconv.ParseInt(user.User_id, 10, 64)
	return User_info_dto{Id: id, Name: user.Username, FollowCount: int64(user.Follow_counts), FollowerCount: int64(user.Fans_counts), IsFollow: false}
}
