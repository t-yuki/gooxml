package opc

import (
	"path"
	"strings"
)

type Package struct {
	Part
	Types ContentTypes
	Parts []*Part
}

func (p *Package) ContentType(partName string) string {
	for _, ov := range p.Types.Override {
		if ov.PartName == partName {
			return ov.ContentType
		}
	}
	ext := strings.TrimLeft(path.Ext(partName), ".")
	for _, def := range p.Types.Default {
		if def.Extension == ext {
			return def.ContentType
		}
	}
	return ""
}

func (p *Package) FindPart(partName string) *Part {
	for _, part := range p.Parts {
		if part.Name == partName {
			return part
		}
	}
	return nil
}

func (p *Package) FindPartsByRelationOn(base *Part, relType RelationType) []*Part {
	dir, _ := path.Split(base.Name)
	parts := make([]*Part, 0, 10)

	for _, rel := range base.Relationships {
		if rel.Type == relType {
			targetName := rel.Target
			if !path.IsAbs(targetName) {
				targetName = path.Clean(path.Join(dir, targetName))
			}
			target := p.FindPart(targetName)
			if target == nil {
				panic(targetName)
			}
			parts = append(parts, target)
		}
	}
	return parts
}
