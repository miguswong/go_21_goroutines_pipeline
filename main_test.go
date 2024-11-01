package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"testing"
)

func TestIsValidImagePath(t *testing.T) {
	paths := []string{
		"images/123.jpeg",
		"images/abcdef.jpeg",
		"images/123.jpg",
		"images/123.png",
	}
	exp := []bool{
		true,
		true,
		false,
		false,
	}
	for i := 0; i < len(paths); i++ {
		if exp[i] != isValidImagePath(paths[i]) {
			t.Errorf("\n%s should be %v, but is %v.", paths[i], exp[i], isValidImagePath(paths[i]))
		}
	}
}

func TestNonConLoadImage(t *testing.T) {
	path := "images/image1.jpeg"
	job := NonConLoadImage(path)
	exp := Job{
		InputPath: "images/image1.jpeg",
		OutPath:   "images/output/image1.jpeg",
		Image:     job.Image, //Need to copy image details since it will be different each time
	}

	if job != exp {
		t.Error("Loaded image struct does not match expected.", job, exp)
	}

}

func TestNonConResize(t *testing.T) {
	job := Job{
		InputPath: "images/image1.jpeg",
		OutPath:   "images/output/image1.jpeg",
		Image:     imageprocessing.ReadImage("images/image1.jpeg"),
	}

	job = nonConResize(job)

	exp := Job{
		InputPath: "images/image1.jpeg",
		OutPath:   "images/output/image1.jpeg",
		Image:     imageprocessing.Resize(imageprocessing.ReadImage("images/image1.jpeg")), //Need to copy image details since it will be different each time
	}

	t.Error(imageprocessing.Resize(imageprocessing.ReadImage("images/image1.jpeg")) == exp.Image)
}
