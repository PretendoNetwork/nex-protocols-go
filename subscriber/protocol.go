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
	Server                       *nex.Server
	helloHandler                 func(err error, packet nex.PacketInterface, callID uint32, unknown string) uint32
	postContentHandler           func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberPostContentParam) uint32
	getContentHandler            func(err error, packet nex.PacketInterface, callID uint32, param *subscriber_types.SubscriberGetContentParam) uint32
	followHandler                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	unfollowAllAndFollowHandler  func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	unfollowHandler              func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getFollowingHandler          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getFollowerHandler           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getNumFollowersHandler       func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTimelineHandler           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	deleteContentHandler         func(err error, packet nex.PacketInterface, callID uint32, unknown1 []string, unknown2 uint64) uint32
	getContentMultiHandler       func(err error, packet nex.PacketInterface, callID uint32, params []*subscriber_types.SubscriberGetContentParam) uint32
	updateUserStatusHandler      func(err error, packet nex.PacketInterface, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8) uint32
	getFriendUserStatusesHandler func(err error, packet nex.PacketInterface, callID uint32, unknown []uint8) uint32
	getUserStatusesHandler       func(err error, packet nex.PacketInterface, callID uint32, pids []uint32, unknown []uint8) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodHello:
				go protocol.handleHello(packet)
			case MethodPostContent:
				go protocol.handlePostContent(packet)
			case MethodGetContent:
				go protocol.handleGetContent(packet)
			case MethodFollow:
				go protocol.handleFollow(packet)
			case MethodUnfollowAllAndFollow:
				go protocol.handleUnfollowAllAndFollow(packet)
			case MethodUnfollow:
				go protocol.handleUnfollow(packet)
			case MethodGetFollowing:
				go protocol.handleGetFollowing(packet)
			case MethodGetFollower:
				go protocol.handleGetFollower(packet)
			case MethodGetNumFollowers:
				go protocol.handleGetNumFollowers(packet)
			case MethodGetTimeline:
				go protocol.handleGetTimeline(packet)
			case MethodDeleteContent:
				go protocol.handleDeleteContent(packet)
			case MethodGetContentMulti:
				go protocol.handleGetContentMulti(packet)
			case MethodUpdateUserStatus:
				go protocol.handleUpdateUserStatus(packet)
			case MethodGetFriendUserStatuses:
				go protocol.handleGetFriendUserStatuses(packet)
			case MethodGetUserStatuses:
				go protocol.handleGetUserStatuses(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Subscriber method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Subscriber protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
