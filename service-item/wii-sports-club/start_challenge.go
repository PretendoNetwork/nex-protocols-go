// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// StartChallenge sets the StartChallenge handler function
func (protocol *Protocol) StartChallenge(handler func(err error, packet nex.PacketInterface, callID uint32, startChallengeParam *service_item_wii_sports_club_types.ServiceItemStartChallengeParam) uint32) {
	protocol.startChallengeHandler = handler
}

func (protocol *Protocol) handleStartChallenge(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.startChallengeHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::StartChallenge not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	startChallengeParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemStartChallengeParam())
	if err != nil {
		errorCode = protocol.startChallengeHandler(fmt.Errorf("Failed to read startChallengeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.startChallengeHandler(nil, packet, callID, startChallengeParam.(*service_item_wii_sports_club_types.ServiceItemStartChallengeParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
