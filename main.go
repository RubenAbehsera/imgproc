package main

import (
	"fmt"
	"github.com/RubenAbehsera/imgproc/filter"
	"github.com/RubenAbehsera/imgproc/task"
	"time"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewChanTask("./images", "output", f, 4)
	start := time.Now()
	err := t.Process()
	if err != nil {
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
