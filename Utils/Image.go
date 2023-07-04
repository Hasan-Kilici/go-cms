package Utils

import (
    "image/jpeg"
    "image/png"
	"path/filepath"
	"image/draw"
	"os"
	"fmt"
	"golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
    "image"
    "image/color"
	"github.com/nfnt/resize"
)


func SaveResizedImage(img image.Image, path string) error {
    resized := resize.Resize(500, 0, img, resize.Lanczos3)

    col := color.RGBA{255, 255, 255, 255}
    point := fixed.Point26_6{fixed.I(20), fixed.I(30)}

    rgbaResized := image.NewRGBA(resized.Bounds())
    draw.Draw(rgbaResized, rgbaResized.Bounds(), resized, image.Point{}, draw.Src)
    d := &font.Drawer{
        Dst:  rgbaResized,
        Src:  image.NewUniform(col),
        Face: basicfont.Face7x13,
        Dot:  point,
    }
    d.DrawString("watermark")

    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer f.Close()

    switch filepath.Ext(path) {
    case ".jpg":
        return jpeg.Encode(f, rgbaResized, nil)
    case ".png":
        return png.Encode(f, rgbaResized)
    default:
        return fmt.Errorf("Desteklenmiyen Dosya Tipi")
    }
}