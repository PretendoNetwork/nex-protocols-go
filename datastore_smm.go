package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreSMMProtocolID is the protocol ID for the DataStore (SMM) protocol. ID is the same as the DataStore protocol
	DataStoreSMMProtocolID = 0x73

	// DataStoreSMMMethodRateCustomRanking is the method ID for the method RateCustomRanking
	DataStoreSMMMethodRateCustomRanking = 0x30

	// DataStoreSMMMethodGetCustomRankingByDataId is the method ID for the method GetCustomRankingByDataId
	DataStoreSMMMethodGetCustomRankingByDataId = 0x32

	// DataStoreSMMMethodGetBufferQueue is the method ID for the method GetBufferQueue
	DataStoreSMMMethodGetBufferQueue = 0x36

	// DataStoreSMMMethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	DataStoreSMMMethodGetApplicationConfig = 0x3D

	// DataStoreSMMMethodFollowingsLatestCourseSearchObject is the method ID for the method FollowingsLatestCourseSearchObject
	DataStoreSMMMethodFollowingsLatestCourseSearchObject = 0x41

	// DataStoreSMMMethodRecommendedCourseSearchObject is the method ID for the method RecommendedCourseSearchObject
	DataStoreSMMMethodRecommendedCourseSearchObject = 0x42

	// DataStoreSMMMethodGetApplicationConfigString is the method ID for the method GetApplicationConfigString
	DataStoreSMMMethodGetApplicationConfigString = 0x4A

	// DataStoreSMMMethodGetMetasWithCourseRecord is the method ID for the method GetMetasWithCourseRecord
	DataStoreSMMMethodGetMetasWithCourseRecord = 0x4E
)

// DataStoreSMMProtocol handles the DataStore (SMM) nex protocol. Embeds DataStoreProtocol
type DataStoreSMMProtocol struct {
	server *nex.Server
	DataStoreProtocol
	RateCustomRankingHandler                  func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)
	GetCustomRankingByDataIdHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)
	GetBufferQueueHandler                     func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)
	GetApplicationConfigHandler               func(err error, client *nex.Client, callID uint32, applicationID uint32)
	FollowingsLatestCourseSearchObjectHandler func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)
	RecommendedCourseSearchObjectHandler      func(err error, client *nex.Client, callID uint32, dataStoreSearchParam *DataStoreSearchParam, extraData []string)
	GetApplicationConfigStringHandler         func(err error, client *nex.Client, callID uint32, applicationID uint32)
	GetMetasWithCourseRecordHandler           func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *DataStoreGetMetaParam)
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
func (dataStoreGetCustomRankingByDataIdParam *DataStoreRateCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreGetCustomRankingByDataIdParam.DataID = stream.ReadUInt64LE()
	dataStoreGetCustomRankingByDataIdParam.ApplicationId = stream.ReadUInt32LE()
	dataStoreGetCustomRankingByDataIdParam.Score = stream.ReadUInt32LE()
	dataStoreGetCustomRankingByDataIdParam.Period = stream.ReadUInt16LE()

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
			case DataStoreMethodPrepareGetObject:
				go dataStoreSMMProtocol.handlePrepareGetObject(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreSMMProtocol.handleChangeMeta(packet)
			case DataStoreSMMMethodRateCustomRanking:
				go dataStoreSMMProtocol.handleRateCustomRanking(packet)
			case DataStoreSMMMethodGetCustomRankingByDataId:
				go dataStoreSMMProtocol.handleGetCustomRankingByDataId(packet)
			case DataStoreSMMMethodGetBufferQueue:
				go dataStoreSMMProtocol.handleGetBufferQueue(packet)
			case DataStoreSMMMethodGetApplicationConfig:
				go dataStoreSMMProtocol.handleGetApplicationConfig(packet)
			case DataStoreSMMMethodFollowingsLatestCourseSearchObject:
				go dataStoreSMMProtocol.handleFollowingsLatestCourseSearchObject(packet)
			case DataStoreSMMMethodRecommendedCourseSearchObject:
				go dataStoreSMMProtocol.handleRecommendedCourseSearchObject(packet)
			case DataStoreSMMMethodGetApplicationConfigString:
				go dataStoreSMMProtocol.handleGetApplicationConfigString(packet)
			case DataStoreSMMMethodGetMetasWithCourseRecord:
				go dataStoreSMMProtocol.handleGetMetasWithCourseRecord(packet)
			default:
				fmt.Printf("Unsupported DataStoreSMM method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// RateCustomRanking sets the RateCustomRanking handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) RateCustomRanking(handler func(err error, client *nex.Client, callID uint32, dataStoreRateCustomRankingParams []*DataStoreRateCustomRankingParam)) {
	dataStoreSMMProtocol.RateCustomRankingHandler = handler
}

// GetCustomRankingByDataId sets the GetCustomRankingByDataId handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetCustomRankingByDataId(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam)) {
	dataStoreSMMProtocol.GetCustomRankingByDataIdHandler = handler
}

// GetBufferQueue sets the GetBufferQueue handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetBufferQueue(handler func(err error, client *nex.Client, callID uint32, bufferQueueParam *BufferQueueParam)) {
	dataStoreSMMProtocol.GetBufferQueueHandler = handler
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

// GetApplicationConfigString sets the GetApplicationConfigString handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetApplicationConfigString(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	dataStoreSMMProtocol.GetApplicationConfigStringHandler = handler
}

// GetMetasWithCourseRecord sets the GetMetasWithCourseRecord handler function
func (dataStoreSMMProtocol *DataStoreSMMProtocol) GetMetasWithCourseRecord(handler func(err error, client *nex.Client, callID uint32, dataStoreGetCourseRecordParams []*DataStoreGetCourseRecordParam, dataStoreGetMetaParam *DataStoreGetMetaParam)) {
	dataStoreSMMProtocol.GetMetasWithCourseRecordHandler = handler
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

// NewDataStoreSMMProtocol returns a new DataStoreSMMProtocol
func NewDataStoreSMMProtocol(server *nex.Server) *DataStoreSMMProtocol {
	dataStoreSMMProtocol := &DataStoreSMMProtocol{server: server}
	dataStoreSMMProtocol.DataStoreProtocol.server = server

	dataStoreSMMProtocol.Setup()

	return dataStoreSMMProtocol
}
