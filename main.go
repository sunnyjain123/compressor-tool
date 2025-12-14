package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sunnyjain123/compressor-tool/compressor"
)

func main() {
	decompress := flag.Bool("d", false, "decompress file")
	format := flag.String("f", "gzip", "compression format: gz | zip")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage:")
		fmt.Println("  compress file.txt")
		fmt.Println("  compress -d file.txt.gz")
		os.Exit(1)
	}

	var c compressor.Compressor
	switch *format {
	case "zip":
		c = &compressor.ZipCompressor{}
	case "gz":
		c = &compressor.GzipCompressor{}
	default:
		c = &compressor.ZipCompressor{}
		log.Println("Default selection of zip compressor")
	}

	input := flag.Arg(0)
	if *decompress {
		output, err := c.Decompress(input)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Decompressed to:", output)
	} else {
		output, err := c.Compress(input)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Compressed to:", output)
	}
}
