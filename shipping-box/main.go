package main

import (
	"sort"
	"sync"
)

type Product struct {
	Name   string
	Width  int
	Height int
	Length int
}

type Box struct {
	Width  int
	Height int
	Length int
}

type ProductsMeasurements struct {
	volume    int
	maxHeigth int
	maxLength int
	maxWidth  int
}

// This func will use the non-concurrent approach to evaluate performance
func getBestBox(availableBoxes []Box, products []Product) Box {
	productsMeasures := getProductsMeasurements(products)

	// Search for the smallest possible box that fits all the products
	for _, box := range availableBoxes {
		if checkBoxFits(box, productsMeasures) != nil {
			return box
		}
	}

	// If no box is found, we return an empty box
	return Box{}
}

func getBestBoxConcurrent(availableBoxes []Box, products []Product) *Box {
	input := make(chan *Box)
	output := make(chan []*Box)
	productsMeasures := getProductsMeasurements(products)
	var wg sync.WaitGroup
	defer close(output)
	// Search for the smallest possible box that fits all the products
	go func() {
		var results []*Box
		for ev := range input {
			if ev != nil {
				results = append(results, ev)
			}
			wg.Done()
		}
		output <- results
	}()

	for _, box := range availableBoxes {
		wg.Add(1)
		go func(box Box) {
			boxFound := checkBoxFits(box, productsMeasures)
			input <- boxFound
		}(box)
	}

	wg.Wait()
	close(input)

	result := <-output
	if result == nil {
		return nil
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Width < result[j].Width
	})
	return result[0]
}

func checkBoxFits(box Box, pm ProductsMeasurements) *Box {
	boxVolume := box.Width * box.Length * box.Height
	// If one side of the product is greater than the same side of the box
	// we don't check the volume because one box cannot fit the other
	if box.Height >= pm.maxHeigth && box.Length >= pm.maxLength && box.Width >= pm.maxWidth { // TODO: What if we turn the box?
		if pm.volume <= boxVolume {
			return &box
		}
	}
	return nil
}

func getProductsMeasurements(products []Product) (productMeasures ProductsMeasurements) {
	var productsVolume int
	var mH, mL, mW int
	for _, product := range products {
		// Check the volume of all of the products
		productsVolume += product.Width * product.Length * product.Height
		// Check the maximum dimensions of each side of the product
		if product.Height > mH {
			mH = product.Height
		}
		if product.Length > mL {
			mL = product.Length
		}
		if product.Width > mW {
			mW = product.Width
		}
	}

	return ProductsMeasurements{
		productsVolume,
		mH,
		mL,
		mW,
	}
}
