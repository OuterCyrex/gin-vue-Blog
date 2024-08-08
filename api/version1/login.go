package version1

import (
	"gin-vue-blog/middleware"
	"gin-vue-blog/model"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)

	code := model.CheckLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		setToken(c, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.ERROR,
			"msg":  errmsg.GetErrMsg(code),
		})
	}
}

func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "GinVueBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
