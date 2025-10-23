package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"path"
	"path/filepath"
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

	storagetype, err := models.GetStorageValue(TenantId)

	if err != nil {

		fmt.Println(err)

		return models.TblStorageType{}, err
	}

	return storagetype, nil

}

func MediaList(c *gin.Context) {

	var (
		limt   int
		offset int
	)

	//get data from html url query
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if limit == "" {
		limt = 40
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	fmt.Println(offset, limit, "offset")

	search := strings.TrimSpace(c.Query("keyword"))

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	_, _, Media, nextcont, _ := GetMedia1(search, "")

	Folder, File, _ := S3ArrPagination(Media, limt, offset)

	fmt.Println(Folder, "totalmediaa")

	ModuleName, TabName, _ := ModuleRouteName(c)

	selectedtype, _ := GetSelectedType()

	var paginationendcount = len(Folder) + len(File) + offset

	paginationstartcount := offset + 1

	fmt.Println(paginationendcount, paginationstartcount, "startendcount")

	Previous, Next, PageCount, Page := Pagination(pageno, int(len(Media)), limt)

	c.HTML(200, "media.html", gin.H{"Menu": menu, "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3}, "title": ModuleName, "csrf": csrf.GetToken(c), "translate": translate, "Folder": Folder, "File": File, "Media": Media, "HeadTitle": "Media", "Count": len(Media), "Mediamenu": true, "Tabmenu": TabName, "StorageType": selectedtype.SelectedType, "Nextcont": nextcont, "Previous": Previous, "Next": Next, "Page": Page, "Limit": limt, "PageCount": PageCount, "CurrentPage": pageno, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "Filter": search, "ListCount": len(Folder) + len(File)})

}

/*ADD FOLDER OR File*/
func AddFolder(c *gin.Context) {

	foldername := c.PostForm("foldername")
	folderpath := c.PostForm("path")

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		_ = storagecontroller.AddFolderMakeDir(foldername, folderpath)
	} else if selectedtype.SelectedType == "aws" {
		_, makedirerr := storagecontroller.CreateFolderToS3(foldername, folderpath)
		if makedirerr != nil {
			ErrorLog.Printf("add folder error: %s", makedirerr)
			json.NewEncoder(c.Writer).Encode("false")
			return
		}

	}

	json.NewEncoder(c.Writer).Encode("true")

}

/*REMOVE FOLDER*/
func DeleteFolderFile(c *gin.Context) {

	deleteid := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(c.PostForm("id"), "[", ""), "]", ""), `"`, "")
	ids := strings.Split(deleteid, ",")
	folderpath := c.PostForm("path")
	selectedtype, _ := GetSelectedType()

	fmt.Println(deleteid, ids, folderpath, "deletedetails")

	var DeleteErr error

	if len(deleteid) > 0 {

		for _, val := range ids {

			if selectedtype.SelectedType == "local" {

				fmt.Println("Local selected")
				DeleteErr = storagecontroller.DeleteImageFolder(folderpath, val)
			} else if selectedtype.SelectedType == "aws" {
				fmt.Println("S3 selected")

				err := storagecontroller.DeleteS3FolderAndContents(val)
				if err != nil {
					fmt.Printf("Error deleting folder and contents %s: %v\n", val, err)
				}

			}
			// } else if selectedtype.SelectedType == "aws" {

			// 	fmt.Println("s3 selected")
			// 	resp1, _ := storagecontroller.ListS3BucketWithPath(val)
			// 	for _, val := range resp1.CommonPrefixes {
			// 		storagecontroller.DeleteS3Files(*val.Prefix)
			// 	}
			// 	DeleteErr = storagecontroller.DeleteS3Files(val)
			// }

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
		nextconf               string
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {
		folderCount, fileCount, Folder, File, err = storagecontroller.FolderDetails(folderpath + "/")
		if err != nil {
			fmt.Println(err)
		}

	} else if selectedtype.SelectedType == "aws" {

		resp1, _ := storagecontroller.ListS3BucketWithPath(folderpath + "/")
		Folder, File = MakeFolderandFileArr(resp1, folderpath, "")
		if resp1.NextContinuationToken != nil {
			nextconf = *resp1.NextContinuationToken
		}

	}

	sort.SliceStable(Folder, func(i, j int) bool {
		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {
		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)
	Media = append(Media, File...)

	json.NewEncoder(c.Writer).Encode(gin.H{"folder": Folder, "file": File, "mediaCount": len(Media), "fileCount": fileCount, "folderCount": folderCount, "Media": Media, "nextcont": nextconf})

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
		return
	}

	var buf bytes.Buffer

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

	if extention == ".gif" {
		gif.Encode(c.Writer, newImage, nil)
		return
	}

	if extention == ".jpeg" || extention == ".jpg" {
		_ = jpeg.Encode(c.Writer, newImage, nil)
	}

}

func AudioResize(c *gin.Context) {
	fileName := c.Query("name")
	filePath := c.Query("path")
	extension := path.Ext(fileName)

	rawObject, err := storagecontroller.GetObjectFromS3(filePath + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the audio file"})
		return
	}
	defer rawObject.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(rawObject.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read audio data"})
		return
	}

	audioData := buf.Bytes()
	switch extension {
	case ".mpeg":
		c.Data(http.StatusOK, "audio/mpeg", audioData)
	case ".mp3":
		c.Data(http.StatusOK, "audio/mpeg", audioData)
	case ".wav":
		c.Data(http.StatusOK, "audio/wav", audioData)
	default:
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported audio format"})
	}
}

func DocumentResize(c *gin.Context) {
	fileName := c.Query("name")
	filePath := c.Query("path")
	extension := path.Ext(fileName)

	rawObject, err := storagecontroller.GetObjectFromS3(filePath + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the document file"})
		return
	}
	defer rawObject.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(rawObject.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read document data"})
		return
	}

	documentData := buf.Bytes()

	switch extension {
	case ".pdf":
		c.Data(http.StatusOK, "application/pdf", documentData)
	case ".doc", ".docx":
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.wordprocessingml.document", documentData)
	case ".txt":
		c.Data(http.StatusOK, "text/plain", documentData)
	case ".zip":
		c.Data(http.StatusOK, "application/zip", documentData)
	default:
		c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported document format"})
	}
}

// make folder and file array
func MakeFolderandFileArr(resp1 *s3.ListObjectsV2Output, parentfoldername string, searchfilter string) (folders []storagecontroller.Medias, files []storagecontroller.Medias) {

	var (
		Folder []storagecontroller.Medias
		File   []storagecontroller.Medias
	)

	if searchfilter != "" {

		for _, val := range resp1.CommonPrefixes {

			if strings.Contains(strings.ToLower(RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))), strings.ToLower(searchfilter)) {

				// obj, _ := storagecontroller.GetObjectFromS3(*val.Prefix)

				var Folde storagecontroller.Medias
				filename := RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))
				Folde.File = true
				Folde.Path = filename
				Folde.AliaseName = *val.Prefix
				Folde.Name = filename
				// Folde.ModTime = *obj.LastModified
				Folder = append(Folder, Folde)
			}

		}

		for _, val := range resp1.Contents {

			if strings.Contains(strings.ToLower(*val.Key), strings.ToLower(searchfilter)) {

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

		}

	} else {

		for _, val := range resp1.CommonPrefixes {

			folderPrefix := *val.Prefix
			var folderCount, imageCount int
			// List objects in the current folder
			resp2, err := storagecontroller.ListS3BucketWithPath(folderPrefix)
			if err != nil {
				// Handle error
				continue
			}

			for _, obj := range resp2.Contents {

				if strings.HasSuffix(*obj.Key, "/") {
					// Increment folder count
					folderCount++
				} else {
					// Increment image count
					imageCount++
				}
			}
			folderCount = len(resp2.CommonPrefixes)

			// obj, _ := storagecontroller.GetObjectFromS3(*val.Prefix)

			var Folde storagecontroller.Medias
			filename := RemoveSpecialCharacter(strings.Replace(*val.Prefix, parentfoldername, "", 1))
			Folde.File = true
			Folde.Path = filename
			Folde.AliaseName = *val.Prefix
			Folde.Name = filename
			Folde.TotalSubMedia = folderCount + imageCount

			// Folde.ModTime = *obj.LastModified
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

	}

	return Folder, File
}

func GetMedia() (Folders []storagecontroller.Medias, Files []storagecontroller.Medias, TotalMedia []storagecontroller.Medias, Nextcontin string, err error) {

	var (
		MediaErr error
		Media    []storagecontroller.Medias
		Folder   []storagecontroller.Medias
		File     []storagecontroller.Medias
		nextcont string
		search   string
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {

		Folder, File, MediaErr = storagecontroller.MediaLocalList(search, "")

		if MediaErr != nil {

			ErrorLog.Println(MediaErr)
		}

	} else if selectedtype.SelectedType == "aws" {

		// resp, _ := storagecontroller.ListS3Bucket()

		resp1, err := storagecontroller.ListS3BucketWithPath("media/")

		nextcont = *resp1.NextContinuationToken

		if err != nil {

			fmt.Println("please check the aws fields")

			return []storagecontroller.Medias{}, []storagecontroller.Medias{}, []storagecontroller.Medias{}, nextcont, err
		}

		Folder, File = MakeFolderandFileArr(resp1, "media/", "")

	}

	sort.SliceStable(Folder, func(i, j int) bool {

		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {

		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)

	Media = append(Media, File...)

	return Folder, File, Media, nextcont, nil

}

func LoadMore(c *gin.Context) {

	offset, _ := strconv.Atoi(c.PostForm("offset"))
	search := c.PostForm("search")
	folderpath := c.PostForm("path")

	var limit = 40
	// offset := limit * offse

	fmt.Println(offset, "offsetvalue")

	_, _, Medias, NextCont, _ := GetMedia1(search, folderpath)

	Folder, File, Media := S3ArrPagination(Medias, limit, offset)

	json.NewEncoder(c.Writer).Encode(gin.H{"folder": Folder, "file": File, "nextcont": NextCont, "Media": Media, "count": len(Medias), "LoadFileCount": len(Folder) + len(File)})
}

func LoadMoreS3(NextCont string) (Folders []storagecontroller.Medias, Files []storagecontroller.Medias, TotalMedia []storagecontroller.Medias, NextContu string, err error) {

	var (
		MediaErr error
		Media    []storagecontroller.Medias
		Folder   []storagecontroller.Medias
		File     []storagecontroller.Medias
		nextcont string
		search   string
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {

		Folder, File, MediaErr = storagecontroller.MediaLocalList(search, "")

		if MediaErr != nil {

			ErrorLog.Println(MediaErr)
		}

	} else if selectedtype.SelectedType == "aws" {

		resp1, err := storagecontroller.LoadMoreListS3BucketWithPath("media/", NextCont)

		if err != nil {

			fmt.Println("please check the aws fields")

			return []storagecontroller.Medias{}, []storagecontroller.Medias{}, []storagecontroller.Medias{}, nextcont, err
		}

		if resp1.NextContinuationToken != nil {

			nextcont = *resp1.NextContinuationToken
		}

		Folder, File = MakeFolderandFileArr(resp1, "media/", "")

	}

	sort.SliceStable(Folder, func(i, j int) bool {

		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {

		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)

	Media = append(Media, File...)

	return Folder, File, Media, nextcont, nil

}

func GetMedia1(search, folderpath string) (Folders []storagecontroller.Medias, Files []storagecontroller.Medias, TotalMedia []storagecontroller.Medias, Nextcontin string, err error) {

	var (
		MediaErr error
		Media    []storagecontroller.Medias
		Folder   []storagecontroller.Medias
		File     []storagecontroller.Medias
		nextcont string
	)

	selectedtype, _ := GetSelectedType()

	if selectedtype.SelectedType == "local" {

		Folder, File, MediaErr = storagecontroller.MediaLocalList(search, folderpath)

		if MediaErr != nil {

			ErrorLog.Println(MediaErr)
		}

	} else if selectedtype.SelectedType == "aws" {

		resp1, err := storagecontroller.ListS3BucketWithPath1("media/" + folderpath)

		if err != nil {

			fmt.Println("please check the aws fields")

			return []storagecontroller.Medias{}, []storagecontroller.Medias{}, []storagecontroller.Medias{}, nextcont, err
		}

		Folder, File = MakeFolderandFileArr(resp1, "media/"+folderpath, search)

	}

	sort.SliceStable(Folder, func(i, j int) bool {

		return Folder[i].ModTime.Unix() > Folder[j].ModTime.Unix()
	})

	sort.SliceStable(File, func(i, j int) bool {

		return File[i].ModTime.Unix() > File[j].ModTime.Unix()
	})

	Media = append(Media, Folder...)

	Media = append(Media, File...)

	return Folder, File, Media, nextcont, nil

}

func S3ArrPagination(Media []storagecontroller.Medias, limit int, offset int) (s []storagecontroller.Medias, f []storagecontroller.Medias, m []storagecontroller.Medias) {

	var (
		Files   []storagecontroller.Medias
		Folders []storagecontroller.Medias
		Medias  []storagecontroller.Medias
	)

	end := limit + offset

	if end > len(Media) {
		end = len(Media)
	}

	for i := offset; i < end; i++ {

		if Media[i].File {

			var Folde storagecontroller.Medias
			Folde.File = true
			Folde.Path = Media[i].Path
			Folde.AliaseName = Media[i].AliaseName
			Folde.TotalSubMedia = Media[i].TotalSubMedia
			Folde.Name = Media[i].Name
			Folders = append(Folders, Folde)

		} else {

			var files storagecontroller.Medias
			files.File = false
			files.Name = Media[i].Name
			files.AliaseName = Media[i].AliaseName
			files.Path = Media[i].Path
			files.ModTime = Media[i].ModTime
			Files = append(Files, files)

		}

		Medias = append(Medias, Media[i])

	}

	return Folders, Files, Medias
}
func RenameMediaPath(c *gin.Context) {

	oldFilename := c.PostForm("oldfilename")
	newFilename := c.PostForm("newfilename")
	filetype := c.PostForm("filetype")

	newFilename1 := filepath.Dir(oldFilename)

	if !strings.HasSuffix(newFilename1, "/") {
		newFilename1 += "/"
	}

	var oldObjectKey string

	var newObjectKey string

	if filetype == "filediv" {

		oldObjectKey = oldFilename
		newObjectKey = newFilename1 + newFilename
	} else if filetype == "folderdiv" {

		oldObjectKey = oldFilename + "/"
		newObjectKey = newFilename1 + newFilename + "/"
	}

	selectedtype, _ := GetSelectedType()

	fmt.Println(oldObjectKey, newObjectKey, "checkrenameee")

	if selectedtype.SelectedType == "local" {
		fmt.Println("Local selected")

	} else if selectedtype.SelectedType == "aws" {
		fmt.Println("s3 selected")
		err1 := storagecontroller.RenameFileS3(oldObjectKey, newObjectKey)

		fmt.Println(err1, "getting error from renameprocess")

	} else if selectedtype.SelectedType == "azure" {

	} else if selectedtype.SelectedType == "drive" {

		fmt.Println("drive selected")
	}
	json.NewEncoder(c.Writer).Encode(true)
}
