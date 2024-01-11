var languagedata
/** */
$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    $.getJSON("/locales/" + languagecode + ".json", function (data) {

        languagedata = data
    })
})

$(document).ready(function () {
    // $('.nav-link.Members').parent().addClass("active");
    $('.Members').addClass('checked');

    // $.validator.addMethod("fname_validator", function (value) {
    //     return /^[\S+(?: \S+)]{3,15}$/.test(value);
    // }, '* Please Enter a Valid Name.');

    // $.validator.addMethod("username_validator", function (value) {
    //     return /^[\S+(?: \S+)]{3,15}$/.test(value);
    // }, ' * Please Enter a Valid Name.');

    // $.validator.addMethod("mob_validator", function (value) {
    //     return /^[6-9]{1}[0-9]{9}$/.test(value);
    // }, '* ' + languagedata.Userss.usrmobnumrgx);

    // jQuery.validator.addMethod("mob_validator", function (value) {
    //     return /^([0|\+[0-9]{1,5})?([6-9]{1}[0-9]{9})$/.test(value)
    // }, "*" + languagedata.Memberss.memmobvailderr);

    // $.validator.addMethod("email_validator", function (value) {
    //     return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    // }, '* ' + languagedata.Memberss.mememailrgx);

    // jQuery.validator.addMethod("pass_validator", function (value, element) {
    //     if (value != "") {
    //         if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
    //             return true;
    //         else return false;
    //     }
    //     else return true;
    // }, "* " + languagedata.Memberss.mempswdrgx
    // );

});

$(document).on('click', '#crop-button', function () {
    $(".name-string").hide()
    // $('#profpic').show()
})

$("#addmember , #clickadd").click(function () {
    $("#title").text(languagedata.Memberss.addmember)
    $("#triggerId").text(languagedata.Memberss.choosegroup)

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
    $("#mem_mobile-error").hide()
    $("#mem_usrname-error").hide()
    // $("#mem_activestat-error").hide()
    $("#membergroupvalue-error").hide()
    $("#memberimg-error").hide()
    $("#mem_pass-error").hide()
    $(".input-group").removeClass("input-group-error")
    $('.update-profile-img>img').attr('src', '/public/img/profile.png')
    $('input[type=hidden][name=crop_data]').val("")
    $(".name-string").hide()


    $("#save").show()
    $("#update").hide()

})

// only allow numbers
// $('#mem_mobile').keyup(function () {
//     this.value = this.value.replace(/[^0-9\.]/g, '');
// });

// delete btn
$(document).on('click', '#del', function () {
    var MemberId = $(this).attr("data-id")
    $(".deltitle").text(languagedata.Memberss.deltitle)
    $('.delname').text($(this).parents('tr').find('td:eq(0)').text())
    $("#content").text(languagedata.Memberss.delmember)
    $('#delete').attr('href', '/member/deletemember?id=' + MemberId);
})

$(document).on("click", "#editmembergroup", function () {
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');
    console.log(pageno,"chech")
    if (pageno == null) {
        pageno = "1";
    }
    $("#memgrbpageno").val(pageno)
    // var memberId = $(this).data("id");
    var editUrl = $(this).attr("href");
    window.location.href = editUrl + "&page=" + pageno;
});

// editbtn
var edit
$(document).on('click', '#editmem', function () {

    $(".input-group").removeClass("input-group-error")

    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')

    $("#memgrbpageno").val(pageno)

    var data = $(this).attr("data-id");
    console.log("id", data);
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
        url: "/member/editmember",
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
                $('#profpic').attr('src',result.Member.ProfileImagePath.replace(/^/, '/')).show();
                $(".name-string").hide()
            } else {
                $('#profpic').hide()
                $(".name-string").text(result.Member.NameString).show()
            }
          

            $("#triggerId").val(result.Group.Name)

            var res = $("#mem_activestat").val(result.Member.IsActive)

            if (res == 1) {
                $("#mem_activestat").prop("checked", true)

            } else {
                $("#mem_activestat").prop("", false)

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
    

    jQuery.validator.addMethod("mob_validator", function (value) {
        return /^([0|\+[0-9]{1,5})?([6-9]{1}[0-9]{9})$/.test(value)
    }, "*" + languagedata.Memberss.memmobvailderr);

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
            membergroupvalue: {
                required: true,
            },
            mem_usrname: {
                required: true,
                //  username_validator: true,
                space: true,
                duplicatename: true
            },
            // mem_pass: {
            //     // required: true,
            //     pass_validator: true,
            // },
            mem_mobile: {
                required: true,
                mob_validator: true,
                duplicatenumber: true
                // number: true,

            },
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
            membergroupvalue: {
                required: "* " + languagedata.memgroup
            },
            mem_usrname: {
                required: "* " + languagedata.Memberss.memname,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Memberss.membernamevaild
            },
            // mem_pass: {
            //     required: "* "+languagedata.Userss.usrpswd
            // },
            mem_mobile: {
                required: "* " + languagedata.Memberss.memmobnum,
                duplicatenumber: "* " + languagedata.Memberss.memmobnumexist
            },
            // mem_activestat: {
            //     required: "* Please checkin the active toggle button",
            // }
        }
    })

    $('input[name=mem_pass]').rules('remove', 'required')

    // $("#newmember").valid();
    // if ($('#newmember').valid()) {
    //     $("#password").removeClass('.error').addClass("valid");
    //     $('#newmember')[0].submit();
    // }

    var formcheck = $("#memberform").valid();
    if (formcheck == true) {
        $('#memberform')[0].submit();
    } else {
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        $('.input-group').each(function () {
            var inputField = $(this).find('input');
            if (!inputField.valid() || inputField.val() === "") {
                $(this).addClass('validate');
            }
        });
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

    jQuery.validator.addMethod("mob_validator", function (value) {
        return /^([0|\+[0-9]{1,5})?([6-9]{1}[0-9]{9})$/.test(value)
    }, "*" + languagedata.Memberss.memmobvailderr);

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
            membergroupvalue: {
                required: true,
            },
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
            mem_mobile: {
                required: true,
                mob_validator: true,
                duplicatenumber: true
                // number: true,

            },
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
            membergroupvalue: {
                required: "* " + languagedata.memgroup
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
                required: "* " + languagedata.Memberss.memmobnum,
                duplicatenumber: "* " + languagedata.Memberss.memmobnumexist
            },
            // mem_activestat: {
            //     required: "* Please checkin the active toggle button",
            // }
        }
    })

    var formcheck = $("#memberform").valid();
    if (formcheck == true) {
        $('#memberform')[0].submit();
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
    $(this).parent().siblings("a").text(text)
    $(this).parent().siblings("input").val(id)
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

/*search */
$(document).on("click", "#filterformsubmit1", function () {
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
    console.log("ssss")
    if ($('#valmem').val() !== '') {
        $(this).parents('.input-group').removeClass('validate')
    }
    $(this).parents('.input-group').removeClass('focus')

})

function Validationcheck() {
    let inputGro = document.querySelectorAll('.input-group');
    inputGro.forEach(inputGroup => {
        let inputField = inputGroup.querySelector('input:not([name="membergroupvalue"])');
        var inputName = inputField.getAttribute('name');
        console.log("input", inputName)

        if (inputField.classList.contains('error')) {
            inputGroup.classList.add('input-group-error');
        } else {
            inputGroup.classList.remove('input-group-error');
        }
        if ($('#membergroupvalue-error').css('display') !== 'none') {
            console.log("check")
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
    console.log(pageno, "chech")
    if (pageno == null) {
        pageno = "1";
    }
    $("#memgrbpageno").val(pageno)
    // var memberId = $(this).data("id");
    var editUrl = $(this).attr("href");
    window.location.href = editUrl + "&page=" + pageno;
});

$("#searchrole").keyup(function () {
    console.log("check")
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

$(document).on('keyup', '#searchroles', function () {

    if ($('.search').val() === "") {
        window.location.href = "/member"

    }

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

$('form[class=filterform]>img').click(function () {

    if ($(this).siblings('input[name=keyword]').val() != "" && $(this).parents('.transitionSearch').hasClass('active')) {

        var keyword = $(this).siblings('input[name=keyword]').val()

        window.location.href = "/membersgroup/?keyword=" + keyword
    }
})

$('#memModal').on('hide.bs.modal', function () {

    $('.input-field-group .input-group').each(function () {

        if ($(this).hasClass('input-group-error')) {

            $(this).removeClass('input-group-error')
        }
    })

    $('label[class=error').removeClass('error').hide()

    $("#memberform").attr("action", "/member/newmember");

    $("#memberform").attr("name", "createmember");

    $('#mem_pass-error').remove()

    $('input[name=mem_activestat]').remove()

    $('<input class="tgl tgl-light" name="mem_activestat" id="mem_activestat" type="checkbox">').insertBefore('label[for=mem_activestat]')

    $('input[type=hidden][name=membergroupvalue]').val("")

})

