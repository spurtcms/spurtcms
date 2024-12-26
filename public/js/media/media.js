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
var oldfilename
var mediafiletype

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

            dimension = '500-250'

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

            dimension = '680-340'

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

            dimension = '680-340'

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

            dimension = '400-280'

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
            dimension = '500-250'

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

            dimension = '500-250'

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

            dimension = '300-300'
        } else if (window.location.href.indexOf("ecommerce") != -1) {

            $('#changepicModal .admin-header >h3').text('Payment Image')

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

            dimension = '680-340'

        } else if (window.location.href.indexOf("blocks") != -1) {
            $('#changepicModal .admin-header >h3').text('Crop Block Image')

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
            dimension = '500-250'
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

    } else if (window.location.href.indexOf("ecommerce") != -1) {

        BindCropImageInPayment(src)

    } else if (window.location.href.indexOf("blocks") != -1) {

        BindCropImageInBlock(src)
    } else if (window.location.href.indexOf("") != -1) {

        BindCropImageInLanguage(src)

    }
}


// $('#changepicModal').on('shown.bs.modal', function () {

//     if ($('#logo-input').val() != "2") {

//         canvas = document.querySelector('canvas[class=cr-image]');

//         console.log(canvas, "canvas")

//         ctx = canvas.getContext('2d');

//         let common_cropper

//         if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() == "1") {

//             common_cropper = logoexpand_cropper

//         } else {

//             common_cropper = default_cropper
//         }

//         console.log("cropper src", html_imgurl);

//         common_cropper.croppie('bind', {

//             url: html_imgurl,

//             orientation: 0,

//             zoom: 0

//         }).then(function () {

//             // cropper.croppie('setZoom',0)

//             console.log("cropper initialized!");
//         })
//     }
// });

// $('#rotateSlider').on('input', function () {

//     var rotationValue = parseInt($(this).val());

//     ctx.clearRect(0, 0, canvas.width, canvas.height);

//     ctx.save()

//     ctx.translate(canvas.width / 2, canvas.height / 2);

//     ctx.rotate((rotationValue * Math.PI) / 180);

//     ctx.drawImage(img, -canvas.width / 2, -canvas.height / 2);

//     ctx.restore()

// });


// $('#crop-button').click(function () {

//     if ($('#logo-input').val() != "2") {

//         let common_cropper

//         if (window.location.href.indexOf("general-settings") != -1 && $('#logo-input').val() == "1" && $('.admincropimg-btns-rht:visible').length == 0) {

//             common_cropper = logoexpand_cropper

//         } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.portrait-btn').hasClass('active')) {

//             common_cropper = portrait_cropper

//         } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.square-btn').hasClass('active')) {

//             common_cropper = square_cropper

//         } else if ($('.admincropimg-btns-rht:visible').length == 1 && $('.landscape-btn').hasClass('active')) {

//             common_cropper = default_cropper

//         } else {

//             common_cropper = default_cropper

//         }

//         common_cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg' }).then(function (dataUrl) {

//             var imgData = ""

//             console.log("storagetype", StorageType);

//             if (StorageType == "aws") {

//                 imgData = GetImageNameS3(html_imgurl)

//                 console.log("chk--", imgData);

//             } else if (StorageType == "local") {

//                 imgData = GetImageProcessData(html_imgurl)
//             }

//             var final_imgName = ''

//             const currentTimestamp = new Date().toISOString();

//             if (imgData[0].indexOf(`(${dimension})`) == -1) {

//                 final_imgName = "IMG-" + "(" + dimension + ")(" + currentTimestamp + ").jpg"

//             } else {

//                 console.log("else--");

//                 final_imgName = imgData[0].split(dimension)[0] + dimension + ")(" + currentTimestamp + ").jpg"

//             }

//             console.log("final_name", final_imgName);

//             if (imgData[0] != "" && imgData[1] != "" && dimension != "") {

//                 var formData = new FormData()

//                 formData.append('csrf', $("input[name='csrf']").val())

//                 formData.append('file', dataUrl)

//                 formData.append('path', html_imgurl.split('/')[1] + "/" + html_imgurl.split('/')[2])

//                 formData.append('name', final_imgName)

//                 $.ajax({

//                     type: "post",

//                     url: "/media/uploadimage",

//                     data: formData,

//                     datatype: "json",

//                     cache: false,

//                     processData: false,

//                     contentType: false,

//                     success: function (result) {

//                         html_imgurl = ""

//                         parsed_obj = JSON.parse(result)

//                         if (StorageType == "local") {

//                             IntegrateMediaImages(parsed_obj.Path)

//                         } else if (StorageType == "aws") {

//                             IntegrateMediaImages(S3GetRouteName + ParentRoute + parsed_obj.Path)
//                         }

//                         $("#rightModal").modal('show')
//                         $("#righrModal2").modal('show')
//                         $('#changepicModal').modal('hide')

//                         $('#crop-container').removeClass('croppie-container').empty().show()

//                         $('#crop-container1').removeClass('croppie-container').empty().hide()

//                         $('#addnewimageModal').modal('hide')
//                     }

//                 })

//             } else {

//                 console.log("imgdata", imgData[0], imgData[1], dimension);

//             }

//         });
//     }
// });

// $('#close,#crop-cancel').click(function () {

//     html_imgurl = ""

//     $('canvas[class=cr-image]').css('opacity', '0')

//     $("#changepicModal").modal('hide');

//     $('#crop-container').removeClass('croppie-container').empty().hide()

//     $('#crop-container1').removeClass('croppie-container').empty().hide()

// });

// $('#changepicModal').on('hide.bs.modal', function (event) {

//     html_imgurl = ""

//     $('canvas[class=cr-image]').css('opacity', '0')

//     $('#crop-container').removeClass('croppie-container').empty().hide()

//     $('#crop-container1').removeClass('croppie-container').empty().hide()

// })

// $("#addnewimageModal").on('show.bs.modal', function () {

//     if (breadcrumbLength > 0) {

//         $('#Refreshdiv').click();
//     }

// })

// $("#addnewimageModal").on('hide.bs.modal', function () {

//     breadcrumbLength = MediaBreadcrumbRoot.length

//     MediaBreadcrumbRoot = []

// })

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

function BindCropImageInBlock(src) {

    alert("ddf", src)

    var data = $("#uploadimg").attr("src", src);

    $("#coverimg").val(src)

    if (data != "") {

        $("p[id=imgtitle]").hide();

        $("#browse").hide();

        // $("#ctimagehide").attr("src", src).show()

        $("#catdel-img").show()

    }

    $("#mediaModal").hide()

    $("#createModal").modal('show')


}

// function BindCropImageInLanguage(src) {

//     $('.file-upload input[name=flag_imgpath]').parents('.upload-json').find('.file-upload').remove()

//     var flag_uploadHtml = `<div class="uploaded-file">

//             <input type="hidden" name="flag_imgpath" value="`+ src + `">

//             <img src="`+ src + `" alt="" class="uploaded-img-flag">

//             <button class="delete-flag"><img id="delete-flag" src="/public/img/delete-white-icon.svg" alt=""></button>

//             <div class="hover-delete-img"></div>

//             </div>`

//     $(flag_uploadHtml).insertBefore('#flag_imgpath-error')

//     $('#flag_imgpath-error').hide()

//     $("#addnewimageModal").modal('hide')

//     $('#laguangeModal').modal('show')

// }

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

                        $('#addFolder').css('display', 'none')

                        $('.modal-backdrop').remove()

                        var message = languagedata.Toast.foldercreated

                        if (window.location.href.indexOf("media") != -1) {

                            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

                            $(notify_content).insertBefore(".header-rht");

                            // setCookie("get-toast", "foldercreated")

                            // setCookie("Alert-msg", "success", 1)

                            // window.location.href="/media/"

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

                        var message = languagedata.Toast.foldercreateerrmsg

                        if (window.location.href.indexOf("media") != -1) {

                            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast[key] + `</p ></div ></div ></li></ul> `;

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

        console.log("trueee")

        $('#deselectid').text("Unselect All")

        $('#deselectid').removeClass("selectall")

        $('#deselectid').addClass("unselectall")

    } else {

        $('#deselectid').text("Select All")

        $('#deselectid').addClass("selectall")

        $('#deselectid').removeClass("unselectall")

        DeleteListArray = jQuery.grep(DeleteListArray, function (value) {

            return value != name;

        });


    }

    if ($(this).is(':checked')) {

        // $('#deselectid').text("Select All")

        DeleteListArray.push(name)

        $(this).parents('.chk-group-box ').addClass('select-all')

        $('.selected-numbers').removeClass("hidden")

        $('.renameimg').removeClass("hidden")

    } else {

        console.log("checkecondition", name)

        $(this).parents('.chk-group-box ').removeClass('select-all')


        DeleteListArray = jQuery.grep(DeleteListArray, function (value) {

            return value != name;

        });

        console.log("checking mediay", DeleteListArray);

    }

    $('.checkboxlength').text(DeleteListArray.length + " " + "itesm selected")



    $('#deselectid').show()

    $('#seleccheckboxdelete').addClass('dis')

    if (DeleteListArray.length > 1) {

        $('#unbulishslt').hide()

        $('#seleccheckboxdelete').removeClass("border-r border-[#717171]")
    } else {
        $('#unbulishslt').show()

        $('#seleccheckboxdelete').addClass("border-r border-[#717171]")

        $('#unbulishslt').attr('data-bs-target', '#addFolder')




    }

    if (DeleteListArray.length == 0) {

        $('.selected-numbers').addClass("hidden")
    }
    filename = $(this).attr("data-name")

    filetype = $(this).attr("data-fname")

    mediafiletype = filetype

    if (mediafiletype == "folderdiv") {

        oldfilename = name.slice(0, -1);
    } else {
        oldfilename = name
    }



    $('#foldername').val(filename)

    $('#renamepath').val(name)

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

        $('#deleteModal').modal('show');

        $(".deltitle").text(languagedata.Mediaa.deltitle)

        if (DeleteListArray.length == 1) {

            $('#content').text(languagedata.Mediaa.delmediacontent)

        } else if (DeleteListArray.length > 1) {

            $('#content').text(languagedata.Mediaa.delmediacontents)
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

            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast[key] + `</p ></div ></div ></li></ul> `;

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
$(document).on('click', "#delid", function () {

    // $(this).attr("data-bs-dismiss","modal")
    $("#deleteModal").hide()

    $(".modal-open").css("overflow", "auto")

    $('.modal-backdrop').remove()

    $('.selected-numbers').addClass('hidden')

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


                    console.log(response, "result")
                    // $('.dis').removeClass('dis').addClass('disabled');

                    $('#Refreshdiv').click()

                    // $("#delcancel").trigger('click')

                    DeleteListArray = []

                    var message = languagedata.Toast.deletemsg

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

                        $(notify_content).insertBefore(".header-rht");

                        // setCookie("get-toast", "deletemsg")

                        // setCookie("Alert-msg", "success", 1)

                        // window.location.href="/media/"

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

                    var message = languagedata.Toast.deleteerrmsg

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

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

                    var message = languagedata.Toast.imgupload

                    if (window.location.href.indexOf("media") != -1) {

                        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

                        $(notify_content).insertBefore(".header-rht");

                        // setCookie("get-toast", "imgupload")

                        // setCookie("Alert-msg", "success", 1)

                        // window.location.href="/media/"

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

                        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

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

                        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

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

        var message = languagedata.Toast.errmsgupload

        if (window.location.href.indexOf("media") != -1) {

            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

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
// $(document).on('click', '#Refreshdiv', function () {

//     var folderpath

//     if (StorageType == "local") {

//         folderpath = ""

//     } else if (StorageType == "aws") {

//         folderpath = ParentRoute
//     }

//     var bcrumb = ` <li><a href="javascript:void(0)" data-id="/" class="bclick">` + languagedata.Mediaa.medialibrary + `</a></li>`

//     for (let x in MediaBreadcrumbRoot) {

//         if (StorageType == "local") {

//             if (MediaBreadcrumbRoot.length == 1) {

//                 folderpath += MediaBreadcrumbRoot[x]

//             } else {

//                 if (x == 0) {

//                     folderpath += MediaBreadcrumbRoot[x] + "/"

//                 } else if (x == MediaBreadcrumbRoot.length - 1) {

//                     folderpath += MediaBreadcrumbRoot[x]

//                 } else {

//                     folderpath += MediaBreadcrumbRoot[x] + "/"
//                 }
//             }


//         } else if (StorageType == "aws") {

//             folderpath += "/" + MediaBreadcrumbRoot[x]
//         }

//         bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

//     }

//     $('.mediabreadcrumb').children('ul').html(bcrumb);

//     $.ajax({

//         type: "post",

//         url: "/media/singlefolder",

//         data: {

//             path: folderpath,

//             csrf: $("input[name='csrf']").val()

//         },

//         datatype: "json",

//         cache: false,

//         success: function (result) {

//             var res = JSON.parse(result)

//             $('.media-library-list').html('');

//             $(".media-error-design").html('')

//             $("#mediasearch").val("")

//             str = "";

//             var count_text = ''

//             if (res.mediaCount != null) {

//                 if (res.mediaCount > 1) {

//                     count_text = languagedata.total + " " + languagedata.Mediaa.mediafilesavailable + ":" + ` <span class="para">${res.mediaCount}</span> `

//                 } else {

//                     count_text = languagedata.total + " " + languagedata.Mediaa.mediafileavailable + ":" + ` <span class="para">${res.mediaCount}</span> `
//                 }
//             }

//             $('.leftHeader>.caption-flexx>p').html(count_text)

//             if (res.Media != null) {

//                 // $('.media-select-del').children().not('#unselectall').show()

//                 $('#unselectall').hide()

//                 $('#selectall').show()

//                 $('#media-del-btn').show()

//                 for (let x of res.Media) {

//                     var mediaClass, src, fileType, totalSubMedia, mediaNameHtml = ''

//                     var subMediaHtml = ''

//                     if (x.File === true) {

//                         mediaClass = 'folderdiv'

//                         src = '/public/img/folder-media.svg'

//                         fileType = 'folder'

//                         if (x.TotalSubMedia > 1) {

//                             totalSubMedia = x.TotalSubMedia + ' Files'

//                         } else {

//                             totalSubMedia = x.TotalSubMedia + ' File'

//                         }

//                         subMediaHtml = '<p>' + totalSubMedia + '</p>'

//                     } else {

//                         mediaClass = 'filediv'

//                         if (StorageType == "local") {

//                             src = "/" + x.Path + x.Name

//                         } else if (StorageType == "aws") {

//                             if (MediaBreadcrumbRoot.length >= 1) {

//                                 src = S3GetRouteName + "/" + x.Path + "/" + x.Name

//                             } else {

//                                 src = S3GetRouteName + "/" + x.Path + x.Name
//                             }

//                         }

//                         fileType = 'file'

//                     }

//                     var Name = x.Name.replace("/", "")

//                     if (x.Name.length >= 20) {

//                         mediaNameHtml = `<h3 data-bs-toggle="tooltip" data-bs-custom-class="lms-tooltip" data-bs-html="true" data-bs-placement="top" title="${Name}">${Name}</h3>`

//                     } else {

//                         mediaNameHtml = `<h3>${Name}</h3>`

//                     }

//                     str += `<div class="upload-folders ${mediaClass}">

//                                 <p class="forsearch" style="display: none;">`+ x.AliaseName + `</p>

//                                 <div class="chk-group chk-group-box">

//                                     <input type="checkbox" class="ckbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  for="` + x.AliaseName + `">

//                                     <label for="`+ x.AliaseName + `"></label>

//                                 </div>

//                                 <div class="upload-folder-img ${fileType}">

//                                     <img src="${src}" alt="">

//                                 </div>

//                                 <div class="file-detail media-library-content">

//                                     ${mediaNameHtml}

//                                     ${subMediaHtml}

//                                 </div>

//                             </div>`
//                 }

//                 $("#drivelist").html(str)

//                 $('[data-bs-toggle="tooltip"]').tooltip();

//             } else {

//                 $('.media-select-del').children().hide()

//                 $('#addfoldervalues').hide()

//                 $('#errorhtml').html(`<div class="noData-foundWrapper">

//                 <div class="empty-folder">

//                 <img src="/public/img/folder-sh.svg" alt="">
//                 <img src="/public/img/shadow.svg" alt="">

//                 </div>

//                 <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

//             </div>`)

//             }

//         }
//     })

// })


/*Folder Inside */
$(document).on('click', '.folder', function () {


    console.log("checkfolderclick")

    DeleteListArray = [];

    TrashButton()

    var foldername = $(this).siblings("p").text().trim();

    console.log(foldername, "foldername")

    MediaBreadcrumbRoot.push(foldername)

    var folderpath

    if (StorageType == "local") {

        folderpath = ""

    } else if (StorageType == "aws") {

        folderpath = ""
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

            folderpath += MediaBreadcrumbRoot[x] + "/"
        }

        bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

    }

    $("#path").val(folderpath)

    $('.mediabreadcrumb').children('ul').html(bcrumb);

    $.ajax({

        type: "post",

        url: "/media/loadmore",

        data: {

            path: $("#path").val(),

            csrf: $("input[name='csrf']").val()

        },

        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },

        datatype: "json",

        cache: false,

        success: function (result) {

            console.log(result, "resultdata")

            $("#overlay").hide();

            var res = JSON.parse(result)

            $("#mediacount").val(parseInt(res.count))

            $('#drivelist').html('');

            $('#drivelist1').html('');

            $('#drivelist2').html('');

            $("#errorhtml").html('')

            $("#mediasearch").val("")

            $(".mediapagination").addClass("hidden")

            $("#drivelist1").siblings("h3").hide()

            str = "";

            strr = "";

            var count_text = ''

            if (res.count != null) {

                if (res.count > 1) {

                    count_text = languagedata.total + " " + languagedata.Mediaa.mediafilesavailable + ":" + ` <span class="para tcount">${res.count}</span> `

                } else {

                    count_text = languagedata.total + " " + languagedata.Mediaa.mediafileavailable + ":" + ` <span class="para tcount">${res.count}</span> `
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

                            src = x.Path + "/" + x.Name

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


                    if (fileType == 'folder') {

                        console.log("folderdiv")

                        str += `      <div class="block  ${mediaClass}" >

                <a href="javascript:void(0)" class="block">
                    <div class="chk-group chk-group-label">

                        <div
                            class=" group p-[16px_16px_16px_8px] border border-[#ECECEC] gap-[16px] rounded-[4px] h-[72px] relative  flex items-center mb-0 text-[14px] font-[400] leading-[1] text-[#262626] tracking-[0.005em]  ">
                            <input type="checkbox" id="`+ x.AliaseName + `" data-fname="folderdiv" data-id="` + x.AliaseName + `" data-name=" ${Name}" class="hidden peer ckbox">
                            <label for="`+ x.AliaseName + `"
                                class=" before:appearance-none before:inline-block  before:min-w-[14px] group-hover:before:w-[14px] before:h-[14px] before:relative before:align-middle before:cursor-pointer group-hover:before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:bg-none  peer-checked:before:bg-[url('/public/img/checked-box.svg')]"></label>
                            <div class="min-w-10 folder">
                                <img src="/public/img/sample-folder.svg" alt="sample-folder">
                            </div>
                            <p
                                class="text-[14px] font-[400] leading-[17.5px] text-[#252525] flex-grow overflow-hidden flex-[1_1_auto] line-clamp-1">
                                 ${Name}</p>
                            <span class="ml-auto text-[12px] font-[400] leading-[15px] text-[#717171]">
                                `+ totalSubMedia + `
                            </span>
                        </div>
                    </div>



                </a>

            </div>`




                    }

                    if (fileType == 'file') {

                        console.log(" inside file")

                        strr += `<div class="block ${mediaClass}">
                <div class="block">
                    <input type="checkbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  data-fname="filediv" data-name="${Name}" class="hidden peer ckbox">
                    <label for="`+ x.AliaseName + `"
                        class="flex flex-col border border-[#ECECEC] rounded-[4px] p-[8px] items-start relative cursor-pointer gap-[6px] mb-0 text-[14px] font-[400] leading-[1] text-[#262626] tracking-[0.005em] before:appearance-none  before:min-w-[14px] before:w-[14px] before:h-[14px] before:relative before:align-middle before:cursor-pointer hover:before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:bg-none  peer-checked:before:bg-[url('/public/img/checked-box.svg')]">
                        <div class="m-[0_-8px_4px_-8px] h-[116px] bg-[#FBFBFB] w-[calc(100%+16px)] imgname">
                                           
                            <img src="${src}" data-id="` + x.AliaseName + `"  class="w-full h-full object-contain object-center" alt="" id="selectimg">
                                          
                        </div>
                        <p data-fullname="`+ x.AliaseName + `" class="whitespace-nowrap overflow-hidden text-ellipsis text-[12px] font-[400] leading-[16px] text-[#152027] w-full">
                           ${Name}</p>

                    </label>
                </div>

            </div>`

                    }

                }


                $('#drivelist1').html(str);
                $('#drivelist2').html(strr);

                $('[data-bs-toggle="tooltip"]').tooltip();

            } else {

                $('.media-select-del').children().hide()

                $('#addfoldervalues').hide()

                $('#errorhtml').html(`   <div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                <div class="text-center w-fit mx-auto">
                                    <img src="/public/img/noData.svg" alt="nodate">
                                </div>
                                <h2
                                    class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
                                   `+ languagedata.oopsnodata + `</h2>
                             
                            </div>`)

            }

        }

    })

    $("#offset").val("1");
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
        folderpath = ""
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

            folderpath += MediaBreadcrumbRoot[x] + "/"
        }

        bcrumb += `<li><a href="javascript:void(0)" data-id="` + MediaBreadcrumbRoot[x] + `" class="bclick">` + MediaBreadcrumbRoot[x] + `</a></li>`

    }

    $('.mediabreadcrumb').children('ul').html(bcrumb);
    $("#path").val(folderpath)
    $("#offset").val("1");

    BreadCrumbAjax()

})

function BreadCrumbAjax() {

    $.ajax({

        type: "post",
        url: "/media/loadmore",
        data: {
            path: $("#path").val(),
            csrf: $("input[name='csrf']").val()
        },

        datatype: "json",
        cache: false,
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        success: function (result) {
            var res = JSON.parse(result)
            $("#overlay").hide();
            $('#drivelist').html('');
            $("#errorhtml").html('')
            $("#mediasearch").val("")
            str = "";
            var count_text = ''
            if (res.count != null) {
                if (res.count > 1) {
                    count_text = languagedata.total + " " + languagedata.Mediaa.mediafilesavailable + ":" + ` <span class="para tcount">${res.count}</span> `
                } else {
                    count_text = languagedata.total + " " + languagedata.Mediaa.mediafileavailable + ":" + ` <span class="para tcount">${res.count}</span> `
                }
            }
            $("#mediacount").val(res.LoadFileCount)
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
                            src = x.Path + "/" + x.Name
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

                $('#errorhtml').html(`<div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                <div class="text-center w-fit mx-auto">
                                    <img src="/public/img/noData.svg" alt="nodate">
                                </div>
                                <h2
                                    class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
                                   `+ languagedata.oopsnodata + `</h2>
                             
                            </div>`)
            }

        }
    })

}

$(document).on("click", ".selectall", function () {


    console.log("checkselectallfunc")


    $('#deselectid').removeClass("selectall")

    $('#deselectid').addClass("unselectall")


    $('.ckbox').each(function () {

        var name = $(this).attr('data-id');

        console.log("name", name);

        if ($(this).parents('.block').css('display') == 'block' && !$(this).is(":checked")) {

            DeleteListArray.push(name)

            $('.checkboxlength').text(DeleteListArray.length + " " + "itesm selected")

            $('#seleccheckboxdelete').addClass('dis')

            if (DeleteListArray.length > 1) {

                $('#unbulishslt').hide()

                $('#seleccheckboxdelete').removeClass("border-r border-[#717171]")
            } else {
                $('#unbulishslt').show()

                $('#seleccheckboxdelete').addClass("border-r border-[#717171]")

                $('#unbulishslt').attr('data-bs-target', '#addFolder')

            }

            if (DeleteListArray.length == 0) {

                $('.selected-numbers').addClass("hidden")
            }

            $(this).parents('.chk-group-box ').addClass('select-all')

            $(this).prop('checked', true)

            $(this).parents('.chk-group-box ').addClass('select-all')

        } else {


        }

    })

    TrashButton()

    $("#deselectid").text("Unselect All")

    $("#unselectall").show()

    console.log("media delete array", DeleteListArray);

})

$(document).on('click', '.unselectall', function () {


    DeleteListArray = []

    $('.ckbox').prop('checked', false)

    $('.selected-numbers').addClass("hidden")


    TrashButton()

    $(this).hide()

    $("#selectall").show()

    $('.chk-group-box').each(function () {

        $(this).removeClass('select-all')
    })

    console.log("media delete array", DeleteListArray);

})

// $("#mediasearch").keyup(function () {

//     var keyword = $(this).val().trim().toLowerCase()

//     if (keyword != "") {

//         PerformMediaSearch(keyword)

//     } else {

//         $('.noData-foundWrapper').remove()

//         $('.upload-folders').show()

//         $('.media-select-del>#unselectall').hide()

//         $('.media-select-del').children().not('#unselectall').show()

//         if ($('.upload-folders:visible').length == 0) {

//             $('.media-select-del').children().hide()

//             $('#errorhtml').html(`<div class="noData-foundWrapper">

//             <div class="empty-folder">

//                 <img src="/public/img/nodatafilter.svg" alt="">

//             </div>

//             <h1 class="heading">`+ languagedata.oopsnodata + `</h1>

//         </div>`)

//         }

//     }

//     DeleteListArray = []

//     TrashButton()

//     $('.upload-folders>.chk-group>.ckbox').prop('checked', false)

//     CrossCheckCount()

// })

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

                $('#errorhtml').html(`<div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                <div class="text-center w-fit mx-auto">
                                    <img src="/public/img/noData.svg" alt="nodate">
                                </div>
                                <h2
                                    class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
                                   `+ languagedata.oopsnodata + `</h2>
                             
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

            count_text = languagedata.total + " " + languagedata.Mediaa.mediafilesavailable + ":" + ` <span class="para tcount">` + count + `</span> `

        } else {

            count_text = languagedata.total + " " + languagedata.Mediaa.mediafileavailable + ":" + ` <span class="para tcount">` + count + `</span> `
        }

    } else {

        count_text = languagedata.total + " " + languagedata.Mediaa.mediafileavailable + ":" + ` <span class="para tcount">` + count + `</span> `

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

// Ecommerce payment image

function BindCropImageInPayment(src) {

    $("#paymentimage").val(src)

    var data = $("#paytimagehide").attr("src", src);
    if (data != "") {
        $("h3[id=mediadesc]").hide();
        $("#browse").hide();
        $("#ctimagehide").attr("src", src).show()
        $("#catdel-img").show()
    }

    $("#addnewimageModal").hide()
    $("#rightModal2").modal('show')

}

function LoadMore(count) {

    $.ajax({
        type: "post",
        url: "/media/loadmore",
        data: {
            csrf: $("input[name='csrf']").val(),
            nextconf: $("#nextcont").val(),
            offset: count,
            search: $("#mediasearch").val(),
            path: $("#path").val()
        },
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        datatype: "json",
        success: function (result) {
            var res = JSON.parse(result)
            $("#nextcont").val(res.nextcont)
            $("#mediacount").val(parseInt($("#mediacount").val()) + parseInt(res.LoadFileCount))
            $("#overlay").hide();
            BindLoadmore(res, "load")
        }

    })

}

function BindLoadmore(res, action) {

    $("#drivelist").empty()

    $("#drivelist1").empty()

    $("#drivelist2").empty()
    str = "";

    strr = "";

    if (res.Media != null) {

        $('.media-select-del').children().not('#unselectall').show()

        for (let x of res.Media) {

            console.log(res.Media, "media values")

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

                console.log(" media path check ")

                mediaClass = 'filediv'

                if (StorageType == "local") {

                    src = x.Path + "/" + x.Name

                    console.log("local path", src);

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

            console.log("media class", mediaClass)

            if (fileType == 'folder') {

                console.log("folderdiv")

                str += ` <div class="block  ${mediaClass}" >

                <a href="javascript:void(0)" class="block">
                    <div class="chk-group chk-group-label">

                        <div
                            class=" group p-[16px_16px_16px_8px] border border-[#ECECEC] gap-[16px] rounded-[4px] h-[72px] relative  flex items-center mb-0 text-[14px] font-[400] leading-[1] text-[#262626] tracking-[0.005em]  ">
                            <input type="checkbox" id="`+ x.AliaseName + `" data-fname="folderdiv" data-id="` + x.AliaseName + `" data-name="${Name}" class="hidden peer ckbox">
                            <label for="`+ x.AliaseName + `"
                                class=" before:appearance-none before:inline-block  before:min-w-[14px] group-hover:before:w-[14px] before:h-[14px] before:relative before:align-middle before:cursor-pointer group-hover:before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:bg-none  peer-checked:before:bg-[url('/public/img/checked-box.svg')]"></label>
                            <div class="min-w-10 folder">
                                <img src="/public/img/sample-folder.svg" alt="sample-folder">
                            </div>
                            <p
                                class="text-[14px] font-[400] leading-[17.5px] text-[#252525] flex-grow overflow-hidden flex-[1_1_auto] line-clamp-1">
                                ${Name}</p>
                            <span class="ml-auto text-[12px] font-[400] leading-[15px] text-[#717171]">
                               `+ totalSubMedia + `
                            </span>
                        </div>
                    </div>



                </a>

            </div>`




            }

            if (fileType == 'file') {

                console.log(" inside file")

                strr += `<div class="block ${mediaClass}">
                <div class="block">
                    <input type="checkbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `" data-fname="filediv" data-name="${Name}" class="hidden peer ckbox">
                    <label for="`+ x.AliaseName + `"
                        class="flex flex-col border border-[#ECECEC] rounded-[4px] p-[8px] items-start relative cursor-pointer gap-[6px] mb-0 text-[14px] font-[400] leading-[1] text-[#262626] tracking-[0.005em] before:appearance-none  before:min-w-[14px] before:w-[14px] before:h-[14px] before:relative before:align-middle before:cursor-pointer hover:before:bg-[url('/public/img/unchecked-box.svg')] before:bg-no-repeat before:bg-contain before:bg-none  peer-checked:before:bg-[url('/public/img/checked-box.svg')]">
                        <div class="m-[0_-8px_4px_-8px] h-[116px] bg-[#FBFBFB] w-[calc(100%+16px)] imgname">
                                           
                            <img src="${src}" data-id="` + x.AliaseName + `"  class="w-full h-full object-contain object-center" alt="" id="selectimg">
                                          
                        </div>
                        <p data-fullname="`+ x.AliaseName + `" class="whitespace-nowrap overflow-hidden text-ellipsis text-[12px] font-[400] leading-[16px] text-[#152027] w-full">
                           ${Name}</p>

                    </label>
                </div>

            </div>`

            }

        }

        $('#errorhtml').html(``);

        if (action == "refresh") {

            console.log("refresh div");

            $('#drivelist1').html(str);

            $('#drivelist2').html(strr)

        } else {
            console.log("media append in mediamodal");

            $('#drivelist1').append(str);

            $('#drivelist2').append(strr);

            // $('#drivelist').append(str);
        }


        $('[data-bs-toggle="tooltip"]').tooltip();

        $('.tcount').text(res.count)

    } else {

        console.log("loadnodata")

        $('.media-select-del').children().hide()

        $('#addfoldervalues').hide()

        // if ($("#drivelist").children('.upload-folders').length == 0) {

        $("#drivelist").html(``);

        $('#errorhtml').html(

            ` <div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                <div class="text-center w-fit mx-auto">
                                    <img src="/public/img/noData.svg" alt="nodate">
                                </div>
                                <h2
                                    class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
                                   `+ languagedata.oopsnodata + `</h2>
                             
                            </div>`

        )

        // }
    }
}

// if (window.location.href.indexOf("media") > -1) {
//     // Detect scroll to bottom
//     $(window).scroll(function () {
//         if (Math.ceil($(window).scrollTop() + $(window).height()) + 1 >= $(document).height()) {
//             if ($("#mediacount").val().trim() != $(".tcount").text().trim()) {
//                 var count = parseInt($('#offset').val());
//                 $('#offset').val(count + 1);
//                 LoadMore(count);
//             }
//         }
//     });
// }

$(document).on('click', "#Refreshdiv", function () {

    $.ajax({
        type: "post",
        url: "/media/loadmore",
        data: {
            csrf: $("input[name='csrf']").val(),
            nextconf: $("#nextcont").val(),
            path: $("#path").val(),
            offset: $("#offset").val()
        },
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        datatype: "json",
        success: function (result) {

            console.log(result, "resulttt")
            var res = JSON.parse(result)
            BindLoadmore(res, "refresh")
            $('#offset').val('1');
            $("#overlay").hide();
            $('#mediacount').val(res.LoadFileCount)
        }

    })
})

//----------------- Entries Media started -------------------

function Ajax() {

    $.ajax({
        type: "post",
        url: "/media/loadmore",
        data: {
            search: $("#mediasearch").val(),
            path: "",
            csrf: $("input[name='csrf']").val(),
            offset: 0
        },
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        datatype: "json",
        success: function (result) {
            var res = JSON.parse(result)
            $("#nextcont").val(res.nextcont);
            $("#overlay").hide();
            $("#mediacount").val(res.LoadFileCount)
            $("#totalmediacount").val(res.count)
            BindLoadmore(res, "load")
        }

    })
}
// Entries media load

$(document).on('click', '.imgSelect-btn', function () {
    Ajax()
})

// Product media load

$(document).on('click', '.productmedia', function () {
    Ajax()
})


$(document).ready(function () {

    var $scrollableDiv = $('.model-medialib');
    $scrollableDiv.on('scroll', function () {
        if (Math.ceil($scrollableDiv.scrollTop() + $scrollableDiv.innerHeight()) + 3 >= $scrollableDiv[0].scrollHeight) {
            if ($("#mediacount").val().trim() != $("#totalmediacount").val().trim()) {
                var count = parseInt($('#offset').val());
                $('#offset').val(count + 1);
                LoadMore(count)
            }
        }
    });
})

//----------------- Entries Media end ---------------

//---------------- Category Media start -------------


$(document).on('click', '#browse', function () {

    Ajax()

})

// $(document).ready(function () {
//     // var count = 1
//     var $scrollableDiv = $('#browse');
//     $scrollableDiv.on('scroll', function () {
//         if (Math.ceil($scrollableDiv.scrollTop() + $scrollableDiv.innerHeight()) + 3 >= $scrollableDiv[0].scrollHeight) {
//             $('#offset').val(count + 1);
//             LoadMore(count)
//         }
//     });
// })

//--------------- Category Media end ------------

//--------------- Media Search start ------------

function search(keyword) {

    $.ajax({
        type: "post",
        url: "/media/loadmore",
        data: {
            search: keyword,
            csrf: $("input[name='csrf']").val(),
            offset: 0
        },
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        datatype: "json",
        success: function (result) {
            var res = JSON.parse(result)
            $("#nextcont").val(res.nextcont);
            $("#overlay").hide();
            BindLoadmore(res, "refresh")
        }

    })
}

$(document).on('keypress', '#mediasearch', function (e) {

    if (e.which == 13) {

        search($(this).val())

    }

})

//--------------- Media Search end --------------

//------------- hide & show (delete/selectall) ----

// $(document).on('click', '.ckbox', function () {

//     $('.ckbox').each(function () {

//         console.log($(this).is(":checked"));

//         if ($(this).is(":checked") == true) {

//             console.log("inside");

//             $(".media-select-del").css('opacity', '1')

//             return false;

//         } else {

//             $(".media-select-del").css('opacity', '0')
//         }
//     })

// })
$(document).on('click', '.add-galleryfolder', function () {
    $("#addnewimageentriesModal").modal('show')
    Ajaxload()
})
function BindLoadmoreaddtionalfield(res, action) {

    str = "";

    if (res.Media != null) {

        $('.media-select-del').children().not('#unselectall').show()

        for (let x of res.Media) {

            console.log("newfolder upload", x);
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

            }

            var Name = x.Name.replace("/", "")

            if (x.Name.length >= 20) {

                mediaNameHtml = `<h3 data-bs-toggle="tooltip" data-bs-custom-class="lms-tooltip" data-bs-html="true" data-bs-placement="top" title="${Name}">${Name}</h3>`

            } else {

                mediaNameHtml = `<h3 id="fl-name">${Name}</h3>`

            }

            if (x.File === true) {
                str += `<div class="upload-folders ${mediaClass}">

                        <p class="forsearch" style="display: none;">`+ x.AliaseName + `</p>

                        <div class="chk-group chk-group-box">

                            <input type="checkbox" class="ckbox" id="`+ x.AliaseName + `" data-id="` + x.AliaseName + `"  for="` + x.AliaseName + `">

                            <label for="`+ x.AliaseName + `"></label>

                        </div>

                        <div class="upload-folder-img mediagallery-folder ${fileType}">

                            <img src="${src}" alt="">

                        </div>

                        <div class="file-detail media-library-content">

                            ${mediaNameHtml}

                            ${subMediaHtml}

                        </div>

                    </div>`
            }
        }

        $('#errorhtml').html(``);

        if (action == "refresh") {

            $('#drivelistdata').html(str);

        } else {

            $('#drivelistdata').append(str);
        }


        $('[data-bs-toggle="tooltip"]').tooltip();

        $('.tcount').text(res.count)

    } else {

        $('.media-select-del').children().hide()

        $('#addfoldervalues').hide()

        // if ($("#drivelist").children('.upload-folders').length == 0) {

        $("#drivelistdata").html(``);

        $('#errorhtml').html(`<div class="max-w-[328px] mx-auto text-center m-[120px_16px]">
                                <div class="text-center w-fit mx-auto">
                                    <img src="/public/img/noData.svg" alt="nodate">
                                </div>
                                <h2
                                    class="text-[#262626] text-center text-[18px] font-medium leading-[22.5px] mb-[6px] ">
                                   `+ languagedata.oopsnodata + `</h2>
                             
                            </div>`)

        // }
    }
}
function Ajaxload() {

    $.ajax({
        type: "post",
        url: "/media/loadmore",
        data: {
            search: $("#mediasearch").val(),
            path: "",
            csrf: $("input[name='csrf']").val(),
            offset: 0
        },
        beforeSend: function () {
            $("#overlay").show();
        },
        error: function () { // if error occured
            $("#overlay").hide();
        },
        datatype: "json",
        success: function (result) {
            var res = JSON.parse(result)
            $("#nextcont").val(res.nextcont);
            $("#overlay").hide();
            $("#mediacount").val(res.LoadFileCount)
            $("#totalmediacount").val(res.count)
            BindLoadmoreaddtionalfield(res, "load")
        }

    })
}

// add folder
$(document).on('click', '.addfolderbtn', function () {

    $('.foldereditcon').text("Folder Name")

    $('#newfileadd').removeClass('hidden')

    $('#updatefileadd').addClass('hidden')

    $('#foldername').val("")

    $('#foldername-error').hide()
})

// rename in folder and image
$(document).on('click', '#unbulishslt', function () {

    $('.foldereditcon').text("Edit Folder Name")

    $('#newfileadd').addClass('hidden')

    $('#updatefileadd').removeClass('hidden')


})


$(document).on("click", '#updatefileadd', function () {

    newfilename = $('#foldername').val()


    filepath = $('#renamepath').val()

    $.ajax({

        type: "post",
        url: "/media/renamemediapath",
        data: {
            oldfilename: oldfilename,
            newfilename: newfilename,
            path: ParentRoute,
            filetype: mediafiletype,
            csrf: $("input[name='csrf']").val(),

        },

        datatype: "json",
        success: function (result) {

            $('.modal-backdrop').remove()

            $("#addFolder").hide()

            $(".selected-numbers").addClass("hidden")

            $(".modal-open").css("overflow", "auto")

            DeleteListArray = [];

            if (result) {

                $('#Refreshdiv').click()

                var message = languagedata.Toast.Mediarename

                if (window.location.href.indexOf("media") != -1) {

                    notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + message + `</p ></div ></div ></li></ul> `;

                    $(notify_content).insertBefore(".header-rht");

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

// media model refresh in block
// $(document).on("click", "#blockbrowse", function () {
//     $('#Refreshdiv').click();

// })