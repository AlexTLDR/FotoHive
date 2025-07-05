package models

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Image struct {
	GalleryID int
	Path      string
	Filename  string
}

type Gallery struct {
	ID     int
	UserID int
	Title  string
}

type GalleryService struct {
	DB *sql.DB
	// used to tell the GalleryService where to store/locate images. If not set, defaults to "images/"
	ImagesDir string
}

func (service *GalleryService) Create(title string, userID int) (*Gallery, error) {
	gallery := Gallery{
		Title:  title,
		UserID: userID,
	}
	row := service.DB.QueryRow(`
		INSERT INTO galleries (title, user_id) VALUES ($1, $2) RETURNING id;`, gallery.Title, gallery.UserID)
	err := row.Scan(&gallery.ID)
	if err != nil {
		return nil, fmt.Errorf("create gallery: %w", err)
	}
	return &gallery, nil
}

func (service *GalleryService) ByID(id int) (*Gallery, error) {
	// TODO: Add validation on the passed in ID
	gallery := Gallery{
		ID: id,
	}
	row := service.DB.QueryRow(`
		SELECT title, user_id FROM galleries WHERE id = $1;`, gallery.ID)
	err := row.Scan(&gallery.Title, &gallery.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("query gallery by id: %w", err)
	}
	return &gallery, nil
}

func (service *GalleryService) ByUserID(userID int) ([]Gallery, error) {
	rows, err := service.DB.Query(`
		SELECT id, title FROM galleries WHERE user_id = $1;`, userID)
	if err != nil {
		return nil, fmt.Errorf("query galleries by user id: %w", err)
	}
	var galleries []Gallery
	for rows.Next() {
		gallery := Gallery{
			UserID: userID,
		}
		err := rows.Scan(&gallery.ID, &gallery.Title)
		if err != nil {
			return nil, fmt.Errorf("query galleries by user id: %w", err)
		}
		galleries = append(galleries, gallery)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("query galleries by user id: %w", rows.Err())
	}
	return galleries, nil
}

func (service *GalleryService) Update(gallery *Gallery) error {
	_, err := service.DB.Exec(`
		UPDATE galleries SET title = $1 WHERE id = $2;`, gallery.Title, gallery.ID)
	if err != nil {
		return fmt.Errorf("update gallery: %w", err)
	}
	return nil
}

func (service *GalleryService) Delete(id int) error {
	_, err := service.DB.Exec(`
		DELETE FROM galleries WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("delete gallery: %w", err)
	}
	dir := service.galleryDir(id)
	err = os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("delete gallery: %w", err)
	}
	return nil
}

func (service *GalleryService) Images(galleryID int) ([]Image, error) {
	globPattern := filepath.Join(service.galleryDir(galleryID), "*")
	allFiles, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, fmt.Errorf("glob retrieving gallery images: %w", err)
	}
	var images []Image
	for _, file := range allFiles {
		if hasExtension(file, service.extensions()) {
			images = append(images, Image{
				GalleryID: galleryID,
				Path:      file,
				Filename:  filepath.Base(file),
			})
		}
	}
	return images, nil
}

func (service *GalleryService) Image(galleryID int, filename string) (Image, error) {
	imagePath := filepath.Join(service.galleryDir(galleryID), filename)
	_, err := os.Stat(imagePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Image{}, ErrNotFound
		}
		return Image{}, fmt.Errorf("querying for image: %w", err)
	}
	return Image{
		Filename:  filename,
		Path:      imagePath,
		GalleryID: galleryID,
	}, nil

}

func (service *GalleryService) CreateImage(galleryID int, filename string, contents io.Reader) error {
	// Sanitize and validate filename first
	safeFilename, err := service.validateAndSanitizeFilename(filename)
	if err != nil {
		return fmt.Errorf("invalid filename: %w", err)
	}

	readBytes, err := checkContentType(contents, service.imageContentTypes())
	if err != nil {
		return fmt.Errorf("creating image %v: %w", safeFilename, err)
	}
	err = checkExtension(safeFilename, service.extensions())
	if err != nil {
		return fmt.Errorf("creating image %v: %w", safeFilename, err)
	}

	galleryDir := service.galleryDir(galleryID)
	err = os.MkdirAll(galleryDir, 0755)
	if err != nil {
		return fmt.Errorf("creating gallery-%d directory: %w", galleryID, err)
	}

	// Create the safe image path and double-check it's within gallery directory
	imagePath, err := service.createSafeImagePath(galleryDir, safeFilename)
	if err != nil {
		return fmt.Errorf("creating safe image path: %w", err)
	}

	dst, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("creating image: %w", err)
	}
	defer dst.Close()

	completeFile := io.MultiReader(
		bytes.NewReader(readBytes),
		contents,
	)
	_, err = io.Copy(dst, completeFile)
	if err != nil {
		return fmt.Errorf("copying to image: %w", err)
	}
	return nil
}

func (service *GalleryService) CreateImageViaURL(galleryID int, url string) error {
	filename := path.Base(url)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("downloading image: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("downloading image: invalid status code %d", resp.StatusCode)
	}
	return service.CreateImage(galleryID, filename, resp.Body)
}

func (service *GalleryService) DeleteImage(galleryID int, filename string) error {
	// Validate and sanitize filename
	safeFilename, err := service.validateAndSanitizeFilename(filename)
	if err != nil {
		return fmt.Errorf("invalid filename: %w", err)
	}

	image, err := service.Image(galleryID, safeFilename)
	if err != nil {
		return fmt.Errorf("deleting image: %w", err)
	}

	err = os.Remove(image.Path)
	if err != nil {
		return fmt.Errorf("deleting image: %w", err)
	}
	return nil
}

func (service *GalleryService) extensions() []string {
	return []string{".jpg", ".png", ".jpeg", ".gif"}
}

func (service *GalleryService) imageContentTypes() []string {
	return []string{"image/png", "image/jpeg", "image/gif"}
}

func (service *GalleryService) galleryDir(id int) string {
	imagesDir := service.ImagesDir
	if imagesDir == "" {
		imagesDir = "images"
	}
	return filepath.Join(imagesDir, fmt.Sprintf("gallery-%d", id))
}

func hasExtension(file string, extensions []string) bool {
	for _, ext := range extensions {
		file = strings.ToLower(file)
		ext = strings.ToLower(ext)
		if filepath.Ext(file) == ext {
			return true
		}
	}
	return false
}

// Helper functions
// validateAndSanitizeFilename performs comprehensive filename validation
func (service *GalleryService) validateAndSanitizeFilename(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("filename cannot be empty")
	}

	// Extract just the filename, removing any directory components
	safeFilename := filepath.Base(filename)

	// Additional sanitization
	safeFilename = strings.ReplaceAll(safeFilename, "..", "")
	safeFilename = strings.TrimSpace(safeFilename)

	// Validate filename characteristics
	if safeFilename == "" || safeFilename == "." || safeFilename == ".." {
		return "", fmt.Errorf("invalid filename after sanitization")
	}

	// Check for hidden files (starting with dot)
	if strings.HasPrefix(safeFilename, ".") {
		return "", fmt.Errorf("hidden files not allowed")
	}

	// Validate length (reasonable limit)
	if len(safeFilename) > 255 {
		return "", fmt.Errorf("filename too long")
	}

	// Check for null bytes and other control characters
	if strings.ContainsAny(safeFilename, "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f") {
		return "", fmt.Errorf("filename contains invalid characters")
	}

	return safeFilename, nil
}

// createSafeImagePath creates and validates the complete image path
func (service *GalleryService) createSafeImagePath(galleryDir, filename string) (string, error) {
	// Ensure gallery directory is absolute
	absGalleryDir, err := filepath.Abs(galleryDir)
	if err != nil {
		return "", fmt.Errorf("getting absolute gallery directory: %w", err)
	}

	// Create the target path
	imagePath := filepath.Join(absGalleryDir, filename)

	// Get absolute target path
	absImagePath, err := filepath.Abs(imagePath)
	if err != nil {
		return "", fmt.Errorf("getting absolute image path: %w", err)
	}

	// Verify the target path is within the gallery directory
	relPath, err := filepath.Rel(absGalleryDir, absImagePath)
	if err != nil {
		return "", fmt.Errorf("calculating relative path: %w", err)
	}

	// Check if the relative path goes outside the gallery directory
	if strings.HasPrefix(relPath, "..") || strings.Contains(relPath, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("path traversal attempt detected")
	}

	// Ensure the path doesn't contain any directory separators (should be just filename)
	if strings.Contains(relPath, string(filepath.Separator)) {
		return "", fmt.Errorf("subdirectories not allowed")
	}

	return absImagePath, nil
}
