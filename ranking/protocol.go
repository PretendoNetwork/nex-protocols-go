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
	Server                          *nex.Server
	uploadScoreHandler              func(err error, client *nex.Client, callID uint32, scoreData *ranking_types.RankingScoreData, uniqueID uint64)
	deleteScoreHandler              func(err error, client *nex.Client, callID uint32, category uint32, uniqueID uint64)
	deleteAllScoresHandler          func(err error, client *nex.Client, callID uint32, uniqueID uint64)
	uploadCommonDataHandler         func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64)
	deleteCommonDataHandler         func(err error, client *nex.Client, callID uint32, uniqueID uint64)
	getCommonDataHandler            func(err error, client *nex.Client, callID uint32, uniqueID uint64)
	changeAttributesHandler         func(err error, client *nex.Client, callID uint32, category uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64)
	changeAllAttributesHandler      func(err error, client *nex.Client, callID uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64)
	getRankingHandler               func(err error, client *nex.Client, callID uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64, principalID uint32)
	getApproxOrderHandler           func(err error, client *nex.Client, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, score uint32, uniqueID uint64, principalID uint32)
	getStatsHandler                 func(err error, client *nex.Client, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam, flags uint32)
	getRankingByPIDListHandler      func(err error, client *nex.Client, callID uint32, principalIDList []uint32, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64)
	getRankingByUniqueIDListHandler func(err error, client *nex.Client, callID uint32, nexUniqueIDList []uint64, rankingMode uint8, category uint32, orderParam *ranking_types.RankingOrderParam, uniqueID uint64)
	getCachedTopXRankingHandler     func(err error, client *nex.Client, callID uint32, category uint32, orderParam *ranking_types.RankingOrderParam)
	getCachedTopXRankingsHandler    func(err error, client *nex.Client, callID uint32, categories []uint32, orderParams []*ranking_types.RankingOrderParam)
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

	if request.ProtocolID() == ProtocolID {
		switch request.MethodID() {
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
			go globals.RespondNotImplemented(packet, ProtocolID)
			fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID())
		}
	}
}

// NewProtocol returns a new Ranking protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
