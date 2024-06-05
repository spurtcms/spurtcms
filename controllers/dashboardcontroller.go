package controllers

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spurtcms/pkgcontent/channels"
	spaces "github.com/spurtcms/pkgcontent/spaces"
	"github.com/spurtcms/pkgcore/member"
	"github.com/spurtcms/pkgcore/teams"
	csrf "github.com/utrack/gin-csrf"
)

func DashboardView(c *gin.Context) {

	var (
		user_id     = c.GetInt("userid")
		memlist     []member.TblMember
		newEntries  []channels.TblChannelEntries
		allchnentry []channels.TblChannelEntries
		chaid       int
	)

	teamAuth := teams.TeamAuth{Authority: &AUTH}
	memAuth := member.Memberauth{Authority: &AUTH}
	Channel := channels.Channel{Authority: &AUTH}
	space := spaces.Space{Authority: &AUTH}

	menu := NewMenuController(c)

	usercount, Newusercount, _ := teamAuth.DashboardUserCount()
	membercount, Newmembercount, _ := memAuth.DashboardMemberCount()
	Totalentris, Newentrycount, _ := Channel.DashboardEntriesCount()
	Pagescount, Newpagescount, _ := space.DashboardPagesCount()

	NewActive, _ := Channel.DashboardRecentActivites()

	memfilter := member.Filter{Keyword: strings.Trim(c.DefaultQuery("keyword", ""), " ")}
	memflag := false

	memberlist, _, _ := memAuth.ListMembers(0, 0, memfilter, memflag)

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

	Finalchannelslist, _, _ := Channel.GetPermissionChannels(100, 0, channels.Filter{}, flag1)

	sort.Slice(Finalchannelslist, func(i, j int) bool {
		return Finalchannelslist[i].EntriesCount > Finalchannelslist[j].EntriesCount
	})

	maxRecords := 6
	if len(Finalchannelslist) < maxRecords {
		maxRecords = len(Finalchannelslist)
	}

	topsix_channels := Finalchannelslist[:maxRecords]

	chnentry, _, _, _ := Channel.GetAllChannelEntriesList(chaid, 0, 0, channels.EntriesFilter{})

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

	mostlyviewlist, _ := space.MostlyViewList(MLimit)
	recentlyviewlist, _ := space.RecentlyViewList(MLimit)
	activememberlist, _ := memAuth.ActiveMemberList(MLimit)

	var activememberlist1 []member.TblMember

	for _, val := range activememberlist {
		val.DateString = val.LoginTime.In(TZONE).Format(Datelayout)
		activememberlist1 = append(activememberlist1, val)
	}

	translate, _ := TranslateHandler(c)

	c.HTML(200, "dashboard.html", gin.H{"Menu": menu, "csrf": csrf.GetToken(c), "translate": translate, "memberlist": memlist, "recentactivity": NewActive, "Lastlogin": Lastlogin[user_id], "channelslist": topsix_channels, "newentries": newEntries, "usercount": usercount, "Newusercount": Newusercount, "membercount": membercount, "Newmembercount": Newmembercount, "entriscount": Totalentris, "Newentrycount": Newentrycount, "pagescount": Pagescount, "Newpagescount": Newpagescount, "HeadTitle": translate.Dashboard, "title": "Dashboard", "mostlyviewlist": mostlyviewlist, "recentlyviewlist": recentlyviewlist, "activememberlist": activememberlist1, "Dashboardmenu": true})

}

func TruncateDescription(description string, limit int) string {
	if len(description) <= limit {
		return description
	}

	truncated := description[:limit] + "..."
	return truncated
}
