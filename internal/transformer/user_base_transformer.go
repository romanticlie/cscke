package transformer

import (
	"github.com/ibllex/go-fractal"
	"cscke/internal/model"
)

type UserBaseTransformer struct {
	fractal.BaseTransformer
}

func NewUserBaseTransformer() *UserBaseTransformer {
	return &UserBaseTransformer{}
}

func (t *UserBaseTransformer) Transform(data fractal.Any) fractal.M {

	result := fractal.M{}

	if user, ok := data.(*model.User); ok {
		result["nickname"] = user.Nickname
		result["avatar"] = user.Avatar
		result["gender"] = user.Gender
	}

	return result
}
