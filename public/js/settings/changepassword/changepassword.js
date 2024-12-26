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

                    console.log("checkformsub")
                    $('#passform')[0].reset();
                    notify_content = '<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">New password must be different from the old password</p></div></div> </li></ul>';

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
      
        $(this).find('img').attr('src', '/public/img/eye-opened.svg');

        $(This).attr('type', 'text');

    } else {

        $(this).find('img').attr('src',"/public/img/eye-closed.svg")

        $(This).attr('type', 'password');


    }
})

// Password Change
$(document).on('click', '#eye2', function () {

    var This = $("#cpass")

    if ($(This).attr('type') === 'password') {

        $(this).find('img').attr('src', '/public/img/eye-opened.svg');

        $(This).attr('type', 'text');


    } else {
        $(this).find('img').attr('src',"/public/img/eye-closed.svg")

        $(This).attr('type', 'password');

    }
})



