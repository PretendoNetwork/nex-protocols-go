package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayAction is a type within the Uplay Win protocol
type UplayAction struct {
	types.Structure
	Code           types.String
	Name           types.String
	Description    types.String
	Value          types.Int32
	RewardTypeName types.String
	GameCode       types.String
	Platforms      types.List[UplayActionPlatform]
}

// ObjectID returns the object identifier of the type
func (upa UplayAction) ObjectID() types.RVType {
	return upa.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (upa UplayAction) DataObjectID() types.RVType {
	return types.NewString("UplayAction")
}

// WriteTo writes the UplayAction to the given writable
func (upa UplayAction) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	upa.Code.WriteTo(contentWritable)
	upa.Name.WriteTo(contentWritable)
	upa.Description.WriteTo(contentWritable)
	upa.Value.WriteTo(contentWritable)
	upa.RewardTypeName.WriteTo(contentWritable)
	upa.GameCode.WriteTo(contentWritable)
	upa.Platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	upa.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayAction from the given readable
func (upa *UplayAction) ExtractFrom(readable types.Readable) error {
	if err := upa.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction Header. %s", err.Error())
	}

	if err := upa.Code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Code. %s", err.Error())
	}

	if err := upa.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Name. %s", err.Error())
	}

	if err := upa.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Description. %s", err.Error())
	}

	if err := upa.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Value. %s", err.Error())
	}

	if err := upa.RewardTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.RewardTypeName. %s", err.Error())
	}

	if err := upa.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.GameCode. %s", err.Error())
	}

	if err := upa.Platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.Platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayAction
func (upa UplayAction) Copy() types.RVType {
	copied := NewUplayAction()

	copied.StructureVersion = upa.StructureVersion
	copied.Code = upa.Code
	copied.Name = upa.Name
	copied.Description = upa.Description
	copied.Value = upa.Value
	copied.RewardTypeName = upa.RewardTypeName
	copied.GameCode = upa.GameCode
	copied.Platforms = upa.Platforms

	return copied
}

// Equals checks if the given UplayAction contains the same data as the current UplayAction
func (upa UplayAction) Equals(o types.RVType) bool {
	if _, ok := o.(UplayAction); !ok {
		return false
	}

	other := o.(UplayAction)

	if upa.StructureVersion != other.StructureVersion {
		return false
	}

	if !upa.Code.Equals(other.Code) {
		return false
	}

	if !upa.Name.Equals(other.Name) {
		return false
	}

	if !upa.Description.Equals(other.Description) {
		return false
	}

	if !upa.Value.Equals(other.Value) {
		return false
	}

	if !upa.RewardTypeName.Equals(other.RewardTypeName) {
		return false
	}

	if !upa.GameCode.Equals(other.GameCode) {
		return false
	}

	return upa.Platforms.Equals(other.Platforms)
}

// CopyRef copies the current value of the UplayAction
// and returns a pointer to the new copy
func (upa UplayAction) CopyRef() types.RVTypePtr {
	copied := upa.Copy().(UplayAction)
	return &copied
}

// Deref takes a pointer to the UplayAction
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (upa *UplayAction) Deref() types.RVType {
	return *upa
}

// String returns the string representation of the UplayAction
func (upa UplayAction) String() string {
	return upa.FormatToString(0)
}

// FormatToString pretty-prints the UplayAction using the provided indentation level
func (upa UplayAction) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayAction{\n")
	fmt.Fprintf(&b, "%sm_code: %v,\n", indentationValues, upa.Code)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, upa.Name)
	fmt.Fprintf(&b, "%sm_description: %v,\n", indentationValues, upa.Description)
	fmt.Fprintf(&b, "%sm_value: %v,\n", indentationValues, upa.Value)
	fmt.Fprintf(&b, "%sm_rewardTypeName: %v,\n", indentationValues, upa.RewardTypeName)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, upa.GameCode)
	fmt.Fprintf(&b, "%sm_platforms: %v,\n", indentationValues, upa.Platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayAction returns a new UplayAction
func NewUplayAction() UplayAction {
	return UplayAction{
		Code:           "code",
		Name:           "name",
		Description:    "desc",
		Value:          0,
		RewardTypeName: "Reward",
		GameCode:       "Uplay",
		Platforms:      types.NewList[UplayActionPlatform](),
	}
}
