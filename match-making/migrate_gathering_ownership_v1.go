// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleMigrateGatheringOwnershipV1(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.MigrateGatheringOwnershipV1 == nil {
		globals.Logger.Warning("MatchMaking::MigrateGatheringOwnershipV1 not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.MigrateGatheringOwnershipV1(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstPotentialNewOwnersID := types.NewList[*types.PID]()
	lstPotentialNewOwnersID.Type = types.NewPID(0)
	err = lstPotentialNewOwnersID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.MigrateGatheringOwnershipV1(fmt.Errorf("Failed to read lstPotentialNewOwnersID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.MigrateGatheringOwnershipV1(nil, packet, callID, gid, lstPotentialNewOwnersID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
