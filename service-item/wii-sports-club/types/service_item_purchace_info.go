// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaceInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPurchaceInfo struct {
	nex.Structure
	TransactionID    string
	ExtTransactionID string
	ItemCode         string
	PostBalance      *ServiceItemAmount
}

// ExtractFromStream extracts a ServiceItemPurchaceInfo structure from a stream
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPurchaceInfo.TransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.TransactionID from stream. %s", err.Error())
	}

	serviceItemPurchaceInfo.ExtTransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ExtTransactionID from stream. %s", err.Error())
	}

	serviceItemPurchaceInfo.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ItemCode from stream. %s", err.Error())
	}

	postBalance, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.PostBalance from stream. %s", err.Error())
	}

	serviceItemPurchaceInfo.PostBalance = postBalance.(*ServiceItemAmount)

	return nil
}

// Bytes encodes the ServiceItemPurchaceInfo and returns a byte array
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemPurchaceInfo.TransactionID)
	stream.WriteString(serviceItemPurchaceInfo.ExtTransactionID)
	stream.WriteString(serviceItemPurchaceInfo.ItemCode)
	stream.WriteStructure(serviceItemPurchaceInfo.PostBalance)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaceInfo
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaceInfo()

	copied.SetStructureVersion(serviceItemPurchaceInfo.StructureVersion())

	copied.TransactionID = serviceItemPurchaceInfo.TransactionID
	copied.ExtTransactionID = serviceItemPurchaceInfo.ExtTransactionID
	copied.ItemCode = serviceItemPurchaceInfo.ItemCode
	copied.PostBalance = serviceItemPurchaceInfo.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaceInfo)

	if serviceItemPurchaceInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemPurchaceInfo.TransactionID != other.TransactionID {
		return false
	}

	if serviceItemPurchaceInfo.ExtTransactionID != other.ExtTransactionID {
		return false
	}

	if serviceItemPurchaceInfo.ItemCode != other.ItemCode {
		return false
	}

	if !serviceItemPurchaceInfo.PostBalance.Equals(other.PostBalance) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) String() string {
	return serviceItemPurchaceInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaceInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPurchaceInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTransactionID: %q,\n", indentationValues, serviceItemPurchaceInfo.TransactionID))
	b.WriteString(fmt.Sprintf("%sExtTransactionID: %q,\n", indentationValues, serviceItemPurchaceInfo.ExtTransactionID))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemPurchaceInfo.ItemCode))

	if serviceItemPurchaceInfo.PostBalance != nil {
		b.WriteString(fmt.Sprintf("%sPostBalance: %s\n", indentationValues, serviceItemPurchaceInfo.PostBalance.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPostBalance: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaceInfo returns a new ServiceItemPurchaceInfo
func NewServiceItemPurchaceInfo() *ServiceItemPurchaceInfo {
	return &ServiceItemPurchaceInfo{}
}
