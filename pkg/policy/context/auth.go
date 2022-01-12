package context


import "sync"

var (
	authContext     *AuthContext
	authContextOnce sync.Once
)

type AuthContext struct {
}

func GetAuthContext() *AuthContext {
	if authContext == nil {
		authContextOnce.Do(func() {
			authContext = &AuthContext{}
		})
	}
	return authContext
}