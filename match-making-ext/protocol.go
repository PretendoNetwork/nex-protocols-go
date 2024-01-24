// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the MatchMakingExt protocol
	ProtocolID = 0x32

	// MethodEndParticipation is the method ID for the method EndParticipation
	MethodEndParticipation = 0x1

	// MethodGetParticipants is the method ID for the method GetParticipants
	MethodGetParticipants = 0x2

	// MethodGetDetailedParticipants is the method ID for the method GetDetailedParticipants
	MethodGetDetailedParticipants = 0x3

	// MethodGetParticipantsURLs is the method ID for the method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x4

	// MethodGetGatheringRelations is the method ID for the method GetGatheringRelations
	MethodGetGatheringRelations = 0x5

	// MethodDeleteFromDeletions is the method ID for the method DeleteFromDeletions
	MethodDeleteFromDeletions = 0x6
)

// Protocol handles the MatchMakingExt protocol
type Protocol struct {
	server                  nex.ServerInterface
	EndParticipation        func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)
	GetParticipants         func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetDetailedParticipants func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetParticipantsURLs     func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	GetGatheringRelations   func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, uint32)
	DeleteFromDeletions     func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Match Making Ext protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerEndParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32))
	SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerGetGatheringRelations(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, uint32))
	SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerEndParticipation sets the handler for the EndParticipation method
func (protocol *Protocol) SetHandlerEndParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)) {
	protocol.EndParticipation = handler
}

// SetHandlerGetParticipants sets the handler for the GetParticipants method
func (protocol *Protocol) SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.GetParticipants = handler
}

// SetHandlerGetDetailedParticipants sets the handler for the GetDetailedParticipants method
func (protocol *Protocol) SetHandlerGetDetailedParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering *types.PrimitiveU32, bOnlyActive *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.GetDetailedParticipants = handler
}

// SetHandlerGetParticipantsURLs sets the handler for the GetParticipantsURLs method
func (protocol *Protocol) SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.GetParticipantsURLs = handler
}

// SetHandlerGetGatheringRelations sets the handler for the GetGatheringRelations method
func (protocol *Protocol) SetHandlerGetGatheringRelations(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, descr *types.String) (*nex.RMCMessage, uint32)) {
	protocol.GetGatheringRelations = handler
}

// SetHandlerDeleteFromDeletions sets the handler for the DeleteFromDeletions method
func (protocol *Protocol) SetHandlerDeleteFromDeletions(handler func(err error, packet nex.PacketInterface, callID uint32, lstDeletions *types.List[*types.PrimitiveU32], pid *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.DeleteFromDeletions = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodEndParticipation:
		protocol.handleEndParticipation(packet)
	case MethodGetParticipants:
		protocol.handleGetParticipants(packet)
	case MethodGetDetailedParticipants:
		protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		protocol.handleGetParticipantsURLs(packet)
	case MethodGetGatheringRelations:
		protocol.handleGetGatheringRelations(packet)
	case MethodDeleteFromDeletions:
		protocol.handleDeleteFromDeletions(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported MatchMakingExt method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Match Making Ext protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
