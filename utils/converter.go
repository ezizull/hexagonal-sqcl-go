package utils

import (
	"database/sql"
	"strconv"
)

func ConvertToNullInt32(activityGroupIDStr string) sql.NullInt32 {
	activityGroupID, err := strconv.ParseInt(activityGroupIDStr, 10, 64)
	if err != nil || activityGroupID == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: int32(activityGroupID),
		Valid: true,
	}
}
