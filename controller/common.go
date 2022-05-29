package controller

import "time"

var user User                    //定义一个结构体修改用户数据，对应71行和63行，gorm写法&user会自动去user表里找
var relationsFollow []Relation   //放当前用户的结构体，根据follower来找，方便取出对方用户的ID，也就是关注对象的ID
var relationsFollower []Relation //放当前用户的结构体，根据followerid来找，方便取出对方用户的ID，也就是粉丝对象的ID
var UserListFollow []User        //根据对方用户的ID，从用户表里找出来的结构体，用来返回给软件关注列表展示
var UserListFollower []User      //根据对方用户的ID，从用户表里找出来的结构体，用来返回给软件粉丝列表展示
var toUserIdsFollow []int64      //放对方用户的ID，关注的对方用户的ID
var toUserIdsFollower []int64    //放对方用户的ID，粉丝的对方用户的ID

var relationsUser []ClickRelation  //放当前用户的结构体，根据userId来找，方便取出视频的ID
var count int
var userList []ClickRelation//存储当前用户的点赞结构体
var videosList []int64//存储当前用户点赞视频id列表
var DemoVideos []Video//当前用户点赞的视频列表

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	PublisherToken string `json:"publisher_token,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Token         string `json:"token,omitempty"`
}
