// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMySubscriptionData sets the UpdateMySubscriptionData handler function
func (protocol *SubscriptionProtocol) UpdateMySubscriptionData(handler func(err error, client *nex.Client, callID uint32, unk uint32, content []byte)) {
	protocol.updateMySubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.updateMySubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::UpdateMySubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	//This is done since the server doesn't need to care about the data here (it's game-specific), so we just pass it along to store however the handler wants
	content := parameters[4:]
	go protocol.updateMySubscriptionDataHandler(nil, client, callID, 0, content)
}