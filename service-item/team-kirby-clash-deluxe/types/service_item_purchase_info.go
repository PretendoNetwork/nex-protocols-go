// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseInfo is a type within the ServiceItem protocol
type ServiceItemPurchaseInfo struct {
	types.Structure
	TransactionID    types.String
	ExtTransactionID types.String
	ItemCode         types.String
	PostBalance      ServiceItemAmount
}

// WriteTo writes the ServiceItemPurchaseInfo to the given writable
func (sipi ServiceItemPurchaseInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipi.TransactionID.WriteTo(contentWritable)
	sipi.ExtTransactionID.WriteTo(contentWritable)
	sipi.ItemCode.WriteTo(contentWritable)
	sipi.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseInfo from the given readable
func (sipi *ServiceItemPurchaseInfo) ExtractFrom(readable types.Readable) error {
	if err := sipi.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseInfo header. %s", err.Error())
	}

	if err := sipi.TransactionID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseInfo.TransactionID. %s", err.Error())
	}

	if err := sipi.ExtTransactionID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseInfo.ExtTransactionID. %s", err.Error())
	}

	if err := sipi.ItemCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseInfo.ItemCode. %s", err.Error())
	}

	if err := sipi.PostBalance.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseInfo.PostBalance. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseInfo
func (sipi ServiceItemPurchaseInfo) Copy() types.RVType {
	copied := NewServiceItemPurchaseInfo()

	copied.StructureVersion = sipi.StructureVersion
	copied.TransactionID = sipi.TransactionID.Copy().(types.String)
	copied.ExtTransactionID = sipi.ExtTransactionID.Copy().(types.String)
	copied.ItemCode = sipi.ItemCode.Copy().(types.String)
	copied.PostBalance = sipi.PostBalance.Copy().(ServiceItemAmount)

	return copied
}

// Equals checks if the given ServiceItemPurchaseInfo contains the same data as the current ServiceItemPurchaseInfo
func (sipi ServiceItemPurchaseInfo) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemPurchaseInfo); !ok {
		return false
	}

	other := o.(ServiceItemPurchaseInfo)

	if sipi.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipi.TransactionID.Equals(other.TransactionID) {
		return false
	}

	if !sipi.ExtTransactionID.Equals(other.ExtTransactionID) {
		return false
	}

	if !sipi.ItemCode.Equals(other.ItemCode) {
		return false
	}

	return sipi.PostBalance.Equals(other.PostBalance)
}

// CopyRef copies the current value of the ServiceItemPurchaseInfo
// and returns a pointer to the new copy
func (sipi ServiceItemPurchaseInfo) CopyRef() types.RVTypePtr {
	copied := sipi.Copy().(ServiceItemPurchaseInfo)
	return &copied
}

// Deref takes a pointer to the ServiceItemPurchaseInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sipi *ServiceItemPurchaseInfo) Deref() types.RVType {
	return *sipi
}

// String returns the string representation of the ServiceItemPurchaseInfo
func (sipi ServiceItemPurchaseInfo) String() string {
	return sipi.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseInfo using the provided indentation level
func (sipi ServiceItemPurchaseInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseInfo{\n")
	fmt.Fprintf(&b, "%sTransactionID: %s,\n", indentationValues, sipi.TransactionID)
	fmt.Fprintf(&b, "%sExtTransactionID: %s,\n", indentationValues, sipi.ExtTransactionID)
	fmt.Fprintf(&b, "%sItemCode: %s,\n", indentationValues, sipi.ItemCode)
	fmt.Fprintf(&b, "%sPostBalance: %s,\n", indentationValues, sipi.PostBalance.FormatToString(indentationLevel+1))
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewServiceItemPurchaseInfo returns a new ServiceItemPurchaseInfo
func NewServiceItemPurchaseInfo() ServiceItemPurchaseInfo {
	return ServiceItemPurchaseInfo{
		TransactionID:    types.NewString(""),
		ExtTransactionID: types.NewString(""),
		ItemCode:         types.NewString(""),
		PostBalance:      NewServiceItemAmount(),
	}

}
