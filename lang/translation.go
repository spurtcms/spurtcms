package lang

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"spurt-cms/models"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/team"
)

type Translation struct {
	Hello                  string `json:"hello"`
	Signin                 string `json:"signin"`
	Newto                  string `json:"newto"`
	Add                    string `json:"add"`
	Signup                 string `json:"signup"`
	Email                  string `json:"email"`
	Forgotpassword         string `json:"forgotpassword"`
	Login                  string `json:"login"`
	Submit                 string `json:"submit"`
	Back                   string `json:"back"`
	Passresetmsg           string `json:"passresetmsg"`
	Resetpass              string `json:"resetpass"`
	Newpass                string `json:"newpass"`
	Confirmpass            string `json:"confirmpass"`
	Dashboard              string `json:"dashboard"`
	Applications           string `json:"applications"`
	Marketplace            string `json:"marketplace"`
	Roles                  string `json:"roles"`
	Categories             string `json:"categories"`
	Settings               string `json:"settings"`
	Channelentries         string `json:"channelentries"`
	Myprofilequotes        string `json:"myprofilequotes"`
	Manage                 string `json:"manage"`
	Language               string `json:"language"`
	Groups                 string `json:"groups"`
	Customfields           string `json:"customfields"`
	App                    string `json:"app"`
	Search                 string `json:"search"`
	Users                  string `json:"users"`
	Name                   string `json:"name"`
	Role                   string `json:"role"`
	Action                 string `json:"action"`
	Actions                string `json:"actions"`
	Download               string `json:"download"`
	Browse                 string `json:"browse"`
	Oopsnodata             string `json:"oopsnodata"`
	Pagesorry              string `json:"pagesorry"`
	Gobackhome             string `json:"gobackhome"`
	Clickadd               string `json:"clickadd"`
	Admin                  string `json:"admin"`
	Edit                   string `json:"edit"`
	Delete                 string `json:"delete"`
	Records                string `json:"records"`
	Available              string `json:"available"`
	Recordsavailable       string `json:"recordsavailable"`
	Itemsperpage           string `json:"itemsperpage"`
	Nodatafound            string `json:"nodatafound"`
	Data                   string `json:"data"`
	No                     string `json:"no"`
	Found                  string `json:"found"`
	Entries                string `json:"entries"`
	Media                  string `json:"media"`
	Booking                string `json:"booking"`
	Showtime               string `json:"showtime"`
	Create                 string `json:"create"`
	Createuser             string `json:"createuser"`
	Profile                string `json:"profile"`
	Picture                string `json:"picture"`
	First                  string `json:"first"`
	Last                   string `json:"last"`
	Buildataarchitecture   string `json:"buildataarchitecture"`
	Allrightsreserved      string `json:"allrightsreserved"`
	Approver               string `json:"approver"`
	Manager                string `json:"manager"`
	Publisher              string `json:"publisher"`
	English                string `json:"english"`
	French                 string `json:"french"`
	Spanish                string `json:"spanish"`
	Myprofile              string `json:"myprofile"`
	Logout                 string `json:"logout"`
	Save                   string `json:"save"`
	Cancel                 string `json:"cancel"`
	Updatedon              string `json:"updatedon"`
	Update                 string `json:"update"`
	Delmsg                 string `json:"delmsg"`
	Yes                    string `json:"yes"`
	Useforedit             string `json:"useforedit"`
	Usefordelete           string `json:"usefordelete"`
	Rolepermassign         string `json:"rolepermassign"`
	Useforpublish          string `json:"useforpublish"`
	Useforunpublish        string `json:"useforunpublish"`
	Useforcopy             string `json:"useforcopy"`
	Configuration          string `json:"configuration"`
	Selected               string `json:"selected"`
	Changepassword         string `json:"changepassword"`
	Changeyourloginpass    string `json:"changeyourloginpass"`
	Copyrights             string `json:"copyrights"`
	Uploadimage            string `json:"uploadimage"`
	Uploadneworchoose      string `json:"uploadneworchoose"`
	Opendirectory          string `json:"opendirectory"`
	Chooseparentcategories string `json:"chooseparentcategories"`
	Enterdescription       string `json:"enterdescription"`
	Description            string `json:"description"`
	Uploadcoverimage       string `json:"uploadcoverimage"`
	Createdby              string `json:"createdby"`
	Loading                string `json:"loading"`
	Of                     string `json:"of"`
	Preview                string `json:"preview"`
	Config                 string `json:"config"`
	Cropimage              string `json:"cropimage"`
	Crop                   string `json:"crop"`
	Updaterole             string `json:"updaterole"`
	Forgotmailmismatch     string `json:"forgotmailmismatch"`
	Invalidpass            string `json:"invalidpass"`
	Usrrole                string `json:"usrrole"`
	Copy                   string `json:"copy"`
	Lastupdatedon          string `json:"lastupdatedon"`
	Updated                string `json:"updated"`
	Selectall              string `json:"selectall"`
	Next                   string `json:"next"`
	Previous               string `json:"previous"`
	Cropimgdesc            string `json:"cropimgdesc"`
	Zoom                   string `json:"zoom"`
	Rotate                 string `josn:"rotate"`
	Thereare               string `json:"thereare"`
	Thereis                string `json:"thereis"`
	Total                  string `json:"total"`
	Categoriess            string `json:"categoriess"`
	Membergroups           string `json:"membergroups"`
	Memberrestrict         string `json:"memberrestrict"`
	Memberaccess           string `json:"memberaccess"`
	Csearch                string `json:"csearch"`
	Deselectall            string `json:"deselectall"`
	Activateclaim          string `json:"activateclaim"`
	Companyprofile         string `json:"companyprofile"`
	Others                 string `json:"others"`
	Filtersearch           string `json:"filternodatafound"`
	Filterkeyword          string `json:"filtertrykeyword"`
	Itemselected           string `json:"itemselected"`
	Dataconnect            string `json:"dataconnect"`
	ChooseFormats          string `json:"chooseformats"`
	Forms                  string `json:"forms"`
	Aiwriting              string `json:"aiwriting"`
	Formbuilder            string `json:"formbuilder"`
	Notification           string `json:"notification"`
	Contentwriting         string `json:"contentwriting"`
	Getsupport             string `json:"getsupport"`
	Explore                string `json:"explore"`
	Aiwritingheading       string `json:"aiwritingheading"`
	Selling                string `json:"selling"`
	Courses                string `json:"courses"`
	Listing                string `json:"listing"`

	Permission struct {
		Searchroles       string `json:"searchroles"`
		Permissions       string `json:"permissions"`
		Searchpermname    string `json:"searchpermname"`
		Assigntotherole   string `json:"assigntotherole"`
		Searchpermissions string `json:"searchpermissions"`
		Roleexist         string `json:"roleexist"`
		Rolechat          string `json:"rolechat"`
		Descriptionchat   string `json:"descriptionchat"`
	} `json:"Permission"`

	Rolecontent struct {
		Newrole        string `json:"newrole"`
		Lastmodifiedon string `json:"lastmodifiedon"`
		Yettoconfigure string `json:"yettoconfigure"`
		Manage         string `json:"manage"`
		Addnewrole     string `json:"addnewrole"`
		Setpermisson   string `json:"setpermisson"`
		Roledetails    string `json:"roledetails"`
		Enterrole      string `json:"enterrole"`
		Optional       string `json:"optional"`
	} `json:"Rolecontent"`

	Setting struct {
		Title                 string `json:"title"`
		Profiletitle          string `json:"profiletitle"`
		Permissionstitle      string `json:"permissionstitle"`
		Rolesandpermissions   string `json:"rolesandpermissions"`
		Email                 string `json:"email"`
		Emailcontent          string `json:"emailcontent"`
		Data                  string `json:"data"`
		Datacontent           string `json:"datacontent"`
		Basicinfo             string `json:"basicinfo"`
		Myprofilesubheading   string `json:"myprofilesubheading"`
		Firstname             string `json:"firstname"`
		Lastname              string `json:"lastname"`
		Username              string `json:"username"`
		Manageroles           string `json:"manageroles"`
		Managepermissions     string `json:"managepermissions"`
		Managechannels        string `json:"managechannels"`
		Manageusers           string `json:"manageusers"`
		Managecategories      string `json:"managecategories"`
		Managelanguages       string `json:"managelanguages"`
		Contentaccesscontrol  string `json:"contentaccesscontrol"`
		Profilecontent        string `json:"profilecontent"`
		Permissionscontent    string `json:"permissionscontent"`
		System                string `json:"system"`
		Systemcontent         string `json:"systemcontent"`
		Myaccount             string `json:"myaccount"`
		Myaccountcontent      string `json:"myaccountcontent"`
		Personalize           string `json:"personalize"`
		Personalizecontent    string `json:"personalizecontent"`
		Security              string `json:"security"`
		Securitycontent       string `json:"securitycontent"`
		Rolescontent          string `json:"rolescontent"`
		Team                  string `json:"team"`
		Teams                 string `json:"teams"`
		Teamcontent           string `json:"teamcontent"`
		Languagecontent       string `json:"languagecontent"`
		Emailtemplates        string `json:"emailtemplates"`
		Emailtemplatescontent string `json:"emailtemplatescontent"`
		Myprofiletooltip      string `json:"myprofiletooltip"`
		Personalizetooltip    string `json:"personalizetooltip"`
		Securitytooltip       string `json:"securitytooltip"`
		Rolestooltip          string `json:"rolestooltip"`
		Teamstooltip          string `json:"teamstooltip"`
		Languagetooltip       string `json:"languagetooltip"`
		Emailtooltip          string `json:"emailtooltip"`
		Showpassword          string `json:"showpassword"`
		Hidepassword          string `json:"hidepassword"`
		Businessname          string `json:"businessname"`
		Choosecompany         string `json:"choosecompany"`
		Chooselang            string `json:"chooselang"`
		Choosedefault         string `json:"choosedefault"`
		Datetimeformet        string `json:"datetimeformet"`
		Choosedateformat      string `json:"choosedateformat"`
		Setyourdate           string `json:"setyourdate"`
		Setyourtime           string `json:"setyourtime"`
		Timezone              string `json:"timezone"`
		Choosetimezone        string `json:"choosetimezone"`
		Selecttimezone        string `json:"selecttimezone"`
		Searchtimezone        string `json:"searchtimezone"`
		Notimezone            string `json:"notimezone"`
		Companyplcholder      string `json:"companyplcholder"`
		Generalsettings       string `json:"generalsettings"`
		Emailconfig           string `json:"emailconfig"`
		Emailconfcontent      string `json:"emailconfcontent"`
		Smtpcontent           string `json:"smtpcontent"`
		Senderemail           string `json:"senderemail"`
		Emailpass             string `json:"emailpass"`
		Smtphost              string `json:"smtphost"`
		Smtpport              string `json:"smtpport"`
		Smtpplcholder         string `json:"smtpplcholder"`
		Smtpemailplcholder    string `json:"smtpemailplcholder"`
		Smtphostplcholder     string `json:"smtphostplcholder"`
		Smtpportplcholder     string `json:"smtpportplcholder"`
		Smtp                  string `json:"smtp"`
		Environment           string `json:"environment"`
		Imagetypeerror        string `json:"imagetypeerror"`

		Firstnamelable string `json:"firstnamelable"`
		Lastlamelable  string `json:"lastlamelable"`
		Usernamelable  string `json:"usernamelable"`
		Rolelable      string `json:"rolelable"`
		Emaillable     string `json:"emaillable"`
		Mobilelable    string `json:"mobilelable"`
		Aimodule       string `json:"aimodule"`
	} `json:"Setting"`

	Emailtemplate struct {
		Searchtemplates      string `json:"searchtemplates"`
		Updatetemplate       string `json:"updatetemplate"`
		Templatename         string `json:"templatename"`
		Templatesubject      string `json:"templatesubject"`
		Entertemplatesubject string `json:"entertemplatesubject"`
		Templatecontent      string `json:"templatecontent"`
		Entertemplatecontent string `json:"entertemplatecontent"`
		Entertemplatename    string `json:"entertemplatename"`
		EmailName            string `json:"emailname"`
		EmailDesc            string `json:"emaildesc"`
		EmailSub             string `json:"emailsub"`
		EmailCont            string `json:"emailcont"`
		Emailnameplcholder   string `json:"emailnameplcholder"`
		Emaildesplcholder    string `json:"emaildesplcholder"`
		Emailsubplcholder    string `json:"emailsubplcholder"`
		Emailcontplcholder   string `json:"emailcontplcholder"`
		Updatetemplateemail  string `json:"updatetemplateemail"`
		Emailcount           string `json:"emailcount"`
		Emailcounts          string `json:"emailcounts"`
	} `json:"Emailtemplate"`

	Channell struct {
		Headingcontent           string `json:"headingcontent"`
		Entryheadingcont         string `json:"entryheadingcont"`
		Newchannel               string `json:"newchannel"`
		Searchchannelname        string `json:"searchchannelname"`
		Channelconfiguration     string `json:"channelconfiguration"`
		Channelname              string `json:"channelname"`
		Fields                   string `json:"fields"`
		Properties               string `json:"properties"`
		Label                    string `json:"label"`
		Text                     string `json:"text"`
		Datetime                 string `json:"datetime"`
		Link                     string `json:"link"`
		Select                   string `json:"select"`
		Date                     string `json:"date"`
		Textarea                 string `json:"textarea"`
		Textbox                  string `json:"textbox"`
		Checkbox                 string `json:"checkbox"`
		Texteditor               string `json:"texteditor"`
		Radiobutton              string `json:"radiobutton"`
		Section                  string `json:"section"`
		Break                    string `json:"break"`
		Sectionbreak             string `json:"sectionbreak"`
		Displaytext              string `json:"displaytext"`
		Mandatory                string `json:"mandatory"`
		Initialvalue             string `json:"initialvalue"`
		Placeholder              string `json:"placeholder"`
		Customisable             string `json:"customisable"`
		Options                  string `json:"options"`
		Dateformat               string `json:"dateformat"`
		Timeformat               string `json:"timeformat"`
		Small                    string `json:"small"`
		Long                     string `json:"long"`
		Externallink             string `json:"externallink"`
		Slug                     string `json:"slug"`
		Multitext                string `json:"multitext"`
		Multiselectoption        string `json:"multiselectoption"`
		Headertext               string `json:"headertext"`
		Withformatting           string `json:"withformatting"`
		Hours                    string `json:"hours"`
		Visibility               string `json:"visibility"`
		Optionslist              string `json:"optionslist"`
		Grouping                 string `json:"grouping"`
		Searchchanentryname      string `json:"searchchanentryname"`
		Id                       string `json:"id"`
		Filter                   string `json:"filter"`
		Filterby                 string `json:"filterby"`
		Daterange                string `json:"daterange"`
		Fromdate                 string `json:"fromdate"`
		Todate                   string `json:"todate"`
		Applyfilters             string `json:"applyfilters"`
		Clearfilter              string `json:"clearfilters"`
		Newentry                 string `json:"newentry"`
		Createddatetime          string `json:"createddatetime"`
		Time                     string `json:"time"`
		Title                    string `json:"title"`
		Channelselect            string `json:"channelselect"`
		Publish                  string `json:"publish"`
		Unpublish                string `json:"unpublish"`
		Published                string `json:"published"`
		Unpublished              string `json:"unpublished"`
		Copy                     string `json:"copy"`
		Status                   string `json:"status"`
		Draft                    string `json:"draft"`
		Contentc                 string `json:"contentc"`
		Categoryc                string `json:"categoryc"`
		Additionalfields         string `json:"additionalfields"`
		Clear                    string `json:"clear"`
		Relatedarticles          string `json:"relatedarticles"`
		Choosearticlefromlist    string `json:"choosearticlefromlist"`
		Arthur                   string `json:"arthur"`
		Article                  string `json:"article"`
		Entryinformation         string `json:"entryinformation"`
		Createddate              string `json:"createddate"`
		Entrypublished           string `json:"entrypublished"`
		Entrynotpublished        string `json:"entrynotpublished"`
		Selectedrelatedarticle   string `json:"selectedrelatedarticle"`
		Content                  string `json:"content"`
		Category                 string `json:"category"`
		Required                 string `json:"required"`
		Metatitle                string `json:"metatitle"`
		Metadescription          string `json:"metadescription"`
		Metakeyword              string `json:"metakeyword"`
		Metaslug                 string `json:"metaslug"`
		Titletooltip             string `json:"titletooltip"`
		Desctooltip              string `json:"desctooltip"`
		Keywordtooltip           string `json:"keywordtooltip"`
		Slugtooltip              string `json:"slugtooltip"`
		Keywords                 string `json:"keywords"`
		Selectedcategories       string `json:"selectedcategories"`
		Availablecategories      string `json:"availablecategories"`
		Lastupdate               string `json:"lastupdate"`
		Channel                  string `json:"channel"`
		Entrylist                string `json:"entrylist"`
		Entry                    string `json:"entry"`
		Channels                 string `json:"channels"`
		Createentry              string `json:"createentry"`
		Blogavailable            string `json:"blogavailable"`
		Addcategories            string `json:"addcategories"`
		Additionaldata           string `json:"additionaldata"`
		Seo                      string `json:"seo"`
		Saveasdraft              string `json:"savedraft"`
		Generateai               string `json:"generateai"`
		Uploadordragging         string `json:"choosedirectory"`
		Pltypehere               string `json:"pltypehere"`
		Draftsave                string `json:"draftsave"`
		Grade                    string `json:"grade"`
		Searchentries            string `json:"searchentries"`
		Entriestitle             string `json:"entriestitle"`
		Entriesavailable         string `json:"entriesavailable"`
		Createchannel            string `json:"createchannel"`
		Pltitle                  string `json:"pltitle"`
		Pldesc                   string `json:"pldesc"`
		Plkeyword                string `json:"plkeyword"`
		Plslug                   string `json:"plslug"`
		Plsearchcategory         string `json:"plsearchcategory"`
		Categoriesavailable      string `json:"categoryavailable"`
		Titlevaild               string `json:"titleerrvaild"`
		Descerrvaild             string `json:"descerrvaild"`
		Aititle                  string `json:"aititle"`
		Articletitle             string `json:"articeltitle"`
		Heading                  string `json:"heading"`
		Generate                 string `json:"generate"`
		Articlestitle            string `json:"articlestitle"`
		Articlename              string `json:"articlename"`
		Plarticle                string `json:"plarticle"`
		Selectlanguage           string `json:"selectlanguage"`
		Noofheading              string `json:"noofheading"`
		Customheading            string `json:"customheading"`
		Plstypehere              string `json:"pltexthere"`
		Option                   string `json:"option"`
		Url                      string `json:"url"`
		Plurl                    string `json:"plurl"`
		Addfield                 string `json:"addfield"`
		Newsection               string `json:"newsection"`
		Challcategory            string `json:"challcategory"`
		Challdesc                string `json:"challdesc"`
		Pladdoption              string `json:"pladdoption"`
		Characterallow           string `json:"characterallow"`
		Fieldname                string `json:"fieldname"`
		Plcharallow              string `json:"plcharallow"`
		Channelfield             string `json:"channelproperty"`
		Chosefld                 string `json:"chosefld"`
		Channeldesc              string `json:"channeldesc"`
		Plchlname                string `json:"plchlname"`
		Createdesc               string `json:"createchldesc"`
		Createchl                string `json:"createchl"`
		Dateerr                  string `json:"dateerr"`
		Timeerr                  string `json:"timeerr"`
		Optionerr                string `json:"optionerr"`
		Optionserr               string `json:"optionserr"`
		Fielderr                 string `json:"fielderr"`
		Urlerr                   string `json:"urlerr"`
		Selectchannel            string `json:"selectchannel"`
		Configtooltip            string `json:"configtooltip"`
		Editchannel              string `json:"editchannel"`
		Backtoprevious           string `json:"backtoprevious"`
		Descriptionentry         string `json:"descriptionentry"`
		Mychannel                string `json:"mychannel"`
		MyChannels               string `json:"mychannels"`
		Mycollectiondescst       string `json:"mycollectiondescst"`
		Mycollectiondescentry    string `json:"mycollectiondescentry"`
		Mycollectiondescmid      string `json:"mycollectiondescmid"`
		Mycollectiondescchannel  string `json:"mycollectiondescchannel"`
		Mycollectiondescmidnext  string `json:"mycollectiondescmidnext"`
		Mycollectiondescategorie string `json:"mycollectiondescategorie"`
		Mycollectiondescend      string `json:"mycollectiondescend"`

		ChannelNamedesc        string `json:"channelnamedesc"`
		ChannelFieldsdesc      string `json:"channelfieldsdesc"`
		ChannelCatdesc         string `json:"channelcatdesc"`
		AvailableCatdesc       string `json:"availablecatdesc"`
		SelectedCatdesc        string `json:"selectedcatdesc"`
		Step                   string `json:"steps"`
		AddFields              string `json:"addfields"`
		ThisContainsEntries    string `json:"thischannelcontainsentries"`
		Duplicate              string `json:"duplicate"`
		Nochannelsyet          string `json:"nochannelsyet"`
		Createfirstchannels    string `json:"createfirstchannels"`
		Nochannels             string `json:"nochannels"`
		Noentrysyet            string `json:"noentrysyet"`
		Noentrys               string `json:"noentrys"`
		Createfirstentry       string `json:"createfirstentry"`
		Nodraftsyet            string `json:"nodraftsyet"`
		Nodraft                string `json:"nodraft"`
		Createfirstdraft       string `json:"createfirstdraft"`
		Nounpublished          string `json:"nounpublished"`
		Createfirstunpublished string `json:"createfirstunpublished"`
		Channelhead            string `json:"channelhead"`
		Channelnamelable       string `json:"channelnamelable"`
		Descriptionlable       string `json:"descriptionlable"`
	} `json:"Channell"`

	Userss struct {
		User                string `json:"user"`
		Addnewuser          string `json:"addnewuser"`
		Newuser             string `json:"newuser"`
		Searchusers         string `json:"searchusers"`
		Accesspermission    string `json:"accesspermission"`
		Accessgranted       string `json:"accessgranted"`
		Accessgrantallusers string `json:"accessgrantallusers"`
		Accessgrantthisuser string `json:"accessgrantthisuser"`
		Firstname           string `json:"firstname"`
		Lastname            string `json:"lastname"`
		Logincredentials    string `json:"logincredentials"`
		Enterfname          string `json:"enterfname"`
		Enterlname          string `json:"enterlname"`
		Enteremail          string `json:"enteremail"`
		Entermobnum         string `json:"entermobnum"`
		Chooserole          string `json:"chooserole"`
		Enterusername       string `json:"enterusername"`
		Enterpassword       string `json:"enterpassword"`
		Onlyaccessdata      string `json:"onlyaccessdata"`
		Profilepicture      string `json:"profilepicture"`
		Clicktoupload       string `json:"clicktoupload"`
		Active              string `json:"active"`
		Emailid             string `json:"emailid"`
		Mobilenumber        string `json:"mobilenumber"`
		Mobile              string `json:"mobile"`
		Group               string `json:"group"`
		Username            string `json:"username"`
		Password            string `json:"password"`
		Labelfname          string `json:"labelfname"`
		Labellname          string `json:"labellname"`
		Labeluname          string `json:"labeluname"`
		Labelemail          string `json:"labelemail"`
		Labelmobile         string `json:"labelmobile"`
		Entermobilelabel    string `json:"entermobilelabel"`
		Selectrole          string `json:"selectrole"`
		Activecontent       string `json:"activecontent"`
		Permission          string `json:"permission"`
		Updateuser          string `json:"updateuser"`
		Nameexist           string `json:"nameexist"`
		Deactive            string `json:"deactive"`
	} `json:"Userss"`

	Personalize struct {
		Themebackground  string `json:"themebackground"`
		Themecontent     string `json:"themecontent"`
		Logoupload       string `json:"logoupload"`
		Expandlogoupload string `json:"expandlogoupload"`
		Logocontent      string `json:"logocontent"`
		Browsecontent    string `json:"browsecontent"`
	} `json:"Personalize"`

	Security struct {
		Changepassword    string `json:"changepassword"`
		Changepasscontent string `json:"changepasscontent"`
		Newpass           string `json:"newpass"`
		Enternewpass      string `json:"enternewpass"`
		Confirmnewpass    string `json:"confirmnewpass"`
	} `json:"Security"`

	Memberss struct {
		Membershipheading  string `json:"membershipheading"`
		Usersheading       string `json:"usersheading"`
		Settings           string `json:"settings"`
		Alreadyclaimed     string `json:"alreadyclaimed"`
		Activeclaim        string `json:"activeclaim"`
		Addmember          string `json:"addmember"`
		Searchmembers      string `json:"searchmembers"`
		Entermemname       string `json:"entermemname"`
		Entermememail      string `json:"entermememail"`
		Entermemmobnum     string `json:"entermemmobnum"`
		Enterusername      string `json:"enterusrname"`
		Enterlastname      string `json:"enterlstname"`
		Enterpswd          string `json:"enterpswd"`
		Choosegroup        string `json:"choosegroup"`
		Member             string `json:"member"`
		Members            string `json:"members"`
		Name               string `json:"name"`
		Group              string `json:"group"`
		Createdon          string `json:"createdon"`
		Updatedon          string `json:"updatedon"`
		Basicinfo          string `json:"basicinfo"`
		Firstname          string `json:"firstname"`
		Lastname           string `json:"lastname"`
		Email              string `json:"email"`
		Mobile             string `json:"mobile"`
		Username           string `json:"username"`
		Password           string `json:"password"`
		Membergroup        string `json:"membergroup"`
		Active             string `json:"active"`
		Activetoggle       string `json:"activetoggle"`
		Selectmemerr       string `json:"memselerr"`
		Companyprofile     string `json:"companyprofile"`
		Editmember         string `json:"editmember"`
		Updatemember       string `json:"uptmember"`
		Allowereg          string `json:"allowereg"`
		Alloweregcont      string `json:"alloweregcont"`
		Optionmeberlogin   string `json:"optionmeberlogin"`
		Optioncont         string `json:"optioncont"`
		Receivenot         string `json:"receivenot"`
		Reveivecont        string `json:"reveivecont"`
		Selectusers        string `json:"selectusers"`
		Emailnotification  string `josn:"emailnotification"`
		Companyname        string `json:"companyname"`
		Location           string `json:"location"`
		Profilename        string `json:"profilename"`
		Profileslug        string `json:"profileslug"`
		Emailid            string `json:"emailid"`
		Mobilenumber       string `json:"mobilenumber"`
		Dateofclaim        string `json:"dateofclaim"`
		Companyinfo        string `json:"companyinfo"`
		About              string `json:"about"`
		Linkedin           string `json:"linkedin"`
		Twitter            string `json:"twitter"`
		Website            string `json:"website"`
		Metainfo           string `json:"metainfo"`
		Metatagtitle       string `json:"metatagtitle"`
		Metatagkeyword     string `json:"metatagkeyword"`
		Metatagdescription string `json:"metatagdescription"`
		Companylocation    string `json:"companylocation"`
		Enteryour          string `json:"enteryour"`
		Membresupdate      string `json:"membresupdate"`
		Deactive           string `json:"deactive"`
		Membersavilable    string `json:"membersavilable"`
		Memberavilable     string `json:"memberavilable"`
		Nomemberyet        string `json:"nomemberyet"`
		Nomember           string `json:"nomember"`
		Createfirstmember  string `json:"createfirstmember"`

		Firstnamelable   string `json:"firstnamelable"`
		Lastmamelable    string `json:"lastmamelable"`
		Emaillable       string `json:"emaillable"`
		Mobilelable      string `json:"mobilelable"`
		Usernamelable    string `json:"usernamelable"`
		Passwordlable    string `json:"passwordlable"`
		Membergrouplable string `json:"membergrouplable"`

		Companynamelable     string `json:"companynamelable"`
		Profilenamelable     string `json:"profilenamelable"`
		Companysluglable     string `json:"companysluglable"`
		Companylocationlable string `json:"companylocationlable"`
		Aboutlable           string `json:"aboutlable"`
		LinkedInlable        string `json:"linkedInlable"`
		Twitterlable         string `json:"twitterlable"`
		Websitelable         string `json:"websitelable"`
		Metatagtitlelable    string `json:"metatagtitlelable"`
		Metatagkeywordlable  string `json:"metatagkeywordlable"`
		Metadescriptionlable string `json:"metadescriptionlable"`
	} `json:"Memberss"`

	MembersGroup struct {
		Membersgroup           string `json:"membersgroup"`
		Searchgrpname          string `json:"searchgrpname"`
		Membergrpname          string `json:"membergrpname"`
		Membergrpdesc          string `json:"membergrpdesc"`
		Membergroup            string `json:"membergroup"`
		Updategroup            string `json:"updmemgrp"`
		Addmembergrp           string `json:"addmembergrp"`
		Plmembergrpname        string `json:"plmembergrpname"`
		Plmembergrpdesc        string `json:"plmembergrpdesc"`
		MemberGroupNameValid   string `json:"memgrpnamevalid"`
		Lastupdatedon          string `json:"lastupdatedon"`
		Status                 string `json:"status"`
		Group                  string `json:"group"`
		Groups                 string `json:"groups"`
		Deletemembergroup      string `json:"delmembergrp"`
		Membersgroupavilable   string `json:"membersgroupavilable"`
		Membergroupavilable    string `json:"membergroupavilable"`
		Nomembergroupyet       string `json:"nomembergroupyet"`
		Nomembergroup          string `json:"nomembergroup"`
		Createfirstmembergroup string `json:"createfirstmembergroup"`

		Membergroupnamelable        string `json:"membergroupnamelable"`
		Membergroupdescriptionlable string `json:"membergroupdescriptionlable"`
	} `json:"Members_Group"`

	Categoryy struct {
		Headingcontent           string `json:"headingcont"`
		Newgroup                 string `json:"newgroup"`
		Searchcategoryname       string `json:"searchcategoryname"`
		Addnewcategorygrp        string `json:"addnewcategorygrp"`
		Categorygrpname          string `json:"categorygrpname"`
		Categorygrp              string `json:"categorygrp"`
		Categorygrpdescription   string `json:"categorygrpdescription"`
		Addcategory              string `json:"addcategory"`
		Parentcategory           string `json:"parentcategory"`
		Categoryname             string `json:"categoryname"`
		Entercategoryname        string `json:"entercategoryname"`
		Entercategorydescription string `json:"entercategorydescription"`
		Lastupdatedon            string `json:"lastupdatedon"`
		Categorys                string `json:"categorys"`
		Category                 string `json:"category"`
		Searchcategories         string `json:"searchcategories"`
		Addcategry               string `json:"addcategry"`
		Selectcategry            string `json:"selctcategry"`
		Categories               string `json:"categories"`
		Availablecategory        string `json:"availablecategory"`
		Selcatvailderr           string `json:"selcaterr"`

		Nocategoryyet                 string `json:"nocategoryyet"`
		Nocategory                    string `json:"nocategory"`
		Createfirstcategory           string `json:"createfirstcategory"`
		Nocategorygroupyet            string `json:"nocategorygroupyet"`
		Nocategorygroup               string `json:"nocategorygroup"`
		Createfirstcategorygroup      string `json:"createfirstcategorygroup"`
		Categorynamelable             string `json:"categorynamelable"`
		Descriptionlable              string `json:"descriptionlable"`
		Availablecategorylable        string `json:"availablecategorylable"`
		Uploadimagelable              string `json:"uploadimagelable"`
		Categorygrouplable            string `json:"categorygrouplable"`
		Categorygroupdescriptionlable string `json:"categorygroupdescriptionlable"`
	} `json:"Categoryy"`

	Mediaa struct {
		Medialibrary        string `json:"medialibrary"`
		Searchmedianame     string `json:"searchmedianame"`
		Refresh             string `json:"refresh"`
		Addfolder           string `json:"addfolder"`
		Trash               string `json:"trash"`
		Upload              string `json:"upload"`
		Library             string `json:"library"`
		Enterfoldername     string `json:"enterfoldername"`
		Image               string `json:"image"`
		Images              string `json:"images"`
		SearchImg           string `json:"searchimg"`
		Unselectall         string `json:"unselectall"`
		Plenterfolder       string `json:"plenterfolder"`
		Vailderr            string `json:"vailderr"`
		Folder              string `json:"folder"`
		Medialibray         string `json:"medialibray"`
		Mediafile           string `json:"mediafileavailable"`
		Mediafiles          string `json:"mediafilesavailable"`
		MediaSettingUp      string `json:"mediasettingupdate"`
		Storageprovider     string `json:"storageprovider"`
		Storagecont         string `json:"storagecont"`
		Localstorage        string `json:"localstorage"`
		Localstoragecont    string `json:"localstoragecont"`
		Foldername          string `json:"foldername"`
		Amazon              string `json:"amozon"`
		Amazoncont          string `json:"amazoncont"`
		Acckeyid            string `json:"acckeyid"`
		Accidplcholder      string `json:"accidplcholder"`
		Acckeyplcholder     string `json:"acckeyplcholder"`
		Accesskey           string `json:"accesskey"`
		Region              string `json:"region"`
		Regionplcholder     string `json:"regionplcholder"`
		Buckername          string `json:"buckername"`
		Bucketnameplcholder string `json:"bucketnameplcholder"`
		Azure               string `json:"azure"`
		Azurecont           string `json:"azurecont"`
		Azureaccname        string `json:"azureaccname"`
		Azurestoragekey     string `json:"azurestoragekey"`
		Containername       string `json:"containername"`
		Containerplcholder  string `json:"containerplcholder"`
	} `json:"Mediaa"`

	Roless struct {
		Addrole       string `json:"addrole"`
		Rolename      string `json:"rolename"`
		Newrole       string `json:"newrole"`
		Enterrolename string `json:"enterrolename"`
		Enterroledesc string `json:"enterroledesc"`
	} `json:"Roless"`

	Languages struct {
		Searchlanguages     string `json:"searchlanguages"`
		Default             string `json:"default"`
		Code                string `json:"code"`
		Flag                string `json:"flag"`
		Inactive            string `json:"inactive"`
		Languagename        string `json:"languagename"`
		Languagecode        string `json:"languagecode"`
		Makedefault         string `json:"makedefault"`
		Uploadjsonfile      string `json:"uploadjsonfile"`
		Uploadflagpicture   string `json:"uploadflagpicture"`
		Addlanguage         string `json:"addlanguage"`
		NewLanguage         string `json:"newlanguage"`
		Languages           string `json:"languages"`
		Language            string `json:"language"`
		Lastupdatedon       string `json:"lastupdatedon"`
		Status              string `json:"status"`
		Upload              string `json:"upload"`
		Jsonfile            string `json:"jsonfile"`
		Flagpicture         string `json:"flagpicture"`
		Active              string `json:"active"`
		Addnewlanguage      string `json:"addnewlanguage"`
		Basicinfo           string `json:"basicinfo"`
		Uploadimg           string `json:"uploadimg"`
		Uploadjson          string `json:"uploadjson"`
		Selectlangcode      string `json:"selectlangcode"`
		Searchpl            string `json:"searchpl"`
		Langvailderr        string `json:"langcodevalid"`
		Choosefromdirectory string `json:"choosefromdirectory"`
		Browse              string `json:"browse"`
		Langcodeerr         string `json:"langcodeerr"`
		Jsonfileerr         string `json:"jsonfileerr"`
		Choosefrommedia     string `json:"choosefrommedia"`
		Flagfileerr         string `json:"flagfileerr"`
		Updatelanguage      string `json:"updatelanguage"`
		Setasdefault        string `json:"setasdefault"`
	} `json:"Languages"`

	Spaces struct {
		Newspace                                    string `json:"newspace"`
		Availablespaces                             string `json:"availablespaces"`
		Spacename                                   string `json:"spacename"`
		Choosecategoryforthisspace                  string `json:"choosecategoryforspace"`
		Shortdescription                            string `json:"shortdescription"`
		Clone                                       string `json:"clone"`
		Createaspace                                string `json:"createaspace"`
		Aspaceiswhereyoucanstartorganisingyourideas string `json:"spacedesc"`
		Space                                       string `json:"space"`
		Spacess                                     string `json:"spaces"`
		Searchspacename                             string `json:"searchspacename"`
		Filter                                      string `json:"filter"`
		Searchcategory                              string `json:"searchcategory"`
		Addpage                                     string `json:"addpage"`
		Plspacename                                 string `json:"plspacename"`
		Plspacedesc                                 string `json:"plspacedesc"`
		Category                                    string `json:"category"`
		Selectcategory                              string `json:"slectcatgory"`
		Categoryerr                                 string `json:"spacecategoryvalid"`
		Spaceimg                                    string `json:"spaceimg"`

		Lms             string `json:"lms"`
		Movepage        string `json:"movepage"`
		Choosepagegroup string `json:"choosepagegroup"`
		Publish         string `json:"publish"`
		Revision        string `json:"revision"`
		Pllckeditor     string `json:"pllckeditor"`
		Newpagegroup    string `json:"newpagegroup"`
		Newpage         string `json:"newpage"`
		Addnew          string `json:"addnew"`
		Pages           string `json:"pages"`
		Createpages     string `json:"createpages"`
		Plentertitle    string `json:"plentertitle"`
		Managepages     string `json:"managepages"`
		Readingtime     string `json:"readingtime"`
		Close           string `json:"expandclose"`
	} `json:"Spaces"`

	ContentAccessControl struct {
		Memberrestrict            string `json:"memberrestrict"`
		Contentaccess             string `json:"contentaccess"`
		Grantaccess               string `json:"grantaccess"`
		Accessgrantedmembergroups string `json:"accessgrantedmember"`
		Space                     string `json:"space"`
		Permissionaccess          string `json:"permissionaccess"`
		Title                     string `json:"title"`
		Searchmemberaccess        string `json:"searchmemberaccess"`
		Totalrecordsavailable     string `json:"totalrecordsavailable"`
		Channels                  string `json:"channels"`
		Accesstomembergroups      string `json:"accesstomembergroups"`
		Spaces                    string `json:"spaces"`
		Content                   string `json:"content"`
		Titlevailderr             string `json:"titleerr"`
		Memberaccesstitle         string `json:"plmemtitle"`
		Contentdesc               string `json:"contentdesc"`
		Spacedesc                 string `json:"spacedesc"`
		Memberdesc                string `json:"memberdesc"`
		Creatememberacc           string `json:"creatememberacc"`
		Updatedon                 string `json:"updatedon"`
		Memberrestrictsavailble   string `json:"memberrestrictsavailble"`
		Memberretrictavailble     string `json:"memberretrictavailble"`
		Norestrictyet             string `json:"norestrictyet"`
		Norestrict                string `json:"norestrict"`
		Createfirstrestrict       string `json:"createfirstrestrict"`

		Titlelable string `json:"titlelable"`
	} `json:"ContentAccessControl"`

	DashBoard struct {
		Lastactive            string `json:"lastactive"`
		Welcome               string `json:"welcome"`
		Descwelcome           string `json:"contents"`
		Userwelcome           string `json:"userwelcome"`
		Editprofile           string `json:"editprofile"`
		Increased             string `json:"increased"`
		Days                  string `json:"days"`
		Createnewcontent      string `json:"createnewcontent"`
		Createdesc            string `json:"createnewcontentdesc"`
		Readyaddchannel       string `json:"readyaddchannel"`
		Readyaddchanneldesc   string `json:"readyaddchanneldesc"`
		Clicktocreate         string `json:"clicktocreate"`
		Recentlyadded         string `json:"recentlyadded"`
		Mostviewcontent       string `json:"mostviewcontent"`
		Recentlyviewedcontent string `json:"recentlyviewedcontent"`
		User                  string `json:"user"`
		Nocontentfound        string `json:"nocontentfound"`
		Recentactivities      string `json:"recentactivities"`
		Recentactivitiesdesc  string `json:"recentactivitiesdesc"`
		Contenttype           string `json:"contenttype"`
		Publishedan           string `json:"publishedan"`
		Ago                   string `json:"ago"`
		Emptyactivity         string `json:"emptyactivity"`
		Memberscontent        string `json:"memberscontent"`
		Emptymember           string `json:"emptymember"`
		Accesspermission      string `json:"accesspermission"`
	} `json:"DashBoard"`

	Datas struct {
		Datadesc       string `json:"datadesc"`
		Dataimport     string `json:"dataimport"`
		Dataexport     string `json:"dataexport"`
		Importfile     string `json:"importfile"`
		Importimage    string `json:"importimage"`
		Fieldsmapping  string `json:"fieldsmapping"`
		Datamigration  string `json:"datamigration"`
		Dataimportdesc string `json:"dataimportdesc"`
		Chosechannel   string `json:"chosechannel"`
		Selchannel     string `json:"selchannel"`
		Pleaseselchl   string `json:"pleaseselchl"`
		Uploadtitle    string `json:"uploadtitle"`
		Browsefile     string `json:"browsefile"`
		Uploadyour     string `json:"uploadyour"`
		Xlsx           string `json:"xlsx"`
		Filehere       string `json:"filehere"`
		Fileerr        string `json:"fileerr"`
		Importimgdesc  string `json:"importimgdesc"`
		Choseimg       string `json:"choseimg"`
		Fieldmapdesc   string `json:"fieldmapdesc"`
		Fieldinfile    string `json:"fieldinfile"`
		Fieldinchannel string `json:"fieldinchannel"`
		Datamigdesc    string `json:"datamigdesc"`
		Successmsg     string `json:"successmsg"`
		Started        string `json:"started"`
		Filename       string `json:"filename"`
		Uploadfile     string `json:"uploadfile"`
		Total          string `json:"total"`
		Mapfield       string `json:"mapfield"`
		Unmapfield     string `json:"unmapfield"`
		Resetmapping   string `json:"resetmapping"`
		Automapping    string `json:"automapping"`
		Exportfile     string `json:"exportfile"`
		Exportdesc     string `json:"exportdesc"`
		EntryName      string `json:"entryname"`
		Plsearch       string `json:"plsearch"`
		Sno            string `json:"sno"`
	} `json:"Datas"`

	Ecommerce struct {
		Ecommerce                 string `json:"ecommerce"`
		Orderid                   string `json:"orderid"`
		Customername              string `json:"customername"`
		Price                     string `json:"price"`
		Orderedcreated            string `json:"orderedcreated"`
		Orderedupdated            string `json:"orderedupdated"`
		Memberid                  string `json:"memberid"`
		Emailid                   string `json:"emailid"`
		Ordercount                string `json:"ordercount"`
		Status                    string `json:"status"`
		Customers                 string `json:"customers"`
		TotalCustomer             string `json:"totalcustomer"`
		Info                      string `json:"info"`
		CustomerInfo              string `json:"customerinfo"`
		Addressinfo               string `json:"addressinfo"`
		Addcustomer               string `json:"Addcustomer"`
		FirstName                 string `json:"firstname"`
		Lastname                  string `json:"lastname"`
		Email                     string `json:"email"`
		Mobile                    string `json:"mobile"`
		Username                  string `json:"username"`
		Password                  string `json:"password"`
		Streetadd                 string `json:"streetaddress"`
		City                      string `json:"city"`
		State                     string `json:"state"`
		Country                   string `json:"country"`
		Zipcode                   string `json:"zipcode"`
		Active                    string `json:"active"`
		StatusContent             string `json:"statuscontent"`
		Firstnameerr              string `json:"firstnameerr"`
		Emailerr                  string `json:"emailerr"`
		Mobileerr                 string `json:"mobileerr"`
		Usernameerr               string `json:"usernameerr"`
		Passerr                   string `json:"passerr"`
		Updatecustomer            string `json:"updatecustomer"`
		Customerinformation       string `json:"customerinformation"`
		Orderitems                string `json:"orderitems"`
		Orderdate                 string `json:"orderdate"`
		Customer                  string `json:"customer"`
		Billingadd                string `json:"billingadd"`
		Shippingadd               string `json:"shippingadd"`
		Itemsordered              string `json:"itemsordered"`
		Item                      string `json:"item"`
		Items                     string `json:"items"`
		Quantity                  string `json:"quantity"`
		Total                     string `json:"total"`
		Subtotal                  string `json:"subtotal"`
		Salestax                  string `json:"salestax"`
		Grandtotal                string `json:"grandtotal"`
		Orderstatus               string `json:"orderstatus"`
		Orderconfirmed            string `json:"orderconfirmed"`
		Shipped                   string `json:"shipped"`
		Outofdelivery             string `json:"outofdelivery"`
		Delivered                 string `json:"delivered"`
		Invoice                   string `json:"invoice"`
		Addplaceholder            string `json:"addplaceholder"`
		Cityplaceholder           string `json:"cityplaceholder"`
		Stateplaceholder          string `json:"stateplaceholder"`
		Countryplaceholder        string `json:"countryplaceholder"`
		Zipplaceholder            string `json:"zipplaceholder"`
		Inactive                  string `json:"inactive"`
		Address                   string `json:"address"`
		Orders                    string `json:"orders"`
		Order                     string `json:"order"`
		Products                  string `json:"products"`
		Product                   string `json:"product"`
		Productname               string `json:"productname"`
		Addproduct                string `json:"addproduct"`
		Generalinfo               string `json:"generalinfo"`
		Productnameplch           string `json:"productnameplch"`
		Sku                       string `json:"sku"`
		Skuplcholder              string `json:"skuplcholder"`
		Skuerr                    string `json:"skuerr"`
		Skuexistserr              string `json:"skuexistserr"`
		Selectstatus              string `json:"selectstatus"`
		Moreinformation           string `json:"moreinformation"`
		Pricing                   string `json:"pricing"`
		Videos                    string `json:"videos"`
		Productcost               string `json:"productcost"`
		Enteramount               string `json:"enteramount"`
		Productamterr             string `json:"productamterr"`
		Tax                       string `json:"tax"`
		Totalcost                 string `json:"totalcost"`
		Discount                  string `json:"discount"`
		Priceerr                  string `json:"priceerr"`
		Priority                  string `json:"priority"`
		Priorityerr               string `json:"priorityerr"`
		Startdate                 string `json:"startdate"`
		Startdateerr              string `json:"startdateerr"`
		Enddate                   string `json:"enddate"`
		Enddateerr                string `json:"enddateerr"`
		Special                   string `json:"special"`
		Productimgerr             string `json:"productimgerr"`
		Youtubeurl                string `json:"youtubeurl"`
		Urlplcholder              string `json:"urlplcholder"`
		Vimeourl                  string `json:"vimeourl"`
		Editproduct               string `json:"editproduct"`
		Qty                       string `json:"qty"`
		Orderinfo                 string `json:"orderinfo"`
		Catelogue                 string `json:"catelogue"`
		StoreInfo                 string `json:"storeinfo"`
		CurrencyInfo              string `json:"currencyinfo"`
		PaymentInfo               string `json:"paymentinfo"`
		StatusInfo                string `json:"statusinfo"`
		Stock                     string `json:"stock"`
		DisplayStock              string `json:"displaystock"`
		Warning                   string `json:"warning"`
		Checkout                  string `json:"checkout"`
		StoreInformation          string `json:"storeinformation"`
		StoreName                 string `json:"storename"`
		Currency                  string `json:"currency"`
		CurrencyTitle             string `json:"currencytitle"`
		Code                      string `json:"code"`
		Symbol                    string `json:"symbol"`
		LastUpdatedon             string `json:"lastupdatedon"`
		Payment                   string `json:"payment"`
		PaymentMethod             string `json:"paymentmethod"`
		Method                    string `json:"method"`
		OrderStatusTitle          string `json:"orderstatustitle"`
		DisplayStockTitle         string `json:"displaystocktitle"`
		ShowOutOfStockWarning     string `json:"showoutofstockwarning"`
		StockCheckoutTitle        string `json:"stockcheckouttitle"`
		AddCurrency               string `json:"addcurrency"`
		CurrencyCode              string `json:"currencycode"`
		CurrencySymbol            string `json:"currencysymbol"`
		SetasDefault              string `json:"setasdefault"`
		AddPaymentMethod          string `json:"addpaymentmethod"`
		PaymentName               string `json:"paymentname"`
		AddOrderStatus            string `json:"addorderstatus"`
		StatusName                string `json:"statusname"`
		ColorCode                 string `json:"colorcode"`
		Productcaterr             string `json:"productcaterr"`
		StockPlacholder           string `json:"stockplaceholder"`
		StockErr                  string `json:"stockerr"`
		PriceRange                string `json:"pricerange"`
		Lowtohigh                 string `json:"lowtohigh"`
		Hightolow                 string `json:"hightolow"`
		StatusNamePlaceholder     string `json:"statusnameplholder"`
		StatusDescPlaceholder     string `json:"statusdescplholder"`
		StatusPriorityPlaceholder string `json:"statuspriorityplholder"`
		PaymentDescPlaceholder    string `json:"paymentdescplholder"`
		PaymentNamePlaceholder    string `json:"paymentnameplholder"`
		CurrencyNamePlaceholder   string `json:"currencynameplholder"`
		CurrencyTypePlaceholder   string `json:"currencytypeplholder"`
		CurrencySymbolPlaceholder string `json:"currencysymbolplholder"`
		StoreNamePlacholder       string `json:"storenameplac"`
		ProductSlug               string `json:"productslug"`
		ProductSlugPlaceholder    string `json:"productslugplaceholder"`
		ProductDesc               string `json:"productdesc"`
	} `json:"Ecommerce"`

	Jobs struct {
		AddNewJobs          string `json:"addnewjobs"`
		JobId               string `json:"jobid"`
		Categories          string `json:"categories"`
		JobType             string `json:"jobtype"`
		PostedDate          string `json:"posteddate"`
		Status              string `json:"status"`
		Info                string `json:"info"`
		AddNewJob           string `json:"addnewjob"`
		JobInfo             string `json:"jobinformation"`
		JobTitle            string `json:"jobtitle"`
		JobtitlePlcHolder   string `json:"jobtitleplcholder"`
		Location            string `json:"location"`
		JobLocPlcHolder     string `json:"joblocationplcholder"`
		JobDesc             string `json:"jobdesc"`
		JobDescPlcholder    string `json:"jobdescplcholder"`
		Keywords            string `json:"keywords"`
		KeywordPlcholder    string `json:"keywordplcholder"`
		SkillContent        string `json:"skillcontent"`
		Skill               string `json:"skill"`
		SkilPlcholder       string `json:"skilplcholder"`
		Yearsofexp          string `json:"yearsofexp"`
		Expcontent          string `json:"expcontent"`
		Minimumyrs          string `json:"minimumyrs"`
		Miniplcholder       string `json:"miniplcholder"`
		Maximumyrs          string `json:"maximumyrs"`
		Maxplcholder        string `json:"maxplcholder"`
		Salary              string `json:"salary"`
		Salaryplcholder     string `json:"salaryplcholder"`
		Salaryinlpa         string `json:"salaryinlpa"`
		Salarycontent       string `json:"salarycontent"`
		Qualification       string `json:"qualification"`
		Qualificontent      string `json:"qualificontent"`
		Education           string `json:"education"`
		Educationplcholder  string `json:"educationplcholder"`
		Jobperiod           string `json:"jobperiod"`
		Jobperiodcontent    string `json:"jobperiodcontent"`
		Validthrough        string `json:"validthrough"`
		Jobstatus           string `json:"jobstatus"`
		Jobstatuscontent    string `json:"jobstatuscontent"`
		Active              string `json:"active"`
		Activecontent       string `json:"activecontent"`
		Editjob             string `json:"editjob"`
		Jobinfo             string `json:"jobinfo"`
		Back                string `json:"back"`
		Applicants          string `json:"applicants"`
		Applicantname       string `json:"applicantname"`
		Userid              string `json:"userid"`
		Emailid             string `json:"emailid"`
		Jobcode             string `json:"jobcode"`
		TotalJobs           string `json:"totaljobs"`
		Totaljob            string `json:"totaljob"`
		Totalapplicants     string `json:"totalapplicants"`
		Totalapplicant      string `json:"totalapplicant"`
		Name                string `json:"name"`
		Nameplcholder       string `json:"nameplcholder"`
		Addapplicants       string `json:"addapplicants"`
		Personalinfo        string `json:"personalinfo"`
		Mobile              string `json:"mobile"`
		Mobileplcholder     string `json:"mobileplcholder"`
		Password            string `json:"password"`
		Appsalcontent       string `json:"appsalcontent"`
		Currentsalary       string `json:"currentsalary"`
		Currentsalplcholder string `json:"currentsalplcholder"`
		Expectedsalary      string `json:"expectedsalary"`
		Expecplcholder      string `json:"expecplcholder"`
		Yrofgraduation      string `json:"yrofgraduation"`
		Graplcholder        string `json:"graplcholder"`
		Appquacontent       string `json:"appquacontent"`
		Appstatus           string `json:"appstatus"`
		Appstatuscontent    string `json:"appstatuscontent"`
		Appactivecont       string `json:"appactivecont"`
		Editapplicant       string `json:"editapplicant"`
		Appinformation      string `json:"appinformation"`
		Resume              string `json:"resume"`
		Applieddate         string `json:"applieddate"`
		Appliedjobs         string `json:"appliedjobs"`
		Type                string `json:"type"`
		Selectcategory      string `json:"selectcategory"`
		Selectjobtype       string `json:"selectjobtype"`
		Jobtypeval          string `json:"jobtypeval"`
		Addapplicant        string `json:"addapplicant"`
		Emailplcholder      string `json:"emailplcholder"`
		Locationplcholder   string `json:"locationplcholder"`
		Experienceplcholder string `json:"experiencecplcholder"`
		Skills              string `json:"skills"`
		Appskilplcholder    string `json:"appskilplcholder"`
		About               string `json:"about"`
		Applicantinfo       string `json:"applicantinfo"`
		Job                 string `json:"job"`
	} `json:"Jobs"`

	Blocks struct {
		Blockheading            string `json:"blockheading"`
		Block                   string `json:"block"`
		NewBlock                string `json:"newblock"`
		MyCollection            string `json:"mycollection"`
		Mycollectiondescst      string `json:"mycollectiondescst"`
		Mycollectiondescentry   string `json:"mycollectiondescentry"`
		Mycollectiondescmid     string `json:"mycollectiondescmid"`
		Mycollectiondescchannel string `json:"mycollectiondescchannel"`
		Mycollectiondescend     string `json:"mycollectiondescend"`
		HowToUse                string `json:"howtouse"`
		ExploreDocumentation    string `json:"exploredocumentation"`
		Setaspremium            string `json:"setaspremium"`
		Blockhtml               string `json:"blockhtml"`
		Blockcss                string `json:"blockcss"`
		Tags                    string `json:"tags"`
		Note                    string `json:"note"`
		Notedesc                string `json:"notedesc"`
		Titleplaceholder        string `json:"titleplaceholder"`
		HtmlPlaceholder         string `json:"htmlplaceholder"`
		CssPlaceholder          string `json:"cssplaceholder"`
		TagPlaceholder          string `json:"tagplaceholder"`
		AddCollectiontooltip    string `json:"addcollectiontooltip"`
		Removetooltip           string `json:"removetooltip"`
		Blocks                  string `json:"blocks"`
		Noblockyet              string `json:"noblockyet"`
		Noblock                 string `json:"noblock"`
		Createfirstblock        string `json:"createfirstblock"`
		MyCollections           string `json:"mycollections"`
		Customblocks            string `json:"customblocks"`

		Titlelable            string `json:"titlelable"`
		Blockhtmllable        string `json:"blockhtmllable"`
		Selectchannellable    string `json:"selectchannellable"`
		Uploadcoverimagelable string `json:"uploadcoverimagelable"`
	} `json:"Blocks"`

	FormBuilder struct {
		Ctaheading             string `json:"ctaheading"`
		Cta                    string `json:"cta"`
		FormAvailable          string `json:"formavailable"`
		FormsAvailable         string `json:"formsavailable"`
		FormName               string `json:"formname"`
		TotalResponse          string `json:"totalresponse"`
		CreatedBy              string `json:"createdby"`
		Action                 string `json:"action"`
		CreateForms            string `json:"createforms"`
		Save                   string `json:"save"`
		Publish                string `json:"publish"`
		CreateForm             string `json:"createform"`
		UpdateForm             string `json:"updateform"`
		SelectFilters          string `json:"selectfilter"`
		FromDate               string `json:"fromdate"`
		ToDate                 string `json:"todate"`
		Clear                  string `json:"clear"`
		ApplyFilter            string `json:"applyfilter"`
		DeleteForm             string `json:"deleteform"`
		Delete                 string `json:"delete"`
		AreyouSureWantDelete   string `json:"areyousurewantdelete"`
		Noformyet              string `json:"noformyet"`
		Noform                 string `json:"noform"`
		Createfirstform        string `json:"createfirstform"`
		Nodraftsyet            string `json:"nodraftsyet"`
		Nodraft                string `json:"nodraft"`
		Createfirstdraft       string `json:"createfirstdraft"`
		Nounpublished          string `json:"nounpublished"`
		Createfirstunpublished string `json:"createfirstunpublished"`
		Ctadesc                string `json:"ctadesc"`
		Fbanner                string `json:"fbanner"`
		Sbanner                string `json:"sbanner"`
		Tbanner                string `json:"tbanner"`
		Fobanner               string `json:"fobanner"`
	} `json:"FormBuilder"`

	Support struct {
		Supporthead  string `json:"supporthead"`
		Title        string `json:"title"`
		Supportdec   string `json:"supportdec"`
		Servicehead  string `json:"servicehead"`
		Servicelist1 string `json:"servicelist1"`
		Servicelist2 string `json:"servicelist2"`
		Servicelist3 string `json:"servicelist3"`
		Servicelist4 string `json:"servicelist4"`

		Namehead        string `json:"namehead"`
		Nameplaceholder string `json:"nameplaceholder"`

		Emailhead         string `json:"emailplaceholder"`
		Emailplaceholders string `json:"emailplaceholders"`

		Contanthead        string `json:"contenthead"`
		Contactplaceholder string `json:"contactplaceholder"`

		Timezonehead        string `json:"timezonehead"`
		Timezoneplaceholder string `json:"timezoneplaceholder"`

		Countryhead        string `json:"countryhead"`
		Countryplaceholder string `json:"countryplaceholder"`

		Dechead        string `json:"dechead"`
		Decplaceholder string `json:"decplaceholder"`

		Submitbtn string `json:"submitbtn"`
	} `json:"Support"`

	Templates struct {
		Templatesheading        string `json:"templatesheading"`
		Title                   string `json:"title"`
		BannerHeading           string `json:"bannerheading"`
		Mycollectiondescst      string `json:"mycollectiondescst"`
		Mycollectiondescentry   string `json:"mycollectiondescentry"`
		Mycollectiondescmid     string `json:"mycollectiondescmid"`
		Mycollectiondescchannel string `json:"mycollectiondescchannel"`
		Mycollectiondescend     string `json:"mycollectiondescend"`
		PlayButtonDescription   string `json:"playbuttondescription"`
		BookButtonDescription   string `json:"bookbuttondescription"`
		GithubButton            string `json:"githubbutton"`
		DeployButton            string `json:"deploybutton"`
		TotalTemplates          string `json:"totaltemplates"`
		TotalTemplate           string `json:"totaltemplate"`
		Helpcenter              string `json:"helpcenter"`
		Helpcenterdes           string `json:"helpcenterdes"`
		Jobs                    string `json:"jobs"`
		Jobsdes                 string `json:"jobsdes"`
		Blogs                   string `json:"blogs"`
		Blogdes                 string `json:"blogdes"`
		News                    string `json:"news"`
		Newsdes                 string `json:"newsdes"`
	} `json:"Templates"`
	Graphql struct {
		Graphqlplayheading      string `json:"graphqlplayheading"`
		Graphqlapiheading       string `json:"graphqlapiheading"`
		Title                   string `json:"title"`
		NewUserHeading          string `json:"newuserheading"`
		NewUserDescription      string `json:"newuserdescription"`
		CreateToken             string `json:"createtoken"`
		CancelButton            string `json:"cancelbutton"`
		GenerateButton          string `json:"generatebutton"`
		NameFieldHeading        string `json:"namefieldheading"`
		NameFieldError          string `json:"namefielderror"`
		DescriptionFieldHeading string `json:"descriptionfieldheading"`
		DurationFieldHeading    string `json:"durationfieldheading"`
		DurationFieldError      string `json:"durationfielderror"`
		DurationSelect          string `json:"durationselect"`
		SevenDays               string `json:"sevendays"`
		FifteenDays             string `json:"fifteendays"`
		ThirtyDays              string `json:"thirtydays"`
		UnlimitedDays           string `json:"unlimiteddays"`
		SubHeading              string `json:"subheading"`
		RecordsHeading          string `json:"recordsheading"`
		TableNameHeading        string `json:"tablenameheading"`
		TableDescriptionHeading string `json:"tabledescriptionheading"`
		TableDurationHeading    string `json:"tabledurationheading"`
		TableLastHeading        string `json:"tablelastheading"`
		TableActionHeading      string `json:"tableactionheading"`
		ActionEditHeading       string `json:"actioneditheading"`
		ActionDeleteHeading     string `json:"actiondeleteheading"`
		PaginationBackButton    string `json:"paginationbackbutton"`
		PaginationNextButton    string `json:"paginationnextbutton"`
		Updatetokenheading      string `json:"updatetokenheading"`
		Deletetoken             string `json:"deletetoken"`
		Deletesubheading        string `json:"deletesubheading"`
		Copied                  string `json:"copied"`
		Copyheading             string `json:"copyheading"`
		Copydetails             string `json:"copydetails"`
		Notes                   string `json:"notes"`
		Apikey                  string `json:"apikey"`
		Graphqlplayground       string `json:"graphqlplayground"`

		Apitokennamelable  string `json:"apitokennamelable"`
		Descriptionlable   string `json:"descriptionlable"`
		Tokendurationlable string `json:"tokendurationlable"`
	} `json:"Graphql"`

	Webhooks struct {
		Webhookheading string `json:"webhookheading"`
		Webhook        string `json:"webhook"`
		HeadingDesc    string `json:"headingDesc"`
		NewWebhook     string `json:"newWebhook"`
		Edit           string `json:"edit"`
		Delete         string `json:"delete"`
		Back           string `json:"back"`
		Next           string `json:"next"`
		EventEnable    string `json:"eventEnable"`
		Cancel         string `json:"cancel"`
		Update         string `json:"update"`
		CreateWebhook  string `json:"createWebhook"`
		Create         string `json:"create"`
		WebhookName    string `json:"webhookName"`
		Event          string `json:"event"`
		EndpointUrl    string `json:"endpointUrl"`
		Method         string `json:"method"`
		PayloadType    string `json:"payloadType"`
		Headers        string `json:"headers"`
		PayloadFields  string `json:"payloadFields"`
		EnableWebhook  string `json:"enableWebhook"`
		FilterNoData   string `json:"filterNoData"`
		ChangeKeywords string `json:"changeKeywords"`

		Webhooknamelable   string `json:"webhooknamelable"`
		Eventlable         string `json:"eventlable"`
		Endpointurllable   string `json:"endpointurllable"`
		Methodlable        string `json:"methodlable"`
		Payloadtypelable   string `json:"payloadtypelable"`
		Headerslable       string `json:"headerslable"`
		Payloadfieldslable string `json:"payloadfieldslable"`

		Placeholders struct {
			Header           string `json:"header"`
			Value            string `json:"value"`
			FieldName        string `json:"fieldName"`
			SlctFieldValue   string `json:"slctFieldValue"`
			SlctPayloadType  string `json:"slctPayloadType"`
			SlctMethod       string `json:"slctMethod"`
			SlctEventType    string `json:"slctEventType"`
			EnterEndpointUrl string `json:"endpoinUrl"`
			EnterWebhookName string `json:"enterWebhookName"`
		} `json:"placeholders"`
		Dropdowns struct {
			EntryCreate        string `json:"entryCreate"`
			EntryPublish       string `json:"entryPublish"`
			EntryUnpublish     string `json:"entryUnpublish"`
			EntryDelete        string `json:"entryDelete"`
			EntryTitle         string `json:"entryTitle"`
			ChannelName        string `json:"channelName"`
			EventInitiatorName string `json:"eventInitatorName"`
			EventInitiatorRole string `json:"eventInitiatorRole"`
		} `json:"dropdowns"`
		ValidationErrors struct {
			ReqWebhookName      string `json:"reqWebhookName"`
			ReqUrl              string `json:"reqUrl"`
			ReqMethod           string `json:"reqMethod"`
			ReqEventType        string `json:"reqEventType"`
			ReqFieldName        string `json:"reqFieldName"`
			ReqFieldValue       string `json:"reqFieldValue"`
			ValidateWebhookName string `json:"validateWebhookName"`
		} `json:"validationErrors"`
	} `json:"webhooks"`

	Memberships struct {
		Membershiphead         string `json:"membershiphead"`
		Membershipheaddesc     string `json:"membershipheaddesc"`
		Membershipmember       string `json:"membershipmember"`
		Membershipgroup        string `json:"membershipgroup"`
		Membershiplevel        string `json:"membershiplevel"`
		Membershipsubscription string `json:"membershipsubscription"`
		Membershiporders       string `json:"membershiporders"`

		Recordsavailable string `json:"recordsavailable"`
		Recordavailable  string `json:"recordavailable"`

		Member struct {
			Recordsavailable string `json:"recordsavailable"`
			Recordavailable  string `json:"recordavailable"`
			CreateMember     string `json:"createmember"`
			UpdateMember     string `json:"updatemember"`

			Selectfilters string `json:"selectfilters"`

			Createdon   string `json:"createdon"`
			Clear       string `json:"clear"`
			Applyfilter string `json:"applyfilter"`

			Searchplacehoder string `json:"searchplacehoder"`

			Id           string `json:"id"`
			Name         string `json:"name"`
			Email        string `json:"email"`
			Level        string `json:"level"`
			Subscription string `json:"subscription"`
			Startdate    string `json:"startdate"`
			Enddate      string `json:"enddate"`
			Status       string `json:"status"`
			Action       string `json:"action"`

			Edit   string `json:"edit"`
			Delete string `json:"delete"`

			Memberemptytitle string `json:"memberemptytitle"`
			Memberemptydesc  string `json:"memberemptydesc"`

			// create

			Cancel                string `json:"cancel"`
			Save                  string `json:"save"`
			Active                string `json:"active"`
			Firstname             string `json:"firstname"`
			Firstnameinfo         string `json:"firstnameinfo"`
			Firstnameplacceholder string `json:"firstnameplacceholder"`
			Lastname              string `json:"lastname"`
			Lastnameinfo          string `json:"lastnameinfo"`

			Lastnameplacceholder string `json:"lastnameplacceholder"`
			Emailinfo            string `json:"emailinfo"`
			Emailplacceholder    string `json:"emailplacceholder"`
			Mobile               string `json:"mobile"`
			Mobileinfo           string `json:"mobileinfo"`

			Mobileplacehoder string `json:"mobileplacehoder"`
			Username         string `json:"username"`
			Usernameinfo     string `json:"usernameinfo"`

			Usernameplacehoder string `json:"usernameplacehoder"`
			Password           string `json:"password"`
			Passwordinfo       string `json:"passwordinfo"`

			Passwordplacehoder string `json:"passwordplacehoder"`
		} `json:"member"`

		MembershipGroups struct {
			Recordsavailable      string `json:"recordsavailable"`
			Recordavailable       string `json:"recordavailable"`
			CreateMembersiphgroup string `json:"createMembersiphgroup"`
			UpdateMembersiphgroup string `json:"updateMembersiphgroup"`
			Selectfilters         string `json:"selectfilters"`
			Lastupdatedon         string `json:"lastupdatedon"`
			Clear                 string `json:"clear"`
			Applyfilter           string `json:"applyfilter"`
			Status                string `json:"status"`
			Active                string `json:"active"`
			Inactive              string `json:"inactive"`
			Searchplacehoder      string `json:"searchplacehoder"`

			Group       string `json:"group"`
			Description string `json:"description"`
			Action      string `json:"action"`

			Membershipgroupemptytitle string `json:"membershipgroupemptytitle"`
			Membershipgroupemptydesc  string `json:"membershipgroupemptydesc"`

			Cancel                  string `json:"cancel"`
			Save                    string `json:"save"`
			Groupname               string `json:"groupname"`
			Groupnameinfo           string `json:"groupnameinfo"`
			Groupnameplacehoder     string `json:"groupnameplacehoder"`
			Groupnamedesc           string `json:"groupnamedesc"`
			Groupnamedescinfo       string `json:"groupnamedescinfo"`
			Groupnamedescplacehoder string `json:"groupnamedescplacehoder"`
		} `json:"membershipgroups"`

		MemberShipLevels struct {
			Create               string `json:"create"`
			Membershipavailable  string `json:"membershipavailable"`
			Membershipavailables string `json:"membershipavailables"`
			Id                   string `json:"id"`
			Level                string `json:"level"`
			Fee                  string `json:"fee"`
			Status               string `json:"status"`
			Action               string `json:"action"`
			Edit                 string `json:"edit"`
			Delete               string `json:"delete"`
			Cancel               string `json:"cancel"`
			Save                 string `json:"save"`

			Basicinfo     string `json:"basicinfo"`
			Basicinfodesc string `json:"basicinfodesc"`

			Name                       string `json:"name"`
			Nameinfo                   string `json:"nameinfo"`
			Selectmembership           string `json:"selectmembership"`
			Selectmembershipinfo       string `json:"selectmembershipinfo"`
			Selectmembershipplaceholde string `json:"selectmembershipplaceholde"`

			Selectmembershiplevelgroup            string `json:"selectmembershiplevelgroup"`
			Selectmembershiplevelgroupinfo        string `json:"selectmembershiplevelgroupinfo"`
			Selectmembershiplevelgroupplaceholder string `json:"Selectmembershiplevelgroupplaceholder"`

			Message            string `json:"message"`
			Messageinfo        string `json:"messageinfo"`
			Billingdetails     string `json:"billingdetails"`
			Billingdetailsdesc string `json:"billingdetailsdesc"`

			Initialpayment     string `json:"initialpayment"`
			Initialpaymentinfo string `json:"initialpaymentinfo"`
			Initialpaymentdesc string `json:"initialpaymentdesc"`

			Recurringsubscription     string `json:"recurringsubscription"`
			Recurringsubscriptioninfo string `json:"recurringsubscriptioninfo"`
			Recurringsubscriptiondesc string `json:"recurringsubscriptiondesc"`

			Billingamount     string `json:"billingamount"`
			Billingamountinfo string `json:"billingamountinfo"`
			Billingamountdesc string `json:"billingamountdesc"`

			Per   string `json:"per"`
			Day   string `json:"day"`
			Week  string `json:"week"`
			Year  string `json:"year"`
			Month string `json:"month"`

			Billingcyclelimit     string `json:"billingcyclelimit"`
			Billingcyclelimitinfo string `json:"billingcyclelimitinfo"`
			Billingcyclelimitdesc string `json:"billingcyclelimitdesc"`

			Customtrial     string `json:"Customtrial"`
			Customtrialinfo string `json:"Customtrialinfo"`
			Customtrialdesc string `json:"Customtrialdesc"`

			Trailbillingamount     string `json:"Trailbillingamount"`
			Trailbillingamountinfo string `json:"Trailbillingamountinfo"`
			Trailbillingamountdesc string `json:"Trailbillingamountdesc"`
			Forthefirst            string `json:"forthefirst"`
			Subscriptionpayment    string `json:"Subscriptionpayment"`

			Membershiplevelemptytitle string `json:"membershiplevelemptytitle"`
			Membershiplevelemptydesc  string `json:"membershiplevelemptydesc"`
			Membershipdesc            string `json:"membershipdesc"`
		} `json:"membershiplevels"`

		Subscriptions struct {
			Create                 string `json:"create"`
			Subscriptionavailable  string `json:"subscriptionavailable"`
			Subscriptionavailables string `json:"subscriptionavailables"`

			Subcriptionid             string `json:"subcriptionid"`
			User                      string `json:"user"`
			Level                     string `json:"level"`
			Fee                       string `json:"fee"`
			Gateway                   string `json:"gateway"`
			Paymenton                 string `json:"paymenton"`
			Status                    string `json:"status"`
			Action                    string `json:"action"`
			Edit                      string `json:"edit"`
			Delete                    string `json:"delete"`
			Cancel                    string `json:"cancel"`
			Save                      string `json:"save"`
			Subscriptionemptytitle    string `json:"subscriptionemptytitle"`
			Subscriptionemptydesc     string `json:"subscriptionemptydesc"`
			Selectgateway             string `json:"selectgateway"`
			MemberShiplevel           string `json:"memberShiplevel"`
			Selectlevel               string `json:"Selectlevel"`
			Subscriptiontransactionid string `json:"subscriptiontransactionid"`
			Entertransactionid        string `json:"entertransactionid"`

			Linksubscription     string `json:"linksubscription"`
			Linksubscriptiondesc string `json:"linksubscriptiondesc"`

			Transactionid     string `json:"transactionid"`
			Transactionidinfo string `json:"transactionidinfo"`

			Subscriptiontransactionidinfo string `json:"subscriptiontransactionidinfo"`
			Subscriptiontransactioniddesc string `json:"subscriptiontransactioniddesc"`

			Gatewayinfo string `json:"gatewayinfo"`

			Gatewayenvironment     string `json:"Gatewayenvironment"`
			Gatewayenvironmentinfo string `json:"Gatewayenvironmentinfo"`

			Memberid     string `json:"memberid"`
			Memberidinfo string `json:"memberidinfo"`

			Membershiplevelid     string `json:"membershiplevelid"`
			Membershiplevelidinfo string `json:"membershiplevelidinfo"`
			Selectmember          string `json:"selectmember"`
		} `json:"subscriptions"`
		Orders struct {
			Create                    string `json:"create"`
			Orderavailable            string `json:"orderavailable"`
			Orderavailables           string `json:"orderavailables"`
			Orderid                   string `json:"orderid"`
			User                      string `json:"user"`
			Level                     string `json:"level"`
			Total                     string `json:"total"`
			Gateway                   string `json:"gateway"`
			Paymenton                 string `json:"paymenton"`
			Subscriptiontransactionid string `json:"subscriptiontransactionid"`
			Status                    string `json:"status"`
			Action                    string `json:"action"`
			Edit                      string `json:"edit"`
			Delete                    string `json:"delete"`
			Cancel                    string `json:"cancel"`
			Save                      string `json:"save"`
			Orderemptytitle           string `json:"orderemptytitle"`
			Orderemptydesc            string `json:"orderemptydesc"`
			Enterorderid              string `json:"enterorderid"`
			Entertransactionid        string `json:"entertransactionid"`
			MemberShiplevel           string `json:"memberShiplevel"`
			Selectlevel               string `json:"selectlevel"`

			Memberinformation     string `json:"memberinformation"`
			Memberinformationdesc string `json:"memberinformationdesc"`
			Userid                string `json:"userid"`
			Useridinfo            string `json:"useridinfo"`
			Membershiplevelid     string `json:"membershiplevelid"`
			Membershiplevelidinfo string `json:"membershiplevelidinfo"`
			Selectmember          string `json:"selectmember"`
			Selectmembershiplevel string `json:"selectmembershiplevel"`
			Searchmembershiplevel string `json:"searchmembershiplevel"`

			Billingaddress     string `json:"billingaddress"`
			Billingaddressdesc string `json:"billingaddressdesc"`

			Billingname           string `json:"billingname"`
			Billingnameinfo       string `json:"billingnameinfo"`
			Billingstreet         string `json:"billingstreet"`
			Billingstreetinfo     string `json:"billingstreetinfo"`
			Billingstreet2        string `json:"billingstreet2"`
			Billingstreetinfo2    string `json:"billingstreetinfo2"`
			Billingcity           string `json:"billingcity"`
			Billingcityinfo       string `json:"billingcityinfo"`
			Billingstate          string `json:"billingstate"`
			Billingstateinfo      string `json:"billingstateinfo"`
			Billingpostelcode     string `json:"billingpostelcode"`
			Billingpostelcodeinfo string `json:"billingpostelcodeinfo"`
			Billingcounty         string `json:"billingcounty"`
			Billingcountyinfo     string `json:"billingcountyinfo"`
			Billingphone          string `json:"billingphone"`
			Billingphoneinfo      string `json:"Billingphoneinfo"`

			Paymentinformation     string `json:"paymentinformation"`
			Paymentinformationdesc string `json:"paymentinformationdesc"`

			Subtotal                      string `json:"subtotal"`
			Subtotalinfo                  string `json:"subtotalinfo"`
			Tax                           string `json:"tax"`
			Taxinfo                       string `json:"taxinfo"`
			Totalinfo                     string `json:"totalinfo"`
			Entertotalamount              string `json:"entertotalamount"`
			Paymenttype                   string `json:"paymenttype"`
			Paymenttypeinfo               string `json:"paymenttypeinfo"`
			Enterpaymenttype              string `json:"enterpaymenttype"`
			Paymenttypedesc               string `json:"paymenttypedesc"`
			Statusinfo                    string `json:"statusinfo"`
			Success                       string `json:"success"`
			Failed                        string `json:"failed"`
			Paymentgatewayinformation     string `json:"paymentgatewayinformation"`
			Paymentgatewayinformationdesc string `json:"paymentgatewayinformationdesc"`
			Gatewayinfo                   string `json:"gatewayinfo"`
			Gatewayenvironment            string `json:"Gatewayenvironment"`
			Gatewayenvironmentinfo        string `json:"Gatewayenvironmentinfo"`
			Paymenttransactionid          string `json:"paymenttransactionid"`
			Paymenttransactionidinfo      string `json:"paymenttransactionidinfo"`
			Paymenttransactioniddesc      string `json:"paymenttransactioniddesc"`
			Subscriptiontransactionidinfo string `json:"subscriptiontransactionidinfo"`
			Subscriptiontransactioniddesc string `json:"subscriptiontransactioniddesc"`

			Selectgateway            string `json:"selectgateway"`
			Selectgatewayenvironment string `json:"selectgatewayenvironment"`
		} `json:"orders"`
	} `json:"memberships"`

	Integrations struct {
		Integration string `json:"integration"`
		Headtitle   string `json:"headtitle"`
		Title       string `json:"title"`
		Keypoint1   string `json:"keypoint1"`
		Keypoint2   string `json:"keypoint2"`
		Keypoint3   string `json:"keypoint3"`
		Keypoint4   string `json:"keypoint4"`
		Keypoint5   string `json:"keypoint5"`
		Keypoint6   string `json:"keypoint6"`
		Keypoint7   string `json:"keypoint7"`
		Cloudtite   string `json:"cloudtite"`
		Deschead    string `json:"deschead"`
		List1       string `json:"list1"`
		Sublist1    string `json:"sublist1"`
		List2       string `json:"list2"`
		Sublist2    string `json:"sublist2"`
		List3       string `json:"list3"`
		Sublist3    string `json:"sublist3"`
		List4       string `json:"list4"`
		Sublist4    string `json:"sublist4"`
		List5       string `json:"list5"`
		Sublist5    string `json:"sublist5"`
		Sublist6    string `json:"sublist6"`
		Sublist7    string `json:"sublist7"`
		Sublist8    string `json:"sublist8"`
		Sublist9    string `json:"sublist9"`
		Sublist10   string `json:"sublist10"`
		List6       string `json:"list6"`
		Sublist11   string `json:"sublist11"`
		List7       string `json:"list7"`
		Sublist12   string `json:"sublist12"`
		List8       string `json:"list8"`
		Sublist13   string `json:"sublist13"`
		List9       string `json:"list9"`
		Sublist14   string `json:"sublist14"`
		List10      string `json:"list10"`
		Sublist15   string `json:"sublist15"`
		Sublist16   string `json:"sublist16"`
		Sublist17   string `json:"sublist17"`
		Sublist18   string `json:"sublist18"`
	} `json:"integrations"`

	AImodulesettings struct {
		Cancel                 string `json:"cancel"`
		Save                   string `json:"save"`
		Create                 string `json:"create"`
		Aimodule               string `json:"aimodule"`
		Description            string `json:"description"`
		Lastupdatedon          string `json:"lastupdatedon"`
		Status                 string `json:"status"`
		Model                  string `json:"model"`
		Action                 string `json:"action"`
		Edit                   string `json:"edit"`
		Delete                 string `json:"delete"`
		Aiemptytitle           string `json:"aiemptytitle"`
		Aiemptydesc            string `json:"aiemptydesc"`
		Selectmodule           string `json:"selectmodule"`
		Selectmoduledesc       string `json:"selectmoduledesc"`
		Openai                 string `json:"openai"`
		Openaidesc             string `json:"openaidesc"`
		Google                 string `json:"google"`
		Googledesc             string `json:"googledesc"`
		Deepseek               string `json:"deepseek"`
		Deepseekdesc           string `json:"deepseekdesc"`
		Apikey                 string `json:"apikey"`
		Apikeyplaceholder      string `json:"apikeyplaceholder"`
		Descriptionplaceholder string `json:"descriptionplaceholder"`
		Aimodel                string `json:"aimodel"`
		Selectaimodel          string `json:"selectaimodel"`
	} `json:"AImodulesettings"`

	Securitys struct {
		Security            string `json:"security"`
		Changepassword      string `json:"changepassword"`
		Changepassworddesc  string `json:"changepassworddesc"`
		Newpassword         string `json:"newpassword"`
		Newpasswordinfo     string `json:"Newpasswordinfo"`
		Confirmpassword     string `json:"confirmpassword"`
		Confirmpasswordinfo string `json:"Confirmpasswordinfo"`
		Cancel              string `json:"cancel"`
		Save                string `json:"save"`
	} `json:"Securitys"`

	Course struct {
		Course                string `json:"course"`
		Manageandorganize     string `json:"manageandorganize"`
		Createcourse          string `json:"createcourse"`
		RecordsAvailable      string `json:"recordsavailable"`
		All                   string `json:"all"`
		SelectFilters         string `json:"selectfilters"`
		CourseTitle           string `json:"coursetitle"`
		EnterCourseTitle      string `json:"entercoursetitle"`
		Status                string `json:"status"`
		SelectStatus          string `json:"selectstatus"`
		Pricing               string `json:"pricing"`
		SelectPricing         string `json:"selectpricing"`
		Clear                 string `json:"clear"`
		ApplyFilter           string `json:"applyfilter"`
		SortBy                string `json:"sortby"`
		LastUpdated           string `json:"lastupdated"`
		CreatedDate           string `json:"createddate"`
		Ascending             string `json:"ascending"`
		Descending            string `json:"descending"`
		CreatedBy             string `json:"createdBy"`
		Action                string `json:"action"`
		LastUpdatedOn         string `json:"lastupdatedon"`
		CreatedOn             string `json:"createdon"`
		Publish               string `json:"publish"`
		UnPublish             string `json:"unpublish"`
		Delete                string `json:"delete"`
		Cancel                string `json:"cancel"`
		Save                  string `json:"save"`
		Title                 string `json:"title"`
		EnterYourTitle        string `json:"enteryourtitle"`
		Titleinfo             string `json:"titleinfo"`
		Description           string `json:"description"`
		EnterYourDescription  string `json:"enteryourdescription"`
		Descriptioninfo       string `json:"descriptioninfo"`
		Category              string `json:"category"`
		Categoryinfo          string `json:"categoryinfo"`
		SelectCategory        string `json:"selectcategory"`
		SearchCategory        string `json:"searchcategory"`
		UploadCoverImage      string `json:"uploadcoverimage"`
		UploadCoverImageinfo  string `json:"uploadcoverimageinfo"`
		Browse                string `json:"browse"`
		Uploadnewfiles        string `json:"uploadnewfiles"`
		Chooseonly            string `json:"chooseonly"`
		NoCoursesyet          string `json:"nocoursesyet"`
		Startcreating         string `json:"startcreating"`
		DeleteCourse          string `json:"deletecourse"`
		Areyousure            string `json:"areyousure"`
		Lesson                string `json:"lesson"`
		Section               string `json:"section"`
		Settings              string `json:"settings"`
		NoLessonYet           string `json:"nolessonsyet"`
		Startbuilding         string `json:"startbuilding"`
		CreateYourFirstLesson string `json:"createyourfirstlesson"`
		Newlesson             string `json:"newlesson"`
		NewSection            string `json:"newSection"`
		Addtext               string `json:"addtext"`
		Addembed              string `json:"addembed"`
		Addquiz               string `json:"addquiz"`
		AddFiles              string `json:"addFiles"`
		Createlesson          string `json:"createlesson"`
		Updatelesson          string `json:"updatelesson"`
		Createsection         string `json:"createsection"`
		Updatesection         string `json:"updatesection"`
	} `json:"Course"`

	Coursessettings struct {
		Settings        string `json:"settings"`
		Details         string `json:"details"`
		Title           string `json:"title"`
		Titledesc       string `json:"titledesc"`
		Description     string `json:"description"`
		Desccont        string `json:"desccont"`
		Image           string `json:"image"`
		Imagedesc       string `json:"imagedesc"`
		Chooseimg       string `json:"chooseimg"`
		Category        string `json:"category"`
		Catdesc         string `json:"catdesc"`
		Searchcat       string `json:"searchcat"`
		Options         string `json:"options"`
		Certificate     string `json:"certificate"`
		Cerdesc         string `json:"cerdesc"`
		Cerstatus       string `json:"cerstatus"`
		Comments        string `json:"comments"`
		Commentsdesc    string `json:"commentsdesc"`
		Active          string `json:"active"`
		Inactive        string `json:"inactive"`
		Pricing         string `json:"pricing"`
		Offer           string `json:"offer"`
		Offerdesc       string `json:"offerdesc"`
		Offerchange     string `json:"offerchange"`
		Free            string `json:"free"`
		Freedesc        string `json:"freedesc"`
		Waitlist        string `json:"waitlist"`
		Waitlistdesc    string `json:"waitlistdesc"`
		Paid            string `json:"paid"`
		Paiddesc        string `json:"paiddesc"`
		Availability    string `json:"availability"`
		Status          string `json:"status"`
		Statusdesc      string `json:"statusdesc"`
		Visibility      string `json:"visibility"`
		Visibilitydesc  string `json:"visibilitydesc"`
		Published       string `json:"published"`
		Unpublish       string `json:"unpublish"`
		Draft           string `json:"draft"`
		Selectvisible   string `json:"selectvisible"`
		Public          string `json:"public"`
		Private         string `json:"private"`
		Startdate       string `json:"startdate"`
		Datedesc        string `json:"datedesc"`
		Singuplimit     string `json:"singuplimit"`
		Singuplimitdesc string `json:"singuplimitdesc"`
		Enterlimits     string `json:"enterlimits"`
		Duration        string `json:"duration"`
		Durationdesc    string `json:"durationdesc"`
		Selectduration  string `json:"selectduration"`
		Month           string `json:"month"`
		Year            string `json:"year"`
		Visiblityerr    string `json:"visiblityerr"`
		Dateerr         string `json:"dateerr"`
		Limiterr        string `json:"limiterr"`
		Durationerr     string `json:"durationerr"`
	} `json:"Coursessettings"`
	Menu struct {
		Website            string `json:"website"`
		Templates          string `json:"templates"`
		SEO                string `json:"seo"`
		Setting            string `json:"setting"`
		WebDesc            string `json:"webdesc"`
		BannerTitle        string `json:"bannertitle"`
		BannerDesc         string `json:"bannerdesc"`
		Menu               string `json:"menu"`
		MenuName           string `json:"menuname"`
		TotalItems         string `json:"totalitems"`
		Status             string `json:"status"`
		LastModified       string `json:"lastmodified"`
		Action             string `json:"action"`
		CreateMenu         string `json:"createmenu"`
		CreateNewMenu      string `json:"createnewmenu"`
		Description        string `json:"description"`
		ManageMenu         string `json:"managemenu"`
		ManageMenuDesc     string `json:"managemenudesc"`
		ChooseMenuItem     string `json:"choosemenuitem"`
		AddMenuItems       string `json:"addmenuitems"`
		AddToMenu          string `json:"addtomenu"`
		MostRecent         string `json:"mostrecent"`
		ViewAll            string `json:"viewall"`
		SelectAll          string `json:"selectall"`
		AddMenuItem        string `json:"addmenuitem"`
		LabelName          string `json:"labelname"`
		LabelNamePlcholder string `json:"labelnameplcholder"`
		NavigationPath     string `json:"navigationpath"`
		NavPlcholder       string `json:"navplcholder"`
		AddItem            string `json:"additem"`
		LabelNameErr       string `json:"labelnameerr"`
		LinkPathErr        string `json:"linkpatherr"`
	} `json:"Menu"`
	Seo struct {
		HomePage             string `json:"homepage"`
		PageTitle            string `json:"pagetitle"`
		HomePageTitle        string `json:"homepagetitle"`
		PageDescription      string `json:"pagedescription"`
		HomePageDescription  string `json:"homepagedescription"`
		PageKeyword          string `json:"pagekeyword"`
		HomePageKeyword      string `json:"homepagekeyword"`
		Save                 string `json:"save"`
		Store                string `json:"store"`
		StoreTitle           string `json:"storetitle"`
		StorePageTitle       string `json:"storepagetitle"`
		StoreDescription     string `json:"storedescription"`
		StorePageDescription string `json:"storepagedescription"`
		StoreKeyword         string `json:"storekeyword"`
		StorePageKeyword     string `json:"storepagekeyword"`
		SiteMap              string `json:"sitemap"`
		Choosefile           string `json:"choosefile"`
		Nofilechose          string `json:"nofilechosen"`
		Update               string `json:"update"`
	} `json:"Seo"`

	WebsiteSettings struct {
		SiteName    string `json:"sitename"`
		Spurtcms    string `json:"spurtcms"`
		SiteLogo    string `json:"sitelogo"`
		SiteFavIcon string `json:"sitefavicon"`
		WebsiteUrl  string `json:"websiteurl"`
		Subdomain   string `json:"subdomain"`
		Spurtcmscom string `json:"spurtcmscom"`
	} `json:"WebsiteSettings"`
}

func LoadTranslation(filepath string) (Translation, error) {

	fileContents, err := ioutil.ReadFile(filepath)

	if err != nil {

		return Translation{}, err
	}

	var translation Translation

	err = json.Unmarshal(fileContents, &translation)

	if err != nil {

		return Translation{}, err
	}

	return translation, nil
}

func TranslateHandler(c *gin.Context) {

	json_folder := os.Getenv("LOCAL_LANGUAGE_PATH")

	lan, _ := c.Cookie("lang")

	userDetails, err := GetRequestScopedTenantDetails(c)

	if err != nil {

		log.Println(err)
	}

	var defaultLanguage models.TblLanguage

	err = models.GetLanguageById(&defaultLanguage, userDetails.DefaultLanguageId)

	if err != nil {

		log.Println(err)
	}

	if lan == "" {

		translation, err := LoadTranslation(defaultLanguage.JsonPath)

		if err != nil {

			translation, err = LoadTranslation(json_folder + "en.json")

			if err != nil {

				c.AbortWithError(http.StatusInternalServerError, err)

				return
			}
		}

		c.Set("currentLanguage", defaultLanguage)

		c.Set("translation", translation)

	} else {

		var language models.TblLanguage

		langId, _ := strconv.Atoi(lan)

		err := models.GetLanguageById(&language, langId)

		if err != nil {

			log.Println(err)
		}

		translation, err := LoadTranslation(language.JsonPath)

		if err != nil {

			translation, err = LoadTranslation(defaultLanguage.JsonPath)

			if err != nil {

				c.AbortWithError(http.StatusInternalServerError, err)

				return
			}
		}

		c.Set("currentLanguage", language)

		c.Set("translation", translation)

	}

	c.Next()

}

func GetRequestScopedTenantDetails(c *gin.Context) (team.TblUser, error) {

	userData, exists := c.Get("userDetails")

	if !exists {

		return team.TblUser{}, errors.New("failed to fetch tenant details")
	}

	return userData.(team.TblUser), nil

}
