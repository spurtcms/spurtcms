var languagedata
var selectedcheckboxarr = []


/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {


        languagedata = data
    })
})




$(document).ready(function () {


    // Mobile num valiadte

    jQuery.validator.addMethod("mob_validator", function (value, element) {
        if (value.length >= 7)
            return true;
        else return false;
    }, "*" + "Mobile number must contain 7 digits");

    // Email validate

    $.validator.addMethod("emails_validator", function (value) {
        return /(^[a-zA-Z_0-9\.-]+)@([a-z]+)\.([a-z]+)(\.[a-z]+)?$/.test(value);
    }, '* ' + " Please enter a valid email address");



    jQuery.validator.addMethod("pass_validator", function (value, element) {
        if (value != "") {
            if (/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[\W_]).{8,}$/.test(value))
                return true;
            else return false;
        }
        else return true;
    }, "* " + " Please Enter at Least 1 Uppercase, 1 Lowercase, 1 Number,1 Special Character($,@),and 8 characters long"
    );


    $(document).on('click', '#membershipcreatemember', function () {
        // Form Validation

        var formAction = $("#membershipform").attr('action');
        var passwordRequired = (formAction !== '/admin/membership/updatemember');

      
        $("#membershipform").validate({
            ignore: [],
            rules: {
                prof_pic: {
                    extension: "jpg|png|jpeg"
                },
                mem_name: {
                    required: true,
                    space: true
                },
                mem_email: {
                    required: true,
                    emails_validator: true,
                    // duplicateemail: true
                },
                mem_usrname: {
                    required: true,
                    space: true,
                    // duplicatename: true
                },
                mem_pass: {
                    required: passwordRequired,
                    pass_validator: passwordRequired,
                },
                mem_mobile: {
                    required: true,
                    mob_validator: true,
                    // duplicatenumber: true
                }


            },
            messages: {
                prof_pic: {
                    extension: "* " + languagedata.profextension
                },
                mem_name: {
                    required: "* " + "Please enter the First name",
                    space: "* " + languagedata.spacergx
                },

                mem_email: {
                    required: "* " + " Please enter the Email ID",
                    // duplicateemail: "* " + languagedata.Memberss.emailexist
                },

                mem_usrname: {
                    required: "* " + "Please enter the User name",
                    space: "* " + languagedata.spacergx,
                    // duplicatename: "*" + languagedata.Memberss.membernamevaild
                },
                mem_pass: {
                    required: "* " + "Please enter the Password",
                },
                mem_mobile: {
                    required: "* " + "Enter the mobile number",
                }

            }
        })

        var formcheck = $("#membershipform").valid();

        if (formcheck) {
            $('#membershipform')[0].submit();
        }
        else {
            console.log("form nt valid");

        }

    })


    // Edit

    $(document).on('click', '#EditmembershipMember', function () {

        $('#modalId2').modal('show');
        $('#modalTitleId').text("Update Member")
        $('#membershipcreatemember').text("Update")
       
        var member_id = $(this).attr('data-id')

        $.ajax({
            url: "/admin/membership/EditMember",
            type: "GET",
            async: false,
            data: { "id": member_id },
            datatype: "json",
            caches: false,
            success: function (data) {

                var Eidtmember = data.EditMember
                var EndPoint = data.endpoint

                $('#mem_name').val(Eidtmember.FirstName)                
                $('#mem_lname').val(Eidtmember.LastName)
                $('#mem_email').val(Eidtmember.Email)
                $('#mem_usrname').val(Eidtmember.Username)
                // $('#mem_pass').val(Eidtmember.Password)
                $('#mem_mobile').val(Eidtmember.MobileNo)

                $('#membershipform').attr('action', EndPoint);

                $('#Update_memberid').val(Eidtmember.Id)

                if (Eidtmember.IsActive == 1) {
                    $('#mem_activestatus').prop('checked', true).val(1);
                } else {
                    $('#mem_activestatus').prop('checked', false).val('');
                }


            }
        })


    })


    // Is_active

    $("input[name=mem_activestatus]").click(function () {
        if ($(this).prop('checked')) {
            $(this).attr("value", "1")
        } else {
            $(this).removeAttr("value")
        }
    })


    // click create member then 

    $(document).on('click', '#AddnewMember,#AddnewMembers', function () {        

        $('#mem_activestatus').prop('checked', true).val('1')

        $('#modalTitleId').text("Add Member")
        $('#membershipcreatemember').text("Save")
        $('#modalId2').modal('show')

        $('#membershipform').attr('action', '/admin/membership/newmember')

        $('#mem_name').val("")
        $('#last_name').val('')
        $('#mem_email').val('')
        $('#mem_usrname').val('')
        $('#mem_pass').val('')
        $('#mem_mobile').val('')

        $('#mem_name-error').hide()
        $('#mem_email-error').hide()
        $('#mem_usrname-error').hide()
        $('#mem_pass-error').hide()
        $('#mem_mobile-error').hide()
        $('.lengthErr').addClass('hidden')



        $('#Update_memberid').val('')

    })

//cancel button empty error func//

$(document).on('click','.cancelbtn',function(){

    $('#mem_name').val("")
    $('#last_name').val('')
    $('#mem_email').val('')
    $('#mem_usrname').val('')
    $('#mem_pass').val('')
    $('#mem_mobile').val('')

    $('#mem_name-error').hide()
    $('#mem_email-error').hide()
    $('#mem_usrname-error').hide()
    $('#mem_pass-error').hide()
    $('#mem_mobile-error').hide()
    $('.lengthErr').addClass('hidden')
    $('#membershipform').attr('action',"")
    $("#membershipform")[0].reset();

})
    // Delete work flow 

    $(document).on('click', '#DeleteMembershipMember', function () {
        var MemberId = $(this).attr("data-id")
        $(".deltitle").text(languagedata.memberships.member.deletemember)
        // $('.delname').text($(this).parents('tr').find('#membername').text())
        $("#content").text(languagedata.memberships.member.areyousure)

        $('#delid').attr('href', '/admin/membership/deletemembershipmember?Id=' + MemberId);

    })


    // pasword show and hide

    $(document).on('click', '#eye', function () {

        var This = $("#mem_pass")

        if ($(This).attr('type') === 'password') {

            $(This).attr('type', 'text');

            $(this).addClass('open')

            $("#img").attr("src", "/public/img/eye-opened.svg")

        } else {

            $(This).attr('type', 'password');

            $(this).removeClass('close')

            $("#img").attr("src", "/public/img/eye-closed.svg")

        }
    })


    // only allow numbers (mobile number validation)

    $('#mem_mobile').keyup(function () {
        this.value = this.value.replace(/[^0-9\.]/g, '');
    });


    // search funtionss


    $(document).on('keyup', '#searchmembershiplevel', function (event) {
        const searchInput = $(this).val();

        if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

            if ($(this).val() == "") {

                window.location.href = "/admin/membership/";
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
        window.location.href = "/admin/membership/"
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

            $('.deltitle').text(languagedata.memberships.member.deletemembers)

            $('#content').text(languagedata.memberships.member.areyousures)

        } else {

            $('.deltitle').text(languagedata.memberships.member.deletemember)

            $('#content').text(languagedata.memberships.member.areyousure)
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
            url: '/admin/membership/multiselectdeletemember',
            type: 'post',
            dataType: 'json',
            async: false,
            data: {
                "membersids": selectedcheckboxarr,
                csrf: $("input[name='csrf']").val(),
                "page": pageno
            },
            success: function (data) {

                console.log(data.value, "data.value");


                if (data.value) {

                    setCookie("get-toast", "MembershipMemberUpdaredSuccessfully")

                    window.location.href = data.url
                } else {

                    // setCookie("Alert-msg", "Internal Server Error")

                    window.location.href = data.url

                }

            }
        })

    })





    // ready end 
})

// is active works


function MembershipStatus(id) {
    $('#toggleTwo' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();

    var isactive = $('#toggleTwo' + id).val();


    $.ajax({
        url: '/admin/membership/membershipisactive',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Member updated successfully</p ></div ></div ></li></ul> `;
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






$(document).on('keypress', '.mobileInput', function () {
    var inputVal = $(this).val()

    var inputLength = inputVal.length

    if (inputLength == 15) {
        $(this).siblings('.lengthErr').removeClass('hidden')
        $(this).prop('readonly', true)
    }


})


$(document).on('keyup', '.mobileInput', function (e) {
    

    if (e.which == 8) {
        
        var inputVal = $(this).val()

        var inputLength = inputVal.length

        if (inputLength <= 15) {
            $(this).siblings('.lengthErr').addClass('hidden')
            $(this).prop('readonly', false)
        }
    }
})

$(document).on('click',".filterlevel",function(){

   

    var levelvalue=$(this).data("value")

    $("#filterlevelspan").text(levelvalue)

    $("#filter-level").val(levelvalue)

    $(".filterleveldropdown").removeClass("show")


    
})