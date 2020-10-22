package controllers
import (
	"context"
	"strconv"
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/pmn-kao/app/ent"
	"github.com/pmn-kao/app/ent/doctor"
	"github.com/pmn-kao/app/ent/degree"
	"github.com/pmn-kao/app/ent/department"
	"github.com/pmn-kao/app/ent/nametitle"
)
//DoctorController struct
type DoctorController struct {
	client *ent.Client
	router gin.IRouter
}
//Doctor struct
type Doctor struct {
	Degree          int
	Department      int
	Nametitle       int
	Email     string
	Password  string
	Name      string
	Tel       string
}
// CreateDoctor handles POST requests for adding doctor entities
// @Summary Create doctor
// @Description Create doctor
// @ID create-doctor
// @Accept   json
// @Produce  json
// @Param doctor body Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [post]
func (ctl *DoctorController) CreateDoctor(c *gin.Context) {
	obj := Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctor video binding failed",
		})
		return
	}

	p, err := ctl.client.Nametitle.
		Query().
		Where(nametitle.IDEQ(int(obj.Nametitle))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nametitle not found",
		})
		return
	}

	v, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(obj.Degree))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "degree not found",
		})
		return
	}

	r, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "department not found",
		})
		return
	}
	d, err := ctl.client.Doctor.
		Create().
		SetDegree(v).
		SetDepartment(r).
		SetNametitle(p).
		SetTel(obj.Tel).
		SetPassword(obj.Password).
		SetEmail(obj.Email).
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200,d)
	
}
// DeleteDoctor handles DELETE requests to delete a doctor entity
// @Summary Delete a doctor entity by ID
// @Description get doctor by ID
// @ID delete-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [delete]
func (ctl *DoctorController) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Doctor.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
// GetDoctor handles GET requests to retrieve a doctor entity
// @Summary Get a doctor entity by ID
// @Description get doctor by ID
// @ID get-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /departments/{id} [get]
func (ctl *DoctorController) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	dp, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, dp)
 }

// ListDoctor handles request to get a list of doctor entities
// @Summary List doctor entities
// @Description list doctor entities
// @ID list-doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [get]
func (ctl *DoctorController) ListDoctor(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	doctors, err := ctl.client.Doctor.
		Query().
		WithDegree().
		WithDepartment().
		WithNametitle().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, doctors)
}
// UpdateDoctor handles PUT requests to update a doctor entity
// @Summary Update a doctor entity by ID
// @Description update doctor by ID
// @ID update-doctor
// @Accept   json
// @Produce  json
// @Param id path int true "Doctor ID"
// @Param doctor body ent.Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [put]
func (ctl *DoctorController) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nametitle binding failed",
		})
		return
	}
	obj.ID = int(id)
	nt, err := ctl.client.Doctor.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, nt)
 }

// NewDoctorController creates and registers handles for the doctor controller
func NewDoctorController(router gin.IRouter, client *ent.Client) *DoctorController {
	pvc := &DoctorController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc
}
func (ctl *DoctorController) register() {
	doctors := ctl.router.Group("/doctors")
	doctors.GET("", ctl.ListDoctor)
	// CRUD
	doctors.POST("", ctl.CreateDoctor)
	doctors.GET(":id", ctl.GetDoctor)
	doctors.PUT(":id", ctl.UpdateDoctor)
}