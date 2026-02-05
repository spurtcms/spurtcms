 
var selectedcheckboxarr = []


$(document).ready(function(){


    $(document).on('click', '#Createsubscription', function (e) {


        $("#Createsubscriptionform").validate({
            ignore: [],
            rules: {

                gateway: {
                    required: true,
                   
                },
                gatewayenvironment: {
                    required: true,
                   
                },
                userid: {
                    required: true,
                   
                },
                membershiplevelid: {
                    required: true,
                   
                },
                subscriptiontransactionid: {
                    required: true,
                   
                }
         
                
            },
            messages: {

                gateway: {
                    required: "* " + "Please Select the payment gateway",
                   
                },
                gatewayenvironment: {
                    required: "* " + "Please Select the payment gateway environment",
                   
                },
                userid: {
                    required: "* " + "Please select the Member",
                   
                },
                membershiplevelid: {
                    required: "* " + "Please select the Membership Level ",
                   
                },
                subscriptiontransactionid: {
                    required: "* " + "Please Enter the subscription transaction id",
                   
                }

         

            }
        })

        e.preventDefault();

        if ($('#Createsubscriptionform').valid()) {
            $('#Createsubscriptionform')[0].submit();
        } else {
            console.log('Form is not valid');
        }
    });



    // membership userid search 
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


    $('.membershiplevelid').on('click', function () {

        var selectedValue = $(this).data('value');

        var selectedid = $(this).data('id');

        $('#membershiplevelidvalue').text(selectedValue);

        $('#membershiplevelid').val(selectedid);

        $('#membershiplevelid-error').addClass("hidden");

        $('.dropdown').dropdown('hide');
    });


    $('.order-dropdownlist a').on('click', function () {

        var memberName = $(this).data('value');

        var memberid = $(this).data('id');

        $('#searchmemberlists').text(memberName);

        $('#userid').val(memberid);

        $('#userid-error').addClass("hidden");
    });



    $('.gateway p').on('click', function () {

        var selectedValue = $(this).find(".option").data('value');        

        $('#gatewayvalue').text(selectedValue);

        $('#gateway').val(selectedValue);

        $('#gateway-error').addClass("hidden");

        $('.dropdown').dropdown('hide');
    });

    $('.gatewayenvironment p').on('click', function () {

        var selectedValue =  $(this).find(".option").data('value');

        $('#gatewayenvironmentvalue').text(selectedValue);

        $('#gatewayenvironment').val(selectedValue);

        $('.dropdown').dropdown('hide');
    });



// search funtionss


$(document).on('keyup', '#searchsubscription', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/admin/subscription/";
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
    window.location.href = "/admin/subscription/"
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


// multi select delete

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



 //MULTI SELECT DELETE FUNCTION//
 $(document).on('click', '.checkboxdelete', function () {
    

    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')


    $('.selected-numbers').hide()
    $.ajax({
        url: '/admin/subscription/multiselectdeletesubscription',
        type: 'post',
        dataType: 'json',
        async: false,
        cache: false,
        data: {
            "subsriptionids": selectedcheckboxarr,
            csrf: $("input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            console.log(data.value,"data.value");
            

            if (data.value) {

                setCookie("get-toast", "SubscriptionUpdaredSuccessfully")

                window.location.href = data.url
            } else {

                // setCookie("Alert-msg", "Internal Server Error")

                window.location.href = data.url

            }

        }
    })

})



$(document).on('click', '#seleccheckboxdelete', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text(languagedata.memberships.subscriptions.Deletesubscriptions)

        $('#content').text(languagedata.memberships.subscriptions.areyousures)

    } else {

        $('.deltitle').text(languagedata.memberships.subscriptions.Deletesubscription)

        $('#content').text(languagedata.memberships.subscriptions.areyousure)
    }


    $('#delid').addClass('checkboxdelete')
})


    // Delete work flow 

    $(document).on('click','.DeleteSubscription',function(){
        var SubscriptionId = $(this).attr("data-id")
        $(".deltitle").text(languagedata.memberships.subscriptions.Deletesubscription)
        // $('.delname').text($(this).parents('tr').find('#membername').text())
        $("#content").text(languagedata.memberships.subscriptions.areyousures)

        $('#delid').attr('href', '/admin/subscription/deletesubscription?Id=' + SubscriptionId);

})


})


function SubscriptionStatus(id) {

    
    $('#check' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();

    var isactive = $('#check' + id).val()    


    $.ajax({
        url: '/admin/subscription/subscriptionisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            
            if (result.value) {
        
                setCookie("get-toast", "SubscriptionUpdaredSuccessfully")

                window.location.href = data.url
            } else {

                window.location.href = data.url
            }
        }
    });


}

$(".filterlevel").on("click",function(){

    var levelvalue=$(this).data("value")

    $("#filterlevelspan").text(levelvalue)

    $("#filter-level").val(levelvalue)

    $(".filterleveldropdown").removeClass("show")

})

$('#subscriptiontransactionid').on('input', function() {
    // Get the input value and remove non-digit characters
    let value = $(this).val().replace(/[^0-9]/g, '');
    // Limit the value to 10 digits
    if (value.length > 10) {
        value = value.slice(0, 10);
    }
    $(this).val(value);
    
    // Show or hide the error message
    if (value.length === 10) {
        $('#error-message').show();
    } else {
        $('#error-message').hide();
    }
});