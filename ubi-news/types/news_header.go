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
	recipientID     types.UInt32
	recipientType   types.UInt32
	publisherPID    types.PID
	publisherName   types.String
	publicationTime types.DateTime
	displayTime     types.DateTime
	expirationTime  types.DateTime
	title           types.String
	link            types.String
}

// WriteTo writes the NewsHeader to the given writable
func (nh NewsHeader) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nh.ID.WriteTo(contentWritable)
	nh.recipientID.WriteTo(contentWritable)
	nh.recipientType.WriteTo(contentWritable)
	nh.publisherPID.WriteTo(contentWritable)
	nh.publisherName.WriteTo(contentWritable)
	nh.publicationTime.WriteTo(contentWritable)
	nh.displayTime.WriteTo(contentWritable)
	nh.expirationTime.WriteTo(contentWritable)
	nh.title.WriteTo(contentWritable)
	nh.link.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nh.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

func (nh *NewsHeader) ExtractFrom(readable types.Readable) error {
	if err := nh.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader header. %s", err.Error())
	}

	if err := nh.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_ID. %s", err.Error())
	}

	if err := nh.recipientID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_recipientID. %s", err.Error())
	}

	if err := nh.recipientType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_recipientType. %s", err.Error())
	}

	if err := nh.publisherPID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publisherPID. %s", err.Error())
	}

	if err := nh.publisherName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publisherName. %s", err.Error())
	}

	if err := nh.publicationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_publicationTime. %s", err.Error())
	}

	if err := nh.displayTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_displayTime. %s", err.Error())
	}

	if err := nh.expirationTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_expirationTime. %s", err.Error())
	}

	if err := nh.title.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_title. %s", err.Error())
	}

	if err := nh.link.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract NewsHeader.m_link. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instanced of NewsHeader
func (nh NewsHeader) Copy() types.RVType {
	copied := NewNewsHeader()

	copied.StructureVersion = nh.StructureVersion
	copied.ID = nh.ID
	copied.recipientID = nh.recipientID
	copied.recipientType = nh.recipientType
	copied.publisherPID = nh.publisherPID
	copied.publisherName = nh.publisherName
	copied.publicationTime = nh.publicationTime
	copied.displayTime = nh.displayTime
	copied.expirationTime = nh.expirationTime
	copied.title = nh.title
	copied.link = nh.link

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

	if nh.recipientID != other.recipientID {
		return false
	}

	if nh.recipientType != other.recipientType {
		return false
	}

	if nh.publisherPID != other.publisherPID {
		return false
	}

	if nh.publisherName != other.publisherName {
		return false
	}

	if nh.publicationTime != other.publicationTime {
		return false
	}

	if nh.displayTime != other.displayTime {
		return false
	}

	if nh.expirationTime != other.expirationTime {
		return false
	}

	if nh.title != other.title {
		return false
	}

	return nh.link == other.link
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

func (nh NewsHeader) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NewsHeader{\n}")
	fmt.Fprintf(&b, "%sm_ID: %s,\n", indentationValues, nh.ID)
	fmt.Fprintf(&b, "%sm_recipientID: %s,\n", indentationValues, nh.recipientID)
	fmt.Fprintf(&b, "%sm_recipientType: %s,\n", indentationValues, nh.recipientType)
	fmt.Fprintf(&b, "%sm_publisherPID: %s,\n", indentationValues, nh.publisherPID)
	fmt.Fprintf(&b, "%sm_publisherName: %s,\n", indentationValues, nh.publisherName)
	fmt.Fprintf(&b, "%sm_publicationTime: %s,\n", indentationValues, nh.publicationTime)
	fmt.Fprintf(&b, "%sm_type: %s,\n", indentationValues, nh.displayTime)
	fmt.Fprintf(&b, "%sm_expirationTime: %s,\n", indentationValues, nh.expirationTime)
	fmt.Fprintf(&b, "%sm_title: %s,\n", indentationValues, nh.title)
	fmt.Fprintf(&b, "%sm_link: %s,\n", indentationValues, nh.link)

	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewNewsHeader returns a new NewsHeader
func NewNewsHeader() NewsHeader {
	return NewsHeader{
		ID:              types.NewUInt32(0),
		recipientID:     types.NewUInt32(0),
		recipientType:   types.NewUInt32(0),
		publisherPID:    types.NewPID(0),
		publisherName:   types.NewString(""),
		publicationTime: types.NewDateTime(0),
		displayTime:     types.NewDateTime(0),
		expirationTime:  types.NewDateTime(0),
		title:           types.NewString(""),
		link:            types.NewString(""),
	}

}
