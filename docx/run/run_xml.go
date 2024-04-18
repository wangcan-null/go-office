package run

import (
	"bytes"
	"github.com/wangcan-null/go-office/docx/template"
)

func (r *Run) GetXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.RunStart)

	body, err := r.GetProperties().GetInnerXmlBytes()
	if nil != err {
		return nil, err
	}

	buffer.Write(body)

	buffer.Write(r.body.Bytes())
	buffer.WriteString(template.RunEnd)

	return buffer.Bytes(), nil
}
