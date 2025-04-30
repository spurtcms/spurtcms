var languagedata
/**This if for Language json file */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })

});



// Email format validation
$.validator.addMethod("email_validator", function (value) {
    return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
}, '* ' + "Please enter the vaild email");
$.validator.addMethod("space", function (value) {
    return /^[^-\s]/.test(value);
});
// Format vaildation


    $("#addemailconfig").on("click" , function(){

        $("#emailconfigform").validate({
    
            rules:{
    
                email: {
                    required: true,
                    space: true,
                    email_validator:true
                },
                password: {
                    required: true,
                    space: true
                },
                host: {
                    required: true,
                    space: true
                },
                port: {
                    required: true,
                    space: true
                },
    
            },
            messages:{
                email: {
                    required: "* " + languagedata.Setting.emailerror,
                    space: "* " + languagedata.Spaces.spacergx,
                    
                },
                password: {
                    required: "* " + languagedata.Setting.passworderr,
                    space: "* " + languagedata.spacergx,
                    
                },
                host: {
                    required: "* " + languagedata.Setting.hosterr,
                    space: "* " + languagedata.spacergx,
                    
                },
                port: {
                    required: "* " + languagedata.Setting.porterr,
                    space: "* " + languagedata.spacergx,
                    
                },
            },
    
        });

        // if (("#radiosmtp").prop("checked" ,true)){
        //     var formcheck = $("#emailconfigform").valid();

        // }else{
        //     $('#emailconfigform')[0].submit();
        // }
                var formcheck = $("#emailconfigform").valid();

        if (formcheck == true) {
            
            $('#emailconfigform')[0].submit();

        }
        else {
            Validationcheck()
            $(document).on('keyup', ".field", function () {
                Validationcheck()
            })
        }
    
    })




// Input error class add and remove

function Validationcheck() {

    if ($('#emailaddress').hasClass('error')) {
        $('#config_email').addClass('input-group-error');
    } else {
        $('#config_email').removeClass('input-group-error');
    }

    if ($('#mailpassword').hasClass('error')) {
        $('#config_password').addClass('input-group-error');
    } else {
        $('#config_password').removeClass('input-group-error');
    }

    if ($('#mailhost').hasClass('error')) {
        $('#config_host').addClass('input-group-error');
    } else {
        $('#config_host').removeClass('input-group-error');
    }

    if ($('#mailport').hasClass('error')) {
        $('#config_port').addClass('input-group-error');
    } else {
        $('#config_port').removeClass('input-group-error');
    }
}

// only allow numbers
$('#mailport').keyup(function () {
    this.value = this.value.replace(/[^0-9\.]/g, '');
});


// Choose smtp format mail send
$("#smtpmail").click(function (){
    $("#radioenv").prop('checked',false)
    $("#emailseltype").val("smtp")
    $("#emailconfigdiv").show()
})

// Choose environment format mail send
$("#environmentmail").click(function(){
    $("#radiosmtp").prop('checked',false)
    $("#emailseltype").val("environment")
    $("#emailconfigdiv").hide()
})

function setroute(){

    window.location.href="/settings/emails/"
}

function setroute1(){

    window.location.href="/settings/emails/emailconfig/"
}