// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateMii == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateMii not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	miiV2 := friends_wiiu_types.NewMiiV2()
	err = miiV2.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateMii(fmt.Errorf("Failed to read miiV2 from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateMii(nil, packet, callID, miiV2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
