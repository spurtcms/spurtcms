var languagedata

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
        })

            $('.Content').addClass('checked');

})


//savebutton function
$(document).on('click','#savebtn,#savemenubtn',function(){

      
        jQuery.validator.addMethod("duplicatename", function (value) {

            var result;
            menu_id = $("#menu_id").val()
    
            $.ajax({
                url: "/admin/website/menu/checkmenuname",
                type: "POST",
                async: false,
                data: { "menu_name": value,"webid":$(".templateid").val(), "menu_id": menu_id, csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })
    $("#menuform").validate({
        rules: {
            menu_name: {
                required: true,
                space: true,
                duplicatename:true,
               
            },
            menu_title: {
                required: true,
                space: true,
                duplicatename:true,
            },
            menu_desc: {
                // required: true,
                maxlength: 250,
                // space: true,
            }
        },
        messages: {
            menu_name: {
                required: "*"+languagedata.Menu.menunameerr ,
                space: "* " + languagedata.spacergx,
                duplicatename: "*"+languagedata.Menu.nameduplicateerr 
            },
            menu_title: {
                required: "* Please Enter Menu Title",
                space: "* " + languagedata.spacergx,
            },
            menu_desc: {
                // required: "*"+languagedata.Menu.descriptionerror ,
                maxlength: "* " + languagedata.Permission.descriptionchat,
                // space: "* " + languagedata.spacergx,
            },
        }
    });

    var formcheck = $("#menuform").valid();
    if (formcheck == true) {
        $('#menuform')[0].submit();
    }
   
    return false

})

//Status button change function

$("input[name=menustatus]").click(function () {
    if ($(this).prop('checked') == true) {
        $(this).attr("value", "1")
    } else {
        $(this).attr("value", "0")
    }
})


//Edit Menu function
$(document).on('click','#editbtn',function(){


    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    var data = $(this).attr("data-id");
    edit = $(this).closest("tr");
    $("#menuform").attr("name", "editmenu").attr("action", "/admin/website/menu/updatemenu/?webid="+$('.templateid').val())
    var name = edit.find("td:eq(0) a").attr('data-name');
    var title = edit.find("td:eq(0) a").text();
    var desc = edit.find("td:eq(1)").attr('data-desc');
    var checkbox = edit.find("td:eq(2) input[type='checkbox']").is(':checked');
    // var statusValue = checkbox.val();
    console.log("name,",checkbox)
    $("input[name=menu_name]").val(name.trim());
    $("input[name=menu_title]").val(title.trim());

    $("textarea[name=menu_desc]").val(desc.trim());

    if ((name.trim()=="Aside")||(name.trim()=="Headers") ||(name.trim()=="Footers")||(name.trim()=="SERVICES")||(name.trim()=="SOLUTIONS")||(name.trim()=="EXPERTISE")){

        $('#menu_name').addClass('pointer-events-none opacity-50')
    }else{
         $('#menu_name').removeClass('pointer-events-none opacity-50')
    }
    if (checkbox == 1) {
        $('.menustatus').prop('checked', true).val('1');
    } else {
        $('.menustatus').prop('checked', false).val('0');
    }
    
    $("input[name=menu_id]").val(data);
    $("#savebtn").text(languagedata.update)

    console.log(languagedata.update,languagedata.Menu.menu,"updatedfdfdf")
    $('#modalTitleId').text(languagedata.update+languagedata.Menu.menu)
    $("#menu_name-error").hide()
    $("#menu_desc-error").hide()
    $('#menupageno').val(pageno)
});

//delete func//

$(document).on('click', '#deletebtn', function () {

    var menuid = $(this).attr("data-id")
    $("#content").text(languagedata.Menu.removeconfirmation)
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    templateid =$('.webid').attr('data-id')

console.log(templateid)

    if (pageno == null) {
        $('#delid').attr('href', "/admin/website/menu/deletemenu/" + menuid +"?webid="+templateid);

    } else {
        $('#delid').attr('href', "/admin/website/menu/deletemenu/" + menuid + "?page=" + pageno & "?webid="+templateid);

    }

   
    $(".deltitle").text(languagedata.Menu.deletemenu +"?")
    $('.delname').text($(this).parents('tr').find('td:eq(0) a').text())

})

//Cancel button function//

$(document).on('click', '.cancelbtn', function () {
    $("#add-btn").show()
    $("#savebtn").text(languagedata.Save)
    
    $("#menuform").attr("action", "/admin/website/menu/createmenu/?webid="+$('.templateid').val())
    $("input[name=menu_id]").val("")
    $("input[name=menu_name]").val("")
    // $("input[name=menu_desc]").val("")
    $("#menu_name-error").hide()
    $("#menu_desc-error").hide()
    $("#error-messagename").html("")
    $("#error-messagedesc").html("")
    $('.menustatus').prop('checked', false).val('0');
    $('#modalTitleId').text(languagedata.createnewmenu)
    $('#menu_desc').val("")
    $('#menu_name').removeClass('pointer-events-none')

})

//search cancel button function//
$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/admin/website/menu?webid="+$('.webid').attr('data-id')
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

//Filter dropdown function//

$(document).on('click','.statuscls',function(){

    statusval =$(this).text().trim()
    $(".filterleveldropdown").removeClass("show")
    $('.slctstatus').text(statusval)
    $('#statusid').val(statusval)
})

//Menu Status Change Functionlity//


function MenuStatus(id) {
    $('#cb' + id).on('change', function () {
        console.log("printf");
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb' + id).val();

    console.log("check", isactive, id)

    $.ajax({
        url: '/admin/website/menu/menustatuschange',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Menu Updated Successfully</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            } else {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
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

$('.hd-crd-btn').click(function () {
  
    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show"); 
        document.cookie = `webbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `webbanner=true; path=/;`; 
    }
});


$('.createmenuclass').click(function(){

    console.log("hhhh")

    $('#modalTitleId').text('Create New Menu')

    $('#savebtn').text('Save')

    $('#menu_name').removeClass(' pointer-events-none opacity-50')
})