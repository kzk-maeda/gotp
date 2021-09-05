package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func truncate(hs []byte) uint32 {
	offset := int(hs[len(hs)-1] & 0x0F)
	p := hs[offset : offset+4]
	return (binary.BigEndian.Uint32(p) & 0x7FFFFFFF) % 1000000
}

func hmacsha1(k []byte, c uint64) []byte {
	cb := make([]byte, 8)
	binary.BigEndian.PutUint64(cb, c)

	mac := hmac.New(sha1.New, k)
	mac.Write(cb)

	return mac.Sum(nil)
}

func execTOTP(k string, x uint64) string {
	// fmt.Println(strings.ToUpper(k))
	key, err := base32.StdEncoding.DecodeString(strings.ToUpper(k))
	if err != nil {
		fmt.Printf("Decode Error: %v\n", err)
		return ""
	}

	totp := hotp(key, t(0, x))
	return paddingZero(totp)
}

func hotp(k []byte, c uint64) uint32 {
	return truncate(hmacsha1(k, c))
}

func t(t0, x uint64) uint64 {
	return (uint64(time.Now().Unix()) - t0) / x
}

func paddingZero(totp uint32) string {
	s := fmt.Sprintf("%06d", totp)
	return s
}
