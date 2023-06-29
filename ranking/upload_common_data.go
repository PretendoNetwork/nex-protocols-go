package ranking

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCommonData sets the UploadCommonData handler function
func (protocol *RankingProtocol) UploadCommonData(handler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64)) {
	protocol.UploadCommonDataHandler = handler
}

func (protocol *RankingProtocol) handleUploadCommonData(packet nex.PacketInterface) {
	if protocol.UploadCommonDataHandler == nil {
		globals.Logger.Warning("Ranking::UploadCommonData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	commonData, err := parametersStream.ReadBuffer()
	if err != nil {
		go protocol.UploadCommonDataHandler(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.UploadCommonDataHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, nil, 0)
		return
	}

	go protocol.UploadCommonDataHandler(nil, client, callID, commonData, uniqueID)
}
