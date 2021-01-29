package main

import (
"fmt"

qrcode "github.com/skip2/go-qrcode"
)





func main() {
	var png []byte
	module, err := qrcode.GetModule("https://www.reddit.com/r/eurovision/comments/knnr5x/piccolo_grandi_little_big_mamma_maria_ricchi_e/", qrcode.Low, 256, "test.png")

	if module == nil {
		fmt.Print("could not generate QRCode: module is nil")
	}
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	fmt.Sprint(png)
}
