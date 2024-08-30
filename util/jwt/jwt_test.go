package jwt

import (
	"testing"
)

func TestGenToken(t *testing.T) {
	payLoad := &JwtPayLoad{
		UserID:   1,
		Nickname: "sa",
		Role:     1,
	}
	token, err := GenToken(*payLoad, "secret", 3600)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
	}
	claims, err := ParseToken(token, "secret")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(claims)
	}
}
