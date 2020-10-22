package main
import (
	"context"
	"log"
	"github.com/pmn-kao/app/controllers"
	_ "github.com/pmn-kao/app/docs"
	"github.com/pmn-kao/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
 )
 
//Degrees struct
type Degrees struct {
	 Degree []Degree
}
//Degree struct
type Degree struct {
	degree  string
}

//Departments struct
type Departments struct {
	Department []Department
}
//Department struct
type Department struct {
	department  string
}

//Nametitles struct
type Nametitles struct {
	Nametitle []Nametitle
}
//Nametitle struct 
type Nametitle struct {
	 nametitle  string
   
}

// @title SUT SA Example API Doctor 
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDegreeController(v1, client)
	controllers.NewDepartmentController(v1, client)
	controllers.NewNametitleController(v1, client)
	controllers.NewDoctorController(v1, client)

	// Set Degrees Data
	degrees := Degrees{
		Degree: []Degree{
			Degree{"ปฏิบัติการ"},
			Degree{"ชำนาญการ"},
			Degree{"ชำนาญการพิเศษ"},
			Degree{"เชี่ยวชาญ"},
		},
	}

	for _, de := range degrees.Degree {
		client.Degree.
			Create().
			SetDegree(de.degree).
			Save(context.Background())
	}
	// Set Departments Data
	departments := Departments{
		Department: []Department{
			Department{"Emergency Room"},
			Department{"Outpatient Department"},
			Department{"Intensive Care Unitr"},
			Department{"Medicine"},
			Department{"Pediatric"},
			Department{"Orthopedic"},
			Department{"Obstretic Gynecology"},
		},
	}

	for _, dp := range departments.Department {
		client.Department.
			Create().
			SetDepartment(dp.department).
			Save(context.Background())
		}

	// Set Nametitles Data
	nametitles := Nametitles{
		Nametitle: []Nametitle{
			Nametitle{"นายแพทย์"},
			Nametitle{"แพทย์หญิง"},
		},
	}
	for _, nt := range nametitles.Nametitle {
		client.Nametitle.
			Create().
			SetNameTitle(nt.nametitle).
			Save(context.Background())
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}