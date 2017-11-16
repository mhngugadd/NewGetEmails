package main

import "fmt"

func main() {
	maxThread := 10
	maxWorkers := 50
	jobs := make(chan Job)
	exit := make(chan bool)
	g := GetJob{ Jobs : jobs}
	g.GetJobs(maxWorkers)
	for i := 0; i < maxThread; i++ {
		go ProcessJob(jobs,exit)
	}
	for _,e := range exit {
		fmt.Println(e)
	}
}

