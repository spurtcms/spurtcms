var languagedata

$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })

})

$("input[name=mem_activestat]").click(function () {
    if ($(this).prop('checked') == true) {
        $(this).attr("value", "1")
    } else {
        $(this).attr("value", "0")
    }
})

$("input[name=mem_data_access]").click(function () {
    if ($(this).prop('checked')) {
        $(this).attr("value", "1")
    } else {
        $(this).attr("value", "0")
    }
})

//**Dropdown role//
$(document).on('click', '.dropdown-item', function () {

    role = $(this).text()
    roleid = $(this).attr('data-id')
    $('#triggerId').text(role)
    $('#rolen').val(roleid)
    if ($('#rolen').val() !== '') {
        $('#rolen-error').hide()
        $('.user-drop-down').removeClass('input-group-error')

    }
})

/* Create User */
$("#saveuser").click(function () {
    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {
            if (/(^[a-zA-Z_0-9\.-]+)@([a-z]{2,})\.([a-z]{2,})(\.[a-z]{2,})?$/.test(value))
                return true;
            else return false;
        },
        "* " + languagedata.Userss.usremailrgx
    );

    // jQuery.validator.addMethod(
    //     "username_validator",
    //     function (value, element) {
    //         if (/^[\S+(?: \S+)]{3,15}$/.test(value))
    //             return true;
    //         else return false;
    //     },
    //     "* "+languagedata.usrnamergx
    // );

    jQuery.validator.addMethod(
        "pass_validator",
        function (value, element) {
            if (value != "") {
                if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                    return true;
                else return false;
            }
            else return true;
        },
        "* " + languagedata.Userss.usrpswdrgx
    );

    jQuery.validator.addMethod(
        "mob_validator",
        function (value, element) {
            if (/^[6-9]{1}[0-9]{9}$/.test(value))
                return true;
            else return false;
        },
        "* " + languagedata.Userss.usrmobnumrgx
    );
    $("form[name='userform']").validate({

        ignore: [],
        rules: {
            prof_pic: {

                extension: "jpg|png|jpeg"
            },
            user_fname: {
                required: true,
                space: true,
            },

            user_email: {
                required: true,
                email_validator: true,
                // duplicateemail:true
            },
            user_role: {
                required: true,
            },
            user_name: {
                required: true,
                // duplicateusername:true,
                space: true,
            },
            user_pass: {
                required: true,
                pass_validator: true,
            },
            user_mob: {
                required: true,
                mob_validator: true,
                // duplicatenumber:true
            },


        },
        messages: {
            prof_pic: {
                extension: "* " + languagedata.profextension
            },
            user_fname: {
                required: "* " + languagedata.Userss.usrfname,
                space: "* " + languagedata.spacergx,
            },
            user_email: {
                required: "* " + languagedata.Userss.usrmail,
                // duplicateemail:"* "+languagedata.Userss.emailexist
            },
            user_role: {
                required: "* " + languagedata.Userss.usrrole
            },
            user_name: {
                required: "* " + languagedata.Userss.usrname,
                // duplicateusername:"*"+languagedata.Memberss.membernamevaild,
                space: "* " + languagedata.spacergx,
            },
            user_pass: {
                required: "* " + languagedata.Userss.usrpswd
            },
            user_mob: {
                required: "* " + languagedata.Userss.usrmobnum,
                // duplicatenumber:"* "+languagedata.Userss.mobnumexist,
            },



        }
    })
    var formcheck = $("#userform").valid();
    if (formcheck == true) {
        var email = $("#user_email").val()
        var mob = $("#user_mob").val()
        var uname = $('#user_name').val()
        var user_id = $("#userid").val()
        $.ajax({
            url: "/settings/users/checkuserdata",
            type: "POST",
            async: false,
            data: { "email": email, "mobile": mob, "username": uname, "id": user_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                console.log("data", data);
                if (data.email == true) {

                    $('#user_email-error').text("* "+languagedata.Userss.emailexist).show();
                    $('#emailgrp').addClass('input-group-error');
                }
                if (data.number == true) {
                    $('#user_mob-error').text("* "+languagedata.Userss.mobnumexist).show();
                    $('#mobilegrp').addClass('input-group-error');
                }
                if (data.user == true) {
                    $('#user_name-error').text("* "+languagedata.Userss.nameexist).show();
                    $('#usergrp').addClass('input-group-error');
                }
                if (data.user == false && data.email == false && data.number == false) {
                    $('#userform')[0].submit();
                }

            }
        })
    } else {
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        $('.input-group').each(function () {
            var inputField = $(this).find('input');
            var inputName = inputField.attr('name');


            if (inputName !== 'user_role' && !inputField.valid()) {
                $(this).addClass('input-group-error');

            } else {
                $(this).removeClass('input-group-error');
            }
        });
        if ($('#rolen-error').css('display') !== 'none') {
            $('.user-drop-down').addClass('input-group-error')
        } else {
            $('.user-drop-down').removeClass('input-group-error')
        }

    }
    return false;
});
$(document).on('click', '.cancel', function () {
    $('#changepicModal').modal('hide');
})
$(document).on('click', '.btn-close', function () {
    $('#changepicModal').modal('hide');
})

$(document).on('click', '#adduser', function () {
    $('#heading').text(languagedata.Userss.addnewuser)
    $(".name-string").hide()
    $('#profpic').attr('src', '/public/img/default profile .svg').show()
    $('#saveuser').show()
    $('#updateuser').hide()
})
$(document).on('click', '#clickadd', function () {
    $('#heading').text(languagedata.Userss.addnewuser)
    $('#saveuser').show()
    $('#updateuser').hide()
})
// ** edituser*//
$(document).on('click', '#edit-btn', function () {
    $('#saveuser').hide()
    $('#updateuser').show()
    var data = $(this).attr("data-id");
    $('#heading').text(languagedata.Userss.updateuser)
    $("#userModal").show()
    $("#userModal").attr("action", "");
    $("#userModal").attr("name", "edit");
    $('#userform').attr("action", "/settings/users/update-user")
    $('#userform').attr("name", "editform")
    $('#userid').val(data)
    $.ajax({
        url: "/settings/users/edit-user",
        type: "GET",
        dataType: "json",
        data: { "id": data },
        success: function (result) {
            console.log("user",result);
            if (result.Id != 0) {
                console.log("result", result)
                var fname = $("#user_fname").val(result.FirstName);
                var lname = $("#user_lname").val(result.LastName);
                var email = $("#user_email").val(result.Email)
                var mob = $("#user_mob").val(result.MobileNo)
                var uname = $('#user_name').val(result.Username)
                var role = $('#rolen').val(result.RoleId)
                var isactive = $('#cb1').val(result.IsActive)
                var dataacc = $('#chk1').val(result.DataAccess)
                $("#user_fname").text(fname);
                $("#user_lname").text(lname);
                $("#user_email").text(email)
                $("#user_mob").text(mob)
                $('#user_name').text(uname)
                $('#rolen').text(role)
                $('#rolen').siblings('.dropdown-menu').find('button[data-id="' + result.RoleId + '"]').click();
                if (result.ProfileImagePath != "") {
                    $('#profpic').attr('src', result.ProfileImagePath.replace(/^/, '/')).show();
                    $(".name-string").hide()
                } else {
                    $('#profpic').hide()
                    $(".name-string").text(result.NameString).show()
                }
                $('.tgl-btn').val(isactive)
                $('.chk1').val(dataacc)
                if ($('#cb1').val() == 1) {
                    console.log("box")
                    $('input[name=mem_activestat]').prop('checked', true)
                }
                if ($('#chk1').val() == 1) {
                    $('input[name=mem_data_access]').prop('checked', true)

                }

            }
        }
    })
})
/* Update User */
$(document).on('click', '#updateuser', function () {
    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {
            if (/(^[a-zA-Z_0-9\.-]+)@([a-z]{2,})\.([a-z]{2,})(\.[a-z]{2,})?$/.test(value))
                return true;
            else return false;
        },
        "* " + languagedata.Userss.usremailrgx
    );

    // jQuery.validator.addMethod(
    //     "username_validator",
    //     function (value, element) {
    //         if (/^[\S+(?: \S+)]{3,15}$/.test(value))
    //             return true;
    //         else return false;
    //     },
    //     "* "+languagedata.usrnamergx
    // );

    jQuery.validator.addMethod(
        "pass_validator",
        function (value, element) {
            if (value != "") {
                if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                    return true;
                else return false;
            }
            else return true;
        },
        "* " + languagedata.Userss.usrpswdrgx
    );

    jQuery.validator.addMethod(
        "mob_validator",
        function (value, element) {
            if (/^[6-9]{1}[0-9]{9}$/.test(value))
                return true;
            else return false;
        },
        "* " + languagedata.Userss.usrmobnumrgx
    );
    $("form[name='editform']").validate({

        ignore: [],
        rules: {
            prof_pic: {

                extension: "jpg|png|jpeg"
            },
            user_fname: {
                required: true,
                space: true,
            },

            user_email: {
                required: true,
                email_validator: true,
                // duplicateemail: true
            },
            // user_role: {
            //     required: true,
            // },
            user_name: {
                required: true,
                // duplicateusername: true,
                space: true,
            },
            user_pass: {
                // required: true,
                pass_validator: true,
            },
            user_mob: {
                required: true,
                mob_validator: true,
                // duplicatenumber: true
            },


        },
        messages: {
            prof_pic: {
                extension: "* " + languagedata.profextension
            },
            user_fname: {
                required: "* " + languagedata.Userss.usrfname,
                space: "* " + languagedata.spacergx,
            },
            user_email: {
                required: "* " + languagedata.Userss.usrmail,
                // duplicateemail: "* " + languagedata.Userss.emailexist
            },
            // user_role: {
            //     required: "* " + languagedata.Userss.usrrole
            // },
            user_name: {
                required: "* " + languagedata.Userss.usrname,
                // duplicateusername: "*" + languagedata.Memberss.membernamevaild,
                space: "* " + languagedata.spacergx,
            },
            // user_pass: {
            //     required: "* " + languagedata.Userss.usrpswd
            // },
            user_mob: {
                required: "* " + languagedata.Userss.usrmobnum,
                // duplicatenumber: "* " + languagedata.Userss.mobnumexist,
            },


        }


    })
    $('input[name=user_pass]').rules('remove', 'required')
    var formcheck = $("#userform").valid();
    if (formcheck == true) {
        var email = $("#user_email").val()
        var mob = $("#user_mob").val()
        var uname = $('#user_name').val()
        console.log("dfgh", email, mob, uname);
        var user_id = $("#userid").val()
        $.ajax({
            url: "/settings/users/checkuserdata",
            type: "POST",
            async: false,
            data: { "email": email, "mobile": mob, "username": uname, "id": user_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                if (data.email == true) {

                    $('#user_email-error').text("* "+languagedata.Userss.emailexist).show();
                    $('#emailgrp').addClass('input-group-error');
                }
                if (data.number == true) {
                    $('#user_mob-error').text("* "+languagedata.Userss.mobnumexist).show();
                    $('#mobilegrp').addClass('input-group-error');
                }
                if (data.user == true) {
                    $('#user_name-error').text("* "+languagedata.Userss.nameexist).show();
                    $('#usergrp').addClass('input-group-error');
                }
                if (data.user == false && data.email == false && data.number == false) {
                    $('#userform')[0].submit();
                }

            }
        })
    } else {
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        Validationcheck()
        // $('.input-group').each(function () {
        //     var inputField = $(this).find('input');
        //     var inputName = inputField.attr('name');


        //     if (!inputField.valid()) {
        //         $(this).addClass('input-group-error');

        //     } else {
        //         $(this).removeClass('input-group-error');
        //     }
        // });
        // if ($('#rolen-error').css('display') !=='none'){
        //     console.log("check")
        //         $('.user-drop-down').addClass('input-group-error')
        //     }

    }
    return false;
})

$(document).on('click', '#del-btn', function () {

    var userId = $(this).attr("data-id")
    $("#content").text(languagedata.Userss.deluser)
    $("#delete").attr("href", "/settings/users/delete-user/" + userId)
})

/*search */
$(document).on("click", "#filterformsubmit", function () {
    var key = $(this).siblings().children(".search").val();
    if (key == "") {
        window.location.href = "/settings/users/"
    } else {
        $('.filterform').submit();
    }
})

$('input[name=user_mob]').keyup(function () {
    this.value = this.value.replace(/[^0-9\.]/g, '');
});

$(document).on('click', '.newck-group', function () {
    var role = $(this).find('label>h4').text()
    var roleId = $(this).attr('data-id')
    if (roleId != '') {
        $('input[name=mem_role]').val(roleId)
        $('#mem_role-error').hide()
    } else {
        $('input[name=mem_role]').val("")
    }
    $('.newselect .dd-a > span').text(role)

    if ($("#mem_role") !== "") {
        $(this).parents('.dd-c').css('display', 'none')
    }
    $("#searchrole").val("");
    // $('#newdd-input').prop('checked',false)
    // $(this).parents('.dd-c').css('display','none')  
    refreshdiv()
})

$(document).on('click', '#newdd-input', function () {
    $(this).siblings('.dd-c').css('display', 'block')
    $("#searchrole").val("");
})

$('.page-wrapper').on('click', function (event) {
    if ($('.dd-c').css('visibility') == 'visible' && !$(event.target).is('#newdd-input') && !$(event.target).is('.dd-c')) {
        // $('#newdd-input').prop('checked',false)
        // $('.dd-c').css('display','none')
    }
})

// $(document).on('click', '#Eye', function () {

//     var pass_field = $(this).parent().siblings('input')

//     if ($(pass_field).attr('type') == 'password') {

//         $(pass_field).attr('type', 'text')

//         $(this).removeClass('fa-eye-slash').addClass('fa-eye')

//     } else {

//         $(pass_field).attr('type', 'password')

//         $(this).removeClass('fa-eye').addClass('fa-eye-slash')

//     }

// })



$('#triggerId').blur(function () {
    if ($('#rolen').val() !== '') {
        $('.user-drop-down').removeClass('input-group-error')
        $('#rolen-error').hide()
    }
    $(this).parents('.input-group').removeClass('focus')

})

function Validationcheck() {
    let inputGro = document.querySelectorAll('.input-group');
    inputGro.forEach(inputGroup => {
        let inputField = inputGroup.querySelector('input:not([name="user_role"])');
        var inputName = inputField.getAttribute('name');


        if (inputField.classList.contains('error')) {
            inputGroup.classList.add('input-group-error');
        } else {
            inputGroup.classList.remove('input-group-error');
        }

        if ($('#rolen-error').css('display') !== 'none') {
            $('.user-drop-down').addClass('input-group-error')
        }
        else {
            $('.user-drop-down').removeClass('input-group-error')
        }
    });
}
$(document).on("click", "#edituser", function () {

    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');

    if (pageno == null) {
        pageno = "1";
    }
    $("#userpageno").val(pageno)

    var editUrl = $(this).attr("href");

    window.location.href = editUrl + "?page=" + pageno;
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

$(document).on('keyup', '#searchroles', function () {

    if ($('.search').val() === "") {
        window.location.href = "/settings/users/"

    }

})

//**close model */

$(document).on('click', '.close', function () {
    $('#userform')[0].reset();
    $('label.error').remove()
    $('.input-group').removeClass('input-group-error')
    $('.field').removeClass('error')
    $('#triggerId').text('Select Role')
    // $('#profpic').attr('src', '')


})

// $(document).on('click', '#triggerId', function() {
//     console.log("check");
//     $(this).closest('.input-group').addClass('input-group-focused');
//   });


// **eye open **//

// $(document).on('click', '.eye', function () {
//     $(this).children('img').attr('src', "/public/img/eye-opened.svg")
//     $(this).siblings('input').attr('type', "text")
// })
$(document).on('click', '#crop-button', function () {
    $(".name-string").hide()
    $('#profpic').show()
})
$(document).on('click', '#eye', function () {

    var This = $("#user_pass")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');
        

    } else {

        $(This).attr('type', 'password');

        $(this).find('img').attr('src', '/public/img/eye-closed.svg');
    }
})