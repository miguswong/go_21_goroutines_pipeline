package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"strings"
)

func NonConLoadImage(input string) Job {
	//Create job struct for path
	job := Job{InputPath: input,
		OutPath: strings.Replace(input, "images/", "images/output/", 1)}
	job.Image = imageprocessing.ReadImage(input)
	return job
}

func nonConResize(input Job) Job {
	input.Image = imageprocessing.Resize(input.Image)
	return input
}

func nonConConvertToGrayScale(input Job) Job {
	input.Image = imageprocessing.Grayscale(input.Image)
	return input
}

func nonConSaveImage(input Job) bool {
	imageprocessing.WriteImage(input.OutPath, input.Image)
	return true
}
