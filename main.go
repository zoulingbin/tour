package main

import (
	"github.com/zoulingbin/tour/cmd"
	"log"
)

func main(){
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v\n",err.Error())
	}
}