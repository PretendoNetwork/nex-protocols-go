package nexproto

import (
	nex "github.com/PretendoNetwork/nex-go"
)

// StreamIn is an abstraction of StreamIn from github.com/PretendoNetwork/nex-go
// Adds protocol-specific Structure list support
type StreamIn struct {
	*nex.StreamIn
}

// ReadListPersistentNotification reads a list of PersistentNotification structures
func (stream *StreamIn) ReadListPersistentNotification() ([]*PersistentNotification, error) {
	length := stream.ReadUInt32LE()
	persistentNotifications := make([]*PersistentNotification, 0)

	for i := 0; i < int(length); i++ {
		persistentNotificationStructureInterface, err := stream.ReadStructure(NewPersistentNotification())
		if err != nil {
			return nil, err
		}

		persistentNotification := persistentNotificationStructureInterface.(*PersistentNotification)
		persistentNotifications = append(persistentNotifications, persistentNotification)
	}

	return persistentNotifications, nil
}

// ReadListStationURL reads a list of PersistentNotification structures
func (stream *StreamIn) ReadListStationURL() ([]*nex.StationURL, error) {
	length := stream.ReadUInt32LE()
	stationUrls := make([]*nex.StationURL, 0)

	for i := 0; i < int(length); i++ {
		stationString, err := stream.ReadString()

		if err != nil {
			return nil, err
		}

		station := nex.NewStationURL(stationString)
		stationUrls = append(stationUrls, station)
	}

	return stationUrls, nil
}

// NewStreamIn returns a new nexproto output stream
func NewStreamIn(data []byte, server *nex.Server) *StreamIn {
	return &StreamIn{
		StreamIn: nex.NewStreamIn(data, server),
	}
}
