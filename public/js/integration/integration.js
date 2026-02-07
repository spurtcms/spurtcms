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

                console.log(data.Integration,"data.Integration");
                
    
               var Editintegration=data.Integration

               console.log(Editintegration.credentials.credentials.gateway_name, "→ Gateway Name");

            
               
               $('#IntegrationId').val(Editintegration.id)
               

               $('#GatewayName').val(Editintegration.credentials.credentials.gateway_name)
               $('#GatewayDesc').val(Editintegration.credentials.credentials.gateway_desc)

            

               var integrationType = Editintegration.integration_type; // ✅ define before use
               $("#IntegrationType").val(integrationType)

               if (Editintegration.is_active == 1) {
                $('#integration_activestatus').prop('checked', true).val(1);
            } else {
                $('#integration_activestatus').prop('checked', false).val('');
            }

            var $label = $("#GatewayName").prev("label");
            var $label2 = $("#GatewayDesc").prev("label");

   

            if (integrationType.toLowerCase() === "storage") {
                $('#AccessKeyId').val(Editintegration.credentials.credentials.access_key)
                $('#SecretKey').val(Editintegration.credentials.credentials.secret_key)
                $('#Region').val(Editintegration.credentials.credentials.region)
                $('#BucketName').val(Editintegration.credentials.credentials.bucket_name)

                $label.contents().first().replaceWith("Storage Name ");
                $label2.contents().first().replaceWith("Storage Desc ");
                $("#forpayment").addClass("hidden")
                $("#forstorage").removeClass("hidden")

            } else {
                $('#ClientId').val(Editintegration.credentials.credentials.access_key)
               
                $('#ClientSecret').val(Editintegration.credentials.credentials.secret_key)
                $label.contents().first().replaceWith("Gateway Name ");
                $label2.contents().first().replaceWith("Gateway Desc ");
                $("#forpayment").removeClass("hidden")
                $("#forstorage").addClass("hidden")

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



        // $(document).on('click', '#UpdateIntegration', function () {
        //     // Form Validation
    
        //     // console.log("checking::");
        //     // var formAction = $("#membershipform").attr('action');
        //     // var passwordRequired = (formAction !== '/membership/updatemember');
    
          
        //     $("#IntegrationForm").validate({
        //         ignore: [],
        //         rules: {
                    
        //             GatewayName: {
        //                 required: true,
        //                 space: true
        //             },
        //             GatewayDesc: {
        //                 required: true,
        //             },
        //             ClientId: {
        //                 required: true,
        //                 space: true,
        //             },
                    
        //             ClientSecret: {
        //                 required: true,
        //                 space: true,
        //             }
    
    
        //         },
        //         messages: {
                  
        //             GatewayName: {
        //                 required: "* " + "Please enter the Gateway Name",
        //                 space: "* " + languagedata.spacergx
        //             },
    
        //             GatewayDesc: {
        //                 required: "* " + " Please enter the Gateway Desc",
        //             },
    
        //             ClientId: {
        //                 required: "* " + "Please enter the Client Id",
        //                 space: "* " + languagedata.spacergx,
        //             },
        //             ClientSecret: {
        //                 required: "* " + "Please enter the Client Secret",
        //                 space: "* " + languagedata.spacergx,
        //             }
        //         }
        //     })
    
        //     var formcheck = $("#IntegrationForm").valid();
    
        //     if (formcheck) {
        //         $('#IntegrationForm')[0].submit();
        //     }
        //     else {
        //         console.log("form nt valid");
    
        //     }
    
        // })

        $(document).on('click', '#UpdateIntegration', function () {
            // Get integration type dynamically (from hidden input or variable)
            var integrationType = $("#IntegrationType").val() || ""; // e.g., "Payment" or "Storage"

            console.log("integrationType",integrationType);
            
        
            // First, clear old validation if any
            $("#IntegrationForm").validate().destroy();
        
            // Initialize validation
            $("#IntegrationForm").validate({
                ignore: [],
                rules: {
                    GatewayName: {
                        required: true
                    },
                    GatewayDesc: {
                        required: true
                    },
        
                    // Common payment fields
                    ClientId: {
                        required: function () {
                            return integrationType.toLowerCase() === "payment";
                        }
                    },
                    ClientSecret: {
                        required: function () {
                            return integrationType.toLowerCase() === "payment";
                        }
                    },
        
                    // Storage-specific fields
                    Region: {
                        required: function () {
                            return integrationType.toLowerCase() === "storage";
                        }
                    },
                    AccessKeyId: {
                        required: function () {
                            return integrationType.toLowerCase() === "storage";
                        }
                    },
                    SecretKey: {
                        required: function () {
                            return integrationType.toLowerCase() === "storage";
                        }
                    },
                    BucketName: {
                        required: function () {
                            return integrationType.toLowerCase() === "storage";
                        }
                    }
                },
        
                messages: {
                    GatewayName: {
                        required: "* Please enter the Gateway Name"
                    },
                    GatewayDesc: {
                        required: "* Please enter the Gateway Description"
                    },
        
                    ClientId: {
                        required: "* Please enter the Client ID"
                    },
                    ClientSecret: {
                        required: "* Please enter the Client Secret"
                    },
        
                    Region: {
                        required: "* Please enter the AWS region"
                    },
                    AccessKey: {
                        required: "* Please enter the Access Key"
                    },
                    SecretKey: {
                        required: "* Please enter the Secret Key"
                    },
                    BucketName: {
                        required: "* Please enter the Bucket Name"
                    }
                }
            });
        
            // Validate the form
            var formcheck = $("#IntegrationForm").valid();
        
            if (formcheck) {
                $('#IntegrationForm')[0].submit();
            } else {
                console.log("Form is not valid");
            }
        });
        


})
