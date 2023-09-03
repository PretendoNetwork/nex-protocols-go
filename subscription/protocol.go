// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// SubscriptionProtocolID is the protocol ID for the Subscription protocol
	SubscriptionProtocolID = 0x75

	// SubscriptionCreateMySubscriptionDataID is the method ID for the method SubscriptionCreateMySubscriptionData
	SubscriptionCreateMySubscriptionDataID = 0x1

	// SubscriptionUpdateMySubscriptionDataID is the method ID for the method SubscriptionUpdateMySubscriptionData
	SubscriptionUpdateMySubscriptionDataID = 0x2

	// SubscriptionClearMySubscriptionDataID is the method ID for the method SubscriptionClearMySubscriptionDataID
	SubscriptionClearMySubscriptionDataID = 0x3

	// SubscriptionAddTargetID is the method ID for the method SubscriptionAddTarget
	SubscriptionAddTargetID = 0x4

	// SubscriptionDeleteTargetID is the method ID for the method SubscriptionDeleteTarget
	SubscriptionDeleteTargetID = 0x5

	// SubscriptionClearTargetID is the method ID for the method SubscriptionClearTarget
	SubscriptionClearTargetID = 0x6

	// SubscriptionGetFriendSubscriptionDataID is the method ID for the method SubscriptionGetFriendSubscriptionData
	SubscriptionGetFriendSubscriptionDataID = 0x7

	// SubscriptionGetTargetSubscriptionDataID is the method ID for the method SubscriptionGetTargetSubscriptionData
	SubscriptionGetTargetSubscriptionDataID = 0x8

	// SubscriptionGetActivePlayerSubscriptionDataID is the method ID for the method SubscriptionGetActivePlayerSubscriptionData
	SubscriptionGetActivePlayerSubscriptionDataID = 0x9

	// SubscriptionGetSubscriptionDataID is the method ID for the method SubscriptionGetSubscriptionData
	SubscriptionGetSubscriptionDataID = 0xA

	// SubscriptionReplaceTargetAndGetSubscriptionDataID is the method ID for the method SubscriptionReplaceTargetAndGetSubscriptionData
	SubscriptionReplaceTargetAndGetSubscriptionDataID = 0xB

	// SubscriptionSetPrivacyLevelID is the method ID for the method SubscriptionSetPrivacyLevel
	SubscriptionSetPrivacyLevelID = 0xC

	// SubscriptionGetPrivacyLevelID is the method ID for the method SubscriptionGetPrivacyLevel
	SubscriptionGetPrivacyLevelID = 0xD

	// SubscriptionGetSubscriptionUserFriendListID is the method ID for the method SubscriptionGetSubscriptionUserFriendList
	SubscriptionGetSubscriptionUserFriendListID = 0xE

	// SubscriptionGetPrivacyLevelsID is the method ID for the method SubscriptionGetPrivacyLevels
	SubscriptionGetPrivacyLevelsID = 0xF
)

// SubscriptionProtocol handles the Subscription nex protocol
type SubscriptionProtocol struct {
	Server                                      *nex.Server
	createMySubscriptionDataHandler             func(err error, client *nex.Client, callID uint32, unk uint64, content []byte)
	updateMySubscriptionDataHandler             func(err error, client *nex.Client, callID uint32, unk uint32, content []byte)
	getFriendSubscriptionDataHandler            func(err error, client *nex.Client, callID uint32)
	getTargetSubscriptionDataHandler            func(err error, client *nex.Client, callID uint32)
	getActivePlayerSubscriptionDataHandler      func(err error, client *nex.Client, callID uint32)
	getSubscriptionDataHandler                  func(err error, client *nex.Client, callID uint32, pids []uint32)
	replaceTargetAndGetSubscriptionDataHandler  func(err error, client *nex.Client, callID uint32)
	getPrivacyLevelsHandler                     func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *SubscriptionProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if SubscriptionProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case SubscriptionCreateMySubscriptionDataID:
				go protocol.handleCreateMySubscriptionData(packet)
			case SubscriptionUpdateMySubscriptionDataID:
				go protocol.handleUpdateMySubscriptionData(packet)
			case SubscriptionGetFriendSubscriptionDataID:
				go protocol.handleGetFriendSubscriptionData(packet)
			case SubscriptionGetTargetSubscriptionDataID:
				go protocol.handleGetTargetSubscriptionData(packet)
			case SubscriptionGetActivePlayerSubscriptionDataID:
				go protocol.handleGetActivePlayerSubscriptionData(packet)
			case SubscriptionGetSubscriptionDataID:
				go protocol.handleGetSubscriptionData(packet)
			case SubscriptionReplaceTargetAndGetSubscriptionDataID:
				go protocol.handleReplaceTargetAndGetSubscriptionData(packet)
			case SubscriptionGetPrivacyLevelsID:
				go protocol.handleGetPrivacyLevels(packet)
			default:
				go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
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
