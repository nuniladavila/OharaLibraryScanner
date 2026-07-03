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

type RichTextFilter struct {
	Equals string `json:"equals"`
}

type TitleFilter struct {
	Equals string `json:"equals"`
}

type FilterCondition struct {
	Title    *TitleFilter    `json:"title,omitempty"`
	Property string          `json:"property,omitempty"`
	Type     string          `json:"type,omitempty"`
	RichText *RichTextFilter `json:"rich_text,omitempty"`
}

type Filter struct {
	Or       []FilterCondition `json:"or,omitempty"`
	RichText *RichTextFilter   `json:"rich_text,omitempty"`
	Property string            `json:"property,omitempty"`
}

type Sort struct {
	Property string `json:"property"`
}

type QueryPayload struct {
	Sorts       []Sort `json:"sorts,omitempty"`
	Filter      Filter `json:"filter,omitempty"`
	StartCursor string `json:"start_cursor,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
	IsArchived  bool   `json:"is_archived,omitempty"`
	ResultType  string `json:"result_type,omitempty"`
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

type NotionQueryResponse struct {
	Object           string                 `json:"object"`
	Results          []NotionPage           `json:"results"`
	NextCursor       interface{}            `json:"next_cursor"`
	HasMore          bool                   `json:"has_more"`
	Type             string                 `json:"type"`
	PageOrDataSource map[string]interface{} `json:"page_or_data_source"`
	RequestID        string                 `json:"request_id"`
}

type NotionUser struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type NotionPage struct {
	Object         string               `json:"object"`
	ID             string               `json:"id"`
	CreatedTime    string               `json:"created_time"`
	LastEditedTime string               `json:"last_edited_time"`
	CreatedBy      NotionUser           `json:"created_by"`
	LastEditedBy   NotionUser           `json:"last_edited_by"`
	Cover          interface{}          `json:"cover"`
	Icon           interface{}          `json:"icon"`
	Parent         NotionParent         `json:"parent"`
	InTrash        bool                 `json:"in_trash"`
	IsArchived     bool                 `json:"is_archived"`
	IsLocked       bool                 `json:"is_locked"`
	Properties     NotionPageProperties `json:"properties"`
	URL            string               `json:"url"`
	PublicURL      string               `json:"public_url"`
}

type NotionParent struct {
	Type         string `json:"type"`
	DataSourceID string `json:"data_source_id"`
	DatabaseID   string `json:"database_id"`
}

type NotionPageProperties struct {
	ISBN  NotionRichTextProperty `json:"ISBN"`
	Title NotionTitleProperty    `json:"Title"`
}

type NotionRichTextProperty struct {
	ID       string               `json:"id"`
	Type     string               `json:"type"`
	RichText []NotionRichTextItem `json:"rich_text"`
}

type NotionTitleProperty struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Title []NotionTitleItem `json:"title"`
}

type NotionRichTextItem struct {
	Type        string                `json:"type"`
	Text        NotionTextContent     `json:"text"`
	Annotations NotionTextAnnotations `json:"annotations"`
	PlainText   string                `json:"plain_text"`
	Href        interface{}           `json:"href"`
}

type NotionTitleItem struct {
	Type        string                `json:"type"`
	Text        NotionTextContent     `json:"text"`
	Annotations NotionTextAnnotations `json:"annotations"`
	PlainText   string                `json:"plain_text"`
	Href        interface{}           `json:"href"`
}

type NotionTextContent struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}

type NotionTextAnnotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}
