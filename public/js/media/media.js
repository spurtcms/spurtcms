var languagedata

var DeleteListArray = [];

var MediaBreadcrumbRoot = [];

let cropHeight, cropWidth, viewportType, dimension, html_imgurl;

var img, canvas, ctx;

let logoexpand_cropper,cropper;

/** */
$(document).ready(async function () {

    $('.Content').addClass('checked');

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/" + languagecode + ".json", function (data) {

        languagedata = data
    })

})

$(document).on('click', '.upload-folder-img>img', function () {

    if (window.location.href.indexOf("media") == -1 && !$(this).parent().hasClass('folder')) {

        decoded_url = decodeURIComponent($(this).attr('src'))

        html_imgurl = "/"

        var cuts =  decoded_url.split("/")

        for(let x in cuts){

            if(cuts[x]!=""){

                if(x==cuts.length-1){

                    html_imgurl += cuts[x]

                }else{

                    html_imgurl += cuts[x]+"/"
                }
            }
        }

        img = new Image()

        img.src = html_imgurl

        if (window.location.href.indexOf("spaces") != -1) {

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 500,

                    height: 250,

                    type: viewportType

                },

                showZoomer: false,
                // fit: true,

            });

            dimension = '500*250'

        } else if (window.location.href.indexOf("entry") != -1) {

            $('#changepicModal .modal-content').addClass('banner_crop')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 740,

                    height: 370,
                },

                showZoomer: false,

                // fit: true,
            });

            dimension = '740*370'

        } else if (window.location.href.indexOf("languages") != -1) {

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 360,

                    height: 240,

                },

                showZoomer: false,

                // fit: true,
            });

            dimension = '360*240'

        } else if (window.location.href.indexOf("category") != -1) {

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 500,

                    height: 250,
                },

                showZoomer: false,
                
                // fit: true,
            });
            dimension = '500*250'

        }else if (window.location.href.indexOf("personalize") != -1 && $('#logo-input').val()=="1"){

            $('#crop-container').removeClass('croppie-container').empty().hide()

            $('#crop-container1').removeClass('croppie-container').empty().show()

            logoexpand_cropper = $('#crop-container1').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 500,

                    height: 250,
                },

                showZoomer: false,
                // fit: true,
            });

            dimension = '500*250'

        }else if (window.location.href.indexOf("personalize") != -1 && $('#logo-input').val()!="1"){

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()
            
            // $('.cr-slider-wrap').remove()

            cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 300,

                    height: 300,

                    type: viewportType
                },
                
                showZoomer: false,
                // fit: true,
            });

            dimension = '300*300'
        }

        if($('.cr-slider-wrap').length>1){

            $('.cr-slider-wrap')[1].remove()
        }

        $('.cr-slider-wrap').appendTo($('#zoom-div'));

        $('canvas[class=cr-image]').css('opacity', '0')

        $('.cr-slider').css('display', 'block')

        $('#rotateSlider').val("0")

        $('#changepicModal').modal('show')
    }
})

$('#changepicModal').on('shown.bs.modal', function () {

    canvas = document.querySelector('canvas[class=cr-image]');

    ctx = canvas.getContext('2d');

    if (window.location.href.indexOf("personalize") != -1 && $('#logo-input').val()=="1") {

        logoexpand_cropper.croppie('bind', {

            url: html_imgurl,

            orientation: 0,

            zoom: 0

        }).then(function () {

            // cropper.croppie('setZoom',0)

            console.log("cropper initialized!");
        })

    }else{

        cropper.croppie('bind', {

            url: html_imgurl,

            orientation: 0,

            zoom: 0

        }).then(function () {

            // cropper.croppie('setZoom',0)

            console.log("cropper initialized!");
        })
    }

});

$('#rotateSlider').on('input', function () {

    var rotationValue = parseInt($(this).val());

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    ctx.save()

    ctx.translate(canvas.width / 2, canvas.height / 2);

    ctx.rotate((rotationValue * Math.PI) / 180);

    ctx.drawImage(img, -canvas.width / 2, -canvas.height / 2);

    ctx.restore()

});


$('#crop-button').click(function () {

    if (window.location.href.indexOf("personalize") != -1 && $('#logo-input').val()=="1"){

        logoexpand_cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg' }).then(function (dataUrl) {

            imgData = GetImageProcessData(html_imgurl)

            console.log("chk",imgData);

            var final_imgName = ''

            const currentTimestamp = new Date().toISOString();

            if (imgData[0].indexOf(`(${dimension})`) == -1) {

                final_imgName = imgData[0] + "(" + dimension + ")(" + currentTimestamp + ")." + imgData[1]

            } else {

                final_imgName = imgData[0].split(dimension)[0] + dimension + ")(" + currentTimestamp + ")." + imgData[1]

            }

            console.log("final_name",final_imgName);

            if (imgData[0] != "" && imgData[1] != "" && dimension != "") {

                var formData = new FormData()

                formData.append('csrf', $("input[name='csrf']").val())

                formData.append('file', dataUrl)

                formData.append('path', html_imgurl.split('/')[1] + "/" + html_imgurl.split('/')[2])

                formData.append('name', final_imgName)

                $.ajax({

                    type: "post",

                    url: "/media/uploadimage",

                    data: formData,

                    datatype: "json",

                    cache: false,

                    processData: false,

                    contentType: false,

                    success: function (result) {

                        html_imgurl = ""

                        parsed_obj = JSON.parse(result)

                        if (window.location.href.indexOf("spaces") != -1) {

                            BindCropImageInSpace(parsed_obj.Path)

                        } else if (window.location.href.indexOf("entry") != -1) {

                            BindCropImageInEntry(parsed_obj.Path)

                        } else if (window.location.href.indexOf("personalize") != -1) {

                            BindCropImageInPersonalize(parsed_obj.Path)

                        } else if (window.location.href.indexOf("addcategory") != -1) {

                            BindCropImageInCategory(parsed_obj.Path)

                        } else if (window.location.href.indexOf("") != -1) {

                            BindCropImageInLanguage(parsed_obj.Path)

                        }

                        $("#rightModal").modal('show')

                        $('#changepicModal').modal('hide')

                        $('#crop-container').removeClass('croppie-container').empty().show()

                        $('#crop-container1').removeClass('croppie-container').empty().hide()

                        $('#addnewimageModal').modal('hide')

                    }

                })

            } else {

                console.log("imgdata", imgData[0], imgData[1], dimension);
            }

        });


    }else{

        cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg' }).then(function (dataUrl) {

            imgData = GetImageProcessData(html_imgurl)

            console.log("chk",imgData);

            var final_imgName = ''

            const currentTimestamp = new Date().toISOString();

            if (imgData[0].indexOf(`(${dimension})`) == -1) {

                final_imgName = imgData[0] + "(" + dimension + ")(" + currentTimestamp + ")." + imgData[1]

            } else {

                final_imgName = imgData[0].split(dimension)[0] + dimension + ")(" + currentTimestamp + ")." + imgData[1]

            }

            console.log("final_name",final_imgName);

            if (imgData[0] != "" && imgData[1] != "" && dimension != "") {

                var formData = new FormData()

                formData.append('csrf', $("input[name='csrf']").val())

                formData.append('file', dataUrl)

                formData.append('path', html_imgurl.split('/')[1] + "/" + html_imgurl.split('/')[2])

                formData.append('name', final_imgName)

                $.ajax({

                    type: "post",

                    url: "/media/uploadimage",

                    data: formData,

                    datatype: "json",

                    cache: false,

                    processData: false,

                    contentType: false,

                    success: function (result) {

                        html_imgurl = ""

                        parsed_obj = JSON.parse(result)

                        if (window.location.href.indexOf("spaces") != -1) {

                            BindCropImageInSpace(parsed_obj.Path)

                        } else if (window.location.href.indexOf("entry") != -1) {

                            BindCropImageInEntry(parsed_obj.Path)

                        } else if (window.location.href.indexOf("personalize") != -1) {

                            BindCropImageInPersonalize(parsed_obj.Path)

                        } else if (window.location.href.indexOf("addcategory") != -1) {

                            BindCropImageInCategory(parsed_obj.Path)

                        } else if (window.location.href.indexOf("") != -1) {

                            BindCropImageInLanguage(parsed_obj.Path)

                        }

                        $("#rightModal").modal('show')

                        $('#changepicModal').modal('hide')

                        $('#crop-container').removeClass('croppie-container').empty().show()

                        $('#crop-container1').removeClass('croppie-container').empty().hide()

                        $('#addnewimageModal').modal('hide')
                    }

                })

            } else {

                console.log("imgdata", imgData[0], imgData[1], dimension);

            }

        });
    }

});

$('#close,#crop-cancel').click(function () {

    html_imgurl = ""

    $('canvas[class=cr-image]').css('opacity', '0')

    $("#changepicModal").modal('hide');

    $('#crop-container').removeClass('croppie-container').empty().hide()

    $('#crop-container1').removeClass('croppie-container').empty().hide()

});

$('#changepicModal').on('hide.bs.modal', function (event) {

    html_imgurl = ""

    $('canvas[class=cr-image]').css('opacity', '0')

    $('#crop-container').removeClass('croppie-container').empty().hide()

    $('#crop-container1').removeClass('croppie-container').empty().hide()

})

$("#addnewimageModal").on('show.bs.modal', function () {

    $('#Refreshdiv').click();

})

function GetImageProcessData(html_imgurl) {

    console.log("raw",html_imgurl);

    var ext = html_imgurl.split('/')[3].split('.').pop()

    var imgName = ""

    if (html_imgurl.split('/')[3].split('.').length > 2) {

        for (let part in html_imgurl.split('/')[3].split('.')) {

            if(part<html_imgurl.split('/')[3].split('.').length-1){

                 imgName = imgName + html_imgurl.split('/')[3].split('.')[part]
            }

        }

    } else {

        imgName = html_imgurl.split('/')[3].split('.')[0]

    }

    return [imgName, ext]
}

function BindCropImageInSpace(src) {

    $("#spimage").val(src)

    var data = $("#spimagehide").attr("src", src);

    if (data != "") {

        $(".heading-three").hide();

        $("#browse").hide();

        $("#spimagehide").attr("src", src).show()
        
        $("#spacedel-img").show();

        $("#addnewimageModal").modal('hide')

    }

}

function BindCropImageInEntry(src) {

    $("#spimage").val(src)

    var data = $("#spimagehide").attr("src", src);

    if (data != "") {

        $('#cenimg').hide()

        $("#spimagehide").css("height", "23.120rem")

        $("#spimagehide").css("width", "46.25rem")

        $("#spimagehide").attr("src", src).show()

        $("#spacedel-img").show();

    }

    $("#addnewimageModal").hide()

}

function BindCropImageInPersonalize(src) {

    var logo_select = $("#logo-input").val();

    if (logo_select == '1') {

        $("#upload-expandlogo").remove()

        var html = `<div class="upload-logo" id="upload-expandlogo"><img src="` + src + `" alt=""><input type="hidden" id="expandlogo-input" name="expandlogo_imgpath" value="` + src + `"><button class="delete-flag" id="delete-expandlogo"><img src="/public/img/delete-white-icon.svg" alt=""></button><div class="hover-delete-img"></div>`

        $(html).insertAfter('#expandlogo-insert')

        $("#addnewimageModal").hide()

        $('.modal-backdrop').remove()

    }else{

        $("#upload-logo").remove()

        var html = `<div class="upload-logo" id="upload-logo"><img src="` + src + `" alt=""><input type="hidden" id="logo-input" name="logo_imgpath" value="` + src + `"><button class="delete-flag" id="delete-logo"><img src="/public/img/delete-white-icon.svg" alt=""></button><div class="hover-delete-img"></div>`

        $(html).insertAfter('#logo-insert')

        $("#addnewimageModal").hide()

        $('.modal-backdrop').remove()

    }
}

function BindCropImageInCategory(src) {

    $("#categoryimage").val(src)

    var data = $("#ctimagehide").attr("src", src);

    if (data != "") {

        $(".heading-three").hide();

        $("#browse").hide();

        $("#ctimagehide").attr("src", src).show()

        $("#catdel-img").show()

    }

    $("#addnewimageModal").hide()

    $("#categoryModal").modal('show')

}

function BindCropImageInLanguage(src) {

    $('.file-upload input[name=flag_imgpath]').parents('.upload-json').find('.file-upload').remove()

    var flag_uploadHtml = `<div class="uploaded-file">

            <input type="hidden" name="flag_imgpath" value="`+ src + `">

            <img src="`+ src + `" alt="" class="uploaded-img-flag">

            <button class="delete-flag" onclick="RemoveFlagImage(this)" ><img src="/public/img/delete-white-icon.svg" alt=""></button>

            <div class="hover-delete-img"></div>

            </div>`

    $(flag_uploadHtml).insertBefore('#flag_imgpath-error')

    $('#flag_imgpath-error').hide()

    $("#addnewimageModal").modal('hide')

    $('#laguangeModal').modal('show')

}

/*Add Folder Button hide and show*/
$(document).on("click", "#addfolder", function () {

    $("#addfoldervalues").toggle();

    $("#foldername").val('');

})


// add button 
$("#foldername").keyup(function () {

    $("#namecheck").removeClass('input-group-error')

    $("#foldername-error").hide()

})

/*Add Folder function */
$(document).on('click', '#newfileadd', function () {

    var name = $("#foldername").val().trim();

    var folderpath = ""

    for (let x of MediaBreadcrumbRoot) {

        folderpath += x + "/"

    }

    if (name == "") {

        $('#foldername-error').show();

        $("#namecheck").addClass('input-group-error')


    } else {

        $("#namecheck").removeClass('input-group-error')

        $("#foldername-error").hide()

        $("#errorhtml").html('')

        $.ajax({

            type: "post",

            url: "/media/createfolder",

            datatype: "json",

            data: {

                foldername: name,

                path: folderpath,

                csrf: $("input[name='csrf']").val()

            },

            success: function (response) {

                if (response) {

                    $('#foldername-error').hide();

                    $("#addfoldervalues").hide();

                    $('#Refreshdiv').trigger('click');

                    var message = `Folder Created Successfully`
    
                    if(window.location.href.indexOf("media")!=-1){
        
                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';
        
                        $(notify_content).insertBefore(".header-rht");
                
                    }else{
        
                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';
                
                        $(notify_content).insertBefore('.modal-content')
                    }

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {
  
                          $(this).remove();
  
                        });
  
                     }, 5000);
  

                } else {
                   
                   var message = `Unable to create Folder,Please try again later`

                   if(window.location.href.indexOf("media")!=-1){

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".header-rht");

                   }else{

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".modal-content");

                   }

                   setTimeout(function () {

                      $('.toast-msg').fadeOut('slow', function () {

                        $(this).remove();

                      });

                   }, 5000);

                }
            }
        })

    }

})


/*ckbox*/
$(document).on('click', '.ckbox', function () {

    var name = $(this).parent().siblings('.file-detail').find('h3').text();

    var isAllChecked

    $('.ckbox').each(function () {

        if ($(this).is(':checked')) {

            isAllChecked = true

        } else {

            isAllChecked = false

            return false
        }
    })

    if (isAllChecked) {

        $("#selectall").hide()

        $("#unselectall").show()

    } else {

        $("#selectall").show()

        $("#unselectall").hide()

    }

    if ($(this).is(':checked')) {

        DeleteListArray.push(name)

    } else {

        DeleteListArray = jQuery.grep(DeleteListArray, function (value) {

            return value != name;

        });
    }

    console.log("checking media delete array",DeleteListArray);

    TrashButton()

})

function TrashButton() {

    if (DeleteListArray.length > 0) {

        $('.disabled').addClass('dis').removeClass('disabled');

    } else {

        $('.dis').removeClass('dis').addClass('disabled');
    }
}

$(document).on('click', '#upload', function (event) {

    event.preventDefault()

    $('#imgupload').click()

})

/*Trash */
$(document).on('click', '.dis', function () {

    var isLeastOneChecked

    $('.ckbox').each(function () {

        if ($(this).is(':checked')) {

            isLeastOneChecked = true

            return false

        } else {

            isLeastOneChecked = false
        }
    })

    if (isLeastOneChecked) {

        $('#centerModal').modal('show');

        $(".deltitle").text(languagedata.Mediaa.deltitle)

        $('#content').text(languagedata.Mediaa.delmediacontent);

        // $(".deltitle").text('Delete Media File ?')

        // $('#centerModal .modal-body>#content').text('Are you sure, you want to delete this media files ?')

        $('#delid').show();

        $('#delete').attr('href', 'javascript:void(0)');

        $('#delcancel').text(languagedata.no);

    }
})

// no btn functionality
$(document).on('click', '#delcancel', function () {

    $('#centerModal').modal('hide');

})

/*Trash */
$(document).on('click', "#delete", function () {

    var folderpath = ""

    for (let x of MediaBreadcrumbRoot) {

        folderpath +=  x + "/"
    }

    console.log("delelist", DeleteListArray);

    $.ajax({

        type: "post",

        url: "/media/deletefolfile",

        data: {

            csrf: $("input[name='csrf']").val(),

            id: JSON.stringify(DeleteListArray),

            path: folderpath

        },

        datatype: "json",

        success: function (response) {

            if(response){

                $('.dis').removeClass('dis').addClass('disabled');
    
                $('#Refreshdiv').trigger('click');
    
                $("#delcancel").trigger('click')
    
                DeleteListArray = []
    
                var message = `Mediafile Deleted Successfully`
    
                if(window.location.href.indexOf("media")!=-1){
    
                    notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';
    
                    $(notify_content).insertBefore(".header-rht");
            
                }else{
    
                    notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';
            
                    $(notify_content).insertBefore('.modal-content')
                }
            
                setTimeout(function () {
    
                    $('.toast-msg').fadeOut('slow', function () {
    
                        $(this).remove();
                    });
    
                }, 5000); // 5000 milliseconds = 5 seconds

            }else{

                   var message = `Unable to delete,Please try again later`

                   if(window.location.href.indexOf("media")!=-1){

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".header-rht");

                   }else{

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".modal-content");

                   }

                   setTimeout(function () {

                      $('.toast-msg').fadeOut('slow', function () {

                        $(this).remove();

                      });

                   }, 5000);

            }

        }
    })
})

/*Upload file*/
$(document).on('change', '#imgupload', function () {

    var folderpath = ""

    for (let x of MediaBreadcrumbRoot) {

        folderpath +=  x + "/"

    }

    var filename = $(this).val();

    var ext = filename.split(".").pop().toLowerCase();

    if (($.inArray(ext, ["jpg", "png", "jpeg"]) != -1)) {

        var formdata = new FormData();

        formdata.append('path', folderpath)

        formdata.append("image", this.files[0])

        formdata.append('csrf', $("input[name='csrf']").val())

        $.ajax({

            url: "/media/upload",

            type: "POST",

            data: formdata,

            processData: false,

            contentType: false,

            success: function (response) {

                if(response){

                   $('#Refreshdiv').click()

                   var message = `Image Uploaded Successfully`

                   if(window.location.href.indexOf("media")!=-1){

                    notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".header-rht");
        
                   }else{

                    notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';
        
                     $(notify_content).insertBefore('.modal-content')
                   }
        
                   setTimeout(function () {

                    $('.toast-msg').fadeOut('slow', function () {

                        $(this).remove();
                    });

                  }, 5000); // 5000 milliseconds = 5 seconds

                }else{

                    var message = `Unable to upload image,Please try again later`

                   if(window.location.href.indexOf("media")!=-1){

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".header-rht");

                   }else{

                     notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                     $(notify_content).insertBefore(".modal-content");

                   }

                   setTimeout(function () {

                      $('.toast-msg').fadeOut('slow', function () {

                        $(this).remove();

                      });

                   }, 5000);

                }

            },

            error: function () {

                $("#result").text(languagedata.Mediaa.uploaderrmsg);
            }
        });
    }
    else {

        var message = "Please choose images with jpg || jpeg || png formats only"

         if(window.location.href.indexOf("media")!=-1){

            var notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore(".header-rht");

         }else{

            var notify_content = '<div style="top:2px" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore('.modal-content')
         }

        setTimeout(function () {

            $('.toast-msg').fadeOut('slow', function () {

              $(this).remove();

            });

        }, 5000);
    }
})

$(document).on('click', '#imgupload', function () {

    if($(this).val()!=""){

        $(this).val("")
    } 
})

/*Refresh */
$(document).on('click', '#Refreshdiv', function () {

    var folderpath = ""

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata.Mediaa.medialibrary + `</a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if(MediaBreadcrumbRoot.length==1){

            folderpath += MediaBreadcrumbRoot[x]
            
        }else{

            if(x==0){

                folderpath +=  MediaBreadcrumbRoot[x] + "/"
    
            }else if(x==MediaBreadcrumbRoot.length-1){
    
                folderpath +=  MediaBreadcrumbRoot[x] 
    
            }else{
    
                folderpath +=  MediaBreadcrumbRoot[x] + "/" 
            }
        }

        bcrumb += `<li>/</li><li><a href="javascript:void(0)" data-id="` +  MediaBreadcrumbRoot[x] + `" class="bclick">` +  MediaBreadcrumbRoot[x] + `</a></li>`

    }

    $('.mediabreadcrumb').children('ul').html(bcrumb);

    $.ajax({

        type: "post",

        url: "/media/singlefolder",

        data: {

            path: folderpath,

            csrf: $("input[name='csrf']").val()

        },

        datatype: "json",

        cache: false,

        success: function (result) {

            var res = JSON.parse(result)

            $('.media-library-list').html('');

            $(".media-error-design").html('')

            $("#mediasearch").val("")

            str = "";

            if (res.folder != null || res.file != null) {

                $('.media-select-del').children().not('#unselectall').show()

                if (res.folder != null) {

                    for (let x of res.folder) {

                        str += `   <div class="upload-folders folderdiv ">

               <p class="forsearch" style="display: none;">`+ x.Name + `</p>

               <div class="chk-group chk-group-box">

                   <input type="checkbox" class="ckbox" id="`+ x.Name + `" data-id="` + x.Name + `"  for="` + x.Name + `">

                   <label for="`+ x.Name + `"></label>

               </div>

               <div class="upload-folder-img folder" style="cursor: pointer;">

                   <img src="/public/img/folder-media.svg" alt="">

               </div>

               <div class="file-detail media-library-content">

                   <h3>`+ x.Name + `</h3>

                   <p></p>

               </div>

           </div>`

                    }
                }

                if (res.file != null) {

                    for (let j of res.file) {

                        str += `

                    <div class="upload-folders filediv ">

                    <p class="forsearch" style="display: none;">`+ j.Name + `</p>


                    <div class="chk-group chk-group-box">

                        <input type="checkbox" class="ckbox" id="`+ j.Name + `" data-id="` + j.Name + `" for="` + j.Name + `">
                        
                        <label for="`+ j.Name + `"></label>

                    </div>

                    <div class="upload-folder-img" style="cursor: pointer;">

                    <img src="/`+ encodeURIComponent(j.Path + '/' + j.Name) + `" alt="" id="cimg" >

                    </div>

                    <div class="file-detail media-library-content">

                        <h3>`+ j.Name + `</h3>

                        <p></p>

                    </div>

                </div>
                    `
                    }
                }

                $("#drivelist").html(str)

            } else {

                $('.media-select-del').children().hide()

                $('#errorhtml').html(`<div class="noData-foundWrapper">

                <div class="empty-folder">

                    <img src="/public/img/folder-sh.svg" alt="">

                    <img src="/public/img/shadow.svg" alt="">

                </div>

                <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

            </div>`)

            }

        }
    })

})


/*Folder Inside */
$(document).on('click', '.folder', function () {

    DeleteListArray = [];

    TrashButton()

    var foldername = $(this).siblings(".file-detail").children('h3').text();

    MediaBreadcrumbRoot.push(foldername)

    var folderpath = ""

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata.Mediaa.medialibrary + `</a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if(MediaBreadcrumbRoot.length==1){

            folderpath += MediaBreadcrumbRoot[x]
            
        }else{

            if(x==0){

                folderpath +=  MediaBreadcrumbRoot[x] + "/"
    
            }else if(x==MediaBreadcrumbRoot.length-1){
    
                folderpath +=  MediaBreadcrumbRoot[x] 
    
            }else{
    
                folderpath +=  MediaBreadcrumbRoot[x] + "/" 
            }
        }

        bcrumb += `<li>/</li><li><a href="javascript:void(0)" data-id="` +  MediaBreadcrumbRoot[x] + `" class="bclick">` +  MediaBreadcrumbRoot[x] + `</a></li>`

    }

    $('.mediabreadcrumb').children('ul').html(bcrumb);

    $.ajax({
 
        type: "post",

        url: "/media/singlefolder",

        data: {

            path: folderpath,

            csrf: $("input[name='csrf']").val()

        },

        datatype: "json",

        cache: false,

        success: function (result) {

            var res = JSON.parse(result)

            $('#drivelist').html('');

            $("#errorhtml").html('')

            $("#mediasearch").val("")

            str = "";

            if (res.folder != null || res.file != null) {

                $('.media-select-del').children().not('#unselectall').show()

                if (res.folder != null) {

                    for (let x of res.folder) {

                        str += `   <div class="upload-folders folderdiv ">

                        <p class="forsearch" style="display: none;">`+ x.Name + `</p>

                        <div class="chk-group chk-group-box">

                            <input type="checkbox" class="ckbox" id="`+ x.Name + `" data-id="` + x.Name + `"  for="` + x.Name + `">

                            <label for="`+ x.Name + `"></label>

                        </div>

                        <div class="upload-folder-img folder" style="cursor: pointer;">

                            <img src="/public/img/folder-media.svg" alt="">

                        </div>

                        <div class="file-detail media-library-content">

                            <h3>`+ x.Name + `</h3>

                            <p></p>

                        </div>

                    </div>`

                    }

                }

                if (res.file != null) {

                    for (let j of res.file) {
                        
                        str += `

                        <div class="upload-folders filediv">

                        <p class="forsearch" style="display: none;">`+ j.Name + `</p>

                        <div class="chk-group chk-group-box">

                            <input type="checkbox" class="ckbox" id="`+ j.Name + `" data-id="` + j.Name + `" for="` + j.Name + `">

                            <label for="`+ j.Name + `"></label>

                        </div>

                        <div class="upload-folder-img" style="cursor: pointer;">

                        <img src="/`+ encodeURIComponent(j.Path + '/' + j.Name) + `" alt="" id="cimg" >

                        </div>

                        <div class="file-detail media-library-content">

                            <h3>`+ j.Name + `</h3>

                            <p></p>

                        </div>

                    </div>`

                    }
                }

                $('#drivelist').html(str);

            } else {

                $('.media-select-del').children().hide()

                $('#errorhtml').html(`<div class="noData-foundWrapper">

                <div class="empty-folder">

                    <img src="/public/img/folder-sh.svg" alt="">

                    <img src="/public/img/shadow.svg" alt="">

                </div>

                <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

            </div>`)

            }

        }
    })

})

$(document).on('click', '.bclick', function () {

    DeleteListArray = [];

    TrashButton()

    var index = MediaBreadcrumbRoot.indexOf($(this).attr('data-id'))

    MediaBreadcrumbRoot.splice(index + 1, MediaBreadcrumbRoot.length);

    var folderpath = ""

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick"> ` + languagedata.Mediaa.medialibrary + ` </a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if(MediaBreadcrumbRoot.length==1){

            folderpath += MediaBreadcrumbRoot[x]
            
        }else{

            if(x==0){

                folderpath +=  MediaBreadcrumbRoot[x] + "/"
    
            }else if(x==MediaBreadcrumbRoot.length-1){
    
                folderpath +=  MediaBreadcrumbRoot[x] 
    
            }else{
    
                folderpath +=  MediaBreadcrumbRoot[x] + "/" 
            }
        }

        bcrumb += `<li>/</li><li><a href="javascript:void(0)" data-id="` +  MediaBreadcrumbRoot[x] + `" class="bclick">` +  MediaBreadcrumbRoot[x] + `</a></li>`

    }

    $('.mediabreadcrumb').children('ul').html(bcrumb);

    $.ajax({

        type: "post",

        url: "/media/singlefolder",

        data: {

            path: folderpath,

            csrf: $("input[name='csrf']").val()

        },

        datatype: "json",

        cache: false,

        success: function (result) {

            var res = JSON.parse(result)

            $('#drivelist').html('');

            $("#errorhtml").html('')

            $("#mediasearch").val("")

            str = "";

            if (res.folder != null || res.file != null) {

                $('.media-select-del').children().not('#unselectall').show()

                if (res.folder != null) {

                    for (let x of res.folder) {

                        str += `   <div class="upload-folders folderdiv ">

                        <p class="forsearch" style="display: none;">`+ x.Name + `</p>

                        <div class="chk-group chk-group-box">

                            <input type="checkbox" class="ckbox" id="`+ x.Name + `" data-id="` + x.Name + `"  for="` + x.Name + `">

                            <label for="`+ x.Name + `"></label>

                        </div>

                        <div class="upload-folder-img folder" style="cursor: pointer;">

                            <img src="/public/img/folder-media.svg" alt="">

                        </div>

                        <div class="file-detail media-library-content">

                            <h3>`+ x.Name + `</h3>

                            <p></p>

                        </div>

                    </div>`

                    }
                }

                if (res.file != null) {

                    for (let j of res.file) {

                        str += `

                 <div class="upload-folders filediv ">

                 <p class="forsearch" style="display: none;">`+ j.Name + `</p>

                 <div class="chk-group chk-group-box">

                     <input type="checkbox" class="ckbox" id="`+ j.Name + `" data-id="` + j.Name + `" for="` + j.Name + `">

                     <label for="`+ j.Name + `"></label>

                 </div>

                 <div class="upload-folder-img" style="cursor: pointer;">

                 <img src="/`+ encodeURIComponent(j.Path + '/' + j.Name) + `" alt="" id="cimg">

                 </div>

                 <div class="file-detail media-library-content">

                     <h3>`+ j.Name + `</h3>

                     <p></p>

                 </div>

             </div>`
             
                    }

                }

                $('#drivelist').html(str);

            } else {

                $('.media-select-del').children().hide()

                $('#errorhtml').html(`<div class="noData-foundWrapper">

                <div class="empty-folder">

                    <img src="/public/img/folder-sh.svg" alt="">

                    <img src="/public/img/shadow.svg" alt="">

                </div>

                <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

            </div>`)
            }

        }
    })

})

$("#selectall").click(function () {

    $('.ckbox').each(function () {

        var name = $(this).parent().siblings('.file-detail').find('h3').text();

        console.log("name",name);

        if($(this).parents('.upload-folders').css('display')=='block'){

            DeleteListArray.push(name)

            $(this).prop('checked',true)

        }

    })

    TrashButton()

    $(this).hide()

    $("#unselectall").show()

    console.log("media delete array",DeleteListArray);

})

$("#unselectall").click(function () {

    DeleteListArray = []

    $('.ckbox').prop('checked',false)

    TrashButton()

    $(this).hide()

    $("#selectall").show()

    console.log("media delete array",DeleteListArray);

})

$("#mediasearch").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()

    if(keyword!=""){

        PerformMediaSearch(keyword)

    }else{

        $('.noData-foundWrapper').remove()

        $('.upload-folders').show()

        $('.media-select-del>#unselectall').hide()

        $('.media-select-del').children().not('#unselectall').show()

    }

    DeleteListArray = []

    TrashButton()

    $('.upload-folders>.chk-group>.ckbox').prop('checked',false)

})

function PerformMediaSearch(keyword){

    $('.upload-folders').each(function(){

        mediaName = $(this).find('.file-detail>h3').text()

        if(mediaName.includes(keyword)){

            $('.noData-foundWrapper').remove()

            $(this).show()

        }else{

            $(this).hide()

            if($('.upload-folders:visible').length==0){

                $('.media-select-del').children().hide()
    
                $('#errorhtml').html(`<div class="noData-foundWrapper">
    
                <div class="empty-folder">
    
                    <img src="/public/img/folder-sh.svg" alt="">
    
                    <img src="/public/img/shadow.svg" alt="">
    
                </div>
    
                <h1 class="heading">`+ languagedata.oopsnodata + `</h1>
    
            </div>`)
    
            }else{

                $('.media-select-del>#unselectall').hide()

                $('.media-select-del').children().not('#unselectall').show()
    
                $('.noData-foundWrapper').remove()

            }

        }

    })
}





