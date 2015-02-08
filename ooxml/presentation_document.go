package ooxml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/t-yuki/gooxml/opc"
)

type PresentationDocument struct {
	Package *opc.Package
	Part    *opc.Part

	XMLName   xml.Name  `xml:"http://schemas.openxmlformats.org/presentationml/2006/main presentation"`
	SlideList []SlideID `xml:"http://schemas.openxmlformats.org/presentationml/2006/main sldIdLst>sldId"`
}

type SlideID struct {
	ID         string `xml:"http://schemas.openxmlformats.org/presentationml/2006/main id,attr"` // TODO: encoding/xml can't catch id attr
	RelationID string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
}

func (d *PresentationDocument) Slides() []*PresentationSlide {
	slides := make([]*PresentationSlide, 0, len(d.SlideList))
	for _, slide := range d.SlideList {
		parts := d.Package.FindPartsByRelationOn(d.Part, func(rel *opc.Relationship) bool { return rel.ID == slide.RelationID })
		for _, part := range parts {
			if part == nil || part.Content == nil {
				panic(fmt.Sprintf("%+v", part))
			}
			slides = append(slides, part.Content.(*PresentationSlide))
		}
	}
	return slides
}

func (d *PresentationDocument) String() string {
	buf := &bytes.Buffer{}
	for _, slide := range d.Slides() {
		buf.WriteString(slide.String())
	}
	return buf.String()
}

func buildPresentationDocument(pkg *opc.Package, partName string, in io.Reader) error {
	part := pkg.FindPart(partName)
	doc := &PresentationDocument{
		Package: pkg,
		Part:    part,
	}
	dec := xml.NewDecoder(in)
	if err := dec.Decode(doc); err != nil {
		return err
	}
	part.Content = doc
	return nil
}
