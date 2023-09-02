// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateMySubscriptionData sets the CreateMySubscriptionData handler function
func (protocol *SubscriptionProtocol) CreateMySubscriptionData(handler func(err error, client *nex.Client, callID uint32, content []byte)) {
	protocol.CreateMySubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.CreateMySubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::CreateMySubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	content := parameters[8:]

	go protocol.CreateMySubscriptionDataHandler(nil, client, callID, content)
}
