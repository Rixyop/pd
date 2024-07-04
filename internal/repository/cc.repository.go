package repository

import (
	"database/sql"
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
)

func (c *seenRepository) DeleteCourseCoachByCourseId(courseId string) *types.Error {
	query := `DELETE FROM course_coaches WHERE course_id = $1`
	_, err := c.db.Exec(query, courseId)
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 152")
	}
	return nil
}

func (c *seenRepository) CheckCourseCoachByCourseIdAndCoachId(courseId string, coachId string) (bool, *types.Error) {
	query := `SELECT 1 FROM course_coaches WHERE course_id = $1 AND coach_id = $2`
	var result int32
	err := c.db.QueryRow(query, courseId, coachId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 153")
	}
	return true, nil
}

func (c *seenRepository) AddCourseCoach(courseId string, coachId string, garrisonId int32) *types.Error {
	query := `INSERT INTO course_coaches (course_id, coach_id, garrison_id) VALUES ($1,$2,$3)`

	_, err := c.db.Exec(query, courseId, coachId, garrisonId)
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 154")
	}

	return nil
}
func (c *seenRepository) GetAllCourseCoachesByGarrisonId(garrisonId int32) ([]models.CourseCoach, *types.Error) {
	query := `SELECT course_id, coach_id, garrison_id FROM course_coaches WHERE garrison_id = $1`
	rows, err := c.db.Query(query, garrisonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("درس/مربی پیدا نشد. کد خطا 155")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 156")
	}

	defer rows.Close()
	courseCoaches := make([]models.CourseCoach, 0)
	for rows.Next() {
		var courseCoach models.CourseCoach
		if err := rows.Scan(
			&courseCoach.CourseId,
			&courseCoach.CoachId,
			&courseCoach.GarrisonId); err != nil {
			fmt.Println(err)
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 157")
		}
		courseCoaches = append(courseCoaches, courseCoach)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 158")
	}
	return courseCoaches, nil
}
