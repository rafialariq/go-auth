package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/go-auth/config"
	"github.com/rafialariq/go-auth/entities"
	"github.com/rafialariq/go-auth/entities/dto"
	"github.com/rafialariq/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string
	Password string
}

var userModel = models.NewUserModel()

func Index(c *gin.Context) {
	session, _ := config.Store.New(c.Request, config.SESSION_ID)

	if len(session.Values) == 0 {

		c.Redirect(http.StatusSeeOther, "/login")
	} else {

		if session.Values["loggedIn"] != true {

			c.Redirect(http.StatusSeeOther, "/login")
		} else {

			c.HTML(http.StatusOK, "index.html", gin.H{
				"firstName": session.Values["firstName"],
				"lastName":  session.Values["lastName"],
			})
		}
	}
}

func Login(c *gin.Context) {

	loginInput := &LoginInput{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	errorMessage := make(map[string]interface{})

	if loginInput.Username == "" {
		errorMessage["Username"] = "username wajib di isi"
	}

	if loginInput.Password == "" {
		errorMessage["Password"] = "password wajib di isi"
	}

	if len(errorMessage) > 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"validation": errorMessage,
		})
	} else {

		var user entities.User
		userModel.Where(&user, "username", loginInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("username atau kata sandi salah")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
			if errPassword != nil {
				message = errors.New("username atau kata sandi salah")
			}
		}

		if message != nil {

			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"error": message,
			})
		} else {

			session, _ := config.Store.Get(c.Request, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["firstName"] = user.FirstName
			session.Values["lastName"] = user.LastName
			session.Values["username"] = user.Username
			session.Values["email"] = user.Email

			session.Save(c.Request, c.Writer)

			c.Redirect(http.StatusSeeOther, "/")
		}
	}
}

func Logout(c *gin.Context) {
	session, _ := config.Store.Get(c.Request, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/login")
}

func Register(c *gin.Context) {
	registerInput := &dto.RegisterDTO{
		FirstName: c.PostForm("first_name"),
		LastName:  c.PostForm("last_name"),
		Username:  c.PostForm("username"),
		Email:     c.PostForm("email"),
		Password:  c.PostForm("password"),
		Cpassword: c.PostForm("cpassword"),
	}

	errorMessage := make(map[string]interface{})

	if registerInput.FirstName == "" {
		errorMessage["FirstName"] = "first name wajib di isi"
	}
	if registerInput.LastName == "" {
		errorMessage["LastName"] = "last name wajib di isi"
	}
	if registerInput.Username == "" {
		errorMessage["Username"] = "username wajib di isi"
	}
	if registerInput.Email == "" {
		errorMessage["Email"] = "email wajib di isi"
	}
	if registerInput.Password == "" {
		errorMessage["Password"] = "password wajib di isi"
	}
	if registerInput.Cpassword == "" {
		errorMessage["Cpassword"] = "konfirmasi password wajib di isi"
	} else if registerInput.Password != registerInput.Cpassword {
		errorMessage["Cpassword"] = "konfirmasi password tidak sesuai"
	}

	if len(errorMessage) > 0 {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"validation": errorMessage,
			"user":       registerInput,
		})
		return
	} else {

		// hashed password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)

		// assign dto to user object
		userData := &entities.User{
			FirstName: registerInput.FirstName,
			LastName:  registerInput.LastName,
			Username:  registerInput.Username,
			Email:     registerInput.Email,
			Password:  string(hashedPassword),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		userModel.Create(userData)

		c.HTML(http.StatusCreated, "register.html", gin.H{
			"pesan":     "Registrasi Berhasil!",
			"informasi": "Silahkan Login menggunakan akun yang sudah terdaftar.",
		})

	}
}
