
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
var formid
var entryctaid
var editresult


var mapentryandmembershiplevel = []

$(document).ready(function () {
    $(document).on('click', '.mapmembershiplevels', function () {
        var membershiplevelId = $(this).data('id')
        if ($(this).prop('checked')) {

            mapentryandmembershiplevel.push(membershiplevelId)

        } else {

            mapentryandmembershiplevel.pop(membershiplevelId)

            $('#Check').prop('checked', false)

        }
    })

    $(document).on('click', function (event) {


        if (!$(event.target).closest('.userlistdiv').length) {

            $('.userlistdiv').removeClass('show').removeClass('nouserdata');
        }
    });
})

// logout dropdwon
$(document).on("click", "#lg-btn", function () {
    $("#log-btn").toggle("show")
})

$(document).ready(async function () {

    var gurl = window.location.href;

    var channelid = $("#chnid").val()

    if (gurl.includes('generatearticle')) {

        var newUrl = '/admin/entries/create/' + channelid + '/';

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

            if (url.includes('edit') || url.includes("edits")) {

                // var urlvalue = url.split('#')[0]
                // console.log("urlvalue",urlvalue);

                var eid = $("#eid").val()

            }
            console.log("pageno", pageno);

            var urlredirect = $("#dltCancelBtn").data('save')

            if (urlredirect == "save") {

                homeurl = $("#dltCancelBtn").attr('href')

            } else {

                if (pageno == null) {
                    homeurl = "/admin/entries/entrylist"
                } else {
                    homeurl = "/admin/entries/entrylist?page=" + pageno;
                }

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
                ChannelFieldsGet()
                authername = $("#author").val()
                createtime = $("#cdtime").val()
                publishtime = $("#publishtime").val()
                createtime = createtime.replace(" ", "T")
                publishtime = publishtime.replace(" ", "T")
                readingtime = $("#readingtime").val()
                orderindex = $("#orderindex").val()
                // sortorder = $("#article").val()
                tagname = $("#tagname").val()
                extxt = $("#extxt").val()
                languageid = $("#languageid").val()
                $.ajax({
                    url: "/admin/entries/draftentry/" + eid,
                    type: "POST",
                    dataType: "json",
                    data: { "id": $("#slchannel").attr('data-id'), "cname": channelname, "image": spurtdata.image, "title": spurtdata.title, "status": 0, "text": spurtdata.html, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": "", "tagname": tagname, "extxt": extxt, "orderindex": orderindex, "ctaid": formid, "membershiplevels": mapentryandmembershiplevel, "languageid": languageid },
                    success: function (result) {
                        window.location.href = homeurl;
                    }
                })
            }
        }

        if ($("#savetype").val() == "savePreview") {
            var url = window.location.href;

            var pageurl = window.location.search
            const urlpar = new URLSearchParams(pageurl)
            pageno = urlpar.get('page')

            if (url.includes('edit') || url.includes("edits")) {

                // var urlvalue = url.split('#')[0]
                // console.log("urlvalue",urlvalue);

                var eid = $("#eid").val()

            }
            console.log("pageno", pageno);

            if (pageno == null) {
                homeurl = "/admin/entries/entrylist"
            } else {
                homeurl = "/admin/entries/entrylist?page=" + pageno;
            }
            Getselectedcatagories()
            if (categoryIds.length == 0) {
                categoryIds.push($("#cat0").attr("data-id"));
            }
            channelname = $("#slchannel").val().trim().replace(/\s+/g, '-').toLowerCase();

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
                ChannelFieldsGet()
                authername = $("#author").val()
                createtime = $("#cdtime").val()
                publishtime = $("#publishtime").val()
                createtime = createtime.replace(" ", "T")
                publishtime = publishtime.replace(" ", "T")
                readingtime = $("#readingtime").val()
                orderindex = $("#orderindex").val()
                // sortorder = $("#article").val()
                tagname = $("#tagname").val()
                extxt = $("#extxt").val()
                languageid = $("#languageid").val()
                $.ajax({
                    url: "/admin/entries/draftentry/" + eid,
                    type: "POST",
                    dataType: "json",
                    data: { "id": $("#slchannel").attr('data-id'), "cname": channelname, "image": spurtdata.image, "title": spurtdata.title, "status": 0, "text": spurtdata.html, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": "", "tagname": tagname, "extxt": extxt, "orderindex": orderindex, "ctaid": formid, "membershiplevels": mapentryandmembershiplevel, "languageid": languageid },
                    success: function (result) {

                        var anchor = document.createElement('a');
                        anchor.href = result.previewurl; // Replace with your desired link
                        anchor.target = '_blank'; // Open in a new tab/window
                        document.body.appendChild(anchor);
                        anchor.click(); // Simulate a click on the anchor

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


            if (url.includes('edit') || url.includes("edits")) {

                var eid = $("#eid").val()
            }

            if (pageno == null) {
                homeurl = "/admin/entries/entrylist"
            } else {
                homeurl = "/admin/entries/entrylist?page=" + pageno;
            }
            channelname = $("#slchannel").val()
            Getselectedcatagories()
            if (categoryIds.length == 0) {
                categoryIds.push($("#cat0").attr("data-id"));
            }
            if ($("#slchannel").attr("data-id") == "") {
                $("#sl-chn-error").show()
                $('.editor-tabs').removeClass('translate-x-[120%]');
                $('#editingArea').addClass(' w-full ');
            }

            if (categoryIds.length == 0) {
                $("#sl-cat-error").show()
                $('.editor-tabs').removeClass('translate-x-[120%]');
                $('#editingArea').addClass(' w-full ');

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
                ChannelFieldsGet()
                authername = $("#author").val()
                createtime = $("#cdtime").val()
                publishtime = $("#publishtime").val()
                createtime = createtime.replace(" ", "T")
                publishtime = publishtime.replace(" ", "T")
                readingtime = $("#readingtime").val()
                // sortorder = $("#article").val()
                tagname = $("#tagname").val()
                extxt = $("#extxt").val()
                orderindex = $("#orderindex").val()
                languageid = $("#languageid").val()


                $.ajax({
                    url: "/admin/entries/publishentry/" + eid,
                    type: "POST",
                    dataType: "json",
                    data: { "id": $("#slchannel").attr('data-id'), "cname": channelname, "image": spurtdata.image, "title": spurtdata.title, "status": 1, "text": spurtdata.html, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val(), "author": authername, "createtime": createtime, "publishtime": publishtime, "readingtime": readingtime, "sortorder": "", "tagname": tagname, "extxt": extxt, "orderindex": orderindex, "ctaid": formid, "membershiplevels": mapentryandmembershiplevel, "languageid": languageid },
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

    console.log("url:", url);



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
            editurl = "/admin/entries/editentrydetails/" + channelname + "/" + id
        } else {
            editurl = "/admin/entries/editentrydetails/" + channelname + "/" + id + "?page=" + pageno
        }

    } else {

        if (pageno == null) {
            editurl = "/admin/entries/editentrydetail/" + id
        } else {
            editurl = "/admin/entries/editentrydetail/" + id + "?page=" + pageno
        }
    }

    if (url.includes('edit') || url.includes("edits")) {
        if (/^\d+$/.test(id)) {

            $.ajax({
                url: editurl,
                type: "GET",
                dataType: "json",
                data: { "id": id, csrf: $("input[name='csrf']").val() },
                success: function (result) {
                    editresult =result
                    var currentDate = new Date();
                    var formattedDate = new Date(currentDate.getTime() - currentDate.getTimezoneOffset() * 60000).toISOString().slice(0, 16);



                    $('#metatitle').val(result.Entries.MetaTitle)
                    $('#metadesc').val(result.Entries.MetaDescription)
                    $('#metakey').val(result.Entries.Keyword)
                    $('#metaslug').val(result.Entries.Slug)
                    Channelid = result.Entries.ChannelId
                    $("#tagname").val(result.Entries.Tags)
                    $("#extxt").val(result.Entries.Excerpt)
                    $("#orderindex").val(result.Entries.OrderIndex)
                    formid = result.Entries.CtaId
                    if (result.Entries.Author != "") {
                        $("#author").val(result.Entries.Author)
                    }
                    if (result.Createtime != "") {

                        const date = new Date(result.Createtime);
                        function formatDate(date) {
                            const year = date.getFullYear();
                            const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
                            const day = String(date.getDate()).padStart(2, '0');
                            const hours = String(date.getHours()).padStart(2, '0');
                            const minutes = String(date.getMinutes()).padStart(2, '0');

                            return `${year}-${month}-${day} ${hours}:${minutes}`;
                        }
                        const datetime = formatDate(date);

                        $("#cdtime").val(datetime)

                    } else {

                        $("#cdtime").val(formattedDate)

                    }

                    if (result.Publishedtime != "") {

                        const date = new Date(result.Publishedtime);
                        function formatDate(date) {
                            const year = date.getFullYear();
                            const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
                            const day = String(date.getDate()).padStart(2, '0');
                            const hours = String(date.getHours()).padStart(2, '0');
                            const minutes = String(date.getMinutes()).padStart(2, '0');

                            return `${year}-${month}-${day} ${hours}:${minutes}`;
                        }
                        const datetime = formatDate(date);



                        $("#publishtime").val(datetime)


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
                    if (result.Entries.MemebrshipLevelIds != "") {
                        mapentryandmembershiplevel = result.Entries.MemebrshipLevelIds.split(",")

                    }


                    for (let x of mapentryandmembershiplevel) {

                        $('.mapmembershiplevels').each(function () {

                            const $this = $(this);

                            if ($this.attr('data-id') === x) {

                                $this.prop('checked', true);

                            }


                        });
                    }

                    if (result.Entries.CtaId != 0) {

                        entryctaid = result.Entries.CtaId
                        console.log("ctacheck", result.Entries.CtaId)

                    }

                    for (const [index, y] of result.CategoryName.entries()) {

                        if (newcatarray.length != 0) {
                            $('.itemselecteddiv').removeClass('hidden')
                            $('#catcount').text(newcatarray.length)
                        }


                    }
                    if (result.Entries.TblChannelEntryField != null) {
                        result.Entries.TblChannelEntryField.forEach(function (field) {
                            console.log(field.FieldName,"ggggg")
                            $(`#${field.FieldId}`).val(field.FieldValue)
                            $(`#${field.FieldId}`).attr('data-id', field.Id)

                            if (field.FieldName=="File Upload"){
                                 console.log("qqqqq")
                                $('.fileuploaddiv').addClass('hidden')
                                $('.zip-preview').removeClass('hidden')
                                
                            }
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


$(document).on('click','.tab-togg3',function(){
 
 
   var url = window.location.href;
     if (url.includes('edit') || url.includes("edits")) {
    
  if (editresult.Entries.TblChannelEntryField.length > 0) {
 
 
                        editresult.Entries.TblChannelEntryField.forEach(function (field) {
                           
                            
                            $(`#${field.FieldId}`).val(field.FieldValue)
                            $(`#${field.FieldId}`).attr('data-id', field.Id)
                            
                            if (field.FieldName=="File Upload"){
                                
                                $('.fileuploaddiv').addClass('hidden')
                                $('.zip-preview').removeClass('hidden')
                                $('.uploaded-filepath').val(field.FieldValue)
                                
                            }
                        })
                    }
                }
})
 
// New code 
$(document).on("click", "#sl-chn", function () {
    $("#chn-list").toggleClass("show")
})


$(document).on('click', '.tab-togg55', function () {

    $('.ctacard').each(function () {
        if ($(this).attr('data-id') == entryctaid) {
            console.log("ctacheck")
            $(this).find('h5').removeClass('text-[#717171]').addClass('text-[#10A37F]')
            $(this).removeClass('border-[#EDEDED]')
            $(this).addClass('border-[#10A37F]')
        }
    })

})
// Get channel fields
$(document).on('click', '.select-chn', function () {

    categoryIds = []

    seodetails = {};

    channeldata = []
    $(".sl-list").addClass("hidden")
    $(".selected-cat").empty()
    $("#chn-list").removeClass("show")
    $("#defaultChannel").modal('hide')
    $("#defaultChannel").removeClass("show")

    chanid = $(this).attr('data-id')

    $("#chn-list").removeClass("show")
    $("#sl-chn-error").hide()
    $("#chn-name").text($(this).text()).addClass("text-bold-black")
    $("#chn-name").attr("title", $(this).text())
    $("#chn-name").attr("data-id", chanid)
    $("#slchannel").val($(this).text())
    $("#slchannel").attr("data-id", chanid)
    $('.itemselecteddiv').addClass('hidden')

    $('#catcount').text("")
    $.ajax({
        url: "/admin/entries/channelfields/" + chanid,
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
                <li class="p-[9px_24px] border-b border-[#EDEDED]" data-id="`+ index + `">
                   <div class="chk-group chk-group-label">
                  <input type="checkbox" id="Check` + index + `" class="hidden peer selectcheckbox" data-id="` + index + `">
                                                      <label for="Check` + index + `" class="relative cursor-pointer flex space-x-[6px] items-center mb-0 text-[14px] font-normal leading-none text-[#262626] tracking-[0.005em]
                                before:bg-transparent before:min-w-[14px] before:w-[14px] before:h-[14px] before:inline-block before:relative before:align-middle before:cursor-pointer before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:mr-[6px] before:-webkit-appearance-none peer-checked:before:bg-[url('/public/img/checked-box.svg')]  
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
                    <input type="hidden" class="ncati" id="cat` + index + `" data-id="` + id + `">
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
            $('.ctacard').remove()
            $('.listempydes').remove()
            $('.ctanodata').addClass('hidden')
            $('.initaildatadiv').addClass('hidden')
            if (result.ctalist !== null && result.ctalist !== "" && result.ctalist.length != 0) {

                console.log(result.ctalist, "listcta")

                for (let x of result.ctalist) {
                    console.log("x", x.Id)
                    let str = `<div class="flex flex-col items-center space-y-2  border-[1px]  p-[12px] rounded-[4px] cursor-pointer ctacard" data-id="${x.Id}">`;

                    if (x.FormImagePath != "") {
                        console.log("jjjjiii")
                        str += `   <img src="${x.FormImagePath}" alt="meta title">`;
                    }

                    str += `   <h5 class="mt-[8px] mb-[4px] font-normal text-[#717171] text-center text-sm">${x.FormTitle}</h5>
                              <p class="font-light text-[#717171] text-[10px] text-center">${x.FormDescription}</p>
                          </div>`;
                    if ($('.calltoactdiv').is(':visible')) {

                        $('.calltoactdiv').append(str);
                    } else {

                        const newDiv = $('<div class="gap-[12px] grid grid-cols-2 max-[350px]:grid-cols-1 calltoactdiv"></div>');
                        newDiv.append(str);
                        $('.ctalistdiv').append(newDiv);
                    }

                }


                // $('.calltoactdiv').append(str)
            } else {
                nodata = `  <div
            class="noData-foundWrapper flex flex-col justify-content-center items-center h-full listempydes ">

            <div class="empty-folder mb-[10px]">
                <img style="max-width: 50px;" src="/public/img/Formsnodata.svg" alt="">
                <img src="/public/img/shadow.svg" alt="">
            </div>
            <h1 style="text-align: center;font-size: 15px;" class="heading">
                 No CTA's yet?</h1>

        </div>`
                $('.ctalistdiv').append(nodata)
            }
        }
    })

})


$(document).on('click', '#availblecate', function () {

    // $('.tabclose').hide()

    for (let x of newcatarray) {


        $('.ncati').each(function () {

            const $this = $(this);

            if ($this.attr('data-id') === x) {

                $this.closest('li').find(".selectcheckbox").prop('checked', true);
                $this.addClass('cat-val');
                const $li = $this.closest('li');
                $li.prependTo($('.categry-lst'));
            }


        });

    }
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
$(document).ready(function () {
    $(document).on('click', '.selectcheckbox', function () {

        const $this = $(this);
        const cid = $this.attr("data-id");
        const $li = $this.closest('li');
        const $ul = $li.parent('ul');
        const $hiddenInput = $("#cat" + cid);
        const catid = $hiddenInput.attr("data-id");


        if ($this.prop('checked')) {

            $('.itemselecteddiv').removeClass("hidden")
            $ul.prepend($li);
            $hiddenInput.addClass('cat-val');
        } else {

            const originalIndex = $li.data('order-index');

            if (originalIndex !== undefined) {
                let $lis = $ul.children('li');
                if (originalIndex === 0) {
                    $ul.prepend($li);
                } else {
                    if (originalIndex >= $lis.length) {
                        $ul.append($li);
                    } else {
                        $li.insertBefore($lis.eq(originalIndex));
                    }
                }
            } else {

                $ul.append($li);
            }


            $hiddenInput.removeClass('cat-val');

        }
        updateCategoryCount();
    });

    // Store original order on page load: very important
    $('.categry-lst > li').each(function (index) {
        $(this).data('order-index', index);
    });


    updateCategoryCount();
    if ($('.selectcheckbox:checked').length === 0) {
        console.log("lll")
        $('.itemselecteddiv').addClass("hidden");
    }
});
function updateCategoryCount() {
    const count = $('.selectcheckbox:checked').length;
    $('#catcount').text(count);
    if (count > 0) {
        $('.itemselecteddiv').removeClass('hidden');
    } else {
        $('.itemselecteddiv').addClass('hidden');
    }
}


$(document).on('click', '.sl-cat', function () {

    console.log("removesdfdsfdsfdsfd")
    cid = $(this).attr("data-id")
    $(this).remove()
    console.log("cid", cid)
    $("#Check" + cid).prop("checked", false)
    $("#Check" + cid).parents("li").removeClass("hidden")
    if ($(".selected-cat").find('li').length === 0) {
        $(".sl-list").addClass("hidden"); // Hide the empty category div
    }
    let liCount = 0;
    $('.sl-cat').each(function () {
        liCount++;

    })

    $('#catcount').text(liCount)
})
$(document).on('click', '.catclearbtn', function () {

    $('.selectcheckbox:checked').each(function () {
        const $checkbox = $(this);
        const cid = $checkbox.attr("data-id");
        const $hiddenInput = $("#cat" + cid);

        $checkbox.prop('checked', false);
        $hiddenInput.removeClass('cat-val');
    });

    const $ul = $('.categry-lst');
    const $listItems = $ul.children('li');

    $listItems.sort(function (a, b) {
        const dataIdA = parseInt($(a).find('.selectcheckbox').data('id'));
        const dataIdB = parseInt($(b).find('.selectcheckbox').data('id'));
        return dataIdA - dataIdB;
    });

    $ul.empty().append($listItems);
    $('.itemselecteddiv').addClass('hidden');
    $('#catcount').text("")
});

// $(document).on('click','.catclearbtn',function(){

//     $('.sl-cat').click()
// })
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
        console.log(x.MasterFieldId, "checkdifdfdf")
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
        if (x.MasterFieldId == 17) {
            $(".add-fl").append(`<div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                <input type="number" placeholder="Enter Here" class="bg-[#F7F7F5] p-[8px_12px] rounded-[4px] text-[14px] font-normal leading-[17.5px] tracking-[0.005em] border-none outline-none h-[34px] block w-full placeholder:font-normal placeholder:text-[#B2B2B2]" id="${x.FieldId}">
                                                    <label   class="text-[#f26674] font-normal text-xs error manerr" id="opterrr" style="display: none">*`+ languagedata?.Channell?.errmsg + `</label>
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

        if (x.MasterFieldId == 15) {


            $(".add-fl").append(` <div class="mb-[24px] last-of-type:mb-0 getvalue" mandatory-fl="${x.Mandatory}">
                               <label class="fl-name text-[14px] font-normal leading-[17.5px] text-[#262626] mb-[6px]" data-id="${x.FieldId}">${x.FieldName}</label>
                               <input type="hidden" id="${x.FieldId}" class="uploaded-filepath" name="uploadedFilePath">

                                <div class="fileuploaddiv btn-group w-full block">
                                    <button class="p-[8px_12px] w-full flex justify-center items-center relative  rounded cursor-pointer gap-1 bg-[#F9F9F9] hover:bg-[#e0e0e0] border-[#E6E6E6] border border-solid"
                                        >
                                        <img src="/public/img/upload-folder.svg" alt="">
                                        <p class="text-[14px] font-normal leading-[18px] text-[#616161]">Upload File</p>
                                        <input type="file" 
                                        class="fileupload absolute w-full h-full opacity-0 text-[0]" accept=".zip">
                                       
                                    </button>
                                     
                                 <p class="file-error mt-2 text-[#F26674] text-xs hidden"></p>
                                </div>
                                 <div class="zip-preview hidden flex items-center space-x-[8px] mb-[8px] bg-[#FBFBFB] px-3 py-2 mt-2 rounded w-full border-[#E6E6E6] border border-solid">
                                    <div class="w-[32px] h-[32px] rounded  grid place-items-center min-w-[32px]">
                                        <img src="/public/img/Export file icon.svg" alt="">
                                    </div>
                                    <div>
                                        <h4 class="text-[13px] font-normal leading-[20px] text-[#262626] !me-auto">.zip
                                        </h4>
                                       
                                    </div>
                                    <button class="flex items-center ms-auto cursor-pointer select-none text-dark">
                                        <img src="/public/img/close-toast.svg" alt="">
                                    </button>
                                </div>
                            </div>`)
        }
    }
}


//file upload function of additional field//
$(document).on('change', '.fileupload', function () {
    var file = this.files[0];
   console.log(file,"fileeee")
 $('.uploaded-filepath').val("")
    if (!file) return;

    var ext = file.name.split('.').pop().toLowerCase();
    if (ext !== 'zip') {
        $('.file-error').removeClass("hidden").text('Only .zip files are allowed.');
        this.value = '';
        return false;
    }

    var maxSize = 100 * 1024 * 1024; // 100 MB
    if (file.size > maxSize) {
        $('.file-error').removeClass("hidden").text('Zip file must be 100 MB or less.');
        this.value = '';
        return false;
    }
  

  
    $('.file-error').addClass("hidden").text('')
        $(this).parents('.btn-group').addClass('hidden')
    $('.zip-preview').removeClass("hidden")

              var formData = new FormData();
            formData.append('documentdata', file);

            $.ajax({
                url: '/uploadb64document',
                method: 'POST',
                contentType: false,
                processData: false,
                data: formData,
                success: function (response) {


                    console.log(response,"responsedatadfdfd")

                    const urlObj = new URL(response.documentpath);

                    const finalpath = urlObj.pathname + urlObj.search;

                    $('.uploaded-filepath').val(finalpath)
                }
            })

});
$(document).on('click', '.zip-preview button', function () {
    var $preview = $(this).closest('.zip-preview');
    $preview.addClass('hidden')
    $('.fileupload').val('');
     $preview.siblings('.btn-group').removeClass('hidden')
});
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
        console.log(checkboxarr, "checkboxarray")
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
// $(document).on("click", "#fieldsave", function () {

function ChannelFieldsGet() {
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

}

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
        url: '/admin/channels/getfields',
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
    // $("#author").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()


    if (keyword != "") {

        fetch(`/admin/entries/userdetails/?keyword=${keyword}`, {
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

                $('.userlistdiv li:not(.nodata-userlistdiv)').remove();

                if (data == null) {
                    $(".nodata-userlistdiv").removeClass("hidden")
                    $(".userlistdiv").addClass("show nouserdata")
                }
                if (data != null) {
                    $(".nodata-userlistdiv").addClass("hidden")
                    $(".userlistdiv").addClass("show").removeClass("nouserdata")


                    for (let x of data) {

                        console.log(x);

                        $('.userlistdiv').append(`<li style="width: 100%"><a class="optionsid2 dropdown-item p-[12px_16px] border-b border-solid border-[#EDEDED] text-[12px] font-normal leading-[16px] text-[#152027] hover:bg-[#F5F5F5]" href="javascript:void(0);" data-id=${x.Id}>${x.FirstName} ${x.LastName}</a></li>`)

                    }
                }

            })
            .catch(error => {
                console.error('There was a problem with your fetch operation:', error);
            });
    } else {

        $(".userlistdiv").removeClass("show nouserdata")

        $(".nodata-userlistdiv").addClass("hidden")

        $('.userlistdiv li:not(.nodata-userlistdiv)').remove();


    }
})

// })

$(document).on('click', '.optionsid2', function () {


    test = $(this).text()

    $('#author').val(test)

    $(".userlistdiv").removeClass("show")


})

// dropdown filter data categories
$("#Searchcategorie").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    var found = false;
    $(".categry-lst > li").each(function (index, element) {
        var title = $(element).text().toLowerCase()
        if (title.includes(keyword)) {
            $(element).show()
            $('.selected-cat').addClass('hidden')
            found = true;

        } else {

            $(element).hide()

        }
        if (found) {
            updateCategoryCount()
            $('.itemselecteddiv').removeClass('hidden')
            $("#nodatafounddesign").addClass("hidden");
        } else {

            $('.itemselecteddiv').addClass('hidden')
            $("#nodatafounddesign").removeClass("hidden");
        }
    })
})


// dropdown filter Membership level
$("#Searchmembership").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    var found = false;
    $(".chk-group-label > div").each(function (index, element) {
        var title = $(element).text().toLowerCase()
        if (title.includes(keyword)) {
            $(element).show()
            // $('.selected-cat').addClass('hidden')
            found = true;

        } else {

            $(element).hide()

        }
        if (found) {
            // updateCategoryCount()
            // $('.itemselecteddivs').removeClass('hidden')
            $("#membership_nodatafound").addClass("hidden");
        } else {

            // $('.itemselecteddivs').addClass('hidden')
            $("#membership_nodatafound").removeClass("hidden");
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

            url: '/admin/channels/Updatechannelfields',
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
                "entryid": $('#eid').val()
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


$(document).on('click', '.canceladditional', function () {

    $('.tabclose').show()


    // $('.getvalue').each(function () {
    //     $(this).find('input,textarea,.fieldid').val('')
    //     // $(this).find('a').text('')
    //     $(this).find('.radbtn').prop('checked', false)
    //     $(this).find('.checkboxid').prop('checked', false)

    // })

    $('.error').hide()
})

$(document).on('click', '#save-preview', function () {

    $("#savetype").val("savePreview")
    //use this in the submit button 
    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);

})

$(document).on('click', '#cancel-entry', function () {

    $('#content').text("Your changes will be lost. Do you wish to continue?");
    $(".deltitle").text("Discard Unsaved Entry?")
    $("#delid").text("Continue")
    $('#delid').attr('href', '/admin/entries/entrylist');
    $("#dynamicImage").attr('src', "/public/img/Cancel.svg")

})


// $(document).on('click', '#chn-list .dropdown-item', function () {

//     $('#chn-list li').click(function () {
//         notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px]  border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/delete-icon.svg" alt = "toast success"></div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Info !</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ` + "Switching channels will reset selected categories" + `</p ></div ></div ></li></ul> `;
//         $(notify_content).insertBefore(".header-rht");

//         setTimeout(function () {
//             $('.toast-msg').fadeOut('slow', function () {
//                 $(this).remove();

//             });
//         }, 5000);
//     });
// })


// $(document).on('click','.ctacard',function(){

//     $('.ctacard h5').removeClass('text-[#10A37F]').addClass('text-[#717171]');
//     $('.ctacard').removeClass('border-[#10A37F]').addClass('border-[#EDEDED]');
//     formid =$(this).attr('data-id')

//     $(this).find('h5').removeClass('text-[#717171]').addClass('text-[#10A37F]')
//     $(this).removeClass('border-[#EDEDED]').addClass('border-[#10A37F]')

// })

$(document).on('click', '.ctacard', function () {

    formid = $(this).attr('data-id');

    $(this).addClass("bg-[#FBFFFE]")
    if ($(this).find('h5').hasClass('text-[#10A37F]')) {

        $(this).find('h5').removeClass('text-[#10A37F]').addClass('text-[#717171]');
        $(this).removeClass('border-[#10A37F]').addClass('border-[#EDEDED]');
        formid = 0;
    } else {

        $('.ctacard h5').removeClass('text-[#10A37F]').addClass('text-[#717171]');
        $('.ctacard').removeClass('border-[#10A37F]').addClass('border-[#EDEDED]');


        $(this).find('h5').removeClass('text-[#717171]').addClass('text-[#10A37F]');
        $(this).removeClass('border-[#EDEDED]').addClass('border-[#10A37F]');
    }


});
$(document).on('keyup', '.ctasearch', function () {
    var keyword = $(this).val().toLowerCase();

    var $ctacards = $('.ctacard');
    var hasVisibleCards = false;

    if (keyword !== "") {

        $ctacards.each(function () {
            var title = $(this).find('h5').text().toLowerCase();
            if (title.includes(keyword)) {
                $(this).removeClass('hidden');
                hasVisibleCards = true;
            } else {
                $(this).addClass('hidden');
            }
        });


        if (hasVisibleCards) {
            $('.ctanodata').addClass('hidden');
        } else {
            $('.ctanodata').removeClass('hidden');
            $('.listempydes').addClass('hidden')
        }
    } else {

        $ctacards.removeClass('hidden');
        $('.ctanodata').addClass('hidden');
        $('.listempydes').removeClass('hidden')
    }
});


flatpickr("#publishtime", {
    enableTime: true,
    dateFormat: "d M Y h:i K",
});


flatpickr("#cdtime", {
    enableTime: true,
    dateFormat: "d M Y h:i K",
});


$(document).on('click', '#AltText', function () {

    $('body').addClass('overflow-hidden')
})

$(document).on('click', '.altSave', function () {

    $('body').removeClass('overflow-hidden')
})

$(document).on('click', '.altCancel', function () {

    $('body').removeClass('overflow-hidden')
})

$(document).ready(function () {

    $(document).on('click', ".doc-lst", function (event) {

        value = $("#editingArea").data('edited')

        if ($('#editingArea').hasClass('data-edited')) {

            event.preventDefault();

            $("#deleteModal").modal('show')

            $('#dltCancelBtn').removeAttr('data-bs-dismiss');

            tablurl = $(this).find("a").attr("href")

            $('#dynamicImage').attr('src', '/public/img/info-icon.svg')

            $('.deldesc').removeClass("text-center")

            $('.deldesc').html(`
                <div class="text-center">Are you sure you want to leave this page?</div><br>
                <p style="margin-left:36px;">- Click 'Save' to save it as draft.</p><br>
                <p style="margin-left:36px;">- Click 'Cancel' to proceed without saving.</p>
            `);


            $('.deltitle').text("Exit Entry")
            $("#delid").text("Save")
            $("#delid").addClass("featurebtn")
            $("#delcancel2").text(value)
            $('#delid').attr('id', 'save-entry');
            var currentUrl = window.location.href;

            // Remove the specific part of the URL
            var index = currentUrl.indexOf('/entries/');

            var newUrl = currentUrl.substring(0, index);


            $("#dltCancelBtn").attr('href', tablurl)
            $('#dltCancelBtn').attr('data-save', 'save');
        }


    })

});

$(document).on('click', '.membershipclose', function () {
    $('.tabclose').show()
})

$(document).on('click', '.languageselect', function () {
    let id = $(this).data('id');

    $("#languageid").val(id)

    let languagename = $(this).text()

    console.log("languagename::", languagename);

    $("#selectedLanguage").text(languagename)

    // Close the Bootstrap dropdown properly
    const dropdownToggle = $(this).closest('.dropdown').find('[data-bs-toggle="dropdown"]')[0];
    const dropdownInstance = bootstrap.Dropdown.getInstance(dropdownToggle)
        || new bootstrap.Dropdown(dropdownToggle);
    dropdownInstance.hide();
});