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

// FindMatchmakeSessionByParticipantParam holds parameters for a matchmake session
type FindMatchmakeSessionByParticipantParam struct {
	nex.Structure
	PrincipalIDList []*nex.PID
	ResultOptions   uint32
	BlockListParam  *MatchmakeBlockListParam
}

// ExtractFromStream extracts a FindMatchmakeSessionByParticipantParam structure from a stream
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	findMatchmakeSessionByParticipantParam.PrincipalIDList, err = stream.ReadListPID()
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.PrincipalIDList. %s", err.Error())
	}

	findMatchmakeSessionByParticipantParam.ResultOptions, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.ResultOptions. %s", err.Error())
	}

	blockListParam, err := stream.ReadStructure(NewMatchmakeBlockListParam())
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.BlockListParam. %s", err.Error())
	}

	findMatchmakeSessionByParticipantParam.BlockListParam = blockListParam.(*MatchmakeBlockListParam)

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantParam
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) Copy() nex.StructureInterface {
	copied := NewFindMatchmakeSessionByParticipantParam()

	copied.SetStructureVersion(findMatchmakeSessionByParticipantParam.StructureVersion())

	copied.PrincipalIDList = make([]*nex.PID, len(findMatchmakeSessionByParticipantParam.PrincipalIDList))

	for i := 0; i < len(findMatchmakeSessionByParticipantParam.PrincipalIDList); i++ {
		copied.PrincipalIDList[i] = findMatchmakeSessionByParticipantParam.PrincipalIDList[i].Copy()
	}

	copied.ResultOptions = findMatchmakeSessionByParticipantParam.ResultOptions

	if findMatchmakeSessionByParticipantParam.BlockListParam != nil {
		copied.BlockListParam = findMatchmakeSessionByParticipantParam.BlockListParam.Copy().(*MatchmakeBlockListParam)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FindMatchmakeSessionByParticipantParam)

	if findMatchmakeSessionByParticipantParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if len(findMatchmakeSessionByParticipantParam.PrincipalIDList) != len(other.PrincipalIDList) {
		return false
	}

	for i := 0; i < len(findMatchmakeSessionByParticipantParam.PrincipalIDList); i++ {
		if !findMatchmakeSessionByParticipantParam.PrincipalIDList[i].Equals(other.PrincipalIDList[i]) {
			return false
		}
	}

	if findMatchmakeSessionByParticipantParam.ResultOptions != other.ResultOptions {
		return false
	}

	if findMatchmakeSessionByParticipantParam.BlockListParam != nil && other.BlockListParam == nil {
		return false
	}

	if findMatchmakeSessionByParticipantParam.BlockListParam == nil && other.BlockListParam != nil {
		return false
	}

	if findMatchmakeSessionByParticipantParam.BlockListParam != nil && other.BlockListParam != nil {
		if findMatchmakeSessionByParticipantParam.BlockListParam.Equals(other.BlockListParam) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) String() string {
	return findMatchmakeSessionByParticipantParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FindMatchmakeSessionByParticipantParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, findMatchmakeSessionByParticipantParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPrincipalIDList: %v,\n", indentationValues, findMatchmakeSessionByParticipantParam.PrincipalIDList))
	b.WriteString(fmt.Sprintf("%sResultOptions: %d,\n", indentationValues, findMatchmakeSessionByParticipantParam.ResultOptions))

	if findMatchmakeSessionByParticipantParam.BlockListParam != nil {
		b.WriteString(fmt.Sprintf("%sBlockListParam: %s\n", indentationValues, findMatchmakeSessionByParticipantParam.BlockListParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sBlockListParam: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFindMatchmakeSessionByParticipantParam returns a new FindMatchmakeSessionByParticipantParam
func NewFindMatchmakeSessionByParticipantParam() *FindMatchmakeSessionByParticipantParam {
	return &FindMatchmakeSessionByParticipantParam{}
}
