// Package subscription implements the Subscription NEX protocol
package subscription

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCreateMySubscriptionData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CreateMySubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::CreateMySubscriptionData not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unk, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.CreateMySubscriptionData(fmt.Errorf("Failed to read unk from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	// * This is done since the server doesn't need to care about the data here (it's game-specific)
	// * so we just pass it along to store however the handler wants
	rmcMessage, errorCode := protocol.CreateMySubscriptionData(nil, packet, callID, unk, parametersStream.ReadRemaining())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
