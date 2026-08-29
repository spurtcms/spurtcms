$(document).ready(function () {

    $("#searchcatlists").keyup(function () {
        var found = false;
        var searchTerm = $(this).val().trim().toLowerCase()
        $(".countrieslist").filter(function () {
            var isVisible = $(this).text().toLowerCase().indexOf(searchTerm) > -1;
            $(this).toggle(isVisible);
            if (isVisible) found = true;
        })
        if (found) {
            $('.noData-foundWrapper').hide();
        } else {
            $('.noData-foundWrapper').show();
        }

    })

    $(".countrieslist").click(function () {

        let text = $(this).text().trim();

        $("#selectedCountry").html(text)

        $("#country").val(text)

        $(".error").addClass("hidden")


    })

    $("#Next").click(function () {

        $("#CountryList").validate({

            ignore: [],

            rules: {
                country: {
                    required: true
                }
            },
            messages: {
                country: {
                    required: "* Please Choose a Country"
                }
            }
        })

        var formcheck = $("#CountryList").valid();
        if (formcheck == true) {
            $('#CountryList')[0].submit();

        } else {
            $(".error").addClass("[&+label]:text-[#F26674] [&+label]:font-normal [&+label]:text-xs rounded-[4px]")
        }


    })



})