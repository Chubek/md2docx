package containers

import (
	"md2docx/bolditalic"
	"md2docx/headers"
	"md2docx/parse"
	"md2docx/link"
	"md2docx/list"
	"md2docx/patterns"
	"md2docx/misc"
	"md2docx/table"
	"md2docx/code"
	"regexp"
	"github.com/unidoc/unioffice/document"
)


type ParserFunction = func(pattern *regexp.Regexp, text string) (x int) 
type AdderFunction = func(text string, para document.Paragraph, doc *document.Document, level int) int

type MDType struct {
	Pattern			string
	ParseFunc		ParserFunction
	AddFunc			AdderFunction
}


var (
	HeaderType 		= MDType{patterns.HeaderAll, parse.ParseHeader, headers.AddHeader}
	BoldType 		= MDType{patterns.BoldItalicText, parse.ParseBold, bolditalic.AddBoldText}
	ItalicType 		= MDType{patterns.BoldItalicText, parse.ParseItalic, bolditalic.AddItalic}
	BoldItalicType 	= MDType{patterns.BoldItalicText, parse.ParseBoldItalic, bolditalic.AddBoldItalic}
	LinkType 		= MDType{patterns.LinkText, parse.ParseNormal, link.AddLink}
	EmailUrlType 	= MDType{patterns.EmailUrlText, parse.ParseNormal, link.AddEmailUrl}
	ListTextType	= MDType{patterns.ListText, parse.ParseNormal, list.AddList}
	BQType			= MDType{patterns.BlockQuote, parse.ParseNormal, misc.AddBlockQuote}
	HLType			= MDType{patterns.HorizontalLine, parse.ParseNormal, misc.AddHorizontalLine}
	TableType		= MDType{patterns.TableText, parse.ParseNormal, table.AddTable}
	CodeBlockType	= MDType{patterns.CodeBlock, parse.ParseNormal, code.AddCodeBlock}
	CodeInlineType	= MDType{patterns.InlineCode, parse.ParseNormal, code.AddInlineCode}

)