package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayRewardPlatform is a type within the Uplay Win protocol
type UplayRewardPlatform struct {
	types.Structure
	PlatformCode types.String
	Purchased    types.Bool
}

// ObjectID returns the object identifier of the type
func (urp UplayRewardPlatform) ObjectID() types.RVType {
	return urp.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (urp UplayRewardPlatform) DataObjectID() types.RVType {
	return types.NewString("UplayRewardPlatform")
}

// WriteTo writes the UplayRewardPlatform to the given writable
func (urp UplayRewardPlatform) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	urp.PlatformCode.WriteTo(contentWritable)
	urp.Purchased.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	urp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayRewardPlatform from the given readable
func (urp *UplayRewardPlatform) ExtractFrom(readable types.Readable) error {
	if err := urp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform Header. %s", err.Error())
	}

	if err := urp.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform.PlatformCode. %s", err.Error())
	}

	if err := urp.Purchased.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayRewardPlatform.Purchased. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayRewardPlatform
func (urp UplayRewardPlatform) Copy() types.RVType {
	copied := NewUplayRewardPlatform()

	copied.StructureVersion = urp.StructureVersion
	copied.PlatformCode = urp.PlatformCode
	copied.Purchased = urp.Purchased

	return copied
}

// Equals checks if the given UplayRewardPlatform contains the same data as the current UplayRewardPlatform
func (urp UplayRewardPlatform) Equals(o types.RVType) bool {
	if _, ok := o.(UplayRewardPlatform); !ok {
		return false
	}

	other := o.(UplayRewardPlatform)

	if urp.StructureVersion != other.StructureVersion {
		return false
	}

	if !urp.PlatformCode.Equals(other.PlatformCode) {
		return false
	}

	return urp.Purchased.Equals(other.Purchased)
}

// CopyRef copies the current value of the UplayRewardPlatform
// and returns a pointer to the new copy
func (urp UplayRewardPlatform) CopyRef() types.RVTypePtr {
	copied := urp.Copy().(UplayRewardPlatform)
	return &copied
}

// Deref takes a pointer to the UplayRewardPlatform
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (urp *UplayRewardPlatform) Deref() types.RVType {
	return *urp
}

// String returns the string representation of the UplayRewardPlatform
func (urp UplayRewardPlatform) String() string {
	return urp.FormatToString(0)
}

// FormatToString pretty-prints the UplayRewardPlatform using the provided indentation level
func (urp UplayRewardPlatform) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayRewardPlatform{\n")
	fmt.Fprintf(&b, "%sm_platformCode: %v,\n", indentationValues, urp.PlatformCode)
	fmt.Fprintf(&b, "%sm_purchased: %v,\n", indentationValues, urp.Purchased)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayRewardPlatform returns a new UplayRewardPlatform
func NewUplayRewardPlatform() UplayRewardPlatform {
	return UplayRewardPlatform{
		PlatformCode: "PC",
		Purchased:    false,
	}
}
