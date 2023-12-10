// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

const (
	// ProtocolID is the protocol ID for the Subscriber protocol
	ProtocolID = 0x6E

	// MethodHello is the method ID for the method Hello
	MethodHello = 0x1

	// MethodPostContent is the method ID for the method PostContent
	MethodPostContent = 0x2

	// MethodGetContent is the method ID for the method GetContent
	MethodGetContent = 0x3

	// MethodFollow is the method ID for the method Follow
	MethodFollow = 0x4

	// MethodUnfollowAllAndFollow is the method ID for the method UnfollowAllAndFollow
	MethodUnfollowAllAndFollow = 0x5

	// MethodUnfollow is the method ID for the method Unfollow
	MethodUnfollow = 0x6

	// MethodGetFollowing is the method ID for the method GetFollowing
	MethodGetFollowing = 0x7

	// MethodGetFollower is the method ID for the method GetFollower
	MethodGetFollower = 0x8

	// MethodGetNumFollowers is the method ID for the method GetNumFollowers
	MethodGetNumFollowers = 0x9

	// MethodGetTimeline is the method ID for the method GetTimeline
	MethodGetTimeline = 0xA

	// MethodDeleteContent is the method ID for the method DeleteContent
	MethodDeleteContent = 0xB

	// MethodGetContentMulti is the method ID for the method GetContentMulti
	MethodGetContentMulti = 0xC

	// MethodUpdateUserStatus is the method ID for the method UpdateUserStatus
	MethodUpdateUserStatus = 0xD

	// MethodGetFriendUserStatuses is the method ID for the method GetFriendUserStatuses
	MethodGetFriendUserStatuses = 0xE

	// MethodGetUserStatuses is the method ID for the method GetUserStatuses
	MethodGetUserStatuses = 0xF
)

// Protocol stores all the RMC method handlers for the Subscriber protocol and listens for requests
type Protocol struct {
	server                nex.ServerInterface
	Hello                 func(err error, packet nex.PacketInterface, callID uint32, unknown string) (*nex.RMCMessage, uint32)
	PostContent           func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, uint32)
	GetContent            func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32)
	Follow                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	UnfollowAllAndFollow  func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	Unfollow              func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetFollowing          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetFollower           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetNumFollowers       func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTimeline           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DeleteContent         func(err error, packet nex.PacketInterface, callID uint32, unknown1 []string, unknown2 uint64) (*nex.RMCMessage, uint32)
	GetContentMulti       func(err error, packet nex.PacketInterface, callID uint32, params []*subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32)
	UpdateUserStatus      func(err error, packet nex.PacketInterface, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8) (*nex.RMCMessage, uint32)
	GetFriendUserStatuses func(err error, packet nex.PacketInterface, callID uint32, unknown []uint8) (*nex.RMCMessage, uint32)
	GetUserStatuses       func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown []uint8) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Subscriber protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerHello(handler func(err error, packet nex.PacketInterface, callID uint32, unknown string) (*nex.RMCMessage, uint32))
	SetHandlerPostContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, uint32))
	SetHandlerGetContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32))
	SetHandlerFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerUnfollowAllAndFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerUnfollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetFollowing(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetFollower(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetNumFollowers(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetTimeline(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 []string, unknown2 uint64) (*nex.RMCMessage, uint32))
	SetHandlerGetContentMulti(handler func(err error, packet nex.PacketInterface, callID uint32, params []*subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32))
	SetHandlerUpdateUserStatus(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, unknown []uint8) (*nex.RMCMessage, uint32))
	SetHandlerGetUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown []uint8) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerHello sets the handler for the Hello method
func (protocol *Protocol) SetHandlerHello(handler func(err error, packet nex.PacketInterface, callID uint32, unknown string) (*nex.RMCMessage, uint32)) {
	protocol.Hello = handler
}

// SetHandlerPostContent sets the handler for the PostContent method
func (protocol *Protocol) SetHandlerPostContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) (*nex.RMCMessage, uint32)) {
	protocol.PostContent = handler
}

// SetHandlerGetContent sets the handler for the GetContent method
func (protocol *Protocol) SetHandlerGetContent(handler func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32)) {
	protocol.GetContent = handler
}

// SetHandlerFollow sets the handler for the Follow method
func (protocol *Protocol) SetHandlerFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.Follow = handler
}

// SetHandlerUnfollowAllAndFollow sets the handler for the UnfollowAllAndFollow method
func (protocol *Protocol) SetHandlerUnfollowAllAndFollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.UnfollowAllAndFollow = handler
}

// SetHandlerUnfollow sets the handler for the Unfollow method
func (protocol *Protocol) SetHandlerUnfollow(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.Unfollow = handler
}

// SetHandlerGetFollowing sets the handler for the GetFollowing method
func (protocol *Protocol) SetHandlerGetFollowing(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetFollowing = handler
}

// SetHandlerGetFollower sets the handler for the GetFollower method
func (protocol *Protocol) SetHandlerGetFollower(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetFollower = handler
}

// SetHandlerGetNumFollowers sets the handler for the GetNumFollowers method
func (protocol *Protocol) SetHandlerGetNumFollowers(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetNumFollowers = handler
}

// SetHandlerGetTimeline sets the handler for the GetTimeline method
func (protocol *Protocol) SetHandlerGetTimeline(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetTimeline = handler
}

// SetHandlerDeleteContent sets the handler for the DeleteContent method
func (protocol *Protocol) SetHandlerDeleteContent(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 []string, unknown2 uint64) (*nex.RMCMessage, uint32)) {
	protocol.DeleteContent = handler
}

// SetHandlerGetContentMulti sets the handler for the GetContentMulti method
func (protocol *Protocol) SetHandlerGetContentMulti(handler func(err error, packet nex.PacketInterface, callID uint32, params []*subscriber_types.SubscriberGetContentParam) (*nex.RMCMessage, uint32)) {
	protocol.GetContentMulti = handler
}

// SetHandlerUpdateUserStatus sets the handler for the UpdateUserStatus method
func (protocol *Protocol) SetHandlerUpdateUserStatus(handler func(err error, packet nex.PacketInterface, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8) (*nex.RMCMessage, uint32)) {
	protocol.UpdateUserStatus = handler
}

// SetHandlerGetFriendUserStatuses sets the handler for the GetFriendUserStatuses method
func (protocol *Protocol) SetHandlerGetFriendUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, unknown []uint8) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendUserStatuses = handler
}

// SetHandlerGetUserStatuses sets the handler for the GetUserStatuses method
func (protocol *Protocol) SetHandlerGetUserStatuses(handler func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown []uint8) (*nex.RMCMessage, uint32)) {
	protocol.GetUserStatuses = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodHello:
				protocol.handleHello(packet)
			case MethodPostContent:
				protocol.handlePostContent(packet)
			case MethodGetContent:
				protocol.handleGetContent(packet)
			case MethodFollow:
				protocol.handleFollow(packet)
			case MethodUnfollowAllAndFollow:
				protocol.handleUnfollowAllAndFollow(packet)
			case MethodUnfollow:
				protocol.handleUnfollow(packet)
			case MethodGetFollowing:
				protocol.handleGetFollowing(packet)
			case MethodGetFollower:
				protocol.handleGetFollower(packet)
			case MethodGetNumFollowers:
				protocol.handleGetNumFollowers(packet)
			case MethodGetTimeline:
				protocol.handleGetTimeline(packet)
			case MethodDeleteContent:
				protocol.handleDeleteContent(packet)
			case MethodGetContentMulti:
				protocol.handleGetContentMulti(packet)
			case MethodUpdateUserStatus:
				protocol.handleUpdateUserStatus(packet)
			case MethodGetFriendUserStatuses:
				protocol.handleGetFriendUserStatuses(packet)
			case MethodGetUserStatuses:
				protocol.handleGetUserStatuses(packet)
			default:
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Subscriber method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Subscriber protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
