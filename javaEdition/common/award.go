package common

import "github.com/camdenorrb/minecraftPackets/primitive"

type AwardStatistics []AwardStatistic

type AwardStatistic struct {
	CategoryID  primitive.VarInt
	StatisticID primitive.VarInt
	Value       primitive.VarInt
}
