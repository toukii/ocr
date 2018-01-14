package tesseract

import (
	"github.com/everfore/exc"
)

/*
lang: chi_sim eng
*/
func Text(img, lang string) ([]byte, error) {
	if lang == "" {
		lang = "chi_sim"
	}
	return exc.NewCMDf("tesseract %s stdout -l %s", img, lang).Do()
}
