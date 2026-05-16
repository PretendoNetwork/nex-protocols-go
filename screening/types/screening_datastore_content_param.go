// Package types implements all the types used by the Screening protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ScreeningDataStoreContentParam is a type within the Screening protocol
type ScreeningDataStoreContentParam struct {
	types.Structure
	DataID        types.UInt64
	ContentDataID types.UInt64
	UGCType       types.String
	Language      types.String
	SearchKey     types.String
}

// WriteTo writes the ScreeningDataStoreContentParam to the given writable
func (sdscp ScreeningDataStoreContentParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sdscp.DataID.WriteTo(contentWritable)
	sdscp.ContentDataID.WriteTo(contentWritable)
	sdscp.UGCType.WriteTo(contentWritable)
	sdscp.Language.WriteTo(contentWritable)
	sdscp.SearchKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sdscp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ScreeningDataStoreContentParam from the given readable
func (sdscp *ScreeningDataStoreContentParam) ExtractFrom(readable types.Readable) error {
	if err := sdscp.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam header. %s", err.Error())
	}

	if err := sdscp.DataID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam.DataID. %s", err.Error())
	}

	if err := sdscp.ContentDataID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam.ContentDataID. %s", err.Error())
	}

	if err := sdscp.UGCType.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam.UGCType. %s", err.Error())
	}

	if err := sdscp.Language.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam.Language. %s", err.Error())
	}

	if err := sdscp.SearchKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract ScreeningDataStoreContentParam.SearchKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ScreeningDataStoreContentParam
func (sdscp ScreeningDataStoreContentParam) Copy() types.RVType {
	copied := NewScreeningDataStoreContentParam()

	copied.StructureVersion = sdscp.StructureVersion
	copied.DataID = sdscp.DataID
	copied.ContentDataID = sdscp.ContentDataID
	copied.UGCType = sdscp.UGCType
	copied.Language = sdscp.Language
	copied.SearchKey = sdscp.SearchKey

	return copied
}

// Equals checks if the given ScreeningDataStoreContentParam contains the same data as the current ScreeningDataStoreContentParam
func (sdscp ScreeningDataStoreContentParam) Equals(o types.RVType) bool {
	if _, ok := o.(ScreeningDataStoreContentParam); !ok {
		return false
	}

	other := o.(ScreeningDataStoreContentParam)

	if sdscp.StructureVersion != other.StructureVersion {
		return false
	}

	if sdscp.DataID != other.DataID {
		return false
	}

	if sdscp.ContentDataID != other.ContentDataID {
		return false
	}

	if sdscp.UGCType != other.UGCType {
		return false
	}

	if sdscp.Language != other.Language {
		return false
	}

	return sdscp.SearchKey == other.SearchKey
}

// CopyRef copies the current value of the ScreeningDataStoreContentParam
// and returns a pointer to the new copy
func (sdscp ScreeningDataStoreContentParam) CopyRef() types.RVTypePtr {
	copied := sdscp
	return &copied
}

// Deref takes a pointer to the ScreeningDataStoreContentParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sdscp *ScreeningDataStoreContentParam) Deref() types.RVType {
	return *sdscp
}

// String returns the string representation of the ScreeningDataStoreContentParam
func (sdscp ScreeningDataStoreContentParam) String() string {
	return sdscp.FormatToString(0)
}

// FormatToString pretty-prints the ScreeningDataStoreContentParam using the provided indentation level
func (sdscp ScreeningDataStoreContentParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ScreeningDataStoreContentParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, sdscp.DataID))
	b.WriteString(fmt.Sprintf("%sContentDataID: %s,\n", indentationValues, sdscp.ContentDataID))
	b.WriteString(fmt.Sprintf("%sUGCType: %s,\n", indentationValues, sdscp.UGCType))
	b.WriteString(fmt.Sprintf("%sLanguage: %s,\n", indentationValues, sdscp.Language))
	b.WriteString(fmt.Sprintf("%sSearchKey: %s\n", indentationValues, sdscp.SearchKey))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewScreeningDataStoreContentParam returns a new ScreeningDataStoreContentParam
func NewScreeningDataStoreContentParam() ScreeningDataStoreContentParam {
	return ScreeningDataStoreContentParam{
		DataID:        types.NewUInt64(0),
		ContentDataID: types.NewUInt64(0),
		UGCType:       types.NewString(""),
		Language:      types.NewString(""),
		SearchKey:     types.NewString(""),
	}
}
