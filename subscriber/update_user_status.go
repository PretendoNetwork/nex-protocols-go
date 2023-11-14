// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/subscriber/types"
)

func (protocol *Protocol) handleUpdateUserStatus(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateUserStatus == nil {
		globals.Logger.Warning("Subscriber::UpdateUserStatus not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown1, err := parametersStream.ReadListStructure(subscriber_types.NewUnknown())
	if err != nil {
		errorCode = protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadListUInt8()
	if err != nil {
		errorCode = protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdateUserStatus(nil, packet, callID, unknown1.([]*subscriber_types.Unknown), unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
