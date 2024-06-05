var languagedata
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })

})

$(document).on('click', '.passsave', function (event) {

    event.preventDefault()

    $.validator.addMethod("password1", function (value) {
        return /^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/.test(value);
    }, '* ' + languagedata.Userss.usrpswdrgx);

    $("form[name='passform']").validate({

        rules: {
            pass: {
                required: true,
                password1: true,
                // duplicatepassword:true
            },
            cpass: {
                required: true,
                equalTo: "#pass"
            }
        },
        messages: {
            pass: {
                required: "* " + languagedata.Userss.usrpswd,
                // duplicatepassword: "* New password must be different from the old password"
            },
            cpass: {
                required: "* " + languagedata.confirmpswd,
                equalTo: "* " + languagedata.confirmpswdrgx
            }
        }
    })
    var formcheck = $("#passform").valid();
    console.log("formchk", formcheck == true);
    if (formcheck == true) {
        console.log("okkk");
        $.ajax({
            url: "/settings/checkpassword",
            type: "POST",
            async: false,
            data: { "pass": $("input[name='pass']").val(), csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                console.log("result", data.pass, data.pass == true);
                if (data.pass == true) {
                    $('#passform')[0].reset();
                    notify_content = '<div style="top:2px;" class="toast-msg dang-red"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/danger-group-12.svg" alt="" class="left-img" /> <span> New password must be different from the old password </span></div>';

                    $(notify_content).insertBefore(".header-rht");

                    setTimeout(function () {

                        $('.toast-msg').fadeOut('slow', function () {

                            $(this).remove();

                        });

                    }, 5000);

                } else {
                    $('#passform')[0].submit();
                }
            }
        })
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
    }
})

function Validationcheck() {

    if ($('#pass').hasClass('error')) {
        $('#passgroup').addClass('input-group-error');
    } else {
        $('#passgroup').removeClass('input-group-error');
    }

    if ($('#cpass').hasClass('error')) {
        $('#cpassgroup').addClass('input-group-error');
    } else {
        $('#cpassgroup').removeClass('input-group-error');
    }
}
// Password Change
$(document).on('click', '#eye1', function () {

    var This = $("#pass")

    if ($(This).attr('type') === 'password') {

        $('#eye-close1').hide()

        $('#eye-open1').show()

        $(This).attr('type', 'text');

    } else {

        $('#eye-open1').hide()

        $('#eye-close1').show()

        $(This).attr('type', 'password');

    
    }
})

// Password Change
$(document).on('click', '#eye2', function () {

    var This = $("#cpass")

    if ($(This).attr('type') === 'password') {

        $('#eye-close2').hide()

        $('#eye-open2').show()

        $(This).attr('type', 'text');


    } else {

        $('#eye-open2').hide()

        $('#eye-close2').show()

        $(This).attr('type', 'password');

    }
})

