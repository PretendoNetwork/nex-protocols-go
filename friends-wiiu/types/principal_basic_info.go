package friends_wiiu_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// PrincipalBasicInfo contains user account and Mii data
type PrincipalBasicInfo struct {
	nex.Structure
	PID     uint32
	NNID    string
	Mii     *MiiV2
	Unknown uint8
}

// Bytes encodes the PrincipalBasicInfo and returns a byte array
func (principalInfo *PrincipalBasicInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(principalInfo.PID)
	stream.WriteString(principalInfo.NNID)
	stream.WriteStructure(principalInfo.Mii)
	stream.WriteUInt8(principalInfo.Unknown)

	return stream.Bytes()
}

// ExtractFromStream extracts a PrincipalBasicInfo structure from a stream
func (principalInfo *PrincipalBasicInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	principalInfo.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.PID. %s", err.Error())
	}

	principalInfo.NNID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.NNID. %s", err.Error())
	}

	miiV2, err := stream.ReadStructure(NewMiiV2())
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Mii. %s", err.Error())
	}

	principalInfo.Mii = miiV2.(*MiiV2)
	principalInfo.Unknown, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract PrincipalBasicInfo.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of PrincipalBasicInfo
func (principalInfo *PrincipalBasicInfo) Copy() nex.StructureInterface {
	copied := NewPrincipalBasicInfo()

	copied.PID = principalInfo.PID
	copied.NNID = principalInfo.NNID
	copied.Mii = principalInfo.Mii.Copy().(*MiiV2)
	copied.Unknown = principalInfo.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalInfo *PrincipalBasicInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalBasicInfo)

	if principalInfo.PID != other.PID {
		return false
	}

	if principalInfo.NNID != other.NNID {
		return false
	}

	if !principalInfo.Mii.Equals(other.Mii) {
		return false
	}

	if principalInfo.Unknown != other.Unknown {
		return false
	}

	return true
}

// NewPrincipalBasicInfo returns a new PrincipalBasicInfo
func NewPrincipalBasicInfo() *PrincipalBasicInfo {
	return &PrincipalBasicInfo{}
}
