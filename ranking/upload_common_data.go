package ranking

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadCommonData sets the UploadCommonData handler function
func (protocol *RankingProtocol) UploadCommonData(handler func(err error, client *nex.Client, callID uint32, commonData []byte, uniqueID uint64)) {
	protocol.UploadCommonDataHandler = handler
}

func (protocol *RankingProtocol) HandleUploadCommonData(packet nex.PacketInterface) {
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
		go protocol.UploadCommonDataHandler(err, client, callID, nil, 0)
		return
	}

	uniqueID := parametersStream.ReadUInt64LE()

	go protocol.UploadCommonDataHandler(nil, client, callID, commonData, uniqueID)
}
