package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// 1 新建构造体

type Image struct{}

// 2 实现官方image的三个方法

func (i Image) ColorModel() color.Model {
	// 实现Image包中颜色模式的方法
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	// 实现Image包中生成图片边界的方法
	// 这里的200 （宽、高） 我写死了 仅仅是展示作用 正确的做法是从i中获取
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {
	// 实现Image包中生成图像某个点的方法
	return color.RGBA{uint8(x), uint8(y), uint8(255), uint8(255)}
}
func main() {
	// 可以自己设置宽高，传递进去
	m := Image{}
	// 3 调用
	pic.ShowImage(m)
}
