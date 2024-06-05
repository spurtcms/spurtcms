// dropdown filter input box search
$("#searchdropdownrole").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()
  
    $(".dropdown-role .dropdown-filter-roles .dropdown-item ").each(function (index,element) {

        var title = $(this).text().toLowerCase()

        if (title.includes(keyword)) {

            $(element).show()

            $("#nodatafounddesign").hide()

        } else {

            $(element).hide()

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
    var html = '<div class="upload-logo" id="upload-logo"><img src="" alt=""><input type="hidden" id="logo-input" name="logo_imgpath" value="/public/img/logo.svg"><p class="para">'+languagedata.Personalize.browsecontent+'</p><div class="upload-button"><button class="btn-reg btn-xs primary theme-color browse-btn" type="button" data-bs-toggle="modal" data-bs-target="#addnewimageModal"> '+languagedata.browse+' </button><input type="hidden"></div></div>'
    $(html).insertAfter('.myprouploaddiv')
})

$(document).on('click', '#delete-expandlogo', function () {
    $("#upload-expandlogo").remove()
    var html = '<div class="upload-logo" id="upload-expandlogo"><img src="" alt=""><input type="hidden" id="expandlogo-input" name="expandlogo_imgpath" value="/public/img/logo-bg.svg"><p class="para">'+languagedata.Personalize.browsecontent+'</p><div class="upload-button"><button class="btn-reg btn-xs primary theme-color browse-btn" type="button" data-bs-toggle="modal" data-logo="1" data-bs-target="#addnewimageModal"> '+languagedata.browse+' </button><input type="hidden"></div></div>'

    $(html).insertAfter('.expandlogodiv')
})


$(document).on('click', '#upload-expandlogo>.upload-button>.browse-btn', function () {

    console.log($(this).attr('data-logo'),"value")
    $("#logo-input").val($(this).attr('data-logo'))
})