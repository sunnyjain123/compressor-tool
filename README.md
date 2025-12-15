# compress-go-tool

A simple, user-friendly file compression CLI written in Go, focused not just on working code — but on **being easy to install, use, upgrade, and uninstall**.

This project accompanies the article:
**“How to Not Just Build Software, but Make It Easy to Use”**

---

## Features

* Compress files into `.zip` and `.gz`
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

## Releasing & Updating Homebrew (Maintainers)

This section is intended for **maintainers**, not end users.

The release flow is:

1. Build binaries
2. Create a GitHub release
3. Update the Homebrew formula

---

### Step 1: Create Release Binaries

Ensure you have taken a fork are on a clean main branch:

```bash
git checkout main
git pull origin main
```

Build macOS universal binary:

```bash
# Build ARM64
GOOS=darwin GOARCH=arm64 go build -o build/compress-go-tool-arm64

# Build AMD64
GOOS=darwin GOARCH=amd64 go build -o build/compress-go-tool-amd64

# Merge into universal binary
lipo -create \
  build/compress-go-tool-arm64 \
  build/compress-go-tool-amd64 \
  -output build/compress-go-tool
```

Verify:

```bash
lipo -info build/compress-go-tool
```

Create archives:

```bash
tar -czf compress-go-tool-darwin-universal.tar.gz compress-go-tool
```

Generate SHA256:

```bash
shasum -a 256 compress-go-tool-darwin-universal.tar.gz
```
---

### Step 2: Create GitHub Release

1. Go to **GitHub → Releases → New Release**
2. Tag version (example: `v1.1.0`)
3. Upload release assets:
   * `compress-go-tool-darwin-universal.tar.gz`
4. Publish the release

These assets will be consumed by Homebrew.

---

### Step 3: Update Homebrew Formula

Edit the formula in your tap (you can create a fork from https://github.com/sunnyjain123/homebrew-compress-go-tool):

```bash
cd homebrew-compress-go-tool
vim Formula/compress-go-tool.rb
```

Update:

* `url` → GitHub release asset URL
* `sha256` → value generated above
* `version` → new version

Example:

```ruby
url "https://github.com/sunnyjain123/compress-go-tool/releases/download/v1.1.0/compress-go-tool-darwin-universal.tar.gz"
sha256 "<SHA256_VALUE>"
version "1.1.0"
```

Commit and push:

```bash
git add .
git commit -m "compress-go-tool v1.1.0"
git push origin main
```

---

### Step 4: Verify Homebrew Install

```bash
brew update
brew upgrade compress-go-tool
compress-go-tool --version
```

Once verified, the release is live.

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
