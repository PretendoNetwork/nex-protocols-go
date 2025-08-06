// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/match-making/constants"
)

// UpdateMatchmakeSessionParam is a type within the Matchmaking protocol
type UpdateMatchmakeSessionParam struct {
	types.Structure
	GID                 types.UInt32
	ModificationFlag    constants.MatchmakeSessionModificationFlag
	Attributes          types.List[types.UInt32]
	OpenParticipation   types.Bool
	ApplicationBuffer   types.Buffer
	ProgressScore       types.UInt8
	MatchmakeParam      MatchmakeParam
	StartedTime         types.DateTime
	UserPassword        types.String
	GameMode            types.UInt32
	Description         types.String
	MinParticipants     types.UInt16
	MaxParticipants     types.UInt16
	MatchmakeSystemType constants.MatchmakeSystemType
	ParticipationPolicy constants.ParticipationPolicy
	PolicyArgument      constants.PolicyArgument
	Codeword            types.String
}

// WriteTo writes the UpdateMatchmakeSessionParam to the given writable
func (umsp UpdateMatchmakeSessionParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	umsp.GID.WriteTo(contentWritable)
	umsp.ModificationFlag.WriteTo(contentWritable)
	umsp.Attributes.WriteTo(contentWritable)
	umsp.OpenParticipation.WriteTo(contentWritable)
	umsp.ApplicationBuffer.WriteTo(contentWritable)
	umsp.ProgressScore.WriteTo(contentWritable)
	umsp.MatchmakeParam.WriteTo(contentWritable)
	umsp.StartedTime.WriteTo(contentWritable)
	umsp.UserPassword.WriteTo(contentWritable)
	umsp.GameMode.WriteTo(contentWritable)
	umsp.Description.WriteTo(contentWritable)
	umsp.MinParticipants.WriteTo(contentWritable)
	umsp.MaxParticipants.WriteTo(contentWritable)
	umsp.MatchmakeSystemType.WriteTo(contentWritable)
	umsp.ParticipationPolicy.WriteTo(contentWritable)
	umsp.PolicyArgument.WriteTo(contentWritable)
	umsp.Codeword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	umsp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UpdateMatchmakeSessionParam from the given readable
func (umsp *UpdateMatchmakeSessionParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = umsp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam header. %s", err.Error())
	}

	err = umsp.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GID. %s", err.Error())
	}

	err = umsp.ModificationFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ModificationFlag. %s", err.Error())
	}

	err = umsp.Attributes.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Attributes. %s", err.Error())
	}

	err = umsp.OpenParticipation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.OpenParticipation. %s", err.Error())
	}

	err = umsp.ApplicationBuffer.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ApplicationBuffer. %s", err.Error())
	}

	err = umsp.ProgressScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ProgressScore. %s", err.Error())
	}

	err = umsp.MatchmakeParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeParam. %s", err.Error())
	}

	err = umsp.StartedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.StartedTime. %s", err.Error())
	}

	err = umsp.UserPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.UserPassword. %s", err.Error())
	}

	err = umsp.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.GameMode. %s", err.Error())
	}

	err = umsp.Description.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Description. %s", err.Error())
	}

	err = umsp.MinParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MinParticipants. %s", err.Error())
	}

	err = umsp.MaxParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MaxParticipants. %s", err.Error())
	}

	err = umsp.MatchmakeSystemType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.MatchmakeSystemType. %s", err.Error())
	}

	err = umsp.ParticipationPolicy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.ParticipationPolicy. %s", err.Error())
	}

	err = umsp.PolicyArgument.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.PolicyArgument. %s", err.Error())
	}

	err = umsp.Codeword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract UpdateMatchmakeSessionParam.Codeword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UpdateMatchmakeSessionParam
func (umsp UpdateMatchmakeSessionParam) Copy() types.RVType {
	copied := NewUpdateMatchmakeSessionParam()

	copied.StructureVersion = umsp.StructureVersion
	copied.GID = umsp.GID.Copy().(types.UInt32)
	copied.ModificationFlag = umsp.ModificationFlag
	copied.Attributes = umsp.Attributes.Copy().(types.List[types.UInt32])
	copied.OpenParticipation = umsp.OpenParticipation.Copy().(types.Bool)
	copied.ApplicationBuffer = umsp.ApplicationBuffer.Copy().(types.Buffer)
	copied.ProgressScore = umsp.ProgressScore.Copy().(types.UInt8)
	copied.MatchmakeParam = umsp.MatchmakeParam.Copy().(MatchmakeParam)
	copied.StartedTime = umsp.StartedTime.Copy().(types.DateTime)
	copied.UserPassword = umsp.UserPassword.Copy().(types.String)
	copied.GameMode = umsp.GameMode.Copy().(types.UInt32)
	copied.Description = umsp.Description.Copy().(types.String)
	copied.MinParticipants = umsp.MinParticipants.Copy().(types.UInt16)
	copied.MaxParticipants = umsp.MaxParticipants.Copy().(types.UInt16)
	copied.MatchmakeSystemType = umsp.MatchmakeSystemType
	copied.ParticipationPolicy = umsp.ParticipationPolicy
	copied.PolicyArgument = umsp.PolicyArgument
	copied.Codeword = umsp.Codeword.Copy().(types.String)

	return copied
}

// Equals checks if the given UpdateMatchmakeSessionParam contains the same data as the current UpdateMatchmakeSessionParam
func (umsp UpdateMatchmakeSessionParam) Equals(o types.RVType) bool {
	if _, ok := o.(UpdateMatchmakeSessionParam); !ok {
		return false
	}

	other := o.(UpdateMatchmakeSessionParam)

	if umsp.StructureVersion != other.StructureVersion {
		return false
	}

	if !umsp.GID.Equals(other.GID) {
		return false
	}

	if umsp.ModificationFlag != other.ModificationFlag {
		return false
	}

	if !umsp.Attributes.Equals(other.Attributes) {
		return false
	}

	if !umsp.OpenParticipation.Equals(other.OpenParticipation) {
		return false
	}

	if !umsp.ApplicationBuffer.Equals(other.ApplicationBuffer) {
		return false
	}

	if !umsp.ProgressScore.Equals(other.ProgressScore) {
		return false
	}

	if !umsp.MatchmakeParam.Equals(other.MatchmakeParam) {
		return false
	}

	if !umsp.StartedTime.Equals(other.StartedTime) {
		return false
	}

	if !umsp.UserPassword.Equals(other.UserPassword) {
		return false
	}

	if !umsp.GameMode.Equals(other.GameMode) {
		return false
	}

	if !umsp.Description.Equals(other.Description) {
		return false
	}

	if !umsp.MinParticipants.Equals(other.MinParticipants) {
		return false
	}

	if !umsp.MaxParticipants.Equals(other.MaxParticipants) {
		return false
	}

	if umsp.MatchmakeSystemType != other.MatchmakeSystemType {
		return false
	}

	if umsp.ParticipationPolicy != other.ParticipationPolicy {
		return false
	}

	if umsp.PolicyArgument != other.PolicyArgument {
		return false
	}

	return umsp.Codeword.Equals(other.Codeword)
}

// CopyRef copies the current value of the UpdateMatchmakeSessionParam
// and returns a pointer to the new copy
func (umsp UpdateMatchmakeSessionParam) CopyRef() types.RVTypePtr {
	copied := umsp.Copy().(UpdateMatchmakeSessionParam)
	return &copied
}

// Deref takes a pointer to the UpdateMatchmakeSessionParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (umsp *UpdateMatchmakeSessionParam) Deref() types.RVType {
	return *umsp
}

// String returns the string representation of the UpdateMatchmakeSessionParam
func (umsp UpdateMatchmakeSessionParam) String() string {
	return umsp.FormatToString(0)
}

// FormatToString pretty-prints the UpdateMatchmakeSessionParam using the provided indentation level
func (umsp UpdateMatchmakeSessionParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UpdateMatchmakeSessionParam{\n")
	b.WriteString(fmt.Sprintf("%sGID: %s,\n", indentationValues, umsp.GID))
	b.WriteString(fmt.Sprintf("%sModificationFlag: %s,\n", indentationValues, umsp.ModificationFlag))
	b.WriteString(fmt.Sprintf("%sAttributes: %s,\n", indentationValues, umsp.Attributes))
	b.WriteString(fmt.Sprintf("%sOpenParticipation: %s,\n", indentationValues, umsp.OpenParticipation))
	b.WriteString(fmt.Sprintf("%sApplicationBuffer: %s,\n", indentationValues, umsp.ApplicationBuffer))
	b.WriteString(fmt.Sprintf("%sProgressScore: %s,\n", indentationValues, umsp.ProgressScore))
	b.WriteString(fmt.Sprintf("%sMatchmakeParam: %s,\n", indentationValues, umsp.MatchmakeParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStartedTime: %s,\n", indentationValues, umsp.StartedTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUserPassword: %s,\n", indentationValues, umsp.UserPassword))
	b.WriteString(fmt.Sprintf("%sGameMode: %s,\n", indentationValues, umsp.GameMode))
	b.WriteString(fmt.Sprintf("%sDescription: %s,\n", indentationValues, umsp.Description))
	b.WriteString(fmt.Sprintf("%sMinParticipants: %s,\n", indentationValues, umsp.MinParticipants))
	b.WriteString(fmt.Sprintf("%sMaxParticipants: %s,\n", indentationValues, umsp.MaxParticipants))
	b.WriteString(fmt.Sprintf("%sMatchmakeSystemType: %s,\n", indentationValues, umsp.MatchmakeSystemType))
	b.WriteString(fmt.Sprintf("%sParticipationPolicy: %s,\n", indentationValues, umsp.ParticipationPolicy))
	b.WriteString(fmt.Sprintf("%sPolicyArgument: %s,\n", indentationValues, umsp.PolicyArgument))
	b.WriteString(fmt.Sprintf("%sCodeword: %s,\n", indentationValues, umsp.Codeword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUpdateMatchmakeSessionParam returns a new UpdateMatchmakeSessionParam
func NewUpdateMatchmakeSessionParam() UpdateMatchmakeSessionParam {
	return UpdateMatchmakeSessionParam{
		GID:                 types.NewUInt32(0),
		ModificationFlag:    constants.MatchmakeSessionModificationFlagNone,
		Attributes:          types.NewList[types.UInt32](),
		OpenParticipation:   types.NewBool(false),
		ApplicationBuffer:   types.NewBuffer(nil),
		ProgressScore:       types.NewUInt8(0),
		MatchmakeParam:      NewMatchmakeParam(),
		StartedTime:         types.NewDateTime(0),
		UserPassword:        types.NewString(""),
		GameMode:            types.NewUInt32(0),
		Description:         types.NewString(""),
		MinParticipants:     types.NewUInt16(0),
		MaxParticipants:     types.NewUInt16(0),
		MatchmakeSystemType: constants.MatchmakeSystemTypeInvalid,
		ParticipationPolicy: constants.ParticipationPolicy(0), // TODO - Unsure what the real values are for these
		PolicyArgument:      constants.PolicyArgument(0),      // TODO - Unsure what the real values are for these
		Codeword:            types.NewString(""),
	}

}
