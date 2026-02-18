# 🏴‍☠️ One Piece Wanted Poster Generator

A Go-based CLI tool to automatically generate wanted posters for One Piece pirates. Created as part of a school project to learn Go programming.

![Banner](assets/banner.png)

## 📋 Features

- **CLI Mode**: Generate a single wanted poster via command-line arguments
- **Batch Mode**: Generate multiple posters from a CSV file
- **Custom Template**: Uses the official One Piece wanted poster design
- **PDF Output**: High-quality PDF files ready to print

## 🚀 Installation

### Prerequisites
- Go 1.22+ installed
- Liberation Sans Bold font (auto-downloaded on Linux)

### Build
```bash
git clone https://github.com/yourusername/MyGeneratorPrime.git
cd MyGeneratorPrime
go mod download
go build -o MyGeneratorPrime
```

## 📖 Usage

### CLI Mode (Single Poster)
```bash
./MyGeneratorPrime -cli -name "LUFFY" -prime "1,500,000,000" -img "assets/luffy.png"
```

**Arguments:**
- `-name`: Pirate's name (will be auto-converted to uppercase)
- `-prime`: Bounty amount (e.g., "1,500,000,000")
- `-img`: Path to pirate's image

### Batch Mode (CSV File)
```bash
./MyGeneratorPrime -file -path "pirates.csv"
```

**CSV Format:**
```csv
BAGGY;300,000,000;assets/baggy.png
LUFFY;1,500,000,000;assets/luffy.png
SHANKS;4,048,900,000;assets/shanks.png
```

## 📁 Project Structure
```
MyGeneratorPrime/
├── main.go                  # Entry point, handles CLI arguments
├── pirate/
│   └── pirate.go           # Pirate struct and validation logic
├── pdf/
│   └── pdf.go              # PDF generation using gopdf library
├── csvparser/
│   └── csvparser.go        # CSV file parsing
├── generateprime/
│   └── generateprime.go    # Orchestrates pirate + PDF generation
├── assets/
│   ├── wantedVierge.jpg    # Template background
│   ├── fonts/              # TrueType fonts
│   └── *.png               # Pirate images
└── output/                 # Generated PDF files
```

## 🧩 How It Works
```
User Input
    ↓
main.go (parses -cli or -file)
    ↓
csvparser.Parse() OR direct args
    ↓
pirate.New() → validates & creates Pirate struct
    ↓
pdf.PdfSaver.Save() → generates PDF with template
    ↓
output/PIRATE_NAME.pdf
```

## 🛠️ Technical Details

### Packages Used
- `github.com/signintech/gopdf` - PDF generation
- Standard library: `flag`, `bufio`, `strings`, `os`

### Key Concepts Demonstrated
- **Package organization**: Separation of concerns (pirate logic, PDF logic, CSV parsing)
- **Interfaces**: `Saver` interface for extensibility
- **Error handling**: Proper Go error propagation with `fmt.Errorf`
- **Input validation**: Name uppercase enforcement, empty field checks

## 📸 Examples

| Input | Output |
|-------|--------|
| `LUFFY; 1,500,000,000; luffy.png` | ![Luffy Poster](examples/luffy_poster.png) |

## 🎓 Learning Outcomes

This project helped me understand:
- Go module system (`go.mod`)
- Working with external libraries
- CLI argument parsing with `flag`
- File I/O operations
- PDF manipulation
- Code organization in packages

## 🤝 Contributing

This is a school project, but feel free to fork and experiment!


## 👤 Author

**GOEFFIER MAXIME**  
Student project - Learning Go through practical application
