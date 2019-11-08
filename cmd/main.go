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
		panic(err)
	}

	_, _ = nv.Sync()
	_, _ = nv.ChannelValueRequest()
	ud, err := nv.UnitData()
	if err != nil {
		panic(err)
	}

	fmt.Println(*ud.UnitData)

}
