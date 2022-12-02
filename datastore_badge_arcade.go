package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreBadgeArcadeProtocolID is the Protocol ID for the DataStore (Badge Arcade) protocol. ID is the same as the DataStore protocol
	DataStoreBadgeArcadeProtocolID = 0x73

	// DataStoreBadgeArcadeMethodGetMetaByOwnerId is the method ID for GetMetaByOwnerId
	DataStoreBadgeArcadeMethodGetMetaByOwnerId = 0x2D
)

// DataStoreBadgeArcadeProtocol handles the DataStore (Badge Arcade) nex protocol. Embeds DataStoreProtocol
type DataStoreBadgeArcadeProtocol struct {
	server *nex.Server
	DataStoreProtocol
	GetMetaByOwnerIdHandler func(err error, client *nex.Client, callID uint32, param *DataStoreGetMetaByOwnerIdParam)
}

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

// Setup initializes the protocol
func (dataStoreBadgeArcadeProtocol *DataStoreBadgeArcadeProtocol) Setup() {
	nexServer := dataStoreBadgeArcadeProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreBadgeArcadeProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DataStoreMethodPrepareGetObject:
				go dataStoreBadgeArcadeProtocol.handlePrepareGetObject(packet)
			case DataStoreMethodGetPersistenceInfo:
				go dataStoreBadgeArcadeProtocol.handleGetPersistenceInfo(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreBadgeArcadeProtocol.handleChangeMeta(packet)
			case DataStoreBadgeArcadeMethodGetMetaByOwnerId:
				go dataStoreBadgeArcadeProtocol.handleGetMetaByOwnerId(packet)
			case DataStoreMethodPrepareUpdateObject:
				go dataStoreBadgeArcadeProtocol.handlePrepareUpdateObject(packet)
			case DataStoreMethodCompleteUpdateObject:
				go dataStoreBadgeArcadeProtocol.handleCompleteUpdateObject(packet)
			default:
				go respondNotImplemented(packet, DataStoreBadgeArcadeProtocolID)
				fmt.Printf("Unsupported DataStoreBadgeArcade method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetMetaByOwnerId sets the GetMetaByOwnerId function
func (dataStoreBadgeArcadeProtocol *DataStoreBadgeArcadeProtocol) GetMetaByOwnerId(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetMetaByOwnerIdParam)) {
	dataStoreBadgeArcadeProtocol.GetMetaByOwnerIdHandler = handler
}

func (dataStoreBadgeArcadeProtocol *DataStoreBadgeArcadeProtocol) handleGetMetaByOwnerId(packet nex.PacketInterface) {
	if dataStoreBadgeArcadeProtocol.GetMetaByOwnerIdHandler == nil {
		logger.Warning("DataStoreBadgeArcadeProtocol::GetMetaByOwnerId not implemented")
		go respondNotImplemented(packet, DataStoreBadgeArcadeProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreBadgeArcadeProtocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetMetaByOwnerIdParam())
	if err != nil {
		go dataStoreBadgeArcadeProtocol.GetMetaByOwnerIdHandler(err, client, callID, nil)
		return
	}

	go dataStoreBadgeArcadeProtocol.GetMetaByOwnerIdHandler(nil, client, callID, param.(*DataStoreGetMetaByOwnerIdParam))
}

// NewDataStoreBadgeArcadeProtocol returns a new DataStoreBadgeArcadeProtocol
func NewDataStoreBadgeArcadeProtocol(server *nex.Server) *DataStoreBadgeArcadeProtocol {
	dataStoreBadgeArcadeProtocol := &DataStoreBadgeArcadeProtocol{server: server}
	dataStoreBadgeArcadeProtocol.DataStoreProtocol.server = server

	dataStoreBadgeArcadeProtocol.Setup()

	return dataStoreBadgeArcadeProtocol
}
