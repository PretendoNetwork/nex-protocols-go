// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateApplicationBuffer sets the UpdateApplicationBuffer handler function
func (protocol *Protocol) UpdateApplicationBuffer(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, applicationBuffer []byte) uint32) {
	protocol.updateApplicationBufferHandler = handler
}

func (protocol *Protocol) handleUpdateApplicationBuffer(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateApplicationBufferHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateApplicationBuffer not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.updateApplicationBufferHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	applicationBuffer, err := parametersStream.ReadBuffer()
	if err != nil {
		errorCode = protocol.updateApplicationBufferHandler(fmt.Errorf("Failed to read applicationBuffer from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateApplicationBufferHandler(nil, packet, callID, gid, applicationBuffer)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
