// Package types implements all the types used by the Matchmake Extension (Mario Kart 8) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SimpleSearchParam holds data for the Matchmake Extension (Mario Kart 8) protocol
type SimpleSearchParam struct {
	nex.Structure
	Unknown    uint32
	Unknown2   uint32
	Conditions []*SimpleSearchCondition
	Unknown3   string
	Result     *nex.ResultRange
	Unknown4   *nex.DateTime
}

// ExtractFromStream extracts a SimpleSearchParam structure from a stream
func (simpleSearchParam *SimpleSearchParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simpleSearchParam.Unknown, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown from stream. %s", err.Error())
	}

	simpleSearchParam.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown2 from stream. %s", err.Error())
	}

	conditions, err := stream.ReadListStructure(NewSimpleSearchCondition())
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Conditions from stream. %s", err.Error())
	}

	simpleSearchParam.Conditions = conditions.([]*SimpleSearchCondition)

	simpleSearchParam.Unknown3, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown3 from stream. %s", err.Error())
	}

	result, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Result from stream. %s", err.Error())
	}

	simpleSearchParam.Result = result.(*nex.ResultRange)

	simpleSearchParam.Unknown4, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract SimpleSearchParam.Unknown4 from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimpleSearchParam and returns a byte array
func (simpleSearchParam *SimpleSearchParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(simpleSearchParam.Unknown)
	stream.WriteUInt32LE(simpleSearchParam.Unknown2)
	stream.WriteListStructure(simpleSearchParam.Conditions)
	stream.WriteString(simpleSearchParam.Unknown3)
	stream.WriteStructure(simpleSearchParam.Result)
	stream.WriteDateTime(simpleSearchParam.Unknown4)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimpleSearchParam
func (simpleSearchParam *SimpleSearchParam) Copy() nex.StructureInterface {
	copied := NewSimpleSearchParam()

	copied.Unknown = simpleSearchParam.Unknown
	copied.Unknown2 = simpleSearchParam.Unknown2
	copied.Conditions = make([]*SimpleSearchCondition, len(simpleSearchParam.Conditions))

	for i := 0; i < len(simpleSearchParam.Conditions); i++ {
		copied.Conditions[i] = simpleSearchParam.Conditions[i].Copy().(*SimpleSearchCondition)
	}

	copied.Unknown3 = simpleSearchParam.Unknown3
	copied.Result = simpleSearchParam.Result.Copy().(*nex.ResultRange)
	copied.Unknown4 = simpleSearchParam.Unknown4.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simpleSearchParam *SimpleSearchParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimpleSearchParam)

	if simpleSearchParam.Unknown != other.Unknown {
		return false
	}

	if simpleSearchParam.Unknown2 != other.Unknown2 {
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

	if simpleSearchParam.Unknown3 != other.Unknown3 {
		return false
	}

	if !simpleSearchParam.Result.Equals(other.Result) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simpleSearchParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, simpleSearchParam.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, simpleSearchParam.Unknown2))

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

	if simpleSearchParam.Result != nil {
		b.WriteString(fmt.Sprintf("%sResult: %s\n", indentationValues, simpleSearchParam.Result.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sResult: nil\n", indentationValues))
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
