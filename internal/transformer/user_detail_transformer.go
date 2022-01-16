package transformer

import (
	"cscke/internal/model"
	"cscke/pkg/fun"
	"github.com/ibllex/go-fractal"
	"strconv"
)

const (
	UserModuleBase = "base"
)

var userModuleConfig = map[string]string{
	UserModuleBase: "基础模块",
}

type UserDetailTransformer struct {
	fractal.BaseTransformer
	modules []string
}

func NewUserDetailTransformer(modules ...string) *UserDetailTransformer {
	return &UserDetailTransformer{
		modules: modules,
	}
}

func (t *UserDetailTransformer) Transform(data fractal.Any) fractal.M {

	result := fractal.M{}

	if len(t.modules) == 0 {
		t.modules = fun.MapStringKeys(userModuleConfig)
	}

	if user, ok := data.(*model.User); ok {
		result["userid"] = user.Userid
		result["strid"] = strconv.FormatUint(user.Userid, 10)

		//根据模块来获取
		for _, module := range t.modules {

			if _, isExistMod := userModuleConfig[module]; isExistMod {

				switch module {
				case UserModuleBase:
					result[module] = Item(user, NewUserBaseTransformer())["data"]
				}

			}
		}

	}

	return result
}
