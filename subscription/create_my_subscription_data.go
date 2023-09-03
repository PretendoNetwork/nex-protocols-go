// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateMySubscriptionData sets the CreateMySubscriptionData handler function
func (protocol *SubscriptionProtocol) CreateMySubscriptionData(handler func(err error, client *nex.Client, callID uint32, unk uint64, content []byte)) {
	protocol.createMySubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.createMySubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::CreateMySubscriptionData not implemented")
		go globals.RespondError(packet, SubscriptionProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	unk, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.createMySubscriptionDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	//This is done since the server doesn't need to care about the data here (it's game-specific), so we just pass it along to store however the handler wants
	content := parameters[8:]
	go protocol.createMySubscriptionDataHandler(nil, client, callID, unk, content)
}
