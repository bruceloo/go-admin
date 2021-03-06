package types

import (
	"html/template"
	"github.com/chenhg5/go-admin/modules/menu"
)

type FormAttribute interface {
	SetContent(value []FormStruct) FormAttribute
	SetPrefix(value string) FormAttribute
	SetUrl(value string) FormAttribute
	SetInfoUrl(value string) FormAttribute
	SetMethod(value string) FormAttribute
	SetTitle(value string) FormAttribute
	SetToken(value string) FormAttribute
	GetContent() template.HTML
}

type BoxAttribute interface {
	SetHeader(value template.HTML) BoxAttribute
	SetBody(value template.HTML) BoxAttribute
	SetFooter(value template.HTML) BoxAttribute
	SetTitle(value template.HTML) BoxAttribute
	WithHeadBorder(has bool) BoxAttribute
	GetContent() template.HTML
}

type ColAttribute interface {
	SetWidth(value string) ColAttribute
	SetContent(value template.HTML) ColAttribute
	SetType(value string) ColAttribute
	GetContent() template.HTML
}

type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	SetSrc(value string) ImgAttribute
	GetContent() template.HTML
}

type InfoBoxAttribute interface {
	SetTitle(value string) InfoBoxAttribute
	SetValue(value string) InfoBoxAttribute
	SetUrl(value string) InfoBoxAttribute
	GetContent() template.HTML
}

type LabelAttribute interface {
	SetContent(value string) LabelAttribute
	GetContent() template.HTML
}

type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute
	GetContent() template.HTML
}

type TableAttribute interface {
	SetThead(value []map[string]string) TableAttribute
	SetInfoList(value []map[string]template.HTML) TableAttribute
	SetType(value string) TableAttribute
	GetContent() template.HTML
}

type DataTableAttribute interface {
	GetDataTableHeader() template.HTML
	SetThead(value []map[string]string) DataTableAttribute
	SetInfoList(value []map[string]template.HTML) DataTableAttribute
	SetEditUrl(value string) DataTableAttribute
	SetNewUrl(value string) DataTableAttribute
	GetContent() template.HTML
}

type TreeAttribute interface {
	SetTree(value []menu.MenuItem) TreeAttribute
	GetContent() template.HTML
	GetTreeHeader() template.HTML
}

type PaninatorAttribute interface {
	SetCurPageStartIndex(value string) PaninatorAttribute
	SetCurPageEndIndex(value string) PaninatorAttribute
	SetTotal(value string) PaninatorAttribute
	SetPreviousClass(value string) PaninatorAttribute
	SetPreviousUrl(value string) PaninatorAttribute
	SetPages(value []map[string]string) PaninatorAttribute
	SetNextClass(value string) PaninatorAttribute
	SetNextUrl(value string) PaninatorAttribute
	SetOption(value map[string]template.HTML) PaninatorAttribute
	SetUrl(value string) PaninatorAttribute
	GetContent() template.HTML
}
