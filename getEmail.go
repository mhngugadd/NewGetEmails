package main

import (
	"github.com/mhngugadd/NewGetEmail/file"
	"path/filepath"
	"os"
	"log"
	"strings"
	"io/ioutil"
)

type GetJob struct{
	Jobs chan Job
}

func (j *GetJob)GetJobs(maxWorkers int)  {
	// 获取当前执行的文件夹路径
	dir := GetCurrentDir()
	// 获取文件列表
	files , err := GetAllFile(dir)
	if err != nil {
		log.Fatal(err.Error())
	}
	job := file.NewReanJob(maxWorkers)
	for _ , fileName := range files  {
		job.File <- fileName
	}
}

func GetCurrentDir() string  {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
// 获取指定目录下的文件列表
func GetAllFile(dirName string) ([]string, error ){
	fileName := []string{}
	list , err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, file := range list  {
		isTxt := strings.HasSuffix(file.Name(),"txt")
		if ok := file.IsDir(); !ok && isTxt {
			fileName = append(fileName,file.Name())
		}
	}
	return  fileName, err
}