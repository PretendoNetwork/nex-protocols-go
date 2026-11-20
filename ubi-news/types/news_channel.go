// Package types implements all the Types used by the UbiNews protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsChannel is a type with the UbiNews protocol
type NewsChannel struct {
	types.Structure
	ID             types.UInt32
	OwnerPID       types.PID
	Name           types.String
	Description    types.String
	CreationTime   types.DateTime
	ExpirationTime types.DateTime
	NewsType       types.String
	Locale         types.String
	Subscribable   types.Bool
}

// WriteTo writes the NewsChannel to the given writable
func (nc NewsChannel) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nc.ID.WriteTo(contentWritable)
	nc.OwnerPID.WriteTo(contentWritable)
	nc.Name.WriteTo(contentWritable)
	nc.Description.WriteTo(contentWritable)
	nc.CreationTime.WriteTo(contentWritable)
	nc.ExpirationTime.WriteTo(contentWritable)
	nc.NewsType.WriteTo(contentWritable)
	nc.Locale.WriteTo(contentWritable)
	nc.Subscribable.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsChannel from the given readable
func (nc *NewsChannel) ExtractFrom(readable types.Readable) error {
	if err := nc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel header. %s", err.Error())
	}

	if err := nc.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_ID. %s", err.Error())
	}

	if err := nc.OwnerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_OwnerPID. %s", err.Error())
	}

	if err := nc.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_Name. %s", err.Error())
	}

	if err := nc.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_Description. %s", err.Error())
	}

	if err := nc.CreationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_CreationTime. %s", err.Error())
	}

	if err := nc.ExpirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_ExpirationTime. %s", err.Error())
	}

	if err := nc.NewsType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_type. %s", err.Error())
	}

	if err := nc.Locale.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_Locale. %s", err.Error())
	}

	if err := nc.Subscribable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_Subscribable. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of NewsChannel
func (nc NewsChannel) Copy() types.RVType {
	copied := NewNewsChannel()

	copied.StructureVersion = nc.StructureVersion
	copied.ID = nc.ID
	copied.OwnerPID = nc.OwnerPID
	copied.Name = nc.Name
	copied.Description = nc.Description
	copied.CreationTime = nc.CreationTime
	copied.ExpirationTime = nc.ExpirationTime
	copied.NewsType = nc.NewsType
	copied.Locale = nc.Locale
	copied.Subscribable = nc.Subscribable

	return copied
}

// Equals checks if the given NewsChannel contains the same data as the current NewsChannel
func (nc NewsChannel) Equals(o types.RVType) bool {
	if _, ok := o.(NewsChannel); !ok {
		return false
	}

	other := o.(NewsChannel)

	if nc.StructureVersion != other.StructureVersion {
		return false
	}

	if nc.ID != other.ID {
		return false
	}

	if nc.OwnerPID != other.OwnerPID {
		return false
	}

	if nc.Name != other.Name {
		return false
	}

	if nc.Description != other.Description {
		return false
	}

	if nc.CreationTime != other.CreationTime {
		return false
	}

	if nc.ExpirationTime != other.ExpirationTime {
		return false
	}

	if nc.NewsType != other.NewsType {
		return false
	}

	if nc.Locale != other.Locale {
		return false
	}

	return nc.Subscribable == other.Subscribable
}

// CopyRef copies the current value of the NewsChannel
// and returns a pointer to the new copy
func (nc NewsChannel) CopyRef() types.RVTypePtr {
	copied := nc.Copy().(NewsChannel)
	return &copied
}

// Deref takes a pointer of NewsChannel
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nc *NewsChannel) Deref() types.RVType {
	return *nc
}

// String returns the string representation of the NewsChannel
func (nc NewsChannel) String() string {
	return nc.FormatToString(0)
}

// FormatToString pretty-prints the NewsChannel using the provided indentation level
func (nc NewsChannel) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsChannel{\n}")
	fmt.Fprintf(&b, "%sm_ID: %s,\n", indentationValues, nc.ID)
	fmt.Fprintf(&b, "%sm_OwnerPID: %s,\n", indentationValues, nc.OwnerPID)
	fmt.Fprintf(&b, "%sm_Name: %s,\n", indentationValues, nc.Name)
	fmt.Fprintf(&b, "%sm_Description: %s,\n", indentationValues, nc.Description)
	fmt.Fprintf(&b, "%sm_CreationTime: %s,\n", indentationValues, nc.CreationTime)
	fmt.Fprintf(&b, "%sm_ExpirationTime: %s,\n", indentationValues, nc.ExpirationTime)
	fmt.Fprintf(&b, "%sm_type: %s,\n", indentationValues, nc.NewsType)
	fmt.Fprintf(&b, "%sm_Locale: %s,\n", indentationValues, nc.Locale)
	fmt.Fprintf(&b, "%sm_Subscribable: %s,\n", indentationValues, nc.Subscribable)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsChannel returns a new NewsChannel
func NewNewsChannel() NewsChannel {
	return NewsChannel{
		ID:             types.NewUInt32(0),
		OwnerPID:       types.NewPID(0),
		Name:           types.NewString(""),
		Description:    types.NewString(""),
		CreationTime:   types.NewDateTime(0),
		ExpirationTime: types.NewDateTime(0),
		NewsType:       types.NewString(""),
		Locale:         types.NewString(""),
		Subscribable:   types.NewBool(true),
	}

}
