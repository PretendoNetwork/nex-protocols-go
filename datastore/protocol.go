// Package protocol implements the DataStore protocol
package protocol

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

// Protocol stores all the RMC method handlers for the DataStore protocol and listens for requests
type Protocol struct {
	Server                              *nex.Server
	prepareGetObjectV1Handler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParamV1)
	preparePostObjectV1Handler          func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParamV1)
	completePostObjectV1Handler         func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParamV1)
	deleteObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreDeleteParam)
	deleteObjectsHandler                func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreDeleteParam, transactional bool)
	changeMetaV1Handler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1)
	changeMetasV1Handler                func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParamV1, transactional bool)
	getMetaHandler                      func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetMetaParam)
	getMetasHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *datastore_types.DataStoreGetMetaParam)
	prepareUpdateObjectHandler          func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareUpdateParam)
	completeUpdateObjectHandler         func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam)
	searchObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)
	getNotificationURLHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam)
	getNewArrivedNotificationsV1Handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)
	rateObjectHandler                   func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings bool)
	getRatingHandler                    func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64)
	getRatingsHandler                   func(err error, client *nex.Client, callID uint32, dataIDs []uint64, accessPassword uint64)
	resetRatingHandler                  func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64)
	resetRatingsHandler                 func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional bool)
	getSpecificMetaV1Handler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1)
	postMetaBinaryHandler               func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam)
	touchObjectHandler                  func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreTouchObjectParam)
	getRatingWithLogHandler             func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64)
	preparePostObjectHandler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam)
	prepareGetObjectHandler             func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam)
	completePostObjectHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam)
	getNewArrivedNotificationsHandler   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam)
	getSpecificMetaHandler              func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam)
	getPersistenceInfoHandler           func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)
	getPersistenceInfosHandler          func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotIDs []uint16)
	perpetuateObjectHandler             func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, dataID uint64, deleteLastObject bool)
	unperpetuateObjectHandler           func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, deleteLastObject bool)
	prepareGetObjectOrMetaBinaryHandler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam)
	getPasswordInfoHandler              func(err error, client *nex.Client, callID uint32, dataID uint64)
	getPasswordInfosHandler             func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	getMetasMultipleParamHandler        func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreGetMetaParam)
	completePostObjectsHandler          func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	changeMetaHandler                   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParam)
	changeMetasHandler                  func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParam, transactional bool)
	rateObjectsHandler                  func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, params []*datastore_types.DataStoreRateObjectParam, transactional bool, fetchRatings bool)
	postMetaBinaryWithDataIDHandler     func(err error, client *nex.Client, callID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam)
	postMetaBinariesWithDataIDHandler   func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStorePreparePostParam, transactional bool)
	rateObjectWithPostingHandler        func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings bool)
	rateObjectsWithPostingHandler       func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, rateParams []*datastore_types.DataStoreRateObjectParam, postParams []*datastore_types.DataStorePreparePostParam, transactional bool, fetchRatings bool)
	getObjectInfosHandler               func(err error, client *nex.Client, callID uint32, dataIDs uint64)
	searchObjectLightHandler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodPrepareGetObjectV1:
		protocol.handlePrepareGetObjectV1(packet)
	case MethodPreparePostObjectV1:
		protocol.handlePreparePostObjectV1(packet)
	case MethodCompletePostObjectV1:
		protocol.handleCompletePostObjectV1(packet)
	case MethodDeleteObject:
		protocol.handleDeleteObject(packet)
	case MethodDeleteObjects:
		protocol.handleDeleteObjects(packet)
	case MethodChangeMetaV1:
		protocol.handleChangeMetaV1(packet)
	case MethodChangeMetasV1:
		protocol.handleChangeMetasV1(packet)
	case MethodGetMeta:
		protocol.handleGetMeta(packet)
	case MethodGetMetas:
		protocol.handleGetMetas(packet)
	case MethodPrepareUpdateObject:
		protocol.handlePrepareUpdateObject(packet)
	case MethodCompleteUpdateObject:
		protocol.handleCompleteUpdateObject(packet)
	case MethodSearchObject:
		protocol.handleSearchObject(packet)
	case MethodGetNotificationURL:
		protocol.handleGetNotificationURL(packet)
	case MethodGetNewArrivedNotificationsV1:
		protocol.handleGetNewArrivedNotificationsV1(packet)
	case MethodRateObject:
		protocol.handleRateObject(packet)
	case MethodGetRating:
		protocol.handleGetRating(packet)
	case MethodGetRatings:
		protocol.handleGetRatings(packet)
	case MethodResetRating:
		protocol.handleResetRating(packet)
	case MethodResetRatings:
		protocol.handleResetRatings(packet)
	case MethodGetSpecificMetaV1:
		protocol.handleGetSpecificMetaV1(packet)
	case MethodPostMetaBinary:
		protocol.handlePostMetaBinary(packet)
	case MethodTouchObject:
		protocol.handleTouchObject(packet)
	case MethodGetRatingWithLog:
		protocol.handleGetRatingWithLog(packet)
	case MethodPreparePostObject:
		protocol.handlePreparePostObject(packet)
	case MethodPrepareGetObject:
		protocol.handlePrepareGetObject(packet)
	case MethodCompletePostObject:
		protocol.handleCompletePostObject(packet)
	case MethodGetNewArrivedNotifications:
		protocol.handleGetNewArrivedNotifications(packet)
	case MethodGetSpecificMeta:
		protocol.handleGetSpecificMeta(packet)
	case MethodGetPersistenceInfo:
		protocol.handleGetPersistenceInfo(packet)
	case MethodGetPersistenceInfos:
		protocol.handleGetPersistenceInfos(packet)
	case MethodPerpetuateObject:
		protocol.handlePerpetuateObject(packet)
	case MethodUnperpetuateObject:
		protocol.handleUnperpetuateObject(packet)
	case MethodPrepareGetObjectOrMetaBinary:
		protocol.handlePrepareGetObjectOrMetaBinary(packet)
	case MethodGetPasswordInfo:
		protocol.handleGetPasswordInfo(packet)
	case MethodGetPasswordInfos:
		protocol.handleGetPasswordInfos(packet)
	case MethodGetMetasMultipleParam:
		protocol.handleGetMetasMultipleParam(packet)
	case MethodCompletePostObjects:
		protocol.handleCompletePostObjects(packet)
	case MethodChangeMeta:
		protocol.handleChangeMeta(packet)
	case MethodChangeMetas:
		protocol.handleChangeMetas(packet)
	case MethodRateObjects:
		protocol.handleRateObjects(packet)
	case MethodPostMetaBinaryWithDataID:
		protocol.handlePostMetaBinaryWithDataID(packet)
	case MethodPostMetaBinariesWithDataID:
		protocol.handlePostMetaBinariesWithDataID(packet)
	case MethodRateObjectWithPosting:
		protocol.handleRateObjectWithPosting(packet)
	case MethodRateObjectsWithPosting:
		protocol.handleRateObjectsWithPosting(packet)
	case MethodGetObjectInfos:
		protocol.handleGetObjectInfos(packet)
	case MethodSearchObjectLight:
		protocol.handleSearchObjectLight(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported DataStore method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStore protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
