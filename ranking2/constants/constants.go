// Package constants provides all the constants, enums, etc. for the Ranking2 protocol
package constants

const (
	// MaxBinaryDataLength represents the max size any binary data can be
	MaxBinaryDataLength int = 100

	// MaxChartGetLength represents the max number of charts that GetRankingCharts can handle
	MaxChartGetLength int = 20

	// MaxPutMultiScores represents the max number of scores that can be created at once
	MaxPutMultiScores int = 20

	// MaxRankingLength represents the max "Length" when getting rankings
	MaxRankingLength int = 100

	// MaxUsernameLength represents the max size of Ranking2CommonData.UserName
	MaxUsernameLength int = 16

	// ScoreOrderAsc indicates that ScoreOrder is ascending
	ScoreOrderAsc bool = true

	// ScoreOrderDesc indicates that ScoreOrder is descending
	ScoreOrderDesc bool = false
)
