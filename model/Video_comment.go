package model

import (
	"fmt"
	"time"
)

type Video_comment struct {
	//gorm.Model
	ID         int64     `gorm:"not null;unique;primary_key"`
	Commenter  string    `gorm:"not null"`
	Comment    string    `gorm:"not null"`
	VideoId    string    `gorm:"not null"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	IsDelete   int       `gorm:"not null"`
}

type Video_comment_dto struct {
	//gorm.Model
	Id         int64         `json:"id,omitempty"`
	User       User_info_dto `json:"user"`
	Content    string        `json:"content,omitempty"`
	CreateDate string        `json:"create_date,omitempty"`
}

//新增评论：将用户对视频的评论加入数据表：
//输入：视频id string，用户id string，评论 string New_comment(Video_id string, Commenter_id string, Comment string)
func New_comment(Video_id string, Commenter string, Comment string) int {
	var new_cmt Video_comment
	new_cmt.VideoId = Video_id
	new_cmt.Commenter = Commenter
	new_cmt.Comment = Comment
	new_cmt.CreateTime = time.Now()
	new_cmt.IsDelete = 0
	err := db.Create(&new_cmt)
	if err != nil {
		fmt.Println("插入失败", err.Debug())
	}
	return 200
}

//获取评论：获取视频的全部评论和评论者的id列表；
//输入：视频id Get_Comment_List(Video_id string)
//返回结果：评论列表
func Get_Comment_List(Video_id string) []Video_comment {
	var cmts []Video_comment
	db.Find(&cmts)
	err := db.Model(&cmts).Select("id,commenter,comment,video_id,create_time,is_delete").Where("video_id = ?", Video_id)
	if err != nil {
		fmt.Println("查询失败", err.Debug())
	}
	return cmts
}

//更新视频的评论数量
//输入：视频id Get_and_Update_Comment_Count(Video_id string)
//返回：当前视频的评论数量
func Get_and_Update_Comment_Count(Video_id string) int {
	var usr Video_list
	cmt_list := Get_Comment_List(Video_id)
	comment_count := len(cmt_list)
	db.Model(&usr).Where("video_id = ?", Video_id).Update("comment_counts", comment_count)
	return comment_count
}

//软删除视频评论
//输入：视频id，评论id
func Soft_delete(Video_id string, Comment_id int) {
	var video_record Video_comment
	db.Model(&video_record).Where("video_id = ? and id = ?", Video_id, Comment_id).Update("is_delete", 1)
}
