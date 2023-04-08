package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateURLs sets the UpdateURLs handler function
func (protocol *SecureConnectionProtocol) UpdateURLs(handler func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL)) {
	protocol.UpdateURLsHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleUpdateURLs(packet nex.PacketInterface) {
	if protocol.UpdateURLsHandler == nil {
		globals.Logger.Warning("SecureConnection::UpdateURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[SecureConnection::UpdateURLs] Data missing list length")
		go protocol.UpdateURLsHandler(err, client, callID, make([]*nex.StationURL, 0))
		return
	}

	stationURLCount := parametersStream.ReadUInt32LE()
	stationUrls := make([]*nex.StationURL, 0)

	for i := 0; i < int(stationURLCount); i++ {
		stationString, err := parametersStream.ReadString()

		if err != nil {
			go protocol.UpdateURLsHandler(err, client, callID, stationUrls)
			return
		}

		station := nex.NewStationURL(stationString)
		stationUrls = append(stationUrls, station)
	}

	go protocol.UpdateURLsHandler(nil, client, callID, stationUrls)
}
