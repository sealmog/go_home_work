package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	var bar Bar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		err := Copy(from, to, offset, limit)
		if err != nil {
			fmt.Println("some error is present...", err)
			os.Exit(1)
		}

		bar.Play(int64(i))
	}
	bar.Finish()
}
