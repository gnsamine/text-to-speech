package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	urll "net/url"
	"os"
	"unsafe"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Response string `json:"response"`
}

func main() {
	app := fiber.New()

	app.Get("/saymyname", func(c *fiber.Ctx) error {

		Queryvalue := c.Query("src")

		url := fmt.Sprintf("https://api.voicerss.org/?key=74bf8ba2bad04a27bd4728dff8386828&hl=tr-tr&c=MP3&src=%s&f=16khz_16bit_stereo", urll.QueryEscape(Queryvalue))

		fmt.Println("url", url)
		req, _ := http.NewRequest("GET", url, nil)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Default()
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		fileName := "voices/" + Queryvalue + ".mp3"

		file, _ := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		bodybyte := *(*[]byte)(unsafe.Pointer(&body))
		bodyReader := bytes.NewReader(bodybyte)
		_, err = io.Copy(file, bodyReader)
		if err != nil {
			fmt.Println(err)

		}

		if string(body) == "ERROR: The text is not specified!" {

			url = "https://api.voicerss.org/?key=74bf8ba2bad04a27bd4728dff8386828&hl=tr-tr&c=MP3&f=16khz_16bit_stereo&src=hooooooop"
			req, _ := http.NewRequest("GET", url, nil)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Default()
			}

			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			return c.Send(body)
		}

		return c.Send(body)

	})

	log.Fatal(app.Listen(":3000"))
}
