var languagedata
let ckeditor1;
var categoryIds = [];
var seodetails = {};
var channeldata = []
checkboxarr = []
var cateogryarr = []
selctcategory = []
var Channelid
var flagg = false
var result1 = []
var newchanneldata =[]


console.log(channeldata, 'channeldata')
// Return to entries list function //
$(".searchentry").click(function () {

    id = $(this).attr('id')

    if ($(this).is(':checked')) {

        window.location.href = "/channel/entrylist/" + id
    } else {
        window.location.href = "/channel/entrylist/"
    }

})
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    CKEDITORS()

    if ($('#statusid').val() == '') {
        $('#triggerId').val('Status');
    }

    if ($('#statusid').val() == 'Draft') {
        $('#triggerId').val('Draft');
    }
    if ($('#statusid').val() == 'Published') {
        $('#triggerId').val('Published');
    }
    if ($('#statusid').val() == 'Unpublished') {
        $('#triggerId').val('Unpublished');
    }


    var url = window.location.href;

    console.log("url", url)

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    // if (url.includes('/editentry')) {
        var urlvalue = url.split('?')[0]

        var id = urlvalue.split('/').pop()

        var array = url.split('/');

        var channelname = array[array.length - 2];
        console.log("id", id, channelname);

    // }


    // var id = array[array.length - 1];
    var editurl

    if (pageno == null) {
        editurl = "/channel/editentrydetails/" + channelname + "/" + id
    } else {
        editurl = "/channel/editentrydetails/" + channelname + "/" + id + "?page=" + pageno
    }


    if (/^\d+$/.test(id)) {
        $.ajax({
            url: editurl,
            type: "GET",
            dataType: "json",
            data: { "id": id, csrf: $("input[name='csrf']").val() },
            success: function (result) {
                console.log(result, "result")

                result1 = result

                console.log("result1", result1)

                Channelid = result.Entries.ChannelId

                categoryIds = result.Entries.CategoriesId.split(",")

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

                $('#titleid').val(result.Entries.Title)
                // $('#text').val(result.Description)
                $('#titleid').css('height', '100%')
                ckeditor1.setData(result.Entries.Description)
                $('#chanid').val(result.Entries.ChannelId)
                $('#metatitle').val(result.Entries.MetaTitle)
                $('#metadesc').val(result.Entries.MetaDescription)
                $('#metakey').val(result.Entries.Keyword)
                $('#metaslug').val(result.Entries.Slug)
                $('#seosavebtn').text(languagedata?.update)
                $('#categorysave').text(languagedata?.update)
                // $('#configbtn').attr('href', '/channels/editchannel/' + result.Entries.ChannelId)

                seodetails = { title: result.Entries.MetaTitle, desc: result.Entries.MetaDescription, keyword: result.Entries.Keyword, slug: result.Entries.Slug }
                console.log("seodetails", seodetails)

                //dropdown item click//
                var dropdownItems = document.getElementsByClassName('avaliable-dropdown-item');
                for (var i = 0; i < dropdownItems.length; i++) {
                    if (dropdownItems[i].id == result.Entries.ChannelId) {
                        console.log('checkggg');

                        var channelname = dropdownItems[i].querySelector('.para').textContent;
                        var count = dropdownItems[i].querySelector('.para-extralight').textContent;
                        var itemId = dropdownItems[i].getAttribute('id');
                        console.log("name", channelname, count, itemId);

                        $('#chanid').val(itemId)
                        $('#blogn').text(channelname).css('color', 'black')
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
                $('#channelfieldsave').text(languagedata?.update)

                if (result.Section != null) {
                    result.Section.forEach(function (section) {

                        $('.secid').append(`<div style="margin-bottom:10px" id="${section.SectionId}"><h6 style="margin-bottom:10px" class="sechead">${section.SectionName}</h6></div>`)

                    })
                    if (result.FieldValue != null) {
                        result.FieldValue.forEach(function (fieldValue) {
                            if (fieldValue.MasterFieldId == 4) {
                                $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text" class="date" data-id="" data-date="${fieldValue.DateFormat}" data-time="${fieldValue.TimeFormat}" placeholder="` + languagedata?.Channell?.seldate + ` "  id="date${fieldValue.FieldId}"   value=""/>
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                  </div>
                `);
                            }

                            if (fieldValue.MasterFieldId == 5) {
                                $(`#${fieldValue.SectionId}`).append(` 
            <div class="input-group  user-drop-down getvalue " id="dropoption" data-id="${fieldValue.Mandatory}" style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle" type="button" id="triggerId${fieldValue.FieldId}" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        Option
                                        </a>
                                        <input type="hidden" data-id="" class="fieldid" id="${fieldValue.FieldId}"   value=""  >
        
                                        <div class="dropdown-menu"id="drop${fieldValue.FieldId}" aria-labelledby="triggerId">
                
                                            
                                        </div>
                                        <label class="error manerr" id="opterr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
        
                                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 6) {
                                $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date" data-id="" data-date="${fieldValue.DateFormat}" id="date${fieldValue.FieldId}"  placeholder="` + languagedata?.Channell?.seldate + `" />
                    
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                  </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 9) {
                                $(`#${fieldValue.SectionId}`).append(` 

                <div class="input-group getvalue" id="radiogrb" data-id="${fieldValue.Mandatory}">
                <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
                <div class="ig-row">
                <input type="hidden" data-id="" class="radioval" id="${fieldValue.FieldId}" value="">
                <div class="button-col entry-radio-row flexx" id="radio${fieldValue.FieldId}">
              
            </div>
            <label class="error manerr" id="radioerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
            </div>
            </div>
            
              
            
                `);
                            }
                            if (fieldValue.MasterFieldId == 8) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" >
                <label  class="input-label"  data-id="${fieldValue.FieldId}" id="textarealabel" for="">${fieldValue.FieldName}</label>
             
                <div class="ig-row">
                  <textarea placeholder="`+ languagedata?.Channell?.pltexthere + `" data-id="" id="${fieldValue.FieldId}" ></textarea>
                </div>
                <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
              </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 2) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata?.Channell?.pltexthere + `" id="${fieldValue.FieldId}"   value="" />
                      </div>
                      <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 7) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata?.Channell?.pltexthere + `" id="${fieldValue.FieldId}"   value=""   />
                      </div>
                      <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                            }

                            if (fieldValue.MasterFieldId == 10) {
                                $(`#${fieldValue.SectionId}`).append(`
                <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id="${fieldValue.Mandatory}"  style="margin-bottom: 20px">
         
                                         <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
         
                                         <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                             aria-haspopup="true" aria-expanded="false">
                                         Option
                                         </a>
                                         <input type="hidden" data-id="" class="checkfieldid" id="${fieldValue.FieldId}"  value=""  >
                                         <div class="dropdown-menu additional-entry-drop" id="check${fieldValue.FieldId}" aria-labelledby="triggerId">
                 
                                             
                                         </div>
                                         <label class="error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
         
                                     </div>
                `)


                            }

                            if (result.Entries.TblChannelEntryField != null) {
                                result.Entries.TblChannelEntryField.forEach(function (field) {

                                    $(`#${field.FieldId}`).val(field.FieldValue)

                                    $(`#date${field.FieldId}`).attr('data-id', field.Id)

                                    $(`#date${field.FieldId}`).val(field.FieldValue)

                                    $(`#${field.FieldId}`).attr('data-id', field.Id)

                                    $(`#triggerId${field.FieldId}`).text(field.FieldValue)


                                })
                            }
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

                $('.getvalue').each(function () {
                    obj = {}
                    obj = {
                        "name": $(this).children('.input-label').text().trim(),
                        "fid": $(this).children('.input-label').attr('data-id'),
                        "value": $(this).find('input').val() || $(this).find('textarea').val() || $(this).find('.fieldid').val(),
                        "fieldid": $(this).find('input').attr('data-id') || $(this).find('textarea').attr('data-id') || $(this).find('.fieldid').attr('data-id')

                    }
                    console.log("object", obj.value)

                    if ((obj.fieldid == undefined) || (obj.fieldid == '')) {
                        obj.fieldid = '0'

                    }
                    if ((obj.value != '') && (obj.value != undefined)) {
                        channeldata.push(obj)
                        newchanneldata.push(obj)
                    }


                    console.log("channeldata", channeldata)
                })
                var radval = $('.radioval').val()
                console.log(radval, "radiovalue")
                $('.radiodiv').each(function () {
                    if ($(this).find('label').text() == radval) {
                        console.log("uuuu")
                        $(this).find('input').prop('checked', true)
                    }
                })

                var cheval = $('.checkfieldid').val()
                var values = cheval.split(','); // Split the string by comma
                for (var i = 0; i < values.length; i++) {
                    checkboxarr.push(values[i]); // Push each value to the array
                }
                console.log(checkboxarr, "array");
                console.log("fieldval", cheval)
                $('.checkboxdiv').each(function (index) {

                    if (cheval.includes($(this).find('label').text())) {
                        $(this).find('input').prop('checked', true);
                    }
                })

            }
        })

    }
})


$(document).ready(function () {

    console.log("checksdfdf")
    if ($('.channeltitle1').text() == '') {

        console.log("$sdfdfdf")

        $('.channeltitle1').css('display', 'none')
    }
})
$(document).on('keyup', '#searchkey', function () {

    if ($('.search').val() === "") {
        console.log("check")
        window.location.href = "/channel/entrylist/"

    }

})


$(document).on('click', '#filtercancel', function () {

    $('#triggerId').val('');
    var url = window.location.href;
    var cleanUrl = url.split("?")[0];
    window.location.href = cleanUrl
});

//**Dropdown role//
$(document).on('click', '.statuss', function () {

    drop = $(this).text()
    $('#triggerId').val(drop)
    $('#statusid').val(drop)
    $('.filter-dropdown-menu').addClass('show').css({
        position: 'absolute',
        inset: '0px 0px auto auto',
        margin: '0px',
        transform: 'translate(0px, 34px)'
    });


})

$(document).on('click', '.deleteentry', function () {

    $('.filter-dropdown-menu ').hide()

    var entryId = $(this).attr("data-id")

    var channame = $(this).attr("data-name")

    console.log("entryid", entryId)

    var del = $(this).closest("tr");

    $('.delname').text(del.find('td:eq(2)').text())

    $('.deltitle').text(languagedata?.Channell?.delentrytitle)

    $('.deldesc').text(languagedata?.Channell?.delentrycontent)

    $('#delid').attr('data-id', $(this).attr('data-id'))

    $('#delid').text(languagedata?.yes)

    $('#delcancel').text(languagedata?.no)

    var pageno = $(this).attr("data-page")

    if (pageno == "") {
        $('#delid').parent('#delete').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame);

    } else {
        $('#delid').parent('#delete').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame + "&page=" + pageno);

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

    $('#content').text(languagedata.Channell.publishcontent);
    $('.delname').text(channelname)
    $('.deltitle').text(languagedata.Channell.publishtitle)
    $("#delid").text(languagedata.Channell.publish)
    $("#delcancel").text(languagedata.cancel)

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


})

$(document).on("click", "#delid", function () {
    if (window.location.href.includes('/channel/entrylist/')) {

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
                    // setCookie("get-toast", languagedata?.Toast?.entryupdatenotify)
                    // setCookie("Alert-msg", "success", 1)
                    window.location.reload()



                } else {

                    setCookie("Alert-msg", languagedata?.Toast?.internalserverr)



                }
            }
        })
    }
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

    console.log("url", url)

    var pageurl = window.location.search
    const urlpar = new URLSearchParams(pageurl)
    pageno = urlpar.get('page')



    if (url.includes('editentry')) {

        var urlvalue = url.split('?')[0]
        var eid = urlvalue.split('/').pop()
        // var eid = url.split('/').pop();

        console.log("id", eid);
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

    var channelname = $('#blogn').text()

    console.log("channelname", channelname)

    //  text =$('#text').val()

    console.log(title, text, "ggg")

    blog = $('#encount').val()

    blogcount = $('#encount').attr('data-id')

    Removeslashcategory()

    if (title === '' && data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descerr + ' </span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descimgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.imgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds



    } else if (data === '') {

        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descriptionentry + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds


    }

    if (title !== '' && data !== '' && img !== '' && entryId == '') {

        $('#channelslct-error').show()
    }

    if (title !== '' && data !== '' && img !== '' && entryId !== '') {
        var flag = channelfieldvalidation()

        channelfieldkeyup()

        if (flag == false) {
            $('#channelModal').modal('show');
            return
        }

        var flag1 = categoryvalidation()
    }
    if (title && data != '' && img != '' && (flag == true) && (flag1 == true)) {
        $.ajax({
            url: drafturl,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "text": data, "status": 0, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), "page": pageno, csrf: $("input[name='csrf']").val() },
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
                    // window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id
                    if (pageno == null) {
                        window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id

                    } else {
                        window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id + "?page=" + pageno

                    }
                }, 2000);


            }
        })
    }
})
// return flagg
// }


$('#previewbtn').click(function () {

    img = $('#spimagehide').attr('src')

    title = $('#titleid').val()

    var data = ckeditor1.getData();

    $('#preview-img').attr('src', img)

    $('.preview-head').text(title)

    $('.heading-four').html(data)

})


$('.avaliable-dropdown-item').click(function () {

    var id = $(this).find('input').data('id');

    var channelname = $(this).find('.para').text();

    var count = $(this).find('.para-extralight').text();

    $('.blogDropdown1').text(channelname).css('color', 'black');

    $('.blogDropdown1').attr('data-bs-original-title', channelname)

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

    if (url.includes('editentry')) {
        var urlvalue = url.split('?')[0]
        var eid = urlvalue.split('/').pop()

    }
    var data = ckeditor1.getData();

    var entryId = $('#chanid').val()

    var channelname = $('#blogn').text()

    var img = $('#spimagehide').attr('src')

    var title = $('#titleid').val()

    console.log("daata", img, data, title)

    text = $('#text').val()

    blog = $('#encount').val()

    blogcount = $('#encount').attr('data-id')

    Removeslashcategory()


    if (title === '' && data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descimgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if (img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.imgerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    } else if (data === '') {

        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata?.Channell?.descriptionentry + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds


    }

    if (title !== '' && data !== '' && img !== '' && entryId == '') {

        $('#channelslct-error').show()
    }
    if (title !== '' && data !== '' && img !== '' && entryId !== '') {
        console.log("kkkkk")
        var flag = channelfieldvalidation()

        channelfieldkeyup()

        if (flag == false) {
            $('#channelModal').modal('show');
            return
        }

        var flag1 = categoryvalidation()


    }

    if (title != '' && data != '' && img != '' && (flag == true) && (flag1 == true)) {

        $.ajax({
            url: "/channel/publishentry/" + eid,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "status": 1, "text": data, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val() },
            success: function (result) {

                window.location.href = homeurl;
            }
        })

    }


})

//EDIT FUNCTION IN ENTRY PAGE//
$(document).on('click', '#editbtn', function () {

    $('.filter-dropdown-menu ').hide()

})


function CKEDITORS() {

    var url = $('#urlid').val();

    console.log("urlasas", url)

    CKEDITOR.ClassicEditor.create(document.getElementById("text"), {
        toolbar: {
            items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'horizontalLine', 'link', 'code'],
            shouldNotGroupWhenFull: true
        },
        list: {
            properties: {
                styles: true,
                startIndex: true,
                reversed: true
            }
        },
        heading: {
            options: [
                { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
                { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
                { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
                { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
                { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
                { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
                { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
            ]
        },
        placeholder: languagedata?.Channell?.plckeditor,
        fontFamily: {
            options: [
                'default',
                'Arial, Helvetica, sans-serif',
                'Courier New, Courier, monospace',
                'Georgia, serif',
                'Lucida Sans Unicode, Lucida Grande, sans-serif',
                'Tahoma, Geneva, sans-serif',
                'Times New Roman, Times, serif',
                'Trebuchet MS, Helvetica, sans-serif',
                'Verdana, Geneva, sans-serif'
            ],
            supportAllValues: true
        },
        fontSize: {
            options: [10, 12, 14, 'default', 18, 20, 22],
            supportAllValues: true
        },
        htmlSupport: {
            allow: [
                {
                    name: /.*/,
                    attributes: true,
                    classes: true,
                    styles: true
                }
            ]
        },
        htmlEmbed: {
            showPreviews: true
        },
        link: {
            decorators: {
                addTargetToExternalLinks: true,
                defaultProtocol: 'https://',
                toggleDownloadable: {
                    mode: 'manual',
                    label: 'Downloadable',
                    attributes: {
                        download: 'file'
                    }
                }
            }
        },
        mention: {
            feeds: [
                {
                    marker: '@',
                    feed: [
                        '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
                        '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
                        '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
                        '@sugar', '@sweet', '@topping', '@wafer'
                    ],
                    minimumCharacters: 1
                }
            ]
        },
        removePlugins: [
            'ExportPdf',
            'ExportWord',
            'CKBox',
            'CKFinder',
            'EasyImage',
            'RealTimeCollaborativeComments',
            'RealTimeCollaborativeTrackChanges',
            'RealTimeCollaborativeRevisionHistory',
            'PresenceList',
            'Comments',
            'TrackChanges',
            'TrackChangesData',
            'RevisionHistory',
            'Pagination',
            'WProofreader',
            'MathType',
            'SlashCommand',
            'Template',
            'DocumentOutline',
            'FormatPainter',
            'TableOfContents',
            'PasteFromOfficeEnhanced'
        ]
    }).then(ckeditor => {
        ckeditor1 = ckeditor
        ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
            return {
                upload: () => {
                    return loader.file.then(file => {
                        return new Promise((resolve, reject) => {
                            const formData = new FormData();
                            formData.append('file', file);
                            formData.append('csrf', $("input[name='csrf']").val())
                            fetch(url + '/channel/imageupload', {
                                method: 'POST',
                                body: formData
                            })
                                .then(response => response.json())
                                .then(data => {
                                    if (data.success) {
                                        resolve({
                                            default: data.url // URL to the uploaded image

                                        });
                                    } else {
                                        reject(data.error);
                                    }
                                })
                                .catch(error => {
                                    reject(error);
                                });
                        });
                    });
                }
            };
        };
    })
        .catch(error => {
            console.error(error);
        });

}

//CATEGORY SAVE FUNCTION//

$(document).on('click', '#categorysave', function () {
    categoryIds = [];

    categoryvalidation()

    $('.category-unselect-btn').each(function () {
        var categoryId = $(this).attr('data-id');
        categoryIds.push(categoryId);
    });

    if (categoryIds.length !== 0) {



        // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata?.Channell?.categoryerr + '</span></div>';
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
                customLength: "*" + languagedata?.Channell?.metatitlemsg,
                titlemethod: "*Please Enter Meta tag title",
                titlespace: "* " + languagedata.spacergx
            },
            metadesc: {
                customLength: "*" + languagedata?.Channell?.metadescmsg,

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

        if (title !== '' || desc !== '' || keyword !== '' || slug !== '') {
            // seodetails.push({ title: title, desc: desc, keyword: keyword, slug: slug });
            seodetails = { title: title, desc: desc, keyword: keyword, slug: slug }
        }


        console.log("seodetails", seodetails)
        if (Object.keys(seodetails).length !== 0) {

            // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata?.Channell?.seoerr + '</span></div>';
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
        console.log("focusss")
        event.target.closest('.input-group').classList.add('input-group-focused');
    });

    input.addEventListener('blur', (event) => {
        event.target.closest('.input-group').classList.remove('input-group-focused');
    });
});

document.querySelectorAll('.secid input').forEach(input => {
    console.log("vvvvv")

    input.addEventListener('focus', (event) => {
        console.log("focusss")
        event.target.closest('.input-group').classList.add('input-group-focused');
    });

    input.addEventListener('blur', (event) => {
        event.target.closest('.input-group').classList.remove('input-group-focused');
    });
});

//**description focus function */
const Desci = document.getElementById('metadesc');


Desci.addEventListener('focus', () => {

    Desci.closest('.input-group').classList.add('input-group-focused');
});
Desci.addEventListener('blur', () => {
    Desci.closest('.input-group').classList.remove('input-group-focused');
});

$(document).on('click', '#seocancelbtn', function () {

    $('#metatitle').val('')
    $('#metadesc').val('')
    $('#metakey').val('')
    $('#metaslug').val('')
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

$(document).on('click', '.avaliable-dropdown-item ', function () {

    console.log("checkss")

    categoryIds = []

    seodetails = {};

    channeldata = []

    // $('.channeltitle1').show()

    channelid = $('#chanid').val()

    // $('#configbtn').attr('href', '/channels/editchannel/' + channelid)

    chanid = $(this).attr('id')

    if (channelid != '') {

        $('#channelslct-error').hide()
    }

    $('.avaliable-dropdown-item').removeClass('active')

    if (channelid == chanid) {

        $(this).addClass('active')
    }

    console.log("jiddd", channelid)

    $('.secid').empty();

    $.ajax({
        url: "/channel/channelfields/" + channelid,
        type: "GET",
        dataType: "json",
        data: { "id": channelid },
        success: function (result) {
            console.log(result, "resultddddd")
            cateogryarr = []
            //category section//

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

            if (result.Section != null) {
                result.Section.forEach(function (section) {

                    $('.secid').append(`<div style="margin-bottom:10px" id="${section.SectionId}"><h6 class="sechead" style="margin-bottom:10px">${section.SectionName}</h6></div>`)

                })
                if (result.FieldValue != null) {
                    result.FieldValue.forEach(function (fieldValue) {


                        if (fieldValue.MasterFieldId == 4) {
                            $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date date-start" name="todate" data-name="${fieldValue.FieldName}"  id="dateend-${fieldValue.FieldId}" data-date="${fieldValue.DateFormat}" data-time="${fieldValue.TimeFormat}" placeholder="${fieldValue.DateFormat}" value="" />
                    
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                  </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 5) {
                            $(`#${fieldValue.SectionId}`).append(` 
            
            <div class="input-group  user-drop-down getvalue" data-id="${fieldValue.Mandatory}"  id="dropoption" style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        Option
                                        </a>
                                        <div class="ig-row">
                                        <input type="hidden" class="fieldid" id="${fieldValue.FieldId}"  value=""  >
                                        </div>
                                        <label class="error manerr" id="opterr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
        
                                        <div class="dropdown-menu" id="drop${fieldValue.FieldId}" aria-labelledby="triggerId">
                
                                            
                                        </div>
                                      
                                    </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 6) {
                            $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" >
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date"  data-date="${fieldValue.DateFormat}" id="${fieldValue.FieldId}"  value=""  placeholder="${fieldValue.DateFormat}" />
                     
                    
                    </div>
                    <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                  </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 9) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" id="radiogrb" data-id="${fieldValue.Mandatory}" >
                <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
                <div class="ig-row">
                <input type="hidden" class="radioval" value="">
                <div class="button-col entry-radio-row flexx" id="radio${fieldValue.FieldId}">
              
            </div>
            <label class="error manerr" id="radioerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
            </div>
            </div>
            
               
                `);
                        }
                        if (fieldValue.MasterFieldId == 8) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" >
                <label  class="input-label"  data-id="${fieldValue.FieldId}" id="textarealabel" for="">${fieldValue.FieldName}</label>
             
                <div class="ig-row">
                  <textarea placeholder="`+ languagedata?.Channell?.pltexthere + `" id="chtextarea"></textarea>
                </div>
                <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
              </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 2) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}"  data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text"  placeholder="`+ languagedata?.Channell?.pltexthere + `" id="chtextbox" />
                      </div>
                      <label class="error manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>

                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 7) {
                            $(`#${fieldValue.SectionId}`).append(`  
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" placeholder="`+ languagedata?.Channell?.pltexthere + `"  />
                      </div>
                      <label class="error  manerr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 10) {
                            $(`#${fieldValue.SectionId}`).append(`
               <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id="${fieldValue.Mandatory}"  style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        Option
                                        </a>
                                      
                                        <input type="hidden" class="checkfieldid" id="${fieldValue.FieldId}"  value=""  >
                                      
                                        <label class="error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                                        <div class="dropdown-menu additional-entry-drop" id="check${fieldValue.FieldId}" aria-labelledby="triggerId">
                
                                            
                                        </div>
                                       
        
                                    </div>
               `)


                        }

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
                 <div class="radio radio-label" >
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

        }
    })

})

$(document).on('click', '#channelfieldsave', function () {

    // channeldata = [];

    var flag = channelfieldvalidation()

    channelfieldkeyup()

    $('.getvalue').each(function () {
        obj = {}
        obj = {
            "name": $(this).children('.input-label').text().trim(),
            "fid": $(this).children('.input-label').attr('data-id'),
            "value": $(this).find('input').val() || $(this).find('textarea').val() || $(this).find('.fieldid').val(),
            "fieldid": $(this).find('input').attr('data-id') || $(this).find('textarea').attr('data-id') || $(this).find('.fieldid').attr('data-id')

        }
        console.log("object", obj.value)

        if ((obj.fieldid == undefined) || (obj.fieldid == '')) {
            obj.fieldid = '0'

        }
        if ((obj.value != '') && (obj.value != undefined)) {
            channeldata.push(obj)
        }


        console.log("channeldata", channeldata)
    })

    if ((channeldata.length !== 0) && (flag == true)) {

        // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span> ' + languagedata?.Channell?.saveerr + ' </span></div>';
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

    if (flag == true) {
        $('#channelModal').modal('hide');
    }
})

$(document).on('click', '#addcategory', function () {


    channelid = $('#chanid').val()

    title = $('#titleid').val()
    if (title == '') {
        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata?.Channell?.titleerr + `</span></div>`;
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


    if (newchanneldata != '') {
        bindchanneldata()
    }


    channelid = $('#chanid').val()

    title = $('#titleid').val()

    if (title == '') {
        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata?.Channell?.titleerr + `</span></div>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    } else if ((title != '') && (channelid == '')) {

        // console.log("alertt")
        // notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata?.Channell?.channelselect+ `</span></div>`;
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

        notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>` + languagedata?.Channell?.titleerr + `</span></div>`;
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

    console.log("congiggg")
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

            if(data.FieldValue !=null){

                fiedlvalue = data.FieldValue

            }else{

                fiedlvalue =[]
            }


            if (data.SelectedCategory != null) {

                SelectedCategoryValue = data.SelectedCategory
            }

            var sectionstr = ``

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
if(data.FieldValue !=null){


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

                    console.log("draggggggg")
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

        console.log(charallowed, "yyyy")
        if ((($(this).find('input').val() === '') && ($(this).data('id') == 1)) || (($(this).find('textarea').val() === '') && ($(this).data('id') == 1))) {

            $(this).addClass('input-group-error');
            $(this).find('.manerr').show()
            flag = false
        }
        if (($(this).find('input,textarea').val() !== '') && charallowed != 0) {
            console.log("dfdfdfdff")
            if ($(this).find('input,textarea').val().length > charallowed) {
                console.log("gggg")
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

        console.log("yyyy")

        $('#dropoption').removeClass('input-group-error')
        $('#opterr').css('display', 'none');
    }

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

        console.log("sddddddd")
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

    console.log("necheck")
    if ($(this).is(':checked')) {

        option = $(this).next('label').text()
        console.log("option", option)

        checkboxarr.push(option)
        $('.checkfieldid').val(checkboxarr.join(','))

    }
    else {
        var option = $(this).next('label').text();
        console.log(option, "hhhhhh")
        checkboxarr.splice(checkboxarr.indexOf(option), 1);
        console.log("may", option, checkboxarr)
        $('.checkfieldid').val(checkboxarr.join(','));
    }


    if ($('.checkfieldid').val() != '') {

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

    console.log("focus")

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
  

     if(fiedlvalue ==''){

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

        console.log(fiedlvalue,"arrayfield")
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
                    console.log("dropdown")
                    if ($(this).attr('id') === $('#chanid').val()) {
                        $(this).click();
                    }
                });


                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata?.Toast?.Channelupt + '</span></div>';
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


$(document).on('click', '#Check2', function () {

    console.log("checked")

    if ($(this).is(':checked')) {

        $(this).prop('checked', true);

    } else {

        $(this).prop('checked', false);

    }

    var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

    var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

    const index = fiedlvalue.findIndex(obj => {

        return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

    });

    console.log(index, $(this).is(':checked'));

    if (fiedlvalue.length == 0 || index < 0) {

        const obj = CreatePropertiesObjec()

        obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

        fiedlvalue.push(obj)

    } else {

        fiedlvalue[index].Mandatory = $(this).is(':checked') == true ? 1 : 0;

    }

    console.log(fiedlvalue[index], "mandatory")

})




$(document).on('click', ".propdrop", function () {

    console.log("dateformatsssssssss")

    var value = $(this).parents('.dropdown-menu').siblings('.fidvalinput').val()

    if (value != "") {

        $(this).parents('.dropdown-menu').siblings('label.error').hide();

        $(this).parents('.input-group').removeClass('input-group-error');


    }


})


function bindchanneldata() {
    console.log(result1, "log")

    $('.getvalue').each(function () {

        var id = $(this).find('.input-label').attr('data-id')

        const changedataindex = newchanneldata.findIndex(obj => {

            return obj.fid == id

        })

        console.log(newchanneldata[changedataindex], newchanneldata);

        if (changedataindex >= 0) {


            $(this).find('.input-label').siblings('.ig-row').children('input').val(newchanneldata[changedataindex].value)
        }

    })

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

                console.log("fieldname empty");

                return false

            }

            if (x["MasterFieldId"] == 3) {

                if (x['Url'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    $('#url-error').show();

                    $('#url-error').parents('.input-group').addClass('input-group-error')

                    console.log("url empty");

                    return false

                }



            } else if (x["MasterFieldId"] == 4) {

                var flg = true
                var flg2 = true

                if (x['DateFormat'] == "") {

                    $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

                    $('#date-error').show();

                    $('#date-error').parent('.input-group').addClass('input-group-error')

                    console.log("dateformat empty");

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