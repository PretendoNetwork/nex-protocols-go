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
	prepareGetObjectV1Handler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParamV1) uint32
	preparePostObjectV1Handler          func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParamV1) uint32
	completePostObjectV1Handler         func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParamV1) uint32
	deleteObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreDeleteParam) uint32
	deleteObjectsHandler                func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreDeleteParam, transactional bool) uint32
	changeMetaV1Handler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParamV1) uint32
	changeMetasV1Handler                func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParamV1, transactional bool) uint32
	getMetaHandler                      func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetMetaParam) uint32
	getMetasHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64, param *datastore_types.DataStoreGetMetaParam) uint32
	prepareUpdateObjectHandler          func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareUpdateParam) uint32
	completeUpdateObjectHandler         func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) uint32
	searchObjectHandler                 func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam) uint32
	getNotificationURLHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNotificationURLParam) uint32
	getNewArrivedNotificationsV1Handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) uint32
	rateObjectHandler                   func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, param *datastore_types.DataStoreRateObjectParam, fetchRatings bool) uint32
	getRatingHandler                    func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64) uint32
	getRatingsHandler                   func(err error, client *nex.Client, callID uint32, dataIDs []uint64, accessPassword uint64) uint32
	resetRatingHandler                  func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64) uint32
	resetRatingsHandler                 func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, transactional bool) uint32
	getSpecificMetaV1Handler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1) uint32
	postMetaBinaryHandler               func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam) uint32
	touchObjectHandler                  func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreTouchObjectParam) uint32
	getRatingWithLogHandler             func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, accessPassword uint64) uint32
	preparePostObjectHandler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePreparePostParam) uint32
	prepareGetObjectHandler             func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam) uint32
	completePostObjectHandler           func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam) uint32
	getNewArrivedNotificationsHandler   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetNewArrivedNotificationsParam) uint32
	getSpecificMetaHandler              func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParam) uint32
	getPersistenceInfoHandler           func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16) uint32
	getPersistenceInfosHandler          func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotIDs []uint16) uint32
	perpetuateObjectHandler             func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, dataID uint64, deleteLastObject bool) uint32
	unperpetuateObjectHandler           func(err error, client *nex.Client, callID uint32, persistenceSlotID uint16, deleteLastObject bool) uint32
	prepareGetObjectOrMetaBinaryHandler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam) uint32
	getPasswordInfoHandler              func(err error, client *nex.Client, callID uint32, dataID uint64) uint32
	getPasswordInfosHandler             func(err error, client *nex.Client, callID uint32, dataIDs []uint64) uint32
	getMetasMultipleParamHandler        func(err error, client *nex.Client, callID uint32, params []*datastore_types.DataStoreGetMetaParam) uint32
	completePostObjectsHandler          func(err error, client *nex.Client, callID uint32, dataIDs []uint64) uint32
	changeMetaHandler                   func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreChangeMetaParam) uint32
	changeMetasHandler                  func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStoreChangeMetaParam, transactional bool) uint32
	rateObjectsHandler                  func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, params []*datastore_types.DataStoreRateObjectParam, transactional bool, fetchRatings bool) uint32
	postMetaBinaryWithDataIDHandler     func(err error, client *nex.Client, callID uint32, dataID uint64, param *datastore_types.DataStorePreparePostParam) uint32
	postMetaBinariesWithDataIDHandler   func(err error, client *nex.Client, callID uint32, dataIDs []uint64, params []*datastore_types.DataStorePreparePostParam, transactional bool) uint32
	rateObjectWithPostingHandler        func(err error, client *nex.Client, callID uint32, target *datastore_types.DataStoreRatingTarget, rateParam *datastore_types.DataStoreRateObjectParam, postParam *datastore_types.DataStorePreparePostParam, fetchRatings bool) uint32
	rateObjectsWithPostingHandler       func(err error, client *nex.Client, callID uint32, targets []*datastore_types.DataStoreRatingTarget, rateParams []*datastore_types.DataStoreRateObjectParam, postParams []*datastore_types.DataStorePreparePostParam, transactional bool, fetchRatings bool) uint32
	getObjectInfosHandler               func(err error, client *nex.Client, callID uint32, dataIDs uint64) uint32
	searchObjectLightHandler            func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam) uint32
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
		go protocol.handlePrepareGetObjectV1(packet)
	case MethodPreparePostObjectV1:
		go protocol.handlePreparePostObjectV1(packet)
	case MethodCompletePostObjectV1:
		go protocol.handleCompletePostObjectV1(packet)
	case MethodDeleteObject:
		go protocol.handleDeleteObject(packet)
	case MethodDeleteObjects:
		go protocol.handleDeleteObjects(packet)
	case MethodChangeMetaV1:
		go protocol.handleChangeMetaV1(packet)
	case MethodChangeMetasV1:
		go protocol.handleChangeMetasV1(packet)
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
	case MethodGetRating:
		go protocol.handleGetRating(packet)
	case MethodGetRatings:
		go protocol.handleGetRatings(packet)
	case MethodResetRating:
		go protocol.handleResetRating(packet)
	case MethodResetRatings:
		go protocol.handleResetRatings(packet)
	case MethodGetSpecificMetaV1:
		go protocol.handleGetSpecificMetaV1(packet)
	case MethodPostMetaBinary:
		go protocol.handlePostMetaBinary(packet)
	case MethodTouchObject:
		go protocol.handleTouchObject(packet)
	case MethodGetRatingWithLog:
		go protocol.handleGetRatingWithLog(packet)
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
	case MethodGetPersistenceInfos:
		go protocol.handleGetPersistenceInfos(packet)
	case MethodPerpetuateObject:
		go protocol.handlePerpetuateObject(packet)
	case MethodUnperpetuateObject:
		go protocol.handleUnperpetuateObject(packet)
	case MethodPrepareGetObjectOrMetaBinary:
		go protocol.handlePrepareGetObjectOrMetaBinary(packet)
	case MethodGetPasswordInfo:
		go protocol.handleGetPasswordInfo(packet)
	case MethodGetPasswordInfos:
		go protocol.handleGetPasswordInfos(packet)
	case MethodGetMetasMultipleParam:
		go protocol.handleGetMetasMultipleParam(packet)
	case MethodCompletePostObjects:
		go protocol.handleCompletePostObjects(packet)
	case MethodChangeMeta:
		go protocol.handleChangeMeta(packet)
	case MethodChangeMetas:
		go protocol.handleChangeMetas(packet)
	case MethodRateObjects:
		go protocol.handleRateObjects(packet)
	case MethodPostMetaBinaryWithDataID:
		go protocol.handlePostMetaBinaryWithDataID(packet)
	case MethodPostMetaBinariesWithDataID:
		go protocol.handlePostMetaBinariesWithDataID(packet)
	case MethodRateObjectWithPosting:
		go protocol.handleRateObjectWithPosting(packet)
	case MethodRateObjectsWithPosting:
		go protocol.handleRateObjectsWithPosting(packet)
	case MethodGetObjectInfos:
		go protocol.handleGetObjectInfos(packet)
	case MethodSearchObjectLight:
		go protocol.handleSearchObjectLight(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported DataStore method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new DataStore protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
