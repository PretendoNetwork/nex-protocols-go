// Package types implements all the Types used by the UbiNews protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NewsHeader is a type with the UbiNews protocol
type NewsHeader struct {
	types.Structure
	ID              types.UInt32
	RecipientID     types.UInt32
	RecipientType   types.UInt32
	PublisherPID    types.PID
	PublisherName   types.String
	PublicationTime types.DateTime
	DisplayTime     types.DateTime
	ExpirationTime  types.DateTime
	Title           types.String
	Link            types.String
}

// WriteTo writes the NewsHeader to the given writable
func (nh NewsHeader) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nh.ID.WriteTo(contentWritable)
	nh.RecipientID.WriteTo(contentWritable)
	nh.RecipientType.WriteTo(contentWritable)
	nh.PublisherPID.WriteTo(contentWritable)
	nh.PublisherName.WriteTo(contentWritable)
	nh.PublicationTime.WriteTo(contentWritable)
	nh.DisplayTime.WriteTo(contentWritable)
	nh.ExpirationTime.WriteTo(contentWritable)
	nh.Title.WriteTo(contentWritable)
	nh.Link.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nh.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NewsHeader from the given readable
func (nh *NewsHeader) ExtractFrom(readable types.Readable) error {
	if err := nh.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader header. %s", err.Error())
	}

	if err := nh.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_ID. %s", err.Error())
	}

	if err := nh.RecipientID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_recipientID. %s", err.Error())
	}

	if err := nh.RecipientType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_recipientType. %s", err.Error())
	}

	if err := nh.PublisherPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publisherPID. %s", err.Error())
	}

	if err := nh.PublisherName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publisherName. %s", err.Error())
	}

	if err := nh.PublicationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publicationTime. %s", err.Error())
	}

	if err := nh.DisplayTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_displayTime. %s", err.Error())
	}

	if err := nh.ExpirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_expirationTime. %s", err.Error())
	}

	if err := nh.Title.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_title. %s", err.Error())
	}

	if err := nh.Link.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_link. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of NewsHeader
func (nh NewsHeader) Copy() types.RVType {
	copied := NewNewsHeader()

	copied.StructureVersion = nh.StructureVersion
	copied.ID = nh.ID
	copied.RecipientID = nh.RecipientID
	copied.RecipientType = nh.RecipientType
	copied.PublisherPID = nh.PublisherPID
	copied.PublisherName = nh.PublisherName
	copied.PublicationTime = nh.PublicationTime
	copied.DisplayTime = nh.DisplayTime
	copied.ExpirationTime = nh.ExpirationTime
	copied.Title = nh.Title
	copied.Link = nh.Link

	return copied
}

// Equals checks if the given NewsHeader contains the same data as the current NewsHeader
func (nh NewsHeader) Equals(o types.RVType) bool {
	if _, ok := o.(NewsHeader); !ok {
		return false
	}

	other := o.(NewsHeader)

	if nh.StructureVersion != other.StructureVersion {
		return false
	}

	if nh.ID != other.ID {
		return false
	}

	if nh.RecipientID != other.RecipientID {
		return false
	}

	if nh.RecipientType != other.RecipientType {
		return false
	}

	if nh.PublisherPID != other.PublisherPID {
		return false
	}

	if nh.PublisherName != other.PublisherName {
		return false
	}

	if nh.PublicationTime != other.PublicationTime {
		return false
	}

	if nh.DisplayTime != other.DisplayTime {
		return false
	}

	if nh.ExpirationTime != other.ExpirationTime {
		return false
	}

	if nh.Title != other.Title {
		return false
	}

	return nh.Link == other.Link
}

// CopyRef copies the current value of the NewsHeader
// and returns a pointer to the new copy
func (nh NewsHeader) CopyRef() types.RVTypePtr {
	copied := nh.Copy().(NewsHeader)
	return &copied
}

// Deref takes a pointer of NewsHeader
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (nc *NewsHeader) Deref() types.RVType {
	return *nc
}

// String returns the string representation of the NewsHeader
func (nh NewsHeader) String() string {
	return nh.FormatToString(0)
}

// FormatToString pretty-prints the NewsHeader using the provided indentation level
func (nh NewsHeader) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsHeader{\n}")
	fmt.Fprintf(&b, "%sm_ID: %s,\n", indentationValues, nh.ID)
	fmt.Fprintf(&b, "%sm_recipientID: %s,\n", indentationValues, nh.RecipientID)
	fmt.Fprintf(&b, "%sm_recipientType: %s,\n", indentationValues, nh.RecipientType)
	fmt.Fprintf(&b, "%sm_publisherPID: %s,\n", indentationValues, nh.PublisherPID)
	fmt.Fprintf(&b, "%sm_publisherName: %s,\n", indentationValues, nh.PublisherName)
	fmt.Fprintf(&b, "%sm_publicationTime: %s,\n", indentationValues, nh.PublicationTime)
	fmt.Fprintf(&b, "%sm_type: %s,\n", indentationValues, nh.DisplayTime)
	fmt.Fprintf(&b, "%sm_expirationTime: %s,\n", indentationValues, nh.ExpirationTime)
	fmt.Fprintf(&b, "%sm_title: %s,\n", indentationValues, nh.Title)
	fmt.Fprintf(&b, "%sm_link: %s,\n", indentationValues, nh.Link)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsHeader returns a new NewsHeader
func NewNewsHeader() NewsHeader {
	return NewsHeader{
		ID:              types.NewUInt32(0),
		RecipientID:     types.NewUInt32(0),
		RecipientType:   types.NewUInt32(0),
		PublisherPID:    types.NewPID(0),
		PublisherName:   types.NewString(""),
		PublicationTime: types.NewDateTime(0),
		DisplayTime:     types.NewDateTime(0),
		ExpirationTime:  types.NewDateTime(0),
		Title:           types.NewString(""),
		Link:            types.NewString(""),
	}

}
