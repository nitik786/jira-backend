// controllers/issues.go

package controllers

import (
	"jira-mock-app/app/models"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Issues struct {
	*revel.Controller
	DB *gorm.DB // You can access the DB instance here
}

// List all issues
func (c Issues) Index() revel.Result {
    // Check authorization as needed

    // Fetch the list of issues from the database
    var issues []models.Issue
    if err := c.DB.Find(&issues).Error; err != nil {
        return c.RenderJSON(map[string]interface{}{
            "error": "Failed to fetch issues",
        })
    }

    return c.RenderJSON(issues)
}

// Display a form to create a new issue
func (c Issues) New() revel.Result {
    // Check authorization as needed

    // You can render a new issue creation form here

    return c.Render()
}

// View a specific issue
func (c Issues) Show(issueID uint) revel.Result {
    // Check authorization and perform validations as needed

    // Fetch the issue from the database
    var issue models.Issue
    if err := c.DB.Where("id = ?", issueID).First(&issue).Error; err != nil {
        return c.RenderJSON(map[string]interface{}{
            "error": "Issue not found",
        })
    }

    return c.RenderJSON(issue)
}

// Display a form to edit a specific issue
func (c Issues) Edit(issueID uint) revel.Result {
    // Check authorization and perform validations as needed

    // Fetch the issue from the database
    var issue models.Issue
    if err := c.DB.Where("id = ?", issueID).First(&issue).Error; err != nil {
        return c.RenderJSON(map[string]interface{}{
            "error": "Issue not found",
        })
    }

    // You can render an edit form here

    return c.Render()
}


// Create a new issue
func (c Issues) Create(title, description, status, priority string, assigneeID, reporterID, projectID, sprintID uint, dueDate time.Time) revel.Result {
	// Check authorization and perform validations as needed

	// Create the issue in the database
	if err := models.CreateIssue(c.DB, title, description, status, priority, assigneeID, reporterID, projectID, sprintID, dueDate); err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Failed to create issue",
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Issue created successfully",
	})
}

// List issues
func (c Issues) List() revel.Result {
	// Check authorization as needed

	// Fetch the list of issues from the database
	var issues []models.Issue
	if err := c.DB.Find(&issues).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Failed to fetch issues",
		})
	}

	return c.RenderJSON(issues)
}

// View a specific issue
func (c Issues) View(issueID uint) revel.Result {
	// Check authorization and perform validations as needed

	// Fetch the issue from the database
	var issue models.Issue
	if err := c.DB.Where("id = ?", issueID).First(&issue).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Issue not found",
		})
	}

	return c.RenderJSON(issue)
}

// Update an issue
func (c Issues) Update(issueID uint, title, description, status, priority string, assigneeID, reporterID uint, dueDate time.Time) revel.Result {
	// Check authorization and perform validations as needed

	// Fetch the issue from the database
	var issue models.Issue
	if err := c.DB.Where("id = ?", issueID).First(&issue).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Issue not found",
		})
	}

	// Update the issue fields
	issue.Title = title
	issue.Description = description
	issue.Status = status
	issue.Priority = priority
	issue.AssigneeID = assigneeID
	issue.ReporterID = reporterID
	issue.DueDate = dueDate

	// Save the updated issue to the database
	if err := c.DB.Save(&issue).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Failed to update issue",
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Issue updated successfully",
	})
}

// Delete an issue
func (c Issues) Delete(issueID uint) revel.Result {
	// Check authorization and perform validations as needed

	// Fetch the issue from the database
	var issue models.Issue
	if err := c.DB.Where("id = ?", issueID).First(&issue).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Issue not found",
		})
	}

	// Delete the issue
	if err := c.DB.Delete(&issue).Error; err != nil {
		return c.RenderJSON(map[string]interface{}{
			"error": "Failed to delete issue",
		})
	}

	return c.RenderJSON(map[string]interface{}{
		"message": "Issue deleted successfully",
	})
}
