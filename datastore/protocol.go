package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the DataStore protocol
	ProtocolID = 0x73

	// MethodPrepareGetObjectV1 is the method ID for the method PrepareGetObjectV1
	MethodPrepareGetObjectV1 = 0x1

	// MethodPreparePostObjectV1 is the method ID for the method PreparePostObjectV1
	MethodPreparePostObjectV1 = 0x2

	// MethodCompletePostObjectV1 is the method ID for the method CompletePostObjectV1
	MethodCompletePostObjectV1 = 0x3

	// MethodDeleteObject is the method ID for the method DeleteObject
	MethodDeleteObject = 0x4

	// MethodDeleteObjects is the method ID for the method DeleteObjects
	MethodDeleteObjects = 0x5

	// MethodChangeMetaV1 is the method ID for the method ChangeMetaV1
	MethodChangeMetaV1 = 0x6

	// MethodChangeMetasV1 is the method ID for the method ChangeMetasV1
	MethodChangeMetasV1 = 0x7

	// MethodGetMeta is the method ID for the method GetMeta
	MethodGetMeta = 0x8

	// MethodGetMetas is the method ID for the method GetMetas
	MethodGetMetas = 0x9

	// MethodPrepareUpdateObject is the method ID for the method PrepareUpdateObject
	MethodPrepareUpdateObject = 0xA

	// MethodCompleteUpdateObject is the method ID for the method CompleteUpdateObject
	MethodCompleteUpdateObject = 0xB

	// MethodSearchObject is the method ID for the method SearchObject
	MethodSearchObject = 0xC

	// MethodGetNotificationUrl is the method ID for the method GetNotificationUrl
	MethodGetNotificationUrl = 0xD

	// MethodGetNewArrivedNotificationsV1 is the method ID for the method GetNewArrivedNotificationsV1
	MethodGetNewArrivedNotificationsV1 = 0xE

	// MethodRateObject is the method ID for the method RateObject
	MethodRateObject = 0xF

	// MethodGetRating is the method ID for the method GetRating
	MethodGetRating = 0x10

	// MethodGetRatings is the method ID for the method GetRatings
	MethodGetRatings = 0x11

	// MethodResetRating is the method ID for the method ResetRating
	MethodResetRating = 0x12

	// MethodResetRatings is the method ID for the method ResetRatings
	MethodResetRatings = 0x13

	// MethodGetSpecificMetaV1 is the method ID for the method GetSpecificMetaV1
	MethodGetSpecificMetaV1 = 0x14

	// MethodPostMetaBinary is the method ID for the method PostMetaBinary
	MethodPostMetaBinary = 0x15

	// MethodTouchObject is the method ID for the method TouchObject
	MethodTouchObject = 0x16

	// MethodGetRatingWithLog is the method ID for the method GetRatingWithLog
	MethodGetRatingWithLog = 0x17

	// MethodPreparePostObject is the method ID for the method PreparePostObject
	MethodPreparePostObject = 0x18

	// MethodPrepareGetObject is the method ID for the method PrepareGetObject
	MethodPrepareGetObject = 0x19

	// MethodCompletePostObject is the method ID for the method CompletePostObject
	MethodCompletePostObject = 0x1A

	// MethodGetNewArrivedNotifications is the method ID for the method GetNewArrivedNotifications
	MethodGetNewArrivedNotifications = 0x1B

	// MethodGetSpecificMeta is the method ID for the method GetSpecificMeta
	MethodGetSpecificMeta = 0x1C

	// MethodGetPersistenceInfo is the method ID for the method GetPersistenceInfo
	MethodGetPersistenceInfo = 0x1D

	// MethodGetPersistenceInfos is the method ID for the method GetPersistenceInfos
	MethodGetPersistenceInfos = 0x1E

	// MethodPerpetuateObject is the method ID for the method PerpetuateObject
	MethodPerpetuateObject = 0x1F

	// MethodUnperpetuateObject is the method ID for the method UnperpetuateObject
	MethodUnperpetuateObject = 0x20

	// MethodPrepareGetObjectOrMetaBinary is the method ID for the method PrepareGetObjectOrMetaBinary
	MethodPrepareGetObjectOrMetaBinary = 0x21

	// MethodGetPasswordInfo is the method ID for the method GetPasswordInfo
	MethodGetPasswordInfo = 0x22

	// MethodGetPasswordInfos is the method ID for the method GetPasswordInfos
	MethodGetPasswordInfos = 0x23

	// MethodGetMetasMultipleParam is the method ID for the method GetMetasMultipleParam
	MethodGetMetasMultipleParam = 0x24

	// MethodCompletePostObjects is the method ID for the method CompletePostObjects
	MethodCompletePostObjects = 0x25

	// MethodChangeMeta is the method ID for the method ChangeMeta
	MethodChangeMeta = 0x26

	// MethodChangeMetas is the method ID for the method ChangeMetas
	MethodChangeMetas = 0x27

	// MethodRateObjects is the method ID for the method RateObjects
	MethodRateObjects = 0x28

	// MethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataId
	MethodPostMetaBinaryWithDataID = 0x29

	// MethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataId
	MethodPostMetaBinariesWithDataID = 0x2A

	// MethodRateObjectWithPosting is the method ID for the method RateObjectWithPosting
	MethodRateObjectWithPosting = 0x2B

	// MethodRateObjectsWithPosting is the method ID for the method RateObjectsWithPosting
	MethodRateObjectsWithPosting = 0x2C

	// MethodGetObjectInfos is the method ID for the method GetObjectInfos
	MethodGetObjectInfos = 0x2D

	// MethodSearchObjectLight is the method ID for the method SearchObjectLight
	MethodSearchObjectLight = 0x2E
)

// DataStoreProtocol handles the DataStore nex protocol
type DataStoreProtocol struct {
	Server                              *nex.Server
	PrepareGetObjectV1Handler           func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1)
	PreparePostObjectV1Handler          func(err error, client *nex.Client, callID uint32, dataStorePreparePostParamV1 *DataStorePreparePostParamV1)
	CompletePostObjectV1Handler         func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1)
	GetNotificationUrlHandler           func(err error, client *nex.Client, callID uint32, param *DataStoreGetNotificationUrlParam)
	GetNewArrivedNotificationsV1Handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetNewArrivedNotificationsParam)
	DeleteObjectHandler                 func(err error, client *nex.Client, callID uint32, param *DataStoreDeleteParam)
	GetMetaHandler                      func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)
	GetMetasHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *DataStoreGetMetaParam)
	PrepareUpdateObjectHandler          func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *DataStorePrepareUpdateParam)
	CompleteUpdateObjectHandler         func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam)
	SearchObjectHandler                 func(err error, client *nex.Client, callID uint32, param *DataStoreSearchParam)
	RateObjectHandler                   func(err error, client *nex.Client, callID uint32, target *DataStoreRatingTarget, param *DataStoreRateObjectParam, fetchRatings bool)
	PostMetaBinaryHandler               func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)
	PreparePostObjectHandler            func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePreparePostParam)
	PrepareGetObjectHandler             func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePrepareGetParam)
	CompletePostObjectHandler           func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)
	GetNewArrivedNotificationsHandler   func(err error, client *nex.Client, callID uint32, param *DataStoreGetNewArrivedNotificationsParam)
	GetPersistenceInfoHandler           func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)
	GetMetasMultipleParamHandler        func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParams []*DataStoreGetMetaParam)
	CompletePostObjectsHandler          func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	ChangeMetaHandler                   func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)
	RateObjectsHandler                  func(err error, client *nex.Client, callID uint32, targets []*DataStoreRatingTarget, params []*DataStoreRateObjectParam, transactional bool, fetchRatings bool)
	SearchObjectLightHandler            func(err error, client *nex.Client, callID uint32, param *DataStoreSearchParam)
}

// Setup initializes the protocol
func (protocol *DataStoreProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *DataStoreProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPrepareGetObjectV1:
		go protocol.HandlePrepareGetObjectV1(packet)
	case MethodPreparePostObjectV1:
		go protocol.HandlePreparePostObjectV1(packet)
	case MethodCompletePostObjectV1:
		go protocol.HandleCompletePostObjectV1(packet)
	case MethodDeleteObject:
		go protocol.HandleDeleteObject(packet)
	case MethodGetMeta:
		go protocol.HandleGetMeta(packet)
	case MethodGetMetas:
		go protocol.HandleGetMetas(packet)
	case MethodPrepareUpdateObject:
		go protocol.HandlePrepareUpdateObject(packet)
	case MethodCompleteUpdateObject:
		go protocol.HandleCompleteUpdateObject(packet)
	case MethodSearchObject:
		go protocol.HandleSearchObject(packet)
	case MethodGetNotificationUrl:
		go protocol.HandleGetNotificationUrl(packet)
	case MethodGetNewArrivedNotificationsV1:
		go protocol.HandleGetNewArrivedNotificationsV1(packet)
	case MethodRateObject:
		go protocol.HandleRateObject(packet)
	case MethodPostMetaBinary:
		go protocol.HandlePostMetaBinary(packet)
	case MethodPreparePostObject:
		go protocol.HandlePreparePostObject(packet)
	case MethodPrepareGetObject:
		go protocol.HandlePrepareGetObject(packet)
	case MethodCompletePostObject:
		go protocol.HandleCompletePostObject(packet)
	case MethodGetNewArrivedNotifications:
		go protocol.HandleGetNewArrivedNotifications(packet)
	case MethodGetPersistenceInfo:
		go protocol.HandleGetPersistenceInfo(packet)
	case MethodGetMetasMultipleParam:
		go protocol.HandleGetMetasMultipleParam(packet)
	case MethodCompletePostObjects:
		go protocol.HandleCompletePostObjects(packet)
	case MethodChangeMeta:
		go protocol.HandleChangeMeta(packet)
	case MethodRateObjects:
		go protocol.HandleRateObjects(packet)
	case MethodSearchObjectLight:
		go protocol.HandleSearchObjectLight(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStore method ID: %#v\n", request.MethodID())
	}
}

// NewDataStoreProtocol returns a new DataStoreProtocol
func NewDataStoreProtocol(server *nex.Server) *DataStoreProtocol {
	protocol := &DataStoreProtocol{Server: server}

	protocol.Setup()

	return protocol
}
