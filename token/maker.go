/*
package token contains methods required to create jwt tokens
*/
package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
