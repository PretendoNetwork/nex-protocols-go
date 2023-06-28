package shop_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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

// NewShopItem returns a new ShopItem
func NewShopItem() *ShopItem {
	return &ShopItem{}
}
