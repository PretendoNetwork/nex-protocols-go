// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
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
	Server                     nex.ServerInterface
	AddFriend                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string) (*nex.RMCMessage, uint32)
	AddFriendByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName string, uiDetails uint32, strMessage string) (*nex.RMCMessage, uint32)
	AddFriendWithDetails       func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string) (*nex.RMCMessage, uint32)
	AddFriendByNameWithDetails func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string) (*nex.RMCMessage, uint32) // TODO - Is this the right signature?
	AcceptFriendship           func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32) (*nex.RMCMessage, uint32)
	DeclineFriendship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32) (*nex.RMCMessage, uint32)
	BlackList                  func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32) (*nex.RMCMessage, uint32)
	BlackListByName            func(err error, packet nex.PacketInterface, callID uint32, strPlayerName string, uiDetails uint32) (*nex.RMCMessage, uint32)
	ClearRelationship          func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32) (*nex.RMCMessage, uint32)
	UpdateDetails              func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32) (*nex.RMCMessage, uint32)
	GetList                    func(err error, packet nex.PacketInterface, callID uint32, byRelationship uint8, bReversed bool) (*nex.RMCMessage, uint32)
	GetDetailedList            func(err error, packet nex.PacketInterface, callID uint32, byRelationship uint8, bReversed bool) (*nex.RMCMessage, uint32)
	GetRelationships           func(err error, packet nex.PacketInterface, callID uint32, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
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
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Friends method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Friends protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
