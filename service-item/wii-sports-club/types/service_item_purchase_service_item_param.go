// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPurchaseServiceItemParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPurchaseServiceItemParam struct {
	nex.Structure
	ItemCode string
	PriceID string
	ReferenceID string
	Balance string
	ItemName string
	EcServiceToken string
	Language string
	TitleID string
}

// ExtractFromStream extracts a ServiceItemPurchaseServiceItemParam structure from a stream
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPurchaseServiceItemParam.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemCode from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.PriceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.PriceID from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.ReferenceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ReferenceID from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.Balance, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Balance from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.ItemName, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemName from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.EcServiceToken, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.EcServiceToken from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.Language, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Language from stream. %s", err.Error())
	}

	serviceItemPurchaseServiceItemParam.TitleID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.TitleID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemPurchaseServiceItemParam and returns a byte array
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemPurchaseServiceItemParam.ItemCode)
	stream.WriteString(serviceItemPurchaseServiceItemParam.PriceID)
	stream.WriteString(serviceItemPurchaseServiceItemParam.ReferenceID)
	stream.WriteString(serviceItemPurchaseServiceItemParam.Balance)
	stream.WriteString(serviceItemPurchaseServiceItemParam.ItemName)
	stream.WriteString(serviceItemPurchaseServiceItemParam.EcServiceToken)
	stream.WriteString(serviceItemPurchaseServiceItemParam.Language)
	stream.WriteString(serviceItemPurchaseServiceItemParam.TitleID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemParam
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) Copy() nex.StructureInterface {
	copied := NewServiceItemPurchaseServiceItemParam()

	copied.ItemCode = serviceItemPurchaseServiceItemParam.ItemCode
	copied.PriceID = serviceItemPurchaseServiceItemParam.PriceID
	copied.ReferenceID = serviceItemPurchaseServiceItemParam.ReferenceID
	copied.Balance = serviceItemPurchaseServiceItemParam.Balance
	copied.ItemName = serviceItemPurchaseServiceItemParam.ItemName
	copied.EcServiceToken = serviceItemPurchaseServiceItemParam.EcServiceToken
	copied.Language = serviceItemPurchaseServiceItemParam.Language
	copied.TitleID = serviceItemPurchaseServiceItemParam.TitleID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPurchaseServiceItemParam)

	if serviceItemPurchaseServiceItemParam.ItemCode != other.ItemCode {
		return false
	}

	if serviceItemPurchaseServiceItemParam.PriceID != other.PriceID {
		return false
	}

	if serviceItemPurchaseServiceItemParam.ReferenceID != other.ReferenceID {
		return false
	}

	if serviceItemPurchaseServiceItemParam.Balance != other.Balance {
		return false
	}

	if serviceItemPurchaseServiceItemParam.ItemName != other.ItemName {
		return false
	}

	if serviceItemPurchaseServiceItemParam.EcServiceToken != other.EcServiceToken {
		return false
	}

	if serviceItemPurchaseServiceItemParam.Language != other.Language {
		return false
	}

	if serviceItemPurchaseServiceItemParam.TitleID != other.TitleID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) String() string {
	return serviceItemPurchaseServiceItemParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPurchaseServiceItemParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ItemCode))
	b.WriteString(fmt.Sprintf("%sPriceID: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.PriceID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sBalance: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.Balance))
	b.WriteString(fmt.Sprintf("%sItemName: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ItemName))
	b.WriteString(fmt.Sprintf("%sEcServiceToken: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.EcServiceToken))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.Language))
	b.WriteString(fmt.Sprintf("%sTitleID: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.TitleID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseServiceItemParam returns a new ServiceItemPurchaseServiceItemParam
func NewServiceItemPurchaseServiceItemParam() *ServiceItemPurchaseServiceItemParam {
	return &ServiceItemPurchaseServiceItemParam{}
}
