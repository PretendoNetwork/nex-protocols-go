// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetDeletionReason sets the SetDeletionReason handler function
func (protocol *Protocol) SetDeletionReason(handler func(err error, client *nex.Client, callID uint32, dataIDLst []uint64, deletionReason uint32) uint32) {
	protocol.setDeletionReasonHandler = handler
}

func (protocol *Protocol) handleSetDeletionReason(packet nex.PacketInterface) {
	if protocol.setDeletionReasonHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetDeletionReason not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.setDeletionReasonHandler(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	deletionReason, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.setDeletionReasonHandler(fmt.Errorf("Failed to read deletionReason from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.setDeletionReasonHandler(nil, client, callID, dataIDLst, deletionReason)
}
