// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRequestBlockSettings sets the GetRequestBlockSettings handler function
func (protocol *Protocol) GetRequestBlockSettings(handler func(err error, packet nex.PacketInterface, callID uint32, unknowns []uint32) uint32) {
	protocol.getRequestBlockSettingsHandler = handler
}

func (protocol *Protocol) handleGetRequestBlockSettings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getRequestBlockSettingsHandler == nil {
		globals.Logger.Warning("FriendsWiiU::GetRequestBlockSettings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getRequestBlockSettingsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getRequestBlockSettingsHandler(nil, packet, callID, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
