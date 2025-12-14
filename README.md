# compress-go-tool

A simple, fast **file compression & decompression CLI** written in Go.

* ✅ macOS support (Intel + Apple Silicon)
* ✅ Single binary
* ✅ Homebrew install
* ✅ Manual install via script

This README walks you through **every supported way to install and use the tool**, step by step.

---

## ✨ Features

* Compress files to `.zip` and `.gzip`, default is `.zip`
* Decompress back to original filename and extension
* Clean CLI interface
* No external dependencies

---

## 📦 Prerequisites

* macOS
* Go (only required for building from source)

Check Go:

```bash
go version
```

---

## 🚀 Option 1: Install via Homebrew (Recommended)

This is the **best and easiest way** for most users.

### Step 1: Install Homebrew (if not installed)

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Add Homebrew to PATH (Apple Silicon Macs):

```bash
echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
eval "$(/opt/homebrew/bin/brew shellenv)"
```

Verify:

```bash
brew --version
```

---

### Step 2: Install compress-go-tool

```bash
brew tap sunnyjain123/compress-go-tool
brew install compress-go-tool
```

Verify:

```bash
compress-go-tool --help
```

---

### Upgrade

```bash
brew update
brew upgrade compress-go-tool
```

### Uninstall

```bash
brew uninstall compress-go-tool
```

---

## 🛠 Option 2: Install via install.sh (Manual)

Use this if you don’t want Homebrew.

### Step 1: Clone the repo

```bash
git clone https://github.com/sunnyjain123/compress-go-tool.git
cd compress-go-tool
```

---

### Step 2: Run install script

```bash
chmod +x install.sh
./install.sh
```

This installs the binary to:

```
/usr/local/bin/compress-go-tool
```

Verify:

```bash
compress-go-tool --help
```

---

### Uninstall

```bash
sudo rm /usr/local/bin/compress-go-tool
```

---

## 🧑‍💻 Option 3: Build From Source

### Step 1: Clone

```bash
git clone https://github.com/sunnyjain123/compress-go-tool.git
cd compress-go-tool
```

---

### Step 2: Build

```bash
go build -o compress-go-tool
```

---

### Step 3: Run

```bash
./compress-go-tool --help
```

```bash
  -d bool decompress file
  -f string
        compression format: gz | zip (default "zip")
```

---

## 📂 Usage

### Compress a file

```bash
compress-go-tool -f gz example.txt
```

Creates:

```
example.txt.gz
```

---

### Decompress a file

```bash
compress-go-tool -d -f gz example.txt.gz
```

Restores:

```
example.txt
```

---

## 🧠 Design Notes

* Output filenames are auto-generated
* Original filename and extension are preserved
* Avoids conflicts with system binaries (`compress`)

---

## 🔐 Security & Trust

* Homebrew verifies SHA256 checksums
* Releases are immutable
* No runtime downloads

---

## 🤝 Contributing

PRs and issues are welcome.

---

## 📜 License

MIT License
