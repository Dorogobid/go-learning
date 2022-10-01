package main

import (
    "crypto/md5"
    "encoding/hex"
)

func StrToHash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}