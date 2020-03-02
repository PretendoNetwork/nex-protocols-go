package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	DataStoreProtocolID = 0x73

	DataStoreMethodPrepareGetObjectV1           = 0x1
	DataStoreMethodPreparePostObjectV1          = 0x2
	DataStoreMethodCompletePostObjectV1         = 0x3
	DataStoreMethodDeleteObject                 = 0x4
	DataStoreMethodDeleteObjects                = 0x5
	DataStoreMethodChangeMetaV1                 = 0x6
	DataStoreMethodChangeMetasV1                = 0x7
	DataStoreMethodGetMeta                      = 0x8
	DataStoreMethodGetMetas                     = 0x9
	DataStoreMethodPrepareUpdateObject          = 0xA
	DataStoreMethodCompleteUpdateObject         = 0xB
	DataStoreMethodSearchObject                 = 0xC
	DataStoreMethodGetNotificationUrl           = 0xD
	DataStoreMethodGetNewArrivedNotificationsV1 = 0xE
	DataStoreMethodRateObject                   = 0xF
	DataStoreMethodGetRating                    = 0x10
	DataStoreMethodGetRatings                   = 0x11
	DataStoreMethodResetRating                  = 0x12
	DataStoreMethodResetRatings                 = 0x13
	DataStoreMethodGetSpecificMetaV1            = 0x14
	DataStoreMethodPostMetaBinary               = 0x15
	DataStoreMethodTouchObject                  = 0x16
	DataStoreMethodGetRatingWithLog             = 0x17
	DataStoreMethodPreparePostObject            = 0x18
	DataStoreMethodPrepareGetObject             = 0x19
	DataStoreMethodCompletePostObject           = 0x1A
	DataStoreMethodGetNewArrivedNotifications   = 0x1B
	DataStoreMethodGetSpecificMeta              = 0x1C
	DataStoreMethodGetPersistenceInfo           = 0x1D
	DataStoreMethodGetPersistenceInfos          = 0x1E
	DataStoreMethodPerpetuateObject             = 0x1F
	DataStoreMethodUnperpetuateObject           = 0x20
	DataStoreMethodPrepareGetObjectOrMetaBinary = 0x21
	DataStoreMethodGetPasswordInfo              = 0x22
	DataStoreMethodGetPasswordInfos             = 0x23
	DataStoreMethodGetMetasMultipleParam        = 0x24
	DataStoreMethodCompletePostObjects          = 0x25
	DataStoreMethodChangeMeta                   = 0x26
	DataStoreMethodChangeMetas                  = 0x27
	DataStoreMethodRateObjects                  = 0x28
	DataStoreMethodPostMetaBinaryWithDataId     = 0x29
	DataStoreMethodPostMetaBinariesWithDataId   = 0x2A
	DataStoreMethodRateObjectWithPosting        = 0x2B
	DataStoreMethodRateObjectsWithPosting       = 0x2C
	DataStoreMethodGetObjectInfos               = 0x2D
	DataStoreMethodSearchObjectLight            = 0x2E
)

type DataStoreProtocol struct {
	server         *nex.Server
	GetMetaHandler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)
}

type DataStoreGetMetaParam struct {
	dataID            uint64
	persistenceTarget *DataStorePersistenceTarget
	resultOption      uint8
	accessPassword    uint64

	hierarchy []nex.StructureInterface
	nex.Structure
}

func (dataStoreGetMetaParam *DataStoreGetMetaParam) GetDataID() uint64 {
	return dataStoreGetMetaParam.dataID
}

func (dataStoreGetMetaParam *DataStoreGetMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 23 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStoreGetMetaParam::ExtractFromStream] Data size too small")
	}

	dataID := stream.ReadUInt64LE()
	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())

	if err != nil {
		return err
	}

	dataStoreGetMetaParam.dataID = dataID
	dataStoreGetMetaParam.persistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStoreGetMetaParam.resultOption = stream.ReadUInt8()
	dataStoreGetMetaParam.accessPassword = stream.ReadUInt64LE()

	return nil
}

func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	dataStoreGetMetaParam := &DataStoreGetMetaParam{}

	return dataStoreGetMetaParam
}

type DataStorePersistenceTarget struct {
	ownerID           uint32
	persistenceSlotID uint16

	hierarchy []nex.StructureInterface
	nex.Structure
}

func (dataStorePersistenceTarget *DataStorePersistenceTarget) GetOwnerID() uint32 {
	return dataStorePersistenceTarget.ownerID
}

func (dataStorePersistenceTarget *DataStorePersistenceTarget) GetPersistenceSlotID() uint16 {
	return dataStorePersistenceTarget.persistenceSlotID
}

func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	dataStorePersistenceTarget.ownerID = stream.ReadUInt32LE()
	dataStorePersistenceTarget.persistenceSlotID = stream.ReadUInt16LE()

	return nil
}

func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	dataStorePersistenceTarget := &DataStorePersistenceTarget{}

	return dataStorePersistenceTarget
}

type DataStoreMetaInfo struct {
}

func (dataStoreMetaInfo *DataStoreMetaInfo) Bytes() []byte {
	return make([]byte, 1)
}

func (dataStoreProtocol *DataStoreProtocol) Setup() {
	nexServer := dataStoreProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if DataStoreProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreProtocol.handleGetMeta(packet)
			default:
				fmt.Printf("Unsupported DataStore method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (dataStoreProtocol *DataStoreProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(DataStoreProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	dataStoreProtocol.server.Send(responsePacket)
}

func (dataStoreProtocol *DataStoreProtocol) GetMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)) {
	dataStoreProtocol.GetMetaHandler = handler
}

func (dataStoreProtocol *DataStoreProtocol) handleGetMeta(packet nex.PacketInterface) {
	if dataStoreProtocol.GetMetaHandler == nil {
		fmt.Println("[Warning] DataStoreProtocol::GetMeta not implemented")
		go dataStoreProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreGetMetaParam, err := parametersStream.ReadStructure(NewDataStoreGetMetaParam())

	if err != nil {
		go dataStoreProtocol.GetMetaHandler(err, client, callID, &DataStoreGetMetaParam{})
		return
	}

	go dataStoreProtocol.GetMetaHandler(nil, client, callID, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
}

func NewDataStoreProtocol(server *nex.Server) *DataStoreProtocol {
	dataStoreProtocol := &DataStoreProtocol{server: server}

	dataStoreProtocol.Setup()

	return dataStoreProtocol
}
