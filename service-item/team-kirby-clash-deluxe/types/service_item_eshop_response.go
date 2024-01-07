// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemEShopResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemEShopResponse struct {
	types.Structure
	HTTPStatus    *types.PrimitiveU32
	ErrorCode     *types.PrimitiveU32
	CorrelationID string
}

// ExtractFrom extracts the ServiceItemEShopResponse from the given readable
func (serviceItemEShopResponse *ServiceItemEShopResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemEShopResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemEShopResponse header. %s", err.Error())
	}

	err = serviceItemEShopResponse.HTTPStatus.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.HTTPStatus from stream. %s", err.Error())
	}

	err = serviceItemEShopResponse.ErrorCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.ErrorCode from stream. %s", err.Error())
	}

	err = serviceItemEShopResponse.CorrelationID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.CorrelationID from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemEShopResponse to the given writable
func (serviceItemEShopResponse *ServiceItemEShopResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemEShopResponse.HTTPStatus.WriteTo(contentWritable)
	serviceItemEShopResponse.ErrorCode.WriteTo(contentWritable)
	serviceItemEShopResponse.CorrelationID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemEShopResponse.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemEShopResponse
func (serviceItemEShopResponse *ServiceItemEShopResponse) Copy() types.RVType {
	copied := NewServiceItemEShopResponse()

	copied.StructureVersion = serviceItemEShopResponse.StructureVersion

	copied.HTTPStatus = serviceItemEShopResponse.HTTPStatus
	copied.ErrorCode = serviceItemEShopResponse.ErrorCode
	copied.CorrelationID = serviceItemEShopResponse.CorrelationID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemEShopResponse *ServiceItemEShopResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemEShopResponse); !ok {
		return false
	}

	other := o.(*ServiceItemEShopResponse)

	if serviceItemEShopResponse.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemEShopResponse.HTTPStatus.Equals(other.HTTPStatus) {
		return false
	}

	if !serviceItemEShopResponse.ErrorCode.Equals(other.ErrorCode) {
		return false
	}

	if !serviceItemEShopResponse.CorrelationID.Equals(other.CorrelationID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemEShopResponse *ServiceItemEShopResponse) String() string {
	return serviceItemEShopResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemEShopResponse *ServiceItemEShopResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEShopResponse{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemEShopResponse.StructureVersion))
	b.WriteString(fmt.Sprintf("%sHTTPStatus: %d,\n", indentationValues, serviceItemEShopResponse.HTTPStatus))
	b.WriteString(fmt.Sprintf("%sErrorCode: %d,\n", indentationValues, serviceItemEShopResponse.ErrorCode))
	b.WriteString(fmt.Sprintf("%sCorrelationID: %q,\n", indentationValues, serviceItemEShopResponse.CorrelationID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEShopResponse returns a new ServiceItemEShopResponse
func NewServiceItemEShopResponse() *ServiceItemEShopResponse {
	return &ServiceItemEShopResponse{}
}
