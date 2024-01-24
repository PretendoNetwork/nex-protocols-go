// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Friends protocol
	ProtocolID = 0x14

	// MethodAddFriend is the method ID for method AddFriend
	MethodAddFriend = 0x1

	// MethodAddFriendByName is the method ID for method AddFriendByName
	MethodAddFriendByName = 0x2

	// MethodAddFriendWithDetails is the method ID for method AddFriendWithDetails
	MethodAddFriendWithDetails = 0x3

	// MethodAddFriendByNameWithDetails is the method ID for method AddFriendByNameWithDetails
	MethodAddFriendByNameWithDetails = 0x4

	// MethodAcceptFriendship is the method ID for method AcceptFriendship
	MethodAcceptFriendship = 0x5

	// MethodDeclineFriendship is the method ID for method DeclineFriendship
	MethodDeclineFriendship = 0x6

	// MethodBlackList is the method ID for method BlackList
	MethodBlackList = 0x7

	// MethodBlackListByName is the method ID for method BlackListByName
	MethodBlackListByName = 0x8

	// MethodClearRelationship is the method ID for method ClearRelationship
	MethodClearRelationship = 0x9

	// MethodUpdateDetails is the method ID for method UpdateDetails
	MethodUpdateDetails = 0xA

	// MethodGetList is the method ID for method GetList
	MethodGetList = 0xB

	// MethodGetDetailedList is the method ID for method GetDetailedList
	MethodGetDetailedList = 0xC

	// MethodGetRelationships is the method ID for method GetRelationships
	MethodGetRelationships = 0xD
)

// Protocol handles the Friends QRV protocol
type Protocol struct {
	server                     nex.ServerInterface
	AddFriend                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)
	AddFriendByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)
	AddFriendWithDetails       func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)
	AddFriendByNameWithDetails func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32) // TODO - Is this the right signature?
	AcceptFriendship           func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	DeclineFriendship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	BlackList                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	BlackListByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	ClearRelationship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	UpdateDetails              func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	GetList                    func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetDetailedList            func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetRelationships           func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Friends protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32))
	SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32))
	SetHandlerAddFriendWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32))
	SetHandlerAddFriendByNameWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32))
	SetHandlerAcceptFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerDeclineFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerBlackListByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerClearRelationship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerGetList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetDetailedList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerAddFriend sets the handler for the AddFriend method
func (protocol *Protocol) SetHandlerAddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)) {
	protocol.AddFriend = handler
}

// SetHandlerAddFriendByName sets the handler for the AddFriendByName method
func (protocol *Protocol) SetHandlerAddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)) {
	protocol.AddFriendByName = handler
}

// SetHandlerAddFriendWithDetails sets the handler for the AddFriendWithDetails method
func (protocol *Protocol) SetHandlerAddFriendWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)) {
	protocol.AddFriendWithDetails = handler
}

// SetHandlerAddFriendByNameWithDetails sets the handler for the AddFriendByNameWithDetails method
func (protocol *Protocol) SetHandlerAddFriendByNameWithDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32, strMessage *types.String) (*nex.RMCMessage, uint32)) {
	protocol.AddFriendByNameWithDetails = handler
}

// SetHandlerAcceptFriendship sets the handler for the AcceptFriendship method
func (protocol *Protocol) SetHandlerAcceptFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.AcceptFriendship = handler
}

// SetHandlerDeclineFriendship sets the handler for the DeclineFriendship method
func (protocol *Protocol) SetHandlerDeclineFriendship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.DeclineFriendship = handler
}

// SetHandlerBlackList sets the handler for the BlackList method
func (protocol *Protocol) SetHandlerBlackList(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.BlackList = handler
}

// SetHandlerBlackListByName sets the handler for the BlackListByName method
func (protocol *Protocol) SetHandlerBlackListByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName *types.String, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.BlackListByName = handler
}

// SetHandlerClearRelationship sets the handler for the ClearRelationship method
func (protocol *Protocol) SetHandlerClearRelationship(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ClearRelationship = handler
}

// SetHandlerUpdateDetails sets the handler for the UpdateDetails method
func (protocol *Protocol) SetHandlerUpdateDetails(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer *types.PrimitiveU32, uiDetails *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.UpdateDetails = handler
}

// SetHandlerGetList sets the handler for the GetList method
func (protocol *Protocol) SetHandlerGetList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.GetList = handler
}

// SetHandlerGetDetailedList sets the handler for the GetDetailedList method
func (protocol *Protocol) SetHandlerGetDetailedList(handler func(err error, packet nex.PacketInterface, callID uint32, byRelationship *types.PrimitiveU8, bReversed *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.GetDetailedList = handler
}

// SetHandlerGetRelationships sets the handler for the GetRelationships method
func (protocol *Protocol) SetHandlerGetRelationships(handler func(err error, packet nex.PacketInterface, callID uint32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.GetRelationships = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodAddFriend:
		protocol.handleAddFriend(packet)
	case MethodAddFriendByName:
		protocol.handleAddFriendByName(packet)
	case MethodAddFriendWithDetails:
		protocol.handleAddFriendWithDetails(packet)
	case MethodAddFriendByNameWithDetails:
		protocol.handleAddFriendByNameWithDetails(packet)
	case MethodAcceptFriendship:
		protocol.handleAcceptFriendship(packet)
	case MethodDeclineFriendship:
		protocol.handleDeclineFriendship(packet)
	case MethodBlackList:
		protocol.handleBlackList(packet)
	case MethodBlackListByName:
		protocol.handleBlackListByName(packet)
	case MethodClearRelationship:
		protocol.handleClearRelationship(packet)
	case MethodUpdateDetails:
		protocol.handleUpdateDetails(packet)
	case MethodGetList:
		protocol.handleGetList(packet)
	case MethodGetDetailedList:
		protocol.handleGetDetailedList(packet)
	case MethodGetRelationships:
		protocol.handleGetRelationships(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported Friends method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Friends protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
