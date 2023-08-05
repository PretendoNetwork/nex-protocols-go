// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// StartChallenge sets the StartChallenge handler function
func (protocol *Protocol) StartChallenge(handler func(err error, client *nex.Client, callID uint32, startChallengeParam *service_item_wii_sports_club_types.ServiceItemStartChallengeParam) uint32) {
	protocol.startChallengeHandler = handler
}

func (protocol *Protocol) handleStartChallenge(packet nex.PacketInterface) {
	if protocol.startChallengeHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::StartChallenge not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	startChallengeParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemStartChallengeParam())
	if err != nil {
		go protocol.startChallengeHandler(fmt.Errorf("Failed to read startChallengeParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.startChallengeHandler(nil, client, callID, startChallengeParam.(*service_item_wii_sports_club_types.ServiceItemStartChallengeParam))
}
