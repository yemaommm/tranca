package image

import (
	"code.google.com/p/graphics-go/graphics"
	// "github.com/go-martini/martini"
	"image"
	"net/http"
	// "image/bmp"
	// "bytes"
	// "image/draw"
	"fmt"
	"image/color/palette"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// const (
// 	path     = "2.png"
// 	savepath = "x.jpg"
// 	width    = 160
// 	height   = 220
// )

func Thumbnail(path, savepath string, width, height int) error {
	var dst *image.RGBA
	var img image.Image
	var err error
	var gifimage *gif.GIF
	var imgtype = ""
	if len(strings.Split(path, ".")) >= 2 {
		imgtype = strings.Split(path, ".")[1]
	}
	dx := 0
	dy := 0
	file, err2 := os.Open(path)
	if err2 != nil {
		return err2
	}
	defer file.Close()
	if imgtype != "gif" {
		img, _, err = image.Decode(file)
		imgtype = "image"
		fmt.Println(err)
		if err != nil {
			img, err = png.Decode(file)
			imgtype = "png"
			fmt.Println(err)
			if err != nil {
				img, err = jpeg.Decode(file)
				imgtype = "jpeg"
				fmt.Println(err)
				if err != nil {
					return err
				}
			}
		}
	} else {
		gifimage, err = gif.DecodeAll(file)
		// fmt.Println(err)
		if err == nil {
			imgtype = "gif"
		}
	}

	if imgtype == "gif" {
		var retgif *gif.GIF
		retgif = &gif.GIF{Delay: gifimage.Delay, LoopCount: gifimage.LoopCount, Image: make([]*image.Paletted, 0)}
		for _, i := range gifimage.Image {
			x := i.Bounds().Dx()
			y := i.Bounds().Dy()
			if width != 0 {
				dx = width
				dy = width * y / x
			}
			if height != 0 && dy != 0 && height < dy {
				dx = height * dx / dy
				dy = height
			} else if dy == 0 {
				dx = height * x / y
				dy = height
			}
			if err != nil {
				return err
			}
			if x < dx || y < dy {
				return nil
			}
			gdst := image.NewPaletted(image.Rect(0, 0, dx, dy), palette.Plan9)
			// draw.Draw(dst, r, src, sp, op)
			err = graphics.Thumbnail(gdst, i)
			if err != nil {
				return err
			}
			retgif.Image = append(retgif.Image, gdst)
		}
		f3, err := os.Create(savepath)
		if err != nil {
			return err
		}
		gif.EncodeAll(f3, retgif)
		return nil
	}

	x := img.Bounds().Dx()
	y := img.Bounds().Dy()
	if width != 0 {
		dx = width
		dy = width * y / x
	}
	if height != 0 && dy != 0 && height < dy {
		dx = height * dx / dy
		dy = height
	} else if dy == 0 {
		dx = height * x / y
		dy = height
	}
	if err != nil {
		return err
	}
	if x < dx || y < dy {
		return nil
	}

	dst = image.NewRGBA(image.Rect(0, 0, dx, dy))
	err = graphics.Thumbnail(dst, img)

	if err != nil {
		return err
	}

	f3, err := os.Create(savepath)
	if err != nil {
		return err
	}
	// png.Encode(f3, dst)
	switch imgtype {
	case "image":
		png.Encode(f3, dst)
	case "png":
		png.Encode(f3, dst)
	case "jpeg":
		jpeg.Encode(f3, dst, &jpeg.Options{90})
	}
	return nil
}

func ThumbnailHttp(path string, res http.ResponseWriter, width, height int) error {
	var dst *image.RGBA
	var img image.Image
	var err error
	var gifimage *gif.GIF
	var imgtype = ""
	if len(strings.Split(path, ".")) >= 2 {
		imgtype = strings.Split(path, ".")[1]
	}
	dx := 0
	dy := 0
	file, err2 := os.Open(path)
	if err2 != nil {
		return err2
	}
	defer file.Close()
	if imgtype != "gif" {
		img, _, err = image.Decode(file)
		imgtype = "image"
		// fmt.Println(err)
		if err != nil {
			img, err = png.Decode(file)
			imgtype = "png"
			// fmt.Println(err)
			if err != nil {
				img, err = jpeg.Decode(file)
				imgtype = "jpeg"
				// fmt.Println(err)
				if err != nil {
					return err
				}
			}
		}
	} else {
		gifimage, err = gif.DecodeAll(file)
		// fmt.Println(err)
		if err == nil {
			imgtype = "gif"
		}
	}

	// fmt.Println("type", imgtype, gifimage)
	if imgtype == "gif" {
		var retgif *gif.GIF
		retgif = &gif.GIF{Delay: gifimage.Delay, LoopCount: gifimage.LoopCount, Image: make([]*image.Paletted, 0)}
		for _, i := range gifimage.Image {
			x := i.Bounds().Dx()
			y := i.Bounds().Dy()
			if width != 0 {
				dx = width
				dy = width * y / x
			}
			if height != 0 && dy != 0 && height < dy {
				dx = height * dx / dy
				dy = height
			} else if dy == 0 {
				dx = height * x / y
				dy = height
			}
			if err != nil {
				return err
			}
			if x < dx || y < dy {
				return nil
			}
			gdst := image.NewPaletted(image.Rect(0, 0, dx, dy), palette.Plan9)
			// draw.Draw(dst, r, src, sp, op)
			err = graphics.Thumbnail(gdst, i)
			if err != nil {
				return err
			}
			retgif.Image = append(retgif.Image, gdst)
		}

		return gif.EncodeAll(res, retgif)
	}

	x := img.Bounds().Dx()
	y := img.Bounds().Dy()
	if width != 0 {
		dx = width
		dy = width * y / x
	}
	if height != 0 && dy != 0 && height < dy {
		dx = height * dx / dy
		dy = height
	} else if dy == 0 {
		dx = height * x / y
		dy = height
	}
	if err != nil {
		return err
	}
	if x < dx || y < dy {
		return jpeg.Encode(res, img, &jpeg.Options{90})
	}

	dst = image.NewRGBA(image.Rect(0, 0, dx, dy))
	err = graphics.Thumbnail(dst, img)
	if err != nil {
		return err
	}

	// png.Encode(f3, dst)
	// b := bytes.NewBuffer(nil)

	switch imgtype {
	case "image":
		png.Encode(res, dst)
	case "png":
		png.Encode(res, dst)
	case "jpeg":
		jpeg.Encode(res, dst, &jpeg.Options{90})
	}
	return nil
}
