package DTOs

type ImageDTO struct {
	Image string
}

func (d *ImageDTO) setImage(path string) {
	d.Image = path
}
