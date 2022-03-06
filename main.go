package main

import (
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments")
	}
	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	black := color.NRGBA{0, 0, 0, 255}
	bounds := img.Bounds()
	maxy := bounds.Max.Y
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < maxy; y++ {
			if reflect.DeepEqual(img.At(x, y), black) {
				fmt.Println(x, maxy-y)
			}
		}
	}
}
