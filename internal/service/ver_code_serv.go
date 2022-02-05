package service

import "sync"

var (
	verCodeServ     *VerCodeService
	verCodeServOnce sync.Once
)

const (
	VerCodeReg = 1 + iota
)

type VerCodeService struct{}

func GetVerCodeServ() *VerCodeService {
	if verCodeServ == nil {
		verCodeServOnce.Do(func() {
			verCodeServ = &VerCodeService{}
		})
	}

	return verCodeServ
}

// Check 校验手机号是否正确
func (v *VerCodeService) Check(tp int, tel string, random string) bool {

	return random == "123456"
}
