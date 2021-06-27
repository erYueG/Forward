package viewModel

type BaseViewModel struct {
	Title string
}

func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}
