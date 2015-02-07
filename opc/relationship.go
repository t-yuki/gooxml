package opc

import "encoding/xml"

type RelationType string

// From ECMA-376 Annex F. Standard Namespaces and Content Types
const (
	RelationTypeCoreProperties              RelationType = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	RelatinoTypeDigitalSignature            RelationType = "http://schemas.openxmlformats.org/package/2006/relationships/digitalsignature/signature"
	RelatinoTypeDigitalSignatureCertificate RelationType = "http://schemas.openxmlformats.org/package/2006/relationships/digitalsignature/certificate"
	RelatinoTypeDigitalSignatureOrigin      RelationType = "http://schemas.openxmlformats.org/package/2006/relationships/digitalsignature/origin"
	RelationTypeThumbnail                   RelationType = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/thumbnail"
)

type Relationship struct {
	Id     string       `xml:",attr"`
	Type   RelationType `xml:",attr"`
	Target string       `xml:",attr"`
}

type relationships struct {
	XMLName      xml.Name `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationship []*Relationship
}
