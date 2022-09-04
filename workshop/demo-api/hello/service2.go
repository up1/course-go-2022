package hello

type HelloService struct {
	Repo IRepository
}

func (hs HelloService) GetDataById(id int) (string, error) {
    return hs.Repo.Get(id)
}



type IRepository interface {
	Get(id int) (string,error)
}

type HelloRepository struct{
}

func (hr HelloRepository) Get(id int) (string,error) {
    res := "result"
	return res, nil
}