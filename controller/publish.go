package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}


// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	user := User{}
	token := c.Query("token")

	dao.DB.Where("token=?",token).Find(&user).Count(&count)
	if count == 0{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}else if count == 1{
		data, err := c.FormFile("data")
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}

		filename := filepath.Base(data.Filename)
		finalName := fmt.Sprintf("%d_%s", user.Id, filename)
		saveFile := filepath.Join("./public/", finalName)
		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}

		video :=Video{
			Author: 	user,
			PlayUrl:  	"http://192.168.1.127:8080/static/" + finalName,
			CoverUrl: 	"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount:  	0,
			CommentCount: 0,
			IsFavorite: false,
			PublisherToken: token,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
        dao.DB.Create(&video)
		c.JSON(http.StatusOK, Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded successfully",
		})
	}


}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	user := User{}
	token := c.Query("token")

	dao.DB.Where("publisher_token=?",token).Find(&DemoVideos)

	for i := 0; i < len(DemoVideos); i++{
		DemoVideos[i].Author = user
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
