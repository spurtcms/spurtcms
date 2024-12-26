
$(document).ready(function() {

  currentloc= window.location.href 

  if (currentloc.includes("/login/")) {

    setTimeout(function() {
        $('#resendbtn').removeClass("text-[#ACABA9] pointer-events-none ").addClass("text-[#05ACD7] hover:underline  hover:underline"); 
    }, 30000);

  }
   
    
});

function handleKeyPress(event) {
    
    if (event.key === 'Enter') {
      
        event.preventDefault();
      
        document.getElementById('rloginform').submit();
    }

    const input = event.target;
    
    if (!/[0-9]/.test(event.key) || input.value.length >= 6) {
        event.preventDefault(); 
    }
}

$(document).on('click', '.loginf', function () { 
    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {
          
            const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})?$/;
            return this.optional(element) || emailPattern.test(value);
        },
        "* " + languagedata.Userss.usremailrgx
    );
    
    $('#pas-error').hide();
    $('#em-error').hide();
    $('#em-error1').hide();
    $("form[name='rloginform']").validate({
        rules: {
            emailid: {
                required: true,
                email_validator: true,
            },

        },
        messages: {
            emailid: {
                required: "* Please enter your EmailId",
            },

        }
    })
    var formcheck = $("#rloginform").valid();
    if (formcheck == true) {
        $('#rloginform')[0].submit();
        setTimeout(function() {
           
            $('#resendbtn').removeClass("text-[#ACABA9]").addClass("text-[#05ACD7]"); 
        }, 30000);

    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
    }

    return false
})
// $(document).on('click', '.otplogin', function () {

//     $("form[name='rloginform']").validate({
//         rules: {
//             otpno: {
//                 required: true,
//             },

//         },
//         messages: {
//             otpno: {
//                 required: "* Please enter OTP",
//             },

//         }
//     })
//     var formcheck = $("#rloginform").valid();
//     if (formcheck == true) {
//         $('#rloginform')[0].submit();

//     }
//     else {
//         Validationcheck()
//         $(document).on('keyup', ".field", function () {
//             Validationcheck()
//         })
//     }

//     return false

// })

$(document).ready(function () {
    $(document).on('click', '.otplogin', function () {
       
        $('#otp-errorr').hide();

      
        let allFilled = true;
        $('.otp-input').each(function() {
            if ($(this).val() === '') {
                allFilled = false;
                return false; 
            }
        });

        if (!allFilled) {
            $('#otp-errorr').show(); 
            return false; 
        }

       
        let otp = '';
        $('.field').each(function() {
            otp += $(this).val(); 
        });

       
        $('#rloginform')[0].submit(); 

        return false; 
    });

   
    $(document).on('keyup', '.otp-input', function () {
        $('#otp-errorr').hide(); 
        $('#otp-error').hide();
    });
});
function Validationcheck() {
    if ($('#usrid').hasClass('error')) {
        $('#usergrp').addClass('input-group-error');
    } else {
        $('#usergrp').removeClass('input-group-error');
    }

}

document.addEventListener("DOMContentLoaded", function () {

    var Cookie = getCookie("log-toast");
    var content = getCookie("pass-toast");
    var username = getCookie("username");

    console.log("fsrzg", Cookie);

    if (Cookie == "You+are+not+registered+with+us") {
        $('#em-error').show();
        $('#usergrp').addClass('input-group-error');
    } else if (Cookie == "This+account+is+inactive+please+contact+the+admin") {
        $('#em-error1').show();
        $('#usergrp').addClass('input-group-error');
    } else if (Cookie == "Invalid+Otp") {

        $('#otp-error').show()
    } else if (Cookie == "Otp+Expired") {

        $('#otp-error1').show()
        $('#resendbtn').removeClass("text-[#ACABA9] pointer-events-none ").addClass("text-[#05ACD7] hover:underline  hover:underline"); 
    }
    else if (content != "") {
        $("#usrid").val(username)
        $('#pas-error').show();
        $('#passgrp').addClass('input-group-error');
    }

    delete_cookie("log-toast");
    delete_cookie("username");
    delete_cookie("pass-toast");
});

$(document).on('keyup', ".field", function () {
    $('#em-error').hide();
    $('#em-error1').hide();
    $('#pas-error').hide();
    $('#passgrp').removeClass('input-group-error');
    $('#usergrp').removeClass('input-group-error');
})



$(document).on('click', '.copybtn', function () {


    var copyval = $(this).siblings('p').text()

    navigator.clipboard.writeText(copyval).then(() => {

        $(this).tooltip('dispose');
        $(this).attr("data-bs-title", "Copied");
        $(this).tooltip();
        $(this).tooltip('show');
        setTimeout(() => {
            $(this).tooltip('dispose');
            $(this).removeAttr("data-bs-title");
        }, 1000);
    }, () => {
        console.error('Failed to copy');
    });
    $(this).off('mouseenter mouseleave');
})

$(document).on('click', '.otpno', function () {

    $('#otp-error').hide()
})

$(document).on('click', '#resendbtn', function () {

    event.preventDefault();

    $(this).removeClass("text-[#05ACD7] hover:underline  hover:underline ").addClass("text-[#ACABA9] pointer-events-none");

    email = $("#useremail").val()
    $.ajax({
        url: "/login/resendotp",
        type: "POST",
        async: false,
        data: { "emailid": email, csrf: $("input[name='csrf']").val() },
        datatype: "json",
        caches: false,
        success: function (data) {


    setTimeout(function() {
        $('#resendbtn').removeClass("text-[#ACABA9] pointer-events-none").addClass("text-[#05ACD7] hover:underline  hover:underline"); 
    }, 30000);

        }
    })
})

$(document).on('click', '#cancelbtn', function () {

    event.preventDefault();

    window.location.href = "/weblogin"
})



document.addEventListener('DOMContentLoaded', function () {
    const otpInputs = document.querySelectorAll('.otp-input');
    const combinedOtpInput = document.getElementById('otpno');

    otpInputs.forEach((input, index) => {
        input.addEventListener('input', function () {
          
            this.value = this.value.replace(/[^0-9]/g, '');

            if (this.value.length === 1 && index < otpInputs.length - 1) {
                otpInputs[index + 1].focus();
            }

          
            updateCombinedOtp();
        });

       
        input.addEventListener('keydown', function (e) {
            if (e.key === 'Backspace' && this.value.length === 0 && index > 0) {
                otpInputs[index - 1].focus();
            }
        });
        input.addEventListener('paste', function (event) {
            event.preventDefault(); 
            let pasteData = (event.clipboardData || window.clipboardData).getData('text'); 
            pasteData = pasteData.replace(/[^0-9]/g, ''); 
         
            combinedOtpInput.value = pasteData;
           
            for (let i = 0; i < pasteData.length && index + i < otpInputs.length; i++) {
                otpInputs[index + i].value = pasteData[i]; 
            }
    
           
            if (pasteData.length > 0) {
                otpInputs[Math.min(index + pasteData.length - 1, otpInputs.length - 1)].focus();
            }
        });
    });
    
    function updateCombinedOtp() {
        let otp = '';
        otpInputs.forEach(input => {
            otp += input.value;
        });
        combinedOtpInput.value = otp; 
    }
});