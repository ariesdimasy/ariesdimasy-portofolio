package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"ariesdimasy-portofolio/configs"
	"ariesdimasy-portofolio/controllers"
	"ariesdimasy-portofolio/helpers"
	"ariesdimasy-portofolio/repositories"
	"ariesdimasy-portofolio/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ─── JWT Middleware ───────────────────────────────────────────────────────────

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := helpers.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id in token"})
			return
		}

		c.Set("user_id", uint(userIDFloat))
		c.Next()
	}
}

// ─── Main ─────────────────────────────────────────────────────────────────────

func main() {
	// 1. Load environment variables
	configs.LoadEnv()

	// 2. Connect & migrate database
	configs.InitPgDB()

	// 3. Dependency injection: Repository → Service → Controller

	// User
	userRepo := repositories.NewUserRepository(configs.DB)
	userService := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userService)

	// Biodata
	biodataRepo := repositories.NewBiodataRepository(configs.DB)
	biodataService := services.NewBiodataService(biodataRepo)
	biodataCtrl := controllers.NewBiodataController(biodataService)

	// Skill
	skillRepo := repositories.NewSkillRepository(configs.DB)
	skillService := services.NewSkillService(skillRepo)
	skillCtrl := controllers.NewSkillController(skillService)

	// Sosmed
	sosmedRepo := repositories.NewSosmedRepository(configs.DB)
	sosmedService := services.NewSosmedService(sosmedRepo)
	sosmedCtrl := controllers.NewSosmedController(sosmedService)

	// Education
	educationRepo := repositories.NewEducationRepository(configs.DB)
	educationService := services.NewEducationService(educationRepo)
	educationCtrl := controllers.NewEducationController(educationService)

	// Experience
	experienceRepo := repositories.NewExperienceRepository(configs.DB)
	experienceService := services.NewExperienceService(experienceRepo)
	experienceCtrl := controllers.NewExperienceController(experienceService)

	// Certificate
	certificateRepo := repositories.NewCertificateRepository(configs.DB)
	certificateService := services.NewCertificateService(certificateRepo)
	certificateCtrl := controllers.NewCertificateController(certificateService)

	// Project
	projectRepo := repositories.NewProjectRepository(configs.DB)
	projectService := services.NewProjectService(projectRepo)
	projectCtrl := controllers.NewProjectController(projectService)

	// 4. Setup Gin router
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Portofolio API",
			"status":  "success",
			"endpoints": map[string]string{
				"auth":         "/api/auth",
				"users":        "/api/users",
				"biodata":      "/api/biodata",
				"skills":       "/api/skills",
				"sosmeds":      "/api/sosmeds",
				"educations":   "/api/educations",
				"experiences":  "/api/experiences",
				"certificates": "/api/certificates",
				"projects":     "/api/projects",
			},
		})
	})

	// 5. Register routes
	api := r.Group("/api")
	{
		// ── Auth (public) ─────────────────────────────────────────────────────
		auth := api.Group("/auth")
		{
			auth.POST("/register", userCtrl.Register)
			auth.POST("/login", userCtrl.Login) // login + JWT handled in UserController
		}

		// ── Users ─────────────────────────────────────────────────────────────
		users := api.Group("/users", JWTMiddleware())
		{
			users.POST("", userCtrl.CreateUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("", userCtrl.DeleteUser)
			users.GET("/:id", userCtrl.GetUserByID)
			users.GET("", userCtrl.GetAllUsers)
		}

		// ── Biodata ───────────────────────────────────────────────────────────
		biodata := api.Group("/biodata", JWTMiddleware())
		{
			biodata.POST("", biodataCtrl.CreateBiodata)
			biodata.PUT("/:id", biodataCtrl.UpdateBiodata)
			biodata.DELETE("", biodataCtrl.DeleteBiodata)
			biodata.GET("/:id", biodataCtrl.GetBiodataByID)
		}

		// ── Skills (GET = public, mutate = protected) ─────────────────────────
		api.GET("/skills", skillCtrl.GetAllSkills)
		api.GET("/skills/:id", skillCtrl.GetSkillByID)
		skills := api.Group("/skills", JWTMiddleware())
		{
			skills.POST("", skillCtrl.CreateSkill)
			skills.PUT("/:id", skillCtrl.UpdateSkill)
			skills.DELETE("", skillCtrl.DeleteSkill)
		}

		// ── Social Media ──────────────────────────────────────────────────────
		sosmeds := api.Group("/sosmeds", JWTMiddleware())
		{
			sosmeds.POST("", sosmedCtrl.CreateSosmed)
			sosmeds.PUT("/:id", sosmedCtrl.UpdateSosmed)
			sosmeds.DELETE("", sosmedCtrl.DeleteSosmed)
			sosmeds.GET("/:id", sosmedCtrl.GetSosmedByID)
			sosmeds.GET("", sosmedCtrl.GetAllSosmeds)
		}

		// ── Education ─────────────────────────────────────────────────────────
		educations := api.Group("/educations", JWTMiddleware())
		{
			educations.POST("", educationCtrl.CreateEducation)
			educations.PUT("/:id", educationCtrl.UpdateEducation)
			educations.DELETE("", educationCtrl.DeleteEducation)
			educations.GET("/:id", educationCtrl.GetEducationByID)
			educations.GET("", educationCtrl.GetAllEducations)
		}

		// ── Experience ────────────────────────────────────────────────────────
		experiences := api.Group("/experiences", JWTMiddleware())
		{
			experiences.POST("", experienceCtrl.CreateExperience)
			experiences.PUT("/:id", experienceCtrl.UpdateExperience)
			experiences.DELETE("", experienceCtrl.DeleteExperience)
			experiences.GET("/:id", experienceCtrl.GetExperienceByID)
			experiences.GET("", experienceCtrl.GetAllExperiences)
		}

		// ── Certificates ──────────────────────────────────────────────────────
		certificates := api.Group("/certificates", JWTMiddleware())
		{
			certificates.POST("", certificateCtrl.CreateCertificate)
			certificates.PUT("/:id", certificateCtrl.UpdateCertificate)
			certificates.DELETE("", certificateCtrl.DeleteCertificate)
			certificates.GET("/:id", certificateCtrl.GetCertificateByID)
			certificates.GET("", certificateCtrl.GetAllCertificates)
		}

		// ── Projects ──────────────────────────────────────────────────────────
		projects := api.Group("/projects", JWTMiddleware())
		{
			projects.POST("", projectCtrl.CreateProject)
			projects.PUT("/:id", projectCtrl.UpdateProject)
			projects.DELETE("", projectCtrl.DeleteProject)
			projects.GET("/:id", projectCtrl.GetProjectByID)
			projects.GET("", projectCtrl.GetAllProjects)
		}
	}

	// 6. Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	r.Run(":" + port)
}
