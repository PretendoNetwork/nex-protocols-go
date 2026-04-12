// Package types implements all the types used by the TicketGranting protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ValidateAndRequestTicketResult is a type within the TicketGranting protocol
type ValidateAndRequestTicketResult struct {
	types.Structure
	SourcePID      types.PID
	BufResponse    types.Buffer
	ServiceNodeURL types.StationURL
	CurrentUTCTime types.DateTime
	ReturnMsg      types.String
	SourceKey      types.String
	PlatformPID    types.PID // * Only present on games with crossplay between Switch and 3DS/Wii U
}

// WriteTo writes the ValidateAndRequestTicketResult to the given writable
func (vartr ValidateAndRequestTicketResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	vartr.SourcePID.WriteTo(contentWritable)
	vartr.BufResponse.WriteTo(contentWritable)
	vartr.ServiceNodeURL.WriteTo(contentWritable)
	vartr.CurrentUTCTime.WriteTo(contentWritable)
	vartr.ReturnMsg.WriteTo(contentWritable)
	vartr.SourceKey.WriteTo(contentWritable)

	// * This PID, so if it's 0 we can safely assume it isn't
	// * being used. This field is only present on games with
	// * crossplay between Switch and 3DS/Wii U
	if vartr.PlatformPID != 0 {
		vartr.PlatformPID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	vartr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ValidateAndRequestTicketResult from the given readable
func (vartr *ValidateAndRequestTicketResult) ExtractFrom(readable types.Readable) error {
	if err := vartr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult header. %s", err.Error())
	}

	if err := vartr.SourcePID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.PlatformType. %s", err.Error())
	}

	if err := vartr.BufResponse.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.Username. %s", err.Error())
	}

	if err := vartr.ServiceNodeURL.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.ExtraData. %s", err.Error())
	}

	if err := vartr.CurrentUTCTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.IgnoreAPIVersionCheck. %s", err.Error())
	}

	if err := vartr.ReturnMsg.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.APIVersionGeneral. %s", err.Error())
	}

	if err := vartr.SourceKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.APIVersionCustom. %s", err.Error())
	}

	// * This is a hack. This field is not based on a NEX
	// * version difference or a structure version update.
	// * It is only present on games with crossplay between
	// * Switch and 3DS/Wii U. Since we don't know whether
	// * or not a game has crossplay at this stage, this is
	// * the best we can do
	if readable.Remaining() != 0 {
		if err := vartr.PlatformPID.ExtractFrom(readable); err != nil {
			return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.PlatformTypeForPlatformPID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ValidateAndRequestTicketResult
func (vartr ValidateAndRequestTicketResult) Copy() types.RVType {
	copied := NewValidateAndRequestTicketResult()

	copied.StructureVersion = vartr.StructureVersion
	copied.SourcePID = vartr.SourcePID
	copied.BufResponse = vartr.BufResponse.Copy().(types.Buffer)
	copied.ServiceNodeURL = vartr.ServiceNodeURL.Copy().(types.StationURL)
	copied.CurrentUTCTime = vartr.CurrentUTCTime
	copied.ReturnMsg = vartr.ReturnMsg
	copied.SourceKey = vartr.SourceKey
	copied.PlatformPID = vartr.PlatformPID

	return copied
}

// Equals checks if the given ValidateAndRequestTicketResult contains the same data as the current ValidateAndRequestTicketResult
func (vartr ValidateAndRequestTicketResult) Equals(o types.RVType) bool {
	if _, ok := o.(ValidateAndRequestTicketResult); !ok {
		return false
	}

	other := o.(ValidateAndRequestTicketResult)

	if vartr.StructureVersion != other.StructureVersion {
		return false
	}

	if vartr.SourcePID != other.SourcePID {
		return false
	}

	if !vartr.BufResponse.Equals(other.BufResponse) {
		return false
	}

	if !vartr.ServiceNodeURL.Equals(other.ServiceNodeURL) {
		return false
	}

	if vartr.CurrentUTCTime != other.CurrentUTCTime {
		return false
	}

	if vartr.ReturnMsg != other.ReturnMsg {
		return false
	}

	if vartr.SourceKey != other.SourceKey {
		return false
	}

	return vartr.PlatformPID == other.PlatformPID
}

// CopyRef copies the current value of the ValidateAndRequestTicketResult
// and returns a pointer to the new copy
func (vartr ValidateAndRequestTicketResult) CopyRef() types.RVTypePtr {
	copied := vartr.Copy().(ValidateAndRequestTicketResult)
	return &copied
}

// Deref takes a pointer to the ValidateAndRequestTicketResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (vartr *ValidateAndRequestTicketResult) Deref() types.RVType {
	return *vartr
}

// String returns the string representation of the ValidateAndRequestTicketResult
func (vartr ValidateAndRequestTicketResult) String() string {
	return vartr.FormatToString(0)
}

// FormatToString pretty-prints the ValidateAndRequestTicketResult using the provided indentation level
func (vartr ValidateAndRequestTicketResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ValidateAndRequestTicketResult{\n")
	b.WriteString(fmt.Sprintf("%sSourcePID: %s,\n", indentationValues, vartr.SourcePID))
	b.WriteString(fmt.Sprintf("%sBufResponse: %s,\n", indentationValues, vartr.BufResponse))
	b.WriteString(fmt.Sprintf("%sServiceNodeURL: %s,\n", indentationValues, vartr.ServiceNodeURL.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCurrentUTCTime: %s,\n", indentationValues, vartr.CurrentUTCTime))
	b.WriteString(fmt.Sprintf("%sReturnMsg: %s,\n", indentationValues, vartr.ReturnMsg))

	// * This is a hack. This field is not based on a NEX
	// * version difference or a structure version update.
	// * It is only present on games with crossplay between
	// * Switch and 3DS/Wii U. Since we don't know whether
	// * or not a game has crossplay at this stage, this is
	// * the best we can do
	if vartr.PlatformPID != 0 {
		b.WriteString(fmt.Sprintf("%sSourceKey: %s,\n", indentationValues, vartr.SourceKey))
		b.WriteString(fmt.Sprintf("%sPlatformPID: %s\n", indentationValues, vartr.PlatformPID))
	} else {
		b.WriteString(fmt.Sprintf("%sSourceKey: %s\n", indentationValues, vartr.SourceKey))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewValidateAndRequestTicketResult returns a new ValidateAndRequestTicketResult
func NewValidateAndRequestTicketResult() ValidateAndRequestTicketResult {
	return ValidateAndRequestTicketResult{
		SourcePID:      types.NewPID(0),
		BufResponse:    types.NewBuffer(nil),
		ServiceNodeURL: types.NewStationURL(""),
		CurrentUTCTime: types.NewDateTime(0),
		ReturnMsg:      types.NewString(""),
		SourceKey:      types.NewString(""),
		PlatformPID:    types.NewPID(0),
	}

}
