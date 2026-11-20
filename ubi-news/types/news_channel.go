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
	ownerPID       types.PID
	name           types.String
	description    types.String
	creationTime   types.DateTime
	expirationTime types.DateTime
	newsType       types.String
	locale         types.String
	subscribable   types.Bool
}

// WriteTo writes the NewsChannel to the given writable
func (nc NewsChannel) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nc.ID.WriteTo(contentWritable)
	nc.ownerPID.WriteTo(contentWritable)
	nc.name.WriteTo(contentWritable)
	nc.description.WriteTo(contentWritable)
	nc.creationTime.WriteTo(contentWritable)
	nc.expirationTime.WriteTo(contentWritable)
	nc.newsType.WriteTo(contentWritable)
	nc.locale.WriteTo(contentWritable)
	nc.subscribable.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

func (nc *NewsChannel) ExtractFrom(readable types.Readable) error {
	if err := nc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel header. %s", err.Error())
	}

	if err := nc.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_ID. %s", err.Error())
	}

	if err := nc.ownerPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_ownerPID. %s", err.Error())
	}

	if err := nc.name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_name. %s", err.Error())
	}

	if err := nc.description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_description. %s", err.Error())
	}

	if err := nc.creationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_creationTime. %s", err.Error())
	}

	if err := nc.expirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_expirationTime. %s", err.Error())
	}

	if err := nc.newsType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_type. %s", err.Error())
	}

	if err := nc.locale.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_locale. %s", err.Error())
	}

	if err := nc.subscribable.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsChannel.m_subscribable. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of NewsChannel
func (nc NewsChannel) Copy() types.RVType {
	copied := NewNewsChannel()

	copied.StructureVersion = nc.StructureVersion
	copied.ID = nc.ID
	copied.ownerPID = nc.ownerPID
	copied.name = nc.name
	copied.description = nc.description
	copied.creationTime = nc.creationTime
	copied.expirationTime = nc.expirationTime
	copied.newsType = nc.newsType
	copied.locale = nc.locale
	copied.subscribable = nc.subscribable

	return copied
}

// Equal checks if the given NewsChannel contains the same data as the current NewsChannel
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

	if nc.ownerPID != other.ownerPID {
		return false
	}

	if nc.name != other.name {
		return false
	}

	if nc.description != other.description {
		return false
	}

	if nc.creationTime != other.creationTime {
		return false
	}

	if nc.expirationTime != other.expirationTime {
		return false
	}

	if nc.newsType != other.newsType {
		return false
	}

	if nc.locale != other.locale {
		return false
	}

	return nc.subscribable == other.subscribable
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

func (nc NewsChannel) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsChannel{\n}")
	fmt.Fprintf(&b, "%sm_ID: %s,\n", indentationValues, nc.ID)
	fmt.Fprintf(&b, "%sm_ownerPID: %s,\n", indentationValues, nc.ownerPID)
	fmt.Fprintf(&b, "%sm_name: %s,\n", indentationValues, nc.name)
	fmt.Fprintf(&b, "%sm_description: %s,\n", indentationValues, nc.description)
	fmt.Fprintf(&b, "%sm_creationTime: %s,\n", indentationValues, nc.creationTime)
	fmt.Fprintf(&b, "%sm_expirationTime: %s,\n", indentationValues, nc.expirationTime)
	fmt.Fprintf(&b, "%sm_type: %s,\n", indentationValues, nc.newsType)
	fmt.Fprintf(&b, "%sm_locale: %s,\n", indentationValues, nc.locale)
	fmt.Fprintf(&b, "%sm_subscribable: %s,\n", indentationValues, nc.subscribable)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsChannel returns a new NewsChannel
func NewNewsChannel() NewsChannel {
	return NewsChannel{
		ID:             types.NewUInt32(0),
		ownerPID:       types.NewPID(0),
		name:           types.NewString(""),
		description:    types.NewString(""),
		creationTime:   types.NewDateTime(0),
		expirationTime: types.NewDateTime(0),
		newsType:       types.NewString(""),
		locale:         types.NewString(""),
		subscribable:   types.NewBool(true),
	}

}
