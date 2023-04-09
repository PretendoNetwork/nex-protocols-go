package datastore_nintendo_badge_arcade

import nex "github.com/PretendoNetwork/nex-go"

type DataStoreGetMetaByOwnerIDParam struct {
	nex.Structure
	OwnerIDs     []uint32
	DataTypes    []uint16
	ResultOption uint8
	ResultRange  *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreGetMetaByOwnerIDParam structure from a stream
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetMetaByOwnerIDParam.OwnerIDs = stream.ReadListUInt32LE()
	dataStoreGetMetaByOwnerIDParam.DataTypes = stream.ReadListUInt16LE()
	dataStoreGetMetaByOwnerIDParam.ResultOption = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return err
	}

	dataStoreGetMetaByOwnerIDParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreGetMetaByOwnerIDParam and returns a byte array
func (dataStoreGetMetaByOwnerIDParam *DataStoreGetMetaByOwnerIDParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetMetaByOwnerIDParam.OwnerIDs)
	stream.WriteListUInt16LE(dataStoreGetMetaByOwnerIDParam.DataTypes)
	stream.WriteUInt8(dataStoreGetMetaByOwnerIDParam.ResultOption)
	stream.WriteStructure(dataStoreGetMetaByOwnerIDParam.ResultRange)

	return stream.Bytes()
}

// NewDataStoreGetMetaByOwnerIDParam returns a new DataStoreGetMetaByOwnerIDParam
func NewDataStoreGetMetaByOwnerIDParam() *DataStoreGetMetaByOwnerIDParam {
	return &DataStoreGetMetaByOwnerIDParam{}
}
