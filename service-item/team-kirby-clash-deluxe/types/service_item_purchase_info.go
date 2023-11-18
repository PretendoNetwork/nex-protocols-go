// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaseInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseInfo struct {
	nex.Structure
	TransactionID    string
	ExtTransactionID string
	ItemCode         string
	PostBalance      *ServiceItemAmount
}

// ExtractFromStream extracts a ServiceItemPurchaseInfo structure from a stream
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPurchaseInfo.TransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.TransactionID from stream. %s", err.Error())
	}

	serviceItemPurchaseInfo.ExtTransactionID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.ExtTransactionID from stream. %s", err.Error())
	}

	serviceItemPurchaseInfo.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.ItemCode from stream. %s", err.Error())
	}

	serviceItemPurchaseInfo.PostBalance, err = nex.StreamReadStructure(stream, NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.PostBalance from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemPurchaseInfo and returns a byte array
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemPurchaseInfo.TransactionID)
	stream.WriteString(serviceItemPurchaseInfo.ExtTransactionID)
	stream.WriteString(serviceItemPurchaseInfo.ItemCode)
	stream.WriteStructure(serviceItemPurchaseInfo.PostBalance)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaseInfo
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaseInfo()

	copied.SetStructureVersion(serviceItemPurchaseInfo.StructureVersion())

	copied.TransactionID = serviceItemPurchaseInfo.TransactionID
	copied.ExtTransactionID = serviceItemPurchaseInfo.ExtTransactionID
	copied.ItemCode = serviceItemPurchaseInfo.ItemCode
	copied.PostBalance = serviceItemPurchaseInfo.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaseInfo)

	if serviceItemPurchaseInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemPurchaseInfo.TransactionID != other.TransactionID {
		return false
	}

	if serviceItemPurchaseInfo.ExtTransactionID != other.ExtTransactionID {
		return false
	}

	if serviceItemPurchaseInfo.ItemCode != other.ItemCode {
		return false
	}

	if !serviceItemPurchaseInfo.PostBalance.Equals(other.PostBalance) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) String() string {
	return serviceItemPurchaseInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPurchaseInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTransactionID: %q,\n", indentationValues, serviceItemPurchaseInfo.TransactionID))
	b.WriteString(fmt.Sprintf("%sExtTransactionID: %q,\n", indentationValues, serviceItemPurchaseInfo.ExtTransactionID))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemPurchaseInfo.ItemCode))

	if serviceItemPurchaseInfo.PostBalance != nil {
		b.WriteString(fmt.Sprintf("%sPostBalance: %s\n", indentationValues, serviceItemPurchaseInfo.PostBalance.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPostBalance: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseInfo returns a new ServiceItemPurchaseInfo
func NewServiceItemPurchaseInfo() *ServiceItemPurchaseInfo {
	return &ServiceItemPurchaseInfo{}
}
