// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTargetSubscriptionData sets the GetTargetSubscriptionData handler function
func (protocol *SubscriptionProtocol) GetTargetSubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32)) {
	protocol.getTargetSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetTargetSubscriptionData(packet nex.PacketInterface) {
	if protocol.getTargetSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetTargetSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	go protocol.getTargetSubscriptionDataHandler(nil, packet, callID)
}
