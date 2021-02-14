package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golearn/bluebell/dao/mysql"
	"golearn/bluebell/logic"
	"golearn/bluebell/models"
	"golearn/bluebell/pkg/snowflake"
	"strconv"
)

// CommunityHandler 社区分区
// @Summary 社区列表接口
// @Description 可获取所有社区list
// @Tags 社区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.Community false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunity
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		fmt.Println("look err", err)
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
// @Summary 社区分类详情接口
// @Description 可获取所有社区分类详情
// @Tags 社区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.CommunityDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunityList
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id

	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CreatePostHandler 创建帖子
// @Summary 创建帖子接口
// @Description 创建帖子
// @Tags post帖子
// @Accept application/json
// @Produce application/json
// @Param Authorization header string True "Bearer 用户令牌"
// @Param object query models.Post false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePost
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	//拿到用户信息
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//创建帖子
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
// @Summary 获取帖子详情接口
// @Description 获取帖子详情
// @Tags post帖子
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ApiPostDetail false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseGetPostDetail
// @Router /posts/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//此为查询回复数量的页数,和每页的大小
	page:=int64(1)
	size:=int64(10)
	page, size = getPageInfo(c)
	// 2. 根据id取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid,page,size)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

//GetPostListHandler 帖子列表
func GetPostListHandler(c *gin.Context) {

	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList true "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

//CreateReply 帖子回复功能
func CreateReplyHandler(c *gin.Context) {
	r := new(models.Reply)
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err := c.ShouldBindJSON(r); err != nil {
		zap.L().Error("c.ShouldBindJSON(p) error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	author,err:=mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) error", zap.Error(err))
		return
	}
	r.AuthorID=author.AuthorID
	r.ReAuthorID = userID
	r.PostID=pid
	r.ReplyID=snowflake.GenID()
		if err := logic.CreateReply(r); err != nil {
		zap.L().Error("logic.CreateReply(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
