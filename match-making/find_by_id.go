// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByID sets the FindByID handler function
func (protocol *Protocol) FindByID(handler func(err error, client *nex.Client, callID uint32, lstID []uint32) uint32) {
	protocol.findByIDHandler = handler
}

func (protocol *Protocol) handleFindByID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByIDHandler == nil {
		globals.Logger.Warning("MatchMaking::FindByID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.findByIDHandler(fmt.Errorf("Failed to read lstID from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByIDHandler(nil, client, callID, lstID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
