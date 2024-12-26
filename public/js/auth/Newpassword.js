
$(document).on('click', '#passsave', function () {
    $("#passform").validate({
        rules: {
            pass: {
                required: true,
                pass_validator: true,

            },
            cpass: {
                required: true,
                equalTo: '[name=pass]'  //'#new_pass'

            }
        },
        messages: {
            pass: {
                required: "* Please enter your new password",
            },
            cpass: {
                required: "* Please enter your confirm password",
                equalTo: " * Password Mismatch"
            }
        }
    })

    jQuery.validator.addMethod(
        "pass_validator", function (value, element) {
            if (value != "") {
                if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/.test(value))
                    return true;
                else return false;
            }
            else return true;
        },
        "* Please Enter at Least 1 Uppercase, 1 Lowercase, 1 Number,1 Special Character($,@),and 8 characters long"
    );

    var formcheck = $("#passform").valid();
    if (formcheck == true) {
        $('#passform')[0].submit();
    }
    else {
        NewValidationCheck()
        $(document).on('keyup', ".field", function () {
            NewValidationCheck()
        })
    }

    return false


})
function NewValidationCheck() {
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

document.addEventListener("DOMContentLoaded", function () {


    var Cookie = getCookie("log-toast");

    if (Cookie != "") {
        $("input[name=emailid]").val("")
        notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>controllers/authcontroller.go</span></div>'
        $(notify_content).insertBefore(".login-header");
        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
            });
        }, 5000);
    }

    delete_cookie("log-toast");
});
$(document).on('click', '#eye1', function () {

    var This = $("#pass")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');
        

    } else {

        $(This).attr('type', 'password');

        $(this).find('img').attr('src', '/public/img/eye-closed.svg');
    }
})
$(document).on('click', '#eye2', function () {

    var This = $("#cpass")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');
        

    } else {

        $(This).attr('type', 'password');

        $(this).find('img').attr('src', '/public/img/eye-closed.svg');
    }
})