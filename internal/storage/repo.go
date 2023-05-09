package storage

type Repo struct {
	U UsersDataRepo
}

func NewRepo(ud UsersDataRepo) *Repo {
	return &Repo{
		U: ud,
	}
}
