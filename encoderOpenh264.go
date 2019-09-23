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

type Openh264Data struct {
	key string
	val interface{}
}

type eoh264UsageType struct {
	CameraVideoRealTime    Openh264Data
	ScreenContentRealTime  Openh264Data
	CameraVideoNonRealTime Openh264Data
}

type eoh264RCModeType struct {
	RcQualityMode         Openh264Data
	RcBitrateMode         Openh264Data
	RcBufferbasedMode     Openh264Data
	RcTimestampMode       Openh264Data
	RcBitrateModePostSkip Openh264Data
	RcOffMode             Openh264Data
}

type eoh264ComplexityMode struct {
	LowComplexity    Openh264Data
	MediumComplexity Openh264Data
	HighComplexity   Openh264Data
}

type eoh264SpsPpsIDStrategy struct {
	ConstantID                 Openh264Data
	IncreasingID               Openh264Data
	SPSListing                 Openh264Data
	SPSListingAndPPSIncreasing Openh264Data
	SPSPPSListing              Openh264Data
}

// Openh264Params openh264Encoderで利用するParam設定項目
var Openh264Params = struct {
	UsageType                  eoh264UsageType
	ITargetBitrate             func(int) Openh264Data
	RCMode                     eoh264RCModeType
	FMaxFrameRate              func(float32) Openh264Data
	ITemporalLayerNum          func(int) Openh264Data
	ISpatialLayerNum           func(int) Openh264Data
	ComplexityMode             eoh264ComplexityMode
	UIIntraPeriod              func(uint) Openh264Data
	INumRefFrame               func(int) Openh264Data
	SpsPpsIDStrategy           eoh264SpsPpsIDStrategy
	BPrefixNalAddingCtrl       func(bool) Openh264Data
	BEnableSSEI                func(bool) Openh264Data
	BSimulcastAVC              func(bool) Openh264Data
	IPaddingFlag               func(int) Openh264Data
	IEntropyCodingModeFlag     func(int) Openh264Data
	BEnableFrameSkip           func(bool) Openh264Data
	IMaxBitrate                func(int) Openh264Data
	IMaxQp                     func(int) Openh264Data
	IMinQp                     func(int) Openh264Data
	UIMaxNalSize               func(uint) Openh264Data
	BEnableLongTermReference   func(bool) Openh264Data
	ILTRRefNum                 func(int) Openh264Data
	ILtrMarkPeriod             func(int) Openh264Data
	IMultipleThreadIdc         func(int) Openh264Data
	BUseLoadBalancing          func(bool) Openh264Data
	ILoopFilterDisableIdc      func(int) Openh264Data
	ILoopFilterAlphaC0Offset   func(int) Openh264Data
	ILoopFilterBetaOffset      func(int) Openh264Data
	BEnableDenoise             func(bool) Openh264Data
	BEnableBackgroundDetection func(bool) Openh264Data
	BEnableAdaptiveQuant       func(bool) Openh264Data
	BEnableFrameCroppingFlag   func(bool) Openh264Data
	BEnableSceneChangeDetect   func(bool) Openh264Data
	BIsLosslessLink            func(bool) Openh264Data
}{
	UsageType: eoh264UsageType{
		Openh264Data{key: "iUsageType", val: "CAMERA_VIDEO_REAL_TIME"},
		Openh264Data{key: "iUsageType", val: "SCREEN_CONTENT_REAL_TIME"},
		Openh264Data{key: "iUsageType", val: "CAMERA_VIDEO_NON_REAL_TIME"},
	},
	ITargetBitrate: func(value int) Openh264Data { return Openh264Data{key: "iTargetBitrate", val: value} },
	RCMode: eoh264RCModeType{
		Openh264Data{key: "iRCMode", val: "RC_QUALITY_MODE"},
		Openh264Data{key: "iRCMode", val: "RC_BITRATE_MODE"},
		Openh264Data{key: "iRCMode", val: "RC_BUFFERBASED_MODE"},
		Openh264Data{key: "iRCMode", val: "RC_TIMESTAMP_MODE"},
		Openh264Data{key: "iRCMode", val: "RC_BITRATE_MODE_POST_SKIP"},
		Openh264Data{key: "iRCMode", val: "RC_OFF_MODE"},
	},
	FMaxFrameRate:     func(value float32) Openh264Data { return Openh264Data{key: "fMaxFrameRate", val: value} },
	ITemporalLayerNum: func(value int) Openh264Data { return Openh264Data{key: "iTemporalLayerNum", val: value} },
	ISpatialLayerNum:  func(value int) Openh264Data { return Openh264Data{key: "iSpatialLayerNum", val: value} },
	ComplexityMode: eoh264ComplexityMode{
		Openh264Data{key: "iComplexityMode", val: "LOW_COMPLEXITY"},
		Openh264Data{key: "iComplexityMode", val: "MEDIUM_COMPLEXITY"},
		Openh264Data{key: "iComplexityMode", val: "HIGH_COMPLEXITY"},
	},
	UIIntraPeriod: func(value uint) Openh264Data { return Openh264Data{key: "uiIntraPeriod", val: value} },
	INumRefFrame:  func(value int) Openh264Data { return Openh264Data{key: "iNumRefFrames", val: value} },
	SpsPpsIDStrategy: eoh264SpsPpsIDStrategy{
		Openh264Data{key: "eSpsPpsIdStrategy", val: "CONSTANT_ID"},
		Openh264Data{key: "eSpsPpsIdStrategy", val: "INCREASING_ID"},
		Openh264Data{key: "eSpsPpsIdStrategy", val: "SPS_LISTING"},
		Openh264Data{key: "eSpsPpsIdStrategy", val: "SPS_LISTING_AND_PPS_INCREASING"},
		Openh264Data{key: "eSpsPpsIdStrategy", val: "SPS_PPS_LISTING"},
	},
	BPrefixNalAddingCtrl:       func(value bool) Openh264Data { return Openh264Data{key: "bPrefixNalAddingCtrl", val: value} },
	BEnableSSEI:                func(value bool) Openh264Data { return Openh264Data{key: "bEnableSSEI", val: value} },
	BSimulcastAVC:              func(value bool) Openh264Data { return Openh264Data{key: "bSimulcastAVC", val: value} },
	IPaddingFlag:               func(value int) Openh264Data { return Openh264Data{key: "iPaddingFlag", val: value} },
	IEntropyCodingModeFlag:     func(value int) Openh264Data { return Openh264Data{key: "iEntropyCodingModeFlag", val: value} },
	BEnableFrameSkip:           func(value bool) Openh264Data { return Openh264Data{key: "bEnableFrameSkip", val: value} },
	IMaxBitrate:                func(value int) Openh264Data { return Openh264Data{key: "iMaxBitrate", val: value} },
	IMaxQp:                     func(value int) Openh264Data { return Openh264Data{key: "iMaxQp", val: value} },
	IMinQp:                     func(value int) Openh264Data { return Openh264Data{key: "iMinQp", val: value} },
	UIMaxNalSize:               func(value uint) Openh264Data { return Openh264Data{key: "uiMaxNalSize", val: value} },
	BEnableLongTermReference:   func(value bool) Openh264Data { return Openh264Data{key: "bEnableLongTermReference", val: value} },
	ILTRRefNum:                 func(value int) Openh264Data { return Openh264Data{key: "iLTRRefNum", val: value} },
	ILtrMarkPeriod:             func(value int) Openh264Data { return Openh264Data{key: "iLtrMarkPeriod", val: value} },
	IMultipleThreadIdc:         func(value int) Openh264Data { return Openh264Data{key: "iMultipleThreadIdc", val: value} },
	BUseLoadBalancing:          func(value bool) Openh264Data { return Openh264Data{key: "bUseLoadBalancing", val: value} },
	ILoopFilterDisableIdc:      func(value int) Openh264Data { return Openh264Data{key: "iLoopFilterDisableIdc", val: value} },
	ILoopFilterAlphaC0Offset:   func(value int) Openh264Data { return Openh264Data{key: "iLoopFilterAlphaC0Offset", val: value} },
	ILoopFilterBetaOffset:      func(value int) Openh264Data { return Openh264Data{key: "iLoopFilterBetaOffset", val: value} },
	BEnableDenoise:             func(value bool) Openh264Data { return Openh264Data{key: "bEnableDenoise", val: value} },
	BEnableBackgroundDetection: func(value bool) Openh264Data { return Openh264Data{key: "bEnableBackgroundDetection", val: value} },
	BEnableAdaptiveQuant:       func(value bool) Openh264Data { return Openh264Data{key: "bEnableAdaptiveQuant", val: value} },
	BEnableFrameCroppingFlag:   func(value bool) Openh264Data { return Openh264Data{key: "bEnableFrameCroppingFlag", val: value} },
	BEnableSceneChangeDetect:   func(value bool) Openh264Data { return Openh264Data{key: "bEnableSceneChangeDetect", val: value} },
	BIsLosslessLink:            func(value bool) Openh264Data { return Openh264Data{key: "bIsLosslessLink", val: value} },
}

type eoh264ProfileIdc struct {
	ProUNKNOWN          Openh264Data
	ProBASELINE         Openh264Data
	ProMAIN             Openh264Data
	ProEXTENDED         Openh264Data
	ProHIGH             Openh264Data
	ProHIGH10           Openh264Data
	ProHIGH422          Openh264Data
	ProHIGH444          Openh264Data
	ProCAVLC444         Openh264Data
	ProScalableBASELINE Openh264Data
	ProScalableHIGH     Openh264Data
}

type eoh264LevelIdc struct {
	LevelUNKNOWN Openh264Data
	Level1_0     Openh264Data
	Level1_B     Openh264Data
	Level1_1     Openh264Data
	Level1_2     Openh264Data
	Level1_3     Openh264Data
	Level2_0     Openh264Data
	Level2_1     Openh264Data
	Level2_2     Openh264Data
	Level3_0     Openh264Data
	Level3_1     Openh264Data
	Level3_2     Openh264Data
	Level4_0     Openh264Data
	Level4_1     Openh264Data
	Level4_2     Openh264Data
	Level5_0     Openh264Data
	Level5_1     Openh264Data
	Level5_2     Openh264Data
}

// 1.5.0以下用
type eoh264SliceCfgSliceMode struct {
	SmSingleSLICE      Openh264Data
	SmFixedslcnumSLICE Openh264Data
	SmRasterSLICE      Openh264Data
	SmRowmbSLICE       Openh264Data
	SmDynSLICE         Openh264Data
	SmAutoSLICE        Openh264Data
	SmRESERVED         Openh264Data
}

type eoh264SliceCfgSliceArgument struct {
	UISliceNum            func(uint) Openh264Data
	UISliceSizeConstraint func(uint) Openh264Data
}

type eoh264SliceCfg struct {
	SliceMode     eoh264SliceCfgSliceMode
	SliceArgument eoh264SliceCfgSliceArgument
}

// 1.5.0以下用ここまで

var Openh264SpatialParams_Before15 = struct {
	IVideoWidth        func(int) Openh264Data
	IVideoHeight       func(int) Openh264Data
	FFrameRate         func(float32) Openh264Data
	ISpatialBitrate    func(int) Openh264Data
	IMaxSpatialBitrate func(int) Openh264Data
	ProfileIdc         eoh264ProfileIdc
	LevelIdc           eoh264LevelIdc
	IDLayerQp          func(int) Openh264Data
	SliceCfg           eoh264SliceCfg
}{
	IVideoWidth:        func(value int) Openh264Data { return Openh264Data{key: "iVideoWidth", val: value} },
	IVideoHeight:       func(value int) Openh264Data { return Openh264Data{key: "iVideoHeight", val: value} },
	FFrameRate:         func(value float32) Openh264Data { return Openh264Data{key: "fFrameRate", val: value} },
	ISpatialBitrate:    func(value int) Openh264Data { return Openh264Data{key: "iSpatialBitrate", val: value} },
	IMaxSpatialBitrate: func(value int) Openh264Data { return Openh264Data{key: "iMaxSpatialBitrate", val: value} },
	ProfileIdc: eoh264ProfileIdc{
		Openh264Data{key: "uiProfileIdc", val: "PRO_UNKNOWN"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_BASELINE"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_MAIN"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_EXTENDED"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH10"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH422"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH444"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_CAVLC444"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_SCALABLE_BASELINE"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_SCALABLE_HIGH"},
	},
	LevelIdc: eoh264LevelIdc{
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_UNKNOWN"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_B"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_3"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_2"},
	},
	IDLayerQp: func(value int) Openh264Data { return Openh264Data{key: "iDLayerQp", val: value} },
	SliceCfg: eoh264SliceCfg{
		SliceMode: eoh264SliceCfgSliceMode{
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_SINGLE_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_FIXEDSLCNUM_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_RASTER_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_ROWMB_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_DYN_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_AUTO_SLICE"},
			Openh264Data{key: "sSliceCfg.uiSliceMode", val: "SM_RESERVEDs"},
		},
		SliceArgument: eoh264SliceCfgSliceArgument{
			UISliceNum: func(value uint) Openh264Data {
				return Openh264Data{key: "sSliceCfg.sSliceArgument.uiSliceNum", val: value}
			},
			UISliceSizeConstraint: func(value uint) Openh264Data {
				return Openh264Data{key: "sSliceCfg.sSliceArgument.uiSliceSizeConstraint", val: value}
			},
		},
	},
}

type eoh264SliceArgumentSliceMode struct {
	SmSingleSLICE      Openh264Data
	SmFixedslcnumSLICE Openh264Data
	SmRasterSLICE      Openh264Data
	SmSizelimitedSLICE Openh264Data
	SmRESERVED         Openh264Data
}

type eoh264SliceArgument struct {
	SliceMode             eoh264SliceArgumentSliceMode
	UISliceNum            func(uint) Openh264Data
	UISliceSizeConstraint func(uint) Openh264Data
}

type eoh264VideoFormat struct {
	VfCOMPONENT Openh264Data
	VfPAL       Openh264Data
	VfNTSC      Openh264Data
	VfSECAM     Openh264Data
	VfMAC       Openh264Data
	VfUNDEF     Openh264Data
	VfNumENUM   Openh264Data
}

type eoh264ColorPrimaries struct {
	CpRESERVED0 Openh264Data
	CpBT709     Openh264Data
	CpUNDEF     Openh264Data
	CpRESERVED3 Openh264Data
	CpBT470M    Openh264Data
	CpBT470BG   Openh264Data
	CpSMPTE170M Openh264Data
	CpSMPTE240M Openh264Data
	CpFILM      Openh264Data
	CpBT2020    Openh264Data
	CpNumENUM   Openh264Data
}

type eoh264TransferCharacteristics struct {
	TrcRESERVED0    Openh264Data
	TrcBT709        Openh264Data
	TrcUNDEF        Openh264Data
	TrcRESERVED3    Openh264Data
	TrcBT470M       Openh264Data
	TrcBT470BG      Openh264Data
	TrcSMPTE170M    Openh264Data
	TrcSMPTE240M    Openh264Data
	TrcLINEAR       Openh264Data
	TrcLOG100       Openh264Data
	TrcLOG316       Openh264Data
	TrcIEC61966_2_4 Openh264Data
	TrcBT1361E      Openh264Data
	TrcIEC61966_2_1 Openh264Data
	TrcBT2020_10    Openh264Data
	TrcBT2020_12    Openh264Data
	TrcNumENUM      Openh264Data
}

type eoh264ColorMatrix struct {
	CmGBR       Openh264Data
	CmBT709     Openh264Data
	CmUNDEF     Openh264Data
	CmRESERVED3 Openh264Data
	CmFCC       Openh264Data
	CmBT470BG   Openh264Data
	CmSMPTE170M Openh264Data
	CmSMPTE240M Openh264Data
	CmYCGCO     Openh264Data
	CmBT2020NC  Openh264Data
	CmBT2020C   Openh264Data
	CmNumENUM   Openh264Data
}

var Openh264SpatialParams = struct {
	IVideoWidth              func(int) Openh264Data
	IVideoHeight             func(int) Openh264Data
	FFrameRate               func(float32) Openh264Data
	ISpatialBitrate          func(int) Openh264Data
	IMaxSpatialBitrate       func(int) Openh264Data
	ProfileIdc               eoh264ProfileIdc
	LevelIdc                 eoh264LevelIdc
	IDLayerQp                func(int) Openh264Data
	SliceArgument            eoh264SliceArgument
	BVideoSignalTypePresent  func(bool) Openh264Data
	VideoFormat              eoh264VideoFormat
	BFullRange               func(bool) Openh264Data
	BColorDescriptionPresent func(bool) Openh264Data
	ColorPrimaries           eoh264ColorPrimaries
	TransferCharacteristics  eoh264TransferCharacteristics
	ColorMatrix              eoh264ColorMatrix
}{
	IVideoWidth:        func(value int) Openh264Data { return Openh264Data{key: "iVideoWidth", val: value} },
	IVideoHeight:       func(value int) Openh264Data { return Openh264Data{key: "iVideoHeight", val: value} },
	FFrameRate:         func(value float32) Openh264Data { return Openh264Data{key: "fFrameRate", val: value} },
	ISpatialBitrate:    func(value int) Openh264Data { return Openh264Data{key: "iSpatialBitrate", val: value} },
	IMaxSpatialBitrate: func(value int) Openh264Data { return Openh264Data{key: "iMaxSpatialBitrate", val: value} },
	ProfileIdc: eoh264ProfileIdc{
		Openh264Data{key: "uiProfileIdc", val: "PRO_UNKNOWN"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_BASELINE"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_MAIN"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_EXTENDED"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH10"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH422"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_HIGH444"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_CAVLC444"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_SCALABLE_BASELINE"},
		Openh264Data{key: "uiProfileIdc", val: "PRO_SCALABLE_HIGH"},
	},
	LevelIdc: eoh264LevelIdc{
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_UNKNOWN"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_B"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_1_3"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_2_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_3_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_4_2"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_0"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_1"},
		Openh264Data{key: "uiLevelIdc", val: "LEVEL_5_2"},
	},
	IDLayerQp: func(value int) Openh264Data { return Openh264Data{key: "iDLayerQp", val: value} },
	SliceArgument: eoh264SliceArgument{
		SliceMode: eoh264SliceArgumentSliceMode{
			SmSingleSLICE:      Openh264Data{key: "sSliceArgument.uiSliceMode", val: "SM_SINGLE_SLICE"},
			SmFixedslcnumSLICE: Openh264Data{key: "sSliceArgument.uiSliceMode", val: "SM_FIXEDSLCNUM_SLICE"},
			SmRasterSLICE:      Openh264Data{key: "sSliceArgument.uiSliceMode", val: "SM_RASTER_SLICE"},
			SmSizelimitedSLICE: Openh264Data{key: "sSliceArgument.uiSliceMode", val: "SM_SIZELIMITED_SLICE"},
			SmRESERVED:         Openh264Data{key: "sSliceArgument.uiSliceMode", val: "SM_RESERVED"},
		},
		UISliceNum: func(value uint) Openh264Data { return Openh264Data{key: "sSliceArgument.uiSliceNum", val: value} },
		UISliceSizeConstraint: func(value uint) Openh264Data {
			return Openh264Data{key: "sSliceArgument.uiSliceSizeConstraint", val: value}
		},
	},
	BVideoSignalTypePresent: func(value bool) Openh264Data { return Openh264Data{key: "bVideoSignalTypePresent", val: value} },
	VideoFormat: eoh264VideoFormat{
		VfCOMPONENT: Openh264Data{key: "uiVideoFormat", val: "VF_COMPONENT"},
		VfPAL:       Openh264Data{key: "uiVideoFormat", val: "VF_PAL"},
		VfNTSC:      Openh264Data{key: "uiVideoFormat", val: "VF_NTSC"},
		VfSECAM:     Openh264Data{key: "uiVideoFormat", val: "VF_SECAM"},
		VfMAC:       Openh264Data{key: "uiVideoFormat", val: "VF_MAC"},
		VfUNDEF:     Openh264Data{key: "uiVideoFormat", val: "VF_UNDEF"},
		VfNumENUM:   Openh264Data{key: "uiVideoFormat", val: "VF_NUM_ENUM"},
	},
	BFullRange:               func(value bool) Openh264Data { return Openh264Data{key: "bFullRange", val: value} },
	BColorDescriptionPresent: func(value bool) Openh264Data { return Openh264Data{key: "bColorDescriptionPresent", val: value} },
	ColorPrimaries: eoh264ColorPrimaries{
		CpRESERVED0: Openh264Data{key: "uiColorPrimaries", val: "CP_RESERVED0"},
		CpBT709:     Openh264Data{key: "uiColorPrimaries", val: "CP_BT709"},
		CpUNDEF:     Openh264Data{key: "uiColorPrimaries", val: "CP_UNDEF"},
		CpRESERVED3: Openh264Data{key: "uiColorPrimaries", val: "CP_RESERVED3"},
		CpBT470M:    Openh264Data{key: "uiColorPrimaries", val: "CP_BT470M"},
		CpBT470BG:   Openh264Data{key: "uiColorPrimaries", val: "CP_BT470BG"},
		CpSMPTE170M: Openh264Data{key: "uiColorPrimaries", val: "CP_SMPTE170M"},
		CpSMPTE240M: Openh264Data{key: "uiColorPrimaries", val: "CP_SMPTE240M"},
		CpFILM:      Openh264Data{key: "uiColorPrimaries", val: "CP_FILM"},
		CpBT2020:    Openh264Data{key: "uiColorPrimaries", val: "CP_BT2020"},
		CpNumENUM:   Openh264Data{key: "uiColorPrimaries", val: "CP_NUM_ENUM"},
	},
	TransferCharacteristics: eoh264TransferCharacteristics{
		TrcRESERVED0:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_RESERVED0"},
		TrcBT709:        Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT709"},
		TrcUNDEF:        Openh264Data{key: "uiTransferCharacteristics", val: "TRC_UNDEF"},
		TrcRESERVED3:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_RESERVED3"},
		TrcBT470M:       Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT470M"},
		TrcBT470BG:      Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT470BG"},
		TrcSMPTE170M:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_SMPTE170M"},
		TrcSMPTE240M:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_SMPTE240M"},
		TrcLINEAR:       Openh264Data{key: "uiTransferCharacteristics", val: "TRC_LINEAR"},
		TrcLOG100:       Openh264Data{key: "uiTransferCharacteristics", val: "TRC_LOG100"},
		TrcLOG316:       Openh264Data{key: "uiTransferCharacteristics", val: "TRC_LOG316"},
		TrcIEC61966_2_4: Openh264Data{key: "uiTransferCharacteristics", val: "TRC_IEC61966_2_4"},
		TrcBT1361E:      Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT1361E"},
		TrcIEC61966_2_1: Openh264Data{key: "uiTransferCharacteristics", val: "TRC_IEC61966_2_1"},
		TrcBT2020_10:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT2020_10"},
		TrcBT2020_12:    Openh264Data{key: "uiTransferCharacteristics", val: "TRC_BT2020_12"},
		TrcNumENUM:      Openh264Data{key: "uiTransferCharacteristics", val: "TRC_NUM_ENUM"},
	},
	ColorMatrix: eoh264ColorMatrix{
		CmGBR:       Openh264Data{key: "uiColorMatrix", val: "CM_GBR"},
		CmBT709:     Openh264Data{key: "uiColorMatrix", val: "CM_BT709"},
		CmUNDEF:     Openh264Data{key: "uiColorMatrix", val: "CM_UNDEF"},
		CmRESERVED3: Openh264Data{key: "uiColorMatrix", val: "CM_RESERVED3"},
		CmFCC:       Openh264Data{key: "uiColorMatrix", val: "CM_FCC"},
		CmBT470BG:   Openh264Data{key: "uiColorMatrix", val: "CM_BT470BG"},
		CmSMPTE170M: Openh264Data{key: "uiColorMatrix", val: "CM_SMPTE170M"},
		CmSMPTE240M: Openh264Data{key: "uiColorMatrix", val: "CM_SMPTE240M"},
		CmYCGCO:     Openh264Data{key: "uiColorMatrix", val: "CM_YCGCO"},
		CmBT2020NC:  Openh264Data{key: "uiColorMatrix", val: "CM_BT2020NC"},
		CmBT2020C:   Openh264Data{key: "uiColorMatrix", val: "CM_BT2020C"},
		CmNumENUM:   Openh264Data{key: "uiColorMatrix", val: "CM_NUM_ENUM"},
	},
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
	callback func(frame *Frame) bool) bool {
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
