package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Resgister(c *gin.Context)  {
	db := common.DB
	// 获取参数
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	// 数据验证

	if len(pwd) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码必须为6位",
		})
		return
	}
	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11 位",
		})
		return
	}
	if isPhoneExist(db, phone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号已存在，请登录！",
		})
		return
	}
	// 随机给一个名字
	if len(name) == 0 {
		name = utils.RandString(10)
	}
	// 加密密码
	//screctPasspord, err := Gen
	newUser := model.User{
		Name:  name,
		Phone: phone,
		Pwd:   pwd,
	}
	db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	fmt.Println(name, phone, pwd)
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}



