// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetRivToken(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetRivToken == nil {
		globals.Logger.Warning("ShopNintendoBadgeArcade::GetRivToken not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	itemCode := types.NewString("")
	err = itemCode.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRivToken(fmt.Errorf("Failed to read itemCode from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	referenceID := types.NewQBuffer(nil)
	err = referenceID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRivToken(fmt.Errorf("Failed to read referenceID from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRivToken(nil, packet, callID, itemCode, referenceID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
