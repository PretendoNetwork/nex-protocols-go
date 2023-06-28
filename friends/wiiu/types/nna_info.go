package friends_wiiu_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// NNAInfo contains information about a Nintendo Network Account
type NNAInfo struct {
	nex.Structure
	PrincipalBasicInfo *PrincipalBasicInfo
	Unknown1           uint8
	Unknown2           uint8
}

// Bytes encodes the NNAInfo and returns a byte array
func (nnaInfo *NNAInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(nnaInfo.PrincipalBasicInfo)
	stream.WriteUInt8(nnaInfo.Unknown1)
	stream.WriteUInt8(nnaInfo.Unknown2)

	return stream.Bytes()
}

// ExtractFromStream extracts a NNAInfo structure from a stream
func (nnaInfo *NNAInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalBasicInfo, err := stream.ReadStructure(NewPrincipalBasicInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.PrincipalBasicInfo. %s", err.Error())
	}

	nnaInfo.PrincipalBasicInfo = principalBasicInfo.(*PrincipalBasicInfo)
	nnaInfo.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown1. %s", err.Error())
	}

	nnaInfo.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NNAInfo.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NNAInfo
func (nnaInfo *NNAInfo) Copy() nex.StructureInterface {
	copied := NewNNAInfo()

	copied.PrincipalBasicInfo = nnaInfo.PrincipalBasicInfo.Copy().(*PrincipalBasicInfo)
	copied.Unknown1 = nnaInfo.Unknown1
	copied.Unknown2 = nnaInfo.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nnaInfo *NNAInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NNAInfo)

	if !nnaInfo.PrincipalBasicInfo.Equals(other.PrincipalBasicInfo) {
		return false
	}

	if nnaInfo.Unknown1 != other.Unknown1 {
		return false
	}

	if nnaInfo.Unknown2 != other.Unknown2 {
		return false
	}

	return true
}

// NewNNAInfo returns a new NNAInfo
func NewNNAInfo() *NNAInfo {
	return &NNAInfo{}
}
