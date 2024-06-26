# go-office

golang的开源办公套件生成器（非编辑器），支持列表如下：

- Word(段落:o:、表格:o:、页头页脚:o:、基础样式:o:、公式:x:、图表:x:)
- Excel(基本数据设置:o:、样式:x:、公式:x:、图表:x:)
- PowerPoint(:x:)

## 安装

`go get -u github.com/wangcan-null/go-office`

## 使用

可以在`example`目录下查看更多示例

```go
import (
    "github.com/wangcan-null/go-office/docx"
    "github.com/wangcan-null/go-office/docx/paragraph"
)

// 创建一个新的文档
doc := docx.New()

// 设置文档属性
doc.GetSection().GetPageSize().SetWidth(10000).SetHeight(200000)

// 设置默认段落属性
doc.GetProperties().GetDefaultParagraphProperties().GetSpacing().SetSpace(360)

// 添加一个段落
p1 := doc.AddParagraph()

// 设置段落属性
p1.GetProperties().SetHorizontalAlignment(paragraph.HorizontalAlignmentCenter)

// 添加一个文本段
r1 := p1.AddRun()

// 设置文本段属性
r1.GetProperties().GetBackground().SetBackgroundColor("FF0000")

// 在文本段内添加文本
r1.AddText("hello")
r1.AddText(" ")
r1.AddText("word")

// 保存文档
doc.Save("example.docx")
```
