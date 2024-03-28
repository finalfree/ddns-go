package main

import (
	"io/ioutil"
	"net"
	"net/http"
)

func GetPublicIpv4() (*net.IP, error) {
	httpResp, err := http.Get("https://cogear.cn/ip")
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	ipStr := string(body)
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, err
	}
	return &ip, nil
}
