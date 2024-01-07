// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPurchaceInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPurchaceInfo struct {
	types.Structure
	TransactionID    string
	ExtTransactionID string
	ItemCode         string
	PostBalance      *ServiceItemAmount
}

// ExtractFrom extracts the ServiceItemPurchaceInfo from the given readable
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPurchaceInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPurchaceInfo header. %s", err.Error())
	}

	err = serviceItemPurchaceInfo.TransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.TransactionID from stream. %s", err.Error())
	}

	err = serviceItemPurchaceInfo.ExtTransactionID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ExtTransactionID from stream. %s", err.Error())
	}

	err = serviceItemPurchaceInfo.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemPurchaceInfo.PostBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaceInfo.PostBalance from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemPurchaceInfo to the given writable
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPurchaceInfo.TransactionID.WriteTo(contentWritable)
	serviceItemPurchaceInfo.ExtTransactionID.WriteTo(contentWritable)
	serviceItemPurchaceInfo.ItemCode.WriteTo(contentWritable)
	serviceItemPurchaceInfo.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPurchaceInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPurchaceInfo
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) Copy() types.RVType {
	copied := NewServiceItemPurchaceInfo()

	copied.StructureVersion = serviceItemPurchaceInfo.StructureVersion

	copied.TransactionID = serviceItemPurchaceInfo.TransactionID
	copied.ExtTransactionID = serviceItemPurchaceInfo.ExtTransactionID
	copied.ItemCode = serviceItemPurchaceInfo.ItemCode
	copied.PostBalance = serviceItemPurchaceInfo.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaceInfo *ServiceItemPurchaceInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaceInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaceInfo)

	if serviceItemPurchaceInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPurchaceInfo.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !serviceItemPurchaceInfo.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !serviceItemPurchaceInfo.ItemCode.Equals(other.ItemCode) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPurchaceInfo.StructureVersion))
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
