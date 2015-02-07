package opc

import "testing"

func TestOpen(t *testing.T) {
	pkg, err := Open("../testdata/06.Hook into SharePoint APIs with Android.pptx")
	if err != nil {
		t.Log(err)
	}
	if pkg == nil {
		t.Fatal(err)
	}
	if pkg.Part.Name != "/" {
		t.FailNow()
	}
	if len(pkg.Types.Default) == 0 {
		t.Fatal(pkg.Types.Default)
	}
	if len(pkg.Types.Override) == 0 {
		t.Fatal(pkg.Types.Override)
	}
	if len(pkg.Parts) == 0 {
		t.Fatal(pkg.Parts)
	}
	for _, p := range pkg.Parts {
		t.Logf("%+v", p)
	}
}
