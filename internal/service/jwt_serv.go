package service

import (
	"github.com/dgrijalva/jwt-go"
	"cscke/pkg/fun"
	"sync"
)

var (
	jwtServ *JwtService
	jwtServOnce sync.Once
)

// GetJwtServ  jwt 服务
func GetJwtServ() *JwtService{

	if jwtServ == nil {
		jwtServOnce.Do(func() {
			jwtServ = &JwtService{}
		})
	}

	return jwtServ
}

type JwtService struct {

}

// ParseToken 解析jwtToken 成sub
func (j *JwtService) ParseToken(jwtToken string) (string,error){

	claims := &jwt.StandardClaims{}

	token,err := jwt.ParseWithClaims(jwtToken,claims, func(token *jwt.Token) (interface{}, error) {
		//获取密钥
		v,err := fun.GetYamlCfg("jwt")

		if err != nil {
			return []byte{},err
		}

		return []byte(v.GetString("secret")),nil
	})

	if err != nil {
		return "",err
	}

	if !token.Valid {
		return "",err
	}

	return claims.Subject,nil
}

// GenerateToken 直接生成jwtToken
func (j *JwtService) GenerateToken(sub string) (string,error){

	//获取密钥
	v,err := fun.GetYamlCfg("jwt")

	if err != nil {
		return "",err
	}

	secret := v.GetString("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.StandardClaims{
		Subject: sub,
	})

	return token.SignedString([]byte(secret))
}
