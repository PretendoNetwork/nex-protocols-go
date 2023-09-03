// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceTargetAndGetSubscriptionData sets the ReplaceTargetAndGetSubscriptionData handler function
func (protocol *SubscriptionProtocol) ReplaceTargetAndGetSubscriptionData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.replaceTargetAndGetSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleReplaceTargetAndGetSubscriptionData(packet nex.PacketInterface) {
	if protocol.replaceTargetAndGetSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::ReplaceTargetAndGetSubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.replaceTargetAndGetSubscriptionDataHandler(nil, client, callID)
}