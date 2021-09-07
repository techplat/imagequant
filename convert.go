package imagequant

import (
	"image"
	"image/color"
)

// GoImageToRgba32 convert Go Image to RGBA32 bytes
func GoImageToRgba32(im image.Image) []byte {
	w := im.Bounds().Max.X
	h := im.Bounds().Max.Y
	ret := make([]byte, w*h*4)

	p := 0

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r16, g16, b16, a16 := im.At(x, y).RGBA()

			ret[p+0] = uint8(r16 >> 8)
			ret[p+1] = uint8(g16 >> 8)
			ret[p+2] = uint8(b16 >> 8)
			ret[p+3] = uint8(a16 >> 8)
			p += 4
		}
	}

	return ret
}

// Rgb8PaletteToGoImage convert from RBG8 byte to Go Image
func Rgb8PaletteToGoImage(w, h int, rgb8data []byte, pal color.Palette) image.Image {
	rect := image.Rectangle{
		Max: image.Point{
			X: w,
			Y: h,
		},
	}

	ret := image.NewPaletted(rect, pal)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ret.SetColorIndex(x, y, rgb8data[y*w+x])
		}
	}

	return ret
}

func ImageOptimizer(img image.Image) (image.Image, error) {
	ww := img.Bounds().Max.X
	hh := img.Bounds().Max.Y
	attr, e := NewAttributes()
	if e != nil {
		return nil, e
	}
	defer attr.Release()

	e = attr.SetSpeed(10)
	if e != nil {
		return nil, e
	}
	rgba32data := GoImageToRgba32(img)
	iqm, e := NewImage(attr, string(rgba32data), ww, hh, 0)
	if e != nil {
		return nil, e
	}
	defer iqm.Release()

	res, e := iqm.Quantize(attr)
	if e != nil {
		return nil, e
	}
	defer res.Release()

	rgb8data, e := res.WriteRemappedImage()
	if e != nil {
		return nil, e
	}

	img = Rgb8PaletteToGoImage(res.GetImageWidth(), res.GetImageHeight(), rgb8data, res.GetPalette())

	return img, nil
}
