package file

import (
	"io/ioutil"
	"log"
	"regexp"
	"errors"
)

type ReadJob struct {
	File chan string
	Content chan []byte
	Emails chan []string
}

func NewReanJob(maxWorkers int) *ReadJob  {
	return &ReadJob{
		File : make(chan string , maxWorkers),
		Content : make(chan []byte , maxWorkers),
		Emails : make(chan []string , maxWorkers),
	}
}

func (j ReadJob)Read() ( error)  {
	f := <-j.File
	content , err := ioutil.ReadFile(f)
	if err !=nil {
		log.Fatal(err.Error())
	}
	j.Content <-content
	return  err
}
// get emails
func (j ReadJob)Filter()( error)  {
	regEmail := regexp.MustCompile("^\\w+@\\w+\\.\\w{2,4}$")
	c := <-j.Content
	email := regEmail.FindAllString(string(c),-1)

	if email == nil {
		err := errors.New("此文件中没有找到Email")
		return   err
	}
	j.Emails <- email
	return  nil
}

func (j ReadJob)Write()error  {
	e := <-j.Emails
	err := Writer(e)
	if err != nil {
		log.Fatal(err.Error())
	}
	return err
}