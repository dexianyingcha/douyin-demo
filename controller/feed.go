package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}
// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
    user := User{}
    count :=0
    dao.DB.Where("token=?",token).Find(&user).Count(&count)
	dao.DB.Find(&DemoVideos)
    //将视频Video中的is_favorite初始化为false
    for i := 0; i < len(DemoVideos); i++{
    	dao.DB.Model(&Video{}).Update("is_favorite",false)
	}
	//匹配视频与作者
	//发布的视频有publish_token
	//根据publish_token取出对应的user结构体，赋值给DemoVideos的Author
    for j := 0; j < len(DemoVideos); j++{
    	user := User{}
    	dao.DB.Where("token=?",DemoVideos[j].PublisherToken).Find(&user)
    	DemoVideos[j].Author = user
	}
    //用户没有登录的状态下
    if count == 0{
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
	}else if count == 1{//如果有用户登录，根据用户点赞关系表正确显示当前用户是否已点赞该视频
		userId := user.Id
		favoriteVideoIdSlice := []int64{}
		dao.DB.Where("user_id=?",userId).Find(&relationsUser)
		//将当前用户点赞过的视频id放入favoriteVideoIdSlice中
		for i := 0; i < len(relationsUser); i++{
			favoriteVideoIdSlice = append(favoriteVideoIdSlice,relationsUser[i].VideoId)
		}
		//根据favoriteVideoIdSlice中的视频id，更新is_favorite为true
		for j := 0; j < len(favoriteVideoIdSlice); j++{
			dao.DB.Model(&Video{}).Where("id=?",favoriteVideoIdSlice[j]).Update("is_favorite",true)
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: DemoVideos,
			NextTime:  time.Now().Unix(),
		})
	}
}
