package file

import (
	"os"
	"errors"
	"io"
)



func Writer(emails []string) error  {
	fileName  := "./output.txt"
	var err  error
	for _,email := range emails {
		err = Write(fileName , email)
	}
	return err
}

// 判断文件是否存在
func CheckFileExist(FileName string) bool  {
	exist := true
	if _, err := os.Stat(FileName); os.IsNotExist(err){
		exist = false
	}
	return exist
}

func Write(FileName string, word string) error  {
	var f *os.File
	var err error

	if CheckFileExist(FileName) {
		f , err = os.OpenFile(FileName , os.O_APPEND,0666)
	}else {
		f , err = os.Create(FileName)
	}

	if err != nil {
		return errors.New("file not exist or create file failed")
	}

	_ , err = io.WriteString(f,word)

	if err != nil {
		return errors.New("write error")
	}
	return nil
}