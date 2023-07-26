// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// MigrateGatheringOwnershipV1 sets the MigrateGatheringOwnershipV1 handler function
func (protocol *Protocol) MigrateGatheringOwnershipV1(handler func(err error, client *nex.Client, callID uint32, gid uint32, lstPotentialNewOwnersID []uint32)) {
	protocol.migrateGatheringOwnershipV1Handler = handler
}

func (protocol *Protocol) handleMigrateGatheringOwnershipV1(packet nex.PacketInterface) {
	if protocol.migrateGatheringOwnershipV1Handler == nil {
		globals.Logger.Warning("MatchMaking::MigrateGatheringOwnershipV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.migrateGatheringOwnershipV1Handler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	lstPotentialNewOwnersID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.migrateGatheringOwnershipV1Handler(fmt.Errorf("Failed to read lstPotentialNewOwnersID from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.migrateGatheringOwnershipV1Handler(nil, client, callID, gid, lstPotentialNewOwnersID)
}
