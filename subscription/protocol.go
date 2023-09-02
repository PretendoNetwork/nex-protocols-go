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

	// SubscriptionGetFriendSubscriptionDataID is the method ID for the method SubscriptionGetFriendSubscriptionData
	SubscriptionGetFriendSubscriptionDataID = 0x7

	// SubscriptionGetTargetSubscriptionDataID is the method ID for the method SubscriptionGetTargetSubscriptionData
	SubscriptionGetTargetSubscriptionDataID = 0x8

	// SubscriptionGetActivePlayerSubscriptionDataID is the method ID for the method SubscriptionGetActivePlayerSubscriptionData
	SubscriptionGetActivePlayerSubscriptionDataID = 0x9

	// SubscriptionGetSubscriptionDataID is the method ID for the method SubscriptionGetSubscriptionData
	SubscriptionGetSubscriptionDataID = 0xA

	// SubscriptionReplaceTargetAndGetSubscriptionDataID is the method ID for the method SubscriptionReplaceTargetAndGetSubscriptionData
	SubscriptionReplaceTargetAndGetSubscriptionDataID = 0xb

	// SubscriptionGetPrivacyLevelsID is the method ID for the method SubscriptionGetPrivacyLevels
	SubscriptionGetPrivacyLevelsID = 0xf
)

// SubscriptionProtocol handles the Subscription nex protocol
type SubscriptionProtocol struct {
	Server										*nex.Server
	CreateMySubscriptionDataHandler				func(err error, client *nex.Client, callID uint32, content []byte)
	UpdateMySubscriptionDataHandler				func(err error, client *nex.Client, callID uint32, content []byte)
	GetFriendSubscriptionDataHandler			func(err error, client *nex.Client, callID uint32)
	GetTargetSubscriptionDataHandler			func(err error, client *nex.Client, callID uint32)
	GetActivePlayerSubscriptionDataHandler		func(err error, client *nex.Client, callID uint32)
	GetSubscriptionDataHandler					func(err error, client *nex.Client, callID uint32, pids []uint32)
	ReplaceTargetAndGetSubscriptionDataHandler	func(err error, client *nex.Client, callID uint32)
	GetPrivacyLevelsHandler						func(err error, client *nex.Client, callID uint32)
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
