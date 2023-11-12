// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MigrateGatheringOwnershipV1 sets the MigrateGatheringOwnershipV1 handler function
func (protocol *Protocol) MigrateGatheringOwnershipV1(handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []uint32) uint32) {
	protocol.migrateGatheringOwnershipV1Handler = handler
}

func (protocol *Protocol) handleMigrateGatheringOwnershipV1(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.migrateGatheringOwnershipV1Handler == nil {
		globals.Logger.Warning("MatchMaking::MigrateGatheringOwnershipV1 not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.migrateGatheringOwnershipV1Handler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstPotentialNewOwnersID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.migrateGatheringOwnershipV1Handler(fmt.Errorf("Failed to read lstPotentialNewOwnersID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.migrateGatheringOwnershipV1Handler(nil, packet, callID, gid, lstPotentialNewOwnersID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
