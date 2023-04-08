package datastore_nintendo_badge_arcade

import nex "github.com/PretendoNetwork/nex-go"

type DataStoreGetMetaByOwnerIdParam struct {
	nex.Structure
	OwnerIDs     []uint32
	DataTypes    []uint16
	ResultOption uint8
	ResultRange  *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreGetMetaByOwnerIdParam structure from a stream
func (dataStoreGetMetaByOwnerIdParam *DataStoreGetMetaByOwnerIdParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetMetaByOwnerIdParam.OwnerIDs = stream.ReadListUInt32LE()
	dataStoreGetMetaByOwnerIdParam.DataTypes = stream.ReadListUInt16LE()
	dataStoreGetMetaByOwnerIdParam.ResultOption = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return err
	}

	dataStoreGetMetaByOwnerIdParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreGetMetaByOwnerIdParam and returns a byte array
func (dataStoreGetMetaByOwnerIdParam *DataStoreGetMetaByOwnerIdParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetMetaByOwnerIdParam.OwnerIDs)
	stream.WriteListUInt16LE(dataStoreGetMetaByOwnerIdParam.DataTypes)
	stream.WriteUInt8(dataStoreGetMetaByOwnerIdParam.ResultOption)
	stream.WriteStructure(dataStoreGetMetaByOwnerIdParam.ResultRange)

	return stream.Bytes()
}

// NewDataStoreGetMetaByOwnerIdParam returns a new DataStoreGetMetaByOwnerIdParam
func NewDataStoreGetMetaByOwnerIdParam() *DataStoreGetMetaByOwnerIdParam {
	return &DataStoreGetMetaByOwnerIdParam{}
}
