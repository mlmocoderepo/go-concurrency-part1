package main

import (
	"fmt"
	"image"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

type result struct {
	srcImagePath string
	thumbnail    *image.NRGBA
	err          error
}

func main() {

	// check if there are two arguments sent from the command-line
	if len(os.Args) < 2 {
		log.Fatal("Image directory not found\n")
	}

	start := time.Now()

	// set up the pipeline to obtain the folder's name submitted via the command line
	// then check, process (thumbnail) and save the images
	err := pipelineSetup(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time taken to complete: %v", time.Since(start))
}

type StageError struct {
	Stage string
	Err   error
}

func (s *StageError) Error() string {
	return fmt.Sprintf("Stage: %s ,error: %s", s.Stage, s.Err)
}

func pipelineSetup(root string) error {

	// create a channel that discontinues the goroutine process if needed
	done := make(chan struct{})
	// defer close(done)

	// 1. set up stage to walk through the files in folder
	paths, errc := walkFiles(done, root)

	select {
	case err := <-errc:
		if err != nil {
			return &StageError{
				Stage: "WalkFiles",
				Err:   err,
			}
		}
	default:
	}

	// 2. set up stage to process the images in the folder
	results := processImage(done, paths)

	// create a counter to send a close(done) prematurely
	count := 0

	// 3. save the result of images to the folder
	// only if the result does not have any error
	for r := range results {
		if r.err != nil {
			return &StageError{
				Stage: "Processing",
				Err:   r.err,
			}
		}

		// prematurely close the counter by closing the done() channel
		count++
		if count == 5 {
			close(done)
		}

		saveThumbnail(r.srcImagePath, r.thumbnail)
	}

	if err := <-errc; err != nil {
		return &StageError{
			Stage: "WalkFiles",
			Err:   err,
		}
	}

	return nil
}

// walkFiles walks through the files of the directory passed in and returns the path->filename of each file
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {

	// return the paths to be processed later on
	// return any error(s) catptured while detecting the file type (image/jpeg)
	paths := make(chan string)
	errc := make(chan error, 1)

	// invoke a goroutine to walk thorugh the files of the folder sent via the command line
	go func() {
		// will need to close the receiver channel so that it may be read after it is returned
		defer close(paths)

		// walk through the folder that was sent over
		errc <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

			// checks for errors visiting the folder
			if err != nil {
				return err
			}

			// checks if it is a regular file
			if !info.Mode().IsRegular() {
				return nil
			}

			// checks the content type for 'image/jpeg'
			contentType, _ := getContentType(path)
			if contentType != "image/jpeg" {
				return nil
			}

			// select statement checks:
			// proceed if it is a send operation
			// if the process is disrupted (via <-done), return an error indicating the walk function was disrupted
			select {
			case paths <- path:
			case <-done:
				return fmt.Errorf("walk error")
			}

			return nil
		})

	}()

	return paths, errc
}

// getContentType looks for the valid file type "image/jpeg"
func getContentType(path string) (string, error) {

	// open the file received
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	// create buffer no more than 200MB to read the file in bytes
	// read the file into the buffer, and returns if there's an error
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)

	if err != nil {
		return "", fmt.Errorf("reading file %v", err)
	}

	// use http packages to detect the content type: image/jpeg
	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" {
		return "", fmt.Errorf("image file %v", err)
	}

	return contentType, nil
}

// processImage converts all imagepaths into thumbnails and stores into memory first
// fanIn stage happens here
func processImage(done <-chan struct{}, paths <-chan string) <-chan *result {

	// create a channel that stores the imgSrcPath, *image.Image in memory and err - if any
	results := make(chan *result)
	var wg sync.WaitGroup = sync.WaitGroup{}

	// thumbnailed goroutine creates a thumbnail of the image received
	thumbnailed := func() {
		defer wg.Done()

		for imgSrc := range paths {

			// open the file to be thumbnailed
			imageSrcPath, err := imaging.Open(imgSrc)

			// if an error occurs opening the file, store the result without the image and return
			if err != nil {
				select {
				case results <- &result{imgSrc, nil, err}:
				case <-done:
					return
				}
			}

			// if there are no errors, process the image into a 100x100 thumbnail
			thumbnail := imaging.Thumbnail(imageSrcPath, 100, 100, imaging.Lanczos)

			// stores the imgSrcPath, *image.Image in memory and err - if any
			select {
			case results <- &result{imgSrc, thumbnail, err}:
			case <-done:
				return
			}
		}
	}

	// run 5 goroutines to process all the directories in path
	const numThumbnailer = 5
	for i := 0; i < numThumbnailer; i++ {
		wg.Add(1)
		go thumbnailed()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// save image takes the image in memory and saves it to the constructed destination folder
func saveThumbnail(filename string, thumbnail *image.NRGBA) error {

	// create destination path and filename for saving the thumbnail
	srcImagePath := filepath.Base(filename)
	dstImagepath := "thumbnail/" + srcImagePath

	// save the thumbnail in memory sent over to the destination path
	err := imaging.Save(thumbnail, dstImagepath)
	if err != nil {
		return err
	}

	fmt.Printf("%s -> %s\n", srcImagePath, dstImagepath)

	return nil
}
