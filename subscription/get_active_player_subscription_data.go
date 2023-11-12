// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetActivePlayerSubscriptionData sets the GetActivePlayerSubscriptionData handler function
func (protocol *SubscriptionProtocol) GetActivePlayerSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32)) {
	protocol.getActivePlayerSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetActivePlayerSubscriptionData(packet nex.PacketInterface) {
	if protocol.getActivePlayerSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetActivePlayerSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.getActivePlayerSubscriptionDataHandler(nil, packet, callID)
}
