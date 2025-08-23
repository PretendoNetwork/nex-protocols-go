package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// ModificationFlag indicates what fields of an object have been modified
type ModificationFlag uint16

const (
	// ModificationFlagNone means that nothing has changed
	ModificationFlagNone ModificationFlag = 0x0

	// ModificationFlagName means that the object name has changed
	ModificationFlagName ModificationFlag = 0x1

	// ModificationFlagAccessPermission means that the object access permission has changed
	ModificationFlagAccessPermission ModificationFlag = 0x2

	// ModificationFlagUpdatePermission means that the object update permission has changed
	ModificationFlagUpdatePermission ModificationFlag = 0x4

	// ModificationFlagPeriod means that the object data expiration period has changed
	ModificationFlagPeriod ModificationFlag = 0x8

	// ModificationFlagMetaBinary means that the object MetaBinary has changed
	ModificationFlagMetaBinary ModificationFlag = 0x10

	// ModificationFlagTags means that the object tags have changed.
	// Tags are replaced, not appended
	ModificationFlagTags ModificationFlag = 0x20

	// ModificationFlagUpdatedTime means that the object itself has been updated.
	// This updates the object "updated" timestamp and refreshes the expiration date
	ModificationFlagUpdatedTime ModificationFlag = 0x40

	// ModificationFlagDataType means that the object data type has changed
	ModificationFlagDataType ModificationFlag = 0x80

	// ModificationFlagReferredCount means that the object referred count has changed
	ModificationFlagReferredCount ModificationFlag = 0x100

	// ModificationFlagStatus means that the object status has changed
	ModificationFlagStatus ModificationFlag = 0x200
)

// WriteTo writes the ModificationFlag to the given writable
func (mf ModificationFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt16LE(uint16(mf))
}

// ExtractFrom extracts the ModificationFlag value from the given readable
func (mf *ModificationFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt16LE()
	if err != nil {
		return err
	}

	*mf = ModificationFlag(value)
	return nil
}
