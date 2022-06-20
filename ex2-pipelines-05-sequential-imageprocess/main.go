package main

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
)

// 1. finds the path of the folder containing the images
// 2. check the file types for images/jpeg
// 3. saves the images as thumbnail into memory
// 4. saves the images into a thumbnail folder

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Filepath not found.")
	}

	start := time.Now()

	if err := walkFiles(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Program completed in %s", time.Since(start))

}

func walkFiles(root string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		if contentType, _ := getContentType(path); contentType != "image/jpeg" {
			return nil
		}

		thumbnail, err := processImage(path)
		if err != nil {
			return err
		}

		if err := saveThumbnail(path, thumbnail); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// getContentType
func getContentType(path string) (string, error) {

	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	// get no more than 200 bytes to det. the file type
	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" {
		return "", err
	}

	return contentType, nil
}

// process the image and return as thumnbanil saved in memory
func processImage(path string) (*image.NRGBA, error) {

	file, err := imaging.Open(path)

	if err != nil {
		return nil, err
	}

	thumbnail := imaging.Thumbnail(file, 100, 100, imaging.Lanczos)

	return thumbnail, nil
}

// save an image to from the source to the destination
func saveThumbnail(path string, thumbnail *image.NRGBA) error {

	filename := filepath.Base(path)
	dstImagePath := "thumbnail/" + filename

	fmt.Println(filename)
	fmt.Println(dstImagePath)

	err := imaging.Save(thumbnail, dstImagePath)
	if err != nil {
		return err
	}

	return nil
}
