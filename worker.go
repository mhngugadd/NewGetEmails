package main

import (
	"log"
)

type Job interface{
	Read() ( error) // file string
	Filter() ( error)
	Write() error
}

func ProcessJob(jobs chan Job , exit chan bool)  {
	j := <- jobs
	err := j.Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = j.Filter()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = j.Write()
	if err != nil {
		log.Fatal(err.Error())
	}
	exit <- true
}
