
$(document).ready(function () {

    $("#websiteUrlSave").click(function (e) {
        e.preventDefault();


        let websiteInput = $("#websiteInput").val();
        const isValid = /^[a-z0-9]*$/.test(websiteInput);

        if (websiteInput === "") {
            $("#websiteInput-error").show().text("* Please Enter Your Website URL");
            return;
        }

        if (!isValid) {
            $("#websiteInput-error").show().text("Only lowercase letters and numbers are allowed");
            return;
        }


        $.ajax({
            url: "/admin/website/checksitename",
            type: "POST",
            data: {
                "sitename": websiteInput,
                csrf: $("input[name='csrf']").val(),
            },
            dataType: "json",
            cache: false,
            success: function (result) {
                if (result) {
                    $("#websiteInput-error").show().text("Name Already Exists");
                } else {
                    $("#websiteInput-error").hide();
                    $('#websiteUrlForm')[0].submit();
                }
            },

        });
    });


    $("#websiteInput").on('input', function () {

        let value = $(this).val();

        const isValid = /^[a-z0-9]*$/.test(value);

        if (value === '') {

            $("#websiteInput-error").show().text("* Please Enter Your Website URL")

        } else if (value.length >= 20) {

            $("#websiteInput-error").show().text("Maximum length must be 20 characters")

        } else if (!isValid) {

            $("#websiteInput-error").show().text("Only lowercase letters and numbers are allowed");

        } else {

            $("#websiteInput-error").hide().text("");

        }
    })

})


$(document).on('click','#skipnow',function(){

    
})