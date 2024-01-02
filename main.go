package main

import (
	"flag"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// 定义命令行参数，设置默认值和说明
	ssid := flag.String("s", "your_ssid", "WiFi SSID")
	t := flag.String("t", "WPA2", "WiFi security type (WPA2/WPA/WEP/NONE)")
	p := flag.String("p", "12345678", "WiFi password")
	h := flag.Bool("h", false, "Whether the network is hidden")
	flag.Parse()

	// 生成二维码字符串并写入文件
	str := fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;H:%t;", *ssid, *t, *p, *h)
	qrcode.WriteFile(str, qrcode.Medium, 256, *ssid+".png")
}
