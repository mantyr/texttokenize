package texttokenize

import (
    "github.com/mantyr/goquery"
    "os"
)

func (f *Fields) LoadFile(address string) error {
    var file *os.File
    file, f.Error = os.Open(address)
    if f.Error != nil {
        return f.Error
    }
    defer file.Close()

    var doc *goquery.Document
    doc, f.Error = goquery.NewDocumentFromReader(file)
    if f.Error == nil {
        f.doc = doc.Clone()
    }
    return f.Error
}