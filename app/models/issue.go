package models

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Issue struct {
    gorm.Model

    Title       string `gorm:"not null"`
    Description string
    Status      string `gorm:"not null"`
    Priority    string `gorm:"not null"`
    AssigneeID  uint   // You can add a foreign key relationship if needed
    ReporterID  uint   // You can add a foreign key relationship if needed
    DueDate     time.Time
    ProjectID   uint // Foreign key to associate with a project
    SprintID    uint // Foreign key to associate with a sprint (if applicable)
}

// Define functions to retrieve, update, or delete issues as needed

func CreateIssue(db *gorm.DB, title, description, status, priority string, assigneeID, reporterID, projectID, sprintID uint, dueDate time.Time) error {
    issue := Issue{
        Title:       title,
        Description: description,
        Status:      status,
        Priority:    priority,
        AssigneeID:  assigneeID,
        ReporterID:  reporterID,
        DueDate:     dueDate,
        ProjectID:   projectID,
        SprintID:    sprintID,
    }
    return db.Create(&issue).Error
}

