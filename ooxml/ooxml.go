package ooxml

import "github.com/t-yuki/gooxml/opc"

const (
	RelationTypeOfficeDocument     opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	RelationTypeCustomProperties   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/custom-properties"
	RelationTypeExtendedProperties opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"

	RelationTypeViewProps   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"
	RelationTypeCustomXML   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml"
	RelationTypeTableStyles opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"
	RelationTypeTheme       opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"
	RelationTypePresProps   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"

	RelationTypeSlide         opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	RelationTypeSlideMaster   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
	RelationTypeHandoutMaster opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"
	RelationTypeNotesMaster   opc.RelationType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"
)
