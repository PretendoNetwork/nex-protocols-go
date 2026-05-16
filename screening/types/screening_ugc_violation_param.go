// Package types implements all the types used by the Screening protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/screening/constants"
)

// ScreeningUGCViolationParam is a type within the Screening protocol
type ScreeningUGCViolationParam struct {
	types.Structure
	Category         constants.ReportCategory
	Reason           types.String
	Context          types.List[ScreeningContextInfo]
	ScreenshotDataID types.UInt64
}

// WriteTo writes the ScreeningUGCViolationParam to the given writable
func (suvp ScreeningUGCViolationParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	suvp.Category.WriteTo(contentWritable)
	suvp.Reason.WriteTo(contentWritable)
	suvp.Context.WriteTo(contentWritable)
	suvp.ScreenshotDataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	suvp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ScreeningUGCViolationParam from the given readable
func (suvp *ScreeningUGCViolationParam) ExtractFrom(readable types.Readable) error {
	if err := suvp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningUGCViolationParam header. %s", err.Error())
	}

	if err := suvp.Category.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningUGCViolationParam.Category. %s", err.Error())
	}

	if err := suvp.Reason.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningUGCViolationParam.Reason. %s", err.Error())
	}

	if err := suvp.Context.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningUGCViolationParam.Context. %s", err.Error())
	}

	if err := suvp.ScreenshotDataID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningUGCViolationParam.ScreenshotDataID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ScreeningUGCViolationParam
func (suvp ScreeningUGCViolationParam) Copy() types.RVType {
	copied := NewScreeningUGCViolationParam()

	copied.StructureVersion = suvp.StructureVersion
	copied.Category = suvp.Category
	copied.Reason = suvp.Reason
	copied.Context = suvp.Context
	copied.ScreenshotDataID = suvp.ScreenshotDataID

	return copied
}

// Equals checks if the given ScreeningUGCViolationParam contains the same data as the current ScreeningUGCViolationParam
func (suvp ScreeningUGCViolationParam) Equals(o types.RVType) bool {
	if _, ok := o.(ScreeningUGCViolationParam); !ok {
		return false
	}

	other := o.(ScreeningUGCViolationParam)

	if suvp.StructureVersion != other.StructureVersion {
		return false
	}

	if suvp.Category != other.Category {
		return false
	}

	if suvp.Reason != other.Reason {
		return false
	}

	if !suvp.Context.Equals(other.Context) {
		return false
	}

	return suvp.ScreenshotDataID == other.ScreenshotDataID
}

// CopyRef copies the current value of the ScreeningUGCViolationParam
// and returns a pointer to the new copy
func (suvp ScreeningUGCViolationParam) CopyRef() types.RVTypePtr {
	copied := suvp
	return &copied
}

// Deref takes a pointer to the ScreeningUGCViolationParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (suvp *ScreeningUGCViolationParam) Deref() types.RVType {
	return *suvp
}

// String returns the string representation of the ScreeningUGCViolationParam
func (suvp ScreeningUGCViolationParam) String() string {
	return suvp.FormatToString(0)
}

// FormatToString pretty-prints the ScreeningUGCViolationParam using the provided indentation level
func (suvp ScreeningUGCViolationParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ScreeningUGCViolationParam{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, suvp.Category))
	b.WriteString(fmt.Sprintf("%sReason: %s,\n", indentationValues, suvp.Reason))
	b.WriteString(fmt.Sprintf("%sContext: %s,\n", indentationValues, suvp.Context))
	b.WriteString(fmt.Sprintf("%sScreenshotDataID: %s\n", indentationValues, suvp.ScreenshotDataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewScreeningUGCViolationParam returns a new ScreeningUGCViolationParam
func NewScreeningUGCViolationParam() ScreeningUGCViolationParam {
	return ScreeningUGCViolationParam{
		Category:         constants.ReportCategoryInvalid, // TODO - Find the real default
		Reason:           types.NewString(""),
		Context:          types.NewList[ScreeningContextInfo](),
		ScreenshotDataID: types.NewUInt64(0),
	}
}
