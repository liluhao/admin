package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string) string {
	//func Sum256(data []byte) [32]byte:返回数据的SHA256校验和
	obj := sha256.Sum256([]byte(data))
	//func EncodeToString(src []byte) string:将数据src编码为字符串s
	return hex.EncodeToString(obj[:])
}

func Sha1(data string) string {
	//func New() hash.Hash:返回一个新的使用SHA1校验的hash.Hash接口
	obj := sha1.New()
	//Write(p []byte) (n int, err error):通过嵌入的匿名io.Writer接口的Write方法向hash中添加更多数据
	obj.Write([]byte(data))
	// Sum(b []byte) []byte : 返回"添加b到当前的hash值"后的新切片，不会改变底层的hash状态
	return hex.EncodeToString(obj.Sum([]byte("")))
}

func Hmac(key, data string) string {
	//func New(h func() hash.Hash, key []byte) hash.Hash:返回一个采用hash.Hash作为底层hash接口、key作为密钥的HMAC算法的hash接口
	obj := hmac.New(md5.New, []byte(key))
	obj.Write([]byte(data))
	return hex.EncodeToString(obj.Sum([]byte("")))
}

func Md5(data string) string {
	//func New() hash.Hash: 返回一个新的使用MD5校验的hash.Hash接口
	obj := md5.New()
	obj.Write([]byte(data))
	return hex.EncodeToString(obj.Sum([]byte("")))
}
