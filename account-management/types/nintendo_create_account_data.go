// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
)

// NintendoCreateAccountData contains data for creating a new NNID on the network
type NintendoCreateAccountData struct {
	nex.Structure
	NNAInfo  *friends_wiiu_types.NNAInfo
	Token    string
	Birthday *nex.DateTime
	Unknown  uint64
}

// ExtractFromStream extracts a NintendoCreateAccountData structure from a stream
func (nintendoCreateAccountData *NintendoCreateAccountData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	nnaInfo, err := stream.ReadStructure(friends_wiiu_types.NewNNAInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.NNAInfo from stream. %s", err.Error())
	}

	nintendoCreateAccountData.NNAInfo = nnaInfo.(*friends_wiiu_types.NNAInfo)
	nintendoCreateAccountData.Token, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Token from stream. %s", err.Error())
	}

	nintendoCreateAccountData.Birthday, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Birthday from stream. %s", err.Error())
	}

	nintendoCreateAccountData.Unknown, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoCreateAccountData.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoCreateAccountData
func (nintendoCreateAccountData *NintendoCreateAccountData) Copy() nex.StructureInterface {
	copied := NewNintendoCreateAccountData()

	if nintendoCreateAccountData.NNAInfo != nil {
		copied.NNAInfo = nintendoCreateAccountData.NNAInfo.Copy().(*friends_wiiu_types.NNAInfo)
	}

	copied.Token = nintendoCreateAccountData.Token

	if nintendoCreateAccountData.Birthday != nil {
		copied.Birthday = nintendoCreateAccountData.Birthday.Copy()
	}

	copied.Unknown = nintendoCreateAccountData.Unknown

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoCreateAccountData *NintendoCreateAccountData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoCreateAccountData)

	if nintendoCreateAccountData.NNAInfo == nil && other.NNAInfo != nil {
		return false
	}

	if nintendoCreateAccountData.NNAInfo != nil && other.NNAInfo == nil {
		return false
	}

	if !nintendoCreateAccountData.NNAInfo.Equals(other.NNAInfo) {
		return false
	}

	if nintendoCreateAccountData.Token != other.Token {
		return false
	}

	if nintendoCreateAccountData.Birthday == nil && other.Birthday != nil {
		return false
	}

	if nintendoCreateAccountData.Birthday != nil && other.Birthday == nil {
		return false
	}

	if !nintendoCreateAccountData.Birthday.Equals(other.Birthday) {
		return false
	}

	if nintendoCreateAccountData.Unknown != other.Unknown {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, nintendoCreateAccountData.StructureVersion()))

	if nintendoCreateAccountData.NNAInfo != nil {
		b.WriteString(fmt.Sprintf("%sNNAInfo: %s,\n", indentationValues, nintendoCreateAccountData.NNAInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sNNAInfo: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, nintendoCreateAccountData.Token))

	if nintendoCreateAccountData.Birthday != nil {
		b.WriteString(fmt.Sprintf("%sBirthday: %s,\n", indentationValues, nintendoCreateAccountData.Birthday.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBirthday: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown: %d\n", indentationValues, nintendoCreateAccountData.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoCreateAccountData returns a new NintendoCreateAccountData
func NewNintendoCreateAccountData() *NintendoCreateAccountData {
	return &NintendoCreateAccountData{}
}
