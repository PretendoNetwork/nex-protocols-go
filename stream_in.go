package nexproto

import (
	nex "github.com/PretendoNetwork/nex-go"
)

// Adds protocol-specific Structure list methods

// ReadListPersistentNotification reads a list of PersistentNotification structures
func ReadListPersistentNotification(stream *nex.StreamIn) ([]*PersistentNotification, error) {
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
func ReadListStationURL(stream *nex.StreamIn) ([]*nex.StationURL, error) {
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
func ReadListDataStoreRateCustomRankingParam(stream *nex.StreamIn) ([]*DataStoreRateCustomRankingParam, error) {
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
func ReadListDataStoreRatingInfoWithSlot(stream *nex.StreamIn) ([]*DataStoreRatingInfoWithSlot, error) {
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
func ReadListDataStoreGetCourseRecordParam(stream *nex.StreamIn) ([]*DataStoreGetCourseRecordParam, error) {
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
func ReaListDataStoreGetMetaParam(stream *nex.StreamIn) ([]*DataStoreGetMetaParam, error) {
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

// ReadListDataStoreRatingInitParamWithSlot reads a list of DataStoreRatingInitParamWithSlot structures
func ReadListDataStoreRatingInitParamWithSlot(stream *nex.StreamIn) ([]*DataStoreRatingInitParamWithSlot, error) {
	length := stream.ReadUInt32LE()
	dataStoreRatingInitParamWithSlots := make([]*DataStoreRatingInitParamWithSlot, 0)

	for i := 0; i < int(length); i++ {
		dataStoreRatingInitParamWithSlot, err := stream.ReadStructure(NewDataStoreRatingInitParamWithSlot())

		if err != nil {
			return nil, err
		}

		dataStoreRatingInitParamWithSlots = append(dataStoreRatingInitParamWithSlots, dataStoreRatingInitParamWithSlot.(*DataStoreRatingInitParamWithSlot))
	}

	return dataStoreRatingInitParamWithSlots, nil
}

// ReadListDataStoreRatingTarget reads a list of DataStoreRatingTarget structures
func ReadListDataStoreRatingTarget(stream *nex.StreamIn) ([]*DataStoreRatingTarget, error) {
	length := stream.ReadUInt32LE()
	dataStoreRatingTargets := make([]*DataStoreRatingTarget, 0)

	for i := 0; i < int(length); i++ {
		dataStoreRatingTarget, err := stream.ReadStructure(NewDataStoreRatingTarget())
		if err != nil {
			return nil, err
		}

		dataStoreRatingTargets = append(dataStoreRatingTargets, dataStoreRatingTarget.(*DataStoreRatingTarget))
	}

	return dataStoreRatingTargets, nil
}

// ReadListDataStoreRateObjectParam reads a list of DataStoreRateObjectParam structures
func ReadListDataStoreRateObjectParam(stream *nex.StreamIn) ([]*DataStoreRateObjectParam, error) {
	length := stream.ReadUInt32LE()
	dataStoreRateObjectParams := make([]*DataStoreRateObjectParam, 0)

	for i := 0; i < int(length); i++ {
		dataStoreRateObjectParam, err := stream.ReadStructure(NewDataStoreRateObjectParam())
		if err != nil {
			return nil, err
		}

		dataStoreRateObjectParams = append(dataStoreRateObjectParams, dataStoreRateObjectParam.(*DataStoreRateObjectParam))
	}

	return dataStoreRateObjectParams, nil
}

// ReadListBufferQueueParam reads a list of BufferQueueParam structures
func ReadListBufferQueueParam(stream *nex.StreamIn) ([]*BufferQueueParam, error) {
	length := stream.ReadUInt32LE()
	bufferQueueParams := make([]*BufferQueueParam, 0)

	for i := 0; i < int(length); i++ {
		bufferQueueParam, err := stream.ReadStructure(NewBufferQueueParam())
		if err != nil {
			return nil, err
		}

		bufferQueueParams = append(bufferQueueParams, bufferQueueParam.(*BufferQueueParam))
	}

	return bufferQueueParams, nil
}

// ReadListDataStoreKeyValue reads a list of DataStoreKeyValue structures
func ReadListDataStoreKeyValue(stream *nex.StreamIn) ([]*DataStoreKeyValue, error) {
	length := stream.ReadUInt32LE()
	dataStoreKeyValues := make([]*DataStoreKeyValue, 0)

	for i := 0; i < int(length); i++ {
		dataStoreKeyValue, err := stream.ReadStructure(NewDataStoreKeyValue())
		if err != nil {
			return nil, err
		}

		dataStoreKeyValues = append(dataStoreKeyValues, dataStoreKeyValue.(*DataStoreKeyValue))
	}

	return dataStoreKeyValues, nil
}

// ReadListDataStoreMetaInfo reads a list of DataStoreMetaInfo structures
func ReadListDataStoreMetaInfo(stream *nex.StreamIn) ([]*DataStoreMetaInfo, error) {
	length := stream.ReadUInt32LE()
	dataStoreMetaInfos := make([]*DataStoreMetaInfo, 0)

	for i := 0; i < int(length); i++ {
		dataStoreMetaInfo, err := stream.ReadStructure(NewDataStoreMetaInfo())
		if err != nil {
			return nil, err
		}

		dataStoreMetaInfos = append(dataStoreMetaInfos, dataStoreMetaInfo.(*DataStoreMetaInfo))
	}

	return dataStoreMetaInfos, nil
}

// ReadListDataStoreReplayPlayer reads a list of DataStoreReplayPlayer structures
func ReadListDataStoreReplayPlayer(stream *nex.StreamIn) ([]*DataStoreReplayPlayer, error) {
	length := stream.ReadUInt32LE()
	dataStoreReplayPlayers := make([]*DataStoreReplayPlayer, 0)

	for i := 0; i < int(length); i++ {
		dataStoreReplayPlayer, err := stream.ReadStructure(NewDataStoreReplayPlayer())
		if err != nil {
			return nil, err
		}

		dataStoreReplayPlayers = append(dataStoreReplayPlayers, dataStoreReplayPlayer.(*DataStoreReplayPlayer))
	}

	return dataStoreReplayPlayers, nil
}

// ReadListDataStoreFightingPowerScore reads a list of DataStoreFightingPowerScore structures
func ReadListDataStoreFightingPowerScore(stream *nex.StreamIn) ([]*DataStoreFightingPowerScore, error) {
	length := stream.ReadUInt32LE()
	dataStoreFightingPowerScores := make([]*DataStoreFightingPowerScore, 0)

	for i := 0; i < int(length); i++ {
		dataStoreFightingPowerScore, err := stream.ReadStructure(NewDataStoreFightingPowerScore())
		if err != nil {
			return nil, err
		}

		dataStoreFightingPowerScores = append(dataStoreFightingPowerScores, dataStoreFightingPowerScore.(*DataStoreFightingPowerScore))
	}

	return dataStoreFightingPowerScores, nil
}

// ReadListDataStorePostFightingPowerScoreParam reads a list of DataStorePostFightingPowerScoreParam structures
func ReadListDataStorePostFightingPowerScoreParam(stream *nex.StreamIn) ([]*DataStorePostFightingPowerScoreParam, error) {
	length := stream.ReadUInt32LE()
	dataStorePostFightingPowerScoreParams := make([]*DataStorePostFightingPowerScoreParam, 0)

	for i := 0; i < int(length); i++ {
		dataStorePostFightingPowerScoreParam, err := stream.ReadStructure(NewDataStorePostFightingPowerScoreParam())
		if err != nil {
			return nil, err
		}

		dataStorePostFightingPowerScoreParams = append(dataStorePostFightingPowerScoreParams, dataStorePostFightingPowerScoreParam.(*DataStorePostFightingPowerScoreParam))
	}

	return dataStorePostFightingPowerScoreParams, nil
}
