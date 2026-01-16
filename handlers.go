package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile returns profile information
func GetProfile(c *gin.Context) {
	profile := map[string]interface{}{
		"name":        "Your Name",
		"title":       "Your Title/Position",
		"bio":         "Brief bio about yourself",
		"email":       "your.email@example.com",
		"phone":       "+62xxx",
		"location":    "Your Location",
		"avatar":      "/images/avatar.jpg",
		"social": map[string]string{
			"github":   "https://github.com/yourusername",
			"linkedin": "https://linkedin.com/in/yourusername",
			"twitter":  "https://twitter.com/yourusername",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    profile,
	})
}

// GetExperiences returns all work experiences
func GetExperiences(c *gin.Context) {
	experiences := []map[string]interface{}{
		{
			"id":          1,
			"company":     "Company Name",
			"position":    "Position Title",
			"location":    "Location",
			"startDate":   "2023-01",
			"endDate":     "Present",
			"description": "Job description and responsibilities",
			"technologies": []string{"Go", "React", "PostgreSQL"},
		},
		// Add more experiences
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    experiences,
	})
}

// GetExperience returns single experience by ID
func GetExperience(c *gin.Context) {
	id := c.Param("id")
	
	// TODO: Implement database query
	experience := map[string]interface{}{
		"id":          id,
		"company":     "Company Name",
		"position":    "Position Title",
		"location":    "Location",
		"startDate":   "2023-01",
		"endDate":     "Present",
		"description": "Detailed job description",
		"technologies": []string{"Go", "React", "PostgreSQL"},
		"achievements": []string{
			"Achievement 1",
			"Achievement 2",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    experience,
	})
}

// GetEducation returns education history
func GetEducation(c *gin.Context) {
	education := []map[string]interface{}{
		{
			"id":          1,
			"institution": "University Name",
			"degree":      "Bachelor of Science",
			"field":       "Computer Science",
			"startDate":   "2015",
			"endDate":     "2019",
			"gpa":         "3.8",
			"description": "Relevant coursework and achievements",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    education,
	})
}

// GetEducationDetail returns single education detail
func GetEducationDetail(c *gin.Context) {
	id := c.Param("id")
	
	education := map[string]interface{}{
		"id":          id,
		"institution": "University Name",
		"degree":      "Bachelor of Science",
		"field":       "Computer Science",
		"startDate":   "2015",
		"endDate":     "2019",
		"gpa":         "3.8",
		"description": "Detailed education information",
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    education,
	})
}

// GetSkills returns skills categorized
func GetSkills(c *gin.Context) {
	skills := map[string]interface{}{
		"technical": []map[string]interface{}{
			{"name": "Go", "level": 90},
			{"name": "JavaScript", "level": 85},
			{"name": "SvelteKit", "level": 80},
			{"name": "PostgreSQL", "level": 85},
		},
		"soft": []string{
			"Problem Solving",
			"Team Collaboration",
			"Communication",
		},
		"tools": []string{
			"Git",
			"Docker",
			"VS Code",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    skills,
	})
}

// GetProjects returns all projects
func GetProjects(c *gin.Context) {
	projects := []map[string]interface{}{
		{
			"id":          1,
			"title":       "Project Name",
			"description": "Project description",
			"image":       "/images/project1.jpg",
			"technologies": []string{"Go", "SvelteKit", "PostgreSQL"},
			"githubUrl":   "https://github.com/yourusername/project",
			"liveUrl":     "https://project-demo.com",
			"featured":    true,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    projects,
	})
}

// GetProject returns single project
func GetProject(c *gin.Context) {
	id := c.Param("id")
	
	project := map[string]interface{}{
		"id":          id,
		"title":       "Project Name",
		"description": "Detailed project description",
		"image":       "/images/project1.jpg",
		"technologies": []string{"Go", "SvelteKit", "PostgreSQL"},
		"githubUrl":   "https://github.com/yourusername/project",
		"liveUrl":     "https://project-demo.com",
		"features": []string{
			"Feature 1",
			"Feature 2",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    project,
	})
}

// SendContactMessage handles contact form submissions
func SendContactMessage(c *gin.Context) {
	var message struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Subject string `json:"subject"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	// TODO: Implement email sending or save to database
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Message received successfully",
	})
}
