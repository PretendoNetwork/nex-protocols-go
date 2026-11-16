package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayReward is a type within the Uplay Win protocol
type UplayReward struct {
	types.Structure
	Code           types.String
	Name           types.String
	Description    types.String
	Value          types.Int32
	RewardTypeName types.String
	GameCode       types.String
	Platforms      types.List[UplayRewardPlatform]
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

	upr.Code.WriteTo(contentWritable)
	upr.Name.WriteTo(contentWritable)
	upr.Description.WriteTo(contentWritable)
	upr.Value.WriteTo(contentWritable)
	upr.RewardTypeName.WriteTo(contentWritable)
	upr.GameCode.WriteTo(contentWritable)
	upr.Platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	upr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayReward from the given readable
func (upr *UplayReward) ExtractFrom(readable types.Readable) error {
	if err := upr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward Header. %s", err.Error())
	}

	if err := upr.Code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Code. %s", err.Error())
	}

	if err := upr.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Name. %s", err.Error())
	}

	if err := upr.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Description. %s", err.Error())
	}

	if err := upr.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Value. %s", err.Error())
	}

	if err := upr.RewardTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.RewardTypeName. %s", err.Error())
	}

	if err := upr.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.GameCode. %s", err.Error())
	}

	if err := upr.Platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayReward.Platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayReward
func (upr UplayReward) Copy() types.RVType {
	copied := NewUplayReward()

	copied.StructureVersion = upr.StructureVersion
	copied.Code = upr.Code
	copied.Name = upr.Name
	copied.Description = upr.Description
	copied.Value = upr.Value
	copied.RewardTypeName = upr.RewardTypeName
	copied.GameCode = upr.GameCode
	copied.Platforms = upr.Platforms

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

	if !upr.Code.Equals(other.Code) {
		return false
	}

	if !upr.Name.Equals(other.Name) {
		return false
	}

	if !upr.Description.Equals(other.Description) {
		return false
	}

	if !upr.Value.Equals(other.Value) {
		return false
	}

	if !upr.RewardTypeName.Equals(other.RewardTypeName) {
		return false
	}

	if !upr.GameCode.Equals(other.GameCode) {
		return false
	}

	return upr.Platforms.Equals(other.Platforms)
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
	fmt.Fprintf(&b, "%sm_code: %v,\n", indentationValues, upr.Code)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, upr.Name)
	fmt.Fprintf(&b, "%sm_description: %v,\n", indentationValues, upr.Description)
	fmt.Fprintf(&b, "%sm_value: %v,\n", indentationValues, upr.Value)
	fmt.Fprintf(&b, "%sm_rewardTypeName: %v,\n", indentationValues, upr.RewardTypeName)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, upr.GameCode)
	fmt.Fprintf(&b, "%sm_platforms: %v,\n", indentationValues, upr.Platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayReward returns a new UplayReward
func NewUplayReward() UplayReward {
	return UplayReward{
		Code:           "code",
		Name:           "name",
		Description:    "desc",
		Value:          0,
		RewardTypeName: "Reward",
		GameCode:       "Uplay",
		Platforms:      types.NewList[UplayRewardPlatform](),
	}
}
