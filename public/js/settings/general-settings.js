// dropdown filter input box search
$("#searchdropdownrole").keyup(function () {

    var keyword = $(this).val()

    keyword = $.trim(keyword.toLowerCase())

    $("#timezoneDropdown .timeZoneVal ").each(function (index, element) {

        var title = $(this).text().toLowerCase()

        console.log(title);


        if (title.includes(keyword)) {

            // $(element).show()

            // $("#nodatafounddesign").hide()

            $(element).closest('.chk-group').show() // Show the parent element that contains the radio button
            $("#nodatafounddesign").hide()

        } else {

            // $(element).hide()

            // if ($('.dropdown-filter-roles .dropdown-item:visible').length == 0) {

            //     $("#nodatafounddesign").show()

            // }

            $(element).closest('.chk-group').hide() // Hide the parent element that contains the radio button
            if ($('.dropdown-filter-roles .dropdown-item:visible').length == 0) {
                $("#nodatafounddesign").show()
            }
        }
    })

})

$(".drop-timezone").on("click", function () {
    var text = $(this).text()
    var id = $(this).attr("data-id")
    $(this).parents().siblings("a").text(text)
    $(this).parents().siblings("input[name='timezon']").val(text)
})


$(document).on('click', '#delete-logo', function () {

    $("#upload-logo").remove()
    var html = '<div class="upload-logo" id="upload-logo"><img src="" alt=""><input type="hidden" id="logo-input" name="logo_imgpath" value="/public/img/logo.svg"><p class="para">' + languagedata.Personalize.browsecontent + '</p><div class="upload-button"><button class="btn-reg btn-xs primary theme-color browse-btn" type="button" data-bs-toggle="modal" data-bs-target="#addnewimageModal"> ' + languagedata.browse + ' </button><input type="hidden"></div></div>'
    $(html).insertAfter('.myprouploaddiv')
})

$(document).on('click', '#delete-expandlogo', function () {
    $("#upload-expandlogo").remove()
    var html = '<div class="upload-logo" id="upload-expandlogo"><img src="" alt=""><input type="hidden" id="expandlogo-input" name="expandlogo_imgpath" value="/public/img/logo-bg.svg"><p class="para">' + languagedata.Personalize.browsecontent + '</p><div class="upload-button"><button class="btn-reg btn-xs primary theme-color browse-btn" type="button" data-bs-toggle="modal" data-logo="1" data-bs-target="#addnewimageModal"> ' + languagedata.browse + ' </button><input type="hidden"></div></div>'

    $(html).insertAfter('.expandlogodiv')
})


$(document).on('click', '#upload-expandlogo>.upload-button>.browse-btn', function () {

    console.log($(this).attr('data-logo'), "value")
    $("#logo-input").val($(this).attr('data-logo'))
})


$(document).on('click', '.timezoneInputs', function () {
    if ($(this).prop('checked')) {
        $('#timezoneText').text($(this).val())
    }
})

$(document).on('click', '#settingsUpdate[data-status=1]', function () {

    var formData = new FormData
    formData.append('csrf', $('#csrf-value').val())
    formData.append('companyname', $('#companyName').val())
    formData.append('expandlogo_imgpath', $('#prof-crop').attr('src'))

    var langVal, dateFormat, timeFormat, timeZones;
    $('.langInput').each(function (index, element) {

        if ($(this).prop('checked')) {
            langVal = $(this).val();
        }
    });
    formData.append('language', langVal)

    $('.dateFormat').each(function (index, element) {

        if ($(this).prop('checked')) {
            dateFormat = $(this).val();
        }
    });
    formData.append('dateformat', dateFormat)

    $('.timeFormat').each(function (index, element) {

        if ($(this).prop('checked')) {
            timeFormat = $(this).val();
        }
    });
    formData.append('timeformat', timeFormat)



    // $('.timezoneInputs').each(function (index, element) {

    //     if ($(this).prop('checked')) {
    //         timeZones = $(this).val();
    //     }
    // });
    timeZones = $.trim($('#timezoneText').text())
    formData.append('timezon', timeZones)

    $.ajax({
        url: "/settings/general-settings/update",
        type: "POST",
        data: formData,
        contentType: false, // Prevent jQuery from setting Content-Type
        processData: false,
        success: function (result) {
            if (result.status == 1) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >General Settings Updated Successfully</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000);

                window.location.href = "/settings/general-settings/"

            } else {
                setCookie("Alert-msg", "success", 1)
            }
        }
    })
})


//for users to restrict image change
$(document).on('click', "#logoLabel", function (event) {

    var roleId = $(this).attr('data-role-id')


    if (roleId != 2 && roleId != 1) {
        event.preventDefault()


    }
})

$(document).on('click', ".checkRole", function (event) {

    var roleId = $(this).attr('data-role-id')


    if (roleId != 2 && roleId != 1) {
        event.preventDefault()


    }
})
const isChanged = false;
const roleId = $('#genSettingsCancelBtn').attr('data-role-id')


$(document).ready(function (e) {
    if (!isChanged) {
        $('#genSettingsCancelBtn').attr('href', '#')
        $('#genSettingsCancelBtn').addClass('cursor-not-allowed')
    }
})

// only enabling the cancel button if any thing changes in the general settings page

$(document).on('change', '#companyName,#logoLabel', function () {
    $('#genSettingsCancelBtn').attr('href', '/settings/general-settings/')
    $('#genSettingsCancelBtn').removeClass('cursor-not-allowed')
})

$('.checkRole,.timeZone').click(function () {
    $('#genSettingsCancelBtn').attr('href', '/settings/general-settings/')
    $('#genSettingsCancelBtn').removeClass('cursor-not-allowed')
})



