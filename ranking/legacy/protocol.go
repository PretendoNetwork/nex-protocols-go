// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/legacy/types"
)

const (
	// ProtocolID is the protocol ID for the legacy Ranking protocol
	ProtocolID = 0x70

	// MethodUploadScore is the method ID for the method UploadScore
	MethodUploadScore = 0x1

	// MethodUploadScores is the method ID for the method UploadScores
	MethodUploadScores = 0x2

	// MethodDeleteScore is the method ID for the method DeleteScore
	MethodDeleteScore = 0x3

	// MethodDeleteAllScore is the method ID for the method DeleteAllScore
	MethodDeleteAllScore = 0x4

	// MethodUploadCommonData is the method ID for the method UploadCommonData
	MethodUploadCommonData = 0x5

	// MethodDeleteCommonData is the method ID for the method DeleteCommonData
	MethodDeleteCommonData = 0x6

	// MethodUnk0x7 is the method ID for the method Unk0x7
	// TODO - Find name if possible
	MethodUnk0x7 = 0x7

	// MethodUnk0x8 is the method ID for the method Unk0x8
	// TODO - Find name if possible
	MethodUnk0x8 = 0x8

	// MethodUnk0x9 is the method ID for the method Unk0x9
	// TODO - Find name if possible
	MethodUnk0x9 = 0x9

	// MethodGetTopScore is the method ID for the the method GetTopScore
	MethodGetTopScore = 0xA

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0xB

	// MethodUnk0xC is the method ID for the method Unk0xC
	// TODO - Find name if possible
	MethodUnk0xC = 0xC

	// MethodUnk0xD is the method ID for the method Unk0xD
	// TODO - Find name if possible
	MethodUnk0xD = 0xD

	// MethodGetScore is the method ID for the method GetScore
	MethodGetScore = 0xE

	// MethodGetSelfScore is the method ID for the method GetSelfScore
	MethodGetSelfScore = 0xF

	// MethodGetTotal is the method ID for method GetTotal
	MethodGetTotal = 0x10

	// MethodUploadScoreWithLimit is the method ID for method UploadScoreWithLimit
	MethodUploadScoreWithLimit = 0x11

	// MethodUploadScoresWithLimit is the method ID for method UploadScoresWithLimit
	MethodUploadScoresWithLimit = 0x12

	// MethodUnk0x13 is the method ID for the method Unk0x13
	// TODO - Find name if possible
	MethodUnk0x13 = 0x13

	// * NEX 1 method IDs

	// MethodDeleteScoreNEX1 is the method ID for the method DeleteScore on NEX 1
	MethodDeleteScoreNEX1 = 0x2

	// MethodDeleteAllScoreNEX1 is the method ID for the method DeleteAllScore on NEX 1
	MethodDeleteAllScoreNEX1 = 0x3

	// MethodUploadCommonDataNEX1 is the method ID for the method UploadCommonData on NEX 1
	MethodUploadCommonDataNEX1 = 0x4

	// MethodDeleteCommonDataNEX1 is the method ID for the method DeleteCommonData on NEX 1
	MethodDeleteCommonDataNEX1 = 0x5

	// MethodUnk0x7NEX1 is the method ID for the method Unk0x7 on NEX 1
	// TODO - Find name if possible
	MethodUnk0x7NEX1 = 0x6

	// MethodUnk0x8NEX1 is the method ID for the method Unk0x8 on NEX 1
	// TODO - Find name if possible
	MethodUnk0x8NEX1 = 0x7

	// MethodGetTopScoreNEX1 is the method ID for the the method GetTopScore on NEX 1
	MethodGetTopScoreNEX1 = 0x8

	// MethodGetCommonDataNEX1 is the method ID for the method GetCommonData on NEX 1
	MethodGetCommonDataNEX1 = 0x9

	// MethodUnk0xCNEX1 is the method ID for the method Unk0xC on NEX 1
	// TODO - Find name if possible
	MethodUnk0xCNEX1 = 0xA

	// MethodUnk0xDNEX1 is the method ID for the method Unk0xD on NEX 1
	// TODO - Find name if possible
	MethodUnk0xDNEX1 = 0xB

	// MethodGetScoreNEX1 is the method ID for the method GetScore on NEX 1
	MethodGetScoreNEX1 = 0xC

	// MethodGetSelfScoreNEX1 is the method ID for the method GetSelfScore on NEX 1
	MethodGetSelfScoreNEX1 = 0xD

	// MethodGetTotalNEX1 is the method ID for method GetTotal on NEX 1
	MethodGetTotalNEX1 = 0xE
)

// Protocol stores all the RMC method handlers for the Ranking protocol and listens for requests
type Protocol struct {
	endpoint              nex.EndpointInterface
	UploadScore           func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32) (*nex.RMCMessage, *nex.Error)
	UploadScores          func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScore]) (*nex.RMCMessage, *nex.Error)
	DeleteScore           func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)
	DeleteAllScore        func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)
	UploadCommonData      func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, commonData types.Buffer) (*nex.RMCMessage, *nex.Error)
	DeleteCommonData      func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)
	Unk0x7                func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	Unk0x8                func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	Unk0x9                func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8]) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	GetTopScore           func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetCommonData         func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)
	Unk0xC                func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	Unk0xD                func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	GetScore              func(err error, packet nex.PacketInterface, callID uint32, rankingMode types.UInt8, category types.UInt32, orderParam ranking_types.RankingOrderParam, offset types.UInt32, length types.UInt8) (*nex.RMCMessage, *nex.Error)
	GetSelfScore          func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam, length types.UInt8) (*nex.RMCMessage, *nex.Error)
	GetTotal              func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown1 types.UInt8, unknown2 types.UInt8, unknown3 types.UInt8, unknown4 types.UInt32) (*nex.RMCMessage, *nex.Error)
	UploadScoreWithLimit  func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32, limit types.UInt16) (*nex.RMCMessage, *nex.Error)
	UploadScoresWithLimit func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScoreWithLimit]) (*nex.RMCMessage, *nex.Error)
	Unk0x13               func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8], unknown5 types.Bool, unknown6 types.UInt16) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	Patches               nex.ServiceProtocol
	PatchedMethods        []uint32
}

// Interface implements the methods present on the legacy Ranking protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScore]) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteAllScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, commonData types.Buffer) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnk0x7(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerUnk0x8(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerUnk0x9(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8]) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerGetTopScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnk0xC(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerUnk0xD(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerGetScore(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode types.UInt8, category types.UInt32, orderParam ranking_types.RankingOrderParam, offset types.UInt32, length types.UInt8) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSelfScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam, length types.UInt8) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetTotal(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown1 types.UInt8, unknown2 types.UInt8, unknown3 types.UInt8, unknown4 types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadScoreWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32, limit types.UInt16) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadScoresWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScoreWithLimit]) (*nex.RMCMessage, *nex.Error))
	SetHandlerUnk0x13(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8], unknown5 types.Bool, unknown6 types.UInt16) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerUploadScore sets the handler for the UploadScore method
func (protocol *Protocol) SetHandlerUploadScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadScore = handler
}

// SetHandlerUploadScores sets the handler for the UploadScores method
func (protocol *Protocol) SetHandlerUploadScores(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScore]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadScores = handler
}

// SetHandlerDeleteScore sets the handler for the DeleteScore method
func (protocol *Protocol) SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteScore = handler
}

// SetHandlerDeleteAllScore sets the handler for the DeleteAllScore method
func (protocol *Protocol) SetHandlerDeleteAllScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteAllScore = handler
}

// SetHandlerUploadCommonData sets the handler for the UploadCommonData method
func (protocol *Protocol) SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, commonData types.Buffer) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadCommonData = handler
}

// SetHandlerDeleteCommonData sets the handler for the DeleteCommonData method
func (protocol *Protocol) SetHandlerDeleteCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteCommonData = handler
}

// SetHandlerUnk0x7 sets the handler for the Unk0x7 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0x7(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0x7 = handler
}

// SetHandlerUnk0x8 sets the handler for the Unk0x8 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0x8(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0x8 = handler
}

// SetHandlerUnk0x9 sets the handler for the Unk0x9 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0x9(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8]) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0x9 = handler
}

// SetHandlerGetTopScore sets the handler for the GetTopScore method
func (protocol *Protocol) SetHandlerGetTopScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTopScore = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCommonData = handler
}

// SetHandlerUnk0xC sets the handler for the Unk0xC method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0xC(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0xC = handler
}

// SetHandlerUnk0xD sets the handler for the Unk0xD method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0xD(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0xD = handler
}

// SetHandlerGetScore sets the handler for the GetScore method
func (protocol *Protocol) SetHandlerGetScore(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode types.UInt8, category types.UInt32, orderParam ranking_types.RankingOrderParam, offset types.UInt32, length types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetScore = handler
}

// SetHandlerGetSelfScore sets the handler for the GetSelfScore method
func (protocol *Protocol) SetHandlerGetSelfScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, orderParam ranking_types.RankingOrderParam, length types.UInt8) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSelfScore = handler
}

// SetHandlerGetTotal sets the handler for the GetTotal method
func (protocol *Protocol) SetHandlerGetTotal(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, unknown1 types.UInt8, unknown2 types.UInt8, unknown3 types.UInt8, unknown4 types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetTotal = handler
}

// SetHandlerUploadScoreWithLimit sets the handler for the UploadScoreWithLimit method
func (protocol *Protocol) SetHandlerUploadScoreWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, category types.UInt32, scores types.List[types.UInt32], unknown1 types.UInt8, unknown2 types.UInt32, limit types.UInt16) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadScoreWithLimit = handler
}

// SetHandlerUploadScoresWithLimit sets the handler for the UploadScoresWithLimit method
func (protocol *Protocol) SetHandlerUploadScoresWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt32, scores types.List[ranking_types.RankingScoreWithLimit]) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadScoresWithLimit = handler
}

// SetHandlerUnk0x13 sets the handler for the Unk0x13 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk0x13(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 types.UInt32, unknown2 types.UInt32, unknown3 types.List[types.UInt32], unknown4 types.List[types.UInt8], unknown5 types.Bool, unknown6 types.UInt16) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk0x13 = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	endpoint := packet.Sender().Endpoint()
	rankingVersion := endpoint.LibraryVersions().Ranking

	// * NEX 1 doesn't implement UploadScores, Unk0x9 and any method after UploadScoresWithLimit
	if rankingVersion.GreaterOrEqual("2.0.0") {
		switch message.MethodID {
		case MethodUploadScore:
			protocol.handleUploadScore(packet)
		case MethodUploadScores:
			protocol.handleUploadScores(packet)
		case MethodDeleteScore:
			protocol.handleDeleteScore(packet)
		case MethodDeleteAllScore:
			protocol.handleDeleteAllScore(packet)
		case MethodUploadCommonData:
			protocol.handleUploadCommonData(packet)
		case MethodDeleteCommonData:
			protocol.handleDeleteCommonData(packet)
		case MethodUnk0x7:
			protocol.handleUnk0x7(packet) // TODO - Find name if possible
		case MethodUnk0x8:
			protocol.handleUnk0x8(packet) // TODO - Find name if possible
		case MethodUnk0x9:
			protocol.handleUnk0x9(packet) // TODO - Find name if possible
		case MethodGetTopScore:
			protocol.handleGetTopScore(packet)
		case MethodGetCommonData:
			protocol.handleGetCommonData(packet)
		case MethodUnk0xC:
			protocol.handleUnk0xC(packet) // TODO - Find name if possible
		case MethodUnk0xD:
			protocol.handleUnk0xD(packet) // TODO - Find name if possible
		case MethodGetScore:
			protocol.handleGetScore(packet)
		case MethodGetSelfScore:
			protocol.handleGetSelfScore(packet)
		case MethodGetTotal:
			protocol.handleGetTotal(packet)
		case MethodUploadScoreWithLimit:
			protocol.handleUploadScoreWithLimit(packet)
		case MethodUploadScoresWithLimit:
			protocol.handleUploadScoresWithLimit(packet)
		case MethodUnk0x13:
			protocol.handleUnk0x13(packet) // TODO - Find name if possible
		default:
			errMessage := fmt.Sprintf("Unsupported Ranking method ID: %#v\n", message.MethodID)
			err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

			globals.RespondError(packet, ProtocolID, err)
			globals.Logger.Warning(err.Message)
		}
	} else {
		switch message.MethodID {
		case MethodUploadScore:
			protocol.handleUploadScore(packet)
		case MethodDeleteScoreNEX1:
			protocol.handleDeleteScore(packet)
		case MethodDeleteAllScoreNEX1:
			protocol.handleDeleteAllScore(packet)
		case MethodUploadCommonDataNEX1:
			protocol.handleUploadCommonData(packet)
		case MethodDeleteCommonDataNEX1:
			protocol.handleDeleteCommonData(packet)
		case MethodUnk0x7NEX1:
			protocol.handleUnk0x7(packet) // TODO - Find name if possible
		case MethodUnk0x8NEX1:
			protocol.handleUnk0x8(packet) // TODO - Find name if possible
		case MethodGetTopScoreNEX1:
			protocol.handleGetTopScore(packet)
		case MethodGetCommonDataNEX1:
			protocol.handleGetCommonData(packet)
		case MethodUnk0xCNEX1:
			protocol.handleUnk0xC(packet) // TODO - Find name if possible
		case MethodUnk0xDNEX1:
			protocol.handleUnk0xD(packet) // TODO - Find name if possible
		case MethodGetScoreNEX1:
			protocol.handleGetScore(packet)
		case MethodGetSelfScoreNEX1:
			protocol.handleGetSelfScore(packet)
		case MethodGetTotalNEX1:
			protocol.handleGetTotal(packet)
		default:
			errMessage := fmt.Sprintf("Unsupported Ranking method ID: %#v\n", message.MethodID)
			err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

			globals.RespondError(packet, ProtocolID, err)
			globals.Logger.Warning(err.Message)
		}
	}
}

// NewProtocol returns a new Ranking protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
