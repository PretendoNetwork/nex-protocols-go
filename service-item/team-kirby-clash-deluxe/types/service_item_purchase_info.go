// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPurchaseInfo holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseInfo struct {
	types.Structure
	TransactionID    string
	ExtTransactionID string
	ItemCode         string
	PostBalance      *ServiceItemAmount
}

// ExtractFrom extracts the ServiceItemPurchaseInfo from the given readable
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPurchaseInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPurchaseInfo header. %s", err.Error())
	}

	err = serviceItemPurchaseInfo.TransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.TransactionID from stream. %s", err.Error())
	}

	err = serviceItemPurchaseInfo.ExtTransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.ExtTransactionID from stream. %s", err.Error())
	}

	err = serviceItemPurchaseInfo.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemPurchaseInfo.PostBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseInfo.PostBalance from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemPurchaseInfo to the given writable
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPurchaseInfo.TransactionID.WriteTo(contentWritable)
	serviceItemPurchaseInfo.ExtTransactionID.WriteTo(contentWritable)
	serviceItemPurchaseInfo.ItemCode.WriteTo(contentWritable)
	serviceItemPurchaseInfo.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPurchaseInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPurchaseInfo
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) Copy() types.RVType {
	copied := NewServiceItemPurchaseInfo()

	copied.StructureVersion = serviceItemPurchaseInfo.StructureVersion

	copied.TransactionID = serviceItemPurchaseInfo.TransactionID
	copied.ExtTransactionID = serviceItemPurchaseInfo.ExtTransactionID
	copied.ItemCode = serviceItemPurchaseInfo.ItemCode
	copied.PostBalance = serviceItemPurchaseInfo.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseInfo *ServiceItemPurchaseInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseInfo)

	if serviceItemPurchaseInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPurchaseInfo.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !serviceItemPurchaseInfo.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !serviceItemPurchaseInfo.ItemCode.Equals(other.ItemCode) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPurchaseInfo.StructureVersion))
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
