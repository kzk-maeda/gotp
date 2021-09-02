package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
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

func execTOTP(k string, x uint64) uint32 {
	fmt.Println(k)
	key, err := base32.StdEncoding.DecodeString(string(k))
	if err != nil {
		fmt.Printf("Decode Error: %v\n", err)
		return 0
	}

	return hotp(key, t(0, x))
}

func hotp(k []byte, c uint64) uint32 {
	return truncate(hmacsha1(k, c))
}

func t(t0, x uint64) uint64 {
	return (uint64(time.Now().Unix()) - t0) / x
}
