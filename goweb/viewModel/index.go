package viewModel

import "Forward/goweb/model"

type IndexViewModel struct {
	BaseViewModel
	User  *model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (ivmo *IndexViewModelOp) GetView() *IndexViewModel {
	user1 := &model.User{
		UserName: "Owen",
	}
	user2 := &model.User{
		UserName: "Jessica",
	}
	posts := []model.Post{
		model.Post{
			user1,
			"I will commit code every day!",
		},
		model.Post{
			user2,
			"I want to be beautiful every day!",
		},
	}

	ivm := &IndexViewModel{
		BaseViewModel: BaseViewModel{
			Title: "HomePage",
		},
		User:  user1,
		Posts: posts,
	}
	return ivm
}
