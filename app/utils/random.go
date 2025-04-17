package utils

import "math/rand"

var runes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func GenerateTicketID(size int) string{
 str := make([]rune, size)

 for i := range str{
  str[i] = runes[rand.Intn(len(runes))]
 }

 return string(str)
}