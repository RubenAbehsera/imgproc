package filter

import (
	"github.com/disintegration/imaging"
	"image/jpeg"
	"os"
)

type Filter interface {
	Process(srcPath, destPath string) error
}

type Grayscale struct{}

type Blur struct{}

func (g Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Grayscale(src)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}

func (b Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	dst := imaging.Blur(src, 15)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}
