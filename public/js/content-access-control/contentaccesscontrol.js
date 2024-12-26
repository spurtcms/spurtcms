var languagedata

var spaces = []

var pagegroups = []

var pages = []

var subpages = []

var access_granted_memgrps = []

var channels = []

var channelEntries = []

var content_access_id

var languagepath = $('.language-group>button').attr('data-path')

$.getJSON(languagepath, function (data) {

  languagedata = data
})

/** */
$(document).ready(function () {

  if ($("#route-type").val() == "Update" || $("#route-type").val() == "Copy") {

    var urlPath = window.location.pathname

    var url_parts = urlPath.split('/')

    content_access_id = url_parts[url_parts.length - 1]

    $.ajax({

      url: '/memberaccess/getaccess-pages',

      type: 'get',

      data: { content_access_id: content_access_id },

      dataType: 'json',

      success: function (data) {

        if (data.Pages != null || data.Subpages != null || data.Pagegroups != null || data.MemberGroups != null || data.Spaces != null) {

          if (data.Pages != null) {

            pages = data.Pages
          }

          if (data.Subpages != null) {

            subpages = data.Subpages
          }

          if (data.Pagegroups != null) {

            pagegroups = data.Pagegroups
          }

          if (data.Spaces != null) {

            spaces = data.Spaces

            for (let spaceId of spaces) {

              $('.spacechkbox[data-id=' + spaceId + ']').prop('checked', true)
            }
          }


          $('.space-chkbox').each(function () {

            var spaceId = $(this).attr('data-id')

            if ($(this).is(':checked')) {

              spaces.push(spaceId)
            }

          })

          if (data.MemberGroups != null) {

            access_granted_memgrps = data.MemberGroups

            var isAllMembersChkd = true

            for (let memgrpId of access_granted_memgrps) {

              $('.memgrp-chkboxes').each(function () {

                if ($(this).attr('data-id') == memgrpId) {

                  $(this).prop('checked', true)

                }


              })
              console.log("condtionchecksss", isAllMembersChkd)



            }
            $('.memgrp-chkboxes').each(function () {

              if (!$(this).is(':checked')) {

                isAllMembersChkd = false;
              }
            });

            if (isAllMembersChkd) {


              $('#memgrp-slctall').prop('checked', true)
            }

          }

          if (data.Channels != null) {

            channels = data.Channels



            // if ($('#default-mod-chk[type=hidden]').val() == "channel") {

            for (let channelId of channels) {

              $('.channelchkbox[data-id=' + channelId + ']').prop('checked', true)
            }
            // }
          }

          if (data.ChannelEntries != null) {

            channelEntries = data.ChannelEntries

            // if ($('#default-mod-chk[type=hidden]').val() == "channel") {

            for (let entry of channelEntries) {
              console.log("loop", entry.id);


              $('.chanEntry-chkbox[data-id=' + entry.id + ']').prop('checked', true)
            }
            // }

          }

          console.log("pages", pages, "subpages", subpages, "pagegroups", pagegroups, "spaces", spaces, "memgrp", access_granted_memgrps, "channels", channels, "entries", channelEntries);

        }
      }

    })

  }

})

$('.transitionSearch').click(function () {

  $(this).addClass('active')
})

$('form[name=contentaccess-form]>img').click(function () {

  if ($(this).siblings('input[name=keyword]').val() != "" && $(this).parents('.transitionSearch').hasClass('active')) {

    var keyword = $(this).siblings('input[name=keyword]').val()

    window.location.href = "/memberaccess/?keyword=" + keyword
  }
})

$('#space-img-div>img').click(function () {

  var spaceid = $(this).attr('data-toggle')

  console.log("spaceid", spaceid);

  if ($(this).parent().attr('aria-expanded') == 'false' && $('#collapse' + spaceid + '>.accessAccord-child-container').html() == "") {

    $(this).parent().attr('aria-expanded', 'true')

    $.ajax({

      type: "get",

      url: "/memberaccess/getpages",

      dataType: 'json',

      data: { sid: spaceid },

      cache: false,

      success: function (result) {

        if (result.Pagegroups != null || result.Pages != null || result.Subpages != null) {

          var pagegroupz = []

          var pagez = []

          var subpagez = []

          var space_contents = []

          if (result.Pagegroups != null) {

            pagegroupz = result.Pagegroups

          }
          if (result.Pages != null) {

            pagez = result.Pages

          }
          if (result.Subpages != null) {

            subpagez = result.Subpages

          }

          space_contents = space_contents.concat(pagez, pagegroupz)

          space_contents.sort((a, b) => {

            return a.OrderIndex - b.OrderIndex;

          })

          subpagez.sort((a, b) => {

            return a.OrderIndex - b.OrderIndex;

          })

          var spaceChecked

          if ($('#space' + spaceid).is(':checked')) {

            spaceChecked = true

            IntegrateSpaceContents(space_contents, subpagez, spaceid, pagez, spaceChecked)

          } else {

            spaceChecked = false

            IntegrateSpaceContents(space_contents, subpagez, spaceid, pagez, spaceChecked)

          }

          $('#collapse' + spaceid).slideToggle()
        }
      }
    })

  } else if ($(this).parent().attr('aria-expanded') == 'false' && $('#collapse' + spaceid + '>.accessAccord-child-container').html() != "") {

    $(this).parent().attr('aria-expanded', 'true')

    $('#collapse' + spaceid).slideToggle()

  } else {

    $(this).parent().attr('aria-expanded', 'false')

    $('#collapse' + spaceid).slideToggle()
  }

})

function IntegrateSpaceContents(overall_array, subpagez, spid, pagez, spaceChecked) {

  for (let data of overall_array) {

    if (data.PgId !== undefined && data.Pgroupid == 0 && data.GroupId === undefined) {

      var containsPgid = pages.some(obj => {

        return obj.id == data.PgId

      });

      if (containsPgid == true) {

        var pageChecked = true

      } else {

        var pageChecked = false

      }

      var containsSubpage = subpagez.some(obj => {

        return obj.ParentId == data.PgId
      })

      var pagehtml = GetPageHtml(data, spid, spaceChecked, pageChecked, containsSubpage)

      $('#collapse' + spid + '>.accessAccord-child-container').append(pagehtml)

      for (let spg of subpagez) {

        if (spg.ParentId == data.PgId) {

          var spgChecked = false

          var containsSpgid = subpages.some(obj => {

            return obj.id == spg.SpgId

          });

          if (containsSpgid) {

            spgChecked = true

          }

          var subpage_html = GetSubpageHtml(spg, spid, spaceChecked, spgChecked)

          $('#collapseChild' + spg.ParentId + '>.accessAccord-child-row').append(subpage_html)

        }
      }
    }

    if (data.GroupId !== undefined && data.GroupId != 0 && data.PgId === undefined) {

      var pggChecked = false

      var containsPggid = pagegroups.some(obj => {

        return obj.id == data.GroupId

      });

      if (containsPggid) {

        pggChecked = true

      }

      var grouphtml = GetPageGroupHtml(data, spid, spaceChecked, pggChecked)

      $('#collapse' + spid + '>.accessAccord-child-container').append(grouphtml)

      for (let page of pagez) {

        if (page.Pgroupid == data.GroupId) {

          var pgChecked = false

          var containsPgid = pages.some(obj => {

            return obj.id == page.PgId

          });

          if (containsPgid) {

            pgChecked = true

          }

          var contains_subpage = subpagez.some(obj => {

            return obj.ParentId == page.PgId
          })

          var pg_underpgg_html = GetPageForPageGroup(page, spid, data.GroupId, spaceChecked, pgChecked, contains_subpage)

          $('#pagegroup' + data.GroupId).append(pg_underpgg_html)

          for (let spg of subpagez) {

            if (spg.ParentId == page.PgId) {

              var containsSpgid = subpages.some(obj => {

                return obj.id == spg.SpgId

              });

              if (containsSpgid) {

                var spgzChecked = true

                var spg_underpg_html = GetSubpageForPageGroup(spg, spid, data.GroupId, page.PgId, spaceChecked, spgzChecked)

                $('#collapseChild' + spg.ParentId + '>.accessAccord-child-row').append(spg_underpg_html)

              } else {

                var spgzChecked = false

                var spg_underpg_html = GetSubpageForPageGroup(spg, spid, data.GroupId, page.PgId, spaceChecked, spgzChecked)

                $('#collapseChild' + spg.ParentId + '>.accessAccord-child-row').append(spg_underpg_html)

              }

            }

          }

        }

      }

    }

  }

}

function GetPageHtml(data, spid, spaceChecked, pageChecked, containsSubpage) {

  var page_chkbox

  if (spaceChecked || pageChecked) {

    page_chkbox = `<input type="checkbox" id="chkpg${data.PgId}" class="dirpage-chkbox SPID${spid}" data-spid="${spid}" data-pgid="${data.PgId}" checked/>`

  } else {

    page_chkbox = `<input type="checkbox" id="chkpg${data.PgId}" class="dirpage-chkbox SPID${spid}" data-spid="${spid}" data-pgid="${data.PgId}"/>`

  }

  var subpagedecider

  if (containsSubpage) {

    subpagedecider = `<button class="accord-collapse" data-bs-toggle="collapse" data-bs-target="#collapseChild${data.PgId}" aria-expanded="false" aria-controls="collapseChildone">
      <a href="javascript:void(0)" class="img-div">
       <img src="/public/img/arrow-down-dark.svg" alt="">
      </a>
      <span class="para">${data.Name}</span>
     </button>`

  } else {

    subpagedecider = `<button class="accord-collapse" data-bs-toggle="collapse" data-bs-target="#collapseChild${data.PgId}" aria-expanded="false" aria-controls="collapseChildone">
      <span class="para">${data.Name}</span>
      </button>`

  }

  return `<div class="accessAccord-child-div accessAccord-child-page" id="directpage${data.PgId}">

            <div class="accessAccord-child">
                <div class="chk-group">
                  ${page_chkbox}
                  <label for="chkpg${data.PgId}"></label>
                </div>

                ${subpagedecider}
            </div>


           <div id="collapseChild${data.PgId}" class="accordion-collapse collapse  accessAccord-child-content"
            aria-labelledby="headingOne" data-bs-parent="#accordionExample">
              <div class="accessAccord-child-row"></div>
            </div>
    </div>`
}

function GetSubpageHtml(data, spid, spaceChecked, spgChecked) {

  var spg_chkbox

  if (spaceChecked || spgChecked) {

    spg_chkbox = `<input type="checkbox" id="chkspg${data.SpgId}" data-spid="${spid}" class="subpage-chkbox SPID${spid} PGID${data.ParentId}" data-spgid="${data.SpgId}" data-pgid="${data.ParentId}" checked/>`

  } else {

    spg_chkbox = `<input type="checkbox" id="chkspg${data.SpgId}" data-spid="${spid}" class="subpage-chkbox SPID${spid} PGID${data.ParentId}" data-spgid="${data.SpgId}" data-pgid="${data.ParentId}"/>`

  }

  return ` <div class="chk-group chk-group-label">
                ${spg_chkbox}
                <label for="chkspg${data.SpgId}" class="para">${data.Name}</label>
             </div>`
}

function GetPageGroupHtml(data, spid, spaceChecked, pggChecked) {

  var pgg_chkbox

  if (spaceChecked || pggChecked) {

    pgg_chkbox = `<input type="checkbox" id="chkpgg${data.GroupId}" data-spid="${spid}" class="pagegroup-chkbox SPID${spid}"  data-pggid="${data.GroupId}" checked/>`

  } else {

    pgg_chkbox = `<input type="checkbox" id="chkpgg${data.GroupId}" data-spid="${spid}" class="pagegroup-chkbox SPID${spid}"  data-pggid="${data.GroupId}"/>`

  }

  return `<div class="accessAccord-child-div accessAccord-child-page" id="pagegroup${data.GroupId}">
                <div class="accessAccord-childgrp">
                  <div class="chk-group">
                     ${pgg_chkbox}
                     <label for="chkpgg${data.GroupId}"></label>
                  </div>
                  <span class="para-light">${data.Name}</span>
                </div>
            </div>`
}

function GetPageForPageGroup(data, spid, pggid, spaceChecked, pgChecked, containsSubpage) {

  var page_chkbox

  if (spaceChecked || pgChecked) {

    page_chkbox = `<input type="checkbox" id="chkpg${data.PgId}" data-spid="${spid}" class="pgunderpgg-chkbox SPID${spid} GRPID${pggid}"  data-pgid="${data.PgId}" data-pggid="${pggid}" checked/>`

  } else {

    page_chkbox = `<input type="checkbox" id="chkpg${data.PgId}" data-spid="${spid}" class="pgunderpgg-chkbox SPID${spid} GRPID${pggid}" data-pgid="${data.PgId}" data-pggid="${pggid}"/>`

  }

  var subpagedecider

  if (containsSubpage) {

    subpagedecider = `<button class="accord-collapse" data-bs-toggle="collapse"
      data-bs-target="#collapseChild${data.PgId}" aria-expanded="false"
      aria-controls="collapseChild${data.PgId}">
      <a href="javascript:void(0)" class="img-div">
        <img src="/public/img/arrow-down-dark.svg" alt="">
       </a>
       <span class="para">${data.Name}</span>
    </button>`

  } else {

    subpagedecider = `<button class="accord-collapse" data-bs-toggle="collapse"
      data-bs-target="#collapseChild${data.PgId}" aria-expanded="false"
      aria-controls="collapseChild${data.PgId}">
       <span class="para">${data.Name}</span>
    </button>`

  }

  return `<div class="accessAccord-child" id="pageunderpgg${data.PgId}">
              <div class="chk-group">
                  ${page_chkbox}
                 <label for="chkpg${data.PgId}"></label>
              </div>
              ${subpagedecider}
            </div>

           <div id="collapseChild${data.PgId}" class="accordion-collapse collapse  accessAccord-child-content"
           aria-labelledby="headingOne" data-bs-parent="#accordionExample">
           <div class="accessAccord-child-row"></div>
           </div>`

}

function GetSubpageForPageGroup(data, spid, pggid, pgid, spaceChecked, spgzChecked) {

  var spg_chkbox

  if (spaceChecked || spgzChecked) {

    spg_chkbox = `<input type="checkbox" id="chkspg${data.SpgId}" data-spid="${spid}" class="spgunderpgg-chkbox SPID${spid}  GRPID${pggid} PGID${pgid}"  data-spgid="${data.SpgId}" data-pggid="${pggid}" data-pgid="${pgid}" checked/>`

  } else {

    spg_chkbox = `<input type="checkbox" id="chkspg${data.SpgId}" data-spid="${spid}" class="spgunderpgg-chkbox SPID${spid}  GRPID${pggid} PGID${pgid}" data-spgid="${data.SpgId}" data-pggid="${pggid}" data-pgid="${pgid}"/>`

  }

  return ` <div class="chk-group chk-group-label">
              ${spg_chkbox}
              <label for="chkspg${data.SpgId}" class="para">${data.Name}</label>
             </div>`

}

$('.spacechkbox').click(function () {

  var space_id = $(this).attr('data-id')

  console.log("spid", space_id);

  if ($(this).is(':checked')) {

    subpages = subpages.filter(obj => obj.spaceId !== space_id)

    pages = pages.filter(obj => obj.spaceId !== space_id)

    pagegroups = pagegroups.filter(obj => obj.spaceId !== space_id)

    spaces.push(space_id)

    $('.SPID' + space_id).prop('checked', true)

  } else {

    for (let x in spaces) {

      if (spaces[x] == space_id) {

        spaces.splice(x, 1)
      }
    }

    subpages = subpages.filter(obj => obj.spaceId !== space_id)

    pages = pages.filter(obj => obj.spaceId !== space_id)

    pagegroups = pagegroups.filter(obj => obj.spaceId !== space_id)

    $('.SPID' + space_id).prop('checked', false)
  }

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);
})

/**page click**/
$(document).on('click', '.dirpage-chkbox', function () {

  var pg_id = $(this).attr('data-pgid');

  var spaceid = $(this).attr('data-spid');

  var page = {}

  page.id = pg_id

  page.groupId = '0'

  page.spaceId = spaceid

  if ($(this).is(':checked')) {

    pages.push(page)

  } else {

    for (let x in pages) {

      if (pages[x].id == pg_id) {

        pages.splice(x, 1)

      }
    }

  }

  var containerObj = HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces)

  subpages = containerObj.subpages

  pages = containerObj.pages

  pagegroups = containerObj.pagegroups

  spaces = containerObj.spaces

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);

})

/**subpage click**/
$(document).on('click', '.subpage-chkbox', function () {

  var spgid = $(this).attr('data-spgid')

  var pg_id = $(this).attr('data-pgid')

  var spaceid = $(this).attr('data-spid');

  var subpage = {}

  subpage.id = spgid

  subpage.groupId = '0'

  subpage.parentId = pg_id

  subpage.spaceId = spaceid

  if ($(this).is(':checked')) {

    subpages.push(subpage)

  } else {

    for (let x in subpages) {

      if (subpages[x].id == spgid) {

        subpages.splice(x, 1)

      }
    }

  }

  var containerObj = HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces);

  subpages = containerObj.subpages

  pages = containerObj.pages

  pagegroups = containerObj.pagegroups

  spaces = containerObj.spaces

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);

})

/**Group Click **/
$(document).on('click', '.pagegroup-chkbox', function () {

  var pgg_id = $(this).attr('data-pggid')

  var spaceid = $(this).attr('data-spid');

  var pagegroup = {}

  pagegroup.id = pgg_id

  pagegroup.spaceId = spaceid

  if ($(this).is(':checked')) {

    $('.GRPID' + pgg_id).each(function () {

      if ($(this).hasClass('pgunderpgg-chkbox')) {

        var pg_id = $(this).attr('data-pgid')

        var new_page = {}

        new_page.id = pg_id

        new_page.groupId = pgg_id

        new_page.spaceId = spaceid

        pages.push(new_page)

      } else {

        var subpg_id = $(this).attr('data-spgid')

        var parent_id = $(this).attr('data-pgid')

        var new_subpage = {}

        new_subpage.id = subpg_id

        new_subpage.parentId = parent_id

        new_subpage.groupId = pgg_id

        new_subpage.spaceId = spaceid

        subpages.push(new_subpage)
      }

      $(this).prop('checked', true)
    })

    pagegroups.push(pagegroup)

  } else {

    for (let x in pagegroups) {

      if (pagegroups[x].id == pgg_id) {

        pagegroups.splice(x, 1)

      }
    }

    pages = pages.filter(obj => { obj.groupId != pgg_id })

    subpages = subpages.filter(obj => { obj.groupId != pgg_id })

    $('.GRPID' + pgg_id).prop('checked', false)

  }

  var containerObj = HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces)

  subpages = containerObj.subpages

  pages = containerObj.pages

  pagegroups = containerObj.pagegroups

  spaces = containerObj.spaces

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);

})

/** page under page group click **/
$(document).on('click', '.pgunderpgg-chkbox', function () {

  var pg_id = $(this).attr('data-pgid')

  var pggid = $(this).attr('data-pggid')

  var spaceid = $(this).attr('data-spid');

  var page = {}

  page.id = pg_id

  page.groupId = pggid

  page.spaceId = spaceid

  if ($(this).is(':checked')) {

    pages.push(page)

    var isAllPagesChecked = true

    $('.GRPID' + pggid).each(function () {

      if (!$(this).is(':checked')) {

        isAllPagesChecked = false

      }

    })

    if (isAllPagesChecked) {

      $('.pagegroup-chkbox[data-pggid=' + pggid + ']').prop('checked', true)
    }

  } else {

    for (let x in pages) {

      if (pages[x].id == pg_id) {

        pages.splice(x, 1)
      }
    }

    for (let x in pagegroups) {

      if (pagegroups[x].id == pggid) {

        pagegroups.splice(x, 1)
      }
    }

    $('.pagegroup-chkbox[data-pggid=' + pggid + ']').prop('checked', false)

  }

  var containerObj = HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces)

  subpages = containerObj.subpages

  pages = containerObj.pages

  pagegroups = containerObj.pagegroups

  spaces = containerObj.spaces

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);

})

/** subpage under page group click **/
$(document).on('click', '.spgunderpgg-chkbox', function () {

  var spgid = $(this).attr('data-spgid')

  var pg_id = $(this).attr('data-pgid')

  var pggid = $(this).attr('data-pggid')

  var spaceid = $(this).attr('data-spid');

  var subpage = {}

  subpage.id = spgid

  subpage.groupId = pggid

  subpage.parentId = pg_id

  subpage.spaceId = spaceid

  if ($(this).is(':checked')) {

    subpages.push(subpage)

    var isAllPagesChecked = true

    $('.GRPID' + pggid).each(function () {

      if (!$(this).is(':checked')) {

        isAllPagesChecked = false

      }

    })

    if (isAllPagesChecked) {

      $('.pagegroup-chkbox[data-pggid=' + pggid + ']').prop('checked', true)
    }

  } else {

    for (let x in subpages) {

      if (subpages[x].id == spgid) {

        subpages.splice(x, 1)
      }
    }

    for (let x in pagegroups) {

      if (pagegroups[x].id == pggid) {

        pagegroups.splice(x, 1)
      }

    }

    $('.pagegroup-chkbox[data-pggid=' + pggid + ']').prop('checked', false)

  }

  var containerObj = HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces)

  subpages = containerObj.subpages

  pages = containerObj.pages

  pagegroups = containerObj.pagegroups

  spaces = containerObj.spaces

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces);

})

function HandleSpaceCheckboxesAndSpaceArray(spaceid, subpages, pages, pagegroups, spaces) {

  var isAllChecked = true

  $('.SPID' + spaceid).each(function () {

    if ($(this).is(':checked')) {

      $('#space' + spaceid).prop('checked', true)

    } else {

      isAllChecked = false

      $('#space' + spaceid).prop('checked', false)

      return false
    }
  })

  if (isAllChecked) {

    subpages = subpages.filter(obj => { obj.spaceId != spaceid })

    pages = pages.filter(obj => { obj.spaceId != spaceid })

    pagegroups = pagegroups.filter(obj => { obj.spaceId != spaceid })

    spaces.push(spaceid)

  } else {

    for (let x in spaces) {

      if (spaces[x] == spaceid) {

        spaces.splice(x, 1)

      }
    }

    var pages_chkd = []

    var pages_unchkd = []

    $('.dirpage-chkbox[data-spid=' + spaceid + ']').each(function () {

      var id = $(this).attr('data-pgid')

      if ($(this).is(':checked')) {

        var isIdExist = pages.some(item => item.id === id);

        var new_page = {}

        new_page.id = id

        new_page.groupId = '0'

        new_page.spaceId = spaceid

        if (!isIdExist) {

          pages.push(new_page)

          pages_chkd.push(id)
        }

      } else {

        pages_unchkd.push(id)

      }
    })


    for (let x of pages_chkd) {

      $('.subpage-chkbox[data-spid=' + spaceid + '][data-pgid=' + x + ']').each(function () {

        var subpage_id = $(this).attr('data-spgid')

        if ($(this).is(':checked')) {

          var isSpgidExist = subpages.some(item => item.id === subpage_id);

          var new_subpage = {}

          new_subpage.id = subpage_id

          new_subpage.groupId = '0'

          new_subpage.parentId = x

          new_subpage.spaceId = spaceid

          if (!isSpgidExist) {

            subpages.push(new_subpage)
          }

        }

      })
    }

    for (let x of pages_unchkd) {

      $('.subpage-chkbox[data-spid=' + spaceid + '][data-pgid=' + x + ']').each(function () {

        var spgid = $(this).attr('data-spgid')

        if ($(this).is(':checked')) {

          var isSpgidExist = subpages.some(item => item.id === spgid);

          var new_subpage = {}

          new_subpage.id = spgid

          new_subpage.groupId = '0'

          new_subpage.parentId = x

          new_subpage.spaceId = spaceid

          if (!isSpgidExist) {

            subpages.push(new_subpage)
          }
        }
      })
    }

    var pgg_chkd = []

    $('.pagegroup-chkbox[data-spid=' + spaceid + ']').each(function () {

      var id = $(this).attr('data-pggid')

      if ($(this).is(':checked')) {

        var isPggIdExist = pagegroups.some(item => item.id === id);

        var new_pagegroup = {}

        new_pagegroup.id = id

        new_pagegroup.spaceId = spaceid

        if (!isPggIdExist) {

          pagegroups.push(new_pagegroup)

          pgg_chkd.push(id)
        }

      } else {

        $('.pgunderpgg-chkbox[data-spid=' + spaceid + '][data-pggid=' + id + ']').each(function () {

          var pgInPgg_id = $(this).attr('data-pgid')

          if ($(this).is(':checked')) {

            var isPgIdExist = pages.some(item => item.id === pgInPgg_id);

            var new_page = {}

            new_page.id = pgInPgg_id

            new_page.groupId = id

            new_page.spaceId = spaceid

            if (!isPgIdExist) {

              pages.push(new_page)

            }

          }

        })
      }

    })

    for (let x of pgg_chkd) {

      var isAllPagesChecked = true

      var pginpgg_chkd = []

      var pginpgg_unchkd = []

      $('.pgunderpgg-chkbox[data-spid=' + spaceid + '][data-pggid=' + x + ']').each(function () {

        var pgInPgg_id = $(this).attr('data-pgid')

        if ($(this).is(':checked')) {

          var isPgIdExist = pages.some(item => item.id === pgInPgg_id);

          var new_page = {}

          new_page.id = pgInPgg_id

          new_page.groupId = x

          new_page.spaceId = spaceid

          if (!isPgIdExist) {

            pages.push(new_page)

            pginpgg_chkd.push(pgInPgg_id)

          }

        } else {

          isAllPagesChecked = false

          pginpgg_unchkd.push(pgInPgg_id)

        }
      })

      for (let y of pginpgg_chkd) {

        $('.spgunderpgg-chkbox[data-pggid=' + x + '][data-pgid=' + y + ']').each(function () {

          var spgInPgg_id = $(this).attr('data-spgid')

          if ($(this).is(':checked')) {

            var isSpgidExist = subpages.some(item => item.id === spgInPgg_id);

            var new_subpage = {}

            new_subpage.id = spgInPgg_id

            new_subpage.groupId = x

            new_subpage.parentId = y

            new_subpage.spaceId = spaceid

            if (!isSpgidExist) {

              subpages.push(new_subpage)
            }

          }

        })

      }

      for (let z of pginpgg_unchkd) {

        $('.spgunderpgg-chkbox[data-pggid=' + x + '][data-pgid=' + z + ']').each(function () {

          var spgInPgg_id = $(this).attr('data-spgid')

          if ($(this).is(':checked')) {

            var isSpgidExist = subpages.some(item => item.id === spgInPgg_id);

            var new_subpage = {}

            new_subpage.id = spgInPgg_id

            new_subpage.groupId = x

            new_subpage.parentId = z

            new_subpage.spaceId = spaceid

            if (!isSpgidExist) {

              subpages.push(new_subpage)
            }

          }

        })

      }

    }

    $('.subpage-chkbox[data-spid=' + spaceid + ']').each(function () {

      var spgid = $(this).attr('data-spgid')

      var parent_id = $(this).attr('data-pgid')

      if ($(this).is(':checked')) {

        var isSpgidExist = subpages.some(item => item.id === spgid);

        var new_subpage = {}

        new_subpage.id = spgid

        new_subpage.groupId = '0'

        new_subpage.parentId = parent_id

        new_subpage.spaceId = spaceid

        if (!isSpgidExist) {

          subpages.push(new_subpage)
        }

      }

    })

    $('.spgunderpgg-chkbox[data-spid=' + spaceid + ']').each(function () {

      var spgid = $(this).attr('data-spgid')

      var parent_id = $(this).attr('data-pgid')

      var pgg_id = $(this).attr('data-pggid')

      if ($(this).is(':checked')) {

        var isSpgidExist = subpages.some(item => item.id === spgid);

        var new_subpage = {}

        new_subpage.id = spgid

        new_subpage.groupId = pgg_id

        new_subpage.parentId = parent_id

        new_subpage.spaceId = spaceid

        if (!isSpgidExist) {

          subpages.push(new_subpage)
        }

      }

    })

  }

  var container = { "subpages": subpages, "pages": pages, "pagegroups": pagegroups, "spaces": spaces }

  console.log("func chk", container);

  return container

}

$('#memgrp-slctall').click(function () {

  if ($(this).prop("checked")) {

    $('.memgrp-chkboxes').each(function (chk_index, chk_element) {

      var memgrp_id = $(chk_element).attr('data-id')

      if ($(chk_element).prop('checked') == false) {

        access_granted_memgrps.push(memgrp_id)

        $(chk_element).prop('checked', true)

      }
    })

  } else {

    // access_granted_memgrps.splice(0,access_granted_memgrp

    access_granted_memgrps.length = 0

    $('.memgrp-chkboxes').prop("checked", false)

  }

  console.log("memgrps", access_granted_memgrps);
})



$('.memgrp-chkboxes').each(function (chk_index, chk_element) {

  $(chk_element).click(function () {

    console.log('checked', $(this).prop('checked'));

    var memgrp_id = $(this).attr('data-id')

    if ($(this).prop('checked')) {

      var allChecked = true

      $('.memgrp-chkboxes').each(function (cb_index, cb_element) {

        if ($(cb_element).prop('checked') == false) {

          allChecked = false

          return false

        }
      })


      if (!$('#memgrp-slctall').prop('checked')) {

        if (allChecked) {

          access_granted_memgrps.push(memgrp_id)

          $(this).prop('checked', true)

          $('#memgrp-slctall').prop('checked', true)

        } else {

          access_granted_memgrps.push(memgrp_id)

          $(this).prop('checked', true)

        }

      } else {

        if (allChecked) {

          access_granted_memgrps.push(memgrp_id)

          $(this).prop('checked', true)

          $('#memgrp-slctall').prop('checked', true)

        } else {

          access_granted_memgrps.push(memgrp_id)

          $(this).prop('checked', true)

        }

      }

    } else {

      if ($('#memgrp-slctall').prop('checked')) {

        for (let x in access_granted_memgrps) {

          if (access_granted_memgrps[x] == memgrp_id) {

            access_granted_memgrps.splice(x, 1)
          }
        }

        $(this).prop('checked', false)

        $('#memgrp-slctall').prop('checked', false)

      } else {

        for (let x in access_granted_memgrps) {

          if (access_granted_memgrps[x] == memgrp_id) {

            access_granted_memgrps.splice(x, 1)

          }
        }

        $(this).prop('checked', false)

      }
    }

    console.log("memgrps", access_granted_memgrps);
  })
})

var pageno

$('.configurationContent-btm > button').click(function () {

  var url = window.location.search

  const urlpar = new URLSearchParams(url);


  pageno = urlpar.get('page');

  // pages = pages.filter(obj => !spaces.includes(obj.spaceId))

  // subpages = subpages.filter(obj => !spaces.includes(obj.spaceId))

  // pagegroups = pagegroups.filter(obj => !spaces.includes(obj.spaceId))

  console.log("subpages", subpages, "pages", pages, "pagegroups", pagegroups, "spaces", spaces, "memgrps", access_granted_memgrps, "channels", channels, "channelEntries", channelEntries);

  if ($('#ca-inpt').val() != "") {
    console.log("titile:");

    if ($('#ca-inpt').val() != "" && $('input[type=hidden][id=csrf-input]').val() != "") {

      var data = {}

      data.title = $('#ca-inpt').val()

      // data.spaces = JSON.stringify({ spaces })

      // data.pages = JSON.stringify({ pages })

      // data.subpages = JSON.stringify({ subpages })

      // data.pagegroups = JSON.stringify({ pagegroups })

      data.channels = JSON.stringify({ channels })

      data.entries = JSON.stringify({ channelEntries })

      data.memgrps = JSON.stringify({ access_granted_memgrps })

      data.csrf = $('input[type=hidden][id=csrf-input]').val()

      if (content_access_id != "") {
        console.log(content_access_id, "mem access");

        data.content_access_id = content_access_id

      }

      // var data = {
      //   "title": title,
      //   "spaces":JSON.stringify({spaces}),
      //   "pages":JSON.stringify({pages}),
      //   "subpages":JSON.stringify({subpages}),
      //   "pagegroups":JSON.stringify({pagegroups}),
      //   "memgrps":JSON.stringify({access_granted_memgrps}),
      //   "csrf":csrf
      // }

      if ((spaces.length > 0 || pages.length > 0 || subpages.length > 0 || pagegroups.length > 0 || channelEntries.length > 0 || channels.length > 0) && access_granted_memgrps.length > 0) {
        var callUrl

        if ($('#route-type').val() == "Save") {

          callUrl = '/memberaccess/grant-accesscontrol'

        } else if ($('#route-type').val() == "Update") {

          callUrl = '/memberaccess/update-accesscontrol'

        } else if ($('#route-type').val() == "Copy") {

          callUrl = '/memberaccess/grant-accesscontrol'

        }

        PassContentAccessGrantedIds(data, callUrl, pageno)

      } else {

        var message = ''

        if ($('#default-mod-chk[type=hidden]').val() == "space" && spaces.length == 0 && pages.length == 0 && subpages.length == 0 && pagegroups.length == 0 && channels.length == 0 && channelEntries.length == 0) {

          message = languagedata.ContentAccessControl.pagechannelgrantaccess

        } else if ($('#default-mod-chk[type=hidden]').val() == "channel" && channels.length == 0 && channelEntries.length == 0) {

          message = languagedata.ContentAccessControl.channelgrantaccess

        } else if (access_granted_memgrps.length == 0 && $('#default-mod-chk[type=hidden]').val() == "space") {

          message = languagedata.ContentAccessControl.spacedefaultmembergroupaccess

        } else if (access_granted_memgrps.length == 0 && $('#default-mod-chk[type=hidden]').val() == "channel") {

          message = languagedata.ContentAccessControl.channeldefaultmembergroupaccess

        }


        notify_content = `<ul class=" warn-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="/ flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/toast-error.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + "please select entry and member group properly" + `</p></div></div> </li></ul>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
          $('.warn-msg').fadeOut('slow', function () {
            $(this).remove();
          });
        }, 5000);

      }
    }

  } else {

    $('#ca-inpt-error').show()
  }

})


$('#ca-inpt').keyup(function () {

  if ($(this).val() != "") {

    $('#ca-inpt-error').hide()

  } else {

    $('#ca-inpt-error').show()
  }

})

function PassContentAccessGrantedIds(data, url, pageno) {
  console.log("hir", url);

  $.ajax({

    url: url,

    type: 'POST',

    data: data,

    dataType: 'json',

    success: function (data) {

      console.log("datass", data)

      if (data) {

        console.log("data", data);

        if (pageno == null) {

          window.location.href = "/memberaccess/"

        } else {
          window.location.href = "/memberaccess/?page=" + pageno

        }

      } else {

        var message = languagedata.Toast.internalserverr

        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

        $(notify_content).insertBefore(".header-rht");

        setTimeout(function () {

          $('.toast-msg').fadeOut('slow', function () {

            $(this).remove();

          });

        }, 5000);
      }

    }

  })
}


// $('.ava-sel-srch>input').keyup(function () {

//   var keyword = $(this).val().trim().toLowerCase()

//   if ($(this).hasClass("space-srch-input") && $('.space-row').length > 0) {

//     PerformSpaceSearch(keyword)

//   } else if ($(this).hasClass("channel-srch-input") && $('.channel-row').length > 0) {

//     PerformChannelSearch(keyword)

//   }

// })


$('.channel-srch-input').keyup(function () {
  var keyword = $(this).val().trim().toLowerCase()

  if (keyword !== "") {

    $('.closesearch-btn').show()
    $('.srchBtn-togg').addClass('pointer-events-none')

}else{

    $('.closesearch-btn').hide()
    $('.srchBtn-togg').removeClass('pointer-events-none')

}
  if ($('.channel-row').length > 0) {
    PerformChannelSearch(keyword)
  }


})



$('.ava-sel-srch>img').click(function () {

  var keyword = $(this).siblings('input').val().trim().toLowerCase()

  if ($(this).hasClass("space-srch-img") && $('.space-row').length > 0) {

    PerformSpaceSearch(keyword)

  } else if ($(this).hasClass("channel-srch-img") && $('.channel-row').length > 0) {

    PerformChannelSearch(keyword)

  }

})

// function PerformSpaceSearch(keyword) {

//   if (keyword !== "") {

//     $('.space-row').each(function () {

//       var regex = new RegExp(keyword, "i"); // "i" makes the search case-insensitive

//       var spacename = $(this).find('.accessAccord-parent>.chk-group>label').text()

//       if (regex.test(spacename)) {

//         $(this).show()

//         $('.noData-foundWrapper').remove()

//       } else {

//         $(this).hide()

//         if ($('.space-row:visible').length == 0) {

//           var nodata_html = `<div class="noData-foundWrapper">
//                                     <div class="empty-folder">
//                                         <img style="max-width: 35px;" src="/public/img/nodatafilter.svg" alt="">

//                                     </div>
//                                    <h1 style="text-align: center;font-size: 14px;" class="heading"> `+ languagedata.nodatafound + ` </h1>
//                                 </div>`

//           if ($('.noData-foundWrapper').length == 0) {

//             $(this).parent().append(nodata_html)
//           }

//         } else {

//           $('.noData-foundWrapper').remove()
//         }

//       }
//     })

//   } else {

//     $('.noData-foundWrapper').remove()

//     $('.space-row').show()

//   }

// }

function PerformChannelSearch(keyword) {

  if (keyword !== "") {

    $('.channel-row').each(function () {

      var regex = new RegExp(keyword, "i"); // "i" makes the search case-insensitive
      console.log(regex, "regx");

      var channelname = $(this).find('.accessAccord-parent p').text();

      if (regex.test(channelname)) {

        $(this).show()

        $('.noData-foundWrapper').remove()

      } else {

        $(this).hide()

        if ($('.channel-row:visible').length == 0) {

          var nodata_html = `<div class="noData-foundWrapper m-auto" style="margin:auto;">
                                  <div class="empty-folder" >
                                      <img style="max-width: 100px; margin:auto" src="/public/img/nodatafilter.svg" alt="">

                                  </div>
                                 <h1 style="text-align: center;font-size: 14px;" class="heading"> `+ languagedata.nodatafound + ` </h1>
                              </div>`

          if ($('.noData-foundWrapper').length == 0) {

            $(this).parent().append(nodata_html)
          }

        } else {

          $('.noData-foundWrapper').remove()
        }

      }
    })

  } else {

    $('.noData-foundWrapper').remove()

    $('.channel-row').show()

  }

}

$('.delete-access').click(function () {

  var url = window.location.search;

  const urlpar = new URLSearchParams(url);

  var pageno = urlpar.get('page');

  var accessId = $(this).attr("data-id")

  console.log("langid", accessId);

  $('.deltitle').text(languagedata.ContentAccessControl.deltitle)

  $('.deldesc').text(languagedata.ContentAccessControl.deldesc)

  $('.delname').text($(this).parents('tr').find('td:first>.title-cell>h3').text())

  if (pageno == null) {

    $('#delid').attr('href', "/memberaccess/delete-accesscontrol/" + accessId)

  } else {

    $('#delid').attr('href', "/memberaccess/delete-accesscontrol/" + accessId + "?page=" + pageno)

  }

  $('#deleteModal').modal('show')

})

$(document).on('click', '#channel-access', function () {

  $("#description").text(languagedata.ContentAccessControl.channelaccessdesc)

  console.log("channel-access clicked");

  if ($('.channel-row').length == 0) {

    if ($('#default-mod-chk[type=hidden]').val() != "channel") {

      $.ajax({

        url: '/memberaccess/get-channels',

        dataType: 'json',

        success: function (data) {

          console.log("channel data", data);

          if (data.Channels != null) {

            $('.noData-foundWrapper').hide()

            for (let channel of data.Channels) {

              var channelEntry_html = ""

              for (let channelEntry of channel.ChannelEntries) {

                var isEntryChkd = false

                for (let readyEntry of channelEntries) {

                  if (readyEntry.id == channelEntry.Id) {

                    isEntryChkd = true

                  }

                }

                var entry_chkbox = ""

                if (isEntryChkd) {

                  entry_chkbox = ` <input type="checkbox" id="ChannelEntry${channelEntry.Id}" class="chanEntry-chkbox" data-id="${channelEntry.Id}" data-chanid="${channel.Id}" checked>`

                } else {

                  entry_chkbox = ` <input type="checkbox" id="ChannelEntry${channelEntry.Id}" class="chanEntry-chkbox" data-id="${channelEntry.Id}" data-chanid="${channel.Id}">`

                }

                channelEntry_html += `<div class="accessAccord-child">
                                                 <div class="chk-group">
                                                 ${entry_chkbox}
                                                    <label for="ChannelEntry${channelEntry.Id}"></label>
                                                  </div>

                                                  <button class="accord-collapse">
                                                     <span class="para">${channelEntry.Title}</span>
                                                   </button>
                                                   </div>`

              }

              var channelEntryDiv = ""

              if (channelEntry_html != "") {

                channelEntryDiv = `<div id="channel${channel.Id}" class="accordion-collapse collapse  accessAccord-parent-content" data-toggle="${channel.Id}" aria-labelledby="headingOne" data-bs-parent="#accordionExample">
                                 <div class="accessAccord-child-container channel-childrow">${channelEntry_html}</div>
                                 <div>`

              } else {

                channelEntryDiv = ''

              }

              var channel_chkbox = ""

              if (channels.length > 0) {

                for (let channelid of channels) {

                  if (channelid == channel.Id) {

                    channel_chkbox = `<input type="checkbox" id="channel${channel.Id}" class="channelchkbox" data-id="${channel.Id}" checked>`

                  } else {

                    channel_chkbox = `<input type="checkbox" id="channel${channel.Id}" class="channelchkbox" data-id="${channel.Id}">`

                  }

                }

              } else {

                channel_chkbox = `<input type="checkbox" id="channel${channel.Id}" class="channelchkbox" data-id="${channel.Id}">`

              }

              var channelBindHtml = `<div class="accessAccord-row channel-row">

              <div class="accessAccord-parent">

                    <div class="chk-group chk-group-label">
                          ${channel_chkbox}
                          <label for="channel${channel.Id}" class="para">${channel.ChannelName}</label>
                    </div>

                    <div class="accessAccord-parent-end">

                       <div class="card-breadCrumb">
                           <ul class="breadcrumb-container"></ul>
                       </div>

                       <a href="javascript:void(0)" class="img-div accord-collapse" data-bs-toggle="collapse" data-bs-target="#channel${channel.Id}" aria-expanded="false" aria-controls="channel${channel.Id}">
                        <img src="/public/img/arrow-rgt.svg" alt="" data-toggle="${channel.Id}">
                      </a>

                   </div>
              </div>
              ${channelEntryDiv}
            </div>`

              $('.configurationContent-section-lft>.accessAccord-container').append(channelBindHtml)

            }

          } else {

            if ($('.channel-row').length == 0) {

              var nodata_html = `<div class="noData-foundWrapper">
                                  <div class="empty-folder">
                                       <img style="max-width: 35px;" src="/public/img/folder-sh.svg" alt="">
                                       <img src="/public/img/shadow.svg" alt="">
                                  </div>
                                  <h1 style="text-align: center;font-size: 14px;" class="heading"> `+ languagedata.nodatafound + ` </h1>
                                </div>`

              if ($('.noData-foundWrapper').length == 0) {

                $('accessAccord-container').append(nodata_html)

              } else {

                $('.noData-foundWrapper').show()
              }
            }

          }

        }

      })

    }

  } else if ($('.channel-row').length > 0) {

    $('.noData-foundWrapper').hide()

  }

  $('.space-childrow').each(function () {

    if ($(this).parent().css('display') == 'block') {

      $(this).parent().hide()

      $(this).parents('.accessAccord-row').find('.space-row>.accessAccord-parent-end>a').attr('aria-expanded', 'false')

    }

  })

  $(this).addClass('active')

  $('#space-access').removeClass('active')

  $('.accessHead-rgt-start>div>h3').text('Channels')

  $('.ava-sel-srch>img').removeClass('space-srch-img').addClass('channel-srch-img')

  $('.ava-sel-srch>input').attr('placeholder', 'Search Channels').removeClass('space-srch-input').addClass('channel-srch-input').val("")

  $('.configurationContent-section-lft>.accessAccord-container .space-row').hide()

  $('.configurationContent-section-lft>.accessAccord-container .channel-row').show()

})

$(document).on('click', '#space-access', function () {

  console.log("space-access clicked");

  $("#description").text(languagedata.ContentAccessControl.spacedesc)

  $(this).addClass('active')

  $('#channel-access').removeClass('active')

  $('.accessHead-rgt-start>div>h3').text('Spaces')

  $('.ava-sel-srch>img').removeClass('channel-srch-img').addClass('space-srch-img')

  $('.ava-sel-srch>input').attr('placeholder', 'Search Spaces').removeClass('channel-srch-input').addClass('space-srch-input').val("")

  $('.configurationContent-section-lft>.accessAccord-container .channel-row').hide()

  $('.configurationContent-section-lft>.accessAccord-container .space-row').show()

  if ($('.space-row').length > 0) {

    $('.noData-foundWrapper').hide()

    $('.channel-childrow').each(function () {

      if ($(this).parent().css('display') == 'block') {

        $(this).parent().removeClass('show')

        $(this).parents('.channel-row').find('.accessAccord-parent-end>a').addClass('collapsed').attr('aria-expanded', 'false')

      }

    })

  } else if ($('.space-row').length == 0) {

    var nodata_html = `<div class="noData-foundWrapper">
                        <div class="empty-folder">
                             <img style="max-width: 35px;" src="/public/img/folder-sh.svg" alt="">
                             <img src="/public/img/shadow.svg" alt="">
                        </div>
                        <h1 style="text-align: center;font-size: 14px;" class="heading"> `+ languagedata.nodatafound + ` </h1>
                      </div>`

    if ($('.noData-foundWrapper').length == 0) {

      $('accessAccord-container').append(nodata_html)

    } else {

      $('.noData-foundWrapper').show()
    }
  }

})


$(document).on('click', '.channelchkbox', function () {
  console.log("hKHF");

  var channelId = $(this).attr('data-id')

  if ($(this).is(':checked')) {

    channels.push(channelId)

    $(this).prop('checked', true)

    $('.chanEntry-chkbox[data-chanid=' + channelId + ']').each(function () {

      var entryId = $(this).attr('data-id')

      var entry = {}

      entry.id = entryId

      entry.channelId = channelId

      channelEntries.push(entry)

      $(this).prop('checked', true)

    })

  } else {

    for (let x in channels) {

      if (channels[x] == channelId) {

        channels.splice(x, 1)


      }

    }

    // channelEntries = channelEntries.filter(obj => { obj.channelId != channelId })

    channelEntries = channelEntries.filter(function(obj) {
      return obj.channelId !== channelId; // Keep only objects whose channelId is not equal to the unchecked one
  });

    $('.chanEntry-chkbox[data-chanid=' + channelId + ']').prop('checked', false)

    $(this).prop('checked', false)

  }

  console.log("chk channel", channels, channelEntries);

})




$(document).on('click', '.chanEntry-chkbox', function () {

  var channelId = $(this).attr('data-chanid')

  var entryId = $(this).attr('data-id')

  if ($(this).is(':checked')) {

    var entry = {}

    entry.id = entryId

    entry.channelId = channelId

    channelEntries.push(entry)

    $(this).prop('checked', true)

  } else {

    for (let x in channelEntries) {

      if (channelEntries[x].id == entryId) {

        channelEntries.splice(x, 1)
      }

    }

    $(this).prop('checked', false)
  }

  var isAllEntrieschkd

  $('.chanEntry-chkbox[data-chanid=' + channelId + ']').each(function () {

    if ($(this).is(':checked')) {

      isAllEntrieschkd = true

    } else {

      isAllEntrieschkd = false

      return false
    }

  })

  if (isAllEntrieschkd) {

    channels.push(channelId)

    $('.channelchkbox[data-id=' + channelId + ']').prop('checked', true)

  } else {

    for (let x in channels) {

      if (channels[x] == channelId) {

        channels.splice(x, 1)
      }

    }

    $('.channelchkbox[data-id=' + channelId + ']').prop('checked', false)

  }

  console.log("chk channel", channels, channelEntries);

})


// Search return to home page

$(document).on('keyup', '#memberrestrictsearch', function () {

  if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

    if ($('.search').val() === "") {
      window.location.href = "/memberaccess"

    }
  }

})

$(document).on("click", ".Closebtn", function () {
  $(".search").val('')
  $(".Closebtn").addClass("hidden")
  $(".srchBtn-togg").removeClass("pointer-events-none")

})

$(document).on("click", ".searchClosebtn", function () {
  $(".search").val('')
  window.location.href = "/memberaccess/"
})

$(document).ready(function () {

  $('.search').on('input', function () {

      if ($(this).val().length >= 1) {
          $(".Closebtn").removeClass("hidden")
          $(".srchBtn-togg").addClass("pointer-events-none")

      }else{
        $(".Closebtn").addClass("hidden")
        $(".srchBtn-togg").removeClass("pointer-events-none")

      }
  });
})

$(document).on("click", ".hovericon", function () {
  $(".search").val('')
  $(".Closebtn").addClass("hidden")
})

$(document).on('keyup', '.checklength', function () {
  var inputVal = $(this).val()

  // console.log(inputVal.length);

  var inputLength = inputVal.length

  if (inputLength == 25) {
    $(this).siblings('.lengthErr').removeClass('hidden')
  } else {
    $(this).siblings('.lengthErr').addClass('hidden')
  }

})


$(document).on('click','.closesearch-btn',function(){
    $('#restrict-searchinput').val('')
    $(this).hide()
    $('.noData-foundWrapper').remove()

    $('.channel-row').show()
})