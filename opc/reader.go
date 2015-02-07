package opc

import (
	"archive/zip"
	"io"
)

var DefaultReader = &ZipReader{}

func RegisterReadFormat(contentType string, fn func(pkg *Package, partName string, in io.Reader) error) {
	DefaultReader.RegisterFormat(contentType, fn)
}

func Open(name string) (*Package, error) {
	return DefaultReader.Open(name)
}

type Reader struct {
	formats map[string]builderFunc
}

func (rd *Reader) RegisterFormat(contentType string, read func(*Package, string, io.Reader) error) {
	if rd.formats == nil {
		rd.formats = make(map[string]builderFunc)
	}
	rd.formats[contentType] = read
}

type ZipReader struct {
	Reader
}

func (rd *ZipReader) Open(name string) (*Package, error) {
	pkg := &packageBuilder{
		Package:  &Package{Part: Part{Name: "/"}},
		Builders: rd.formats,
		Files:    map[string]opener{},
	}

	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		pkg.Files["/"+f.Name] = f
	}
	if err := pkg.buildTypes(); err != nil {
		return nil, err
	}

	err = pkg.buildAllRels()
	if err != nil {
		return nil, err
	}

	err = pkg.buildAllFiles()
	return pkg.Package, err
}
