package main

import (
	"errors"
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// createClient 创建发起请求的client
func createClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

// 封装获取对应解析记录的方法
func getRecordIp(client *alidns20150109.Client, domain *string, recordId *string) *string {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: domain,
	}
	// 复制代码运行请自行打印 API 的返回值
	result, _err := client.DescribeDomainRecords(describeDomainRecordsRequest)
	if _err != nil {
		return nil
	}
	records := result.Body.DomainRecords.Record
	var currentIP *string
	for _, record := range records {
		if *record.RecordId == *recordId {
			currentIP = record.Value
		}
	}
	return currentIP
}

// 执行比对当前ip和dns值，并更新操作
func UpdateDomain(config *AlidnsConfig, newIP *string) (bool, error) {
	var updated = false
	client, _err := createClient(&config.AccessKeyId, &config.AccessKeySecret)
	if _err != nil {
		return updated, _err
	}
	currentIP := getRecordIp(client, &config.Domain, &config.RecordId)

	if currentIP == nil {
		return updated, errors.New("获取Ip解析记录失败")
	} else {
		fmt.Printf("Current record value of domain is %s\n", *currentIP)
	}

	if *currentIP != *newIP {
		updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
			RecordId: tea.String(config.RecordId),
			RR:       tea.String(config.RR),
			Type:     tea.String(config.RecordType),
			Value:    newIP,
		}
		// 复制代码运行请自行打印 API 的返回值
		_, _err := client.UpdateDomainRecord(updateDomainRecordRequest)
		if _err != nil {
			return updated, _err
		} else {
			updated = true
			fmt.Printf("Update dns record successfully, now record ip is %s\n", *getRecordIp(client, &config.Domain, &config.RecordId))
		}
	} else {
		fmt.Println("IP not changed, no need to update.")
	}
	return updated, _err
}
