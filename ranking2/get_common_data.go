// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommonData sets the GetCommonData handler function
func (protocol *Protocol) GetCommonData(handler func(err error, client *nex.Client, callID uint32, optionFlags uint32, principalID uint32, nexUniqueID uint64)) {
	protocol.getCommonDataHandler = handler
}

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	if protocol.getCommonDataHandler == nil {
		globals.Logger.Warning("Ranking2::GetCommonData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	optionFlags, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getCommonDataHandler(fmt.Errorf("Failed to read optionFlags from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	principalID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getCommonDataHandler(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getCommonDataHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), client, callID, 0, 0, 0)
		return
	}

	go protocol.getCommonDataHandler(nil, client, callID, optionFlags, principalID, nexUniqueID)
}