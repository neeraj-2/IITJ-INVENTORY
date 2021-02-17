package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User Model
type User struct {
	gorm.Model
	Name     		string		`gorm:"not null"`
	Email    		string		`gorm:"primaryKey;not null"`
	Password 		string		`gorm:"not null"`
}

type Items struct {
	gorm.Model
	ItemKey      	string		`gorm:"primaryKey;not null"`		// identifier or ItemKey
	Society		    string		`gorm:"not null"`		// society or society key
	Details		    string		
	Quantity		int			`gorm:"not null"`
	Availabe		int			`gorm:"not null"`
}

type Societies struct {
	gorm.Model
	Email	 		string		`gorm:"not null"`
	Password		string		`gorm:"not null"`
	Details			string		
}

type Issued struct {
	gorm.Model
	Itemkey     	string		`gorm:"primaryKey;not null"`
	UserKey			string		`gorm:"not null"`		// UserKey or user email??
	IssueDate    	time.Time	`gorm:"not null"`
	DueDate			time.Time	`gorm:"not null"`
	Approved		string		`gorm:"not null"`
	Denied			string		`gorm:"not null"`
	Purpose			string		`gorm:"not null"`
}

type Inbound struct {
	gorm.Model
	DateOfInbound	time.Time	`gorm:"not null"`
	DateOfPurchase	time.Time	`gorm:"not null"`
	Bill			string		`gorm:"not null"`		// ??? kya h ye 
	Society			string		`gorm:"not null"`
	Price			int			`gorm:"not null"`
	Quantity		int			`gorm:"not null"`
	ItemKey			string		`gorm:"primaryKey;not null"`
}

type Defective struct {
	gorm.Model
	DateOfDest		time.Time	`gorm:"not null"`
	Quantity		string		`gorm:"not null"`
	ItemKey			string		`gorm:"primaryKey;not null"`
}