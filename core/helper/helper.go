package helper

import (
	"errors"
	"net"
)

func Ip2bytes(ips string)(res []byte, err error){
    ip := net.ParseIP(ips)
    if ip ==nil {
    	err = errors.New("invalid ip")
    	return
	}
    res = ip[:]
	return 
}