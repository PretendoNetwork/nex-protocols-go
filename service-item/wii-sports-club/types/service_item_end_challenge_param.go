// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemEndChallengeParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemEndChallengeParam struct {
	types.Structure
	ChallengeScheduleID *types.PrimitiveU32
	UserInfo            *ServiceItemUserInfo
}

// ExtractFrom extracts the ServiceItemEndChallengeParam from the given readable
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemEndChallengeParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemEndChallengeParam header. %s", err.Error())
	}

	err = serviceItemEndChallengeParam.ChallengeScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.ChallengeScheduleID from stream. %s", err.Error())
	}

	err = serviceItemEndChallengeParam.UserInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.UserInfo from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemEndChallengeParam to the given writable
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemEndChallengeParam.ChallengeScheduleID.WriteTo(contentWritable)
	serviceItemEndChallengeParam.UserInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemEndChallengeParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemEndChallengeParam
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) Copy() types.RVType {
	copied := NewServiceItemEndChallengeParam()

	copied.StructureVersion = serviceItemEndChallengeParam.StructureVersion

	copied.ChallengeScheduleID = serviceItemEndChallengeParam.ChallengeScheduleID
	copied.UserInfo = serviceItemEndChallengeParam.UserInfo.Copy().(*ServiceItemUserInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemEndChallengeParam); !ok {
		return false
	}

	other := o.(*ServiceItemEndChallengeParam)

	if serviceItemEndChallengeParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemEndChallengeParam.ChallengeScheduleID.Equals(other.ChallengeScheduleID) {
		return false
	}

	if !serviceItemEndChallengeParam.UserInfo.Equals(other.UserInfo) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) String() string {
	return serviceItemEndChallengeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEndChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemEndChallengeParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %d,\n", indentationValues, serviceItemEndChallengeParam.ChallengeScheduleID))

	if serviceItemEndChallengeParam.UserInfo != nil {
		b.WriteString(fmt.Sprintf("%sUserInfo: %s\n", indentationValues, serviceItemEndChallengeParam.UserInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUserInfo: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEndChallengeParam returns a new ServiceItemEndChallengeParam
func NewServiceItemEndChallengeParam() *ServiceItemEndChallengeParam {
	return &ServiceItemEndChallengeParam{}
}
