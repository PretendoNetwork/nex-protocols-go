// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PutCommonData sets the PutCommonData handler function
func (protocol *Protocol) PutCommonData(handler func(err error, client *nex.Client, callID uint32, commonData *ranking2_types.Ranking2CommonData, nexUniqueID uint64)) {
	protocol.putCommonDataHandler = handler
}

func (protocol *Protocol) handlePutCommonData(packet nex.PacketInterface) {
	if protocol.putCommonDataHandler == nil {
		globals.Logger.Warning("Ranking2::PutCommonData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	commonData, err := parametersStream.ReadStructure(ranking2_types.NewRanking2CommonData())
	if err != nil {
		go protocol.putCommonDataHandler(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.putCommonDataHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.putCommonDataHandler(nil, client, callID, commonData.(*ranking2_types.Ranking2CommonData), nexUniqueID)
}