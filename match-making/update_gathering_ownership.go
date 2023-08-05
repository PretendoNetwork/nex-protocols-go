// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateGatheringOwnership sets the UpdateGatheringOwnership handler function
func (protocol *Protocol) UpdateGatheringOwnership(handler func(err error, client *nex.Client, callID uint32, gid uint32, participantsOnly bool) uint32) {
	protocol.updateGatheringOwnershipHandler = handler
}

func (protocol *Protocol) handleUpdateGatheringOwnership(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateGatheringOwnershipHandler == nil {
		globals.Logger.Warning("MatchMaking::UpdateGatheringOwnership not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.updateGatheringOwnershipHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participantsOnly, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updateGatheringOwnershipHandler(fmt.Errorf("Failed to read participantsOnly from parameters. %s", err.Error()), client, callID, 0, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateGatheringOwnershipHandler(nil, client, callID, gid, participantsOnly)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
