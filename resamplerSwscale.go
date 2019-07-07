package ttLibGo

// SwscaleModes Swscaleのモードリスト
var SwscaleModes = struct {
	X            subType
	Area         subType
	Sinc         subType
	Point        subType
	Gauss        subType
	Spline       subType
	Bicubic      subType
	Lanczos      subType
	Bilinear     subType
	Bicublin     subType
	FastBilinear subType
}{
	subType{"X"},
	subType{"Area"},
	subType{"Sinc"},
	subType{"Point"},
	subType{"Gauss"},
	subType{"Spline"},
	subType{"Bicubic"},
	subType{"Lanczos"},
	subType{"Bilinear"},
	subType{"Bicublin"},
	subType{"FastBilinear"},
}
