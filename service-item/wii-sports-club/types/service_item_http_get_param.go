// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemHTTPGetParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemHTTPGetParam struct {
	nex.Structure
	URL string
}

// ExtractFromStream extracts a ServiceItemHTTPGetParam structure from a stream
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemHTTPGetParam.URL, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetParam.URL from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemHTTPGetParam and returns a byte array
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemHTTPGetParam.URL)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemHTTPGetParam
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) Copy() nex.StructureInterface {
	copied := NewServiceItemHTTPGetParam()

	copied.URL = serviceItemHTTPGetParam.URL

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemHTTPGetParam)

	return serviceItemHTTPGetParam.URL == other.URL
}

// String returns a string representation of the struct
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) String() string {
	return serviceItemHTTPGetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemHTTPGetParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, serviceItemHTTPGetParam.URL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetParam returns a new ServiceItemHTTPGetParam
func NewServiceItemHTTPGetParam() *ServiceItemHTTPGetParam {
	return &ServiceItemHTTPGetParam{}
}
