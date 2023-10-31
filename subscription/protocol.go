// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Subscription protocol
	ProtocolID = 0x75

	// MethodCreateMySubscriptionData is the method ID for the method CreateMySubscriptionDataID
	MethodCreateMySubscriptionData = 0x1

	// MethodUpdateMySubscriptionData is the method ID for the method UpdateMySubscriptionData
	MethodUpdateMySubscriptionData = 0x2

	// MethodClearMySubscriptionData is the method ID for the method ClearMySubscriptionDataID
	MethodClearMySubscriptionData = 0x3

	// MethodAddTarget is the method ID for the method AddTarget
	MethodAddTarget = 0x4

	// MethodDeleteTarget is the method ID for the method DeleteTarget
	MethodDeleteTarget = 0x5

	// MethodClearTarget is the method ID for the method ClearTarget
	MethodClearTarget = 0x6

	// MethodGetFriendSubscriptionData is the method ID for the method GetFriendSubscriptionData
	MethodGetFriendSubscriptionData = 0x7

	// MethodGetTargetSubscriptionData is the method ID for the method GetTargetSubscriptionData
	MethodGetTargetSubscriptionData = 0x8

	// MethodGetActivePlayerSubscriptionData is the method ID for the method GetActivePlayerSubscriptionData
	MethodGetActivePlayerSubscriptionData = 0x9

	// MethodGetSubscriptionData is the method ID for the method GetSubscriptionData
	MethodGetSubscriptionData = 0xA

	// MethodReplaceTargetAndGetSubscriptionData is the method ID for the method ReplaceTargetAndGetSubscriptionData
	MethodReplaceTargetAndGetSubscriptionData = 0xB

	// MethodSetPrivacyLevel is the method ID for the method SetPrivacyLevel
	MethodSetPrivacyLevel = 0xC

	// MethodGetPrivacyLevel is the method ID for the method GetPrivacyLevel
	MethodGetPrivacyLevel = 0xD

	// MethodGetSubscriptionUserFriendList is the method ID for the method GetSubscriptionUserFriendList
	MethodGetSubscriptionUserFriendList = 0xE

	// MethodGetPrivacyLevels is the method ID for the method GetPrivacyLevels
	MethodGetPrivacyLevels = 0xF
)

// SubscriptionProtocol handles the Subscription nex protocol
type SubscriptionProtocol struct {
	Server                                     *nex.Server
	createMySubscriptionDataHandler            func(err error, packet nex.PacketInterface, callID uint32, unk uint64, content []byte)
	updateMySubscriptionDataHandler            func(err error, packet nex.PacketInterface, callID uint32, unk uint32, content []byte)
	getFriendSubscriptionDataHandler           func(err error, packet nex.PacketInterface, callID uint32)
	getTargetSubscriptionDataHandler           func(err error, packet nex.PacketInterface, callID uint32)
	getActivePlayerSubscriptionDataHandler     func(err error, packet nex.PacketInterface, callID uint32)
	getSubscriptionDataHandler                 func(err error, packet nex.PacketInterface, callID uint32, pids []uint32)
	replaceTargetAndGetSubscriptionDataHandler func(err error, packet nex.PacketInterface, callID uint32)
	getPrivacyLevelsHandler                    func(err error, packet nex.PacketInterface, callID uint32)
}

// Setup initializes the protocol
func (protocol *SubscriptionProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if ProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MethodCreateMySubscriptionData:
				go protocol.handleCreateMySubscriptionData(packet)
			case MethodUpdateMySubscriptionData:
				go protocol.handleUpdateMySubscriptionData(packet)
			case MethodGetFriendSubscriptionData:
				go protocol.handleGetFriendSubscriptionData(packet)
			case MethodGetTargetSubscriptionData:
				go protocol.handleGetTargetSubscriptionData(packet)
			case MethodGetActivePlayerSubscriptionData:
				go protocol.handleGetActivePlayerSubscriptionData(packet)
			case MethodGetSubscriptionData:
				go protocol.handleGetSubscriptionData(packet)
			case MethodReplaceTargetAndGetSubscriptionData:
				go protocol.handleReplaceTargetAndGetSubscriptionData(packet)
			case MethodGetPrivacyLevels:
				go protocol.handleGetPrivacyLevels(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported Subscription method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewSubscriptionProtocol returns a new SubscriptionProtocol
func NewSubscriptionProtocol(server *nex.Server) *SubscriptionProtocol {
	protocol := &SubscriptionProtocol{Server: server}

	protocol.Setup()

	return protocol
}
