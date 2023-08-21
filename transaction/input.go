package transaction

import "crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount    int `json:"amount" binding:"required"`
	CampaigID int `json:"campaign_id" binding:"required"`
	User      user.User
}
