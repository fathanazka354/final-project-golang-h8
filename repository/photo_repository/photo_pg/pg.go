package photo_pg

import (
	"database/sql"
	"errors"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/photo_repository"
)

const (
	// Title     string    `json:"title" validate:"required"`
	// Caption   string    `json:"caption"`
	// PhotoUrl  string    `json:"photoUrl" validate:"required"`
	// UserId    int       `json:"userId" validate:"required,numeric"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	createPhotoQuery = `
		INSERT INTO "photos"
		(
			title,
			caption,
			photourl,
			userid
		) VALUES ($1,$2,$3,$4) 
		RETURNING id,title, caption, photourl, userId;
	`

	getPhotosQuery = `
		SELECT "id", "title", "caption", "photourl", "userid", "createdat", "updatedat" FROM "photos"
		ORDER BY "id" ASC
	`

	getPhotoByIdQuery = `
		SELECT "id", "title", "caption", "photourl", "userid", "createdat", "updatedat" FROM "photos"
		WHERE "id" = $1
	`

	updatePhotoByIdQuery = `
		UPDATE "photos" SET "title" = $1, "caption" = $2, "photourl" = $3
		WHERE "id" = $4
	`
	deletePhotoByIdQuery = `
		DELETE FROM "photos"
		WHERE "id" = $1
	`
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.PhotoRepository {
	return &photoPG{db: db}
}

// CreateNewPhoto(payload entity.Photo) errs.MessageErr
// GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
// GetPhotos() ([]*entity.Photo, errs.MessageErr)
// UpdatePhoto(photoId int, payload entity.Photo) errs.MessageErr
// DeletePhoto(photoId int) errs.MessageErr

func (photoPG *photoPG) CreateNewPhoto(payload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	row := photoPG.db.QueryRow(createPhotoQuery, payload.Title, payload.Caption, payload.PhotoUrl, payload.UserId)

	var photo entity.Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &photo, nil
}

func (photoPG *photoPG) GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr) {
	row := photoPG.db.QueryRow(getPhotoByIdQuery, photoId)

	var photo entity.Photo
	// "id", "title", "caption", "photourl", "userid", "createdAt", "updatedAt"
	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo is not exists")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &photo, nil
}

func (photoPG *photoPG) GetPhotos() ([]*entity.Photo, errs.MessageErr) {
	rows, err := photoPG.db.Query(getPhotosQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()

	var photos []*entity.Photo

	for rows.Next() {
		var photo entity.Photo

		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

		if err != nil {
			if err.Error() == sql.ErrNoRows.Error() {
				return nil, errs.NewNotContent("photo still empty")
			}
			return nil, errs.NewInternalServerError("something went wrong")
		}
		photos = append(photos, &photo)
	}
	return photos, nil
}

func (photoPG *photoPG) UpdatePhoto(photoId int, payload entity.Photo) errs.MessageErr {
	_, err := photoPG.db.Exec(updatePhotoByIdQuery, payload.Title, payload.Caption, payload.PhotoUrl, photoId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func (photoPG *photoPG) DeletePhoto(photoId int) errs.MessageErr {
	_, err := photoPG.db.Exec(deletePhotoByIdQuery, photoId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}
