// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceTargetAndGetSubscriptionData sets the ReplaceTargetAndGetSubscriptionData handler function
func (protocol *SubscriptionProtocol) ReplaceTargetAndGetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32)) {
	protocol.replaceTargetAndGetSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleReplaceTargetAndGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.replaceTargetAndGetSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::ReplaceTargetAndGetSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.replaceTargetAndGetSubscriptionDataHandler(nil, packet, callID)
}
