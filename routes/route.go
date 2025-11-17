package routes

import (
	"latihan-hris/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	role := r.Group("/roles")
	{
		role.GET("/", controllers.GetRoles)
		role.POST("/", controllers.CreateRole)
		role.DELETE("/:id", controllers.DeleteRole)
	}

	user := r.Group("/users")
	{
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserById)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	department := r.Group("/departments")
	{
		department.GET("/", controllers.GetDepartments)
		department.GET("/:id", controllers.GetDepartmentById)
		department.POST("/", controllers.CreateDepartment)
		department.DELETE("/:id", controllers.DeleteDepartment)
	}

	division := r.Group("/divisions")
	{
		division.POST("/", controllers.CreateDivision)
		division.GET("/", controllers.GetDivisions)
		division.DELETE("/:id", controllers.DeleteDivision)
	}

	employee := r.Group("/employees")
	{
		employee.POST("/", controllers.CreateEmployee)
		employee.GET("/:id", controllers.GetEmployeeById)
		employee.GET("/", controllers.GetEmployees)
		employee.PUT("/:id", controllers.UpdateEmployee)
		employee.DELETE("/:id", controllers.DeleteEmployee)
	}

	employeeDetail := r.Group("/employee/details")
	{
		employeeDetail.PUT("/:id", controllers.UpdateEmployeeDetail)
	}

	photo := r.Group("/photos")
	{
		photo.POST("/", controllers.UploadPhoto)
		photo.DELETE("/:id", controllers.DeletePhoto)
		photo.PUT("/:id", controllers.UpdateIsProfile)
		photo.GET("/:employee_id", controllers.GetPhotos)
		photo.GET("/profile/:employee_id", controllers.GetPhotoProfile)
	}

	document := r.Group("/documents") //cek
	{
		document.POST("/", controllers.UploadDocument)
		document.DELETE("/:id", controllers.DeleteDocument)
	}

	position := r.Group("/positions")
	{
		position.POST("/", controllers.CreatePosition)
		position.GET("/", controllers.GetPositions)
		position.GET("/:id", controllers.GetDepartmentById)
		position.PUT("/:id", controllers.UpdatePosition)
		position.DELETE("/:id", controllers.DeletePosition)
	}
}

func AuthRoute(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.POST("/logout", controllers.Logout)
	r.GET("/verifikasi", controllers.VerifiedUser) // cek
}
