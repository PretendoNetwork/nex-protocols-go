// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MatchmakeParam holds parameters for a matchmake session
type MatchmakeParam struct {
	nex.Structure
	Parameters map[string]*nex.Variant
}

// ExtractFromStream extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	matchmakeParam.Parameters, err = nex.StreamReadMap(stream, stream.ReadString, stream.ReadVariant)

	return err
}

// Bytes extracts a MatchmakeParam structure from a stream
func (matchmakeParam *MatchmakeParam) Bytes(stream *nex.StreamOut) []byte {
	nex.StreamWriteMap(stream, matchmakeParam.Parameters)

	return stream.Bytes()
}

// Copy returns a new copied instance of MatchmakeParam
func (matchmakeParam *MatchmakeParam) Copy() nex.StructureInterface {
	copied := NewMatchmakeParam()

	copied.SetStructureVersion(matchmakeParam.StructureVersion())

	copied.Parameters = make(map[string]*nex.Variant, len(matchmakeParam.Parameters))

	for key, value := range matchmakeParam.Parameters {
		copied.Parameters[key] = value.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeParam *MatchmakeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MatchmakeParam)

	if matchmakeParam.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (matchmakeParam *MatchmakeParam) String() string {
	return matchmakeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeParam *MatchmakeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationMapValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, matchmakeParam.StructureVersion()))

	if len(matchmakeParam.Parameters) == 0 {
		b.WriteString(fmt.Sprintf("%sParameters: {}\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sParameters: {\n", indentationValues))

		for k, v := range matchmakeParam.Parameters {
			// TODO - Special handle the the last item to not add the comma on last item
			b.WriteString(fmt.Sprintf("%s%q: %s,\n", indentationMapValues, k, v.FormatToString(indentationLevel+2)))
		}

		b.WriteString(fmt.Sprintf("%s}\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeParam returns a new MatchmakeParam
func NewMatchmakeParam() *MatchmakeParam {
	return &MatchmakeParam{}
}
