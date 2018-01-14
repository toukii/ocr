package ocr

import (
	"fmt"

	cr "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/toukii/ocr/gosseract"
	"github.com/toukii/ocr/tesseract"
)

var Command = &cobra.Command{
	Use:   "ocr",
	Short: "ocr",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		size := len(args)
		if size > 0 {
			viper.Set("image", args[0])
		} else {
			return
		}
		if err := Excute(); err != nil {
			cr.Red("%+v", err)
		}
	},
}

const (
	ModeGs      = "gs"
	ModeImagick = "imagick"
)

func init() {
	Command.PersistentFlags().StringP("lang", "l", "chi_sim", "lang:chi_sim|eng")
	Command.PersistentFlags().StringP("mode", "m", "tesseract", "mode:tesseract|cgo")

	viper.BindPFlag("image", Command.PersistentFlags().Lookup("image"))
	viper.BindPFlag("lang", Command.PersistentFlags().Lookup("lang"))
	viper.BindPFlag("mode", Command.PersistentFlags().Lookup("mode"))
}

func Excute() error {
	if viper.GetString("mode") == "cgo" {
		return Gosseract()
	}
	return Tesseract()
}

func Tesseract() error {
	fmt.Println("Tesseract")
	bs, err := tesseract.Text(viper.GetString("image"), viper.GetString("lang"))
	if err != nil {
		return err
	}
	cr.Green("%s", bs)
	return nil
}

func Gosseract() error {
	fmt.Println("Gosseract")
	c := gosseract.NewClient()
	c.SetLanguage(viper.GetString("lang"))
	c.SetImage(viper.GetString("image"))

	txt, err := c.Text()
	if err != nil {
		return err
	}

	cr.Green("%s", txt)
	return nil
}
