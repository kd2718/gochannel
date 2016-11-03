package main

import (
	"fmt"
	"github.com/kd2718/gochannel/personChannel"
	"github.com/kd2718/gomarshal/person"
	"time"
)

func main() {
	k := &person.Person{33, "kory", "donati", true}
	defer fmt.Println("END")
	defer func() {
		if r := recover(); r != nil{
			fmt.Printf("%v has died...", k.First)
		}
	}()

	c := personChannel.PersonChannel(k)
	for {
		guy, ok := <-c
		if !ok {
			fmt.Println("that's it for", k.First)
			break
		}
		fmt.Println(guy)
		<-time.After(30000000)
	}

}
