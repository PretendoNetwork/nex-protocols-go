// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
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

	// MethodGetNotificationURL is the method ID for the method GetNotificationURL
	MethodGetNotificationURL = 0xD

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

	// MethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataID
	MethodPostMetaBinaryWithDataID = 0x29

	// MethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataID
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

// DataStoreProtocol handles the DataStore NEX protocol
type DataStoreProtocol struct {
	Server                              *nex.Server
	PrepareGetObjectV1Handler           func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParamV1 *datastore_types.DataStorePrepareGetParamV1)
	PreparePostObjectV1Handler          func(err error, client *nex.Client, callID uint32, dataStorePreparePostParamV1 *datastore_types.DataStorePreparePostParamV1)
	CompletePostObjectV1Handler         func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParamV1 *datastore_types.DataStoreCompletePostParamV1)
	GetNotificationURLHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam)
	GetNewArrivedNotificationsV1Handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)
	DeleteObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreDeleteParam)
	GetMetaHandler                      func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *datastore_types.DataStoreGetMetaParam)
	GetMetasHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *datastore_types.DataStoreGetMetaParam)
	PrepareUpdateObjectHandler          func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *datastore_types.DataStorePrepareUpdateParam)
	CompleteUpdateObjectHandler         func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *datastore_types.DataStoreCompleteUpdateParam)
	SearchObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)
	RateObjectHandler                   func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings bool)
	GetSpecificMetaV1Handler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1)
	PostMetaBinaryHandler               func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *datastore_types.DataStorePreparePostParam)
	PreparePostObjectHandler            func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *datastore_types.DataStorePreparePostParam)
	PrepareGetObjectHandler             func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *datastore_types.DataStorePrepareGetParam)
	CompletePostObjectHandler           func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *datastore_types.DataStoreCompletePostParam)
	GetNewArrivedNotificationsHandler   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)
	GetSpecificMetaHandler              func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam)
	GetPersistenceInfoHandler           func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)
	GetMetasMultipleParamHandler        func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParams []*datastore_types.DataStoreGetMetaParam)
	CompletePostObjectsHandler          func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	ChangeMetaHandler                   func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *datastore_types.DataStoreChangeMetaParam)
	RateObjectsHandler                  func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, params []*datastore_types.DataStoreRateObjectParam, transactional bool, fetchRatings bool)
	SearchObjectLightHandler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)
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

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *DataStoreProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPrepareGetObjectV1:
		go protocol.handlePrepareGetObjectV1(packet)
	case MethodPreparePostObjectV1:
		go protocol.handlePreparePostObjectV1(packet)
	case MethodCompletePostObjectV1:
		go protocol.handleCompletePostObjectV1(packet)
	case MethodDeleteObject:
		go protocol.handleDeleteObject(packet)
	case MethodGetMeta:
		go protocol.handleGetMeta(packet)
	case MethodGetMetas:
		go protocol.handleGetMetas(packet)
	case MethodPrepareUpdateObject:
		go protocol.handlePrepareUpdateObject(packet)
	case MethodCompleteUpdateObject:
		go protocol.handleCompleteUpdateObject(packet)
	case MethodSearchObject:
		go protocol.handleSearchObject(packet)
	case MethodGetNotificationURL:
		go protocol.handleGetNotificationURL(packet)
	case MethodGetNewArrivedNotificationsV1:
		go protocol.handleGetNewArrivedNotificationsV1(packet)
	case MethodRateObject:
		go protocol.handleRateObject(packet)
	case MethodGetSpecificMetaV1:
		go protocol.handleGetSpecificMetaV1(packet)
	case MethodPostMetaBinary:
		go protocol.handlePostMetaBinary(packet)
	case MethodPreparePostObject:
		go protocol.handlePreparePostObject(packet)
	case MethodPrepareGetObject:
		go protocol.handlePrepareGetObject(packet)
	case MethodCompletePostObject:
		go protocol.handleCompletePostObject(packet)
	case MethodGetNewArrivedNotifications:
		go protocol.handleGetNewArrivedNotifications(packet)
	case MethodGetSpecificMeta:
		go protocol.handleGetSpecificMeta(packet)
	case MethodGetPersistenceInfo:
		go protocol.handleGetPersistenceInfo(packet)
	case MethodGetMetasMultipleParam:
		go protocol.handleGetMetasMultipleParam(packet)
	case MethodCompletePostObjects:
		go protocol.handleCompletePostObjects(packet)
	case MethodChangeMeta:
		go protocol.handleChangeMeta(packet)
	case MethodRateObjects:
		go protocol.handleRateObjects(packet)
	case MethodSearchObjectLight:
		go protocol.handleSearchObjectLight(packet)
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
