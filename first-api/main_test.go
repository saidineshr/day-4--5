package main

import (
	"first-api/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}

//func TestCreateUser(t *testing.T) {
	//r := SetUpRouter()
	//r.POST("/user-api/user", Controllers.CreateUser)
//	req:=httptest.NewRequest(http.MethodPost,"/user-api/user",nil)
//	w:=httptest.NewRecorder()
//
//	Controllers.CreateUser(w,req)
//
//
//	company := Models.User{
//		Id     : uint(1),
//		Name    :"xx",
//		Email   :"xx",
//		LastName :"xx",
//		DOB      :"xx",
//		Subject :"xx",
//		Marks   :int32(0),
//		Address:"xx",
//	}
//
//	jsonValue, _ := json.Marshal(company)
//	fmt.Println("..",jsonValue)
//	req, _ := http.NewRequest("POST", "http://localhost:8080/user-api/user", bytes.NewBuffer(jsonValue))
//
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	assert.Equal(t, http.StatusOK, w.Code)
//}
func TestGetUserById(t *testing.T) {
	rPath := "/user-api/user/:id"
	router := gin.Default()
	router.Run(":8080")8080
	router.GET(rPath, Controllers.GetUserByID)

	req, _ := http.NewRequest("GET", rPath, strings.NewReader(`{"id": "1"}`))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	//t.Logf("status: %d", w.Code)
	//t.Logf("response: %s", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}