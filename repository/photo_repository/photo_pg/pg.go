package photo_pg

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
	`

	getPhotosQuery = `
		SELECT "id", "title", "caption", "photourl", "userid", "createdAt", "updatedAt" FROM "photos"
		ORDER BY "id" ASC
	`

	getPhotoByIdQuery = `
		SELECT "id", "title", "caption", "photourl", "userid", "createdAt", "updatedAt" FROM "photos"
		WHERE "id" = $1
	`

	updatePhotoByIdQuery = `
		UPDATE "photos" SET "title" = $1, "caption" = $2, "photourl" = $3
		WHERE "id" = $4
	`
)
