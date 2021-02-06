package xlsx

import (
    "bytes"
    "github.com/Luna-CY/go-office/xlsx/template"
    "strconv"
)

// App app.xml
type App struct {
    // Application 生产者应用的名称
    Application string

    // Security 加密级别
    Security int
}

func (a *App) GetXmlBytes() ([]byte, error) {
    buffer := new(bytes.Buffer)

    buffer.WriteString(template.Xml)
    buffer.WriteString(template.AppXmlStart)

    buffer.WriteString(template.AppApplicationStart)
    buffer.WriteString(a.Application)
    buffer.WriteString(template.AppApplicationEnd)

    buffer.WriteString(template.AppSecurityStart)
    buffer.WriteString(strconv.Itoa(a.Security))
    buffer.WriteString(template.AppSecurityEnd)

    buffer.WriteString(template.AppXmlEnd)

    return buffer.Bytes(), nil
}
