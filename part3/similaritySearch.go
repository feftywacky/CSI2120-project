package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	// "sort"
	"sync"
	"time"
)

type Histo struct {
	Name string
	H    []float64
}

func computeHistogram(imagePath string, depth int) (Histo, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return Histo{}, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return Histo{}, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	histogram := make([]float64, 1<<(3*depth))
	mask := (1 << depth) - 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r = (r >> (16 - depth)) & uint32(mask)
			g = (g >> (16 - depth)) & uint32(mask)
			b = (b >> (16 - depth)) & uint32(mask)
			index := (int(r) << (2*depth)) + (int(g) << depth) + int(b)
			histogram[index]++
		}
	}

	// normalize the histogram
	totalPixels := float64(width * height)
	for i := range histogram {
		histogram[i] /= totalPixels
	}

	return Histo{Name: filepath.Base(imagePath), H: histogram}, nil
}

func computeHistograms(imagePaths []string, depth int, hChan chan<- Histo) {
	for _, imagePath := range imagePaths {
		histo, err := computeHistogram(imagePath, depth)
		if err != nil {
			fmt.Println("Error computing histogram:", err)
			continue
		}
		hChan <- histo
	}
}

// func main() {
// 	var mutex sync.Mutex
// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: go run similaritySearch.go queryImageFilename imageDatasetDirectory")
// 		return
// 	}

// 	queryImageFilename := os.Args[1]
// 	imageDatasetDirectory := os.Args[2]

// 	// compute the histogram for the query image
// 	queryHisto, err := computeHistogram(queryImageFilename, 3)
// 	if err != nil {
// 		fmt.Println("Error computing histogram for query image:", err)
// 		return
// 	}

// 	imageFiles, err := filepath.Glob(filepath.Join(imageDatasetDirectory, "*.jpg"))
// 	if err != nil {
// 		fmt.Println("Error reading image dataset directory:", err)
// 		return
// 	}

// 	hChan := make(chan Histo, len(imageFiles))
// 	defer close(hChan)

// 	var wg sync.WaitGroup

// 	K := 1048 // NUMBER OF SLICES, CHANGE THIS VALUE TO EXPERIMENT WITH EXECUTION TIMEs
	
// 	// split the list of image files into K slices and send each slice to the go function computeHistograms
// 	sliceSize := len(imageFiles) / K
// 	for i := 0; i < K; i++ {
// 		start := i * sliceSize
// 		end := start + sliceSize
// 		if i == K-1 {
// 			end = len(imageFiles) // include all files in the last slice
// 		}
// 		wg.Add(1)
// 		go func(imagePaths []string) {
// 			defer wg.Done()
// 			computeHistograms(imagePaths, 3, hChan)
// 		}(imageFiles[start:end])
// 	}

// 	// hashmap to store the similarity scores
// 	similarityScores := make(map[string]float64)

// 	// read the channel of histograms and compare them to the query histogram
// 	go func() {
// 		for histo := range hChan {
// 			score := compareHistograms(queryHisto.H, histo.H)
// 			mutex.Lock()
// 			similarityScores[histo.Name] = score
// 			mutex.Unlock()
// 		}
// 	}()

// 	wg.Wait() // wait for all goroutines to finish

// 	// sort the images based on similarity score
// 	type imageScore struct {
// 		Name  string
// 		Score float64
// 	}
// 	var sortedScores []imageScore

// 	mutex.Lock()
// 	for name, score := range similarityScores {
// 		sortedScores = append(sortedScores, imageScore{Name: name, Score: score})
// 	}
// 	mutex.Unlock()

// 	sort.Slice(sortedScores, func(i, j int) bool {
// 		return sortedScores[i].Score > sortedScores[j].Score
// 	})

// 	fmt.Println("The 5 most similar images to", queryImageFilename, "are:")
// 	for i := 4; i >= 0; i-- {
// 		fmt.Printf("#%d: %s (similarity: %f)\n", i+1, sortedScores[i].Name, sortedScores[i].Score)
// 	}
// }

func main() {
    var mutex sync.Mutex
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run similaritySearch.go imageDatasetDirectory")
        return
    }

    imageDatasetDirectory := os.Args[1]

    // Get the list of all query image filenames
    queryImageFiles, err := filepath.Glob("../part1/queryImages/*.jpg")
    if err != nil {
        fmt.Println("Error reading query image directory:", err)
        return
    }

    // Get the list of all image filenames in the dataset directory
    imageFiles, err := filepath.Glob(filepath.Join(imageDatasetDirectory, "*.jpg"))
    if err != nil {
        fmt.Println("Error reading image dataset directory:", err)
        return
    }

    var totalExecutionTime time.Duration // Variable to accumulate total execution time
	K := 1048 // Change this value for each experiment
    // Iterate over each query image
    for _, queryImageFilename := range queryImageFiles {
        // Compute the histogram for the query image
        queryHisto, err := computeHistogram(queryImageFilename, 3)
        if err != nil {
            fmt.Println("Error computing histogram for query image:", err)
            continue
        }

        // Create the channel of histograms
        hChan := make(chan Histo, len(imageFiles))
        defer close(hChan)

        // Start timing
        startTime := time.Now()

        // Split the list of image files into K slices and send each slice to the go function computeHistograms
        var wg sync.WaitGroup
        sliceSize := len(imageFiles) / K
        for i := 0; i < K; i++ {
            start := i * sliceSize
            end := start + sliceSize
            if i == K-1 {
                end = len(imageFiles) // Make sure to include all files in the last slice
            }
            wg.Add(1)
            go func(imagePaths []string) {
                defer wg.Done()
                computeHistograms(imagePaths, 3, hChan)
            }(imageFiles[start:end])
        }

        // Create a map to store the similarity scores
        similarityScores := make(map[string]float64)

        // Read the channel of histograms and compare them to the query histogram
        go func() {
            for histo := range hChan {
                score := compareHistograms(queryHisto.H, histo.H)
                mutex.Lock()
                similarityScores[histo.Name] = score
                mutex.Unlock()
            }
        }()

        wg.Wait() // Wait for all goroutines to finish

        // End timing
        endTime := time.Now()
        executionTime := endTime.Sub(startTime)

        // Accumulate the total execution time
        totalExecutionTime += executionTime

        // Print the execution time for the current query image
        fmt.Println("Execution time for query image", filepath.Base(queryImageFilename), "with K =", K, ":", executionTime)
    }

    // Compute and print the average execution time
    averageExecutionTime := totalExecutionTime / time.Duration(len(queryImageFiles))
    fmt.Println("Average execution time for", len(queryImageFiles), "query images with K =", K, ":", averageExecutionTime)
}


func compareHistograms(h1, h2 []float64) float64 {
	// Compare the histograms using histogram intersection
	minSum := 0.0
	for i := 0; i < len(h1); i++ {
		if h1[i] < h2[i] {
			minSum += h1[i]
		} else {
			minSum += h2[i]
		}
	}
	return minSum
}

