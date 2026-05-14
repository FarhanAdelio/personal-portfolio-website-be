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
			"twitter":  "https://www.instagram.com/farhanadelio/",
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
			"startDate":   "2024-12",
			"endDate":     "2026-2",
			"logo":        "/img/logo/logodinas.png",
			"description": "Assisted in processing and compiling educational data, analyzing system/application requirements, and supporting system and data documentation. Collaborated with technical and non-technical teams to improve business process efficiency.",
			"achievements": []string{
				"Analyzed and processed raw data into actionable information for decision-making",
				"Assisted in mapping business processes and workflows for system digitalization",
				"Created requirement analysis documentation for internal application development",
				"Contributed to cross-departmental communication (IT, public relations, data, management)",
			},
			"technologies": []string{"SQL", "PostgreSQL", "Python", "Excel", "Data Analysis"},
			"learnings": []string{
				"Understanding government systems workflow and bureaucracy",
				"Experience in processing data from raw to ready-to-use",
				"The importance of data accuracy for decision making",
				"Communication skills with non-technical stakeholders",
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
			"field":       "Information Systems",
			"startDate":   "2023",
			"endDate":     "Present",
			"status":      "Ongoing",
			"description": "Focused on analysis, system design, and business process understanding with an approach that aligns business needs with technology solutions.",
			"relevantCourses": []string{
				"Information Systems Analysis & Design",
				"Information Systems Management",
				"Human Resource Information Systems",
				"Introduction to Statistics",
				"Application Architecture & Programming",
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
			{"name": "Python", "level": 70, "logo": "/img/logo/python.png"},
			{"name": "SQL", "level": 80, "logo": "/img/logo/sql.png"},
			{"name": "Java", "level": 65, "logo": "/img/logo/java.png"},
			{"name": "Golang", "level": 60, "logo": "/img/logo/go.png"},
			{"name": "Vue.js", "level": 65, "logo": "/img/logo/vue.png"},
			{"name": "Flutter", "level": 60, "logo": "/img/logo/flutter.png"},
			{"name": "Svelte", "level": 60, "logo": "/img/logo/svelte.png"},
		},
		"softSkills": []string{
			"Komunikatif & Articulate",
			"Analytical Thinking",
			"Team Collaboration",
			"Stakeholder Communication",
			"Conceptual Problem Solving",
			"Documentation",
		},
		"tools": []map[string]interface{}{
			{"name": "PostgreSQL", "logo": "/img/logo/postgre.png"},
			{"name": "Figma", "logo": "/img/logo/figma.png"},
			{"name": "Excel", "logo": "/img/logo/exel.png"},
			{"name": "Github", "logo": "/img/logo/github.png"},
			{"name": "VS Code", "logo": "/img/logo/vscode.png"},
			{"name": "Postman", "logo": "/img/logo/postman.png"},
			{"name": "Docker", "logo": "/img/logo/docker.png"},
			{"name": "DBeaver", "logo": "/img/logo/Dbvear.png"},
			{"name": "Gitlab", "logo": "/img/logo/gitlab.png"},
			{"name": "Draw", "logo": "/img/logo/draw.png"},

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
			"title":       "Smart Waste Management System",
			"category":    "Fullstack",
			"type":        []string{"Web App", "IoT System"},
			"date":        "01/2026",
			"description": "A comprehensive waste management system that streamlines garbage collection, tracking, and reporting. Features real-time monitoring, route optimization for waste collectors, and data analytics dashboard for waste management insights.",
			"highlights": []string{
				"Real-time Tracking",
				"Route Optimization",
				"Analytics Dashboard",
				"Mobile-responsive",
			},
			"technologies": []string{"Go", "PostgreSQL", "SvelteKit", "TypeScript", "Docker", "Google Maps API"},
			// "techIcons": []string{"🔷", "🐘", "🔥", "📘", "🐳", "🗺️"},
			"githubUrl":   "https://github.com/FarhanAdelio/Retrash",
			"liveUrl":     "https://waste-mgmt-demo.com",
			"featured":    true,
			"video":       "/videos/Screen Recording 2026-01-21 at 12.55.41.mp4",
			"videoType":   "local",
			"features": []string{
				"User authentication and role-based access (Admin, Collector, Resident)",
				"Real-time garbage collection status tracking with live updates",
				"Interactive map showing collection routes and waste bin locations",
				"Automated scheduling system for collection routes",
				"Mobile app for waste collectors with offline support",
				"Analytics dashboard with waste statistics and trends",
				"Notification system for collection schedules and updates",
				"Report generation for waste management insights",
			},
		},
		{
			"id":          2,
			"title":       "REST API Service",
			"description": "A scalable REST API built with Golang and Gin framework. Includes authentication, database integration, and comprehensive documentation.",
			"technologies": []string{"Go", "PostgreSQL", "Docker"},
			"githubUrl":   "https://github.com/farhanadelio/api-service",
			"featured":    false,
			"video":       "",  // Bisa kosong jika tidak ada video
			"videoType":   "local",
		},
		// {
		// 	"id":          3,
		// 	"title":       "Task Management App",
		// 	"description": "Full-stack task management application with real-time updates and collaborative features.",
		// 	"technologies": []string{"Go", "SvelteKit", "WebSocket"},
		// 	"githubUrl":   "https://github.com/farhanadelio/task-app",
		// 	"featured":    false,
		// 	"video":       "dQw4w9WgXcQ",  // YouTube video ID (contoh)
		// 	"videoType":   "youtube",
		// },
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
