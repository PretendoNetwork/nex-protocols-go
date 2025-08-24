package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// NotificationCategorySigned represents a signed NotificationCategory.
//
// Note: This is not a real type. This is a bespoke type created solely for
// use with MatchmakeExtension::GetFriendNotificationData.
type NotificationCategorySigned int32

// WriteTo writes the NotificationCategorySigned to the given writable
func (nc NotificationCategorySigned) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(nc))
}

// ExtractFrom extracts the NotificationCategorySigned value from the given readable
func (nc *NotificationCategorySigned) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*nc = NotificationCategorySigned(value)
	return nil
}

// ToUnsigned converts a NotificationCategorySigned to a NotificationCategory
func (ncs NotificationCategorySigned) ToUnsigned() NotificationCategory {
	return NotificationCategory(ncs)
}
