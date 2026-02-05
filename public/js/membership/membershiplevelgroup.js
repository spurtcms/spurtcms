
var languagedata
var selectedcheckboxarr = []




$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
})

$(document).ready(function () {

    var $textarea = $('#plangroupdesc');
    var $errorMessage = $('#error-plangroupdesc');
    var maxLength = 250;

    $textarea.on('input', function () {

        console.log("checkdesc")
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessage.text("Maximum 250 character allowed").show();
        } else {
            // Clear error message if under the limit
            $errorMessage.text('');
        }
    });

    $(document).on('click', '#createmembershipgroup', function (e) {

        $("#membershipgroupcreateform").validate({
            ignore: [],
            rules: {

                plangroupname: {
                    required: true,
                    space: true
                },
                plangroupdesc: {
                    required: true,
                    maxlength: 250,
                    space: true
                }

            },
            messages: {

                plangroupname: {
                    required: "* " + "Please Enter the Membership Level group name",
                    space: "* " + languagedata.spacergx
                },
                plangroupdesc: {
                    required: "* " + "Please Enter the Membership Level Group Desc",
                    maxlength: "* " + "Maximum 250 character allowed",
                    space: "* " + languagedata.spacergx
                }

            }
        })

        e.preventDefault();

        if ($('#membershipgroupcreateform').valid()) {
            $('#membershipgroupcreateform')[0].submit();
        } else {
            console.log('Form is not valid');
        }

    })


   

    $('#addmember').click(function(){
        $('#error-plangroupdesc').hide()
        $('#plangroupname-error').hide()
        $('#plangroupdesc-error').hide()

        $('#plangroupname').val('')
        $('#plangroupdesc').val('')
        
    })


    $('#plangroupdesc').on('input', function () {

        let lines = $(this).val().split('\n').length;

        if (lines > 5) {

            let value = $(this).val();

            let linesArray = value.split('\n').slice(0, 5);

            $(this).val(linesArray.join('\n'));
        }
    });

    

    // Open membership Group Model

    $('.membershipgrouEdit').click(function () {



        edit = $(this).closest("tr");
        var name = edit.find("td:eq(1)").text().trim()
        var desc = edit.find("td:eq(2)").text().trim()
        var Id = $(this).attr("data-id");

      
        $('#membershipgroupcreateform').attr('action', '/admin/membershipgroup/updatemembershiplevelgroup')
        $("#plangroupname").val(name);
        $("#plangroupdesc").val(desc);
        $('#membershipgroupid').val(Id);
        $("#createmembershipgroup").text(languagedata.update)
        $('#modalTitleId').text(languagedata.memberships.membershipgroups.updateMembersiphgroup)

        // $(this).data('bs-toggle', 'modal').data('bs-target', '#modalId2')

    })

    // delete membership group flow

    $(document).on('click', '.MembersgipGroupdelete', function () {

        var MembershipLevelGroupId = $(this).attr("data-id")
        $(".deltitle").text(languagedata.memberships.membershipgroups.deletemembergroup)
        // $('.delname').text($(this).parents('tr').find('#membername').text())
        $("#content").text(languagedata.memberships.membershipgroups.areyousure)

        $('#delid').attr('href', '/admin/membershipgroup/deletemembershiplevelgroup?Id=' + MembershipLevelGroupId);
    })


// search funtionss

$(document).on('keyup', '#searchmembershiplevel', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/admin/membershipgroup/";
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
    window.location.href = "/admin/membershipgroup/"
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
// SEARCH END


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

        $('.deltitle').text("Delete Membership Level Groups?")

        $('#content').text('Are you sure want to delete selected Membership Level Groups?')

    } else {

        $('.deltitle').text("Delete Membership Level Group?")

        $('#content').text('Are you sure want to delete selected Membership Level Group?')
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
        url: '/admin/membershipgroup/multiselectdeletemembershipgroup',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "membershipgroupids": selectedcheckboxarr,
            csrf: $("input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            console.log(data.value,"data.value");
            

            if (data.value) {

                setCookie("get-toast", "MembershiplevelGroupDeletedSuccessfully")

                window.location.href = data.url
            } else {

                // setCookie("Alert-msg", "Internal Server Error")

                window.location.href = data.url

            }

        }
    })

})

// reay end 
})


// is active works


function MembershipLevelGroupStatus(id) {

    
    $('#check' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();

    var isactive = $('#check' + id).val()    


    $.ajax({
        url: '/admin/membershipgroup/membershipgroupisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >"Membership level group updated successfully"</p ></div ></div ></li></ul> `;
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


$(document).on('click','.statuscls',function(){

    statusval =$(this).text().trim()
    $(".filterleveldropdown").removeClass("show")
    $('.slctstatus').text(statusval)
    $('#statusid').val(statusval)
})