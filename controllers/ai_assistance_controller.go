package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"spurt-cms/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/spurtcms/auth"
	"github.com/spurtcms/blocks"
	chn "github.com/spurtcms/channels"
	csrf "github.com/utrack/gin-csrf"
)

func Aicontents(c *gin.Context) {

	channelid := c.Query("id")

	ChannelId, _ := strconv.Atoi(channelid)

	ChannelList, _, _ := ChannelConfig.GetChannelsById(ChannelId, TenantId)

	keyword := strings.TrimSpace(c.Request.URL.Query().Get("keyword"))

	menu := NewMenuController(c)

	ModuleName, TabName, _ := ModuleRouteName(c)

	translate, _ := TranslateHandler(c)

	myapps, _, _ := models.GetAppList(keyword, TenantId)

	var filterflag bool

	if keyword != "" {
		filterflag = true
	} else {
		filterflag = false
	}

	var applist []models.TblApps

	for _, app := range myapps {

		fields, _ := readJSONFile(app.FieldJsonPath)

		var data map[string]interface{}

		json.Unmarshal(fields, &data)

		app.CustomFields = data

		applist = append(applist, app)

	}

	userid := c.GetInt("userid")

	UserDetails, _ := models.GetArticleCount(userid, TenantId)

	c.HTML(200, "ai-assistance.html", gin.H{"articlecount": UserDetails.ArticleCount, "Menu": menu, "csrf": csrf.GetToken(c), "ChanSlug": ChannelList.SlugName, "Searchtrue": filterflag, "translate": translate, "title": ModuleName, "linktitle": "AI assistance", "Tabmenu": TabName, "myapps": applist, "filter": keyword})

}

func Receivetopicdata(c *gin.Context) {

	topicvalue := c.PostForm("topic")

	var aiprompt models.TblAiPrompt

	perr := models.GetAiPrompts(&aiprompt, 2)

	systemPrompt := aiprompt.SystemPrompt

	tmpl := aiprompt.UserPrompt

	if perr != nil {
		ErrorLog.Printf("get aiprompt error: %s", perr)
	}
	newString := strings.Replace(tmpl, "{topic}", topicvalue, 1)

	list, _, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	var ApiModule models.TblAiSettingsModule

	for _, value := range list {

		if value.IsActive == 1 {
			ApiModule = value
		}
	}

	var client *openai.Client
	var Model string

	if ApiModule.ApiKey != "" {

		client = openai.NewClient(ApiModule.ApiKey)
		Model = ApiModule.AiModel
		fmt.Println("ApiModule1:")

	} else {

		client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		Model = "gpt-4o-mini"

	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: newString,
		},
	}

	// Create a context
	ctx := context.Background()

	// Call the Chat API
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:          Model, // Specify the model to use
		Messages:       messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{Type: "json_object"},
	})
	if err != nil {
		fmt.Println("Error calling ChatCompletion: %v", err)
	}
	var data string
	// Print the response
	for _, choice := range response.Choices {
		data = choice.Message.Content
		fmt.Println("Response:", choice.Message.Content)

	}
	c.JSON(200, gin.H{"data": data})
}

func Generatetopics(c *gin.Context) {

	keywordvalue := c.PostForm("keyword")

	var aiprompt models.TblAiPrompt

	perr := models.GetAiPrompts(&aiprompt, 3)

	systemPrompt := aiprompt.SystemPrompt

	tmpl := aiprompt.UserPrompt

	if perr != nil {
		ErrorLog.Printf("get aiprompt error: %s", perr)
	}

	newString := strings.Replace(tmpl, "{keyword}", keywordvalue, 1)

	list, _, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	var ApiModule models.TblAiSettingsModule

	for _, value := range list {

		if value.IsActive == 1 {
			ApiModule = value
		}
	}

	var client *openai.Client
	var Model string

	if ApiModule.ApiKey != "" {

		client = openai.NewClient(ApiModule.ApiKey)
		Model = ApiModule.AiModel
		fmt.Println("ApiModule2:")

	} else {

		client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		Model = "gpt-4o-mini"

	}


	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: newString,
		},
	}

	// Create a context
	ctx := context.Background()

	// Call the Chat API
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:          Model, // Specify the model to use
		Messages:       messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{Type: "json_object"},
	})
	if err != nil {
		fmt.Println("Error calling ChatCompletion: %v", err)
	}
	// Print the response
	var data string
	for _, choice := range response.Choices {
		data = choice.Message.Content
	}
	c.JSON(200, gin.H{"data": data})
}
func Generatearticle(c *gin.Context) {

	chnid := c.PostForm("chnid")

	topic := c.PostForm("titletopic")

	keywordvalue := c.PostForm("keywords")

	articallenth := c.PostForm("articallenth")

	tonetype := c.PostForm("tonetype")

	var aiprompt models.TblAiPrompt

	perrs := models.GetAiPrompts(&aiprompt, 1)

	systemPrompt := aiprompt.SystemPrompt

	tmpl := aiprompt.UserPrompt

	if perrs != nil {
		ErrorLog.Printf("get aiprompt error: %s", perrs)
	}

	tmpl = strings.Replace(tmpl, "{topic}", topic, 1)
	tmpl = strings.Replace(tmpl, "{keywords}", keywordvalue, 1)
	tmpl = strings.Replace(tmpl, "{articleLength}", articallenth, 1)
	tmpl = strings.Replace(tmpl, "{tone}", tonetype, 2)
	fmt.Println("end prompt", tmpl)

	list, _, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	var ApiModule models.TblAiSettingsModule

	for _, value := range list {

		if value.IsActive == 1 {
			ApiModule = value
		}
	}

	var client *openai.Client
	var Model string

	if ApiModule.ApiKey != "" {

		client = openai.NewClient(ApiModule.ApiKey)
		Model = ApiModule.AiModel
		fmt.Println("ApiModule3:")

	} else {

		client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		Model = "gpt-4o-mini"
		
		userid := c.GetInt("userid")

		models.ArticleCountUpdate(userid, TenantId)

	}


	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: tmpl,
		},
	}

	// Create a context
	ctx := context.Background()

	// Call the Chat API
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    Model, // Specify the model to use
		Messages: messages,
	})
	if err != nil {
		fmt.Println("Error calling ChatCompletion: %v", err)
	}
	// Print the response
	var htmldata string
	for _, choice := range response.Choices {
		htmldata = choice.Message.Content
		htmldata = strings.Replace(htmldata, "```html", "", 1)
		htmldata = strings.Replace(htmldata, "```", "", 1)
		fmt.Println("Response:", htmldata)

	}

	id, _ := strconv.Atoi(c.Param("id"))
	// flag := true

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("create Entry authorization error: %s", perr)
	}

	// if permisison {

	channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	if clerr != nil {
		ErrorLog.Printf("create Entry listchannel error: %s", clerr)
	}

	_, entrcount, _, cherr := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0}, TenantId)
	if cherr != nil {
		ErrorLog.Printf("createEntry entrylist error: %s", cherr)
	}

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)
	if cerr != nil {
		ErrorLog.Printf("createEntry channel data error: %s", cerr)
	}

	var idstr []string

	for _, val := range FinalSelectedCategories {
		var str string
		for index, cat := range val.Categories {
			if index+1 != len(val.Categories) {
				str += strconv.Itoa(cat.Id) + ","
			} else {
				str += strconv.Itoa(cat.Id)
			}
		}
		idstr = append(idstr, str)
	}

	field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
	if err != nil {
		ErrorLog.Printf("createEntry get master fields error: %s", perr)
	}

	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02T15:04")
	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	// Folder, File, Media, NextCont, merr := GetMedia()
	// if merr != nil {
	// 	ErrorLog.Printf("createEntry media error: %s", perr)
	// }
	var (
		filter         blocks.Filter
		collectionlist []blocks.TblBlock
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
	var finalcollectionlist []blocks.TblBlock

	if permisison {

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		collectionlist, _, err = BlockConfig.CollectionList(blocks.Filter(filter), TenantId, "")

		for _, value := range collectionlist {

			img := value.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				value.CoverImage = "/image-resize?name=" + value.CoverImage
			}

			finalcollectionlist = append(finalcollectionlist, value)
		}

		if err != nil {
			fmt.Println("collection list", err)
		}

	}
	selectedtype, _ := GetSelectedType()

	data := map[string]interface{}{"data": finalcollectionlist}

	bytedata, _ := json.Marshal(data)

	baseurl := os.Getenv("BASE_URL")

	urlpath := map[string]interface{}{"path": baseurl + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	ubyte, _ := json.Marshal(urlpath)

	ModuleName, _, _ := ModuleRouteName(c)

	selectedchannel := channelist[len(channelist)-1]

	slchannelid := selectedchannel.Id

	// fmt.Println("chnid:", chnid)

	ChannelId, _ := strconv.Atoi(chnid)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(200, "addentry.html", gin.H{"Menu": menu, "EntryId": ChannelId, "linktitle": "Create Entry", "title": ModuleName, "Fields": field, "translate": translate, "csrf": csrf.GetToken(c), "channellist": channelist, "CategoryName": FinalSelectedCategories, "entrycount": entrcount, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "currentdate": cdate, "StorageType": selectedtype.SelectedType, "blocks": string(bytedata), "Mode": "create", "Slchannelid": slchannelid, "Storagepath": string(ubyte), "htmldata": htmldata})

	// c.Redirect(301,"/entries/create")
}
func Receiveproductdata(c *gin.Context) {

	topicvalue := c.PostForm("product")

	tmpl := `Looking for a reliable, high-quality product that can meet all your needs? {product} might just be the perfect choice. Designed to provide exceptional value and performance, {product} is gaining popularity for its innovative features and affordable price point. reply only with json: { topic: "topic",keywords:["keywords"],articleLength:"count of words"}`

	newString := strings.Replace(tmpl, "{product}", topicvalue, 1)

	list, _, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	var ApiModule models.TblAiSettingsModule

	for _, value := range list {

		if value.IsActive == 1 {
			ApiModule = value
		}
	}

	var client *openai.Client
	var Model string

	if ApiModule.ApiKey != "" {

		client = openai.NewClient(ApiModule.ApiKey)
		Model = ApiModule.AiModel
		fmt.Println("ApiModule4:")

	} else {

		client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		Model = "gpt-4o-mini"

	}


	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are the best content creator for a common blog website",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: newString,
		},
	}

	// Create a context
	ctx := context.Background()

	// Call the Chat API
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:          Model, // Specify the model to use
		Messages:       messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{Type: "json_object"},
	})
	if err != nil {
		fmt.Println("Error calling ChatCompletion: %v", err)
	}
	var data string
	// Print the response
	for _, choice := range response.Choices {
		data = choice.Message.Content
		// fmt.Println("Response:", choice.Message.Content)

	}
	c.JSON(200, gin.H{"data": data})
}
func GenerateEcommercearticle(c *gin.Context) {

	topic := c.PostForm("titleproduct")

	keywordvalue := c.PostForm("keyfeature")

	articallenth := c.PostForm("descriptionlength")

	tonetype := c.PostForm("ectonetype")

	tmpl := `create a set of content for the ecommerce blog  with the given product name: {topic} in {articleLength}. The article should be in the tone of {tone}. The article should be based on the keywords: "{keywords}", and you may consider those keywords as sub-concepts. Generate the article effectively and exactly in the {tone} tone. The output should be in the format of innerHTML with Tailwind CSS and the response doesn't contain any additional texts or notes`

	tmpl = strings.Replace(tmpl, "{topic}", topic, 1)
	tmpl = strings.Replace(tmpl, "{keywords}", keywordvalue, 1)
	tmpl = strings.Replace(tmpl, "{articleLength}", articallenth, 1)
	tmpl = strings.Replace(tmpl, "{tone}", tonetype, 2)
	// fmt.Println("end prompt", tmpl)

	list, _, _ := models.ListAiModule(0, 0, models.Filter{}, TenantId)

	var ApiModule models.TblAiSettingsModule

	for _, value := range list {

		if value.IsActive == 1 {
			ApiModule = value
		}
	}

	var client *openai.Client
	var Model string

	if ApiModule.ApiKey != "" {

		client = openai.NewClient(ApiModule.ApiKey)
		Model = ApiModule.AiModel
		fmt.Println("ApiModule5:")

	} else {

		client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		Model = "gpt-4o-mini"

	}


	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are the best content creator for a common blog website",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: tmpl,
		},
	}

	// Create a context
	ctx := context.Background()

	// Call the Chat API
	response, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    Model, // Specify the model to use
		Messages: messages,
	})
	if err != nil {
		fmt.Println("Error calling ChatCompletion: %v", err)
	}
	// Print the response
	var htmldata string
	for _, choice := range response.Choices {
		htmldata = choice.Message.Content
		htmldata = strings.Replace(htmldata, "```html", "", 1)
		htmldata = strings.Replace(htmldata, "```", "", 1)
		fmt.Println("Response:", htmldata)

	}

	id, _ := strconv.Atoi(c.Param("id"))
	// flag := true

	_, perr := NewAuth.IsGranted("Entries", auth.CRUD, TenantId)
	if perr != nil {
		ErrorLog.Printf("create Entry authorization error: %s", perr)
	}

	// if permisison {

	channelist, clerr := ChannelConfig.GetPermissionChannel(TenantId)
	if clerr != nil {
		ErrorLog.Printf("create Entry listchannel error: %s", clerr)
	}

	_, entrcount, _, cherr := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: id, Limit: 0, Offset: 0}, TenantId)
	if cherr != nil {
		ErrorLog.Printf("createEntry entrylist error: %s", cherr)
	}

	_, FinalSelectedCategories, cerr := ChannelConfig.GetChannelsById(id, TenantId)
	if cerr != nil {
		ErrorLog.Printf("createEntry channel data error: %s", cerr)
	}

	var idstr []string

	for _, val := range FinalSelectedCategories {
		var str string
		for index, cat := range val.Categories {
			if index+1 != len(val.Categories) {
				str += strconv.Itoa(cat.Id) + ","
			} else {
				str += strconv.Itoa(cat.Id)
			}
		}
		idstr = append(idstr, str)
	}

	field, err := ChannelConfig.GetAllMasterFieldType(TenantId)
	if err != nil {
		ErrorLog.Printf("createEntry get master fields error: %s", perr)
	}

	currentTime := time.Now()
	cdate := currentTime.Format("2006-01-02T15:04")
	translate, _ := TranslateHandler(c)
	menu := NewMenuController(c)

	// Folder, File, Media, NextCont, merr := GetMedia()
	// if merr != nil {
	// 	ErrorLog.Printf("createEntry media error: %s", perr)
	// }
	var (
		filter         blocks.Filter
		collectionlist []blocks.TblBlock
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
	var finalcollectionlist []blocks.TblBlock

	if permisison {

		BlockConfig.DataAccess = c.GetInt("dataaccess")
		BlockConfig.UserId = c.GetInt("userid")

		collectionlist, _, err = BlockConfig.CollectionList(blocks.Filter(filter), TenantId, "")

		for _, value := range collectionlist {

			img := value.CoverImage
			imgcontain := "/image-resize?name="
			flag := strings.Contains(img, imgcontain)
			if !flag {
				value.CoverImage = "/image-resize?name=" + value.CoverImage
			}

			finalcollectionlist = append(finalcollectionlist, value)
		}

		if err != nil {
			fmt.Println("collection list", err)
		}

	}
	selectedtype, _ := GetSelectedType()

	data := map[string]interface{}{"data": finalcollectionlist}

	bytedata, _ := json.Marshal(data)

	baseurl := os.Getenv("BASE_URL")

	urlpath := map[string]interface{}{"path": baseurl + "uploadb64image", "payload": "imagedata", "success": map[string]interface{}{"imagepath": "imagepath", "imagename": "imagename"}}

	ubyte, _ := json.Marshal(urlpath)

	ModuleName, _, _ := ModuleRouteName(c)

	selectedchannel := channelist[len(channelist)-1]

	slchannelid := selectedchannel.Id

	c.HTML(200, "addentry.html", gin.H{"Menu": menu, "linktitle": "Create Entry", "title": ModuleName, "Fields": field, "translate": translate, "csrf": csrf.GetToken(c), "channellist": channelist, "CategoryName": FinalSelectedCategories, "entrycount": entrcount, "HeadTitle": translate.Channell.Channels, "Cmsmenu": true, "Entriestab": true, "currentdate": cdate, "StorageType": selectedtype.SelectedType, "blocks": string(bytedata), "Mode": "create", "Slchannelid": slchannelid, "Storagepath": string(ubyte), "htmldata": htmldata})

	return

	// c.Redirect(301,"/entries/create")
}

func readJSONFile(filePath string) ([]byte, error) {

	file, err := os.ReadFile(filePath)
	if err != nil {
		return file, fmt.Errorf("could not open file: %v", err)
	}

	return file, nil
}
