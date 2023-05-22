package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
)

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	nex.Structure
	NexUniqueID         uint64
	NexUniqueIDPassword uint64
}

// Bytes encodes the UniqueIDInfo and returns a byte array
func (uniqueIDInfo *UniqueIDInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueID)
	stream.WriteUInt64LE(uniqueIDInfo.NexUniqueIDPassword)

	return stream.Bytes()
}

// ExtractFromStream extracts a UniqueIDInfo structure from a stream
func (uniqueIDInfo *UniqueIDInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	uniqueIDInfo.NexUniqueID = stream.ReadUInt64LE()
	uniqueIDInfo.NexUniqueIDPassword = stream.ReadUInt64LE()

	if err != nil {
		return err
	}

	return nil
}

// Copy returns a new copied instance of UniqueIDInfo
func (uniqueIDInfo *UniqueIDInfo) Copy() nex.StructureInterface {
	copied := NewUniqueIDInfo()

	copied.NexUniqueID = uniqueIDInfo.NexUniqueID
	copied.NexUniqueIDPassword = uniqueIDInfo.NexUniqueIDPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (uniqueIDInfo *UniqueIDInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*UniqueIDInfo)

	if uniqueIDInfo.NexUniqueID != other.NexUniqueID {
		return false
	}

	if uniqueIDInfo.NexUniqueIDPassword != other.NexUniqueIDPassword {
		return false
	}

	return true
}

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	return &UniqueIDInfo{}
}
