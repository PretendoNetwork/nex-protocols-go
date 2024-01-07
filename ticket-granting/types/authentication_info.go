// Package types implements all the types used by the Ticket Granting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// AuthenticationInfo holds information about an authentication request
type AuthenticationInfo struct {
	types.Structure
	*types.Data
	Token         string
	NGSVersion    *types.PrimitiveU32
	TokenType     *types.PrimitiveU8
	ServerVersion *types.PrimitiveU32
}

// ExtractFrom extracts the AuthenticationInfo from the given readable
func (authenticationInfo *AuthenticationInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = authenticationInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read AuthenticationInfo header. %s", err.Error())
	}

	err = authenticationInfo.Token.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.Token. %s", err.Error())
	}

	err = authenticationInfo.NGSVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract AccountExtraInfo.NGSVersion. %s", err.Error())
	}

	if authenticationInfo.NGSVersion > 2 {
	err = 	authenticationInfo.TokenType.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract AccountExtraInfo.TokenType. %s", err.Error())
		}

	err = 	authenticationInfo.ServerVersion.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract AccountExtraInfo.ServerVersion. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of AuthenticationInfo
func (authenticationInfo *AuthenticationInfo) Copy() types.RVType {
	copied := NewAuthenticationInfo()

	copied.StructureVersion = authenticationInfo.StructureVersion

	copied.Data = authenticationInfo.Data.Copy().(*types.Data)
	copied.Token = authenticationInfo.Token
	copied.TokenType = authenticationInfo.TokenType
	copied.NGSVersion = authenticationInfo.NGSVersion
	copied.ServerVersion = authenticationInfo.ServerVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (authenticationInfo *AuthenticationInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*AuthenticationInfo); !ok {
		return false
	}

	other := o.(*AuthenticationInfo)

	if authenticationInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !authenticationInfo.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !authenticationInfo.Token.Equals(other.Token) {
		return false
	}

	if !authenticationInfo.TokenType.Equals(other.TokenType) {
		return false
	}

	if !authenticationInfo.NGSVersion.Equals(other.NGSVersion) {
		return false
	}

	if !authenticationInfo.ServerVersion.Equals(other.ServerVersion) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, authenticationInfo.StructureVersion))
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
	authenticationInfo.Data = types.NewData()
	authenticationInfo.SetParentType(authenticationInfo.Data)

	return authenticationInfo
}
