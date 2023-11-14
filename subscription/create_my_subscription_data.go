// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.CreateMySubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::CreateMySubscriptionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	unk, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil)
		return
	}

	//This is done since the server doesn't need to care about the data here (it's game-specific), so we just pass it along to store however the handler wants
	content := parametersStream.ReadRemaining()
	go protocol.CreateMySubscriptionData(nil, packet, callID, unk, content)
}
