package blockchain

import (
	"time"
)

/*
 *区块结构体的定义
 */
type Block struct {
	Height int64 //区块高度
	//Size int  //区块大小
	TimeStamp int64  //时间戳
	Hash []byte  //区块的hash
	Data []byte  //数据
	PrevHash []byte  // 上一个区块的hash
	Version string  //版本号
	Nonce int64
}

func NewBlock(height int64 ,data []byte , prevHash []byte) (Block) {
	//1、构建一个block实例，用于生成区块
	block := Block{
		Height:    height ,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}
	//2。为新生的block，寻找合适的nonce
	pow := NewPow(block)
	nonce,bolcksha256hash:= pow.Run()
	block.Nonce = nonce
	block.Hash = bolcksha256hash
	//3、将block的nonce设置


	//heightBytes,_ := unti.IntToBytes(block.Height)
	//timeBytes,_ := unti.IntToBytes(block.TimeStamp)
	//versionBytes := unti.StringToBytes(block.Version)
	//nonceBytes,_ := unti.IntToBytes(block.Nonce)
	//
	//blockBytes := bytes.Join([][]byte{
	//	heightBytes,
	//	timeBytes,
	//	data,
	//	prevHash,
	//	versionBytes,
	//	nonceBytes,
	//},[]byte{})
	//block.Hash = unti.SHA256Hash(blockBytes)
	return block
}