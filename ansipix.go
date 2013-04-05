package ansipix

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

type TermColor struct {
	color.Color
	Text string
	FG   int
	BG   int
}

func (t *TermColor) String() string {
	return fmt.Sprintf("\033[%d;%dm%s\033[0m", t.FG, t.BG, t.Text)
}

var (
	TermPalette = &color.Palette{
		//Block colors
		&TermColor{&color.RGBA{0, 0, 0, 255}, " ", 30, 40},
		&TermColor{&color.RGBA{225, 0, 0, 255}, " ", 30, 41},
		&TermColor{&color.RGBA{0, 225, 0, 255}, " ", 30, 42},
		&TermColor{&color.RGBA{225, 225, 0, 255}, " ", 30, 43},
		&TermColor{&color.RGBA{0, 0, 225, 255}, " ", 30, 44},
		&TermColor{&color.RGBA{225, 0, 225, 255}, " ", 30, 45},
		&TermColor{&color.RGBA{0, 225, 225, 255}, " ", 30, 46},
		&TermColor{&color.RGBA{255, 255, 255, 255}, " ", 30, 47},

		//Half block
		&TermColor{&color.RGBA{80, 0, 0, 255}, "░", 31, 40},
		&TermColor{&color.RGBA{0, 80, 0, 255}, "░", 32, 40},
		&TermColor{&color.RGBA{80, 80, 0, 255}, "░", 33, 40},
		&TermColor{&color.RGBA{0, 0, 80, 255}, "░", 34, 40},
		&TermColor{&color.RGBA{80, 0, 80, 255}, "░", 35, 40},
		&TermColor{&color.RGBA{0, 80, 80, 255}, "░", 36, 40},
		&TermColor{&color.RGBA{80, 80, 80, 255}, "░", 37, 40},

		//Quarter block
		&TermColor{&color.RGBA{153, 0, 0, 255}, "▒", 31, 40},
		&TermColor{&color.RGBA{0, 153, 0, 255}, "▒", 32, 40},
		&TermColor{&color.RGBA{153, 153, 0, 255}, "▒", 33, 40},
		&TermColor{&color.RGBA{0, 0, 153, 255}, "▒", 34, 40},
		&TermColor{&color.RGBA{153, 0, 153, 255}, "▒", 35, 40},
		&TermColor{&color.RGBA{0, 153, 153, 255}, "▒", 36, 40},
		&TermColor{&color.RGBA{153, 153, 153, 255}, "▒", 37, 40},

		//Half block black
		&TermColor{&color.RGBA{102, 0, 0, 255}, "░", 30, 41},
		&TermColor{&color.RGBA{0, 102, 0, 255}, "░", 30, 42},
		&TermColor{&color.RGBA{102, 102, 0, 255}, "░", 30, 43},
		&TermColor{&color.RGBA{0, 0, 102, 255}, "░", 30, 44},
		&TermColor{&color.RGBA{102, 0, 102, 255}, "░", 30, 45},
		&TermColor{&color.RGBA{0, 102, 102, 255}, "░", 30, 46},
		&TermColor{&color.RGBA{102, 102, 102, 255}, "░", 30, 47},

		//Half block red
		&TermColor{&color.RGBA{153, 102, 0, 255}, "░", 32, 41},
		&TermColor{&color.RGBA{204, 102, 0, 255}, "░", 33, 41},
		&TermColor{&color.RGBA{153, 0, 102, 255}, "░", 34, 41},
		&TermColor{&color.RGBA{204, 0, 102, 255}, "░", 35, 41},
		&TermColor{&color.RGBA{153, 102, 102, 255}, "░", 36, 41},
		&TermColor{&color.RGBA{204, 102, 102, 255}, "░", 37, 41},
		//Quarter block red
		&TermColor{&color.RGBA{102, 153, 0, 255}, "▒", 32, 41},
		&TermColor{&color.RGBA{204, 153, 0, 255}, "▒", 33, 41},
		&TermColor{&color.RGBA{102, 0, 153, 255}, "▒", 34, 41},
		&TermColor{&color.RGBA{204, 0, 153, 255}, "▒", 35, 41},
		&TermColor{&color.RGBA{102, 153, 153, 255}, "▒", 36, 41},
		&TermColor{&color.RGBA{204, 153, 153, 255}, "▒", 37, 41},

		//Half Block Colors Green
		&TermColor{&color.RGBA{102, 153, 0, 255}, "░", 31, 42},
		&TermColor{&color.RGBA{102, 204, 0, 255}, "░", 33, 42},
		&TermColor{&color.RGBA{0, 153, 102, 255}, "░", 34, 42},
		&TermColor{&color.RGBA{102, 153, 102, 255}, "░", 35, 42},
		&TermColor{&color.RGBA{0, 204, 102, 255}, "░", 36, 42},
		&TermColor{&color.RGBA{102, 204, 102, 255}, "░", 37, 42},
		//Quarter block colors Green
		&TermColor{&color.RGBA{153, 102, 0, 255}, "▒", 31, 42},
		&TermColor{&color.RGBA{102, 204, 0, 255}, "▒", 33, 42},
		&TermColor{&color.RGBA{0, 102, 153, 255}, "▒", 34, 42},
		&TermColor{&color.RGBA{153, 102, 153, 255}, "▒", 35, 42},
		&TermColor{&color.RGBA{0, 204, 153, 255}, "▒", 36, 42},
		&TermColor{&color.RGBA{153, 204, 153, 255}, "▒", 37, 42},

		//Half Block Colors Yellow
		&TermColor{&color.RGBA{204, 153, 0, 255}, "░", 31, 43},
		&TermColor{&color.RGBA{153, 204, 0, 255}, "░", 32, 43},
		&TermColor{&color.RGBA{153, 153, 102, 255}, "░", 34, 43},
		&TermColor{&color.RGBA{204, 153, 102, 255}, "░", 35, 43},
		&TermColor{&color.RGBA{153, 204, 102, 255}, "░", 36, 43},
		&TermColor{&color.RGBA{204, 204, 102, 255}, "░", 37, 43},
		//Quarter block colors Yellow
		&TermColor{&color.RGBA{204, 102, 0, 255}, "▒", 31, 43},
		&TermColor{&color.RGBA{102, 204, 0, 255}, "▒", 32, 43},
		&TermColor{&color.RGBA{102, 102, 153, 255}, "▒", 34, 43},
		&TermColor{&color.RGBA{204, 102, 153, 255}, "▒", 35, 43},
		&TermColor{&color.RGBA{102, 204, 153, 255}, "▒", 36, 43},
		&TermColor{&color.RGBA{204, 204, 153, 255}, "▒", 37, 43},

		//Half Block Blue 
		&TermColor{&color.RGBA{102, 0, 153, 255}, "░", 31, 44},
		&TermColor{&color.RGBA{0, 102, 153, 255}, "░", 32, 44},
		&TermColor{&color.RGBA{102, 102, 153, 255}, "░", 33, 44},
		&TermColor{&color.RGBA{102, 0, 204, 255}, "░", 35, 44},
		&TermColor{&color.RGBA{0, 102, 204, 255}, "░", 36, 44},
		&TermColor{&color.RGBA{102, 102, 204, 255}, "░", 37, 44},
		//Quarter block Blue
		&TermColor{&color.RGBA{153, 0, 102, 255}, "▒", 31, 44},
		&TermColor{&color.RGBA{0, 153, 102, 255}, "▒", 32, 44},
		&TermColor{&color.RGBA{153, 153, 102, 255}, "▒", 33, 44},
		&TermColor{&color.RGBA{153, 0, 204, 255}, "▒", 35, 44},
		&TermColor{&color.RGBA{0, 153, 204, 255}, "▒", 36, 44},
		&TermColor{&color.RGBA{153, 153, 204, 255}, "▒", 37, 44},

		//Half Block Magenta
		&TermColor{&color.RGBA{204, 0, 153, 255}, "░", 31, 45},
		&TermColor{&color.RGBA{153, 102, 153, 255}, "░", 32, 45},
		&TermColor{&color.RGBA{204, 102, 153, 255}, "░", 33, 45},
		&TermColor{&color.RGBA{153, 0, 204, 255}, "░", 34, 45},
		&TermColor{&color.RGBA{153, 102, 204, 255}, "░", 36, 45},
		&TermColor{&color.RGBA{204, 102, 204, 255}, "░", 37, 45},
		//Quarter block Magenta
		&TermColor{&color.RGBA{204, 0, 102, 255}, "▒", 31, 45},
		&TermColor{&color.RGBA{102, 153, 102, 255}, "▒", 32, 45},
		&TermColor{&color.RGBA{204, 153, 102, 255}, "▒", 33, 45},
		&TermColor{&color.RGBA{102, 0, 204, 255}, "▒", 34, 45},
		&TermColor{&color.RGBA{102, 153, 204, 255}, "▒", 36, 45},
		&TermColor{&color.RGBA{204, 153, 204, 255}, "▒", 37, 45},

		//Half Block Cyan
		&TermColor{&color.RGBA{102, 153, 153, 255}, "░", 31, 46},
		&TermColor{&color.RGBA{0, 204, 153, 255}, "░", 32, 46},
		&TermColor{&color.RGBA{102, 204, 153, 255}, "░", 33, 46},
		&TermColor{&color.RGBA{0, 153, 204, 255}, "░", 34, 46},
		&TermColor{&color.RGBA{102, 153, 204, 255}, "░", 35, 46},
		&TermColor{&color.RGBA{102, 204, 204, 255}, "░", 37, 46},
		//Quarter block Cyan
		&TermColor{&color.RGBA{153, 102, 102, 255}, "▒", 31, 46},
		&TermColor{&color.RGBA{0, 204, 102, 255}, "▒", 32, 46},
		&TermColor{&color.RGBA{153, 204, 102, 255}, "▒", 33, 46},
		&TermColor{&color.RGBA{0, 102, 204, 255}, "▒", 34, 46},
		&TermColor{&color.RGBA{153, 102, 204, 255}, "▒", 35, 46},
		&TermColor{&color.RGBA{153, 204, 204, 255}, "▒", 37, 46},

		//Half Block white
		&TermColor{&color.RGBA{204, 153, 153, 255}, "░", 31, 47},
		&TermColor{&color.RGBA{153, 204, 153, 255}, "░", 32, 47},
		&TermColor{&color.RGBA{204, 204, 153, 255}, "░", 33, 47},
		&TermColor{&color.RGBA{153, 153, 204, 255}, "░", 34, 47},
		&TermColor{&color.RGBA{204, 153, 204, 255}, "░", 35, 47},
		&TermColor{&color.RGBA{153, 204, 204, 255}, "░", 36, 47},
		//Quarter block white
		&TermColor{&color.RGBA{204, 102, 102, 255}, "▒", 31, 47},
		&TermColor{&color.RGBA{102, 204, 102, 255}, "▒", 32, 47},
		&TermColor{&color.RGBA{204, 204, 102, 255}, "▒", 33, 47},
		&TermColor{&color.RGBA{102, 102, 204, 255}, "▒", 34, 47},
		&TermColor{&color.RGBA{204, 102, 204, 255}, "▒", 35, 47},
		&TermColor{&color.RGBA{102, 204, 204, 255}, "▒", 36, 47},
	}
)

// Takes in a reference to an image file, and gives it back.. Ansi style!
func Blockify(r io.Reader, columns int) (ret string, err error) {
	im, _, err := image.Decode(r)
	if err != nil {
		return
	}

	ratio := float64(im.Bounds().Max.X-im.Bounds().Min.X) / float64(columns)
	rows := int(float64(im.Bounds().Max.Y-im.Bounds().Min.Y) / ratio)

	for j := 0; j < rows; j++ {
		for i := 0; i < columns; i++ {
			xOffset := im.Bounds().Min.X + int(ratio*float64(i))
			yOffset := im.Bounds().Min.Y + int(ratio*float64(j))
			ret += TermPalette.Convert(im.At(xOffset, yOffset)).(*TermColor).String()
		}
		ret += "\n"
	}
	return
}
