// Package types implements all the types used by the TicketGranting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// AuthenticationInfo is a type within the TicketGranting protocol
type AuthenticationInfo struct {
	types.Structure
	types.Data
	Token         types.String
	NGSVersion    types.UInt32
	TokenType     types.UInt8
	ServerVersion types.UInt32
}

// ObjectID returns the object identifier of the type
func (ai AuthenticationInfo) ObjectID() types.RVType {
	return ai.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (ai AuthenticationInfo) DataObjectID() types.RVType {
	return types.NewString("AuthenticationInfo")
}

// WriteTo writes the AuthenticationInfo to the given writable
func (ai AuthenticationInfo) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.Main

	ai.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	ai.Token.WriteTo(contentWritable)
	ai.NGSVersion.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("3.0.0") {
		ai.TokenType.WriteTo(contentWritable)
		ai.ServerVersion.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	ai.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the AuthenticationInfo from the given readable
func (ai *AuthenticationInfo) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.Main

	var err error

	err = ai.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AuthenticationInfo.Data. %s", err.Error())
	}

	err = ai.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AuthenticationInfo header. %s", err.Error())
	}

	err = ai.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AuthenticationInfo.Token. %s", err.Error())
	}

	err = ai.NGSVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AuthenticationInfo.NGSVersion. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("3.0.0") {
		err = ai.TokenType.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract AuthenticationInfo.TokenType. %s", err.Error())
		}

		err = ai.ServerVersion.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract AuthenticationInfo.ServerVersion. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of AuthenticationInfo
func (ai AuthenticationInfo) Copy() types.RVType {
	copied := NewAuthenticationInfo()

	copied.StructureVersion = ai.StructureVersion
	copied.Data = ai.Data.Copy().(types.Data)
	copied.Token = ai.Token.Copy().(types.String)
	copied.NGSVersion = ai.NGSVersion.Copy().(types.UInt32)
	copied.TokenType = ai.TokenType.Copy().(types.UInt8)
	copied.ServerVersion = ai.ServerVersion.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given AuthenticationInfo contains the same data as the current AuthenticationInfo
func (ai AuthenticationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(AuthenticationInfo); !ok {
		return false
	}

	other := o.(AuthenticationInfo)

	if ai.StructureVersion != other.StructureVersion {
		return false
	}

	if !ai.Data.Equals(other.Data) {
		return false
	}

	if !ai.Token.Equals(other.Token) {
		return false
	}

	if !ai.NGSVersion.Equals(other.NGSVersion) {
		return false
	}

	if !ai.TokenType.Equals(other.TokenType) {
		return false
	}

	return ai.ServerVersion.Equals(other.ServerVersion)
}

// CopyRef copies the current value of the AuthenticationInfo
// and returns a pointer to the new copy
func (ai AuthenticationInfo) CopyRef() types.RVTypePtr {
	copied := ai.Copy().(AuthenticationInfo)
	return &copied
}

// Deref takes a pointer to the AuthenticationInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ai *AuthenticationInfo) Deref() types.RVType {
	return *ai
}

// String returns the string representation of the AuthenticationInfo
func (ai AuthenticationInfo) String() string {
	return ai.FormatToString(0)
}

// FormatToString pretty-prints the AuthenticationInfo using the provided indentation level
func (ai AuthenticationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AuthenticationInfo{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, ai.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, ai.Token))
	b.WriteString(fmt.Sprintf("%sNGSVersion: %s,\n", indentationValues, ai.NGSVersion))
	b.WriteString(fmt.Sprintf("%sTokenType: %s,\n", indentationValues, ai.TokenType))
	b.WriteString(fmt.Sprintf("%sServerVersion: %s,\n", indentationValues, ai.ServerVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAuthenticationInfo returns a new AuthenticationInfo
func NewAuthenticationInfo() AuthenticationInfo {
	return AuthenticationInfo{
		Data:          types.NewData(),
		Token:         types.NewString(""),
		NGSVersion:    types.NewUInt32(0),
		TokenType:     types.NewUInt8(0),
		ServerVersion: types.NewUInt32(0),
	}

}
