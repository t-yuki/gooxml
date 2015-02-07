package opc

import "encoding/xml"

// From ECMA-376 Annex F. Standard Namespaces and Content Types
const (
	ContentTypeCoreProperties              = "application/vnd.openxmlformats-package.core-properties+xml"
	ContentTypeDigitalSignatureCertificate = "application/vnd.openxmlformats-package.digital-signaturecertificate"
	ContentTypeDigitalSignatureOrigin      = "application/vnd.openxmlformats-package.digital-signature-origin"
	ContentTypeXMLSignature                = "application/vnd.openxmlformats-package.digital-signaturexmlsignature+xml"
	ContentTypeRelationships               = "application/vnd.openxmlformats-package.relationships+xml"
)

type ContentTypes struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Default  []ContentTypeDefault
	Override []ContentTypeOverride
}

type ContentTypeDefault struct {
	Extension   string `xml:",attr"`
	ContentType string `xml:",attr"`
}

type ContentTypeOverride struct {
	PartName    string `xml:",attr"`
	ContentType string `xml:",attr"`
}
