package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_NO_PRIVILIGE     = 1008

	//code = 2000... 分类模块错误

	ERROR_CATEGORY_USED      = 2001
	ERROR_CATEGORY_NOT_EXIST = 2002

	//code = 3000... 文章模块错误

	ERROR_ARTICLE_NOT_EXIST = 3001

	//code = 4000... 评论模块错误

	ERROR_COMMENT_NOT_EXIST = 4001
)

var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USERNAME_USED:      "用户已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_NO_PRIVILIGE:       "用户权限不足",
	ERROR_TOKEN_RUNTIME:      "TOKEN已过期",
	ERROR_TOKEN_NOT_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_WRONG:        "TOKEN出错",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误",
	ERROR_CATEGORY_USED:      "分类已被使用",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
	ERROR_ARTICLE_NOT_EXIST:  "文章不存在",
	ERROR_COMMENT_NOT_EXIST:  "评论不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
