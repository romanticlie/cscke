package validate

type Register struct {
	Nickname string `form:"nickname" binding:"required"`
	Gender int  `form:"gender" binding:"required"`
}
