package  cmd

import (
    "fmt"
    "github.com/slowmoon/modtest/core/helper"
)


func  Printf(mes string)  {
	ip, err := helper.Ip2bytes(mes)
    if err != nil {
    	panic(err)
	}
	fmt.Println(ip)
}
