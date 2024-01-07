// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FindMatchmakeSessionByParticipantParam holds parameters for a matchmake session
type FindMatchmakeSessionByParticipantParam struct {
	types.Structure
	PrincipalIDList *types.List[*types.PID]
	ResultOptions   *types.PrimitiveU32
	BlockListParam  *MatchmakeBlockListParam
}

// ExtractFrom extracts the FindMatchmakeSessionByParticipantParam from the given readable
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = findMatchmakeSessionByParticipantParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FindMatchmakeSessionByParticipantParam header. %s", err.Error())
	}

	err = findMatchmakeSessionByParticipantParam.PrincipalIDList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.PrincipalIDList. %s", err.Error())
	}

	err = findMatchmakeSessionByParticipantParam.ResultOptions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.ResultOptions. %s", err.Error())
	}

	err = findMatchmakeSessionByParticipantParam.BlockListParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.BlockListParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantParam
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) Copy() types.RVType {
	copied := NewFindMatchmakeSessionByParticipantParam()

	copied.StructureVersion = findMatchmakeSessionByParticipantParam.StructureVersion

	copied.PrincipalIDList = make(*types.List[*types.PID], len(findMatchmakeSessionByParticipantParam.PrincipalIDList))

	for i := 0; i < len(findMatchmakeSessionByParticipantParam.PrincipalIDList); i++ {
		copied.PrincipalIDList[i] = findMatchmakeSessionByParticipantParam.PrincipalIDList[i].Copy()
	}

	copied.ResultOptions = findMatchmakeSessionByParticipantParam.ResultOptions

	copied.BlockListParam = findMatchmakeSessionByParticipantParam.BlockListParam.Copy().(*MatchmakeBlockListParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (findMatchmakeSessionByParticipantParam *FindMatchmakeSessionByParticipantParam) Equals(o types.RVType) bool {
	if _, ok := o.(*FindMatchmakeSessionByParticipantParam); !ok {
		return false
	}

	other := o.(*FindMatchmakeSessionByParticipantParam)

	if findMatchmakeSessionByParticipantParam.StructureVersion != other.StructureVersion {
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

	if !findMatchmakeSessionByParticipantParam.ResultOptions.Equals(other.ResultOptions) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, findMatchmakeSessionByParticipantParam.StructureVersion))
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
