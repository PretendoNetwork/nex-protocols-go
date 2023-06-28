package friends_3ds_types

import "github.com/PretendoNetwork/nex-go"

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	nex.Structure
	PID              uint32
	LFC              uint64
	RelationshipType uint8
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PID)
	stream.WriteUInt64LE(relationship.LFC)
	stream.WriteUInt8(relationship.RelationshipType)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRelationship
func (relationship *FriendRelationship) Copy() nex.StructureInterface {
	copied := NewFriendRelationship()

	copied.PID = relationship.PID
	copied.LFC = relationship.LFC
	copied.RelationshipType = relationship.RelationshipType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (relationship *FriendRelationship) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRelationship)

	if relationship.PID != other.PID {
		return false
	}

	if relationship.LFC != other.LFC {
		return false
	}

	if relationship.RelationshipType != other.RelationshipType {
		return false
	}

	return true
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}
