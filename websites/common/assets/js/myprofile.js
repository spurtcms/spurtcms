// update membes
$(document).on('click', '#update', function () {

    console.log("cccccc")

     var mem_id = $("#memberid").val()

      console.log(mem_id,"checkmemberid")

    jQuery.validator.addMethod("duplicateemail", function (value) {
 
            var result;

            $.ajax({
                url: "/checkemailinmember",
                type: "POST",
                async: false,
                data: { "email": value, "id": mem_id },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                    console.log(data, "data");
 
                }
            })
            return result.trim() != "true"
        })

    jQuery.validator.addMethod("duplicatenumber", function (value) {

        var result;
     

        $.ajax({
            url: "/checknumberinmember",
            type: "POST",
            async: false,
            data: { "number": value, "id": mem_id },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();

            }
        })
        return result.trim() != "true"
    })


    jQuery.validator.addMethod("mob_validator", function (value, element) {
        if (value.length >= 7)
            return true;
        else return false;
    }, "* Mobile number must contain  7 digits");

    $.validator.addMethod("email_validator", function (value) {
        return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    }, '*Please enter a valid email address');



    $.validator.addMethod("space", function (value) {
        return /^[^-\s]/.test(value);
    });


    $("#myprofileform").validate({
        ignore: [],
        rules: {
            prof_pic: {
                extension: "jpg|png|jpeg"
            },
            mem_name: {
                required: true,
                space: true
            },
            mem_email: {
                required: true,
                email_validator: true,
                duplicateemail: true
            },

            mem_mobile: {
                required: true,
                mob_validator: true,
                duplicatenumber: true
            },

        },
        messages: {
            prof_pic: {
                extension: "* Please select jpg /png /jpeg files only"
            },
            mem_name: {
                required: "* Please enter the FirstName",
                space: "* No space Prefix"
            },

            mem_email: {
                required: "* Please enter the EmailId",
                duplicateemail: "* Emailid Already Exists "
            },

            mem_mobile: {
                required: "* Please enter the MobileNo",
                duplicatenumber: "* MobileNo Already Exists",
            },

        }
    })


    var formcheck = $("#myprofileform").valid();
    console.log("working this ", formcheck);
    if (formcheck == true) {
        $('#myprofileform')[0].submit();
    } else {
        console.log(" not working this ");

    }

    return false
})

$('#mem_mobile').keyup(function () {
    this.value = this.value.replace(/[^0-9\.]/g, '');
});


$(document).on('click','#removebtn',function(){

    $('#profileimg').attr('src',"")
    $('#crop_data').val("")
    $('#imgdiv').text($('#imgdiv').attr('namestring'))
    $('#remove_image').val("1")
})

$('#myfile').on('change', function(event) {
    var file = event.target.files[0];
    if (file) {
        var reader = new FileReader();
        reader.onload = function(e) {
            $('#imgdiv').empty(); 
       
            $('#imgdiv').append('<img src="' + e.target.result + '" id="profileimg" class="h-[90px] object-cover rounded-full" alt="Profile Image">');

            $('#crop_data').val(e.target.result)
        };
        reader.readAsDataURL(file);
    }
});

