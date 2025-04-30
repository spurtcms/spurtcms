var languagedata
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })

})

$(document).ready(function(){
    var theme_clr = $("#color-val").val()
    console.log("theme", theme_clr);
    $(".circle-change").each(function () {
        console.log($(this).css("background-color") == theme_clr, theme_clr, $(this).css("background-color"));
        if ($(this).css("background-color") == theme_clr) {
            // $(this).find(".clr").show();
            $(this).addClass('active')
        }
    })
})

$(document).on('click', '.circle-change', function () {
  $(".circle-change").each(function(){

    $(".circle-change").removeClass('active')
  })
    $(this).addClass('active')
    var color = $(this).css("background-color")
    console.log("clr", color);
    $("#color-change").val(color)
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

$(document).on('click', '#upload-logo>.upload-button>.browse-btn', function () {
    $("#logo-input").val("0")
})
// $(document).on('click', ".upload-folders", function () {
//     var src = $(this).find('img').attr("src");
//     $(".upload-logo").remove()
//     var html = `<div class="upload-logo"><img src="` + src + `" alt=""><input type="hidden" id="logo-input" name="logo_imgpath" value="` + src + `"><button class="delete-flag"  ><img src="/public/img/delete-white-icon.svg" alt=""></button><div class="hover-delete-img"></div>`
//     $(html).insertAfter('#logo-insert')
//     console.log("ok", $("#logo-input").val());
//     $("#addnewimageModal").hide()
//     $('.modal-backdrop').remove()
// })

