package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool Openh264Encoder_setRCMode(void *encoder, const char *mode);
extern bool Openh264Encoder_setIDRInterval(void *encoder, uint32_t interval);
extern bool Openh264Encoder_forceNextKeyFrame(void *encoder);
*/
import "C"
import (
	"unsafe"
)

type openh264Encoder encoder

type eoh264UsageType struct {
	value                  string
	CameraVideoRealTime    string
	ScreenContentRealTime  string
	CameraVideoNonRealTime string
}

type eoh264RCModeType struct {
	value                 string
	RcQualityMode         string
	RcBitrateMode         string
	RcBufferbasedMode     string
	RcTimestampMode       string
	RcBitrateModePostSkip string
	RcOffMode             string
}

type eoh264ComplexityMode struct {
	value            string
	LowComplexity    string
	MediumComplexity string
	HighComplexity   string
}

type eoh264SpsPpsIDStrategy struct {
	value                      string
	ConstantID                 string
	IncreasingID               string
	SPSListing                 string
	SPSListingAndPPSIncreasing string
	SPSPPSListing              string
}

// Openh264Params openh264Encoderで利用するParam設定項目
var Openh264Params = struct {
	UsageType                  eoh264UsageType
	ITargetBitrate             subType
	RCMode                     eoh264RCModeType
	FMaxFrameRate              subType
	ITemporalLayerNum          subType
	ISpatialLayerNum           subType
	ComplexityMode             eoh264ComplexityMode
	UIIntraPeriod              subType
	INumRefFrame               subType
	SpsPpsIDStrategy           eoh264SpsPpsIDStrategy
	BPrefixNalAddingCtrl       subType
	BEnableSSEI                subType
	BSimulcastAVC              subType
	IPaddingFlag               subType
	IEntropyCodingModeFlag     subType
	BEnableFrameSkip           subType
	IMaxBitrate                subType
	IMaxQp                     subType
	IMinQp                     subType
	UIMaxNalSize               subType
	BEnableLongTermReference   subType
	ILTRRefNum                 subType
	ILtrMarkPeriod             subType
	IMultipleThreadIdc         subType
	BUseLoadBalancing          subType
	ILoopFilterDisableIdc      subType
	ILoopFilterAlphaC0Offset   subType
	ILoopFilterBetaOffset      subType
	BEnableDenoise             subType
	BEnableBackgroundDetection subType
	BEnableAdaptiveQuant       subType
	BEnableFrameCroppingFlag   subType
	BEnableSceneChangeDetect   subType
	BIsLosslessLink            subType
}{
	UsageType: eoh264UsageType{
		"iUsageType",
		"CAMERA_VIDEO_REAL_TIME",
		"SCREEN_CONTENT_REAL_TIME",
		"CAMERA_VIDEO_NON_REAL_TIME",
	},
	ITargetBitrate: subType{"iTargetBitrate"},
	RCMode: eoh264RCModeType{
		"iRCMode",
		"RC_QUALITY_MODE",
		"RC_BITRATE_MODE",
		"RC_BUFFERBASED_MODE",
		"RC_TIMESTAMP_MODE",
		"RC_BITRATE_MODE_POST_SKIP",
		"RC_OFF_MODE",
	},
	FMaxFrameRate:     subType{"fMaxFrameRate"},
	ITemporalLayerNum: subType{"iTemporalLayerNum"},
	ISpatialLayerNum:  subType{"iSpatialLayerNum"},
	ComplexityMode: eoh264ComplexityMode{
		"iComplexityMode",
		"LOW_COMPLEXITY",
		"MEDIUM_COMPLEXITY",
		"HIGH_COMPLEXITY",
	},
	UIIntraPeriod: subType{"uiIntraPeriod"},
	INumRefFrame:  subType{"iNumRefFrame"},
	SpsPpsIDStrategy: eoh264SpsPpsIDStrategy{
		"eSpsPpsIdStrategy",
		"CONSTANT_ID",
		"INCREASING_ID",
		"SPS_LISTING",
		"SPS_LISTING_AND_PPS_INCREASING",
		"SPS_PPS_LISTING",
	},
	BPrefixNalAddingCtrl:       subType{"bPrefixNalAddingCtrl"},
	BEnableSSEI:                subType{"bEnableSSEI"},
	BSimulcastAVC:              subType{"bSimulcastAVC"},
	IPaddingFlag:               subType{"iPaddingFlag"},
	IEntropyCodingModeFlag:     subType{"iEntropyCodingModeFlag"},
	BEnableFrameSkip:           subType{"bEnableFrameSkip"},
	IMaxBitrate:                subType{"iMaxBitrate"},
	IMaxQp:                     subType{"iMaxQp"},
	IMinQp:                     subType{"iMinQp"},
	UIMaxNalSize:               subType{"uiMaxNalSize"},
	BEnableLongTermReference:   subType{"bEnableLongTermReference"},
	ILTRRefNum:                 subType{"iLTRRefNum"},
	ILtrMarkPeriod:             subType{"iLtrMarkPeriod"},
	IMultipleThreadIdc:         subType{"iMultipleThreadIdc"},
	BUseLoadBalancing:          subType{"bUseLoadBalancing"},
	ILoopFilterDisableIdc:      subType{"iLoopFilterDisableIdc"},
	ILoopFilterAlphaC0Offset:   subType{"iLoopFilterAlphaC0Offset"},
	ILoopFilterBetaOffset:      subType{"iLoopFilterBetaOffset"},
	BEnableDenoise:             subType{"bEnableDenoise"},
	BEnableBackgroundDetection: subType{"bEnableBackgroundDetection"},
	BEnableAdaptiveQuant:       subType{"bEnableAdaptiveQuant"},
	BEnableFrameCroppingFlag:   subType{"bEnableFrameCroppingFlag"},
	BEnableSceneChangeDetect:   subType{"bEnableSceneChangeDetect"},
	BIsLosslessLink:            subType{"bIsLosslessLink"},
}

type eoh264ProfileIdc struct {
	value               string
	ProUNKNOWN          string
	ProBASELINE         string
	ProMAIN             string
	ProEXTENDED         string
	ProHIGH             string
	ProHIGH10           string
	ProHIGH422          string
	ProHIGH444          string
	ProCAVLC444         string
	ProScalableBASELINE string
	ProScalableHIGH     string
}

type eoh264LevelIdc struct {
	value        string
	LevelUNKNOWN string
	Level1_0     string
	Level1_B     string
	Level1_1     string
	Level1_2     string
	Level1_3     string
	Level2_0     string
	Level2_1     string
	Level2_2     string
	Level3_0     string
	Level3_1     string
	Level3_2     string
	Level4_0     string
	Level4_1     string
	Level4_2     string
	Level5_0     string
	Level5_1     string
	Level5_2     string
}

// openh264 1 5以下
type eoh264SliceCfgSliceMode struct {
	value              string
	SmSingleSLICE      string
	SmFixedslcnumSLICE string
	SmRasterSLICE      string
	SmRowmbSLICE       string
	SmDynSLICE         string
	SmAutoSLICE        string
	SmRESERVED         string
}

type eoh264SliceCfg struct {
	SliceMode     eoh264SliceCfgSliceMode
	SliceArgument eoh264SliceCfgSliceArgument
}

type eoh264SliceCfgSliceArgument struct {
	UISliceNum            subType
	UISliceSizeConstraint subType
}

// openh264 1 5以下、ここまで

// openh264 1 6以上

type eoh264SliceArgumentSliceMode struct {
	value              string
	SmSingleSLICE      string
	SmFixedslcnumSLICE string
	SmRasterSLICE      string
	SmSizelimitedSLICE string
	SmRESERVED         string
}

type eoh264SliceArgument struct {
	SliceMode             eoh264SliceArgumentSliceMode
	UISliceNum            subType
	UISliceSizeConstraint subType
}

type eoh264VideoFormat struct {
	value       string
	VfCOMPONENT string
	VfPAL       string
	VfNTSC      string
	VfSECAM     string
	VfMAC       string
	VfUNDEF     string
	VfNumENUM   string
}

type eoh264ColorPrimaries struct {
	value       string
	CpRESERVED0 string
	CpBT709     string
	CpUNDEF     string
	CpRESERVED3 string
	CpBT470M    string
	CpBT470BG   string
	CpSMPTE170M string
	CpSMPTE240M string
	CpFILM      string
	CpBT2020    string
	CpNumENUM   string
}

type eoh264TransferCharacteristics struct {
	value           string
	TrcRESERVED0    string
	TrcBT709        string
	TrcUNDEF        string
	TrcRESERVED3    string
	TrcBT470M       string
	TrcBT470BG      string
	TrcSMPTE170M    string
	TrcSMPTE240M    string
	TrcLINEAR       string
	TrcLOG100       string
	TrcLOG316       string
	TrcIEC61966_2_4 string
	TrcBT1361E      string
	TrcIEC61966_2_1 string
	TrcBT2020_10    string
	TrcBT2020_12    string
	TrcNumENUM      string
}

type eoh264ColorMatrix struct {
	value       string
	CmGBR       string
	CmBT709     string
	CmUNDEF     string
	CmRESERVED3 string
	CmFCC       string
	CmBT470BG   string
	CmSMPTE170M string
	CmSMPTE240M string
	CmYCGCO     string
	CmBT2020NC  string
	CmBT2020C   string
	CmNumENUM   string
}

// openh264 1 6以上、ここまで

// Openh264SpatialParams openh264Encoderで利用するspatial paramの設定項目
var Openh264SpatialParams = struct {
	IVideoWidth        subType
	IVideoHeight       subType
	FFrameRate         subType
	ISpatialBitrate    subType
	IMaxSpatialBitrate subType
	ProfileIdc         eoh264ProfileIdc
	LevelIdc           eoh264LevelIdc
	IDLayerQp          subType
	// 古いopenh264用 1.5以下
	SliceCfg eoh264SliceCfg
	// 新しいopenh264用
	SliceArgument            eoh264SliceArgument
	BVideoSignalTypePresent  subType
	VideoFormat              eoh264VideoFormat
	BFullRange               subType
	BColorDescriptionPresent subType
	ColorPrimaries           eoh264ColorPrimaries
	TransferCharacteristics  eoh264TransferCharacteristics
	ColorMatrix              eoh264ColorMatrix
}{
	IVideoWidth:        subType{"iVideoWidth"},
	IVideoHeight:       subType{"iVideoHeight"},
	FFrameRate:         subType{"fFrameRate"},
	ISpatialBitrate:    subType{"iSpatialBitrate"},
	IMaxSpatialBitrate: subType{"iMaxSpatialBitrate"},
	ProfileIdc: eoh264ProfileIdc{
		"uiProfileIdc",
		"PRO_UNKNOWN",
		"PRO_BASELINE",
		"PRO_MAIN",
		"PRO_EXTENDED",
		"PRO_HIGH",
		"PRO_HIGH10",
		"PRO_HIGH422",
		"PRO_HIGH444",
		"PRO_CAVLC444",
		"PRO_SCALABLE_BASELINE",
		"PRO_SCALABLE_HIGH",
	},
	LevelIdc: eoh264LevelIdc{
		"uiLevelIdc",
		"LEVEL_UNKNOWN",
		"LEVEL_1_0",
		"LEVEL_1_B",
		"LEVEL_1_1",
		"LEVEL_1_2",
		"LEVEL_1_3",
		"LEVEL_2_0",
		"LEVEL_2_1",
		"LEVEL_2_2",
		"LEVEL_3_0",
		"LEVEL_3_1",
		"LEVEL_3_2",
		"LEVEL_4_0",
		"LEVEL_4_1",
		"LEVEL_4_2",
		"LEVEL_5_0",
		"LEVEL_5_1",
		"LEVEL_5_2",
	},
	IDLayerQp: subType{"iDLayerQp"},
	// これは古い設定openh264 1.5以下
	SliceCfg: eoh264SliceCfg{
		SliceMode: eoh264SliceCfgSliceMode{
			"sSliceCfg.uiSliceMode",
			"SM_SINGLE_SLICE",
			"SM_FIXEDSLCNUM_SLICE",
			"SM_RASTER_SLICE",
			"SM_ROWMB_SLICE",
			"SM_DYN_SLICE",
			"SM_AUTO_SLICE",
			"SM_RESERVEDs",
		},
		SliceArgument: eoh264SliceCfgSliceArgument{
			UISliceNum:            subType{"sSliceCfg.sSliceArgument.uiSliceNum"},
			UISliceSizeConstraint: subType{"sSliceCfg.sSliceArgument.uiSliceSizeConstraint"},
		},
	},
	// --------
	// こっちは新しい設定openh264 1.6以上
	SliceArgument: eoh264SliceArgument{
		SliceMode: eoh264SliceArgumentSliceMode{
			"sSliceArgument.uiSliceMode",
			"SM_SINGLE_SLICE",
			"SM_FIXEDSLCNUM_SLICE",
			"SM_RASTER_SLICE",
			"SM_SIZELIMITED_SLICE",
			"SM_RESERVED",
		},
		UISliceNum:            subType{"sSliceArgument.uiSliceNum"},
		UISliceSizeConstraint: subType{"sSliceArgument.uiSliceSizeConstraint"},
	},
	BVideoSignalTypePresent: subType{"bVideoSignalTypePresent"},
	VideoFormat: eoh264VideoFormat{
		"uiVideoFormat",
		"VF_COMPONENT",
		"VF_PAL",
		"VF_NTSC",
		"VF_SECAM",
		"VF_MAC",
		"VF_UNDEF",
		"VF_NUM_ENUM",
	},
	BFullRange:               subType{"bFullRange"},
	BColorDescriptionPresent: subType{"bColorDescriptionPresent"},
	ColorPrimaries: eoh264ColorPrimaries{
		"uiColorPrimaries",
		"CP_RESERVED0",
		"CP_BT709",
		"CP_UNDEF",
		"CP_RESERVED3",
		"CP_BT470M",
		"CP_BT470BG",
		"CP_SMPTE170M",
		"CP_SMPTE240M",
		"CP_FILM",
		"CP_BT2020",
		"CP_NUM_ENUM",
	},
	TransferCharacteristics: eoh264TransferCharacteristics{
		"uiTransferCharacteristics",
		"TRC_RESERVED0",
		"TRC_BT709",
		"TRC_UNDEF",
		"TRC_RESERVED3",
		"TRC_BT470M",
		"TRC_BT470BG",
		"TRC_SMPTE170M",
		"TRC_SMPTE240M",
		"TRC_LINEAR",
		"TRC_LOG100",
		"TRC_LOG316",
		"TRC_IEC61966_2_4",
		"TRC_BT1361E",
		"TRC_IEC61966_2_1",
		"TRC_BT2020_10",
		"TRC_BT2020_12",
		"TRC_NUM_ENUM",
	},
	ColorMatrix: eoh264ColorMatrix{
		"uiColorMatrix",
		"CM_GBR",
		"CM_BT709",
		"CM_UNDEF",
		"CM_RESERVED3",
		"CM_FCC",
		"CM_BT470BG",
		"CM_SMPTE170M",
		"CM_SMPTE240M",
		"CM_YCGCO",
		"CM_BT2020NC",
		"CM_BT2020C",
		"CM_NUM_ENUM",
	},
	// こっちは新しい設定openh264 1.6以上 ここまで
}

// Openh264RCModes rcModeで設定できる値
var Openh264RCModes = struct {
	QualityMode         subType
	BitrateMode         subType
	BufferbasedMode     subType
	TimestampMode       subType
	BitrateModePostSkip subType
	OffMode             subType
}{
	subType{"QualityMode"},
	subType{"BitrateMode"},
	subType{"BufferbasedMode"},
	subType{"TimestampMode"},
	subType{"BitrateModePostSkip"},
	subType{"OffMode"},
}

// EncodeFrame エンコードを実行
func (openh264Encoder *openh264Encoder) EncodeFrame(
	frame IFrame,
	callback FrameCallback) bool {
	return encoderEncodeFrame((*encoder)(openh264Encoder), frame, callback)
}

// Close 閉じる
func (openh264Encoder *openh264Encoder) Close() {
	encoderClose((*encoder)(openh264Encoder))
}

// SetRCMode rcModeを設定する
func (openh264Encoder *openh264Encoder) SetRCMode(mode subType) bool {
	cmode := C.CString(mode.value)
	defer C.free(unsafe.Pointer(cmode))
	return bool(C.Openh264Encoder_setRCMode(unsafe.Pointer(openh264Encoder.cEncoder), cmode))
}

// SetIDRInterval keyFrame間隔を数値で設定する
func (openh264Encoder *openh264Encoder) SetIDRInterval(interval uint32) bool {
	return bool(C.Openh264Encoder_setIDRInterval(unsafe.Pointer(openh264Encoder.cEncoder), C.uint32_t(interval)))
}

// ForceNextKeyFrame 次のencode結果を強制でkeyFrameにする
func (openh264Encoder *openh264Encoder) ForceNextKeyFrame() bool {
	return bool(C.Openh264Encoder_forceNextKeyFrame(unsafe.Pointer(openh264Encoder.cEncoder)))
}
