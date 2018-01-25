package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
	"crypto/sha1"
	"math"
	"math/big"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

func Sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func RandomInfo(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func IntTo52(length int, seed int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sc := big.NewInt(int64(int(math.Pow(52, float64(length-1))) + seed))
	var code string
	nData := big.NewInt(52)
	nRem := big.NewInt(0)
	for i := length; i > 0; i-- {
		sc.DivMod(sc, nData, nRem)
		rem := nRem.Int64()
		code = code + string(str[rem])
	}
	return code
}
