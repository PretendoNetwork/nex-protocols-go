package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreProtocolID is the protocol ID for the DataStore protocol
	DataStoreProtocolID = 0x73

	// DataStoreMethodPrepareGetObjectV1 is the method ID for the method PrepareGetObjectV1
	DataStoreMethodPrepareGetObjectV1 = 0x1

	// DataStoreMethodPreparePostObjectV1 is the method ID for the method PreparePostObjectV1
	DataStoreMethodPreparePostObjectV1 = 0x2

	// DataStoreMethodCompletePostObjectV1 is the method ID for the method CompletePostObjectV1
	DataStoreMethodCompletePostObjectV1 = 0x3

	// DataStoreMethodDeleteObject is the method ID for the method DeleteObject
	DataStoreMethodDeleteObject = 0x4

	// DataStoreMethodDeleteObjects is the method ID for the method DeleteObjects
	DataStoreMethodDeleteObjects = 0x5

	// DataStoreMethodChangeMetaV1 is the method ID for the method ChangeMetaV1
	DataStoreMethodChangeMetaV1 = 0x6

	// DataStoreMethodChangeMetasV1 is the method ID for the method ChangeMetasV1
	DataStoreMethodChangeMetasV1 = 0x7

	// DataStoreMethodGetMeta is the method ID for the method GetMeta
	DataStoreMethodGetMeta = 0x8

	// DataStoreMethodGetMetas is the method ID for the method GetMetas
	DataStoreMethodGetMetas = 0x9

	// DataStoreMethodPrepareUpdateObject is the method ID for the method PrepareUpdateObject
	DataStoreMethodPrepareUpdateObject = 0xA

	// DataStoreMethodCompleteUpdateObject is the method ID for the method CompleteUpdateObject
	DataStoreMethodCompleteUpdateObject = 0xB

	// DataStoreMethodSearchObject is the method ID for the method SearchObject
	DataStoreMethodSearchObject = 0xC

	// DataStoreMethodGetNotificationURL is the method ID for the method GetNotificationUrl
	DataStoreMethodGetNotificationURL = 0xD

	// DataStoreMethodGetNewArrivedNotificationsV1 is the method ID for the method GetNewArrivedNotificationsV1
	DataStoreMethodGetNewArrivedNotificationsV1 = 0xE

	// DataStoreMethodRateObject is the method ID for the method RateObject
	DataStoreMethodRateObject = 0xF

	// DataStoreMethodGetRating is the method ID for the method GetRating
	DataStoreMethodGetRating = 0x10

	// DataStoreMethodGetRatings is the method ID for the method GetRatings
	DataStoreMethodGetRatings = 0x11

	// DataStoreMethodResetRating is the method ID for the method ResetRating
	DataStoreMethodResetRating = 0x12

	// DataStoreMethodResetRatings is the method ID for the method ResetRatings
	DataStoreMethodResetRatings = 0x13

	// DataStoreMethodGetSpecificMetaV1 is the method ID for the method GetSpecificMetaV1
	DataStoreMethodGetSpecificMetaV1 = 0x14

	// DataStoreMethodPostMetaBinary is the method ID for the method PostMetaBinary
	DataStoreMethodPostMetaBinary = 0x15

	// DataStoreMethodTouchObject is the method ID for the method TouchObject
	DataStoreMethodTouchObject = 0x16

	// DataStoreMethodGetRatingWithLog is the method ID for the method GetRatingWithLog
	DataStoreMethodGetRatingWithLog = 0x17

	// DataStoreMethodPreparePostObject is the method ID for the method PreparePostObject
	DataStoreMethodPreparePostObject = 0x18

	// DataStoreMethodPrepareGetObject is the method ID for the method PrepareGetObject
	DataStoreMethodPrepareGetObject = 0x19

	// DataStoreMethodCompletePostObject is the method ID for the method CompletePostObject
	DataStoreMethodCompletePostObject = 0x1A

	// DataStoreMethodGetNewArrivedNotifications is the method ID for the method GetNewArrivedNotifications
	DataStoreMethodGetNewArrivedNotifications = 0x1B

	// DataStoreMethodGetSpecificMeta is the method ID for the method GetSpecificMeta
	DataStoreMethodGetSpecificMeta = 0x1C

	// DataStoreMethodGetPersistenceInfo is the method ID for the method GetPersistenceInfo
	DataStoreMethodGetPersistenceInfo = 0x1D

	// DataStoreMethodGetPersistenceInfos is the method ID for the method GetPersistenceInfos
	DataStoreMethodGetPersistenceInfos = 0x1E

	// DataStoreMethodPerpetuateObject is the method ID for the method PerpetuateObject
	DataStoreMethodPerpetuateObject = 0x1F

	// DataStoreMethodUnperpetuateObject is the method ID for the method UnperpetuateObject
	DataStoreMethodUnperpetuateObject = 0x20

	// DataStoreMethodPrepareGetObjectOrMetaBinary is the method ID for the method PrepareGetObjectOrMetaBinary
	DataStoreMethodPrepareGetObjectOrMetaBinary = 0x21

	// DataStoreMethodGetPasswordInfo is the method ID for the method GetPasswordInfo
	DataStoreMethodGetPasswordInfo = 0x22

	// DataStoreMethodGetPasswordInfos is the method ID for the method GetPasswordInfos
	DataStoreMethodGetPasswordInfos = 0x23

	// DataStoreMethodGetMetasMultipleParam is the method ID for the method GetMetasMultipleParam
	DataStoreMethodGetMetasMultipleParam = 0x24

	// DataStoreMethodCompletePostObjects is the method ID for the method CompletePostObjects
	DataStoreMethodCompletePostObjects = 0x25

	// DataStoreMethodChangeMeta is the method ID for the method ChangeMeta
	DataStoreMethodChangeMeta = 0x26

	// DataStoreMethodChangeMetas is the method ID for the method ChangeMetas
	DataStoreMethodChangeMetas = 0x27

	// DataStoreMethodRateObjects is the method ID for the method RateObjects
	DataStoreMethodRateObjects = 0x28

	// DataStoreMethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataId
	DataStoreMethodPostMetaBinaryWithDataID = 0x29

	// DataStoreMethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataId
	DataStoreMethodPostMetaBinariesWithDataID = 0x2A

	// DataStoreMethodRateObjectWithPosting is the method ID for the method RateObjectWithPosting
	DataStoreMethodRateObjectWithPosting = 0x2B

	// DataStoreMethodRateObjectsWithPosting is the method ID for the method RateObjectsWithPosting
	DataStoreMethodRateObjectsWithPosting = 0x2C

	// DataStoreMethodGetObjectInfos is the method ID for the method GetObjectInfos
	DataStoreMethodGetObjectInfos = 0x2D

	// DataStoreMethodSearchObjectLight is the method ID for the method SearchObjectLight
	DataStoreMethodSearchObjectLight = 0x2E
)

// DataStoreProtocol handles the DataStore nex protocol
type DataStoreProtocol struct {
	server         *nex.Server
	GetMetaHandler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)
}

// DataStoreGetMetaParam is sent in the GetMeta method
type DataStoreGetMetaParam struct {
	DataID            uint64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      uint8
	AccessPassword    uint64

	hierarchy []nex.StructureInterface
	nex.Structure
}

// ExtractFromStream extracts a DataStoreGetMetaParam structure from a stream
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

	dataStoreGetMetaParam.DataID = dataID
	dataStoreGetMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStoreGetMetaParam.ResultOption = stream.ReadUInt8()
	dataStoreGetMetaParam.AccessPassword = stream.ReadUInt64LE()

	return nil
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{}
}

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	OwnerID           uint32
	PersistenceSlotID uint16

	hierarchy []nex.StructureInterface
	nex.Structure
}

// ExtractFromStream extracts a DataStorePersistenceTarget structure from a stream
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	dataStorePersistenceTarget.OwnerID = stream.ReadUInt32LE()
	dataStorePersistenceTarget.PersistenceSlotID = stream.ReadUInt16LE()

	return nil
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{}
}

// DataStoreMetaInfo contains DataStore meta information
type DataStoreMetaInfo struct {
}

// Bytes encodes the DataStoreMetaInfo and returns a byte array
func (dataStoreMetaInfo *DataStoreMetaInfo) Bytes() []byte {
	return make([]byte, 1)
}

// Setup initializes the protocol
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

// GetMeta sets the GetMeta handler function
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
		go dataStoreProtocol.GetMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.GetMetaHandler(nil, client, callID, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
}

// NewDataStoreProtocol returns a new DataStoreProtocol
func NewDataStoreProtocol(server *nex.Server) *DataStoreProtocol {
	dataStoreProtocol := &DataStoreProtocol{server: server}

	dataStoreProtocol.Setup()

	return dataStoreProtocol
}
