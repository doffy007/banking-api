package model

import "time"

type BankBalace struct {
	ID                int               `json:"id" gorm:"primaryKey"`
	Balance           int               `json:"balance"`
	BalanceArchive    int               `json:"balance_archive"`
	Code              string            `json:"code"`
	Enable            bool              `json:"enable"`
	BankBalaceHistory BankBalaceHistory `gorm:"foreignKey:BankBalanceID"`
	CreatedAt         time.Time         `json:"created_at"`
}

type BankBalaceHistory struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	BankBalanceID int       `json:"bank_balance_id"`
	BalanceBefore int       `json:"balance_before"`
	BalanceAfter  int       `json:"balance_after"`
	Activity      string    `json:"activity"`
	Type          string    `json:"type"`
	IP            string    `json:"ip"`
	Location      string    `json:"location"`
	UserAgent     string    `json:"user_agent"`
	Author        string    `json:"author"`
	CreatedAt     time.Time `json:"created_at"`
}
