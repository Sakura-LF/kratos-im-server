package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	pwd := "123456"
	hashPwd := HashPwd(pwd)
	fmt.Println(hashPwd)
}

func TestCheckPwd(t *testing.T) {
	//$2a$04$Wa38056lMnmp39.tuyXTRurMqL9z7O7mWD0Cb5NfGx3clP9.3EREm
	pwd := "123456"
	ok := CheckPwd("$2a$04$Wa38056lMnmp39.tuyXTRurMqL9z7O7mWD0Cb5NfGx3clP9.3EREm", pwd)
	fmt.Println(ok)
}
