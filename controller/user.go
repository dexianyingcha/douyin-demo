package controller

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin


var userIdSequence = int64(1)


type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}


//将用户信息存入usersLoginInfo,用于用户登录校验，同时将用户信息存入数据库。
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username+password

    dao.DB.Where("token=?",token).Find(&user).Count(&count)
	if count == 1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		//如果查询到不存在，则往数据库里添加对应的用户信息
	} else if count == 0 {
		newUser := User{
			Name: username,
			FollowCount: 0,
			FollowerCount: 0,
			IsFollow: false,
			Token: token,
		}
		dao.DB.Create(&newUser)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	user := User{}//需要user := User{}，才能重置count
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
    dao.DB.Where("token=?",token).Find(&user).Count(&count)
	if count == 0{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}else if count == 1{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    user.Token,
		})
	}
}

func UserInfo(c *gin.Context) {
	user := User{}
	token := c.Query("token")

	dao.DB.Where("token=?",token).Find(&user).Count(&count)
	if count == 0{
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}else if count == 1{
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	}
}
