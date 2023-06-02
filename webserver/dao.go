package webserver

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type MemberReq struct {
	Token string `json:"token"`
}

type MemberRes struct {
	TodayBet int64 `json:"todayBet"`
	TodayPay int64 `json:"todayPay"`
	Status   int   `json:"status"`
	TotalBet int64 `json:"totalBet"`
	TotalPay int64 `json:"totalPay"`
}
