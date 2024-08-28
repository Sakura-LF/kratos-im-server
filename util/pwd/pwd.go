package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd Hash 密码
func HashPwd(pwd string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hashPwd)
}

// CheckPwd 密码 hash之后的密码, 输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
