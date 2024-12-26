var languagedata

var selectedcheckboxarr = []
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    console.log("lang", languagepath);

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

})

$(document).on('click', '#cancel-notify', function () {
    delete_cookie("Alert-msg");
    $(".toast-msg").remove();
})

//setting the json file name after update 

$(document).on('change', '#jsonFile', function () {

    $('#jsonFileErr').addClass('hidden')
    $('#jsonFile-error').addClass('hidden')

    var jsonFile = $(this)[0].files[0]

    var fileName = jsonFile.name

    var fileExtension = fileName.split(".").pop().toLowerCase()

    console.log(fileExtension);

    if (fileExtension == "json") {
        console.log("test");

        $('#jsonDesc').text(fileName)

    } else {

        $('#jsonFileErr').removeClass('hidden')
    }

})

//after uploading flag image functionality for delete button
$(document).on('click', '#deleteFlag', function (e) {
    e.preventDefault()

    $('canvas[class=cr-image]').css('opacity', '0')
    $("#changepicModal").modal('hide');
    $('#crop-container').removeClass('croppie-container').empty()
    $('#flagSpan').removeClass('hidden')
    $('#flaggFileLabel').removeClass('hidden')
    $('#flagPara').removeClass('hidden')
    $('#prof-crop').removeClass('h-[180px]')
    $('#prof-crop').attr('src', "")
    $('#cropData').attr('value', "")
    $('.deleteFlag').remove();
    $('.hover-delete-img').remove();

})

//on clicking lang create modal close button, modal will return to mormal 

var originalModal = $('#LanguageModal').clone()

$(document).on('click', '#langCancelBtn', function () {
    $('#LanguageModal').html(originalModal.html())
    $('canvas[class=cr-image]').css('opacity', '0')
    $("#changepicModal").modal('hide');
    $('#imageSpace').removeClass('croppie-container').empty()
})


var langCodeArr = [];
$(document).ready(function () {

    $('.languageCode').each(function (index, element) {
        langCodeArr.push($.trim($(this).text()));

    });

    console.log(langCodeArr);

})

$(document).on('click', '.langCodeGrp', function () {
    var langCode = $(this).attr('data-code')

    $('#lang_code-error').addClass('hidden')
    if (langCodeArr.includes(langCode)) {

        $('#langCodeErr').removeClass('hidden')
        $('#langCodePara').text("")

    } else {
        $('#lang_code').val(langCode)
        $('#langCodePara').text(langCode)
        $('#langCodeErr').addClass('hidden')
    }
})

$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $("#languagesearch").focus().select();
    }
});

$(document).ready(function () {


    $(document).on('click', '#langCreateBtn', function () {

        jQuery.validator.addMethod("duplicatelangname", function (value) {

            var result;
            var id = $("#langid").val()

            $.ajax({
                url: "/settings/languages/checklanguagename",
                type: "POST",
                async: false,
                data: { "langname": value, "id": id, csrf: $("#csrfValue").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    console.log("data", data);
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })

        $('form[name=language-form]').validate({
            ignore: [],
            rules: {
                lang_name: {
                    required: true,
                    duplicatelangname: true
                },
                lang_code: {
                    required: true,
                },
                lang_json: {
                    required: true,
                },
                flag_imgPath: {
                    required: true,
                }
            },
            messages: {
                lang_name: {
                    required: languagedata.Languages.langnamevalid,
                    duplicatelangname: languagedata.Languages.duplicatenameerr
                },
                lang_code: {
                    required: languagedata.Languages.langcodevalid,
                },
                lang_json: {
                    required: languagedata.Languages.seljsonvalid,
                },
                flag_imgPath: {
                    required: languagedata.Languages.selflagimgvalid,
                },
            }
        })

        var isFormValid = $('form[name=language-form]').valid()

        if (!isFormValid) {

            $('.input-group').addClass('input-group-error')

            langcode = $("#triggerId").text()

            if (langcode != "") {

                $('.input-group').removeClass('input-group-error')

            }

            if ($('input[name=lang_name]').hasClass('valid')) {

                $('.input-group').removeClass('input-group-error')
            }

        } else {

            $('.input-group').removeClass('input-group-error')

            $('form[name=language-form]').submit()
        }

    })

    $('.input-group input').keyup(function () {

        $(this).parents('.input-group').removeClass('input-group-error')

    })

})

//search language

$(document).on('keypress', '#langSearchBar', function (e) {

    if (e.which == 13) {
        var searchText = $('#langSearchBar').val()
        console.log(searchText, "searchText");


        if (searchText != "") {
            $('#langSearchlink').attr('href', "/settings/languages/?keyword=" + searchText)
            $('#langSearchBar').attr('value', searchText)
            $('#langSearchlink').get(0).click()
        } else {
            window.location.href = "/settings/languages/"
        }

    }
})

/*search redirect home page */

$(document).on('keyup', '#langSearchBar', function (e) {
    var searchVal = $('#langSearchBar').attr('value')
    if (e.which == 8 && searchVal != "") {
        var searchKey = $('#langSearchBar').val()
        if (searchKey == "") {
            window.location.href = "/settings/languages/"
        }
    }
})

$(document).on('click', '#setdefault', function () {

    if ($(this).is(':checked')) {

        $(this).val("1")
    } else {

        $(this).val("0")
    }
})


$(document).on('click', '.langEditBtn', function () {
    var langId = $(this).attr('data-id')
    console.log(langId);

    $('#langId').val(langId)


    $.ajax({
        url: '/settings/languages/editlanguage/' + langId,
        type: 'GET',
        dataType: 'JSON',
        success: function (data) {
            console.log("data", data)
            $('form[name=language-form]').validate({
                ignore: [],
                rules: {
                    lang_name: {
                        required: true,
                        duplicatelangname: false
                    },
                    lang_code: {
                        required: true,
                    },
                    lang_json: {
                        required: false,
                    },
                    flag_imgPath: {
                        required: false,
                    }
                },
                messages: {
                    lang_name: {
                        required: languagedata.Languages.langnamevalid,
                        duplicatelangname: languagedata.Languages.duplicatenameerr
                    },
                    lang_code: {
                        required: languagedata.Languages.langcodevalid,
                    },
                    lang_json: {
                        required: languagedata.Languages.seljsonvalid,
                    },
                    flag_imgPath: {
                        required: languagedata.Languages.selflagimgvalid,
                    },
                }
            })
            $('#LanguageModal #staticBackdropLabel').text(languagedata.updatelanguage)
            $('#LanguageModal #langCreateBtn').text(languagedata.update)
            $('form[name=language-form]').attr("action", data.UploadUrl)

            if (data.Language.DefaultLanguageId != 0 && data.Language.DefaultLanguageId == data.Language.Id) {
                $('#LanguageModal #setdefault').val("1").attr('checked', true)
            } else {
                $('#LanguageModal #setdefault').val("0").attr("checked", false)
            }



            $('#langName').val(data.Language.LanguageName)
            $('#langCodePara').text(data.Language.LanguageCode)
            $("#lang_code").val(data.Language.LanguageCode)
            $('#languageCodeDropdown').attr('data-bs-toggle', '')
            $('#languageCodeDropdown').removeClass("bg-[url('/public/img/property-arrow.svg')]")
            $('#langStatus').val(data.Language.IsStatus)
            $('#setdefault').val(data.Language.IsDefault)

            var jsonPath = data.Language.JsonPath

            var jsonFileName = jsonPath.split("/").pop()

            $('#jsonDesc').text(jsonFileName)
            console.log(data.Language.ImagePath);
            $('#flagSpan').addClass('hidden')
            $('#flaggFileLabel').addClass('hidden')
            $('#flagPara').addClass('hidden')
            $('#prof-crop').addClass('h-[180px]')

            $('#LanguageModal #prof-crop').attr('src', data.Language.ImagePath)
            $('#cropData').attr('value', data.Language.ImagePath)
            $('#flagDiv').append('<div class="hover-delete-img "></div>');
            $('#flagDiv').append('<button class="deleteFlag"><img id="deleteFlag" src="/public/img/deleteWhiteIcon.svg" alt=""></button>')

        }

    })
})

// status change on clicking the status button
// $(document).on('click', '.langStatusBtn', function () {
//     if ($(this).prop('checked')) {

//         $(this).attr('value', '1')
//     } else {
//         $(this).attr('value', '0')
//     }

//     var lang_id = $(this).attr('data-id')

//     var isactive = $(this).val();
//     $.ajax({
//         url: '/settings/languages/languageisactive',
//         type: 'POST',
//         async: false,
//         data: {
//             id: lang_id,
//             isactive: isactive,
//             csrf: $("input[name='csrf']").val()
//         },
//         dataType: 'json',
//         cache: false,
//         success: function (result) {
//             console.log(result);

//             if (result.status == true) {
//                 notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast.languageupdated + `</p ></div ></div ></li></ul> `;
//                 $(notify_content).insertBefore(".header-rht");
//                 setTimeout(function () {
//                     $('.toast-msg').fadeOut('slow', function () {
//                         $(this).remove();
//                     });
//                 }, 5000);

//             } else {

//                 notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/toast-error.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + languagedata.Toast.languageupdated + `</p></div></div> </li></ul>`;
//                 $(notify_content).insertBefore(".header-rht");
//                 setTimeout(function () {
//                     $('.toast-msg').fadeOut('slow', function () {
//                         $(this).remove();
//                     });
//                 }, 5000);

//             }


//         }
//     });
// })


$(document).on('click', '.langDeleteBtn', function () {

    var langId = $(this).attr("data-id")

    console.log("langid", langId);

    $('.deltitle').text(languagedata.Languages.deletelanguage + "?")

    $('.deldesc').text(languagedata.Languages.dellanguage)

    $('.delname').text($(this).parents('tr').find('td:first>.flexx>h4').text())

    $('#delid').attr('href', "/settings/languages/deletelanguage/" + langId)

    // $('#centerModal').modal('show')
})

$('#addnewimageModal').on('hide.bs.modal', function () {

    $('#laguangeModal').css('opacity', 1)

})


$(document).on('click', '.selectcheckbox', function () {

    console.log(selectedcheckboxarr);

    var languageId = $(this).attr('data-id')

    var status = $(`#toggle${languageId}`).val();

    if ($(this).prop('checked')) {

        selectedcheckboxarr.push({ "languageId": languageId, "status": status })

    } else {

        const index = selectedcheckboxarr.findIndex(item => item.languageId === languageId);

        if (index !== -1) {

            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {
        $('.selected-numbers').removeClass("hidden")

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus
        var img;

        if (selectedcheckboxarr[0].status === '1') {

            setstatus = "Deactive";

            img = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = "Active";

            img = "/public/img/Active.svg";

        }

        var htmlContent = '';

        if (allSame) {

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + setstatus;

        } else {

            htmlContent = '';

        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

        if (!allSame) {

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

            $('.unbulishslt').html("")
        } else {

            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')
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

    $('#Check').prop('checked', allChecked);

    console.log(selectedcheckboxarr, "checkkkk")
})


//  //ALL CHECKBOX CHECKED FUNCTION//



$(document).on('click', '#Check', function () {

    selectedcheckboxarr = []

    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            var languageId = $(this).attr('data-id')

            var status = $(`#toggle${languageId}`).val();

            console.log(status, "state");

            selectedcheckboxarr.push({ "languageId": languageId, "status": status })
        })


        $('.selected-numbers').removeClass("hidden")

        var allSame = selectedcheckboxarr.every(function (item) {

            return item.status === selectedcheckboxarr[0].status;
        });

        console.log(allSame, "allsome");



        var img
        if (selectedcheckboxarr[0].status === '1') {

            setstatus = "Deactive";

            img = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = "Active";

            img = "/public/img/Active.svg";
        }

        // }
        var htmlContent = '';

        if (allSame) {

            htmlContent = '<img src="' + img + '" style="width: 14px; height: 14px;">' + setstatus;

            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        } else {

            htmlContent = '';

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

    } else {
        $('.selected-numbers').addClass("hidden")

        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);
    }

})

// ------------------------------------------------------------------


$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text("Delete Languages?")

        $('#content').text('Are you sure want to delete selected Languages?')

    } else {

        $('.deltitle').text("Delete Languages?")

        $('#content').text('Are you sure want to delete selected Languages?')
    }

    $("#delid").text($(this).text());
    $('#delid').addClass('checkboxdelete')
})

$(document).on('click', '#unbulishslt', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text($(this).text() + " " + "Languages?")

        $('#content').text("Are you sure want to " + $(this).text() + " " + "selected Languages?")
    } else {

        $('.deltitle').text($(this).text() + " " + "Languages?")

        $('#content').text("Are you sure want to " + $(this).text() + " " + "selected Languages?")
    }
    // $('#delid').text(setstatus)
    $("#delid").text($(this).text());

    $('#delid').addClass('selectedunpublish')

})

//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {

    var url = window.location.href;

    console.log("url", url)

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/settings/languages/multiselectlanguagedelete',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "languageids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {

            console.log(data, "result")

            if (data.value == true) {

                setCookie("get-toast", "Language Deleted Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})
//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass("hidden")

})

//multi select active and deactive function//

$(document).on('click', '.selectedunpublish', function () {

    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/settings/languages/multiselectlanguagestatus',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "languageids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "languageupdated")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})

// stop changing default language for users and allow only for tenants

$(document).on('click', ".langStatusBtn", function (event) {

    var roleId = $(this).attr('data-role-id')


    if (roleId != 2 && roleId != 1) {
        event.preventDefault()
        notify_content = `<ul class=" warn-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="/ flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/toast-error.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + "please contact admin for changing default language" + `</p></div></div> </li></ul>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.warn-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000);

    }
})

// changing default language by tenant 

function SetDefaultLang(langId, isDefault, roleId) {

    if (roleId != 2 && roleId != 1) {
        notify_content = `<ul class=" warn-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/toast-error.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + "please contact admin for changing default language" + `</p></div></div> </li></ul>`;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.warn-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000);

        return
    }

    $.ajax({
        url: '/settings/languages/setdefaultlanguage',
        type: 'POST',
        dataType: 'json',
        cache: false,
        data: {
            langId: langId,
            isDefault: isDefault,
            csrf: $("input[name='csrf']").val()
        },
        success: function (result) {
            if (result.value) {
                setCookie("get-toast", "Language Updated Successfully")
                window.location.href = "/settings/languages/"
            }

        }
    })


}

$(document).on('click', '.langStatusBtn', function () {
    var langId = $(this).attr("data-id")
    var dataRoleId = $(this).attr("data-role-id")
    var isDefault = $(this).val()


    if ($(this).is(':checked')) {
        $(this).val("1")
        $('.langStatusBtn').each(function (index, element) {
            // element == this
            if (dataRoleId == 2 || dataRoleId == 1) {
                if ($(element).attr('data-id') != langId) {
                    $(element).prop('checked', false)
                }

            }

        });
        if ($(this).val() == 1) {
            SetDefaultLang(langId, isDefault, dataRoleId)
        }
    } else {
        console.log("ulla vra");

        var isChecked = false
        $('.langStatusBtn').each(function (index, element) {
            // element == this
            if ($(this).is(':checked')) {
                isChecked = true
            }
        });
        if (!isChecked) {
            $(this).prop('checked', true)
            $(this).val("1")
        }
    }
    // SetDefaultLang(langId, isDefault, dataRoleId)
})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')

    window.location.href = "/settings/languages/"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
        } else {
            $(".Closebtn").addClass("hidden")
        }
    });
})

$(document).on("click", ".langHoverIcon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})



//old Code ......................... down

// $(document).on('click', '.selectcheckbox', function () {

//     languageid = $(this).attr('data-id')

//     var status = $(this).parents('td').siblings('td').find('.tgl-light').prop('checked');

//     console.log(status, "status")

//     var sstatus


//     if (status == true) {

//         sstatus = '1'

//     } else {

//         sstatus = '0'

//     }

//     if ($(this).prop('checked')) {

//         selectedcheckboxarr.push({ "languageid": languageid, "status": sstatus })

//     } else {

//         const index = selectedcheckboxarr.findIndex(item => item.languageid === languageid);

//         if (index !== -1) {

//             console.log(index, "sssss")
//             selectedcheckboxarr.splice(index, 1);
//         }

//         $('#Check').prop('checked', false)

//     }


//     if (selectedcheckboxarr.length != 0) {

//         $('.selected-numbers').show()

//         var allSame = selectedcheckboxarr.every(function (item) {
//             return item.status === selectedcheckboxarr[0].status;
//         });

//         var setstatus
//         var img;

//         if (selectedcheckboxarr[0].status === '1') {

//             setstatus = "Deactive";

//             img = "/public/img/In-Active (1).svg";

//         } else if (selectedcheckboxarr[0].status === '0') {

//             setstatus = "Active";

//             img = "/public/img/Active (1).svg";

//         }

//         var htmlContent = '';

//         if (allSame) {

//             htmlContent = '<img src="' + img + '">' + setstatus;

//         } else {

//             htmlContent = '';

//         }

//         $('#unbulishslt').html(htmlContent);

//         $('.checkboxlength').text(selectedcheckboxarr.length + " " + 'items selected')

//         if (!allSame) {

//             $('#seleccheckboxdelete').removeClass('border-end')

//             $('.unbulishslt').html("")
//         } else {

//             $('#seleccheckboxdelete').addClass('border-end')
//         }


//     } else {

//         $('.selected-numbers').hide()
//     }

//     var allChecked = true;

//     $('.selectcheckbox').each(function () {

//         if (!$(this).prop('checked')) {

//             allChecked = false;

//             return false;
//         }
//     });

//     $('#Check').prop('checked', allChecked);

//     console.log(selectedcheckboxarr, "checkkkk")
// })

// //ALL CHECKBOX CHECKED FUNCTION//

// $(document).on('click', '#Check', function () {

//     selectedcheckboxarr = []

//     var isChecked = $(this).prop('checked');

//     if (isChecked) {

//         $('.selectcheckbox').prop('checked', isChecked);

//         $('.selectcheckbox').each(function () {

//             languageid = $(this).attr('data-id')

//             var status = $(this).parents('td').siblings('td').find('.tgl-light').prop('checked');

//             var sstatus

//             if (status == true) {

//                 sstatus = '1'

//             } else {

//                 sstatus = '0'

//             }


//             selectedcheckboxarr.push({ "languageid": languageid, "status": sstatus })
//         })

//         $('.selected-numbers').show()

//         var allSame = selectedcheckboxarr.every(function (item) {

//             return item.status === selectedcheckboxarr[0].status;
//         });

//         var setstatus

//         var img


//         if (selectedcheckboxarr.length != 0) {

//             if (selectedcheckboxarr[0].status === '1') {

//                 setstatus = "Deactive";

//                 img = "/public/img/In-Active (1).svg";

//             } else if (selectedcheckboxarr[0].status === '0') {

//                 setstatus = "Active";

//                 img = "/public/img/Active (1).svg";

//             }
//         }

//         var htmlContent = '';

//         if (allSame) {

//             htmlContent = '<img src="' + img + '">' + setstatus;

//             $('#seleccheckboxdelete').addClass('border-end')

//         } else {

//             htmlContent = '';

//             $('#seleccheckboxdelete').removeClass('border-end')

//         }

//         $('#unbulishslt').html(htmlContent);

//         $('.checkboxlength').text(selectedcheckboxarr.length + " " + 'items selected')

//     } else {


//         selectedcheckboxarr = []

//         $('.selectcheckbox').prop('checked', isChecked);

//         $('.selected-numbers').hide()
//     }

//     if (selectedcheckboxarr.length == 0) {

//         $('.selected-numbers').hide()
//     }
// })

// $(document).on('click', '#seleccheckboxdelete', function () {

//     if (selectedcheckboxarr.length > 1) {

//         $('.deltitle').text("Delete Languages?")

//         $('#content').text('Are you sure want to delete selected Languages?')

//     } else {

//         $('.deltitle').text("Delete Language?")

//         $('#content').text('Are you sure want to delete selected Language?')
//     }


//     $('#delete').addClass('checkboxdelete')
// })

// $(document).on('click', '#unbulishslt', function () {

//     if (selectedcheckboxarr.length > 1) {

//         $('.deltitle').text($(this).text() + " " + "Languages?")

//         $('#content').text("Are you sure want to " + $(this).text() + " " + "selected Languages?")
//     } else {

//         $('.deltitle').text($(this).text() + " " + "Language?")

//         $('#content').text("Are you sure want to " + $(this).text() + " " + "selected Language?")
//     }

//     $('#delete').addClass('selectedunpublish')

// })

// //MULTI SELECT DELETE FUNCTION//
// $(document).on('click', '.checkboxdelete', function () {

//     var pageurl = window.location.search

//     const urlpar = new URLSearchParams(pageurl)

//     pageno = urlpar.get('page')

//     $('.selected-numbers').hide()
//     $.ajax({
//         url: '/settings/languages/multiselectlanguagedelete',
//         type: 'post',
//         dataType: 'json',
//         async: false,
//         data: {
//             "languageids": JSON.stringify(selectedcheckboxarr),
//             csrf: $("input[name='csrf']").val(),
//             "page": pageno


//         },
//         success: function (data) {

//             console.log(data, "result")

//             if (data.value == true) {

//                 setCookie("get-toast", "Language Deleted Successfully")

//                 window.location.href = data.url
//             } else {

//                 setCookie("Alert-msg", "Internal Server Error")

//             }

//         }
//     })

// })
// //Deselectall function//

// $(document).on('click', '#deselectid', function () {

//     $('.selectcheckbox').prop('checked', false)

//     $('#Check').prop('checked', false)

//     selectedcheckboxarr = []

//     $('.selected-numbers').hide()

// })

// //multi select active and deactive function//

// $(document).on('click', '.selectedunpublish', function () {

//     var pageurl = window.location.search

//     const urlpar = new URLSearchParams(pageurl)

//     pageno = urlpar.get('page')

//     $('.selected-numbers').hide()
//     $.ajax({
//         url: '/settings/languages/multiselectlanguagestatus',
//         type: 'post',
//         dataType: 'json',
//         async: false,
//         data: {
//             "languageids": JSON.stringify(selectedcheckboxarr),
//             csrf: $("input[name='csrf']").val(),
//             "page": pageno


//         },
//         success: function (data) {

//             console.log(data, "result")

//             if (data.value == true) {

//                 setCookie("get-toast", "languageupdated")

//                 window.location.href = data.url
//             } else {

//                 setCookie("Alert-msg", "Internal Server Error")

//             }

//         }
//     })


// })
