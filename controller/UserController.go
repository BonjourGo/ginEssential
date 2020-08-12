package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/dto"
	"ginEssential/model"
	"ginEssential/response"
	"ginEssential/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context)  {
	db := common.DB
	// 获取参数
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	// 数据验证

	if len(pwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "密码必须为6位",
		//})
		return
	}
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号必须为11 位",
		//})
		return
	}
	if isPhoneExist(db, phone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号已存在，请登录！")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号已存在，请登录！",
		//})
		return
	}
	// 随机给一个名字
	if len(name) == 0 {
		name = utils.RandString(10)
	}
	// 加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{
		//	"code": 500,
		//	"msg":  "手机号已存在，请登录！",
		//})
		panic(err)
		return
	}
	newUser := model.User{
		Name:  name,
		Phone: phone,
		Pwd:   string(password),
	}
	db.Create(&newUser)
	response.Success(c, gin.H{}, "注册成功！请登录")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})
	fmt.Println(name, phone, pwd)
}

func Login(c *gin.Context)  {
	DB := common.GetDB()
	// 获取参数
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	// 验证数据
	//if len(pwd) < 6 {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code": 422,
	//		"msg":  "密码必须为6位",
	//	})
	//	return
	//}
	if len(phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号必须为11位",
		//})
		return
	}

	var user model.User
	DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号未注册")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号未注册！",
		//})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(pwd)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误！")
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"code": 400,
		//	"msg":  "密码错误！",
		//})
		return
	}
	token, err:= common.GetToken(user)
	if err != nil {
		response.Fail(c, gin.H{"code": 500}, "系统异常！")
	//	response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常！")
		//c.JSON(http.StatusInternalServerError, gin.H{
		//	"code": 500,
		//	"msg":  "系统异常！",
		//})
		return
	}
	response.Success(c, gin.H{"token":token}, "登陆成功")
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": gin.H{"token":token},
	//	"msg":  "登陆成功",
	//})
}

func Info(c *gin.Context)  {
	// 从上下文获取user
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": dto.UserToUserDTO(user.(model.User))})
}

func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}



