package datastore_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreChangeMetaParamV1 is a data structure used by the DataStore protocol
type DataStoreChangeMetaParamV1 struct {
	nex.Structure
	DataID         uint64
	ModifiesFlag   uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStoreChangeMetaParamV1 structure from a stream
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreChangeMetaParamV1.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DataID. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.ModifiesFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.ModifiesFlag. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Name. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Permission. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.DelPermission. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaParamV1.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Period. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.MetaBinary. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.Tags. %s", err.Error())
	}

	dataStoreChangeMetaParamV1.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreChangeMetaParamV1.UpdatePassword. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreChangeMetaParamV1 and returns a byte array
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreChangeMetaParamV1.DataID)
	stream.WriteUInt32LE(dataStoreChangeMetaParamV1.ModifiesFlag)
	stream.WriteString(dataStoreChangeMetaParamV1.Name)
	stream.WriteStructure(dataStoreChangeMetaParamV1.Permission)
	stream.WriteStructure(dataStoreChangeMetaParamV1.DelPermission)
	stream.WriteUInt16LE(dataStoreChangeMetaParamV1.Period)
	stream.WriteQBuffer(dataStoreChangeMetaParamV1.MetaBinary)
	stream.WriteListString(dataStoreChangeMetaParamV1.Tags)
	stream.WriteUInt64LE(dataStoreChangeMetaParamV1.UpdatePassword)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParamV1()

	copied.DataID = dataStoreChangeMetaParamV1.DataID
	copied.ModifiesFlag = dataStoreChangeMetaParamV1.ModifiesFlag
	copied.Name = dataStoreChangeMetaParamV1.Name
	copied.Permission = dataStoreChangeMetaParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParamV1.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaParamV1.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaParamV1.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaParamV1.Tags))

	copy(copied.Tags, dataStoreChangeMetaParamV1.Tags)

	copied.UpdatePassword = dataStoreChangeMetaParamV1.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParamV1)

	if dataStoreChangeMetaParamV1.DataID != other.DataID {
		return false
	}

	if dataStoreChangeMetaParamV1.ModifiesFlag != other.ModifiesFlag {
		return false
	}

	if dataStoreChangeMetaParamV1.Name != other.Name {
		return false
	}

	if !dataStoreChangeMetaParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStoreChangeMetaParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaParamV1.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaParamV1.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaParamV1.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaParamV1.Tags); i++ {
		if dataStoreChangeMetaParamV1.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaParamV1.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) String() string {
	return dataStoreChangeMetaParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreChangeMetaParamV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreChangeMetaParamV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreChangeMetaParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sModifiesFlag: %d,\n", indentationValues, dataStoreChangeMetaParamV1.ModifiesFlag))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStoreChangeMetaParamV1.Name))

	if dataStoreChangeMetaParamV1.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreChangeMetaParamV1.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStoreChangeMetaParamV1.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreChangeMetaParamV1.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStoreChangeMetaParamV1.Period))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStoreChangeMetaParamV1.MetaBinary))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreChangeMetaParamV1.Tags))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %d\n", indentationValues, dataStoreChangeMetaParamV1.UpdatePassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreChangeMetaParamV1 returns a new DataStoreChangeMetaParamV1
func NewDataStoreChangeMetaParamV1() *DataStoreChangeMetaParamV1 {
	return &DataStoreChangeMetaParamV1{}
}
