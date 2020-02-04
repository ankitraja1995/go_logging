package main

import (
	"fmt"
	"time"
)

func main() {
	var z int = 2

	a := LogConfig{true, "127.0.0.1", 24224, "logs/testTolerant1.log",
		1, 10, 11, true, "Debug", 100000,
		"tlogs/sdk.log", 3, 6000}
	b := a.GetLogger()

	//c := Logger.LogConfig{true, "127.0.0.1", 24224, "logs/testTolerant2.log",
	//	1, 10, 11, true, Error, 100000,
	//	"tlogs/sdk.log", 3, 6000}
	//d := c.GetLogger()



	type Sample struct {
		X int `json:"x"`
		Y string `json:"y"`
	}

	//fmt.Println(fmt.Sprintf("aaaaaa  %v bbbb %v " , "111", 1234))

	//some_sample := Sample{
	//	X: 1 ,
	//	Y: "aaaaa",
	//
	//}
	//d.Log(Error, "kafka.1", "{\"lg5\": \"%v\"}", z)
	//b.EventLog("kafka.1", some_sample)
	//fmt.Println("sleeping")
	//time.Sleep(10 * time.Second)
	//d.Log(Error, "kafka.1", "{\"lg7\": \"%v\"}", z)
	b.Error("kafka.1", "{\"lg6\": \"%v--%v\"}", z,"cc")
	fmt.Println("sleeping")
	time.Sleep(10 * time.Second)

}
