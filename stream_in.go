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

// ReadListStationURL reads a list of StationURL structures
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

// ReadListDataStoreRateCustomRankingParam reads a list of DataStoreRateCustomRankingParam structures
func (stream *StreamIn) ReadListDataStoreRateCustomRankingParam() ([]*DataStoreRateCustomRankingParam, error) {
	length := stream.ReadUInt32LE()
	dataStoreRateCustomRankingParams := make([]*DataStoreRateCustomRankingParam, 0)

	for i := 0; i < int(length); i++ {
		dataStoreRateCustomRankingParam, err := stream.ReadStructure(NewDataStoreRateCustomRankingParam())

		if err != nil {
			return nil, err
		}

		dataStoreRateCustomRankingParams = append(dataStoreRateCustomRankingParams, dataStoreRateCustomRankingParam.(*DataStoreRateCustomRankingParam))
	}

	return dataStoreRateCustomRankingParams, nil
}

// ReadListDataStoreRatingInfoWithSlot reads a list of DataStoreRatingInfoWithSlot structures
func (stream *StreamIn) ReadListDataStoreRatingInfoWithSlot() ([]*DataStoreRatingInfoWithSlot, error) {
	length := stream.ReadUInt32LE()
	dataStoreRatingInfoWithSlots := make([]*DataStoreRatingInfoWithSlot, 0)

	for i := 0; i < int(length); i++ {
		dataStoreRatingInfoWithSlot, err := stream.ReadStructure(NewDataStoreRatingInfoWithSlot())

		if err != nil {
			return nil, err
		}

		dataStoreRatingInfoWithSlots = append(dataStoreRatingInfoWithSlots, dataStoreRatingInfoWithSlot.(*DataStoreRatingInfoWithSlot))
	}

	return dataStoreRatingInfoWithSlots, nil
}

// ReadListDataStoreGetCourseRecordParam reads a list of DataStoreGetCourseRecordParam structures
func (stream *StreamIn) ReadListDataStoreGetCourseRecordParam() ([]*DataStoreGetCourseRecordParam, error) {
	length := stream.ReadUInt32LE()
	dataStoreGetCourseRecordParams := make([]*DataStoreGetCourseRecordParam, 0)

	for i := 0; i < int(length); i++ {
		dataStoreGetCourseRecordParam, err := stream.ReadStructure(NewDataStoreGetCourseRecordParam())

		if err != nil {
			return nil, err
		}

		dataStoreGetCourseRecordParams = append(dataStoreGetCourseRecordParams, dataStoreGetCourseRecordParam.(*DataStoreGetCourseRecordParam))
	}

	return dataStoreGetCourseRecordParams, nil
}

// ReaListDataStoreGetMetaParam reads a list of DataStoreGetMetaParam structures
func (stream *StreamIn) ReaListDataStoreGetMetaParam() ([]*DataStoreGetMetaParam, error) {
	length := stream.ReadUInt32LE()
	dataStoreGetMetaParams := make([]*DataStoreGetMetaParam, 0)

	for i := 0; i < int(length); i++ {
		dataStoreGetMetaParam, err := stream.ReadStructure(NewDataStoreGetMetaParam())

		if err != nil {
			return nil, err
		}

		dataStoreGetMetaParams = append(dataStoreGetMetaParams, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
	}

	return dataStoreGetMetaParams, nil
}

// NewStreamIn returns a new nexproto output stream
func NewStreamIn(data []byte, server *nex.Server) *StreamIn {
	return &StreamIn{
		StreamIn: nex.NewStreamIn(data, server),
	}
}
