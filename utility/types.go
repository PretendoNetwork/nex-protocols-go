package utility

import nex "github.com/PretendoNetwork/nex-go"

// UniqueIDInfo holds parameters for a matchmake session
type UniqueIDInfo struct {
	NexUniqueID         uint64
	NexUniqueIDPassword uint64

	*nex.Structure
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

// NewUniqueIDInfo returns a new UniqueIDInfo
func NewUniqueIDInfo() *UniqueIDInfo {
	return &UniqueIDInfo{}
}
