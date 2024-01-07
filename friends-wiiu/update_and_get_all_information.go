// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateAndGetAllInformation == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateAndGetAllInformation not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	nnaInfo := friends_wiiu_types.NewNNAInfo()
	err = nnaInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read nnaInfo from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	presence := friends_wiiu_types.NewNintendoPresenceV2()
	err = presence.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read presence from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	birthday := types.NewDateTime(0)
	err = birthday.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read birthday from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateAndGetAllInformation(nil, packet, callID, nnaInfo, presence, birthday)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
