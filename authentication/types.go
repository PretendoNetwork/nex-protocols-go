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

// NewAuthenticationInfo returns a new NintendoLoginData
func NewNintendoLoginData() *NintendoLoginData {
	return &NintendoLoginData{}
}

// AuthenticationInfo holds information about an authentication request
type AuthenticationInfo struct {
	*nex.Data
	hierarchy     []nex.StructureInterface
	Token         string
	NGSVersion    uint32
	TokenType     uint8
	ServerVersion uint32
}

// Hierarchy returns the Structure hierarchy
func (authenticationInfo *AuthenticationInfo) Hierarchy() []nex.StructureInterface {
	return authenticationInfo.hierarchy
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

// NewAuthenticationInfo returns a new AuthenticationInfo
func NewAuthenticationInfo() *AuthenticationInfo {
	data := nex.NewData()

	authenticationInfo := &AuthenticationInfo{}
	authenticationInfo.Data = data
	authenticationInfo.hierarchy = []nex.StructureInterface{data}

	return authenticationInfo
}
