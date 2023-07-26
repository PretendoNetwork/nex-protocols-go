// Package types implements all the types used by the Ticket Granting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// AuthenticationInfo holds information about an authentication request
type AuthenticationInfo struct {
	nex.Structure
	*nex.Data
	Token         string
	NGSVersion    uint32
	TokenType     uint8
	ServerVersion uint32
}

// ExtractFromStream extracts a AuthenticationInfo structure from a stream
func (authenticationInfo *AuthenticationInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	authenticationInfo.Token, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Token. %s", err.Error())
	}

	authenticationInfo.NGSVersion, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.NGSVersion. %s", err.Error())
	}

	if authenticationInfo.NGSVersion > 2 {
		authenticationInfo.TokenType, err = stream.ReadUInt8()
		if err != nil {
			return fmt.Errorf("Failed to extract AccountExtraInfo.TokenType. %s", err.Error())
		}

		authenticationInfo.ServerVersion, err = stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract AccountExtraInfo.ServerVersion. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of AuthenticationInfo
func (authenticationInfo *AuthenticationInfo) Copy() nex.StructureInterface {
	copied := NewAuthenticationInfo()

	copied.Data = authenticationInfo.Data.Copy().(*nex.Data)
	copied.SetParentType(copied.Data)
	copied.Token = authenticationInfo.Token
	copied.TokenType = authenticationInfo.TokenType
	copied.NGSVersion = authenticationInfo.NGSVersion
	copied.ServerVersion = authenticationInfo.ServerVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (authenticationInfo *AuthenticationInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*AuthenticationInfo)

	if !authenticationInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if authenticationInfo.Token != other.Token {
		return false
	}

	if authenticationInfo.TokenType != other.TokenType {
		return false
	}

	if authenticationInfo.NGSVersion != other.NGSVersion {
		return false
	}

	if authenticationInfo.ServerVersion != other.ServerVersion {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (authenticationInfo *AuthenticationInfo) String() string {
	return authenticationInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (authenticationInfo *AuthenticationInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("AuthenticationInfo{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, authenticationInfo.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, authenticationInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sToken: %s,\n", indentationValues, authenticationInfo.Token))
	b.WriteString(fmt.Sprintf("%sTokenType: %d,\n", indentationValues, authenticationInfo.TokenType))
	b.WriteString(fmt.Sprintf("%sNGSVersion: %d,\n", indentationValues, authenticationInfo.NGSVersion))
	b.WriteString(fmt.Sprintf("%sServerVersion: %d\n", indentationValues, authenticationInfo.ServerVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewAuthenticationInfo returns a new AuthenticationInfo
func NewAuthenticationInfo() *AuthenticationInfo {
	authenticationInfo := &AuthenticationInfo{}
	authenticationInfo.Data = nex.NewData()
	authenticationInfo.SetParentType(authenticationInfo.Data)

	return authenticationInfo
}
