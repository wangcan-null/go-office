package row

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
)

func (r *Row) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteByte('<')
    buffer.WriteString(template.TableRowTag)
    buffer.WriteByte('>')

    if nil != r.pr {
        body, err := r.pr.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    for _, cell := range r.GetCells() {
        body, err := cell.GetXmlBytes()
        if nil != err {
            return nil, err
        }

        buffer.Write(body)
    }

    buffer.Write([]byte{'<', '/'})
    buffer.WriteString(template.TableRowTag)
    buffer.WriteByte('>')

    return buffer.Bytes(), nil
}
