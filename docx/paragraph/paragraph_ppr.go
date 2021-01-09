package paragraph

import (
    "bytes"
    "github.com/Luna-CY/go-office/docx/template"
    "strings"
)

// PPr 段落的样式属性定义
type PPr struct {

    // horizontalAlignment 水平对齐方式
    horizontalAlignment *HorizontalAlignment

    // borderManager 边框管理器
    borderManager *BorderManager

    // keepLines 该段落是否尽可能的在一个页面上显示
    keepLines bool

    // keepNext 该段落与下一个段落是否尽可能的在一个页面上显示
    keepNext bool

    // identity 缩进配置结构
    identity *Identity

    // background 背景配置结构
    background *Background
}

// GetBorderManager 获取边框管理器
func (p *PPr) GetBorderManager() *BorderManager {
    if nil == p.borderManager {
        p.borderManager = new(BorderManager)
        p.borderManager.isSet = false
    }

    return p.borderManager
}

// GetIdentity 获取缩进配置结构指针
func (p *PPr) GetIdentity() *Identity {
    if nil == p.identity {
        p.identity = new(Identity)
        p.identity.isSet = false
    }

    return p.identity
}

// GetBackground 获取背景配置结构指针
func (p *PPr) GetBackground() *Background {
    if nil == p.background {
        p.background = new(Background)
        p.background.isSet = false
    }

    return p.background
}

// SetKeepLines
func (p *PPr) SetKeepLines(keepLines bool) *PPr {
    p.keepLines = keepLines

    return p
}

// SetKeepNext
func (p *PPr) SetKeepNext(keepNext bool) *PPr {
    p.keepNext = keepNext

    return p
}

func (p *PPr) GetBody() ([]byte, error) {
    buffer := new(bytes.Buffer)

    if nil != p.horizontalAlignment {
        buffer.WriteString(strings.Replace(template.ParagraphPPrHorizontalAlignment, "{{TYPE}}", string(*p.horizontalAlignment), 1))
    }

    if nil != p.borderManager {
        body, err := p.borderManager.GetBody()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.identity {
        body, err := p.identity.GetBody()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if nil != p.background {
        body, err := p.background.GetBody()
        if nil != err {
            return nil, nil
        }

        buffer.Write(body)
    }

    if p.keepLines {
        buffer.WriteString(template.ParagraphPPrKeepLines)
    }

    if p.keepNext {
        buffer.WriteString(template.ParagraphPPrKeepNext)
    }

    return buffer.Bytes(), nil
}

// SetHorizontalAlignment 设置水平对齐方式
func (p *PPr) SetHorizontalAlignment(alignment HorizontalAlignment) *PPr {
    p.horizontalAlignment = &alignment

    return p
}

type HorizontalAlignment string

const (
    // HorizontalAlignmentStart 左对齐
    HorizontalAlignmentStart = HorizontalAlignment("start")

    // HorizontalAlignmentEnd 右对齐
    HorizontalAlignmentEnd = HorizontalAlignment("end")

    // HorizontalAlignmentCenter 居中对齐
    HorizontalAlignmentCenter = HorizontalAlignment("center")

    // HorizontalAlignmentBoth 左右对齐，不改变字符间距
    HorizontalAlignmentBoth = HorizontalAlignment("both")

    // HorizontalAlignmentDistribute 左右对齐，改变字符间距
    HorizontalAlignmentDistribute = HorizontalAlignment("distribute")
)
