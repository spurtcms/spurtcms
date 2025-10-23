var languagedata
var finaltags = []
var finaldeletetags = []
var tags = [];
var joindata
var channelnamearr = []
var ChannelIdarr = []
// loading language data
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    if (localStorage.getItem('showToast') === 'removetrue') {

        notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ` + languagedata.Blocks.removesuccessfully + `</p ></div ></div ></li></ul> `;
        $(notify_content).insertBefore(".header-rht");

        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();

            });
        }, 5000);

        localStorage.removeItem('showToast');
    }
    if (localStorage.getItem('showToast') === 'addtrue') {

        notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ` + languagedata.Blocks.collectionsuccessfully + `</p ></div ></div ></li></ul> `;
        $(notify_content).insertBefore(".header-rht");

        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();

            });
        }, 5000);

        localStorage.removeItem('showToast');
    }



})





$('.hd-crd-btn').click(function () {

    if ($('#hd-crd').is(':visible')) {

        $('#hd-crd').addClass('hidden').removeClass("show");
        document.cookie = `blockbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `blockbanner=true; path=/;`;
    }
});


// Media model open

// $(document).on("click", "#blockbrowse", function () {
//     $("#mediaModal").addClass('show')
// })


// image upload in block

$(document).on("click", "#selectimg", function () {
    $(".selected-numbers").hide()
    src = $(this).attr("src")
    $("#coverimg").val(src)
    names = $(this).attr("data-id")
    if (src != "") {
        $("#coverimg-error").css("display", "none")
    }

    if (src != "") {

        html = `  <div class="border border-[#ECECEC] p-[8px] px-[12px] flex items-center gap-[12px] rounded-[5px]">
             <span class="w-[24px] min-w-[24px] h-[24px]"><img id="uploadimg" src="` + src + ` " alt=""></span>
             <p class="text-[14px] font-normal leading-[17.5px] text-[#252525] text-left" id="filename" > `+ names + ` </p>
             <a href="javascript:void(0);" class="min-w-[16px] ml-auto" id="removeimg">
                 <img src="/public/img/removeCover.svg" alt="remove">
             </a>
         </div> `

        $(html).insertAfter(".imagediv");
    }

    $("#mediaModal").hide()
    $("#createModal").addClass('show')
})

//  remove image
$(document).on("click", "#removeimg", function () {
    $(this).parent("div").remove()
    $("#coverimg").val("")
    $(this).attr("src", "")
    $("#browse").removeAttr("style")
    $('#browse').removeClass('hidden')
    $("#blockchosimg").removeClass('hidden')
})


// save Form vaildation
$("#createblock").on("click", function () {

    jQuery.validator.addMethod("duplicatetitle", function (value) {

        var result;
        id = $("#blockid").val()
        $.ajax({
            url: "/admin/blocks/checktitle",
            type: "POST",
            async: false,
            data: { "block_title": value, "block_id": id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim() != "true"
    })

    $("#blockform").validate({
        ignore: [],
        rules: {
            title: {
                required: true,
                space: true,
                duplicatetitle: true
            },
            html: {
                required: true,
                space: true,
            },
            css: {
                // required: true,
                // space: true,
            },
            chennelname: {
                required: true,
            },

            coverimg: {
                required: true,
            },
            channelids: {

            }
        },
        messages: {
            title: {
                required: "*" + languagedata.Blocks.titleerr,
                space: "* " + languagedata.spacergx,
                duplicatetitle: "*" + languagedata.Blocks.titleduplicateerr,
            },
            html: {
                required: "*" + languagedata.Blocks.htmlerr,
                space: "* " + languagedata.spacergx,
            },
            css: {
                // required: "* Please Enter Css " ,
                // space: "* " + languagedata.spacergx,
            },
            chennelname: {
                required: "*" + "Please select Channel",
            },
            coverimg: {
                required: "*" + languagedata.Blocks.imgerr,
            },
            channelids: {

            }
        },


    });

    $("#chennelname-error").css("margin-top", "6px");

    var formcheck = $("#blockform").valid();

    if (formcheck) {

        const form = document.getElementById('blockform');

        // Create a FormData object from the form
        const formData = new FormData(form);

        fetch('/admin/blocks/create', {
            method: 'POST',
            body: formData

        }).then(response => response.json())
            .then(data => {
                const editorElement = document.getElementById('spurt-editor');
                if (editorElement) {

                    editorElement.setAttribute('block', data.blocks);
                } else {
                    console.error('Element not found: spurt-editor');
                }
            })
        $('#createModal').modal('hide');
        const basicBlockModalElement = document.getElementById('basicBlock');
        if (basicBlockModalElement) {
            const basicBlockModal = new window.bootstrap.Modal(basicBlockModalElement);
            basicBlockModal.show();
        }


    }
    $('#categoryimage-error').addClass("hidden!")
    $('#categoryimage-error').css('display', 'none')
    return false

})

// Update Form vaildation
$("#updateblock").on("click", function () {



    jQuery.validator.addMethod("duplicatetitle", function (value) {

        var result;
        id = $("#blockid").val()

        $.ajax({
            url: "/admin/blocks/checktitle",
            type: "POST",
            async: false,
            data: { "block_title": value, "block_id": id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim() != "true"
    })
    $("#tagvalues-error").css("display", "none")
    $("#blockform").validate({
        ignore: [],
        rules: {
            title: {
                required: true,
                space: true,
                duplicatetitle: true
            },
            html: {
                required: true,
                space: true,
            },
            css: {
                // required: true,
                // space: true,
            },
            // tag: {
            //     required: true,
            //     space: true,
            // },

            coverimg: {
                required: true,
            },
            chennelname: {
                required: true,
            },

            coverimg: {
                required: true,
            },
            channelids: {

            }
        },
        messages: {
            title: {
                required: "*" + languagedata.Blocks.titleerr,
                space: "* " + languagedata.spacergx,
                duplicatetitle: "*" + languagedata.Blocks.titleduplicateerr,
            },
            html: {
                required: "*" + languagedata.Blocks.htmlerr,
                space: "* " + languagedata.spacergx,
            },
            css: {
                // required: "* Please Enter Css " ,
                // space: "* " + languagedata.spacergx,
            },
            // tag: {
            //     required: "*" + languagedata.Blocks.tagerr,
            //     space: "* " + languagedata.  spacergx,
            // },
            coverimg: {
                required: "*" + languagedata.Blocks.imgerr,
            },
            chennelname: {
                required: "*" + "Please select Channel",
            },
            coverimg: {
                required: "*" + languagedata.Blocks.imgerr,
            },
            channelids: {

            }
        },

    });
    var formcheck = $("#blockform").valid();
    if (formcheck) {

        const form = document.getElementById('blockform');

        // Create a FormData object from the form
        const formData = new FormData(form);

        fetch('/admin/blocks/update', {
            method: 'POST',
            body: formData

        }).then(response => response.json())
            .then(data => {
                const editorElement = document.getElementById('spurt-editor');
                if (editorElement) {
                    editorElement.setAttribute('block', data.blocks);

                } else {
                    console.error('Element not found: spurt-editor');
                }

            })
        $('#createModal').modal('hide');

        const basicBlockModalElement = document.getElementById('basicBlock');
        if (basicBlockModalElement) {
            const basicBlockModal = new window.bootstrap.Modal(basicBlockModalElement);
            basicBlockModal.show();
        }

    }
    return false

})


// // tag creating

// document.addEventListener('keyup', function (event) {
//     var tagval = $("#tagvalues").val()
//     $("#tagvalues-error").hide()
//     if (event.key === "Enter") {
//         console.log("if");
//         tagname = $("#tag").val()
//         console.log(tagname, "tagname")
//         if (tagname != "") {
//             if (tags.indexOf(tagname) == -1) {
//                 tags.push(tagname)
//                 arr = tagname.split(",")

//                 console.log(tags, "arraydiv")

//                 // $(".tagdiv").addClass('show')
//                 $(".tagdiv").removeClass('hidden')

//                 if (tagval != "") {
//                     joindata = tagval + ',' + tags
//                 } else {
//                     joindata = tags
//                 }

//                 if (arr == "") {
//                     tagdesign = `
//                     <li class="bg-[#ECECEC] p-[6px_12px] flex items-center gap-3 w-fit tagdivli">
//                         <a href="javascript:void(0);" id="tagdeletebtn" data-id=`+ tagname + `>
//                             <img src="/public/img/pill-close.svg" alt="close">
//                         </a>
//                         <span class="text-[14px] font-[300] leading-[17.5px] tracking-[0.01em] text-[#252525]">`+ tagname + `</span>
//                     </li>

//                 `
//                     $(".taglist").append(tagdesign)


//                 } else {
//                     arr.forEach(element => {

//                         tagdesign = `
//                         <li class="bg-[#ECECEC] p-[6px_12px] flex items-center gap-3 w-fit tagdivli">
//                             <a href="javascript:void(0);" id="tagdeletebtn" data-id=`+ element + `>
//                                 <img src="/public/img/pill-close.svg" alt="close">
//                             </a>
//                             <span class="text-[14px] font-[300] leading-[17.5px] tracking-[0.01em] text-[#252525]">`+ element + `</span>
//                         </li>

//                     `
//                         $(".taglist").append(tagdesign)

//                         $("#tagvalues").val(element)
//                     });

//                 }

//                 $("#tag").val("")
//             }
//             $("#tagvalues").val(tags)
//         }
//     } else {

//         taglist = $("#tag").val()

//         $("#tagvalues").val(taglist)
//     }


// })

// Delete tags

$(document).on("click", "#tagdeletebtn", function () {

    deltagname = $(this).attr("data-id").trim()

    spanval = $(this).siblings('span').text().trim()

    console.log(spanval, "val")

    if (deltagname == spanval) {

        $(this).parents("li").remove()
    }


    finaldeletetags.push(deltagname)

    $("#deletetag").val(finaldeletetags)

    finaltags.forEach((link, index) => {

        if (link == deltagname) {

            finaltags.splice(index, 1)
        }

    })

    $("#tagvalues").val(finaltags)

    if ($(".tagdivli:visible").length == 0) {
        $(".tagdiv").addClass('hidden')
    }

    // if ($('#tagvalues').val() == "") {
    //     $(".tagdiv").addClass('hidden')
    //     $(".tagdiv").removeClass('show')

    // }


})



// add Collection

$(document).on("click", "#collectionblock", function () {

    blockid = $(this).attr("data-blockid")

    $.ajax({
        url: "/admin/blocks/createcollection",
        type: "POST",
        async: false,
        data: { "blockid": blockid, csrf: $("input[name='csrf']").val() },
        datatype: "json",
        caches: false,
        success: function (result) {


            if (result) {

                localStorage.setItem('showToast', 'addtrue');

                window.location.href = "/blocks";


            } else {
                message = "This Collection Already Exists"

                notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"> <li><div class="flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FFB74D] bg-[#FFF7EB]"><a href="javascript:void(0)" class="absolute right-[8px] top-[8px] " id="cancel-notify"><img src="/public/img/close-toast.svg" alt="close"></a><div><img src="/public/img/toast-wraning.svg" alt="toastwraning"></div><div><h3 class="text-[#FFB74D] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] "> ` + message + ` </p></div></div></li></ul>;`
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }


        }

    })
})

// preview design
$(document).on("click", "#blockpreview", function () {
    var blockimg = $(this).attr("data-blockimg")
    var blocktitle = $(this).attr("data-title")
    var username = $(this).attr("data-username")
    var userimg = $(this).attr("data-userimg")
    var namestring = $(this).attr("data-namestring")

    $("#previewtitle").text(blocktitle)
    $("#previewname").text(username)
    $("#previewblogimg").attr("src", blockimg)

    if (userimg != "") {
        $("#previewuserimg").attr("src", userimg)
        $(".previewusername").css("display", "none")
    } else {
        $("#previewnamestring").text(namestring)
        $(".previewimgblock").css("display", "none")
        $(".previewusername").removeAttr('style')
    }
})

// preview close empty the all text

$(document).on("click", "#previewclose", function () {
    $("#previewtitle").text("")
    $("#previewname").text("")
    $("#previewuserimg").attr("src", "")
    $("#previewblogimg").attr("src", "")
    $("#data-blockid").val("")
    $("#data-userid").val("")
    $("#previewnamestring").val("")
})

// Block delete functionality

// $(document).on("click", "#removeblock", function () {
//     id = $(this).data("id")
//     $('#content').text(languagedata.Blocks.delcontent);
//     $(".deltitle").text(languagedata.Blocks.deltitle)
//     blockname = $("#blocktitle").text()
//     $('.delname').text(blockname)
//     $('#delid').show();
//     $('#delcancel').text(languagedata.no);
//     $('#delid').attr('href', '/blocks/deleteblock?id=' + id);

// })



$(document).on("click", ".deleteblock", function () {
    id = $(this).attr("data-id")
  console.log(id,"dataid");
  
    fetch(`/admin/blocks/deleteblock?id=${id}`, {
        method: "GET"
    })
        .then(response => response.json())
        .then(data => {

            const editorElement = document.getElementById('spurt-editor');
            if (editorElement) {
                editorElement.setAttribute('block', data.blocks);


            } else {
                console.error('Element not found: spurt-editor');
            }
        })

})




// search functionality in block

$(document).on('keyup', '#blocksearch', function (event) {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/blocks";
        }
    }
});

// Default search list
$(document).on('keyup', '#blockdefaultsearch', function (event) {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/blocks/defaultlist";
        }
    }
});


// set as premium

function BlockPremium() {
    $('#premium').on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();

}

// block modal cancel button functionlaity
$(document).on("click", "#Closeblockmodal,#CreateBlock", function () {
    $(".coverimgblock").remove()
    $("#createModal").addClass('hidden')
    $("#createModal").removeClass('show')
    $("#backgroundfade").removeClass("offcanvas-backdrop fade show")
    $("#premium").val("")
    $("#premium").attr("checked", false)
    $("#title").val("")
    $("#html").val("")
    $("#css").val("")
    $("#tag").val("")
    $("#coverimg").attr("src", "")
    $("#uploadimg").attr("src", "")
    $("#coverimg-error").hide()
    $("#tagvalues-error").hide()
    $("#html-error").hide()
    $("#title-error").hide()
    $("#blockfilename").parent('').empty()
    $('body').css("overflow", "visible")
    $("#deletetag").val("")
    $("#showgroup").text("Choose Channel")
    $('#showgroup').removeClass('text-bold').addClass('text-bold-gray')
    $('#chennelname').val('');
    $('#chennelname-error').hide()
    $(".select-channel").prop("checked", false)
    $('#channelids').val('')
    channelnamearr = ['']
    ChannelIdarr = ['']

    $("#CreateBlock").attr('data-block', false)

})


// new block button functionality
$(document).on("click", "#CreateBlock", function () {
    $("#createModal").removeClass('hidden')
    $("#createModal").addClass('show')
    $("#blockform").attr("action", "/blocks/create")
    $("#createblock").show()
    $("#updateblock").css("display", "none")
    $("#titleblock").text("Create Block")
    $('body').css("overflow", "hidden")
    $(".tagdiv").addClass("hidden")
    $(".tagdiv").removeClass("show")
    $("#tagvalues").val("")
    $("#tagdeletebtn").attr("data-id", "")
    $(".taglist").html("")
    $("#browse").css("display", "")
    $("#blockchosimg").css("display", "")
    $("#premium").prop("checked", false)
    $("#deletetag").val("")
    $("#blockcoverimg").hide()
    $("#catdel-img").hide()
    $("#uploadLine").show()
    $("#uploadFormat").show()
    $("#coverimg").val("")
    $("#blockcoverimg").attr("src", "")
    $(this).attr('data-block', true)
    $("#showgroup").text("Choose Channel")
    $('#showgroup').removeClass('text-bold').addClass('text-bold-gray')
})



// $(document).on('click','.BlockStatus',function(){

//     var id=$(this).data('id')

//     $('#BlockStatus' + id).on('change', function () {
//         this.value = this.checked ? 1 : 0;
//     }).change();
//     var isactive = $(this).val();

//     $.ajax({
//         url: '/blocks/blockisactive',
//         type: 'POST',
//         async: false,
//         data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
//         dataType: 'json',
//         cache: false,
//         success: function (result) {

//             console.log(result);


//             if (result) {

//                 notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > Block Status Updated Successfully </p ></div ></div ></li></ul> `;
//                 $(notify_content).insertBefore(".header-rht");

//             } else {

//                 notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"> <li><div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"><a href="javascript:void(0)" class="absolute right-[8px] top-[8px] "><img src="/public/img/close-toast.svg" alt="close"></a><div><img src="/public/img/toast-error.svg" alt="toast error"></div><div><h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Error</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] "> ` + languagedata.Toast.internalserverr + ` </p></div></div></li><ul>`;
//                 $(notify_content).insertBefore(".header-rht");

//             }

//             setTimeout(function () {
//                 $('.toast-msg').fadeOut('slow', function () {
//                     $(this).remove();
//                 });
//             }, 5000); // 5000 milliseconds = 5 seconds

//         }
//     });
// })


// Edit functionality
$(document).on("click", "#editblock", function () {

    $("#tagvalues").val("")
    $('.coverimgblock').remove()
    blockid = $(this).attr("data-id")
    $('body').css("overflow", "hidden")
    $("#titleblock").text("Update Block")
    $("#updateblock").show()
    $("#createblock").hide()
    $(".taglist").html("")
    $('#updateblock').attr('data-block', true)

    $("#backgroundfade").addClass('offcanvas-backdrop fade show')


    $.ajax({
        url: "/admin/blocks/blockdetails",
        type: "GET",
        async: false,
        data: { "id": blockid },
        datatype: "json",
        caches: false,
        success: function (result) {

            if (result) {

                $("#blockid").val(result.Id)
                $("#title").val(result.Title)
                $("#html").val(result.BlockContent)
                $("#css").val(result.BlockCss)
                $("#isactive").val(result.IsActive)
                $("#tag").val(result.TagValue)
                $("#tagvalues").val(result.TagValue)
                $("#deletetag").val(result.TagValue)
                $("#coverimg").val(result.CoverImage)
                $("#chennelname").val(result.ChannelSlugname);
                $("#showgroup").text(result.ChannelSlugname);
                $('#showgroup').addClass('text-bold').removeClass('text-bold-gray')
                $('#channelids').val(result.ChannelID)
                if (result.ChannelSlugname) {
                    channelnamearr = result.ChannelSlugname.split(',');
                }
                if (result.ChannelID) {

                    ChannelIdarr = result.ChannelID.split(',')

                    $('.select-channel').each(function () {
                        var channelid = $(this).data('id').toString().trim();;
                        if (ChannelIdarr.includes(channelid)) {
                            $(this).prop('checked', true);
                        } else {
                            $(this).prop('checked', false);
                        }
                    });
                }
                // $("#showgroup").text(result.Channel  $('#showgroup').addClass('text-bold').removeClass('text-bold-gray')Slugname);

                $("#blockcoverimg").attr("src", "/image-resize?name="+result.CoverImage)


                if (result.CoverImage != "") {
                    $("#browse").hide();
                    $("#blockcoverimg").show();
                    $("#catdel-img").show()
                    $('#uploadLine').hide()
                    $('#uploadFormat').hide()

                } else {
                    $("#browse").show();
                    $("#blockcoverimg").hide();
                    $("#catdel-img").hide()
                    $('#uploadLine').show()
                    $('#uploadFormat').show()

                }
                // Show tag data in box design

                // tag = result.TagValue
                // tagvalues = tag.split(",")

                // tagvalues.forEach((tagname, index) => {
                //     names = tagname.trim()
                //     finaltags.push(names)
                //     if (tagname != "") {
                //         $(".tagdiv").addClass('show')
                //         $(".tagdiv").removeClass('hidden')

                //         tagdesign = `
                //         <li class="bg-[#ECECEC] p-[6px_12px] flex items-center gap-3 w-fit">
                //          <a href="javascript:void(0);" id="tagdeletebtn" data-id=`+tagname +` >
                //             <img src="/public/img/pill-close.svg" alt="close">
                //          </a>
                //          <span class="text-[14px] font-[300] leading-[17.5px] tracking-[0.01em] text-[#252525]">`+ tagname + `</span>
                //         </li>   `

                //         $(".taglist").append(tagdesign)

                //     }
                // })



                // if (result.CoverImage != "") {
                //     html = `  <div class="border border-[#ECECEC] p-[8px] px-[12px] flex items-center gap-[12px] rounded-[5px] coverimgblock">
                //     <span class="w-[24px] min-w-[24px] h-[24px]"><img id="uploadimg" src="/image-resize?name=` + result.CoverImage + ` " alt=""></span>
                //     <p class="text-[14px] font-normal leading-[17.5px] text-[#252525] text-left" id="filename" > `+ result.CoverImage + ` </p>
                //     <a href="javascript:void(0);" class="min-w-[16px] ml-auto" id="removeimg">
                //        <img src="/public/img/removeCover.svg" alt="remove">
                //      </a>
                //     </div> `

                //     $(html).insertAfter(".imagediv");
                //     $("#browse").css("display", "none")
                //     $("#blockchosimg").css("display", "none")
                // }
            }
        }
    })
})
$('#catdel-img').click(function () {
    $('#categoryimages').val("")
    $('#blockcoverimg').attr('src', '')
    $('#blockcoverimg').hide()
    $('#browse').show()
    $("#mediadesc").css("margin-top", "1%")
    $("#coverimg").val("")
    $(this).siblings('p,button').show()
    $(this).hide()

})
// Cover Image crop 
$(document).on('click', '#blockimg', function () {
    console.log("proidd::");

    $("#prof-crop").val("8")
})

// Collection remove functionality
$(document).on("click", "#collectionremove", function () {

    blockid = $(this).attr("data-blockid")

    $.ajax({
        url: "/admin/blocks/removecollection",
        type: "GET",
        async: false,
        data: { "blockid": blockid },
        datatype: "json",
        caches: false,
        success: function (result) {

            if (result) {

                localStorage.setItem('showToast', 'removetrue');

                window.location.href = "/blocks";


            } else {

                notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"> <li><div class="flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FFB74D] bg-[#FFF7EB]"><a href="javascript:void(0)" class="absolute right-[8px] top-[8px] " id="cancel-notify"><img src="/public/img/close-toast.svg" alt="close"></a><div><img src="/public/img/toast-wraning.svg" alt="toastwraning"></div><div><h3 class="text-[#FFB74D] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] "> ` + languagedata.Toast.internalserverr + ` </p></div></div></li></ul>;`
                $(notify_content).insertBefore(".header-rht");
                // window.location.href = "/blocks"
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }
        }

    })

})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/admin/blocks"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            var value = $(".search").val();
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")
            $(".SearchClosebtn").addClass("hidden")
        } else {
            $(".SearchClosebtn").removeClass("hidden")
            $(".Closebtn").addClass("hidden")
            $(".srchBtn-togg").removeClass("pointer-events-none")
        }
    });
})

$(document).on("click", ".SearchClosebtn", function () {
    $(".SearchClosebtn").addClass("hidden")
    $(".transitionSearch").removeClass("w-[300px] justify-start p-2.5 border border-[#ECECEC] rounded-sm gap-3 overflow-hidden")
    $(".transitionSearch").addClass("w-[32px]")


})

$(document).on("click", ".searchopen", function () {

    $(".SearchClosebtn").removeClass("hidden")

})


$(document).on('click', '.select-channel', function () {

    $('.dropdown-menu').css('position: absolute; inset: 0px auto auto 0px; margin: 0px; transform: translate(-1px, 69px);')
    var Selecttext = $(this).text()
    var channelid = $(this).data('id')
    var channelName = $(this).data('cname')
    if ($(this).is(':checked')) {

        channelnamearr.push(channelName);
        ChannelIdarr.push(channelid)

    } else {

        var index = channelnamearr.indexOf(channelName);
        var index1 = ChannelIdarr.indexOf(channelid)
        if (index > -1) {
            channelnamearr.splice(index, 1);
        }
        if (index1 > -1) {
            ChannelIdarr.splice(index1, 1);
        }
    }
    
    if (channelnamearr[0] === '') {
       
        channelnamearr.shift()
    }
    
    $("#chennelname").val(channelnamearr);
    $('#channelids').val(ChannelIdarr)
    $("#showgroup").text(channelnamearr)
    $('#showgroup').addClass('text-bold').removeClass('text-bold-gray')
    $("#chennelname-error").hide();
    if(channelnamearr.length === 0) {
        
        $('#showgroup').addClass('text-bold-gray').removeClass('text-bold');
        $('#showgroup').text('Choose Channel');
    }

})

$("#searchmembergrp").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()
    console.log(keyword);
    $(".channel-list-row").each(function (index, element) {
        var title = $(element).text().toLowerCase()

        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafounddesign").addClass("hidden")

        } else {
            $(element).hide()
            if ($('.channel-list-row:visible').length == 0) {
                $("#nodatafounddesign").removeClass("hidden")
            }

        }
    })

})


//add to my collection function//

$(document).on('click', '#addtocollect', function () {

    console.log("checkajax")

    blockimg = $(this).attr('data-blockimg')
    blocktitle = $(this).attr('data-title')
    blockslug = $(this).attr('data-slugname')
    blockdata = $(this).attr('data-content')
    blockchannelname = $(this).attr('data-channelname')


    $.ajax({
        url: "/admin/blocks/addtomycollection",
        type: "POST",
        async: false,
        data: { csrf: $("input[name='csrf']").val(), "blocktitle": blocktitle, "blockimg": blockimg, "blockslug": blockslug, "blockdata": blockdata, "blockchannel": blockchannelname },
        datatype: "json",
        caches: false,
        success: function (result) {

            console.log(result.data, "reedf")
            if (result.data == true) {



                notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ` + languagedata.Blocks.collectionsuccessfully + `</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");

                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();

                    });
                }, 5000);

                setTimeout(() => {
                    window.location.href = "/admin/blocks"
                }, 2000);

            } else if (result.data == "channelmissing") {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + languagedata.Toast.Pleaseaddtherequiredchannelbeforeproceeding + `</p></div></div> </li></ul>`;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds
            } else {
                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + languagedata.internalserverr + `</p></div></div> </li></ul>`;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }


        }
    })
})
