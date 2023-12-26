package progress_bar

type Size int

func (size Size) Equal(compareSize Size) bool {
	return size == compareSize
}

type Fill Size

func (fill Fill) Equal(fillSize Fill) bool {
	return fill == fillSize
}

type ProgressBar struct {
	Size  Size
	Fills []Fill
}

func Create(size Size) ProgressBar {
	return ProgressBar{
		Size: size,
	}
}

// Fill returns a new ProgressBar with the additional fill
func (bar ProgressBar) Fill(fills ...Fill) ProgressBar {
	bar.Fills = append(bar.Fills, fills...)
	return bar
}
