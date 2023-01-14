package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ipv6Addresses, err := GetPublicIpv6()
	if err != nil {
		fmt.Println("Get Ipv6 addresses failed", err)
		return
	}

	ipv6 := ipv6Addresses[0]
	fmt.Println(ipv6.IP)

	configFileName := os.Args[1]

	config, file := ReadConfig(configFileName)

	config.LastIP = ipv6.IP.String()
	config.Token = "test"
	config.ExpireAt = time.Now().Unix()

	fmt.Println(file.Name())
	fmt.Println(config)

	err = WriteConfig(file, config)
	if err != nil {
		fmt.Println("Write Config failed", err)
		return
	}
}
