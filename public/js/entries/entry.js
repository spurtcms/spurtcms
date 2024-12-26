var languagedata
let ckeditor1;
var categoryIds = [];
var seodetails = {};
var channeldata = []
var checkboxarr = []
var membercheckarr = []
var cateogryarr = []
selctcategory = []
var Channelid
var flagg = false
var result1 = []
var newchanneldata = []
var membername
var authername
var createtime
var publishtime
var readingtime
var sortorder
var tagname
var extxt
var selectedcheckboxarr = []
var videourlarr
var urlvalues
var youtubelinks = []
var finalyoutubelinks = []
var joindata
var newOrder = [];
var access_granted_memgrps = []
var pentryid


$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $("#searchkey").focus().select();
    }
});
// Return to entries list function //
$(".searchentry").click(function () {

    id = $(this).attr('id')

    if ($(this).is(':checked')) {

        window.location.href = "/channel/entrylist/" + id
    } else {
        window.location.href = "/channel/entrylist/"
    }

})

$(document).on('click', '#removecategory', function () {

    var currentLocation = window.location.href;


    if (currentLocation.includes("unpublishentrieslist")) {

        window.location.href = "/channel/unpublishentries"

    } else if (currentLocation.includes("draftentrieslist")) {

        window.location.href = "/channel/draftentries"

    } else if (currentLocation.includes("entrylist")) {

        window.location.href = "/channel/entrylist"
    }


})

$(document).on("keyup", "#article", function () {


    $("#sort-error").hide()

})

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    // You get innerHTML here    
    document.addEventListener('change', function (event) {
        console.log("artsdhj", event.detail);
    });


    $(document).on("click", "#save-entry", function () {

        //use this in the submit button 
        const event = new CustomEvent("getHTML", {
        });
        document.dispatchEvent(event);
    })

    if ($('#statusid').val() == 'Draft') {
        $('#slctstatus').text('Draft');
    }
    if ($('#statusid').val() == 'Published') {
        $('#slctstatus').text('Published');
    }
    if ($('#statusid').val() == 'Unpublished') {
        $('#slctstatus').text('Unpublished');
    }


    var url = window.location.href;


    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    // if (url.includes('/editentry')) {
    var urlvalue = url.split('?')[0]

    var id = urlvalue.split('/').pop()

    var array = url.split('/');

    var channelname = array[array.length - 2];

    // }


    // var id = array[array.length - 1];
    var editurl

    if (channelname != "") {
        if (pageno == null) {
            editurl = "/channel/editentrydetails/" + channelname + "/" + id
        } else {
            editurl = "/channel/editentrydetails/" + channelname + "/" + id + "?page=" + pageno
        }

    } else {

        if (pageno == null) {
            editurl = "/channel/editentrydetail/" + id
        } else {
            editurl = "/channel/editentrydetail/" + id + "?page=" + pageno
        }
    }




    if (/^\d+$/.test(id)) {
        $.ajax({
            url: editurl,
            type: "GET",
            dataType: "json",
            data: { "id": id, csrf: $("input[name='csrf']").val() },
            success: function (result) {

                result1 = result


                Channelid = result.Entries.ChannelId

                if (result.Entries.CategoriesId != "") {

                    categoryIds = result.Entries.CategoriesId.split(",")
                }

                if (result.Entries.CoverImage == "") {
                    $('#cenimg').show()

                }
                else {
                    $(".imgSelect-image").hide()

                    $(".uplclass").hide()

                    $(".imgSelect-btn").hide()
                    // $("#spimagehide").css("height", "23.120rem")
                    // $("#spimagehide").css("width", "46.25rem")
                    $(".blogBottom-imgSelect").css("padding-block", "0rem")
                    $("#spacedel-img").show();
                    $('#spimagehide').attr('src', result.Entries.CoverImage)
                }
                var currentDate = new Date();
                var formattedDate = new Date(currentDate.getTime() - currentDate.getTimezoneOffset() * 60000).toISOString().slice(0, 16);
                $('#current-time-input').val(formattedDate);

                $('#titleid').val(result.Entries.Title)
                // $('#text').val(result.Description)
                $('#titleid').css('height', '100%')
                ckeditor1.setData(result.Entries.Description)
                $('#chanid').val(result.Entries.ChannelId)
                $('#metatitle').val(result.Entries.MetaTitle)
                $('#metadesc').val(result.Entries.MetaDescription)
                $('#metakey').val(result.Entries.Keyword)
                $('#metaslug').val(result.Entries.Slug)
                $("#imgtag").val(result.Entries.ImageAltTag)
                $('#seosavebtn').text(languagedata.update)
                $('#categorysave').text(languagedata.update)
                if (result.Entries.Author != "") {

                    $("#author").val(result.Entries.Author)
                    $("#cn-user").show()
                }
                if (result.Createtime != "") {

                    $("#cdtime").val(result.Createtime)

                } else {

                    $("#cdtime").val(formattedDate)

                }

                if (result.Publishedtime != "") {

                    $("#publishtime").val(result.Publishedtime)


                } else {

                    $("#publishtime").val(formattedDate)

                }

                if (result.Entries.ReadingTime != 0) {

                    $("#readingtime").val(result.Entries.ReadingTime)

                }
                if (result.Entries.SortOrder != 0) {

                    $("#article").val(result.Entries.SortOrder)
                }
                $("#tagname").val(result.Entries.Tags)
                $("#extxt").val(result.Entries.Excerpt)
                membername = result.Memberlist
                // $('#configbtn').attr('href', '/channels/editchannel/' + result.Entries.ChannelId)

                seodetails = { title: result.Entries.MetaTitle, desc: result.Entries.MetaDescription, keyword: result.Entries.Keyword, slug: result.Entries.Slug, imgtag: result.Entries.ImageAltTag }

                //dropdown item click//
                var dropdownItems = document.getElementsByClassName('avaliable-dropdown-item');
                for (var i = 0; i < dropdownItems.length; i++) {
                    if (dropdownItems[i].id == result.Entries.ChannelId) {

                        var channelname = dropdownItems[i].querySelector('.para').textContent;
                        var count = dropdownItems[i].querySelector('.para-extralight').textContent;
                        var itemId = dropdownItems[i].getAttribute('id');

                        $('#chanid').val(itemId)
                        $('#blogn').text(channelname).css('color', 'black')
                        $('#blogn').attr('data-char', channelname)
                        $('.blogDropdown1').attr('data-bs-original-title', channelname)
                        $('#blogcount').text(count)
                        // $('.blogDropdown').append('<img class="blogimg" src="/public/img/blog-drop.svg" alt="">');
                    }
                }
                if (result.CategoryName !== null) {

                    var div = ""

                    for (let y of result.CategoryName) {

                        var categoriess = ""

                        var ids = ""

                        var id = ""

                        div +=
                            `    <div class="categories-list-child categorypdiv" id="category-">
                        <div class="choose-cat-list-col" style="display: flex;" data-id=>`

                        for (let x of y.Categories) {

                            div += `<h3 class="para categoryname" data-id="` + x.Id + `" style="font-weight: 400;">  ` + x.Category + `</h3>
                              
                                <h3 class="para">&nbsp;/&nbsp;</h3>`

                            categoriess += x.Category + "~"

                            ids += x.Id

                            id = x.Id
                        }

                        div += `</div>
                        <a href="javascript:void(0)" class="category-select-btn" data-id="`+ id + `" data-categoryid="` + ids + `">
                            <img src="/public/img/add-category.svg" alt="" />
                        </a>
                        <p class="forsearch" style="display: none;">`+ categoriess + `</p>
                    </div>
                    <div class="noData-foundWrapper" id="categorynodatafound" style="display: none;">
    
                    <div class="empty-folder">
                        <img src="/public/img/folder-sh.svg" alt="">
                        <img src="/public/img/shadow.svg" alt="">
                    </div>
                    <h1 class="heading">Oops No Data Found</h1>
                    </div>  
                      `
                    }
                    $('.slist').html(div);
                }



                if (result.SelectedCategory !== null) {
                    result.SelectedCategory.forEach(function (category) {


                        newcategory = category.split(',').join('');

                        cateogryarr.push(newcategory)

                    })
                }

                var count = 0

                $('.categorypdiv').each(function () {

                    if ($(this).css('display') != "none") {

                        count = count + 1

                    }
                })

                $('#avacatcount').text(count);

                //channel field section//
                $('#channelfieldsave').text(languagedata.update)

                if (result.Section != null) {
                    result.Section.forEach(function (section) {

                        $('.secid').append(`<div style="margin-bottom:10px" id="section${section.SectionId}"><h6 style="margin-bottom:10px" class="sechead">${section.SectionName}</h6></div>`)

                    })
                    if (result.FieldValue != null) {
                        result.FieldValue.forEach(function (fieldValue) {

                            if (fieldValue.MasterFieldId == 15) {

                                $(`#section${fieldValue.SectionId}`).append(` 
                                     <div class="getvalue">
                                <h3 class="input-label para" data-id="${fieldValue.FieldId}">${fieldValue.FieldName}</h3>
                                 <div class="uploadFolders add-galleryfolder"  id="addnewimageentriesModalmedia">
                <img src="/public/img/uploadFolder.svg" alt="Upload Folder" id="fl-img">
                <p class="heading-three fl-path${fieldValue.FieldId}" id="fl-path">Upload Folder</p>
                <a href="javascript:void(0);" id="fl-remove" style="display:none;"><img src="/public/img/remove-folder.svg" alt="remove" id="remove-ga-folder"></a>
                <span class="heading-three"><input type="hidden" id="${fieldValue.FieldId}" data-id="" class="fl-pathfld fl-path${fieldValue.FieldId}"></span>
                  </div>
                  </div>
                                `)
                            }
                            if (fieldValue.MasterFieldId == 16) {

                                $(`#section${fieldValue.SectionId}`).append(` 
                         <div class="input-group getvalue">
                <label class="input-label" data-id="${fieldValue.FieldId}">${fieldValue.FieldName}</label>
                <div class="ig-row pre-imgCont date-input">
                  <input type="text" id="youtubeurl" placeholder="Paste your link here" />
                                    <input type="hidden" name="urldata"  data-id="" class="${fieldValue.FieldId}"  id="urldatas" value="` + urlvalues + `">
                  <span class="pre-img">
                    <img src="/public/img/pasteLink.svg" alt="paste link">
                  </span>
                  <a href="javascript:void(0);" class="para" id="videourl">Add</a>
                  
                </div>
               <ul class="vdo-list video-url-list">

                </ul>
                  </div>
                                `)
                            }

                            if (fieldValue.MasterFieldId == 14 && fieldValue.Mandatory == 1) {

                                $(`#section${fieldValue.SectionId}`).append(` 

                                <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id=""  style="margin-bottom: 20px">
         
                                         <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
         
                                         <a class="dropdown-toggle " type="button" id="triggerId" data-bs-toggle="dropdown"
                                             aria-haspopup="true" aria-expanded="false">
                                         Select Members
                                         </a>
                                         <input type="hidden" data-id="" class="checkfieldid1" id="${fieldValue.FieldId}"  value=""  >
                                         <div class="dropdown-menu additional-entry-drop" id="check${fieldValue.FieldId}" aria-labelledby="triggerId">
                                         <div class="ig-row ig-channel-input">

                                         <input type="text" id="searchdropdownrole" class="search" name="keyword"
                                             placeholder="Search Members" value="">
                                     </div>
                                     <div class="noData-foundWrapper" id="nodatafounddesign"
                                     style="margin-top: -40px;display: none;">
 
                                     <div class="empty-folder">
                                         <img style="max-width: 20px;" src="/public/img/folder-sh.svg" alt="">
                                         <img src="/public/img/shadow.svg" alt="">
                                     </div>
                                     <h1 style="text-align: center;font-size: 10px;" class="heading">
                                    ` + languagedata.oopsnodata + `</h1>
 
                                 </div>
                                             
                                         </div>
                                         <label class="error manerr" id="opterrr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
         
                                     </div>
          
                `);
                            }
                            if (fieldValue.MasterFieldId == 14) {

                                $(`#section${fieldValue.SectionId}`).append(` 
                                <div class="input-group getvalue" data-id="" data-char="">
                                <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                                <input type="hidden" data-id="" class="fieldid1" id="${fieldValue.FieldId}"  value=""  >
                                <div class="ig-row">
                                <button class="closemember" style="display: none"><img src="/public/img/close-1234.svg" /></button>
                                  <input type="text" class="searchval"  id="searchdropdownrole" placeholder="Search Members"  />
                                  <div class="drop-menu memberlistdiv" style="display:none">
                                 
                              </div>
                              </div>
                               
                              </div>

                                `)
                            }
                            if (fieldValue.MasterFieldId == 4) {
                                $(`#section${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text" class="date" data-id="" data-date="${fieldValue.DateFormat}" data-time="${fieldValue.TimeFormat}" placeholder="` + languagedata.Channell.seldate + ` "  id="date${fieldValue.FieldId}"   value="" readonly/>
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                  </div>
                `);
                            }

                            if (fieldValue.MasterFieldId == 5) {
                                $(`#section${fieldValue.SectionId}`).append(` 
            <div class="input-group  user-drop-down getvalue " id="dropoption" data-id="${fieldValue.Mandatory}" style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle atag${fieldValue.FieldId}" type="button" id="triggerId${fieldValue.FieldId}" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        Option
                                        </a>
                                        <input type="hidden" data-id="" class="fieldid" id="${fieldValue.FieldId}"   value=""  >
        
                                        <div class="dropdown-menu"id="drop${fieldValue.FieldId}" aria-labelledby="triggerId">
                
                                            
                                        </div>
                                        <label class="error manerr" id="opterr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
        
                                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 6) {
                                $(`#section${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date" data-id="" data-date="${fieldValue.DateFormat}" id="date${fieldValue.FieldId}"  placeholder="` + languagedata.Channell.seldate + `" readonly/>
                    
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                  </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 9) {
                                $(`#section${fieldValue.SectionId}`).append(` 

                <div class="input-group getvalue" id="radiogrb" data-id="${fieldValue.Mandatory}">
                <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
                <div class="ig-row">
                <input type="hidden" data-id="" class="radioval" id="${fieldValue.FieldId}" value="">
                <div class="button-col entry-radio-row flexx" id="radio${fieldValue.FieldId}">
              
            </div>
            <label class="error manerr" id="radioerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
            </div>
            </div>
            
              
            
                `);
                            }
                            if (fieldValue.MasterFieldId == 8) {
                                $(`#section${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" >
                <label  class="input-label"  data-id="${fieldValue.FieldId}" id="textarealabel" for="">${fieldValue.FieldName}</label>
             
                <div class="ig-row">
                  <textarea placeholder="`+ languagedata.Channell.pltexthere + `" data-id="" id="${fieldValue.FieldId}" ></textarea>
                </div>
                <label class="error manerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
              </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 2) {
                                $(`#section${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata.Channell.pltexthere + `" id="${fieldValue.FieldId}"   value="" />
                      </div>
                      <label class="error manerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 7) {
                                $(`#section${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata.Channell.pltexthere + `" id="${fieldValue.FieldId}"   value=""   />
                      </div>
                      <label class="error manerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                            }

                            if (fieldValue.MasterFieldId == 10) {
                                $(`#section${fieldValue.SectionId}`).append(`
                <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id="${fieldValue.Mandatory}"  style="margin-bottom: 20px">
         
                                         <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
         
                                         <a class="dropdown-toggle " type="button" id="triggerId" data-bs-toggle="dropdown"
                                             aria-haspopup="true" aria-expanded="false">
                                         Option
                                         </a>
                                         <input type="hidden" data-id="" class="checkfieldid" id="${fieldValue.FieldId}"  value=""  >
                                         <div class="dropdown-menu additional-entry-drop" id="check${fieldValue.FieldId}" aria-labelledby="triggerId">
                 
                                             
                                         </div>
                                         <label class="error manerr" id="opterrr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
         
                                     </div>
                `)

                            }

                            if (result.Entries.TblChannelEntryField != null) {
                                result.Entries.TblChannelEntryField.forEach(function (field) {

                                    $(`#${field.FieldId}`).val(field.FieldValue)

                                    if (`${field.FieldName}` == "Media Gallery" && `${field.FieldValue}` != "") {
                                        $(".uploadFolders").removeClass("add-galleryfolder")
                                        $(".uploadFolders").addClass("removeFolder")
                                        $("#fl-img").attr("src", "/public/img/folder-icon.svg")

                                        var fpath = field.FieldValue

                                        var parts = fpath.split('/');

                                        parts = parts.filter(function (part) {
                                            return part.length > 0;
                                        });

                                        var lastIndex = parts[parts.length - 1];

                                        $(`.fl-path${field.FieldId}`).text(lastIndex)
                                        $("#fl-remove").show()
                                    }

                                    $(`#date${field.FieldId}`).attr('data-id', field.Id)

                                    $(`#date${field.FieldId}`).val(field.FieldValue)

                                    $(`#${field.FieldId}`).attr('data-id', field.Id)

                                    $(`#triggerId${field.FieldId}`).text(field.FieldValue)

                                    $(`.${field.FieldId}`).attr('data-id', field.Id)

                                    // $(`.memberinput${field.FieldId}`).val(field.FieldValue)

                                    if (field.FieldName == "Video URL" && `${field.FieldValue}` != "") {

                                        urlvalues = $("#urldatas").val(field.FieldValue)

                                        videourlarr = field.FieldValue.split(",")

                                    }
                                })
                            }
                            $(".searchval").val(result.Memberlist)

                            if (fieldValue.OptionValue !== null) {
                                fieldValue.OptionValue.forEach(function (option) {


                                    if (fieldValue.MasterFieldId == 5) {
                                        $(`#drop${fieldValue.FieldId}`).append(`
                    <button type="button" id="optionsid" class="dropdown-item" >${option.Value}</button>
                  `);
                                    }
                                    if (fieldValue.MasterFieldId == 10) {
                                        $(`#check${fieldValue.FieldId}`).append(`
                        <div class="chk-group chk-group-label checkboxdiv">
                        <input type="checkbox" class="checkboxid" id="Check${option.Id}">
                        <label for="Check${option.Id}" class="checkboxcls">${option.Value}</label>
                    </div>
                    `)
                                    }

                                    if (fieldValue.MasterFieldId == 9) {

                                        $(`#radio${fieldValue.FieldId}`).append(`

                     <div class="radio radio-label radiodiv" >
                     <input id="radio-${option.Id}" name="radio" class="radbtn"  type="radio">
                     <label for="radio-${option.Id}" class="radio-label">${option.Value}</label>
                   
                 </div>
                    `)
                                    }

                                })
                            }


                        });

                    }




                }
                if (videourlarr != "" && videourlarr != undefined) {

                    videourlarr.forEach((element, index) => {

                        $(".video-url-list").append(` <li  id="urlvalues">
                       <p id="youtubevalue">`+ element + `</p>
                       <a href="javascript:void(0);" data-id=`+ element + ` id="urldeletebtn"> <img src="/public/img/remove-folder.svg" alt="remove"></a>
                   </li>`)
                    })
                }



                // dropdown filter input box search

                $("#searchdropdownrole").keyup(function () {

                    var keyword = $(this).val().trim().toLowerCase()


                    if (keyword == "") {

                        $('.memberlistdiv').hide()
                        $('.closemember').hide()
                        $('.fieldid1').val('')
                    } else {

                        $('.memberlistdiv').show()
                        $('.closemember').show()
                    }

                    fetch(`/channel/memberdetails/?keyword=${keyword}`, {
                        method: "GET"
                    })
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Network response was not ok');
                            }
                            return response.json();
                        })
                        .then(data => {

                            $('.memberlistdiv').empty();
                            if (data == null) {

                                $('.memberlistdiv').append(` <div class="noData-foundWrapper" id="nodatafounddesign"
                                    style="margin-top: -40px;display: none;">
        
                                    <div class="empty-folder">
                                        <img style="max-width: 20px;" src="/public/img/folder-sh.svg" alt="">
                                        <img src="/public/img/shadow.svg" alt="">
                                    </div>
                                    <h1 style="text-align: center;font-size: 10px;" class="heading">
                                    ` + languagedata.oopsnodata + `</h1>
                                    </div>
                                     `)
                                $("#nodatafounddesign").show()
                            }
                            if (data != null) {
                                $("#nodatafounddesign").hide()
                                for (let x of data) {
                                    $('.memberlistdiv').append(` <button type="button" id="optionsid1" class="dropdown-item optionsid1" data-id=${x.Id} >${x.FirstName} ${x.LastName}</button>`)

                                }
                            }

                        })
                        .catch(error => {
                            console.error('There was a problem with your fetch operation:', error);
                        });


                })


                $('.getvalue').each(function () {
                    obj = {}
                    obj = {
                        "name": $(this).children('.input-label').text().trim(),
                        "fid": $(this).children('.input-label').attr('data-id'),
                        "value": $(this).find('input').val() || $(this).find('textarea').val() || $(this).find('.fieldid').val() || $(this).find('.fieldid1').val() || $(this).find(".fl-pathfld").val() || $(this).find('#urldatas').val(),
                        "fieldid": $(this).find('input').attr('data-id') || $(this).find('textarea').attr('data-id') || $(this).find('.fieldid').attr('data-id') || $(this).find('.fieldid1').attr('data-id') || $(this).find(".fl-pathfld").attr('data-id') || $(this).find('#urldatas').attr('data-id'),

                    }

                    if ((obj.fieldid == undefined) || (obj.fieldid == '')) {
                        obj.fieldid = '0'

                    }
                    if ((obj.value != '') && (obj.value != undefined)) {
                        channeldata.push(obj)
                        newchanneldata.push(obj)
                    }


                })
                var radval = $('.radioval').val()
                $('.radiodiv').each(function () {
                    if ($(this).find('label').text() == radval) {
                        $(this).find('input').prop('checked', true)
                    }
                })

                var cheval = $('.checkfieldid').val()

                if (cheval != "" && cheval != undefined) {
                    var values = cheval.split(','); // Split the string by comma
                    for (var i = 0; i < values.length; i++) {
                        checkboxarr.push(values[i]); // Push each value to the array
                    }

                    $('.checkboxdiv').each(function (index) {

                        if (cheval.includes($(this).find('label').text())) {
                            $(this).find('input').prop('checked', true);
                        }
                    })
                    $('.checkboxdiv').each(function (index) {

                        if (cheval.includes($(this).find('input').attr('data-id'))) {
                            $(this).find('input').prop('checked', true);
                        }
                    })
                }
                var cheval1 = $('.checkfieldid1').val()

                if (cheval1 != "" && cheval1 != undefined) {

                    var values = cheval1.split(','); // Split the string by comma
                    for (var i = 0; i < values.length; i++) {
                        membercheckarr.push(values[i]); // Push each value to the array
                    }

                    $('.checkboxdiv').each(function (index) {

                        if (cheval1.includes($(this).find('label').text())) {
                            $(this).find('input').prop('checked', true);
                        }
                    })
                    $('.checkboxdiv').each(function (index) {

                        if (cheval1.includes($(this).find('input').attr('data-id'))) {
                            $(this).find('input').prop('checked', true);
                        }
                    })
                }
            }
        })

    }
})


$(document).ready(function () {

    if ($('.channeltitle1').text() == '') {


        $('.channeltitle1').css('display', 'none')
    }
})
$(document).on('keyup', '#searchkey', function () {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('#searchkey').val() === "") {

            var currentLocation = window.location.href;


            if (currentLocation.includes("unpublish")) {

                window.location.href = "/channel/unpublishentries"

            } else if (currentLocation.includes("draft")) {

                window.location.href = "/channel/draftentries"

            } else if (currentLocation.includes("entrylist")) {

                window.location.href = "/channel/entrylist"
            }



        }

    }

})


$(document).on('click', '#filtercancel', function () {

    // $('#slctstatus').text('');
    var url = window.location.href;
    var cleanUrl = url.split("?")[0];
    window.location.href = cleanUrl
});

//**Dropdown role//
$(document).on('click', '.statuss', function () {

    drop = $(this).text()
    $('#slctstatus').text(drop)
    $('#statusid').val(drop)

    // $('.filter-dropdown-menu').addClass('show').css({
    //     position: 'absolute',
    //     inset: '0px 0px auto auto',
    //     margin: '0px',
    //     transform: 'translate(0px, 34px)'
    // });


})

$(document).on('click', '.deleteentry', function () {

    console.log("deleteentryyyy")

    var currentLocation = window.location.href;

    var pname


    if (currentLocation.includes("unpublishentries")) {

        pname = "unpublish"

    } else if (currentLocation.includes("draftentries")) {

        pname = "draft"

    } else if (currentLocation.includes("entrylist")) {

        pname = "publish"
    }

    $('.filter-dropdown-menu ').hide()

    var entryId = $(this).attr("data-id")

    var channame = $(this).attr("data-name")

    var del = $(this).closest("tr");

    $('.delname').text(del.find('td:eq(4)').text())

    $('.deltitle').text(languagedata.Channell.delentrytitle)

    $('.deldesc').text(languagedata.Channell.delentrycontent)

    $('#delid').attr('data-id', $(this).attr('data-id'))

    $('#delid').text("Delete")

    $('#delcancel').text(languagedata.no)

    var pageno = $(this).attr("data-page")

    if (pageno == "") {
        $('#delid').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame + "&pname=" + pname);

    } else {
        $('#delid').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame + "&pname=" + pname + "&page=" + pageno);

    }

    // $('#delid').parent('#delete').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame + "&page="+pageno);

    // $('#delid').parent('#delete').attr('href', "/channel/deleteentries/"+ entryId + "/" + channelname);


})


//*Publish and Unpublish status change */

var entryid

var channelname

var chlstatus

$(document).on("click", "#publish", function () {

    entryid = $(this).attr("data-id")
    channelname = $(this).attr("data-channelname")
    chlstatus = $(this).attr("data-status")


    $.ajax({

        url: "/channel/checkmandatoryfields/" + entryid,
        datatype: "json",
        type: "GET",
        data: {
            csrf: $("input[name='csrf']").val(),
            entryid: entryid
        },
        success: function (result) {


            if (result.Title == "" || result.Description == "" || result.CategoriesId == "") {

                $('#content').text("You could not Publish this entry please add mandatory fields");
                $('.delname').text(channelname)
                $('.deltitle').text("Add Mandatory Fields")
                $("#delid").text("Edit Entry")
                $("#delcancel").text(languagedata.cancel)
                if (channelname != "") {
                    $('#delid').attr('href', "/channel/editentry/" + channelname + "/" + entryid)

                } else {

                    $('#delid').attr('href', "/channel/editsentry/" + entryid)
                }

            } else {

                $('#content').text(languagedata.Channell.publishcontent);
                $('.delname').text(channelname)
                $('.deltitle').text(languagedata.Channell.publishtitle)
                $("#delid").text(languagedata.Channell.publish)
                $("#delcancel").text(languagedata.cancel)
                $('#delid').addClass('entrystatuschange')
            }
        }

    })


})

$(document).on("click", "#dltCancelBtn", function () {

    $('#delid').removeClass('entrystatuschange')
    $('#delid').removeClass('featurebtn')

})

$(document).on("click", "#unpublish", function () {

    entryid = $(this).attr("data-id")
    channelname = $(this).attr("data-channelname")
    chlstatus = $(this).attr("data-status")

    $('#content').text(languagedata.Channell.unpublishcontent);
    $('.delname').text(channelname)
    $('.deltitle').text(languagedata.Channell.unpublishtitle)
    $("#delid").text(languagedata.Channell.unpublish)
    $("#delcancel").text(languagedata.cancel)
    $('#delid').addClass('entrystatuschange')


})


$(document).on("click", "#feature", function () {

    console.log("feature")

    // entryid = $(this).attr("data-id")
    // channelname = $(this).attr("data-channelname")

    $('.deldesc').text("Are you Sure you Want to Feature this Entries into Website?");
    $('.delname').text(channelname)
    $('.deltitle').text("Feature Entry")
    $("#delid").text("Feature")
    $("#delid").addClass("featurebtn")
    $("#delcancel2").text(languagedata.cancel)
    $('#delid').attr('data-id', $(this).attr("data-id")).attr('data-channelname', $(this).attr("data-channelname")).attr('data-feature', $(this).attr("data-status"))

})

$(document).on("click", "#Unfeature", function () {

    // entryid = $(this).attr("data-id")
    // channelname = $(this).attr("data-channelname")

    $('#content').text("Are you Sure you Want to Unfeature this Entries into Website?");
    $('.delname').text(channelname)
    $('.deltitle').text("Unfeature Entry")
    $("#delid").text("Unfeature")
    $("#delid").addClass("featurebtn")
    $("#delcancel").text(languagedata.cancel)
    $('#delid').attr('data-id', $(this).attr("data-id")).attr('data-channelname', $(this).attr("data-channelname")).attr('data-feature', $(this).attr("data-status"))
})

$(document).on("click", ".featurebtn", function () {

    entryid = $(this).attr("data-id")
    channelname = $(this).attr("data-channelname")
    featurestatus = $(this).attr("data-feature")

    $.ajax({
        url: "/channel/feature",
        type: "post",
        data: {
            entryid: entryid,
            cname: channelname,
            feature: featurestatus,
            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        success: function (result) {

            window.location.reload()
        }
    })
})

$(document).on("click", ".entrystatuschange", function () {
    // if (window.location.href.includes('/channel/entrylist')) {

    $.ajax({
        url: "/channel/changestatus/" + entryid + "?" + "entry=" + entryid + "cname=" + channelname + "&&status=" + chlstatus,
        type: "post",
        data: {
            entryid: entryid,
            status: chlstatus,
            cname: channelname,
            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        success: function (result) {

            if (result) {

                if (chlstatus == 1) {

                    $("#statusname-" + entryid).text("Published");

                    $("#statusname-" + entryid).removeClass('unpublished').addClass('published')

                    $("#statusname-" + entryid).removeClass('draft').addClass('published')

                    $("#unpublish").show()

                    $("#publish").hide()

                    setCookie("get-toast", "Entry Published Successfully")

                    setCookie("Alert-msg", "success", 1)

                } else if (chlstatus == 2) {

                    $("#statusname-" + entryid).text("Unpublished");

                    $("#statusname-" + entryid).removeClass('published').addClass('unpublished')

                    $("#statusname-" + entryid).removeClass('draft').addClass('unpublished')

                    $("#unpublish").hide()

                    $("#publish").show()

                    setCookie("get-toast", "Entry Unpublished Successfully")

                    setCookie("Alert-msg", "success", 1)

                }
                // setCookie("get-toast", languagedata.Toast.entryupdatenotify)
                // setCookie("Alert-msg", "success", 1)
                window.location.reload()



            } else {

                setCookie("Alert-msg", languagedata.Toast.internalserverr)



            }
        }
    })
    // }
})


$('.delete-flag').click(function () {
    $('input[name=spaceimagepath]').val("")
    $('#spimagehide').attr('src', '')
    $(this).hide()
    // $('#cenimg').show()
    $('#spimagehide').hide()

    $(".imgSelect-image").show()

    $(".uplclass").show()

    $(".imgSelect-btn").show()

    $(".blogBottom-imgSelect").css("padding-block", "5.44rem 7.75rem")

    $("#spacedel-img").hide()

    $(this).hide()
})

//Save to Draft Function//

$('#draftbtn').click(function () {

    var url = window.location.href;


    var pageurl = window.location.search
    const urlpar = new URLSearchParams(pageurl)
    pageno = urlpar.get('page')



    if (url.includes('editentry') || url.includes("editsentry")) {

        var urlvalue = url.split('?')[0]
        var eid = urlvalue.split('/').pop()
        // var eid = url.split('/').pop();

    }
    var drafturl

    if (pageno == null) {
        drafturl = "/channel/draftentry/" + eid
    } else {
        drafturl = "/channel/draftentry/" + eid + "?page=" + pageno

    }
    var data = ckeditor1.getData();

    var entryId = $('#chanid').val()

    var img = $('#spimagehide').attr('src')

    var title = $('#titleid').val()

    var channelname = $('#blogn').attr('data-char')


    //  text =$('#text').val()


    blog = $('#encount').val()

    blogcount = $('#encount').attr('data-id')

    Removeslashcategory()

    // if (title === '' && data === '' && img === '') {
    //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descerr + ' </span></div>';
    //     $(notify_content).insertBefore(".header-rht");
    //     setTimeout(function () {
    //         $('.toast-msg').fadeOut('slow', function () {
    //             $(this).remove();
    //         });
    //     }, 5000); // 5000 milliseconds = 5 seconds

    // } else if (data === '' && img === '') {
    //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descimgerr + '</span></div>';
    //     $(notify_content).insertBefore(".header-rht");
    //     setTimeout(function () {
    //         $('.toast-msg').fadeOut('slow', function () {
    //             $(this).remove();
    //         });
    //     }, 5000); // 5000 milliseconds = 5 seconds

    // } else if (img === '') {
    //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.imgerr + '</span></div>';
    //     $(notify_content).insertBefore(".header-rht");
    //     setTimeout(function () {
    //         $('.toast-msg').fadeOut('slow', function () {
    //             $(this).remove();
    //         });
    //     }, 5000); // 5000 milliseconds = 5 seconds



    // } else if (title === '') {

    //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter The Title</span></div>';
    //     $(notify_content).insertBefore(".header-rht");
    //     setTimeout(function () {
    //         $('.toast-msg').fadeOut('slow', function () {
    //             $(this).remove();
    //         });
    //     }, 5000); // 5000 milliseconds = 5 seconds
    // }
    // else if (data === '') {

    //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descriptionentry + '</span></div>';
    //     $(notify_content).insertBefore(".header-rht");
    //     setTimeout(function () {
    //         $('.toast-msg').fadeOut('slow', function () {
    //             $(this).remove();
    //         });
    //     }, 5000); // 5000 milliseconds = 5 seconds


    // }

    // if (title !== '' && data !== '' && img !== '' && entryId == '') {

    //     $('#channelslct-error').show()
    // }

    // if (title !== '' && data !== '' && img !== '' && entryId !== '') {
    //     var flag = channelfieldvalidation()

    //     channelfieldkeyup()

    //     if (flag == false) {
    //         $('#channelModal').modal('show');
    //         return
    //     }

    //     var flag1 = categoryvalidation()
    // }
    // if (title && data != '' && img != '' && (flag == true) && (flag1 == true)) {

    if (data == '' && img == '' && title == '') {

        // if (title === '') {

        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter The Title or Description or Image</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    }
    // }
    if (title != '' || img != '' || data != '') {

        authername = $("#author").val()
        createtime = $("#cdtime").val()
        publishtime = $("#publishtime").val()
        readingtime = $("#readingtime").val()
        sortorder = $("#article").val()
        tagname = $("#tagname").val()
        extxt = $("#extxt").val()

        $.ajax({
            url: drafturl,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "text": data, "status": 0, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), "page": pageno, csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": sortorder, "tagname": tagname, "extxt": extxt },
            success: function (result) {

                // if (result == true) {

                // $('.blog-status-container').show()
                // setTimeout(function () {
                //     $('.blog-status-container').hide();
                // }, 5000);

                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>entry draft saved successfully</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

                setTimeout(function () {

                    if (result.Channelname == "") {
                        if (pageno == null) {
                            window.location.href = "/channel/editsentry/" + result.id

                        } else {
                            window.location.href = "/channel/editsentry/" + result.id + "?page=" + pageno

                        }

                    } else {



                        if (pageno == null) {
                            window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id

                        } else {
                            window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id + "?page=" + pageno

                        }
                    }
                }, 2000);


            }
        })
        // }

    }
})
// return flagg
// }


// $('#previewbtn').click(function () {

//     img = $('#spimagehide').attr('src')

//     title = $('#titleid').val()

//     var data = ckeditor1.getData();

//     $('#preview-img').attr('src', img)

//     $('.preview-head').text(title)

//     $('.heading-four').html(data)

// })


$('.avaliable-dropdown-item').click(function () {

    var id = $(this).find('input').data('id');

    var channelname = $(this).find('.para').text();

    var count = $(this).find('.para-extralight').text();

    $('.blogDropdown1').text(channelname).css('color', 'black');

    $('.blogDropdown1').attr('data-bs-original-title', channelname)

    $('.blogDropdown1').attr('data-char', channelname)


    // $('.blogDropdown').append('<img class="blogimg" src="/public/img/blog-drop.svg" alt="">');

    $('#blogcount').text(count)

    $('#chanid').val(id)


})

$('#publishbtn').click(function () {

    flagg = true

    var url = window.location.href;

    var pageurl = window.location.search
    const urlpar = new URLSearchParams(pageurl)
    pageno = urlpar.get('page')

    var homeurl

    if (pageno == null) {
        homeurl = "/channel/entrylist/"
    } else {
        homeurl = "/channel/entrylist/?page=" + pageno;
    }

    if (url.includes('editentry') || url.includes('editsentry')) {
        var urlvalue = url.split('?')[0]
        var eid = urlvalue.split('/').pop()

    }
    var data = ckeditor1.getData();

    var entryId = $('#chanid').val()

    var channelname = $('#blogn').attr('data-char')

    var img = $('#spimagehide').attr('src')

    var title = $('#titleid').val()

    text = $('#text').val()

    blog = $('#encount').val()

    blogcount = $('#encount').attr('data-id')

    Removeslashcategory()


    if (title === '' && data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descimgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.imgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    } else if (title === '') {

        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>Please Enter The Title</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    }

    else if (data === '') {

        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descriptionentry + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds


    }

    if (title !== '' && data !== '' && img !== '' && entryId == "0" || entryId == "") {

        $('#channelslct-error').show()
    }
    if (title !== '' && data !== '' && img !== '' && entryId != "0" && entryId != "") {


        var flag = channelfieldvalidation()

        channelfieldkeyup()

        var flag1 = categoryvalidation()

        if (flag == false) {
            $('#channelModal').modal('show');
            return
        }


    }

    if (title != '' && data != '' && img != '' && entryId != "0" && (flag == true) && (flag1 == true)) {

        authername = $("#author").val()
        createtime = $("#cdtime").val()
        publishtime = $("#publishtime").val()
        readingtime = $("#readingtime").val()
        sortorder = $("#article").val()
        tagname = $("#tagname").val()
        extxt = $("#extxt").val()

        $.ajax({
            url: "/channel/publishentry/" + eid,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "status": 1, "text": data, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": sortorder, "tagname": tagname, "extxt": extxt },
            success: function (result) {

                window.location.href = homeurl;
            }
        })

    }


})
//ENTRY SELECTED FUNCTION//

$(document).on('click', '.selectcheckbox', function () {

    entryid = $(this).attr('data-id')

    var status = $('.pb-tab').text()

    var htmlContent = '';

    if ($(this).prop('checked')) {

        selectedcheckboxarr.push({ "entryid": entryid, "status": status })

    } else {

        const index = selectedcheckboxarr.findIndex(item => item.entryid === entryid);

        if (index !== -1) {

            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }
    console.log(selectedcheckboxarr, "arrayvalue")

    if (selectedcheckboxarr.length != 0) {

        $('.selected-numbers').removeClass("hidden")

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus
        var img;

        if (selectedcheckboxarr[0].status === "Unpublished") {

            setstatus = "Publish";

            img = "/public/img/publish.png";

        } else if (selectedcheckboxarr[0].status === "Published") {

            setstatus = "Unpublish";

            img = "/public/img/unpublish-select.svg";

        }
        else if (selectedcheckboxarr[0].status === "Draft") {

            htmlContent = '';

        }


        if (allSame) {

            console.log("checksame")

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';


        } else {

            htmlContent = '';

        }

        console.log(htmlContent, "htmlcontent")

        $('#unbulishslt').html(htmlContent);


        var items

        if (selectedcheckboxarr.length == 1) {

            items = "Item Selected"
        } else {

            items = languagedata.itemselected
        }

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

        $('#deselectid').text(languagedata.selectall)

        $('#deselectid').addClass("selectall")

        if (!allSame || selectedcheckboxarr[0].status === "Draft") {

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171]')

            $('#unbulishslt').html("")
        } else {

            $('#seleccheckboxdelete').addClass('border-r border-[#717171]')
        }

    } else {

        $('.selected-numbers').addClass("hidden")
    }

    var allChecked = true;

    $('.selectcheckbox').each(function () {

        if (!$(this).prop('checked')) {

            allChecked = false;

            return false;
        }
    });

    if (allChecked) {

        $('#deselectid').text(languagedata.deselectall)

        $('#deselectid').removeClass("selectall")

        $('#deselectid').addClass("deselectid")
    }



    $('#Check').prop('checked', allChecked);

})
//ALL CHECKBOX CHECKED FUNCTION//


$(document).on('click', '#Check', function () {

    selectedcheckboxarr = []

    var htmlContent = '';


    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            entryid = $(this).attr('data-id')

            var status = $(this).parents('td').siblings('td').find('span').text().trim();

            selectedcheckboxarr.push({ "entryid": entryid, "status": status })

        })

        $('.selected-numbers').removeClass("hidden")

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus

        var img

        if (selectedcheckboxarr.length != 0 && allSame) {

            if (selectedcheckboxarr[0].status == "Unpublished") {

                setstatus = "Published"

                img = "/public/img/publish.png"


                htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">"' + setstatus + '"</span>';


            } else if (selectedcheckboxarr[0].status == "Published") {

                setstatus = "Unpublished"

                img = "/public/img/unpublish-select.svg";


                htmlContent = '<img src="' + img + '">' + setstatus;


            }
            else if (selectedcheckboxarr[0].status == "Draft") {

                $('#unbulishslt').html("")

                htmlContent = "";
            }

        } else {


            htmlContent = '';

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171]')
        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

    } else {

        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selected-numbers').addClass("hidden")
    }

    if (selectedcheckboxarr.length == 0) {

        $('.selected-numbers').addClass("hidden")
    }
})

$(document).on('click', '#seleccheckboxdelete', function () {



    $('.deltitle').text("Delete Entries?")

    $('#content').text('Are you sure want to delete selected Entries?')

    $('#delid').addClass('checkboxdelete')

    $('#delid').text("Delete")
})

$(document).on('click', '#unbulishslt', function () {

    $('.deltitle').text($(this).text() + " " + "Entries?")

    $('#content').text("Are you sure want to " + $(this).text() + " " + "selected Entries?")

    $('#delid').addClass('selectedunpublish')

    $('#delid').removeClass("entrystatuschange")

    textval = $(this).text()

    if (textval == "Publish") {

        $('#delid').text("Publish")
    }
    if (textval == "Unpublish") {

        $('#delid').text("Unpublish")
    }


})
//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {


    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    var newUrl = window.location.pathname; 

    $('.selected-numbers').hide()
    $.ajax({
        url: '/channel/deleteselectedentry',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "entryids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno,
            "url": newUrl


        },
        success: function (data) {


            if (data.value == true) {

                setCookie("get-toast", "Entry Deleted Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})
//Deselectall function//

$(document).on('click', '.deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass("hidden")

})

//multi select unpublish function//

$(document).on('click', '.selectedunpublish', function () {

    var url = window.location.href;


    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/channel/unpublishselectedentry',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "entryids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno,
            "url": url


        },
        success: function (data) {

            console.log(data, "result")

            var datastatus

            if (data.status == 2) {

                datastatus = "Unpublished"
            } else if (data.status == 1) {

                datastatus = "Published"
            } else if (data.status == 0) {

                datastatus = "Draft"
            }
            if (data.value == true) {

                setCookie("get-toast", "Entry " + datastatus + " Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })


})
//EDIT FUNCTION IN ENTRY PAGE//
$(document).on('click', '#editbtn', function () {

    $('.filter-dropdown-menu ').hide()

})


// function CKEDITORS() {

//     var url = $('#urlid').val();


//     CKEDITOR.ClassicEditor.create(document.getElementById("text"), {
//         toolbar: {
//             items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'horizontalLine', 'link', 'code'],
//             shouldNotGroupWhenFull: true
//         },
//         list: {
//             properties: {
//                 styles: true,
//                 startIndex: true,
//                 reversed: true
//             }
//         },
//         heading: {
//             options: [
//                 { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
//                 { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
//                 { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
//                 { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
//                 { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
//                 { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
//                 { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
//             ]
//         },
//         placeholder: languagedata.Channell.plckeditor,
//         fontFamily: {
//             options: [
//                 'default',
//                 'Arial, Helvetica, sans-serif',
//                 'Courier New, Courier, monospace',
//                 'Georgia, serif',
//                 'Lucida Sans Unicode, Lucida Grande, sans-serif',
//                 'Tahoma, Geneva, sans-serif',
//                 'Times New Roman, Times, serif',
//                 'Trebuchet MS, Helvetica, sans-serif',
//                 'Verdana, Geneva, sans-serif'
//             ],
//             supportAllValues: true
//         },
//         fontSize: {
//             options: [10, 12, 14, 'default', 18, 20, 22],
//             supportAllValues: true
//         },
//         htmlSupport: {
//             allow: [
//                 {
//                     name: /.*/,
//                     attributes: true,
//                     classes: true,
//                     styles: true
//                 }
//             ]
//         },
//         htmlEmbed: {
//             showPreviews: true
//         },
//         link: {
//             decorators: {
//                 addTargetToExternalLinks: true,
//                 defaultProtocol: 'https://',
//                 toggleDownloadable: {
//                     mode: 'manual',
//                     label: 'Downloadable',
//                     attributes: {
//                         download: 'file'
//                     }
//                 }
//             }
//         },
//         mention: {
//             feeds: [
//                 {
//                     marker: '@',
//                     feed: [
//                         '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
//                         '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
//                         '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
//                         '@sugar', '@sweet', '@topping', '@wafer'
//                     ],
//                     minimumCharacters: 1
//                 }
//             ]
//         },
//         removePlugins: [
//             'ExportPdf',
//             'ExportWord',
//             'CKBox',
//             'CKFinder',
//             'EasyImage',
//             'RealTimeCollaborativeComments',
//             'RealTimeCollaborativeTrackChanges',
//             'RealTimeCollaborativeRevisionHistory',
//             'PresenceList',
//             'Comments',
//             'TrackChanges',
//             'TrackChangesData',
//             'RevisionHistory',
//             'Pagination',
//             'WProofreader',
//             'MathType',
//             'SlashCommand',
//             'Template',
//             'DocumentOutline',
//             'FormatPainter',
//             'TableOfContents',
//             'PasteFromOfficeEnhanced'
//         ]
//     }).then(ckeditor => {
//         ckeditor1 = ckeditor
//         ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
//             return {
//                 upload: () => {
//                     return loader.file.then(file => {
//                         return new Promise((resolve, reject) => {
//                             const formData = new FormData();
//                             formData.append('file', file);
//                             formData.append('csrf', $("input[name='csrf']").val())
//                             fetch(url + '/channel/imageupload', {
//                                 method: 'POST',
//                                 body: formData
//                             })
//                                 .then(response => response.json())
//                                 .then(data => {
//                                     if (data.success) {
//                                         resolve({
//                                             default: data.url // URL to the uploaded image

//                                         });
//                                     } else {
//                                         reject(data.error);
//                                     }
//                                 })
//                                 .catch(error => {
//                                     reject(error);
//                                 });
//                         });
//                     });
//                 }
//             };
//         };
//     })
//         .catch(error => {
//             console.error(error);
//         });

// }

//CATEGORY SAVE FUNCTION//

$(document).on('click', '#categorysave', function () {
    categoryIds = [];

    categoryvalidation()

    $('.category-unselect-btn').each(function () {
        var categoryId = $(this).attr('data-id');
        categoryIds.push(categoryId);
    });

    if (categoryIds.length !== 0) {



        // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Channell.categoryerr + '</span></div>';
        // $(notify_content).insertBefore(".header-rht");
        // setTimeout(function () {
        //     $('.toast-msg').fadeOut('slow', function () {
        //         $(this).remove();
        //     });
        // }, 5000); // 5000 milliseconds = 5 seconds
        $('#categoryModal').modal('hide');
        $("#category-error").hide()

        if (flagg == false) {
            $('#draftbtn').click()
        } else {
            $('#publishbtn').click()
        }
    }

})

$.validator.addMethod("customLength", function (value, element, params) {
    var minLength = params[0];
    var maxLength = params[1];
    return this.optional(element) || (value.length >= minLength && value.length <= maxLength);
}, $.validator.format("Please enter between {0} and {1} characters."));

//SEO DETAILS SAVE FUNCTION//
$(document).on('click', '#seosavebtn', function () {



    seodetails = {}

    $("#seoform").validate({
        rules: {
            metatitle: {
                customLength: [1, 60],
                titlemethod: true,
                titlespace: true
            },
            metadesc: {
                customLength: [1, 150]
            },



        },
        messages: {
            metatitle: {
                customLength: "*" + languagedata.Channell.metatitlemsg,
                titlemethod: "*Please Enter Meta tag title",
                titlespace: "* " + languagedata.spacergx
            },
            metadesc: {
                customLength: "*" + languagedata.Channell.metadescmsg,

            }
        }
    });
    var formcheck = $("#seoform").valid();
    if ((formcheck == true)) {
        // $('#seoform')[0].submit();
        $('#seoModal').modal('hide');

        title = $('#metatitle').val()
        desc = $('#metadesc').val()
        keyword = $('#metakey').val()
        slug = $('#metaslug').val()
        imgtag = $("#imgtag").val()

        if (title !== '' || desc !== '' || keyword !== '' || slug !== '') {
            // seodetails.push({ title: title, desc: desc, keyword: keyword, slug: slug });
            seodetails = { title: title, desc: desc, keyword: keyword, slug: slug, imgtag: imgtag }
        }


        if (Object.keys(seodetails).length !== 0) {

            // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Channell.seoerr + '</span></div>';
            // $(notify_content).insertBefore(".header-rht");
            // setTimeout(function () {
            //     $('.toast-msg').fadeOut('slow', function () {
            //         $(this).remove();
            //     });
            // }, 5000); // 5000 milliseconds = 5 seconds

            if (flagg == false) {
                $('#draftbtn').click()
            } else {
                $('#publishbtn').click()
            }
        }
    }
    else {


        $('#seoModal').modal('show');
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })

    }
})

$(document).on('click', '.btn-close', function () {
    $('#changepicModal').modal('hide');
})

$(document).ready(function () {



    $('.categorypdiv').each(function () {

        var length = $(this).children('.choose-cat-list-col').children('.categoryname').length - 1

        $(this).children('.choose-cat-list-col').children('.categoryname').each(function (index) {

            if (length == index) {

                $(this).next().remove();
            }

        })

    })

    $('.selectedcategorydiv').each(function () {

        var length = $(this).children('.choose-cat-list-col').children('.categoryname').length - 1

        $(this).children('.choose-cat-list-col').children('.categoryname').each(function (index) {

            if (length == index) {

                $(this).next().remove();
            }

        })

    })

})


document.querySelectorAll('.input-group input').forEach(input => {

    input.addEventListener('focus', (event) => {
        event.target.closest('.input-group').classList.add('input-group-focused');
    });

    input.addEventListener('blur', (event) => {
        event.target.closest('.input-group').classList.remove('input-group-focused');
    });
});

document.querySelectorAll('.secid input').forEach(input => {

    input.addEventListener('focus', (event) => {
        event.target.closest('.input-group').classList.add('input-group-focused');
    });

    input.addEventListener('blur', (event) => {
        event.target.closest('.input-group').classList.remove('input-group-focused');
    });
});

//**description focus function */
const Desci = document.getElementById('metadesc');


// Desci.addEventListener('focus', () => {

//     Desci.closest('.input-group').classList.add('input-group-focused');
// });
// Desci.addEventListener('blur', () => {
//     Desci.closest('.input-group').classList.remove('input-group-focused');
// });

$(document).on('click', '#seocancelbtn', function () {

    $('#metatitle').val('')
    $('#metadesc').val('')
    $('#metakey').val('')
    $('#metaslug').val('')
    $('#imgtag').val('')
    $('#metatitle-error').hide()
    $('.input-group').removeClass('input-group-error')
    $('#prefixerr').hide()
    $('.field').removeClass('error')

})

$(document).on('click', '#seoclose', function () {

    // $('#metatitle').val('')
    // $('#metadesc').val('')
    // $('#metakey').val('')
    // $('#metaslug').val('')
    $('#metatitle-error').hide()
    $('.input-group').removeClass('input-group-error')
    $('#prefixerr').hide()
    $('.field').removeClass('error')


})

// $(document).ready(function () {

//     var url = window.location.href;

//     var id = url.split('/').pop();
//     console.log("id", id);
//     if (id === 'newentry') {
//         console.log("hhkkkk")
//         var dropdownItems = $('.avaliable-dropdown-item');
//         dropdownItems.eq(0).click();
//     }
// });

$(document).on('click', '#channelfieldsave', function () {

    channeldata = [];
    authername = $("#author").val()
    createtime = $("#cdtime").val()
    publishtime = $("#publishtime").val()
    readingtime = $("#readingtime").val()
    sortorder = $("#article").val()
    tagname = $("#tagname").val()
    extxt = $("#extxt").val()
    videourl = $("#urldatas").val()

    var flag = channelfieldvalidation()
    var url = window.location.href;
    channelfieldkeyup()
    if (sortorder != "") {


        if (url.includes('editentry') || url.includes('editsentry')) {
            var urlvalue = url.split('?')[0]
            var eid = urlvalue.split('/').pop()

        }

        $.ajax({
            url: "/channel/checkentriesorder",
            type: "POST",
            async: false,
            data: { "chid": $("#chanid").val(), csrf: $("input[name='csrf']").val(), "order": sortorder, "eid": eid },
            success: function (data) {

                console.log("result", data.value, data.value == true);

                if (data.value == true) {

                    $("#sort-error").show()

                    flag = false

                } else {

                    $("#sort-error").hide()

                    flag = true

                }
            }
        })
    }
    $('.getvalue').each(function () {

        obj = {}
        obj = {
            "name": $(this).children('.input-label').text().trim(),
            "fid": $(this).children('.input-label').attr('data-id'),
            "value": $(this).find('input').val() || $(this).find('textarea').val() || $(this).find('.fieldid').val() || $(this).find('.fieldid1').val() || $(this).find('#urldatas').val() || $(this).find("#fl-pathfld").val(),
            "fieldid": $(this).find('input').attr('data-id') || $(this).find('textarea').attr('data-id') || $(this).find('.fieldid').attr('data-id') || $(this).find('.fieldid1').attr('data-id') || $(this).find(".fl-pathfld").attr('data-id') || $(this).find('#urldatas').attr('data-id'),

        }
        console.log("object", obj.value)
        if ((obj.fieldid == undefined) || (obj.fieldid == '')) {
            obj.fieldid = '0'

        }
        // if ((obj.value != '') && (obj.value != undefined)) {
        // }
        channeldata.push(obj)



    })

    if ((flag == true)) {

        if (flagg == false) {
            $('#draftbtn').click()
        } else {
            $('#publishbtn').click()
        }




    }

    if (flag == true) {
        $('#channelModal').modal('hide');
    }
})

$(document).on('click', '#addcategory', function () {


    channelid = $('#chanid').val()

    title = $('#titleid').val()
    if (title == '') {
        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Channell.titleerr + `</span></div>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    }
    else if ((title != '') && (channelid == '')) {

        $('#channelslct-error').show()

    } else {

        $('#categoryModal').modal('show')
    }


    // categoryIds = catid.split(',');

    console.log("categorids", categoryIds)
    var catselect = document.querySelectorAll('.categories-list-child');
    console.log(catselect, "catelist");
    if (Channelid == channelid) {
        categoryIds.forEach(function (categoryId) {
            var matchedH3 = document.querySelector('a[data-id="' + categoryId + '"]');
            console.log("match", matchedH3);

            matchedH3.click()

        });
    }
    Removeslashcategory()


})

$(document).on('click', '#channelfield', function () {

    console.log(newchanneldata, "channeldata")

    var radval = $('.radioval').val()
    console.log(radval, "radiovalue")

    if (newchanneldata != '') {
        bindchanneldata()
    }


    channelid = $('#chanid').val()

    title = $('#titleid').val()

    if (title == '') {
        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Channell.titleerr + `</span></div>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if ((title != '') && (channelid == '')) {

        // console.log("alertt")
        // notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Channell.channelselect+ `</span></div>`;
        // $(notify_content).insertBefore(".header-rht");
        // setTimeout(function () {
        //     $('.toast-msg').fadeOut('slow', function () {
        //         $(this).remove();
        //     });
        // }, 5000); // 5000 milliseconds = 5 seconds
        $('#channelslct-error').show()

    }
    else {
        $('#channelModal').modal('show')
    }
})

$(document).on('click', '#seo', function () {
    title = $('#titleid').val()
    if (title == '') {

        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata.Channell.titleerr + `</span></div>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    }
    else {
        $('#seoModal').modal('show')
    }
})


$(document).on('click', '#configbtn', function () {


    $('#Sortsection').html("");

    id = $('#chanid').val()

    // $('#ChannelEditmodal').modal('show');
    $.ajax({
        url: '/channels/getfields',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "id": id,
            csrf: $("input[name='csrf']").val()
        },
        success: function (data) {

            console.log("dataaaa", data)

            sections = data.Section

            if (sections == null) {

                sections = []
            }


            if (data.FieldValue != null) {

                fiedlvalue = data.FieldValue

            } else {

                fiedlvalue = []
            }


            if (data.SelectedCategory != null) {

                SelectedCategoryValue = data.SelectedCategory
            }

            var sectionstr = ``

            if (data.section != "") {


                for (let x of data.Section) {

                    sectionstr += `<div class="section-fields-content drop-able ui-droppable" style="margin-bottom: 10px;" id="section-` + x['SectionId'] + `` + x['SectionNewId'] + `" data-id ="` + x['SectionId'] + `"  data-newid="` + x['SectionNewId'] + `" data-orderindex="` + x['OrderIndex'] + `">
            <div class="section-class">
              <input type="text" class="sectionname" placeholder="" style="font-size:0.825rem" value="`+ x['SectionName'] + `">
              <button class="sectiondel" style="margin-right: 11px;"><img src="/public/img/bin.svg"></button>   
            </div>
          </div>`

                }

                $('#Sortsection').append(sectionstr);

                for (let x of data.Section) {

                    SortableSection(x['SectionId'], x['SectionNewId'])

                }
            }
            if (data.FieldValue != null) {


                for (let [index, x] of data.FieldValue.entries()) {

                    $('.section-fields-content').each(function () {

                        if ($(this).attr('data-id') == x['SectionId'] && $(this).attr('data-newid') == x['SectionNewId']) {

                            var fieldstr = `<div class="section-fields-content-child fieldapp" data-masterfieldid=` + x.MasterFieldId + ` data-newfieldid=` + x.NewFieldId + ` data-fieldid=` + x.FieldId + ` data-orderindex=` + x.OrderIndex + ` id=fieldapp` + x.FieldId + `` + x.NewFieldId + `>
                <input type="hidden" class="orderval" value="">
                    <div class="img-div">
                        <img src="`+ x.IconPath + `" alt="">
                    </div>
                    <h3 class="heading-three" id=fieldtitle`+ x.FieldId + `` + x.NewFieldId + `>` + x.FieldName + `</h3>
                    <a href="javascript:void(0)" class="field-delete delete" data-id="`+ id + `" >
                        <img src="/public/img/bin.svg" alt="">
                    </a>
                </div> `

                            $(this).append(fieldstr);

                        }

                    })
                    if (fiedlvalue[index].OptionValue == null) {

                        fiedlvalue[index].OptionValue = []

                    }


                }
            }
            if (sections.length == 0) {

                var sectionobj = SectionObjectCreate()

                sectionobj.SectionName = "New Section"
                sections.push(sectionobj)

                $("#Sortsection").append(`<div class="section-fields-content drop-able active ui-droppable" style="margin-bottom: 10px;" id="section-0` + sectionid + `" data-newid=${sectionid} data-id=0 data-orderindex=${orderindex}>
          <div class="section-class" >
            <input type="text" class="sectionname" placeholder="" style="font-size:0.825rem" value="New Section"/>
            <button class="sectiondel" style="margin-right: 11px" ><img src="/public/img/bin.svg"></button>   
          </div>
        </div>`)

                sectionid++;

                orderindex++;

            }


            // drop
            $(".drop-able").droppable({
                tolerance: 'pointer',

                drop: function (ev, ui) {

                    var id = ui.draggable[0]['dataset'].id;
                    var name = ui.draggable.text().trim().split("\n");
                    var img = ui.draggable.find("img").attr("src");
                    var newsection = ui.draggable[0]['id']
                    console.log("newsectiondrop1111", id, name, img, newsection)
                    if (newsection != "newsection") {
                        $(this).append(`
                    <div class="section-fields-content-child fieldapp" data-masterfieldid=`+ id + ` data-newfieldid=` + newFieldid + ` data-fieldid=` + fielid + ` data-orderindex=` + orderindex + ` id=fieldapp` + fielid + `` + newFieldid + `>
                    <input type="hidden" class="orderval" value="">
                        <div class="img-div">
                            <img src="`+ img + `" alt="">
                        </div>
                        <h3 class="heading-three" id=fieldtitle`+ fielid + `` + newFieldid + `>` + name[0] + `</h3>
                        <a href="javascript:void(0)" class="field-delete delete" data-id="`+ id + `" >
                            <img src="/public/img/bin.svg" alt="">
                        </a>
                    </div>    
                `)


                        newFieldid = newFieldid + 1;
                        orderindex = orderindex + 1;
                    }

                },
            });


        }
    })

})

function channelfieldvalidation() {

    var flag = true;
    $('.getvalue').each(function () {

        charallowed = $(this).data('char')

        if ((($(this).find('input').val() === '') && ($(this).data('id') == 1)) || (($(this).find('textarea').val() === '') && ($(this).data('id') == 1))) {

            $(this).addClass('input-group-error');
            $(this).find('.manerr').show()
            flag = false
        }
        if (($(this).find('input,textarea').val() !== '') && charallowed != 0) {
            if ($(this).find('input,textarea').val().length > charallowed) {
                $(this).addClass('input-group-error');
                $(this).find('.cerr').show()
                flag = false
            }
        }
    })
    return flag;
}

function channelfieldkeyup() {

    $('.getvalue').each(function () {

        $(this).find('input,textarea').keyup(function () {
            if (($(this).val() !== '') || ($(this).text() !== '')) {

                $(this).closest('.input-group').removeClass('input-group-error');
                $(this).parent('.ig-row').siblings('.error').hide();
            } else {

                $(this).closest('.input-group').addClass('input-group-error');
                $(this).parent('.ig-row').siblings('.manerr').show();


            }

        })
        $(this).find('.date').change(function () {
            if (($(this).val() !== '')) {

                $(this).closest('.input-group').removeClass('input-group-error');
                $(this).parent('.ig-row').siblings('.error').hide();
            } else {

                $(this).closest('.input-group').addClass('input-group-error');
                $(this).parent('.ig-row').siblings('.manerr').show();


            }

        })

    })


}

$(document).on('click', '#optionsid', function () {
    option = $(this).text()
    console.log("dropdown", option)
    $('.fieldid').val(option)

    if (($('.fieldid').val()) !== '') {


        $('#dropoption').removeClass('input-group-error')
        $('#opterr').css('display', 'none');
    }

})
$(document).on('click', '#optionsid1', function () {
    option = $(this).attr('data-id')
    console.log("dropdown", option)
    $('.fieldid1').val(option)
    test = $(this).text()
    $('#searchdropdownrole').val(test)
    $('.memberlistdiv').hide()

    // if (($('.fieldid1').val()) !== '') {

    //     console.log("yyyy")

    //     $('#dropoption').removeClass('input-group-error')
    //     $('#opterr').css('display', 'none');
    // }

})
//seo validation//
function Validationcheck() {

    if ($('#metatitle').hasClass('error')) {
        $('#grbname').addClass('input-group-error');
    } else {
        $('#grbname').removeClass('input-group-error');
    }

    if ($('#metadesc').hasClass('error')) {
        $('#grbdesc').addClass('input-group-error');
    } else {
        $('#grbdesc').removeClass('input-group-error');
    }

}

function Seovalidationspce() {
    var flag = true

    // $(document).on('keyup','#metatitle',function(){
    var value = $('#metatitle').val();

    if (value[0] === ' ') {

        $('#prefixerr').show();
        $('#prefixerr').text("No Space in Prefix ");
        $('#grbname').addClass('input-group-error');
        flag = false;


    } else {
        $('#prefixerr').hide();
        $('#grbname').removeClass('input-group-error');
        flag = true;

    }


    // });

    return flag;
}
$(document).on('click', '#channelfieldcancel', function () {
    $('.getvalue').each(function () {
        $(this).find('input,textarea,.fieldid').val('')
        $(this).find('a').text('')
        $(this).find('.radbtn').prop('checked', false)
        $(this).find('.checkboxid').prop('checked', false)

    })
    $('#channelModal').modal('hide')

})

$(document).on('click', '#categorycancel', function () {

    // $('.selectedcategorydiv').remove()
    $('.category-unselect-btn').click()

})




$(document).on('click', '.checkboxid', function () {

    if ($(this).is(':checked')) {

        option = $(this).next('label').text()
        console.log("option", option)

        checkboxarr.push(option)
        $('.checkfieldid').val(checkboxarr.join(','))

    }
    else {
        var option = $(this).next('label').text();
        checkboxarr.splice(checkboxarr.indexOf(option), 1);
        console.log("may", option, checkboxarr)
        $('.checkfieldid').val(checkboxarr.join(','));
    }


    // if ($('.checkfieldid').val() != '') {

    //     $('#opterrr').hide()
    //     $('#checkdrop').removeClass('input-group-error')

    // } else {
    //     $('#opterrr').show()
    //     $('#checkdrop').addClass('input-group-error')

    // }
})

$(document).on('click', '.checkboxid1', function () {

    if ($(this).is(':checked')) {

        option = $(this).attr('data-id')
        console.log("option", option)

        membercheckarr.push(option)
        $('.checkfieldid1').val(membercheckarr.join(','))

    }
    else {
        var option = $(this).attr('data-id');
        membercheckarr.splice(membercheckarr.indexOf(option), 1);
        console.log("may", option, membercheckarr)
        $('.checkfieldid1').val(membercheckarr.join(','));
    }


    if ($('.checkfieldid1').val() != '') {

        $('#opterrr').hide()
        $('#checkdrop').removeClass('input-group-error')

    } else {
        $('#opterrr').show()
        $('#checkdrop').addClass('input-group-error')

    }
})
$(document).on('click', '.radbtn', function () {

    option = $(this).next('label').text()
    console.log("optiongggg", option)
    $('.radioval').val(option)
    if ($('.radioval').val() != '') {
        $('#radiogrb').removeClass('input-group-error')
        $('#radioerr').hide()
    } else {
        $('#radiogrb').addClass('input-group-error')
        $('#radioerr').show()
    }

})

//date picker function//
$(document).on('focus', '.date', function () {


    // $(this).trigger('click');

    var dateformat = $(this).attr('data-date');

    var timeformat = $(this).attr('data-time');

    console.log(timeformat, dateformat, "format")

    if (dateformat != "" && timeformat != "" && timeformat != undefined) {

        var formatset

        formatset = dateformat + " hh:mm a "

        console.log(formatset, "set")

        $('#' + $(this).attr('id')).bootstrapMaterialDatePicker({
            setMaxDate: moment(),
            shortTime: false, weekStart: 0, time: true, format: formatset, maxDate: new Date()
        });


    } else if (dateformat != "") {

        $('#' + $(this).attr('id')).bootstrapMaterialDatePicker({ setMaxDate: moment(), weekStart: 0, shortTime: false, time: false, format: dateformat, maxDate: new Date() });

    }

    //  $('#'+$(this).attr('id')).bootstrapMaterialDatePicker({ "setMaxDate": moment(),time:false});
})

$(document).on('click', '#fieldcancelid', function () {
    $('.getvalue').removeClass('input-group-error')
    $('.error').hide()
})


function categoryvalidation() {

    var flag = true;

    if (categoryIds.length == 0) {

        flag = false;

        $('#categoryModal').modal('show');

        $("#category-error").show()


    } else {
        $('#categoryModal').modal('hide');
    }
    return flag;
}

$(document).on('click', '.category-select-btn', function () {

    $("#category-error").hide()

})

$(document).on('click', '#cateclose', function () {

    $('.category-unselect-btn').click()

})


const averageWordWidth = 15;

function adjustRows() {
    var textarea = document.getElementById('titleid');
    textarea.style.height = '100% ';
    // console.log(textarea,"jjjj")
    var textareaWidth = textarea.clientWidth;
    var averageWordsPerLine = Math.floor(textareaWidth / averageWordWidth);
    var lineCount = Math.ceil(textarea.value.length / averageWordsPerLine);
    textarea.rows = lineCount > 1 ? lineCount : 2;
}



$.validator.addMethod("titlemethod", function (value) {

    var flag = true;

    if (($('#metatitle').val() == '') && ($('#metadesc').val() != '' || $('#metakey').val() != '' || $('#metaslug').val() != '')) {
        flag = false;
    }

    return flag;
});

$.validator.addMethod('titlespace', function () {


    var value = $('#metatitle').val();

    var flag = true;

    if ($('#metatitle').val() !== '' && (value[0] === ' ')) {
        flag = false;
    }
    return flag;
})

$(document).on('keyup', '#titleid', function () {

    if ($('#titleid').val() == '') {

        var textarea = document.getElementById('titleid');
        textarea.style.height = '3rem ';

    }

})


function Removeslashcategory() {

    $('.categorypdiv').each(function () {

        var length = $(this).find('.categoryname').length - 1

        var name = $(this).find('.categoryname')

        $(this).find('.categoryname').each(function (index) {

            if (length == index) {

                $(this).next().remove();
            }

        })

    })

    $('.selected_category').each(function () {

        var length = $(this).find('.categoryname').length - 1

        $(this).find('.categoryname').each(function (index) {

            if (length == index) {

                $(this).next().remove();
            }

        })

    })
}



$(document).on('focus', '.getvalue input', function () {

    $(this).parents('.getvalue').addClass("input-group-focused");
})
$(document).on('blur', '.getvalue input', function () {

    $(this).parents('.getvalue').removeClass("input-group-focused");
})
$(document).on('focus', '.getvalue textarea', function () {

    $(this).parents('.getvalue').addClass("input-group-focused");
})
$(document).on('blur', '.getvalue textarea', function () {

    $(this).parents('.getvalue').removeClass("input-group-focused");
})


$(document).on('click', '#uptchannelfield', function () {



    if (fiedlvalue == '') {

        $('.fieldapp').each(function () {

            var obj = CreatePropertiesObjec()

            obj.MasterFieldId = parseInt($(this).attr('data-masterfieldid'))

            obj.NewFieldId = parseInt($(this).attr('data-newfieldid'))

            obj.FieldName = $(this).children('.heading-three').text();

            obj.IconPath = $(this).children('.img-div').children('img').attr('src')

            obj.SectionId = parseInt($(this).parent('.section-fields-content').attr('data-id'));

            obj.SectionNewId = parseInt($(this).parent('.section-fields-content').attr('data-newid'));

            obj.SectionName = $(this).parent('.section-fields-content').children('.section-class').children('input').val();

            fiedlvalue.push(obj)

            console.log(fiedlvalue, "arrayfield")
        })
    }

    $('.fieldapp').each(function () {



        const index = fiedlvalue.findIndex(obj => {

            return obj.FieldId == $(this).attr('data-fieldid') && obj.NewFieldId == $(this).attr('data-newfieldid');

        });

        if (index < 0) {

            var obj = CreatePropertiesObjec()

            obj.MasterFieldId = parseInt($(this).attr('data-masterfieldid'))

            obj.NewFieldId = parseInt($(this).attr('data-newfieldid'))

            obj.FieldName = $(this).children('.heading-three').text();

            obj.IconPath = $(this).children('.img-div').children('img').attr('src')

            obj.SectionId = parseInt($(this).parent('.section-fields-content').attr('data-id'));

            obj.SectionNewId = parseInt($(this).parent('.section-fields-content').attr('data-newid'));

            obj.SectionName = $(this).parent('.section-fields-content').children('.section-class').children('input').val();

            fiedlvalue.push(obj)
        }

    })

    SortFieldsandSection()

    flg = FieldValidation1()

    console.log("flg==", flg);

    if (!flg) {

        return false
    }


    //   var name = $('#channelname').val();

    //   var desc = $('#channeldesc').val();

    $.ajax({

        url: '/channels/Updatechannelfields',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "id": $("#chanid").val(),

            "sections": JSON.stringify({ sections }),
            "deletesections": JSON.stringify({ deletesecion }),
            "deletefields": JSON.stringify({ deletefields }),
            "deleteoptions": JSON.stringify({ deleteoption }),
            "fiedlvalue": JSON.stringify({ fiedlvalue }),

            csrf: $("input[name='csrf']").val()
        },
        success: function (data) {

            if (data == true) {

                $('#ChannelEditmodal').modal('hide')

                // $('.avaliable-dropdown-item.active').click()
                $('.avaliable-dropdown-item').each(function () {
                    if ($(this).attr('id') === $('#chanid').val()) {
                        $(this).click();
                    }
                });


                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Toast.Channelupt + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }
            else {

                setCookie("Alert-msg", "Internal Server Error")

            }
            // window.location.href = "/channels/"

        }
    })





})


// $(document).on('click', '#Check2', function () {


//     if ($(this).is(':checked')) {

//         $(this).prop('checked', true);

//     } else {

//         $(this).prop('checked', false);

//     }

//     var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

//     var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

//     const index = fiedlvalue.findIndex(obj => {

//         return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

//     });

//     console.log(index, $(this).is(':checked'));

//     if (fiedlvalue.length == 0 || index < 0) {

//         const obj = CreatePropertiesObjec()

//         obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

//         fiedlvalue.push(obj)

//     } else {

//         fiedlvalue[index].Mandatory = $(this).is(':checked') == true ? 1 : 0;

//     }

//     console.log(fiedlvalue[index], "mandatory")

// })




$(document).on('click', ".propdrop", function () {


    var value = $(this).parents('.dropdown-menu').siblings('.fidvalinput').val()

    if (value != "") {

        $(this).parents('.dropdown-menu').siblings('label.error').hide();

        $(this).parents('.input-group').removeClass('input-group-error');


    }


})


function bindchanneldata() {


    $('.getvalue').each(function () {

        var id = $(this).find('.input-label').attr('data-id')

        const changedataindex = newchanneldata.findIndex(obj => {

            return obj.fid == id

        })



        if (changedataindex >= 0) {

            $(this).find('.input-label').siblings('.ig-row').children('input').val(newchanneldata[changedataindex].value)
            $(this).find('.input-label').siblings('.ig-row').children('textarea').val(newchanneldata[changedataindex].value)
            $(this).find('.input-label').siblings('.atag' + newchanneldata[changedataindex].fid).text(newchanneldata[changedataindex].value)
            $("#youtubeurl").val("")
        }

    })

    $('#searchdropdownrole').val(membername)
    if ($('#searchdropdownrole').val() != "") {

        $('.closemember').show()
    } else {

        $('.closemember').hide()
    }


}
function FieldValidation1() {

    console.log(fiedlvalue, "validation")

    // if (fiedlvalue.length == 0) {

    //   notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span> ` + languagedata?.Channell?.fieldselcterr + `</span></div>`;
    //   $(notify_content).insertBefore("#ChannelEditmodal .modal-content");
    //   setTimeout(function () {
    //     $('.toast-msg').fadeOut('slow', function () {
    //       $(this).remove();
    //     });
    //   }, 5000); // 5000 milliseconds = 5 seconds

    //   return false

    // }

    if (fiedlvalue != null) {

        for (let x of fiedlvalue) {

            if (x['FieldName'] == "") {

                $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                $('#fieldname-error').show();

                $('#fieldname-error').parents('.input-group').addClass('input-group-error')


                return false

            }

            if (x["MasterFieldId"] == 3) {

                if (x['Url'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    $('#url-error').show();

                    $('#url-error').parents('.input-group').addClass('input-group-error')


                    return false

                }



            } else if (x["MasterFieldId"] == 4) {

                var flg = true
                var flg2 = true

                if (x['DateFormat'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    $('#date-error').show();

                    $('#date-error').parent('.input-group').addClass('input-group-error')


                    flg = false

                }

                if (x['TimeFormat'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    console.log("timeformat empty");

                    $('#time-error').show();

                    $('#time-error').parents('.input-group').addClass('input-group-error')

                    flg2 = false
                }

                if (!flg || !flg2) {

                    console.log("11");

                    return false

                }


            } else if (x["MasterFieldId"] == 5 || x["MasterFieldId"] == 9 || x["MasterFieldId"] == 10) {

                if (x['OptionValue'].length == 0) {

                    $('#select-error').show();

                    $('#select-error').parents('.input-group').addClass('input-group-error')

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    console.log("options empty");

                    return false

                }

            } else if (x["MasterFieldId"] == 6) {

                if (x['DateFormat'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    $('#date-error').parents('.input-group').addClass('input-group-error')

                    $('#date-error').show();

                    console.log("dateformat empty");

                    return false

                }

            }

        }

    }
    return true
}

// search key value empty in dropdown search

$(document).on("click", "#triggerId", function () {



    var keyword = $("#searchdropdownrole").val()
    $(".memberlistdiv  button").each(function (index, element) {


        if (keyword != "") {
            $("#searchdropdownrole").val("")
            $(".memberlistdiv").empty()
            $("#nodatafounddesign").hide()
            $(element).show()
        } else {

            $("#searchdropdownrole").val()
            $("#nodatafounddesign").hide()
            $(element).show()

        }
    })

})

// $(document).on('focus', '.dateinfo', function () {

//     var dateformat = $(this).attr('data-date');

//     // var timeformat = $(this).attr('data-time');

//     if (dateformat != "") {

//         $('#' + $(this).attr('id')).bootstrapMaterialDatePicker({ setMaxDate: moment(), weekStart: 0, shortTime: false, time: false, format: dateformat, maxDate: new Date() });

//     }

// })

// $(document).on('focus', '.timeinfo', function () {


//     // var timeformat = $(this).attr('data-time');

//     $('#' + $(this).attr('id')).bootstrapMaterialDatePicker({ setMaxDate: moment(), weekStart: 0, shortTime: false, time: true, format: "HH:mm:ss", maxDate: new Date(), date: false });

// })

$(document).on("keyup", "#author", function () {

    // dropdown filter input box search
    $("#author").keyup(function () {

        var keyword = $(this).val().trim().toLowerCase()


        if (keyword != "") {

            $("#cn-user").show()

            fetch(`/channel/userdetails/?keyword=${keyword}`, {
                method: "GET"
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    // Handle the data received from the API

                    $('.userlistdiv').empty();

                    if (data == null) {

                        $(".nodata-userlistdiv").show()

                    }
                    if (data != null) {

                        $(".nodata-userlistdiv").hide()

                        $(".userlistdiv").show()

                        for (let x of data) {

                            $('.userlistdiv').append(` <button type="button" id="optionsid2" class="dropdown-item" data-id=${x.Id} >${x.FirstName}</button>`)

                        }
                    }

                })
                .catch(error => {
                    console.error('There was a problem with your fetch operation:', error);
                });
        } else {

            $(".userlistdiv").hide()

            $("#cn-user").hide()

            $(".nodata-userlistdiv").hide()
        }
    })

})

$(document).on('click', '#optionsid2', function () {


    test = $(this).text()

    $('#author').val(test)

    $(".userlistdiv").hide()

    $("#cn-user").show()

})
$(document).on('click', '#cn-user', function () {

    $('#author').val("")

    $("#cn-user").hide()

})

$(document).on('click', '.closemember', function () {

    $('#searchdropdownrole').val('')
    $('.memberlistdiv').hide()
    $(this).hide()
    $('.fieldid1').val('')
})

// list video url list 
$(document).on("click", "#videourl", function () {

    url = $("#youtubeurl").val()

    values = $("#urldatas").val()

    if (youtubelinks.indexOf(url) == -1) {

        youtubelinks.push(url)
        if (url != "") {
            $(".video-url-list").append(` <li  id="urlvalues">
            <p id="youtubevalue">` + url + ` </p>
         <a href="javascript:void(0);" data-id=`+ url + ` id="urldeletebtn"> <img src="/public/img/remove-folder.svg" alt="remove"></a>
          </li>`)


        }
        if (values != "") {
            joindata = values + "," + url
        } else {
            joindata = url
        }

        $("#urldatas").val(joindata)

        $("#youtubeurl").val("")
    }


})
// Delete video link

$(document).on("click", "#urldeletebtn", function () {

    val = $(this).siblings("#youtubevalue").text().trim()

    $(this).parents("li").remove()

    values = $("#urldatas").val()

    finalyoutubelinks = values.split(",")

    finalyoutubelinks.forEach((link, index) => {

        if (link == val) {

            finalyoutubelinks.splice(index, 1)
        }

    })

    $("#urldatas").val(finalyoutubelinks)
})

$(document).on("click", "#enmediamodalclose", function () {

    $("#addnewimageentriesModal").modal('hide')
    $("#drivelistdata").empty()

})

// set media gallery path
$(document).on("click", ".mediagallery-folder", function () {
    var folderpath = $(this).siblings(".forsearch").text()
    $(".uploadFolders").removeClass("add-galleryfolder")
    $(".uploadFolders").addClass("removeFolder")
    $("#fl-img").attr("src", "/public/img/folder-icon.svg")

    $(".fl-pathfld").val(folderpath)
    var parts = folderpath.split('/');
    parts = parts.filter(function (part) {
        return part.length > 0;
    });
    var lastIndex = parts[parts.length - 1];
    $("#fl-path").text(lastIndex)
    $("#fl-remove").show()
    $("#addnewimageentriesModal").modal('hide')
    $("#drivelistdata").empty()

})

// set media gallery path
$(document).on("click", "#remove-ga-folder", function () {
    $(".uploadFolders").removeClass("removeFolder")
    $(".uploadFolders").addClass("add-galleryfolder")
    $("#fl-img").attr("src", "/public/img/uploadFolder.svg")
    $("#fl-path").text("Upload Folder")
    $(".fl-pathfld").val("")
    $("#fl-remove").hide()
})



//preview page function//

$(document).on('click', '#previewbtn', function () {


    entryid = $(this).attr('data-id')

    $.ajax({
        url: "/channel/previewdetails/" + entryid,
        type: "GET",
        dataType: "json",
        data: { "id": entryid, csrf: $("input[name='csrf']").val() },
        success: function (result) {



            // $('#previewimg').attr("src", result.CoverImage)
            // $("#previewtitle").text(result.Title)
            $('#previewdesc').html(result.Description)
        }
    })
})

// entry list tab 
// $(document).on('click', '.pb-tab', function () {
//     $(this).addClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").removeClass("text-[#717171]")
//     $(".up-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".da-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".publish-tab").removeClass("hidden")
//     $(".unpublish-tab").addClass("hidden")
//     $(".draft-tab").addClass("hidden")
// })
// $(document).on('click','.up-tab', function () {
//     $(this).addClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").removeClass("text-[#717171]")
//     $(".pb-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".da-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".publish-tab").addClass("hidden")
//     $(".unpublish-tab").removeClass("hidden")
//     $(".draft-tab").addClass("hidden")
// })

// $(document).on('click','.da-tab', function () {    
//     $(this).addClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").removeClass("text-[#717171]")
//     $(".pb-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".up-tab").removeClass("after:inline-block after:w-full after:h-[2px] after:bg-[#262626] after:rounded-t-[18px] after:absolute after:bottom-0 after:left-0 text-[#262626]").addClass("text-[#717171]")
//     $(".publish-tab").addClass("hidden")
//     $(".unpublish-tab").addClass("hidden")
//     $(".draft-tab").removeClass("hidden")
// })


// Copy link function
$(document).on("click", '.copyButton', function () {

    var dataValue = $(this).attr('data-id');

    // Copy the data value to the clipboard
    navigator.clipboard.writeText(dataValue).then(function () {
        console.log('Data attribute value copied to clipboard: ' + dataValue);
        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex  w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Link Copied</p ></div ></div ></li></ul> `;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    }, function (err) {
        console.error('Could not copy text: ', err);
    });
});

function EntryStatus(id) {
    $('#Status' + id).on('change', function () {
        console.log("printf");
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#Status' + id).val();

    console.log("check", isactive, id)

    $.ajax({
        url: '/channel/entryisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > Entry Status Updated Successfully</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            } else {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }
        }
    });
}

$(document).on("click", ".selectall", function () {

    $(this).removeClass('selectall')

    $('#deselectid').addClass('deselectid')

    $(this).text(languagedata.deselectall)

    $('.selectcheckbox').prop('checked', true);

    $('.selectcheckbox').each(function () {

        entryid = $(this).attr('data-id')

        var status = $('.pb-tab').text()

        var exists = selectedcheckboxarr.some(function (item) {
            return item.entryid === entryid;
        });

        if (!exists) {
            selectedcheckboxarr.push({ "entryid": entryid, "status": status });
        }
    });




    if (selectedcheckboxarr.length != 0) {

        $('.selected-numbers').removeClass("hidden")

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus
        var img;

        if (selectedcheckboxarr[0].status === "Unpublished") {

            setstatus = "Publish";

            img = "/public/img/publish.png";

        } else if (selectedcheckboxarr[0].status === "Published") {

            setstatus = "Unpublish";

            img = "/public/img/unpublish-select.svg";

        }
        else if (selectedcheckboxarr[0].status === "Draft") {

            htmlContent = '';

        }


        if (allSame) {

            console.log("checksame")

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';


        } else {

            htmlContent = '';

        }

        console.log(htmlContent, "htmlcontent")

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)



        if (!allSame || selectedcheckboxarr[0].status === "Draft") {

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171]')

            $('#unbulishslt').html("")
        } else {

            $('#seleccheckboxdelete').addClass('border-r border-[#717171]')
        }

    } else {

        $('.selected-numbers').addClass("hidden")
    }
})


$(".sortable").sortable({

    update: function (event, ui) {

        newOrder = $(this).sortable('toArray', { attribute: 'data-id' })

        var pageno = $(".pagno").attr("data-page")

        console.log(pageno, "pageno")

        $.ajax({
            url: '/channel/reorder',
            method: 'POST',
            dataType: 'application/json',
            data: { "neworder": newOrder, csrf: $("input[name='csrf']").val(), pageno: pageno },
            success: function (response) {
                console.log("resposne")
            },

        });
    }
});

$(document).on('click', '#memgrp-slctall', function () {

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


    })
})

//Make it Private click function//
$(document).on('click', '#makeprivate', function () {

    pentryid = $(this).attr('data-id')

    $.ajax({
        url: "/channel/previewdetails/" + pentryid,
        type: "GET",
        dataType: "json",
        data: { "id": pentryid, csrf: $("input[name='csrf']").val() },
        success: function (result) {

            if (result.MembergroupId != "") {

                if (result.MembergroupId.includes(",")) {

                    access_granted_memgrps = result.MembergroupId.split(",")
                } else {


                    access_granted_memgrps.push(result.MembergroupId)
                }

                if (access_granted_memgrps != "") {

                    for (x of access_granted_memgrps) {

                        $('.memgrp-chkboxes').each(function () {

                            if ($(this).attr("data-id") == x) {

                                $(this).prop('checked', true)
                            }
                        })
                    }

                }

                $('.memgrp-chkboxes').each(function () {

                    if ($(this).prop('checked')) {

                        $('#memgrp-slctall').prop('checked', $('.memgrp-chkboxes:checked').length === $('.memgrp-chkboxes').length);
                    } else {

                        $('#memgrp-slctall').prop('checked', false);
                    }
                })
            }


        }
    })

})
//Update membergroupid in Entry
$(document).on('click', '#permissionupdate', function () {


    $.ajax({
        url: "/channel/updatepermissionmembergroupid",
        type: "POST",
        dataType: "json",
        data: { "entryid": pentryid, csrf: $("input[name='csrf']").val(), "memgrpids": access_granted_memgrps },
        success: function (result) {

            console.log("resultt", result)
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > Entry Updated Successfully</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            } else {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }

        }
    })



})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    var value = $(".entryclosebutton").val()
    console.log("value:", value);
    if (value == 1) {
        window.location.href = "/channel/entrylist"
    } else if (value == 2) {
        window.location.href = "/channel/unpublishentries"
    } else if (value == 3) {
        window.location.href = "/channel/draftentries"
    }
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")
        } else {
            $(".Closebtn").addClass("hidden")
            $(".srchBtn-togg").removeClass("pointer-events-none")
        }
    });
})

