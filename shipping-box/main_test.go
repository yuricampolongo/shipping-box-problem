package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	availableBoxes = []Box{
		{10, 10, 20},
		{10, 20, 20},
		{15, 20, 25},
		{15, 30, 50},
		{30, 30, 60},
		{40, 40, 40},
		{50, 40, 45},
		{60, 60, 50},
	}

	productsSmallSize = []Product{ //Expected box 3
		{"Product 1", 10, 8, 21},
		{"Product 2", 10, 8, 20},
		{"Product 3", 10, 4, 20},
	}

	productsMediumSize = []Product{ //Expected box 4
		{"Product 1", 15, 10, 50},
		{"Product 2", 15, 10, 50},
		{"Product 3", 15, 10, 50},
	}

	productsLargeSize = []Product{ //Expected box 8
		{"Product 1", 20, 60, 50},
		{"Product 2", 20, 60, 50},
		{"Product 3", 20, 60, 50},
	}

	productsExtraLarge = []Product{ //Expected no box to be found
		{"Product 1", 20, 60, 70},
		{"Product 2", 20, 60, 70},
		{"Product 3", 20, 60, 70},
	}
)

func TestGetBestBoxNonConcurrent(t *testing.T) {
	box := getBestBox(availableBoxes, productsSmallSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[2].Height, box.Height)
	assert.EqualValues(t, availableBoxes[2].Length, box.Length)
	assert.EqualValues(t, availableBoxes[2].Width, box.Width)

	box = getBestBox(availableBoxes, productsMediumSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[3].Height, box.Height)
	assert.EqualValues(t, availableBoxes[3].Length, box.Length)
	assert.EqualValues(t, availableBoxes[3].Width, box.Width)

	box = getBestBox(availableBoxes, productsLargeSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[7].Height, box.Height)
	assert.EqualValues(t, availableBoxes[7].Length, box.Length)
	assert.EqualValues(t, availableBoxes[7].Width, box.Width)
}

func TestGetBestBoxNonConcurrentNoBoxAvailable(t *testing.T) {
	box := getBestBox(availableBoxes, productsExtraLarge)
	assert.NotNil(t, box)
	assert.EqualValues(t, 0, box.Height)
	assert.EqualValues(t, 0, box.Length)
	assert.EqualValues(t, 0, box.Width)
}

func TestGetBestBoxConcurrent(t *testing.T) {
	box := getBestBoxConcurrent(availableBoxes, productsSmallSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[2].Height, box.Height)
	assert.EqualValues(t, availableBoxes[2].Length, box.Length)
	assert.EqualValues(t, availableBoxes[2].Width, box.Width)

	box = getBestBoxConcurrent(availableBoxes, productsMediumSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[3].Height, box.Height)
	assert.EqualValues(t, availableBoxes[3].Length, box.Length)
	assert.EqualValues(t, availableBoxes[3].Width, box.Width)

	box = getBestBoxConcurrent(availableBoxes, productsLargeSize)
	assert.NotNil(t, box)
	assert.EqualValues(t, availableBoxes[7].Height, box.Height)
	assert.EqualValues(t, availableBoxes[7].Length, box.Length)
	assert.EqualValues(t, availableBoxes[7].Width, box.Width)
}

func TestGetBestBoxConcurrentNoBoxAvailable(t *testing.T) {
	box := getBestBoxConcurrent(availableBoxes, productsExtraLarge)
	assert.Nil(t, box)
}

func BenchmarkNonConcurrentBoxAvailable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getBestBox(availableBoxes, productsSmallSize)
	}
}

func BenchmarkConcurrentBoxAvailable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getBestBoxConcurrent(availableBoxes, productsSmallSize)
	}
}
