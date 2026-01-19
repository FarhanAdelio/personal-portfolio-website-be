package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile returns profile information
func GetProfile(c *gin.Context) {
	profile := map[string]interface{}{
		"name":     "Farhan Adelio Prayata",
		"title":    "Information Systems Student",
		"subtitle": "System Analyst | Business Analyst | Data Analyst",
		"tagline":  "Transforming complex problems into actionable system solutions",
		"bio":      "Mahasiswa Sistem Informasi dengan fokus pada analisis, perancangan sistem, dan pemahaman proses bisnis. Memiliki pengalaman magang di instansi pemerintahan dan passionate dalam mengubah data menjadi insight yang valuable untuk pengambilan keputusan.",
		"email":    "farhan.adelio@ui.ac.id",
		"phone":    "+62 877-7547-5299",
		"location": "Jakarta, Indonesia",
		"careerInterests": []string{
			"System Analyst",
			"Business Analyst",
			"Data Analyst",
			"Information Analyst",
			"Product Support",
		},
		"social": map[string]string{
			"github":   "https://github.com/farhanadelio",
			"linkedin": "https://linkedin.com/in/farhanadelio",
			"twitter":  "https://twitter.com/farhanadelio",
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
			"company":     "Dinas Pendidikan Provinsi DKI Jakarta",
			"position":    "IT Support & Data Analysis Intern",
			"type":        "Internship",
			"location":    "Jakarta, Indonesia",
			"startDate":   "2024-06",
			"endDate":     "2024-12",
			"description": "Membantu dalam pengolahan dan penyusunan data pendidikan, analisis kebutuhan sistem/aplikasi, serta mendukung dokumentasi sistem dan data. Berkolaborasi dengan tim teknis dan non-teknis untuk meningkatkan efisiensi proses bisnis.",
			"achievements": []string{
				"Menganalisis dan mengolah data mentah menjadi informasi yang actionable untuk pengambilan keputusan",
				"Membantu mapping business process dan workflow untuk digitalisasi sistem",
				"Membuat dokumentasi requirement analysis untuk pengembangan aplikasi internal",
				"Berkontribusi dalam komunikasi lintas departemen (IT, humas, data, manajemen)",
			},
			"technologies": []string{"SQL", "PostgreSQL", "Python", "Excel", "Data Analysis"},
			"learnings": []string{
				"Memahami cara kerja sistem pemerintahan dan alur birokrasi",
				"Pengalaman mengolah data dari mentah hingga siap pakai",
				"Pentingnya data accuracy untuk decision making",
				"Skill komunikasi dengan stakeholder non-teknis",
			},
		},
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
			"institution": "Universitas Indonesia",
			"degree":      "Bachelor of Information Systems",
			"field":       "Sistem Informasi",
			"startDate":   "2023",
			"endDate":     "Present",
			"status":      "Ongoing",
			"description": "Fokus pada analisis, perancangan sistem, dan pemahaman proses bisnis dengan pendekatan yang menyelaraskan kebutuhan bisnis dan teknologi.",
			"relevantCourses": []string{
				"Analisis & Perancangan Sistem Informasi",
				"Manajemen Sistem Informasi",
				"Sistem Informasi SDM",
				"Pengantar Statistika",
				"Arsitektur & Pemrograman Aplikasi",
				"Database Management",
				"Business Process Management",
			},
			"coreCompetencies": []string{
				"SDLC (Software Development Life Cycle)",
				"UML (Unified Modeling Language)",
				"Business Process Modeling",
				"Database Design & SQL",
				"Data Analysis Fundamentals",
				"Requirements Engineering",
			},
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
		"analytical": []map[string]interface{}{
			{"name": "Requirements Analysis", "level": 85, "icon": "📋"},
			{"name": "Business Process Mapping", "level": 80, "icon": "🔄"},
			{"name": "UML Modeling", "level": 75, "icon": "📐"},
			{"name": "SDLC Understanding", "level": 80, "icon": "♻️"},
			{"name": "Problem Solving", "level": 85, "icon": "🧩"},
		},
		"dataSkills": []map[string]interface{}{
			{"name": "SQL & PostgreSQL", "level": 80, "icon": "🗄️"},
			{"name": "Data Cleaning & Processing", "level": 75, "icon": "🧹"},
			{"name": "Statistical Analysis", "level": 70, "icon": "📊"},
			{"name": "Data Visualization", "level": 70, "icon": "📈"},
			{"name": "Excel/Spreadsheet", "level": 85, "icon": "📑"},
		},
		"programming": []map[string]interface{}{
			{"name": "Python", "level": 70},
			{"name": "SQL", "level": 80},
			{"name": "Java", "level": 65},
			{"name": "Golang", "level": 60},
			{"name": "JavaScript/Vue", "level": 65},
			{"name": "Dart/Flutter", "level": 60},
			{"name": "Svelte", "level": 60},
		},
		"softSkills": []string{
			"Komunikatif & Articulate",
			"Analytical Thinking",
			"Team Collaboration",
			"Stakeholder Communication",
			"Conceptual Problem Solving",
			"Documentation",
		},
		"tools": []string{
			"PostgreSQL",
			"Draw.io / Lucidchart",
			"Figma",
			"Excel / Google Sheets",
			"Git & GitHub",
			"VS Code",
			"Postman",
			"DBeaver",
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
			"title":       "Personal Portfolio Website",
			"description": "A modern, elegant personal website built with Go backend and SvelteKit frontend. Features dark theme, smooth animations, and responsive design.",
			"technologies": []string{"Go", "SvelteKit", "TypeScript"},
			"githubUrl":   "https://github.com/farhanadelio/portfolio",
			"featured":    true,
		},
		{
			"id":          2,
			"title":       "REST API Service",
			"description": "A scalable REST API built with Golang and Gin framework. Includes authentication, database integration, and comprehensive documentation.",
			"technologies": []string{"Go", "PostgreSQL", "Docker"},
			"githubUrl":   "https://github.com/farhanadelio/api-service",
			"featured":    false,
		},
		{
			"id":          3,
			"title":       "Task Management App",
			"description": "Full-stack task management application with real-time updates and collaborative features.",
			"technologies": []string{"Go", "SvelteKit", "WebSocket"},
			"githubUrl":   "https://github.com/farhanadelio/task-app",
			"featured":    false,
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
