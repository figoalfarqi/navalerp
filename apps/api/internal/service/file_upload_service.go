package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/figoalfarqi/apipml/config"
	"github.com/figoalfarqi/apipml/internal/helper"
)

type FileUploadService struct {
	Cfg *config.Config
}

func NewFileUploadService(c *config.Config) *FileUploadService {
	return &FileUploadService{Cfg: c}
}

func (s *FileUploadService) GeneratePresignedURL(ctx context.Context, folder, originalFilename, action string, userID int) (map[string]string, error) {
	// buat filename sesuai format custom
	newFilename := helper.GenerateFilename(userID, originalFilename)

	// expiry 5 menit
	expiry := time.Now().Add(5 * time.Minute).Unix()

	// ambil secret
	secret := s.Cfg.FileServiceSecretKey
	if secret == "" {
		secret = os.Getenv("SECRET_KEY")
	}

	// generate token
	token := helper.GeneratePresigned(secret, folder, newFilename, action, expiry)

	// return hasil
	return map[string]string{
		"folder":    folder,
		"filename":  newFilename,
		"action":    action,
		"expiry":    strconv.FormatInt(expiry, 10),
		"presigned": token,
	}, nil
}

type FinalizeRequest struct {
	Folder   string `json:"folder"`
	Filename string `json:"filename"`
}

type DeleteRequest struct {
	Folder   string `json:"folder"`
	Filename string `json:"filename"`
}

func (s *FileUploadService) getFolderAndFileName(FullUrl, baseUrl string) (string, string) {
	prefixUrl1 := baseUrl + "/uploads/temp/"
	prefixUrl2 := baseUrl + "/uploads/"
	var path string
	if strings.Contains(FullUrl, prefixUrl1) {
		path = strings.TrimPrefix(FullUrl, prefixUrl1)
	} else if strings.Contains(FullUrl, prefixUrl2) {
		path = strings.TrimPrefix(FullUrl, prefixUrl2)
	} else {
		return "WRONG", "SALAH"
	}

	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]
	folder := strings.Join(parts[:len(parts)-1], "/")
	return folder, filename
}

func (s *FileUploadService) FinalizeFile(baseURL, folder, filename string) error {
	// request body
	reqBody := FinalizeRequest{
		Folder:   folder,
		Filename: filename,
	}
	body, _ := json.Marshal(reqBody)

	// buat request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/finalize_system", baseURL), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	secret := s.Cfg.FileServiceSecretKey
	// tambahkan header auth
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+secret)

	// kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to finalize file %s:, status: %s body: %s", baseURL, resp.Status, respBody)
	}

	return nil
}

func (s *FileUploadService) DeleteFile(baseURL, folder, filename string) error {
	// request body
	reqBody := DeleteRequest{
		Folder:   folder,
		Filename: filename,
	}
	body, _ := json.Marshal(reqBody)

	// buat request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/delete_system", baseURL), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	secret := s.Cfg.FileServiceSecretKey
	// tambahkan header auth
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer "+secret)

	// kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete file, status: %s", resp.Status)
	}

	return nil
}

// err := FinalizeFile("http://fileservice:5000", "superinternal123", "transaction/wings/loading/surat_jalan", "20250927190732fyTEu1RgtxgvHX.jpg")
// if err != nil {
//     fmt.Println("error finalize:", err)
// }

// FinalizeManyFiles memproses banyak file URL sekaligus.
// Parameter berupa map[string]*string (label → pointer URL aslinya)
func (s *FileUploadService) FinalizeManyFiles(items map[string]*string) error {
	baseUrl := s.Cfg.FileServiceUrl
	for key, ptr := range items {

		// Skip jika nil
		if ptr == nil {
			continue
		}

		// Skip jika kosong
		if *ptr == "" || *ptr == "-1" {
			continue
		}

		folder, filename := s.getFolderAndFileName(*ptr, baseUrl)

		err := s.FinalizeFile(baseUrl, folder, filename)
		if err != nil {
			return fmt.Errorf("failed to finalize %s: %s: %w", baseUrl, key, err)
		}

		// Update pointer URL final
		finalURL := baseUrl + "/uploads/" + folder + "/" + filename
		*ptr = finalURL
	}

	return nil
}

// UpdateFile handles delete old file, finalize new file, and update pointer URL.
func (s *FileUploadService) UpdateFile(
	oldURL *string,
	newURL **string, // pointer to pointer -> supaya bisa update nilai
) error {
	baseUrl := s.Cfg.FileServiceUrl
	// --- DELETE FILE LAMA ---
	if oldURL != nil && newURL != nil && !helper.StrPtrEqual(*newURL, oldURL) {
		folder, filename := s.getFolderAndFileName(*oldURL, baseUrl)
		_ = s.DeleteFile(baseUrl, folder, filename) // ignore delete error
	}

	// --- FINALIZE FILE BARU ---
	if newURL != nil && *newURL != nil && !helper.StrPtrEqual(*newURL, oldURL) {
		folder, filename := s.getFolderAndFileName(**newURL, baseUrl)

		if err := s.FinalizeFile(baseUrl, folder, filename); err != nil {
			return err
		}

		// set ke URL final
		finalURL := baseUrl + "/uploads/" + folder + "/" + filename
		**newURL = finalURL
	}

	return nil
}
