package pdfgen

import (
    "github.com/jung-kurt/gofpdf"
    "fmt"
)

func GenerateBusinessCard(name, email, phone, outputPath string) error {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    // Set name
    pdf.SetXY(10, 10)
    pdf.Cell(40, 10, fmt.Sprintf("Name: %s", name))

    // Set email
    pdf.SetXY(10, 20)
    pdf.Cell(40, 10, fmt.Sprintf("Email: %s", email))

    // Set phone
    pdf.SetXY(10, 30)
    pdf.Cell(40, 10, fmt.Sprintf("Phone: %s", phone))

    // Save PDF
    return pdf.OutputFileAndClose(outputPath)
}
