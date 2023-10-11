package common

// Version is a type that represents the version of the protocol.
type Version uint16

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	Version_1_7_10 Version = 5
	Version_1_8    Version = 47
	Version_1_9    Version = 107
	Version_1_9_1  Version = 108
	Version_1_9_2  Version = 109
	Version_1_9_3  Version = 110
	Version_1_9_4  Version = 111
	Version_1_10   Version = 210
	Version_1_11   Version = 315
	Version_1_11_1 Version = 316
	Version_1_12   Version = 335
	Version_1_12_1 Version = 338
	Version_1_12_2 Version = 340
	Version_1_13   Version = 393
	Version_1_13_1 Version = 401
	Version_1_13_2 Version = 404
	Version_1_14   Version = 477
	Version_1_14_1 Version = 480
	Version_1_14_2 Version = 485
	Version_1_14_3 Version = 490
	Version_1_14_4 Version = 498
	Version_1_15   Version = 573
	Version_1_15_1 Version = 575
	Version_1_15_2 Version = 578
	Version_1_16   Version = 735
	Version_1_16_1 Version = 736
	Version_1_16_2 Version = 751
	Version_1_16_3 Version = 753
	Version_1_16_4 Version = 754
	Version_1_17   Version = 755
	Version_1_17_1 Version = 756
	Version_1_18_1 Version = 757
	Version_1_18_2 Version = 758
	Version_1_19   Version = 759
	Version_1_19_1 Version = 760
	Version_1_19_2 Version = 760
	Version_1_19_3 Version = 761
	Version_1_19_4 Version = 762
	Version_1_20   Version = 763
	Version_1_20_1 Version = 763
	Version_1_20_2 Version = 764
)

//goland:noinspection GoUnusedConst,GoSnakeCaseUsage
const Version_Latest = Version_1_20_2

// TODO: Version.HasNettySupport

func (v Version) IsConfigStateSupported() bool {
	return v >= Version_1_20_2
}
