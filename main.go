package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	//ipv6Addresses, err := GetPublicIpv6()
	//if err != nil {
	//	fmt.Println("Get Ipv6 addresses failed", err)
	//	return
	//}
	//
	//ipv6 := ipv6Addresses[0].IP.String()
	//for _, address := range ipv6Addresses {
	//	if address.Mask[15] == 0xff {
	//		ipv6 = address.IP.String()
	//		break
	//	}
	//}
	ipv4, err := GetPublicIpv4()
	if err != nil {
		fmt.Println("Get Ipv4 addresses failed", err)
		return
	}
	ipv4str := ipv4.String()

	configFileName := os.Args[1]
	tempFileName := path.Dir(configFileName) + "/temp.yaml"

	tempConfig, file := ReadTempConfig(tempFileName)
	defer file.Close()

	config := ReadConfig(configFileName)
	updated, err := UpdateDomain(&config.AlidnsConfig, &ipv4str)
	if err != nil {
		fmt.Println("Update domain failed", err)
	} else if updated {
		var message Message
		wechatConfig := config.WechatConfig
		message.ToUser = wechatConfig.ToUser
		message.AgentID = wechatConfig.AgentID
		message.MsgType = "text"
		if time.Now().Unix() > tempConfig.ExpireAt {
			wechatToken, err := Get_Token(&wechatConfig)
			if err != nil {
				fmt.Println("Get wechat token failed", err)
			} else {
				tempConfig.Token = wechatToken.Token
				tempConfig.ExpireAt = time.Now().Unix() + wechatToken.ExpireAt

				file.Truncate(0)
				file.Seek(0, 0)
				err = WriteTempConfig(file, tempConfig)
				if err != nil {
					fmt.Println("Write TempConfig failed", err)
				}
			}
		}
		var domain = config.AlidnsConfig.RR + "." + config.AlidnsConfig.Domain
		message.Text = Text{
			Content: fmt.Sprintf("Successfully update domain %s dns record to %s", domain, ipv4str),
		}
		err = SendMessage(tempConfig.Token, &message)
		if err != nil {
			fmt.Println(err)
		}
	}
}
