// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMySubscriptionData sets the UpdateMySubscriptionData handler function
func (protocol *SubscriptionProtocol) UpdateMySubscriptionData(handler func(err error, packet nex.PacketInterface, callID uint32, unk uint32, content []byte)) {
	protocol.updateMySubscriptionDataHandler = handler
}

func (protocol *SubscriptionProtocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.updateMySubscriptionDataHandler == nil {
		fmt.Println("[Warning] SubscriptionProtocol::UpdateMySubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	unk, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.updateMySubscriptionDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil)
		return
	}

	//This is done since the server doesn't need to care about the data here (it's game-specific), so we just pass it along to store however the handler wants
	content := parametersStream.ReadRemaining()
	go protocol.updateMySubscriptionDataHandler(nil, packet, callID, unk, content)
}
