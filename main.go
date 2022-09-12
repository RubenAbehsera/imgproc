package main

import (
	"flag"
	"fmt"
	"github.com/RubenAbehsera/imgproc/filter"
	"github.com/RubenAbehsera/imgproc/task"
	"time"
)

func main() {

	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolsize = flag.Int("poolsize", 4, "Workers poolsize for the channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "chan":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolsize)
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	}

	start := time.Now()
	err := t.Process()
	if err != nil {
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
