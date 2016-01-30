// bgp.go
package packet

import (
	"l3/bgp/config"
)

type AFI uint16
type SAFI uint8

const (
	AfiIP AFI = iota + 1
	AfiIP6
)

const (
	SafiUnicast SAFI = iota + 1
	SafiMulticast
)

var ProtocolFamilyMap = map[string]uint32{
	"ipv4-unicast":   GetProtocolFamily(AfiIP, SafiUnicast),
	"ipv6-unicast":   GetProtocolFamily(AfiIP6, SafiUnicast),
	"ipv4-multicast": GetProtocolFamily(AfiIP, SafiMulticast),
	"ipv6-multicast": GetProtocolFamily(AfiIP6, SafiMulticast),
}

func GetProtocolFromConfig(afiSafis *[]config.AfiSafiConfig) (map[uint32]bool, bool) {
	afiSafiMap := make(map[uint32]bool)
	rv := true
	for _, afiSafi := range *afiSafis {
		if afiSafiVal, ok := ProtocolFamilyMap[afiSafi.AfiSafiName]; ok {
			afiSafiMap[afiSafiVal] = true
		} else {
			rv = false
			break
		}
	}

	if len(afiSafiMap) == 0 {
		afiSafiMap[ProtocolFamilyMap["ipv4-unicast"]] = true
	}
	return afiSafiMap, rv
}

func GetProtocolFamily(afi AFI, safi SAFI) uint32 {
	return uint32(afi<<8) | uint32(safi)
}

func GetAfiSafi(protocolFamily uint32) (AFI, SAFI) {
	return AFI(protocolFamily >> 8), SAFI(protocolFamily & 0xFF)
}

func GetProtocolFromOpenMsg(openMsg *BGPOpen) map[uint32]bool {
	afiSafiMap := make(map[uint32]bool)
	for _, optParam := range openMsg.OptParams {
		if capabilities, ok := optParam.(*BGPOptParamCapability); ok {
			for _, capability := range capabilities.Value {
				if val, ok := capability.(*BGPCapMPExt); ok {
					afiSafiMap[GetProtocolFamily(val.AFI, val.SAFI)] = true
				}
			}
		}
	}

	return afiSafiMap
}