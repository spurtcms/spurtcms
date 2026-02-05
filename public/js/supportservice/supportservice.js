var languagedata

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
})


$(document).ready(function () {
    $('#timezone-view').addClass('text-bold-gray').removeClass('text-bold');

    // $('#radio').prop('checked', true);
    $('#services').val('Integrate spurtCMS seamlessly into your project with our expert assistance.');

})

$(document).ready(function () {

    $(document).on('click', '#supportformsubmit', function () {


        // suppot email validate regex
        $.validator.addMethod("email_validator", function (value) {
            return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
        }, '* ' + languagedata.Memberss.mememailrgx);


        // content number validate regex 
        jQuery.validator.addMethod("mob_validator", function (value, element) {
            if (value.length >= 7)
                return true;
            else return false;
        }, "* " + languagedata.Memberss.memmobnumrgx);


        $.validator.addMethod("noSpaceStart", function (value, element) {
            return this.optional(element) || value.trim() === value;
        }, "The first letter cannot be a space.");

        $("#supportserviceform").validate({
            ignore: [],
            rules: {
                companyname: {
                    required: true,
                },
                contectemail: {
                    required: true,
                    email_validator: true,
                },
                contectnumber: {
                    required: true,
                    mob_validator: true,
                },
                // timezone: {
                //     required: true,
                // },
                Country: {
                    required: true,
                },
                // describe: {
                //     required: true,
                // }

            },
            messages: {

                companyname: {
                    required: "* " + languagedata.Support.nameerror,
                },

                contectemail: {
                    required: "* " + languagedata.Support.emailerror,
                },
                contectnumber: {
                    required: "* " + languagedata.Support.contacterror,
                },
                // timezone: {
                //     required: "* " + languagedata.Support.timezoneerror,
                // },
                Country: {
                    required: "* " + languagedata.Support.countryerror,
                },
                // describe: {
                //     required: "* " + languagedata.Support.decerror,
                // }

            }
        })

        var formcheck = $("#supportserviceform").valid();

        if (formcheck) {
            $('#supportserviceform')[0].submit();
        } else {
            console.log("support geting error");
        }

    })


    // function togglePlanClasses(activePlan, inactivePlan) {
    //     $(activePlan).removeClass('bg-transparent').addClass('bg-[#262626] text-white');
    //     $(inactivePlan).removeClass('bg-[#262626] text-white').addClass('bg-transparent');
    // }
    
    // $('#Yearlyplan').click(function(){
    //     togglePlanClasses('#Yearlyplan', '#monthlyplan');
    //     $('#subscriptionamount').text('$799')
    //     $('#subscriptionintervel').text('/year')
    //     // $('#Billedintervel').text('Billed yearly')
    //     $('#interval').val('year')

    // });
    
    // $('#monthlyplan').click(function(){
    //     togglePlanClasses('#monthlyplan', '#Yearlyplan');
    //     $('#subscriptionamount').text('$99')
    //     $('#subscriptionintervel').text('/month')
    //     // $('#Billedintervel').text('Billed monthly')
    //     $('#interval').val('month')
    // });

    


    $('#contectnumber').on('keypress', function (event) {
        // Allow only numbers (key codes for 0-9 are 48-57)
        if (event.which < 48 || event.which > 57) {
            event.preventDefault();
        }
    });



$(document).on('click','#CountrySelect',function(){
    
    var selectedCountry = $(this).val();  
    $('#Country').val(selectedCountry)
})


    // dropdown Time zone filter input box search
    $("#searchtimezone").keyup(function () {

        var keyword = $(this).val().trim().toLowerCase()
        $(".timezonelist").each(function (index, element) {
            var title = $(element).text().toLowerCase()

            if (title.includes(keyword)) {
                $(element).show()
                $("#nodatafounddesign").addClass("hidden")

            } else {
                $(element).hide()
                if ($('.timezonelist:visible').length == 0) {
                    $("#nodatafounddesign").removeClass("hidden")
                }

            }
        })

    })


    // Select types of support service

    $(document).on('click', '.service-type', function () {
        var Service = $(this).text().trim()
        $('#services').val(Service)
        console.log($('#services').val());

    })


      // only allow numbers (content number validation)
      $('#contectnumber').keyup(function () {
        this.value = this.value.replace(/[^0-9\.]/g, '');
    });


// ----------------
})





