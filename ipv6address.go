package main

import (
	"fmt"
	"net"
)

func GetPublicIpv6() ([]*net.IPNet, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ipv6addresses []*net.IPNet
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() == nil {
			if !ipnet.IP.IsLinkLocalUnicast() {
				fmt.Println(ipnet.IP.String())
				ipv6addresses = append(ipv6addresses, ipnet)
			}
		}
	}
	return ipv6addresses, nil
}
