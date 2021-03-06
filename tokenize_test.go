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

func TestParseTokenize_MultiParse(t *testing.T) {
    fields := NewTokenize()
    fields.LoadFile("./testdata/valuefields.html")
    fields.Parse("#main_test3 .shortText p")
    fields.Parse("#main_test3 .shortText2")

    val := fields.Get("main material")
    if val != "100% polyester" {
        t.Errorf("Error get field, %q", val)
    }

    val = fields.Get("material 2")
    if val != "110% cotton" {
        t.Errorf("Error get field, %q", val)
    }
}

func TestParseBR(t *testing.T) {
    fields := NewTokenize()
    fields.LoadFile("./testdata/valuefields.html")
    fields.Parse("#main_test4 p")

    val := fields.Get("описание")
    if val != "" {
        t.Errorf("Error get field, %q", val)
    }

    val = fields.Get("цвет")
    if val != "Синий" {
        t.Errorf("Error get field, %q", val)
    }
    val = fields.Get("производитель")
    if val != "Вьетнам" {
        t.Errorf("Error get field, %q", val)
    }
    val = fields.Get("релиз! условия заказа товара")
    if val != "- не доступен для заказа по Москве\r\n- только по полной предоплате\r\n- по 1 паре в руки\r\n- акции на данный товар не распространяются" {
        t.Errorf("Error get field, %q", val)
    }
}
