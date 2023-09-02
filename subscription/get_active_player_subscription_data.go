// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetActivePlayerSubscriptionData sets the GetActivePlayerSubscriptionData handler function
func (protocol *SubscriptionProtocol) GetActivePlayerSubscriptionData(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetActivePlayerSubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleGetActivePlayerSubscriptionData(packet nex.PacketInterface) {
	if protocol.GetActivePlayerSubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::GetActivePlayerSubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	
	go protocol.GetActivePlayerSubscriptionDataHandler(nil, client, callID)
}