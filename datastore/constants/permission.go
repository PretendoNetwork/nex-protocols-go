package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Permission defines what users can perform access and update
// operations for an object. Access and update operations are
// stored separately
type Permission uint8

// WriteTo writes the Permission to the given writable
func (p Permission) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(p))
}

// ExtractFrom extracts the Permission value from the given readable
func (p *Permission) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*p = Permission(value)
	return nil
}

// String returns a human-readable representation of the Permission.
func (p Permission) String() string {
	switch p {
	case PermissionPublic:
		return "Public"
	case PermissionFriend:
		return "Friend"
	case PermissionSpecified:
		return "Specified"
	case PermissionPrivate:
		return "Private"
	case PermissionSpecifiedFriend:
		return "SpecifiedFriend"
	default:
		return fmt.Sprintf("Permission(%d)", int(p))
	}
}

const (
	// PermissionPublic means that anyone is allowed
	// to perform the operation
	PermissionPublic Permission = iota

	// PermissionFriend means that only the owner and
	// their friends are allowed to perform the operation
	PermissionFriend

	// PermissionSpecified means that only the owner and
	// the users whose PIDs are specified are allowed to
	// perform the operation
	PermissionSpecified

	// PermissionPrivate means that only the owner is allowed
	// to perform the operation
	PermissionPrivate

	// PermissionSpecifiedFriend means that only the owner and
	// friends of the owner whose PIDs are specified are allowed
	// to perform the operation
	PermissionSpecifiedFriend
)
