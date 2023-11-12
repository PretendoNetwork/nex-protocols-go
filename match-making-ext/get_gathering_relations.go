// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetGatheringRelations sets the GetGatheringRelations handler function
func (protocol *Protocol) GetGatheringRelations(handler func(err error, packet nex.PacketInterface, callID uint32, id uint32, descr string) uint32) {
	protocol.getGatheringRelationsHandler = handler
}

func (protocol *Protocol) handleGetGatheringRelations(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getGatheringRelationsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetGatheringRelations not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getGatheringRelationsHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	descr, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.getGatheringRelationsHandler(fmt.Errorf("Failed to read descr from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getGatheringRelationsHandler(nil, packet, callID, id, descr)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
