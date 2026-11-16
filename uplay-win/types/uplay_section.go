package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySection is a type within the Uplay Win protocol
type UplaySection struct {
	types.Structure
	Key          types.String
	Name         types.String
	TypeName     types.String
	MenuTypeName types.String
	ContentList  types.List[UplaySectionContent]
	GameCode     types.String
	PlatformCode types.String
}

// ObjectID returns the object identifier of the type
func (us UplaySection) ObjectID() types.RVType {
	return us.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (us UplaySection) DataObjectID() types.RVType {
	return types.NewString("UplaySection")
}

// WriteTo writes the UplaySection to the given writable
func (us UplaySection) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	us.Key.WriteTo(contentWritable)
	us.Name.WriteTo(contentWritable)
	us.TypeName.WriteTo(contentWritable)
	us.MenuTypeName.WriteTo(contentWritable)
	us.ContentList.WriteTo(contentWritable)
	us.GameCode.WriteTo(contentWritable)
	us.PlatformCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	us.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySection from the given readable
func (us *UplaySection) ExtractFrom(readable types.Readable) error {
	if err := us.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection Header. %s", err.Error())
	}

	if err := us.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.Key. %s", err.Error())
	}

	if err := us.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.Name. %s", err.Error())
	}

	if err := us.TypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.TypeName. %s", err.Error())
	}

	if err := us.MenuTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.MenuTypeName. %s", err.Error())
	}

	if err := us.ContentList.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.ContentList. %s", err.Error())
	}

	if err := us.GameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.GameCode. %s", err.Error())
	}

	if err := us.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.PlatformCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySection
func (us UplaySection) Copy() types.RVType {
	copied := NewUplaySection()

	copied.StructureVersion = us.StructureVersion
	copied.Key = us.Key
	copied.Name = us.Name
	copied.TypeName = us.TypeName
	copied.MenuTypeName = us.MenuTypeName
	copied.ContentList = us.ContentList
	copied.GameCode = us.GameCode
	copied.PlatformCode = us.PlatformCode

	return copied
}

// Equals checks if the given UplaySection contains the same data as the current UplaySection
func (us UplaySection) Equals(o types.RVType) bool {
	if _, ok := o.(UplaySection); !ok {
		return false
	}

	other := o.(UplaySection)

	if us.StructureVersion != other.StructureVersion {
		return false
	}

	if !us.Key.Equals(other.Key) {
		return false
	}

	if !us.Name.Equals(other.Name) {
		return false
	}

	if !us.TypeName.Equals(other.TypeName) {
		return false
	}

	if !us.MenuTypeName.Equals(other.MenuTypeName) {
		return false
	}

	if !us.ContentList.Equals(other.ContentList) {
		return false
	}

	if !us.GameCode.Equals(other.GameCode) {
		return false
	}

	return us.PlatformCode.Equals(other.PlatformCode)
}

// CopyRef copies the current value of the UplaySection
// and returns a pointer to the new copy
func (us UplaySection) CopyRef() types.RVTypePtr {
	copied := us.Copy().(UplaySection)
	return &copied
}

// Deref takes a pointer to the UplaySection
// and dereferences it to the raw value.
func (us *UplaySection) Deref() types.RVType {
	return *us
}

// String returns the string representation of the UplaySection
func (us UplaySection) String() string {
	return us.FormatToString(0)
}

// FormatToString pretty-prints the UplaySection using the provided indentation level
func (us UplaySection) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySection{\n")
	fmt.Fprintf(&b, "%sm_key: %v,\n", indentationValues, us.Key)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, us.Name)
	fmt.Fprintf(&b, "%sm_typeName: %v,\n", indentationValues, us.TypeName)
	fmt.Fprintf(&b, "%sm_menuTypeName: %v,\n", indentationValues, us.MenuTypeName)
	fmt.Fprintf(&b, "%sm_contentList: %v,\n", indentationValues, us.ContentList)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, us.GameCode)
	fmt.Fprintf(&b, "%sm_platformCode: %v,\n", indentationValues, us.PlatformCode)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySection returns a new UplaySection
func NewUplaySection() UplaySection {
	return UplaySection{
		Key:          "key",
		Name:         "name",
		TypeName:     "type",
		MenuTypeName: "menu",
		ContentList:  types.NewList[UplaySectionContent](),
		GameCode:     "Uplay",
		PlatformCode: "PC",
	}
}
