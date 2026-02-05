
var languagedata
var selectedcheckboxarr = []

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
})

$(document).ready(function () {

    // inital values 

    $('#billingfrequenttype').val(30)

    //

    // Get Default Membership Template and load

    // $(document).on('click', '#loadtemplate', function () {

    //     var TemplateId = $(this).attr('data-id');

    //     $.ajax({
    //         url: '/membershiplevel/loadmembershiptemplate',
    //         type: 'GET',
    //         dataType: 'json',
    //         cache: false,
    //         data: {
    //             "membershipid": TemplateId
    //         },
    //         contentType: "application/json",
    //         success: function (data) {

    //             var MembershipLevel = data.SelectedMembership[0]


    //             $('#subscriptionname').val(MembershipLevel.SubscriptionName)
    //             $('#subscriptiondesc').val(MembershipLevel.MembershiplevelDetails)
    //             $('#subscriptiondescription').val(MembershipLevel.Description)
    //             $('#initialpayment').val(MembershipLevel.InitialPayment)
    //             if (MembershipLevel.RecurrentSubscription === 1) {
    //                 $('#isrecurring').prop('checked', true);
    //                 $('#recurringsubscriptiondev').removeClass('hidden')
    //             } else {
    //                 $('#isrecurring').prop('checked', false);
    //             }


    //             $('#billingamount').val(MembershipLevel.BillingAmount)
    //             $('#billingfrequentvalue').val(MembershipLevel.BillingfrequentValue)
    //             $('#billingfrequenttype').val(MembershipLevel.BillingfrequentType)
    //             $('#billingcyclelimit').val(MembershipLevel.BillingCyclelimit)


    //             if (MembershipLevel.CustomTrial === 1) {
    //                 $('#customtrial').prop('checked', true);
    //                 $('#customtraildiv').removeClass('hidden')

    //             } else {
    //                 $('#customtrial').prop('checked', false);
    //                 $('#customtraildiv').addClass('hidden')
    //             }
    //             $('#trirbillingamount').val(MembershipLevel.TrialBillingAmount)
    //             $('#trialbillinglimit').val(MembershipLevel.TrialBillingLimit)

    //         }
    //     })

    // })


    // discount calculation //
    $(document).on('click', '.Defaulltdescount', function () {
        var DiscountPercent = $(this).attr('data-value');
        var initialpayment = $('#initialpayment').val().trim();

        if (initialpayment !== "" && !isNaN(initialpayment) && !isNaN(DiscountPercent)) {
            initialpayment = parseFloat(initialpayment);
            DiscountPercent = parseFloat(DiscountPercent);
            $('#discountpercentage').val(DiscountPercent)

            var Discount = initialpayment * (DiscountPercent / 100);

            var DiscountedPrice = initialpayment - Discount;

            $('#discountedprice').val(DiscountedPrice.toFixed(2));
        } else {
            $('#error-initailpayment').removeClass('hidden')
            // console.log("Please Enter Initial Payment.");
        }
    });
    
    $(document).on('input', '#discountpercentage', function () {
        var DiscountPercent = $(this).val().trim();
        var initialpayment = $('#initialpayment').val().trim();

        if(DiscountPercent == ""){
            $('#discountedprice').val('');
        }

        if (initialpayment !== "" && DiscountPercent !== "" && !isNaN(initialpayment) && !isNaN(DiscountPercent)) {
            initialpayment = parseFloat(initialpayment);
            DiscountPercent = parseFloat(DiscountPercent);

            var Discount = initialpayment * (DiscountPercent / 100);

            var DiscountedPrice = initialpayment - Discount;

            $('#discountedprice').val(DiscountedPrice.toFixed(2));
        } else if(initialpayment==''){
            $('#error-initailpayment').removeClass('hidden')
            // console.log("Please Enter Initial Payment.");
        }
    });



    $(document).on('input', '#initialpayment', function () {
        $('#error-initailpayment').addClass('hidden')

        if($(this).val()==''){
         $('#discountedprice').val('')
         $('#discountpercentage').val('');
        }
    })




    // ----- //

    $.validator.addMethod("decimal_validator", function (value) {
        return /^\d{1,5}(\.\d{1,2})?$/.test(value);
    }, '* ' + "Give Valid input");


    //  Only allow numberonly not a text
    $('#initialpayment, #billingamount, #trirbillingamount, #billingfrequentvalue, #trialbillinglimit,#discountpercentage').on('keypress', function (event) {
        // Allow only numbers (key codes for 0-9 are 48-57)
        if (event.which < 48 || event.which > 57) {
            event.preventDefault();
        }
    });


    // Custom validation rule to ensure trialbillinglimit <= billingcyclelimit
    $.validator.addMethod("billingLimitCheck", function (value, element) {
        var billingcyclelimit = parseFloat($('#billingcyclelimit').val());
        var trialbillinglimit = parseFloat(value);
        // Return true if trialbillinglimit is less than or equal to billingcyclelimit, false otherwise
        return this.optional(element) || trialbillinglimit <= billingcyclelimit;
    }, "Trial Billing Limit must be less than or equal to Billing Cycle Limit.");


    // Create Membership level 

    $(document).on('click', '#membershipsave', function (e) {


        $("#membershipcreateform").validate({
            ignore: [],
            rules: {

                subscriptionname: {
                    required: true,
                    space: true
                },
                subscriptiondescription: {
                    required: true,
                    space: true
                },
                subscriptiondesc: {
                    required: true,
                    space: true
                },
                initialpayment: {
                    required: true,
                    decimal_validator: true

                },
                discountpercentage: {
                    required: {
                        depends: function () {
                            return $('#Discount').is(':checked')
                        }
                    }
                },
                billingamount: {
                    required: {
                        depends: function () {
                            return $('#isrecurring').is(':checked')
                        }
                    },
                    decimal_validator: {
                        depends: function () {
                            return $('#isrecurring').is(':checked')
                        }
                    }
                },
                billingfrequentvalue: {
                    required: {
                        depends: function () {
                            return $('#isrecurring').is(':checked')
                        }
                    }
                },
                billingfrequenttype: {
                    required: {
                        depends: function () {
                            return $('#isrecurring').is(':checked')
                        }
                    }
                },
                billingcyclelimit: {
                    required: {
                        depends: function () {
                            return $('#isrecurring').is(':checked')
                        }
                    }
                },
                trirbillingamount: {
                    required: {
                        depends: function () {
                            return $('#customtrial').is(':checked')
                        }
                    },
                    decimal_validator: {
                        depends: function () {
                            return $('#customtrial').is(':checked')
                        }
                    }
                },
                trialbillinglimit: {
                    required: {
                        depends: function () {
                            return $('#customtrial').is(':checked')
                        }
                    },
                    billingLimitCheck: true
                },


            },
            messages: {

                subscriptionname: {
                    required: "* " + "Please Enter the Membership Level name",
                    space: "* " + languagedata.spacergx
                },
                subscriptiondescription: {
                    required: "* " + "Please Enter the Membership Level Description",
                    space: "* " + languagedata.spacergx
                },
                subscriptiondesc: {
                    required: "* " + "Please Enter the Membership Level Details",
                    space: "* " + languagedata.spacergx
                },

                initialpayment: {
                    required: "* " + "Please Enter the initialpayment",
                },
                discountpercentage: {
                    required: "* " + "Please Enter the Discount %",
                },
                billingamount: {
                    required: "* " + "Please Enter the billingamount",
                },
                billingfrequentvalue: {
                    required: "* " + "Please Enter the billingfrequentvalue",
                },

                billingfrequenttype: {
                    required: "* " + "Please Enter the billingfrequenttype",
                },
                billingcyclelimit: {
                    required: "* " + "Please Enter the billingcyclelimit",
                },
                trirbillingamount: {
                    required: "* " + "Please Enter the trirbillingamount",
                },
                trialbillinglimit: {
                    required: "* " + "Please Enter the trialbillinglimit",
                }

            }

        })

        e.preventDefault();

        if ($('#membershipcreateform').valid()) {
            $('#membershipcreateform')[0].submit();
        } else {
            console.log('Form is not valid');
        }
    });

    // ----- //




    // Hide and show in recurrent subscription and custom trail input fields

    // $('#isrecurring').click(function(){

    $(document).on('click', '#isrecurring', function () {

        if ($(this).prop('checked')) {
            $('#recurringsubscriptiondev').removeClass('hidden')
            $(this).val(1)

        } else {
            $(this).val('')
            $('#recurringsubscriptiondev').addClass('hidden')
        }

    })

    $(document).on('click', '#Discount', function () {
        if ($(this).prop('checked')) {
            $('#Discountdiv').removeClass('hidden')
            $(this).val(1)

        } else {
            $(this).val('')
            $('#Discountdiv').addClass('hidden')
        }
    })

    $(document).on('click', '#customtrial', function () {

        if ($(this).prop('checked')) {
            $('#customtraildiv').removeClass('hidden')
            $(this).val(1)

        } else {
            $(this).val("")

            $('#customtraildiv').addClass('hidden')
        }

    })

    $(document).on('click', '.Membershipdeletes', function () {
        var MembershipLevelId = $(this).attr("data-id")
        $(".deltitle").text(languagedata.memberships.membershiplevels.deletemembershiplevel)
        // $('.delname').text($(this).parents('tr').find('#membername').text())
        $("#content").text(languagedata.memberships.membershiplevels.areyousure)

        $('#delid').attr('href', '/admin/membershiplevel/deletemembershiplevel?Id=' + MembershipLevelId);
    })




    // search funtionss


    $(document).on('keyup', '#searchmembershiplevel', function (event) {
        const searchInput = $(this).val();

        if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

            if ($(this).val() == "") {

                window.location.href = "/admin/membershiplevel/";
            }
        }

        $('.searchClosebtn').toggleClass('hidden', searchInput === "");

    })


    $(document).on("click", ".Closebtn", function () {
        $(".Searchmembership").val('')
        $(".Closebtn").addClass("hidden")
        $(".SearchClosebtn").removeClass("hidden")
        $(".srchBtn-togg").removeClass("pointer-events-none")
    })

    $(document).on("click", ".searchClosebtn", function () {
        $(".Searchmembership").val('')
        window.location.href = "/admin/membershiplevel/"
    })

    $(document).ready(function () {

        $('.Searchmembership').on('input', function () {
            if ($(this).val().length >= 1) {
                $(".Closebtn").removeClass("hidden")
                $(".srchBtn-togg").addClass("pointer-events-none")
                $(".SearchClosebtn").addClass("hidden")
            } else {
                $(".SearchClosebtn").removeClass("hidden")
                $(".Closebtn").addClass("hidden")
                $(".srchBtn-togg").removeClass("pointer-events-none")
            }
        });
    })
    $(document).on("click", ".SearchClosebtn", function () {
        $(".SearchClosebtn").addClass("hidden")
        $(".transitionSearch").removeClass("w-[300px] justify-start p-2.5 border border-[#ECECEC] rounded-sm gap-3 overflow-hidden")
        $(".transitionSearch").addClass("w-[32px]")


    })

    $(document).on("click", ".searchopen", function () {

        $(".SearchClosebtn").removeClass("hidden")

    })
    // search end 


    // Membership time select flow 

    $(document).on('click', '.PlanvalidTime', function () {
        var MembershipDuration = $(this).data('id')
        var time = $(this).text()

        console.log(time, "time");


        if (MembershipDuration == "Day(s)") {

            $('#billingfrequenttype').val(1)
            $('#PlanTimeShow').text(time)
        } else if (MembershipDuration == "Week(s)") {
            $('#billingfrequenttype').val(7)
            $('#PlanTimeShow').text(time)
        }
        else if (MembershipDuration == "Month(s)") {
            $('#billingfrequenttype').val(30)
            $('#PlanTimeShow').text(time)
        } else if (MembershipDuration == "Year(s)") {
            $('#billingfrequenttype').val(365)
            $('#PlanTimeShow').text(time)
        }

    })





    // select checkbox delete flow 


    $(document).on('click', '.selectcheckbox', function () {

        $('#unbulishslt').hide()
        $('#seleccheckboxdelete').removeClass('border-r');

        orderid = $(this).attr('data-id')


        if ($(this).prop('checked')) {

            selectedcheckboxarr.push(orderid)

        } else {

            const index = selectedcheckboxarr.indexOf(orderid);

            if (index !== -1) {

                selectedcheckboxarr.splice(index, 1);
            }

            $('#Check').prop('checked', false)

        }


        if (selectedcheckboxarr.length != 0) {

            $('.selected-numbers').removeClass('hidden')

            var items

            if (selectedcheckboxarr.length == 1) {

                items = "Item Selected"
            } else {

                items = languagedata.itemselected
            }

            $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

            $('#seleccheckboxdelete').removeClass('border-end')

            $('.unbulishslt').html("")

            if (selectedcheckboxarr.length > 1) {

                $('#deselectid').text("Deselect All")

            } else {
                $('#deselectid').text("Deselect")
            }

        } else {

            $('.selected-numbers').addClass('hidden')
        }

        var allChecked = true;

        $('.selectcheckbox').each(function () {

            if (!$(this).prop('checked')) {

                allChecked = false;

                return false;
            }
        });

        $('#Check').prop('checked', allChecked);

    })

    //ALL CHECKBOX CHECKED FUNCTION//

    $(document).on('click', '#Check', function () {


        $('#unbulishslt').hide()
        $('#seleccheckboxdelete').removeClass('border-r');

        selectedcheckboxarr = []

        var isChecked = $(this).prop('checked');

        if (isChecked) {

            $('.selectcheckbox').prop('checked', isChecked);

            $('.selectcheckbox').each(function () {

                orderid = $(this).attr('data-id')

                selectedcheckboxarr.push(orderid)
            })

            $('.selected-numbers').removeClass('hidden')

            var items

            if (selectedcheckboxarr.length == 1) {

                items = "Item Selected"
            } else {

                items = languagedata.itemselected
            }
            $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

        } else {


            selectedcheckboxarr = []

            $('.selectcheckbox').prop('checked', isChecked);

            $('.selected-numbers').addClass('hidden')
        }

        if (selectedcheckboxarr.length == 0) {

            $('.selected-numbers').addClass('hidden')
        }


    })


    $(document).on('click', '#seleccheckboxdelete', function () {

        if (selectedcheckboxarr.length > 1) {

            $('.deltitle').text("Delete Membership Levels?")

            $('#content').text('Are you sure want to delete selected Membership Levels?')

        } else {

            $('.deltitle').text("Delete Membership Level?")

            $('#content').text('Are you sure want to delete selected Membership Level?')
        }


        $('#delid').addClass('checkboxdelete')
    })


    //MULTI SELECT DELETE FUNCTION//
    $(document).on('click', '.checkboxdelete', function () {


        var url = window.location.href;

        var pageurl = window.location.search

        const urlpar = new URLSearchParams(pageurl)

        pageno = urlpar.get('page')


        $('.selected-numbers').hide()
        $.ajax({
            url: '/admin/membershiplevel/multiselectdeletemembershiplevel',
            type: 'post',
            dataType: 'json',
            async: false,
            data: {
                "membershiplevelids": selectedcheckboxarr,
                csrf: $("input[name='csrf']").val(),
                "page": pageno
            },
            success: function (data) {

                console.log(data.value, "data.value");


                if (data.value) {

                    setCookie("get-toast", "Category Group Deleted Successfully")

                    window.location.href = data.url
                } else {

                    // setCookie("Alert-msg", "Internal Server Error")

                    window.location.href = data.url

                }

            }
        })

    })


    // doc load end
})


function MembershipLevelStatus(id) {


    $('#check' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();

    var isactive = $('#check' + id).val()


    $.ajax({
        url: '/admin/membershiplevel/membershiplevelisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            console.log("working::");

            if (result.value) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >"Membership level updated successfully"</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            } else {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }
        }
    });

}

$(document).on('click', ".filterlevel", function () {

    var levelvalue = $(this).data("value")

    $("#filterlevelspan").text(levelvalue)

    $("#filter-level").val(levelvalue)

    $(".filterleveldropdown").removeClass("show")



})


// Select membergroup

function SelectMembershipGroup(id, groupname) {

    $('#subscriptiongroupid').val(id)

    $('#membershipgrouptext').text(groupname)

}