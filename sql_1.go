package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 首先用sql建好表，每一张表以结构体的形式列出
// 建立结构体，对应映射关系

type User_fans struct {
	//gorm.Model
	ID      int    `gorm:"not null;unique;primary_key"`
	User_id string `gorm:"not null"`
	Fans_id string `gorm:"not null"`
}
type User_follow struct {
	//gorm.Model
	ID            string `gorm:"not null;unique"`
	Usr_id        string `gorm:"not null"`
	Usr_follow_id string
}
type User_info struct {
	//gorm.Model
	ID            int    `gorm:"not null;unique;primaryKey;autoIncrement:true"`
	User_id       string `gorm:"not null;unique"`
	Password      string `gorm:"not null"`
	Username      string `gorm:"not null"`
	Fans_counts   int
	Follow_counts int
}

type User_like struct {
	//gorm.Model
	ID       int    `gorm:"not null;unique;primary_key;autoincrement:true"`
	User_id  string `gorm:"not null"`
	Video_id string `gorm:"not null"`
}

type Video_comment struct {
	//gorm.Model
	ID        int    `gorm:"not null;unique;primary_key"`
	Commenter string `gorm:"not null"`
	Comment   string `gorm:"not null"`
	Video_id  string `gorm:"not null"`
}

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

// 连接数据库
func Conn() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:wtt991221@tcp(192.168.1.235:3306)/douyin") //实验室的ip
	//db, err := gorm.Open("mysql", "root:wtt991221@tcp(172.26.84.206:3306)/douyin") //宿舍的ip
	db, err := gorm.Open("mysql", "root:wtt991221@tcp(172.19.37.110:0)/douyin") //教学楼的ip

	if err != nil {
		fmt.Print("sql open error", err)
		return nil
	}
	db.SingularTable(true)
	return db
}

func Update_Follow_Count(db *gorm.DB, user_id string, operation bool) {
	var usr User_info
	if operation {
		db.Model(&usr).Where("user_id = ?", user_id).Update("follow_counts", gorm.Expr("follow_counts+1"))
	} else {
		db.Model(&usr).Where("user_id = ?", user_id).Update("follow_counts", gorm.Expr("follow_counts-1"))
	}
	//db.Model(&usr).Where().Update("follow_counts", usr.Follow_counts+1)
}

func Update_Fans_Count(db *gorm.DB, fans_id string, operation bool) {
	var usr User_info
	if operation {
		db.Model(&usr).Where("user_id = ?", fans_id).Update("fans_counts", gorm.Expr("fans_counts+1"))
	} else {
		db.Model(&usr).Where("user_id = ?", fans_id).Update("fans_counts", gorm.Expr("fans_counts-1"))
	}
}

func Get_and_Update_Follow_count(db *gorm.DB, user_id string) int {
	// var follow_list []string
	var usr User_info
	follow_list := Get_Follow_List(db, user_id)
	follow_count := len(follow_list)
	db.Model(&usr).Where("user_id = ?", user_id).Update("follow_counts", follow_count)
	return follow_count
}

func Get_and_Update_Fans_count(db *gorm.DB, user_id string) int {
	// var follow_list []string
	var usr User_info
	fans_list := Get_Fans_List(db, user_id)
	fans_count := len(fans_list)
	db.Model(&usr).Where("user_id = ?", user_id).Update("fans_counts", fans_count)
	return fans_count
}

func New_fans(db *gorm.DB, User_id string, fans_id string) {
	var new_usr_fans User_fans
	//new_usr_fans.ID = id
	new_usr_fans.User_id = User_id
	new_usr_fans.Fans_id = fans_id
	Update_Fans_Count(db, fans_id, true)
	Update_Follow_Count(db, User_id, true)
	if err1 := db.Create(&new_usr_fans).Error; err1 != nil {
		fmt.Println("插入失败", err1)
	}
}

func Get_Follow_List(db *gorm.DB, User_id string) []string {
	var fans []User_fans
	var flrs_list []string
	result := db.Select("fans_id").Where("user_id = ?", User_id).Find(&fans)
	for _, usr := range fans {
		flrs_list = append(flrs_list, usr.Fans_id)
	}
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	// fmt.Println(fans_list)
	return flrs_list
}

func Get_Fans_List(db *gorm.DB, User_id string) []string {
	var fans []User_fans
	var fans_list []string
	result := db.Select("user_id").Where("fans_id = ?", User_id).Find(&fans)
	for _, fan := range fans {
		fans_list = append(fans_list, fan.User_id)
	}
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	// fmt.Println(fans_list)
	return fans_list
}

func New_Video(db *gorm.DB, User_id string, Video_id string, Video_url string) {
	var new_video Video_list
	if Video_url == "" { //Go 不允许定义默认参数，只能出此下策
		new_video.Play_url = "default"
	}
	new_video.User_id = User_id
	new_video.Video_id = Video_id
	new_video.Creat_time = time.Now()
	if err := db.Create(&new_video).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func Delete_Fans(db *gorm.DB, User_id string, fans_id string) {
	var user User_fans
	result := db.Where("fans_id = ?", fans_id).Delete(&user)
	Update_Fans_Count(db, fans_id, false)
	Update_Follow_Count(db, User_id, false)
	if result.Error != nil {
		fmt.Println("删除失败", result.Error)
	}
}

func New_Usr(db *gorm.DB, Usr_id string, Usr_name string, Password string) {
	var new_usr User_info
	new_usr.User_id = Usr_id
	new_usr.Username = Usr_name
	new_usr.Password = Password
	if err := db.Create(&new_usr).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func Get_and_Update_Favorite_Count(db *gorm.DB, Video_id string) int {
	var usr Video_list
	liker_list := Get_Favorite_List(db, Video_id)
	favorite_count := len(liker_list)
	db.Model(&usr).Where("video_id = ?", Video_id).Update("favorite_counts", favorite_count)
	return favorite_count
}

func Like_This_Video(db *gorm.DB, User_id string, Video_id string) {
	var new_like User_like
	new_like.User_id = User_id
	new_like.Video_id = Video_id
	if err := db.Create(&new_like).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func Unlike_This_Video(db *gorm.DB, User_id string, Video_id string) {
	var like User_like
	result := db.Where("video_id = ?", Video_id).Delete(&like)
	if result.Error != nil {
		fmt.Println("删除失败", result.Error)
	}
}

func Get_Like_List(db *gorm.DB, User_id string) []string {
	var likes []User_like
	var likes_list []string
	result := db.Select("video_id").Where("user_id = ?", User_id).Find(&likes)
	for _, like := range likes {
		likes_list = append(likes_list, like.Video_id)
	}
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	// fmt.Println(fans_list)
	return likes_list
}

func Get_Favorite_List(db *gorm.DB, Video_id string) []string {
	var likers []User_like
	var liker_list []string
	result := db.Select("user_id").Where("Video_id = ?", Video_id).Find(&likers)
	for _, liker := range likers {
		liker_list = append(liker_list, liker.User_id)
	}
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	// fmt.Println(fans_list)
	return liker_list
}
func New_comment(db *gorm.DB, Video_id string, Commenter_id string, Comment string) {
	var new_cmt Video_comment
	new_cmt.Video_id = Video_id
	new_cmt.Commenter = Commenter_id
	new_cmt.Comment = Comment
	if err := db.Create(&new_cmt).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func Get_Comment_List(db *gorm.DB, Video_id string) ([]string, []string) {
	var cmts []Video_comment
	var cmter_list []string
	var cmt_list []string
	result := db.Select("commenter", "comment").Where("Video_id = ?", Video_id).Find(&cmts)

	for _, cmt := range cmts {
		cmter_list = append(cmter_list, cmt.Commenter)
		cmt_list = append(cmt_list, cmt.Comment)
	}
	if result.Error != nil {
		fmt.Println("查询失败", result.Error)
	}
	// fmt.Println(fans_list)
	return cmter_list, cmt_list
}

func Get_and_Update_Comment_Count(db *gorm.DB, Video_id string) int {
	var usr Video_list
	_, cmt_list := Get_Comment_List(db, Video_id)
	comment_count := len(cmt_list)
	db.Model(&usr).Where("video_id = ?", Video_id).Update("comment_counts", comment_count)
	return comment_count
}

// 获取某一列的所有数据
func GetallData(db *gorm.DB, something string) *gorm.DB {
	var usrs []User_info
	return db.Select(something).Find(&usrs)
}

func main() {
	conn := Conn()
	conn.AutoMigrate(&User_fans{}, &User_follow{}, &User_info{}, &User_like{}, &Video_comment{}, &Video_list{})
	//New_comment(conn, "1's Vid", "1", "this is cmt1.")
	// Unlike_This_Video(conn, "1", "1's Vid")
	// fmt.Println(Get_Favorite_List(conn, "1's Vid"))
	// fmt.Println(Get_and_Update_Favorite_Count(conn, "1's Vid"))
	// fmt.Println(Get_Like_list(conn, "1"))
	// New_Video(conn, "4", "4's Vid")
	// Like_This_Video(conn, "4", "2's Vid")
	// New_Usr(conn, "4", "C++", "444")
	// conn.AutoMigrate(&User_fans{})
	// New_fans(conn, "4", "1")
	// fmt.Println(Get_Fans_List(conn, "1"))
	// fmt.Println(Get_Follow_List(conn, "1"))
	// fmt.Println(Get_and_Update_Fans_count(conn, "1"))
	// fmt.Println(Get_and_Update_Follow_count(conn, "1"))
	// Delete_Fans(conn, "2", "1")
	defer conn.Close()
}
