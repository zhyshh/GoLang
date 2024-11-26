package main

import (
    "fmt"
)

// Existing interface that needs to be adapted
type LegacyPrinter interface {
    Print(s string) string
}

// Existing implementation of the LegacyPrinter interface
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) string {
    newMsg := fmt.Sprintf("Legacy Printer: %s", s)
    fmt.Println(newMsg)
    return newMsg
}

// New interface that the client expects to use
type ModernPrinter interface {
    PrintStored() string
}

// Adapter which adapts LegacyPrinter to ModernPrinter
type PrinterAdapter struct {
    legacyPrinter LegacyPrinter
    msg           string
}

func (p *PrinterAdapter) PrintStored() string {
    // Check if the legacy printer is null (nil) and create a new instance if it is
    if p.legacyPrinter == nil {
        p.legacyPrinter = &MyLegacyPrinter{}
    }
    return p.legacyPrinter.Print(p.msg)
}

func main() {
    legacyPrinter := &MyLegacyPrinter{}
    adapter := &PrinterAdapter{
        legacyPrinter: legacyPrinter,
        msg:           "Hello, World!",
    }

    adapter.PrintStored() // Output: Legacy Printer: Hello, World!
}