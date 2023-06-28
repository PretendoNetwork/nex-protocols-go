package friends_wiiu_types

import "github.com/PretendoNetwork/nex-go"

// PrincipalRequestBlockSetting contains unknow data
type PrincipalRequestBlockSetting struct {
	nex.Structure
	PID       uint32
	IsBlocked bool
}

// Bytes encodes the PrincipalRequestBlockSetting and returns a byte array
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(principalRequestBlockSetting.PID)
	stream.WriteBool(principalRequestBlockSetting.IsBlocked)

	return stream.Bytes()
}

// Copy returns a new copied instance of PrincipalRequestBlockSetting
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Copy() nex.StructureInterface {
	copied := NewPrincipalRequestBlockSetting()

	copied.PID = principalRequestBlockSetting.PID
	copied.IsBlocked = principalRequestBlockSetting.IsBlocked

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (principalRequestBlockSetting *PrincipalRequestBlockSetting) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PrincipalRequestBlockSetting)

	if principalRequestBlockSetting.PID != other.PID {
		return false
	}

	if principalRequestBlockSetting.IsBlocked != other.IsBlocked {
		return false
	}

	return true
}

// NewPrincipalRequestBlockSetting returns a new PrincipalRequestBlockSetting
func NewPrincipalRequestBlockSetting() *PrincipalRequestBlockSetting {
	return &PrincipalRequestBlockSetting{}
}
