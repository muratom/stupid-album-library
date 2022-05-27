package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"web-service-gin/models"
)

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (models.Album, error) {
	// An album to hold data from the returned row.
	var alb models.Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func allAlbums() ([]models.Album, error) {
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("allAlbums: %v", err)
	}
	defer rows.Close()

	albums := make([]models.Album, 0)
	for rows.Next() {
		var alb models.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("allAlbums: %v", err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("allAlbums: %v", err)
	}

	return albums, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb models.AlbumRequest) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// GetAlbums godoc
// @Summary Retrieves all albums from DB
// @Produce json
// @Success 200 {object} models.Album
// @Router /albums [get]
func GetAlbums(c *gin.Context) {
	albums, err := allAlbums()
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumById godoc
// @Summary Retrieves album based on given ID
// @Produce json
// @Param id path integer true "Album ID"
// @Success 200 {object} models.Album "OK"
// @Failure	500 					  "Fail"
// @Router /albums/{id} [get]
func GetAlbumById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	album, err := albumByID(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// PostAlbums adds an album from JSON received in the request body.
// GetAlbumsById godoc
// @Summary Adds an album from JSON
// @Description  Adds an album from JSON
// @Accept       json
// @Produce      json
// @Param        album    body      models.AlbumRequest   true  "Album info"
// @Success      201      {object}  models.Album    	  true  "created"
// @Failure      500         				              "fail"
// @Router       /albums [post]
// @Security ApiKeyAuth
func PostAlbums(c *gin.Context) {
	var newAlbum models.AlbumRequest

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.Error(err)
		return
	}

	// Add the new album to the slice.
	_, err := addAlbum(newAlbum)
	if err != nil {
		c.Error(err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
