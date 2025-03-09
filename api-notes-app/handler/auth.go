package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/azkaainurridho514/api_notes_app/constant"
	"github.com/azkaainurridho514/api_notes_app/database"
	"github.com/azkaainurridho514/api_notes_app/model"
)

func Register(c *fiber.Ctx) error {
	db := database.DB.Db
	userFromFE := new(model.UserFromFeRegister)
	var user model.User
	err := c.BodyParser(userFromFE)
	if err != nil{
		return c.Status(400).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_400, 
			"message":  constant.STATUS_CODE_400_MASSAGE, 
			"data": []string{}})
	}
	bytes, errs := bcrypt.GenerateFromPassword([]byte(userFromFE.Password), bcrypt.DefaultCost)
	if errs != nil{
		return c.Status(500).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_500, 
			"message":  constant.STATUS_CODE_500_MASSAGE, 
			"data": []string{}})
	}
	userToCreate := model.User{
		Email: userFromFE.Email,
		Username: userFromFE.Username,
		Phone: userFromFE.Phone,
		Password:  string(bytes),
		Address: userFromFE.Address,
	}
   	err = db.Create(&userToCreate).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_500, 
			"message":  constant.STATUS_CODE_500_MASSAGE,
			"data": []string{}})
	}
	db.Where("email = ?", userFromFE.Email).First(&user)
	var userID = userToCreate.ID
	if userID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_404, 
			"message": constant.STATUS_CODE_404_MASSAGE, 
			"data": []string{},
		})
	}
	register := model.RegisterUser{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Phone: user.Phone,
		Address: user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return c.Status(201).JSON(fiber.Map{
		"status_code": constant.STATUS_CODE_201, 
		"message":  constant.STATUS_CODE_201_MASSAGE, 
		"data": register})
}
func Login(c *fiber.Ctx) error {
	db:= database.DB.Db
	email := c.Query("email")
	password := c.Query("password")
	var user model.User
	db.Where("email = ?", email).First(&user)
	var userID = user.ID
	if userID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_404, 
			"message": constant.STATUS_CODE_404_MASSAGE, 
			"data": []string{},
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status_code": constant.STATUS_CODE_401, 
			"message":  constant.STATUS_CODE_401_MASSAGE, 
			"data": []string{}})
	}
	db.Where("email = ?", user.Email).First(&user)
	toFE := model.UserToFE{
		ID: user.ID,
		Username: user.Username,
		Phone: user.Phone,
		Email: user.Email,
		Address: user.Address,
	}
	return c.Status(200).JSON(fiber.Map{
		"status_code": constant.STATUS_CODE_200, 
		"message":  constant.STATUS_CODE_200_MASSAGE,  
		"data": toFE,
	})
}