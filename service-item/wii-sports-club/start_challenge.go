// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleStartChallenge(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.StartChallenge == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::StartChallenge not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	startChallengeParam, err := nex.StreamReadStructure(parametersStream, service_item_wii_sports_club_types.NewServiceItemStartChallengeParam())
	if err != nil {
		_, errorCode = protocol.StartChallenge(fmt.Errorf("Failed to read startChallengeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.StartChallenge(nil, packet, callID, startChallengeParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
