package xerr

var message map[uint32]string = make(map[uint32]string)

func init() {
	message[SERVER_COMMON_ERROR] = "服务器开小差了，请稍后重试。"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[DB_NOTFOUNR] = "未查找到数据"
	message[DB_DELETE_NOTEXIST] = "数据库中没有此条数据"
	message[FILE_UPLOAD_ERROR] = "文件上传失败"
	message[TOKEN_NOT_EXIST] = "为获取到 token"
	message[TOKEN_INBLACKLIST_ERROT] = "黑名单或异地登录"
	message[TOKEN_EXPIRE_ERROR] = "登录已过期，请重新登录"
	message[USERNAME_IS_EXIST] = "用户名已存在"
	message[TOKEN_PARSE_ERROR] = "token解析错误"
}

func MapErrMsg(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器开小差了，请稍后重试。"
	}
}

func IsCodeErr(code uint32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}
