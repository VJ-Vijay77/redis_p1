package db


type Leaderboard struct {
	Count int `json:"count"`
	User []*User
}

func(db *Database) GetLeaderboard() (*Leaderboard,error) {
	scores := db.Client.ZRangeWithScores(Ctx,"LKEY",0,-1)
	if scores == nil {
		return nil,ErrNil
	}

	count := len(scores.Val())
	users := make([]*User,count)
	for i,member := range scores.Val() {
		users[i] = &User{
			Username: member.Member.(string),
			Points: int(member.Score),
			Rank: i,
		}
	}

	leaderboard := &Leaderboard{
		Count: count,
		User: users,
	}

	return leaderboard,nil

}