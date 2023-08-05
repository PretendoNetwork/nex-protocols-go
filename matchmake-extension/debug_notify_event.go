// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DebugNotifyEvent sets the DebugNotifyEvent handler function
func (protocol *Protocol) DebugNotifyEvent(handler func(err error, client *nex.Client, callID uint32, pid uint32, mainType uint32, subType uint32, param1 uint64, param2 uint64, stringParam string) uint32) {
	protocol.debugNotifyEventHandler = handler
}

func (protocol *Protocol) handleDebugNotifyEvent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.debugNotifyEventHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::DebugNotifyEvent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	mainType, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read mainType from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	subType, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read subType from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param1, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read param1 from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param2, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read param2 from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	stringParam, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.debugNotifyEventHandler(fmt.Errorf("Failed to read stringParam from parameters. %s", err.Error()), client, callID, 0, 0, 0, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.debugNotifyEventHandler(nil, client, callID, pid, mainType, subType, param1, param2, stringParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
