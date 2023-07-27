// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRequestBlockSettings sets the GetRequestBlockSettings handler function
func (protocol *Protocol) GetRequestBlockSettings(handler func(err error, client *nex.Client, callID uint32, unknowns []uint32)) {
	protocol.getRequestBlockSettingsHandler = handler
}

func (protocol *Protocol) handleGetRequestBlockSettings(packet nex.PacketInterface) {
	if protocol.getRequestBlockSettingsHandler == nil {
		globals.Logger.Warning("FriendsWiiU::GetRequestBlockSettings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getRequestBlockSettingsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getRequestBlockSettingsHandler(nil, client, callID, pids)
}
