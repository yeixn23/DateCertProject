package blockchain

import (
	"DataCertProject/unti"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const DIFFFICULTY  = 12
/*
 *工作量证明结构体
 */
type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明对应的那个区块
	Block Block
}
/*
 *实例化一个pow算法实例
 */
func NewPow(block Block) ProofOfWork {
	target := big.NewInt(1)//初始值
	target.Lsh(target, 255-DIFFFICULTY)
	pow := ProofOfWork{
		Target:target,
		Block:block,
		}
	return pow
}

/*
 *pow算法：寻找符合条件的nonce值
 */
func (p ProofOfWork) Run()  (int64,[]byte){
	var nonce int64
	var bigBlock *big.Int
	bigBlock = new(big.Int)
	var bolcksha256hash []byte
	for  {

		block := p.Block

		heightBytes,_ := unti.IntToBytes(block.Height)
		timeBytes,_ :=unti.IntToBytes(block.TimeStamp)
		versionBytes:= unti.StringToBytes(block.Version)
		nonceBytes,_ := unti.IntToBytes(nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})

		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		bolcksha256hash = sha256Hash.Sum(nil)
		fmt.Printf("挖矿中,当前正在尝试Nonce值:%d\n",nonce)
		bigBlock := bigBlock.SetBytes(bolcksha256hash)
		fmt.Printf("目标值：%X\n",p.Target)
		fmt.Printf("hash值：%X\n",bigBlock)
		if p.Target.Cmp(bigBlock) == 1{//如果满足条件时，退出循环} {
			break
		}
		nonce++//如果条件不满足，nonce值+1，继续下次循环
	}
	return nonce,bolcksha256hash
}