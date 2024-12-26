
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
var spurtdata
var Channelcategories
var newcatarray = []
// logout dropdwon
$(document).on("click", "#lg-btn", function () {
    $("#log-btn").toggle("show")
})

$(document).ready(async function () {

    var gurl = window.location.href;

    if (gurl.includes('generatearticle')) {

        var newUrl = '/channel/newentry';

        window.history.pushState({}, '', newUrl);
    }

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    // You get innerHTML here    
    document.addEventListener('saveChange', function (event) {
        spurtdata = event.detail

        console.log("spurtdata", spurtdata.title, spurtdata.image);

        if ($("#savetype").val() == "draft") {
            var url = window.location.href;


            var pageurl = window.location.search
            const urlpar = new URLSearchParams(pageurl)
            pageno = urlpar.get('page')

            if (url.includes('editentry') || url.includes("editsentry")) {

                // var urlvalue = url.split('#')[0]
                // console.log("urlvalue",urlvalue);

                var eid = $("#eid").val()

            }
            console.log("pageno", pageno);

            if (pageno == null) {
                homeurl = "/channel/draftentries"
            } else {
                homeurl = "/channel/draftentries?page=" + pageno;
            }
            Getselectedcatagories()
            if (categoryIds.length == 0) {
                categoryIds.push($("#cat0").attr("data-id"));
            }
            channelname = $("#slchannel").val()
            if ($.trim(spurtdata.title) === "") {

                spurtdata.title = "Untitled Entries"
            }
            if (spurtdata.title != "") {
                if (containsHtmlElements(spurtdata.title)) {

                    spurtdata.title = "Untitled Entries"

                    console.log("The string contains HTML elements");
                } else {

                    var wordsArray = spurtdata.title.split(' ');

                    var first30Words = wordsArray.slice(0, 20);

                    spurtdata.title = first30Words.join(' ');

                    console.log("The string does not contain HTML elements");
                }
                Getseodetails()
                authername = $("#author").val()
                createtime = $("#cdtime").val()
                publishtime = $("#publishtime").val()
                readingtime = $("#readingtime").val()
                orderindex = $("#orderindex").val()
                // sortorder = $("#article").val()
                tagname = $("#tagname").val()
                extxt = $("#extxt").val()
                $.ajax({
                    url: "/channel/draftentry/" + eid,
                    type: "POST",
                    dataType: "json",
                    data: { "id": $("#slchannel").attr('data-id'), "cname": channelname, "image": spurtdata.image, "title": spurtdata.title, "status": 0, "text": spurtdata.html, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": "", "tagname": tagname, "extxt": extxt, "orderindex": orderindex },
                    success: function (result) {
                        window.location.href = homeurl;
                    }
                })
            }
        }

        if ($("#savetype").val() == "publish") {

            var url = window.location.href;
            var homeurl

            var pageurl = window.location.search
            const urlpar = new URLSearchParams(pageurl)
            pageno = urlpar.get('page')


            if (url.includes('editentry') || url.includes("editsentry")) {

                var eid = $("#eid").val()
            }

            if (pageno == null) {
                homeurl = "/channel/entrylist"
            } else {
                homeurl = "/channel/entrylist?page=" + pageno;
            }
            channelname = $("#slchannel").val()
            Getselectedcatagories()
            if (categoryIds.length == 0) {
                categoryIds.push($("#cat0").attr("data-id"));
            }
            if ($("#slchannel").attr("data-id") == "") {
                $("#sl-chn-error").show()
                $('.editor-tabs').removeClass('translate-x-[100%]');
                $('#editingArea').addClass('mr-[387px] w-full ');
            }

            if (categoryIds.length == 0) {
                $("#sl-cat-error").show()
                $('.editor-tabs').removeClass('translate-x-[100%]');
                $('#editingArea').addClass('mr-[387px] w-full ');

            }
            if ($.trim(spurtdata.title) === "" && $("#slchannel").attr("data-id") != "") {
                spurtdata.title = "Untitled Entries"
            }
            if (spurtdata.title != "" && $("#slchannel").attr("data-id") != "" && categoryIds.length != 0) {

                if (containsHtmlElements(spurtdata.title)) {

                    spurtdata.title = "Untitled Entries"

                    console.log("The string contains HTML elements");
                } else {

                    var wordsArray = spurtdata.title.split(' ');

                    var first30Words = wordsArray.slice(0, 20);

                    spurtdata.title = first30Words.join(' ');

                    console.log("The string does not contain HTML elements");
                }

                Getseodetails()

                authername = $("#author").val()
                createtime = $("#cdtime").val()
                publishtime = $("#publishtime").val()
                readingtime = $("#readingtime").val()
                // sortorder = $("#article").val()
                tagname = $("#tagname").val()
                extxt = $("#extxt").val()
                orderindex = $("#orderindex").val()

                $.ajax({
                    url: "/channel/publishentry/" + eid,
                    type: "POST",
                    dataType: "json",
                    data: { "id": $("#slchannel").attr('data-id'), "cname": channelname, "image": spurtdata.image, "title": spurtdata.title, "status": 1, "text": spurtdata.html, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": "", "tagname": tagname, "extxt": extxt, "orderindex": orderindex },
                    success: function (result) {
                        window.location.href = homeurl;
                    }
                })
            }
        }
    });

    if ($("#chnid").val() != "") {
        $(".select-chn").each(function () {
            if ($(this).attr("data-id") == $("#chnid").val()) {
                $(this).click()
            }
        })
    }
    var url = window.location.href;


    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    var urlvalue = url.split('?')[0]

    var id = urlvalue.split('/').pop()

    var array = url.split('/');

    var channelname = array[array.length - 2];

    var editurl


    if (channelname == "") {
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

                var currentDate = new Date();
                var formattedDate = new Date(currentDate.getTime() - currentDate.getTimezoneOffset() * 60000).toISOString().slice(0, 16);
                console.log("formattedDate", formattedDate);

                $('#metatitle').val(result.Entries.MetaTitle)
                $('#metadesc').val(result.Entries.MetaDescription)
                $('#metakey').val(result.Entries.Keyword)
                $('#metaslug').val(result.Entries.Slug)
                Channelid = result.Entries.ChannelId
                $("#tagname").val(result.Entries.Tags)
                $("#extxt").val(result.Entries.Excerpt)
                $("#orderindex").val(result.Entries.OrderIndex)
                if (result.Entries.Author != "") {
                    $("#author").val(result.Entries.Author)
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
                if (result.Entries.CategoriesId != "") {
                    $(".sl-list").removeClass("hidden")
                    newcatarray = result.Entries.CategoriesId.split(",")

                }

                for (const [index, y] of result.CategoryName.entries()) {

                    var div = ""

                    var id = ""

                    div += `<li class="cursor-pointer rounded-[4px] p-[8px_12px] border border-[#EDEDED] flex space-x-[6px] justify-between mb-[8px] last-of-type:mb-[0] sl-cat sl-cat` + index + `" data-id="` + index + `">
                    <div class="relative flex space-x-[6px] items-center mb-0 text-[14px] font-normal leading-none text-[#262626] tracking-[0.005em] line-clamp-1 flex-wrap">`

                    for (let x of y.Categories) {
                        div += `<span class="inline-flex items-center  text-[12px] font-normal leading-none text-[#555555]
            after:inline-block after:ml-[6px] after:bg-[url('/public/img/path-arrow.svg')] after:w-[8px] after:h-[8px] after:bg-no-repeat after:bg-[length:6px_8px] after:bg-center last-of-type:after:hidden" data-id ="` + x.Id + `">
            `+ x.Category + `</span>`

                        id = x.Id
                    }
                    div += `</div><a href="javascript:void(0);" class="bg-[#F7F7F5]  p-[3px] rounded-[3px] min-w-[16px] w-[16px] h-[16px] grid place-items-center hover:bg-[#F0F0F0] rsl-cat">
                    <img src="/public/img/remove-categories.svg" alt="remove">
                </a>
                <input type="hidden" data-id="`+ id + `" class="cat-val">
                </li>`
                    for (let x of newcatarray) {
                        if (x == id) {
                            $(".selected-cat").append(div)
                            $("#Check" + index).prop("checked", true)
                        }
                    }
                }
                if (result.Entries.TblChannelEntryField != null) {
                    result.Entries.TblChannelEntryField.forEach(function (field) {
                        $(`#${field.FieldId}`).val(field.FieldValue)
                        $(`#${field.FieldId}`).attr('data-id', field.Id)
                    })
                }
                var radval = $('.radioval').val()
                $('.radbtn').each(function () {
                    if ($(this).next('label').text() == radval) {
                        $(this).prop('checked', true)
                    }
                })

                var selectval = $('.fieldid').val()
                $('#selectval').text(selectval)
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
            }
        })

    }
    // Allow only five enter key 
    $('#metadesc').on('input', function () {

        let lines = $(this).val().split('\n').length;

        if (lines > 5) {

            let value = $(this).val();

            let linesArray = value.split('\n').slice(0, 5);

            $(this).val(linesArray.join('\n'));
        }
    });


    var $textarea = $('#metadesc');
    var $errorMessage = $('#metadesc-error');
    var maxLength = 250;

    $textarea.on('input', function () {
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessage.text(languagedata.Permission.descriptionchat);
        } else {
            // Clear error message if under the limit
            $errorMessage.text('');
        }

        if ($(this).val().startsWith(" ")) {
            $errorMessage.text(languagedata.Permission.space);
        } else {
            $errorMessage.text("");
        }
    });

    // meta title validation space validation

    $("#metatitle").on('input', function () {
        var metatilte = $(this).val();
        if (metatilte.startsWith(" ")) {

            $("#metatitle-error").text(languagedata.Permission.space);
        } else {
            $("#metatitle-error").text("");
        }

    })

})

// New code 
$(document).on("click", "#sl-chn", function () {
    $("#chn-list").toggleClass("show")
})

// Get channel fields
$(document).on('click', '.select-chn', function () {

    categoryIds = []

    seodetails = {};

    channeldata = []
    $(".sl-list").addClass("hidden")
    $(".selected-cat").empty()
    $("#chn-list").removeClass("show")

    chanid = $(this).attr('data-id')

    $("#chn-list").removeClass("show")
    $("#sl-chn-error").hide()
    $("#chn-name").text($(this).text()).addClass("text-bold-black")
    $("#chn-name").attr("title", $(this).text())
    $("#chn-name").attr("data-id", chanid)
    $("#slchannel").val($(this).text())
    $("#slchannel").attr("data-id", chanid)
    $.ajax({
        url: "/channel/channelfields/" + chanid,
        type: "GET",
        dataType: "json",
        data: { "id": chanid },
        success: function (result) {
            cateogryarr = []
            $(".getvalue").remove()
            //category section//
            txt = `<span class="text-[#262626]">${result.Categorycount}</span> Categories Available`
            $("#av-cat").html(txt)

            if (result.CategoryName !== null) {
                Channelcategories = result.CategoryName
                var div = ""
                for (const [index, y] of result.CategoryName.entries()) {

                    var categoriess = ""

                    var ids = ""

                    var id = ""

                    div +=
                        `
                <li class="p-[9px_24px] border-b border-[#EDEDED]">
                   <div class="chk-group chk-group-label">
                  <input type="checkbox" id="Check` + index + `" class="hidden peer selectcheckbox" data-id="` + index + `">
                                                      <label for="Check` + index + `" class="relative cursor-pointer flex space-x-[6px] items-center mb-0 text-[14px] font-normal leading-none text-[#262626] tracking-[0.005em]
                                before:bg-transparent before:w-[14px] before:h-[14px] before:inline-block before:relative before:align-middle before:cursor-pointer before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:mr-[6px] before:-webkit-appearance-none peer-checked:before:bg-[url('/public/img/checked-box.svg')]  
                                    ">`

                    for (let x of y.Categories) {

                        div += `<span class="inline-flex items-center space-x-[6px] text-[12px] font-normal leading-none text-[#555555]
                    after:inline-block after:ml-[6px] after:bg-[url('/public/img/path-arrow.svg')] after:w-[8px] after:h-[8px] after:bg-no-repeat after:bg-[length:6px_8px] after:bg-center 
                    last-of-type:after:hidden" data-id ="` + x.Id + `">` + x.Category + `</span>`

                        categoriess += x.Category + "~"

                        ids += x.Id

                        id = x.Id
                    }
                    div += ` </label>
                    <input type="hidden" id="cat` + index + `" data-id="` + id + `">
                </div>
              </li>`
                }
                $('.categry-lst').html(div);
            }

            if (result.SelectedCategory !== null) {
                result.SelectedCategory.forEach(function (category) {

                    newcategory = category.split(',').join('');

                    cateogryarr.push(newcategory)
                })
            }

            if (result.FieldValue !== null) {
                GetDaynamicChannelFields(result.FieldValue)
            }

        }
    })

})

// publish btn
$(document).on("click", "#publishbtn", function () {

    $("#savetype").val("publish")
    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);

})

// save btn
$(document).on("click", "#save-entry", function () {

    $("#savetype").val("draft")
    //use this in the submit button 
    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);

})

function containsHtmlElements(str) {   // check the string contains html elements

    var regex = /<[^>]*>/;
    return regex.test(str);
}

// select categories using check box
$(document).on('click', '.selectcheckbox', function () {
    cid = $(this).attr("data-id")
    catid = $("#cat" + cid).attr("data-id")
    if ($(this).prop('checked')) {
        var div = ""
        $("#sl-cat-error").hide()
        for (const [index, y] of Channelcategories.entries()) {
            if (index == cid) {
                $(".sl-list").removeClass("hidden")
                div += `<li class="cursor-pointer rounded-[4px] p-[8px_12px] border border-[#EDEDED] flex space-x-[6px] justify-between mb-[8px] last-of-type:mb-[0] sl-cat sl-cat` + index + `" data-id="` + index + `">
                                        <div class="relative flex space-x-[6px] items-center mb-0 text-[14px] font-normal leading-none text-[#262626] tracking-[0.005em] line-clamp-1 flex-wrap">`

                for (let x of y.Categories) {
                    div += `<span class="inline-flex items-center space-x-[6px] text-[12px] font-normal leading-none text-[#555555]
                                after:inline-block after:ml-[6px] after:bg-[url('/public/img/path-arrow.svg')] after:w-[8px] after:h-[8px] after:bg-no-repeat after:bg-[length:6px_8px] after:bg-center last-of-type:after:hidden" data-id ="` + x.Id + `">
                                `+ x.Category + `</span>`
                }
                div += `</div><a href="javascript:void(0);" class="bg-[#F7F7F5]  p-[3px] rounded-[3px] min-w-[16px] w-[16px] h-[16px] grid place-items-center hover:bg-[#F0F0F0] rsl-cat">
                    <img src="/public/img/remove-categories.svg" alt="remove">
                </a>
                <input type="hidden" data-id="`+ catid + `" class="cat-val">
                </li>`
            }
        }
        $(".selected-cat").append(div)

    } else {
        $(".sl-cat" + cid).remove()

        if ($(".selected-cat").find('li').length === 0) {
            $(".sl-list").addClass("hidden"); // Hide the empty category div
        }
    }
})

$(document).on('click', '.sl-cat', function () {
    cid = $(this).attr("data-id")
    $(this).remove()
    $("#Check" + cid).prop("checked", false)
    if ($(".selected-cat").find('li').length === 0) {
        $(".sl-list").addClass("hidden"); // Hide the empty category div
    }
})

function Getselectedcatagories() {   // check the string contains html elements
    $('.cat-val').each(function () {
        categoryIds.push($(this).attr("data-id"));
    })
}

// Get seo details
function Getseodetails() {
    title = $('#metatitle').val()
    desc = $('#metadesc').val()
    keyword = $('#metakey').val()
    slug = $('#metaslug').val()
    // imgtag = $("#imgtag").val()
    seodetails = { title: title, desc: desc, keyword: keyword, slug: slug, imgtag: "" }
}

function GetDaynamicChannelFields(additionalfields) {  // get dynamic additional fields based on channel 
    for (let x of additionalfields) {

        if (x.MasterFieldId == 2 || x.MasterFieldId == 7) {

            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                                <input type="text" placeholder="Text Here" class="bg-[#F7F7F5] p-[8px_12px] rounded-[4px] text-[14px] font-normal leading-[17.5px] tracking-[0.005em] border-none outline-none h-[34px] block w-full placeholder:font-normal placeholder:text-[#B2B2B2]" id="${x.FieldId}">
                                                                    <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                                </div>`)
        }
        if (x.MasterFieldId == 8) {

            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                                <textarea placeholder="Enter Text" class="bg-white  border border-[#EDEDED] p-[8px_12px] rounded-[4px] text-[14px] font-normal leading-[17.5px] tracking-[0.005em] border-none outline-none h-[120px] resize-none block w-full  placeholder:text-[#B2B2B2]" id="${x.FieldId}"></textarea>
                           <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label> </div>`)
        }
        if (x.MasterFieldId == 6) {

            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                                <input type="date" placeholder="Text Here" class="bg-[#F7F7F5] p-[8px_12px] rounded-[4px] text-[14px] font-normal leading-[17.5px] tracking-[0.005em] border-none outline-none h-[34px] block w-full placeholder:font-normal placeholder:text-[#B2B2B2] date" id="${x.FieldId}">
                          <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                                </div>`)
        }
        if (x.MasterFieldId == 4) {

            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                                <input type="datetime-local" placeholder="Text Here" class="bg-[#F7F7F5] p-[8px_12px] rounded-[4px] text-[14px] font-normal leading-[17.5px] tracking-[0.005em] border-none outline-none h-[34px] block w-full placeholder:font-normal placeholder:text-[#B2B2B2] date" id="${x.FieldId}">
                          <label class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                                </div>`)
        }
        if (x.MasterFieldId == 10) {

            var txt = ""

            if (x.OptionValue != null) {


                for (let y of x.OptionValue) {
                    txt += ` <a class="chk-group chk-group-label dropdown-item  lg items-center flex border-b border-solid border-[#ECECEC] last-of-type:border-b-0 !p-[8px_0] bg-white checkboxdiv">
                                            <input type="checkbox"  id="Check${y.Id}" class="hidden peer checkboxid">
                                            <label for="Check${y.Id}" class="checkbox-field h-[14px] relative cursor-pointer flex gap-[6px] w-full  items-center text-[14px] font-normal leading-[1] text-[#262626] tracking-[0.005em]
                                            before:bg-transparent before:w-[14px] before:h-[14px] before:inline-block before:relative before:align-middle before:cursor-pointer before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:-webkit-appearance-none peer-checked:before:bg-[url('/public/img/checked-box.svg')]  ">${y.Value}</label>
                                        </a>`
                }
            }
            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                <input type="hidden" class="checkfieldid" id="${x.FieldId}">
<label class=" fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
<div class="dropdown open ">
                                    <a href="javascript:void(0);" class="rounded-[4px] p-[8px_12px] h-[34px] bg-white border border-[#EDEDED] no-underline w-full flex justify-between items-center show drop-open" type="button" aria-haspopup="true" data-bs-auto-close="outside" aria-expanded="true">
                                        <span class="text-[#262626] leading-[17.5px] text-sm font-normal">Option</span>
                                        <img src="/public/img/drop-down-img.svg">
                                    </a>
                                    <div class="dropdown-menu min-w-full w-full border-0 p-[8px_12px] focus:bg-transparent drop-shadow-4 rounded-[4px] overflow-auto scrollbar-thin option-drop" aria-labelledby="triggerId" data-popper-placement="bottom-start" style="position: absolute; inset: 0px auto auto 0px; margin: 0px; transform: translate3d(0px, 36px, 0px);">${txt}</div>
                                </div>
                                <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                            </div>`)
        }
        if (x.MasterFieldId == 5) {

            var txt = ""

            if (x.OptionValue != null) {


                for (let y of x.OptionValue) {
                    txt += ` <a class=" chk-group chk-group-label dropdown-item  lg items-center flex  !p-[0px_0] bg-white checkboxdiv">
                                           
                                             <button id="optionsid"
                                                class=" !p-[2px] chk-group chk-group-label dropdown-items dropdown-item text-[14px] font-normal leading-[1] text-[#262626] dropdown-filter-roless cursor-pointer"
                                                type="button" data-id="${y.Id}">${y.Value}</button>
                                            
                                        </a>`
                }
            }
            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                <input type="hidden" class="fieldid" id="${x.FieldId}">
<label class=" fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
<div class="dropdown open ">
                                    <a href="javascript:void(0);" class="rounded-[4px] p-[8px_12px] h-[34px] bg-white border border-[#EDEDED] no-underline w-full flex justify-between items-center show drop-open" type="button" aria-haspopup="true" data-bs-auto-close="outside" aria-expanded="true">
                                        <span class="text-[#262626] leading-[17.5px] text-sm font-normal"  id="selectval">Option</span>
                                        <img src="/public/img/drop-down-img.svg">
                                    </a>
                                    <div class="dropdown-menu min-w-full w-full border-0 p-[8px_12px] focus:bg-transparent drop-shadow-4 rounded-[4px] overflow-auto scrollbar-thin option-drop" aria-labelledby="triggerId" data-popper-placement="bottom-start" style="position: absolute; inset: 0px auto auto 0px; margin: 0px; transform: translate3d(0px, 36px, 0px);">${txt}</div>
                                </div>
                                <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
                            </div>`)
        }
        if (x.MasterFieldId == 9) {

            var txt = ""

            for (let y of x.OptionValue) {
                txt += `<div class="mb-[10px]"><input type="radio" name="radio" id="radio${y.Id}" class="hidden peer radbtn">
                        <label for="radio${y.Id}" class="radio-field   relative cursor-pointer flex gap-[8px] items-center mb-[16px]     last-of-type:mb-0 text-sm font-normal text-[#222222] peer-checked:text-[#262626] tracking-[0.005em]
                                    before:bg-transparent before:w-[16px] before:h-[16px] before:inline-block before:relative before:align-middle before:cursor-pointer before:bg-[url('/public/img/radio-unchecked.svg')] before:bg-no-repeat before:bg-contain before:-webkit-appearance-none peer-checked:before:bg-[url('/public/img/radio-checked.svg')] leading-none ">${y.Value}</label></div>`
            }
            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                <input type="hidden" id="${x.FieldId}" class="radioval">
                                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px] " data-id="${x.FieldId}">${x.FieldName}</label>${txt}  <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*` + languagedata?.Channell?.errmsg + `</label></div>
                                `)
        }
    }
}

// field dropdown 
$(document).on('click', '.drop-open', function () {

    $(this).siblings(".option-drop").toggleClass("show")
})

function GetFieldValue() {
    channeldata = []

    $('.getvalue').each(function () {

        obj = {}
        obj = {
            "name": $(this).children('.fl-name').text().trim(),
            "fid": $(this).children('.fl-name').attr('data-id'),
            "value": $(this).find('input').val() || $(this).find('textarea').val(),
            "fieldid": $(this).find('input').attr('data-id') || $(this).find('textarea').attr('data-id'),

        }
        if ((obj.fieldid == undefined) || (obj.fieldid == '')) {
            obj.fieldid = '0'

        }

        channeldata.push(obj)



    })
}
$(document).on('click', '.checkboxid', function () {

    if ($(this).is(':checked')) {

        option = $(this).next('label').text()

        checkboxarr.push(option)
        $(this).parents(".getvalue").find('.checkfieldid').val(checkboxarr.join(','))
        $(this).parents(".getvalue").find('.manerr').hide()

    }
    else {
        var option = $(this).next('label').text();
        checkboxarr.splice(checkboxarr.indexOf(option), 1);
        console.log(checkboxarr,"checkboxarray")
        $(this).parents(".getvalue").find('.checkfieldid').val(checkboxarr.join(','));
        // $(this).parents(".getvalue").find('.manerr').show()
      

    }

})
$(document).on('click', '.radbtn', function () {

    option = $(this).next('label').text()
    $('.radioval').val(option)

    if ($('.radioval').val() == "") {

        // $(this).parents(".getvalue").find('.manerr').show()
    } else {
        $(this).parents(".getvalue").find('.manerr').hide()
    }


})

$(document).on('click', '#optionsid', function () {
    option = $(this).text()
    
    $('.fieldid').val(option)
    $('#selectval').text(option)

    if (($('.fieldid').val()) !== '') {


        $(this).parents(".getvalue").find('.manerr').hide()
    }

})
$(document).on("click", "#fieldsave", function () {
    channeldata = []
    GetFieldValue()
    var flag = channelfieldvalidation()



    if (flag == true) {

        console.log("check11")

        // $('.tab-togg3').click(function () {
        $('.editor-tabs3').removeClass('translate-x-[-378px] max-lg:translate-x-[0] max-lg:z-[999]  max-lg:min-w-[378px]');
        $('.editor-tabs3').addClass('translate-x-[100%]');
        $('.editor-tabs').removeClass('shadow-[-2px_0px_6px_0px_#0000000D]');
        $('.editor-tabs3').addClass(' shadow-[-8px_0px_16px_0px_#0000000D]');


    } else {

        console.log("check22")
        $('.editor-tabs3').addClass('translate-x-[-378px] max-lg:translate-x-[0] max-lg:z-[999]  max-lg:min-w-[378px]');
        $('.editor-tabs3').removeClass('translate-x-[100%]');
        $('.editor-tabs').addClass('shadow-[-2px_0px_6px_0px_#0000000D]');
        $('.editor-tabs3').removeClass(' shadow-[-8px_0px_16px_0px_#0000000D]');
    }

    channelfieldkeyup()
    console.log("new", channeldata);

})

function channelfieldvalidation() {

  

    var flag = true;
    $('.getvalue').each(function () {

        charallowed = $(this).data('char')

        if ((($(this).find('input').val() === '') && ($(this).attr('mandatory-fl') == 1)) || (($(this).find('textarea').val() === '') && ($(this).attr('mandatory-fl') == 1))) {

            $(this).find('.manerr').show()
            flag = false
        }
        if (($(this).find('input,textarea').val() !== '') && charallowed != 0) {
            if ($(this).find('input,textarea').val().length > charallowed) {
              
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

           
                $(this).siblings('.error').hide()
            } 
         

        })
        $(this).find('.date').change(function () {
            if (($(this).val() !== '')) {

                $(this).siblings('.error').hide()
            } 
            

        })

    })


}
// channel field config
$(document).on("click", "#field-config", function () {
    $('.sl-fields').remove()
    $("#ed-section").addClass("hidden")
    $("#field-section").removeClass("hidden")
    id = $('#slchannel').attr('data-id')
    orderindex = 1
    fiedlvalue = []
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

            for (let [index, x] of data.FieldValue.entries()) {

                AddFieldString(x.MasterFieldId, x.FieldId, x.FieldName, x.IconPath, x.DateFormat, x.TimeFormat, x.Mandatory, 1, id)

                if (x.OptionValue != null) {
                    for (let y of x.OptionValue) {

                        var intext = `<input type="hidden" class="opt-input" value="${y.Value}" data-id="${y.Id}" field-id="${x.FieldId}">`

                        $(".new-field" + (orderindex - 1)).find(".field-name").append(intext)
                    }
                }
            }


        }
    })

})
// channel field config
$(document).on("click", "#field-cancel", function () {
    $("#field-section").addClass("hidden")
    $("#ed-section").removeClass("hidden")
})

$(document).on("keyup", "#author", function () {

    // dropdown filter input box search
    $("#author").keyup(function () {

        var keyword = $(this).val().trim().toLowerCase()


        if (keyword != "") {

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

                    $('.optionsid2').remove();

                    if (data == null) {

                        $(".nodata-userlistdiv").removeClass("hidden")
                        $(".userlistdiv").addClass("show nouserdata")
                    }
                    if (data != null) {

                        $(".nodata-userlistdiv").addClass("hidden")

                        $(".userlistdiv").addClass("show").removeClass("nouserdata")

                        for (let x of data) {

                            console.log(x);

                            $('.userlistdiv').append(`<li><a class="optionsid2 dropdown-item p-[12px_16px] border-b border-solid border-[#EDEDED] text-[12px] font-normal leading-[16px] text-[#152027] hover:bg-[#F5F5F5]" href="javascript:void(0);" data-id=${x.Id}>${x.FirstName} ${x.LastName}</a></li>`)

                        }
                    }

                })
                .catch(error => {
                    console.error('There was a problem with your fetch operation:', error);
                });
        } else {

            $(".userlistdiv").removeClass("show nouserdata")

            $(".nodata-userlistdiv").addClass("hidden")
        }
    })

})

$(document).on('click', '.optionsid2', function () {


    test = $(this).text()

    $('#author').val(test)

    $(".userlistdiv").removeClass("show")


})

// dropdown filter data categories
$("#Searchcategorie").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".categry-lst > li").each(function (index, element) {
        var title = $(element).text().toLowerCase()
        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafounddesign").addClass("hidden")

        } else {
            $(element).hide()

            $("#nodatafounddesign").removeClass("hidden")
        }
    })

})




$(document).on('click', '#triggerId', function () {

    if ($('.addfdrop').hasClass('show')) {

        $('.addfdrop').removeClass('hide').addClass('show');
    } else {

        $('.addfdrop').removeClass('show').addClass('hide');
    }


});

$(document).on("click", '#field-update', function () {
    FieldBasedPropertiesValidation()

    if (isvalied == false) {
        GetfieldsValue()

        $.ajax({

            url: '/channels/Updatechannelfields',
            type: 'post',
            dataType: 'json',
            async: false,
            data: {
                "id": $("#slchannel").attr("data-id"),
                "sections": JSON.stringify({ sections }),
                "deletesections": JSON.stringify({ deletesecion }),
                "deletefields": JSON.stringify({ deletefields }),
                "deleteoptions": JSON.stringify({ deleteoption }),
                "fiedlvalue": JSON.stringify({ fiedlvalue }),
                csrf: $("input[name='csrf']").val(),
                "entryid":$('#eid').val()
            },
            success: function (result) {

              
                $(".select-chn").each(function () {
                    if ($(this).attr("data-id") == $("#slchannel").attr("data-id")) {
                        $(this).click()
                        
                        if (result && result.TblChannelEntryField) {
                           setTimeout(() => {
                            bindData(result.TblChannelEntryField)
                           }, 100);
                        } else {
                            console.error("No data available to bind.");
                        }
                        
                    }
                })

           
                $("#field-section").addClass("hidden")
                $("#ed-section").removeClass("hidden")
            
            }
        })

       
    }
})

function bindData(fields) {
    fields.forEach(function (field) {
     
        $("#" + field.FieldId).val(field.FieldValue);
        $("#" + field.FieldId).attr('data-id', field.Id);

    });

 
    var radval = $('.radioval').val();
    $('.radbtn').each(function () {
        if ($(this).next('label').text() === radval) {
            $(this).prop('checked', true);
        }
    });

   
    var selectval = $('.fieldid').val();
    $('#selectval').text(selectval);

  
    var cheval = $('.checkfieldid').val();
    if (cheval && cheval !== "") {
        var values = cheval.split(','); 
       
        $('.checkboxdiv').each(function () {
            var labelText = $(this).find('label').text();
            var inputDataId = $(this).find('input').attr('data-id');

            if (values.includes(labelText) || values.includes(inputDataId)) {
                $(this).find('input').prop('checked', true);
            }
        });
    }
}
$(document).on('click', '.date-drop', function () {


    if ($('.date-dropdown').hasClass('show')) {
        $('.date-dropdown').removeClass('hide').addClass('show');
    } else {
        $('.date-dropdown').removeClass('show').addClass('hide');
    }
})
$(document).on('click', '.time-drop', function () {

    if ($('.time-dropdown').hasClass('show')) {
        $('.time-dropdown').removeClass('hide').addClass('show');
    } else {
        $('.time-dropdown').removeClass('show').addClass('hide');
    }
})
$(document).on('click', '.dt-format', function () {

    $('.date-dropdown').addClass('hide')

})
$(document).on('click', '.tm-format', function () {
    $('.time-dropdown').addClass('hide');
})

$(document).on('click', '.dropup ', function () {
    $('.myprofilearrow').toggleClass('show')
})

$(document).ready(function () {

    $(document).on('click', function (event) {
        if (!$(event.target).closest('.dropdown').length) {
            $('#chn-list').removeClass("show");
        }
    });
});


$(document).on('click','.canceladditional',function(){


    $('.getvalue').each(function () {
        $(this).find('input,textarea,.fieldid').val('')
        // $(this).find('a').text('')
        $(this).find('.radbtn').prop('checked', false)
        $(this).find('.checkboxid').prop('checked', false)

    })

    $('.error').hide()
})