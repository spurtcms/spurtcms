var languagedata
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

})

$(document).on('click', '#uptprofile', function () {

    console.log("checking")

    jQuery.validator.addMethod("duplicateemail", function (value) {

        var result;
        user_id = $("#id").val()
        console.log("id", user_id)
        $.ajax({
            url: "/settings/checkemail",
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

    jQuery.validator.addMethod("duplicateusername", function (value) {

        var result;
        user_id = $("#id").val()
        $.ajax({
            url: "/settings/checkusername",
            type: "POST",
            async: false,
            data: { "username": value, "id": user_id, csrf: $("input[name='csrf']").val() },
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
        user_id = $("#id").val()
        $.ajax({
            url: "/settings/checknumber",
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
    // $.validator.addMethod("mob_validator", function (value) {
    //     return /^[6-9]{1}[0-9]{9}$/.test(value);
    // }, '* ' + languagedata.Userss.usrmobnumrgx);
    jQuery.validator.addMethod(
        "mob_validator",
        function (value, element) {
            if (/^[6-9]{1}[0-9]{9}$/.test(value))
                return true;
            else return false;
        },
        "* " + languagedata.Userss.usrmobnumrgx
    );

    $.validator.addMethod("email_validator", function (value) {
        return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    }, '* ' + languagedata.Userss.usremailrgx);

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
                duplicateemail: true
            },
            // user_role: {
            //     required: true,
            // },
            user_name: {
                required: true,
                duplicateusername: false,
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
            prof_pic: {
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
                duplicateusername: "*" + languagedata.Userss.nameexist,
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
    var formcheck = $("#userform").valid();

    if (formcheck == true) {
        $('#userform')[0].submit();
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
        // if ($('#rolen-error').css('display') !=='none'){
        //     console.log("check")
        //         $('.user-drop-down').addClass('input-group-error')
        //     }

    }
    return false;
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


    });
}

$(document).on('click', '.btn-close', function () {

    $('#changepicModal').modal('hide');
})

$(document).on('click', '#crop-button', function () {
    $(".name-string").hide()
    $('#profpic').show()
})

$('input[name=user_mob]').keyup(function () {
    this.value = this.value.replace(/[^0-9\.]/g, '');
});

// $(document).on('change', '#myfile', function () {
//    $("#prof-crop").val("2")
// })

$(document).on('click', '#back', function () {

    window.location.href = "/settings/myprofile"
})


$(document).on('keypress', '.checklength', function () {
    var inputVal = $(this).val()

    // console.log(inputVal.length);

    var inputLength = inputVal.length

    if (inputLength == 25) {
        $(this).siblings('.lengthErr').removeClass('hidden')
    }


})


$(document).on('keyup', '.checklength', function (e) {

    if (e.which == 8) {
        var inputVal = $(this).val()

        var inputLength = inputVal.length

        if (inputLength != 25) {
            $(this).siblings('.lengthErr').addClass('hidden')
        }
    }
})


$(document).on('keypress', '.mobileInput', function () {
    var inputVal = $(this).val()

    // console.log(inputVal.length);

    var inputLength = inputVal.length

    if (inputLength == 10) {
        $(this).siblings('.lengthErr').removeClass('hidden')
        $(this).prop('readonly', true)
    }


})


$(document).on('keyup', '.mobileInput', function (e) {

    if (e.which == 8) {
        var inputVal = $(this).val()

        var inputLength = inputVal.length

        if (inputLength <= 10) {
            $(this).siblings('.lengthErr').addClass('hidden')
            $(this).prop('readonly', false)
        }
    }
})





