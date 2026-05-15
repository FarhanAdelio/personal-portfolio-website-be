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
		"subtitle": "System Analyst & Business Intelligence Enthusiast",
		"tagline":  "Bridging the gap between business strategy and technology implementation",
		"bio":      "Information Systems student with a strong foundation in system analysis, business process mapping, and data-driven decision making. Experienced in collaborating with government and industry stakeholders to translate business needs into effective system solutions. Technically capable in backend development and database management.",
		"email":    "farhanadelio10@gmail.com",
		"phone":    "087775475299",
		"location": "Jakarta, Indonesia",
		"careerInterests": []string{
			"System Analyst",
			"Business Analyst",
			"Data Analyst",
			"Information Analyst",
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
			"position":    "System Analyst (Intern)",
			"type":        "Internship",
			"location":    "Jakarta, Indonesia",
			"startDate":   "2025-12",
			"endDate":     "2026-02",
			"logo":        "/img/logo/logodinas.png",
			"description": "Contributed to the development and improvement of education information systems, focusing on system scalability and implementation support within a government environment.",
			"achievements": []string{
				"Designed system solution for teacher recruitment exam platform to handle high concurrent users",
				"Analyzed system requirements and applied SDLC approach to ensure scalable design",
				"Created sequence diagrams to model system workflows and optimize request handling",
				"Performed data preparation, filtering, and validation using SQL",
				"Supported backend integration using Golang and frontend development using SvelteKit",
				"Collaborated with cross-functional teams to align business requirements",
			},
			"technologies": []string{"Golang", "SvelteKit", "SQL", "SDLC", "System Design"},
		},
		{
			"id":          2,
			"company":     "PT Petro One Indonesia",
			"position":    "Freelance Backend Developer",
			"type":        "Contract",
			"location":    "Jakarta, Indonesia",
			"startDate":   "2026-01",
			"endDate":     "2026-03",
			"logo":        "/img/logo/logopetroone.png",
			"description": "Developed a web-based ERP system to support company operations and asset management within an IT services and system integration environment.",
			"achievements": []string{
				"Designed and implemented backend services using Java Spring Boot",
				"Structured and managed PostgreSQL database for operational data processing",
				"Built RESTful APIs to handle asset management and workflow processes",
				"Collaborated with stakeholders to gather system requirements",
				"Contributed to system digitalization, enhancing efficiency and data organization",
			},
			"technologies": []string{"Java Spring Boot", "PostgreSQL", "RESTful API", "System Integration"},
		},
		{
			"id":          3,
			"company":     "Dinas Kebudayaan Provinsi DKI Jakarta",
			"position":    "Freelance Web Developer",
			"type":        "Contract",
			"location":    "Jakarta, Indonesia",
			"startDate":   "2025-04",
			"endDate":     "2025-09",
			"logo":        "/img/logo/logokebudayaan.png",
			"description": "Developed a web-based waste management system based on direct collaboration with government stakeholders, focusing on translating business needs into a functional digital solution.",
			"achievements": []string{
				"Conducted requirement gathering through discussions with Cultural Department head",
				"Translated business requirements into system design and technical implementation",
				"Designed system features to support waste tracking and management processes",
				"Developed application using Django with structured database modeling",
				"Delivered functional system ready for real-world implementation",
				"Collaborated with stakeholders to ensure alignment with operational needs",
			},
			"technologies": []string{"Django", "Python", "PostgreSQL", "System Design"},
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
		"programming": []map[string]interface{}{
			{"name": "SQL & PostgreSQL", "level": 85, "logo": "/img/logo/sql.png"},
			{"name": "Python", "level": 70, "logo": "/img/logo/python.png"},
			{"name": "Golang", "level": 65, "logo": "/img/logo/go.png"},
			{"name": "Java", "level": 65, "logo": "/img/logo/java.png"},
			{"name": "Vue.js", "level": 65, "logo": "/img/logo/vue.png"},
			{"name": "Svelte", "level": 60, "logo": "/img/logo/svelte.png"},
			{"name": "Flutter", "level": 60, "logo": "/img/logo/flutter.png"},
		},
		"dataSkills": []map[string]interface{}{
			{"name": "Power BI", "level": 80, "icon": "📊"},
			{"name": "Tableau", "level": 80, "icon": "📈"},
			{"name": "SQL & Database", "level": 85, "icon": "🗄️"},
			{"name": "Excel", "level": 85, "icon": "📑"},
			{"name": "Data Analysis", "level": 75, "icon": "🔍"},
		},
		"analytical": []map[string]interface{}{
			{"name": "Requirements Analysis", "level": 85, "icon": "📋"},
			{"name": "Business Process Mapping", "level": 85, "icon": "🔄"},
			{"name": "UML Modeling", "level": 80, "icon": "📐"},
			{"name": "SDLC", "level": 80, "icon": "♻️"},
			{"name": "System Design", "level": 80, "icon": "🏗️"},
			{"name": "Problem Solving", "level": 85, "icon": "🧩"},
		},
		"tools": []map[string]interface{}{
			{"name": "PostgreSQL", "logo": "/img/logo/postgre.png"},
			{"name": "Excel", "logo": "/img/logo/exel.png"},
			{"name": "Github", "logo": "/img/logo/github.png"},
			{"name": "VS Code", "logo": "/img/logo/vscode.png"},
			{"name": "Postman", "logo": "/img/logo/postman.png"},
			{"name": "Docker", "logo": "/img/logo/docker.png"},
			{"name": "Figma", "logo": "/img/logo/figma.png"},
		},
		"softSkills": []string{
			"Stakeholder Communication",
			"Analytical Problem Solving",
			"Team Collaboration",
			"Documentation",
			"Clear Communication",
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
			"title":       "ThinkNalyze – Internal Sales & Management System",
			"category":    "Full-Stack",
			"date":        "2026",
			"description": "Led system analysis and full-cycle development for a real client, an AI-based stock market analysis platform. Produced complete system documentation and integrated internal system with existing ThinkNalyze product.",
			"highlights": []string{
				"System Analysis & Design",
				"Full-Stack Development",
				"Microservices Architecture",
				"Dashboard Implementation",
			},
			"achievements": []string{
				"Produced complete system documentation (requirements, architecture, design)",
				"Conducted stakeholder discussions to gather and validate business requirements",
				"Designed internal sales management system and management dashboard",
				"Integrated internal system with existing ThinkNalyze product",
				"Applied microservices architecture",
			},
			"technologies": []string{"Golang", "HTML/CSS", "System Design", "Microservices"},
			"featured":    true,
			"role":        "System Analyst & Full-Stack Developer",
		},
		{
			"id":          2,
			"title":       "Hotel Management System – Microservices Architecture",
			"category":    "Backend/Architecture",
			"date":        "2025",
			"description": "Built an enterprise-grade hotel management platform with 5 independent microservices covering internal operations and a public-facing website.",
			"highlights": []string{
				"5 Independent Microservices",
				"Enterprise Architecture",
				"REST APIs",
				"Database Design",
			},
			"achievements": []string{
				"Designed and implemented 5 microservices with inter-service communication",
				"Developed backend using Java Spring Boot",
				"Developed frontend using Vue.js",
				"Applied microservices architecture",
				"Produced system design and documentation",
			},
			"technologies": []string{"Java Spring Boot", "Vue.js", "PostgreSQL", "Microservices", "REST API"},
			"featured":    true,
			"role":        "Backend Developer & System Architect",
		},
		{
			"id":          3,
			"title":       "Education Information System – UI/UX & System Design",
			"category":    "System Design",
			"date":        "2024",
			"description": "Designed a system solution to improve internal education administration efficiency for Dinas Pendidikan DKI Jakarta, focusing on UI/UX and system architecture.",
			"highlights": []string{
				"System Analysis",
				"Requirement Gathering",
				"UI/UX Design",
				"Business Process Mapping",
			},
			"achievements": []string{
				"Conducted stakeholder interviews and requirement gathering sessions",
				"Produced comprehensive system documentation and business process mapping",
				"Designed end-to-end UI/UX prototype using Figma",
				"Applied UML modeling and SDLC principles",
				"Delivered solution ready for implementation",
			},
			"technologies": []string{"Figma", "UML", "SDLC", "Business Analysis", "System Design"},
			"featured":    false,
			"role":        "System Analyst",
		},
		{
			"id":          4,
			"title":       "Shoe E-Commerce Platform – Web & Mobile Application",
			"category":    "Full-Stack",
			"date":        "2024",
			"description": "Developed a comprehensive e-commerce platform for shoe sales with web and mobile applications, featuring user authentication, product management, and payment integration.",
			"highlights": []string{
				"E-Commerce Platform",
				"Web & Mobile Apps",
				"Payment Integration",
				"User Authentication",
			},
			"achievements": []string{
				"Designed system architecture for web and mobile platforms",
				"Implemented user authentication and authorization system",
				"Developed product catalog with search and filtering functionality",
				"Integrated payment gateway for secure transactions",
				"Created responsive UI/UX design for both platforms",
				"Deployed application with scalable infrastructure",
			},
			"technologies": []string{"Flutter", "Vue.js", "Node.js", "PostgreSQL", "REST API"},
			"featured":    false,
			"role":        "Full-Stack Developer",
		},
		{
			"id":          5,
			"title":       "Zoo Management System – Database-Focused Web Application",
			"category":    "Backend/Database",
			"date":        "2023",
			"description": "Built a comprehensive zoo management system focusing on database design and backend services for managing animals, staff, visitors, and operations.",
			"highlights": []string{
				"Database Design",
				"Backend Services",
				"Data Management",
				"Reporting System",
			},
			"achievements": []string{
				"Designed comprehensive database schema for zoo operations",
				"Implemented backend APIs for animal management and inventory",
				"Created reporting and analytics dashboard",
				"Developed staff management and scheduling system",
				"Implemented visitor management and ticketing system",
				"Optimized database queries for performance",
			},
			"technologies": []string{"Java", "PostgreSQL", "Spring Boot", "SQL", "REST API"},
			"featured":    false,
			"role":        "Backend Developer & Database Designer",
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
