package containers

import (
	"md2docx/bolditalic"
	"md2docx/code"
	"md2docx/headers"
	"md2docx/image"
	"md2docx/link"
	"md2docx/list"
	"md2docx/misc"
	"md2docx/parse"
	"md2docx/patterns"
	"md2docx/table"
	"regexp"

	"github.com/unidoc/unioffice/document"
)

type ParserFunction = func(pattern *regexp.Regexp, text string) (x int)
type AdderFunction = func(text string, para document.Paragraph, doc *document.Document, level int) int
type MDTypeName string

type MDType struct {
	Name      MDTypeName
	Pattern   string
	ParseFunc ParserFunction
	AddFunc   AdderFunction
}

const (
	Header     MDTypeName = "Header"
	Bold       MDTypeName = "Bold"
	Italic     MDTypeName = "Italic"
	BoldItalic MDTypeName = "BoldItalic"
	Link       MDTypeName = "Link"
	ImageLink  MDTypeName = "ImageLink"
	EmailURL   MDTypeName = "EmailUrl"
	BlockQ     MDTypeName = "BlockQ"
	List       MDTypeName = "List"
	Table      MDTypeName = "Table"
	HorizL     MDTypeName = "HorizL"
	CodeBlock  MDTypeName = "CodeBlock"
	CodeInline MDTypeName = "CodeInline"
	Image      MDTypeName = "Image"
)

var (
	HeaderType     = MDType{Header, patterns.HeaderAll, parse.ParseHeader, headers.AddHeader}
	BoldType       = MDType{Bold, patterns.BoldItalicText, parse.ParseBold, bolditalic.AddBoldText}
	ItalicType     = MDType{Italic, patterns.BoldItalicText, parse.ParseItalic, bolditalic.AddItalic}
	BoldItalicType = MDType{BoldItalic, patterns.BoldItalicText, parse.ParseBoldItalic, bolditalic.AddBoldItalic}
	LinkType       = MDType{Link, patterns.LinkText, parse.ParseNormal, link.AddLink}
	EmailUrlType   = MDType{EmailURL, patterns.EmailUrlText, parse.ParseNormal, link.AddEmailUrl}
	ListTextType   = MDType{List, patterns.ListText, parse.ParseNormal, list.AddList}
	BQType         = MDType{BlockQ, patterns.BlockQuote, parse.ParseNormal, misc.AddBlockQuote}
	HLType         = MDType{HorizL, patterns.HorizontalLine, parse.ParseNormal, misc.AddHorizontalLine}
	TableType      = MDType{Table, patterns.TableText, parse.ParseNormal, table.AddTable}
	CodeBlockType  = MDType{CodeBlock, patterns.CodeBlock, parse.ParseNormal, code.AddCodeBlock}
	CodeInlineType = MDType{CodeInline, patterns.InlineCode, parse.ParseNormal, code.AddInlineCode}
	ImageTypee     = MDType{Image, patterns.ImageFile, parse.ParseNormal, image.AddImage}
)
