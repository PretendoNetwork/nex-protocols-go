package shop_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ShopItemRights is a data structure used by the Shop protocol
type ShopItemRights struct {
	nex.Structure
	ReferenceID []byte
	ItemType    int8
	Attribute   uint32
}

// ExtractFromStream extracts a ShopItemRights structure from a stream
func (shopItemRights *ShopItemRights) ExtractFromStream(stream *nex.StreamIn) error {
	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ReferenceID from stream. %s", err.Error())
	}

	itemType, err := stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ItemType from stream. %s", err.Error())
	}

	attribute, err := stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.Attribute from stream. %s", err.Error())
	}

	shopItemRights.ReferenceID = referenceID
	shopItemRights.ItemType = itemType
	shopItemRights.Attribute = attribute

	return nil
}

// Bytes encodes the ShopItemRights and returns a byte array
func (shopItemRights *ShopItemRights) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(shopItemRights.ReferenceID)
	stream.WriteUInt8(uint8(shopItemRights.ItemType))
	stream.WriteUInt32LE(shopItemRights.Attribute)

	return stream.Bytes()
}

// Copy returns a new copied instance of ShopItemRights
func (shopItemRights *ShopItemRights) Copy() nex.StructureInterface {
	copied := NewShopItemRights()

	copied.ReferenceID = make([]byte, len(shopItemRights.ReferenceID))

	copy(copied.ReferenceID, shopItemRights.ReferenceID)

	copied.ItemType = shopItemRights.ItemType
	copied.Attribute = shopItemRights.Attribute

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItemRights *ShopItemRights) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ShopItemRights)

	if !bytes.Equal(shopItemRights.ReferenceID, other.ReferenceID) {
		return false
	}

	if shopItemRights.ItemType != other.ItemType {
		return false
	}

	if shopItemRights.Attribute != other.Attribute {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (shopItemRights *ShopItemRights) String() string {
	return shopItemRights.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (shopItemRights *ShopItemRights) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItem{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, shopItemRights.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sReferenceID: %x,\n", indentationValues, shopItemRights.ReferenceID))
	b.WriteString(fmt.Sprintf("%sItemType: %d,\n", indentationValues, shopItemRights.ItemType))
	b.WriteString(fmt.Sprintf("%sAttribute: %d\n", indentationValues, shopItemRights.Attribute))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItemRights returns a new ShopItemRights
func NewShopItemRights() *ShopItemRights {
	return &ShopItemRights{}
}
