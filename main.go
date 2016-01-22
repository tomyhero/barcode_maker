package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tomyhero/barcode_maker/bcode"
	"image/png"
	"net/url"
	"strconv"
)

// コマンドラインから読み取り値
var flagValue struct {
	Bind string
}

// コマンドラインからの値の受け取り。
func init() {
	flag.StringVar(&flagValue.Bind, "bind", ":9999", "e.g. :8080")
	flag.Parse()
}

func main() {
	/*
		http.HandleFunc("/", handler)
		http.HandleFunc("/image/", barcodeImage)
		http.ListenAndServe(flagValue.Bind, nil)
	*/

	r := gin.Default()
	r.GET("/", handler)
	r.POST("/", handler)
	r.GET("/image/", barcodeImage)
	r.Run(flagValue.Bind) // listen and serve on 0.0.0.0:8080

}

//func handler(w http.ResponseWriter, r *http.Request) {
func handler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")

	tmpl := "<html><head><title>BCODE</title></head><h1>Barcode Maker(code128)</h1><form method='post' action='/'><input type='text' name='code' /><input type='submit' value='Generate' /></form> %s </body></html>"
	if c.Request.Method == "POST" {
		code := c.PostForm("code")
		c.String(200, fmt.Sprintf(tmpl, "<img width='250px' src='/image/?code="+url.QueryEscape(code)+"'/>"))

	} else {
		c.String(200, fmt.Sprintf(tmpl, ""))
	}
}

func barcodeImage(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		panic("no code")
	}
	b := bcode.Bcode{}
	img := b.Generate(code)

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		panic(err)
	}

	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := c.Writer.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}
