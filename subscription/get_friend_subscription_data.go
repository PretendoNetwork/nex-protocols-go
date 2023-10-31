// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendSubscriptionData sets GetFriendSubscriptionData Unk1 handler function
func (protocol *SubscriptionProtocol) GetFriendSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32)) {
	protocol.getFriendSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetFriendSubscriptionData(packet nex.PacketInterface) {
	if protocol.getFriendSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetFriendSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getFriendSubscriptionDataHandler(nil, packet, callID)
}
