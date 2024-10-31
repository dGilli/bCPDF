package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/dgilli/bCPDF/internal/pdfgen"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: generate_pdf <name> <email> <phone>")
        return
    }

    name := os.Args[1]
    email := os.Args[2]
    phone := os.Args[3]

    outputDir := "./output"
    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        os.Mkdir(outputDir, os.ModePerm)
    }

    outputFile := filepath.Join(outputDir, "business_card.pdf")
    if err := pdfgen.GenerateBusinessCard(name, email, phone, outputFile); err != nil {
        fmt.Println("Error generating PDF:", err)
        return
    }

    fmt.Println("PDF generated at:", outputFile)
}
