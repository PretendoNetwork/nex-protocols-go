package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// RelationshipType defines the status of a 3DS friend relationship
type RelationshipType uint8

// WriteTo writes the RelationshipType to the given writable
func (rt RelationshipType) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(rt))
}

// ExtractFrom extracts the RelationshipType value from the given readable
func (rt *RelationshipType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*rt = RelationshipType(value)
	return nil
}

const (
	// RelationshipTypeIncomplete means that the relationship is provisional (one party has added
	// the other, but the other has not added them back)
	RelationshipTypeIncomplete RelationshipType = iota

	// RelationshipTypeComplete means that both users have added each other as friends
	RelationshipTypeComplete

	// RelationshipTypeInvalid means that there is no relationship between the users
	RelationshipTypeInvalid
)
