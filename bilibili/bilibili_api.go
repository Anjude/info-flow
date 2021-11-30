package bilibili

import "InfoFlow/utils"

const (
	pgcUrl = "https://api.bilibili.com/pgc/web/rank/list?day=3&season_type=1"
)

// GetPgcRank 获取知乎热榜数据
func GetPgcRank() (*PgcRank, error) {
	res, err := utils.GetReq(pgcUrl, "https://www.bilibili.com/")
	if err != nil {
		return nil, nil
	}
	pgcRank := PgcRank{}
	err = res.ToJSON(&pgcRank)
	if err != nil {
		return nil, err
	}

	return &pgcRank, nil
}
