package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// AuthenticationTokenType denotes the type of authentication token
// the client is using.
//
// Note: This is not a real type. This is a bespoke type created
// for our convenience.
type AuthenticationTokenType uint8

const (
	// AuthenticationTokenTypeNASC means that the connecting client used NASC
	// to authenticate themselves.
	AuthenticationTokenTypeNASC AuthenticationTokenType = iota

	// AuthenticationTokenTypeNNAS means that the connecting client used NNAS
	// to authenticate themselves.
	AuthenticationTokenTypeNNAS

	// AuthenticationTokenTypeSwitch means that the connecting client used the
	// Nintendo Switch authentication servers to authenticate themselves.
	AuthenticationTokenTypeSwitch
)

// WriteTo writes the AuthenticationTokenType to the given writable
func (att AuthenticationTokenType) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(att))
}

// ExtractFrom extracts the AuthenticationTokenType value from the given readable
func (att *AuthenticationTokenType) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*att = AuthenticationTokenType(value)
	return nil
}
