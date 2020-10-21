package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/unti"
)

type UploadRecord struct {
	Id        int
	FileName  string
	FileSize  int64
	FileCert  string //认证号
	FileTitle string
	CertTime  int64
	Phone     string //对应的用户的phone
	FormateCertTime string
}

/*
*保存上传记录到数据库
 */

func (u UploadRecord) SaveRecord() (int64, error) {
	row, err := db_mysql.Db.Exec("insert into user_upload(file_name,file_size,file_cert,file_title,cert_time,phone) values (?,?,?,?,?,?)",
		u.FileName,
		u.FileSize,
		u.FileCert,
		u.FileTitle,
		u.CertTime,
		u.Phone )
	if err != nil { //保存数据时遇到错误
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/*
*读取数据库中phone用户对应的所有认证数据
 */
func QueryRecordByPhone(phone string) ([]UploadRecord, error) {
	rs, err := db_mysql.Db.Query("select id, file_name, file_size, file_cert, file_title, cert_time, phone from user_upload where phone = ?", phone)
	if err != nil {
		return nil, err
	}
	records := make([]UploadRecord, 0)
	for rs.Next() {
		var record UploadRecord
		rs.Scan(&record.Id,&record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle,&record.CertTime, &record.Phone)
		if err != nil {
			return nil, err
		}
		record.FormateCertTime = unti.TimeFormat(record.CertTime,0,unti.TIME_FORMAT_ONE)
		records = append(records, record)
	}
	return records, nil
}
