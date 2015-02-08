package ooxml

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestPresentationSlide_decode1(t *testing.T) {
	var buf = bytes.NewBufferString(`
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sld xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:spTree>
      <p:nvGrpSpPr><p:cNvPr id="1" name=""/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr>
      <p:grpSpPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="0" cy="0"/><a:chOff x="0" y="0"/><a:chExt cx="0" cy="0"/></a:xfrm></p:grpSpPr>
      <p:sp><p:nvSpPr><p:cNvPr id="4" name="Subtitle 3"/><p:cNvSpPr><a:spLocks noGrp="1"/></p:cNvSpPr><p:nvPr><p:ph type="subTitle" idx="1"/></p:nvPr></p:nvSpPr><p:spPr/><p:txBody><a:bodyPr/><a:lstStyle/><a:p><a:r><a:rPr lang="en-US" dirty="0" smtClean="0"/><a:t>Authentication with azure AD</a:t></a:r><a:endParaRPr lang="en-US" dirty="0"/></a:p></p:txBody></p:sp>
      <p:sp><p:nvSpPr><p:cNvPr id="2" name="Text Placeholder 1"/><p:cNvSpPr><a:spLocks noGrp="1"/></p:cNvSpPr><p:nvPr><p:ph type="body" sz="quarter" idx="10"/></p:nvPr></p:nvSpPr><p:spPr/><p:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" indent="0"><a:buNone/></a:pPr><a:r><a:rPr lang="en-US" dirty="0" smtClean="0"/><a:t>demo</a:t></a:r><a:endParaRPr lang="en-US" dirty="0"/></a:p></p:txBody></p:sp>
    </p:spTree>
    <p:extLst>
      <p:ext uri="{BB962C8B-B14F-4D97-AF65-F5344CB8AC3E}"><p14:creationId xmlns:p14="http://schemas.microsoft.com/office/powerpoint/2010/main" val="3536787970"/></p:ext>
    </p:extLst>
  </p:cSld>
  <p:clrMapOvr><a:masterClrMapping/></p:clrMapOvr>
  <p:transition><p:fade/></p:transition>
  <p:timing><p:tnLst><p:par><p:cTn id="1" dur="indefinite" restart="never" nodeType="tmRoot"/></p:par></p:tnLst></p:timing>
</p:sld>
	`)
	x := &PresentationSlide{}
	dec := xml.NewDecoder(buf)
	if err := dec.Decode(x); err != nil {
		t.Fatal(err)
	}
	if txt := x.String(); !strings.Contains(txt, "Authentication") {
		t.Fatal(txt)
	}
}

func TestPresentationSlide_decode2(t *testing.T) {
	var buf = bytes.NewBufferString(`
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<p:sld xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main">
  <p:cSld>
    <p:spTree>
      <p:nvGrpSpPr><p:cNvPr id="1" name=""/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr>
      <p:grpSpPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="0" cy="0"/><a:chOff x="0" y="0"/><a:chExt cx="0" cy="0"/></a:xfrm></p:grpSpPr>
      <p:sp><p:nvSpPr><p:cNvPr id="5" name="Title 4"/><p:cNvSpPr><a:spLocks noGrp="1"/></p:cNvSpPr><p:nvPr><p:ph type="title"/></p:nvPr></p:nvSpPr><p:spPr/><p:txBody><a:bodyPr/><a:lstStyle/><a:p><a:r><a:rPr lang="en-US" dirty="0" smtClean="0"/><a:t>Course Agenda</a:t></a:r><a:endParaRPr lang="en-US" dirty="0"/></a:p></p:txBody></p:sp>
      <p:graphicFrame>
        <p:nvGraphicFramePr><p:cNvPr id="10" name="Content Placeholder 9"/><p:cNvGraphicFramePr><a:graphicFrameLocks noGrp="1"/></p:cNvGraphicFramePr><p:nvPr><p:ph sz="quarter" idx="10"/><p:extLst><p:ext uri="{D42A27DB-BD31-4B8C-83A1-F6EECF244321}"><p14:modId xmlns:p14="http://schemas.microsoft.com/office/powerpoint/2010/main" val="1978659899"/></p:ext></p:extLst></p:nvPr></p:nvGraphicFramePr><p:xfrm><a:off x="351383" y="1063255"/><a:ext cx="11225057" cy="4633690"/></p:xfrm>
        <a:graphic><a:graphicData uri="http://schemas.openxmlformats.org/drawingml/2006/table">
          <a:tbl>
            <a:tblPr firstRow="1" bandRow="1"><a:tableStyleId>{5C22544A-7EE6-4342-B048-85BDC9FD1C3A}</a:tableStyleId></a:tblPr>
            <a:tblGrid><a:gridCol w="11225057"><a:extLst><a:ext uri="{9D8B030D-6E8A-4147-A177-3AD203B41FA5}"><a16:colId xmlns="" xmlns:a16="http://schemas.microsoft.com/office/drawing/2014/main" val="1253488153"/></a:ext></a:extLst></a:gridCol></a:tblGrid>
            <a:tr h="273317">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:r><a:rPr lang="en-US" sz="2800" dirty="0" smtClean="0"/><a:t>Office Camp</a:t></a:r><a:endParaRPr lang="en-US" sz="2800" dirty="0"/></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
              <a:extLst><a:ext uri="{0D108BD9-81ED-4DB2-BD59-A6C34878D82A}"><a16:rowId xmlns="" xmlns:a16="http://schemas.microsoft.com/office/drawing/2014/main" val="829859176"/></a:ext></a:extLst>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:r><a:rPr lang="en-US" sz="2400" b="0" dirty="0" smtClean="0"/><a:t>Module 1: Introduction to the Day</a:t></a:r></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
              <a:extLst><a:ext uri="{0D108BD9-81ED-4DB2-BD59-A6C34878D82A}"><a16:rowId xmlns="" xmlns:a16="http://schemas.microsoft.com/office/drawing/2014/main" val="1946132611"/></a:ext></a:extLst>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" marR="0" indent="0" algn="l" defTabSz="932559" rtl="0" eaLnBrk="1" fontAlgn="auto" latinLnBrk="0" hangingPunct="1"><a:lnSpc><a:spcPct val="100000"/></a:lnSpc><a:spcBef><a:spcPts val="0"/></a:spcBef><a:spcAft><a:spcPts val="0"/></a:spcAft><a:buClrTx/><a:buSzTx/><a:buFontTx/><a:buNone/><a:tabLst/><a:defRPr/></a:pPr><a:r><a:rPr lang="en-US" sz="2400" dirty="0" smtClean="0"/><a:t>Module 2: Setting up the Environments</a:t></a:r></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
              <a:extLst><a:ext uri="{0D108BD9-81ED-4DB2-BD59-A6C34878D82A}"><a16:rowId xmlns="" xmlns:a16="http://schemas.microsoft.com/office/drawing/2014/main" val="3204002662"/></a:ext></a:extLst>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" marR="0" indent="0" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" fontAlgn="auto" latinLnBrk="0" hangingPunct="1"><a:lnSpc><a:spcPct val="100000"/></a:lnSpc><a:spcBef><a:spcPts val="0"/></a:spcBef><a:spcAft><a:spcPts val="0"/></a:spcAft><a:buClrTx/><a:buSzTx/><a:buFontTx/><a:buNone/><a:tabLst/><a:defRPr/></a:pPr><a:r><a:rPr lang="en-US" sz="2400" dirty="0" smtClean="0"/><a:t>Module 3: Hooking into Apps for SharePoint</a:t></a:r></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
              <a:extLst><a:ext uri="{0D108BD9-81ED-4DB2-BD59-A6C34878D82A}"><a16:rowId xmlns="" xmlns:a16="http://schemas.microsoft.com/office/drawing/2014/main" val="4266278162"/></a:ext></a:extLst>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" marR="0" indent="0" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" fontAlgn="auto" latinLnBrk="0" hangingPunct="1"><a:lnSpc><a:spcPct val="100000"/></a:lnSpc><a:spcBef><a:spcPts val="0"/></a:spcBef><a:spcAft><a:spcPts val="0"/></a:spcAft><a:buClrTx/><a:buSzTx/><a:buFontTx/><a:buNone/><a:tabLst/><a:defRPr/></a:pPr><a:r><a:rPr lang="en-US" sz="2400" b="0" dirty="0" smtClean="0"/><a:t>Module 4: </a:t></a:r><a:r><a:rPr lang="en-US" sz="2400" dirty="0" smtClean="0"/><a:t>Hooking into Office </a:t></a:r><a:r><a:rPr lang="en-US" sz="2400" smtClean="0"/><a:t>365 </a:t></a:r><a:r><a:rPr lang="en-US" sz="2400" smtClean="0"/><a:t>APIs</a:t></a:r><a:endParaRPr lang="en-US" sz="2400" dirty="0" smtClean="0"/></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" marR="0" indent="0" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" fontAlgn="auto" latinLnBrk="0" hangingPunct="1"><a:lnSpc><a:spcPct val="100000"/></a:lnSpc><a:spcBef><a:spcPts val="0"/></a:spcBef><a:spcAft><a:spcPts val="0"/></a:spcAft><a:buClrTx/><a:buSzTx/><a:buFontTx/><a:buNone/><a:tabLst/><a:defRPr/></a:pPr><a:r><a:rPr lang="en-US" sz="2400" b="0" dirty="0" smtClean="0"/><a:t>Module 5: Hooking into Apps for Office</a:t></a:r></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
            </a:tr>
            <a:tr h="685928">
              <a:tc><a:txBody><a:bodyPr/><a:lstStyle/><a:p><a:pPr marL="0" marR="0" indent="0" algn="l" defTabSz="914363" rtl="0" eaLnBrk="1" fontAlgn="auto" latinLnBrk="0" hangingPunct="1"><a:lnSpc><a:spcPct val="100000"/></a:lnSpc><a:spcBef><a:spcPts val="0"/></a:spcBef><a:spcAft><a:spcPts val="0"/></a:spcAft><a:buClrTx/><a:buSzTx/><a:buFontTx/><a:buNone/><a:tabLst/><a:defRPr/></a:pPr><a:r><a:rPr lang="en-US" sz="2400" b="0" dirty="0" smtClean="0"/><a:t>Module 6:</a:t></a:r><a:r><a:rPr lang="en-US" sz="2400" b="0" baseline="0" dirty="0" smtClean="0"/><a:t> Hooking into SharePoint APIs with Android</a:t></a:r><a:endParaRPr lang="en-US" sz="2400" b="0" dirty="0" smtClean="0"/></a:p></a:txBody><a:tcPr marL="91403" marR="91403" marT="45701" marB="45701" anchor="ctr"/></a:tc>
            </a:tr>
          </a:tbl>
        </a:graphicData></a:graphic>
      </p:graphicFrame>
    </p:spTree>
    <p:extLst>
      <p:ext uri="{BB962C8B-B14F-4D97-AF65-F5344CB8AC3E}"><p14:creationId xmlns:p14="http://schemas.microsoft.com/office/powerpoint/2010/main" val="2628257742"/></p:ext>
    </p:extLst>
  </p:cSld>
  <p:clrMapOvr><a:masterClrMapping/></p:clrMapOvr>
  <p:transition><p:fade/></p:transition>
  <p:timing><p:tnLst><p:par><p:cTn id="1" dur="indefinite" restart="never" nodeType="tmRoot"/></p:par></p:tnLst></p:timing>
</p:sld>
	`)
	x := &PresentationSlide{}
	dec := xml.NewDecoder(buf)
	if err := dec.Decode(x); err != nil {
		t.Fatal(err)
	}
	if txt := x.String(); !strings.Contains(txt, "Module 1: Introduction") {
		t.Fatal(txt)
	}
}
