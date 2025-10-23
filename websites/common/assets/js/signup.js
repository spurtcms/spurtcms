$(document).ready(function () {

    $("#gettingstarted").click(function () {

        jQuery.validator.addMethod("duplicatename", function (value) {

            var result;
            var adminuserid = $("#adminuserid").val()

            $.ajax({
                url: "/checknameinmember",
                type: "POST",
                async: false,
                data: { "name": value, "id": adminuserid },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                    console.log("result:", result);

                }
            })
            return result.trim() != "true"
        })

        $.validator.addMethod("email_validator", function (value) {
            return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
        }, "* Please enter a valid email address");

        jQuery.validator.addMethod("duplicateemail", function (value) {

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
            return result.trim() != "true"
        })

        jQuery.validator.addMethod("pass_validator", function (value, element) {
            if (value != "") {
                if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                    return true;
                else return false;
            }
            else return true;
        }, "* Password must have at least 1 uppercase, 1 lowercase, 1 number, 1 special character($, @), and be at least 8 characters long");

        $("#templatesignup").validate({
            ignore: [],
            rules: {
                username: {
                    required: true,
                    duplicatename: true
                },
                email: {
                    required: true,
                    email_validator: true,
                    duplicateemail: true
                },
                password: {
                    required: true,
                    pass_validator: true
                }
            },
            messages: {
                username: {
                    required: "* Please Enter UserName",
                    duplicatename: "* This UserName already exists"
                },
                email: {
                    required: "* Please Enter Email",
                    duplicateemail: "* This Email already exists"
                },
                password: {
                    required: "* Please Enter Password"
                }
            },
            errorPlacement: function (error, element) {
                if (element.attr("id") === "password") {
                    element.closest(".password").after(error);
                } else {
                    error.insertAfter(element);
                }
            }
        });

        var formcheck = $("#templatesignup").valid();

        if (formcheck == true) {
            $('#templatesignup')[0].submit();
        }
    });

    $("#eyeicon").click(function () {

        value = $("#password").attr('type')

        if (value == "text") {

            $("#password").attr('type', "password")

            $(this).children('img').attr('src', '/websites/public/img/hide-password.svg')

        } else if (value == "password") {

            $("#password").attr('type', "text")

            $(this).children('img').attr('src', '/websites/public/img/show-password.svg')
        }

    })
});