// Package protocol implements the Subscription protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateMySubscriptionData == nil {
		fmt.Println("[Warning] SubscriptionProtocol::UpdateMySubscriptionData not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unk := types.NewPrimitiveU32(0)
	err = unk.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateMySubscriptionData(fmt.Errorf("Failed to read unk from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	// * This is done since the server doesn't need to care about the data here (it's game-specific)
	// * so we just pass it along to store however the handler wants
	rmcMessage, errorCode := protocol.UpdateMySubscriptionData(nil, packet, callID, unk, parametersStream.ReadRemaining())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
