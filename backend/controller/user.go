package controller

import (
	"github.com/gin-gonic/gin"
	"whiteboard/service"
	"whiteboard/utils/jwt"
	"whiteboard/utils/res"
)

func Register(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	err := service.Register(name, pwd)
	if err != nil {
		res.Ok(c, 400, "注册失败, 账号已注册", nil)
		return
	}
	res.Ok(c, 200, "注册成功", nil)
}

func Login(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	user, err := service.GetUserByName(name)
	if err != nil {
		res.Ok(c, 400, "用户不存在", nil)
		return
	}
	if user == nil || user.Pwd != pwd {
		res.Ok(c, 400, "账号或密码错误", nil)
		return
	}
	token, _ := jwt.GenToken(user.Id, user.Name)
	res.Ok(c, 200, "登录成功", gin.H{
		"token": token,
	})

}