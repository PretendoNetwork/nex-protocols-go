// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePrivacySetting sets the UpdatePrivacySetting handler function
func (protocol *Protocol) UpdatePrivacySetting(handler func(err error, client *nex.Client, callID uint32, onlineStatus bool, participationCommunity bool)) {
	protocol.updatePrivacySettingHandler = handler
}

func (protocol *Protocol) handleUpdatePrivacySetting(packet nex.PacketInterface) {
	if protocol.updatePrivacySettingHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdatePrivacySetting not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	onlineStatus, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePrivacySettingHandler(fmt.Errorf("Failed to read onlineStatus from parameters. %s", err.Error()), client, callID, false, false)
		return
	}

	participationCommunity, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePrivacySettingHandler(fmt.Errorf("Failed to read participationCommunity from parameters. %s", err.Error()), client, callID, false, false)
		return
	}

	go protocol.updatePrivacySettingHandler(nil, client, callID, onlineStatus, participationCommunity)
}
