package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetTokenResponse struct {
	Token    string `json:"access_token"`
	ExpireAt int64  `json:"expires_in"`
}

type Message struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentID int    `json:"agentid"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

func Get_Token(config *Config) (*GetTokenResponse, error) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + config.CorpID + "&corpsecret=" + config.AppSecret
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//get json data from response
	decoder := json.NewDecoder(resp.Body)
	var data GetTokenResponse
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func SendMessage(token string, message *Message) error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token
	//send http post with body message
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	//print response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(resp.Status, buf.String())
	return nil
}
