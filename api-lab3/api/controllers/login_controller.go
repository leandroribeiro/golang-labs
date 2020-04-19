package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/go-labs/api-lab3/api/auth"
	"github.com/leandroribeiro/go-labs/api-lab3/api/models"
	"github.com/leandroribeiro/go-labs/api-lab3/api/responses"
	"github.com/leandroribeiro/go-labs/api-lab3/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)

	fmt.Println(token)
	fmt.Println(err)

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formatedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	//TODO Debug
	fmt.Println("VerifyPassword")
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", nil
	}

	//TODO Debug
	fmt.Println("CreateToken")
	return auth.CreateToken(user.ID)
}
