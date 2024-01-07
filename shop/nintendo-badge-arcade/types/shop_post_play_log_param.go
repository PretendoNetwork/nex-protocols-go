// Package types implements all the types used by the Shop Nintendo Badge Arcade protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ShopPostPlayLogParam is a data structure used by the Nintendo Badge Arcade Shop protocol
type ShopPostPlayLogParam struct {
	types.Structure
	Unknown1  *types.List[*types.PrimitiveU32]
	Timestamp *types.DateTime
	Unknown2  string
}

// ExtractFrom extracts the ShopPostPlayLogParam from the given readable
func (shopPostPlayLogParam *ShopPostPlayLogParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = shopPostPlayLogParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ShopPostPlayLogParam header. %s", err.Error())
	}

	err = shopPostPlayLogParam.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown1 from stream. %s", err.Error())
	}

	err = shopPostPlayLogParam.Timestamp.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Timestamp from stream. %s", err.Error())
	}

	err = shopPostPlayLogParam.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ShopPostPlayLogParam.Unknown2 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ShopPostPlayLogParam to the given writable
func (shopPostPlayLogParam *ShopPostPlayLogParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	shopPostPlayLogParam.Unknown1.WriteTo(contentWritable)
	shopPostPlayLogParam.Timestamp.WriteTo(contentWritable)
	shopPostPlayLogParam.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	shopPostPlayLogParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ShopPostPlayLogParam
func (shopPostPlayLogParam *ShopPostPlayLogParam) Copy() types.RVType {
	copied := NewShopPostPlayLogParam()

	copied.StructureVersion = shopPostPlayLogParam.StructureVersion

	copied.Unknown1 = make(*types.List[*types.PrimitiveU32], len(shopPostPlayLogParam.Unknown1))

	copy(copied.Unknown1, shopPostPlayLogParam.Unknown1)

	copied.Timestamp = shopPostPlayLogParam.Timestamp.Copy()
	copied.Unknown2 = shopPostPlayLogParam.Unknown2

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopPostPlayLogParam *ShopPostPlayLogParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopPostPlayLogParam); !ok {
		return false
	}

	other := o.(*ShopPostPlayLogParam)

	if shopPostPlayLogParam.StructureVersion != other.StructureVersion {
		return false
	}

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

	if !shopPostPlayLogParam.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (shopPostPlayLogParam *ShopPostPlayLogParam) String() string {
	return shopPostPlayLogParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (shopPostPlayLogParam *ShopPostPlayLogParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopPostPlayLogParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, shopPostPlayLogParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %v,\n", indentationValues, shopPostPlayLogParam.Unknown1))

	if shopPostPlayLogParam.Timestamp != nil {
		b.WriteString(fmt.Sprintf("%sTimestamp: %s,\n", indentationValues, shopPostPlayLogParam.Timestamp.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTimestamp: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown2: %q\n", indentationValues, shopPostPlayLogParam.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopPostPlayLogParam returns a new ShopPostPlayLogParam
func NewShopPostPlayLogParam() *ShopPostPlayLogParam {
	return &ShopPostPlayLogParam{}
}
