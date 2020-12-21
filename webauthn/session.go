package webauthn

import (
	"encoding/binary"
	"github.com/duo-labs/webauthn/protocol"
)

// SessionData is the data that should be stored by the Relying Party for
// the duration of the web authentication ceremony
type SessionData struct {
	Challenge            string                               `json:"challenge"`
	UserID               []byte                               `json:"user_id"`
	AllowedCredentialIDs [][]byte                             `json:"allowed_credentials,omitempty"`
	UserVerification     protocol.UserVerificationRequirement `json:"userVerification"`
}

// UserID returns the user's ID (so the corresponsing User struct can be found); returns 0 on error
func (d SessionData) GetUserID() uint64 {
	//return binary.BigEndian.Uint64(d.UserID)
	id, bytes := binary.Uvarint(d.UserID)
	if bytes <= 0 {
		return 0
	}
	return id
	
	/*
	buf := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(buf, uint64(u.id))
	return buf
	*/
}
