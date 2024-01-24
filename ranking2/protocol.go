// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

const (
	// ProtocolID is the protocol ID for the Ranking2 protocol
	ProtocolID = 0x7A

	// MethodPutScore is the method ID for the method PutScore
	MethodPutScore = 0x1

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0x2

	// MethodPutCommonData is the method ID for the method PutCommonData
	MethodPutCommonData = 0x3

	// MethodDeleteCommonData is the method ID for the method DeleteCommonData
	MethodDeleteCommonData = 0x4

	// MethodGetRanking is the method ID for the method GetRanking
	MethodGetRanking = 0x5

	// MethodGetRankingByPrincipalID is the method ID for the method GetRankingByPrincipalID
	MethodGetRankingByPrincipalID = 0x6

	// MethodGetCategorySetting is the method ID for the method GetCategorySetting
	MethodGetCategorySetting = 0x7

	// MethodGetRankingChart is the method ID for the method GetRankingChart
	MethodGetRankingChart = 0x8

	// MethodGetRankingCharts is the method ID for the method GetRankingCharts
	MethodGetRankingCharts = 0x9

	// MethodGetEstimateScoreRank is the method ID for the method GetEstimateScoreRank
	MethodGetEstimateScoreRank = 0xA
)

// Protocol stores all the RMC method handlers for the Ranking2 protocol and listens for requests
type Protocol struct {
	server                  nex.ServerInterface
	PutScore                func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetCommonData           func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	PutCommonData           func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	DeleteCommonData        func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetRanking              func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, uint32)
	GetRankingByPrincipalID func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	GetCategorySetting      func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	GetRankingChart         func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, uint32)
	GetRankingCharts        func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, uint32)
	GetEstimateScoreRank    func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Ranking2 protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerPutScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerPutCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, uint32))
	SetHandlerGetRankingByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerGetCategorySetting(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerGetRankingChart(handler func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, uint32))
	SetHandlerGetRankingCharts(handler func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, uint32))
	SetHandlerGetEstimateScoreRank(handler func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerPutScore sets the handler for the PutScore method
func (protocol *Protocol) SetHandlerPutScore(handler func(err error, packet nex.PacketInterface, callID uint32, scoreDataList *types.List[*ranking2_types.Ranking2ScoreData], nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.PutScore = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, optionFlags *types.PrimitiveU32, principalID *types.PID, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetCommonData = handler
}

// SetHandlerPutCommonData sets the handler for the PutCommonData method
func (protocol *Protocol) SetHandlerPutCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.PutCommonData = handler
}

// SetHandlerDeleteCommonData sets the handler for the DeleteCommonData method
func (protocol *Protocol) SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, nexUniqueID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.DeleteCommonData = handler
}

// SetHandlerGetRanking sets the handler for the GetRanking method
func (protocol *Protocol) SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam) (*nex.RMCMessage, uint32)) {
	protocol.GetRanking = handler
}

// SetHandlerGetRankingByPrincipalID sets the handler for the GetRankingByPrincipalID method
func (protocol *Protocol) SetHandlerGetRankingByPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, getParam *ranking2_types.Ranking2GetParam, principalIDList *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.GetRankingByPrincipalID = handler
}

// SetHandlerGetCategorySetting sets the handler for the GetCategorySetting method
func (protocol *Protocol) SetHandlerGetCategorySetting(handler func(err error, packet nex.PacketInterface, callID uint32, category *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.GetCategorySetting = handler
}

// SetHandlerGetRankingChart sets the handler for the GetRankingChart method
func (protocol *Protocol) SetHandlerGetRankingChart(handler func(err error, packet nex.PacketInterface, callID uint32, info *ranking2_types.Ranking2ChartInfoInput) (*nex.RMCMessage, uint32)) {
	protocol.GetRankingChart = handler
}

// SetHandlerGetRankingCharts sets the handler for the GetRankingCharts method
func (protocol *Protocol) SetHandlerGetRankingCharts(handler func(err error, packet nex.PacketInterface, callID uint32, infoArray *types.List[*ranking2_types.Ranking2ChartInfoInput]) (*nex.RMCMessage, uint32)) {
	protocol.GetRankingCharts = handler
}

// SetHandlerGetEstimateScoreRank sets the handler for the GetEstimateScoreRank method
func (protocol *Protocol) SetHandlerGetEstimateScoreRank(handler func(err error, packet nex.PacketInterface, callID uint32, input *ranking2_types.Ranking2EstimateScoreRankInput) (*nex.RMCMessage, uint32)) {
	protocol.GetEstimateScoreRank = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodPutScore:
		protocol.handlePutScore(packet)
	case MethodGetCommonData:
		protocol.handleGetCommonData(packet)
	case MethodPutCommonData:
		protocol.handlePutCommonData(packet)
	case MethodDeleteCommonData:
		protocol.handleDeleteCommonData(packet)
	case MethodGetRanking:
		protocol.handleGetRanking(packet)
	case MethodGetRankingByPrincipalID:
		protocol.handleGetRankingByPrincipalID(packet)
	case MethodGetCategorySetting:
		protocol.handleGetCategorySetting(packet)
	case MethodGetRankingChart:
		protocol.handleGetRankingChart(packet)
	case MethodGetRankingCharts:
		protocol.handleGetRankingCharts(packet)
	case MethodGetEstimateScoreRank:
		protocol.handleGetEstimateScoreRank(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported Ranking2 method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Ranking2 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
