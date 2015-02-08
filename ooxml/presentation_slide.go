package ooxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/t-yuki/gooxml/opc"
)

type PresentationSlide struct {
	XMLName   xml.Name `xml:"sld"`
	SlideData SlideData
}

type SlideData struct {
	XMLName   xml.Name   `xml:"cSld"`
	ShapeTree GroupShape `xml:"spTree"`
}

type GroupShape struct {
	XMLName xml.Name
	Shapes  []AnyShape `xml:",any"`
}

func (g *GroupShape) String() string {
	buf := &bytes.Buffer{}
	for _, sp := range g.Shapes {
		buf.WriteString(sp.String())
	}
	return buf.String()
}

type AnyShape struct {
	XMLName xml.Name
	Shape   Shape        `xml:"sp"`
	Group   GroupShape   `xml:"grpSp"`
	Graphic GraphicFrame `xml:"graphicFrame"`
	fmt.Stringer
}

func (a *AnyShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.XMLName = start.Name
	switch start.Name.Local {
	case "sp":
		a.Stringer = &a.Shape
		return d.DecodeElement(&a.Shape, &start)
	case "grpSp":
		a.Stringer = &a.Group
		return d.DecodeElement(&a.Group, &start)
	case "graphicFrame":
		a.Stringer = &a.Graphic
		return d.DecodeElement(&a.Graphic, &start)
	}
	a.Stringer = &bytes.Buffer{}                // dummy stringer
	return d.DecodeElement(&struct{}{}, &start) // unknown, ignore it
}

type GraphicFrame struct {
	Graphic GraphicalObject `xml:"graphic>graphicData"`
}

func (g *GraphicFrame) String() string {
	return fmt.Sprintf("%v", g.Graphic.Data)
}

type GraphicalObject struct {
	XMLName xml.Name
	URI     string `xml:"uri,attr"`
	Data    interface{}
}

func (g *GraphicalObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	g.XMLName = start.Name
	var data struct {
		Table Table `xml:"tbl"`
	} // TODO switch using URI
	g.Data = &data.Table
	return d.DecodeElement(&data, &start)
}

type Table struct {
	XMLName xml.Name
	Rows    []TableRow `xml:"tr"`
}

func (t *Table) String() string {
	buf := &bytes.Buffer{}
	for _, row := range t.Rows {
		for _, cell := range row.Cells {
			buf.WriteString("| ")
			buf.WriteString(strings.TrimSpace(cell.TextBody.String()))
			buf.WriteString(" ")
		}
		buf.WriteString("|")
		buf.WriteString("\n")
	}
	return buf.String()
}

type TableRow struct {
	XMLName xml.Name
	Cells   []TableCell `xml:"tc"`
}

type TableCell struct {
	XMLName  xml.Name
	TextBody TextBody `xml:"txBody"`
}

type Shape struct {
	XMLName             xml.Name `xml:"sp"`
	NonVisualProperties NonVisualDrawingShapeProps
	TextBody            TextBody `xml:"txBody"`
}

func (s *Shape) String() string {
	return s.TextBody.String()
}

type NonVisualDrawingShapeProps struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main nvSpPr"`
	Application ApplicationNonVisualDrawingProps
}

type ApplicationNonVisualDrawingProps struct {
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main nvPr"`
	PlaceHolder PlaceHolder
}

type PlaceHolder struct {
	XMLName xml.Name        `xml:"http://schemas.openxmlformats.org/presentationml/2006/main ph"`
	Type    PlaceHolderType `xml:"type,attr"`
	Idx     uint            `xml:"idx,attr"`
}

type PlaceHolderType string

const (
	PlaceHolderTypeTitle         = "title"
	PlaceHolderTypeBody          = "body"
	PlaceHolderTypeCenteredTitle = "ctrTitle"
	PlaceHolderTypeSubTitle      = "subTitle"
	PlaceHolderTypeDateTime      = "dt"
	PlaceHolderTypeSlideNo       = "sldNum"
)

func (p *PresentationSlide) Title() string {
	var title, cTitle, subTitle string
	for _, sp := range p.SlideData.ShapeTree.Shapes {
		if sp.Shape.XMLName.Local == "" {
			continue
		}
		phType := sp.Shape.NonVisualProperties.Application.PlaceHolder.Type
		text := sp.Shape.TextBody.String()
		switch phType {
		case PlaceHolderTypeTitle:
			title = text
		case PlaceHolderTypeCenteredTitle:
			cTitle = text
		case PlaceHolderTypeSubTitle:
			subTitle = text
		}
	}
	switch {
	case title != "":
		return title
	case cTitle != "":
		return cTitle
	case subTitle != "":
		return subTitle
	}
	return "[NO TITLE]"
}

func (p *PresentationSlide) String() string {
	// filter out title shape or empty shape
	array := make([]string, 0, len(p.SlideData.ShapeTree.Shapes))
	for _, sp := range p.SlideData.ShapeTree.Shapes {
		if sp.Shape.XMLName.Local != "" {
			phType := sp.Shape.NonVisualProperties.Application.PlaceHolder.Type
			if phType != "" && phType != PlaceHolderTypeBody {
				continue
			}
		}
		txt := sp.String()
		if strings.TrimSpace(txt) == "" {
			continue
		}
		array = append(array, txt)
	}

	buf := &bytes.Buffer{}
	buf.WriteString(strings.TrimSpace(p.Title()))
	buf.WriteString("\n-----\n")
	for i, txt := range array {
		switch i {
		case 0:
			txt = strings.TrimLeftFunc(txt, unicode.IsSpace)
		case len(array) - 1:
			txt = strings.TrimRightFunc(txt, unicode.IsSpace)
		}
		buf.WriteString(txt)
		buf.WriteString("\n")
	}
	buf.WriteString("\n")
	return buf.String()
}

func buildPresentationSlide(pkg *opc.Package, partName string, in io.Reader) error {
	slide := &PresentationSlide{}
	dec := xml.NewDecoder(in)
	if err := dec.Decode(slide); err != nil {
		return err
	}
	part := pkg.FindPart(partName)
	if part == nil {
		panic(partName)
	}
	part.Content = slide
	return nil
}
