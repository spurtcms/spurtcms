package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	"os"
	"spurt-cms/models"

	storagecontroller "spurt-cms/storage-controller"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
	"github.com/spurtcms/channels"
	chn "github.com/spurtcms/channels"
	csrf "github.com/utrack/gin-csrf"
)

type ResponseData struct {
	BlockCount          int                         `json:"block_count"`
	DefaultList         []models.TblMstrBlocks      `json:"default_list"`
	FinalblockList      []models.TblMstrBlocks      `json:"finalblock_list"`
	AllList             []models.TblMstrBlocks      `json:"all_list"`
	DefaultctaList      []models.TblMstrForms       `json:"defaultcta_list"`
	FinalctaList        []models.TblMstrForms       `json:"finalcta_list"`
	AllctaList          []models.TblMstrForms       `json:"allcta_list"`
	Channellist         string                      `json:"channellist"`
	Channeldetail       models.TblMstrchannel       `json:"channeldetail"`
	IntegrationList     []models.TblMstrIntegration `json:"integration_list"`
	IntegrationDetailst models.TblMstrIntegration   `json:"integrationdetail"`
}

// collection list return datatype json (entries layout list)
func CollectionList(c *gin.Context) {

	var (
		filter blocks.Filter
	)

	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		collectionlist, _, err := BlockConfig.CollectionList(blocks.Filter(filter), TenantId, "")

		if err != nil {
			log.Println("collection list", err)
		}
		c.JSON(200, gin.H{"Collection": collectionlist})

	}
}

// block list
func BlockList(c *gin.Context) {

	var (
		filter blocks.Filter
		limt   int
		offset int
	)

	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		blockBannerValue, _ := c.Cookie("blockbanner")

		if blockBannerValue == "" {

			blockBannerValue = "true"
		}
		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		limit := c.Query("limit")
		pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
		filter.Channelid = (c.Query("id"))
		channelid, _ := strconv.Atoi(filter.Channelid)

		var filterflag string
		// if filter.Keyword != ""  {
		// 	filterflag = true
		// } else {
		// 	filterflag = false
		// }
		if filter.Keyword != "" {
			filterflag = "true"
		} else if filter.Channelid != "" {

			filterflag = "ftrue"
		} else {
			filterflag = "false"
		}

		if limit == "" {
			limt = Limit
		} else {
			limt, _ = strconv.Atoi(limit)
		}
		if pageno != 0 {
			offset = (pageno - 1) * limt
		}

		var finalblocklist []blocks.TblBlock

		blocklist, blockcount, err := BlockConfig.BlockList(limt, offset, blocks.Filter(filter), TenantId)

		fmt.Println("blocklist:",blocklist)

		for _, val := range blocklist {
			var first = val.FirstName
			var last = val.LastName
			var firstn = strings.ToUpper(first[:1])
			var lastn string
			if val.LastName != "" {
				lastn = strings.ToUpper(last[:1])
			}

			val.ChannelNames = strings.Split(val.ChannelSlugname, ",")

			var Name = firstn + lastn
			val.NameString = Name

			tagname := strings.Split(val.TagValue, ",")

			val.TagValueArr = append(val.TagValueArr, tagname...)
			img := val.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				val.CoverImage = "/image-resize?name=" + val.CoverImage
			}
			if val.ProfileImagePath != "" {
				userimg := val.ProfileImagePath
				imgflag := strings.Contains(userimg, imgcontain)
				if !imgflag {
					val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
				}
			}

			finalblocklist = append(finalblocklist, val)
		}

		channelslist, _, err := ChannelConfig.ListChannel(chn.Channels{IsActive: false, CreateOnly: true, TenantId: TenantId})

		var finalChannelList []channels.Tblchannel

		for _, chnList := range channelslist {

			chnid := strconv.Itoa(chnList.Id)

			_, blockcount, err := BlockConfig.CollectionList(blocks.Filter{}, TenantId, chnid)

			if err != nil {
				log.Println("collection list", err)
			}
			chnList.EntriesCount = int(blockcount)
			finalChannelList = append(finalChannelList, chnList)
		}

		taglist, _ := BlockConfig.TagList(blocks.Filter(filter), TenantId)

		if err != nil {
			log.Println("collection list", err)
		}

		var paginationendcount = len(blocklist) + offset
		paginationstartcount := offset + 1
		Previous, Next, PageCount, Page := Pagination(pageno, int(blockcount), limt)
		menu := NewMenuController(c)
		ModuleName, TabName, _ := ModuleRouteName(c)
		translate, _ := TranslateHandler(c)
		var active string
		storagetype, err := GetSelectedType()

		if err != nil {
			fmt.Printf("blocklist getting storagetype error: %s", err)
		}

		c.HTML(200, "blocks.html", gin.H{"Menu": menu, "Pagination": PaginationData{
			NextPage:     pageno + 1,
			PreviousPage: pageno - 1,
			TotalPages:   PageCount,
			TwoAfter:     pageno + 2,
			TwoBelow:     pageno - 2,
			ThreeAfter:   pageno + 3,
		}, "blockBannerValue": blockBannerValue, "chnid": channelid, "channelist": finalChannelList, "channel": channelslist, "Count": blockcount, "SearchTrues": filterflag, "linktitle": ModuleName, "Limit": limt, "UserId": c.GetInt("userid"), "Previous": Previous, "Page": Page, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "StorageType": storagetype.SelectedType, "Filter": filter, "title": ModuleName, "csrf": csrf.GetToken(c), "Blocklist": finalblocklist, "Active": active, "Taglist": taglist, "TabName": TabName, "translate": translate})

	}

}

// Default Blocks list
func DefaultBlockList(c *gin.Context) {
	var (
		filter models.Filter
		limt   int
		offset int
	)

	limit := c.Query("limit")
	pageno, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	filter.Keyword = strings.Trim(c.DefaultQuery("keyword", ""), " ")
	filter.ChannelName = strings.Trim(c.DefaultQuery("channel", ""), "")

	var filterflag bool
	var keywordtrue bool
	if filter.Keyword != "" || filter.ChannelName != "" {
		filterflag = true

	} else {
		filterflag = false
	}

	if filter.Keyword != "" {
		keywordtrue = true
	}
	if limit == "" {
		limt = Limit
	} else {
		limt, _ = strconv.Atoi(limit)
	}
	if pageno != 0 {
		offset = (pageno - 1) * limt
	}

	endurl := os.Getenv("MASTER_BLOCKS_ENDPOINTURL")
	req, err := http.NewRequest("GET", endurl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}

	query := req.URL.Query()
	if filterflag {
		query.Add("keyword", filter.Keyword)
		query.Add("channel", filter.ChannelName)
	}
	query.Add("limit", strconv.Itoa(limt))
	query.Add("offset", strconv.Itoa(offset))
	query.Add("page", strconv.Itoa(pageno))
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			InsecureSkipVerify: true, // This disables certificate verification
	// 		},
	// 	},
	// }

	// Use this client to make requests
	// resp, err := client.Do(req)
	masterconnect := true

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error connecting to master server:", err)
		masterconnect = false
	} else {
		defer resp.Body.Close()
	}

	var responseData ResponseData
	if masterconnect {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			fmt.Println("Error response:", err)
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			err = json.NewDecoder(resp.Body).Decode(&responseData)
			if err != nil {
				masterconnect = false
			}
		} else {
			masterconnect = false
		}
	}

	if !masterconnect {
		responseData = ResponseData{
			DefaultList:    []models.TblMstrBlocks{},
			AllList:        []models.TblMstrBlocks{},
			FinalblockList: []models.TblMstrBlocks{},
			BlockCount:     0,
		}
	}

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}
	if !permisison {

		ErrorLog.Printf("Block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	blockBannerValue, _ := c.Cookie("blockbanner")
	if blockBannerValue == "" {
		blockBannerValue = "true"
	}

	channelist, _, err := ChannelConfig.ListChannel(channels.Channels{Limit: 200, Offset: 0, IsActive: true, TenantId: TenantId})
	if err != nil {
		ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)
	}

	var finalChannelList []channels.Tblchannel
	for _, chnList := range channelist {
		blockCount := 0
		for _, blockInfo := range responseData.AllList {
			if blockInfo.ChannelSlugname == chnList.ChannelName {
				blockCount++
			}
		}
		chnList.EntriesCount = blockCount
		finalChannelList = append(finalChannelList, chnList)
	}

	paginationendcount := len(responseData.DefaultList) + offset
	paginationstartcount := offset + 1
	Previous, Next, PageCount, Page := Pagination(pageno, int(responseData.BlockCount), limt)
	menu := NewMenuController(c)
	ModuleName, TabName, _ := ModuleRouteName(c)
	translate, _ := TranslateHandler(c)
	storagetype, _ := GetSelectedType()

	c.HTML(200, "defaultblock.html", gin.H{"Menu": menu, "Pagination": PaginationData{
		NextPage:     pageno + 1,
		PreviousPage: pageno - 1,
		TotalPages:   PageCount,
		TwoAfter:     pageno + 2,
		TwoBelow:     pageno - 2,
		ThreeAfter:   pageno + 3,
	}, "blockBannerValue": blockBannerValue, "chnid": filter.ChannelName, "keywordtrue": keywordtrue, "Count": responseData.BlockCount, "Searchtrue": filterflag, "linktitle": "Blocks", "Limit": limt, "UserId": c.GetInt("userid"), "Defaultlist": responseData.FinalblockList, "Previous": Previous, "Page": Page, "Next": Next, "PageCount": PageCount, "CurrentPage": pageno, "Paginationendcount": paginationendcount, "Paginationstartcount": paginationstartcount, "StorageType": storagetype.SelectedType, "Filter": filter, "title": ModuleName, "csrf": csrf.GetToken(c), "TabName": TabName, "translate": translate, "channelist": finalChannelList})

}

// Create block
func CreateBlock(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		var premium int

		var (
			imageByte []byte
			err       error
		)

		premium, _ = strconv.Atoi(c.PostForm("premiumvalue"))
		chennelname := c.PostForm("chennelname")
		channelids := c.PostForm("channelids")

		// chID := ChannelConfig.GetChannal(chennelname, TenantId)

		imagedata := c.PostForm("coverimg")

		var imageName, imagePath string

		tenantDetails, tenanterr := NewTeam.GetTenantDetails(TenantId)

		if tenanterr != nil {
			fmt.Println(tenanterr)
		}

		if imagedata != "" {
			imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "blocks")
			if err != nil {
				fmt.Println(err)
			}

			imagePath = tenantDetails.S3FolderName + imagePath

			uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
			if uerr != nil {
				c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
				c.Redirect(301, "/blocks")
			}
		}
		// }

		BlockCreate := blocks.BlockCreation{
			Title:        c.PostForm("title"),
			BlockContent: c.PostForm("html"),
			CoverImage:   imagePath,
			TenantId:     TenantId,
			Prime:        premium,
			CreatedBy:    c.GetInt("userid"),
			IsActive:     1,
			ChannelName:  chennelname,
			ChannelId:    channelids,
		}

		createblock, err := BlockConfig.CreateBlock(BlockCreate)

		if err != nil {
			log.Println("collection list", err)
		}
		collection := blocks.CreateCollection{
			UserId:   createblock.CreatedBy,
			BlockId:  createblock.Id,
			TenantId: TenantId,
		}

		err1 := BlockConfig.BlockCollection(collection)

		if err1 != nil {
			log.Println("collection list", err1)
		}

		// var tagnames []string

		tagname := c.PostForm("tag")

		tagnames := strings.Split(tagname, ",")

		for _, val := range tagnames {

			tag, err := BlockConfig.CheckTagName(val, TenantId)

			if err != nil {
				fmt.Println("tag name alreay exists:", tag)
			}

			if tag.TagTitle == "" {

				MstrTag := blocks.MasterTagCreate{

					TagTitle:  val,
					TenantId:  TenantId,
					CreatedBy: c.GetInt("userid"),
				}

				createtags, err := BlockConfig.CreateMasterTag(MstrTag)

				if err != nil {
					fmt.Println("block err", err)
				}

				TagCreate := blocks.CreateTag{
					BlockId:   createblock.Id,
					TagId:     createtags.Id,
					TagName:   createtags.TagTitle,
					TenantId:  TenantId,
					CreatedBy: c.GetInt("userid"),
				}

				err2 := BlockConfig.CreateTag(TagCreate)

				if err2 != nil {
					fmt.Println("block err", err)
				}

			} else {

				TagCreate := blocks.CreateTag{
					BlockId:   createblock.Id,
					TagId:     tag.Id,
					TagName:   tag.TagTitle,
					TenantId:  TenantId,
					CreatedBy: c.GetInt("userid"),
				}

				err2 := BlockConfig.CreateTag(TagCreate)

				if err2 != nil {
					fmt.Println("block err", err)
				}

			}

			if err != nil {
				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
				c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
				c.JSON(500, gin.H{"Error": err})
			}

		}

		var filter blocks.Filter
		var finalblocklist []blocks.TblBlock

		blocklist, _, err := BlockConfig.BlockList(0, 0, blocks.Filter(filter), TenantId)

		if err != nil {
			fmt.Println("collection list", err)
		}

		for _, val := range blocklist {
			var first = val.FirstName
			var last = val.LastName
			var firstn = strings.ToUpper(first[:1])
			var lastn string
			if val.LastName != "" {
				lastn = strings.ToUpper(last[:1])
			}

			val.ChannelNames = strings.Split(val.ChannelSlugname, ",")

			var Name = firstn + lastn
			val.NameString = Name

			tagname := strings.Split(val.TagValue, ",")

			val.TagValueArr = append(val.TagValueArr, tagname...)
			img := val.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				val.CoverImage = "/image-resize?name=" + val.CoverImage
			}
			if val.ProfileImagePath != "" {
				userimg := val.ProfileImagePath
				imgflag := strings.Contains(userimg, imgcontain)
				if !imgflag {
					val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
				}
			}

			finalblocklist = append(finalblocklist, val)
		}

		if err != nil {
			fmt.Println("collection list", err)
		}

		data := map[string]interface{}{"data": finalblocklist}

		bytedata1, _ := json.Marshal(data)

		// c.SetCookie("get-toast", "Block Created Successfully", 3600, "", "", false, false)
		// c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.JSON(500, gin.H{"blocks": string(bytedata1)})

	}
}

// Edit Functionality oin block
func EditBlock(c *gin.Context) {

	//get values from html from data
	id, _ := strconv.Atoi(c.Query("id"))

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Block authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		blocklist, err := BlockConfig.BlockEdit(id, TenantId) //update block details

		if err != nil {
			ErrorLog.Printf("block details error: %s", err)

		} else {

			c.JSON(200, blocklist)
		}

	}

}

// Update block
func UpdateBlock(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		var imageName, imagePath string
		imagedata := c.PostForm("coverimg")
		chennelname := c.PostForm("chennelname")
		channelids := c.PostForm("channelids")

		maxlength := 60
		if len(imagedata) > maxlength {
			_, imgerr := GetSelectedType()
			if imgerr != nil {
				ErrorLog.Printf("error get storage type error: %s", imgerr)
			}

			tenantDetails, tenanterr := NewTeam.GetTenantDetails(TenantId)

			if tenanterr != nil {
				fmt.Println(tenanterr)
			}

			// if storagetype.SelectedType == "local" {
			var (
				imageByte []byte
				err       error
			)
			if imagedata != "" {
				fmt.Println("print1")
				imageName, imagePath, imageByte, err = ConvertBase64toByte(imagedata, "blocks")
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("print2")

				imagePath = tenantDetails.S3FolderName + imagePath

				uerr := storagecontroller.UploadCropImageS3(imageName, imagePath, imageByte)
				fmt.Println("print3")

				if uerr != nil {
					c.SetCookie("Alert-msg", "ERRORAWScredentialsnotfound", 3600, "", "", false, false)
					c.Redirect(301, "/blocks")
				}
			} else {
				imagePath = imagedata
			}
		}
		// }

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		premium, _ := strconv.Atoi(c.PostForm("premiumvalue"))
		isactive, _ := strconv.Atoi(c.PostForm("isactive"))
		id, _ := strconv.Atoi(c.PostForm("blockid"))

		BlockCreate := blocks.BlockCreation{
			Title:        c.PostForm("title"),
			BlockContent: c.PostForm("html"),
			CoverImage:   imagePath,
			TenantId:     TenantId,
			Prime:        premium,
			ModifiedBy:   c.GetInt("userid"),
			IsActive:     isactive,
			ChannelName:  chennelname,
			ChannelId:    channelids,
		}

		err := BlockConfig.UpdateBlock(id, BlockCreate)

		if err != nil {
			log.Println("block update", err)
		}

		deletetag := c.PostForm("deletetag")

		if deletetag != "" {

			deletetags := strings.Split(deletetag, ",")

			for _, deltag := range deletetags {

				tagname := strings.TrimPrefix(deltag, " ")

				tagdelerr := BlockConfig.DeleteTags(id, tagname, TenantId)

				if tagdelerr != nil {
					fmt.Println(tagdelerr)
				}
			}
		}
		var tagname = c.PostForm("tag")

		if tagname != "" {

			tagnames := strings.Split(tagname, ",")

			for _, val := range tagnames {

				trimmedTag := strings.TrimSpace(val)

				tag, err := BlockConfig.CheckTagName(trimmedTag, TenantId)

				if err != nil {
					fmt.Println("tag name alreay exists:", tag)
				}

				if tag.TagTitle == "" {

					fmt.Println(trimmedTag, "tagtrim")

					MstrTag := blocks.MasterTagCreate{

						TagTitle:  trimmedTag,
						TenantId:  TenantId,
						CreatedBy: c.GetInt("userid"),
					}

					createtags, err := BlockConfig.CreateMasterTag(MstrTag)

					if err != nil {
						fmt.Println("block err", err)
					}

					TagCreate := blocks.CreateTag{
						BlockId:   id,
						TagId:     createtags.Id,
						TagName:   createtags.TagTitle,
						TenantId:  TenantId,
						CreatedBy: c.GetInt("userid"),
					}

					err2 := BlockConfig.CreateTag(TagCreate)

					if err2 != nil {
						fmt.Println("block err", err)
					}
				} else {

					TagCreate := blocks.CreateTag{
						BlockId:   id,
						TagId:     tag.Id,
						TagName:   trimmedTag,
						TenantId:  TenantId,
						CreatedBy: c.GetInt("userid"),
					}

					err2 := BlockConfig.CreateTag(TagCreate)

					if err2 != nil {
						fmt.Println("block err", err)
					}

				}

			}
		}

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		var (
			filter blocks.Filter
		)
		var finalblocklist []blocks.TblBlock

		blocklist, _, err := BlockConfig.BlockList(0, 0, blocks.Filter(filter), TenantId)

		if err != nil {
			fmt.Println("collection list", err)
		}

		for _, val := range blocklist {
			var first = val.FirstName
			var last = val.LastName
			var firstn = strings.ToUpper(first[:1])
			var lastn string
			if val.LastName != "" {
				lastn = strings.ToUpper(last[:1])
			}

			val.ChannelNames = strings.Split(val.ChannelSlugname, ",")

			var Name = firstn + lastn
			val.NameString = Name

			tagname := strings.Split(val.TagValue, ",")

			val.TagValueArr = append(val.TagValueArr, tagname...)
			img := val.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				val.CoverImage = "/image-resize?name=" + val.CoverImage
			}
			if val.ProfileImagePath != "" {
				userimg := val.ProfileImagePath
				imgflag := strings.Contains(userimg, imgcontain)
				if !imgflag {
					val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
				}
			}

			finalblocklist = append(finalblocklist, val)
		}


		data := map[string]interface{}{"data": finalblocklist}

		bytedata, _ := json.Marshal(data)

		c.JSON(200, gin.H{"blocks": string(bytedata)})
		// c.SetCookie("get-toast", "Block Update Successfully", 3600, "", "", false, false)
		// c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		// c.Redirect(301, "/blocks")

	}
}

// Delete Block

func DeleteBlock(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		id, _ := strconv.Atoi(c.Query("id"))

		fmt.Println("remove blocks",id)

		BlockConfig.UserId = c.GetInt("userid")

		err := BlockConfig.RemoveBlock(id, TenantId)

		if err != nil {
			// c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
			// c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			return
		}
		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")
		var filter blocks.Filter

		var finalblocklist []blocks.TblBlock

		blocklist, _, err := BlockConfig.BlockList(0, 0, blocks.Filter(filter), TenantId)

		if err != nil {
			fmt.Println("collection list", err)
		}

		for _, val := range blocklist {
			var first = val.FirstName
			var last = val.LastName
			var firstn = strings.ToUpper(first[:1])
			var lastn string
			if val.LastName != "" {
				lastn = strings.ToUpper(last[:1])
			}

			val.ChannelNames = strings.Split(val.ChannelSlugname, ",")

			var Name = firstn + lastn
			val.NameString = Name

			tagname := strings.Split(val.TagValue, ",")

			val.TagValueArr = append(val.TagValueArr, tagname...)
			img := val.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				val.CoverImage = "/image-resize?name=" + val.CoverImage
			}
			if val.ProfileImagePath != "" {
				userimg := val.ProfileImagePath
				imgflag := strings.Contains(userimg, imgcontain)
				if !imgflag {
					val.ProfileImagePath = "/image-resize?name=" + val.ProfileImagePath
				}
			}

			finalblocklist = append(finalblocklist, val)
		}

		data := map[string]interface{}{"data": finalblocklist}

		bytedata, _ := json.Marshal(data)
		c.JSON(200, gin.H{"blocks": string(bytedata)})
        
	}
}

// Block Is active
func BlockIsActive(c *gin.Context) {

	//get values from html from data
	id, _ := strconv.Atoi(c.Request.PostFormValue("id"))
	val, _ := strconv.Atoi(c.Request.PostFormValue("isactive"))
	userid := c.GetInt("userid")

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("Block authorization error: %s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		flg, err := BlockConfig.BlockIsActive(id, val, userid, TenantId) //update block status

		if err != nil {
			ErrorLog.Printf("block status update error: %s", err)
			json.NewEncoder(c.Writer).Encode(flg)

		} else {

			json.NewEncoder(c.Writer).Encode(flg)
		}

	}

}

// Create Collection
func CreateCollection(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block collection list authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		id := c.GetInt("userid")

		blockid, _ := strconv.Atoi(c.PostForm("blockid"))

		collectionflg, _ := BlockConfig.CheckCollection(blockid, id, TenantId)

		if !collectionflg {

			collection := blocks.CreateCollection{
				UserId:   id,
				BlockId:  blockid,
				TenantId: TenantId,
			}

			err := BlockConfig.BlockCollection(collection)

			if err != nil {
				c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
				c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
				c.JSON(200, false)

			}

		} else {
			c.SetCookie("Alert-msg", "Block collection already exists", 3600, "", "", false, false)
			c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
			c.JSON(200, false)
			return
		}

		c.JSON(200, true)

	} else {
		c.Redirect(301, "/403-page")
	}

}

// Remove collection

func CollectionRemove(c *gin.Context) {

	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
	}

	if perr != nil {
		ErrorLog.Printf("block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		id, _ := strconv.Atoi(c.Query("blockid"))

		BlockConfig.UserId = c.GetInt("userid")

		// err := BlockConfig.DeleteCollection(id, TenantId)

		// if err != nil {
		// 	c.SetCookie("Alert-msg", ErrInternalServerError, 3600, "", "", false, false)
		// 	c.SetCookie("Alert-msg", "alert", 3600, "", "", false, false)
		// 	c.JSON(200, false)
		// 	return
		// }

		err1 := BlockConfig.RemoveBlock(id, TenantId)

		if err1 != nil {
			log.Println(err1)
		}

		c.SetCookie("get-toast", "Collection Remove Successfully", 3600, "", "", false, false)
		c.SetCookie("Alert-msg", "success", 3600, "", "", false, false)
		c.JSON(200, true)

	} else {
		c.Redirect(301, "/403-page")
	}
}

func Addtomycollection(c *gin.Context) {

	title := c.PostForm("blocktitle")
	content := c.PostForm("blockdata")
	coverimg := c.PostForm("blockimg")
	channelname := c.PostForm("blockchannel")
	permisison, perr := NewAuth.IsGranted("Blocks", auth.CRUD, TenantId)

	if perr != nil {
		ErrorLog.Printf("block authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {

		channeldetails, err := ChannelConfig.GetchannelByName(channelname, TenantId)
		if err != nil {
			ErrorLog.Printf("%v: %v", ErrGettingTemplateModuleList, err)

		}

		if channeldetails.ChannelName != channelname {

			c.JSON(200, gin.H{"data": "channelmissing"})
			return
		}

		userid := c.GetInt("userid")
		channelid := strconv.Itoa(channeldetails.Id)

		BlockCreate := blocks.BlockCreation{
			Title:        title,
			BlockContent: content,
			CoverImage:   coverimg,
			TenantId:     TenantId,
			CreatedBy:    userid,
			IsActive:     1,
			ChannelName:  channelname,
			ChannelId:    channelid,
			Prime:        0,
		}
		_, err1 := BlockConfig.CreateBlock(BlockCreate)

		if err1 != nil {

			c.JSON(200, gin.H{"data": false})
			return

		}

		c.JSON(200, gin.H{"data": true})

	}
}

// check block title name alreay exists
func CheckTitleInBlock(c *gin.Context) {

	title := c.PostForm("block_title")

	id, _ := strconv.Atoi(c.PostForm("block_id"))

	permisison, perr := NewAuth.IsGranted("Blocks", auth.Read, TenantId)

	if perr != nil {
		ErrorLog.Printf("block authorization error :%s", perr)
	}

	if !permisison {
		ErrorLog.Printf("Block authorization error: %s", perr)
		c.Redirect(301, "/403-page")
		return
	}

	if permisison {
		flg, err := BlockConfig.CheckTitleInBlock(title, id, TenantId)
		if err != nil {
			ErrorLog.Printf("check title in block error :%s", perr)
			json.NewEncoder(c.Writer).Encode(false)
			return
		}
		json.NewEncoder(c.Writer).Encode(flg)

	}

}
