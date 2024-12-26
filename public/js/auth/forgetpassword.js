$(document).on('click', '#forgotfrm-btn', function () {
    $('#emailid-error').hide();
    $("#forgotmailform").validate({
        rules: {
            emailid: {
                required: true,
                email_validator: true,
            }
        },
        messages: {
            emailid: {
                required: "* Please enter your email",
            }
        }
    })
    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {
            if (/(^[a-zA-Z_0-9\.-]+)@([a-z]{2,})\.([a-z]{2,})(\.[a-z]{2,})?$/.test(value))
                return true;
            else return false;
        },
        "* Please enter valid email"
    );
    var formcheck = $("#forgotmailform").valid();
    if (formcheck == true) {
        $('#forgotmailform')[0].submit();
    }
    else {
        console.log("ug");
        EmailValidationcheck()
        $(document).on('keyup', ".field", function () {
            EmailValidationcheck()
        })
    }

    return false

})
function EmailValidationcheck() {
console.log("dfgh");
    if ($('#emailid').hasClass('error')) {
        $('#emailgrp').addClass('input-group-error');
    } else {
        $('#emailgrp').removeClass('input-group-error');
    }


}

document.addEventListener("DOMContentLoaded", function () {

    var Cookie = getCookie("log-toast");
    var content = getCookie("success");

    console.log("fsrzg", Cookie,content);
    if (Cookie != "") {

        $('#emailid-error').show();
        $('#emailgrp').addClass('input-group-error');
    } else if (content != "") {
        $("input[name=emailid]").val("")

        notify_content = ` <span id="" class="para" for="emailid" style="color:green; font-size:0.75rem">Thank you! A link to reset your password will be sent to your registered 
        email shortly</span>`
        $(notify_content).insertAfter(".ig-row");


        // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>Reset password email send successfully</span></div>'
        // $(notify_content).insertBefore(".login-header");
        // setTimeout(function () {
        //     $('.toast-msg').fadeOut('slow', function () {
        //         $(this).remove();
        //     });
        // }, 5000); // 5000 milliseconds = 5 seconds
    }

    delete_cookie("log-toast");
    delete_cookie("success");
});


