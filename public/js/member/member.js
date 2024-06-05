var languagedata
var selectedcheckboxarr =[]

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })
})

$(document).keydown(function(event) {
    if (event.ctrlKey && event.key === '/') {
        $("#searchmember").focus().select();
    }
});


$("#addmember , #clickadd").click(function () {

    $('#profpic').attr('src', '/public/img/default profile .svg').show()


    $("#title").text(languagedata.Memberss.addmember)
    $("#triggerId").text(languagedata.Memberss.defaultgroup)
    $("#membergroupvalue").val(1)

    $("#mem_id").val("")
    $("#mem_name").val("")
    $("#mem_lname").val("")
    $("#mem_email").val("")
    $("#mem_mobile").val("")
    $("#mem_usrname").val("")
    $("#mem_activestat").val("")
    $("#membergroupid").val("")
    $("#memberimg").val("")

    $("#mem_name-error").hide()
    $("#mem_lname-error").hide()
    $("#mem_email-error").hide()
    $("#mem_usrname-error").hide()
    $("#myfile-error").css("display", "none")
    // $("#mem_activestat-error").hide()
    $("#membergroupvalue-error").hide()
    $("#memberimg-error").hide()
    $("#mem_pass-error").hide()
    $(".input-group").removeClass("input-group-error")
    $('input[type=hidden][name=crop_data]').val("")
    $(".name-string").hide()
    $("#myfile-error").hide()


    $("#save").show()
    $("#update").hide()

})

// only allow numbers
$('#mem_mobile').keyup(function () {
    this.value = this.value.replace(/[^0-9\.]/g, '');
});

// delete btn
$(document).on('click', '#del', function () {
    var MemberId = $(this).attr("data-id")
    $(".deltitle").text(languagedata.Memberss.deltitle)
    $('.delname').text($(this).parents('tr').find('#membername').text())
    $("#content").text(languagedata.Memberss.delmember)
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')

    if (pageno == null) {
        $('#delete').attr('href', '/member/deletemember?id=' + MemberId);

    } else {
        $('#delete').attr('href', '/member/deletemember?id=' + MemberId + "&page=" + pageno);

    }
})

// editbtn
var edit
$(document).on('click', '#editmem', function () {

    $(".input-group").removeClass("input-group-error")
    $("#mem_name-error").hide()
    $("#mem_lname-error").hide()
    $("#mem_email-error").hide()
    $("#mem_usrname-error").hide()
    $("#membergroupvalue-error").hide()
    $("#memberimg-error").hide()
    $("#mem_pass-error").hide()
    $("#myfile-error").css("display", "none")


    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')   
    $("#memgrbpageno").val(pageno)

    var data = $(this).attr("data-id");
    $("#title").text(languagedata.Memberss.uptmember)
    $("#save").hide()
    $("#update").show()
    $("#memberform").attr("action", "/member/updatemember");
    $("#memberform").attr("name", languagedata.Memberss.upttitle);
    var id = $("#mem_id").val(data)

    edit = $(this).closest("tr");
    var desc = edit.find("td:eq(1)").text();
    $("#triggerId").text(desc);

    $.ajax({
        url: "/member/updatemember",
        type: "GET",
        dataType: "json",
        data: { "id": data },
        success: function (result) {
            $("#mem_id").val(result.Member.Id)
            $("#mem_name").val(result.Member.FirstName)
            $("#mem_lname").val(result.Member.LastName)
            $("#mem_email").val(result.Member.Email)
            $("#mem_mobile").val(result.Member.MobileNo)
            $("#mem_usrname").val(result.Member.Username)
            $("#membergroupvalue").val(result.Member.MemberGroupId)
            // $("#profpic").attr("src", "/" + result.Member.ProfileImagePath)

            if (result.Member.ProfileImagePath != "") {
                $('#profpic').attr('src', result.Member.ProfileImagePath.replace(/^/, '/')).show();
                $(".name-string").hide()
            } else {
                $('#profpic').hide()
                $(".name-string").text(result.Member.NameString).show()
            }

            $("#triggerId").val(result.Group.Name)

            var isactive = $("#cb1").val(result.Member.IsActive)
            $('.tgl-btn').val(isactive)

            if ($("#cb1").val() == 1) {
                $('input[name=mem_activestat]').prop('checked', true)

            }

        }
    })

})




$(document).on('click', '#update', function () {

    jQuery.validator.addMethod("duplicateemail", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checkemailinmember",
            type: "POST",
            async: false,
            data: { "email": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    jQuery.validator.addMethod("duplicatenumber", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checknumberinmember",
            type: "POST",
            async: false,
            data: { "number": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checknameinmember",
            type: "POST",
            async: false,
            data: { "name": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    $.validator.addMethod("mob_validator", function (value) {
        if (/^[6-9]{1}[0-9]{9}$/.test(value))
            return true;
        else return false;
    }, "*" + languagedata.Memberss.memmobnumrgx);

    $.validator.addMethod("email_validator", function (value) {
        return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    }, '* ' + languagedata.Memberss.mememailrgx);

    jQuery.validator.addMethod("pass_validator", function (value, element) {
        if (value != "") {
            if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                return true;
            else return false;
        }
        else return true;
    }, "* " + languagedata.Memberss.mempswdrgx
    );

$.validator.addMethod("noSpaceStart", function(value, element) {
    return this.optional(element) || value.trim() === value;
}, "The first letter cannot be a space.");

$.validator.addMethod("customLength", function (value, element, params) {
    var minLength = params[0];
    var maxLength = params[1];
    return this.optional(element) || (value.length >= minLength && value.length <= maxLength);
}, $.validator.format("Please enter between {0} and {1} characters."));

    $("form[name='updatemember']").validate({
        ignore: [],
        rules: {
            prof_pic: {
                // required: true,
                extension: "jpg|png|jpeg"
            },
            mem_name: {
                required: true,
                // fname_validator: true,
                space: true
            },
            // mem_lname:{
            //     required: true,
            //     space: true
            // },
            mem_email: {
                required: true,
                email_validator: true,
                duplicateemail: true
            },
            // membergroupvalue: {
            //     required: true,
            // },
            mem_usrname: {
                required: true,
                //  username_validator: true,
                space: true,
                duplicatename: true
            },
            companyname:{
                noSpaceStart: true,
            },
            profilename:{
                noSpaceStart: true,
            },
            companylocation:{
                noSpaceStart: true,
            },
            aboutcompany:{
                noSpaceStart: true,
            },
            linkedin:{
                noSpaceStart: true,
            },
            twitter:{
                noSpaceStart: true,
            },
            website:{
                noSpaceStart: true,
            },
            metaTitle:{
                customLength: [1, 60],
                noSpaceStart: true,
            },
            metaKeyword:{
                noSpaceStart: true,
            },
            metaDescription:{
                customLength: [1, 150],
                noSpaceStart: true,
            }
            // mem_pass: {
            //     // required: true,
            //     pass_validator: true,
            // },
            // mem_mobile: {
            //     required: true,
            //     mob_validator: true,
            //     duplicatenumber: true
            //     // number: true,

            // },
            // mem_activestat: {
            //     required: true,
            // }

        },
        messages: {
            prof_pic: {
                // required: "* Please select a profile picture",
                extension: "* " + languagedata.profextension
            },
            mem_name: {
                required: "* " + languagedata.Memberss.memfname,
                space: "* " + languagedata.spacergx
            },
            // mem_lname: {
            //     required: "* " + languagedata.usrlname,
            //     space: "* " + languagedata.spacergx
            // },
            mem_email: {
                required: "* " + languagedata.Memberss.memmail,
                duplicateemail: "* " + languagedata.Memberss.emailexist
            },
            // membergroupvalue: {
            //     required: "* " + languagedata.memgroup
            // },
            mem_usrname: {
                required: "* " + languagedata.Memberss.memname,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Memberss.membernamevaild
            },
            companyname:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            profilename:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
         
            companylocation:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            aboutcompany:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            linkedin:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            twitter:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            website:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaTitle:{
                customLength: "*" + languagedata?.Channell?.metatitlemsg,
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaKeyword:{
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaDescription:{
                customLength: "*" + languagedata?.Channell?.metadescmsg,
                noSpaceStart: "* " + languagedata.spacergx,
            }
            // mem_pass: {
            //     required: "* "+languagedata.Userss.usrpswd
            // },
            // mem_mobile: {
            //     required: "* " + languagedata.Memberss.memmobnum,
            //     duplicatenumber: "* " + languagedata.Memberss.memmobnumexist
            // },
            // mem_activestat: {
            //     required: "* Please checkin the active toggle button",
            // }
        }
    })

    var profileflg 

    $.ajax({
        url: "/member/checkprofilesluginmember",
        type: "POST",
        async:false,
        data: { "name": $("input[name='profilepage']").val(), "id":  $("#mem_id").val(), csrf: $("input[name='csrf']").val() },
        datatype: "json",
        caches: false,
        success: function (data) {
            profileflg = data
        }
    })


    if(profileflg.trim()=="true"){

        $("#profilename-error").show();

        return false

    }

    $("#profilename-error").hide();
    

    $('input[name=mem_pass]').rules('remove', 'required')

    var formcheck = $("#memberform").valid();
    if (formcheck == true) {
        $('#memberform')[0].submit();
    } else {

        $('.input-field-group').each(function () {

            $(this).children('.input-group').each(function(){
                var inputField = $(this).find('input');
    
                if (  !inputField.valid()) {
                    $(this).addClass('input-group-error');
        
                } else {
                    $(this).removeClass('input-group-error');
                }

            })
            if($('#aboutcompany').hasClass('error')){

                    $('#aboutgrb').addClass('input-group-error')
                }
                if($('#aboutcompany').hasClass('valid')){
        
                    $('#aboutgrb').removeClass('input-group-error')
                }
                   if($('#metadesc').hasClass('error')){

                        $('#descgrbmeta').addClass('input-group-error')
                    }
                    if($('#metadesc').hasClass('valid')){
            
                        $('#descgrbmeta').removeClass('input-group-error')
                    }
        
        });

    
        $(document).on('keyup', ".field", function () {

               if ( $(this).hasClass('valid')){

                $(this).parents('.input-group').removeClass('input-group-error')
               }

               if ( $(this).hasClass('error')){

                $(this).parents('.input-group').addClass('input-group-error')
               }
           

            Validationcheck()

            
            
        })

        Validationcheck()
    }

    return false
})
// savebtn

$("#save").click(function () {

    jQuery.validator.addMethod("duplicateemail", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checkemailinmember",
            type: "POST",
            async: false,
            data: { "email": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    jQuery.validator.addMethod("duplicatenumber", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checknumberinmember",
            type: "POST",
            async: false,
            data: { "number": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        var mem_id = $("#mem_id").val()

        $.ajax({
            url: "/member/checknameinmember",
            type: "POST",
            async: false,
            data: { "name": value, "id": mem_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })

    $.validator.addMethod("mob_validator", function (value) {
        if (/^[6-9]{1}[0-9]{9}$/.test(value))
            return true;
        else return false;
    }, "*" + languagedata.Memberss.memmobnumrgx);

    $.validator.addMethod("email_validator", function (value) {
        return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    }, '* ' + languagedata.Memberss.mememailrgx);

    jQuery.validator.addMethod("pass_validator", function (value, element) {
        if (value != "") {
            if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                return true;
            else return false;
        }
        else return true;
    }, "* " + languagedata.Memberss.mempswdrgx
    );

    // Form Validation

    $("form[name='createmember']").validate({

        ignore: [],
        rules: {
            prof_pic: {
                // required: true,
                extension: "jpg|png|jpeg"
            },
            mem_name: {
                required: true,
                // fname_validator: true,
                space: true
            },
            // mem_lname:{
            //     required: true,
            //     space: true
            // },
            mem_email: {
                required: true,
                email_validator: true,
                duplicateemail: true
            },
            // membergroupvalue: {
            //     required: true,
            // },
            mem_usrname: {
                required: true,
                //  username_validator: true,
                space: true,
                duplicatename: true
            },
            mem_pass: {
                required: true,
                pass_validator: true,
            },
           
            // mem_mobile: {
            //     required: true,
            //     mob_validator: true,
            //     duplicatenumber: true
            //     // number: true,

            // },
            // mem_activestat: {
            //     required: true,
            // }

        },
        messages: {
            prof_pic: {
                // required: "* Please select a profile picture",
                extension: "* " + languagedata.profextension
            },
            mem_name: {
                required: "* " + languagedata.Memberss.memfname,
                space: "* " + languagedata.spacergx
            },
            // mem_lname: {
            //     required: "* " + languagedata.usrlname,
            //     space: "* " + languagedata.spacergx
            // },
            mem_email: {
                required: "* " + languagedata.Memberss.memmail,
                duplicateemail: "* " + languagedata.Memberss.emailexist
            },
            // membergroupvalue: {
            //     required: "* " + languagedata.memgroup
            // },
            mem_usrname: {
                required: "* " + languagedata.Memberss.memname,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Memberss.membernamevaild
            },
            mem_pass: {
                required: "* " + languagedata.Memberss.mempswd
            },
         

            // mem_mobile: {
            //     required: "* " + languagedata.Memberss.memmobnum,
            //     duplicatenumber: "* " + languagedata.Memberss.memmobnumexist
            // },
            // mem_activestat: {
            //     required: "* Please checkin the active toggle button",
            // }
        }
    })

    var formcheck = $("#memberform").valid();
    if (formcheck == true) {
        $('#memberform')[0].submit();
        $('#save').prop('disabled', true);
    }
    else {

        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        $('.input-group').each(function () {
            var inputField = $(this).find('input');
            var inputName = inputField.attr('name');


            if (inputName !== 'membergroupvalue' && !inputField.valid()) {
                $(this).addClass('input-group-error');

            } else {
                $(this).removeClass('input-group-error');
            }
        });

    }

    return false
})

$(".dropdown-values").on("click", function () {
    var text = $(this).text()
    var id = $(this).attr("data-id")
    $(this).parents().siblings("a").text(text)
    $(this).parents().siblings("input[name='membergroupvalue']").val(id)
    if ($('#membergroupvalue').val() !== '') {
        $('#membergroupvalue-error').hide()
        $('.user-drop-down').removeClass('input-group-error')

    }
})

// Is_active

$("input[name=mem_activestat]").click(function () {
    if ($(this).prop('checked') == true) {
        $(this).attr("value", "1")
    } else {
        $(this).removeAttr("value")
    }
})

// Claim status
$("input[name=com_activestat]").click(function () {
    if ($(this).prop('checked') == true) {
        $(this).attr("value", "1")
    } else {
        $(this).removeAttr("value")
    }
})
// Email status
$("input[name=mem_emailactive]").click(function () {
    if ($(this).prop('checked') == true) {
        $(this).attr("value", "1")
    } else {
        $(this).removeAttr("value")
    }
})
/*search */
$(document).on("submit", "#filterformsubmit", function () {
    var key = $(this).siblings().children(".search").val();
    if (key == "") {
        window.location.href = "/member/"
    } else {
        $('.filterform').submit();
    }
})

// cancel btm

$("#cancel").click(function () {
    window.location.href = "/member/";
})

$(document).on('click', '#newdd-input', function () {
    $(this).siblings('.dd-c').css('display', 'block')
    $("#searchrole").val("");
})

// $('.page-wrapper').on('click', function (event) {
//     if ($('.dd-c').css('visibility') == 'visible' && !$(event.target).is('#newdd-input') && !$(event.target).is('.dd-c')) {
// $('#newdd-input').prop('checked',false)
// $('.dd-c').css('display','none')
//     }
// })


$('#newdd-input').blur(function () {
    if ($('#valmem').val() !== '') {
        $(this).parents('.input-group').removeClass('validate')
    }
    $(this).parents('.input-group').removeClass('focus')

})

function Validationcheck() {
    let inputGro = document.querySelectorAll('.input-group');
    inputGro.forEach(inputGroup => {
        let inputField = inputGroup.querySelector('input:not([name="membergroupvalue"])');
      

        if (inputField.classList.contains('error')) {
            inputGroup.classList.add('input-group-error');
        } else {
            inputGroup.classList.remove('input-group-error');
        }
        if ($('#membergroupvalue-error').css('display') !== 'none') {
            $('#memgrp').addClass('input-group-error')
        }
        else {
            $('#memgrp').removeClass('input-group-error')
        }
    });

  
}

$(document).on("click", "#editmembergroup", function () {
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');
    if (pageno == null) {
        pageno = "1";
    }
    $("#memgrbpageno").val(pageno)
    // var memberId = $(this).data("id");
    var editUrl = $(this).attr("href");
    window.location.href = editUrl + "&page=" + pageno;
});

$("#searchrole").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".choose-rel-articles .newck-group").each(function (index, element) {
        var title = $(element).find('h4').text().toLowerCase()
        if (title.includes(keyword)) {
            $(element).show()
        } else {
            $(element).hide()
           
        }
    })
})

function refreshdiv() {
    $('.choose-rel-articles').load(location.href + ' .choose-rel-articles');
}

/*search redirect home page */

$(document).on('keyup', '#searchmember', function () {

    if (event.key === 'Backspace') {
       
        if ($('.search').val() === "") {
        
            window.location.href = "/member/";
        }
    }
})
$(document).on('click','.closemember',function(){


    window.location.href="/member/"
})
// pasword show and close

$(document).on('click', '#eye', function () {

    var This = $("#mem_pass")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).addClass('open')

        $("#img").attr("src", "/public/img/eye-opened.svg")

    } else {

        $(This).attr('type', 'password');

        $(this).removeClass('close')

        $("#img").attr("src", "/public/img/eye-closed.svg")

    }
})

// $('form[class=filterform]>img').click(function () {

//     if ($(this).siblings('input[name=keyword]').val() != "" && $(this).parents('.transitionSearch').hasClass('active')) {

//         var keyword = $(this).siblings('input[name=keyword]').val()

//         window.location.href = "/membersgroup/?keyword=" + keyword
//     }
// })

// $('#memModal').on('hide.bs.modal', function () {

//     $('.input-field-group .input-group').each(function () {

//         if ($(this).hasClass('input-group-error')) {

//             $(this).removeClass('input-group-error')
//         }
//     })

//     $('label[class=error').removeClass('error').hide()

//     $("#memberform").attr("action", "/member/newmember");

//     $("#memberform").attr("name", "createmember");

//     $('#mem_pass-error').remove()

//     // $('input[name=mem_activestat]').remove()

//     // $('<input class="tgl tgl-light" name="mem_activestat" id="mem_activestat" type="checkbox">').insertBefore('label[for=mem_activestat]')

//     $('input[type=hidden][name=membergroupvalue]').val("")

// })

// dropdown filter input box search
$("#searchmembergrp").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".membergrp-list-row button").each(function (index, element) {
        var title = $(element).text().toLowerCase()

        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafoundtext").hide()

        } else {
            $(element).hide()
            if($('.membergrp-list-row button:visible').length==0){
                $("#nodatafoundtext").show()
            }
           
        }
    })

})


// search key value empty in dropdown search

$("#triggerId").on("click",function(){

    var keyword = $("#searchmembergrp").val()
    $(".membergrp-list-row button").each(function (index, element) {
        if(keyword != ""){
        $("#searchmembergrp").val("")
        $("#nodatafoundtext").hide()
        $(element).show()
      }else{
        $("#searchmembergrp").val()
        $("#nodatafoundtext").hide()
        $(element).show()
    
      }
    })
   
  })

  $(document).on('click', '#myfile', function () {
    $("#prof-crop").val("1")
 })

 $(document).on('click', '#cmpymyfile', function () {
    $("#prof-crop").val("2")
 })

 // async function profilenamevalidator(){

//     $.ajax({
//         url: "/member/checkprofilenameinmember",
//         type: "POST",
//         async: false,
//         data: { "name": $("input[name='profilename']").val(), "id":  $("#mem_id").val(), csrf: $("input[name='csrf']").val() },
//         datatype: "json",
//         caches: false,
//         success: function (data) {
           
//             return data

//         }
//     })

//  }
function MemberStatus(id) {
    $('#cbox' + id).on('change', function () {
         this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cbox' + id).val();

    console.log("check",isactive,id)

      $.ajax({
         url: '/member/memberisactive',
         type: 'POST',
         async: false,
         data: {"id": id,"isactive": isactive,csrf: $("input[name='csrf']").val()},
         dataType: 'json',
         cache: false,
         success: function (result) { 
             if (result) {

                   notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>Member Status Updated Successfully</span></div>';
                   $(notify_content).insertBefore(".header-rht");
                   setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                             $(this).remove();
                        });
                   }, 5000); // 5000 milliseconds = 5 seconds

              } else {
                   
                   notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                   $(notify_content).insertBefore(".breadcrumbs");
                   setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                             $(this).remove();
                        });
                   }, 5000); // 5000 milliseconds = 5 seconds

              }
         }
    });
}
$(document).on('click','.selectcheckbox',function(){

    memberid =$(this).attr('data-id')

    var status = $(this).parents('td').siblings('td').find('.tgl-light').val();

    console.log(status,"status")


   if ($(this).prop('checked')){

       selectedcheckboxarr.push({"memberid":memberid,"status":status})
   
   }else{

       const index = selectedcheckboxarr.findIndex(item => item.memberid === memberid);
   
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
    
           memberid= $(this).attr('data-id')

           var status = $(this).parents('td').siblings('td').find('.tgl-light').val();
    
           selectedcheckboxarr.push({"memberid":memberid,"status":status})
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
         
    $('.deltitle').text("Delete Members?")

    $('#content').text('Are you sure want to delete selected Members?')

    }else {

         $('.deltitle').text("Delete Member?")

         $('#content').text('Are you sure want to delete selected Member?')
    }


    $('#delete').addClass('checkboxdelete')
})

$(document).on('click','#unbulishslt',function(){

    if (selectedcheckboxarr.length>1){

         $('.deltitle').text( $(this).text()+" "+"Members?")

         $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Members?")
    }else{

         $('.deltitle').text( $(this).text()+" "+"Member?")

         $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Member?")
    }

    $('#delete').addClass('selectedunpublish')

})

//MULTI SELECT DELETE FUNCTION//
$(document).on('click','.checkboxdelete',function(){

    var url = window.location.href;

    console.log("url", url)

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/member/deleteselectedmember',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "memberids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page":pageno

            
        },
        success: function (data) {

            console.log(data,"result")

            if (data.value==true){

                setCookie("get-toast", "Member Deleted Successfully")

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

    var url = window.location.href;

    console.log("url", url)

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/member/multiselectmemberstatus',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "memberids":JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page":pageno

            
        },
        success: function (data) {

            console.log(data,"result")

            if (data.value==true){

                setCookie("get-toast", "memstatusnotify")

                window.location.href=data.url
            }else{

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})

$('#profile-slugData').on('input', function(event){

    var inputvalue = $(this).val()

    var smallcase = inputvalue.toLowerCase()

    var modifiedText = smallcase.replace(/ /g, '-');

    $(this).val(modifiedText)
})