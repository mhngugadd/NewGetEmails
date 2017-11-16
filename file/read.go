package file

import (
	"path/filepath"
	"os"
	"log"
	"strings"
	"io/ioutil"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetAllFile(dirName string) ([]string, error ){
	fileName := []string{}
	list , err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _,file := range list  {
		isTxt := strings.HasSuffix(file.Name(),"txt")
		if ok := file.IsDir(); !ok && isTxt {
			fileName = append(fileName,file.Name())
		}
	}
	return  fileName, err
}