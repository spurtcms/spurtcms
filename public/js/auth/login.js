//Remember me functionality 
$("#Check2").click(function () {

    if ($(this).is(":checked")) {
        $(this).val("1")
        
        $('.user-select-none').removeClass("text-[#D2D2D2]").addClass("text-[#262626]")
    } else {
        $(this).val("0")
        $('.user-select-none').removeClass("text-[#262626]").addClass("text-[#D2D2D2]")

    }
})

$(document).ready(function() {

   

  
    $('#usrid').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });

    $('#passid').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });
});

$(document).on('click', '#loginf', function () {

   
    $('#pas-error').hide();
    $('#em-error').hide();
    $('#em-error1').hide();
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

        if ($("#Check2").val() == 0) {
            console.log($("#Check2").val());
            sessionStorage.setItem("rememberme", "0");
        } else {

            localStorage.setItem("rememberme", "1");
        }

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
    var passchange =getCookie("pass-change")

    console.log("fsrzg", Cookie);

    if (Cookie == "You+are+not+registered+with+us") {
        $('#em-error').show();
        $('#usergrp').addClass('input-group-error');
    }else if (Cookie == "This+account+is+inactive+please+contact+the+admin"){
        $('#em-error1').show();
        $('#usergrp').addClass('input-group-error');
    } else if (content != "") {
        $("#usrid").val(username)
        $('#pas-error').show();
        $('#passgrp').addClass('input-group-error');
    }
    if (getCookie("pass-change") === 'Password+Updated+Successfully') {

     
        notify_content = `<ul class="toast-msg fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > Password Updated Successfully </p ></div ></div ></li></ul> `;
        $(notify_content).insertBefore(".header-rht");

        setTimeout(function () {
            $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();

            });
        }, 5000);
    }

    

    delete_cookie("log-toast");
    delete_cookie("username");
    delete_cookie("pass-toast");
    delete_cookie("pass-change");
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
    $('#em-error1').hide();
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

// delete the cookies and session storage

$("#logout").click(function(){

    console.log("checkss")

    sessionStorage.removeItem("rememberme");
    localStorage.removeItem("rememberme");
})

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
/**Form design */
let inputGroups = document.querySelectorAll('.input-group');
inputGroups.forEach(inputGroup => {
    
    let inputField = inputGroup.querySelector('input');

    inputField.addEventListener('focus', function (event) {
        if(event.target.id !== 'searchcatlist'){
            inputGroup.classList.add('input-group-focused');

        }
    });
    inputField.addEventListener('blur', function () {
        inputGroup.classList.remove('input-group-focused');
    });

});

$(document).ready(function() {
    $('#usrid').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });

    $('#passid').keypress(function(event) {
        if (event.which === 32 && $(this).val().length === 0) {
            event.preventDefault();
        }
    });
});