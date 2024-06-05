var languagedata

var selectedcheckboxarr=[]
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    console.log("lang", languagepath);

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

})
$(document).keydown(function(event) {
    if (event.ctrlKey && event.key === '/') {
        $("#languagesearch").focus().select();
    }
});
$(document).ready(function () {

   
    $(document).on('click', '.modal-save>button', function () {

        jQuery.validator.addMethod("duplicatelangname", function (value) {

            var result;
            var id = $("#langid").val()
            
                $.ajax({
                url: "/settings/languages/checklanguagename",
                type: "POST",
                async: false,
                data: { "langname": value,"id":id ,csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    console.log("data",data);
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
                    duplicatelangname:true
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
                    required: languagedata.Languages.langnamevalid,
                    duplicatelangname :languagedata.Languages.duplicatenameerr
                },
                lang_code: {
                    required: languagedata.Languages.langcodevalid,
                },
                lang_json: {
                    required: languagedata.Languages.seljsonvalid,
                },
                flag_imgpath: {
                    required:languagedata.Languages.selflagimgvalid,
                },
            }
        })

        var isFormValid = $('form[name=language-form]').valid()

        if (!isFormValid) {

            $('.input-group').addClass('input-group-error')

            langcode = $("#triggerId").text()

            if (langcode != ""){

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
$(document).on('click','.search-grp>a',function(){

    var keyword = $(this).siblings("input").val().trim()

    if (keyword != "") {

        window.location.href = "/settings/languages/?keyword=" + keyword
    }

})

/*search redirect home page */

$(document).on('keyup', '#languagesearch', function () {

    if (event.key === 'Backspace') {

    if ($('.search').val() === "") {

         window.location.href = "/settings/languages/"
    }
}
})

$(document).on('click','input[name=lang_status][type=checkbox]',function(){

    if ($(this).is(':checked')) {

        $(this).val("1")
    } else {

        $(this).val("0")
    }
})

$(document).on('click','input[name=lang_default][type=checkbox]',function(){

    if ($(this).is(':checked')) {

        $(this).val("1")
    } else {

        $(this).val("0")
    }
})
$(document).on('change','.file-upload input[name=lang_json]',function(){
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

$(document).on('click','.file-upload input[name=flag_imgpath]',function(event){
    event.preventDefault()
    $('#laguangeModal').css('opacity', 0)
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
    <p class="para">`+ languagedata.Languages.uploadimg + `</p>
    <div class="upload-button">
        <button class="btn-reg btn-xs primary"> `+ languagedata.browse + ` </button>
        <input type="file" name="lang_json" onchange="UploadJson(this)">
    </div>
    </div>`
    $(html).insertBefore('#lang_json-error')
}

$(document).on('click', '#delete-flag', function () {
    console.log("yesss");
    $(this).parents('.upload-json').find('.uploaded-file').remove()
    var html = `<div class="file-upload">
        <img src="/public/img/upload-logo.svg" alt="">
        <p class="para">`+ languagedata.Languages.uploadimg + `</p>
        <div class="upload-button">
            <button class="btn-reg btn-xs primary"> `+ languagedata.browse + ` </button>
            <input  type="file" name="flag_imgpath" onclick="UploadFlagImage(event)">
        </div>
        </div>`
    $(html).insertBefore('#flag_imgpath-error')

})

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

function UploadFlagImage(event) {
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
    $('#laguangeModal').css('opacity', 0)
    $('#addnewimageModal').modal('show')
}

$(document).on('click', '.edit-language', function () {
    var langId = $(this).attr('data-id')
    $("#langid").val(langId)
    var parentTr = $(this).parents('tr')
    var keyword = $("#searchlanguagecode").val()
    $(".languagecode-list-row button").each (function (index, element) {
        if (keyword != ""){
            $(element).show()

            $("#searchlanguagecode").val("")
            $("#nodatafoundtexts").hide()
            }
    })
   
    $.ajax({
        url: '/settings/languages/editlanguage/' + langId,
        type: 'GET',
        dataType: 'JSON',
        success: function (data) {
            console.log("data", data)
            $('form[name=language-form]').attr("action", data.UploadUrl)
            $('input[name=lang_name]').val(data.Language.LanguageName)
            $('#triggerId').text(data.Language.LanguageCode)
            $("#lang_code").val(data.Language.LanguageCode)

            var value = data.Language.LanguageCode
           
            $(".languagecode-list-row button").each(function (index, element) {

                var title = $(element).text().toLowerCase()

                if (title.includes(value) && langId == data.Language.Id) {
                 
                        $(element).css('background-color','var(--clr-dropdown-hover)')
                    
                }else{
                    $(element).css('background-color','')
                }
            })

            if (data.Language.IsStatus == 1) {
                $('#ckb1[name=lang_status]').val("1").prop('checked', 'true')
            }else{
                $('#ckb1[name=lang_status]').val("0").prop('checked', 'false')
            }

            if (data.Language.DefaultLanguageId != 0 && data.Language.DefaultLanguageId == data.Language.Id) {
                $('#ckb2[name=lang_default]').val("1").attr('checked', true)
            }else{
                console.log("elseif");
                $('#ckb2[name=lang_default]').val("0").attr("checked", false)
            }

            $('.full-upload >.upload-json:first>.file-upload').remove()

            var uploaded_html = `<div class="uploaded-file">
                <input type="hidden"  name="lang_json" value="`+ data.Language.JsonPath + `">
                <img class="lang_json_para" src="/public/img/folder.svg" alt="">
                <p class="para lang_json_para" id="lang_json_para">`+ data.Language.JsonPath.split('/').pop() + `</p>
                </div>`

            $(uploaded_html).insertBefore("#lang_json-error")

            $('#lang_json-error').hide()

            $('.full-upload >.upload-json:last>.file-upload').remove()

            if (data.Language.ImagePath != ""){
                var flag_uploadHtml = `<div class="uploaded-file">
                <input type="hidden" name="flag_imgpath" value="`+ data.Language.ImagePath + `">
                <img src="`+ data.Language.ImagePath + `" alt="" class="uploaded-img-flag">
                <button class="delete-flag" type="button"><img id="delete-flag" src="/public/img/delete-white-icon.svg" alt=""></button>
                <div class="hover-delete-img"></div>
                </div>`
    
                $(flag_uploadHtml).insertBefore('#flag_imgpath-error')
            }else{
                var flag_uploadHtml = `<div class="file-upload">
                <img src="/public/img/upload-logo.svg" alt="">
                <p class="para">`+ languagedata.Languages.uploadimg + `</p>
                <div class="upload-button">
                    <button class="btn-reg btn-xs primary"> `+ languagedata.browse + ` </button>
                    <input  type="file" name="flag_imgpath" onclick="UploadFlagImage(event)">
                </div>
                </div>`
    
                $(flag_uploadHtml).insertBefore('#flag_imgpath-error')
            }
             
            $('#flag_imgpath-error').hide()
            $('.modal-save>button').text(languagedata.update)
            $('form[name=language-form]').prepend('<input type="hidden" name="lang_id" value="' + data.Language.Id + '">')
            $('input[type=hidden][name=csrf]').val(data.csrf)
            $('#lang_json').prop('disabled', true);
            // $('#lang_code').prop('disabled', true);
            // $('#lang_code').css("color", '#d2d2d2')
            $('#lang_json_para').css("color", '#d2d2d2')
            $('#laguangeModal').modal('show')
            $('.modal-header>h3').text(languagedata.Languages.editlanguage)
        }

    })
})

$('#laguangeModal').on('hide.bs.modal', function () {
    $('#lang_json').prop('disabled', false);
    // $('#lang_code').prop('disabled', false);
    // $('#lang_code').css("color", '')
    $('#lang_json_para').css("color", '')
    $('.input-group input[name=lang_name]').val("")
    // $('.input-group input[name=lang_code]').val("")
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
                   <button class="btn-reg btn-xs primary"> `+ languagedata.browse + ` </button>
                   <input type="file" name="lang_json" onchange="UploadJson(this)">
                   </div>
                   </div>`
        $(html).insertBefore('#lang_json-error')
    }
    if ($('.full-upload >.upload-json:last>.uploaded-file').length == 1) {
        $('.full-upload >.upload-json:last>.uploaded-file').remove()
        var html = `<div class="file-upload">
                  <img src="/public/img/upload-logo.svg" alt="">
                  <p class="para">`+ languagedata.Languages.uploadimg + `</p>
                  <div class="upload-button">
                  <button class="btn-reg btn-xs primary"> `+ languagedata.browse + ` </button>
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
    if ($('form[name=language-form]>input[type=hidden][name=lang_id]').length == 1) {
        $('form[name=language-form]>input[type=hidden][name=lang_id]').remove()
    }
    $('.input-group input').parents('.input-group').removeClass('input-group-error')
});

$(document).on('click', '.delete-language', function () {

    var langId = $(this).attr("data-id")

    console.log("langid", langId);

    $('.deltitle').text(languagedata.Languages.deletelanguage + "?")

    $('.deldesc').text(languagedata.Languages.dellanguage)

    $('.delname').text($(this).parents('tr').find('td:first>.flexx>h4').text())

    $('button[id=delid]').parent().attr('href', "/settings/languages/deletelanguage/" + langId)

    $('#centerModal').modal('show')
})

$('#addnewimageModal').on('hide.bs.modal', function () {

    $('#laguangeModal').css('opacity', 1)

})

function LanguageStatus(This, event) {

    console.log("click", $(This).parents('tr').find('td:first>.tbl-img-content>.default').length == 0);

    // if ($(This).parents('tr').find('td:first>.tbl-img-content>.default').length == 0) {
    var lang_id = $(This).attr("data-id")

    if ($(This).is(':checked')) {

        $(This).val("1")
    } else {

        $(This).val("0")
    }
    var isactive = $(This).val();
    $.ajax({
        url: '/settings/languages/languageisactive',
        type: 'POST',
        async: false,
        data: {
            id: lang_id,
            isactive: isactive,
            csrf: $("input[name='csrf']").val()
        },
        dataType: 'json',
        cache: false,
        success: function (result) {
            console.log("language", result.language);
            if (result.language.IsStatus == 0) {

                $(".lang[data-id=" + result.language.Id + "]").remove()

                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Toast.languageupdated + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds
            }
            if (result.language.IsStatus == 1) {
                if (result.language.Id == $("#current-lang").attr("data-id")) {
                    notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>'+ languagedata.Toast.languageupdated +'</span></div>';
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();
                        });
                    }, 5000); // 5000 milliseconds = 5 seconds
                } else {
                    var html = `<li class="lang" data-code="${result.language.LanguageCode}" data-id="${result.language.Id}"><button class="dropdown-item" type="button"> <img src="${result.language.ImagePath}" alt="">${result.language.LanguageName}</button> </li>`
                    $("#lang-menu").append(html)
                    notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Toast.languageupdated + '</span></div>';
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();
                        });
                    }, 5000); // 5000 milliseconds = 5 seconds
                }
            }

        }
    });
    // } else {
    //     event.preventDefault()
    // }
}
// language code dropdown

$(".dropdown-values").on("click", function () {
    var text = $(this).text()
    $(this).parents().siblings("a").text(text)
    $("#lang_code").val(text)
    $("#lang_code-error").hide()
    $("#languagecode").removeClass("input-group-error")
   
})

// language code dropdown filter input box search
$("#searchlanguagecode").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".languagecode-list-row button").each(function (index, element) {
        var title = $(element).text().toLowerCase()

        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafoundtexts").hide()

        } else {
            $(element).hide()
            if($('.languagecode-list-row button:visible').length==0){
                $("#nodatafoundtexts").show()
            }
           
        }
    })

})

$(document).on('click','.selectcheckbox',function(){

    languageid =$(this).attr('data-id')

    var status = $(this).parents('td').siblings('td').find('.tgl-light').prop('checked');

    console.log(status,"status")

    var sstatus

 
    if (status ==true){
       
        sstatus= '1'

        }else{

        sstatus= '0'
        
        }

   if ($(this).prop('checked')){

       selectedcheckboxarr.push({"languageid":languageid,"status":sstatus})
   
   }else{

       const index = selectedcheckboxarr.findIndex(item => item.languageid === languageid);
   
       if (index !== -1) {

           console.log(index,"sssss")
           selectedcheckboxarr.splice(index, 1);
       }
      
       $('#Check').prop('checked',false)

   }

  
   if (selectedcheckboxarr.length !=0){

       $('.selected-numbers').show()

       var allSame = selectedcheckboxarr.every(function(item) {
           return item.status === selectedcheckboxarr[0].status;
       });
       
       var setstatus
       var img;

          if (selectedcheckboxarr[0].status === '1') {
 
            setstatus = "Deactive";

             img = "/public/img/In-Active (1).svg";

       } else if (selectedcheckboxarr[0].status === '0') {

              setstatus = "Active";

              img = "/public/img/Active (1).svg";

         } 

          var htmlContent = '';

            if (allSame) {

             htmlContent = '<img src="' + img + '">' + setstatus;

             } else {

               htmlContent = '';

             }

           $('#unbulishslt').html(htmlContent);
      
       $('.checkboxlength').text(selectedcheckboxarr.length+" " +'items selected')

       if(!allSame){

           $('#seleccheckboxdelete').removeClass('border-end')

           $('.unbulishslt').html("")
       }else{

           $('#seleccheckboxdelete').addClass('border-end')
       }

      
   }else{

       $('.selected-numbers').hide()
   }

       var allChecked = true;

        $('.selectcheckbox').each(function() {

            if (!$(this).prop('checked')) {

            allChecked = false;

           return false; 
        }
   });

         $('#Check').prop('checked', allChecked);

    console.log(selectedcheckboxarr,"checkkkk")
})

//ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click','#Check',function(){

    selectedcheckboxarr=[]

    var isChecked = $(this).prop('checked');

    if (isChecked){

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function(){
    
           languageid= $(this).attr('data-id')

           var status = $(this).parents('td').siblings('td').find('.tgl-light').prop('checked');

           var sstatus

           if (status ==true){
       
            sstatus= '1'
    
            }else{
    
            sstatus= '0'

            }
 
    
           selectedcheckboxarr.push({"languageid":languageid,"status":sstatus})
        })

        $('.selected-numbers').show()

        var allSame = selectedcheckboxarr.every(function(item) {

            return item.status === selectedcheckboxarr[0].status;
        });
        
        var setstatus

        var img
    

        if (selectedcheckboxarr.length !=0){

        if (selectedcheckboxarr[0].status === '1') {
 
            setstatus = "Deactive";

             img = "/public/img/In-Active (1).svg";

       } else if (selectedcheckboxarr[0].status === '0') {

              setstatus = "Active";

              img = "/public/img/Active (1).svg";

         } 
        }

        var htmlContent = '';

        if (allSame) {

         htmlContent = '<img src="' + img + '">' + setstatus;

         $('#seleccheckboxdelete').addClass('border-end')

         } else {

           htmlContent = '';

           $('#seleccheckboxdelete').removeClass('border-end')

         }

       $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length+" " +'items selected')
    
    }else{


        selectedcheckboxarr=[]

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selected-numbers').hide()
    }

    if (selectedcheckboxarr.length ==0){

        $('.selected-numbers').hide()
    }
})

$(document).on('click','#seleccheckboxdelete',function(){

    if (selectedcheckboxarr.length>1){
         
    $('.deltitle').text("Delete Languages?")

    $('#content').text('Are you sure want to delete selected Languages?')

    }else {

         $('.deltitle').text("Delete Language?")

         $('#content').text('Are you sure want to delete selected Language?')
    }


    $('#delete').addClass('checkboxdelete')
})

$(document).on('click','#unbulishslt',function(){

    if (selectedcheckboxarr.length>1){

         $('.deltitle').text( $(this).text()+" "+"Languages?")

         $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Languages?")
    }else{

         $('.deltitle').text( $(this).text()+" "+"Language?")

         $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Language?")
    }

    $('#delete').addClass('selectedunpublish')

})

//MULTI SELECT DELETE FUNCTION//
$(document).on('click','.checkboxdelete',function(){

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
            "page":pageno

            
        },
        success: function (data) {

            console.log(data,"result")

            if (data.value==true){

                setCookie("get-toast", "Language Deleted Successfully")

                window.location.href=data.url
            }else{

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})
//Deselectall function//

$(document).on('click','#deselectid',function(){

    $('.selectcheckbox').prop('checked',false)

    $('#Check').prop('checked',false)

    selectedcheckboxarr=[]

    $('.selected-numbers').hide()
    
})

//multi select active and deactive function//

$(document).on('click','.selectedunpublish',function(){

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
            "languageids":JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page":pageno

            
        },
        success: function (data) {

            console.log(data,"result")

            if (data.value==true){

                setCookie("get-toast", "languageupdated")

                window.location.href=data.url
            }else{

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })


})