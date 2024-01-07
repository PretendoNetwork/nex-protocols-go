// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SimpleSearchParam holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchParam struct {
	types.Structure
	Unknown     *types.PrimitiveU32
	Unknown2    *types.PID
	Conditions  []*SimpleSearchCondition
	Unknown3    string
	ResultRange *types.ResultRange
	Unknown4    *types.DateTime
}

// ExtractFrom extracts the SimpleSearchParam from the given readable
func (simpleSearchParam *SimpleSearchParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simpleSearchParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimpleSearchParam header. %s", err.Error())
	}

	err = simpleSearchParam.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown from stream. %s", err.Error())
	}

	err = simpleSearchParam.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown2 from stream. %s", err.Error())
	}

	conditions, err := nex.StreamReadListStructure(stream, NewSimpleSearchCondition())
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Conditions from stream. %s", err.Error())
	}

	simpleSearchParam.Conditions = conditions

	err = simpleSearchParam.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown3 from stream. %s", err.Error())
	}

	err = simpleSearchParam.ResultRange.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.ResultRange from stream. %s", err.Error())
	}

	err = simpleSearchParam.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown4 from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SimpleSearchParam to the given writable
func (simpleSearchParam *SimpleSearchParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	simpleSearchParam.Unknown.WriteTo(contentWritable)
	simpleSearchParam.Unknown2.WriteTo(contentWritable)
	simpleSearchParam.Conditions.WriteTo(contentWritable)
	simpleSearchParam.Unknown3.WriteTo(contentWritable)
	simpleSearchParam.ResultRange.WriteTo(contentWritable)
	simpleSearchParam.Unknown4.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	simpleSearchParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SimpleSearchParam
func (simpleSearchParam *SimpleSearchParam) Copy() types.RVType {
	copied := NewSimpleSearchParam()

	copied.StructureVersion = simpleSearchParam.StructureVersion

	copied.Unknown = simpleSearchParam.Unknown
	copied.Unknown2 = simpleSearchParam.Unknown2.Copy()
	copied.Conditions = make([]*SimpleSearchCondition, len(simpleSearchParam.Conditions))

	for i := 0; i < len(simpleSearchParam.Conditions); i++ {
		copied.Conditions[i] = simpleSearchParam.Conditions[i].Copy().(*SimpleSearchCondition)
	}

	copied.Unknown3 = simpleSearchParam.Unknown3
	copied.ResultRange = simpleSearchParam.ResultRange.Copy().(*types.ResultRange)
	copied.Unknown4 = simpleSearchParam.Unknown4.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchParam *SimpleSearchParam) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchParam); !ok {
		return false
	}

	other := o.(*SimpleSearchParam)

	if simpleSearchParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !simpleSearchParam.Unknown.Equals(other.Unknown) {
		return false
	}

	if !simpleSearchParam.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if len(simpleSearchParam.Conditions) != len(other.Conditions) {
		return false
	}

	for i := 0; i < len(simpleSearchParam.Conditions); i++ {
		if !simpleSearchParam.Conditions[i].Equals(other.Conditions[i]) {
			return false
		}
	}

	if !simpleSearchParam.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !simpleSearchParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if !simpleSearchParam.Unknown4.Equals(other.Unknown4) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simpleSearchParam *SimpleSearchParam) String() string {
	return simpleSearchParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simpleSearchParam *SimpleSearchParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simpleSearchParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, simpleSearchParam.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, simpleSearchParam.Unknown2.FormatToString(indentationLevel+1)))

	if len(simpleSearchParam.Conditions) == 0 {
		b.WriteString(fmt.Sprintf("%sConditions: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sConditions: [\n", indentationValues))

		for i := 0; i < len(simpleSearchParam.Conditions); i++ {
			str := simpleSearchParam.Conditions[i].FormatToString(indentationLevel + 2)
			if i == len(simpleSearchParam.Conditions)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown3: %q,\n", indentationValues, simpleSearchParam.Unknown3))

	if simpleSearchParam.ResultRange != nil {
		b.WriteString(fmt.Sprintf("%sResultRange: %s\n", indentationValues, simpleSearchParam.ResultRange.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResultRange: nil\n", indentationValues))
	}

	if simpleSearchParam.Unknown4 != nil {
		b.WriteString(fmt.Sprintf("%sUnknown4: %s\n", indentationValues, simpleSearchParam.Unknown4.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown4: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchParam returns a new SimpleSearchParam
func NewSimpleSearchParam() *SimpleSearchParam {
	return &SimpleSearchParam{}
}
