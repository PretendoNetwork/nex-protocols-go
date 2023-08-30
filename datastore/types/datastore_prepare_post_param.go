// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePreparePostParam is sent in the PreparePostObject method
type DataStorePreparePostParam struct {
	nex.Structure
	Size                 uint32
	Name                 string
	DataType             uint16
	MetaBinary           []byte
	Permission           *DataStorePermission
	DelPermission        *DataStorePermission
	Flag                 uint32
	Period               uint16
	ReferDataID          uint32
	Tags                 []string
	RatingInitParams     []*DataStoreRatingInitParamWithSlot
	PersistenceInitParam *DataStorePersistenceInitParam
	ExtraData            []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePreparePostParam structure from a stream
func (dataStorePreparePostParam *DataStorePreparePostParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	dataStorePreparePostParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Size. %s", err.Error())
	}

	dataStorePreparePostParam.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Name. %s", err.Error())
	}

	dataStorePreparePostParam.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DataType. %s", err.Error())
	}

	dataStorePreparePostParam.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.MetaBinary. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Permission. %s", err.Error())
	}

	dataStorePreparePostParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.DelPermission. %s", err.Error())
	}

	dataStorePreparePostParam.DelPermission = delPermission.(*DataStorePermission)
	dataStorePreparePostParam.Flag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Flag. %s", err.Error())
	}

	dataStorePreparePostParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Period. %s", err.Error())
	}

	dataStorePreparePostParam.ReferDataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.ReferDataID. %s", err.Error())
	}

	dataStorePreparePostParam.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.Tags. %s", err.Error())
	}

	ratingInitParams, err := stream.ReadListStructure(NewDataStoreRatingInitParamWithSlot())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.RatingInitParams. %s", err.Error())
	}

	dataStorePreparePostParam.RatingInitParams = ratingInitParams.([]*DataStoreRatingInitParamWithSlot)

	persistenceInitParam, err := stream.ReadStructure(NewDataStorePersistenceInitParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePreparePostParam.PersistenceInitParam. %s", err.Error())
	}

	dataStorePreparePostParam.PersistenceInitParam = persistenceInitParam.(*DataStorePersistenceInitParam)

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStorePreparePostParam.ExtraData, err = stream.ReadListString()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePreparePostParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostParam
func (dataStorePreparePostParam *DataStorePreparePostParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostParam()

	copied.Size = dataStorePreparePostParam.Size
	copied.Name = dataStorePreparePostParam.Name
	copied.DataType = dataStorePreparePostParam.DataType
	copied.MetaBinary = make([]byte, len(dataStorePreparePostParam.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostParam.MetaBinary)

	copied.Permission = dataStorePreparePostParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParam.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParam.Flag
	copied.Period = dataStorePreparePostParam.Period
	copied.ReferDataID = dataStorePreparePostParam.ReferDataID
	copied.Tags = make([]string, len(dataStorePreparePostParam.Tags))

	copy(copied.Tags, dataStorePreparePostParam.Tags)

	copied.RatingInitParams = make([]*DataStoreRatingInitParamWithSlot, len(dataStorePreparePostParam.RatingInitParams))

	for i := 0; i < len(dataStorePreparePostParam.RatingInitParams); i++ {
		copied.RatingInitParams[i] = dataStorePreparePostParam.RatingInitParams[i].Copy().(*DataStoreRatingInitParamWithSlot)
	}

	copied.PersistenceInitParam = dataStorePreparePostParam.PersistenceInitParam.Copy().(*DataStorePersistenceInitParam)
	copied.ExtraData = make([]string, len(dataStorePreparePostParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParam *DataStorePreparePostParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostParam)

	if dataStorePreparePostParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostParam.Name != other.Name {
		return false
	}

	if dataStorePreparePostParam.DataType != other.DataType {
		return false
	}

	if !bytes.Equal(dataStorePreparePostParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParam.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStorePreparePostParam.Flag != other.Flag {
		return false
	}

	if dataStorePreparePostParam.Period != other.Period {
		return false
	}

	if dataStorePreparePostParam.ReferDataID != other.ReferDataID {
		return false
	}

	if len(dataStorePreparePostParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.Tags); i++ {
		if dataStorePreparePostParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(dataStorePreparePostParam.RatingInitParams) != len(other.RatingInitParams) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.RatingInitParams); i++ {
		if !dataStorePreparePostParam.RatingInitParams[i].Equals(other.RatingInitParams[i]) {
			return false
		}
	}

	if !dataStorePreparePostParam.PersistenceInitParam.Equals(other.PersistenceInitParam) {
		return false
	}

	if len(dataStorePreparePostParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.ExtraData); i++ {
		if dataStorePreparePostParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePreparePostParam *DataStorePreparePostParam) String() string {
	return dataStorePreparePostParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePreparePostParam *DataStorePreparePostParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePreparePostParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePreparePostParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStorePreparePostParam.Size))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStorePreparePostParam.Name))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStorePreparePostParam.DataType))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStorePreparePostParam.MetaBinary))

	if dataStorePreparePostParam.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStorePreparePostParam.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStorePreparePostParam.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStorePreparePostParam.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sFlag: %d,\n", indentationValues, dataStorePreparePostParam.Flag))
	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStorePreparePostParam.Period))
	b.WriteString(fmt.Sprintf("%sReferDataID: %d,\n", indentationValues, dataStorePreparePostParam.ReferDataID))
	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStorePreparePostParam.Tags))

	if len(dataStorePreparePostParam.RatingInitParams) == 0 {
		b.WriteString(fmt.Sprintf("%sRatingInitParams: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRatingInitParams: [\n", indentationValues))

		for i := 0; i < len(dataStorePreparePostParam.RatingInitParams); i++ {
			str := dataStorePreparePostParam.RatingInitParams[i].FormatToString(indentationLevel + 2)
			if i == len(dataStorePreparePostParam.RatingInitParams)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if dataStorePreparePostParam.PersistenceInitParam != nil {
		b.WriteString(fmt.Sprintf("%sPersistenceInitParam: %s,\n", indentationValues, dataStorePreparePostParam.PersistenceInitParam.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPersistenceInitParam: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sExtraData: %v,\n", indentationValues, dataStorePreparePostParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePreparePostParam returns a new DataStorePreparePostParam
func NewDataStorePreparePostParam() *DataStorePreparePostParam {
	return &DataStorePreparePostParam{}
}
