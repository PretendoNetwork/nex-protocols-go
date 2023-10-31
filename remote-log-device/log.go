// Package protocol implements the Remote Log Device protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Log sets the Log handler function
func (protocol *Protocol) Log(handler func(err error, packet nex.PacketInterface, callID uint32, strLine string) uint32) {
	protocol.logHandler = handler
}

func (protocol *Protocol) handleLog(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.logHandler == nil {
		globals.Logger.Warning("RemoteLogDevice::Log not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strLine, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.logHandler(fmt.Errorf("Failed to read strLine from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.logHandler(nil, packet, callID, strLine)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
