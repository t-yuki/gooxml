package ooxml

import (
	"bytes"
	"encoding/xml"
)

type TextBody struct {
	Properties TextBodyProperties
	Paragraphs []TextParagraph `xml:"http://schemas.openxmlformats.org/drawingml/2006/main p"`
}

func (t *TextBody) String() string {
	buf := &bytes.Buffer{}
	for _, p := range t.Paragraphs {
		for _, r := range p.Run {
			buf.WriteString(r.String())
		}
		if p.EndProperties.XMLName.Local != "" {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

type TextParagraph struct {
	XMLName       xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main p"`
	Properties    TextParagraphProperties
	Run           []TextRunGroup          `xml:",any"`
	EndProperties TextCharacterProperties `xml:"endParaRPr"`
}

type TextRunGroup struct {
	XMLName xml.Name
	Run     TextRun
	Break   TextLineBreak
	Field   TextField
}

func (t *TextRunGroup) String() string {
	switch t.XMLName.Local {
	case "r":
		return t.Run.Text.Data
	case "br":
		return "\n"
	case "fld":
		return t.Field.Text.Data
	}
	return ""
}

func (t *TextRunGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	t.XMLName = start.Name
	var v interface{}
	switch start.Name.Local {
	case "r":
		v = &t.Run
	case "br":
		v = &t.Break
	case "fld":
		v = &t.Field
	default:
		v = &struct{}{}
	}
	return d.DecodeElement(v, &start)
}

type TextLineBreak struct {
	XMLName    xml.Name                `xml:"http://schemas.openxmlformats.org/drawingml/2006/main br"`
	Properties TextCharacterProperties `xml:"rPr"`
}

type TextField struct {
	XMLName             xml.Name                `xml:"http://schemas.openxmlformats.org/drawingml/2006/main fld"`
	Properties          TextCharacterProperties `xml:"rPr"`
	ParagraphProperties TextParagraphProperties
	Text                Text
}

type TextRun struct {
	XMLName    xml.Name                `xml:"http://schemas.openxmlformats.org/drawingml/2006/main r"`
	Properties TextCharacterProperties `xml:"rPr"`
	Text       Text
}

type Text struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main t"`
	Data    string   `xml:",chardata"`
}

type TextBodyProperties struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main bodyPr"`
}

type TextParagraphProperties struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/drawingml/2006/main pPr"`
}

type TextCharacterProperties struct {
	XMLName xml.Name
	Lang    string `xml:lang,attr`
	Dirty   int    `xml:dirty,attr`
}
