package piscine

import "fmt"

func StrRev(s string) {
	var rev string
	for i := len(s); i >= 0; i-- {
		rev = string(s[i])
	}
	return rev
}
