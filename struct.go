package texttokenize

import (
    "github.com/mantyr/goquery"
)

type Fields struct {
    doc   *goquery.Selection
    d     map[string]Field
    Error error

    ignore_empty_line bool
}

type Field struct {
    key_source string
    value      string
}