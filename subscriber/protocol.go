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
	helloHandler                 func(err error, client *nex.Client, callID uint32, unknown string)
	postContentHandler           func(err error, client *nex.Client, callID uint32, param *subscriber_types.SubscriberPostContentParam)
	getContentHandler            func(err error, client *nex.Client, callID uint32, param *subscriber_types.SubscriberGetContentParam)
	followHandler                func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	unfollowAllAndFollowHandler  func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	unfollowHandler              func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	getFollowingHandler          func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	getFollowerHandler           func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	getNumFollowersHandler       func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	getTimelineHandler           func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	deleteContentHandler         func(err error, client *nex.Client, callID uint32, unknown1 []string, unknown2 uint64)
	getContentMultiHandler       func(err error, client *nex.Client, callID uint32, params []*subscriber_types.SubscriberGetContentParam)
	updateUserStatusHandler      func(err error, client *nex.Client, callID uint32, unknown1 []*subscriber_types.Unknown, unknown2 []uint8)
	getFriendUserStatusesHandler func(err error, client *nex.Client, callID uint32, unknown []uint8)
	getUserStatusesHandler       func(err error, client *nex.Client, callID uint32, pids []uint32, unknown []uint8)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
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
				go globals.RespondNotImplemented(packet, ProtocolID)
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
