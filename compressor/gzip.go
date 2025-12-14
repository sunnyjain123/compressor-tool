package compressor

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"strings"
)

type GzipCompressor struct{}

func (g *GzipCompressor) Compress(inputPath string) (string, error) {
	if strings.HasSuffix(inputPath, ".gz") {
		return "", errors.New("file already appears to be compressed")
	}

	in, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer in.Close()

	outputPath := inputPath + ".gz"

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	writer := gzip.NewWriter(out)
	defer writer.Close()

	_, err = io.Copy(writer, in)
	return outputPath, err
}

func (g *GzipCompressor) Decompress(inputPath string) (string, error) {
	if !strings.HasSuffix(inputPath, ".gz") {
		return "", errors.New("file is not a gzip archive")
	}

	in, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer in.Close()

	reader, err := gzip.NewReader(in)
	if err != nil {
		return "", err
	}
	defer reader.Close()

	outputPath := strings.TrimSuffix(inputPath, ".gz")

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, reader)
	return outputPath, err
}
