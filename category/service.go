package category

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/err"
	"github.com/bhuvnesh13396/PlayMySong/common/id"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

var (
	errInvalidArgument = err.New(401, "Invalid Arguments.")
	errNoCategoryFound = err.New(402, "No category found.")
)

type Service interface {
	Get(ctx context.Context, ID string) (category CategoryResp, err error)
	Add(ctx context.Context, title string, category_type string, songIDs []string) (categoryID string, err error)
	List(ctx context.Context) (categories []CategoryResp, err error)
	Update(ctx context.Context, ID string, songsIDs []string) (err error)
}

type service struct {
	songRepo 			model.SongRepo
	categoryRepo 	model.CategoryRepo
}

func NewService(songRepo model.SongRepo, categoryRepo model.CategoryRepo) Service {
	return &service{
		songRepo: 	songRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *service) Get(ctx context.Context, ID string) (category CategoryResp, err error) {
	if len(ID) < 1 {
		err = errInvalidArgument
		return
	}

	dbcategory, err := s.categoryRepo.Get(ID)
	if err != nil {
		err = errNoCategoryFound
		return
	}

	songIDs := dbcategory.SongIDs
	songsIncategory := make([]Song, 0)

	for _, songID := range songIDs {
		song, err := s.songRepo.Get1(songID)
		if err != nil {
			return CategoryResp {}, err
		}

		songResp := Song{
			ID:     song.ID,
			Title:  song.Title,
			Length: song.Length,
			Path:   song.Path,
			Image:  song.Image,
		}

		songsIncategory = append(songsIncategory, songResp)
	}

	category = CategoryResp{
		Title:       dbcategory.Title,
		Type:        dbcategory.Type,
		Songs:       songsIncategory,
	}

	return
}

func (s *service) Add(ctx context.Context, title string, category_type string, songIDs []string) (categoryID string, err error) {
	category := model.Category{
		ID:          id.New(),
		Title:       title,
		Type:        category_type,
		SongIDs:     songIDs,
	}

	err = s.categoryRepo.Add(category)
	if err != nil {
		return
	}

	return category.ID, nil
}

func (s *service) List(ctx context.Context) (categories []CategoryResp, err error) {
	dbcategories, err := s.categoryRepo.List()
	if err != nil {
		return
	}

	for _, category := range dbcategories {
		songIDs := category.SongIDs
		songsIncategory := make([]Song, 0)

		for _, songID := range songIDs {
			song, err := s.songRepo.Get1(songID)
			if err != nil {
				return nil, err
			}

			songResp := Song{
				ID:     song.ID,
				Title:  song.Title,
				Length: song.Length,
				Path:   song.Path,
				Image:  song.Image,
			}

			songsIncategory = append(songsIncategory, songResp)
		}

		newcategory := CategoryResp{
			Title:       category.Title,
			Type:        category.Type,
			Songs:       songsIncategory,
		}

		categories = append(categories, newcategory)
	}

	return
}

func (s *service) Update(ctx context.Context, ID string, songIDs []string) (err error) {
	err = s.categoryRepo.Update(ID, songIDs)
	return
}
