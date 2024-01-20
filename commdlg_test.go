package win32go

import "testing"

func TestGetOpenFileName(t *testing.T) {
	f := NewOPENFILENAME()
	fileName := NewUint16Ptr(128)
	f.LpstrFile = fileName.Pointer()
	f.NMaxFile = 128
	filter := NewFileFiter()
	filter.Add("txt file", "*.txt")
	// t.Log(filter.String())
	f.LpstrFilter = filter.Pointer()

	if GetOpenFileName(f) {
		t.Log(fileName.String())
	}
}

func TestGetSaveFileName(t *testing.T) {
	f := NewOPENFILENAME()
	fileName := NewUint16Ptr(128)
	f.LpstrFile = fileName.Pointer()
	f.NMaxFile = 128
	if GetSaveFileName(f) {
		t.Log(fileName.String())
	}
}
