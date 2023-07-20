// Package friends implements the Friends QRV protocol
package friends

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

// FriendsProtocol handles the Friends QRV protocol
type FriendsProtocol struct {
	Server                            *nex.Server
	addFriendHandler                  func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string)
	addFriendByNameHandler            func(err error, client *nex.Client, callID uint32, strPlayerName string, uiDetails uint32, strMessage string)
	addFriendWithDetailsHandler       func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string)
	addFriendByNameWithDetailsHandler func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string) // TODO - Is this the right signature?
	acceptFriendshipHandler           func(err error, client *nex.Client, callID uint32, uiPlayer uint32)
	declineFriendshipHandler          func(err error, client *nex.Client, callID uint32, uiPlayer uint32)
	blackListHandler                  func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32)
	blackListByNameHandler            func(err error, client *nex.Client, callID uint32, strPlayerName string, uiDetails uint32)
	clearRelationshipHandler          func(err error, client *nex.Client, callID uint32, uiPlayer uint32)
	updateDetailsHandler              func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32)
	getListHandler                    func(err error, client *nex.Client, callID uint32, byRelationship uint8, bReversed bool)
	getDetailedListHandler            func(err error, client *nex.Client, callID uint32, byRelationship uint8, bReversed bool)
	getRelationshipsHandler           func(err error, client *nex.Client, callID uint32, resultRange *nex.ResultRange)
}

// Setup initializes the protocol
func (protocol *FriendsProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *FriendsProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodAddFriend:
		go protocol.handleAddFriend(packet)
	case MethodAddFriendByName:
		go protocol.handleAddFriendByName(packet)
	case MethodAddFriendWithDetails:
		go protocol.handleAddFriendWithDetails(packet)
	case MethodAddFriendByNameWithDetails:
		go protocol.handleAddFriendByNameWithDetails(packet)
	case MethodAcceptFriendship:
		go protocol.handleAcceptFriendship(packet)
	case MethodDeclineFriendship:
		go protocol.handleDeclineFriendship(packet)
	case MethodBlackList:
		go protocol.handleBlackList(packet)
	case MethodBlackListByName:
		go protocol.handleBlackListByName(packet)
	case MethodClearRelationship:
		go protocol.handleClearRelationship(packet)
	case MethodUpdateDetails:
		go protocol.handleUpdateDetails(packet)
	case MethodGetList:
		go protocol.handleGetList(packet)
	case MethodGetDetailedList:
		go protocol.handleGetDetailedList(packet)
	case MethodGetRelationships:
		go protocol.handleGetRelationships(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Friends method ID: %#v\n", request.MethodID())
	}
}

// NewFriendsProtocol returns a new FriendsProtocol
func NewFriendsProtocol(server *nex.Server) *FriendsProtocol {
	protocol := &FriendsProtocol{Server: server}

	protocol.Setup()

	return protocol
}
