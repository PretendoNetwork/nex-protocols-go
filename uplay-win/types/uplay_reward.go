package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayReward is a type within the Uplay Win protocol
type UplayReward struct {
	types.Structure
	m_code           types.String
	m_name           types.String
	m_description    types.String
	m_value          types.Int32
	m_rewardTypeName types.String
	m_gameCode       types.String
	m_platforms      types.List[UplayRewardPlatform]
}

// ObjectID returns the object identifier of the type
func (upr UplayReward) ObjectID() types.RVType {
	return upr.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (upr UplayReward) DataObjectID() types.RVType {
	return types.NewString("UplayReward")
}

// WriteTo writes the UplayReward to the given writable
func (upr UplayReward) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	upr.m_code.WriteTo(contentWritable)
	upr.m_name.WriteTo(contentWritable)
	upr.m_description.WriteTo(contentWritable)
	upr.m_value.WriteTo(contentWritable)
	upr.m_rewardTypeName.WriteTo(contentWritable)
	upr.m_gameCode.WriteTo(contentWritable)
	upr.m_platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	upr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayReward from the given readable
func (upr *UplayReward) ExtractFrom(readable types.Readable) error {
	if err := upr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward Header. %s", err.Error())
	}

	if err := upr.m_code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_code. %s", err.Error())
	}

	if err := upr.m_name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_name. %s", err.Error())
	}

	if err := upr.m_description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_description. %s", err.Error())
	}

	if err := upr.m_value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_value. %s", err.Error())
	}

	if err := upr.m_rewardTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_rewardTypeName. %s", err.Error())
	}

	if err := upr.m_gameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_gameCode. %s", err.Error())
	}

	if err := upr.m_platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.m_platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayReward
func (upr UplayReward) Copy() types.RVType {
	copied := NewUplayReward()

	copied.StructureVersion = upr.StructureVersion
	copied.m_code = upr.m_code
	copied.m_name = upr.m_name
	copied.m_description = upr.m_description
	copied.m_value = upr.m_value
	copied.m_rewardTypeName = upr.m_rewardTypeName
	copied.m_gameCode = upr.m_gameCode
	copied.m_platforms = upr.m_platforms

	return copied
}

// Equals checks if the given UplayReward contains the same data as the current UplayReward
func (upr UplayReward) Equals(o types.RVType) bool {
	if _, ok := o.(UplayReward); !ok {
		return false
	}

	other := o.(UplayReward)

	if upr.StructureVersion != other.StructureVersion {
		return false
	}

	if !upr.m_code.Equals(other.m_code) {
		return false
	}

	if !upr.m_name.Equals(other.m_name) {
		return false
	}

	if !upr.m_description.Equals(other.m_description) {
		return false
	}

	if !upr.m_value.Equals(other.m_value) {
		return false
	}

	if !upr.m_rewardTypeName.Equals(other.m_rewardTypeName) {
		return false
	}

	if !upr.m_gameCode.Equals(other.m_gameCode) {
		return false
	}

	return upr.m_platforms.Equals(other.m_platforms)
}

// CopyRef copies the current value of the UplayReward
// and returns a pointer to the new copy
func (upr UplayReward) CopyRef() types.RVTypePtr {
	copied := upr.Copy().(UplayReward)
	return &copied
}

// Deref takes a pointer to the UplayReward
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (upr *UplayReward) Deref() types.RVType {
	return *upr
}

// String returns the string representation of the UplayReward
func (upr UplayReward) String() string {
	return upr.FormatToString(0)
}

// FormatToString pretty-prints the UplayReward using the provided indentation level
func (upr UplayReward) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayReward{\n")
	fmt.Fprintf(&b, "%sm_code: %v,\n", indentationValues, upr.m_code)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, upr.m_name)
	fmt.Fprintf(&b, "%sm_description: %v,\n", indentationValues, upr.m_description)
	fmt.Fprintf(&b, "%sm_value: %v,\n", indentationValues, upr.m_value)
	fmt.Fprintf(&b, "%sm_rewardTypeName: %v,\n", indentationValues, upr.m_rewardTypeName)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, upr.m_gameCode)
	fmt.Fprintf(&b, "%sm_platforms: %v,\n", indentationValues, upr.m_platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayReward returns a new UplayReward
func NewUplayReward() UplayReward {
	return UplayReward{
		m_code:           "code",
		m_name:           "name",
		m_description:    "desc",
		m_value:          0,
		m_rewardTypeName: "Reward",
		m_gameCode:       "Uplay",
		m_platforms:      types.NewList[UplayRewardPlatform](),
	}
}
