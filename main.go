package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"msg-crypt/bcrype"
)

const (
	defaultKey = "a very very very very secret key"
)

var (
	operateType string
	msg         string
	key         string
)

func main() {
	initFlag()
	flag.Parse()
	encodeKey := hex.EncodeToString([]byte(defaultKey))
	encrypt := bcrype.Of(encodeKey)
	if operateType == "en" {
		fmt.Println(encrypt.Encrypt(msg))
	} else {
		fmt.Println(encrypt.Decrypt(msg))
	}
}

func initFlag() {
	flag.StringVar(&operateType, "type", "en", "操作分类，en-加密 de-解密")
	flag.StringVar(&msg, "msg", "", "需要加密的文本")
	flag.StringVar(&key, "key", defaultKey, "加解密的key")
}
