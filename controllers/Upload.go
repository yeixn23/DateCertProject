package controllers

import (
	"DataCertProject/models"
	"DataCertProject/unti"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)
type UploadController struct{
	beego.Controller
}
/*
*
 */
func (u *UploadController) Get() {
	phone := u.GetString("phone")
	u.Data["Phone"] =phone
	u.TplName = "home.html"
}
func (u *UploadController) Post() {
	//标题
	fileTitle := u.Ctx.Request.PostFormValue("upload_title")
	phone := u.Ctx.Request.PostFormValue("phone")
	//文件
	file,header,err := u.GetFile("upload_file")

	if err !=nil {
		u.Ctx.WriteString("抱歉用户文件解析失败，请重试！")
		return
	}
	//关闭文件
	defer file.Close()

	fmt.Println("自定义标题",fileTitle)
	fmt.Println("文件名称",header.Filename)
	fmt.Println("文件大小",header.Size)
	fmt.Println(file)

	//文件全路径：文件路径 + 文件名 +"."+ 文件后缀
	uploadDir := "./static/img/" + header.Filename
	//文件权限： a+b+c
	//a:文件所有者拥有的权限, 读、写、执行---4/2/1
	//b：文件所有者所在的组的用户对文件拥有的权限  读、写、执行4/2/1
	//c: 其他用户对文件拥有的权限  读、写、执行4/2/1


	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)


	//创建一个writer用于在硬盘写一个文件
	writer := bufio.NewWriter(saveFile)
	file_size, err := io.Copy(writer,file)
	if err !=nil {
		u.Ctx.WriteString("对不起保存电子数据失败，请重试")
		return
	}
	fmt.Println(file_size)
	defer saveFile.Close()
	//2、计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	hash, err := unti.Md5Hashfile(hashFile)

	//3、将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix() //毫秒数
	record.FileCert = hash
	record.Phone = phone //手机
	_, err = record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉，数据认证错误, 请重试!")
		return
	}

	//4、从数据库中读取phone用户对应的所有认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，获取认证数据失败, 请重试!")
		return
	}
	fmt.Println(records)
	u.Data["Records"] = records
	u.Data["Phone"] = phone
	u.TplName = "list_record.html"
}












//func (this *UploadController) Get(){
//	this.TplName = "login.html"
//}

//func (this *UploadController) Post(){
//
//	f, h, _ := this.GetFile("myfile")//获取上传的文件
//	ext := path.Ext(h.Filename)
//	//验证后缀名是否符合要求
//	var AllowExtMap map[string]bool = map[string]bool{
//		".jpg":true,
//		".jpeg":true,
//		".png":true,
//		".txt":true,
//	}
//	if _,ok:=AllowExtMap[ext];!ok{
//		this.Ctx.WriteString( "后缀名不符合上传要求" )
//		return
//	}
//	//创建目录
//	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
//	err := os.MkdirAll( uploadDir , 777)
//	if err != nil {
//		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
//		return
//	}
//	//构造文件名称
//	rand.Seed(time.Now().UnixNano())
//	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
//	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )
//
//	fileName := fmt.Sprintf("%x",hashName) + ext
//	//this.Ctx.WriteString(  fileName )
//
//	fpath := uploadDir + fileName
//	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
//	err = this.SaveToFile("myfile", fpath)
//	if err != nil {
//		this.Ctx.WriteString( fmt.Sprintf("%v",err) )
//	}
//	this.Ctx.WriteString( "上传成功~！！！！！！！" )
//}