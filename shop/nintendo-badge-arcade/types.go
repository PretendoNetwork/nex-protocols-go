package shop_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

type ShopPostPlayLogParam struct {
	nex.Structure
	Unknown1  []uint32
	Timestamp *nex.DateTime
	Unknown2  string
}

// ExtractFromStream extracts a ShopPostPlayLogParam structure from a stream
func (shopPostPlayLogParam *ShopPostPlayLogParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	shopPostPlayLogParam.Unknown1, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown1 from stream. %s", err.Error())
	}

	shopPostPlayLogParam.Timestamp, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Timestamp from stream. %s", err.Error())
	}

	shopPostPlayLogParam.Unknown2, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown2 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ShopPostPlayLogParam and returns a byte array
func (shopPostPlayLogParam *ShopPostPlayLogParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(shopPostPlayLogParam.Unknown1)
	stream.WriteDateTime(shopPostPlayLogParam.Timestamp)
	stream.WriteString(shopPostPlayLogParam.Unknown2)

	return stream.Bytes()
}

// Copy returns a new copied instance of ShopPostPlayLogParam
func (shopPostPlayLogParam *ShopPostPlayLogParam) Copy() nex.StructureInterface {
	copied := NewShopPostPlayLogParam()

	copied.Unknown1 = make([]uint32, len(shopPostPlayLogParam.Unknown1))

	copy(copied.Unknown1, shopPostPlayLogParam.Unknown1)

	copied.Timestamp = shopPostPlayLogParam.Timestamp.Copy()
	copied.Unknown2 = shopPostPlayLogParam.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopPostPlayLogParam *ShopPostPlayLogParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ShopPostPlayLogParam)

	if len(shopPostPlayLogParam.Unknown1) != len(other.Unknown1) {
		return false
	}

	for i := 0; i < len(shopPostPlayLogParam.Unknown1); i++ {
		if shopPostPlayLogParam.Unknown1[i] != other.Unknown1[i] {
			return false
		}
	}

	if !shopPostPlayLogParam.Timestamp.Equals(other.Timestamp) {
		return false
	}

	if shopPostPlayLogParam.Unknown2 != other.Unknown2 {
		return false
	}

	return true
}

// NewShopPostPlayLogParam returns a new ShopPostPlayLogParam
func NewShopPostPlayLogParam() *ShopPostPlayLogParam {
	return &ShopPostPlayLogParam{}
}
