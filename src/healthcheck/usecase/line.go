package usecase

import (
	"healthcheck/helper/line"
	"time"

	"fmt"
	"healthcheck/helper/encoding"
	"healthcheck/helper/file"

	"github.com/gin-gonic/gin"
)

// GetLineAuthen -
func (u *Usecase) GetLineAuthen(ctx *gin.Context, code, state string) (err error) {
	res, err := line.GetBearer(code)

	if err == nil {
		en, _ := encoding.Encode(res.AccessToken)
		fileName := "line_access_token.txt"
		path := fmt.Sprintf("public/bearer/%s", fileName)
		file.WriteFile(path, en)
	}

	return err
}

// PostLineAuthen -
func (u *Usecase) PostLineAuthen(ctx *gin.Context) error {
	encode, _ := encoding.Encode(time.Now().UnixNano())

	err := line.Authorization(encode)
	if err != nil {
		return err
	}

	return nil
}
