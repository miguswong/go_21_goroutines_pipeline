# MSDS 431 Week 6 Assignment: Data Pipelines with Concurrency

# Original README
[Episode link](https://www.codeheim.io/courses/Episode-21-Concurrency-in-Go-Pipeline-Pattern-65c3ca14e4b0628a4e002201)

Requires Golang 1.20 or higher.

# Assignment Description

Amrit Singh [(CODEHEIML)](https://www.codeheim.io/) offers an example of a Go image processing pipeline with concurrency. He provides a GitHubLinks to an external site. code repository and [video tutorial](https://www.youtube.com/watch?v=8Rn8yOQH62k). Let's replicate his work using image files that we select.

[![codeHein Pipeline Pattern](https://img.youtube.com/vi/Y8Rn8yOQH62k/0.jpg)](https://www.youtube.com/watch?v=8Rn8yOQH62k)

Clone the GitHub repository for image processing. 
Build and run the program in its original form.
Add error checking for image file input and output.
Replace the four input image files with files of your choosing.
Add unit tests to the code repository.
Add benchmark methods for capturing pipeline throughput times. Design the program so it can be run with and without goroutines. 
Make additional code modifications as you see fit.
Build, test, and run the pipeline program with and without goroutines. Compare processing times with and without goroutines.
Prepare a complete README.md file documenting your work.
(Optional) Note that resizing images can cause distortion. How can we preserve the aspect ratio of images? Suppose we detect the size of the input images in pixels and ensure that the output image has the same shape, rather than the 500x500 shape in the image_processing.Resize helper function. 
