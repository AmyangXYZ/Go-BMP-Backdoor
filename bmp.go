package main

import (
	"encoding/base64"
	"image"
	"image/color"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

// func main() {
// 	err := write("/home/amyang/Projects/Backdoor-BMP", "test.bmp")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	text, err := read("test.bmp")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(text)
// }

func write(text, outFileName string) error {
	encryptedBar := image.NewRGBA(image.Rect(0, 0, 64, 16))

	sEnc := base64.StdEncoding.EncodeToString([]byte(text))
	data := []byte(sEnc)
	encryptedBar.Pix[0] = uint8(len(sEnc))
	x := 1
	for i := 0; i < len(sEnc); i += 3 {
		var r, g, b uint8 = 0, 0, 0
		r = data[i]
		if i+1 < len(sEnc) {
			g = data[i+1]
			if i+2 < len(sEnc) {
				b = data[i+2]
			}
		}
		encryptedBar.SetRGBA(x, 0, color.RGBA{
			R: r,
			G: g,
			B: b,
		})
		x++
	}

	outputFile, err := os.Create(outFileName)
	defer outputFile.Close()
	if err != nil {
		return err
	}
	err = bmp.Encode(outputFile, encryptedBar)
	if err != nil {
		return err
	}
	return nil
}

func read(bmpFile string) (string, error) {
	existingImageFile, err := os.Open(bmpFile)
	if err != nil {
		return "", err
	}
	defer existingImageFile.Close()

	_, _, err = image.Decode(existingImageFile)
	if err != nil {
		return "", err
	}
	existingImageFile.Seek(0, 0)
	img, err := bmp.Decode(existingImageFile)
	if err != nil {
		return "", err
	}

	no, _, _, _ := img.At(0, 0).RGBA()
	s := ""
	for i := 0; i < int(uint(no>>8)); i++ {
		r, g, b, _ := img.At(i+1, 0).RGBA()
		s += string(uint8(r>>8)) + string(uint8(g>>8)) + string(uint8(b>>8))
	}
	s = strings.Replace(s, string(byte(00)), "", -1)
	text, err := base64.StdEncoding.DecodeString(s)
	return string(text), nil
}
