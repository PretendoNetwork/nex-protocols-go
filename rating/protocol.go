// Package protocol implements the Rating protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	rating_types "github.com/PretendoNetwork/nex-protocols-go/v2/rating/types"
)

const (
	// ProtocolID is the protocol ID for the Rating protocol
	ProtocolID = 0x76

	// MethodUnk1 is the method ID for the method Unk1
	// TODO - Find name if possible
	MethodUnk1 = 0x1

	// MethodUnk2 is the method ID for the method Unk2
	// TODO - Find name if possible
	MethodUnk2 = 0x2

	// MethodReportRatingStats is the method ID for the method ReportRatingStats
	MethodReportRatingStats = 0x3

	// MethodGetRanking is the method ID for the method GetRanking
	MethodGetRanking = 0x4

	// MethodDeleteScore is the method ID for the method DeleteScore
	MethodDeleteScore = 0x5

	// MethodUploadCommonData is the method ID for the method UploadCommonData
	MethodUploadCommonData = 0x7

	// MethodGetCommonData is the method ID for the method GetCommonData
	MethodGetCommonData = 0x8
)

// Protocol handles the Rating protocol
type Protocol struct {
	endpoint          nex.EndpointInterface
	Unk1              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	Unk2              func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken) (*nex.RMCMessage, *nex.Error) // TODO - Find name if possible
	ReportRatingStats func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken, stats types.List[rating_types.RatingStats]) (*nex.RMCMessage, *nex.Error)
	GetRanking        func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64, principalID types.PID) (*nex.RMCMessage, *nex.Error)
	DeleteScore       func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)
	UploadCommonData  func(err error, packet nex.PacketInterface, callID uint32, commonData types.Buffer, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)
	GetCommonData     func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)
	Patches           nex.ServiceProtocol
	PatchedMethods    []uint32
}

// Interface implements the methods present on the Rating protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken) (*nex.RMCMessage, *nex.Error)) // TODO - Find name if possible
	SetHandlerReportRatingStats(handler func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken, stats types.List[rating_types.RatingStats]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64, principalID types.PID) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error))
	SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData types.Buffer, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerUnk1 sets the handler for the Unk1 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk1(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk1 = handler
}

// SetHandlerUnk2 sets the handler for the Unk2 method
// TODO - Find name if possible
func (protocol *Protocol) SetHandlerUnk2(handler func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken) (*nex.RMCMessage, *nex.Error)) {
	protocol.Unk2 = handler
}

// SetHandlerReportRatingStats sets the handler for the ReportRatingStats method
func (protocol *Protocol) SetHandlerReportRatingStats(handler func(err error, packet nex.PacketInterface, callID uint32, sessionToken rating_types.RatingSessionToken, stats types.List[rating_types.RatingStats]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportRatingStats = handler
}

// SetHandlerGetRanking sets the handler for the GetRanking method
func (protocol *Protocol) SetHandlerGetRanking(handler func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64, principalID types.PID) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetRanking = handler
}

// SetHandlerDeleteScore sets the handler for the DeleteScore method
func (protocol *Protocol) SetHandlerDeleteScore(handler func(err error, packet nex.PacketInterface, callID uint32, category types.UInt32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteScore = handler
}

// SetHandlerUploadCommonData sets the handler for the UploadCommonData method
func (protocol *Protocol) SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, commonData types.Buffer, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)) {
	protocol.UploadCommonData = handler
}

// SetHandlerGetCommonData sets the handler for the GetCommonData method
func (protocol *Protocol) SetHandlerGetCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID types.UInt64) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCommonData = handler
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

	switch message.MethodID {
	case MethodUnk1:
		protocol.handleUnk1(packet) // TODO - Find name if possible
	case MethodUnk2:
		protocol.handleUnk2(packet) // TODO - Find name if possible
	case MethodReportRatingStats:
		protocol.handleReportRatingStats(packet)
	case MethodGetRanking:
		protocol.handleGetRanking(packet)
	case MethodDeleteScore:
		protocol.handleDeleteScore(packet)
	case MethodUploadCommonData:
		protocol.handleUploadCommonData(packet)
	case MethodGetCommonData:
		protocol.handleGetCommonData(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Rating method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Rating protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
