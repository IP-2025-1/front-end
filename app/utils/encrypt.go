package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Encrypt(data string) (hash string){
	hasher := md5.New()
	hasher.Write([]byte(data))
	hash = hex.EncodeToString(hasher.Sum(nil))
	return
}