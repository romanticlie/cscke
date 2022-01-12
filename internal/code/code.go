package code

const (
	None       = 0
	ClientErr  = 1
	ServerBusy = 2
	AuthFailed = 7
	Prompt     = 10
	LogFailed     = 11
	FileErr     = 12
)

var ErrCodeMessage = map[int]string{
	ClientErr:  "输入有误",
	ServerBusy: "网络繁忙，请求稍后再试",
	AuthFailed: "身份已过期",
	LogFailed: "登录失败，账号或者密码错误",
	FileErr: "文件上传失败",
}
