# MSDS 431 Week 6 Assignment: Data Pipelines with Concurrency

## Original README
[Episode link](https://www.codeheim.io/courses/Episode-21-Concurrency-in-Go-Pipeline-Pattern-65c3ca14e4b0628a4e002201)

Requires Golang 1.20 or higher.

## Assignment Description

Amrit Singh [(CODEHEIML)](https://www.codeheim.io/) offers an example of a Go image processing pipeline with concurrency. He provides a GitHubLinks to an external site. code repository and [video tutorial](https://www.youtube.com/watch?v=8Rn8yOQH62k). Let's replicate his work using image files that we select.

[![codeHein Pipeline Pattern](https://img.youtube.com/vi/8Rn8yOQH62k/0.jpg)](https://www.youtube.com/watch?v=8Rn8yOQH62k)

* Clone the GitHub repository for image processing. 
* Build and run the program in its original form.
* * See [ORIGINAL_goroutines_pipeline.exe](https://github.com/miguswong/go_21_goroutines_pipeline/blob/main/ORIGINAL_goroutines_pipeline.exe)
* Add error checking for image file input and output.
* Replace the four input image files with files of your choosing.
* Add unit tests to the code repository.
* Add benchmark methods for capturing pipeline throughput times. 
* Design the program so it can be run with and without goroutines. 
* Make additional code modifications as you see fit.
* Build, test, and run the pipeline program with and without goroutines. 
* Compare processing times with and without goroutines.
* Prepare a complete README.md file documenting your work.

(Optional) Note that resizing images can cause distortion. How can we preserve the aspect ratio of images? Suppose we detect the size of the input images in pixels and ensure that the output image has the same shape, rather than the 500x500 shape in the image_processing.Resize helper function. 

## Overview of Development
The original repository for this assignment is from codeHein's [o_21_goroutines_pipeline](https://github.com/code-heim/go_21_goroutines_pipeline). 

The purpose of the program was to resize images down to a maximum of 500x500 and convert them to grayscale. If the images dimensions are not equal to each other (meaning the images is square-sized) then the image's longest dimension will be resized to 500 px and the smaller dimension will bee resized to maintain the aspect ratio.

The application was also designed so that it can be run with or without concurrency. All functions created without concurrency were placed in [noConcurrency.go](https://github.com/miguswong/go_21_goroutines_pipeline/blob/main/noConcurrency.go).

### Using the Application
[*image processor.exe*](https://github.com/miguswong/go_21_goroutines_pipeline/blob/main/imageProcessor.exe) contains the windows executable application. The application will look for an images folder and specifically .jpeg images named image5,image6,image7,image8, respectively.If running the application via *go run*, be sure to include [noConcurrency.go](https://github.com/miguswong/go_21_goroutines_pipeline/blob/main/noConcurrency.go) file. It should look something like:

```
go run main.go noConcurrency.go
```

The user will be prompted by this screen
```
Would you like to run the program with or without concurrency?
1. With Concurrency
2. WithOUT Concurrency
3. Quit Program
```

User Input is validated and the respective portions of the program will run accordingly.

Assuming all images are found, the the modified images will be generated in the images/output file path of your working directory:
```
4 of 4 images successfully processed
```

### Error Checking
The function *isValidImagePath* was created to ensure that all specified images in main.go contain the ccorrect syntax of a /images prefix and .jpeg suffix.

A counter was also added to the original program so that the user is able to see how many images were successfully processed.

### Chosen Images
The original images from codeHeim's project were kept and 4 images of my choosing were added from [unsplash](https://unsplash.com/s/photos/jpeg) and named image5,image6,image7,image8, respectively.

Higher-resolution photos were chosen to demonstrate the potentail savings of of utilizing go routines and channels.  

### Unit Testing
Unit tests were created to test various functions within the program. They can be found in main_test.go.


### Benchmarking Concurrent Processes vs Non-Concurrent Processes
Go's Standard "time" package was used to benchmark processing speed difference between concurrent and non-concurrent execution.

| Concurrent Process (s) | Non-Concurrent Process (s) | Difference (s) |
|----------|----------|----------|
| 2.92    | 3.45   | 0.53   |

Based off benchmarking results, concurrent processing for this application resulted in ~15% faster compared to non-concurrent processing.


### Changes to Resizing Helper Function
Changes were made to the resizing function so that the aspect ratio of the original image is retained rather than always being resized to 500x500 pixels. In the new program, a factor variable is determined based off the current image's dimensions and is used to refactor photos to have a maximum dimension(either Height or Width) to 500 pixels. 

This was demonstrated via [image 6](https://github.com/miguswong/go_21_goroutines_pipeline/blob/main/images/output/image6.jpeg) wich is a landscape image compared to other converted images who which were all taken in the portrait orientation.

