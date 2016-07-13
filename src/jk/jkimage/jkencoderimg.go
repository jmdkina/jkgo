package jkimage

import (
	"fmt"
	"image"
	// "image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	// "os/exec"
	// "encoding/base64"
	"code.google.com/p/graphics-go/graphics"
	"jk/jklog"
	"path"
	"path/filepath"
	// "strconv"
	"strings"
)

type JKEncoderImage struct {
	ScanPath string // Where to find the images (source image)
	SavePos  string // Where to save the files (dst image)
	Scale    int    // How size to scale (0-100)
	Quality  int    // The quality of compress
}

/**
 * reduce jpeg
 * @save: save new image
 * @name: the full path
 * @enType: 0(Over), 1(Src)
 * @quality: 1-100, greater is better
 */

func (en *JKEncoderImage) JK_reducejpeg(save, name string, enType draw.Op, quality int) error {
	// 1. Open file
	ifile, err := os.Open(name)
	if err != nil {
		return err
	}
	defer ifile.Close()

	// Create file to save
	ofile, err := os.Create(save)
	if err != nil {
		return err
	}
	defer ofile.Close()
	/*
		config, err := jpeg.DecodeConfig(ifile)
		if err != nil {
			jklog.L().Errorln("error : ", err)
		}
		jklog.L().Infof("[%d, %d]\n", config.Width, config.Height)
	*/
	img, err := jpeg.Decode(ifile)
	if err != nil {
		jklog.L().Errorln("error : ", err)
		return err
	}
	bounds := img.Bounds()

	x := bounds.Dx()
	y := bounds.Dy()
	nx := x * en.Scale / 100
	ny := y * en.Scale / 100
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	jklog.L().Infof("[%d, %d] \n", x, y)
	// white := color.RGBA{255, 255, 255, 255}
	// draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	// draw.Draw(m, bounds, img, image.ZP, draw.Src)
	draw.Draw(m, m.Bounds(), img, bounds.Min, enType)

	mm := image.NewRGBA(image.Rect(0, 0, nx, ny))
	graphics.Scale(mm, m)

	err = jpeg.Encode(ofile, mm, &jpeg.Options{quality})
	if err != nil {
		jklog.L().Errorln("error: ", err)
		return err
	}
	return nil
}

func (en *JKEncoderImage) JK_reducepng(save, name string) error {
	ifile, err := os.Open(name)
	if err != nil {
		return err
	}
	defer ifile.Close()

	ofile, err := os.Create(save)
	if err != nil {
		return err
	}
	defer ofile.Close()

	img, err := png.Decode(ifile)
	if err != nil {
		return err
	}
	err = png.Encode(ofile, img)
	if err != nil {
		return err
	}

	return nil
}

func (en *JKEncoderImage) exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

/**
 * @param path:
 * @param postFix: support jpg
 * @param op: draw.Over/draw.Src
 * @param quality: 1-100, greater better
 * @func: generate to a new dir reduced
 */
func (en *JKEncoderImage) JK_convertWithPath() {

	if en.ScanPath == en.SavePos {
		fmt.Println("You must pointer one another place")
		return
	}

	filepath.Walk(en.ScanPath, func(pathf string, f os.FileInfo, err error) error {

		// check if has extern jpg
		var entype = 0
		if strings.HasSuffix(f.Name(), "jpg") || strings.HasSuffix(f.Name(), "JPG") ||
			strings.HasSuffix(f.Name(), "jpeg") || strings.HasSuffix(f.Name(), "JPEG") {
			entype = 1

		} else if strings.HasSuffix(f.Name(), "png") {
			entype = 2

		} else {
			jklog.L().Infoln("I can't convert the img ", f.Name())
			return nil
		}

		pathname := filepath.Dir(pathf)
		// fmt.Println("scaned file : ", srcpath, "-", pathname)

		if !en.exist(en.SavePos) {
			os.MkdirAll(en.SavePos, os.ModeDir|os.ModePerm)
		}
		/*
			savepath := path.Join(en.SavePos, pathname)
			if !en.exist(savepath) {
				os.MkdirAll(savepath, os.ModeType)
			}
		*/

		// newfile := path.Join(en.SavePos, pathname, f.Name())
		newfile := path.Join(en.SavePos, f.Name())
		// newfile := savedir + "/" + pathname + "/" + f.Name()
		// fmt.Println("a file path ", savedir+"/"+pathname+"/")

		srcfile := path.Join(pathname, f.Name())
		// srcfile := pathname + "/" + f.Name()
		fmt.Println("Start convert ", srcfile, " to ", newfile)

		if srcfile == newfile {
			jklog.L().Infoln("Wrong place, please select a another place.")
			return nil
		}

		if entype == 1 {
			err := en.JK_reducejpeg(newfile, srcfile, draw.Src, en.Quality)
			if err != nil {
				fmt.Println(err)
			}
		} else if entype == 2 {
			err := en.JK_reducepng(newfile, srcfile)
			if err != nil {
				fmt.Println(err)
			}
		}

		return nil
	})
}
