var languagedata
/** */
$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })

})

$(document).ready(function () {

    $.validator.addMethod("password1", function (value) {
        return /^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/.test(value);
    }, '* '+languagedata.Userss.usrpswdrgx);
 
});

$(document).on('click','.passsave',function(){

    // jQuery.validator.addMethod("duplicatepassword", function (value) {
    //     var result;
    //     $.ajax({
    //         url:"/settings/checkpassword",
    //         type:"POST",
    //         async:false,
    //         data:{"pass":value,csrf:$("input[name='csrf']").val()},
    //         datatype:"json",
    //         caches:false,
    //         success: function (data) {
    //             result = data.trim();
    //         }
    //     })
    //     return result.trim()!="true" ;
    // });

    $("#passform").validate({
        
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
                required: "* "+languagedata.Userss.usrpswd,
                // duplicatepassword: "* New password must be different from the old password"
            },
            cpass: {
                required: "* "+languagedata.confirmpswd,
                equalTo: "* "+languagedata.confirmpswdrgx
            }
        }
    })
    var formcheck = $("#passform").valid();
    if (formcheck == true) {
        $('#passform')[0].submit();
    }
    else{
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
    }
})

function Validationcheck(){
   
    if ($('#pass').hasClass('error')) {
        $('#passgroup').addClass('input-group-error');
    }else{
        $('#passgroup').removeClass('input-group-error');
    }
    
    if ($('#cpass').hasClass('error')){
        $('#cpassgroup').addClass('input-group-error');
    }else{
        $('#cpassgroup').removeClass('input-group-error');
    }
}
// Password Change
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

// Password Change
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

