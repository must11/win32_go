package win32go

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	/*Flags

	类型：DWORD

	可用于初始化对话框的一组位标志。 当对话框返回时，它会设置这些标志以指示用户的输入。 此成员可以是以下标志的组合。

	*/
	FN_ALLOWMULTISELECT      = 0x00000200
	OFN_CREATEPROMPT         = 0x00002000
	OFN_DONTADDTORECENT      = 0x02000000
	OFN_ENABLEHOOK           = 0x00000020
	OFN_ENABLEINCLUDENOTIFY  = 0x00400000
	OFN_ENABLESIZING         = 0x00800000
	OFN_ENABLETEMPLATE       = 0x00000040
	OFN_ENABLETEMPLATEHANDLE = 0x00000080
	OFN_EXPLORER             = 0x00080000
	OFN_EXTENSIONDIFFERENT   = 0x00000400
	OFN_FILEMUSTEXIST        = 0x00001000
	OFN_FORCESHOWHIDDEN      = 0x10000000
	OFN_HIDEREADONLY         = 0x00000004
	OFN_LONGNAMES            = 0x00200000
	OFN_NOCHANGEDIR          = 0x00000008
	OFN_NODEREFERENCELINKS   = 0x00100000
	OFN_NOLONGNAMES          = 0x00040000
	OFN_NONETWORKBUTTON      = 0x00020000
	OFN_NOREADONLYRETURN     = 0x00008000
	OFN_NOTESTFILECREATE     = 0x00010000
	OFN_NOVALIDATE           = 0x00000100
	OFN_OVERWRITEPROMPT      = 0x00000002
	OFN_PATHMUSTEXIST        = 0x00000800
	OFN_READONLY             = 0x00000001
	OFN_SHAREAWARE           = 0x00004000
	OFN_SHOWHELP             = 0x00000010
)

/*
FlagsEx

类型：DWORD

一组可用于初始化对话框的位标志。 目前，此成员可以是零或以下标志
*/
const OFN_EX_NOPLACESBAR = 0x00000001

type CHOOSECOLOR struct {
}

// https://learn.microsoft.com/zh-cn/windows/win32/api/commdlg/ns-commdlg-openfilenamea
type OPENFILENAME struct {
	LStructSize       uint32
	HwndOwner         windows.HWND
	HInstance         windows.HWND
	LpstrFilter       *uint16
	LpstrCustomFilter *uint16
	NMaxCustFilter    uint32
	NFilterIndex      uint32
	LpstrFile         *uint16
	NMaxFile          uint32
	LpstrFileTitle    *uint16
	NMaxFileTitle     uint32
	LpstrInitialDir   *uint16
	LpstrTitle        *uint16
	Flags             uint32
	NFileOffset       uint16
	NFileExtension    uint16
	LpstrDefExt       *uint16
	LCustData         uintptr
	LpTemplateName    *uint16
	LpstrPrompt       uintptr
	PvReserved        unsafe.Pointer
	DwReserved        uint32
	FlagsEx           uint32
}

func NewOPENFILENAME() *OPENFILENAME {
	o := OPENFILENAME{}
	o.LStructSize = uint32(unsafe.Sizeof(o))
	return &o
}

type PRINTPAGERANGE struct {
	NFromPage uint32
	NToPage   uint32
}

type PRINTDLGEX struct {
	lStructSize         int32
	hwndOwner           windows.HWND
	hDevMode            uintptr
	HDevNames           uintptr
	HDC                 uintptr
	Flags               uint32
	ExclusionFlags      uint32
	NPageRanges         uint32
	NMaxPageRanges      uint32
	LpPageRanges        PRINTPAGERANGE
	NMinPage            uint32
	NMaxPage            uint32
	NCopies             uint32
	HInstance           windows.HWND
	LpPrintTemplateName *uint16
	LpCallback          uintptr
	NPropertyPages      uint32

	LphPropertyPages *HPROPSHEETPAGE
	NStartPage       uint32
	DwResultAction   uint32
}
type HPROPSHEETPAGE uintptr

var chooseColorProc *windows.LazyProc
var openFileNameProc *windows.LazyProc
var saveFileNameProc *windows.LazyProc
var printDlgExProc *windows.LazyProc
var libcomdlg32 *windows.LazyDLL

func init() {
	libcomdlg32 = windows.NewLazySystemDLL("comdlg32.dll")
	chooseColorProc = libcomdlg32.NewProc("ChooseColorW")
	openFileNameProc = libcomdlg32.NewProc("GetOpenFileNameW")
	saveFileNameProc = libcomdlg32.NewProc("GetSaveFileNameW")
	printDlgExProc = libcomdlg32.NewProc("PrintDlgExW")
}

func ChooseColor(choosecolor *CHOOSECOLOR) bool {
	ret, _, _ := syscall.SyscallN(chooseColorProc.Addr(),
		uintptr(unsafe.Pointer(choosecolor)))
	return ret != 0
}

func GetOpenFileName(openfileName *OPENFILENAME) bool {
	ret, _, _ := syscall.SyscallN(openFileNameProc.Addr(),
		uintptr(unsafe.Pointer(openfileName)))
	return ret != 0
}

func GetSaveFileName(openfileName *OPENFILENAME) bool {
	ret, _, _ := syscall.SyscallN(saveFileNameProc.Addr(),
		uintptr(unsafe.Pointer(openfileName)))
	return ret != 0
}
