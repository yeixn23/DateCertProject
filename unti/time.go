package unti

import "time"

const TIME_FORMAT_ONE = "2006/01/09 15:04:05"
const TIME_FORMAT_TWO  = "2006年01月02日 15:04:05"

func TimeNow(format string)string  {
	return time.Now().Format(format)
}
func TimeFormat(t int64,nesc int64,format string)string  {
	return time.Unix(t,nesc).Format(format)
}