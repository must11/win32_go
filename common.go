package win32go

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows"
)

// 辅助类型，用于产生 *uint16 分配内存
type uint16Ptr struct {
	s []uint16
}

func (u *uint16Ptr) Pointer() *uint16 {
	return &u.s[0]
}

func (u *uint16Ptr) String() string {
	return strings.TrimRight(windows.UTF16PtrToString(&u.s[0]), " ")
}

func NewUint16Ptr(size uint32) *uint16Ptr {
	return &uint16Ptr{
		s: make([]uint16, size, size),
	}
}

// 文件类型过滤，用于产生符合win32过滤规则的字符串
type FileFilter struct {
	s []uint16
}

const DEFAULT_FILTER_CAP = 64

func NewFileFiter() *FileFilter {
	return &FileFilter{
		s: make([]uint16, 0, DEFAULT_FILTER_CAP),
	}
}

func (ff *FileFilter) Add(text, rule string) {
	if len(ff.s) > 0 {
		ff.s = ff.s[0 : len(ff.s)-1]
	}
	for _, r := range text {
		ff.s = append(ff.s, uint16(r))
	}
	fmt.Print(ff.s)
	ff.s = append(ff.s, uint16(0))
	for _, r := range rule {
		ff.s = append(ff.s, uint16(r))
	}
	fmt.Print(ff.s)
	ff.s = append(ff.s, uint16(0))
	ff.s = append(ff.s, uint16(0))

}
func (ff *FileFilter) Pointer() *uint16 {
	return &ff.s[0]
}

func (ff *FileFilter) String() string {
	r := make([]rune, 0, len(ff.s))
	for _, v := range ff.s {
		r = append(r, rune(v))
	}
	r = r[0 : len(r)-2]
	return string(r)
}
