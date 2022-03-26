package request

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// TODO: Analizar como conseguir el body de un request
func GetJsonBody(ctx *gin.Context) (interface{}, error) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []byte{}, err
	}
	log.Printf("JSON: %+v", jsonData)

	return jsonData, nil
	/*	var result interface{}
		if err := ctx.BindJSON(&result); err != nil {
			return nil, err
		}
		return result, nil*/
}
