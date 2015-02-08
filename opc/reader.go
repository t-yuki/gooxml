package opc

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
)

var DefaultReadFormatter = &ReadFormatter{}

func RegisterReadFormat(contentType string, fn func(pkg *Package, partName string, in io.Reader) error) {
	DefaultReadFormatter.RegisterFormat(contentType, fn)
}

func Open(name string) (*Package, error) {
	rd := &ZipReader{DefaultReadFormatter}
	return rd.Open(name)
}

func Read(in io.Reader) (*Package, error) {
	rd := &ZipReader{DefaultReadFormatter}
	return rd.Read(in)
}

type ReadFormatter struct {
	formats map[string]builderFunc
}

func (rd *ReadFormatter) RegisterFormat(contentType string, read func(*Package, string, io.Reader) error) {
	if rd.formats == nil {
		rd.formats = make(map[string]builderFunc)
	}
	rd.formats[contentType] = read
}

type ZipReader struct {
	*ReadFormatter
}

func (rd *ZipReader) Open(name string) (*Package, error) {
	r, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return rd.read(&r.Reader)
}

func (rd *ZipReader) Read(in io.Reader) (*Package, error) {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, err
	}
	return rd.read(r)
}

func (rd *ZipReader) read(in *zip.Reader) (*Package, error) {
	pkg := &packageBuilder{
		Package:  &Package{Part: Part{Name: "/"}},
		Builders: rd.formats,
		Files:    map[string]opener{},
	}

	for _, f := range in.File {
		pkg.Files["/"+f.Name] = f
	}
	if err := pkg.buildTypes(); err != nil {
		return nil, err
	}

	if err := pkg.buildAllRels(); err != nil {
		return nil, err
	}

	err := pkg.buildAllFiles()
	return pkg.Package, err
}
