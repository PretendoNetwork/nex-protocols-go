// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendSubscriptionData sets GetFriendSubscriptionData Unk1 handler function
func (protocol *SubscriptionProtocol) GetFriendSubscriptionData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetFriendSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetFriendSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetFriendSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetFriendSubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.GetFriendSubscriptionDataHandler(nil, client, callID)
}