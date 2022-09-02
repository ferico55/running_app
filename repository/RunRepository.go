package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ferico55/running_app/model"
)

func SaveRun(run model.Run, userId int64, context context.Context) error {
	db := openDBConnection()
	defer db.Close()

	route, err := json.Marshal(run.Routes)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(context, `INSERT INTO runs(user_id, location_name, date, duration, distance, total_steps, route) VALUES(?, ?, CURRENT_TIMESTAMP, ?, ?, ?, ?)`,
		userId, run.LocationName, run.Duration, run.Distance, run.TotalSteps, route)
	return err
}

func GetRunForUser(userId int64, context context.Context) ([]model.Run, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `
	SELECT id, location_name, date, duration, distance, total_steps from runs where user_id = ?
	`, userId)
	if err != nil {
		return []model.Run{}, err
	}

	defer rows.Close()
	var runs = []model.Run{}
	for rows.Next() {
		var run model.Run
		err = rows.StructScan(&run)
		if err != nil {
			return []model.Run{}, err
		}

		run.FormattedDate = run.Date.Format("2 Jan 2006")
		run.FormattedDuration = formatDuration(run.Duration)
		runs = append(runs, run)
	}
	return runs, nil
}

func formatDuration(duration int) string {
	second := duration % 60
	duration = duration / 60
	minute := duration % 60
	duration = duration / 60
	hour := duration % 60

	var str = ""
	if hour > 0 {
		str += strconv.Itoa(hour) + ":"
	}
	if minute > 0 && minute < 10 {
		str += "0" + strconv.Itoa(minute) + ":"
	} else if minute > 0 {
		str += strconv.Itoa(minute) + ":"
	}

	if second < 10 {
		str += "0" + strconv.Itoa(second)
	} else {
		str += strconv.Itoa(second)
	}

	return str
}

func GetRunById(id int64, context context.Context) (*model.Run, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `
	SELECT id, location_name, date, duration, distance, total_steps, route from runs where id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var run model.Run
		err = rows.StructScan(&run)
		if err != nil {
			return nil, err
		}

		routes := []model.Coordinate{}
		err = json.Unmarshal([]byte(run.RawRoute), &routes)
		if err != nil {
			return nil, err
		}
		run.Routes = routes
		run.FormattedDate = run.Date.Format("2 Jan 2006")
		run.FormattedDuration = formatDuration(run.Duration)

		return &run, nil
	}
	return nil, nil
}

func GetLastRunForUser(userId int64, context context.Context) (*model.Run, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `
	SELECT id, location_name, date, duration, distance, total_steps, route from runs where user_id = ?
	ORDER BY date DESC
	`, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		var run model.Run
		err = rows.StructScan(&run)
		if err != nil {
			return nil, err
		}

		routes := []model.Coordinate{}
		err = json.Unmarshal([]byte(run.RawRoute), &routes)
		if err != nil {
			return nil, err
		}
		run.Routes = routes
		run.FormattedDate = run.Date.Format("2 Jan 2006")
		run.FormattedDuration = formatDuration(run.Duration)

		return &run, nil
	}
	return nil, nil
}

func GetRunFromUsers(users []model.User, context context.Context) ([]model.Run, error) {
	ids := []string{}
	for _, user := range users {
		ids = append(ids, strconv.FormatInt(user.Id, 10))
	}
	in := strings.Join(ids, ",")
	fmt.Print(in)

	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, "SELECT id, location_name, date, duration, distance, total_steps from runs where user_id IN ("+in+")")
	if err != nil {
		return []model.Run{}, err
	}

	defer rows.Close()
	var runs = []model.Run{}
	for rows.Next() {
		var run model.Run
		err = rows.StructScan(&run)
		if err != nil {
			return []model.Run{}, err
		}

		fmt.Println(run.Id)
		run.FormattedDate = run.Date.Format("2 Jan 2006")
		run.FormattedDuration = formatDuration(run.Duration)
		runs = append(runs, run)
	}
	return runs, nil
}
