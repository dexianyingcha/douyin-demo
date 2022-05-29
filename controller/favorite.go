package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

type ClickRelation struct{
	Id         int64 `json:"id,omitempty"`
	UserId   int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}




// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	user := User{}
	clickRelation := ClickRelation{}
	token := c.Query("token")

	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")
    dao.DB.Where("token=?",token).Find(&user).Count(&count)
	if count == 0{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}else if count ==1{
		var userId int64 = user.Id
		videoId, _ := strconv.Atoi(videoIdStr)
		if actionType == "1" {
			//1表示点赞
			//把当前用户userId和视频videoId添加到点赞列表里
			r := ClickRelation{
				UserId:   userId,
				VideoId: int64(videoId),
			}
			//gorm 增加一行

			dao.DB.Where("user_id = ? and video_id = ?",userId, videoId).Find(&clickRelation).Count(&count)
			if count==0{
				fmt.Println("点赞成功!")
				dao.DB.Create(&r)
				dao.DB.Model(&Video{}).Where("id = ?", videoId).Update("is_favorite", true)
			}

			//修改is_favorite字段为true，表示已点赞


		} else {
			//2表示取消点赞
			//把当前用户userId和视频videoId从点赞列表里删除
			dao.DB.Where("user_id = ? and video_id = ?", userId, videoId).Delete(ClickRelation{})
			//修改is_favorite字段为false，表示未点赞
			dao.DB.Model(&Video{}).Where("id = ?", videoId).Update("is_favorite", false)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}


// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	//token := c.Query("token")

	dao.DB.Where("id=?",userId).Find(&userList)
	//先获得当前用户点赞视频的id列表videosList
    for i := 0; i < len(userList); i++{
    	videosList =append(videosList,userList[i].VideoId)
	}
	//根据视频id将视频video存入点赞视频列表DemoVideos
    for j :=0; j< len(videosList); j++{
    	var video Video
    	dao.DB.Where("id=?",videosList[j]).Find(&video)
    	DemoVideos =append(DemoVideos,video)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
