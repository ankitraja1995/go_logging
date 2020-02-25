package main


import (
	"fmt"
	"go_logging"
	"time"
)

func main(){

	a := go_logging.LogConfig{true, "127.0.0.1", 24224, "logs/testTolerant1.log",
		1, 10, 11, true, "Error", 100000,
		"tlogs/sdk.log", 3, 6000}

	b := a.GetLogger()
	b.Error("kafka.1", "{\"lg6\": \"%v_%v\"}", 2,"cc")

	type pet struct {
		Name string
		Age int
	}
	cat := pet{
		"tom",
		10,
	}
	b.EventLog("kafka.1", cat)

	fmt.Println("sleeping")
	time.Sleep(10 * time.Second)
	fmt.Println("slept")
}
