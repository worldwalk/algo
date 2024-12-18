package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FeedbackVO 用于转换反馈上报请求
type FeedbackVO struct {
	Message        string
	LevelId        string
	OccurrenceTime int64
	PartId         string
	FbType         string
	Custom         string
	Attach         []string // 反馈附件
	Img            []string
	UploadIndexes  []int32
	AppId          string
	T              string                 // 13位毫秒数时间戳
	UserId         string                 // 用户ID，用户唯一标识
	UserName       string                 // 用户名称
	Version        string                 // 应用版本号
	Ip             string                 // 客户端 ip
	Hardware       string                 // 机型
	Os             string                 // 系统版本
	Net            string                 // 网络类型: 1 wifi, 2 2g, 3 3g, 4 4g, 5 5g, 未知的留空
	Imei           string                 // 设备标识
	Brand          string                 // 厂商
	Root           string                 // 是否已root，Android 特有: 0 未知, 1 yes, 2 no
	DeviceId       string                 // 设备的guid
	BuildNo        string                 // 构建号
	CustomField    map[string]interface{} // 新版自定义字段
	UseCaseID      int32                  // 众测用例ID
	TaskID         int32                  // 众测任务ID
	SourceEN       string                 // 来源英文，优先使用 partID
	Url            string                 // 内容链接
	Processor      string                 //处理人
}

// 读取文件并反序列化 JSON 数组
func readAndDeserialize(filename string) ([]FeedbackVO, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 反序列化 JSON 数组
	var people []FeedbackVO
	// fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &people)
	if err != nil {
		return nil, err
	}

	return people, nil
}
