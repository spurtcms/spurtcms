var languagedata



$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {


        languagedata = data
    })
})




$(document).ready(function(){


    $("#itegration-modal").on("hide.bs.modal", function () {
        

        $('#IntegrationId').val('')
        $('#GatewayName').val('')
        $('#GatewayDesc').val('')
        $('#ClientId').val('')
        $('#ClientSecret').val('')

        $('#GatewayName-error').hide()
        $('#GatewayDesc-error').hide()
        $('#ClientId-error').hide()
        $('#ClientSecret-error').hide()

        
        $('#integration_activestatus').prop('checked', false).val(0);


    })

    $(document).on('click','.Integrationmanage',function(){
        
        $('#itegration-modal').modal('show');

        var IntegrationId = $(this).data('id');
        

        $.ajax({
            url: "/admin/integration/edit",
            type: "GET",
            async: false,
            data: { "id": IntegrationId },
            datatype: "json",
            caches: false,
            success: function (data) {
    
               var Editintegration=data.Integration
               
               $('#IntegrationId').val(Editintegration.Id)

               $('#GatewayName').val(Editintegration.GatewayName)
               $('#GatewayDesc').val(Editintegration.GatewayDesc)

               $('#ClientId').val(Editintegration.ClientId)
               
               $('#ClientSecret').val(Editintegration.ClientSecret)

               if (Editintegration.IsActive == 1) {
                $('#integration_activestatus').prop('checked', true).val(1);
            } else {
                $('#integration_activestatus').prop('checked', false).val('');
            }

            }
        })

    })


        // Is_active

        $("input[name=integration_activestatus]").click(function () {
            if ($(this).prop('checked')) {
                $(this).attr("value", "1")
            } else {
                $(this).removeAttr("value")
            }
        })



        $(document).on('click', '#UpdateIntegration', function () {
            // Form Validation
    
            // console.log("checking::");
            // var formAction = $("#membershipform").attr('action');
            // var passwordRequired = (formAction !== '/membership/updatemember');
    
          
            $("#IntegrationForm").validate({
                ignore: [],
                rules: {
                    // prof_pic: {
                    //     extension: "jpg|png|jpeg"
                    // },
                    GatewayName: {
                        required: true,
                        space: true
                    },
                    GatewayDesc: {
                        required: true,
                    },
                    ClientId: {
                        required: true,
                        space: true,
                    },
                    
                    ClientSecret: {
                        required: true,
                        space: true,
                    }
    
    
                },
                messages: {
                    // prof_pic: {
                    //     extension: "* " + languagedata.profextension
                    // },
                    GatewayName: {
                        required: "* " + "Please enter the Gateway Name",
                        space: "* " + languagedata.spacergx
                    },
    
                    GatewayDesc: {
                        required: "* " + " Please enter the Gateway Desc",
                    },
    
                    ClientId: {
                        required: "* " + "Please enter the Client Id",
                        space: "* " + languagedata.spacergx,
                    },
                    ClientSecret: {
                        required: "* " + "Please enter the Client Secret",
                        space: "* " + languagedata.spacergx,
                    }
                }
            })
    
            var formcheck = $("#IntegrationForm").valid();
    
            if (formcheck) {
                $('#IntegrationForm')[0].submit();
            }
            else {
                console.log("form nt valid");
    
            }
    
        })

    


})
