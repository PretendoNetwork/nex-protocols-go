// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleSearchCondition holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchCondition struct {
	nex.Structure
	Value              uint32
	ComparisonOperator uint32
}

// ExtractFromStream extracts a SimpleSearchCondition structure from a stream
func (simpleSearchCondition *SimpleSearchCondition) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleSearchCondition.Value, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.Value from stream. %s", err.Error())
	}

	simpleSearchCondition.ComparisonOperator, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchCondition.ComparisonOperator from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimpleSearchCondition and returns a byte array
func (simpleSearchCondition *SimpleSearchCondition) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchCondition.Value)
	stream.WriteUInt32LE(simpleSearchCondition.ComparisonOperator)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchCondition
func (simpleSearchCondition *SimpleSearchCondition) Copy() nex.StructureInterface {
	copied := NewSimpleSearchCondition()

	copied.SetStructureVersion(simpleSearchCondition.StructureVersion())

	copied.Value = simpleSearchCondition.Value
	copied.ComparisonOperator = simpleSearchCondition.ComparisonOperator

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchCondition *SimpleSearchCondition) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchCondition)

	if simpleSearchCondition.StructureVersion() != other.StructureVersion() {
		return false
	}

	if simpleSearchCondition.Value != other.Value {
		return false
	}

	if simpleSearchCondition.ComparisonOperator != other.ComparisonOperator {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simpleSearchCondition.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sValue: %d,\n", indentationValues, simpleSearchCondition.Value))
	b.WriteString(fmt.Sprintf("%sComparisonOperator: %d,\n", indentationValues, simpleSearchCondition.ComparisonOperator))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimpleSearchCondition returns a new SimpleSearchCondition
func NewSimpleSearchCondition() *SimpleSearchCondition {
	return &SimpleSearchCondition{}
}
