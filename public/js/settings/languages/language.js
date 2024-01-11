var languagedata
/** */
$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })

})

$(document).ready(function () {

    $('form[name=language-form]').validate({
        ignore: [],
        rules: {
            lang_name: {
                required: true,
            },
            lang_code: {
                required: true,
            },
            lang_json: {
                required: true,
            },
            flag_imgpath: {
                required: true,
            }
        },
        messages: {
            lang_name: {
                required: "This field is required",
            },
            lang_code: {
                required: "This field is required",
            },
            lang_json: {
                required: "This field is required",
            },
            flag_imgpath: {
                required: "This field is required",
            },
        }
    })

    $('.modal-save>button').click(function(){

        var isFormValid = $('form[name=language-form]').valid()

        if(!isFormValid){

            $('.input-group').addClass('input-group-error')

            if($('input[name=lang_name]').hasClass('valid') && $('input[name=lang_code]').hasClass('valid')){

                $('.input-group').removeClass('input-group-error')
            }

        }else{

            $('.input-group').removeClass('input-group-error')

            $('form[name=language-form]').submit()
        }

    })

    $('.input-group input').keyup(function(){

        $(this).parents('.input-group').removeClass('input-group-error')

    })

})

//search language
$('.search-grp>a').click(function () {

    var keyword = $(this).siblings("input").val().trim()

    if (keyword != "") {

        window.location.href = "/settings/languages/?keyword=" + keyword
    }

})

$('input[name=lang_status][type=checkbox]').click(function () {

    if ($(this).is(':checked')) {

        $(this).val("1")
    } else {

        $(this).val("0")
    }
})

$('input[name=lang_default][type=checkbox]').click(function () {

    if ($(this).is(':checked')) {

        $(this).val("1")
    } else {

        $(this).val("0")
    }
})

$('.file-upload input[name=lang_json]').change(function () {
    var file = this.files[0]
    var filename = $(this).val();
    var ext = filename.split(".").pop().toLowerCase();
    var nameExtract = filename.split("\\").pop()
    if (ext == "json") {
        $(this).css('display', 'none').appendTo($(this).parents(".upload-json")).insertBefore('#lang_json-error')
        $(this).parents('.upload-json').find('.file-upload').remove()
        var uploaded_html = `<div class="uploaded-file">
                <img src="/public/img/folder.svg" alt="">
                <p class="para">`+ nameExtract + `</p>
                <a onclick="RemoveJsonData(this)"><img src="/public/img/delete-icon.svg" alt=""></a>
                </div>`
        $(uploaded_html).insertBefore("#lang_json-error")
        $('#lang_json-error').hide()
    } else {
        $(this).val("")
    }
})

$('.file-upload input[name=flag_imgpath]').click(function(event){
    event.preventDefault()
    $('#laguangeModal').css('opacity',0)
    $('#addnewimageModal').modal('show')
})

// $('.file-upload input[name=flag_imgpath]').change(function () {
//     var file = this.files[0]
//     var filename = $(this).val();
//     var ext = filename.split(".").pop().toLowerCase();
//     var reader = new FileReader();
//     console.log("chk11", $.inArray(ext, ["jpg", "png", "jpeg", "svg"]) != -1);
//     if (($.inArray(ext, ["jpg", "png", "jpeg", "svg"]) != -1)) {
//         $(this).parents('.upload-json').find('.file-upload').remove()
//         reader.onload = function (e) {
//             var imgurl = e.target.result
//             var flag_uploadHtml = `<div class="uploaded-file">
//             <input type="hidden" name="flag_imgpath" value="`+ imgurl + `">
//             <img src="`+ imgurl + `" alt="" class="uploaded-img-flag">
//             <button class="delete-flag" onclick="RemoveFlagImage(this)" ><img src="/public/img/delete-white-icon.svg" alt=""></button>
//             <div class="hover-delete-img"></div>
//             </div>`
//             $(flag_uploadHtml).insertBefore('#flag_imgpath-error')
//             $('#flag_imgpath-error').hide()
//         }
//         reader.readAsDataURL(file)
//     } else {
//         $(this).val("")
//     }
// })

function RemoveJsonData(This) {
    $(This).parent().siblings('input[name=lang_json]').remove()
    $(This).parents('.upload-json').find('.uploaded-file').remove()
    var html = `<div class="file-upload">
    <img src="/public/img/folder.svg" alt="">
    <p class="para">Drag your JSON file or choose from directory</p>
    <div class="upload-button">
        <button class="btn-reg btn-xs primary"> `+languagedata.browse+` </button>
        <input type="file" name="lang_json" onchange="UploadJson(this)">
    </div>
    </div>`
    $(html).insertBefore('#lang_json-error')
}

function RemoveFlagImage(This) {
    $(This).parents('.upload-json').find('.uploaded-file').remove()
    var html = `<div class="file-upload">
    <img src="/public/img/upload-logo.svg" alt="">
    <p class="para">Drag your JSON file or choose from directory</p>
    <div class="upload-button">
        <button class="btn-reg btn-xs primary"> `+languagedata.browse+` </button>
        <input type="file" name="flag_imgpath" onclick="UploadFlagImage(event,this)">
    </div>
    </div>`
    $(html).insertBefore('#flag_imgpath-error')
}

function UploadJson(This) {
    var file = This.files[0]
    var filename = $(This).val();
    var ext = filename.split(".").pop().toLowerCase();
    var nameExtract = filename.split("\\").pop()
    if (ext == "json") {
        $(This).css('display', 'none').appendTo($(This).parents(".upload-json")).insertBefore('#lang_json-error')
        $(This).parents('.upload-json').find('.file-upload').remove()
        var uploaded_html = `<div class="uploaded-file">
          <img src="/public/img/folder.svg" alt="">
          <p class="para">`+ nameExtract + `</p>
          <a onclick="RemoveJsonData(this)"><img src="/public/img/delete-icon.svg" alt=""></a>
          </div>`
        $(uploaded_html).insertBefore('#lang_json-error')
        $('#lang_json-error').hide()
    }
}

function UploadFlagImage(event,This) {
    // var file = This.files[0]
    // var filename = $(This).val();
    // var ext = filename.split(".").pop().toLowerCase();
    // var reader = new FileReader();
    // if (($.inArray(ext, ["jpg", "png", "jpeg", "svg"]) != -1)) {
    //     $(This).parents('.upload-json').find('.file-upload').remove()
    //     reader.onload = function (e) {
    //         var imgurl = e.target.result
    //         var flag_uploadHtml = `<div class="uploaded-file">
    //         <input type="hidden" name="flag_imgpath" value="`+ imgurl + `">
    //         <img src="`+ imgurl + `" alt="" class="uploaded-img-flag">
    //         <button class="delete-flag" onclick="RemoveFlagImage(this)" ><img src="/public/img/delete-white-icon.svg" alt=""></button>
    //         <div class="hover-delete-img"></div>
    //         </div>`
    //         $(flag_uploadHtml).insertBefore('#flag_imgpath-error')
    //         $('#flag_imgpath-error').hide()
    //     }
    //     reader.readAsDataURL(file)
    // }
    event.preventDefault()
    $('#laguangeModal').css('opacity',0)
    $('#addnewimageModal').modal('show')
}

$('.edit-language').click(function () {
    var langId = $(this).attr('data-id')
    var parentTr = $(this).parents('tr')
    $.ajax({
        url: '/settings/languages/editlanguage/' + langId,
        type: 'GET',
        dataType: 'JSON',
        success: function (data) {
            // console.log("data", data)
            $('form[name=language-form]').attr("action",data.UploadUrl)
            $('input[name=lang_name]').val(data.Language.LanguageName)
            $('input[name=lang_code]').val(data.Language.LanguageCode)
            if (data.Language.IsStatus == 1) {
                $('#ckb1[name=lang_status]').val("1").prop('checked', 'true')
            }
            if (data.Language.DefaultLanguageId != 0) {
                $('#ckb2[name=lang_default]').val("1").prop('checked', "true")
            }
            $('.full-upload >.upload-json:first>.file-upload').remove()
            var uploaded_html = `<div class="uploaded-file">
                <input type="hidden"  name="lang_json" value="`+data.Language.JsonPath+`">
                <img src="/public/img/folder.svg" alt="">
                <p class="para">`+ data.Language.JsonPath.split('/').pop() + `</p>
                <a onclick="RemoveJsonData(this)"><img src="/public/img/delete-icon.svg" alt=""></a>
                </div>`
            $(uploaded_html).insertBefore("#lang_json-error")
            $('#lang_json-error').hide()
            $('.full-upload >.upload-json:last>.file-upload').remove()
            var flag_uploadHtml = `<div class="uploaded-file">
            <input type="hidden" name="flag_imgpath" value="`+ data.Language.ImagePath + `">
            <img src="`+ data.Language.ImagePath + `" alt="" class="uploaded-img-flag">
            <button class="delete-flag" onclick="RemoveFlagImage(this)" ><img src="/public/img/delete-white-icon.svg" alt=""></button>
            <div class="hover-delete-img"></div>
            </div>`
            $(flag_uploadHtml).insertBefore('#flag_imgpath-error')
            $('#flag_imgpath-error').hide()
            $('.modal-save>button').text(languagedata.update)
            $('form[name=language-form]').prepend('<input type="hidden" name="lang_id" value="'+data.Language.Id+'">')
            $('input[type=hidden][name=csrf]').val(data.csrf)
            $('#laguangeModal').modal('show')
            $('.modal-header>h3').text(languagedata.Languages.editlanguage)
        }

    })
})

$('#laguangeModal').on('hide.bs.modal', function () {
    $('.input-group input[name=lang_name]').val("")
    $('.input-group input[name=lang_code]').val("")
    $('.input-field-group #ckb1[name=lang_status]').remove()
    $('.input-field-group #ckb2[name=lang_default]').remove()
    $('.input-field-group>.active-toggle:first>.toggle').prepend('<input class="tgl tgl-light" id="ckb1" name="lang_status" type="checkbox">')
    $('.input-field-group>.active-toggle:last>.toggle').prepend('<input class="tgl tgl-light" id="ckb2" name="lang_default" type="checkbox">')
    if ($('.full-upload >.upload-json:first>.uploaded-file').length == 1) {
        $('.full-upload >.upload-json:first>input[name=lang_json]').remove()
        $('.full-upload >.upload-json:first>.uploaded-file').remove()
        var html = `<div class="file-upload">
                   <img src="/public/img/folder.svg" alt="">
                   <p class="para">Drag your JSON file or choose from directory</p>
                   <div class="upload-button">
                   <button class="btn-reg btn-xs primary"> `+languagedata.browse+` </button>
                   <input type="file" name="lang_json" onchange="UploadJson(this)">
                   </div>
                   </div>`
        $(html).insertBefore('#lang_json-error')
    }
    if ($('.full-upload >.upload-json:last>.uploaded-file').length == 1) {
       $('.full-upload >.upload-json:last>.uploaded-file').remove()
       var html = `<div class="file-upload">
                  <img src="/public/img/upload-logo.svg" alt="">
                  <p class="para">Drag your JSON file or choose from directory</p>
                  <div class="upload-button">
                  <button class="btn-reg btn-xs primary"> `+languagedata.browse+` </button>
                  <input type="file" name="flag_imgpath" onchange="UploadFlagImage(this)">
                  </div>
                  </div>`
    $(html).insertBefore('#flag_imgpath-error')
    }
    $('#lang_name-error').hide()
    $('#lang_code-error').hide()
    $("#lang_json-error").hide()
    $("#flag_imgpath-error").hide()
    $('.modal-save>button').text(languagedata.save)
    $('.modal-header>h3').text(languagedata.Languages.addnewlanguage)
    if($('form[name=language-form]>input[type=hidden][name=lang_id]').length==1){
        $('form[name=language-form]>input[type=hidden][name=lang_id]').remove()
    }
    $('.input-group input').parents('.input-group').removeClass('input-group-error')
});

$('.delete-language').click(function(){

    var langId = $(this).attr("data-id")

    console.log("langid",langId);

    $('.deltitle').text(languagedata.Languages.deletelanguage+"?")

    $('.deldesc').text(languagedata.Languages.dellanguage)

    $('.delname').text($(this).parents('tr').find('td:first>.flexx>h4').text())

    $('button[id=delid]').parent().attr('href',"/settings/languages/deletelanguage/"+langId)

    $('#centerModal').modal('show')
})

$('#addnewimageModal').on('hide.bs.modal',function(){
    $('#laguangeModal').css('opacity',1)
})