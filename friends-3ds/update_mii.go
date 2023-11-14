// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateMii == nil {
		globals.Logger.Warning("Friends3DS::UpdateMii not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mii, err := parametersStream.ReadStructure(friends_3ds_types.NewMii())
	if err != nil {
		errorCode = protocol.UpdateMii(fmt.Errorf("Failed to read mii from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.UpdateMii(nil, packet, callID, mii.(*friends_3ds_types.Mii))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
