// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

func (protocol *Protocol) handleUpdateUserStatus(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateUserStatus == nil {
		globals.Logger.Warning("Subscriber::UpdateUserStatus not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown1 := types.NewList[*subscriber_types.Unknown]()
	unknown1.Type = subscriber_types.NewUnknown()
	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2 := types.NewList[*types.PrimitiveU8]()
	unknown2.Type = types.NewPrimitiveU8(0)
	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateUserStatus(nil, packet, callID, unknown1, unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
