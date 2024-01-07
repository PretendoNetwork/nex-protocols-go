// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking protocol
	ProtocolID = 0x70

	// MethodUploadScore is the method ID for method UploadScore
	MethodUploadScore = 0x1

	// MethodDeleteScore is the method ID for method DeleteScore
	MethodDeleteScore = 0x2

	// MethodDeleteAllScores is the method ID for method DeleteAllScores
	MethodDeleteAllScores = 0x3

	// MethodUploadCommonData is the method ID for method UploadCommonData
	MethodUploadCommonData = 0x4

	// MethodDeleteCommonData is the method ID for method DeleteCommonData
	MethodDeleteCommonData = 0x5

	// MethodGetCommonData is the method ID for method GetCommonData
	MethodGetCommonData = 0x6

	// MethodChangeAttributes is the method ID for method ChangeAttributes
	MethodChangeAttributes = 0x7

	// MethodChangeAllAttributes is the method ID for method ChangeAllAttributes
	MethodChangeAllAttributes = 0x8

	// MethodGetRanking is the method ID for method GetRanking
	MethodGetRanking = 0x9

	// MethodGetApproxOrder is the method ID for method GetApproxOrder
	MethodGetApproxOrder = 0xA

	// MethodGetStats is the method ID for method GetStats
	MethodGetStats = 0xB

	// MethodGetRankingByPIDList is the method ID for method GetRankingByPIDList
	MethodGetRankingByPIDList = 0xC

	// MethodGetRankingByUniqueIDList is the method ID for method GetRankingByUniqueIDList
	MethodGetRankingByUniqueIDList = 0xD

	// MethodGetCachedTopXRanking is the method ID for method GetCachedTopXRanking
	MethodGetCachedTopXRanking = 0xE

	// MethodGetCachedTopXRankings is the method ID for method GetCachedTopXRankings
	MethodGetCachedTopXRankings = 0xF
)

// Protocol stores all the RMC method handlers for the Ranking protocol and listens for requests
type Protocol struct {
	server                   nex.ServerInterface
	UploadScore              func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	DeleteScore              func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	DeleteAllScores          func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	UploadCommonData         func(err error, packet nex.PacketInterface, callID uint32, commonData []byte, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	DeleteCommonData         func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetCommonData            func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	ChangeAttributes         func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	ChangeAllAttributes      func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetRanking               func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32)
	GetApproxOrder           func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32)
	GetStats                 func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	GetRankingByPIDList      func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetRankingByUniqueIDList func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetCachedTopXRanking     func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)
	GetCachedTopXRankings    func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams []*ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Ranking protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerDeleteAllScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData []byte, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerChangeAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerChangeAllAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetApproxOrder(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerGetRankingByPIDList(handler func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetRankingByUniqueIDList(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetCachedTopXRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32))
	SetHandlerGetCachedTopXRankings(handler func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams []*ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerUploadScore sets the handler for the UploadScore method
func (protocol *Protocol) SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.UploadScore = handler
}

// SetHandlerDeleteScore sets the handler for the DeleteScore method
func (protocol *Protocol) SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.DeleteScore = handler
}

// SetHandlerDeleteAllScores sets the handler for the DeleteAllScores method
func (protocol *Protocol) SetHandlerDeleteAllScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.DeleteAllScores = handler
}

// SetHandlerUploadCommonData sets the handler for the UploadCommonData method
func (protocol *Protocol) SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData []byte, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.UploadCommonData = handler
}

// SetHandlerDeleteCommonData sets the handler for the DeleteCommonData method
func (protocol *Protocol) SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.DeleteCommonData = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetCommonData = handler
}

// SetHandlerChangeAttributes sets the handler for the ChangeAttributes method
func (protocol *Protocol) SetHandlerChangeAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.ChangeAttributes = handler
}

// SetHandlerChangeAllAttributes sets the handler for the ChangeAllAttributes method
func (protocol *Protocol) SetHandlerChangeAllAttributes(handler func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.ChangeAllAttributes = handler
}

// SetHandlerGetRanking sets the handler for the GetRanking method
func (protocol *Protocol) SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetRanking = handler
}

// SetHandlerGetApproxOrder sets the handler for the GetApproxOrder method
func (protocol *Protocol) SetHandlerGetApproxOrder(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, score *types.PrimitiveU32, uniqueID *types.PrimitiveU64, principalID *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetApproxOrder = handler
}

// SetHandlerGetStats sets the handler for the GetStats method
func (protocol *Protocol) SetHandlerGetStats(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, flags *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.GetStats = handler
}

// SetHandlerGetRankingByPIDList sets the handler for the GetRankingByPIDList method
func (protocol *Protocol) SetHandlerGetRankingByPIDList(handler func(err error, packet nex.PacketInterface, callID uint32, principalIDList *types.List[*types.PID], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetRankingByPIDList = handler
}

// SetHandlerGetRankingByUniqueIDList sets the handler for the GetRankingByUniqueIDList method
func (protocol *Protocol) SetHandlerGetRankingByUniqueIDList(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList *types.List[*types.PrimitiveU64], rankingMode *types.PrimitiveU8, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam, uniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetRankingByUniqueIDList = handler
}

// SetHandlerGetCachedTopXRanking sets the handler for the GetCachedTopXRanking method
func (protocol *Protocol) SetHandlerGetCachedTopXRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)) {
	protocol.GetCachedTopXRanking = handler
}

// SetHandlerGetCachedTopXRankings sets the handler for the GetCachedTopXRankings method
func (protocol *Protocol) SetHandlerGetCachedTopXRankings(handler func(err error, packet nex.PacketInterface, callID uint32, categories *types.List[*types.PrimitiveU32], orderParams []*ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)) {
	protocol.GetCachedTopXRankings = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	if request.ProtocolID == ProtocolID {
		switch request.MethodID {
		case MethodUploadScore:
			protocol.handleUploadScore(packet)
		case MethodDeleteScore:
			protocol.handleDeleteScore(packet)
		case MethodDeleteAllScores:
			protocol.handleDeleteAllScores(packet)
		case MethodUploadCommonData:
			protocol.handleUploadCommonData(packet)
		case MethodDeleteCommonData:
			protocol.handleDeleteCommonData(packet)
		case MethodGetCommonData:
			protocol.handleGetCommonData(packet)
		case MethodChangeAttributes:
			protocol.handleChangeAttributes(packet)
		case MethodChangeAllAttributes:
			protocol.handleChangeAllAttributes(packet)
		case MethodGetRanking:
			protocol.handleGetRanking(packet)
		case MethodGetApproxOrder:
			protocol.handleGetApproxOrder(packet)
		case MethodGetStats:
			protocol.handleGetStats(packet)
		case MethodGetRankingByPIDList:
			protocol.handleGetRankingByPIDList(packet)
		case MethodGetRankingByUniqueIDList:
			protocol.handleGetRankingByUniqueIDList(packet)
		case MethodGetCachedTopXRanking:
			protocol.handleGetCachedTopXRanking(packet)
		case MethodGetCachedTopXRankings:
			protocol.handleGetCachedTopXRankings(packet)
		default:
			globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
			fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID)
		}
	}
}

// NewProtocol returns a new Ranking protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
