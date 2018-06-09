package base62

import (
	"math"
	"strings"
)

const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func Encode(i int64) (v string) {
	if i == 0 {
		return string(characters[0])
	}
	r := i
	for r != 0 {
		v += string(characters[r%62])
		r = r / 62
	}
	return
}

func Decode(src string) (r int64) {
	if src == "" {
		return 0
	}
	for k, v := range src {
		i := strings.Index(characters, string(v))
		n := math.Pow(62, float64(k))
		r += int64(i) * int64(n)
	}
	return
}
