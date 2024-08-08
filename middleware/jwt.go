package middleware

import (
	"errors"
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
		c.Next()
	}
}
