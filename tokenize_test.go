package texttokenize

import (
    "testing"
)

func TestParseTokenize(t *testing.T) {
    fields := NewTokenize()
    fields.LoadFile("./testdata/valuefields.html")
    fields.SetIgnoreEmptyLine(false)
    fields.Parse("#main_test1 .descriptionText p")

    val := fields.Get("oписание")
    if val != "" {
        t.Errorf("Error get field, %q", val)
    }

    val = fields.Get("цвет")
    if val != "Зелёный, белый" {
        t.Errorf("Error get field, %q", val)
    }

    val = fields.Get("большой текст")
    if val != "line 1\r\n\r\nline 2\r\n\r\nline 3"{
        t.Errorf("Error get field, %q", val)
    }
}

func TestParseTokenizeMultiline(t *testing.T) {
    fields := NewTokenize()
    fields.LoadFile("./testdata/valuefields.html")
//    fields.SetIgnoreEmptyLine(true) // default value "true"
    fields.Parse("#main_test2 .descriptionText p")

    val := fields.Get("описание")
    if val != "Декоративная шкатулка" {
        t.Errorf("Error get field multi-line, %q", val)
    }

    val = fields.Get("цвет")
    if val != "Зелёный, белый" {
        t.Errorf("Error get field, %q", val)
    }
    val = fields.Get("большой текст")
    if val != "line 1\r\nline 2\r\nline 3"{
        t.Errorf("Error get field, %q", val)
    }
}
