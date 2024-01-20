package win32go

import (
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	kernel32          *windows.LazyDLL
	getModuleHandle   *windows.LazyProc
	getModuleFileName *windows.LazyProc
	getConsoleWindow  *windows.LazyProc
)

func init() {
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")
	getModuleHandle = kernel32.NewProc("GetModuleHandleW")
}

func GetModuleHandle(lpModuleName *uint16) HINSTANCE {
	ret, _, _ := syscall.SyscallN(getModuleHandle.Addr(), uintptr(unsafe.Pointer(lpModuleName)))
	return HINSTANCE(ret)
}

func GetModuleFileName(hModule HINSTANCE, lpFilename *uint16, nsize int32) bool {
	if getModuleFileName == nil {
		log.Print("init getModuleFileName")
		getModuleFileName = kernel32.NewProc("GetModuleFileNameW")
	}
	ret, _, _ := syscall.SyscallN(getModuleFileName.Addr(), uintptr(hModule), uintptr(unsafe.Pointer(lpFilename)), uintptr(unsafe.Pointer(&nsize)))
	return ret != 0
}
func GetConsoleWindow() HWND {
	if getConsoleWindow == nil {
		getConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	}
	ret, _, _ := syscall.SyscallN(getConsoleWindow.Addr())

	return HWND(ret)
}
