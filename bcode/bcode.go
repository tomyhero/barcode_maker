package bcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
)

type Bcode struct {
}

func (self *Bcode) Generate(code string) image.Image {

	bcode, err := code128.Encode(code)
	if err != nil {
		panic(err)
	}

	// 大きさを整える
	bcode, err = barcode.Scale(bcode, 250, 30)
	if err != nil {
		panic(err)
	}

	// フォント配備ディレクトリ指定
	draw2d.SetFontFolder("font")

	// テキスト画像作成
	textImage := image.NewRGBA(image.Rect(0, 0, 250, 50))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(textImage, textImage.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(textImage)
	gc.FillStroke()
	gc.SetFillColor(image.Black)
	gc.SetFontSize(12)
	gc.FillStringAt(code, 50, 20)

	// 箱
	containerImage := imaging.New(270, 70, color.NRGBA{255, 255, 255, 255})

	// バーコード合体
	containerImage = imaging.Paste(containerImage, bcode, image.Pt(10, 10))

	// テキスト合体
	containerImage = imaging.Paste(containerImage, textImage, image.Pt(10, 40))

	return containerImage

}
