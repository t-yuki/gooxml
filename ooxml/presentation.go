package ooxml

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"github.com/t-yuki/gooxml/opc"
)

const (
	ContentTypePresentationDocument    = "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"
	ContentTypePresentationSlide       = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	ContentTypePresentationSlideMaster = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
)

func init() {
	opc.RegisterReadFormat(ContentTypePresentationDocument, buildPresentationDocument)
	opc.RegisterReadFormat(ContentTypePresentationSlide, buildPresentationSlide)
}

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

func ReadPresentationDocument(in io.Reader) (*PresentationDocument, error) {
	pkg, err := opc.Read(in)
	// ignore err if pkg is available
	if err != nil && pkg == nil {
		return nil, err
	}
	return findPresentationDocument(pkg)
}

func OpenPresentationDocument(name string) (*PresentationDocument, error) {
	pkg, err := opc.Open(name)
	if err != nil && pkg == nil {
		return nil, err
	}
	// ignore err if pkg is available
	return findPresentationDocument(pkg)
}

func findPresentationDocument(pkg *opc.Package) (*PresentationDocument, error) {
	parts := pkg.FindPartsByRelationOn(&pkg.Part, func(rel *opc.Relationship) bool { return rel.Type == RelationTypeOfficeDocument })
	if len(parts) != 1 || parts[0].Content == nil {
		return nil, errors.New("it is not a PresentationDocument")
	}
	return parts[0].Content.(*PresentationDocument), nil
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
