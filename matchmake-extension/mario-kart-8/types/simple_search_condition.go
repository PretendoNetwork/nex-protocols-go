// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SimpleSearchCondition holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchCondition struct {
	types.Structure
	Value              *types.PrimitiveU32
	ComparisonOperator *types.PrimitiveU32
}

// ExtractFrom extracts the SimpleSearchCondition from the given readable
func (simpleSearchCondition *SimpleSearchCondition) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simpleSearchCondition.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimpleSearchCondition header. %s", err.Error())
	}

	err = simpleSearchCondition.Value.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.Value from stream. %s", err.Error())
	}

	err = simpleSearchCondition.ComparisonOperator.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.ComparisonOperator from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SimpleSearchCondition to the given writable
func (simpleSearchCondition *SimpleSearchCondition) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	simpleSearchCondition.Value.WriteTo(contentWritable)
	simpleSearchCondition.ComparisonOperator.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	simpleSearchCondition.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SimpleSearchCondition
func (simpleSearchCondition *SimpleSearchCondition) Copy() types.RVType {
	copied := NewSimpleSearchCondition()

	copied.StructureVersion = simpleSearchCondition.StructureVersion

	copied.Value = simpleSearchCondition.Value
	copied.ComparisonOperator = simpleSearchCondition.ComparisonOperator

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchCondition *SimpleSearchCondition) Equals(o types.RVType) bool {
	if _, ok := o.(*SimpleSearchCondition); !ok {
		return false
	}

	other := o.(*SimpleSearchCondition)

	if simpleSearchCondition.StructureVersion != other.StructureVersion {
		return false
	}

	if !simpleSearchCondition.Value.Equals(other.Value) {
		return false
	}

	if !simpleSearchCondition.ComparisonOperator.Equals(other.ComparisonOperator) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simpleSearchCondition *SimpleSearchCondition) String() string {
	return simpleSearchCondition.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simpleSearchCondition *SimpleSearchCondition) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimpleSearchCondition{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simpleSearchCondition.StructureVersion))
	b.WriteString(fmt.Sprintf("%sValue: %d,\n", indentationValues, simpleSearchCondition.Value))
	b.WriteString(fmt.Sprintf("%sComparisonOperator: %d,\n", indentationValues, simpleSearchCondition.ComparisonOperator))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchCondition returns a new SimpleSearchCondition
func NewSimpleSearchCondition() *SimpleSearchCondition {
	return &SimpleSearchCondition{}
}
