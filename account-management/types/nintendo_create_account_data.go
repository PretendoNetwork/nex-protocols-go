// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
)

// NintendoCreateAccountData contains data for creating a new NNID on the network
type NintendoCreateAccountData struct {
	types.Structure
	NNAInfo  *friends_wiiu_types.NNAInfo
	Token    *types.String
	Birthday *types.DateTime
	Unknown  *types.PrimitiveU64
}

// WriteTo writes the NintendoCreateAccountData to the given writable
func (nintendoCreateAccountData *NintendoCreateAccountData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	nintendoCreateAccountData.NNAInfo.WriteTo(contentWritable)
	nintendoCreateAccountData.Token.WriteTo(contentWritable)
	nintendoCreateAccountData.Birthday.WriteTo(contentWritable)
	nintendoCreateAccountData.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	nintendoCreateAccountData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoCreateAccountData from the given readable
func (nintendoCreateAccountData *NintendoCreateAccountData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = nintendoCreateAccountData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read NintendoCreateAccountData header. %s", err.Error())
	}

	err = nintendoCreateAccountData.NNAInfo.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.NNAInfo from stream. %s", err.Error())
	}

	err = nintendoCreateAccountData.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Token from stream. %s", err.Error())
	}

	err = nintendoCreateAccountData.Birthday.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Birthday from stream. %s", err.Error())
	}

	err = nintendoCreateAccountData.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoCreateAccountData
func (nintendoCreateAccountData *NintendoCreateAccountData) Copy() types.RVType {
	copied := NewNintendoCreateAccountData()

	copied.StructureVersion = nintendoCreateAccountData.StructureVersion

	copied.NNAInfo = nintendoCreateAccountData.NNAInfo.Copy().(*friends_wiiu_types.NNAInfo)

	copied.Token = nintendoCreateAccountData.Token.Copy().(*types.String)

	copied.Birthday = nintendoCreateAccountData.Birthday.Copy().(*types.DateTime)

	copied.Unknown = nintendoCreateAccountData.Unknown.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoCreateAccountData *NintendoCreateAccountData) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoCreateAccountData); !ok {
		return false
	}

	other := o.(*NintendoCreateAccountData)

	if nintendoCreateAccountData.StructureVersion != other.StructureVersion {
		return false
	}

	if !nintendoCreateAccountData.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if !nintendoCreateAccountData.Token.Equals(other.Token) {
		return false
	}

	if !nintendoCreateAccountData.Birthday.Equals(other.Birthday) {
		return false
	}

	if !nintendoCreateAccountData.Unknown.Equals(other.Unknown) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (nintendoCreateAccountData *NintendoCreateAccountData) String() string {
	return nintendoCreateAccountData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (nintendoCreateAccountData *NintendoCreateAccountData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoCreateAccountData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, nintendoCreateAccountData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNNAInfo: %s,\n", indentationValues, nintendoCreateAccountData.NNAInfo.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, nintendoCreateAccountData.Token.Value))
	b.WriteString(fmt.Sprintf("%sBirthday: %s,\n", indentationValues, nintendoCreateAccountData.Birthday.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s\n", indentationValues, nintendoCreateAccountData.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoCreateAccountData returns a new NintendoCreateAccountData
func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	return &NintendoCreateAccountData{
		NNAInfo: friends_wiiu_types.NewNNAInfo(),
		Token: types.NewString(""),
		Birthday: types.NewDateTime(0),
		Unknown: types.NewPrimitiveU64(0),
	}
}
