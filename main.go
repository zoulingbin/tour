package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)


type Name string

func (i *Name) Set(value string) error{
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name(value)
	return nil
}

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func main()  {
	var name Name
	flag.Var(&name,"name","帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}