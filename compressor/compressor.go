package compressor

type Compressor interface {
	Compress(inputPath string) (string, error)
	Decompress(inputPath string) (string, error)
}
