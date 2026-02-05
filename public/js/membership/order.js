var selectedcheckboxarr = []

$(document).ready(function () {

    $("#save").click(function () {


        $("#create-order").validate({
            ignore: [],
            rules: {
                userid: {
                    required: true,
                },
                membershiplevelid: {
                    required: true,
                },
                total: {
                    required: true,
                },
                status: {
                    required: true,
                },
                gateway: {
                    required: true,
                }
            
            },
            messages: {
                userid: {
                    required: "* Please Enter UserId",
                },
                membershiplevelid: {
                    required: "* Please Choose Membership Level",
                },
                total: {
                    required: "* Please Enter Valid Amount",
                },
                status: {
                    required: "* Please Select Status",
                },
                gateway: {
                    required: "* Please Select Gateway",
                },
            }
        });

        // Check if the form is valid
        var FormCheck = $("#create-order").valid();
        if (FormCheck) {
            $("#create-order").submit();
        }
    });



    //Default Values loaded

    $("#statusvalue").text("Success")

    $("#status").val("Success")

    $('.order-dropdownlist a').on('click', function () {

        var memberName = $(this).data('value');

        var subscription = $(this).data('subscription');

        $("#subscriptiontransactionid").val(subscription)
        
        var memberid = $(this).data('id');

        $('#searchmemberlists').text(memberName);

        $('#userid').val(memberid);

        $('#userid-error').addClass("hidden");
    });

    $('.search').on('keyup', function () {
        var found = false;
        var searchTerm = $(this).val().toLowerCase();
        $('.order-dropdownlist').filter(function () {
            var isVisible = $(this).text().toLowerCase().indexOf(searchTerm) > -1;
            $(this).toggle(isVisible);
            if (isVisible) found = true;
        });

        if (found) {
            $('.noData-foundWrapper').hide();
        } else {
            $('.noData-foundWrapper').show();
        }

    });


    $('.membershiplevelid .option').on('click', function () {

        var selectedValue = $(this).data('value');

        var selectedid = $(this).data('id');

        $('#membershiplevelidvalue').text(selectedValue);

        $('#membershiplevelid').val(selectedid);

        $('#membershiplevelid-error').addClass("hidden");

        $('.membershiplevelid').dropdown('hide');

    });

    $('.searchlevel').on('keyup', function () {
        var found = false;
        var searchTerm = $(this).val().toLowerCase();
        $('.membershiplevelid').filter(function () {
            var isVisible = $(this).text().toLowerCase().indexOf(searchTerm) > -1;
            $(this).toggle(isVisible);
            if (isVisible) found = true;
        });

        if (found) {
            $('.levelnoData-foundWrapper').hide();
        } else {
            $('.levelnoData-foundWrapper').show();
        }

    });


    //  Only allow numberonly not a text
    $("input[name='billingphone']").on('keypress', function (event) {
        // Allow only numbers (key codes for 0-9 are 48-57)
        if (event.which < 48 || event.which > 57) {
            event.preventDefault();
        }
    });


    $('.status .option').on('click', function () {

        var selectedValue = $(this).data('value');

        $('#statusvalue').text(selectedValue);

        $('#status').val(selectedValue);

        $('#status-error').addClass("hidden");

        $('.status').dropdown('hide');
    });

    $('.gateway .option').on('click', function () {

        var selectedValue = $(this).data('value');

        $('#gatewayvalue').text(selectedValue);

        $('#gateway').val(selectedValue);

        $('#gateway-error').addClass("hidden");

        $('.gateway').dropdown('hide');
    });

    $('.gatewayenvironment .option').on('click', function () {

        var selectedValue = $(this).data('value');

        $('#gatewayenvironmentvalue').text(selectedValue);

        $('#gatewayenvironment').val(selectedValue);

        $('.gatewayenvironment').dropdown('hide');
    });

    $(document).on('click', '#delete-btn', function () {

        var orderId = $(this).attr("data-id")
        $("#content").text(languagedata.memberships.orders.areyousure)
        var url = window.location.search
        const urlpar = new URLSearchParams(url)
        pageno = urlpar.get('page')

        if (pageno == null) {
            $('#delid').attr('href', "/admin/order/deleteorder/" + orderId);

        } else {
            $('#delid').attr('href', "/admin/order/deleteorder/" + orderId + "?page=" + pageno);

        }
        $(".deltitle").text(languagedata.memberships.orders.Deletemembershiporder)
        $('.delname').text($(this).parents('tr').find('td:eq(1)').text())

    })

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

            $('.deltitle').text(languagedata.memberships.orders.Deletemembershiporders)

            $('#content').text(languagedata.memberships.orders.areyousures)

        } else {

            $('.deltitle').text(languagedata.memberships.orders.Deletemembershiporder)

            $('#content').text(languagedata.memberships.orders.areyousure)
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
            url: '/admin/order/multiselectdeleteorder',
            type: 'post',
            dataType: 'json',
            async: false,
            data: {
                "orderids": selectedcheckboxarr,
                csrf: $("input[name='csrf']").val(),
                "page": pageno


            },
            success: function (data) {

                if (data.value == true) {

                    setCookie("get-toast", "Order Deleted Successfully")

                    window.location.href = data.url
                } else {

                    // setCookie("Alert-msg", "Internal Server Error")

                    window.location.href = data.url

                }

            }
        })

    })

    //Deselectall function//

    $(document).on('click', '#deselectid', function () {

        $('.selectcheckbox').prop('checked', false)

        $('#Check').prop('checked', false)

        selectedcheckboxarr = []

        $('.selected-numbers').addClass('hidden')

    })

    $(document).on("click", ".Closebtn", function () {
        $(".search").val('')
        $(".Closebtn").addClass("hidden")
        $(".SearchClosebtn").removeClass("hidden")
        $(".srchBtn-togg").removeClass("pointer-events-none")
    })

    $(document).on("click", ".searchClosebtn", function () {
        $(".search").val('')
        window.location.href = "/admin/order/"
    })

    $(document).ready(function () {

        $('.search').on('input', function () {
            if ($(this).val().length >= 1) {
                var value = $(".search").val();
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

    $(".filterlevel").on("click",function(){

        var levelvalue=$(this).data("value")

        $("#filterlevelspan").text(levelvalue)

        $("#filter-level").val(levelvalue)

        $(".filterleveldropdown").removeClass("show")


        
    })



})