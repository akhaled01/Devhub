package chat

import (
	"errors"
	"fmt"

	"RTF/storage"
	"RTF/types/serializers"
	"RTF/utils"
)

const INSERT_MESSAGE = `
	INSERT INTO chat_messages (text, created_at, sender, recipient)
	VALUES ($1, $2, $3, $4)
	`

// This function saves a post object to the DB
func SaveChatInDB(chat serializers.Message) error {
	stmt, err := storage.DB_Conn.Prepare(INSERT_MESSAGE)
	if err != nil {
		return errors.Join(errors.New("error preparing SaveChatMessageInDB query"), err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(chat.Msg_Content, chat.Timestamp, chat.Sender, chat.Recipient); err != nil {
		return errors.Join(errors.New("error executing SaveChatMessageInDB query"), err)
	} else {
		utils.InfoConsoleLog(fmt.Sprintf("New chat/message created: \n{\n%s\n}\n", chat.To_String()))
	}

	return nil
}
