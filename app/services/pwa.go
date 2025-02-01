package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"landingConstructor/app/domain/dao"
	"landingConstructor/app/repositories"
	"landingConstructor/app/utils"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PwaService interface {
	CreatePWA(c *gin.Context)
	CreatePreLanding(c *gin.Context)
	SaveImage(c *gin.Context)
	AddScreenshots(c *gin.Context)
}

type PwaServiceImpl struct {
	pwaRepository repositories.PwaRepository
}

func (p PwaServiceImpl) CreatePWA(c *gin.Context) {
	var bodyRequest dao.PwaCreateRequest

	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var value dao.Pwa
	value = dao.Pwa{
		Name:         bodyRequest.Name,
		TypeCampaign: bodyRequest.TypeCampaign,
		Icon:         bodyRequest.Icon,
		BaseModel:    dao.BaseModel{},
	}

	if err := p.pwaRepository.CreatePwa(&value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PWA создан"})
}

func (p PwaServiceImpl) CreatePreLanding(c *gin.Context) {
	var request dao.PreLandingCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	preLanding := dao.PreLandingCreateRequest{
		PwaId:  request.PwaId,
		Design: request.Design,
		Header: dao.Header{
			Name:            request.Header.Name,
			IconStore:       request.Header.IconStore,
			IconApp:         request.Header.IconApp,
			Developer:       request.Header.Developer,
			Subtitle:        request.Header.Subtitle,
			Rating:          request.Header.Rating,
			NumberOfReviews: request.Header.NumberOfReviews,
		},
		Description: dao.Description{
			AboutThisGame: request.Description.AboutThisGame,
			UpdatedOn:     request.Description.UpdatedOn,
			DataSafety:    request.Description.DataSafety,
		},
		Ratings: dao.Ratings{
			One:   request.Ratings.One,
			Two:   request.Ratings.Two,
			Three: request.Ratings.Three,
			Four:  request.Ratings.Four,
			Five:  request.Ratings.Five,
		},
		Comments: request.Comments,
	}

	pwa, err := p.pwaRepository.GetPwaById(preLanding.PwaId.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if preLandErr := p.pwaRepository.CreatePreLanding(&dao.PreLanding{
		PwaId:     pwa.ID,
		Design:    preLanding.Design,
		Header:    preLanding.Header,
		BaseModel: dao.BaseModel{},
	}, &pwa); preLandErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when create prelanding"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pwa": pwa})

}

func (p PwaServiceImpl) SaveImage(c *gin.Context) {
	uploadDir := "./uploads"
	errMsg := utils.CheckOrCreateDirectory(uploadDir)
	if errMsg != nil {
		c.String(http.StatusInternalServerError, *errMsg)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка чтения файла: %s", err.Error()))
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка открытия файла: %s", err.Error()))
		return
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fileContent)

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка чтения файла: %s", err.Error()))
		return
	}

	mimeType := http.DetectContentType(fileBytes)

	extensions, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(extensions) == 0 {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Не удалось определить расширение файла %s", file.Filename))
		return
	}

	var ext string
	switch mimeType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		ext = extensions[0]
	}

	randomFileName := utils.GenerateRandomFileName() + ext
	filePath := filepath.Join(uploadDir, randomFileName)

	err = os.WriteFile(filePath, fileBytes, os.ModePerm)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Ошибка сохранения файла: %s", err.Error()))
		return
	}

	fileURL := fmt.Sprintf("%s/uploads/%s", os.Getenv("HOST"), randomFileName)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Файл %s успешно загружен.", randomFileName),
		"fileUrl": fileURL,
	})
}

func (p PwaServiceImpl) AddScreenshots(c *gin.Context) {
	uploadDir := "./uploads"
	errMsg := utils.CheckOrCreateDirectory(uploadDir)
	if errMsg != nil {
		c.String(http.StatusInternalServerError, *errMsg)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Ошибка чтения данных формы: %s", err.Error()))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "Необходимо загрузить хотя бы один файл.")
		return
	}

	if len(files) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Максимальное количество файлов - 10."})
	}

	results := make(chan string, len(files))
	errs := make(chan error, len(files))

	validMimeTypes := map[string]bool{
		"image/png":  true,
		"image/jpeg": true,
		"image/gif":  true,
	}

	for _, file := range files {
		go func(file *multipart.FileHeader) {
			fileContent, err := file.Open()
			if err != nil {
				errs <- fmt.Errorf("Ошибка открытия файла %s: %s", file.Filename, err.Error())
				return
			}
			defer func(fileContent multipart.File) {
				err := fileContent.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(fileContent)

			fileBytes, err := io.ReadAll(fileContent)
			if err != nil {
				errs <- fmt.Errorf("Ошибка чтения файла %s: %s", file.Filename, err.Error())
				return
			}

			mimeType := http.DetectContentType(fileBytes)
			if !validMimeTypes[mimeType] {
				errs <- fmt.Errorf("Файл %s не является допустимым изображением (требуется PNG, JPEG, GIF).", file.Filename)
				return
			}

			extensions, err := mime.ExtensionsByType(mimeType)
			if err != nil || len(extensions) == 0 {
				errs <- fmt.Errorf("Не удалось определить расширение файла %s", file.Filename)
				return
			}

			var ext string
			switch mimeType {
			case "image/jpeg":
				ext = ".jpg"
			case "image/png":
				ext = ".png"
			case "image/gif":
				ext = ".gif"
			default:
				ext = extensions[0]
			}

			randomFileName := utils.GenerateRandomFileName() + ext
			filePath := filepath.Join(uploadDir, randomFileName)

			err = os.WriteFile(filePath, fileBytes, os.ModePerm)
			if err != nil {
				errs <- fmt.Errorf("Ошибка сохранения файла %s: %s", file.Filename, err.Error())
				return
			}

			fileURL := fmt.Sprintf("%s/uploads/%s", os.Getenv("HOST"), randomFileName)
			results <- fmt.Sprintf("Файл %s успешно загружен. URL: %s", file.Filename, fileURL)
		}(file)
	}

	var messages []string
	var errorMessages []string

	for i := 0; i < len(files); i++ {
		select {
		case result := <-results:
			messages = append(messages, result)
		case err := <-errs:
			errorMessages = append(errorMessages, err.Error())
		}
	}

	if len(errorMessages) > 0 {
		c.String(http.StatusInternalServerError, strings.Join(errorMessages, "\n"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": strings.Join(messages, "\n"),
	})
}

func PwaServiceInit(pwaRepository repositories.PwaRepository) PwaService {
	return &PwaServiceImpl{
		pwaRepository: pwaRepository,
	}
}
