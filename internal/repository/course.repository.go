package repository

import (
	"database/sql"
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
)

func (c *seenRepository) AddCourse(data *models.CourseDTO) (*models.Course, *types.Error) {
	query := `INSERT INTO courses (id, garrison_id, name, count, priority, class_time) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := c.db.Exec(query, data.Id, data.GarrisonId, data.Name, data.Count, data.Priority, data.ClassTime)
	if err != nil {
		fmt.Println(err)
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 145")
	}
	return &models.Course{
		Id:         data.Id,
		Name:       data.Name,
		GarrisonId: data.GarrisonId,
		Count:      data.Count,
		Priority:   data.Priority,
		ClassTime:  data.ClassTime,
	}, nil
}

func (c *seenRepository) GetAllCoursesByGarrisonId(garrisonId int32) ([]models.Course, *types.Error) {
	query := `SELECT id, garrison_id, name, count, priority, class_time FROM courses WHERE garrison_id = $1`
	rows, err := c.db.Query(query, garrisonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("درسی پیدا نشد. کد خطا 146")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 147")
	}

	defer rows.Close()
	courses := make([]models.Course, 0)
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(
			&course.Id,
			&course.GarrisonId,
			&course.Name,
			&course.Count,
			&course.Priority,
			&course.ClassTime); err != nil {
			fmt.Println(err)
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 148")
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 149")
	}
	return courses, nil
}

func (c *seenRepository) DeleteCourseByGarrisonIdAndCourseId(garrisonId int32, courseId string) *types.Error {
	tx, err := c.db.Begin()
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 150")
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	queries := []string{
		`DELETE FROM classes WHERE garrison_id = $1 AND course_id = $2`,
		`DELETE FROM courses WHERE garrison_id = $1 AND id = $2`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, garrisonId, courseId); err != nil {
			fmt.Println(err)
			return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 151")
		}
	}

	return nil
}

func (c *seenRepository) CheckCourseByCourseId(courseId string) (bool, *types.Error) {
	query := `SELECT 1 FROM courses WHERE id = $1`
	var result int32
	err := c.db.QueryRow(query, courseId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 159")
	}
	return true, nil
}
