package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// NotificationEvents represents the fully built notification type with it's sub-type.
//
// Note: This is not accurate to the true NotificationEvents type. The true
// NotificationEvents type represents the base category names, also called the "main type".
// This was renamed to NotificationCategory, and now NotificationEvents represents the
// fully built type. For categories, see NotificationCategory.
type NotificationEvents uint32

// WriteTo writes the NotificationEvents to the given writable
func (ne NotificationEvents) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(ne))
}

// ExtractFrom extracts the NotificationEvents value from the given readable
func (ne *NotificationEvents) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*ne = NotificationEvents(value)
	return nil
}
