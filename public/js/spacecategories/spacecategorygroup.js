$('#save').click(function(){

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        category_id= $("#category_id").val()
     
        $.ajax({
            url:"/spaces/categorygroup/create",
            type:"POST",
            async:false,
            data:{"category_name":value,"category_id":category_id,csrf:$("input[name='csrf']").val()},
            datatype:"json",
            caches:false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim()!="true"
    })

    $("#category_form").validate({
        rules: {
            category_name: {
                required: true,
                // categoryname_validator: true
                space: true,
                duplicatename:true,
            },
            category_desc: {
                required: true,
                // categorys_desc: true,
                maxlength: 250,
                space :true
            }
        },
        messages: {
            category_name: {
                required: "* " + languagedata.Categoryy.catgrpnamevalid,
                space: "* " + languagedata.spacergx,
                duplicatename:"*"+ languagedata.Categoryy.namevailderr
            },
            category_desc: {
                required: "* " + languagedata.Categoryy.catgrpdescvalid,
                maxlength: "* "+languagedata?.Permission?.descriptionchat

            },
        }
    });

    var formcheck = $("#category_form").valid();

    if (formcheck == true) {
     
        $('#category_form')[0].submit();
    
    }
    
    else{
        
        Validationcheck()
    
        $(document).on('keyup',".field",function(){
            
            Validationcheck()
        
        })

    }

   return false

})
