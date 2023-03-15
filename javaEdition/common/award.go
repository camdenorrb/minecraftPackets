package common

type AwardStatistics []AwardStatistic

type AwardStatistic struct {
	CategoryID  VarInt
	StatisticID VarInt
	Value       VarInt
}
