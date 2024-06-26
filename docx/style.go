package docx

import (
	"bytes"
	"fmt"
	"github.com/wangcan-null/go-office/docx/paragraph"
	"github.com/wangcan-null/go-office/docx/run"
	"github.com/wangcan-null/go-office/docx/table"
	"github.com/wangcan-null/go-office/docx/template"
	"sync"
)

// Styles 样式配置结构
type Styles struct {
	// pPrDefault 段落的默认样式
	pPrDefault *paragraph.PPr

	// rPrDefault 文本的默认样式
	rPrDefault *run.RPr

	sm        sync.RWMutex
	styleList []*Style
}

func (s *Styles) GetDefaultParagraphProperties() *paragraph.PPr {
	if nil == s.pPrDefault {
		s.pPrDefault = new(paragraph.PPr)
	}

	return s.pPrDefault
}

func (s *Styles) GetDefaultRunProperties() *run.RPr {
	if nil == s.rPrDefault {
		s.rPrDefault = new(run.RPr)
	}

	return s.rPrDefault
}

// addParagraphStyle 添加一个段落的样式结构
func (s *Styles) addParagraphStyle(styleId string, pPr *paragraph.PPr) {
	style := &Style{styleId: styleId, styleType: StyleTypeParagraph, pPr: pPr}

	s.sm.Lock()
	defer s.sm.Unlock()

	s.styleList = append(s.styleList, style)
}

// addRunStyle 添加一个文本的样式结构
func (s *Styles) addRunStyle(styleId string, rPr *run.RPr) {
	style := &Style{styleId: styleId, styleType: StyleTypeCharacter, rPr: rPr}

	s.sm.Lock()
	defer s.sm.Unlock()

	s.styleList = append(s.styleList, style)
}

// addTableStyle 添加一个表格样式结构
func (s *Styles) addTableStyle(styleId string, tblPr *table.TblPr) {
	style := &Style{styleId: styleId, styleType: StyleTypeTable, tblPr: tblPr}

	s.sm.Lock()
	defer s.sm.Unlock()

	s.styleList = append(s.styleList, style)
}

func (s *Styles) GetXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteString(template.Xml)
	buffer.WriteString(template.StyleXmlStart)

	// 输出全局默认样式docDefaults
	buffer.WriteString(template.StyleDocDefaultStart)

	// 段落默认样式
	dppr, err := s.GetDefaultParagraphProperties().GetDefaultXmlBytes()
	if nil != err {
		return nil, err
	}

	buffer.WriteString(template.StyleDefaultParagraphStart)
	buffer.Write(dppr)
	buffer.WriteString(template.StyleDefaultParagraphEnd)

	// 文本默认样式
	rppr, err := s.GetDefaultRunProperties().GetDefaultXmlBytes()
	if nil != err {
		return nil, err
	}

	buffer.WriteString(template.StyleDefaultRunStart)
	buffer.Write(rppr)
	buffer.WriteString(template.StyleDefaultRunEnd)

	buffer.WriteString(template.StyleDocDefaultEnd)

	for _, style := range s.styleList {
		body, err := style.GetXmlBytes()
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	buffer.WriteString(template.StyleXmlEnd)

	return buffer.Bytes(), nil
}

// Style 样式结构
type Style struct {
	// styleType 样式类型
	styleType StyleType

	// styleId 样式ID
	styleId string

	// 段落样式属性
	pPr *paragraph.PPr

	// 文本样式属性
	rPr *run.RPr

	// 表格样式
	tblPr *table.TblPr
}

// SetStyleType 设置样式类型
func (s *Style) SetStyleType(styleType StyleType) *Style {
	s.styleType = styleType

	return s
}

// SetStyleId 设置样式ID
func (s *Style) SetStyleId(styleId string) *Style {
	s.styleId = styleId

	return s
}

// SetPPr 设置段落样式
func (s *Style) SetPPr(pPr *paragraph.PPr) *Style {
	s.pPr = pPr

	return s
}

//SetRPr 设置文本样式
func (s *Style) SetRPr(rPr *run.RPr) *Style {
	s.rPr = rPr

	return s
}

// SetTblPr 设置表格样式
func (s *Style) SetTblPr(tblPr *table.TblPr) *Style {
	s.tblPr = tblPr

	return s
}

func (s *Style) GetXmlBytes() ([]byte, error) {
	if nil == s.pPr && nil == s.rPr && nil == s.tblPr {
		return []byte{}, nil
	}

	buffer := new(bytes.Buffer)

	buffer.WriteByte('<')
	buffer.WriteString(template.StyleStyleTag)

	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.StyleStyleType, s.styleType))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.StyleStyleStyleId, s.styleId))
	buffer.WriteByte('>')

	isEmpty := true

	if nil != s.pPr {
		buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.StyleStyleNameTag, template.StyleStyleVal, s.styleId))
		body, err := s.pPr.GetExtraXmlBytes()
		if nil != err {
			return nil, err
		}

		if 0 < len(body) {
			isEmpty = false
		}

		buffer.Write(body)
	}

	if nil != s.rPr {
		buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.StyleStyleNameTag, template.StyleStyleVal, s.styleId))
		body, err := s.rPr.GetExtraXmlBytes()
		if nil != err {
			return nil, err
		}

		if 0 < len(body) {
			isEmpty = false
		}

		buffer.Write(body)
	}

	if nil != s.tblPr {
		buffer.WriteString(fmt.Sprintf(`<%v %v="%v"/>`, template.StyleStyleNameTag, template.StyleStyleVal, s.styleId))
		body, err := s.tblPr.GetExtraXmlBytes()
		if nil != err {
			return nil, err
		}

		if 0 < len(body) {
			isEmpty = false
		}

		buffer.Write(body)
	}

	buffer.Write([]byte{'<', '/'})
	buffer.WriteString(template.StyleStyleTag)
	buffer.WriteByte('>')

	if isEmpty {
		return []byte{}, nil
	}

	return buffer.Bytes(), nil
}

type StyleType string

const (
	StyleTypeParagraph = StyleType("paragraph")
	StyleTypeCharacter = StyleType("character")
	StyleTypeTable     = StyleType("table")
)
