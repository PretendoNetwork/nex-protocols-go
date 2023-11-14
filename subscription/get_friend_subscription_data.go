// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetFriendSubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetFriendSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.GetFriendSubscriptionData(nil, packet, callID)
}
