package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USERNAME_USED       = 1001
	ERROR_USER_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST      = 1003
	ERROR_TOKEN_EXIST         = 1004
	ERROR_TOKEN_RUN_TIME      = 1005
	ERROR_TEKEN_WRONG         = 1006
	ERROR_TEKEN_TYPE_WRONG    = 1007
	// code = 1xxx 用户模块
	// code = 2xxx 文章模块
	// code = 3xxx 分类模块
)

var codeMsg = map[int]string{
	SUCCESS:                   "success",
	ERROR:                     "fail",
	ERROR_USERNAME_USED:       "用户名已存在",
	ERROR_USER_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_TOKEN_EXIST:         "token 不存在",
	ERROR_TOKEN_RUN_TIME:      "token 过期",
	ERROR_TEKEN_WRONG:         "token 不正确",
	ERROR_TEKEN_TYPE_WRONG:    "token 格式错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
