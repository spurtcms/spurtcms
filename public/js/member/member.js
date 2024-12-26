var languagedata
var selectedcheckboxarr = []
var setstatus

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
})

// focus for search data
$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $("#searchmember").focus().select();
    }
});

// Create new member rightside model open
$("#addmember , #clickadd").click(function () {
    $("#title").text(languagedata.Memberss.addmember)
    $("#showgroup").text(languagedata.Memberss.choosegroup)
    $('#showgroup').removeClass('text-bold').addClass('text-bold-gray')

    // $("#membergroupvalue").val(1)
    $("#mem_id").val("")
    $("#mem_name").val("")
    $("#mem_lname").val("")
    $("#mem_email").val("")
    $("#mem_mobile").val("")
    $("#mem_usrname").val("")
    $("#mem_activestat").val("")
    $("#membergroupid").val("")
    $("#memberimg").val("")
    $("#mem_pass").val("");
    $('#membergroupvalue').val('');


    $("#mem_name-error").hide()
    $("#mem_lname-error").hide()
    $("#mem_email-error").hide()
    $("#mem_mobile-error").hide()
    $("#mem_usrname-error").hide()
    $("#membergroupvalue-error").hide()
    $("#memberimg-error").hide()
    $("#mem_pass-error").hide()
    $(".input-group").removeClass("input-group-error")
    $('input[type=hidden][name=crop_data]').val("")
    $(".name-string").hide()
    $("#Save").show()
    $("#update").hide()

})


// when close model then work this --(using calcel btn)
$("#modalId2").on("hide.bs.modal", function () {
    $("#myfile-error").addClass("hidden");
    $(".lengthErr").addClass("hidden");
    $("#profpic-mem").attr('src', '/public/img/default profile .svg');


})

// only allow numbers (mobile number validation)
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
        console.log("varuthu");
        $('#delid').attr('href', '/member/deletemember?id=' + MemberId);

    } else {
        $('#delid').attr('href', '/member/deletemember?id=' + MemberId + "&page=" + pageno);

    }
})

// editbtn (not use now)

// var edit
// $(document).on('click', '#editmem', function () {
// console.log("edit dffgdg");
//     $(".input-group").removeClass("input-group-error")
//     $("#mem_name-error").hide()
//     $("#mem_lname-error").hide()
//     $("#mem_email-error").hide()
//     $("#mem_usrname-error").hide()
//     $("#membergroupvalue-error").hide()
//     $("#memberimg-error").hide()
//     $("#mem_pass-error").hide()
//     $("#myfile-error").css("display", "none")


//     var url = window.location.search
//     const urlpar = new URLSearchParams(url)
//     pageno = urlpar.get('page')
//     $("#memgrbpageno").val(pageno)

//     var data = $(this).attr("data-id");
//     $("#title").text(languagedata.Memberss.uptmember)
//     $("#save").hide()
//     $("#update").show()
//     // $("#memberprofileform").attr("action", "/member/updatemember");
//     // $("#memberprofileform").attr("name", languagedata.Memberss.upttitle);
//     var id = $("#mem_id").val(data)

//     edit = $(this).closest("tr");
//     var desc = edit.find("td:eq(1)").text();
//     console.log("desc",desc);
//     $("#triggerId").text(desc);

//     $.ajax({
//         url: "/member/updatemember",
//         type: "GET",
//         dataType: "json",
//         data: { "id": data },
//         success: function (result) {
//             $("#mem_id").val(result.Member.Id)
//             $("#mem_name").val(result.Member.FirstName)
//             $("#mem_lname").val(result.Member.LastName)
//             $("#mem_email").val(result.Member.Email)
//             $("#mem_mobile").val(result.Member.MobileNo)
//             $("#mem_usrname").val(result.Member.Username)
//             $("#membergroupvalue").val(result.Member.MemberGroupId)
//             // $("#profpic").attr("src", "/" + result.Member.ProfileImagePath)

//             if (result.Member.ProfileImagePath != "") {
//                 $('#profpic').attr('src', result.Member.ProfileImagePath.replace(/^/, '/')).show();
//                 $(".name-string").hide()
//             } else {
//                 $('#profpic').hide()
//                 $(".name-string").text(result.Member.NameString).show()
//             }

//             $("#triggerId").val(result.Group.Name)

//             var isactive = $("#cb1").val(result.Member.IsActive)
//             $('.tgl-btn').val(isactive)

//             if ($("#cb1").val() == 1) {
//                 $('input[name=mem_activestat]').prop('checked', true)

//             }

//         }
//     })

// })



// update membes
$(document).on('click', '#update', function () {
    $("#showgroup").text(languagedata.Memberss.defaultgroup)
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

    $.validator.addMethod("noSpaceStart", function (value, element) {
        return this.optional(element) || value.trim() === value;
    }, "The first letter cannot be a space.");

    $.validator.addMethod("customLength", function (value, element, params) {
        var minLength = params[0];
        var maxLength = params[1];
        return this.optional(element) || (value.length >= minLength && value.length <= maxLength);
    }, $.validator.format("Please enter between {0} and {1} characters."));




    $("#memberprofileform").validate({
        ignore: [],
        rules: {
            prof_pic: {
                extension: "jpg|png|jpeg"
            },
            mem_name: {
                required: true,
                space: true
            },
            mem_email: {
                required: true,
                email_validator: true,
                duplicateemail: true
            },
            mem_usrname: {
                required: true,
                space: true,
                duplicatename: true
            },
            mem_mobile: {
                required: true,
                mob_validator: true,
                duplicatenumber: true
            },
            companyname: {
                noSpaceStart: true,
            },
            profilename: {
                noSpaceStart: true,
            },
            companylocation: {
                noSpaceStart: true,
            },
            aboutcompany: {
                noSpaceStart: true,
            },
            linkedin: {
                noSpaceStart: true,
            },
            twitter: {
                noSpaceStart: true,
            },
            website: {
                noSpaceStart: true,
            },
            metaTitle: {
                customLength: [1, 200],
                noSpaceStart: true,
            },
            metaKeyword: {
                noSpaceStart: true,
            },
            metaDescription: {
                customLength: [1, 500],
                noSpaceStart: true,
            }

        },
        messages: {
            prof_pic: {
                extension: "* " + languagedata.profextension
            },
            mem_name: {
                required: "* " + languagedata.Memberss.memfname,
                space: "* " + languagedata.spacergx
            },

            mem_email: {
                required: "* " + languagedata.Memberss.memmail,
                duplicateemail: "* " + languagedata.Memberss.emailexist
            },

            mem_usrname: {
                required: "* " + languagedata.Memberss.memname,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Memberss.membernamevaild
            },
            mem_mobile: {
                required: "* " + languagedata.Userss.usrmobnum,
                duplicatenumber: "* " + languagedata.Userss.mobnumexist,
            },
            companyname: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            profilename: {
                noSpaceStart: "* " + languagedata.spacergx,
            },

            companylocation: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            aboutcompany: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            linkedin: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            twitter: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            website: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaTitle: {
                customLength: "*" + languagedata.Channell.metatitlemsg,
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaKeyword: {
                noSpaceStart: "* " + languagedata.spacergx,
            },
            metaDescription: {
                customLength: "*" + languagedata.Channell.metadescmsg,
                noSpaceStart: "* " + languagedata.spacergx,
            }
        }
    })




    var profileflg

    $.ajax({
        url: "/member/checkprofilesluginmember",
        type: "POST",
        async: false,
        data: { "name": $("input[name='profilepage']").val(), "id": $("#mem_id").val(), csrf: $("input[name='csrf']").val() },
        datatype: "json",
        caches: false,
        success: function (data) {
            profileflg = data.trim()
            console.log(profileflg, "data");
        }
    })





    if (profileflg.trim() == "true") {

        $("#profilename-error").show();

        return false

    }

    $("#profilename-error").hide();


    $('input[name=mem_pass]').rules('remove', 'required')


    var formcheck = $("#memberprofileform").valid();
    console.log("working this ");
    if (formcheck == true) {
        $('#memberprofileform')[0].submit();
    } else {
        console.log(" not working this ");
        $('.input-field-group').each(function () {

            $(this).children('.input-group').each(function () {
                var inputField = $(this).find('input');

                if (!inputField.valid()) {
                    $(this).addClass('input-group-error');

                } else {
                    $(this).removeClass('input-group-error');
                }

            })
            if ($('#aboutcompany').hasClass('error')) {

                $('#aboutgrb').addClass('input-group-error')
            }
            if ($('#aboutcompany').hasClass('valid')) {

                $('#aboutgrb').removeClass('input-group-error')
            }
            if ($('#metadesc').hasClass('error')) {

                $('#descgrbmeta').addClass('input-group-error')
            }
            if ($('#metadesc').hasClass('valid')) {

                $('#descgrbmeta').removeClass('input-group-error')
            }

        });


        $(document).on('keyup', ".field", function () {

            if ($(this).hasClass('valid')) {

                $(this).parents('.input-group').removeClass('input-group-error')
            }

            if ($(this).hasClass('error')) {

                $(this).parents('.input-group').addClass('input-group-error')
            }


            Validationcheck()



        })

        Validationcheck()
    }

    return false
})


// savebtn
$("#Save").click(function () {

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
                console.log(data, "data");

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

    $("#memberform").validate({
        ignore: [],
        rules: {
            prof_pic: {
                extension: "jpg|png|jpeg"
            },
            mem_name: {
                required: true,
                space: true
            },
            mem_email: {
                required: true,
                email_validator: true,
                duplicateemail: true
            },
            mem_usrname: {
                required: true,
                space: true,
                duplicatename: true
            },
            mem_pass: {
                required: true,
                pass_validator: true,
            },
            mem_mobile: {
                required: true,
                mob_validator: true,
                duplicatenumber: true
            },
            membergroupvalue:{
                required: true
            }

        },
        messages: {
            prof_pic: {
                extension: "* " + languagedata.profextension
            },
            mem_name: {
                required: "* " + languagedata.Memberss.memfname,
                space: "* " + languagedata.spacergx
            },

            mem_email: {
                required: "* " + languagedata.Memberss.memmail,
                duplicateemail: "* " + languagedata.Memberss.emailexist
            },

            mem_usrname: {
                required: "* " + languagedata.Memberss.memname,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Memberss.membernamevaild
            },
            mem_pass: {
                required: "* " + languagedata.Memberss.mempswd
            },
            mem_mobile: {
                required: "* " + languagedata.Userss.usrmobnum,
                duplicatenumber: "* " + languagedata.Userss.mobnumexist,
            },
            membergroupvalue:{
                required: "* " + languagedata.memgroup
            }

        }
    })



    var formcheck = $("#memberform").valid();

    if (formcheck == true) {
        console.log("hi");
        $('#memberform')[0].submit();
        $('#Save').prop('disabled', true);
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


// search functions strat //


/*search redirect home page */

$(document).on('keyup', '#searchmember', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/member/";
        }
    }

    $('.searchClosebtn').toggleClass('hidden', searchInput === "");

})

$(document).on("click", ".Closebtn", function () {
    $(".Searchmem").val('')
    $(".Closebtn").addClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")

  })

  $(document).on("click", ".searchClosebtn", function () {
    $(".Searchmem").val('')
    window.location.href = "/member/"
  })

  $(document).ready(function () {

    $('.Searchmem').on('input', function () {

        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")

        }else{
          $(".Closebtn").addClass("hidden")
          $(".srchBtn-togg").removeClass("pointer-events-none")

        }
    });
  })

  $(document).on("click", ".hovericon", function () {
    $(".Searchmem").val('')
    $(".Closebtn").addClass("hidden")
  })


// search functions end //


// cancel btn
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


// $('#newdd-input').blur(function () {
//     if ($('#valmem').val() !== '') {
//         $(this).parents('.input-group').removeClass('validate')
//     }
//     $(this).parents('.input-group').removeClass('focus')

// })

function Validationcheck() {
    // let inputGro = document.querySelectorAll('.input-group');
    // inputGro.forEach(inputGroup => {
    //     let inputField = inputGroup.querySelector('input:not([name="membergroupvalue"])');


    //     if (inputField.classList.contains('error')) {
    //         inputGroup.classList.add('input-group-error');
    //     } else {
    //         inputGroup.classList.remove('input-group-error');
    //     }
    //     if ($('#membergroupvalue-error').css('display') !== 'none') {
    //         $('#memgrp').addClass('input-group-error')
    //     }
    //     else {
    //         $('#memgrp').removeClass('input-group-error')
    //     }
    // });
}


$(".chk-group").on("click", function () {
    var Selecttext = $(this).text()
    var Selectvalue = $(this).data('id')
    console.log("select group", Selecttext);
    $("#membergroupvalue").val(Selectvalue);
    $("#showgroup").text(Selecttext);
    $('#showgroup').addClass('text-bold').removeClass('text-bold-gray')
    $("#membergroupvalue-error").hide();

});


$(".chk-group1").on("click", function () {
    var selectValue1 = $(this).data('id');
    var Selecttext1 = $(this).text()
    $("#membergroupvalue1").val(selectValue1);
    $("#showgroup1").text(Selecttext1);
});


$(document).on("click", "#editmembergroup", function () {
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');
    if (pageno == null) {
        pageno = "1";
    }
    $("#memgrbpageno").val(pageno)
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





$(document).on('click', '.closemember', function () {


    window.location.href = "/member/"
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


// dropdown filter input box search
$("#searchmembergrp").keyup(function () {

    var keyword = $(this).val().trim().toLowerCase()
    console.log(keyword);
    $(".membergrp-list-row").each(function (index, element) {
        var title = $(element).text().toLowerCase()

        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafounddesign").addClass("hidden")

        } else {
            $(element).hide()
            if ($('.membergrp-list-row:visible').length == 0) {
                $("#nodatafounddesign").removeClass("hidden")
            }

        }
    })

})


$(document).on('click', '#myfile', function () {
    $("#prof-crop").val("1")
})

$(document).on('click', '#cmpymyfile', function () {

    $("#prof-crop").val("2")
})

$(document).on('click', "#profileImgLabel", function () {
    $("#prof-crop").val("3")
})

$(document).on('change', '#profileImgLabel', function () {
    $("#name-string span").remove();

    $("#name-string").append('<img name="profpicture" id="profpicture"/>');
});

$(document).on('change', '#profileImgLabel', function () {
    $("#mem-img").attr('id', 'profpicture').attr('name', 'profpicture')

});

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
        console.log("printf");
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cbox' + id).val();

    console.log("check", isactive, id)

    $.ajax({
        url: '/member/memberisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >${languagedata.Toast.MemberStatusUpdatedSuccessfully}</p ></div ></div ></li></ul> `;
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


// select checkbox to purform actions

$(document).on('click', '.selectcheckbox', function () {

    memberid = $(this).attr('data-id')

    var status = $(this).parents('td').siblings('td').find('.tgl-light').val();


    var hasDelId = $(this).parents('td').siblings("td").find('#del').length > 0;
    var statusid = $(this).parents('td').siblings('td').find('.tgl-light').length >0


    if (hasDelId) {

        $('#seleccheckboxdelete').show()
    } else {
        $('#seleccheckboxdelete').hide()
    }

    if (statusid) {

        $('#unbulishslt').show()

    } else {
        $('#unbulishslt').hide()

        $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')
    }

    if ($(this).prop('checked')) {

        selectedcheckboxarr.push({ "memberid": memberid, "status": status })


    } else {

        const index = selectedcheckboxarr.findIndex(item => item.memberid === memberid);

        if (index !== -1) {
            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {
        $('.selected-numbers').removeClass("hidden")

        if (selectedcheckboxarr.length == 1) {
            $('#deselectid').text(languagedata.deselect)
        } else if (selectedcheckboxarr.length > 1) {
            $('#deselectid').text(languagedata.deselectall)
        }

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus
        var img;


        if (selectedcheckboxarr[0].status === '1') {

            setstatus = languagedata.Memberss.deactive;

            img = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = languagedata.Memberss.active;

            img = "/public/img/Active.svg";

        }

        var htmlContent = '';

        if (allSame) {

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';


        } else {

            htmlContent = '';

        }

        $('#unbulishslt').html(htmlContent);


    var items

    if (selectedcheckboxarr.length==1){

        items ="Item Selected"
    }else{

        items = languagedata.itemselected
    }
    $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)


        if (!allSame || !statusid) {

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

            $('.unbulishslt').html("")
        } else {

            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')
        }


    } else {

        $('.selected-numbers').addClass("hidden")
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



//  //ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click', '#Check', function () {

    selectedcheckboxarr = []
    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            memberid = $(this).attr('data-id')

            var status = $(this).parents('td').siblings('td').find('.tgl-light').val();

            selectedcheckboxarr.push({ "memberid": memberid, "status": status })
        })
        if (selectedcheckboxarr.length == 0) {
            $('.selected-numbers').addClass('hidden');
        } else {
            const deselectText = selectedcheckboxarr.length == 1 ? languagedata.deselect : languagedata.deselectall;
            $('#deselectid').text(deselectText);
            $('.selected-numbers').removeClass('hidden');
        }


        var allSame = selectedcheckboxarr.every(function (item) {

            return item.status === selectedcheckboxarr[0].status;
        });


        var img

        if (selectedcheckboxarr.length > 0 && selectedcheckboxarr[0].status !== undefined) {
            if (selectedcheckboxarr[0].status === '1') {
                setstatus = languagedata.Memberss.deactive;
                img = "/public/img/In-Active.svg";
            } else if (selectedcheckboxarr[0].status === '0') {
                setstatus = languagedata.Memberss.active;
                img = "/public/img/Active.svg";
            }
        } else {
            console.error("selectedcheckboxarr is empty or status is undefined.");
        }

        var htmlContent = '';

        if (allSame) {

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';

            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        } else {

            htmlContent = '';

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.Memberss.itemsselected)

    } else {
        $('.selected-numbers').addClass("hidden")

        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);
    }


})

// ------------------------------------------------------------------

// delete model cuntent updations

$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text(languagedata.Memberss.deletemembers)

        $('#content').text(languagedata.Memberss.deletecontents)

    } else {

        $('.deltitle').text(languagedata.Memberss.deletemember)

        $('#content').text(languagedata.Memberss.deletecontent)
    }

    $("#delid").text($(this).text());
    $('#delid').addClass('checkboxdelete')
})


// status model content updations

$(document).on('click', '#unbulishslt', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text($(this).text() + " " + languagedata.Memberss.memberstitle)

        $('#content').text(languagedata.Memberss.memberstatuscontent + " " + $(this).text() + " " + languagedata.Memberss.selectedmembers)
    } else {

        $('.deltitle').text($(this).text() + " " + languagedata.Memberss.membertitle)

        $('#content').text(languagedata.Memberss.memberstatuscontent + " " + $(this).text() + " " + languagedata.Memberss.selectedmember)
    }
    $("#delid").text($(this).text());

    $('#delid').addClass('selectedunpublish')

})


// remove datas when delete model close

$("#deleteModal").on("hide.bs.modal", function () {
    $('#delid').removeClass('checkboxdelete');
    $('#delid').removeClass('selectedunpublish');
    $('#delid').attr('href', '');
    $('.delname').text("");
})


//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {

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
            "page": pageno

        },
        success: function (data) {

            console.log(data, "result")

            if (data.value == true) {

                setCookie("get-toast", "Member Deleted Successfully")

                window.location.href = data.url
            } else {

                // setCookie("Alert-msg", "Internal Server Error")

                window.location.href = data.url

            }

        }
    })

})


//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass("hidden")

})


//multi select active and deactive function//

$(document).on('click', '.selectedunpublish', function () {

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
            "memberids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {

            console.log(data, "result")

            if (data.value == true) {

                setCookie("get-toast", "memstatusnotify")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})


$('#profile-slugData').on('input', function (event) {

    var inputvalue = $(this).val()

    var smallcase = inputvalue.toLowerCase()

    var modifiedText = smallcase.replace(/ /g, '-');

    $(this).val(modifiedText)
})



$(document).on('click', '.claimupdate', function () {

    var memberid = $(this).attr('data-id')

    console.log("claim mem id", memberid);

    $("#memberid").val(memberid)

    GetMemberDetails(memberid)

})


function GetMemberDetails(memberid) {
    console.log("work");

    $.ajax({
        url: '/member/getmemberdetails',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "memberid": memberid,
            csrf: $("input[name='csrf']").val()
        },
        success: function (data) {
            console.log("sjdfk,jsd");
            console.log(data, "result")

            if (data.flg == false) {

                setCookie("Alert-msg", "Internal Server Error")
                window.location.href = data.url
            } else {

                $("#claim-companyname").text(data.memberprofile.CompanyName);
                $("#claim-companylocation").text(data.memberprofile.CompanyLocation);
                $("#claim-profilename").text(data.memberprofile.ProfileName);
                $("#claim-profileslug").text(data.memberprofile.ProfileSlug);
                $("#claim-email").text(data.member.Email);
                $("#claim-number").text(data.member.MobileNo);
                $("#claim-dateofclaim").text(data.memberprofile.Website)
                $("#claim-companylogo").attr('src', data.memberprofile.CompanyLogo)
                if (data.memberprofile.ClaimStatus == 1) {
                    $("#claimb1").prop('checked', true)
                    $("#claim-heading").text("Already Claimed")
                    // $("#updateclaim").attr('disabled',true).css('opacity','0.5')
                } else {
                    $("#claimb1").prop('checked', false)
                    $("#claim-heading").text("Activate Claim")
                    // $("#updateclaim").attr('disabled',false).css('opacity','1')
                    // $('#updateclaim').attr('data-bs-dismiss','')
                }
            }
        }
    })
}




//profile slug function
$('#companyname').keyup(function () {

    var companyName = $(this).val();

    var cleanedCompanyName = companyName.replace(/[^a-zA-Z0-9\s]/g, '');

    $(this).val(cleanedCompanyName);

    var profileSlug = cleanedCompanyName.trim().replace(/\s+/g, '-').toLowerCase();

    $('#profile-slugData').val(profileSlug);

});


// first name and last name limit of 25 char
$(document).on('keyup', '.checklength', function () {

    var inputVal = $(this).val()

    var inputLength = inputVal.length

    if (inputLength == 25) {
        $(this).siblings('.lengthErr').removeClass('hidden')
    } else {
        $(this).siblings('.lengthErr').addClass('hidden')
    }
})