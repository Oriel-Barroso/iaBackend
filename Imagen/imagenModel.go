package Imagen

import (
	"gorm.io/gorm"
	"image"
	"os"
)

type ImagenStruct struct {
	gorm.Model
	Archivo    *os.File    `json:"archivo"`
	Tipo       string      `json:"tipo"`
	ImagenFile image.Image `json:"imagenFile"`
}

func (t *ImagenStruct) Imagen() image.Image {
	return t.ImagenFile
}

func (t *ImagenStruct) SetImagen(imagen image.Image) {
	t.ImagenFile = imagen
}
