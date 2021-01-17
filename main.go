package main

import (
	"airmirror/airmirror"
	"airmirror/ntpsvr"
	"log"
)

func main() {
	ntpsvr.Start(":7010")

	tv := &airmirror.AirTV{Address: "192.168.1.142:7100"}
	client := airmirror.NewClient(tv)
	err := client.StreamReq()
	if err != nil {
		log.Printf("err:%s", err.Error())
	}
}
