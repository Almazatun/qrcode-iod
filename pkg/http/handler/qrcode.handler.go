package handler

import (
	"bytes"
	"encoding/base64"
	"net/url"

	"github.com/Almazatun/qrcode-iod/pkg/common"
	"github.com/Almazatun/qrcode-iod/pkg/common/helper"
	"github.com/Almazatun/qrcode-iod/pkg/common/input"
	"github.com/gofiber/fiber/v2"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type QRCodeHandlerInstance struct{}

type AuthHandler interface {
	Create(c *fiber.Ctx) error
}

func (qr *QRCodeHandlerInstance) Create(c *fiber.Ctx) error {
	input := input.CreateQrcodeInput{}

	//  Parse body into product struct
	if err := c.BodyParser(&input); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})

		return nil
	}

	_, err := url.ParseRequestURI(input.Url)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})

		return nil
	}

	qrc, err := qrcode.NewWith(input.Url,
		qrcode.WithEncodingMode(qrcode.EncModeByte),
		qrcode.WithErrorCorrectionLevel(qrcode.ErrorCorrectionQuart),
	)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})

		return nil
	}

	fileName := helper.RandomString(32) + ".png"
	// dir := "assets/qrode_imgs/"

	// get bytes
	buf := bytes.NewBuffer(nil)
	wr := common.NopCloser{Writer: buf}
	w2 := standard.NewWithWriter(wr, standard.WithQRWidth(40))

	if err = qrc.Save(w2); err != nil {
		panic(err)
	}

	// save into file if need it
	// w, err := standard.New((dir + fileName), standard.WithQRWidth(20))

	if err != nil {
		panic(err)
	}

	// Convert the image content to base64 string
	imageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	return c.Status(200).JSON(fiber.Map{
		"image":    imageBase64,
		"filename": fileName,
	})

}
