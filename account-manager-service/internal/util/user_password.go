package util

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(plainPwd string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)

	return string(result), err
}

func CompareHashAndPassword(hashPwd string, plainPwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd)) == nil
}
