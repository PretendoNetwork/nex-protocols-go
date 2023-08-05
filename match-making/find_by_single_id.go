// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindBySingleID sets the FindBySingleID handler function
func (protocol *Protocol) FindBySingleID(handler func(err error, client *nex.Client, callID uint32, id uint32) uint32) {
	protocol.findBySingleIDHandler = handler
}

func (protocol *Protocol) handleFindBySingleID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findBySingleIDHandler == nil {
		globals.Logger.Warning("MatchMaking::FindBySingleID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	id, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findBySingleIDHandler(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findBySingleIDHandler(nil, client, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
