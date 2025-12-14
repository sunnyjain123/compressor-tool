package compressor

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ZipCompressor struct{}

func (z *ZipCompressor) Compress(inputPath string) (string, error) {
	if strings.HasSuffix(inputPath, ".zip") {
		return "", errors.New("file already appears to be a zip archive")
	}

	in, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer in.Close()

	outputPath := inputPath + ".zip"
	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	zipWriter := zip.NewWriter(out)
	defer zipWriter.Close()

	fileInfo, err := in.Stat()
	if err != nil {
		return "", err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return "", err
	}

	header.Name = filepath.Base(inputPath)
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(writer, in)
	return outputPath, err
}

func (z *ZipCompressor) Decompress(inputPath string) (string, error) {
	if !strings.HasSuffix(inputPath, ".zip") {
		return "", errors.New("file is not a zip archive")
	}

	r, err := zip.OpenReader(inputPath)
	if err != nil {
		return "", err
	}
	defer r.Close()

	if len(r.File) != 1 {
		return "", errors.New("zip contains multiple files; unsupported")
	}

	f := r.File[0]

	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()

	outputPath := f.Name

	out, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, rc)
	return outputPath, err
}
