package match_making_types

import "github.com/PretendoNetwork/nex-go"

// MatchmakeParam holds parameters for a matchmake session
type MatchmakeParam struct {
	nex.Structure
	Parameters map[string]*nex.Variant
}

// ExtractFromStream extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) ExtractFromStream(stream *nex.StreamIn) error {
	parameters, err := stream.ReadMap(stream.ReadString, stream.ReadVariant)

	if err != nil {
		return err
	}

	matchmakeParam.Parameters = make(map[string]*nex.Variant, len(parameters))

	for key, value := range parameters {
		matchmakeParam.Parameters[key.(string)] = value.(*nex.Variant)
	}

	return nil
}

// Bytes extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteMap(matchmakeParam.Parameters)

	return stream.Bytes()
}

// Copy returns a new copied instance of MatchmakeParam
func (matchmakeParam *MatchmakeParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeParam()

	copied.Parameters = make(map[string]*nex.Variant, len(matchmakeParam.Parameters))

	for key, value := range matchmakeParam.Parameters {
		copied.Parameters[key] = value.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeParam *MatchmakeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeParam)

	if len(matchmakeParam.Parameters) != len(other.Parameters) {
		return false
	}

	for key, value := range matchmakeParam.Parameters {
		if !value.Equals(other.Parameters[key]) {
			return false
		}
	}

	return true
}

// NewMatchmakeParam returns a new MatchmakeParam
func NewMatchmakeParam() *MatchmakeParam {
	return &MatchmakeParam{}
}
