package main

import (
	"fmt"
	"github.com/gosuri/uilive"
	"time"
)

func main() {
	w := uilive.New()
	w.Start()
	defer w.Stop()

	for i := 0; i <= 100; i++ {
		fmt.Fprintf(w, "Downloading... (%d/%d) GB\n", i, 100)
		time.Sleep(time.Millisecond * 5)
	}

	fmt.Fprint(w, "Done downloading")
}
