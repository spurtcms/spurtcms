package controllers

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"spurt-cms/models"

	// "regexp"
	// "spurt-cms/models"
	"strconv"
	"strings"

	"time"

	// "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	chn "github.com/spurtcms/channels"
	csrf "github.com/utrack/gin-csrf"
	"golang.org/x/net/html"
)

type xlhead struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
}

type xlerr struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	Errmsg      []string
}

type Entrieslog struct {
	Title          string
	Createdon      string
	Totalfields    int
	Mappedfields   int
	Unmappedfields int
	Status         string
}

var (
	header    = make(map[int][]xlhead)
	xlerdata  = make(map[int][]xlerr)
	errflg    = make(map[int]bool)
	imgpath   = make(map[int][]string)
	chennalid = make(map[int]int)
)

func Data(c *gin.Context) {

	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("permissison authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Permissison authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		var flag = true
		channelslist, _, err := ChannelConfig.ListChannel(chn.Channels{Limit: 0, Offset: 0, IsActive: flag, TenantId: TenantId})
		if err != nil {
			ErrorLog.Printf("getchannellist error: %s", err)
		}

		c.HTML(200, "data.html", gin.H{"Menu": menu, "linktitle": "Data", "title": "Data", "csrf": csrf.GetToken(c), "Channellist": channelslist, "channelslist": channelslist, "HeadTitle": "Data", "translate": translate, "SettingsHead": true, "Datamenu": true, "Tooltiptitle": "You may import and export data as and when needed"})
	}

}

func ImportPageTwo(c *gin.Context) {
	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	// permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	// if perr != nil {
	// 	ErrorLog.Printf("permissison authorization error: %s", perr)
	// }

	// if !permisison {
	// 	ErrorLog.Printf("Permissison authorization error")
	// 	c.Redirect(301, "/403-page")
	// 	return
	// }

	// if permisison {

	// 	v
	// }

	c.HTML(200, "data-import-2.html", gin.H{"Menu": menu, "title": "Data", "csrf": csrf.GetToken(c), "translate": translate})

}

func ImportPageThree(c *gin.Context) {
	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	// permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	// if perr != nil {
	// 	ErrorLog.Printf("permissison authorization error: %s", perr)
	// }

	// if !permisison {
	// 	ErrorLog.Printf("Permissison authorization error")
	// 	c.Redirect(301, "/403-page")
	// 	return
	// }

	// if permisison {

	// 	v
	// }

	c.HTML(200, "data-import-3.html", gin.H{"Menu": menu, "title": "Data", "csrf": csrf.GetToken(c), "translate": translate})

}

func ExportPage(c *gin.Context) {
	menu := NewMenuController(c)

	translate, _ := TranslateHandler(c)

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("permissison authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Permissison authorization error")
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		var flag = true
		channelslist, _, err := ChannelConfig.ListChannel(chn.Channels{Limit: 0, Offset: 0, IsActive: flag, TenantId: TenantId})
		if err != nil {
			ErrorLog.Printf("getchannellist error: %s", err)
		}

		c.HTML(200, "data-export.html", gin.H{"Menu": menu,"Channellist":channelslist, "title": "Data", "csrf": csrf.GetToken(c), "translate": translate})
	}

}

// create entries using importdata
func ImportData(c *gin.Context) {

	var (
		xlheads   []xlhead
		xldata    []xlerr
		erflg     bool
		imagepath []string
		chnid     int
	)

	file, _ := c.FormFile("image")

	zipFile, fileerr := file.Open()
	if fileerr != nil {
		ErrorLog.Printf("import data file open error: %s", fileerr)
	}

	buf := new(bytes.Buffer)

	_, err1 := io.Copy(buf, zipFile)
	if err1 != nil {
		ErrorLog.Printf("import data copy error: %s", err1)
	}

	zipReader, err2 := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err2 != nil {
		ErrorLog.Printf("zip file reader error: %s", err2)
	}

	for index, file := range zipReader.File {

		if index > 0 {

			Insidefile, fielerr := file.Open()
			if fielerr != nil {
				ErrorLog.Printf("file open error: %s", fielerr)
			}

			filename := strings.Split(file.Name, "/")[1]
			dest := "storage/media/" + filename
			imagepath = append(imagepath, dest)
			imgpath[c.GetInt("userid")] = imagepath

			filedata, err := os.Create(dest)
			if err != nil {
				ErrorLog.Printf("open error: %s", err)
			}

			io.Copy(filedata, Insidefile)

		}

	}

	chnid, _ = strconv.Atoi(c.PostForm("id"))
	chennalid[c.GetInt("userid")] = chnid

	xlfield := c.PostForm("fielddata")

	unmarshall := json.Unmarshal([]byte(xlfield), &xlheads)
	if unmarshall != nil {
		ErrorLog.Printf("unmarshall error while import data reading : %s", unmarshall)
	}

	header[c.GetInt("userid")] = xlheads
	xldata, erflg = XlsxFilevalidation(xlheads)
	xlerdata[c.GetInt("userid")] = xldata
	errflg[c.GetInt("userid")] = erflg

	if erflg {

		c.SetCookie("redirect_after_download", "true", 3600, "", "", false, false)

	}

	/* File log */

	// if errmsg != nil {

	// 	logdata.FileName = c.PostForm("filename")

	// 	logdata.TotalFields, _ = strconv.Atoi(c.PostForm("totalfields"))

	// 	logdata.MappedFields, _ = strconv.Atoi(c.PostForm("mappedfields"))

	// 	logdata.UnmappedFields = logdata.TotalFields - logdata.MappedFields

	// 	logdata.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	// 	logdata.CreatedBy = c.GetInt("userid")

	// 	logdata.Status = 1
	// }
	// if errmsg == nil {
	// 	logdata.FileName = c.PostForm("filename")

	// 	logdata.TotalFields, _ = strconv.Atoi(c.PostForm("totalfields"))

	// 	logdata.MappedFields, _ = strconv.Atoi(c.PostForm("mappedfields"))

	// 	logdata.UnmappedFields = logdata.TotalFields - logdata.MappedFields

	// 	logdata.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	// 	logdata.CreatedBy = c.GetInt("userid")

	// 	logdata.Status = 0
	// }
	// models.CreateFileLog(&logdata)

	// var filelog []models.TblImportfileLogs

	// var filelogs []models.TblImportfileLogs

	// models.GetFileLogs(&filelog)

	// for _, value := range filelog {

	// 	value.DateString = value.CreatedOn.In(TZONE).Format(Datelayout)

	// 	filelogs = append(filelogs, value)
	// }

	// log.Println("filelog",filelogs)

	c.JSON(200, gin.H{"valid": true})

}

func Entrieslist(c *gin.Context) {

	var (
		limt    int
		offset  int
		filters EntriesFilter
	)

	id, _ := strconv.Atoi(c.Query("id"))
	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	if pageno == 0 {
		pageno = 1
	}

	filters.Keyword = strings.TrimSpace(c.Query("keyword"))
	filters.Status = strings.TrimSpace(c.Query("Status"))
	filters.UserName = strings.TrimSpace(c.Query("Username"))
	filters.Title = strings.TrimSpace(c.Query("Title"))

	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}

	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	chnanem, _, err := ChannelConfig.GetChannelsById(id, TenantId)
	if err != nil {
		log.Println(err)

	}
	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("editchannel authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Permissison authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {
		flag := true
		channelist, _, _ := ChannelConfig.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: flag, TenantId: TenantId})

		/*list*/

		query.DataAccess = c.GetInt("dataaccess")

		query.UserId = c.GetInt("userid")

		chnentry, _, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: limt, Offset: offset, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status}, TenantId)

		_, filtercount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0, Keyword: filters.Keyword, ChannelName: filters.ChannelName, Title: filters.Title, Status: filters.Status}, TenantId)

		_, entrcount, _, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0}, TenantId)

		var allchnentry []chn.Tblchannelentries

		for index, entry := range chnentry {

			entry.Cno = strconv.Itoa((offset + index) + 1)

			entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)

			if !entry.ModifiedOn.IsZero() {

				entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)

			} else {

				entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)

			}

			allchnentry = append(allchnentry, entry)

		}

		paginationendcount := len(chnentry) + offset

		paginationstartcount := offset + 1

		Previous, Next, PageCount, Page := Pagination(pageno, int(filtercount), limt)

		// Breadcrumbs := BreadCrumbs(c.Request.URL.Path)

		translate, _ := TranslateHandler(c)

		c.JSON(200, gin.H{"Pagination": PaginationData{
			CurrectPage:  pageno,
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "chcount1": entrcount, "csrf": csrf.GetToken(c), "channelname": chnanem.ChannelName, "chnid": id, "ChanEntrtlist": allchnentry, "entrycount": entrcount, "Next": Next, "Previous": Previous, "PageCount": PageCount, "CurrentPage": pageno, "limit": limt, "paginationendcount": paginationendcount, "paginationstartcount": paginationstartcount, "filter": filters, "title": "Entries", "heading": chnanem.ChannelName, "translate": translate, "channellist": channelist, "Page": Page, "HeadTitle": "Channels", "chentrycount": filtercount, "Cmsmenu": true, "Entriestab": true})

	}

}

func Export(c *gin.Context) {

	var (
		exportid        int
		exportids       []int
		get_data        []chn.Tblchannelentries
		get_all_data    []chn.Tblchannelentries
		Filechannelname string
		exportentry     []chn.Tblchannelentries
	)

	Categories := ""
	Additionaldata := ""

	const Template = `(<\/?[a-zA-A]+?[^>]*\/?>)*`

	// download entries id

	id := c.Query("exportidarr")

	exid := strings.Split(id, ",")
	for _, ids := range exid {
		exportid, _ = strconv.Atoi(ids)
		exportids = append(exportids, exportid)
	}

	channelid, _ := strconv.Atoi(c.Query("getchannelid"))

	permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("permission authorization error: %s", perr)
	}
	if !permisison {
		ErrorLog.Printf("Permissison authorization error")
		c.Redirect(301, "/403-page")
		return
	}
	if permisison {

		_, FinalSelectedCategories, err := ChannelConfig.GetChannelsById(channelid, TenantId)
		if err != nil {
			ErrorLog.Printf("getchannel error: %s", err)
		}
		for _, val := range FinalSelectedCategories {
			for index, category := range val.Categories {
				if index == len(val.Categories)-1 {
					Categories += category.Category
				} else {
					Categories += category.Category + ","
				}
			}
		}

		exportchannelid, _ := strconv.Atoi(c.Query("exportchannelid"))
		if exportchannelid != 0 {

			channelallentrylist, _, _, cerr := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: exportchannelid}, TenantId)
			if cerr != nil {
				ErrorLog.Printf("getchannellist error: %s", cerr)
			}
			for _, val := range channelallentrylist {

				var chlentry chn.Tblchannelentries
				for index, field := range val.TblChannelEntryField {
					if index == len(val.TblChannelEntryField)-1 {
						Additionaldata += field.FieldValue
					} else {
						Additionaldata += field.FieldValue + ","
					}
					chlentry.AdditionalData = Additionaldata
				}

				var (
					Changestatus      string
					Exportdescription string
				)

				r := regexp.MustCompile(Template)
				originaldesc := val.Description

				Exportdescription = r.ReplaceAllString(originaldesc, "")

				chlentry.Title = val.Title
				chlentry.Username = val.Username
				chlentry.CreatedOn = val.CreatedOn

				if val.Status == 0 {
					Changestatus = "Draft"
				} else if val.Status == 1 {
					Changestatus = "Published"
				} else {
					Changestatus = "Unpublished"
				}

				Filechannelname = val.ChannelName
				chlentry.EntryStatus = Changestatus
				chlentry.Description = Exportdescription
				chlentry.CoverImage = val.CoverImage
				chlentry.ChannelName = val.ChannelName
				chlentry.MetaTitle = val.MetaTitle
				chlentry.Keyword = val.Keyword
				chlentry.Slug = val.Slug
				chlentry.MetaDescription = val.MetaDescription
				chlentry.CategoryGroup = Categories

				get_all_data = append(get_all_data, chlentry)

			}

		}

		if len(exportids) != 0 {

			models.Exportentriesdata(&exportentry, exportids, TenantId)

			for _, val := range exportentry {

				var result chn.Tblchannelentries
				for index, field := range val.TblChannelEntryField {
					if index == len(val.TblChannelEntryField)-1 {
						Additionaldata += field.FieldValue
					} else {
						Additionaldata += field.FieldValue + ","
					}
					result.AdditionalData = Additionaldata
				}

				var (
					Changestatus      string
					Exportdescription string
				)

				r := regexp.MustCompile(Template)

				originaldesc := val.Description
				Exportdescription = r.ReplaceAllString(originaldesc, "")
				result.Title = val.Title
				result.Username = val.Username
				result.CreatedOn = val.CreatedOn

				if val.Status == 0 {
					Changestatus = "Draft"
				} else if val.Status == 1 {
					Changestatus = "Published"
				} else {
					Changestatus = "Unpublished"
				}

				Filechannelname = val.ChannelName
				result.EntryStatus = Changestatus
				result.Description = Exportdescription
				result.CoverImage = val.CoverImage
				result.ChannelName = val.ChannelName
				result.MetaTitle = val.MetaTitle
				result.Keyword = val.Keyword
				result.Slug = val.Slug
				result.MetaDescription = val.MetaDescription
				result.CategoryGroup = Categories

				get_data = append(get_data, result)

			}

		}

		file1 := excelize.NewFile()
		var head_row = map[string]string{"A1": "Entry Title", "B1": "Description", "C1": "ChannelName", "D1": "Status", "E1": "CreateBy", "F1": "Meta Title", "G1": "Meta Description", "H1": "Keyword", "I1": "Slug", "J1": "Categories", "K1": "Additional Data", "L1": "CoverImage", "M1": "CreatedDate"}

		for x, y := range head_row {

			file1.SetCellValue("Sheet1", x, y)
			file1.SetColWidth("Sheet1", "A", "M", 50)
			file1.SetSheetName("Sheet1", "Export Entries")

		}
		var count = 2

		var url_prefix = os.Getenv("DOMAIN_URL")

		url_prefix = url_prefix[:len(url_prefix)-1]

		var coverImages = []string{}

		for x := range get_data {

			file1.SetCellValue("Sheet1", "A"+strconv.Itoa(count), get_data[x].Title)
			file1.SetCellValue("Sheet1", "B"+strconv.Itoa(count), get_data[x].Description)
			file1.SetCellValue("Sheet1", "C"+strconv.Itoa(count), get_data[x].ChannelName)
			file1.SetCellValue("Sheet1", "D"+strconv.Itoa(count), get_data[x].EntryStatus)
			file1.SetCellValue("Sheet1", "E"+strconv.Itoa(count), get_data[x].Username)
			file1.SetCellValue("Sheet1", "F"+strconv.Itoa(count), get_data[x].MetaTitle)
			file1.SetCellValue("Sheet1", "G"+strconv.Itoa(count), get_data[x].MetaDescription)
			file1.SetCellValue("Sheet1", "H"+strconv.Itoa(count), get_data[x].Keyword)
			file1.SetCellValue("Sheet1", "I"+strconv.Itoa(count), get_data[x].Slug)
			file1.SetCellValue("Sheet1", "J"+strconv.Itoa(count), get_data[x].CategoryGroup)
			file1.SetCellValue("Sheet1", "K"+strconv.Itoa(count), get_data[x].AdditionalData)
			file1.SetCellValue("Sheet1", "L"+strconv.Itoa(count), get_data[x].CoverImage)
			file1.SetCellValue("Sheet1", "M"+strconv.Itoa(count), get_data[x].CreatedOn.In(TZONE).Format("02/01/2006"))

			coverImageURL := url_prefix + get_data[x].CoverImage

			// Download the cover image and store it in the coverImages slice
			res, err := http.Get(coverImageURL)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Create a temporary file to store the cover image
			sanitizedFileName := strings.ReplaceAll(get_data[x].CoverImage, "/", "-") // Replace '/' with '-'

			// Create a temporary file to store the cover image with the sanitized filename
			tempFile, err := os.CreateTemp("", sanitizedFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer tempFile.Close()

			// Write the cover image to the temporary file
			if _, err := tempFile.Write(body); err != nil {
				log.Fatal(err)
			}

			// Add the temporary file path to the coverImages slice
			coverImages = append(coverImages, tempFile.Name())

			count++
		}

		for x := range get_all_data {
			file1.SetCellValue("Sheet1", "A"+strconv.Itoa(count), get_all_data[x].Title)
			file1.SetCellValue("Sheet1", "B"+strconv.Itoa(count), get_all_data[x].Description)
			file1.SetCellValue("Sheet1", "C"+strconv.Itoa(count), get_all_data[x].ChannelName)
			file1.SetCellValue("Sheet1", "D"+strconv.Itoa(count), get_all_data[x].EntryStatus)
			file1.SetCellValue("Sheet1", "E"+strconv.Itoa(count), get_all_data[x].Username)
			file1.SetCellValue("Sheet1", "F"+strconv.Itoa(count), get_all_data[x].MetaTitle)
			file1.SetCellValue("Sheet1", "G"+strconv.Itoa(count), get_all_data[x].MetaDescription)
			file1.SetCellValue("Sheet1", "H"+strconv.Itoa(count), get_all_data[x].Keyword)
			file1.SetCellValue("Sheet1", "I"+strconv.Itoa(count), get_all_data[x].Slug)
			file1.SetCellValue("Sheet1", "J"+strconv.Itoa(count), get_all_data[x].CategoryGroup)
			file1.SetCellValue("Sheet1", "K"+strconv.Itoa(count), get_all_data[x].AdditionalData)
			file1.SetCellValue("Sheet1", "L"+strconv.Itoa(count), get_all_data[x].CoverImage)
			file1.SetCellValue("Sheet1", "M"+strconv.Itoa(count), get_all_data[x].CreatedOn.In(TZONE).Format("02/01/2006"))

			coverImageURL := url_prefix + get_all_data[x].CoverImage

			// Download the cover image and store it in the coverImages slice
			res, err := http.Get(coverImageURL)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Create a temporary file to store the cover image
			sanitizedFileName := strings.ReplaceAll(get_all_data[x].CoverImage, "/", "-") // Replace '/' with '-'

			// Create a temporary file to store the cover image with the sanitized filename
			tempFile, err := os.CreateTemp("", sanitizedFileName)
			if err != nil {
				log.Fatal(err)
			}
			defer tempFile.Close()

			// Write the cover image to the temporary file
			if _, err := tempFile.Write(body); err != nil {
				log.Fatal(err)
			}

			// Add the temporary file path to the coverImages slice
			coverImages = append(coverImages, tempFile.Name())

			count++
		}
		zipFile, err := os.Create(Filechannelname + "images.zip")
		if err != nil {
			log.Fatal(err)
		}
		defer zipFile.Close()

		// Create a zip writer to write the cover images to the zip file
		zipWriter := zip.NewWriter(zipFile)
		defer zipWriter.Close()

		// Iterate over the coverImages slice and add each file to the zip file
		for _, file := range coverImages {
			// Open the cover image file
			coverImageFile, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer coverImageFile.Close()

			coverImageFileInfo, err := coverImageFile.Stat()
			if err != nil {
				log.Fatal(err)
			}

			// Create a zip header for the cover image file
			zipHeader, err := zip.FileInfoHeader(coverImageFileInfo)
			if err != nil {
				log.Fatal(err)
			}
			zipHeader.Name = filepath.Base(file)

			// Add the cover image file to the zip file
			writer, err := zipWriter.CreateHeader(zipHeader)
			if err != nil {
				log.Fatal(err)
			}

			// Copy the cover image file to the zip file
			if _, err := io.Copy(writer, coverImageFile); err != nil {
				log.Fatal(err)
			}
		}

		// Close the zip writer and zip file
		if err := zipWriter.Close(); err != nil {
			log.Fatal(err)
		}
		if err := zipFile.Close(); err != nil {
			log.Fatal(err)
		}

		// Print a success message
		fmt.Println("Cover images saved as cover_images.zip")

		// Save the Excel file
		excelFileName := Filechannelname + ".xlsx"
		errR := file1.SaveAs(excelFileName)
		if errR != nil {
			log.Fatal(errR)
		}

		// Combine Excel file and ZIP files into a single ZIP archive
		combinedZipFileName := Filechannelname + ".zip"
		combinedZipFile, err := os.Create(combinedZipFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer combinedZipFile.Close()

		zipWriterCombined := zip.NewWriter(combinedZipFile)

		// Add Excel file to the combined ZIP archive
		excelFile, err := zipWriterCombined.Create(excelFileName)
		if err != nil {
			log.Fatal(err)
		}
		excelData, err := ioutil.ReadFile(excelFileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = excelFile.Write(excelData)
		if err != nil {
			log.Fatal(err)
		}

		// Open the cover images zip file
		coverImagesZipFile, err := os.Open(Filechannelname + "images.zip")
		if err != nil {
			log.Fatal(err)
		}
		defer coverImagesZipFile.Close()

		coverImagesZipFileInfo, err := coverImagesZipFile.Stat()
		if err != nil {
			log.Fatal(err)
		}

		// Create a zip header for the cover images zip file
		coverImagesZipHeader, err := zip.FileInfoHeader(coverImagesZipFileInfo)
		if err != nil {
			log.Fatal(err)
		}
		coverImagesZipHeader.Name = filepath.Base(Filechannelname + "images.zip")

		// Add the cover images zip file to the combined zip file
		coverImagesWriter, err := zipWriterCombined.CreateHeader(coverImagesZipHeader)
		if err != nil {
			log.Fatal(err)
		}

		// Copy the cover images zip file to the combined zip file
		if _, err := io.Copy(coverImagesWriter, coverImagesZipFile); err != nil {
			log.Fatal(err)
		}

		err = zipWriterCombined.Close()
		if err != nil {
			log.Fatal(err)
		}

		c.Header("Content-Disposition", "attachment; filename="+combinedZipFileName+"")

		http.ServeFile(c.Writer, c.Request, combinedZipFileName)

		err = os.Remove(excelFileName)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(combinedZipFileName)
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(Filechannelname + "images.zip")
		if err != nil {
			log.Fatal(err)
		}

		err1 := file1.Write(c.Writer)

		if err1 != nil {

			fmt.Println(err1)
		}

	}

}

func ExtractImgSources(htmlContent string) ([]string, error) {
	var sources []string

	doc, err := html.Parse(strings.NewReader(htmlContent))

	if err != nil {

		return nil, err

	}

	var traverse func(*html.Node)

	traverse = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "img" {

			for _, attr := range n.Attr {

				if attr.Key == "src" {

					sources = append(sources, attr.Val)

				}

			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {

			traverse(c)

		}

	}

	traverse(doc)

	return sources, nil
}

func replaceImageSrc(htmlContent, oldSrc, newSrc string) (string, error) {
	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	// Function to recursively traverse and modify the HTML nodes
	var updateImgSrc func(*html.Node)
	updateImgSrc = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for i, a := range n.Attr {
				if a.Key == "src" && a.Val == oldSrc {
					// Replace the src attribute value
					n.Attr[i].Val = newSrc
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			updateImgSrc(c)
		}
	}

	// Traverse the document tree
	updateImgSrc(doc)

	// Render the modified HTML back to a string
	buf := new(bytes.Buffer)
	err = html.Render(buf, doc)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Xlsx data validation
func XlsxFilevalidation(data []xlhead) ([]xlerr, bool) {

	var xlerrs []xlerr

	var erflg bool

	for _, val := range data {

		var xlerr xlerr

		var errmsg []string

		erflg = false

		if val.Title == "" {

			erflg = true

			errmsg = append(errmsg, "Invalid Title,")

		}
		if val.Description == "" {

			erflg = true

			errmsg = append(errmsg, "Invalid Description,")

		}
		if val.Image == "" {

			erflg = true

			errmsg = append(errmsg, "Invalid Image")

		}

		xlerr.Title = val.Title

		xlerr.Description = val.Description

		xlerr.Image = val.Image

		xlerr.Errmsg = errmsg

		xlerrs = append(xlerrs, xlerr)

	}
	return xlerrs, erflg
}

func DownloadXlsx(c *gin.Context, xldata []xlerr) {

	xlfile := excelize.NewFile()

	var head_row = map[string]string{"A1": "Title", "B1": "Description", "C1": "Image", "D1": "Error"}

	for x, y := range head_row {

		xlfile.SetCellValue("Sheet1", x, y)

		xlfile.SetColWidth("Sheet1", "B", "D", 50)

		xlfile.SetSheetName("Sheet1", "Entries")

	}

	var count = 2

	for _, val := range xldata {

		xlfile.SetCellValue("Sheet1", "A"+strconv.Itoa(count), val.Title)
		xlfile.SetCellValue("Sheet1", "B"+strconv.Itoa(count), val.Description)
		xlfile.SetCellValue("Sheet1", "C"+strconv.Itoa(count), val.Image)
		xlfile.SetCellValue("Sheet1", "D"+strconv.Itoa(count), val.Errmsg)

		count++
	}
	c.Header("Content-Type", "application/octet-stream")

	c.Header("Content-Disposition", "attachment; filename="+"Error.xlsx ")

	c.Header("Content-Transfer-Encoding", "binary")

	err := xlfile.Write(c.Writer)

	log.Println("xldata", xldata)

	log.Println(err)
}

func DownloaderrXlsx(c *gin.Context) {

	var enlogs []Entrieslog

	var enlog Entrieslog

	var errmsg error

	user_id := c.GetInt("userid")

	log.Println("errfilg", errflg)

	if erflg, ok := errflg[user_id]; ok {

		if erflg {

			if xlerr, ok := xlerdata[user_id]; ok {

				DownloadXlsx(c, xlerr)

			}

			return
		}
	}
	if xlheads, ok := header[user_id]; ok {

		for _, val := range xlheads {

			var matchedpath string

			var Description string

			Description = val.Description

			sources, err := ExtractImgSources(val.Description)

			if err != nil {
				fmt.Println("Error extracting image sources:", err)
				return
			}
			permisison, perr := NewAuth.IsGranted("Channels", auth.CRUD, TenantId)
			if perr != nil {
				ErrorLog.Printf("editchannel authorization error: %s", perr)
			}
			if !permisison {
				ErrorLog.Printf("Permissison authorization error")
				c.Redirect(301, "/403-page")
				return
			}
			if permisison {
				for _, src := range sources {

					sr := strings.ReplaceAll(src, `”`, "")

					parts := strings.Split(sr, "/")

					filename := parts[len(parts)-1]

					if imagepath, ok := imgpath[user_id]; ok {

						for _, path := range imagepath {

							file := strings.Split(path, "/")

							imagename := file[len(file)-1]

							var url string

							if os.Getenv("BASE_URL") == "" {

								url = os.Getenv("DOMAIN_URL")

							} else {

								url = os.Getenv("BASE_URL")

							}
							if imagename == filename {

								Description, _ = replaceImageSrc(Description, src, url+path)

							}

						}
					}

				}
				if imagepath, ok := imgpath[user_id]; ok {

					for _, path := range imagepath {

						filearray := strings.Split(path, "/")

						filename := filearray[len(filearray)-1]

						if filename == val.Image {

							matchedpath = "/" + path
						}

					}
				}

				if chnid, ok := chennalid[user_id]; ok {

					entries := chn.EntriesRequired{

						Title:      val.Title,
						Content:    Description,
						ChannelId:  chnid,
						CoverImage: matchedpath,
						CreatedBy:  user_id,
					}
					_, _, errmsg = ChannelConfig.CreateEntry(entries, TenantId)

				}

				if errmsg == nil {

					enlog.Title = val.Title

					enlog.Createdon = time.Now().In(TZONE).Format(Datelayout)

					enlog.Totalfields, _ = strconv.Atoi(c.PostForm("totalfields"))

					enlog.Mappedfields, _ = strconv.Atoi(c.PostForm("mappedfields"))

					enlog.Unmappedfields = enlog.Totalfields - enlog.Mappedfields

					enlog.Status = "Completed"
				}
				if errmsg != nil {

					enlog.Title = val.Title

					enlog.Createdon = time.Now().In(TZONE).Format(Datelayout)

					enlog.Totalfields, _ = strconv.Atoi(c.PostForm("totalfields"))

					enlog.Mappedfields, _ = strconv.Atoi(c.PostForm("mappedfields"))

					enlog.Unmappedfields = enlog.Totalfields - enlog.Mappedfields

					enlog.Status = "Incompleted"

				}

				enlogs = append(enlogs, enlog)

			}
		}
		c.Redirect(301, "/channel/entrylist/")

	}

}
