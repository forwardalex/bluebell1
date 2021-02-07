package logic

import (
	"errors"
	"golearn/bluebell/dao/mysql"
	"golearn/bluebell/models"
	"golearn/bluebell/pkg/jwt"
	"golearn/bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//查询用户是否存在
	if exit, _ := mysql.CheckUserExit(p.Username); exit {
		return errors.New("用户已经存在")
	}
	//生成UID 不重复唯一
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {

		return
	}
	user.Token = token
	return
}
