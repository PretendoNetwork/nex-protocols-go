// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemUserInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemUserInfo struct {
	types.Structure
	NumTotalEntryTicket *types.PrimitiveU32
	ApplicationBuffer   []byte
}

// ExtractFrom extracts the ServiceItemUserInfo from the given readable
func (serviceItemUserInfo *ServiceItemUserInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemUserInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemUserInfo header. %s", err.Error())
	}

	err = serviceItemUserInfo.NumTotalEntryTicket.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo.NumTotalEntryTicket from stream. %s", err.Error())
	}

	serviceItemUserInfo.ApplicationBuffer, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemUserInfo.ApplicationBuffer from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemUserInfo to the given writable
func (serviceItemUserInfo *ServiceItemUserInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemUserInfo.NumTotalEntryTicket.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemUserInfo.ApplicationBuffer)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemUserInfo
func (serviceItemUserInfo *ServiceItemUserInfo) Copy() types.RVType {
	copied := NewServiceItemUserInfo()

	copied.StructureVersion = serviceItemUserInfo.StructureVersion

	copied.NumTotalEntryTicket = serviceItemUserInfo.NumTotalEntryTicket
	copied.ApplicationBuffer = serviceItemUserInfo.ApplicationBuffer

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemUserInfo *ServiceItemUserInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemUserInfo); !ok {
		return false
	}

	other := o.(*ServiceItemUserInfo)

	if serviceItemUserInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemUserInfo.NumTotalEntryTicket.Equals(other.NumTotalEntryTicket) {
		return false
	}

	if !serviceItemUserInfo.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemUserInfo *ServiceItemUserInfo) String() string {
	return serviceItemUserInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemUserInfo *ServiceItemUserInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemUserInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemUserInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNumTotalEntryTicket: %d,\n", indentationValues, serviceItemUserInfo.NumTotalEntryTicket))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %x,\n", indentationValues, serviceItemUserInfo.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemUserInfo returns a new ServiceItemUserInfo
func NewServiceItemUserInfo() *ServiceItemUserInfo {
	return &ServiceItemUserInfo{}
}
