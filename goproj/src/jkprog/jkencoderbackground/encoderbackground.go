package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//原始<strong><strong>图片</strong></strong>是sam.jpg
	imgb, _ := os.Open("sam.jpg")
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	wmb, _ := os.Open("text.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把<strong><strong>水印</strong></strong>写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新<strong><strong>图片</strong></strong>new.jpg，并设置<strong><strong>图片</strong></strong>质量..
	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("<strong><strong>水印</strong></strong>添加结束,请查看new.jpg<strong><strong>图片</strong></strong>...")
}
