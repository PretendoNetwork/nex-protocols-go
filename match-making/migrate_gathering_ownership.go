// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MigrateGatheringOwnership sets the MigrateGatheringOwnership handler function
func (protocol *Protocol) MigrateGatheringOwnership(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []*nex.PID, participantsOnly bool) uint32) {
	protocol.migrateGatheringOwnershipHandler = handler
}

func (protocol *Protocol) handleMigrateGatheringOwnership(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.migrateGatheringOwnershipHandler == nil {
		globals.Logger.Warning("MatchMaking::MigrateGatheringOwnership not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstPotentialNewOwnersID, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	participantsOnly, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.migrateGatheringOwnershipHandler(fmt.Errorf("Failed to read participantsOnly from parameters. %s", err.Error()), packet, callID, 0, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.migrateGatheringOwnershipHandler(nil, packet, callID, gid, lstPotentialNewOwnersID, participantsOnly)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
