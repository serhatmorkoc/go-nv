package main

import (
	"fmt"
	"github.com/serhatmorkoc/go-nv"
	"time"
)

func main() {

	nv := nv.NewService(&nv.Config{
		PortName:    "COM3",
		Address:     0,
		ReadTimeout: time.Millisecond * 50,
		BaudRate:    9600,
	})

	err := nv.Connect()
	if err != nil {
		fmt.Println(err)
	}

	nv.Sync()
	nv.Enable()




}
