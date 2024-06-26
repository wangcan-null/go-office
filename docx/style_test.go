package docx

import (
	"fmt"
	"github.com/wangcan-null/go-office/docx/paragraph"
	"github.com/wangcan-null/go-office/docx/run"
	"github.com/wangcan-null/go-office/docx/table"
	"github.com/wangcan-null/go-office/docx/template"
	"testing"
)

func TestStyleConfig_GetDefaultParagraphProperties(t *testing.T) {
	style := Styles{}

	if nil == style.GetDefaultParagraphProperties() {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v<w:docDefaults><w:pPrDefault><w:pPr></w:pPr></w:pPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	style.GetDefaultParagraphProperties().SetKeepNext(true)
	exp = fmt.Sprintf(`%v%v<w:docDefaults><w:pPrDefault><w:pPr><w:keepNext/></w:pPr></w:pPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyleConfig_GetDefaultRunProperties(t *testing.T) {
	style := Styles{}

	if nil == style.GetDefaultRunProperties() {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v<w:docDefaults><w:rPrDefault><w:rPr></w:rPr></w:rPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	style.GetDefaultRunProperties().SetSize(30)
	exp = fmt.Sprintf(`%v%v<w:docDefaults><w:rPrDefault><w:rPr><w:sz w:val="30"/></w:rPr></w:rPrDefault></w:docDefaults>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyleConfig_GetXmlBytes(t *testing.T) {
	style := Styles{}
	if 0 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyleConfig_addParagraphStyle(t *testing.T) {
	style := Styles{}
	if 0 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	style.addParagraphStyle("TEST_ID", nil)
	if 1 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	ppr := new(paragraph.PPr)
	ppr.GetBackground().SetBackgroundColor("FF0000")
	style.addParagraphStyle("TEST ID 2", ppr)

	exp = fmt.Sprintf(`%v%v<w:style w:type="paragraph" w:styleId="TEST ID 2"><w:name w:val="TEST ID 2"/><w:pPr><w:shd w:fill="FF0000" w:color="auto"/></w:pPr></w:style>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyleConfig_addRunStyle(t *testing.T) {
	style := Styles{}
	if 0 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	style.addRunStyle("TEST_ID", nil)
	if 1 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	rpr := new(run.RPr)
	rpr.GetBackground().SetBackgroundColor("FF0000")

	style.addRunStyle("TEST ID 2", rpr)

	exp = fmt.Sprintf(`%v%v<w:style w:type="character" w:styleId="TEST ID 2"><w:name w:val="TEST ID 2"/><w:rPr><w:shd w:fill="FF0000" w:color="auto"/></w:rPr></w:style>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyleConfig_addTableStyle(t *testing.T) {
	style := Styles{}
	if 0 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	style.addTableStyle("TEST_ID", nil)
	if 1 != len(style.styleList) {
		t.Fatal("验证失败")
	}

	exp := fmt.Sprintf(`%v%v%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}

	tblPr := new(table.TblPr)
	tblPr.GetBorder().SetTop(table.BorderStyleSingle, "FF0000", 30, 10, false)

	style.addTableStyle("TEST ID 2", tblPr)

	exp = fmt.Sprintf(`%v%v<w:style w:type="table" w:styleId="TEST ID 2"><w:name w:val="TEST ID 2"/><w:tblPr><w:tblBorders><w:top w:val="single" w:sz="30" w:color="FF0000" w:space="10" w:shadow="false"/></w:tblBorders></w:tblPr></w:style>%v`, template.Xml, template.StyleXmlStart, template.StyleXmlEnd)
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}

func TestStyle_SetStyleId(t *testing.T) {
	style := Style{}
	style.SetStyleId("TEST ID")

	if "TEST ID" != style.styleId {
		t.Fatal("验证失败")
	}
}

func TestStyle_SetStyleType(t *testing.T) {
	style := Style{}
	style.SetStyleType(StyleTypeParagraph)

	if StyleTypeParagraph != style.styleType {
		t.Fatal("验证失败")
	}
}

func TestStyle_GetXmlBytes(t *testing.T) {
	style := Style{}

	act, err := style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	style.SetStyleId("TEST ID").SetStyleType(StyleTypeTable)

	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if "" != string(act) {
		t.Fatal("验证失败")
	}

	tblPr := new(table.TblPr)
	tblPr.GetBorder().SetTop(table.BorderStyleSingle, "FF0000", 30, 10, false)

	style.SetTblPr(tblPr)

	exp := `<w:style w:type="table" w:styleId="TEST ID"><w:name w:val="TEST ID"/><w:tblPr><w:tblBorders><w:top w:val="single" w:sz="30" w:color="FF0000" w:space="10" w:shadow="false"/></w:tblBorders></w:tblPr></w:style>`
	act, err = style.GetXmlBytes()
	if nil != err {
		t.Fatalf("生成XML失败: %v\n", err)
	}

	if exp != string(act) {
		t.Fatal("验证失败")
	}
}
