var languagedata

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

})

// search focus function

$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
});


// status indication

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

$(document).on('click', '.dropdown-items', function () {
    role = $(this).text()
    roleid = $(this).attr('data-id')
    $('#showgroup').text(role)
    $('#showgroup').removeClass('text-bold-gray');
    $('#showgroup').addClass('text-bold');
    $('#rolen').val(roleid)
    if ($('#rolen').val() !== '') {
        $('#rolen-error').hide()
        $('.user-drop-down').removeClass('input-group-error')

    }
})


/* Create User */

$("#saveuser").click(function () {
    console.log($(this).text().trim(),"text lang save",languagedata.save);

    if ($(this).text().trim() == languagedata.save) {

        jQuery.validator.addMethod(
            "email_validator",
            function (value, element) {
                if (/(^[a-zA-Z_0-9\.-]+)@([a-z]{2,})\.([a-z]{2,})(\.[a-z]{2,})?$/.test(value))
                    return true;
                else return false;
            },
            "* " + languagedata.Userss.usremailrgx
        );

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

        jQuery.validator.addMethod("duplicateemail", function (value) {

            var result;
            user_id = $("#userid").val()
            $.ajax({
                url: "/settings/users/checkemail",
                type: "POST",
                async: false,
                data: { "email": value, "id": user_id, csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    console.log(data,"email");
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })

        // jQuery.validator.addMethod("duplicateusername", function (value) {

            // var result;
            // user_id = $("#userid").val()
            // $.ajax({
            //     url: "/settings/users/checkusername",
            //     type: "POST",
            //     async: false,
            //     data: { "username": value, "id": user_id, csrf: $("input[name='csrf']").val() },
            //     datatype: "json",
            //     caches: false,
            //     success: function (data) {
            //         result = data.trim();
            //     }
            // })
            // return result.trim() != "true"

        // })

        jQuery.validator.addMethod("duplicatenumber", function (value) {

            var result;
            user_id = $("#userid").val()
            $.ajax({
                url: "/settings/users/checknumber",
                type: "POST",
                async: false,
                data: { "number": value, "id": user_id, csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })

        $("form[name='userform']").validate({

            ignore: [],
            rules: {
                prof_pics: {

                    extension: "jpg|png|jpeg|svg"
                },
                user_fname: {
                    required: true,
                    space: true,
                },

                user_email: {
                    required: true,
                    email_validator: true,
                    duplicateemail: true
                },
                user_role: {
                    required: true,
                },
                user_name: {
                    required: true,
                    // duplicateusername: true,
                    space: true,
                },
                user_pass: {
                    required: true,
                    pass_validator: true,
                },
                user_mob: {
                    required: true,
                    mob_validator: true,
                    duplicatenumber: true
                },


            },
            messages: {
                prof_pics: {
                    extension: "* " + languagedata.profextension
                },
                user_fname: {
                    required: "* " + languagedata.Userss.usrfname,
                    space: "* " + languagedata.spacergx,
                },
                user_email: {
                    required: "* " + languagedata.Userss.usrmail,
                    duplicateemail: "* " + languagedata.Userss.emailexist
                },
                user_role: {
                    required: "* " + languagedata.Userss.usrrole
                },
                user_name: {
                    required: "* " + languagedata.Userss.usrname,
                    // duplicateusername: "*" + languagedata.Userss.nameexist,
                    space: "* " + languagedata.spacergx,
                },
                user_pass: {
                    required: "* " + languagedata.Userss.usrpswd
                },
                user_mob: {
                    required: "* " + languagedata.Userss.usrmobnum,
                    duplicatenumber: "* " + languagedata.Userss.mobnumexist,
                },



            }
        })


        var formcheck = $("#userform").valid();
        console.log(formcheck,"validat");
        if (formcheck == true) {
            $('#userform')[0].submit();
            $('#saveuser').prop('disabled', true);
            var email = $("#user_email").val()
            var mob = $("#user_mob").val()
            var uname = $('#user_name').val()
            var user_id = $("#userid").val()
            // $.ajax({
            //     url: "/settings/users/checkuserdata",
            //     type: "POST",
            //     async: false,
            //     data: { "email": email, "mobile": mob, "username": uname, "id": user_id, csrf: $("input[name='csrf']").val() },
            //     datatype: "json",
            //     caches: false,
            //     success: function (data) {
            //         console.log("data", data);
            //         if (data.email == true) {

            //             $('#user_email-error').text("* " + languagedata.Userss.emailexist).show();
            //             $('#emailgrp').addClass('input-group-error');
            //         }

            //         if (data.number == true) {
            //             $('#user_mob-error').text("* " + languagedata.Userss.mobnumexist).show();
            //             $('#mobilegrp').addClass('input-group-error');
            //         }
            //         if (data.user == true) {
            //             $('#user_name-error').text("* " + languagedata.Userss.nameexist).show();
            //             $('#usergrp').addClass('input-group-error');
            //         }
            //         if (data.user == false && data.email == false && data.number == false) {
            //             $('#userform')[0].submit();
            //             $('#saveuser').prop('disabled', true);
            //         }

            //     }
            // })

        } else {
            $('#saveuser').prop('disabled', false);
            $(document).on('keyup', ".field", function () {
                Validationcheck()
            })
            $('.input-group').each(function () {
                var inputField = $(this).find('input');
                var inputName = inputField.attr('name');


                if (!inputField.valid()) {
                    $(this).addClass('input-group-error');

                } else {
                    $(this).removeClass('input-group-error');
                }
            });
            // if ($('#rolen-error').css('display') !== 'none') {
            //     $('.user-drop-down').removeClass('input-group-error')
            // }

        }



        return false;
    }
});



$(document).on('click', '.cancel', function () {
    $('#changepicModal').modal('hide');
})
$(document).on('click', '.btn-close', function () {
    $('#changepicModal').modal('hide');
})

// $(document).on('click', '#adduser', function () {
//     $('#heading').text(languagedata.Userss.addnewuser)
//     $(".name-string").hide()
//     $('#profpic').attr('src', '/public/img/default profile .svg').show()
//     $('#saveuser').show()
//     $('#updateuser').hide();
//     $("#userid").val("");
//     $('#userform').attr('action', '/settings/users/createuser')
// })
// $(document).on('click', '#clickadd', function () {
//     $('#heading').text(languagedata.Userss.addnewuser)
//     $('#saveuser').show()
//     $('#updateuser').hide()
// })


// ** edituser*//

$(document).on('click', '#edit-btn', function () {

    // paganation get the pagenumber
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');

    $("#pageno").val(pageno)

    $('#saveuser').text(languagedata.update)
    var data = $(this).attr("data-id");
    $('#addmember-title').text(languagedata.Userss.updateuser)
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
            console.log("user", result);
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
                $('#rolen').text(result.RoleId)
                $('#rolen').siblings('.dropdown-menu').find('button[data-id="' + result.RoleId + '"]').click();
                if (result.ProfileImagePath != "") {
                    console.log(result.ProfileImagePath);
                    $('#profpic-user').attr('src', result.ProfileImagePath).show();
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


// remove data form add user input (cancel btn ckick function)

$("#usermodelclose").on("click", function () {
    $('#addmember-title').text(languagedata.Userss.addnewuser)
    $("#profpic-user").attr('src', '/public/img/defaultprofile .svg');
    $("#user_fname").val("");
    $("#user_lname").val("");
    $("#user_email").val("");
    $("#user_mob").val("");
    $('#user_name').val("");
    $('#rolen').val("");
    var active = $("#cb1").is(":checked");
    if (active) {
        $("#cb1").prop("checked", false);
    }
    $('#showgroup').text("Select Role")
    $('#showgroup').addClass('text-bold-gray');
    $('#showgroup').removeClass('text-bold');
    $('#saveuser').text(languagedata.save)
    $("#userform").attr('name', 'userform');
    $("label.error").hide();
    $("#myfile-error").hide();
    $(".lengthErr").addClass("hidden");
    $('#userid').val("")
    $('#userform').attr("action", "/settings/users/createuser")
});


// delete model cancel function

$("#deleteModal").on("hide.bs.modal", function () {
    $('#delid').removeClass('checkboxdelete');
    $('#delid').removeClass('selectedIsActive');
    $('#delid').attr('href','');
    $('.delname').text('')
})


/* Update User */

$(document).on('click', '#saveuser', function () {
    if ($(this).text().trim() == languagedata.update) {

        jQuery.validator.addMethod("duplicateemail", function (value) {

            var result;
            user_id = $("#userid").val()
            $.ajax({
                url: "/settings/users/checkemail",
                type: "POST",
                async: false,
                data: { "email": value, "id": user_id, csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })


        // jQuery.validator.addMethod("duplicateusername", function (value) {

        //     var result;
        //     user_id = $("#userid").val()
        //     $.ajax({
        //         url: "/settings/users/checkusername",
        //         type: "POST",
        //         async: false,
        //         data: { "username": value, "id": user_id, csrf: $("input[name='csrf']").val() },
        //         datatype: "json",
        //         caches: false,
        //         success: function (data) {
        //             result = data.trim();
        //         }
        //     })
        //     return result.trim() != "true"
        // })


        jQuery.validator.addMethod("duplicatenumber", function (value) {

            var result;
            user_id = $("#userid").val()
            $.ajax({
                url: "/settings/users/checknumber",
                type: "POST",
                async: false,
                data: { "number": value, "id": user_id, csrf: $("input[name='csrf']").val() },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                }
            })
            return result.trim() != "true"
        })
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
                prof_pics: {

                    extension: "jpg|png|jpeg|svg"
                },
                user_fname: {
                    required: true,
                    space: true,
                },

                user_email: {
                    required: true,
                    email_validator: true,
                    duplicateemail: true
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
                    duplicatenumber: true
                },


            },
            messages: {
                prof_pics: {
                    extension: "* " + languagedata.profextension
                },
                user_fname: {
                    required: "* " + languagedata.Userss.usrfname,
                    space: "* " + languagedata.spacergx,
                },
                user_email: {
                    required: "* " + languagedata.Userss.usrmail,
                    duplicateemail: "* " + languagedata.Userss.emailexist
                },
                // user_role: {
                //     required: "* " + languagedata.Userss.usrrole
                // },
                user_name: {
                    required: "* " + languagedata.Userss.usrname,
                    // duplicateusername: "*" + languagedata.Userss.nameexist,
                    space: "* " + languagedata.spacergx,
                },
                // user_pass: {
                //     required: "* " + languagedata.Userss.usrpswd
                // },
                user_mob: {
                    required: "* " + languagedata.Userss.usrmobnum,
                    duplicatenumber: "* " + languagedata.Userss.mobnumexist,
                },


            }


        })
        $('input[name=user_pass]').rules('remove', 'required')
        var formcheck = $("#userform").valid();
        if (formcheck == true) {

            $('#userform')[0].submit();
            var email = $("#user_email").val()
            var mob = $("#user_mob").val()
            var uname = $('#user_name').val()
            console.log("dfgh", email, mob, uname);
            var user_id = $("#userid").val()

            // $.ajax({
            //     url: "/settings/users/checkuserdata",
            //     type: "POST",
            //     async: false,
            //     data: { "email": email, "mobile": mob, "username": uname, "id": user_id, csrf: $("input[name='csrf']").val() },
            //     datatype: "json",
            //     caches: false,
            //     success: function (data) {
            //         if (data.email == true) {

            //             $('#user_email-error').text("* " + languagedata.Userss.emailexist).show();
            //             $('#emailgrp').addClass('input-group-error');
            //             $('#user_email').addClass('error')
            //         }
            //         if (data.number == true) {
            //             $('#user_mob-error').text("* " + languagedata.Userss.mobnumexist).show();
            //             $('#mobilegrp').addClass('input-group-error');
            //             $('#user_mob').addClass('error')
            //         }
            //         if (data.user == true) {
            //             $('#user_name-error').text("* " + languagedata.Userss.nameexist).show();
            //             $('#usergrp').addClass('input-group-error');
            //             $('#user_name').addClass('error')
            //         }
            //         if (data.user == false && data.email == false && data.number == false) {
            //             $('#userform')[0].submit();
            //         }

            //     }
            // })
        } else {
            $(document).on('keyup', ".field", function () {

                $('.user-drop-down').removeClass('input-group-error')

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

    }

})



// delete btn function

$(document).on('click', '#del-btn', function () {

    var userId = $(this).attr("data-id")
    $(".deltitle").text(languagedata.Userss.deleteuser)
    $("#content").text(languagedata.Userss.deluser)
    var del = $(this).parents('tr');
    $('.delname').text(del.find('#username').text())


    // paganation get the pagenumber
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');

    if (pageno == null) {
        $("#delid").attr("href", "/settings/users/delete-user/" + userId)

    } else {
        $("#delid").attr("href", "/settings/users/delete-user/" + userId + "?page=" + pageno)

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
        let inputField = inputGroup.querySelector('input');
        // var inputName = inputField.getAttribute('name');

        if (inputField.classList.contains('error')) {
            console.log("error")
            inputGroup.classList.add('input-group-error');
        } else {
            console.log("errorsssss")
            inputGroup.classList.remove('input-group-error');
        }

    });
    // if ($('#rolen-error').css('display') !== 'none') {
    //     $('.user-drop-down').addClass('input-group-error')
    // }
    // else {
    //     $('.user-drop-down').removeClass('input-group-error')
    // }
}

// edit user function

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


// search role

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


// search functions strat //


/*search redirect home page */

$(document).on('keyup', '#searchroles', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/settings/users/";
        }
    }

    $('.searchClosebtn').toggleClass('hidden', searchInput === "");

})


$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

  $(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/settings/users/"
  })

  $(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")
        } else {
            $(".Closebtn").addClass("hidden")
            $(".srchBtn-togg").removeClass("pointer-events-none")
        }
    });
})

  $(document).on("click", ".hovericon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
  })



//**close model */

$(document).on('click', '.close', function () {
    $('#userform')[0].reset();
    $('#userform').attr('name', "userform");
    $('label.error').remove()
    $('.input-group').removeClass('input-group-error')
    $('.field').removeClass('error')
    $('#triggerId').text('Select Role')
    $('#rolen').val('')
    $('#userid').val('')
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


// eye icon open and close function

$(document).on('click', '#eye', function () {
    console.log("working");

    var This = $("#user_pass")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');


    } else {

        $(This).attr('type', 'password');

        $(this).find('img').attr('src', '/public/img/eye-closed.svg');
    }
})


// dropdown filter input box search

$("#searchdropdownrole").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    console.log("keyword", keyword);
    $(".dropdown-filter-roless").each(function (index, element) {
        var title = $(element).text().toLowerCase()

        console.log("title", title);
        if (title.includes(keyword)) {
            $(element).show()
            $("#nodatafounddesign").addClass("hidden")

        } else {
            $(element).hide()
            if ($('.dropdown-filter-roles button:visible').length == 0) {

                $("#nodatafounddesign").removeClass("hidden")

            }
        }
    })

})


// search key value empty in dropdown search

$("#triggerId").on("click", function () {

    var keyword = $("#searchdropdownrole").val()
    $(".dropdown-role .dropdown-filter-roles button").each(function (index, element) {
        if (keyword != "") {
            $("#searchdropdownrole").val("")
            $("#nodatafounddesign").hide()
            $(element).show()
        } else {
            $("#searchdropdownrole").val()
            $("#nodatafounddesign").hide()
            $(element).show()

        }
    })

})


$(document).on('click', '#myfile', function () {
    $("#prof-crop").val("6")
})

// checkBox selection code below

var selectedcheckboxarr = []
var entryid

$(document).on('click', '.selectcheckbox', function () {

    $(".multidelete").addClass("responsive-width")

    entryid = $(this).attr('data-id')

    var status = $(this).parents('td').siblings('td').find('span.status').text().trim().replace(/\s+/g, ' ');

    var activeStatus = $(this).parents('td').siblings('td').find('.tgl-light').val();



    if ($(this).prop('checked')) {

        selectedcheckboxarr.push({ "entryid": entryid, "status": status, "activeStatus": activeStatus })

    } else {

        const index = selectedcheckboxarr.findIndex(item => item.entryid === entryid);

        if (index !== -1) {

            console.log(index, "sssss")
            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {
        $('.selected-numbers').removeClass("hidden")

       if (selectedcheckboxarr.length == 1){
            $('#deselectid').text(languagedata.deselect)
        }else if(selectedcheckboxarr.length > 1){
            $('#deselectid').text(languagedata.deselectall)
        }

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)


        var allActiveStatusSame = selectedcheckboxarr.every(function (item) {
            return item.activeStatus === selectedcheckboxarr[0].activeStatus;
        });




        var setActiveStatus
        var activeImg

        if (selectedcheckboxarr[0].activeStatus === '1') {

            setActiveStatus = languagedata.Userss.deactive;

            activeImg = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].activeStatus === '0') {

            setActiveStatus = languagedata.Userss.active;

            activeImg = "/public/img/Active.svg";

        }

        var activeStatusHtmlContent = '';

        if (allActiveStatusSame) {

            activeStatusHtmlContent ='<img style="width: 14px; height: 14px;" src="' + activeImg + '" >'  + '<span class="max-sm:hidden @[550px]:inline-block hidden">'+setActiveStatus+'</span>';


        } else {
            activeStatusHtmlContent = ''
        }

        $('#unbulishslt').html(activeStatusHtmlContent);

        if (!allActiveStatusSame) {
            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')


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

//ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click', '#Check', function () {

    $(".multidelete").addClass("responsive-width")

    selectedcheckboxarr = []

    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            entryid = $(this).attr('data-id')

            var status = $(this).parents('td').siblings('td').find('span.status').text().trim().replace(/\s+/g, ' ');

            var activeStatus = $(this).parents('td').siblings('td').find('.tgl-light').val();

            selectedcheckboxarr.push({ "entryid": entryid, "status": status, "activeStatus": activeStatus })
        })
        
        if (selectedcheckboxarr.length == 1){
            $('#deselectid').text(languagedata.deselect)
        }else if(selectedcheckboxarr.length > 1){
            $('#deselectid').text(languagedata.deselectall)
        }
        $('.selected-numbers').removeClass("hidden")

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)


        var allActiveStatusSame = selectedcheckboxarr.every(function (item) {

            return item.activeStatus === selectedcheckboxarr[0].activeStatus;
        });

        var setActiveStatus
        var activeImg


        if (selectedcheckboxarr.length != 0) {

            if (selectedcheckboxarr[0].activeStatus === '1') {

                setActiveStatus = languagedata.Userss.deactive;

                activeImg = "/public/img/In-Active.svg";

            } else if (selectedcheckboxarr[0].activeStatus === '0') {

                setActiveStatus = languagedata.Userss.active;

                activeImg = "/public/img/Active.svg";

            }
        }
        var activeStatusHtmlContent = '';

        if (allActiveStatusSame) {
            activeStatusHtmlContent ='<img style="width: 14px; height: 14px;" src="' + activeImg + '" >'  + '<span class="max-sm:hidden @[550px]:inline-block hidden">'+setActiveStatus+'</span>';

            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }
        else {
            activeStatusHtmlContent = "";
            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }

        $('#unbulishslt').html(activeStatusHtmlContent)

    } else {


        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selected-numbers').addClass("hidden")
    }

    if (selectedcheckboxarr.length == 0) {

        $('.selected-numbers').addClass("hidden")
    }
})


// delete model content update

$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length>1){

    $('.deltitle').text(languagedata.Userss.deleteusers)

    $('#content').text(languagedata.Userss.delmsg)
    }else{

    $('.deltitle').text(languagedata.Userss.deleteuser)

    $('#content').text(languagedata.Userss.delmsg)

    }

    $("#delid").text($(this).text());

    $('#delid').addClass('checkboxdelete')


})


// status model content update

$(document).on('click', '#unbulishslt', function () {

    if (selectedcheckboxarr.length>1){
        $('.deltitle').text(languagedata.Userss.changestatususers)

        $('#content').text(languagedata.Userss.changestatuscontents)
    }else{
        $('.deltitle').text(languagedata.Userss.changestatususer)

        $('#content').text(languagedata.Userss.changestatuscontent)
    }

    $("#delid").text($(this).text());

    $('#delid').addClass('selectedIsActive')

})


//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass("hidden")

})


// Multi Delete users

$(document).on('click', '.checkboxdelete', function () {

    $(".multidelete").addClass("responsive-width")

    var url = window.location.href;

    console.log("clicked the deleted button" + url + " " + JSON.stringify(selectedcheckboxarr))

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    console.log($("#userform input[name='csrf']").val());

    $(".selected-numbers").hide()
    $.ajax({
        url: '/settings/users/deleteselectedusers',
        type: 'post',
        datatype: 'json',
        async: false,
        data: {
            "userIds": JSON.stringify(selectedcheckboxarr),
            csrf: $("#userform input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            if (data.value == true) {
                setCookie("get-toast", "User Deleted Successfully")
                window.location.href = data.url
            } else {
                setCookie("Alert-msg", "Internal Server Error")
            }

        }
    })

})


// Multi select change Access Permission

$(document).on('click', '.selectedunpublish', function () {

    var url = window.location.href

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    var pageno = urlpar.get('page')
    $('.selected-numbers').hide()

    console.log(JSON.stringify(selectedcheckboxarr));


    $.ajax({
        url: "/settings/users/selectedusersaccesschange",
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "userIds": JSON.stringify(selectedcheckboxarr),
            csrf: $("#userform input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            console.log(data, "data");

            var dataStatus

            if (data.status == 1) {
                dataStatus = "User's Only"
            }
            else if (data.status == 0) {
                dataStatus = "All Users"
            }

            if (data.value == true) {

                setCookie("get-toast", "User access permission changed Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})


// Hide checkbox for spurtCms team Member

$(document).ready(function () {
    $("#CheckLabel1").css("display", "none")
    $("#Check1").css("display", "none")
})


// function for single status changer

function UserStatus(id) {
    $('#cbox' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();
    var isActive = $('#cbox' + id).val();
    console.log("isactive", isActive);

    $.ajax({
        url: '/settings/users/changeActiveStatus',
        type: 'POST',
        async: false,
        data: { "id": id, "isActive": isActive, csrf: $("#userform input[name = 'csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {

            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >User Status Updated Successfully</p ></div ></div ></li></ul> `;
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
    })
}

// multi user active status change function

$(document).on('click', '.selectedIsActive', function () {


    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    console.log(selectedcheckboxarr);

    $('.selected-numbers').hide()

    $.ajax({
        url: "/settings/users/selectedUserStatusChange",
        type: "POST",
        dataType: 'json',
        async: false,
        data: {
            "userIds": JSON.stringify(selectedcheckboxarr),
            csrf: $("#userform input[name = 'csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "User status changed Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })
})


 // Group name limit of 25 char
 $(document).on('keyup', '.checklength', function () {
    var inputVal = $(this).val()

    // console.log(inputVal.length);

    var inputLength = inputVal.length

    if (inputLength == 25) {
        $(this).siblings('.lengthErr').removeClass('hidden')
    }else{
        $(this).siblings('.lengthErr').addClass('hidden')
    }


})

