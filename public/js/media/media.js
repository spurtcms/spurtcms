var languagedata

var DeleteListArray = [];

var MediaBreadcrumbRoot = [];

let cropHeight, cropWidth, dimension, html_imgurl;

var img, canvas, ctx;

let default_cropper, logoexpand_cropper, portrait_cropper, square_cropper;

var breadcrumbLength

var S3GetRouteName = "/image-resize/?name="

var ParentRoute = "media"

var StorageType = ""

var LocalStorageRoot = "storage"
var LocalStoragePath = "storage/media"

/** */
$(document).ready(async function () {

    $('.Content').addClass('checked');

    var languagePath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagePath, function (data) {

        languagedata = data
    })

    StorageType = $("#storagetype").val();

    console.log("type");
})

$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
});

$(document).on('click', '.upload-folder-img>img', function () {

    if (window.location.href.indexOf("media") == -1 && !$(this).parent().hasClass('folder')) {

        decoded_url = decodeURIComponent($(this).attr('src'))

        console.log("chking", decoded_url);

        html_imgurl = "/"

        var cuts = decoded_url.split("/")

        console.log("cuts", cuts);

        for (let x in cuts) {

            if (cuts[x] != "") {

                if (x == cuts.length - 1) {

                    html_imgurl += cuts[x]

                } else {

                    html_imgurl += cuts[x] + "/"
                }
            }
        }

        console.log("src", html_imgurl);

        img = new Image()

        img.src = html_imgurl

        if (window.location.href.indexOf("spaces") != -1) {

            $('#changepicModal .admin-header >h3').text('Crop Space Image')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            default_cropper = $('#crop-container').croppie({

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

        } else if (window.location.href.indexOf("product") != -1) {

            $('#changepicModal .admin-header >h3').text('Crop Product Image')

            $('.admincropimg-row').addClass("grid-row")

            $('.admincropimg-wrap').removeClass('profile-cropimg')

            $('.admincropimg-btns-rht').show()

            $('.admincropimg-btns-rht .active').removeClass('active')

            $('.landscape-btn').addClass('active')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            default_cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 680,

                    height: 340,
                },

                showZoomer: false,

                // fit: true,
            });

            dimension = '680*340'

        } else if (window.location.href.indexOf("entry") != -1) {

            $('#changepicModal .admin-header >h3').text('Crop Entries Image')

            $('.admincropimg-row').addClass("grid-row")

            $('.admincropimg-wrap').removeClass('profile-cropimg')

            $('.admincropimg-btns-rht').show()

            $('.admincropimg-btns-rht .active').removeClass('active')

            $('.landscape-btn').addClass('active')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            default_cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 680,

                    height: 340,
                },

                showZoomer: false,

                // fit: true,
            });

            dimension = '680*340'

        } else if (window.location.href.indexOf("languages") != -1) {

            $('#changepicModal .admin-header >h3').text('Crop Flag Image')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            default_cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 400,

                    height: 280,

                },

                showZoomer: false,

                // fit: true,
            });

            dimension = '400*280'

        } else if (window.location.href.indexOf("category") != -1) {

            $('#changepicModal .admin-header >h3').text('Crop Category Image')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            default_cropper = $('#crop-container').croppie({

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

        } else if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() == "1") {

            $('#changepicModal .admin-header >h3').text('Crop Expand Logo Image')

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

        } else if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() != "1") {

            $('#changepicModal .admin-header >h3').text('Crop Logo Image')

            $('#crop-container').removeClass('croppie-container').empty().show()

            $('#crop-container1').removeClass('croppie-container').empty().hide()

            // $('.cr-slider-wrap').remove()

            default_cropper = $('#crop-container').croppie({

                enableExif: true,

                enableResize: false,

                enableOrientation: true,

                viewport: {

                    width: 300,

                    height: 300,
                },

                showZoomer: false,
                // fit: true,
            });

            dimension = '300*300'
        }

        if ($('.cr-slider-wrap').length > 1) {

            $('.cr-slider-wrap')[1].remove()
        }

        $('.cr-slider-wrap').appendTo($('#zoom-div'));

        $('canvas[class=cr-image]').css('opacity', '0')

        $('.cr-slider').css('display', 'block')

        $('#rotateSlider').val("0")

        $('#mediaDecideModal .imgname').text($(this).parents('.upload-folders').find('.file-detail>h3').text())

        $('#mediaDecideModal #imgsrc-input').val(html_imgurl)

        // $('#changepicModal').modal({backdrop: 'static', keyboard: false}) 

        if (window.location.href.indexOf("categories") == -1) {

            $('#changepicModal').modal('show')

        } else {

            BindCropImageInCategory(html_imgurl)
        }

    }

})

$(document).on('click', '#mediaDecideModal #crop-img', function () {

    $('#mediaDecideModal').modal('hide')

    $('#changepicModal').modal('show')

})

$(document).on('click', '#mediaDecideModal #assign-img', function () {

    var src = $('#mediaDecideModal #imgsrc-input').val()

    IntegrateMediaImages(src)

    $('#mediaDecideModal').modal('hide')

    $('#addnewimageModal').modal('hide')

    $("#rightModal").modal('show')

})


function IntegrateMediaImages(src) {

    if (window.location.href.indexOf("spaces") != -1) {

        BindCropImageInSpace(src)

    } else if (window.location.href.indexOf("entry") != -1) {

        BindCropImageInEntry(src)

    } else if (window.location.href.indexOf("general-settings") != -1) {

        BindCropImageInPersonalize(src)

    } else if (window.location.href.indexOf("addcategory") != -1) {

        BindCropImageInCategory(src)

    } else if (window.location.href.indexOf("product") != -1) {

        BindCropImageInProduct(src)

    } else if (window.location.href.indexOf("") != -1) {

        BindCropImageInLanguage(src)

    }
}


$('#changepicModal').on('shown.bs.modal', function () {

    if ($('#logo-input').val() != "2") {

        canvas = document.querySelector('canvas[class=cr-image]');

        console.log(canvas, "canvas")

        ctx = canvas.getContext('2d');

        let common_cropper

        if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() == "1") {

            common_cropper = logoexpand_cropper

        } else {

            common_cropper = default_cropper
        }

        console.log("cropper src", html_imgurl);

        common_cropper.croppie('bind', {

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

    if ($('#logo-input').val() != "2") {

        let common_cropper

        if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() == "1" && $('.admincropimg-btns-rht:visible').length == 0) {

            common_cropper = logoexpand_cropper

        } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.portrait-btn').hasClass('active')) {

            common_cropper = portrait_cropper

        } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.square-btn').hasClass('active')) {

            common_cropper = square_cropper

        } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.landscape-btn').hasClass('active')) {

            common_cropper = default_cropper

        } else {

            common_cropper = default_cropper

        }

        common_cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg' }).then(function (dataUrl) {

            var imgData = ""

            console.log("storagetype", StorageType);

            if (StorageType == "aws") {

                imgData = GetImageNameS3(html_imgurl)

                console.log("chk--", imgData);

            } else if (StorageType == "local") {

                imgData = GetImageProcessData(html_imgurl)
            }

            console.log("chk", imgData);

            var final_imgName = ''

            const currentTimestamp = new Date().toISOString();

            if (imgData[0].indexOf(`(${dimension})`) == -1) {

                console.log("if--");

                final_imgName = "IMG-" + "(" + dimension + ")(" + currentTimestamp + ").jpg"

            } else {

                console.log("else--");

                final_imgName = imgData[0].split(dimension)[0] + dimension + ")(" + currentTimestamp + ").jpg"

            }

            console.log("final_name", final_imgName);

            if (imgData[0] != "" && imgData[1] != "" && dimension != "") {

                var formData = new FormData()

                formData.append('csrf', $("input[name='csrf']").val())

                formData.append('file', dataUrl)

                console.log("---", html_imgurl);

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

                        if (StorageType == "local") {

                            IntegrateMediaImages(parsed_obj.Path)

                        } else if (StorageType == "aws") {

                            IntegrateMediaImages(S3GetRouteName + ParentRoute + parsed_obj.Path)
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

    if (breadcrumbLength > 0) {

        $('#Refreshdiv').click();
    }

})

$("#addnewimageModal").on('hide.bs.modal', function () {

    breadcrumbLength = MediaBreadcrumbRoot.length

    MediaBreadcrumbRoot = []

})

function GetImageProcessData(html_imgurl) {

    console.log("raw", html_imgurl);

    var ext = html_imgurl.split('/')[3].split('.').pop()

    var imgName = ""

    if (html_imgurl.split('/')[3].split('.').length > 2) {

        for (let part in html_imgurl.split('/')[3].split('.')) {

            if (part < html_imgurl.split('/')[3].split('.').length - 1) {

                imgName = imgName + html_imgurl.split('/')[3].split('.')[part]
            }

        }

    } else {

        imgName = html_imgurl.split('/')[3].split('.')[0]

    }

    return [imgName, ext]
}

function GetImageNameS3(html_imgurl) {

    console.log("raw", html_imgurl);

    var imagename = html_imgurl.split("?name=")[1]

    return imagename

}

function BindCropImageInSpace(src) {

    $("#spimage").val(src)

    var data = $("#spimagehide").attr("src", src);

    if (data != "") {

        // $(".heading-three").hide();

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

        // $('#cenimg').hide()

        // $("#spimagehide").css("height", "23.120rem")

        // $("#spimagehide").css("width", "46.25rem")

        $(".imgSelect-image").hide()

        $(".uplclass").hide()

        $(".imgSelect-btn").hide()

        $(".blogBottom-imgSelect").css("padding-block", "0rem")

        $("#spimagehide").attr("src", src).show()

        $("#spacedel-img").show();

    }

    $("#addnewimageModal").hide()

}

function BindCropImageInPersonalize(src) {

    console.log(src);

    var logo_select = $("#logo-input").val();

    if (logo_select == '1') {

        $("#upload-expandlogo").remove()

        var html = `<div class="upload-logo" id="upload-expandlogo"><img src="` + src + `" alt=""><input type="hidden" id="expandlogo-input" name="expandlogo_imgpath" value="` + src + `"><button class="delete-flag" id="delete-expandlogo"><img src="/public/img/delete-white-icon.svg" alt=""></button><div class="hover-delete-img"></div>`

        $(html).insertAfter('.expandlogodiv')

        $("#addnewimageModal").hide()

        $('.modal-backdrop').remove()

    } else {

        $("#upload-logo").remove()

        var html = `<div class="upload-logo" id="upload-logo"><img src="` + src + `" alt=""><input type="hidden" id="logo-input" name="logo_imgpath" value="` + src + `"><button class="delete-flag" id="delete-logo"><img src="/public/img/delete-white-icon.svg" alt=""></button><div class="hover-delete-img"></div>`

        $(html).insertAfter('.myprouploaddiv')

        $("#addnewimageModal").hide()

        $('.modal-backdrop').remove()

    }
}

function BindCropImageInCategory(src) {

    $("#categoryimage").val(src)

    var data = $("#ctimagehide").attr("src", src);

    if (data != "") {

        $("h3[id=mediadesc]").hide();

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

            <button class="delete-flag"><img id="delete-flag" src="/public/img/delete-white-icon.svg" alt=""></button>

            <div class="hover-delete-img"></div>

            </div>`

    $(flag_uploadHtml).insertBefore('#flag_imgpath-error')

    $('#flag_imgpath-error').hide()

    $("#addnewimageModal").modal('hide')

    $('#laguangeModal').modal('show')

}

function BindCropImageInProduct(src) {

    $('#imgerr').hide()

    var img_uploadHtml = `<div class="hover-lay add-img">

            <input type="hidden" class="product-imgpath" value="`+ src + `">

            <img src="`+ src + `" />

            <div class="black-layout"></div>

            <button type="button" class="del-imgpath"><img src="/public/img/delete-round.svg" alt=""></button>

            </div>`

    $(".productimages").append(img_uploadHtml)

    $("#addnewimageModal").modal('hide')

}

/*Add Folder Button hide and show*/
$(document).on("click", "#addfolder", function () {

    $('.input-fixed').addClass('active')

    // var inputdiv = document.querySelector('#addfoldervalues')

    // var inputfield = document.getElementById("foldername")

    // if (window.location.href.indexOf("media") != -1) {

    //     var mediaDiv = document.querySelector('.media-upload')

    //     if (mediaDiv.scrollHeight > mediaDiv.clientHeight) {

    //         var offset = inputdiv.offsetTop - mediaDiv.offsetTop;

    //         if (mediaDiv.scrollTop <= offset || mediaDiv.scrollTop < offset + 8.75 && $(inputdiv).css('display') != 'none') {

    //             $('#foldername-error').parents('.input-group').removeClass("input-group-error")

    //             $(inputdiv).hide()

    //         } else {

    //             $(inputdiv).show()

    //             // Scroll the div to make the input visible
    //             mediaDiv.scrollTo({ top: offset, behavior: 'smooth' });

    //             $(inputfield).focus()

    //         }

    //     } else {

    //         if ($(inputdiv).css('display') == 'none') {

    //             $(inputdiv).show();

    //             $(inputfield).focus()

    //         } else {

    //             $('#foldername-error').parents('.input-group').removeClass("input-group-error")

    //             $(inputdiv).hide();

    //         }

    //     }

    // } else {

    //     var mediaDiv = document.querySelector('.media-library-list')

    //     if ($(inputdiv).css('display') == 'none') {

    //         $(inputdiv).show();

    //         $(inputfield).focus()

    //     } else {

    //         $('#foldername-error').parents('.input-group').removeClass("input-group-error")

    //         $(inputdiv).hide();

    //     }

    // }

    $('#foldername-error').hide()

    // $(inputfield).val('');

})


// add button 
$("#foldername").keyup(function () {

    $("#namecheck").removeClass('input-group-error')

    $("#foldername-error").hide()

    $('#folderspec-error').hide()

})

$('#foldername[type=text]').keypress(function (event) {

    if (event.which == 13) {

        $('#newfileadd').click()
    }
})

/*Add Folder function */
function addfolder() {

    var name = $("#foldername").val().trim();

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

    } else if (StorageType == "aws") {

        folderpath = ParentRoute

    }

    if (MediaBreadcrumbRoot.length == 0) {

        folderpath += "/"
    }

    for (let x of MediaBreadcrumbRoot) {

        folderpath += "/" + x + "/"

    }

    var specialCharsRegex = /[!@#$%^&*(),.?":{}|<>]/;

    var check = true

    if (specialCharsRegex.test(name)) {

        $('#folderspec-error').show()

        check = false

    } else {

        $('#folderspec-error').hide()
    }

    if (name == "") {

        $('#foldername-error').show();

        $("#namecheck").addClass('input-group-error')

        check = false


    } else {

        $("#namecheck").removeClass('input-group-error')

        $("#foldername-error").hide()

        $("#errorhtml").html('')
    }

    if (check == true) {


        console.log(folderpath);

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

                        $('.input-fixed').removeClass('active')

                        $('#foldername').val('');

                        $('#folderspec-error').hide();

                        $('#foldername-error').hide();

                        // $("#addfoldervalues").hide();

                        $('#Refreshdiv').trigger('click');

                        var message = languagedata?.Toast?.foldercreated

                        if (window.location.href.indexOf("media") != -1) {

                            notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                            $(notify_content).insertBefore(".header-rht");

                        } else {

                            notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                            $(notify_content).insertBefore('#addnewimageModal .modal-dialog')
                        }

                        setTimeout(function () {

                            $('.toast-msg').fadeOut('slow', function () {

                                $(this).remove();

                            });

                        }, 5000);


                    } else {

                        var message = languagedata?.Toast?.foldercreateerrmsg

                        if (window.location.href.indexOf("media") != -1) {

                            notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                            $(notify_content).insertBefore(".header-rht");

                        } else {

                            notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                            $(notify_content).insertBefore("#addnewimageModal .modal-dialog");

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

    }
}


/*ckbox*/
$(document).on('click', '.ckbox', function () {

    // var name = $(this).parent().siblings('.file-detail').find('h3').text();

    var name = $(this).attr('data-id');

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

        $(this).parents('.chk-group-box ').addClass('select-all')

    } else {

        $(this).parents('.chk-group-box ').removeClass('select-all')

        DeleteListArray = jQuery.grep(DeleteListArray, function (value) {

            return value != name;

        });
    }

    console.log("checking media delete array", DeleteListArray);

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

        if (DeleteListArray.length == 1) {

            $('#centerModal .modal-body>#content').text(languagedata.Mediaa.delmediacontent)

        } else if (DeleteListArray.length > 1) {

            $('#centerModal .modal-body>#content').text(languagedata.Mediaa.delmediacontents)
        }

        $('#delid').show();

        $('#delete').attr('href', 'javascript:void(0)');

        $('#delcancel').text(languagedata.no);

    }
})

$(document).on('click', '#media-del-btn', function () {

    if (DeleteListArray.length == 0) {

        var message = languagedata.Mediaa.deleteerr

        if (window.location.href.indexOf("media") != -1) {

            notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore(".header-rht");

        } else {

            notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore("#addnewimageModal .modal-dialog");

        }

        setTimeout(function () {

            $('.toast-msg').fadeOut('slow', function () {

                $(this).remove();
            });

        }, 5000); // 5000 milliseconds = 5 seconds

    }
})

// no btn functionality
$(document).on('click', '#delcancel', function () {

    $('#centerModal').modal('hide');

})

/*Trash */
$(document).on('click', "#delete", function () {

    if (window.location.href.indexOf('media') != -1 || $('#addnewimageModal').css('display') == 'block') {

        $('.toast-msg').remove()

        var folderpath

        if (StorageType == "local") {

            folderpath = ""

        } else if (StorageType == "aws") {

            folderpath = ParentRoute

        }

        if (MediaBreadcrumbRoot.length == 0) {

            folderpath += "/"
        }

        for (let x of MediaBreadcrumbRoot) {

            folderpath += x + "/"
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

                if (response) {

                    $('.dis').removeClass('dis').addClass('disabled');

                    $('#Refreshdiv').trigger('click');

                    $("#delcancel").trigger('click')

                    DeleteListArray = []

                    var message = languagedata.Toast.deletemsg

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore(".header-rht");

                    } else {

                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore('#addnewimageModal .modal-dialog')
                    }

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {

                            $(this).remove();
                        });

                    }, 5000); // 5000 milliseconds = 5 seconds

                } else {

                    var message = languagedata?.Toast?.deleteerrmsg

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore(".header-rht");

                    } else {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore("#addnewimageModal .modal-dialog");

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

/*Upload file*/
$(document).on('change', '#imgupload', function () {

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

        for (let x of MediaBreadcrumbRoot) {

            folderpath += x + "/"

        }

    } else if (StorageType == "aws") {

        folderpath = ParentRoute

        for (let x of MediaBreadcrumbRoot) {

            folderpath += "/" + x + "/"

        }
    }


    if (MediaBreadcrumbRoot.length == 0) {

        folderpath += "/"
    }


    var filename = $(this).val();

    var ext = filename.split(".").pop().toLowerCase();

    if (($.inArray(ext, ["jpg", "png", "jpeg", 'svg']) != -1)) {

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

                if (response) {

                    $('#Refreshdiv').click()

                    var message = languagedata?.Toast?.imgupload

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore(".header-rht");

                    } else {

                        notify_content = '<div style="top:2px;" class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore('#addnewimageModal .modal-dialog')
                    }

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {

                            $(this).remove();
                        });

                    }, 5000); // 5000 milliseconds = 5 seconds

                } else {

                    var message = languagedata.Toast.imguploaderrmsg

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore(".header-rht");

                    } else {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore("#addnewimageModal .modal-dialog");

                    }

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {

                            $(this).remove();

                        });

                    }, 5000);

                }

            },

            error: function (xhr) {

                if (xhr) {

                    var message = ''

                    if (xhr.status == 424) {

                        message = "Imagename shouldnot contains only % symbol"

                    } else if (xhr.status == 413) {

                        message = "Image size should be less than 1mb"

                    } else {

                        message = "unable to upload image in the specified path"
                    }

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore(".header-rht");

                    } else {

                        notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

                        $(notify_content).insertBefore("#addnewimageModal .modal-dialog");

                    }

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {

                            $(this).remove();

                        });

                    }, 5000);

                }
            }
        });
    }
    else {

        var message = languagedata?.Toast?.errmsgupload

        if (window.location.href.indexOf("media") != -1) {

            var notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore(".header-rht");

        } else {

            var notify_content = '<div style="top:2px" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span>' + message + '</span></div>';

            $(notify_content).insertBefore('#addnewimageModal .modal-dialog')
        }

        setTimeout(function () {

            $('.toast-msg').fadeOut('slow', function () {

                $(this).remove();

            });

        }, 5000);
    }
})

$(document).on('click', '#imgupload', function () {

    if ($(this).val() != "") {

        $(this).val("")
    }
})

/*Refresh */
$(document).on('click', '#Refreshdiv', function () {

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

    } else if (StorageType == "aws") {

        folderpath = ParentRoute
    }

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata?.Mediaa?.medialibrary + `</a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if (StorageType == "local") {

            if (MediaBreadcrumbRoot.length == 1) {

                folderpath += MediaBreadcrumbRoot[x]

            } else {

                if (x == 0) {

                    folderpath += MediaBreadcrumbRoot[x] + "/"

                } else if (x == MediaBreadcrumbRoot.length - 1) {

                    folderpath += MediaBreadcrumbRoot[x]

                } else {

                    folderpath += MediaBreadcrumbRoot[x] + "/"
                }
            }


        } else if (StorageType == "aws") {

            folderpath += "/" + MediaBreadcrumbRoot[x]
        }

        bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

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

            var count_text = ''

            if (res.mediaCount != null) {

                if (res.mediaCount > 1) {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafilesavailable + ":" + ` <span class="para">${res.mediaCount}</span> `

                } else {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafileavailable + ":" + ` <span class="para">${res.mediaCount}</span> `
                }
            }

            $('.leftHeader>.caption-flexx>p').html(count_text)

            if (res.Media != null) {

                $('#unselectall').hide()

                $('#selectall').show()

                $('#media-del-btn').show()

                for (let x of res.Media) {

                    var mediaClass, src, fileType, totalSubMedia, mediaNameHtml = ''

                    var subMediaHtml = ''

                    if (x.File === true) {

                        mediaClass = 'folderdiv'

                        src = '/public/img/folder-media.svg'

                        fileType = 'folder'

                        if (x.TotalSubMedia > 1) {

                            totalSubMedia = x.TotalSubMedia + ' Files'

                        } else {

                            totalSubMedia = x.TotalSubMedia + ' File'

                        }

                        subMediaHtml = '<p>' + totalSubMedia + '</p>'

                    } else {

                        mediaClass = 'filediv'

                        if (StorageType == "local") {

                            src = "/" + x.Path + x.Name

                        } else if (StorageType == "aws") {

                            if (MediaBreadcrumbRoot.length >= 1) {

                                src = S3GetRouteName + "/" + x.Path + "/" + x.Name

                            } else {

                                src = S3GetRouteName + "/" + x.Path + x.Name
                            }

                        }

                        fileType = 'file'

                    }

                    var Name = x.Name.replace("/", "")

                    if (x.Name.length >= 20) {

                        mediaNameHtml = `<h3 data-bs-toggle="tooltip" data-bs-custom-class="lms-tooltip" data-bs-html="true" data-bs-placement="top" title="${Name}">${Name}</h3>`

                    } else {

                        mediaNameHtml = `<h3>${Name}</h3>`

                    }

                    str += `<div class="upload-folders ${mediaClass}">

                                <p class="forsearch" style="display: none;">`+ x.AliaseName + `</p>

                                <div class="chk-group chk-group-box">

                                    <input type="checkbox" class="ckbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  for="` + x.AliaseName + `">

                                    <label for="`+ x.AliaseName + `"></label>

                                </div>
  
                                <div class="upload-folder-img ${fileType}">

                                    <img src="${src}" alt="">

                                </div>

                                <div class="file-detail media-library-content">

                                    ${mediaNameHtml}

                                    ${subMediaHtml}

                                </div>

                            </div>`
                }

                $("#drivelist").html(str)

                $('[data-bs-toggle="tooltip"]').tooltip();

            } else {

                $('.media-select-del').children().hide()

                $('#addfoldervalues').hide()

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

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

    } else if (StorageType == "aws") {

        folderpath = ParentRoute
    }

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata.Mediaa.medialibrary + `</a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if (StorageType == "local") {

            if (MediaBreadcrumbRoot.length == 1) {

                folderpath += MediaBreadcrumbRoot[x]

            } else {

                if (x == 0) {

                    folderpath += MediaBreadcrumbRoot[x] + "/"

                } else if (x == MediaBreadcrumbRoot.length - 1) {

                    folderpath += MediaBreadcrumbRoot[x]

                } else {

                    folderpath += MediaBreadcrumbRoot[x] + "/"
                }
            }


        } else if (StorageType == "aws") {

            folderpath += "/" + MediaBreadcrumbRoot[x]
        }

        bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

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

            var count_text = ''

            if (res.mediaCount != null) {

                if (res.mediaCount > 1) {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafilesavailable + ":" + ` <span class="para">${res.mediaCount}</span> `

                } else {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafileavailable + ":" + ` <span class="para">${res.mediaCount}</span> `
                }
            }

            $('.leftHeader>.caption-flexx>p').html(count_text)

            if (res.Media != null) {

                $('.media-select-del').children().not('#unselectall').show()

                for (let x of res.Media) {

                    var mediaClass, src, fileType, totalSubMedia, mediaNameHtml = ''

                    var subMediaHtml = ''

                    if (x.File === true) {

                        mediaClass = 'folderdiv'

                        src = '/public/img/folder-media.svg'

                        fileType = 'folder'

                        if (x.TotalSubMedia > 1) {

                            totalSubMedia = x.TotalSubMedia + ' Files'

                        } else {

                            totalSubMedia = x.TotalSubMedia + ' File'

                        }

                        subMediaHtml = '<p>' + totalSubMedia + '</p>'

                    } else {

                        mediaClass = 'filediv'

                        if (StorageType == "local") {

                            src = "/" + x.Path + "/" + x.Name

                        } else if (StorageType == "aws") {

                            if (MediaBreadcrumbRoot.length >= 1) {

                                src = S3GetRouteName + "/" + x.Path + "/" + x.Name

                            } else {

                                src = S3GetRouteName + "/" + x.Path + x.Name
                            }

                        }
                        fileType = 'file'

                    }

                    var Name = x.Name.replace("/", "")

                    if (x.Name.length >= 20) {

                        mediaNameHtml = `<h3 data-bs-toggle="tooltip" data-bs-custom-class="lms-tooltip" data-bs-html="true" data-bs-placement="top" title="${Name}">${Name}</h3>`

                    } else {

                        mediaNameHtml = `<h3>${Name}</h3>`

                    }

                    str += `<div class="upload-folders ${mediaClass}">

                                <p class="forsearch" style="display: none;">`+ x.AliaseName + `</p>

                                <div class="chk-group chk-group-box">

                                    <input type="checkbox" class="ckbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  for="` + x.AliaseName + `">

                                    <label for="`+ x.AliaseName + `"></label>

                                </div>
  
                                <div class="upload-folder-img ${fileType}">

                                    <img src="${src}" alt="">

                                </div>

                                <div class="file-detail media-library-content">

                                    ${mediaNameHtml}

                                    ${subMediaHtml}

                                </div>

                            </div>`
                }

                $('#drivelist').html(str);

                $('[data-bs-toggle="tooltip"]').tooltip();

            } else {

                $('.media-select-del').children().hide()

                $('#addfoldervalues').hide()

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

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

    } else if (StorageType == "aws") {

        folderpath = ParentRoute
    }

    var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata.Mediaa.medialibrary + `</a></li>`

    for (let x in MediaBreadcrumbRoot) {

        if (StorageType == "local") {

            if (MediaBreadcrumbRoot.length == 1) {

                folderpath += MediaBreadcrumbRoot[x]

            } else {

                if (x == 0) {

                    folderpath += MediaBreadcrumbRoot[x] + "/"

                } else if (x == MediaBreadcrumbRoot.length - 1) {

                    folderpath += MediaBreadcrumbRoot[x]

                } else {

                    folderpath += MediaBreadcrumbRoot[x] + "/"
                }
            }


        } else if (StorageType == "aws") {

            folderpath += "/" + MediaBreadcrumbRoot[x]
        }

        bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

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

            var count_text = ''

            if (res.mediaCount != null) {

                if (res.mediaCount > 1) {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafilesavailable + ":" + ` <span class="para">${res.mediaCount}</span> `

                } else {

                    count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafileavailable + ":" + ` <span class="para">${res.mediaCount}</span> `
                }
            }

            $('.leftHeader>.caption-flexx>p').html(count_text)

            if (res.Media != null) {

                $('.media-select-del').children().not('#unselectall').show()

                for (let x of res.Media) {

                    var mediaClass, src, fileType, totalSubMedia, mediaNameHtml = ''

                    var subMediaHtml = ''

                    if (x.File === true) {

                        mediaClass = 'folderdiv'

                        src = '/public/img/folder-media.svg'

                        fileType = 'folder'

                        if (x.TotalSubMedia > 1) {

                            totalSubMedia = x.TotalSubMedia + ' Files'

                        } else {

                            totalSubMedia = x.TotalSubMedia + ' File'

                        }

                        subMediaHtml = '<p>' + totalSubMedia + '</p>'

                    } else {

                        mediaClass = 'filediv'

                        if (StorageType == "local") {

                            src = "/" + x.Path + "/" + x.Name

                        } else if (StorageType == "aws") {

                            if (MediaBreadcrumbRoot.length >= 1) {

                                src = S3GetRouteName + "/" + x.Path + "/" + x.Name

                            } else {

                                src = S3GetRouteName + "/" + x.Path + x.Name
                            }

                        }

                        fileType = 'file'

                    }

                    var Name = x.Name.replace("/", "")

                    if (x.Name.length >= 20) {

                        mediaNameHtml = `<h3 data-bs-toggle="tooltip" data-bs-custom-class="lms-tooltip" data-bs-html="true" data-bs-placement="top" title="${Name}">${Name}</h3>`

                    } else {

                        mediaNameHtml = `<h3>${Name}</h3>`

                    }

                    str += `<div class="upload-folders ${mediaClass}">

                                <p class="forsearch" style="display: none;">`+ x.AliaseName + `</p>

                                <div class="chk-group chk-group-box">

                                    <input type="checkbox" class="ckbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  for="` + x.AliaseName + `">

                                    <label for="`+ x.AliaseName + `"></label>

                                </div>
  
                                <div class="upload-folder-img ${fileType}">

                                    <img src="${src}" alt="">

                                </div>

                                <div class="file-detail media-library-content">

                                    ${mediaNameHtml}

                                    ${subMediaHtml}

                                </div>

                            </div>`
                }

                $('#drivelist').html(str);

                $('[data-bs-toggle="tooltip"]').tooltip();

            } else {

                $('.media-select-del').children().hide()

                $('#addfoldervalues').hide()

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

        console.log("name", name);

        if ($(this).parents('.upload-folders').css('display') == 'block' && !$(this).is(":checked")) {

            DeleteListArray.push(name)

            $(this).parents('.chk-group-box ').addClass('select-all')

            $(this).prop('checked', true)

        }

    })

    TrashButton()

    $(this).hide()

    $("#unselectall").show()

    console.log("media delete array", DeleteListArray);

})

$("#unselectall").click(function () {

    DeleteListArray = []

    $('.ckbox').prop('checked', false)

    TrashButton()

    $(this).hide()

    $("#selectall").show()

    $('.chk-group-box').each(function(){

        $(this).removeClass('select-all')
    })

    console.log("media delete array", DeleteListArray);

})

$("#mediasearch").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()

    if (keyword != "") {

        PerformMediaSearch(keyword)

    } else {

        $('.noData-foundWrapper').remove()

        $('.upload-folders').show()

        $('.media-select-del>#unselectall').hide()

        $('.media-select-del').children().not('#unselectall').show()

        if ($('.upload-folders:visible').length == 0) {

            $('.media-select-del').children().hide()

            $('#errorhtml').html(`<div class="noData-foundWrapper">

            <div class="empty-folder">

                <img src="/public/img/nodatafilter.svg" alt="">

            </div>

            <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

        </div>`)

        }

    }

    DeleteListArray = []

    TrashButton()

    $('.upload-folders>.chk-group>.ckbox').prop('checked', false)

    CrossCheckCount()

})

function PerformMediaSearch(keyword) {

    $('.upload-folders').each(function () {

        mediaName = $(this).find('.file-detail>h3').text().toLowerCase()

        if (mediaName.includes(keyword)) {

            if ($(this).hasClass('folderdiv')) {

                if (mediaName.includes(keyword) === true) {

                    $(this).show()

                }

            } else if ($(this).hasClass('filediv')) {

                var imagefinalkey = ''

                var imageExt = ''

                var lastDotIndex = mediaName.lastIndexOf(".")

                if (lastDotIndex !== -1) {

                    imagefinalkey = mediaName.substring(0, lastDotIndex)

                    imageExt = mediaName.substring(lastDotIndex, mediaName.length - 1)

                }

                if (imagefinalkey.includes(keyword) === true) {

                    $(this).show()

                } else {

                    $(this).hide()
                }

            }

            if ($('.upload-folders:visible').length > 0) {

                $('.noData-foundWrapper').remove()

            }

        } else {

            $(this).hide()

            if ($('.upload-folders:visible').length == 0) {

                $('.media-select-del').children().hide()

                $('#addfoldervalues').hide()

                $('#errorhtml').html(`<div class="noData-foundWrapper">
    
                <div class="empty-folder">
       
                <img src="/public/img/folder-sh.svg" alt="">
                <img src="/public/img/shadow.svg" alt="">
    
                </div>
    
                <h1 class="heading">`+ languagedata.oopsnodata + `</h1>
    
            </div>`)

            } else {

                $('.media-select-del>#unselectall').hide()

                $('.media-select-del').children().not('#unselectall').show()

                $('.noData-foundWrapper').remove()

            }

        }

    })
}

$(document).on('click', '.portrait-btn', function () {

    $('#crop-container').removeClass('croppie-container').empty().hide()

    $('#crop-container1').removeClass('croppie-container').empty().show()

    $('.admincropimg-wrap').removeClass('profile-cropimg')

    $('.admincropimg-btns-rht .active').removeClass('active')

    $(this).addClass('active')

    portrait_cropper = $('#crop-container1').croppie({

        enableExif: true,

        enableResize: false,

        enableOrientation: true,

        viewport: {

            width: 370,

            height: 440,
        },

        showZoomer: false,

        // fit: true,
    });

    dimension = '370*440'

    if ($('.cr-slider-wrap').length > 1) {

        $('.cr-slider-wrap')[1].remove()
    }

    $('.cr-slider-wrap').appendTo($('#zoom-div'));

    $('canvas[class=cr-image]').css('opacity', '0')

    $('.cr-slider').css('display', 'block')

    $('#rotateSlider').val("0")

    $('#mediaDecideModal .imgname').text($(this).parents('.upload-folders').find('.file-detail>h3').text())

    $('#mediaDecideModal #imgsrc-input').val(html_imgurl)

    canvas = document.querySelector('canvas[class=cr-image]');

    ctx = canvas.getContext('2d');

    portrait_cropper.croppie('bind', {

        url: html_imgurl,

        orientation: 0,

        zoom: 0

    }).then(function () {

        // cropper.croppie('setZoom',0)

        console.log("cropper initialized!");
    })

})

$(document).on('click', '.square-btn', function () {

    $('#crop-container').removeClass('croppie-container').empty().hide()

    $('#crop-container1').removeClass('croppie-container').empty().show()

    $('.admincropimg-wrap').removeClass('profile-cropimg')

    $('.admincropimg-btns-rht .active').removeClass('active')

    $(this).addClass('active')

    square_cropper = $('#crop-container1').croppie({

        enableExif: true,

        enableResize: false,

        enableOrientation: true,

        viewport: {

            width: 370,

            height: 370,
        },

        showZoomer: false,

        // fit: true,
    });

    dimension = '370*370'

    if ($('.cr-slider-wrap').length > 1) {

        $('.cr-slider-wrap')[1].remove()
    }

    $('.cr-slider-wrap').appendTo($('#zoom-div'));

    $('canvas[class=cr-image]').css('opacity', '0')

    $('.cr-slider').css('display', 'block')

    $('#rotateSlider').val("0")

    $('#mediaDecideModal .imgname').text($(this).parents('.upload-folders').find('.file-detail>h3').text())

    $('#mediaDecideModal #imgsrc-input').val(html_imgurl)

    canvas = document.querySelector('canvas[class=cr-image]');

    ctx = canvas.getContext('2d');

    square_cropper.croppie('bind', {

        url: html_imgurl,

        orientation: 0,

        zoom: 0

    }).then(function () {

        // cropper.croppie('setZoom',0)

        console.log("cropper initialized!");
    })

})

$(document).on('click', '.landscape-btn', function () {

    $('#crop-container').removeClass('croppie-container').empty().show()

    $('#crop-container1').removeClass('croppie-container').empty().hide()

    $('.admincropimg-wrap').removeClass('profile-cropimg')

    $('.admincropimg-btns-rht .active').removeClass('active')

    $(this).addClass('active')

    default_cropper = $('#crop-container').croppie({

        enableExif: true,

        enableResize: false,

        enableOrientation: true,

        viewport: {

            width: 680,

            height: 340,
        },

        showZoomer: false,

        // fit: true,
    });

    dimension = '680*340'

    if ($('.cr-slider-wrap').length > 1) {

        $('.cr-slider-wrap')[1].remove()
    }

    $('.cr-slider-wrap').appendTo($('#zoom-div'));

    $('canvas[class=cr-image]').css('opacity', '0')

    $('.cr-slider').css('display', 'block')

    $('#rotateSlider').val("0")

    $('#mediaDecideModal .imgname').text($(this).parents('.upload-folders').find('.file-detail>h3').text())

    $('#mediaDecideModal #imgsrc-input').val(html_imgurl)

    canvas = document.querySelector('canvas[class=cr-image]');

    ctx = canvas.getContext('2d');

    default_cropper.croppie('bind', {

        url: html_imgurl,

        orientation: 0,

        zoom: 0

    }).then(function () {

        // cropper.croppie('setZoom',0)

        console.log("cropper initialized!");
    })

})

function CrossCheckCount() {

    var count_text = ''

    var count = $('.upload-folders:visible').length

    if ($('.upload-folders:visible').length > 0) {

        if (count > 1) {

            count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafilesavailable + ":" + ` <span class="para">` + count + `</span> `

        } else {

            count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafileavailable + ":" + ` <span class="para">` + count + `</span> `
        }

    } else {

        count_text = languagedata.total + " " + languagedata?.Mediaa?.mediafileavailable + ":" + ` <span class="para">` + count + `</span> `

    }

    $('.leftHeader>.caption-flexx>p').html(count_text)

}

$('body').on('click', function () {

    $('#foldername').on('click', function () {

        return false
    })

    $('#newfileadd').on('click', function () {

        addfolder()

        return false
    })

    $('.input-fixed').removeClass('active');

});