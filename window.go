package win32go

type RECT struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
}

type POINT struct {
	X, Y int32
}
type SIZE struct {
	Width, Height int32
}

func (s *SIZE) Equals(other *SIZE) bool {
	return s.Width == other.Width && s.Height == other.Height
}

type BOOL int8

const (
	TRUE  BOOL = 1
	FALSE BOOL = 0
)

func BoolToBOOL(b bool) BOOL {
	if b {
		return TRUE
	}
	return FALSE
}
