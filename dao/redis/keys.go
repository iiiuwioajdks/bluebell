package redis

// redis key

const (
	KeyPostTimeZSet  = "bluebell:post:time"  // 帖子以发帖时间
	KeyPostScoreZSet = "bluebell:post:score" // 帖子以投票分数
	KeyPostVotedZSet = "bluebell:post:voted" // 记录用户是否投票，投什么票
)
