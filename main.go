package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strconv"
	"strings"
	"time"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func isValidImagePath(path string) bool {
	//Check path begins with "images/" and ends with .jpg

	if strings.HasPrefix(path, "images/") && strings.HasSuffix(path, ".jpeg") {
		return true
	}
	fmt.Printf("The following image path \"%s\" is not in the correct format. Image will not be processed\n", path)
	return false
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			if !isValidImagePath(p) { //Ensure that image path is a .jpeg
				continue
			}

			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	imagePaths := []string{
		"images/image5.jpeg",
		"images/image6.jpeg",
		"images/image7.jpeg",
		"images/image8.jpeg",
	}
	// prompt user if they would like to run the program with concurrency or without.
	var userInput string
	var imagesProcessed uint = 0

	for {
		fmt.Println("Would you like to run the program with or without concurrency?\n1. With Concurrency\n2. WithOUT Concurrency\n3. Quit Program")

		fmt.Scan(&userInput)
		choice, err := strconv.Atoi(userInput)
		if err == nil && choice >= 1 && choice <= 3 {
			break
		}
		fmt.Println("Please choose a valid response (choices are 1-3).")
	}

	choice, _ := strconv.Atoi(userInput)

	if choice == 1 {
		start := time.Now() //Record Start time
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if success {
				imagesProcessed++
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}
		}

		duration := time.Since(start)
		fmt.Printf("%v of %v images successfully processed\nTotal Execution time: %v\n", imagesProcessed, len(imagePaths), duration)

	}

	if choice == 2 {
		start := time.Now() //Record Start time
		//Insert Non-concurrency code here
		for _, p := range imagePaths {

			if !isValidImagePath(p) { //Ensure that image path is a .jpeg
				continue
			}

			currentJob := NonConLoadImage(p)
			currentJob = nonConResize(currentJob)
			currentJob = nonConConvertToGrayScale(currentJob)
			if nonConSaveImage(currentJob) {
				imagesProcessed++
				fmt.Println("Success!")
			} else {
				fmt.Println("Failed!")
			}

		}

		duration := time.Since(start)
		fmt.Printf("%v of %v images successfully processed\nTotal Execution time: %v\n", imagesProcessed, len(imagePaths), duration)
	}

	if choice == 3 {
		fmt.Println("Exiting program. Goodbye!")
	}

}
