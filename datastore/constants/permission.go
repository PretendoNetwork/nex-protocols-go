package constants

// Permission defines what users can perform access and update
// operations for an object. Access and update operations are
// stored separately
type Permission uint8

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
