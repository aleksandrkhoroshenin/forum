package database

import (
	"forum/src/dicts/models"
	"github.com/jackc/pgx"
)

type ThreadDataManager interface {
	CreateThreadDB(user *models.Thread) ([]*models.Thread, error)
}

func CreateThreadInstance(conn *pgx.ConnPool) ThreadDataManager {
	return service{
		conn: conn,
	}
}

func (s service) CreateThreadDB(thread *models.Thread) (threads []*models.Thread, err error) {
	//rows, err := s.conn.Exec(
	//	createThreadScript,
	//	&thread.Nickname,
	//	&thread.Fullname,
	//	&thread.Email,
	//	&thread.About,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if rows.RowsAffected() == 0 { // пользователь уже есть
	//	thread := models.Thread{}
	//	err := s.conn.QueryRow(
	//		getThreadByNicknameOrEmailScript, &thread.Nickname, &thread.Email).Scan(&thread)
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return threads, ThreadIsExist
	//}

	return nil, nil
}
