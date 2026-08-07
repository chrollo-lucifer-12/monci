package db

import "time"

type Pipeline struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	Repository string    `gorm:"not null"`
	CommitSHA  string    `gorm:"not null"`
	Ref        string    `gorm:"not null"`
	Status     string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type Job struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	PipelineID int64     `gorm:"not null"`
	Name       string    `gorm:"not null"`
	Status     string    `gorm:"not null"`
	RunnerID   *string   `gorm:"column:runner_id"`
	CreatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	StartedAt  *time.Time
	FinishedAt *time.Time

	Pipeline Pipeline `gorm:"foreignKey:PipelineID"`
}

type Step struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	JobID      int64  `gorm:"not null"`
	Name       string `gorm:"not null"`
	Command    string `gorm:"not null"`
	Status     string `gorm:"not null"`
	Output     string
	StartedAt  *time.Time
	FinishedAt *time.Time

	Job Job `gorm:"foreignKey:JobID"`
}
