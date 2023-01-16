package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	ipv6Addresses, err := GetPublicIpv6()
	if err != nil {
		fmt.Println("Get Ipv6 addresses failed", err)
		return
	}

	ipv6 := ipv6Addresses[0].IP.String()

	configFileName := os.Args[1]
	tempFileName := path.Dir(configFileName) + "/temp.yaml"

	tempConfig, file := ReadTempConfig(tempFileName)
	defer file.Close()

	configChanged := false

	var message Message
	config := ReadConfig(configFileName)
	message.ToUser = config.ToUser
	message.AgentID = config.AgentID
	message.MsgType = "text"

	if time.Now().Unix() > tempConfig.ExpireAt {
		wechatToken, err := Get_Token(config)
		if err != nil {
			fmt.Println("Get wechat token failed", err)
			return
		}
		configChanged = true
		tempConfig.Token = wechatToken.Token
		tempConfig.ExpireAt = time.Now().Unix() + wechatToken.ExpireAt
	}

	if tempConfig.LastIP != ipv6 {
		configChanged = true
		tempConfig.LastIP = ipv6
		message.Text = Text{Content: "IP changed to " + ipv6}
	}

	if configChanged {
		file.Truncate(0)
		file.Seek(0, 0)
		err = WriteTempConfig(file, tempConfig)
		if err != nil {
			fmt.Println("Write TempConfig failed", err)
			return
		}
		err := SendMessage(tempConfig.Token, &message)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("IP address changed to " + ipv6)
		}
	} else {
		fmt.Println("IP address has no change.")
	}
}
