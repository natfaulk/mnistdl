package mnist

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const (
	imagesLink string = "http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz"
	labelsLink string = "http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz"
	cacheDir   string = "filecache"
)

// Download - download mnist files
func Download(useCache bool) error {
	if useCache {
		// if cache intact return
		if verifyCache() {
			fmt.Println("Using cached files")
			return nil
		}

		fmt.Println("Invalid cache")
	}

	fmt.Println("Downloading files...")
	// if not use cache or cache not intact delete it
	removeCache()

	// make folder to hold the files, ignore the error if it already exists
	_ = os.Mkdir(cacheDir, os.ModeDir)

	if err := downloadFile(imagesLink); err != nil {
		return err
	}

	if err := downloadFile(labelsLink); err != nil {
		return err
	}

	fmt.Println("Decompressing files...")
	if err := ungzip(filepath.Join(cacheDir, path.Base(imagesLink))); err != nil {
		return err
	}

	if err := ungzip(filepath.Join(cacheDir, path.Base(labelsLink))); err != nil {
		return err
	}

	fmt.Println("Complete")

	return nil
}

func downloadFile(url string) error {
	// Get the data
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Create the file
	out, err := os.Create(filepath.Join(cacheDir, path.Base(url)))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)
	return err
}

func removeCache() {
	if _, err := os.Stat(cacheDir); !os.IsNotExist(err) {
		err = os.RemoveAll(cacheDir)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func verifyCache() bool {
	filename := filepath.Join(cacheDir, trimExtension(path.Base(imagesLink)))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	filename = filepath.Join(cacheDir, trimExtension(path.Base(labelsLink)))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func ungzip(filename string) error {
	fileIn, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	res, err := gzip.NewReader(fileIn)
	if err != nil {
		return err
	}
	defer res.Close()

	// Create the file
	fileOut, err := os.Create(trimExtension(filename))
	if err != nil {
		return err
	}
	defer fileOut.Close()

	io.Copy(fileOut, res)
	return nil
}
