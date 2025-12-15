# compress-go-tool

A simple, user-friendly file compression CLI written in Go, focused not just on working code — but on **being easy to install, use, upgrade, and uninstall**.

This project accompanies the article:
**“How to Not Just Build Software, but Make It Easy to Use”**

---

## Features

* Compress files into `.zip`
* Decompress files back to original name and extension
* Sensible defaults (no output flags required)
* Universal macOS binaries (Intel + Apple Silicon)
* Multiple installation options (Homebrew, install script, source build)
* Zero runtime dependencies

---

## Prerequisites (For Building From Source)

If you only want to **use** the tool via Homebrew, you can skip this section.

### Install Go

Install Go from the official website:

👉 [https://go.dev/dl/](https://go.dev/dl/)

Verify installation:

```bash
go version
```

You’ll also need:

* Git
* macOS

---

## Installation Options

### Option 1: Homebrew (Recommended for End Users)

This is the easiest and cleanest way to install.

```bash
brew tap sunnyjain123/compress-go-tool
brew install compress-go-tool
```

Verify:

```bash
compress-go-tool --help
```

Upgrade:

```bash
brew upgrade compress-go-tool
```

Uninstall:

```bash
brew uninstall compress-go-tool
```

---

### Option 2: Install Script (Developer-Friendly)

This option is useful if you want to inspect the code or experiment locally.

```bash
git clone https://github.com/sunnyjain123/compress-go-tool.git
cd compress-go-tool
chmod +x install.sh
./install.sh
```

After installation:

```bash
compress-go-tool --help
```

---

### Option 3: Build From Source (Manual)

```bash
git clone https://github.com/sunnyjain123/compress-go-tool.git
cd compress-go-tool
go build -o compress-go-tool
```

Run locally:

```bash
./compress-go-tool file.txt
```

(Optional) Install globally:

```bash
sudo mv compress-go-tool /usr/local/bin/
```

---

## Creating Binaries (For Maintainers / Releases)

### Build for Current Platform

```bash
go build -o compress-go-tool
```

---

### Build macOS Universal Binary (Intel + Apple Silicon)

```bash
# Build ARM64
GOOS=darwin GOARCH=arm64 go build -o compress-go-tool-arm64

# Build AMD64
GOOS=darwin GOARCH=amd64 go build -o compress-go-tool-amd64

# Merge into universal binary
lipo -create \
  compress-go-tool-arm64 \
  compress-go-tool-amd64 \
  -output compress-go-tool
```

Verify:

```bash
lipo -info compress-go-tool
```

---

### Build Linux Binary

```bash
GOOS=linux GOARCH=amd64 go build -o compress-go-tool-linux
```

---

## Usage

### Compress a File

```bash
compress-go-tool example.txt
```

Creates:

```text
example.txt.zip
```

---

### Decompress a File

```bash
compress-go-tool example.txt.zip
```

Restores:

```text
example.txt
```

---

## Related Article

📖 **How to Not Just Build Software, but Make It Easy to Use**

[https://medium.com/@jain.sunny301/how-to-not-just-build-software-but-make-it-easy-to-use-f5f2ed03984d](https://medium.com/@jain.sunny301/how-to-not-just-build-software-but-make-it-easy-to-use-f5f2ed03984d)

---

## License

MIT License
