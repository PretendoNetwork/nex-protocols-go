// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemStartChallengeParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemStartChallengeParam struct {
	types.Structure
	ChallengeScheduleID *types.PrimitiveU32
	TicketType          *types.PrimitiveU32
	NumTicket           *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemStartChallengeParam from the given readable
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemStartChallengeParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemStartChallengeParam header. %s", err.Error())
	}

	err = serviceItemStartChallengeParam.ChallengeScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.ChallengeScheduleID from stream. %s", err.Error())
	}

	err = serviceItemStartChallengeParam.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.TicketType from stream. %s", err.Error())
	}

	err = serviceItemStartChallengeParam.NumTicket.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.NumTicket from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemStartChallengeParam to the given writable
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemStartChallengeParam.ChallengeScheduleID.WriteTo(contentWritable)
	serviceItemStartChallengeParam.TicketType.WriteTo(contentWritable)
	serviceItemStartChallengeParam.NumTicket.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemStartChallengeParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemStartChallengeParam
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) Copy() types.RVType {
	copied := NewServiceItemStartChallengeParam()

	copied.StructureVersion = serviceItemStartChallengeParam.StructureVersion

	copied.ChallengeScheduleID = serviceItemStartChallengeParam.ChallengeScheduleID
	copied.TicketType = serviceItemStartChallengeParam.TicketType
	copied.NumTicket = serviceItemStartChallengeParam.NumTicket

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemStartChallengeParam); !ok {
		return false
	}

	other := o.(*ServiceItemStartChallengeParam)

	if serviceItemStartChallengeParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemStartChallengeParam.ChallengeScheduleID.Equals(other.ChallengeScheduleID) {
		return false
	}

	if !serviceItemStartChallengeParam.TicketType.Equals(other.TicketType) {
		return false
	}

	if !serviceItemStartChallengeParam.NumTicket.Equals(other.NumTicket) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) String() string {
	return serviceItemStartChallengeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemStartChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemStartChallengeParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %d,\n", indentationValues, serviceItemStartChallengeParam.ChallengeScheduleID))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemStartChallengeParam.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %d,\n", indentationValues, serviceItemStartChallengeParam.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemStartChallengeParam returns a new ServiceItemStartChallengeParam
func NewServiceItemStartChallengeParam() *ServiceItemStartChallengeParam {
	return &ServiceItemStartChallengeParam{}
}
