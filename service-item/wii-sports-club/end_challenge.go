// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// EndChallenge sets the EndChallenge handler function
func (protocol *Protocol) EndChallenge(handler func(err error, client *nex.Client, callID uint32, endChallengeParam *service_item_wii_sports_club_types.ServiceItemEndChallengeParam) uint32) {
	protocol.endChallengeHandler = handler
}

func (protocol *Protocol) handleEndChallenge(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.endChallengeHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::EndChallenge not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	endChallengeParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemEndChallengeParam())
	if err != nil {
		errorCode = protocol.endChallengeHandler(fmt.Errorf("Failed to read endChallengeParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.endChallengeHandler(nil, client, callID, endChallengeParam.(*service_item_wii_sports_club_types.ServiceItemEndChallengeParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
