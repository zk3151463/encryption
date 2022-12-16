package main

import (
	"encoding/hex"
	"fmt"

	"github.com/druidcaesa/gotool"
	"github.com/zk3151463/encryption/encryption"
)

func main() {
	// 云端数据

	md5 := gotool.BcryptUtils.MD5("123456789")
	data := gotool.BcryptUtils.Generate(md5)
	fmt.Println(data)
	// //进行加密对比
	// hash := gotool.BcryptUtils.CompareHash(generate, md5)
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := gotool.BcryptUtils.GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))
	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := gotool.BcryptUtils.RsaEncrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := gotool.BcryptUtils.RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))

	fmt.Println("mac", encryption.GetMac())
	// fmt.Println(encryption.GetCpuId())

}
