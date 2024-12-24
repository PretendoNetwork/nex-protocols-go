// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FindMatchmakeSessionByParticipantParam is a type within the Matchmaking protocol
type FindMatchmakeSessionByParticipantParam struct {
	types.Structure
	PrincipalIDList types.List[types.PID]
	ResultOptions   types.UInt32
	BlockListParam  MatchmakeBlockListParam
}

// WriteTo writes the FindMatchmakeSessionByParticipantParam to the given writable
func (fmsbpp FindMatchmakeSessionByParticipantParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	fmsbpp.PrincipalIDList.WriteTo(contentWritable)
	fmsbpp.ResultOptions.WriteTo(contentWritable)
	fmsbpp.BlockListParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	fmsbpp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the FindMatchmakeSessionByParticipantParam from the given readable
func (fmsbpp *FindMatchmakeSessionByParticipantParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = fmsbpp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam header. %s", err.Error())
	}

	err = fmsbpp.PrincipalIDList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.PrincipalIDList. %s", err.Error())
	}

	err = fmsbpp.ResultOptions.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.ResultOptions. %s", err.Error())
	}

	err = fmsbpp.BlockListParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FindMatchmakeSessionByParticipantParam.BlockListParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of FindMatchmakeSessionByParticipantParam
func (fmsbpp FindMatchmakeSessionByParticipantParam) Copy() types.RVType {
	copied := NewFindMatchmakeSessionByParticipantParam()

	copied.StructureVersion = fmsbpp.StructureVersion
	copied.PrincipalIDList = fmsbpp.PrincipalIDList.Copy().(types.List[types.PID])
	copied.ResultOptions = fmsbpp.ResultOptions.Copy().(types.UInt32)
	copied.BlockListParam = fmsbpp.BlockListParam.Copy().(MatchmakeBlockListParam)

	return copied
}

// Equals checks if the given FindMatchmakeSessionByParticipantParam contains the same data as the current FindMatchmakeSessionByParticipantParam
func (fmsbpp FindMatchmakeSessionByParticipantParam) Equals(o types.RVType) bool {
	if _, ok := o.(FindMatchmakeSessionByParticipantParam); !ok {
		return false
	}

	other := o.(FindMatchmakeSessionByParticipantParam)

	if fmsbpp.StructureVersion != other.StructureVersion {
		return false
	}

	if !fmsbpp.PrincipalIDList.Equals(other.PrincipalIDList) {
		return false
	}

	if !fmsbpp.ResultOptions.Equals(other.ResultOptions) {
		return false
	}

	return fmsbpp.BlockListParam.Equals(other.BlockListParam)
}

// CopyRef copies the current value of the FindMatchmakeSessionByParticipantParam
// and returns a pointer to the new copy
func (fmsbpp FindMatchmakeSessionByParticipantParam) CopyRef() types.RVTypePtr {
	copied := fmsbpp.Copy().(FindMatchmakeSessionByParticipantParam)
	return &copied
}

// Deref takes a pointer to the FindMatchmakeSessionByParticipantParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (fmsbpp *FindMatchmakeSessionByParticipantParam) Deref() types.RVType {
	return *fmsbpp
}

// String returns the string representation of the FindMatchmakeSessionByParticipantParam
func (fmsbpp FindMatchmakeSessionByParticipantParam) String() string {
	return fmsbpp.FormatToString(0)
}

// FormatToString pretty-prints the FindMatchmakeSessionByParticipantParam using the provided indentation level
func (fmsbpp FindMatchmakeSessionByParticipantParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FindMatchmakeSessionByParticipantParam{\n")
	b.WriteString(fmt.Sprintf("%sPrincipalIDList: %s,\n", indentationValues, fmsbpp.PrincipalIDList))
	b.WriteString(fmt.Sprintf("%sResultOptions: %s,\n", indentationValues, fmsbpp.ResultOptions))
	b.WriteString(fmt.Sprintf("%sBlockListParam: %s,\n", indentationValues, fmsbpp.BlockListParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFindMatchmakeSessionByParticipantParam returns a new FindMatchmakeSessionByParticipantParam
func NewFindMatchmakeSessionByParticipantParam() FindMatchmakeSessionByParticipantParam {
	return FindMatchmakeSessionByParticipantParam{
		PrincipalIDList: types.NewList[types.PID](),
		ResultOptions:   types.NewUInt32(0),
		BlockListParam:  NewMatchmakeBlockListParam(),
	}

}
