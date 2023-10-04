// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTargetSubscriptionData sets the GetTargetSubscriptionData handler function
func (protocol *SubscriptionProtocol) GetTargetSubscriptionData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getTargetSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetTargetSubscriptionData(packet nex.PacketInterface) {
	if protocol.getTargetSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetTargetSubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getTargetSubscriptionDataHandler(nil, client, callID)
}