package ooxml

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"github.com/t-yuki/gooxml/opc"
)

type PresentationSlide struct {
	XMLName   xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sld"`
	SlideData SlideData
}

type SlideData struct {
	XMLName   xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main cSld"`
	ShapeTree ShapeTree
}

type ShapeTree struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main spTree"`
	Shapes  []Shape  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sp"`
}

type Shape struct {
	XMLName             xml.Name `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sp"`
	NonVisualProperties NonVisualDrawingShapeProps
	TextBody            TextBody `xml:"http://schemas.openxmlformats.org/presentationml/2006/main txBody"`
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
		phType := sp.NonVisualProperties.Application.PlaceHolder.Type
		text := sp.TextBody.String()
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
	buf := &bytes.Buffer{}
	buf.WriteString(strings.TrimSpace(p.Title()))
	buf.WriteString("\n-----\n")
	for _, sp := range p.SlideData.ShapeTree.Shapes {
		phType := sp.NonVisualProperties.Application.PlaceHolder.Type
		text := sp.TextBody.String()
		if phType == "" || phType == PlaceHolderTypeBody {
			buf.WriteString(text)
			buf.WriteString("\n")
		}
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
