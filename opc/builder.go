package opc

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
)

type builderFunc func(pkg *Package, partName string, in io.Reader) error

type opener interface {
	Open() (rc io.ReadCloser, err error)
}

type packageBuilder struct {
	*Package
	Builders map[string]builderFunc
	Files    map[string]opener
}

const typesFileName = "/[Content_Types].xml"

func (p *packageBuilder) buildTypes() error {
	f := p.Files[typesFileName]
	if f == nil {
		return errors.New("types not found")
	}
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	dec := xml.NewDecoder(rc)
	if err := dec.Decode(&p.Types); err != nil {
		return err
	}
	return nil
}

func (p *packageBuilder) buildAllRels() error {
	for name, f := range p.Files {
		if name == typesFileName {
			continue
		}
		if p.ContentType(name) != ContentTypeRelationships {
			continue
		}

		err := p.buildRels(name, f)
		if err != nil {
			return err
		}
	}
	return nil
}

type errorSet []error

func (e errorSet) Error() string {
	return fmt.Sprintf("%+v", []error(e))
}

func (p *packageBuilder) buildAllFiles() error {
	var errs errorSet
	for name, f := range p.Files {
		if name == typesFileName {
			continue
		}
		if p.ContentType(name) == ContentTypeRelationships {
			continue
		}
		err := p.buildFile(name, f)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (p *packageBuilder) buildRels(name string, f opener) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	rels := &relationships{}
	dec := xml.NewDecoder(rc)
	if err := dec.Decode(rels); err != nil {
		return err
	}
	dir, file := path.Split(name)
	file = strings.TrimSuffix(file, ".rels")
	dir = strings.TrimSuffix(dir, "_rels/")
	if dir == "/" {
		p.Relationships = rels.Relationship
	} else {
		partName := path.Join(dir, file)
		p.Parts = append(p.Parts, &Part{
			Name:          partName,
			ContentType:   p.ContentType(partName),
			Relationships: rels.Relationship,
		})
	}
	return nil
}

func (p *packageBuilder) buildFile(name string, f opener) error {
	t := p.ContentType(name)
	build := p.Builders[t]
	if build == nil {
		return errors.New("path: `" + name + "` has unknown content type: " + t)
	}

	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	part := p.FindPart(name)
	if part == nil {
		p.Parts = append(p.Parts, &Part{
			Name:        name,
			ContentType: t,
		})
	}

	return build(p.Package, name, rc)
}
