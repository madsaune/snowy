package auth

import (
	"encoding/base64"
	"fmt"
)

type Authorizer struct {
	Username    string
	Password    string
	InstanceURL string
}

func (a *Authorizer) ToBase64() string {
	val := fmt.Sprintf("%s:%s", a.Username, a.Password)
	return base64.StdEncoding.EncodeToString([]byte(val))
}
