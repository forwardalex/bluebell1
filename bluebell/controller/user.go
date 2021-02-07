package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golearn/bluebell/dao/mysql"
	"golearn/bluebell/logic"
	"golearn/bluebell/models"
	"net/http"
)

// SignUpHandler 用户注册
// @Summary 注册接口
// @Description 注册
// @Tags 注册
// @Accept application/json
// @Produce application/json
// @Param Authorization header string ture "Bearer 用户令牌"
// @Param object query models.ParamSignUp false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseCommunity
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		fmt.Println(err)
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//写入逻辑层
	if err := logic.SignUp(p); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeUserExist)
		zap.L().Error("signup failed", zap.String("username", p.Username), zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//登录处理
	user, err := logic.Login(p)
	if err != nil {

		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//ResponseSuccess(http.StatusOK,"s")
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
