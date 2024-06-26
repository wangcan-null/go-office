package table

import (
	"errors"
	"fmt"
	"github.com/wangcan-null/go-office/docx/table/row"
	"sync"
)

// Row 表格行结构定义
type Row struct {
	// pr 表格行的属性指针
	pr *row.TrPr

	cm sync.RWMutex
	// 单元格列表
	cells []*Cell
}

// GetProperties 获取属性配置结构
func (r *Row) GetProperties() *row.TrPr {
	if nil == r.pr {
		r.pr = new(row.TrPr)
	}

	return r.pr
}

// GetCells 获取全部单元格
func (r *Row) GetCells() []*Cell {
	r.cm.RLock()
	r.cm.RUnlock()

	return r.cells
}

// GetCell 获取指定位置的单元格
func (r *Row) GetCell(index uint) (*Cell, error) {
	if index >= uint(len(r.cells)) {
		return nil, errors.New(fmt.Sprintf("索引溢出"))
	}

	r.cm.RLock()
	defer r.cm.RUnlock()

	return r.cells[index], nil
}

// AddCellText 添加指定数量单元格的文本内容
func (r *Row) AddCellText(cells ...interface{}) error {
	r.cm.RLock()
	defer r.cm.RUnlock()

	if len(cells) > len(r.cells) {
		return errors.New(fmt.Sprintf("索引溢出: 单元格数量为 %v 当前设置数量为 %v", len(r.cells), len(cells)))
	}

	for i, text := range cells {
		cell := r.cells[i]
		cell.AddParagraph().AddRun().AddText(text)
	}
	return nil
}

// addCell 添加一个自动宽度的单元格
func (r *Row) addCell() {
	cell := new(Cell)

	r.cm.Lock()
	defer r.cm.Unlock()

	r.cells = append(r.cells, cell)
}

// addCellWithWidth 添加一个单元格并指定宽度
func (r *Row) addCellWithWidth(width int) {
	cell := new(Cell)
	cell.GetProperties().SetWidth(width)

	r.cm.Lock()
	defer r.cm.Unlock()

	r.cells = append(r.cells, cell)
}

// addCellWithIndexAndWidth 添加一列到指定位置以及宽度
func (r *Row) addCellWithIndexAndWidth(index int, width int) {
	cell := new(Cell)
	cell.GetProperties().SetWidth(width)

	r.cm.Lock()
	defer r.cm.Unlock()

	if index >= len(r.cells) {
		r.cells = append(r.cells, cell)

		return
	}

	before := r.cells[:index]
	after := r.cells[index:]

	i := 0
	r.cells = make([]*Cell, len(r.cells)+1)
	for _, c := range before {
		r.cells[i] = c
		i += 1
	}

	r.cells[i] = cell
	i += 1

	for _, c := range after {
		r.cells[i] = c
		i += 1
	}
}
