let searchParams = new URLSearchParams(window.location.search)
var languagedata

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

$(document).keydown(function(event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
});

$(document).ready(function () {

    if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 ||window.location.href.indexOf("customers") != -1 ||window.location.href.indexOf("applicants") != -1) {

        var imgurl, img, canvas, ctx;
        let cropper;

        // Image Value Empty
        $("#myfile,#file,#cmpymyfile").click(function () {
            $(this).val("")
        })

        var viewportType = ""

        // Change Image with cropper
        $("#myfile,#file,#cmpymyfile").change(function () {

            $('#crop-container').removeClass('croppie-container').empty().show()

            if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 ||window.location.href.indexOf("customers") != -1 ||window.location.href.indexOf("applicants") != -1) {


                if (window.location.href.indexOf("myprofile") != -1) {

                    $("#logo-input").val("2")

                    $("#prof-crop").val("1")

                    $('#changepicModal .admin-header >h3').text('Crop Myprofile Image')

                } else if (window.location.href.indexOf("users") != -1) {

                    $('#changepicModal .admin-header >h3').text('Crop User Image')

                } else if (window.location.href.indexOf("member") != -1) {

                    $('#changepicModal .admin-header >h3').text('Crop Member Image')

                }else if(window.location.href.indexOf("customers") !=-1){

                    $('#changepicModal .admin-header >h3').text('Crop Customer Image')

                }else if(window.location.href.indexOf("applicants") !=-1){

                    $('#changepicModal .admin-header >h3').text('Crop Applicant Image')

                }

                $("#myfile-error").hide()
            }

            if($(this).attr("id")=='cmpymyfile'){

                viewportType = 'square'

            }else{

                viewportType = 'circle'
            }
            var file = this.files[0];
            var filename = $(this).val();
            var ext = filename.split(".").pop().toLowerCase();
            if (($.inArray(ext, ["jpg", "png", "jpeg","svg"]) != -1)) {
                var reader = new FileReader();
                reader.onload = function (event) {
                    imgurl = event.target.result
                    img = new Image()
                    img.onload = function () {
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
                        });
                        $('canvas[class=cr-image]').css('opacity', '0')
                        if ($('.cr-slider-wrap').length > 1) {
                            $('.cr-slider-wrap')[1].remove()
                        }
                        $('.cr-slider-wrap').appendTo($('#zoom-div'));
                        $('#rotateSlider').val("0")
                        $('.cr-slider').css('display', 'block')
                    }

                    img.src = imgurl
                }
                reader.readAsDataURL(file);
                $('#changepicModal').modal('show');
            } else {

                if (window.location.href.indexOf("myprofile") != -1 || window.location.href.indexOf("users") != -1 || window.location.href.indexOf("member") != -1 ||window.location.href.indexOf("customers") != -1 ||window.location.href.indexOf("applicants") != -1) {
                    $("#myfile-error").text(languagedata?.Toast?.errmsgupload).show()
                }
            }
        });

        // $("#cmpymyfile").change(function () {

        //     $('#crop-container').removeClass('croppie-container').empty().show()

        //     if(window.location.href.indexOf("myprofile")!=-1 || window.location.href.indexOf("users")!=-1 || window.location.href.indexOf("member")!=-1){

        //         if(window.location.href.indexOf("member")!=-1){

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
        //                     cropper = $('#crop-container').croppie({
        //                         enableExif: true,
        //                         enableResize: false,
        //                         enableOrientation: true,
        //                         viewport: {
        //                             width: 300,
        //                             height: 300,
        //                             type: 'circle'
        //                         },
        //                         showZoomer: false,
        //                     });
        //                 $('canvas[class=cr-image]').css('opacity', '0')
        //                 if($('.cr-slider-wrap').length>1){
        //                    $('.cr-slider-wrap')[1].remove()
        //                 }
        //                 $('.cr-slider-wrap').appendTo($('#zoom-div'));
        //                 $('#rotateSlider').val("0")
        //                 $('.cr-slider').css('display', 'block')
        //             }

        //             img.src = imgurl
        //         }
        //         reader.readAsDataURL(file);
        //         $('#changepicModal').modal('show');
        //     }else{

        //         if(window.location.href.indexOf("myprofile")!=-1 || window.location.href.indexOf("users")!=-1 || window.location.href.indexOf("member")!=-1){
        //             $("#cmpymyfile-error").text( languagedata?.Toast?.errmsgupload).show()
        //         }
        //     }
        // });

        $('#changepicModal').on('shown.bs.modal', function (event) {
            cropper.croppie('bind', {
                url: imgurl
            }).then(function () {
                cropper.croppie('setZoom', 0)
                console.log("cropper initialized!");
            })
            canvas = document.querySelector('canvas[class=cr-image]');
            ctx = canvas.getContext('2d');
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

            var in_val = $("#prof-crop").val()

            if (in_val == 1) {
                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'jpeg', circle: true, }).then(function (dataUrl) {

                    $('#profpic').attr('src', dataUrl)
                    $("#name-string").hide()
                    $('input[name=crop_data]').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#crop-container').removeClass('croppie-container').empty()
                });
            }
            if(in_val == 2){

                cropper.croppie('result', { type: 'base64', size: 'original', quality: 0.5, format: 'png',}).then(function (dataUrl) {
                    $('#profpic2').attr('src', dataUrl)
                    $("#cname-string").hide()
                    $('input[name=crop_data2]').attr('value', dataUrl)
                    $('canvas[class=cr-image]').css('opacity', '0')
                    $("#changepicModal").modal('hide');
                    $('#crop-container').removeClass('croppie-container').empty()
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

    if (Cookie != '') {
        for (let key in languagedata.Toast) {
            if (key == content) {

                notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata.Toast[key] + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds


            }

        }
    } else if (msg != '') {

        for (let key in languagedata.Toast) {
            if (key == msg) {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Toast[key] + '</span></div>';
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
        setCookie("mainmenu","true")
   
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

    } else {

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

    } else {

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

    $(this).siblings('.choose-cat-list-col').children('.para').each(function () {

        if ($(this).hasClass('categoryname')) {

            pstr += `<h3 class="para categoryname"style="font-weight: 400;" data-id=` + $(this).attr('data-id') + `>` + $(this).text() + `</h3>`;
        } else {
            pstr += `<h3 class="para">` + $(this).text() + `</h3>`;
        }


    })

    var div =
        `<div class="categories-list-child selectedcategorydiv" id="selcategory-` + categoriesid + `">
        <div class="choose-cat-list-col" style="display: flex;">
            `+ pstr + `
        </div>
        <a href="javascript:void(0)" class="category-unselect-btn" data-id=`+ id + ` data-categoryid=` + categoriesid + `>
            <img src="/public/img/bin.svg" alt="" />
        </a>
        <p class="forsearches" style="display: none;">`+ $(this).siblings('.choose-cat-list-col')[0].textContent + `</p>
    </div>
    <div class="noData-foundWrapper" id="selcategorynodatafound" style="display: none;">
    
    <div class="empty-folder">
        <img src="/public/img/folder-sh.svg" alt="">
        <img src="/public/img/shadow.svg" alt="">
    </div>
    <h1 class="heading">Oops No Data Found</h1>
    </div>  
    `

    $('.selected_category').append(div);

    $(this).parents('.categorypdiv').remove();

    var idstr = [];

    $(this).siblings('.choose-cat-list-col').children('.categoryname').each(function () {

        var id = $(this).attr('data-id');

        idstr.push(id);

    })

    var idstring = idstr.join(",");



    SelectedCategoryValue.push(idstring);

    selectedcatcount = selectedcatcount + 1

    selcount = parseInt($('#selcatcount').text()) + 1

    $('#selcatcount').text(selcount);

    if ($('#caterror').is(':visible') && $('#selcatcount').text(selectedcatcount) != "0") {

        $('#caterror').removeClass('ig-error').css('visibility', 'hidden')

    }

    $('#avacatcount').text(parseInt($('#avacatcount').text()) - 1);
})


/*UnSelected */
$(document).on('click', '.category-unselect-btn', function () {

    // $('#category-' + $(this).attr('data-categoryid')).show()

    $(this).parents('.selectedcategorydiv').remove();

    var categoriesid = $(this).attr('data-categoryid')

    var id = $(this).attr('data-id')

    var pstr = ``

    $(this).siblings('.choose-cat-list-col').children('.para').each(function () {

        if ($(this).hasClass('categoryname')) {

            pstr += `<h3 class="para categoryname" style="font-weight: 400;" data-id=` + $(this).attr('data-id') + `>` + $(this).text() + `</h3>`;
        } else {
            pstr += `<h3 class="para">` + $(this).text() + `</h3>`;
        }


    })

    var div =
        `    <div class="categories-list-child categorypdiv" id="category-` + categoriesid + `">
        <div class="choose-cat-list-col" style="display: flex;" data-id={{$index}}>
           
        `+ pstr + `
              
            
               
        </div>
        <a href="javascript:void(0)" class="category-select-btn" data-id="`+ id + `" data-categoryid="` + categoriesid + `">
            <img src="/public/img/add-category.svg" alt="" />
        </a>
        <p class="forsearch" style="display: none;">`+ $(this).siblings('.choose-cat-list-col')[0].textContent + `</p>
    </div>
        
      `

    $('.slist').append(div);



    var idstr = [];

    $(this).siblings('.choose-cat-list-col').children('.categoryname').each(function () {

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

    selectedcatcount = selectedcatcount - 1

    selcount = parseInt($('#selcatcount').text()) - 1

    $('#selcatcount').text(selcount);

    if ($('#caterror').is(':visible') && $('#selcatcount').text() == "0") {

        $('#caterror').addClass('ig-error').css('visibility', 'visible')

    }

    $('#avacatcount').text(parseInt($('#avacatcount').text()) + 1);
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



/**DropDown Select */
$(document).on('click', '.dropdown-item', function () {

    $(this).parent('.dropdown-menu').siblings('a').text($(this).text())

    // $(this).parent('.dropdown-menu').siblings('a').html($(this).html(text))

    $(this).parent('.dropdown-menu').siblings('input[type="hidden"]').val($(this).attr('data-id'))

})


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
let inputGroups = document.querySelectorAll('.input-group');
inputGroups.forEach(inputGroup => {

    let inputField = inputGroup.querySelector('input');

    inputField.addEventListener('focus', function (event) {
        if (event.target.id !== 'searchcatlist') {
            inputGroup.classList.add('input-group-focused');

        }
    });
    inputField.addEventListener('blur', function () {
        inputGroup.classList.remove('input-group-focused');
    });

});

$(document).ready(function () {
    window.addEventListener('beforeunload', (event) => {
        $.ajax({
            url: "/lastactive"
        })
    });
});

// logout
$("#logout").click(function(){
   
    sessionStorage.removeItem("rememberme");
    localStorage.removeItem("rememberme");
})

window.onload = function () {
    const sidebar = document.querySelector(".sidebar");
    const closeBtn = document.querySelector("#btn");
    const searchBtn = document.querySelector(".bx-search")

    closeBtn.addEventListener("click", function () {
        sidebar.classList.toggle("open")
        menuBtnChange()
    })

    searchBtn.addEventListener("click", function () {
        sidebar.classList.toggle("open")
        menuBtnChange()
    })

    function menuBtnChange() {
        if (sidebar.classList.contains("open")) {
            closeBtn.classList.replace("right", "left")
        } else {
            closeBtn.classList.replace("left", "right")
        }
    }
}

inputGroups.forEach(inputGroup => {

    let inputField = inputGroup.querySelector('#select_channel');

    inputField.addEventListener('focus', function (event) {
        if (event.target.id !== 'searchcatlist') {
            inputGroup.classList.add('input-group-focused');

        }
    });
    inputField.addEventListener('blur', function () {
        inputGroup.classList.remove('input-group-focused');
    });

});

