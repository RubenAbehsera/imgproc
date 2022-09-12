# imgproc
Apply filters (graysclae, blur) on pics from a folder using golang

# Use

Run in your terminal

<b>Using Go Channels</b></br>
`go run main.go -src images/ -dst output/ -filter grayscale -task chan -poolsize 4`

or

<b>Using Go Waitgroup</b></br>
`go run main.go -src images/ -dst output/ -filter grayscale -task waitgrp`

