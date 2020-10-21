package unti

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/*
*
 */
func Md5Hash(data string)string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	databytes := md5Hash.Sum(nil)
	return hex.EncodeToString(databytes)
}

func Md5Hashfile(reader io.Reader)(string,error) {
	bytes,err :=ioutil.ReadAll(reader)
	fmt.Println("hash输入端内容",bytes)
	if err !=nil {
		return "",err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashByte := md5Hash.Sum(nil)
	return hex.EncodeToString(hashByte),err
}

func SHA256Hash(data []byte) ([]byte) {
	//对block字段进行拼接
	//对拼接后的数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	return sha256Hash.Sum(nil)
}

