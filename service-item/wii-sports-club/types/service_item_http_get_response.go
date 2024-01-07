// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemHTTPGetResponse holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemHTTPGetResponse struct {
	types.Structure
	Response []byte
}

// ExtractFrom extracts the ServiceItemHTTPGetResponse from the given readable
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemHTTPGetResponse.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemHTTPGetResponse header. %s", err.Error())
	}

	serviceItemHTTPGetResponse.Response, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetResponse.Response from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemHTTPGetResponse to the given writable
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	stream.WriteQBuffer(serviceItemHTTPGetResponse.Response)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemHTTPGetResponse
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) Copy() types.RVType {
	copied := NewServiceItemHTTPGetResponse()

	copied.StructureVersion = serviceItemHTTPGetResponse.StructureVersion

	copied.Response = serviceItemHTTPGetResponse.Response

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemHTTPGetResponse); !ok {
		return false
	}

	other := o.(*ServiceItemHTTPGetResponse)

	if serviceItemHTTPGetResponse.StructureVersion != other.StructureVersion {
		return false
	}

	return bytes.Equal(serviceItemHTTPGetResponse.Response, other.Response)
}

// String returns a string representation of the struct
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) String() string {
	return serviceItemHTTPGetResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetResponse{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemHTTPGetResponse.StructureVersion))
	b.WriteString(fmt.Sprintf("%sResponse: %x,\n", indentationValues, serviceItemHTTPGetResponse.Response))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetResponse returns a new ServiceItemHTTPGetResponse
func NewServiceItemHTTPGetResponse() *ServiceItemHTTPGetResponse {
	return &ServiceItemHTTPGetResponse{}
}
