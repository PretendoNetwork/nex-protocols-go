// Package protocol implements the Notifications protocol
package protocol

// BuildNotificationType builds a combined type for NotificationEvents using a category and subtype
func BuildNotificationType(category, subtype *types.PrimitiveU32) *types.PrimitiveU32 {
	return (category * 1000) + subtype
}
