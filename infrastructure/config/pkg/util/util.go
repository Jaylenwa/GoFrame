package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"time"
)

// Getwd 获得项目的根路径
func Getwd() string {
	dir, _ := os.Getwd()
	return dir
}

func CeilPageNum(total, pageSize uint32) uint32 {
	return uint32(int(math.Ceil(float64(total) / float64(pageSize))))
}

// PrintJson 打印Json
func PrintJson(args interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	return out.String()
}

func GetLocalIPv4Address() (ipv4Address string, err error) {
	//获取所有网卡
	addrs, err := net.InterfaceAddrs()

	//遍历
	for _, addr := range addrs {
		//取网络地址的网卡的信息
		ipNet, isIpNet := addr.(*net.IPNet)
		//是网卡并且不是本地环回网卡
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			//能正常转成ipv4
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}

	return
}

func TimeStr2Time(timeStr string) interface{} {
	if timeStr == "" {
		return ""
	}
	loc, _ := time.LoadLocation("Local") //获取当地时区
	location, _ := time.ParseInLocation("2006/01/02 15:04", timeStr, loc)

	return location
}
