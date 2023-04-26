package social_media_pg

import (
	"database/sql"
	"errors"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/repository/social_media_repository"
)

const (
	createSocialMediaQuery = `
		INSERT INTO "socialMedias"
		(
			name,
			social_media_url,
			userid
		)
		VALUES ($1, $2, $3)
	`

	getSocialMediaByIdQuery = `
		SELECT "id", "name", "social_media_url", "userid", "createdat", "updatedat" FROM "socialMedias"
		WHERE "id" = $1
	`

	getSocialMediaQuery = `
		SELECT "id", "name", "social_media_url", "userid", "createdat", "updatedat" FROM "socialMedias"
		ORDER BY "id" ASC
	`

	updateSocialMediaQuery = `
		UPDATE "socialMedias" SET "name" = $1, "social_media_url" = $2 
		WHERE "id" = $3
	`
	deleteSocialMediaQuery = `
		DELETE FROM "socialMedias" 
		WHERE "id" = $1
	`
)

type socialMediaPG struct {
	db *sql.DB
}

// CreateSocialMedia(newProduct entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.MessageErr)
//
//	GetSocialMedias() (*dto.SocialMediaResultResponse, errs.MessageErr)
//	GetSocialMediaById(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr)
//	UpdateSocialMedia(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr)
//	DeleteSocialMedia(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr)
func NewSocialMediaPG(db *sql.DB) social_media_repository.SocialMediaRepository {
	return &socialMediaPG{db: db}
}
func (socialMediaPG *socialMediaPG) CreateSocialMedia(newProduct entity.SocialMedia) errs.MessageErr {
	_, err := socialMediaPG.db.Exec(createSocialMediaQuery, newProduct.Name, newProduct.SocialMediaUrl, newProduct.UserId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (socialMediaPG *socialMediaPG) GetSocialMedias() ([]*entity.SocialMedia, errs.MessageErr) {
	rows, err := socialMediaPG.db.Query(getSocialMediaQuery)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()
	var socialMedias []*entity.SocialMedia

	for rows.Next() {
		// "id", "name", "social_media_url", "userid", "createdat", "updatedat"
		var socialMedia entity.SocialMedia
		err = rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		socialMedias = append(socialMedias, &socialMedia)
	}

	return socialMedias, nil
}

func (socialMediaPG *socialMediaPG) GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr) {
	row := socialMediaPG.db.QueryRow(getSocialMediaByIdQuery, socialMediaId)

	var socialMedia entity.SocialMedia

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &socialMedia, nil
}
func (socialMediaPG *socialMediaPG) UpdateSocialMedia(payload entity.SocialMedia) errs.MessageErr {
	_, err := socialMediaPG.db.Exec(updateSocialMediaQuery, payload.Name, payload.SocialMediaUrl, payload.Id)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (socialMediaPG *socialMediaPG) DeleteSocialMedia(socialMediaId int) errs.MessageErr {
	_, err := socialMediaPG.db.Exec(deleteSocialMediaQuery, socialMediaId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
