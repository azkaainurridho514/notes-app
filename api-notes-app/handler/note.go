package handler

import (
	"github.com/azkaainurridho514/api_notes_app/constant"
	"github.com/azkaainurridho514/api_notes_app/database"
	"github.com/azkaainurridho514/api_notes_app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateNote(c *fiber.Ctx) error {
	db := database.DB.Db
	note := new(model.Note)
	err := c.BodyParser(note)
	if err != nil  {
		return c.Status(400).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_400, 
			"message":  constant.STATUS_CODE_400_MASSAGE, 
			"data": []string{}})
	}
   	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_500, 
			"message":  constant.STATUS_CODE_500_MASSAGE})
	} 
	return c.Status(201).JSON(fiber.Map{
		"status_code": constant.STATUS_CODE_201, 
		"message":  constant.STATUS_CODE_201_MASSAGE})
   }


func GetAllNotes(c *fiber.Ctx) error {
	db := database.DB.Db
	userID := c.Query("user_id")
	var notes []model.Note
	db.Where("user_id = ?", userID).Find(&notes)
	if len(notes) == 0 {
	 return c.Status(404).JSON(fiber.Map{"status_code": constant.STATUS_CODE_404, 
			"message":  constant.STATUS_CODE_404_MASSAGE, "data": []string{}})
	}
	return c.Status(200).JSON(fiber.Map{"status_code": constant.STATUS_CODE_201, 
		"message":  constant.STATUS_CODE_201_MASSAGE, "data": notes})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
	 	Title string `json:"title"`
	 	Description string `json:"description"`
	}
   	db := database.DB.Db
   	var note model.Note
	id := c.Query("id")
	db.Where("id = ?", id).Find(&note, "id = ?", id)
   	if note.ID == uuid.Nil {
	 	return c.Status(404).JSON(fiber.Map{"status_code": constant.STATUS_CODE_404, 
			"message":  constant.STATUS_CODE_404_MASSAGE, "data": nil})
	}
   	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
	 	return c.Status(500).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_500, 
			"message":  constant.STATUS_CODE_500_MASSAGE, 
			"data": []string{}})
	}
	note.Title = updateNoteData.Title
	note.Description = updateNoteData.Description
	db.Save(&note)
	return c.Status(201).JSON(fiber.Map{"status_code": constant.STATUS_CODE_201, 
		"message":  constant.STATUS_CODE_201_MASSAGE, "data": []string{}})
   }

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB.Db
	var note model.Note
	id := c.Query("id")
	db.Find(&note, "id = ?", id)
   	if note.ID == uuid.Nil {
	return c.Status(404).JSON(fiber.Map{
		"status_code": constant.STATUS_CODE_404, 
		"message":  constant.STATUS_CODE_404_MASSAGE})
   	}
   	err := db.Delete(&note, "id = ?", id).Error
   	if err != nil {
	return c.Status(500).JSON(fiber.Map{
		"status_code": constant.STATUS_CODE_500, 
		"message":  constant.STATUS_CODE_500_MASSAGE, 
		"data": []string{}})
	}
   	return c.Status(201).JSON(fiber.Map{"status_code": constant.STATUS_CODE_201, 
		"message":  constant.STATUS_CODE_201_MASSAGE})
   }