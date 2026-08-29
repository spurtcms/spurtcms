$(document).ready(function () {

    $(".language").click(function () {

        let id = $(this).data("id")

        let text = $(this).text().trim();

        $("#selectLanguage").html(text)

        $("#langId").val(id)

        $(".error").addClass("hidden")

    })

    $("#Next").click(function () {

        $("#LanguageList").validate({

            ignore: [],

            rules: {
                langId: {
                    required: true
                }
            },
            messages: {
                langId: {
                    required: "* Please Choose a Language" 
                }
            }
        })

        var formcheck = $("#LanguageList").valid();
        if (formcheck == true) {
            $('#LanguageList')[0].submit();

        }else{
            $(".error").addClass("[&+label]:text-[#F26674] [&+label]:font-normal [&+label]:text-xs rounded-[4px]")
        }


    })

})