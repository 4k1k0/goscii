package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {

	save := flag.Bool("save", false, "Save the ascii art in a file")
	flag.Parse()

	img, err := openFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toASCII(img, *save))

}

func openFile() (image.Image, error) {
	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	extension := http.DetectContentType(file)
	buff := bytes.NewBuffer(file)

	switch extension {
	case "image/jpeg":
		img, err := jpeg.Decode(buff)
		if err != nil {
			log.Fatal(err)
		}
		return img, nil
	case "image/png":
		img, err := png.Decode(buff)
		if err != nil {
			log.Fatal(err)
		}
		return img, nil
	default:
		fmt.Println("unknow")
		return nil, nil
	}
}

func toASCII(img image.Image, save bool) string {

	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	letters := "MND8OZ$7I?+=~:,.."
	table := []byte(letters)
	buff := new(bytes.Buffer)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buff.WriteByte(table[pos])
		}
		_ = buff.WriteByte('\n')
	}

	if save {
		fmt.Println("save was true")
	} else {
		fmt.Println("save was false")
	}

	return string(buff.Bytes())
}
