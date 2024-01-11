var languagedata
let ckeditor1;
var categoryIds = [];
var seodetails = {};
var channeldata = []
checkboxarr = []
var cateogryarr = []

console.log(cateogryarr, seodetails, categoryIds, "kdfjdfdkffdfd")

/** */
// $(document).ready(async function () {

//     var languagecode = $('.language-group>button').attr('data-code')

//     await $.getJSON("/locales/"+languagecode+".json", function (data) {

//         languagedata = data
//     })

// })

$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/" + languagecode + ".json", function (data) {

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

    var urlvalue = url.split('?')[0]
    var id = urlvalue.split('/').pop()

    // var id = url.split('/').pop().split('?')[0];
    // var id = url.split('/').pop();

    // var channelname= url.split('/').pop();

    var array = url.split('/');

    // var id = array[array.length - 1];

    var channelname = array[array.length - 2];
    console.log("id", id, channelname);
    if (/^\d+$/.test(id)) {
        $.ajax({
            url: "/channel/editentrydetails/" + channelname + "/" + id +"?page=" + pageno,
            type: "GET",
            dataType: "json",
            data: { "id": id, csrf: $("input[name='csrf']").val() },
            success: function (result) {
                console.log(result, "result")

                if (result.Entries.CoverImage == "") {
                    $('#cenimg').show()

                }
                else {
                    $('#cenimg').hide()
                    $("#spimagehide").css("height", "23.120rem")
                    $("#spimagehide").css("width", "46.25rem")
                    $("#spacedel-img").show();
                    $('#spimagehide').attr('src', result.Entries.CoverImage)
                }

                $('#titleid').val(result.Entries.Title)
                // $('#text').val(result.Description)
                ckeditor1.setData(result.Entries.Description)
                $('#chanid').val(result.Entries.ChannelId)
                $('#metatitle').val(result.Entries.MetaTitle)
                $('#metadesc').val(result.Entries.MetaDescription)
                $('#metakey').val(result.Entries.Keyword)
                $('#metaslug').val(result.Entries.Slug)
                $('#seosavebtn').text(languagedata.update)
                $('#categorysave').text(languagedata.update)

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
                        $('#blogn').text(channelname)
                        $('#blogcount').text(count)
                        $('.blogDropdown').append('<img class="blogimg" src="/public/img/blog-drop.svg" alt="">');
                    }
                }


                if (result.SelectedCategory !== null) {
                    result.SelectedCategory.forEach(function (category) {


                        newcategory = category.split(',').join('');

                        cateogryarr.push(newcategory)



                    })
                }



                if (cateogryarr !== null) {

                    cateogryarr.forEach(function (category) {

                        $('#category-' + category).show();

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
                $('#channelfieldsave').text('Update')

                if (result.Section != null) {
                    result.Section.forEach(function (section) {

                        $('.secid').append(`<div id="${section.SectionId}"><h6 class="sechead">${section.SectionName}</h6></div>`)

                    })
                    if (result.FieldValue != null) {
                        result.FieldValue.forEach(function (fieldValue) {
                            if (fieldValue.MasterFieldId == 4) {
                                $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text" class="date" data-id="" data-date="${fieldValue.DateFormat}" data-time="${fieldValue.TimeFormat}" placeholder="` + languagedata.Channell.seldate + ` "  id="date${fieldValue.FieldId}"   value=""/>
                    
                    </div>
                    <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
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
                                        <label class="error" id="opterr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
        
                                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 6) {
                                $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date" data-id="" data-date="${fieldValue.DateFormat}" id="date${fieldValue.FieldId}"  placeholder="` + languagedata.Channell.seldate + `" />
                    
                    
                    </div>
                    <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
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
            <label class="error" id="radioerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
            </div>
            </div>
            
              
            
                `);
                            }
                            if (fieldValue.MasterFieldId == 8) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" >
                <label  class="input-label"  data-id="${fieldValue.FieldId}" id="textarealabel" for="">${fieldValue.FieldName}</label>
             
                <div class="ig-row">
                  <textarea placeholder="`+ languagedata.Channell.pltexthere + `" data-id="" id="${fieldValue.FieldId}" ></textarea>
                </div>
                <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
              </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 2) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata.Channell.pltexthere + `" id="${fieldValue.FieldId}"   value="" />
                      </div>
                      <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                    </div>
                `);
                            }
                            if (fieldValue.MasterFieldId == 7) {
                                $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" data-id="" placeholder="`+ languagedata.Channell.pltexthere + `" id="${fieldValue.FieldId}"   value=""   />
                      </div>
                      <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                    </div>
                `);
                            }

                            if (fieldValue.MasterFieldId == 10) {
                                $(`#${fieldValue.SectionId}`).append(`
                <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id="${fieldValue.Mandatory}"  style="margin-bottom: 20px">
         
                                         <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
         
                                         <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                             aria-haspopup="true" aria-expanded="false">
                                         option
                                         </a>
                                         <input type="hidden" data-id="" class="checkfieldid" id="${fieldValue.FieldId}"  value=""  >
                                         <div class="dropdown-menu additional-entry-drop" id="check${fieldValue.FieldId}" aria-labelledby="triggerId">
                 
                                             
                                         </div>
                                         <label class="error" id="opterrr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
         
                                     </div>
                `)


                            }

                            //     })
                            // }
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


                var radval = $('.radioval').val()
                console.log(radval, "radiovalue")
                $('.radiodiv').each(function () {
                    if ($(this).find('label').text() == radval) {
                        console.log("uuuu")
                        $(this).find('input').prop('checked', true)
                    }
                })
                //category bind//

                categoryIds = result.Entries.CategoriesId.split(',');

                console.log("categorids", categoryIds)
                var catselect = document.querySelectorAll('.categories-list-child');
                console.log(catselect, "catelist");
                categoryIds.forEach(function (categoryId) {
                    var matchedH3 = document.querySelector('a[data-id="' + categoryId + '"]');
                    console.log("match", matchedH3);

                    matchedH3.click()

                });

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

/** Category select */
$(document).ready(function () {

    $('.categorypdiv').hide();

    var count = 0

    $('.categorypdiv').each(function () {

        if ($(this).css('display') != "none") {

            count = count + 1

        }

        var id

        $(this).children('.choose-cat-list-col').children('.categoryname').each(function () {

            id = $(this).attr('data-id');

        })

        $(this).children('.category-select-btn').attr('data-id', id).attr('id', 'lastcatid-' + id)
    })

    // $('#avacatcount').text(count);

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

    $('.deltitle').text(languagedata.Channell.delentrytitle)

    $('.deldesc').text(languagedata.Channell.delentrycontent)

    $('#delid').attr('data-id', $(this).attr('data-id'))

    $('#delid').text(languagedata.yes)

    $('#delcancel').text(languagedata.no)

    $('#delid').parent('#delete').attr('href', "/channel/deleteentries/?id=" + entryId + "&cname=" + channame);

    // $('#delid').parent('#delete').attr('href', "/channel/deleteentries/"+ entryId + "/" + channelname);


})


//*Publish and Unpublish status change */
function ChangeStatus(entryid, channelname, status) {

    $('.filter-dropdown-menu ').hide()

    $.ajax({
        url: "/channel/changestatus/" + entryid + "?" + "entry=" + entryid + "cname=" + channelname + "&&status=" + status,
        type: "post",
        data: {
            entryid: entryid,
            status: status,
            cname: channelname,
            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        success: function (result) {

            console.log("result", result)
            if (result) {

                if (status == 1) {

                    $("#statusname-" + entryid).text("Published");

                    $("#statusname-" + entryid).removeClass('unpublished').addClass('published')

                    $("#statusname-" + entryid).removeClass('draft').addClass('published')

                } else if (status == 2) {

                    $("#statusname-" + entryid).text("Unpublished");

                    $("#statusname-" + entryid).removeClass('published').addClass('unpublished')

                    $("#statusname-" + entryid).removeClass('draft').addClass('unpublished')



                }

                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Toast.entryupdatenotify + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds


            } else {


                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Toast.internalserverr + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds


            }
        }
    })

}

// $(document).on('click', ".upload-folder-img", function () {
//     var src = $(this).find('img').attr("src");
//     var data = $("#spimage").val(src)
//     var data1 = $("#spimagehide").attr("src", src);

//     if (data1 != "") {

//         $('#cenimg').hide()
//         $("#spimagehide").css("height", "23.120rem")
//         $("#spimagehide").css("width", "46.25rem")
//         $("#spimagehide").attr("src", src).show()
//         $("#spacedel-img").show();
//     }
//     $("#addnewimageModal").hide()
// })

$('#spacedel-img').click(function () {
    $('input[name=spaceimagepath]').val("")
    $('#spimagehide').attr('src', '')
    $(this).hide()
    $('#cenimg').show()
    $('#spimagehide').hide()
})

//Save to Draft Function//

$('#draftbtn').click(function () {

    var url = window.location.href;

    console.log("url", url)

    if (url.includes('editentry')) {
        var eid = url.split('/').pop();

        console.log("id", eid);
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

    if (title === '' && data === '' && img === '') {
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Channell.descerr + ' </span></div>';
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

    }
    if (title !== '' && data !== '' && img !== '') {
        var flag = channelfieldvalidation()

        channelfieldkeyup()

        if (flag == false) {
            $('#channelModal').modal('show');
        }
        var flag1 = categoryvalidation()
    }
    if (title && data != '' && img != '' && (flag == true) && (flag1 == true)) {
        $.ajax({
            url: "/channel/draftentry/" + eid,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "text": data, "status": 0, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val() },
            success: function (result) {

                console.log("resultiiii", result)


                // if (result == true) {

                $('.blog-status-container').show()
                setTimeout(function () {
                    $('.blog-status-container').hide();
                }, 5000);

                setTimeout(function () {
                    window.location.href = "/channel/editentry/" + result.Channelname + "/" + result.id
                }, 2000);

                //   var  route =eid
                //     console.log(route,"hhhhh")
                //     if (eid =='newentry'){
                //         console.log("eiddd")
                //     bcount = parseInt(blogcount) + 1;
                //     $('#blogcount').text(bcount + ' Blog Available')
                //     $('#encount').text(bcount + ' Blog Available')
                //        }
                //     $('#chanid').val('')
                //     $('#spimagehide').attr('src', '')
                //     $('#titleid').val('')
                //     $('#text').val('')
                //     $('#cenimg').show()
                //     $('#spimagehide').hide()
                //     $("#spacedel-img").hide();
                //     ckeditor1.setData('')

                // } else {

                //     notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>"Internal Server Error"</span></div>';
                //     $(notify_content).insertBefore(".header-rht");
                //     setTimeout(function () {
                //         $('.toast-msg').fadeOut('slow', function () {
                //             $(this).remove();
                //         });
                //     }, 5000); // 5000 milliseconds = 5 seconds
                // }
            }
        })
    }
})



$('#previewbtn').click(function () {

    img = $('#spimagehide').attr('src')

    title = $('#titleid').val()

    // text =$('#text').val()

    var data = ckeditor1.getData();

    $('#preview-img').attr('src', img)

    $('.preview-head').text(title)

    $('.heading-four').html(data)

})


$('.avaliable-dropdown-item').click(function () {

    var id = $(this).find('input').data('id');

    var channelname = $(this).find('.para').text();

    var count = $(this).find('.para-extralight').text();

    $('.blogDropdown').text(channelname)

    $('.blogDropdown').append('<img class="blogimg" src="/public/img/blog-drop.svg" alt="">');

    $('#blogcount').text(count)

    $('#chanid').val(id)


})

$('#publishbtn').click(function () {

    var url = window.location.href;

    console.log("url", url)

    var pageurl = window.location.search
    const urlpar = new URLSearchParams(pageurl)
    pageno = urlpar.get('page')

    if (url.includes('editentry')) {
        // var eid = url.split('/').pop();

        var urlvalue = url.split('?')[0]
        var eid = urlvalue.split('/').pop()

        console.log("id", eid);
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
    }
    if (title !== '' && data !== '' && img !== '') {
        console.log("kkkkk")
        var flag = channelfieldvalidation()

        channelfieldkeyup()

        if (flag == false) {
            $('#channelModal').modal('show');
        }

        var flag1 = categoryvalidation()


    }
    if (title && data != '' && img != '' && (flag == true) && (flag1 == true)) {

        $.ajax({
            url: "/channel/publishentry/" + eid,
            type: "POST",
            dataType: "json",
            data: { "id": entryId, "cname": channelname, "image": img, "title": title, "status": 1, "text": data, "categoryids": categoryIds, "channeldata": JSON.stringify(channeldata), "seodetails": JSON.stringify(seodetails), csrf: $("input[name='csrf']").val() },
            success: function (result) {

                console.log("result", result)
                // if (result == 0) {


                window.location.href = "/channel/entrylist/?page=" + pageno;

                // }


            }
        })

    }

})






//   //EDIT FUNCTION IN ENTRY PAGE//
$(document).on('click', '#editbtn', function () {

    $('.filter-dropdown-menu ').hide()

})


function CKEDITORS() {

    CKEDITOR.ClassicEditor.create(document.getElementById("text"), {
        toolbar: {
            items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'link', 'code'],
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
        placeholder: languagedata.Channell.plckeditor,
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
                            fetch('/channel/imageupload', {
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

    // console.log("flag",flag1)

    // if (flag1==true){

    //     $('#categoryModal').modal('hide');
    // }
    $('.category-unselect-btn').each(function () {
        var categoryId = $(this).attr('data-id');
        categoryIds.push(categoryId);
    });

    if (categoryIds.length !== 0) {

        notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Channell.categoryerr + '</span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
        $('#categoryModal').modal('hide');
        $("#category-error").hide()
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

    //   var flag  = SeoValidation()
    $("#seoform").validate({
        rules: {
            metatitle: {
                customLength: [1, 60]
            },
            metadesc: {
                customLength: [1, 150]
            }
        },
        messages: {
            metatitle: {
                customLength: "*" + languagedata.Channell.metatitlemsg
            },
            metadesc: {
                customLength: "*" + languagedata.Channell.metadescmsg
            }
        }
    });
    var formcheck = $("#seoform").valid();
    if (formcheck == true) {
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

            notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Channell.seoerr + '</span></div>';
            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds

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

//**description focus function */
const Desc = document.getElementById('metadesc');
const inputGroup = document.querySelectorAll('.input-group');

Desc.addEventListener('focus', () => {

    Desc.closest('.input-group').classList.add('input-group-focused');
});
Desc.addEventListener('blur', () => {
    Desc.closest('.input-group').classList.remove('input-group-focused');
});

$(document).on('click', '#seocancelbtn', function () {

    $('#metatitle').val('')
    $('#metadesc').val('')
    $('#metakey').val('')
    $('#metaslug').val('')

})

$(document).ready(function () {

    var url = window.location.href;

    var id = url.split('/').pop();
    console.log("id", id);
    if (id === 'newentry') {
        console.log("hhkkkk")
        var dropdownItems = $('.avaliable-dropdown-item');
        dropdownItems.eq(0).click();
    }
});

$(document).on('click', '.avaliable-dropdown-item ', function () {


    channelid = $('#chanid').val()

    console.log("jiddd", channelid)

    $('.categorypdiv').hide();

    $('.secid').empty();

    $.ajax({
        url: "/channel/channelfields/" + channelid,
        type: "GET",
        dataType: "json",
        data: { "id": channelid, csrf: $("input[name='csrf']").val() },
        success: function (result) {
            console.log(result, "resultddddd")
            cateogryarr = []

            //category section//

            if (result.SelectedCategory !== null) {
                result.SelectedCategory.forEach(function (category) {

                    newcategory = category.split(',').join('');

                    cateogryarr.push(newcategory)



                })
            }
            console.log(cateogryarr, "arrayyyyy")

            if (cateogryarr !== null) {

                cateogryarr.forEach(function (category) {

                    $('#category-' + category).show();

                })
            }

            var count = 0

            $('.categorypdiv').each(function () {

                if ($(this).css('display') != "none") {

                    count = count + 1

                }
            })


            $('#avacatcount').text(count);

            // var count = $('.categorypdiv:visible').length;
            // console.log("count",count)
            // $('#avacatcount').text(count);

            //channel field section//

            if (result.Section != null) {
                result.Section.forEach(function (section) {

                    $('.secid').append(`<div id="${section.SectionId}"><h6 class="sechead">${section.SectionName}</h6></div>`)

                })
                if (result.FieldValue != null) {
                    result.FieldValue.forEach(function (fieldValue) {


                        if (fieldValue.MasterFieldId == 4) {
                            $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date date-start" name="todate" data-name="${fieldValue.FieldName}"  id="dateend-${fieldValue.FieldId}" data-date="${fieldValue.DateFormat}" data-time="${fieldValue.TimeFormat}" placeholder="${fieldValue.DateFormat}" value="" />
                    
                    
                    </div>
                    <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                  </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 5) {
                            $(`#${fieldValue.SectionId}`).append(` 
            
            <div class="input-group  user-drop-down getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" id="dropoption" style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        option
                                        </a>
                                        <div class="ig-row">
                                        <input type="hidden" class="fieldid" id="${fieldValue.FieldId}"  value=""  >
                                        </div>
                                        <label class="error" id="opterr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
        
                                        <div class="dropdown-menu" id="drop${fieldValue.FieldId}" aria-labelledby="triggerId">
                
                                            
                                        </div>
                                      
                                    </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 6) {
                            $(`#${fieldValue.SectionId}`).append(` 
                    <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                 
                    <label class="input-label" id="datelabel"  data-id="${fieldValue.FieldId}" for="">${fieldValue.FieldName}</label>
                    <div class="ig-row date-input">
                      <input type="text"class="date"  data-date="${fieldValue.DateFormat}" id="${fieldValue.FieldId}"  value=""  placeholder="${fieldValue.DateFormat}" />
                     
                    
                    </div>
                    <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                  </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 9) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" id="radiogrb" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
                <div class="ig-row">
                <input type="hidden" class="radioval" value="">
                <div class="button-col entry-radio-row flexx" id="radio${fieldValue.FieldId}">
              
            </div>
            <label class="error" id="radioerr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
            </div>
            </div>
            
               
                `);
                        }
                        if (fieldValue.MasterFieldId == 8) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" >
                <label  class="input-label"  data-id="${fieldValue.FieldId}" id="textarealabel" for="">${fieldValue.FieldName}</label>
             
                <div class="ig-row">
                  <textarea placeholder="`+ languagedata.Channell.pltexthere + `" id="chtextarea"></textarea>
                </div>
                <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
              </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 2) {
                            $(`#${fieldValue.SectionId}`).append(` 
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}"  data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text"  placeholder="`+ languagedata.Channell.pltexthere + `" id="chtextbox" />
                      </div>
                      <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>

                      <label class="error cerr" style="display: none">* Please Enter ${fieldValue.CharacterAllowed} character only</label>
                    </div>
                `);
                        }
                        if (fieldValue.MasterFieldId == 7) {
                            $(`#${fieldValue.SectionId}`).append(`  
                <div class="input-group getvalue" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}">
                      <label for="" class="input-label"   data-id="${fieldValue.FieldId}" id="textboxlabel">${fieldValue.FieldName}</label>
                      <div class="ig-row">
                        <input type="text" placeholder="`+ languagedata.Channell.pltexthere + `"  />
                      </div>
                      <label class="error" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
                    </div>
                `);
                        }

                        if (fieldValue.MasterFieldId == 10) {
                            $(`#${fieldValue.SectionId}`).append(`
               <div class="input-group  user-drop-down getvalue" id="checkdrop" data-id="${fieldValue.Mandatory}" data-char="${fieldValue.CharacterAllowed}" style="margin-bottom: 20px">
        
                                        <label for="" class="input-label"  data-id="${fieldValue.FieldId}" id="slabel">${fieldValue.FieldName}</label>
        
                                        <a class="dropdown-toggle" type="button" id="triggerId" data-bs-toggle="dropdown"
                                            aria-haspopup="true" aria-expanded="false">
                                        option
                                        </a>
                                      
                                        <input type="hidden" class="checkfieldid" id="${fieldValue.FieldId}"  value=""  >
                                      
                                        <label class="error" id="opterrr" style="display: none">*`+ languagedata.Channell.errmsg + `</label>
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

    channeldata = [];

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

        notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span> ' + languagedata.Channell.saveerr + ' </span></div>';
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds

    }

    if (flag == true) {
        $('#channelModal').modal('hide');
    }

})

$(document).on('click', '#addcategory', function () {
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

        $('#categoryModal').modal('show')
    }
})

$(document).on('click', '#channelfield', function () {
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


    console.log("congiggg")
    channelid = $('#chanid').val()

    $(this).attr('href', '/channels/editchannel/' + channelid)

})

function channelfieldvalidation() {

    var flag = true;
    $('.getvalue').each(function () {


        if ((($(this).find('input').val() === '') && ($(this).data('id') == 1)) || (($(this).find('textarea').val() === '') && ($(this).data('id') == 1))) {

            $(this).addClass('input-group-error');
            $(this).find('.error').show()
            flag = false
        }

        // charallowed = $(this).data('char')

        // console.log(charallowed, "yyyy")
        // if (charallowed < $(this).find('input').val().maxLength) {

        //     console.log("gggg")
        //     $(this).addClass('input-group-error');
        //     $(this).find('.cerr').show()
        //     flag = false
        // }
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
                $(this).parent('.ig-row').siblings('.error').show();


            }

        })
        $(this).find('.date').change(function () {
            if (($(this).val() !== '')) {

                $(this).closest('.input-group').removeClass('input-group-error');
                $(this).parent('.ig-row').siblings('.error').hide();
            } else {

                $(this).closest('.input-group').addClass('input-group-error');
                $(this).parent('.ig-row').siblings('.error').show();


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

    $('.selectedcategorydiv').remove()
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