package texttokenize

import (
    "github.com/mantyr/goquery"
    "bufio"
    "strings"
    "io"
    "bytes"
)

func NewTokenize() (f *Fields) {
    f = new(Fields)
    f.d = make(map[string]Field)
    f.ignore_empty_line = true
    return
}

func (f *Fields) SetIgnoreEmptyLine(status bool) *Fields {
    f.ignore_empty_line = status
    return f
}

func (f *Fields) SetSelection(sel *goquery.Selection) *Fields {
    f.doc = sel
    return f
}

// Example:
//  Parse() - for self Selection
//  Parse("selection") - for find selection
//  Parse("selection", "selection2", ..., "selectionN") - for more find selection
func (f *Fields) Parse(params ...string) *Fields {
    text := ""
    if len(params) > 0 {
        for _, selection := range params {
            f.doc.Find(selection).Each(func(i int, s *goquery.Selection){
                text += s.Text()+"\r\n"
            })
        }
    } else {
        f.doc.Each(func(i int, s *goquery.Selection){
            text += s.Text()+"\r\n"
        })
    }

    option := "default"
    buf := bufio.NewReader(bytes.NewBufferString(text))
    for {
        l, buffer_err := buf.ReadString('\n') // parse line-by-line
        l = strings.TrimSpace(l)

        if buffer_err != nil {
            if buffer_err != io.EOF {
                f.Error = buffer_err
                return f
            }

            if len(l) == 0 {
                break
            }
        }
        switch {
            case len(l) == 0: // empty line
                f.AddValueLine(option, "")
                continue
            default:
                i := strings.IndexAny(l, ":")
                switch {
                    case i > 0: // option and value
                        option = strings.TrimSpace(l[0:i])
                        value := strings.TrimSpace(l[i+1:])
                        f.AddValueLine(option, value)
                    case i < 1:
                        f.AddValueLine(option, strings.TrimSpace(l))
                }
        }
        if buffer_err == io.EOF {
            break
        }
    }
    return f
}

func (f *Fields) AddValueLine(key, value string) {
    if key == "" {
        return
    }
    if len(value) == 0 && f.ignore_empty_line {
        return
    }

    key_source := key
    key = strings.Replace(key, ":", "", -1)
    key = strings.ToLower(strings.TrimSpace(key))

    v, ok := f.d[key]
    if !ok {
        f.d[key] = Field{
            key_source: key_source,
            value     : value,
        }
    } else {
        v.value += "\r\n"+value
        f.d[key] = v
    }
}

func (f *Fields) Get(key string) string {
    v, ok := f.d[key]
    if !ok {
        return ""
    }
    return strings.TrimSpace(v.value)
}

func (f *Fields) GetItems() map[string]Field {
    return f.d
}