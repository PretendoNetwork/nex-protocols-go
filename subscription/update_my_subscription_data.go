// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.UpdateMySubscriptionData == nil {
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
		go protocol.UpdateMySubscriptionData(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil)
		return
	}

	//This is done since the server doesn't need to care about the data here (it's game-specific), so we just pass it along to store however the handler wants
	content := parametersStream.ReadRemaining()
	go protocol.UpdateMySubscriptionData(nil, packet, callID, unk, content)
}
