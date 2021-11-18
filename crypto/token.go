package crypto

import "time"


type Token struct{
	IssuedAt time.Time
	Username string
	Role string
}