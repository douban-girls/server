package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/douban-girls/qiniu-migrate/qn"

	"github.com/graphql-go/graphql"

	"strconv"

	"github.com/revel/revel"
)

// Response just return response info
func Response(status int, data interface{}, err error) (result map[string]interface{}) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	result = map[string]interface{}{
		"status": status,
		"data":   data,
		"error":  errMsg,
	}

	return
}

// Md5Encode will encode string with md5 method
func Md5Encode(resource string) string {
	src := []byte(resource)
	result := md5.Sum(src)
	return hex.EncodeToString(result[:16])

}

func Sha256Encode(resource string) string {
	result := sha256.Sum256([]byte(resource))
	return hex.EncodeToString(result[:])
}

// GetUID will return uid from header
func GetUID(request *revel.Request) int {
	token := request.GetHttpHeader("athena-token")
	if token == "" {
		return -1
	}
	uidStr := strings.Split(token, "|")[0]

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return -1
	}
	return uid
}

func GetUserIDFromSession(params graphql.ResolveParams) int {
	if revel.DevMode {
		return 5
	}

	controller := GetController(params)
	uidStr := controller.Session["userID"]
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return -1
	}
	return uid
}

// GetController will return controller by params
func GetController(params graphql.ResolveParams) *revel.Controller {
	return params.Context.Value("controller").(*revel.Controller)
}

// Log just a log wrapper
func Log(tag string, err error) {
	if err != nil {
		revel.INFO.Println(tag, err)
	}
}

func GenQiniuToken() string {
	return qn.SetupQiniu()
}
