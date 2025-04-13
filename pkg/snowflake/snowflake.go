package snowflake

import (
	"time"

	"github.com/sony/sonyflake"

	"github.com/woxQAQ/upload-server/pkg/types"
)

var snowflake *sonyflake.Sonyflake

func InitSnowFlake(cfg *types.AppConfig) {
	set := sonyflake.Settings{
		StartTime: time.Unix(1483228800, 0),
		MachineID: func() (uint16, error) {
			return uint16(cfg.SnowFlakeMachineId), nil
		},
	}

	snowflake = sonyflake.NewSonyflake(set)
	_, err := snowflake.NextID()
	if err != nil {
		panic(err)
	}
}

func GetSnowFlakeID() (id uint64, err error) {
	if snowflake == nil {
		panic("snowflake id generator not init")
	}
	return snowflake.NextID()
}
