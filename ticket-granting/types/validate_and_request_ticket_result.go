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
	enableCrossplay bool
	SourcePID       types.PID
	BufResponse     types.Buffer
	ServiceNodeURL  types.StationURL
	CurrentUTCTime  types.DateTime
	ReturnMsg       types.String
	SourceKey       types.String
	PlatformPID     types.PID
}

// WriteTo writes the ValidateAndRequestTicketResult to the given writable
func (rs ValidateAndRequestTicketResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rs.SourcePID.WriteTo(contentWritable)
	rs.BufResponse.WriteTo(contentWritable)
	rs.ServiceNodeURL.WriteTo(contentWritable)
	rs.CurrentUTCTime.WriteTo(contentWritable)
	rs.ReturnMsg.WriteTo(contentWritable)
	rs.SourceKey.WriteTo(contentWritable)

	if rs.enableCrossplay {
		rs.PlatformPID.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ValidateAndRequestTicketResult from the given readable
func (rs *ValidateAndRequestTicketResult) ExtractFrom(readable types.Readable) error {
	var err error

	err = rs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult header. %s", err.Error())
	}

	err = rs.SourcePID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.SourcePID. %s", err.Error())
	}

	err = rs.BufResponse.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.BufResponse. %s", err.Error())
	}

	err = rs.ServiceNodeURL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.ServiceNodeURL. %s", err.Error())
	}

	err = rs.CurrentUTCTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.CurrentUTCTime. %s", err.Error())
	}

	err = rs.ReturnMsg.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.ReturnMsg. %s", err.Error())
	}

	err = rs.SourceKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.SourceKey. %s", err.Error())
	}

	if rs.enableCrossplay {
		err = rs.PlatformPID.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ValidateAndRequestTicketResult.PlatformPID. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of ValidateAndRequestTicketResult
func (rs ValidateAndRequestTicketResult) Copy() types.RVType {
	copied := NewValidateAndRequestTicketResult(rs.enableCrossplay)

	copied.StructureVersion = rs.StructureVersion
	copied.SourcePID = rs.SourcePID.Copy().(types.PID)
	copied.BufResponse = rs.BufResponse.Copy().(types.Buffer)
	copied.ServiceNodeURL = rs.ServiceNodeURL.Copy().(types.StationURL)
	copied.CurrentUTCTime = rs.CurrentUTCTime.Copy().(types.DateTime)
	copied.ReturnMsg = rs.ReturnMsg.Copy().(types.String)
	copied.SourceKey = rs.SourceKey.Copy().(types.String)
	copied.PlatformPID = rs.PlatformPID.Copy().(types.PID)

	return copied
}

// Equals checks if the given ValidateAndRequestTicketResult contains the same data as the current ValidateAndRequestTicketResult
func (rs ValidateAndRequestTicketResult) Equals(o types.RVType) bool {
	if _, ok := o.(ValidateAndRequestTicketResult); !ok {
		return false
	}

	other := o.(ValidateAndRequestTicketResult)

	if rs.StructureVersion != other.StructureVersion {
		return false
	}

	if !rs.SourcePID.Equals(other.SourcePID) {
		return false
	}

	if !rs.BufResponse.Equals(other.BufResponse) {
		return false
	}

	if !rs.ServiceNodeURL.Equals(other.ServiceNodeURL) {
		return false
	}

	if !rs.CurrentUTCTime.Equals(other.CurrentUTCTime) {
		return false
	}

	if !rs.ReturnMsg.Equals(other.ReturnMsg) {
		return false
	}

	if !rs.SourceKey.Equals(other.SourceKey) {
		return false
	}

	return rs.PlatformPID.Equals(other.PlatformPID)
}

// CopyRef copies the current value of the ValidateAndRequestTicketResult
// and returns a pointer to the new copy
func (rs ValidateAndRequestTicketResult) CopyRef() types.RVTypePtr {
	copied := rs.Copy().(ValidateAndRequestTicketResult)
	return &copied
}

// Deref takes a pointer to the ValidateAndRequestTicketResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rs *ValidateAndRequestTicketResult) Deref() types.RVType {
	return *rs
}

// String returns the string representation of the ValidateAndRequestTicketResult
func (rs ValidateAndRequestTicketResult) String() string {
	return rs.FormatToString(0)
}

// FormatToString pretty-prints the ValidateAndRequestTicketResult using the provided indentation level
func (rs ValidateAndRequestTicketResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ValidateAndRequestTicketResult{\n")
	b.WriteString(fmt.Sprintf("%sSourcePID: %s,\n", indentationValues, rs.SourcePID))
	b.WriteString(fmt.Sprintf("%sBufResponse: %s,\n", indentationValues, rs.BufResponse))
	b.WriteString(fmt.Sprintf("%sServiceNodeURL: %s,\n", indentationValues, rs.ServiceNodeURL.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCurrentUTCTime: %s,\n", indentationValues, rs.CurrentUTCTime))
	b.WriteString(fmt.Sprintf("%sReturnMsg: %s,\n", indentationValues, rs.ReturnMsg))
	b.WriteString(fmt.Sprintf("%sSourceKey: %s,\n", indentationValues, rs.SourceKey))
	b.WriteString(fmt.Sprintf("%sPlatformPID: %s,\n", indentationValues, rs.PlatformPID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewValidateAndRequestTicketResult returns a new ValidateAndRequestTicketResult
func NewValidateAndRequestTicketResult(enableCrossplay bool) ValidateAndRequestTicketResult {
	return ValidateAndRequestTicketResult{
		enableCrossplay: enableCrossplay,
		SourcePID:       types.NewPID(0),
		BufResponse:     types.NewBuffer(nil),
		ServiceNodeURL:  types.NewStationURL(""),
		CurrentUTCTime:  types.NewDateTime(0),
		ReturnMsg:       types.NewString(""),
		SourceKey:       types.NewString(""),
		PlatformPID:     types.NewPID(0),
	}

}
