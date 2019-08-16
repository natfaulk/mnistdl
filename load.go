package mnist

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// Load - Download images and labels
func Load() ([][]uint8, []uint8, error) {
	// load the images
	images, err := ioutil.ReadFile(filepath.Join(cacheDir, "train-images-idx3-ubyte"))
	if err != nil {
		return nil, nil, err
	}

	magicNum := getUint32(images, 0)
	if magicNum != 0x803 {
		return nil, nil, fmt.Errorf("Mnist load: Incorrect magic number. Got 0x%x, Expected 0x%x", magicNum, 0x803)
	}

	nImages := int(getUint32(images, 4))
	nRows := int(getUint32(images, 8))
	nCols := int(getUint32(images, 12))

	imagesOut := make([][]uint8, nImages)

	for n := range imagesOut {
		imagesOut[n] = images[16+n*nRows*nCols : 16+(n+1)*nRows*nCols]
	}

	// load the labels
	labels, err := ioutil.ReadFile(filepath.Join(cacheDir, "train-labels-idx1-ubyte"))
	if err != nil {
		return nil, nil, err
	}

	magicNum = getUint32(labels, 0)
	if magicNum != 0x801 {
		return nil, nil, fmt.Errorf("Mnist load: Incorrect magic number. Got 0x%x, Expected 0x%x", magicNum, 0x801)
	}

	nlabels := int(getUint32(labels, 4))

	labelsOut := labels[8 : 8+nlabels]

	// file, err := os.Open(filepath.Join(cacheDir, "train-images-idx3-ubyte"))
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	return imagesOut, labelsOut, nil
}
