package models

// Typed payload structures for Notion
type TextContent struct {
	Content string `json:"content"`
}

type TitleText struct {
	Text TextContent `json:"text"`
}

type TitleProp struct {
	Title []TitleText `json:"title"`
}

type SelectItem struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type SelectProp struct {
	Select SelectItem `json:"select"`
}

type MultiSelectProp struct {
	MultiSelect []SelectItem `json:"multi_select"`
}

type DateStart struct {
	Start string `json:"start"`
}

type DateProp struct {
	Date DateStart `json:"date"`
}

type RichTextItem struct {
	Text TextContent `json:"text"`
}

type RichTextProp struct {
	RichText []RichTextItem `json:"rich_text"`
}

type CheckboxProp struct {
	Checkbox bool `json:"checkbox"`
}

type NumberProp struct {
	Number float64 `json:"number"`
}

type FileExternal struct {
	URL string `json:"url"`
}

type FileItem struct {
	External FileExternal `json:"external"`
	Name     string       `json:"name"`
}

type FilesProp struct {
	Files []FileItem `json:"files"`
}

type Properties struct {
	Title         TitleProp       `json:"Title"`
	Author        SelectProp      `json:"Author"`
	SubCategory   MultiSelectProp `json:"SubCategory"`
	PublishedDate DateProp        `json:"Published Date"`
	ISBN          RichTextProp    `json:"ISBN"`
	Read          CheckboxProp    `json:"Read"`
	BookCover     FilesProp       `json:"Book Cover"`
	PageCount     NumberProp      `json:"Page Count"`
	Editor        SelectProp      `json:"Editor"`
	ShelfLocation SelectProp      `json:"Shelf Location"`
	Category      SelectProp      `json:"Category"`
	Publisher     SelectProp      `json:"Publisher"`
	Language      SelectProp      `json:"Language"`
	DateAdded     DateProp        `json:"Date Added"`
	DateAcquired  DateProp        `json:"Date Acquired"`
	Edition       RichTextProp    `json:"Edition"`
}

type ParentProp struct {
	DataSourceID string `json:"data_source_id"`
}

type Payload struct {
	Parent     ParentProp `json:"parent"`
	Properties Properties `json:"properties"`
}
