package win32go

import "testing"

func TestGetModuleHandle(t *testing.T) {
	i := GetModuleHandle(nil)
	if i != 0 {
		t.Log(i)
	}
	fileName := NewUint16Ptr(256)
	b := GetModuleFileName(i, fileName.Pointer(), 256)
	if b {
		t.Log(fileName.String())
	}
}

func TestGetConsoleWindow(t *testing.T) {
	h := GetConsoleWindow()
	t.Log(h)
	fileName := NewUint16Ptr(256)
	b := GetModuleFileName(HINSTANCT(h), fileName.Pointer(), 256)
	if b {
		t.Log(fileName.String())
	}
}
