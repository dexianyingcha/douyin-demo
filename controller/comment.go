package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/model"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []model.Video_comment_dto `json:"comment_list,omitempty"`
}

func CommentAction(c *gin.Context) {
	// "/douyin/comment/action/?token=zhangleidouyin&video_id=1&action_type=1&comment_text=test"

	//token := c.Query("token")
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
	Video_id := c.Query("video_id")
	action_type, _ := strconv.Atoi(c.Query("action_type"))

	//发布评论
	if action_type == 1 {
		var data model.Video_comment
		//添加评论
		Comment := c.Query("comment_text")
		token := c.Query("token")
		code := model.New_comment(Video_id, token, Comment)
		//更新视频评论数量
		model.Get_and_Update_Comment_Count(Video_id)

		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": "ok",
		})

	} else if action_type == 2 {
		Commenter_id, _ := strconv.Atoi(c.Param("commenter_id"))
		model.Soft_delete(Video_id, Commenter_id)
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "删除成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"message": "未知错误",
		})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//token := c.Query("token")
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}

	id := c.Param("video_id")
	comment_list := model.Get_Comment_List(id)

	res := []model.Video_comment_dto{}
	for idx, _ := range comment_list {
		//user_id := comment_list[idx].Commenter
		//user_info := model.Get_User_Info(user_id)
		user_name := comment_list[idx].Commenter
		user_info := model.Get_User_Info_By_Username(user_name)
		res = append(res, model.Video_comment_dto{Id: comment_list[idx].ID, User: user_info, Content: comment_list[idx].Comment, CreateDate: comment_list[idx].CreateTime.Format("2006-01-02 15:04:05")})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: res,
	})
}
