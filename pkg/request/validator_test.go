package request

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type UserRequest struct {
	Name     string `uri:"name" form:"name" json:"name" binding:"required"`
	Password string `uri:"password" form:"password" json:"password" binding:"required"`
}

func (r UserRequest) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "请输入用户名",
		"Password.required": "请输入密码",
	}
}

// get请求验证
func TestValidatorGet(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	// get 方式
	req, _ := http.NewRequest(http.MethodGet, "/validateGet", nil)
	params := make(url.Values)
	params.Add("name", "test name")
	//params.Add("password", "test password")
	req.URL.RawQuery = params.Encode()
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"message\":\"请输入密码\"}")
	fmt.Println("get body：" + w.Body.String())
}

// post请求验证
func TestValidatorPost(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	// get 方式
	request := &UserRequest{
		Name:     "tom",
		Password: "123456",
	}
	reqParam, _ := json.Marshal(&request)
	reqBody := strings.NewReader(string(reqParam))
	req, _ := http.NewRequest(http.MethodPost, "/validatePost", reqBody)
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	params := make(url.Values)
	params.Add("name", "test name")
	//params.Add("password", "test password")
	req.URL.RawQuery = params.Encode()
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Body.String(), "{\"message\":\"请输入密码\"}")
	fmt.Println("get body：" + w.Body.String())
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/validateGet", func(c *gin.Context) {
		userRequest := &UserRequest{}
		err := c.ShouldBind(userRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": GetErrorMsg(userRequest, err),
			})
			return
		}
		c.String(200, "pong")
	})
	r.POST("/validatePost", func(c *gin.Context) {
		userRequest := &UserRequest{}
		err := c.ShouldBindQuery(userRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": GetErrorMsg(userRequest, err),
			})
			return
		}
		c.String(200, "pong")
	})
	//soft_delete.DeletedAt
	return r
}
