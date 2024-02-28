
$(document).on('click', '#loginf', function () {
    $('#pas-error').hide();
    $('#em-error').hide();
    $("form[name='loginform']").validate({
        rules: {
            username: {
                required: true,
            },
            pass: {
                required: true,
            }
        },
        messages: {
            username: {
                required: "* Please enter your Username",
            },
            pass: {
                required: "* Please enter your Password",

            }
        }
    })
    var formcheck = $("#loginform").valid();
    if (formcheck == true) {
        $('#loginform')[0].submit();
        $('.spinner-border').show()
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
    }

    return false
})

function Validationcheck() {
    if ($('#usrid').hasClass('error')) {
        $('#usergrp').addClass('input-group-error');
    } else {
        $('#usergrp').removeClass('input-group-error');
    }

    if ($('#passid').hasClass('error')) {
        $('#passgrp').addClass('input-group-error');
    } else {
        $('#passgrp').removeClass('input-group-error');
    }


}
document.addEventListener("DOMContentLoaded", function () {

    var Cookie = getCookie("log-toast");
    var content = getCookie("pass-toast");
    var username = getCookie("username");

    console.log("fsrzg", username);

    if (Cookie != "") {

        $('#em-error').show();
        $('#usergrp').addClass('input-group-error');
    } else if (content != "") {
        $("#usrid").val(username)
        $('#pas-error').show();
        $('#passgrp').addClass('input-group-error');
    }

    delete_cookie("log-toast");
    delete_cookie("username");
    delete_cookie("pass-toast");
});
$(document).on('click', '#eye', function () {

    var This = $("#passid")

    if ($(This).attr('type') === 'password') {

        $(This).attr('type', 'text');

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');
        

    } else {

        $(This).attr('type', 'password');

        $(this).find('img').attr('src', '/public/img/eye-closed.svg');
    }
})

$(document).on('keyup', ".field", function () {
    $('#em-error').hide();
    $('#pas-error').hide();
    $('#passgrp').removeClass('input-group-error');
    $('#usergrp').removeClass('input-group-error');
})


// const rmCheck = document.getElementById("Check2"),
//     emailInput = document.getElementById("usrid");

// if (localStorage.checkbox && localStorage.checkbox !== "") {
//   rmCheck.setAttribute("checked", "checked");
//   emailInput.value = localStorage.username;
// } else {
//   rmCheck.removeAttribute("checked");
//   emailInput.value = "";
// }

// function lsRememberMe() {
//   if (rmCheck.checked && emailInput.value !== "") {
//     localStorage.username = emailInput.value;
//     localStorage.checkbox = rmCheck.value;
//   } else {
//     console.log("uncehckkkk")
//     localStorage.username = "";
//     localStorage.checkbox = "";
//   }
// }
