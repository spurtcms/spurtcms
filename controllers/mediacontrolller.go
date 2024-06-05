package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"path"
	"sort"
	"spurt-cms/models"
	storagecontroller "spurt-cms/storage-controller"

	"strconv"

	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	csrf "github.com/utrack/gin-csrf"
)

type medias struct {
	File          bool
	Name          string
	Path          string
	ModTime       time.Time
	TotalSubMedia int
}

// var StorageType = os.Getenv("STORAGE_TYPE")

func GetSelectedType() (storagetyp models.TblStorageType, err error) {

	storagetype, err := models.GetStorageValue()

	if err != nil {

		fmt.Println(err)

		return models.TblStorageType{}, err
	}

	return storagetype, nil

}

func MediaList(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	Folder, File, Media, _ := GetMedia()
	ModuleName, TabName, _ := ModuleRouteName(c)

	selectedtype, _ := GetSelectedType()

	c.HTML(200, "media.html", gin.H{"Menu": menu, "title": ModuleName, "csrf": csrf.GetToken(c), "translate": translate, "Folder": Folder, "File": File, "Media": Media, "HeadTitle": "Media", "Count": len(Media), "Mediamenu": true, "Tabmenu": TabName, "StorageType": selectedtype.SelectedType})

}

/*ADD FOLDER OR File*/
func AddFolder(c *gin.Context) {

	foldername := c.PostForm("foldername")
	folderpath := c.PostForm("path")

	var makedirerr error

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		makedirerr = storagecontroller.AddFolderMakeDir(foldername, folderpath)
	} else if selectedtype.SelectedType == "aws" {
		makedirerr = storagecontroller.CreateFolderToS3(foldername, folderpath)
	}

	if makedirerr != nil {
		ErrorLog.Printf("add folder error: %s", makedirerr)
		json.NewEncoder(c.Writer).Encode("false")
		return
	}

	json.NewEncoder(c.Writer).Encode("true")

}

/*REMOVE FOLDER*/
func DeleteFolderFile(c *gin.Context) {

	deleteid := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(c.PostForm("id"), "[", ""), "]", ""), `"`, "")
	ids := strings.Split(deleteid, ",")
	folderpath := c.PostForm("path")
	selectedtype, _ := GetSelectedType()

	var DeleteErr error

	if len(deleteid) > 0 {

		for _, val := range ids {

			if selectedtype.SelectedType == "local" {

				fmt.Println("Local selected")
				DeleteErr = storagecontroller.DeleteImageFolder(folderpath, val, &AUTH)
			} else if selectedtype.SelectedType == "aws" {

				fmt.Println("s3 selected")
				resp1, _ := storagecontroller.ListS3BucketWithPath(val)
				for _, val := range resp1.CommonPrefixes {
					storagecontroller.DeleteS3Files(*val.Prefix)
				}
				DeleteErr = storagecontroller.DeleteS3Files(val)
			}

			if DeleteErr != nil {
				json.NewEncoder(c.Writer).Encode(false)
				return

			}
		}
	}

	json.NewEncoder(c.Writer).Encode(true)

}

/*Individual Folder Details*/
func FolderDetails(c *gin.Context) {

	folderpath := c.PostForm("path")

	var (
		Media                  []storagecontroller.Medias
		Folder                 []storagecontroller.Medias
		File                   []storagecontroller.Medias
		err                    error
		folderCount, fileCount int
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		folderCount, fileCount, Folder, File, err = storagecontroller.FolderDetails(folderpath + "/")
		if err != nil {
			fmt.Println(err)
		}

	} else if selectedtype.SelectedType == "aws" {
		resp1, _ := storagecontroller.ListS3BucketWithPath(folderpath + "/")
		Folder, File = MakeFolderandFileArr(resp1, folderpath)
	}

	sort.SliceStable(Folder, func(i, j int) bool {
		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {
		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)
	Media = append(Media, File...)

	json.NewEncoder(c.Writer).Encode(gin.H{"folder": Folder, "file": File, "mediaCount": len(Media), "fileCount": fileCount, "folderCount": folderCount, "Media": Media})

}

/*Upload Image*/
func UploadImageMedia(c *gin.Context) {

	imageData := c.Request.PostFormValue("file")
	imagename := c.Request.PostFormValue("name")

	var Path string

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		Path = storagecontroller.UploadCropImage(imageData, imagename)
	} else if selectedtype.SelectedType == "aws" {
		storagecontroller.StoreS3Base64(strings.ReplaceAll(imageData, "data:image/jpeg;base64,", ""), "media/", imagename)
	}

	json.NewEncoder(c.Writer).Encode(gin.H{"Path": "/" + Path + imagename})

}

/*Upload system image */
func UploadImage(c *gin.Context) {

	var err error

	file, fileheader, err := c.Request.FormFile("image")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err))
		return
	}

	if fileheader.Size > int64(maxSize) {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File size exceeds 8MB"})
		return
	}

	var UploadErr error

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		fmt.Println("Local selected")
		UploadErr = storagecontroller.UploadImageLocal(file, fileheader, c.Request.PostFormValue("path"), c)

	} else if selectedtype.SelectedType == "aws" {
		fmt.Println("s3 selected")
		UploadErr = storagecontroller.UploadFileS3(file, fileheader, c.Request.PostFormValue("path"))

	} else if selectedtype.SelectedType == "azure" {

	} else if selectedtype.SelectedType == "drive" {

		fmt.Println("drive selected")
	}

	if UploadErr != nil {
		json.NewEncoder(c.Writer).Encode(false)
		return
	}

	json.NewEncoder(c.Writer).Encode(true)

}

// view image
func ResizeImage(c *gin.Context) {

	fileName := c.Query("name")
	filePath := c.Query("path")
	width, _ := strconv.ParseUint(c.Query("width"), 10, 64)
	height, _ := strconv.ParseUint(c.Query("height"), 10, 64)
	extention := path.Ext(fileName)

	rawObject, err := storagecontroller.GetObjectFromS3(filePath + fileName)

	if err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)

	buf.ReadFrom(rawObject.Body)

	if extention == ".svg" {
		svgData := buf.String()
		c.Data(http.StatusOK, "image/svg+xml", []byte(svgData))
		return
	}

	Image, _, erri := image.Decode(bytes.NewReader(buf.Bytes()))
	if erri != nil {
		fmt.Println(erri)
		return
	}

	newImage := resize.Resize(uint(width), uint(height), Image, resize.Lanczos3)
	if extention == ".png" {
		png.Encode(c.Writer, newImage)
		return
	}

	if extention == ".jpeg" || extention == ".jpg" {
		_ = jpeg.Encode(c.Writer, newImage, nil)
	}

}

// make folder and file array
func MakeFolderandFileArr(resp1 *s3.ListObjectsV2Output, parentfoldername string) (folders []storagecontroller.Medias, files []storagecontroller.Medias) {

	var (
		Folder []storagecontroller.Medias
		File   []storagecontroller.Medias
	)

	for _, val := range resp1.CommonPrefixes {

		var Folde storagecontroller.Medias
		filename := RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))
		Folde.File = true
		Folde.Path = filename
		Folde.AliaseName = *val.Prefix
		Folde.Name = filename
		Folder = append(Folder, Folde)

	}

	for _, val := range resp1.Contents {

		if strings.Contains(*val.Key, ".jpeg") || strings.Contains(*val.Key, ".png") || strings.Contains(*val.Key, ".jpg") || strings.Contains(*val.Key, ".svg") {

			var file storagecontroller.Medias
			filename := strings.Replace(*val.Key, parentfoldername, "", 1)
			Lastupdate := *val.LastModified
			file.File = false
			file.Name = filename
			file.AliaseName = *val.Key
			file.Path = parentfoldername
			file.ModTime = Lastupdate
			File = append(File, file)

		}

	}

	return Folder, File
}

func GetMedia() (Folders []storagecontroller.Medias, Files []storagecontroller.Medias, TotalMedia []storagecontroller.Medias, err error) {

	var (
		MediaErr error
		Media    []storagecontroller.Medias
		Folder   []storagecontroller.Medias
		File     []storagecontroller.Medias
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {

		Folder, File, MediaErr = storagecontroller.MediaLocalList()

		if MediaErr != nil {

			ErrorLog.Println(MediaErr)
		}

	} else if selectedtype.SelectedType == "aws" {

		// resp, _ := storagecontroller.ListS3Bucket()

		resp1, err := storagecontroller.ListS3BucketWithPath("media/")

		if err != nil {

			fmt.Println("please check the aws fields")

			return []storagecontroller.Medias{}, []storagecontroller.Medias{}, []storagecontroller.Medias{}, err
		}

		Folder, File = MakeFolderandFileArr(resp1, "media/")

	}

	sort.SliceStable(Folder, func(i, j int) bool {

		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {

		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)

	Media = append(Media, File...)

	return Folder, File, Media, nil

}
