package gosseract_test

import "github.com/otiai10/gosseract/v1/gosseract"

import "fmt"
import "image"

func ExampleMust() {
	// TODO: it panics! error handling in *Client.accept
	out := gosseract.Must(gosseract.Params{Src: "./.samples/png/sample002.png", Languages: "eng+heb"})
	fmt.Println(out)
}

func ExampleClient_Src() {
	client, _ := gosseract.NewClient()
	out, _ := client.Src("./samples/png/samples000.png").Out()
	fmt.Println(out)
}

func ExampleClient_Image() {
	client, _ := gosseract.NewClient()
	var img image.Image // any your image instance
	out, _ := client.Image(img).Out()
	fmt.Println(out)
}
