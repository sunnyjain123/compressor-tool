package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sunnyjain123/compress-go-tool/compressor"
)

func main() {
	format := flag.String("f", "zip", "compression format: gz | zip")
	flag.Parse()
	// examples if someone does not provide any args
	if flag.NArg() != 1 {
		fmt.Println("Usage:")
		fmt.Println("  compress-go-tool file.txt")
		fmt.Println("  compress-go-tool file.txt.gz")
		os.Exit(1)
	}

	// Get input file path
	input := flag.Arg(0)
	var decompress bool
	decompress, format = getFormatAndMethod(input, format)

	var c compressor.Compressor
	switch *format {
	case "zip":
		c = &compressor.ZipCompressor{}
	case "gz":
		c = &compressor.GzipCompressor{}
	default:
		log.Println("No format provided")
	}

	if decompress {
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

// Get format and method based on input file
func getFormatAndMethod(inputFile string, format *string) (bool, *string) {
	flag.Parse()

	decompress := false
	if strings.HasSuffix(inputFile, ".gz") {
		decompress = true
		*format = "gz"
	}

	if strings.HasSuffix(inputFile, ".zip") {
		decompress = true
		*format = "zip"
	}

	return decompress, format
}
