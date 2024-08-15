/*
Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

package main // import "golang.org/x/tour/pic"

import (
	"bufio"
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"os"
)

// Show displays a picture defined by the function f
// when executed on the Go Playground.
//
// f should return a slice of length dy,
// each element of which is a slice of dx
// 8-bit unsigned int. The integers are
// interpreted as bluescale values,
// where the value 0 means full blue,
// and the value 255 means full white.
func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

// ShowImage displays the image m
// when executed on the Go Playground.
func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice of uint8
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// Allocate each []uint8 inside the [][]uint8
		image[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// Use the function (x + y) / 2 to generate the pixel values
			image[y][x] = uint8((x + y) / 2)
		}
	}
	return image
}

func main() {
	Show(Pic)
}
