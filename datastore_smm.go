package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreSMMProtocolID is the protocol ID for the DataStore (SMM) protocol. ID is the same as the DataStore protocol
	DataStoreSMMProtocolID = 0x73

	// DataStoreSMMMethodGetObjectInfos is the method ID for the method GetObjectInfos
	DataStoreSMMMethodGetObjectInfos = 0x2D

	// DataStoreSMMMethodRateCustomRanking is the method ID for the method RateCustomRanking
	DataStoreSMMMethodRateCustomRanking = 0x30

	// DataStoreSMMMethodGetCustomRankingByDataId is the method ID for the method GetCustomRankingByDataId
	DataStoreSMMMethodGetCustomRankingByDataId = 0x32

	// DataStoreSMMMethodAddToBufferQueues is the method ID for the method AddToBufferQueues
	DataStoreSMMMethodAddToBufferQueues = 0x35

	// DataStoreSMMMethodGetBufferQueue is the method ID for the method GetBufferQueue
	DataStoreSMMMethodGetBufferQueue = 0x36

	// DataStoreSMMMethodCompleteAttachFile is the method ID for the method CompleteAttachFile
	DataStoreSMMMethodCompleteAttachFile = 0x39

	// DataStoreSMMMethodPrepareAttachFile is the method ID for the method PrepareAttachFile
	DataStoreSMMMethodPrepareAttachFile = 0x3B

	// DataStoreSMMMethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	DataStoreSMMMethodGetApplicationConfig = 0x3D

	// DataStoreSMMMethodFollowingsLatestCourseSearchObject is the method ID for the method FollowingsLatestCourseSearchObject
	DataStoreSMMMethodFollowingsLatestCourseSearchObject = 0x41

	// DataStoreSMMMethodRecommendedCourseSearchObject is the method ID for the method RecommendedCourseSearchObject
	DataStoreSMMMethodRecommendedCourseSearchObject = 0x42

	// DataStoreSMMMethodSuggestedCourseSearchObjectis the method ID for the method SuggestedCourseSearchObject
	DataStoreSMMMethodSuggestedCourseSearchObject = 0x44

	// DataStoreSMMMethodGetCourseRecord is the method ID for the method GetCourseRecord
	DataStoreSMMMethodGetCourseRecord = 0x48

	// DataStoreSMMMethodGetApplicationConfigString is the method ID for the method GetApplicationConfigString
	DataStoreSMMMethodGetApplicationConfigString = 0x4A

	// DataStoreSMMMethodGetMetasWithCourseRecord is the method ID for the method GetMetasWithCourseRecord
	DataStoreSMMMethodGetMetasWithCourseRecord = 0x4E

	// DataStoreSMMMethodCTRPickUpCourseSearchObject is the method ID for the method CTRPickUpCourseSearchObject
	DataStoreSMMMethodCTRPickUpCourseSearchObject = 0x52
)

// DataStoreSMMProtocol handles the DataStore (SMM) nex protocol. Embeds DataStoreProtocol
type DataStoreSMMProtocol struct {
	server *nex.Server
	DataStoreProtocol
	GetObjectInfosHandler                     func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	RateCustomRankingHandler                  func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)
	GetCustomRankingByDataIdHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)
	AddToBufferQueuesHandler                  func(err error, client *nex.Client, callID uint32, params []*BufferQueueParam, buffers [][]byte)
	GetBufferQueueHandler                     func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)
	CompleteAttachFileHandler                 func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)
	PrepareAttachFileHandler                  func(err error, client *nex.Client, callID uint32, dataStoreAttachFileParam *DataStoreAttachFileParam)
	GetApplicationConfigHandler               func(err error, client *nex.Client, callID uint32, applicationID uint32)
	FollowingsLatestCourseSearchObjectHandler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)
	RecommendedCourseSearchObjectHandler      func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)
	SuggestedCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, param *DataStoreSearchParam, extraData []string)
	GetCourseRecordHandler                    func(err error, client *nex.Client, callID uint32, param *DataStoreGetCourseRecordParam)
	GetApplicationConfigStringHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	GetMetasWithCourseRecordHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *DataStoreGetMetaParam)
	CTRPickUpCourseSearchObjectHandler        func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)
}

// DataStoreGetCourseRecordResult is used to send data about a courses world record
type DataStoreGetCourseRecordResult struct {
	nex.Structure
	DataID      uint64
	Slot        uint8
	FirstPID    uint32
	BestPID     uint32
	BestScore   int32
	CreatedTime *nex.DateTime
	UpdatedTime *nex.DateTime
}

// Bytes encodes the DataStoreGetCourseRecordResult and returns a byte array
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.DataID)
	stream.WriteUInt8(dataStoreGetCourseRecordResult.Slot)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.FirstPID)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.BestPID)
	stream.WriteUInt32LE(uint32(dataStoreGetCourseRecordResult.BestScore))
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.CreatedTime.Value())
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.UpdatedTime.Value())

	return stream.Bytes()
}

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	return &DataStoreGetCourseRecordResult{}
}

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	nex.Structure
	DataID  uint64
	GetInfo *DataStoreReqGetInfo
}

// Bytes encodes the DataStoreFileServerObjectInfo and returns a byte array
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreFileServerObjectInfo.DataID)
	stream.WriteStructure(dataStoreFileServerObjectInfo.GetInfo)

	return stream.Bytes()
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{}
}

// DataStoreAttachFileParam is sent in the PrepareAttachFile method
type DataStoreAttachFileParam struct {
	nex.Structure
	PostParam   *DataStorePreparePostParam
	ReferDataID uint64
	ContentType string
}

// ExtractFromStream extracts a DataStoreAttachFileParam structure from a stream
func (dataStoreAttachFileParam *DataStoreAttachFileParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	postParam, err := stream.ReadStructure(NewDataStorePreparePostParam())
	if err != nil {
		return err
	}

	dataStoreAttachFileParam.PostParam = postParam.(*DataStorePreparePostParam)
	dataStoreAttachFileParam.ReferDataID = stream.ReadUInt64LE()

	contentType, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreAttachFileParam.ContentType = contentType

	return nil
}

// NewDataStoreAttachFileParam returns a new DataStoreAttachFileParam
func NewDataStoreAttachFileParam() *DataStoreAttachFileParam {
	return &DataStoreAttachFileParam{}
}

// DataStoreGetCourseRecordParam is sent in the GetMetasWithCourseRecord method
type DataStoreGetCourseRecordParam struct {
	nex.Structure
	DataID uint64
	Slot   uint8
}

// ExtractFromStream extracts a DataStoreGetCourseRecordParam structure from a stream
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreGetCourseRecordParam.DataID = stream.ReadUInt64LE()
	dataStoreGetCourseRecordParam.Slot = stream.ReadUInt8()

	return nil
}

// NewDataStoreGetCourseRecordParamreturns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	return &DataStoreGetCourseRecordParam{}
}

// BufferQueueParam is sent in the GetBufferQueue method
type BufferQueueParam struct {
	nex.Structure
	DataID uint64
	Slot   uint32
}

// ExtractFromStream extracts a BufferQueueParam structure from a stream
func (bufferQueueParam *BufferQueueParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	bufferQueueParam.DataID = stream.ReadUInt64LE()
	bufferQueueParam.Slot = stream.ReadUInt32LE()

	return nil
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() *BufferQueueParam {
	return &BufferQueueParam{}
}

// DataStoreRateCustomRankingParam is sent in the RateCustomRanking method
type DataStoreRateCustomRankingParam struct {
	nex.Structure
	DataID        uint64
	ApplicationId uint32
	Score         uint32
	Period        uint16
}

// ExtractFromStream extracts a DataStoreRateCustomRankingParam structure from a stream
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreRateCustomRankingParam.DataID = stream.ReadUInt64LE()
	dataStoreRateCustomRankingParam.ApplicationId = stream.ReadUInt32LE()
	dataStoreRateCustomRankingParam.Score = stream.ReadUInt32LE()
	dataStoreRateCustomRankingParam.Period = stream.ReadUInt16LE()

	return nil
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	return &DataStoreRateCustomRankingParam{}
}

// DataStoreGetCustomRankingByDataIdParam is sent in the GetCustomRankingByDataId method
type DataStoreGetCustomRankingByDataIdParam struct {
	nex.Structure
	ApplicationId uint32
	DataIdList    []uint64
	ResultOption  uint8
}

// ExtractFromStream extracts a DataStoreGetCustomRankingByDataIdParam structure from a stream
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreGetCustomRankingByDataIdParam.ApplicationId = stream.ReadUInt32LE()
	dataStoreGetCustomRankingByDataIdParam.DataIdList = stream.ReadListUInt64LE()
	dataStoreGetCustomRankingByDataIdParam.ResultOption = stream.ReadUInt8()

	return nil
}

// NewDataStoreGetCustomRankingByDataIdParam returns a new DataStoreGetCustomRankingByDataIdParam
func NewDataStoreGetCustomRankingByDataIdParam() *DataStoreGetCustomRankingByDataIdParam {
	return &DataStoreGetCustomRankingByDataIdParam{}
}

// DataStoreCustomRankingResult is sent in the FollowingsLatestCourseSearchObject method
type DataStoreCustomRankingResult struct {
	Order    uint32
	Score    uint32
	MetaInfo *DataStoreMetaInfo

	nex.Structure
}

// ExtractFromStream extracts a DataStoreCustomRankingResult structure from a stream
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreCustomRankingResult.Order = stream.ReadUInt32LE()
	dataStoreCustomRankingResult.Score = stream.ReadUInt32LE()

	metaInfo, err := stream.ReadStructure(NewDataStoreMetaInfo())
	if err != nil {
		return err
	}

	dataStoreCustomRankingResult.MetaInfo = metaInfo.(*DataStoreMetaInfo)

	return nil
}

// Bytes encodes the DataStoreCustomRankingResult and returns a byte array
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Order)
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Score)
	stream.WriteStructure(dataStoreCustomRankingResult.MetaInfo)

	return stream.Bytes()
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{}
}

// Setup initializes the protocol
func (dataStoreSMMProtocol *DataStoreSMMProtocol) Setup() {
	nexServer := dataStoreSMMProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreSMMProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreSMMProtocol.handleGetMeta(packet)
			case DataStoreMethodGetMetasMultipleParam:
				go dataStoreSMMProtocol.handleGetMetasMultipleParam(packet)
			case DataStoreMethodPreparePostObject:
				go dataStoreSMMProtocol.handlePreparePostObject(packet)
			case DataStoreMethodPrepareGetObject:
				go dataStoreSMMProtocol.handlePrepareGetObject(packet)
			case DataStoreMethodCompletePostObject:
				go dataStoreSMMProtocol.handleCompletePostObject(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreSMMProtocol.handleChangeMeta(packet)
			case DataStoreMethodRateObjects:
				go dataStoreSMMProtocol.handleRateObjects(packet)
			case DataStoreSMMMethodGetObjectInfos:
				go dataStoreSMMProtocol.handleGetObjectInfos(packet)
			case DataStoreSMMMethodRateCustomRanking:
				go dataStoreSMMProtocol.handleRateCustomRanking(packet)
			case DataStoreSMMMethodGetCustomRankingByDataId:
				go dataStoreSMMProtocol.handleGetCustomRankingByDataId(packet)
			case DataStoreSMMMethodAddToBufferQueues:
				go dataStoreSMMProtocol.handleAddToBufferQueues(packet)
			case DataStoreSMMMethodGetBufferQueue:
				go dataStoreSMMProtocol.handleGetBufferQueue(packet)
			case DataStoreSMMMethodCompleteAttachFile:
				go dataStoreSMMProtocol.handleCompleteAttachFile(packet)
			case DataStoreSMMMethodPrepareAttachFile:
				go dataStoreSMMProtocol.handlePrepareAttachFile(packet)
			case DataStoreSMMMethodGetApplicationConfig:
				go dataStoreSMMProtocol.handleGetApplicationConfig(packet)
			case DataStoreSMMMethodFollowingsLatestCourseSearchObject:
				go dataStoreSMMProtocol.handleFollowingsLatestCourseSearchObject(packet)
			case DataStoreSMMMethodRecommendedCourseSearchObject:
				go dataStoreSMMProtocol.handleRecommendedCourseSearchObject(packet)
			case DataStoreSMMMethodSuggestedCourseSearchObject:
				go dataStoreSMMProtocol.handleSuggestedCourseSearchObject(packet)
			case DataStoreSMMMethodGetCourseRecord:
				go dataStoreSMMProtocol.handleGetCourseRecord(packet)
			case DataStoreSMMMethodGetApplicationConfigString:
				go dataStoreSMMProtocol.handleGetApplicationConfigString(packet)
			case DataStoreSMMMethodGetMetasWithCourseRecord:
				go dataStoreSMMProtocol.handleGetMetasWithCourseRecord(packet)
			case DataStoreSMMMethodCTRPickUpCourseSearchObject:
				go dataStoreSMMProtocol.handleCTRPickUpCourseSearchObject(packet)
			default:
				fmt.Printf("Unsupported DataStoreSMM method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetObjectInfos sets the GetObjectInfos handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetObjectInfos(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64)) {
	dataStoreSMMProtocol.GetObjectInfosHandler = handler
}

// RateCustomRanking sets the RateCustomRanking handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) RateCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)) {
	dataStoreSMMProtocol.RateCustomRankingHandler = handler
}

// GetCustomRankingByDataId sets the GetCustomRankingByDataId handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetCustomRankingByDataId(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)) {
	dataStoreSMMProtocol.GetCustomRankingByDataIdHandler = handler
}

// AddToBufferQueues sets the AddToBufferQueues handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) AddToBufferQueues(handler func(err error, client *nex.Client, callID uint32, params []*BufferQueueParam, buffers [][]byte)) {
	dataStoreSMMProtocol.AddToBufferQueuesHandler = handler
}

// GetBufferQueue sets the GetBufferQueue handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetBufferQueue(handler func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)) {
	dataStoreSMMProtocol.GetBufferQueueHandler = handler
}

// CompleteAttachFile sets the CompleteAttachFile handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) CompleteAttachFile(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)) {
	dataStoreSMMProtocol.CompleteAttachFileHandler = handler
}

// PrepareAttachFile sets the PrepareAttachFile handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) PrepareAttachFile(handler func(err error, client *nex.Client, callID uint32, dataStoreAttachFileParam *DataStoreAttachFileParam)) {
	dataStoreSMMProtocol.PrepareAttachFileHandler = handler
}

// GetApplicationConfig sets the GetApplicationConfig handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	dataStoreSMMProtocol.GetApplicationConfigHandler = handler
}

// FollowingsLatestCourseSearchObject sets the FollowingsLatestCourseSearchObject handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) FollowingsLatestCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)) {
	dataStoreSMMProtocol.FollowingsLatestCourseSearchObjectHandler = handler
}

// RecommendedCourseSearchObject sets the RecommendedCourseSearchObject handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) RecommendedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)) {
	dataStoreSMMProtocol.RecommendedCourseSearchObjectHandler = handler
}

// SuggestedCourseSearchObject sets the SuggestedCourseSearchObject handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) SuggestedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchParam, extraData []string)) {
	dataStoreSMMProtocol.SuggestedCourseSearchObjectHandler = handler
}

// GetCourseRecord sets the GetCourseRecord handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetCourseRecord(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetCourseRecordParam)) {
	dataStoreSMMProtocol.GetCourseRecordHandler = handler
}

// GetApplicationConfigString sets the GetApplicationConfigString handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetApplicationConfigString(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	dataStoreSMMProtocol.GetApplicationConfigStringHandler = handler
}

// GetMetasWithCourseRecord sets the GetMetasWithCourseRecord handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetMetasWithCourseRecord(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *DataStoreGetMetaParam)) {
	dataStoreSMMProtocol.GetMetasWithCourseRecordHandler = handler
}

// CTRPickUpCourseSearchObject sets the CTRPickUpCourseSearchObject handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) CTRPickUpCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)) {
	dataStoreSMMProtocol.CTRPickUpCourseSearchObjectHandler = handler
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetObjectInfos(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetObjectInfosHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetObjectInfos not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataIDs := parametersStream.ReadListUInt64LE()

	go dataStoreSMMProtocol.GetObjectInfosHandler(nil, client, callID, dataIDs)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleRateCustomRanking(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.RateCustomRankingHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::RateCustomRanking not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreRateCustomRankingParams, err := parametersStream.ReadListDataStoreRateCustomRankingParam()

	if err != nil {
		go dataStoreSMMProtocol.RateCustomRankingHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.RateCustomRankingHandler(nil, client, callID, dataStoreRateCustomRankingParams)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetCustomRankingByDataId(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetCustomRankingByDataIdHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetCustomRankingByDataId not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreGetCustomRankingByDataIdParam, err := parametersStream.ReadStructure(NewDataStoreGetCustomRankingByDataIdParam())

	if err != nil {
		go dataStoreSMMProtocol.GetCustomRankingByDataIdHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.GetCustomRankingByDataIdHandler(nil, client, callID, dataStoreGetCustomRankingByDataIdParam.(*DataStoreGetCustomRankingByDataIdParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.AddToBufferQueuesHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::AddToBufferQueues not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := NewStreamIn(parameters, dataStoreSMMProtocol.server)

	params, err := parametersStream.ReadListBufferQueueParam()

	if err != nil {
		go dataStoreSMMProtocol.AddToBufferQueuesHandler(err, client, callID, nil, nil)
		return
	}

	buffers := parametersStream.ReadListQBuffer()

	go dataStoreSMMProtocol.AddToBufferQueuesHandler(nil, client, callID, params, buffers)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetBufferQueue(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetBufferQueueHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetBufferQueue not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	bufferQueueParam, err := parametersStream.ReadStructure(NewBufferQueueParam())

	if err != nil {
		go dataStoreSMMProtocol.GetBufferQueueHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.GetBufferQueueHandler(nil, client, callID, bufferQueueParam.(*BufferQueueParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleCompleteAttachFile(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.CompleteAttachFileHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::CompleteAttachFile not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreCompletePostParam, err := parametersStream.ReadStructure(NewDataStoreCompletePostParam())
	if err != nil {
		go dataStoreSMMProtocol.CompleteAttachFileHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.CompleteAttachFileHandler(nil, client, callID, dataStoreCompletePostParam.(*DataStoreCompletePostParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handlePrepareAttachFile(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.PrepareAttachFileHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::PrepareAttachFile not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreAttachFileParam, err := parametersStream.ReadStructure(NewDataStoreAttachFileParam())

	if err != nil {
		go dataStoreSMMProtocol.PrepareAttachFileHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.PrepareAttachFileHandler(nil, client, callID, dataStoreAttachFileParam.(*DataStoreAttachFileParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetApplicationConfig(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetApplicationConfigHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetApplicationConfig not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	applicationID := parametersStream.ReadUInt32LE()

	go dataStoreSMMProtocol.GetApplicationConfigHandler(nil, client, callID, applicationID)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleFollowingsLatestCourseSearchObject(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.FollowingsLatestCourseSearchObjectHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::FollowingsLatestCourseSearchObject not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreSearchParam, err := parametersStream.ReadStructure(NewDataStoreSearchParam())

	if err != nil {
		go dataStoreSMMProtocol.FollowingsLatestCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go dataStoreSMMProtocol.FollowingsLatestCourseSearchObjectHandler(nil, client, callID, dataStoreSearchParam.(*DataStoreSearchParam), extraData)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleRecommendedCourseSearchObject(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.RecommendedCourseSearchObjectHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::RecommendedCourseSearchObject not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreSearchParam, err := parametersStream.ReadStructure(NewDataStoreSearchParam())

	if err != nil {
		go dataStoreSMMProtocol.RecommendedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go dataStoreSMMProtocol.RecommendedCourseSearchObjectHandler(nil, client, callID, dataStoreSearchParam.(*DataStoreSearchParam), extraData)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleSuggestedCourseSearchObject(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.SuggestedCourseSearchObjectHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::SuggestedCourseSearchObject not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreSearchParam())

	if err != nil {
		go dataStoreSMMProtocol.SuggestedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go dataStoreSMMProtocol.SuggestedCourseSearchObjectHandler(nil, client, callID, param.(*DataStoreSearchParam), extraData)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetCourseRecord(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetCourseRecordHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetCourseRecord not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetCourseRecordParam())

	if err != nil {
		go dataStoreSMMProtocol.GetCourseRecordHandler(err, client, callID, nil)
		return
	}

	go dataStoreSMMProtocol.GetCourseRecordHandler(nil, client, callID, param.(*DataStoreGetCourseRecordParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetApplicationConfigString(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetApplicationConfigStringHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetApplicationConfigString not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	applicationID := parametersStream.ReadUInt32LE()

	go dataStoreSMMProtocol.GetApplicationConfigStringHandler(nil, client, callID, applicationID)
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleGetMetasWithCourseRecord(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.GetMetasWithCourseRecordHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::GetMetasWithCourseRecord not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreGetCourseRecordParams, err := parametersStream.ReadListDataStoreGetCourseRecordParam()
	if err != nil {
		go dataStoreSMMProtocol.GetMetasWithCourseRecordHandler(err, client, callID, nil, nil)
		return
	}

	dataStoreGetMetaParam, err := parametersStream.ReadStructure(NewDataStoreGetMetaParam())
	if err != nil {
		go dataStoreSMMProtocol.GetMetasWithCourseRecordHandler(err, client, callID, nil, nil)
		return
	}

	go dataStoreSMMProtocol.GetMetasWithCourseRecordHandler(nil, client, callID, dataStoreGetCourseRecordParams, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
}

func (dataStoreSMMProtocol *DataStoreSMMProtocol) handleCTRPickUpCourseSearchObject(packet nex.PacketInterface) {
	if dataStoreSMMProtocol.CTRPickUpCourseSearchObjectHandler == nil {
		fmt.Println("[Warning] DataStoreSMMProtocol::CTRPickUpCourseSearchObject not implemented")
		go respondNotImplemented(packet, DataStoreSMMProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSMMProtocol.server)

	dataStoreSearchParam, err := parametersStream.ReadStructure(NewDataStoreSearchParam())

	if err != nil {
		go dataStoreSMMProtocol.CTRPickUpCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData := parametersStream.ReadListString()

	go dataStoreSMMProtocol.CTRPickUpCourseSearchObjectHandler(nil, client, callID, dataStoreSearchParam.(*DataStoreSearchParam), extraData)
}

// NewDataStoreSMMProtocol returns a new DataStoreSMMProtocol
func NewDataStoreSMMProtocol(server *nex.Server) *DataStoreSMMProtocol {
	dataStoreSMMProtocol := &DataStoreSMMProtocol{server: server}
	dataStoreSMMProtocol.DataStoreProtocol.server = server

	dataStoreSMMProtocol.Setup()

	return dataStoreSMMProtocol
}
