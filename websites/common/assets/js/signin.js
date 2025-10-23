


$(document).on('click', '#signintbn', function () {

    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {

            const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})?$/;
            return this.optional(element) || emailPattern.test(value);
        },
        "* Please enter a valid email address"
    );
    $("form[name='signinform']").validate({
        ignore: [],
        rules: {
            emailid: {
                required: true,
                email_validator: true,
            },
            password: {
                required: true,
            }
        },
        messages: {
            emailid: {
                required: "* Please enter your Email",
            },
            password: {
                required: "* Please enter your Password",

            }
        },
        errorPlacement: function (error, element) {
            if (element.attr("id") === "password") {
                element.closest(".password").after(error);
            } else {
                error.insertAfter(element);
            }
        }
    })

    console.log("lliiii")
    var formcheck = $("#signinform").valid();
    if (formcheck == true) {
        console.log("truedfdf")
        $('#signinform')[0].submit();
    }
    return false
})

document.addEventListener("DOMContentLoaded", function () {

    var Cookie = getCookie("email-toast");
    var content = getCookie("pass-toast");
    var username = getCookie("email");

    console.log("fsrzg", Cookie);

    if (Cookie == "You+are+not+registered+with+us") {
        $('#em-error').show();

    } else if (Cookie == "This+account+is+inactive+please+contact+the+admin") {
        $('#em-error1').show();

    } else if (content != "") {
        $("#email").val(username)
        $('#pas-error').show();

    }

    delete_cookie("email-toast");

    delete_cookie("pass-toast");
});

function getCookie(cname) {
    let name = cname + "=";
    let ca = document.cookie.split(';');

    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function delete_cookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}


$(document).on('click', '#eyeicon', function () {

    console.log("checklogin")

    value = $("#password").attr('type')

    if (value == "text") {

        $("#password").attr('type', "password")

        $(this).children('img').attr('src', '/websites/public/img/hide-password.svg')

    } else if (value == "password") {

        $("#password").attr('type', "text")

        $(this).children('img').attr('src', '/websites/public/img/show-password.svg')
    }
})

$(document).on('keyup','#email',function(){

    $('#em-error').hide()
    $('#em-error1').hide()
})

$(document).on('keyup','#password',function(){

    $('#pas-error').hide()
   
})

$(document).ready(function() {
    $('#email').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });

    $('#password').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });
});
