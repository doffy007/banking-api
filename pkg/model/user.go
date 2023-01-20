package model

import "time"

type User struct {
	ID          int           `json:"id" gorm:"primaryKey"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Status      bool          `json:"status" gorm:"default:true"`
	CreatedAt   time.Time     `json:"created_at"`
	UserBalance []UserBalance `gorm:"foreignKey:UserID"`
}

type UserBalance struct {
	ID                 int                `json:"id" gorm:"primaryKey"`
	UserID             int                `json:"user_id"`
	Balance            string             `json:"balance"`
	BalanceArchive     int                `json:"balance_archive"`
	UserBalanceHistory UserBalanceHistory `gorm:"foreignKey:UserBalanceID"`
	CreatedAt          time.Time          `json:"created_at"`
}

type UserBalanceHistory struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	UserBalanceID int       `json:"user_balance_id"`
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
