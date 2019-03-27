//Package base62 implements conversion from and to base62.
package base62

import (
	"bytes"
	"math"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//ToBase62 returns base62 formated int as a string
func ToBase62(base10 int) string {
	var base62 []byte
	for base10 > 0 {
		remainder := base10 % 62
		base62 = append(base62, alphabet[remainder])
		base10 /= 62
	}

	for i, j := 0, len(base62)-1; i < j; i, j = i+1, j-1 {
		base62[i], base62[j] = base62[j], base62[i]
	}
	return string(base62)
}

//ToBase10 returns base10 formated string as a int
func ToBase10(base62 string) int {
	var base10 int
	power := float64(len(base62) - 1)
	for _, character := range []byte(base62) {
		index := bytes.IndexByte([]byte(alphabet), character)
		base10 += index * int(math.Pow(62, power))
		power--
	}
	return base10
}
