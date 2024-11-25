package main

import "fmt"

// Document is an interface that all concrete documents will implement
type Document interface {
    PrintDocument()
}

// WordDocument is a concrete type implementing Document interface
type WordDocument struct{}

func (w WordDocument) PrintDocument() {
    fmt.Println("This is a Word document.")
}

// ExcelDocument is a concrete type implementing Document interface
type ExcelDocument struct{}

func (e ExcelDocument) PrintDocument() {
    fmt.Println("This is an Excel document.")
}

// DocumentFactory is an interface defining the factory method
type DocumentFactory interface {
    CreateDocument() Document
}

// WordDocumentFactory is a factory for creating Word documents
type WordDocumentFactory struct{}

func (w WordDocumentFactory) CreateDocument() Document {
    return WordDocument{}
}

// ExcelDocumentFactory is a factory for creating Excel documents
type ExcelDocumentFactory struct{}

func (e ExcelDocumentFactory) CreateDocument() Document {
    return ExcelDocument{}
}

func main() {
    var docFactory DocumentFactory

    // Use WordDocumentFactory to create a Word document
    docFactory = WordDocumentFactory{}
    wordDoc := docFactory.CreateDocument()
    wordDoc.PrintDocument() // Output: This is a Word document.

    // Use ExcelDocumentFactory to create an Excel document
    docFactory = ExcelDocumentFactory{}
    excelDoc := docFactory.CreateDocument()
    excelDoc.PrintDocument() // Output: This is an Excel document.
}