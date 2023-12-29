package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	ssid := "your_ssid"
	t := "WPA2" // WPA2 WPA WEP NONE
	p := "12345678"
	h := "false"
	str := fmt.Sprintf("WIFI:S:%s;T:%s;P:%s;H:%s;", ssid, t, p, h)
	qrcode.WriteFile(str, qrcode.Medium, 256, ssid+".png")
}
