package ttLibGo

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
extern bool FdkaacEncoder_setBitrate(void *ptr, uint32_t value);
*/
import "C"
import (
	"unsafe"
)

type fdkaacEncoder encoder

// FdkAacTypes fdkaacEncoderで利用できる動作タイプ
var FdkAacTypes = struct {
	AotNone          subType
	AotNullObject    subType
	AotAacMAIN       subType
	AotAacLC         subType
	AotAacSSR        subType
	AotAacLTP        subType
	AotSBR           subType
	AotAacSCAL       subType
	AotTwinVQ        subType
	AotCELP          subType
	AotHVXC          subType
	AotRsvd10        subType
	AotRsvd11        subType
	AotTTSI          subType
	AotMainSYNTH     subType
	AotWavTabSYNTH   subType
	AotGenMIDI       subType
	AotAlgSynthAudFX subType
	AotErAacLC       subType
	AotRsvd18        subType
	AotErAacLTP      subType
	AotErAacSCAL     subType
	AotErTwinVQ      subType
	AotErBSAC        subType
	AotErAacLD       subType
	AotErCELP        subType
	AotErHVXC        subType
	AotErHILN        subType
	AotErPARA        subType
	AotRsvd28        subType
	AotPS            subType
	AotMPEGS         subType
	AotESCAPE        subType
	AotMp3onMp4L1    subType
	AotMp3onMp4L2    subType
	AotMp3onMp4L3    subType
	AotRsvd35        subType
	AotRsvd36        subType
	AotAacSLS        subType
	AotSLS           subType
	AotErAacELD      subType
	AotUSAC          subType
	AotSAOC          subType
	AotLdMPEGS       subType
	AotDrmAAC        subType
	AotDrmSBR        subType
	AotDrmMpegPS     subType
}{
	AotNone:          subType{"AOT_NONE"},
	AotNullObject:    subType{"AOT_NULL_OBJECT"},
	AotAacMAIN:       subType{"AOT_AAC_MAIN"},
	AotAacLC:         subType{"AOT_AAC_LC"},
	AotAacSSR:        subType{"AOT_AAC_SSR"},
	AotAacLTP:        subType{"AOT_AAC_LTP"},
	AotSBR:           subType{"AOT_SBR"},
	AotAacSCAL:       subType{"AOT_AAC_SCAL"},
	AotTwinVQ:        subType{"AOT_TWIN_VQ"},
	AotCELP:          subType{"AOT_CELP"},
	AotHVXC:          subType{"AOT_HVXC"},
	AotRsvd10:        subType{"AOT_RSVD_10"},
	AotRsvd11:        subType{"AOT_RSVD_11"},
	AotTTSI:          subType{"AOT_TTSI"},
	AotMainSYNTH:     subType{"AOT_MAIN_SYNTH"},
	AotWavTabSYNTH:   subType{"AOT_WAV_TAB_SYNTH"},
	AotGenMIDI:       subType{"AOT_GEN_MIDI"},
	AotAlgSynthAudFX: subType{"AOT_ALG_SYNTH_AUD_FX"},
	AotErAacLC:       subType{"AOT_ER_AAC_LC"},
	AotRsvd18:        subType{"AOT_RSVD_18"},
	AotErAacLTP:      subType{"AOT_ER_AAC_LTP"},
	AotErAacSCAL:     subType{"AOT_ER_AAC_SCAL"},
	AotErTwinVQ:      subType{"AOT_ER_TWIN_VQ"},
	AotErBSAC:        subType{"AOT_ER_BSAC"},
	AotErAacLD:       subType{"AOT_ER_AAC_LD"},
	AotErCELP:        subType{"AOT_ER_CELP"},
	AotErHVXC:        subType{"AOT_ER_HVXC"},
	AotErHILN:        subType{"AOT_ER_HILN"},
	AotErPARA:        subType{"AOT_ER_PARA"},
	AotRsvd28:        subType{"AOT_RSVD_28"},
	AotPS:            subType{"AOT_PS"},
	AotMPEGS:         subType{"AOT_MPEGS"},
	AotESCAPE:        subType{"AOT_ESCAPE"},
	AotMp3onMp4L1:    subType{"AOT_MP3ONMP4_L1"},
	AotMp3onMp4L2:    subType{"AOT_MP3ONMP4_L2"},
	AotMp3onMp4L3:    subType{"AOT_MP3ONMP4_L3"},
	AotRsvd35:        subType{"AOT_RSVD_35"},
	AotRsvd36:        subType{"AOT_RSVD_36"},
	AotAacSLS:        subType{"AOT_AAC_SLS"},
	AotSLS:           subType{"AOT_SLS"},
	AotErAacELD:      subType{"AOT_ER_AAC_ELD"},
	AotUSAC:          subType{"AOT_USAC"},
	AotSAOC:          subType{"AOT_SAOC"},
	AotLdMPEGS:       subType{"AOT_LD_MPEGS"},
	AotDrmAAC:        subType{"AOT_DRM_AAC"},
	AotDrmSBR:        subType{"AOT_DRM_SBR"},
	AotDrmMpegPS:     subType{"AOT_DRM_MPEG_PSs"},
}

// EncodeFrame エンコードを実行
func (fdkaacEncoder *fdkaacEncoder) EncodeFrame(
	frame IFrame,
	callback func(frame *Frame) bool) bool {
	return encoderEncodeFrame((*encoder)(fdkaacEncoder), frame, callback)
}

// Close 閉じる
func (fdkaacEncoder *fdkaacEncoder) Close() {
	encoderClose((*encoder)(fdkaacEncoder))
}

// SetBitrate fdkaacのencode結果のbitrateを設定する bps
func (fdkaacEncoder *fdkaacEncoder) SetBitrate(value uint32) bool {
	return bool(C.FdkaacEncoder_setBitrate(unsafe.Pointer(fdkaacEncoder.cEncoder), C.uint32_t(value)))
}
