var languagedata
var selectedcheckboxarr = []

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    $('.Content').addClass('checked');
})

$('.hd-crd-btn').click(function () {
  
    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show"); 
        document.cookie = `ctabanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `ctabanner=true; path=/;`; 
    }
});

$(document).on('click','.closeimg',function(){

    $('.firstdiv').addClass('hidden')
    document.cookie = `ctadesc=false; path=/;`;
})

/*Cta Status */
function CtaStatus(id) {
    $('#cb' + id).on('change', function () {
      this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb' + id).val();

    console.log("isactive",isactive)
    $.ajax({
      url: '/admin/cta/isactive',
      type: 'POST',
      async: false,
      data: {
        id: id,
        isactive: isactive,
        csrf: $("input[name='csrf']").val()
      },
      dataType: 'json',
      cache: false,
      success: function (result) {


    
       
        // Removetoast();
        if (result == true) {

          notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast.Ctastatusnotify + `</p ></div ></div ></li></ul>`
          $(notify_content).insertBefore(".header-rht");
          setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
              $(this).remove();
            });
          }, 5000); // 5000 milliseconds = 5 seconds
  
          $(this).val(result.IsActive);
        } else {
          notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast.internalserverr + `</p ></div ></div ></li></ul>`
          $(notify_content).insertBefore(".header-rht");
          setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
              $(this).remove();
            });
          }, 5000); // 5000 milliseconds = 5 seconds
  
        }
      }
    });
  }
$(document).on('keyup', '#formsearch', function (event) {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/admin/cta/";

        }
    }
});

$(document).on('keyup', '#unpublishedformsearch', function (event) {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/admin/cta/unpublished";

        }
    }
});

$(document).on('keyup', '#draftformsearch', function (event) {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('.search').val() === "") {

            window.location.href = "/admin/cta/draft";

        }
    }
});

$(document).on("click", "#unpublish", function () {

    var formid = $(this).attr("data-id")
    var status = $(this).attr("data-status")
    var formName = $(this).attr("data-name")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    console.log("page:", pageno);


    $('#content').text("Are you Sure you Want to Unpublish this Form into Website?");
    $('.delname').text(formName)
    $('.deltitle').text("Unpublish Form")
    $("#delid").text(languagedata.Channell.unpublish)
    $("#delcancel").text(languagedata.cancel)

    if (pageno == null) {

        $('#delid').attr('href', "/admin/cta/formstatus?id=" + formid + "&status=" + status);

    } else {

        $('#delid').attr('href', "/admin/cta/formstatus?id=" + formid + "&status=" + status + "&page=" + pageno);

    }
})

$(document).on("click", "#publish", function () {

    var formid = $(this).attr("data-id")
    var status = $(this).attr("data-status")
    var formName = $(this).attr("data-name")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    console.log("page:", pageno);


    $('#content').text("Are you Sure you Want to Publish this Form into Website?");
    $('.delname').text(formName)
    $('.deltitle').text("Publish Form")
    $("#delid").text("Publish")
    $("#delcancel").text(languagedata.cancel)

    if (pageno == null) {

        $('#delid').attr('href', "/admin/cta/formstatus?id=" + formid + "&status=" + status);

    } else {

        $('#delid').attr('href', "/admin/cta/formstatus?id=" + formid + "&status=" + status + "&page=" + pageno);

    }
})

$(document).on('click', "#delete-btn", function () {

    var formid = $(this).attr("data-id")
    var formName = $(this).attr("data-name")
    var status = $(this).attr("data-status")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    console.log("page:", pageno);

    $('#content').text(languagedata.FormBuilder.areyousurewantdelete);
    $('.delname').text(formName)
    $('.deltitle').text(languagedata.FormBuilder.deleteform)
    $("#delid").text(languagedata.FormBuilder.delete)
    $("#delcancel").text(languagedata.cancel)

    if (pageno == null) {

        $('#delid').attr('href', "/admin/cta/deleteform?id=" + formid + "&status=" + status);

    } else {

        $('#delid').attr('href', "/admin/cta/deleteform?id=" + formid + "&status=" + status + "&page=" + pageno);

    }

})

$(document).on('click', ".selectcheckbox", function () {

    var url = window.location.href;

    var urlvalue = url.substring(url.lastIndexOf('/') + 1);

    var text
    var src

    if (urlvalue == "unpublished") {

        text = "Publish"
        src = "/public/img/publish.png"

    } else if (urlvalue == "draft") {

        text = "Publish"
        src = "/public/img/publish.png"

    } else {

        text = "Unpublish"
        src = "/public/img/Unpublished (1) (1).svg"

    }

    $(".renameimg").removeClass('hidden')
    $(".renameimg").attr("src", src)
    $("#unbulishslt span").text(text)

    var formid = $(this).attr("data-id")

    if ($(this).prop('checked')) {

        selectedcheckboxarr.push(formid)

    } else {

        const index = selectedcheckboxarr.indexOf(formid);

        if (index !== -1) {

            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }

    if (selectedcheckboxarr.length != 0) {

        if (selectedcheckboxarr.length == 1) {

            $('.selected-numbers').removeClass('hidden')

            $('.checkboxlength').text(selectedcheckboxarr.length + " " + 'item selected')

            $("#deselectid").html("Deselect")

        } else {

            $('.selected-numbers').removeClass('hidden')

            $('.checkboxlength').text(selectedcheckboxarr.length + " " + 'items selected')

            $("#deselectid").html("Deselect All")

        }


    } else {

        $('.selected-numbers').addClass('hidden')
    }

    var allChecked = true;

    $('.selectcheckbox').each(function () {

        if (!$(this).prop('checked')) {

            allChecked = false;

            return false;
        }
    });

    $('#Check').prop('checked', allChecked);


})

$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text("Delete Forms?")

        $('#content').text('Are you sure want to delete selected Forms?')

    } else {

        $('.deltitle').text("Delete Forms?")

        $('#content').text('Are you sure want to delete selected Forms?')
    }


    $('#delid').addClass('checkboxdelete')
})

//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {

    var url = window.location.href;

    var urlvalue = url.substring(url.lastIndexOf('/') + 1);

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')


    console.log("selectedcheckboxarr:1", selectedcheckboxarr);

    $('.selected-numbers').hide()

    $.ajax({
        url: '/admin/cta/multiselectformdelete',
        type: 'POST',
        dataType: 'json',
        async: false,
        data: { "formids": selectedcheckboxarr, csrf: $("input[name='csrf']").val(), "page": pageno, "urlvalue": urlvalue },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "Form Deleted Successfully")

                window.location.href = data.url

            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})

//Status Change functionality
$(document).on('click', '#unbulishslt', function () {

    var url = window.location.href;

    var urlvalue = url.substring(url.lastIndexOf('/') + 1);

    console.log("urlvalue:", urlvalue);

    if (urlvalue == "unpublished" || urlvalue == "draft") {

        if (selectedcheckboxarr.length > 1) {

            $('.deltitle').text("Publish Forms?")

            $('#delid').text("Publish")

            $('#content').text('Are you sure want to Publish selected Forms?')

        } else {

            $('.deltitle').text("Publish Forms?")

            $('#delid').text("Publish")

            $('#content').text('Are you sure want to Publish selected Forms?')
        }

    } else {

        if (selectedcheckboxarr.length > 1) {

            $('.deltitle').text("Unpublish Forms?")

            $('#delid').text("Unpublish")

            $('#content').text('Are you sure want to Unpublish selected Forms?')

        } else {

            $('.deltitle').text("Unpublish Forms?")

            $('#delid').text("Unpublish")

            $('#content').text('Are you sure want to Unpublish selected Forms?')
        }

    }

    $('#delid').addClass('checkboxstatuschange')
})

//MULTI SELECT Status change FUNCTION//
$(document).on('click', '.checkboxstatuschange', function () {

    var url = window.location.href;

    var urlvalue = url.substring(url.lastIndexOf('/') + 1);

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')


    console.log("selectedcheckboxarr:1", selectedcheckboxarr);

    $('.selected-numbers').hide()

    $.ajax({
        url: '/admin/cta/multiselectstatuschange',
        type: 'POST',
        dataType: 'json',
        async: false,
        data: { "formids": selectedcheckboxarr, csrf: $("input[name='csrf']").val(), "page": pageno, "urlvalue": urlvalue },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", data.toast)

                window.location.href = data.url

            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})

//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass('hidden')

})

$("#filterform").validate({

    rules: {
        fromdate: {
            required: true,

        },
        todate: {
            required: true,

        }

    },
    messages: {
        fromdate: {
            required: "* Please Enter Form Date",

        },
        todate: {
            required: "* Please Enter To Date",
        }

    }
})

// Copy link function
$(document).on("click", '.copyButton', function () {

    var dataValue = $(this).attr('data-id');

    // Copy the data value to the clipboard
    navigator.clipboard.writeText(dataValue).then(function () {
        console.log('Data attribute value copied to clipboard: ' + dataValue);
        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex  w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Link Copied</p ></div ></div ></li></ul> `;
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    }, function (err) {
        console.error('Could not copy text: ', err);
    });
});

$(document).ready(function () {

  

    $(".responsecount").each(function () {
        var count = $(this).text().trim();
        if (count === "") {
            $(this).text("0");
        } else {
            console.log("count:", count);
        }
    });
});

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
  })

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
      window.location.href = "/admin/cta/"
    // var value = $(".formclosebutton").val()
    // console.log("value:", value);
    // if (value == 1) {
    //     window.location.href = "/cta/"
    // } else if (value == 2) {
    //     window.location.href = "/cta/unpublished"
    // } else if (value == 3) {
    //     window.location.href = "/cta/draft"
    // }
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            var value=$(".search").val();
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


$(document).on('click','#ctapreview',function(){

    $('form-builder').removeClass('hidden')
    $('.ctaTitle').addClass("border-0")
    ctaid =$(this).attr('data-id')

    $('form-builder').each(function () {

        if (ctaid !==$(this).attr('data-id')){

            $(this).addClass('hidden')

        }
    })

  
})


//add to my collection function//

$(document).on('click','#addtocollect',function(){

    console.log("checkajax")

    cimg =$(this).attr('data-img')
    title =$(this).attr('data-title')
    desc =$(this).attr('data-desc')
    smallimg=$(this).attr('data-fimg')
    fdata= $(this).attr('data-fdata')
    channelname =$(this).attr('data-channel')

 

    $.ajax({
        url: "/admin/cta/addtomycollection",
        type: "POST",
        async: false,
        data: {  csrf: $("input[name='csrf']").val() ,"title": title,"cimg":cimg,"desc":desc,"smallimg":smallimg,"fdata":fdata, "channelname":channelname },
        datatype: "json",
        caches: false,
        success: function (result) {
          
            if (result.data==true) {
              
                notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ` + languagedata.Blocks.collectionsuccessfully + `</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
        
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
        
                    });
                }, 5000);

                setTimeout(() => {
                     window.location.href="/admin/cta"
                }, 2000);
               
            }else if (result.data=="Pleaseaddtherequiredchannelbeforeproceeding"){

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` + languagedata.Toast.Pleaseaddtherequiredchannelbeforeproceeding + `</p></div></div> </li></ul>`;
                $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();
                        });
                    }, 5000); // 5000 milliseconds = 5 seconds
            }else{
                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">` +  languagedata.internalserverr+ `</p></div></div> </li></ul>`;
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