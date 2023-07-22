// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePreference sets the UpdatePreference handler function
func (protocol *Friends3DSProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)) {
	protocol.updatePreferenceHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdatePreference(packet nex.PacketInterface) {
	if protocol.updatePreferenceHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdatePreference not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	publicMode, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePreferenceHandler(fmt.Errorf("Failed to read publicMode from parameters. %s", err.Error()), client, callID, false, false, false)
		return
	}

	showGame, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePreferenceHandler(fmt.Errorf("Failed to read showGame from parameters. %s", err.Error()), client, callID, false, false, false)
		return
	}

	showPlayedGame, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.updatePreferenceHandler(fmt.Errorf("Failed to read showPlayedGame from parameters. %s", err.Error()), client, callID, false, false, false)
		return
	}

	go protocol.updatePreferenceHandler(nil, client, callID, publicMode, showGame, showPlayedGame)
}
