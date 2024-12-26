let searchParams = new URLSearchParams(window.location.search)
var languagedata
var filename

//set modulename cookie
// $(document).on('click','.mainmenu',function(){
//     setCookie("modulename",$(this).children('span').text(),3600)
//     delete_cookie("tabname")
//     window.location.href = $(this).attr('data-route')
//     setCookie("tabroute",$(this).attr('data-route'))
// })

// $(document).on('click','.nav li a',function(){
//     setCookie("tabname",$(this).children('span').text(),3600)
// })

// $(document).on('click','.list-1',function(){
//     // setCookie("modulename",'Channels',3600)
//     window.location.href = $(this).attr('data-route')
// })
$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === 'k') {
        event.preventDefault();
            $(".srchBtn-togg").next('input').toggleClass('w-[calc(100%-36px)] h-full block ');
            $(".srchBtn-togg").next('input').toggleClass('hidden ');
            $(".srchBtn-togg").parent('div').toggleClass('w-[32px] w-[300px] justify-start p-2.5 border border-[#ECECEC] rounded-sm gap-3 overflow-hidden ');
        $(".search").focus().select();
    }
});
$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
});

$(document).ready(function () {

    var imgurl, img, canvas, ctx;
    let cropper;


    if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 || window.location.href.indexOf("customer") != -1 || window.location.href.indexOf("general-settings/") != -1 || window.location.href.indexOf("languages/") != -1 || window.location.href.indexOf("blocks") != -1 || window.location.href.indexOf("categories") != -1) {

        console.log("test");


        // Image Value Empty
        $("#profileImgLabel,#myfile,#cmpymyfile,#profileImg,#flagFile,#blockimg,#categoryimage").click(function () {
            $(this).val("")
        })

        var viewportType
        var width
        var height

        // Change Image with cropper
        $("#profileImgLabel,#myfile,#cmpymyfile,#profileImg,#logoLabel,#flagFile,#blockimg,#categoryimage").change(function () {


            var file;

            if ($(this).attr("id") == 'profileImgLabel') {

                // file = $('#profileImage')[0].files[0];
                // console.log(file, "profileImage");

                file = $('#profileImgLabel')[0].files[0];
                console.log(file, "profileImgLabel");

            } else if ($(this).attr("id") == 'myfile') {

                file = $('#myfile')[0].files[0];

                console.log(file, "file");

            } else if ($(this).attr("id") == 'cmpymyfile') {
                file = $('#cmpymyfile')[0].files[0];
                console.log(file, "cmpymyfile");

            } else if ($(this).attr("id") == 'profileImg') {

                file = $('#profileImage')[0].files[0];
                console.log(file, "profileImage");
            } else if ($(this).attr("id") == 'logoLabel') {

                file = $('#logo')[0].files[0];
                console.log(file, "cmpymyfile");

            } else if ($(this).attr("id") == 'flagFile') {

                file = $(this)[0].files[0];
                console.log(file, "flagFile");

            } else if ($(this).attr("id") == 'blockimg') {
                file = $("#blockimg")[0].files[0];
                console.log(file, "BlockFile");
            } else if ($(this).attr("id") == 'categoryimage') {
                file = $("#categoryimage")[0].files[0];
                console.log(file, "categoryimage");
            }

            filename = file.name;
            console.log(filename);
            // $("#blockfilename").text(filename)

            var ext = filename.split(".").pop().toLowerCase();


            if (($.inArray(ext, ["jpg", "png", "jpeg"]) != -1)) {

                $("#myfile-error").addClass('hidden')

                console.log("crop is worked");

                $('#imageSpace').removeClass('croppie-container').empty().show()

                if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 || window.location.href.indexOf("customer") != -1 || window.location.href.indexOf("general-settings/") != -1 || window.location.href.indexOf("languages/") != -1 || window.location.href.indexOf("blocks") != -1 || window.location.href.indexOf("categories") != -1) {
                    console.log("F2");

                    if (window.location.href.indexOf("myprofile") != -1) {

                        // $("#logo-input").val("2")

                        $("#prof-crop").attr('data-id', '5')

                        $('#cricleCropper-head').text('Crop Myprofile Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#myProfileImg').addClass('hidden')


                    } else if (window.location.href.indexOf("users") != -1) {

                        $('#cricleCropper-head').text('Crop User Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#userModal').modal('hide')

                    } else if (window.location.href.indexOf("member") != -1) {

                        $('#cricleCropper-head').text('Crop Member Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#modalId2').modal('hide')
                        $('body').css("overflow", "visible")


                    } else if (window.location.href.indexOf("customer") != -1) {

                        $('#changepicModal .admin-header >h3').text('Crop Customer Image')

                    } else if (window.location.href.indexOf("general-settings/") != -1) {

                        $("#prof-crop").attr('data-id', '4')

                        $('#cricleCropper-head').text('Crop Company Logo Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()

                    } else if (window.location.href.indexOf("languages/") != -1) {
                        $("#prof-crop").attr('data-id', '7')

                        $('#cricleCropper-head').text('Crop Flag Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#LanguageModal').modal('hide')

                    } else if (window.location.href.indexOf("blocks") != -1) {
                        $("#prof-crop").attr('data-id', '8')

                        $('#cricleCropper-head').text('Crop Block Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#createModal').modal('hide')
                        $('body').css("overflow", "visible")

                    } else if (window.location.href.indexOf("categories") != -1) {
                        $("#prof-crop").attr('data-id', '9')

                        $('#cricleCropper-head').text('Crop Category Image')
                        $('.CropperSection').removeClass('hidden')
                        $('.mainsection').hide()
                        $('#addCategory').modal('hide')

                    }

                }

                if ($(this).attr("id") == 'cmpymyfile' || $(this).attr('id') == 'categoryimage') {

                    viewportType = 'square'
                    $('#extrawidth').show()
                    

                } else if ($(this).attr('id') == 'logoLabel' || $(this).attr('id') == 'flagFile' || $(this).attr('id') == 'blockimg') {

                    viewportType = 'square'
                    $('#extrawidth').show()
                } else {

                    viewportType = 'circle'
                    width = 600
                    height = 600,
                        $('#extrawidth').hide()

                }



                console.log(viewportType, "view");

                var reader = new FileReader();
                reader.onload = function (event) {
                    imgurl = event.target.result

                    img = new Image()
                    img.onload = function () {
                        cropper = $('#imageSpace').croppie({
                            url: imgurl,
                            enableExif: true,
                            enableResize: false,
                            enableOrientation: true,
                            viewport: {
                                width: width,
                                height: height,
                                type: viewportType
                            },
                            aspectRatio:16/9,
                            showZoomer: false,
                        });
                        $('canvas[class=cr-image]').css('opacity', '0')
                        if ($('.cr-slider-wrap').length > 1) {
                            $('.cr-slider-wrap')[1].remove()
                        }
                        $('.cr-slider-wrap').prependTo($('#zoom-div'));
                        $('#rotateSlider').val("0")
                        $('.cr-slider').css('display', 'block')
                        $('.cr-slider').css('color', '#D6D6D6')
                        // $('.cr-slider').addClass('bg-[#D6D6D6]')
                        // $('.cr-slider').css('opacity', '0')
                        $('.cr-slider').on('input', function () {
                            var sliderValue = parseFloat($(this).val());
                            var zoomValue = sliderValue.toFixed(1);

                            cropper.croppie('setZoom', parseFloat(zoomValue));

                            $('#zoomvalue').text(zoomValue);
                        });


                        $(document).ready(function () {
                            $('.cr-slider').on('input', function () {
                                var sliderValue = parseFloat($(this).val());
                                var zoomValue = sliderValue.toFixed(2);
                                var zoomValueNumber = parseFloat(zoomValue);
                                cropper.croppie('setZoom', zoomValueNumber);
                                $('#zoomvalue').text(zoomValue);
                                var sliderMin = parseFloat($('.cr-slider').attr('min'));
                                var sliderMax = parseFloat($('.cr-slider').attr('max'));
                                var sliderWidth = $('.cr-slider').width();
                                var zoomValueWidth = $('#zoomvalue').outerWidth();
                                var zoomValuePercent = (sliderValue - sliderMin) / (sliderMax - sliderMin);
                                var newPosition = zoomValuePercent * (sliderWidth - zoomValueWidth);
                                $('#zoomvalue').css('transform', 'translateX(' + newPosition + 'px)');
                            });
                        });



                    }



                    img.src = imgurl
                }
                reader.readAsDataURL(file);

                $('#changepicModal').modal('show');

            } else {


                if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 || window.location.href.indexOf("customer") != -1 || window.location.href.indexOf("general-settings/") != -1 || window.location.href.indexOf("blocks") != -1 || window.location.href.indexOf("categories") != -1 || window.location.href.indexOf("languages/") != -1) {

                    $("#myfile-error").removeClass('hidden')

                    if ((window.location.href.indexOf("myprofile") != -1)) {
                        $('#profPicWarn').addClass('hidden')

                    }



                    console.log("profie have error");
                    // $('"#myfile-error').css('display', 'block !important')
                }
            }
        });


        // $("#cmpymyfile").change(function () {

        //     $('#crop-container').removeClass('croppie-container').empty().show()

        //     if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1) {

        //         if (window.location.href.indexOf("member") != -1) {

        //             $('#changepicModal .admin-header >h3').text('Crop Company Image')

        //         }

        //         $("#cmpymyfile-error").hide()
        //     }
        //     var file = this.files[0];
        //     var filename = $(this).val();
        //     var ext = filename.split(".").pop().toLowerCase();
        //     if (($.inArray(ext, ["jpg", "png", "jpeg"]) != -1)) {
        //         var reader = new FileReader();
        //         reader.onload = function (event) {
        //             imgurl = event.target.result
        //             img = new Image()
        //             img.onload = function () {
        //                 cropper = $('#crop-container').croppie({
        //                     enableExif: true,
        //                     enableResize: false,
        //                     enableOrientation: true,
        //                     viewport: {
        //                         width: 300,
        //                         height: 300,
        //                         type: 'circle'
        //                     },
        //                     showZoomer: false,
        //                 });
        //                 $('canvas[class=cr-image]').css('opacity', '0')
        //                 if ($('.cr-slider-wrap').length > 1) {
        //                     $('.cr-slider-wrap')[1].remove()
        //                 }
        //                 $('.cr-slider-wrap').appendTo($('#zoom-div'));
        //                 $('#rotateSlider').val("0")
        //                 $('.cr-slider').css('display', 'block')
        //             }

        //             img.src = imgurl
        //         }
        //         reader.readAsDataURL(file);
        //         $('#changepicModal').modal('show');
        //     } else {

        //         if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1) {
        //             $("#cmpymyfile-error").text(languagedata?.Toast?.errmsgupload).show()
        //         }
        //     }
        // });

        // $('#open-cricleCropper').on('shown.bs.modal', function (event) {
        //     console.log("testsst");

        //     console.log(imgurl);

        //     $('#open-cricleCropper #cropImage').attr('src', imgurl)
        //     cropper.croppie('bind', {
        //         url: imgurl
        //     }).then(function () {
        //         cropper.croppie('setZoom', 0)
        //         console.log("cropper initialized!");
        //     })
        //     canvas = document.querySelector('canvas[class=cr-image]');
        //     ctx = canvas.getContext('2d');
        // });


        $('#changepicModal').on('shown.bs.modal', function () {

            cropper.croppie('bind', {
                url: imgurl,
            }).then(function () {
                cropper.croppie('setZoom', 0)
                console.log("cropper initialized!");
            })
            canvas = document.querySelector('canvas[class=cr-image]');
            ctx = canvas.getContext('2d');
        });

        $('#rotateSlider').on('input', function () {
            var rotationValue = parseInt($(this).val(), 10);
            $('#rotateValue').text(rotationValue + '°');
            var canvas = document.querySelector('canvas');
            var ctx = canvas.getContext('2d');
            ctx.clearRect(0, 0, canvas.width, canvas.height);
            ctx.save();
            ctx.translate(canvas.width / 2, canvas.height / 2);
            ctx.rotate((rotationValue * Math.PI) / 180);
            ctx.drawImage(img, -canvas.width / 2, -canvas.height / 2);
            ctx.restore();
        });
        $('#rotateSlider').on('input', function () {
            var value = $(this).val();
            $('#rotateValue').css('transform', 'translateX(' + value + 'px)');
        });

        $('#rotateRight').on('click', function () {
            if (cropper) {
                cropper.croppie('rotate', 90); // Rotate the image 90 degrees clockwise
            }
        })

        $('#rotateLeft').on('click', function () {
            if (cropper) {
                cropper.croppie('rotate', -90); // Rotate the image 90 degrees counterclockwise
            }
        });

        // $('#rotateRight').on('click', function () {
        //     var rotationValue = parseInt($('#rotateSlider').val());
        //     ctx.clearRect(0, 0, canvas.width, canvas.height);
        //     ctx.save()
        //     ctx.translate(canvas.width / 2, canvas.height / 2);
        //     ctx.rotate(((rotationValue * Math.PI) / 180) + 30);
        //     ctx.drawImage(img, -canvas.width / 2, -canvas.height / 2);
        //     ctx.restore()
        // })



        $('#crop-button').click(function () {
            console.log("works");

            var in_val
            if (window.location.href.indexOf("member") != -1 || window.location.href.indexOf("users") != -1) {
                in_val = $("#prof-crop").val()
            } else {
                in_val = $("#prof-crop").attr('data-id')
            }

            //Member
            if (in_val == 1) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', circle: true, }).then(function (dataUrl) {

                    $('#profpic-mem').attr('src', dataUrl)
                    // $('#mem-img').attr('src', dataUrl)
                    // $("#name-string").hide()
                    $('#memcrop_base64').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('#myfile-error').addClass('hidden')
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                    $('#modalId2').modal('show')
                });
            }


            if (in_val == 2) {
                console.log("mem edit");
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'png', }).then(function (dataUrl) {
                    $('#profpic2').attr('src', dataUrl)
                    $("#cname-string").hide()
                    $('input[name=crop_data2]').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#crop-container').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                });

            }
            //Member Edit
            if (in_val == 3) {
                console.log("mem edit");
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'png', }).then(function (dataUrl) {
                    $('#profpicture').attr('src', dataUrl)
                    $("#cname-string").hide()
                    $('input[name=crop_data]').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#crop-container').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                });

            }

            if (in_val == 5) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', circle: true, }).then(function (dataUrl) {
                    console.log("mem");
                    $('#prof-crop').attr('src', dataUrl)
                    $('#cropData').attr('value', dataUrl)
                    $('#profPicWarn').addClass('hidden')
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('#myfile-error').addClass('hidden')
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                });
            }
            //Member Profile
            if (in_val == 4) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', }).then(function (dataUrl) {
                    $('#prof-crop').addClass('w-full')
                    $('#prof-crop').addClass('h-[180px]')
                    $('#prof-crop').attr('src', dataUrl)
                    $('#cropData').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                });
            }


            if (in_val == 6) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', }).then(function (dataUrl) {
                    $('#prof-crop').addClass('w-full')
                    $('#prof-crop').addClass('h-[180px]')
                    $('#profpic-user').attr('src', dataUrl)
                    $('#crop_data').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                    $('#userModal').modal('show')
                });
            }

            if (in_val == 7) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', }).then(function (dataUrl) {
                    $('#flagFile-error').css('display', 'none')
                    $('#flagSpan').addClass('hidden')
                    $('#flaggFileLabel').addClass('hidden')
                    $('#flagPara').addClass('hidden')
                    $('#prof-crop').addClass('h-[180px]')
                    $('#prof-crop').attr('src', dataUrl)
                    $('#cropData').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('#flagDiv').append('<div class="hover-delete-img "></div>');
                    $('#flagDiv').append('<button class="deleteFlag"><img id="deleteFlag" src="/public/img/deleteWhiteIcon.svg" alt=""></button>')
                    $('#myfile-error').addClass('hidden')
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                    $('#LanguageModal').modal('show')
                });
            }
            // Block cover Image upload
            if (in_val == 8) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', }).then(function (dataUrl) {

                    $("#coverimg").val(dataUrl)
                    $('#categoryimages').val(dataUrl)
                    $('#blockcoverimg').attr('src', dataUrl)
                    $('#blockcoverimg').show()
                    $('#browse').hide()
                    $('#uploadLine').hide()
                    $('#uploadFormat').hide()
                    $('#catdel-img').show()
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                    $('#addCategory').modal('show')
                    $("#coverimg-error").hide()
                    // if (dataUrl != "") {

                    //     html = `  <div class="border border-[#ECECEC] p-[8px] px-[12px] flex items-center gap-[12px] rounded-[5px] ">
                    //          <span class="w-[24px] min-w-[24px] h-[24px]"><img id="uploadimg" src="` + dataUrl + ` " alt=""></span>
                    //          <p class="text-[14px] font-normal leading-[17.5px] text-[#252525] text-left" id="blockfilename" > `+ filename + ` </p>
                    //          <a href="javascript:void(0);" class="min-w-[16px] ml-auto" id="removeimg">
                    //              <img src="/public/img/removeCover.svg" alt="remove">
                    //          </a>
                    //      </div> `

                    //     $(html).insertAfter(".imagediv");
                    // }
                    // $('canvas[class=cr-image]').css('opacity', '0')
                    // $("#changepicModal").modal('hide');
                    // $('#imageSpace').removeClass('croppie-container').empty()
                    // $("#coverimg-error").css("display", "none")
                    // $('.CropperSection').addClass('hidden')
                    // $('.mainsection').show()
                    // $('body').css("overflow", "hidden")
                    // $('#browse').addClass('hidden')
                    // $('#uploadFormat').hide()
                });
            }

            if (in_val == 9) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', }).then(function (dataUrl) {
                    $('#categoryimages').val(dataUrl)
                    $('#ctimagehide').attr('src', dataUrl)
                    $('#ctimagehide').show()
                    $('#browse').hide()
                    $('#uploadLine').hide()
                    $('#uploadFormat').hide()
                    $('#catdel-img').show()
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#imageSpace').removeClass('croppie-container').empty()
                    $('.CropperSection').addClass('hidden')
                    $('.mainsection').show()
                    $('#addCategory').modal('show')

                });
            }

        });

        $('#close,#crop-cancel').click(function () {
            $('canvas[class=cr-image]').css('opacity', '0')
            $("#changepicModal").modal('hide');
            $('#crop-container').removeClass('croppie-container').empty()
        });

        $('#changepicModal').on('hide.bs.modal', function (event) {
            $('canvas[class=cr-image]').css('opacity', '0')
            $('#crop-container').removeClass('croppie-container').empty()
        })
    }

    $(document).on('click', "#square", function () {
        $('#square').addClass('active')
        $('#portrait').removeClass('active')
        $('#Landscape').removeClass('active')
        $('.cr-viewport').css("width", "auto");
        $('.cr-viewport').css("height", "100%");
        $('.cr-viewport').css("aspect-ratio", "1/1");

    })

    $(document).on('click', "#portrait", function () {
        $('#square').removeClass('active')
        $('#portrait').addClass('active')
        $('#Landscape').removeClass('active')
        $('.cr-viewport').css("width", "auto");
        $('.cr-viewport').css("height", "100%");
        $('.cr-viewport').css("aspect-ratio", "3/4");
    })

    $(document).on('click', "#Landscape", function () {
        $('#square').removeClass('active')
        $('#portrait').removeClass('active')
        $('#Landscape').addClass('active')
        $('.cr-viewport').css("width", "auto");
        $('.cr-viewport').css("height", "75%");
        $('.cr-viewport').css("aspect-ratio", "16/9");
    })

})

$(document).on('click', "#crop-cancel", function () {
    $('.CropperSection').addClass('hidden')
    $('#addCategory').modal('show')
    $('#modalId2').modal('show')
    $('#userModal').modal('show')
    $('#LanguageModal').modal('show')
    $('.mainsection').show()
    $('body').css("overflow", "hidden")
})


//alert messages
document.addEventListener("DOMContentLoaded", async function () {

    var languagepath = $('.language-group>button').attr('data-path')
    await $.getJSON(languagepath, function (data) {
        languagedata = data
    })

    var Cookie = getCookie("get-toast");
    var content = Cookie.replaceAll("+", " ")
    var msg = getCookie("Alert-msg");
    var alertMsg = msg.replaceAll("+", " ")

    if (Cookie != '') {
        for (let key in languagedata.Toast) {
            // console.log("langtoast", key);

            if (key == content) {

                // **Important -- to get the toast msg please add a class .header-rht in your header or things that are in right of the header
                notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast[key] + `</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds


            }

        }
    } else if (alertMsg != '') {

        console.log("langudata",languagedata.Toast)

        for (let key in languagedata.Toast) {
            console.log("checkcondition",key,msg)
            if (key == msg) {

               
                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + languagedata.Toast[key] + `</p></div></div> </li></ul>`;


                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds


            }

        }

    }


    var Cookie1 = getCookie("log-toast");
    var content1 = Cookie1.replaceAll("+", " ")

    delete_cookie("get-toast");
    delete_cookie("Alert-msg");

    $(document).on('click', '#cancel-notify', function () {
        $(".toast-msg").remove();
    })
})


function getCookie(cname) {
    let name = cname + "=";
    let ca = document.cookie.split(';');

    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function setCookie(cname, cvalue, exdays) {
    const d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function delete_cookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

/*CropCancel */
$(document).on('click', '#cropcancel', function () {
    $("#CropModal .close").click();
})

/**new date picker plugin */
$(document).ready(function () {

    var fendDate = $('#date-end').val();
    var ffromDate = $('#date-start').val();


    $('#date-end').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', maxDate: new Date() }).on('change', function (e, date) {
        $('#date-start').bootstrapMaterialDatePicker('setMaxDate', date);
    });
    $('#date-start').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', maxDate: new Date() }).on('change', function (e, date) {
        $('#date-end').bootstrapMaterialDatePicker('setMinDate', date);
    });
    if (fendDate !== '') {
        $('#date-start').bootstrapMaterialDatePicker('setMaxDate', moment(fendDate, "DD MMM YYYY"));
    }
    if (ffromDate !== '') {
        $('#date-end').bootstrapMaterialDatePicker('setMinDate', moment(ffromDate, "DD MMM YYYY"));
    }
    var fSendDate = $('#date-to').val();
    var fSfromDate = $('#date-from').val();


    $('#date-to').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', minDate: moment() }).on('change', function (e, date) {
        $('#date-from').bootstrapMaterialDatePicker('setMaxDate', date);
    });
    $('#date-from').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', minDate: moment() }).on('change', function (e, date) {
        $('#date-to').bootstrapMaterialDatePicker('setMinDate', date);
    });
    if (fSendDate !== '') {
        $('#date-from').bootstrapMaterialDatePicker('setMaxDate', moment(fSendDate, "DD MMM YYYY"));
    }
    if (fSfromDate !== '') {
        $('#date-to').bootstrapMaterialDatePicker('setMinDate', moment(fSfromDate, "DD MMM YYYY"));
    }
    var fendDateclear = $('#date-end-clear').val();
    var ffromDateclear = $('#date-start-clear').val();


    $('#date-end-clear').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', maxDate: moment(), clearButton: true }).on('change', function (e, date) {

        console.log('calendar event:', e, date)
        $('#date-start-clear').bootstrapMaterialDatePicker('setMaxDate', date);
    });
    $('#date-start-clear').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY', maxDate: moment(), clearButton: true }).on('change', function (e, date) {
        console.log('calendar event:', e, date)
        $('#date-end-clear').bootstrapMaterialDatePicker('setMinDate', date);
    });
    if (fendDateclear !== '') {
        $('#date-start-clear').bootstrapMaterialDatePicker('setMaxDate', moment(fendDateclear, "DD MMM YYYY"));
    }
    if (ffromDateclear !== '') {
        $('#date-end-clear').bootstrapMaterialDatePicker('setMinDate', moment(ffromDateclear, "DD MMM YYYY"));
    }
    var startDate = $('#start-date').val();
    var endDate = $('#end-date').val();


    $('#end-date').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY' }).on('change', function (e, date) {
        $('#start-date').bootstrapMaterialDatePicker('setMaxDate', date);
    });
    $('#start-date').bootstrapMaterialDatePicker({ weekStart: 0, time: false, format: 'DD MMM YYYY' }).on('change', function (e, date) {
        $('#end-date').bootstrapMaterialDatePicker('setMinDate', date);
    });
    if (endDate !== '') {
        $('#start-date').bootstrapMaterialDatePicker('setMaxDate', moment(endDate, "DD MMM YYYY"));
    }
    if (startDate !== '') {
        $('#end-date').bootstrapMaterialDatePicker('setMinDate', moment(startDate, "DD MMM YYYY"));
    }
});


/**Click Expand btn */
$(document).on('click', '.expandbtn', function () {

    $('body').toggleClass('expand')

    if ($('body').hasClass('expand')) {
        $('.sidebar').addClass('open')
        $(this).html(`
    
          <p>v1.0</p> <img src="/public/img/right-arrow1.svg" class="right" />
        
        `)
        setCookie("mainmenu", "true")

    } else {
        $(this).html(` <img src="/public/img/right-arrow1.svg" alt=""><p>v1.0</p>`)
        delete_cookie("mainmenu")
        $('.sidebar').removeClass('open')
    }

})

// let localexpandFlg = localStorage.getItem("mainmenu");

// let sessionexpandFlg = sessionStorage.getItem("mainmenu")

// console.log("localexpandFlg",localexpandFlg ,sessionexpandFlg);

// console.log("localexpandFlg",localexpandFlg );


// if (localexpandFlg != ""){

//     $('body').toggleClass('expand')
//     if ($('body').hasClass('expand')) {
//         console.log("if");
//         $(".expandbtn").html(`<a href="javascript:void(0)"><img src="/public/img/right-arrow.svg" alt=""></a>
//         <a href="javascript:void(0)"><img src="/public/img/left-arrow.svg" alt=""> `+languagedata?.Toast?.close + ` </a>`)

//     }else{
//         $('body').removeClass('expand')
//     }
// }


/**Header class add */
$(document).ready(function () {

    if ($('div').hasClass('role-side-bar')) {
        $('header').addClass('roles-head')
    }

    /**Pagination disable */
    if ($('a').has('disabled')) {
        $('.disabled').attr("href", 'javascript:void(0)');
    }

})

$.validator.addMethod("space", function (value) {
    return /^[^-\s]/.test(value);
});


// /**category search */
$(".parentwithchild").on("keyup", function () {

    var value = $(this).val().toLowerCase();


    $(".forsearch").filter(function () {

        $(this).parents('.categorypdiv').toggle($(this).text().toLowerCase().indexOf(value) > -1)

    });

    var count = $('.categorypdiv:visible').length;

    if ($(".categorypdiv:visible").length == 0) {

        $("#categorynodatafound").show();

        $('#avacatcount').text("0");

    }
    if (value == "" || $(".categorypdiv:visible").length != 0) {

        $("#categorynodatafound").hide();

        $('#avacatcount').text(count);

    }

});

$(document).on('keyup', '#selectcategory', function () {

    var value = $(this).val().toLowerCase();

    $(".forsearches").filter(function () {

        $(this).parents('.selectedcategorydiv').toggle($(this).text().toLowerCase().indexOf(value) > -1)

    });

    var count = $('.selectedcategorydiv:visible').length;

    if ($(".selectedcategorydiv:visible").length == 0) {

        $("#selcategorynodatafound").show();

        $('#selcatcount').text("0");

    }

    if (value == "" || $(".selectedcategorydiv:visible").length != 0) {

        $("#selcategorynodatafound").hide();

        $('#selcatcount').text(count);

    }

})

// // /*Available Categories and Selected Categories */
var SelectedCategoryId = []

let SelectedCategoryValue = []

var selectedcatcount = 0

/*Selected */
$(document).on('click', '.category-select-btn', function () {

    var categoriesid = $(this).attr('data-categoryid')

    var id = $(this).attr('data-id')

    var pstr = ``

    $(this).siblings('.cal-list').children('.category-list').each(function () {

        if ($(this).hasClass('categoryname')) {

            pstr += `<p class="text-gray-900 text-xs font-normal mb-0 categoryname category-list" data-id=` + $(this).attr('data-id') + `>` + $(this).text() + `</p>`;
        } else {
            pstr += `<img src="/public/img/crumps.svg" class="category-list"/>`;
        }


    })

    var div =
        `<div class="flex justify-between pa-x-4 pa-y-3 items-center border-b border-[#EDEDED] selectedcategorydiv" id="selcategory-` + categoriesid + `">
        <div class="flex items-center space-x-[6px] cal-list">
            `+ pstr + `
        </div>
        <button class="w-4 h-4 flex items-center justify-center add-hover rounded-sm category-unselect-btn" data-id=`+ id + ` data-categoryid=` + categoriesid + `>
            <img src="/public/img/minus.svg" alt="" />
        </button>
        <p class="forsearches" style="display: none;">`+ $(this).siblings('.cal-list')[0].textContent + `</p>
    </div>
    `
    $("#no-data").hide()

    $('.selected_category').append(div);

    $(this).parent('.categorypdiv').remove();

    var idstr = [];

    $(this).siblings('.cal-list').children('.categoryname').each(function () {

        var id = $(this).attr('data-id');

        idstr.push(id);

    })

    var idstring = idstr.join(",");

    SelectedCategoryValue.push(idstring);

})


/*UnSelected */
$(document).on('click', '.category-unselect-btn', function () {


    $(this).parents('.selectedcategorydiv').remove();

    var categoriesid = $(this).attr('data-categoryid')

    var id = $(this).attr('data-id')

    var pstr = ``

    $(this).siblings('.cal-list').children('.category-list').each(function () {

        if ($(this).hasClass('categoryname')) {

            pstr += `<p class="text-gray-900 text-xs font-normal mb-0 categoryname category-list" data-id=` + $(this).attr('data-id') + `>` + $(this).text() + `</p>`;
        } else {
            pstr += `<img src="/public/img/crumps.svg" class="category-list"/>`;
        }


    })

    var div =
        `    <div class="flex justify-between pa-x-4 pa-y-3 items-center border-b border-[#EDEDED] categorypdiv" id="category-` + categoriesid + `">
        <div class="flex items-center space-x-[6px] cal-list" data-id={{$index}}>
           
        `+ pstr + `
              
            
               
        </div>
        <button class="w-4 h-4 flex items-center justify-center  add-hover rounded-sm category-select-btn" data-id="`+ id + `" data-categoryid="` + categoriesid + `">
            <img src="/public/img/add.svg" alt="" />
        </button>
        <p class="forsearch" style="display: none;">`+ $(this).siblings('.cal-list')[0].textContent + `</p>
    </div>
        
      `

    $('.categories-list').append(div);



    var idstr = [];

    $(this).siblings('.cal-list').children('.categoryname').each(function () {

        var id = $(this).attr('data-id');

        idstr.push(id);

    })


    var idstring = idstr.join(",");

    const index = SelectedCategoryValue.findIndex(function (element) {

        return element == idstring;

    });

    if (index >= 0) {

        SelectedCategoryValue.splice(index, 1);

    }
})

$(document).ready(function () {

    $('.forsearches').each(function () {

        var TEXT = $(this).text();

        $('.forsearch').each(function () {

            if (TEXT == $(this).text()) {

                $(this).parents('.categorypdiv').remove();
            }
        })

    })

})



// /**DropDown Select */
// $(document).on('click', '.dropdown-item', function () {

//     $(this).parent('.dropdown-menu').siblings('a').text($(this).text())

//     // $(this).parent('.dropdown-menu').siblings('a').html($(this).html(text))

//     $(this).parent('.dropdown-menu').siblings('input[type="hidden"]').val($(this).attr('data-id'))

// })


$(document).on('click', '.lang', function () {

    var languageId = $(this).attr('data-id')

    document.cookie = 'lang=' + languageId + ";path=/"

    window.location.reload()

})

/** Img div click */
$(document).on('click', '.profile-crop', function () {

    $('#myfile').click();

})

$(document).on('click', '.companylogo-crop', function () {

    $('#cmpymyfile').click();

})
/**Search Slide */
$(document).on('click', '.transitionSearch', function () {

    $(this).addClass('active');

})


// // Paganation Functionality
// function getPageList(totalPages, page, maxLength) {
//     function range(start, end) {
//         return Array.from(Array(end - start + 1), (_, i) => i + start);
//     }

//     var sideWidth = maxLength < 9 ? 1 : 2;
//     var leftWidth = (maxLength - sideWidth * 2 - 3) >> 1;
//     var rightWidth = (maxLength - sideWidth * 2 - 3) >> 1;


//     if (totalPages <= maxLength) {
//         return range(1, totalPages);
//     }


//     if (page <= maxLength - sideWidth - 1 - rightWidth) {
//         return range(1, maxLength - sideWidth - 1).concat(0, range(totalPages - sideWidth + 1, totalPages));
//     }


//     if (page >= totalPages - sideWidth - 1 - rightWidth) {
//         return range(1, sideWidth).concat(0, range(totalPages - sideWidth - 1 - rightWidth - leftWidth, totalPages));
//     }


//     return range(1, sideWidth).concat(0, range(page - leftWidth, page + rightWidth), 0, range(totalPages - sideWidth + 1, totalPages));

// }

// $(function() {
//     var numberOfItems = $("#count").val();
//     var limitPerPage = $("#limit").val(); //No. of cards to show per page
//     var totalPages = Math.ceil(numberOfItems / limitPerPage);
//     var paginationSize = 5; //pagination items to show
//     var currentPage;

//     function showPage(whichPage) {

//         if (whichPage < 1 || whichPage > totalPages) return false;

//         currentPage = whichPage;

//         $("tbody tr").hide().slice((currentPage - 1) * limitPerPage, currentPage * limitPerPage).show();

//         $(".pagination li").slice(1, -1).remove();

//         getPageList(totalPages, currentPage, paginationSize).forEach(item => {

//             $("<li>").addClass("page-item").addClass(item ? "current-page" : "dots").toggleClass("active", item === currentPage).append($("<a>").addClass("page-link").attr({ href: "javascript:void(0)" }).text(item || '...')).insertBefore(".next-page");
//         });

//         $(".previous-page").toggleClass("disable", currentPage === 1);
//         $(".next-page").toggleClass("disable", currentPage === totalPages);
//         return true;
//     }


//     $(".pagination").append(
//         $("<li>").addClass("page-item").addClass("previous-page").append($("<img>").addClass("page-link").attr("src","/public/img/carat-right.svg")),
//         $("<li>").addClass("page-item").addClass("next-page").append($("<img>").addClass("page-link").attr("src","/public/img/carat-right.svg"))
//     );


//     $("tbody").show();
//     showPage(1);


//     $(document).on("click", ".pagination li.current-page:not(.active)", function() {
//         return showPage(+$(this).text());
//     });

//     $(".next-page").on("click", function() {
//         return showPage(currentPage + 1);
//     });


//     $(".previous-page").on("click", function() {
//         return showPage(currentPage - 1);
//     });

// });



function updateVisiblePages() {

    // How many pages can be around the current page
    let maxAround = 1;
    // How many pages are displayed if in beginning or end range
    let maxRange = 6;
    // The current page (pass this however you like)
    let currentPage = parseInt($('.current').text().trim());

    // The paginator (would suggest using an ID instead)
    let $paginator = $('.pogination');
    // Count amount of pages
    let totalPages = $paginator.children('ul').children('li').length;

    let endRange = totalPages - maxRange;
    // Check if we're in the starting section
    let inStartRange = currentPage <= maxRange;
    // Check if we're in the ending section
    let inEndRange = currentPage >= endRange;

    // We need this for the span(s)
    let lastDotIndex = -1;

    // Remove all dots
    $('span.dots').remove();
    // Loop the pages
    $paginator.children('ul').children('li').children('a').each(function (page, element) {

        // Index starts at 0, pages at 1
        ++page;
        let $element = $(element);
        if (page === 1 || page === totalPages || page === 2 || page === totalPages - 1 || page === totalPages - 2) {
            // Always show first and last
            $element.show();
        } else if (inStartRange && page <= maxRange) {
            // Show element if in start range
            $element.show();
        } else if (inEndRange && page >= endRange) {

            // Show the element if in ending range
            $element.show();
        } else if (page === currentPage - maxAround || page === currentPage || page === currentPage + maxAround || page === currentPage + 2 || page === currentPage + 3) {
            // Show element if in the wrap around
            $element.show();
        } else {
            // Doesn't validate, hide it.
            // Append dot if needed
            $element.hide();

            if (lastDotIndex === -1 || (!inStartRange && !inEndRange && $('span.dots').length < 2 && page > currentPage)) {

                lastDotIndex = page;
                // Insert dots after this page, we only have one or 2, and we can only insert the second one
                // if we're not in the start or end range, and it's past the current page
                $element.after('<span class="dots">...</span>');
            }
        }

    });
}

/**Form design */
// let inputGroups = document.querySelectorAll('.input-group');
// inputGroups.forEach(inputGroup => {

//     let inputField = inputGroup.querySelector('input');

//     inputField.addEventListener('focus', function (event) {
//         if (event.target.id !== 'searchcatlist') {
//             inputGroup.classList.add('input-group-focused');

//         }
//     });
//     inputField.addEventListener('blur', function () {
//         inputGroup.classList.remove('input-group-focused');
//     });

// });

$(document).ready(function () {
    // window.addEventListener('beforeunload', (event) => {
    //     $.ajax({
    //         url: "/lastactive"
    //     })
    // });
});

// logout
// $("#logout").click(function(){

//     sessionStorage.removeItem("rememberme");
//     localStorage.removeItem("rememberme");
// })

// window.onload = function () {
//     const sidebar = document.querySelector(".sidebar");
//     const closeBtn = document.querySelector("#btn");
//     const searchBtn = document.querySelector(".bx-search")

//     closeBtn.addEventListener("click", function () {
//         sidebar.classList.toggle("open")
//         menuBtnChange()
//     })

//     searchBtn.addEventListener("click", function () {
//         sidebar.classList.toggle("open")
//         menuBtnChange()
//     })

//     function menuBtnChange() {
//         if (sidebar.classList.contains("open")) {
//             closeBtn.classList.replace("right", "left")
//         } else {
//             closeBtn.classList.replace("left", "right")
//         }
//     }
// }

// inputGroups.forEach(inputGroup => {

//     let inputField = inputGroup.querySelector('#select_channel');

//     inputField.addEventListener('focus', function (event) {
//         if (event.target.id !== 'searchcatlist') {
//             inputGroup.classList.add('input-group-focused');

//         }
//     });
//     inputField.addEventListener('blur', function () {
//         inputGroup.classList.remove('input-group-focused');
//     });

// });


//settings menu active menu highlighting code
$(document).ready(function () {

    if (window.location.href.indexOf('myprofile') != -1) {
        $('#myProfileLink').addClass('active')

    } else if (window.location.href.indexOf('changepassword') != -1) {
        $('#changePasswordLink').addClass('active')

    } else if (window.location.href.indexOf('roles') != -1) {
        $('#rolesLink').addClass('active')

    } else if (window.location.href.indexOf('users') != -1) {
        $('#usersLink').addClass('active')

    } else if (window.location.href.indexOf('general-settings') != -1) {
        $('#genSettingsLink').addClass('active')

    } else if (window.location.href.indexOf('languages') != -1) {
        $('#langPageLink').addClass('active')

    } else if (window.location.href.indexOf('emails') != -1) {
        $('#emailPageLink').addClass('active')

    } else if (window.location.href.indexOf('data') != -1) {
        $('#dataPageLink').addClass('active')

    }


})