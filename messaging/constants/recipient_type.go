package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// RecipientType determines what to target when sending messages.
type RecipientType uint32

// WriteTo writes the RecipientType to the given writable
func (rt RecipientType) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(rt))
}

// ExtractFrom extracts the RecipientType value from the given readable
func (rt *RecipientType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*rt = RecipientType(value)
	return nil
}

const (
	// RecipientTypePrincipalID means the message is being sent to
	// a specific player. The recipient ID will be the destination
	// users PID.
	RecipientTypePrincipalID RecipientType = iota + 1

	// RecipientTypeGatheringID means the message is being sent to
	// a gathering. The recipient ID will be the gatehering ID. The
	// message is sent to all participants.
	RecipientTypeGatheringID
)
