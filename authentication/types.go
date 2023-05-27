package authentication

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// NintendoLoginData holds a nex auth token
type NintendoLoginData struct {
	nex.Structure
	Token string
}

// ExtractFromStream extracts a AuthenticationInfo structure from a stream
func (nintendoLoginData *NintendoLoginData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error
	var token string

	token, err = stream.ReadString()

	if err != nil {
		return err
	}

	nintendoLoginData.Token = token

	return nil
}

// Copy returns a new copied instance of NintendoLoginData
func (nintendoLoginData *NintendoLoginData) Copy() nex.StructureInterface {
	copied := NewNintendoLoginData()

	copied.Token = nintendoLoginData.Token

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (nintendoLoginData *NintendoLoginData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoLoginData)

	return nintendoLoginData.Token == other.Token
}

// NewAuthenticationInfo returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	return &NintendoLoginData{}
}

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
	var token string

	token, err = stream.ReadString()

	if err != nil {
		return err
	}

	if len(stream.Bytes()[stream.ByteOffset():]) < 9 {
		globals.Logger.Error("Data size too small")
		return nil //technically not needed (for now) and was causing some strangeness with MK7
	}

	authenticationInfo.Token = token
	authenticationInfo.TokenType = stream.ReadUInt8()
	authenticationInfo.NGSVersion = stream.ReadUInt32LE()
	authenticationInfo.ServerVersion = stream.ReadUInt32LE()

	return nil
}

// Copy returns a new copied instance of AuthenticationInfo
func (authenticationInfo *AuthenticationInfo) Copy() nex.StructureInterface {
	copied := NewAuthenticationInfo()

	copied.SetParentType(authenticationInfo.ParentType().Copy())
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

// NewAuthenticationInfo returns a new AuthenticationInfo
func NewAuthenticationInfo() *AuthenticationInfo {
	authenticationInfo := &AuthenticationInfo{}
	authenticationInfo.Data = nex.NewData()
	authenticationInfo.SetParentType(authenticationInfo.Data)

	return authenticationInfo
}
