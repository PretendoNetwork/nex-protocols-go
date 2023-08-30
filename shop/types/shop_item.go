// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ShopItem is a data structure used by the Shop protocol
type ShopItem struct {
	nex.Structure
	ItemID      uint32
	ReferenceID []byte
	ServiceName string
	ItemCode    string
}

// ExtractFromStream extracts a ShopItem structure from a stream
func (shopItem *ShopItem) ExtractFromStream(stream *nex.StreamIn) error {
	itemID, err := stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemID from stream. %s", err.Error())
	}

	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ReferenceID from stream. %s", err.Error())
	}

	serviceName, err := stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ServiceName from stream. %s", err.Error())
	}

	itemCode, err := stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemCode from stream. %s", err.Error())
	}

	shopItem.ItemID = itemID
	shopItem.ReferenceID = referenceID
	shopItem.ServiceName = serviceName
	shopItem.ItemCode = itemCode

	return nil
}

// Bytes encodes the ShopItem and returns a byte array
func (shopItem *ShopItem) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(shopItem.ItemID)
	stream.WriteQBuffer(shopItem.ReferenceID)
	stream.WriteString(shopItem.ServiceName)
	stream.WriteString(shopItem.ItemCode)

	return stream.Bytes()
}

// Copy returns a new copied instance of ShopItem
func (shopItem *ShopItem) Copy() nex.StructureInterface {
	copied := NewShopItem()

	copied.SetStructureVersion(shopItem.StructureVersion())

	copied.ItemID = shopItem.ItemID
	copied.ReferenceID = make([]byte, len(shopItem.ReferenceID))

	copy(copied.ReferenceID, shopItem.ReferenceID)

	copied.ServiceName = shopItem.ServiceName
	copied.ItemCode = shopItem.ItemCode

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItem *ShopItem) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ShopItem)

	if shopItem.StructureVersion() != other.StructureVersion() {
		return false
	}

	if shopItem.ItemID != other.ItemID {
		return false
	}

	if !bytes.Equal(shopItem.ReferenceID, other.ReferenceID) {
		return false
	}

	if shopItem.ServiceName != other.ServiceName {
		return false
	}

	if shopItem.ItemCode != other.ItemCode {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (shopItem *ShopItem) String() string {
	return shopItem.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (shopItem *ShopItem) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItem{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, shopItem.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sItemID: %d,\n", indentationValues, shopItem.ItemID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %x,\n", indentationValues, shopItem.ReferenceID))
	b.WriteString(fmt.Sprintf("%sServiceName: %q,\n", indentationValues, shopItem.ServiceName))
	b.WriteString(fmt.Sprintf("%sItemCode: %q\n", indentationValues, shopItem.ItemCode))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItem returns a new ShopItem
func NewShopItem() *ShopItem {
	return &ShopItem{}
}
