// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetState sets the SetState handler function
func (protocol *MatchMakingProtocol) SetState(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, uiNewState uint32)) {
	protocol.setStateHandler = handler
}

func (protocol *MatchMakingProtocol) handleSetState(packet nex.PacketInterface) {
	if protocol.setStateHandler == nil {
		globals.Logger.Warning("MatchMaking::SetState not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.setStateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, 0)
	}

	uiNewState, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.setStateHandler(fmt.Errorf("Failed to read uiNewState from parameters. %s", err.Error()), client, callID, 0, 0)
	}

	go protocol.setStateHandler(nil, client, callID, idGathering, uiNewState)
}
