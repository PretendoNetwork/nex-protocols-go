// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteContent(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeleteContent == nil {
		globals.Logger.Warning("Subscriber::DeleteContent not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown1 := types.NewList[*types.String]()
	unknown1.Type = types.NewString("")
	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteContent(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2 := types.NewPrimitiveU64(0)
	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteContent(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteContent(nil, packet, callID, unknown1, unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
