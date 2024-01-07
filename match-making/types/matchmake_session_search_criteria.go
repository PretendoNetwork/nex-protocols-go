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

// MatchmakeSessionSearchCriteria holds information about a matchmaking search
type MatchmakeSessionSearchCriteria struct {
	types.Structure
	Attribs                  *types.List[*types.String]
	GameMode                 string
	MinParticipants          string
	MaxParticipants          string
	MatchmakeSystemType      string
	VacantOnly               *types.PrimitiveBool
	ExcludeLocked            *types.PrimitiveBool
	ExcludeNonHostPID        *types.PrimitiveBool
	SelectionMethod          *types.PrimitiveU32           // NEX v3.0.0+
	VacantParticipants       *types.PrimitiveU16           // NEX v3.4.0+
	MatchmakeParam           *MatchmakeParam  // NEX v3.6.0+
	ExcludeUserPasswordSet   *types.PrimitiveBool             // NEX v3.7.0+
	ExcludeSystemPasswordSet *types.PrimitiveBool             // NEX v3.7.0+
	ReferGID                 *types.PrimitiveU32           // NEX v3.8.0+
	CodeWord                 string           // NEX v4.0.0+
	ResultRange              *types.ResultRange // NEX v4.0.0+
}

// ExtractFrom extracts the Gathering from the given readable
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) ExtractFrom(readable types.Readable) error {
	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	var err error

	err = matchmakeSessionSearchCriteria.Attribs.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.Attribs. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.GameMode. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.MinParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MinParticipants. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.MaxParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MaxParticipants. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeSystemType. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.VacantOnly.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantOnly. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.ExcludeLocked.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeLocked. %s", err.Error())
	}

	err = matchmakeSessionSearchCriteria.ExcludeNonHostPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeNonHostPID. %s", err.Error())
	}

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
	err = 	matchmakeSessionSearchCriteria.SelectionMethod.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.SelectionMethod. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
	err = 	matchmakeSessionSearchCriteria.VacantParticipants.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.VacantParticipants. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		matchmakeParam, err := nex.StreamReadStructure(stream, NewMatchmakeParam())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.MatchmakeParam. %s", err.Error())
		}

		matchmakeSessionSearchCriteria.MatchmakeParam = matchmakeParam
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
	err = 	matchmakeSessionSearchCriteria.ExcludeUserPasswordSet.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeUserPasswordSet. %s", err.Error())
		}

	err = 	matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ExcludeSystemPasswordSet. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("3.8.0") {
	err = 	matchmakeSessionSearchCriteria.ReferGID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ReferGID. %s", err.Error())
		}
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
	err = 	matchmakeSessionSearchCriteria.CodeWord.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.CodeWord. %s", err.Error())
		}

		resultRange, err := nex.StreamReadStructure(stream, types.NewResultRange())
		if err != nil {
			return fmt.Errorf("Failed to extract MatchmakeSessionSearchCriteria.ResultRange. %s", err.Error())
		}

		matchmakeSessionSearchCriteria.ResultRange = resultRange
	}

	return nil
}

// WriteTo writes the Gathering to the given writable
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	matchmakingVersion := stream.Server.MatchMakingProtocolVersion()

	matchmakeSessionSearchCriteria.Attribs.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.GameMode.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.MinParticipants.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.MaxParticipants.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.MatchmakeSystemType.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.VacantOnly.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.ExcludeLocked.WriteTo(contentWritable)
	matchmakeSessionSearchCriteria.ExcludeNonHostPID.WriteTo(contentWritable)

	if matchmakingVersion.GreaterOrEqual("3.0.0") {
		matchmakeSessionSearchCriteria.SelectionMethod.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		matchmakeSessionSearchCriteria.VacantParticipants.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.6.0") {
		matchmakeSessionSearchCriteria.MatchmakeParam.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.7.0") {
		matchmakeSessionSearchCriteria.ExcludeUserPasswordSet.WriteTo(contentWritable)
		matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("3.8.0") {
		matchmakeSessionSearchCriteria.ReferGID.WriteTo(contentWritable)
	}

	if matchmakingVersion.GreaterOrEqual("4.0.0") {
		matchmakeSessionSearchCriteria.CodeWord.WriteTo(contentWritable)
		matchmakeSessionSearchCriteria.ResultRange.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Gathering
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Copy() types.RVType {
	copied := NewMatchmakeSessionSearchCriteria()

	copied.StructureVersion = matchmakeSessionSearchCriteria.StructureVersion

	copied.Attribs = make(*types.List[*types.String], len(matchmakeSessionSearchCriteria.Attribs))

	copy(copied.Attribs, matchmakeSessionSearchCriteria.Attribs)

	copied.GameMode = matchmakeSessionSearchCriteria.GameMode
	copied.MinParticipants = matchmakeSessionSearchCriteria.MinParticipants
	copied.MaxParticipants = matchmakeSessionSearchCriteria.MaxParticipants
	copied.MatchmakeSystemType = matchmakeSessionSearchCriteria.MatchmakeSystemType
	copied.VacantOnly = matchmakeSessionSearchCriteria.VacantOnly
	copied.ExcludeLocked = matchmakeSessionSearchCriteria.ExcludeLocked
	copied.ExcludeNonHostPID = matchmakeSessionSearchCriteria.ExcludeNonHostPID
	copied.SelectionMethod = matchmakeSessionSearchCriteria.SelectionMethod
	copied.VacantParticipants = matchmakeSessionSearchCriteria.VacantParticipants

	copied.MatchmakeParam = matchmakeSessionSearchCriteria.MatchmakeParam.Copy().(*MatchmakeParam)

	copied.ExcludeUserPasswordSet = matchmakeSessionSearchCriteria.ExcludeUserPasswordSet
	copied.ExcludeSystemPasswordSet = matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet
	copied.ReferGID = matchmakeSessionSearchCriteria.ReferGID
	copied.CodeWord = matchmakeSessionSearchCriteria.CodeWord

	copied.ResultRange = matchmakeSessionSearchCriteria.ResultRange.Copy().(*types.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) Equals(o types.RVType) bool {
	if _, ok := o.(*MatchmakeSessionSearchCriteria); !ok {
		return false
	}

	other := o.(*MatchmakeSessionSearchCriteria)

	if matchmakeSessionSearchCriteria.StructureVersion != other.StructureVersion {
		return false
	}

	if len(matchmakeSessionSearchCriteria.Attribs) != len(other.Attribs) {
		return false
	}

	for i := 0; i < len(matchmakeSessionSearchCriteria.Attribs); i++ {
		if matchmakeSessionSearchCriteria.Attribs[i] != other.Attribs[i] {
			return false
		}
	}

	if !matchmakeSessionSearchCriteria.GameMode.Equals(other.GameMode) {
		return false
	}

	if !matchmakeSessionSearchCriteria.MinParticipants.Equals(other.MinParticipants) {
		return false
	}

	if !matchmakeSessionSearchCriteria.MaxParticipants.Equals(other.MaxParticipants) {
		return false
	}

	if !matchmakeSessionSearchCriteria.MatchmakeSystemType.Equals(other.MatchmakeSystemType) {
		return false
	}

	if !matchmakeSessionSearchCriteria.VacantOnly.Equals(other.VacantOnly) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ExcludeLocked.Equals(other.ExcludeLocked) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ExcludeNonHostPID.Equals(other.ExcludeNonHostPID) {
		return false
	}

	if !matchmakeSessionSearchCriteria.SelectionMethod.Equals(other.SelectionMethod) {
		return false
	}

	if !matchmakeSessionSearchCriteria.VacantParticipants.Equals(other.VacantParticipants) {
		return false
	}

	if !matchmakeSessionSearchCriteria.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ExcludeUserPasswordSet.Equals(other.ExcludeUserPasswordSet) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet.Equals(other.ExcludeSystemPasswordSet) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ReferGID.Equals(other.ReferGID) {
		return false
	}

	if !matchmakeSessionSearchCriteria.CodeWord.Equals(other.CodeWord) {
		return false
	}

	if !matchmakeSessionSearchCriteria.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) String() string {
	return matchmakeSessionSearchCriteria.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeSessionSearchCriteria{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, matchmakeSessionSearchCriteria.StructureVersion))
	b.WriteString(fmt.Sprintf("%sAttribs: %v,\n", indentationValues, matchmakeSessionSearchCriteria.Attribs))
	b.WriteString(fmt.Sprintf("%sGameMode: %q,\n", indentationValues, matchmakeSessionSearchCriteria.GameMode))
	b.WriteString(fmt.Sprintf("%sMinParticipants: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MinParticipants))
	b.WriteString(fmt.Sprintf("%sMaxParticipants: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MaxParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %q,\n", indentationValues, matchmakeSessionSearchCriteria.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sVacantOnly: %t,\n", indentationValues, matchmakeSessionSearchCriteria.VacantOnly))
	b.WriteString(fmt.Sprintf("%sExcludeLocked: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeLocked))
	b.WriteString(fmt.Sprintf("%sExcludeNonHostPID: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeNonHostPID))
	b.WriteString(fmt.Sprintf("%sSelectionMethod: %d,\n", indentationValues, matchmakeSessionSearchCriteria.SelectionMethod))
	b.WriteString(fmt.Sprintf("%sVacantParticipants: %d,\n", indentationValues, matchmakeSessionSearchCriteria.VacantParticipants))

	if matchmakeSessionSearchCriteria.MatchmakeParam != nil {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, matchmakeSessionSearchCriteria.MatchmakeParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMatchmakeParam: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sExcludeUserPasswordSet: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeUserPasswordSet))
	b.WriteString(fmt.Sprintf("%sExcludeSystemPasswordSet: %t,\n", indentationValues, matchmakeSessionSearchCriteria.ExcludeSystemPasswordSet))
	b.WriteString(fmt.Sprintf("%sReferGID: %d,\n", indentationValues, matchmakeSessionSearchCriteria.ReferGID))
	b.WriteString(fmt.Sprintf("%sCodeWord: %q\n", indentationValues, matchmakeSessionSearchCriteria.CodeWord))

	if matchmakeSessionSearchCriteria.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, matchmakeSessionSearchCriteria.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeSessionSearchCriteria returns a new MatchmakeSessionSearchCriteria
func NewMatchmakeSessionSearchCriteria() *MatchmakeSessionSearchCriteria {
	return &MatchmakeSessionSearchCriteria{}
}
