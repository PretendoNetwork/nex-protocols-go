// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
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
	Server                   nex.ServerInterface
	UploadScore              func(err error, packet nex.PacketInterface, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64) (*nex.RMCMessage, uint32)
	DeleteScore              func(err error, packet nex.PacketInterface, callID uint32, category uint32, uniqueID uint64) (*nex.RMCMessage, uint32)
	DeleteAllScores          func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint64) (*nex.RMCMessage, uint32)
	UploadCommonData         func(err error, packet nex.PacketInterface, callID uint32, commonData []byte, uniqueID uint64) (*nex.RMCMessage, uint32)
	DeleteCommonData         func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint64) (*nex.RMCMessage, uint32)
	GetCommonData            func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint64) (*nex.RMCMessage, uint32)
	ChangeAttributes         func(err error, packet nex.PacketInterface, callID uint32, category uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64) (*nex.RMCMessage, uint32)
	ChangeAllAttributes      func(err error, packet nex.PacketInterface, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64) (*nex.RMCMessage, uint32)
	GetRanking               func(err error, packet nex.PacketInterface, callID uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64, principalID *nex.PID) (*nex.RMCMessage, uint32)
	GetApproxOrder           func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, score uint32, uniqueID uint64, principalID *nex.PID) (*nex.RMCMessage, uint32)
	GetStats                 func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, flags uint32) (*nex.RMCMessage, uint32)
	GetRankingByPIDList      func(err error, packet nex.PacketInterface, callID uint32, principalIDList []*nex.PID, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64) (*nex.RMCMessage, uint32)
	GetRankingByUniqueIDList func(err error, packet nex.PacketInterface, callID uint32, nexUniqueIDList []uint64, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64) (*nex.RMCMessage, uint32)
	GetCachedTopXRanking     func(err error, packet nex.PacketInterface, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)
	GetCachedTopXRankings    func(err error, packet nex.PacketInterface, callID uint32, categories []uint32, orderParams []*ranking_types.RankingOrderParam) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
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
			go protocol.handleUploadScore(packet)
		case MethodDeleteScore:
			go protocol.handleDeleteScore(packet)
		case MethodDeleteAllScores:
			go protocol.handleDeleteAllScores(packet)
		case MethodUploadCommonData:
			go protocol.handleUploadCommonData(packet)
		case MethodDeleteCommonData:
			go protocol.handleDeleteCommonData(packet)
		case MethodGetCommonData:
			go protocol.handleGetCommonData(packet)
		case MethodChangeAttributes:
			go protocol.handleChangeAttributes(packet)
		case MethodChangeAllAttributes:
			go protocol.handleChangeAllAttributes(packet)
		case MethodGetRanking:
			go protocol.handleGetRanking(packet)
		case MethodGetApproxOrder:
			go protocol.handleGetApproxOrder(packet)
		case MethodGetStats:
			go protocol.handleGetStats(packet)
		case MethodGetRankingByPIDList:
			go protocol.handleGetRankingByPIDList(packet)
		case MethodGetRankingByUniqueIDList:
			go protocol.handleGetRankingByUniqueIDList(packet)
		case MethodGetCachedTopXRanking:
			go protocol.handleGetCachedTopXRanking(packet)
		case MethodGetCachedTopXRankings:
			go protocol.handleGetCachedTopXRankings(packet)
		default:
			go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
			fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID)
		}
	}
}

// NewProtocol returns a new Ranking protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
