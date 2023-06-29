package datastore_super_smash_bros_4_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStoreFightingPowerChart struct {
	nex.Structure
	UserNum uint32
	Chart   []*DataStoreFightingPowerScore
}

// ExtractFromStream extracts a DataStoreFightingPowerChart structure from a stream
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreFightingPowerChart.UserNum, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.UserNum. %s", err.Error())
	}

	chart, err := stream.ReadListStructure(NewDataStoreFightingPowerScore())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.Chart. %s", err.Error())
	}

	dataStoreFightingPowerChart.Chart = chart.([]*DataStoreFightingPowerScore)

	return nil
}

// Bytes encodes the DataStoreFightingPowerChart and returns a byte array
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreFightingPowerChart.UserNum)
	stream.WriteListStructure(dataStoreFightingPowerChart.Chart)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFightingPowerChart
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Copy() nex.StructureInterface {
	copied := NewDataStoreFightingPowerChart()

	copied.UserNum = dataStoreFightingPowerChart.UserNum
	copied.Chart = make([]*DataStoreFightingPowerScore, len(dataStoreFightingPowerChart.Chart))

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		copied.Chart[i] = dataStoreFightingPowerChart.Chart[i].Copy().(*DataStoreFightingPowerScore)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFightingPowerChart)

	if dataStoreFightingPowerChart.UserNum != other.UserNum {
		return false
	}

	if len(dataStoreFightingPowerChart.Chart) != len(other.Chart) {
		return false
	}

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		if !dataStoreFightingPowerChart.Chart[i].Equals(other.Chart[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) String() string {
	return dataStoreFightingPowerChart.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFightingPowerChart{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreFightingPowerChart.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUserNum: %d,\n", indentationValues, dataStoreFightingPowerChart.UserNum))

	if len(dataStoreFightingPowerChart.Chart) == 0 {
		b.WriteString(fmt.Sprintf("%sChart: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sChart: [\n", indentationValues))

		for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
			str := dataStoreFightingPowerChart.Chart[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreFightingPowerChart.Chart)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerChart returns a new DataStoreFightingPowerChart
func NewDataStoreFightingPowerChart() *DataStoreFightingPowerChart {
	return &DataStoreFightingPowerChart{}
}
