package redis

// redis key

const (
	KeyPostTimeZSet  = "bluebell:post:time"  // 帖子以发帖时间
	KeyPostScoreZSet = "bluebell:post:score" // 帖子以投票分数
	KeyPostVotedZSet = "bluebell:post:voted" // 记录用户是否投票，投什么票
	// KeyCommunitySet 这个是 ZSet类型，懒得改了key 是 KeyCommunitySet + communityID，value 是帖子ID，score是 time 或 score
	KeyCommunitySet = "bluebell:post:community:" // 保存每个社区下帖子的id,后面跟着 time 或 score ，表示键的分数是time 或score
)
