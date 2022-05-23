# 抖音项目数据库的接口文档
## 1.数据表介绍（TODO）
- 用户信息表相关接口已完成(基础模块)
- 关注表和粉丝表合并，保留粉丝表即可(扩展接口II部分，已完成);
- id 主键需要自增，用int类型;
## 2.接口介绍
### 2.0 连接数据库
首先用SQL语句建立数据库(gorm无法建库),然后建立结构体，和表自动映射;
结构体详情略，参考数据表介绍;
连接数据库:
`func Conn()` 返回结果：数据库 *gorm.DB
指定表名全为单数名：
`db.SingularTable(true)`
自动迁移：
`db.AutoMigrate(&table_name{})`
### 2.1 用户信息表(user_info)
- 新建用户:输入:用户id，用户名，用户密码
  `New_Usr(Usr_id string, Usr_name string, Password string)`;若失败，返回错误信息
- 更新用户关注数:输入：用户id string
  `Update_Follow_Count(user_id string, operation bool)`;operation:true:新增关注；false:取关
- 更新用户粉丝数:输入：用户id string
  `Update_Fans_Count(fans_id string, operation bool)`;operation:true:新增关注；false:取关
  注：以上函数在新增关注和取关时被调用
- 从粉丝表获取并更新用户关注数:输入：用户id string
  `Get_and_Update_Follow_count(user_id string)`;返回结果:用户关注数 int;
- 从粉丝表获取并更新用户粉丝数:
  `Get_and_Update_Follow_count(user_id string)`;返回结果：用户关注数 int;
  注：以上函数直接从粉丝表中获得用户的关注-粉丝数，并且同步到用户信息表;
### 2.2 粉丝-关注表(user_fans)
- 新增粉丝: 将新的粉丝加入数据，用户信息中：输入：用户id string，关注id string
`new_fans(User_id string, to_user_id string)`;若失败，返回错误信息
- 获取全部粉丝ID的列表：输入：用户id string
`Get_Fans_List(User_id string)` ；返回结果：用户关注id列表 []string;若失败，返回错误信息
- 获取全部关注ID的列表：输入：用户id string
`Get_Follow_List(User_id string)` ；返回结果：用户粉丝id列表 []string;若失败，返回错误信息
- 取消关注：输入：用户id string，取消关注id string
`Delete_Fans(User_id string, to_user_id string)`;若失败，返回错误信息
### 2.3 视频信息表
- 新增视频:将新的视频插入数据表：输入:用户id string, 视频id string, Video_url string
  `New_Video(db *gorm.DB, User_id string, Video_id string, Video_url string)`;若失败，返回错误信息 
- 从视频点赞表中获取和更新到视频点赞数:输入：视频id string
  `Get_and_Update_Favorite_count(user_id string)`;返回结果：视频点赞数 int;
  注：以上函数直接从点赞表中获得视频的点赞数，并且同步到视频信息表;
- 更新视频评论数：用视频评论表获取视频的评论数，同时更新到视频信息表;输入：视频id string
  `Get_and_Update_Comment_Count(Video_id string)`
返回结果：视频评论数 int
### 2.4 视频点赞表
- 新增点赞信息：将用户点赞信息插入点赞表：输入：用户id string，视频id string
  `Like_This_Video(User_id string, Video_id string)`;若失败，返回错误信息 
- 获取全部点赞用户ID的列表：输入：视频id string
  `Get_Favorite_List(Video_id string)` ；返回结果：点赞的用户id列表 []string;若失败，返回错误信息
- 获取用户点赞的全部视频ID的列表：输入：用户id string
  `Get_Like_List(User_id string)` ；返回结果：用户点赞的视频id列表 []string;若失败，返回错误信息
- 取消点赞：从表中删除点赞信息 输入：输入：用户id string，视频id string
  `Unlike_This_Video(User_id string, Video_id string)`;若失败，返回错误信息
### 2.5 视频评论表
- 新增评论：将用户对视频的评论加入数据表：输入：视频id string，用户id string，评论 string
  `New_comment(Video_id string, Commenter_id string, Comment string)`
- 获取评论：获取视频的全部评论和评论者的id列表；输入：视频id
  `Get_Comment_List(Video_id string)`
返回结果：评论者id列表，评论列表 ([]string, []string)


  
## 3 TODO
- 关注-粉丝表和用户信息表之间新增外键限制，关联查询；
- 视频的url和顺序如何生成？如何调用？
- 并发情况下的锁处理；
- 同步到北京时间；