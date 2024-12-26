package controllers

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	chn "github.com/spurtcms/channels"
	mem "github.com/spurtcms/member"
	csrf "github.com/utrack/gin-csrf"
)

func DashboardView(c *gin.Context) {

	var (
		user_id        = c.GetInt("userid")
		memlist        []mem.Tblmember
		newEntries     []chn.Tblchannelentries
		allchnentry    []chn.Tblchannelentries
		chaid          int
		usercount      int
		Newusercount   int
		membercount    int
		Newmembercount int
	)
	menu := NewMenuController(c)

	usercount, Newusercount, _ = NewTeamWP.DashboardUserCount(TenantId)

	// usercount, Newusercount, _ := NewTeamWP.DashboardUserCount(TenantId)
	membercount, Newmembercount, _ = MemberConfigWP.DashboardMemberCount(TenantId)

	Totalentris, Newentrycount, _ := ChannelConfigWP.DashboardEntriesCount(TenantId)

	// Pagescount, Newpagescount, _ := CoursesConfigWP.DashboardPagesCount()

	Newactivelist, _ := ChannelConfigWP.DashboardRecentActivites(TenantId)

	var NewActive []chn.RecentActivities

	for _, val := range Newactivelist {

		if val.Imagepath != "" {
			val.Imagepath = "/image-resize?name=" + val.Imagepath
		}

		NewActive = append(NewActive, val)
	}

	memfilter := mem.Filter{Keyword: strings.Trim(c.DefaultQuery("keyword", ""), " ")}
	memflag := false

	memberlist, _, _ := MemberConfigWP.ListMembers(0, 0, memfilter, memflag, TenantId)

	maxmem := 5

	for i := 0; i < len(memberlist) && i < maxmem; i++ {
		memberlist[i].CreatedDate = memberlist[i].CreatedOn.In(TZONE).Format(Datelayout)
		if !memberlist[i].ModifiedOn.IsZero() {
			memberlist[i].ModifiedDate = memberlist[i].ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			memberlist[i].ModifiedDate = memberlist[i].CreatedOn.In(TZONE).Format(Datelayout)
		}
		memlist = append(memlist, memberlist[i])
	}

	flag1 := true

	Finalchannelslist, _, _ := ChannelConfigWP.ListChannel(chn.Channels{Limit: 100, Offset: 0, IsActive: flag1, TenantId: TenantId})

	sort.Slice(Finalchannelslist, func(i, j int) bool {

		return Finalchannelslist[i].EntriesCount > Finalchannelslist[j].EntriesCount

	})

	var chnallist []chn.Tblchannel

	for _, val := range Finalchannelslist {

		_, _, entrcount, _ := ChannelConfig.ChannelEntriesList(chn.Entries{ChannelId: val.Id, Limit: 0, Offset: 0, Keyword: "", ChannelName: "", Title: "", Status: ""}, TenantId)

		val.EntriesCount = entrcount

		chnallist = append(chnallist, val)

	}
	maxRecords := 6
	if len(chnallist) < maxRecords {
		maxRecords = len(chnallist)
	}

	topsix_channels := chnallist[:maxRecords]

	chnentry, _, _, _ := ChannelConfigWP.ChannelEntriesList(chn.Entries{ChannelId: chaid}, TenantId)

	for _, entry := range chnentry {

		entry.CreatedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		if !entry.ModifiedOn.IsZero() {
			entry.ModifiedDate = entry.ModifiedOn.In(TZONE).Format(Datelayout)
		} else {
			entry.ModifiedDate = entry.CreatedOn.In(TZONE).Format(Datelayout)
		}
		allchnentry = append(allchnentry, entry)
	}

	maxLength := 6

	for i := 0; i < len(allchnentry) && i < maxLength; i++ {
		newEntries = append(newEntries, allchnentry[i])
	}

	activememberlist, _ := MemberConfigWP.ActiveMemberList(MLimit, TenantId)

	var activememberlist1 []mem.Tblmember

	for _, val := range activememberlist {
		val.DateString = val.LoginTime.In(TZONE).Format(Datelayout)
		activememberlist1 = append(activememberlist1, val)
	}

	translate, _ := TranslateHandler(c)

	c.HTML(200, "dashboard.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "memberlist": memlist, "recentactivity": NewActive, "Lastlogin": Lastlogin[user_id], "channelslist": topsix_channels, "newentries": newEntries, "usercount": usercount, "Newusercount": Newusercount, "membercount": membercount, "Newmembercount": Newmembercount, "entriscount": Totalentris, "Newentrycount": Newentrycount, "HeadTitle": translate.Dashboard, "title": translate.Dashboard, "linktitle": "Dashboard", "activememberlist": activememberlist1, "Dashboardmenu": true})

}

func TruncateDescription(description string, limit int) string {
	if len(description) <= limit {
		return description
	}

	truncated := description[:limit] + "..."
	return truncated
}
