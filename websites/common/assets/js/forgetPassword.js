$("#verify").click(function () {

    jQuery.validator.addMethod(
        "email_validator",
        function (value, element) {

            const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\.[a-zA-Z]{2,})?$/;
            return this.optional(element) || emailPattern.test(value);
        },
        "* Please enter a valid email address"
    );

    jQuery.validator.addMethod("emailcheck", function (value) {

        var result;
        var adminuserid = $("#adminuserid").val()

        $.ajax({
            url: "/checkemailinmember",
            type: "POST",
            async: false,
            data: { "email": value, "id": adminuserid },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
                console.log(data, "data");

            }
        })
        return result.trim() == "true"
    })

    $("form[name='forget']").validate({
        ignore: [],
        rules: {
            emailid: {
                required: true,
                email_validator: true,
                emailcheck:true,
            }
        },
        messages: {
            emailid: {
                required: "* Please enter your Email",
                emailcheck:"* You are not registered with us",
            }
        },

    })


    var formcheck = $("#forget").valid();
    if (formcheck == true) {
        $('#forget')[0].submit();
    }

})


document.addEventListener("DOMContentLoaded", function () {

    var Cookie = getCookie("log-toast");
    var content = getCookie("success");

    if (Cookie != "") {

        $('#emailid-error').show();
        $('#emailgrp').addClass('input-group-error');
    } else if (content != "") {
        $("input[name=emailid]").val("")

        notify_content = ` <span id="" class="para" for="emailid" style="color:green">Thank you! A link to reset your password will be sent to your registered 
        email shortly</span>`
        $(notify_content).insertAfter(".ig-row");

    }

    delete_cookie("log-toast");
    delete_cookie("success");
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