package section

import (
	"bytes"
	"fmt"
	"github.com/wangcan-null/go-office/docx/template"
)

// Cols 分栏配置结构
// 分栏中所有宽度的单位都是点的二十分之一
type Cols struct {
	// isSet 是否设置分栏
	isSet bool

	// EqualWidth 是否所有分栏等宽
	EqualWidth bool

	// Num 分栏数量
	Num uint

	// Sep 是否显示分栏间的分割线
	Sep bool

	// 分栏间距
	Space uint

	// cols 独立的分栏配置
	// 此属性不为nil时将忽略 EqualWidth 与 Num 两个属性
	cols []*Col
}

// SetNum 设置分栏数量，该配置默认所有分栏等宽
// 如果需要设置不等宽的分栏，需要通过 AddCol 进行独立的分栏配置
func (c *Cols) SetNum(num uint) *Cols {
	c.isSet = true

	c.EqualWidth = true
	c.Num = num

	return c
}

// SetSep 设置分栏间的分割线
func (c *Cols) SetSep(show bool) *Cols {
	c.isSet = true
	c.Sep = show

	return c
}

// SetSpace 设置分栏间距
func (c *Cols) SetSpace(space uint) *Cols {
	c.isSet = true
	c.Space = space

	return c
}

// AddCol 添加一个独立的分栏
func (c *Cols) AddCol(width int, space int) *Cols {
	c.isSet = true

	if nil == c.cols {
		c.cols = make([]*Col, 1)

		c.EqualWidth = false
		c.Num = 0
	}

	col := new(Col)
	col.SetWidth(width).SetSpace(space)

	c.cols = append(c.cols, col)
	c.Num += 1

	return c
}

func (c *Cols) GetXmlBytes() ([]byte, error) {
	if !c.isSet {
		return []byte{}, nil
	}

	buffer := new(bytes.Buffer)

	// 如果没有单独设置分栏，那么返回的是一个单标签
	if nil == c.cols {
		buffer.WriteByte('<')
		buffer.WriteString(template.SectionColsTag)

		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsNum, c.Num))
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsEqualWidth, c.EqualWidth))

		if c.Sep {
			buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsSep, c.Sep))
		}

		if 0 < c.Space {
			buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsSpace, c.Space))
		}

		buffer.WriteString("/>")

		return buffer.Bytes(), nil
	}

	buffer.WriteByte('<')
	buffer.WriteString(template.SectionColsTag)

	// Microsoft Word 2007在分栏宽度不相等时也要求设置num属性
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsNum, c.Num))

	if c.Sep {
		buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColsSep, c.Sep))
	}
	buffer.WriteByte('>')

	// 输出所有分栏的内容
	for _, col := range c.cols {
		body, err := col.GetXmlBytes()
		if nil != err {
			return nil, err
		}

		buffer.Write(body)
	}

	buffer.WriteString("</")
	buffer.WriteString(template.SectionColsTag)
	buffer.WriteByte('>')

	return buffer.Bytes(), nil
}

// Col 独立的分栏配置
type Col struct {
	// Width 分栏宽度
	Width int

	// Space 分栏间距
	Space int
}

// SetWidth 设置分栏宽度
func (c *Col) SetWidth(width int) *Col {
	c.Width = width

	return c
}

// SetSpace 设置分栏间距
func (c *Col) SetSpace(space int) *Col {
	c.Space = space

	return c
}

func (c *Col) GetXmlBytes() ([]byte, error) {
	buffer := new(bytes.Buffer)

	buffer.WriteByte('<')
	buffer.WriteString(template.SectionColTag)

	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColWidth, c.Width))
	buffer.WriteString(fmt.Sprintf(` %v="%v"`, template.SectionColSpace, c.Space))

	buffer.WriteString("/>")

	return buffer.Bytes(), nil
}
