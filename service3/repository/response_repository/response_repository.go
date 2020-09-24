package response_repository

import (
	service1 "github.com/egsam98/microservice/service1/proto/service1"
	"github.com/egsam98/microservice/service3/db"
	service3 "github.com/egsam98/microservice/service3/proto/service3"
	"time"
)

const insertQuery = `
	INSERT INTO responses (serial_num, date_time, temp1, temp2, height, latitude, logitude, status)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`
const selectWhereSerialNumQuery = `SELECT date_time, temp1, temp2, height, latitude, logitude, status
	FROM responses WHERE serial_num = ?`

func Create(res *service1.Response) error {
	p := res.Packet
	_, err := db.DB.Exec(insertQuery,
		res.SerialNum, p.Datetime, p.Temp1, p.Temp2, p.Height, p.Latitude, p.Logitude, p.Status)
	if err != nil {
		return err
	}
	return nil
}

func FindBySerialNum(serialNum *service3.SerialNum) (*service3.Response, error) {
	rows, err := db.DB.Query(selectWhereSerialNumQuery, serialNum.Value)
	if err != nil {
		return nil, err
	}

	res := service3.Response{}
	res.Values = &service3.Values{}

	if !rows.Next() {
		return nil, nil
	}

	var timestamp string

	err = rows.Scan(
		&timestamp,
		&res.Values.Temp1,
		&res.Values.Temp2,
		&res.Values.Height,
		&res.Values.Latitude,
		&res.Values.Logitude,
		&res.Values.Status,
	)

	if res.Ts, err = convertToUnix(timestamp); err != nil {
		return nil, err
	}

	return &res, err
}

func convertToUnix(timestamp string) (int64, error) {
	result, err := time.Parse("02-01-2006,15:04:05", timestamp)
	if err != nil {
		return 0, err
	}
	return result.Unix(), nil
}
