package ooxml

import (
	"bytes"
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
}

func OpenPresentationDocument(name string) (*PresentationDocument, error) {
	pkg, err := opc.Open(name)
	if err != nil && pkg == nil {
		return nil, err
	}
	// ignore err if pkg is available

	parts := pkg.FindPartsByRelationOn(&pkg.Part, RelationTypeOfficeDocument)
	if len(parts) != 1 || parts[0].Content == nil {
		return nil, errors.New("it is not a PresentationDocument")
	}
	return parts[0].Content.(*PresentationDocument), nil
}

func (d *PresentationDocument) Slides() []*PresentationSlide {
	// TODO: use p:sldIdLst to order properly
	parts := d.Package.FindPartsByRelationOn(d.Part, RelationTypeSlide)
	slides := make([]*PresentationSlide, 0, len(parts))
	for _, part := range parts {
		if part == nil || part.Content == nil {
			panic(fmt.Sprintf("%+v", part))
		}
		slides = append(slides, part.Content.(*PresentationSlide))
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
	part.Content = &PresentationDocument{
		Package: pkg,
		Part:    part,
	}
	return nil
}
