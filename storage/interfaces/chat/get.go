package chat

import (
	"RTF/storage"
	"RTF/types"
	"RTF/types/serializers"
	"errors"
)

// EXAMPLE CODE

const (
	GET_CHAT = `SELECT * FROM chat_messages
				WHERE (recipient = ? and sender = ? or recipient = ? and sender = ?)
				ORDER BY id`

	GET_LATEST_DMS = `SELECT users.user_name, MAX(chat_messages.created_at) AS last_message_time
						FROM users
						LEFT JOIN chat_messages ON (users.user_name = chat_messages.sender OR users.user_name = chat_messages.recipient)
						WHERE chat_messages.sender = ? OR chat_messages.recipient = ?
						GROUP BY users.user_name
						ORDER BY last_message_time DESC;
						`
)

func Get_chat(user_name string, requested_user_name string) ([]serializers.Message, error) {

	chat_messages := []serializers.Message{}
	stmt, err := storage.DB_Conn.Prepare(GET_CHAT)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_name, requested_user_name, requested_user_name, user_name)
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}
	defer rows.Close()

	for rows.Next() {
		chat_message := serializers.Message{}

		if err := rows.Scan(&chat_message.Id, &chat_message.Msg_Content, &chat_message.Timestamp, &chat_message.Sender, &chat_message.Recipient); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		chat_messages = append(chat_messages, chat_message)
	}

	return chat_messages, nil
}

func Get_Users_By_Last_Message(user_name string) ([]serializers.DMs_User, error) {

	dm_users := []serializers.DMs_User{}
	stmt, err := storage.DB_Conn.Prepare(GET_LATEST_DMS)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_name, user_name)
	if err != nil {
		return nil, errors.Join(types.ErrExec, err)
	}
	defer rows.Close()

	for rows.Next() {
		dm_user := serializers.DMs_User{}

		if err := rows.Scan(&dm_user.Username, &dm_user.LastMessageTime); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		dm_users = append(dm_users, dm_user)
	}

	return dm_users, nil
}
