// Package protocol implements the Remote Log Device protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleLog(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.Log == nil {
		globals.Logger.Warning("RemoteLogDevice::Log not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strLine, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.Log(fmt.Errorf("Failed to read strLine from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.Log(nil, packet, callID, strLine)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
