package middleware

import (
	"errors"
	"fmt"
	"gin-vue-blog/model"
	"gin-vue-blog/utils"
	"gin-vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//生成token

func (j *JWT) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

//验证token

func (j *JWT) ParserToken(tokenString string) int {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	// 验证token
	if token.Valid {
		return errmsg.SUCCESS
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return errmsg.ERROR_TOKEN_TYPE_WRONG
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return errmsg.ERROR_TOKEN_RUNTIME
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return errmsg.ERROR_TOKEN_WRONG
	} else {
		return errmsg.ERROR_TOKEN_NOT_EXIST
	}
}

//jwt中间件

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		code = j.ParserToken(checkToken[1])
		if code != errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		code = j.CheckAuth(checkToken[1])
		if code != errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

//二次确认用户权限

func (j *JWT) CheckAuth(tokenString string) int {
	var claims Claims
	_, _ = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	fmt.Println(claims.Username)
	var user model.User

	user, code := model.CheckUserByName(claims.Username)

	if code != errmsg.SUCCESS {
		return errmsg.ERROR_USER_NOT_EXIST
	}

	if user.Role != 1 {
		return errmsg.ERROR_NO_PRIVILIGE
	}

	return errmsg.SUCCESS
}

//通过token获取用户信息

func GetUserInfoByToken(tokenHeader string) (model.User, int) {
	var code int
	if tokenHeader == "" {
		code = errmsg.ERROR_TOKEN_NOT_EXIST
		return model.User{}, code
	}
	checkToken := strings.Split(tokenHeader, " ")
	if len(checkToken) == 0 {
		code = errmsg.ERROR_TOKEN_NOT_EXIST
		return model.User{}, code
	}
	if len(checkToken) != 2 || checkToken[0] != "Bearer" {
		code = errmsg.ERROR_TOKEN_TYPE_WRONG
		return model.User{}, code
	}
	j := NewJWT()
	code = j.ParserToken(checkToken[1])
	if code != errmsg.SUCCESS {
		return model.User{}, code
	}

	var claims Claims
	_, _ = jwt.ParseWithClaims(checkToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	fmt.Println(claims.Username)
	var user model.User

	user, code = model.CheckUserByName(claims.Username)

	return user, code
}
