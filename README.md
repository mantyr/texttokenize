# Golang HTML Text Tokenize - alfa-alfa version

[![Build Status](https://travis-ci.org/mantyr/texttokenize.svg?branch=master)](https://travis-ci.org/mantyr/texttokenize) [![GoDoc](https://godoc.org/github.com/mantyr/texttokenize?status.png)](http://godoc.org/github.com/mantyr/texttokenize) [![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This don't stable version

## Installation

    $ go get github.com/mantyr/texttokenize
    $ go get github.com/mantyr/goquery

## Example
```HTML
    <div id="main_test2">
        <div class="descriptionText">
            <p>Описание:</p>
            <p>Декоративная шкатулка<p>
            <p style="margin-top:7px;">Цвет: &nbsp;Зелёный, белый</p>
            <p>Материал:&nbsp;<p>Дерево</p> </p><p>Производитель:&nbsp;&nbsp;Россия</p></div>
    </div>
```


```GO
package main

import (
    "github.com/mantyr/texttokenize"
    "fmt"
)

func main() {
    fields := texttokenize.NewTokenize()
    fields.LoadFile("./testdata/valuefields.html")
//    fields.SetIgnoreEmptyLine(true)                // default value "true"
    fields.Parse("#main_test2 .descriptionText p")

    fmt.Println(fields.Get("описание"))       // print "Декоративная шкатулка"
    fmt.Println(fields.Get("цвет"))           // print "Зелёный, белый"
    fmt.Println(fields.Get("материал"))       // print "Дерево"
    fmt.Println(fields.Get("производитель"))  // print "Россия"

    fmt.Println(fields.Get("большой текст"))  // print "line 1\r\nline 2\r\nline 3" or "line 1\r\n\r\nline 2\r\n\r\nline 3" if fields.SetIgnoreEmptyLine(false)
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
