package chat

import (
	"errors"
	"fmt"

	"RTF/storage"
	"RTF/types"
	"RTF/types/serializers"
)

// EXAMPLE CODE

const (
	GET_CHAT = `SELECT *
	FROM (
		SELECT *
		FROM chat_messages
		WHERE (recipient = ? AND sender = ?) OR (recipient = ? AND sender = ?)
		ORDER BY id DESC
		LIMIT 10
	) AS last_10
	ORDER BY id ASC;
	;`

	UPDATE_MESSAGE_STATUS = `UPDATE chat_messages SET message_status = 1 WHERE (sender = ? and recipient = ?);`

	GET_LATEST_DMS = `SELECT users.user_name, users.user_id,chat_messages.message_status,chat_messages.sender	, MAX(chat_messages.created_at) AS last_message_time
						FROM users
						LEFT JOIN chat_messages ON (users.user_name = chat_messages.sender OR users.user_name = chat_messages.recipient)
						WHERE chat_messages.sender = ? OR chat_messages.recipient = ?
						GROUP BY users.user_name
						ORDER BY last_message_time DESC;`

	LOAD_MESSAGES_IN_BETWEEN = `SELECT id, text, created_at, sender, recipient
	FROM (
		SELECT *
		FROM chat_messages
		WHERE ((recipient = ? AND sender = ?) OR (recipient = ? AND sender = ?)) AND id < ?
		ORDER BY id DESC
		LIMIT 10
	) AS last_10
	ORDER BY id ASC;
	;`
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

		if err := rows.Scan(&chat_message.Id, &chat_message.Msg_Content, &chat_message.Timestamp, &chat_message.Sender, &chat_message.Recipient, &chat_message.Msg_Status); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}
		chat_messages = append(chat_messages, chat_message)
	}

	if len(chat_messages) > 0 {
		stmt, err = storage.DB_Conn.Prepare(UPDATE_MESSAGE_STATUS)

		if err != nil {
			return nil, errors.Join(types.ErrPrepare, err)
		}
		defer stmt.Close()
		fmt.Println(requested_user_name, user_name)
		result, err := stmt.Exec(requested_user_name, user_name)

		if err != nil {
			return nil, errors.Join(types.ErrExec, err)
		}

		// Check the number of affected rows
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return nil, errors.Join(types.ErrExec, err)
		}

		if rowsAffected == 0 {
			fmt.Println("no rows updated")
		}
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

		if err := rows.Scan(&dm_user.Username, &dm_user.ID, &dm_user.Msg_Status, &dm_user.Msg_sender, &dm_user.LastMessageTime); err != nil {
			return nil, errors.Join(types.ErrScan, err)
		}

		dm_users = append(dm_users, dm_user)

	}

	return dm_users, nil
}

/* A function to Load_Messages in between begin_id and end_id */
func Load_Messages(user_name string, requested_user_name string, begin_id int) ([]serializers.Message, error) {
	chat_messages := []serializers.Message{}
	stmt, err := storage.DB_Conn.Prepare(LOAD_MESSAGES_IN_BETWEEN)
	if err != nil {
		return nil, errors.Join(types.ErrPrepare, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_name, requested_user_name, requested_user_name, user_name, begin_id)
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
