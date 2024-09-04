package area

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/mspos-api/database"
	"github.com/kgermando/mspos-api/models"
)

// Get All data
func GetAreas(c *fiber.Ctx) error {

	p, _ := strconv.Atoi(c.Query("page", "1"))
	l, _ := strconv.Atoi(c.Query("limit", "15"))

	return c.JSON(models.Paginate(database.DB, &models.Area{}, p, l))
}

// query data
func GetAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var areas []models.Area
	db.Where("province_id = ?", id).Find(&areas)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "areas by id found", 
		"data": areas,
	})
}

// query data
func GetSupAreaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var areas []models.Area
	db.Where("sup_id = ?", id).Find(&areas)
	 
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "poss by id found", 
		"data": areas,
	})
}


// Get one data
func GetArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var area models.Area
	db.Find(&area, id)
	if area.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No area name found",
				"data":    nil,
			},
		)
	}
	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area found",
			"data":    area,
		},
	)
}

// Create data
func CreateArea(c *fiber.Ctx) error {
	p := &models.Area{}

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	database.DB.Create(p)

	return c.JSON(p)
}

// Update data
func UpdateArea(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	type UpdateData struct {
		Name       string `json:"name"`
		ProvinceID uint   `json:"province_id"`
		SupID      uint   `json:"sup_id"`
		Signature  string `json:"signature"`
	}

	var updateData UpdateData

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": "Review your iunput",
				"data":    nil,
			},
		)
	}

	area := new(models.Area)

	db.First(&area, id)
	area.Name = updateData.Name
	area.ProvinceID = updateData.ProvinceID
	area.SupID = updateData.SupID
	area.Signature = updateData.Signature

	db.Save(&area)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area updated success",
			"data":    area,
		},
	)

}

// Delete data
func DeleteArea(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var area models.Area
	db.First(&area, id)
	if area.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"status":  "error",
				"message": "No area name found",
				"data":    nil,
			},
		)
	}

	db.Delete(&area)

	return c.JSON(
		fiber.Map{
			"status":  "success",
			"message": "area deleted success",
			"data":    nil,
		},
	)
}