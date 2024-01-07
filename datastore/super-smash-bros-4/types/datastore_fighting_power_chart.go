// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreFightingPowerChart is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreFightingPowerChart struct {
	types.Structure
	UserNum *types.PrimitiveU32
	Chart   *types.List[*DataStoreFightingPowerScore]
}

// ExtractFrom extracts the DataStoreFightingPowerChart from the given readable
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreFightingPowerChart.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreFightingPowerChart header. %s", err.Error())
	}

	err = dataStoreFightingPowerChart.UserNum.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.UserNum. %s", err.Error())
	}

	err = dataStoreFightingPowerChart.Chart.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreFightingPowerChart.Chart. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreFightingPowerChart to the given writable
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreFightingPowerChart.UserNum.WriteTo(contentWritable)
	dataStoreFightingPowerChart.Chart.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreFightingPowerChart.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreFightingPowerChart
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Copy() types.RVType {
	copied := NewDataStoreFightingPowerChart()

	copied.StructureVersion = dataStoreFightingPowerChart.StructureVersion

	copied.UserNum = dataStoreFightingPowerChart.UserNum.Copy().(*types.PrimitiveU32)
	copied.Chart = dataStoreFightingPowerChart.Chart.Copy().(*types.List[*DataStoreFightingPowerScore])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreFightingPowerChart); !ok {
		return false
	}

	other := o.(*DataStoreFightingPowerChart)

	if dataStoreFightingPowerChart.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreFightingPowerChart.UserNum.Equals(other.UserNum) {
		return false
	}

	if !dataStoreFightingPowerChart.Chart.Equals(other.Chart) {
		return false
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
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreFightingPowerChart{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreFightingPowerChart.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUserNum: %s,\n", indentationValues, dataStoreFightingPowerChart.UserNum))
	b.WriteString(fmt.Sprintf("%sChart: %s\n", indentationValues, dataStoreFightingPowerChart.Chart))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreFightingPowerChart returns a new DataStoreFightingPowerChart
func NewDataStoreFightingPowerChart() *DataStoreFightingPowerChart {
	dataStoreFightingPowerChart := &DataStoreFightingPowerChart{
		UserNum: types.NewPrimitiveU32(0),
		Chart: types.NewList[*DataStoreFightingPowerScore](),
	}

	dataStoreFightingPowerChart.Chart.Type = NewDataStoreFightingPowerScore()

	return dataStoreFightingPowerChart
}
