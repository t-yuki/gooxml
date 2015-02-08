package ooxml

import (
	"errors"
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
