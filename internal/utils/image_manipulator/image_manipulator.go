package image_manipualtor

import (
	"fmt"
	"image"

	//	"image/draw"
	"image/jpeg"
	"net/http"

	"golang.org/x/image/draw"
)

//array com todos os 9 retangulos para uso futuro

// faz o handle do http request para manipular a imagem
func ToRGBA(src image.Image) *image.RGBA {
	b := src.Bounds()
	rgba := image.NewRGBA(b)
	draw.Draw(rgba, b, src, b.Min, draw.Src)
	return rgba
}
func HandleImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return nil, err
	}
	defer resp.Body.Close()

	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil, err
	}
	return img, nil
}

func PixelateRegion(src image.Image, region image.Rectangle, blockSize int) {
	tmp := image.NewRGBA(image.Rect(0, 0, region.Dx()/blockSize, region.Dy()/blockSize))
	draw.ApproxBiLinear.Scale(tmp, tmp.Bounds(), src, region, draw.Over, nil)
	draw.NearestNeighbor.Scale(src.(draw.Image), region, tmp, tmp.Bounds(), draw.Over, nil)
}

func PixelateNRegions(src image.Image, regions []image.Rectangle, blockSize int) image.Image {
	for _, region := range regions {
		PixelateRegion(src, region, blockSize)
	}
	return src
}

func GetAllRects(img image.Image) []image.Rectangle {
	b := img.Bounds()
	w, h := b.Dx()/3, b.Dy()/3
	rects := []image.Rectangle{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			rect := image.Rect(i*w, j*h, (i+1)*w, (j+1)*h)
			rects = append(rects, rect)
		}
	}
	return rects
}

// exemplo de uso
// img, err := handleImage("https://ik.imagekit.io/8jqrt0u9w/1_WRrci-Pva.jpg")
// if err != nil {
// 	return
// }
// b := img.Bounds()
// w, h := b.Dx()/3, b.Dy()/3
// dst := image.NewRGBA(b)
// draw.Draw(dst, b, img, b.Min, draw.Src)

// // Exemplo: pixelar blocos (0,0) e (1,2)
// bs := 15
// pixelateRegion(dst, image.Rect(0, 0, w, h), bs)         // canto superior esquerdo
// pixelateRegion(dst, image.Rect(w, 2*h, 2*w, 3*h), bs)   // linha inferior coluna meio
// pixelateRegion(dst, image.Rect(2*w, 2*h, 3*w, 3*h), bs) //linha inferior coluna direita
// pixelateRegion(dst, image.Rect(0, 2*h, w, 3*h), bs)     // linha inferior coluna esquerda

// fg, err := os.Create("img.jpg")
// if err != nil {
// 	fmt.Println("Creating file:", err)
// 	return
// }
// defer fg.Close()
// err = jpeg.Encode(fg, dst, nil)
// if err != nil {
// 	fmt.Println("Encoding error:", err)
// 	return
// }
