package chat

import (
	"errors"
	"fmt"

	"RTF/storage"
	"RTF/types"
	"RTF/types/serializers"
	"RTF/utils"

	"github.com/gofrs/uuid"
)

const INSERT_MESSAGE = `
	INSERT INTO chat_messages (text, created_at, sender, recipient, message_status)
	VALUES ($1, $2, $3, $4, $5)
	`

// This function saves a post object to the DB
func SaveChatInDB(chat serializers.Message) error {
	msg_Status := false
	var session_id uuid.UUID

	for _, s := range types.Sessions {
		if s.User.Username == chat.Recipient {
			session_id = s.SessionID
			if types.Sessions[session_id].ChatPartnerID == chat.Sender {
				msg_Status = true
			}
		}
	}


	stmt, err := storage.DB_Conn.Prepare(INSERT_MESSAGE)
	if err != nil {
		return errors.Join(errors.New("error preparing SaveChatMessageInDB query"), err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(chat.Msg_Content, chat.Timestamp, chat.Sender, chat.Recipient, msg_Status); err != nil {
		return errors.Join(errors.New("error executing SaveChatMessageInDB query"), err)
	} else {
		utils.InfoConsoleLog(fmt.Sprintf("New chat/message created: \n{\n%s\n}\n", chat.To_String()))
	}

	return nil
}
