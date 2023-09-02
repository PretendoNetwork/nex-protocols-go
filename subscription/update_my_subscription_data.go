// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMySubscriptionData sets the UpdateMySubscriptionData handler function
func (protocol *SubscriptionProtocol) UpdateMySubscriptionData(handler func(err error, client *nex.Client, callID uint32, content []byte)) {
	protocol.UpdateMySubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.UpdateMySubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::UpdateMySubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	content := parameters[4:]

	go protocol.UpdateMySubscriptionDataHandler(nil, client, callID, content)
}