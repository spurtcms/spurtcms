var languagedata

var selectedcheckboxarr = []

/**This if for Language json file */

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data

    })

    $('#roledesc').on('input', function () {

        let lines = $(this).val().split('\n').length;

        if (lines > 5) {

            let value = $(this).val();

            let linesArray = value.split('\n').slice(0, 5);

            $(this).val(linesArray.join('\n'));
        }
    });

    var $textarea = $('#roledesc');
    var $errorMessage = $('#roledesc-error');
    var maxLength = 250;

    $textarea.on('input', function () {
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessage.text(languagedata.Permission.descriptionchat);
        } else {
            // Clear error message if under the limit
            $errorMessage.text('');
        }
    });
});


// cancel btn function

$("#newroles").on("hide.bs.modal", function () {
    console.log("calcel btn work");

    $("#rolename-error").hide();
    $("#roledesc-error").hide();
    $("#roledesc").val("")
    $("#rolename").val("")
})


// search input focus function

$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
});


/**Save & Update Btn*/

$(document).on('click', '.saverolperm', function () {
    setTimeout(function () {
    }, 500);

    $("#roleform").validate({
        rules: {
            name: {
                required: true,
                space: true,
                maxlength: 100,

            },
            description: {
                required: true,
                space: true,
                maxlength: 250,
            }
        },
        messages: {
            name: {
                required: "* " + languagedata.Roless.rolenamevalid,
                space: "* " + languagedata.spacergx,
                maxlength: "* " + languagedata.Permission.rolechat,
            },
            description: {
                required: "* " + languagedata.Roless.roledescvalid,
                maxlength: "* " + languagedata.Permission.descriptionchat

            }
        }
    });

    var formcheck = $("#roleform").valid();
    if (formcheck == true) {
        var id = $("#rolid").val()
        var value = $("#rolename").val()
        var url = window.location.search;
        const urlpar = new URLSearchParams(url);
        pageno = urlpar.get('page');

        if (pageno == null) {
            pageno = "1";
        }
        $.ajax({
            url: "/settings/roles/checkrole",
            type: "POST",
            async: false,
            data: { "name": value, "id": id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                if (data.role == true) {

                    $('#rolename-error').text("* " + languagedata.Permission.roleexist).show();
                    $('#rolen').addClass('input-group-error');
                }
                if (data.role == false) {
                    $('#roleform')[0].submit();
                    console.log("submit")
                    var url = $("#url").val()

                    var rolename = $("#rolename").val();

                    var roledesc = $('#roledesc').val();

                    var roleid = $('#rolid').val();
                    var roleisactive = $('#roleisactive').val()

                    var permissionid = []

                    $(".roleperm").each(function () {
                        if ($(this).prop("checked")) {
                            permissionid.push($(this).attr('data-id'))
                        }
                    })
                    $.ajax({
                        url: url,
                        type: "POST",
                        async: false,
                        data: { "rolename": rolename, "roledesc": roledesc, "permissionid": permissionid, csrf: $("input[name='csrf']").val(), "roleid": roleid, "roleisactive": roleisactive },
                        datatype: "json",
                        caches: false,
                        success: function (data) {
                            if (data.role == "added") {
                                setCookie('get-toast', 'Role Created Successfully', 1)
                                setCookie('Alert-msg', 'success', 1)
                                window.location.href = "/settings/roles?page=" + pageno
                            }
                            if (data.role == "updated") {
                                setCookie('get-toast', 'Role Updated Successfully', 1)
                                setCookie('Alert-msg', 'success', 1)
                                window.location.href = "/settings/roles?page=" + pageno
                            }
                        }
                    })
                }
            }
        })
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })

    }
})



/**Edit Get Data from backend */

$(document).on('click', '.roledit', function () {
    var id = $(this).attr('data-id');
    $('#rolid').val(id);

    Editrole(id, languagedata)

    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');

    $("#pageno").val(pageno)

})


// create role (model open) function

$(document).on('click', '.add-new', function () {

    $('.saverolperm').text('Save');

    $('.roleperm').attr('checked', false);

    $('#rolename').val("")

    $('#roledesc').text("")

    $('.modal-header').children('h5').text(languagedata.Rolecontent.addnewrole + ' & ' + languagedata.Rolecontent.setpermisson)

    $('#url').val('/settings/roles/createrole');

    $('.saverolperm').removeClass('roldisabled');

    $('.saverolperm').attr('disabled', false);


    //category group restrict


    if ($('#Check2').prop('checked')) {

        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').addClass("pointer-events-none opacity-25");
    }

    //category restrict


    if ($('#Check6').prop('checked')) {

        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').addClass("pointer-events-none opacity-25");
    }

    // member group restrict


    if ($('#Check11').prop('checked')) {

        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').addClass("pointer-events-none opacity-25");
    }

    // member restrict

    if ($('#Check15').prop('checked')) {

        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').addClass("pointer-events-none opacity-25");
    }


    // category (init) restrict

    if ($('#Check2:checked, #Check3:checked, #Check4:checked, #Check5:checked').length === 4) {
        $('#headingCategories').removeClass("pointer-events-none opacity-25");
        $('label[for="Check6"]').removeClass("pointer-events-none opacity-25");
    } else {
        $('#headingCategories').addClass("pointer-events-none opacity-25");
        $('label[for="Check6"]').addClass("pointer-events-none opacity-25");
    }



})


//public/js/settings/users
/** Delete Role & Permission */

$(document).on('click', '#deleterole-btn', function () {

    var id = $(this).attr('data-id')

    $.ajax({
        url: '/settings/roles/chkroleshaveuser',
        type: 'POST',
        async: false,
        data: { "rolesid": id, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        success: function (data) {
            if (data.value) {
                console.log("trrrrr", data.value);

                $('#delid').addClass("hidden");
                $('#dltCancelBtn').text(languagedata.ok);
                $("#content").text(languagedata.Roless.rolesrestrictmsg)
            } else {
                $('#delid').removeClass("hidden")
                $('#content').text(languagedata.Roless.deleterolecontent);
                $('#dltCancelBtn').text(languagedata.cancel);

            }
        }
    })

    var name = $(this).parents('tr').children('td:eq(1)').text();

    $('.deltitle').text(languagedata.Roless.deleterole)

    // $('.deldesc').text(languagedata.Roless.deleterolecontent)

    $('.delname').text(name)

    var url = window.location.search;

    const urlpar = new URLSearchParams(url);

    pageno = urlpar.get('page');

    $('#delid').attr('href', '/settings/roles/deleterole?id=' + id + "&page=" + pageno)

})


// search functions strat //


/*search redirect home page */

$(document).on('keyup', '#searchroles', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/settings/roles/";
        }
    }

    $('.searchClosebtn').toggleClass('hidden', searchInput === "");

})


$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")

})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/settings/roles/"
})

$(document).ready(function () {

    $('.search').on('input', function () {

        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
            $(".srchBtn-togg").addClass("pointer-events-none")

        } else {
            $(".Closebtn").addClass("hidden")
            $(".srchBtn-togg").removeClass("pointer-events-none")

        }
    });
})

$(document).on("click", ".hovericon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})

//   end


/**Role isactive status function */
function RoleStatus(id) {
    $('#cb' + id).on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb' + id).val();
    $.ajax({
        url: '/settings/roles/roleisactive',
        type: 'POST',
        async: false,
        data: {
            id: id,
            isactive: isactive,
            csrf: $("input[name='csrf']").val()
        },
        dataType: 'json',
        cache: false,
        success: function (result) {
            console.log(result,);
            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px]" > Role Status Updated Successfully </p ></div ></div ></li></ul> `;

            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds
        }
    });
}


function Validationcheck() {

    if ($('#rolename').hasClass('error')) {
        $('#rolen').addClass('input-group-error');
    } else {
        $('#rolen').removeClass('input-group-error');
    }

    if ($('#roledesc').hasClass('error')) {
        $('#roledes').addClass('input-group-error');
    } else {
        $('#roledes').removeClass('input-group-error');
    }
}


//**close button function */
$(document).on('click', '.close', function () {
    $('#roleform')[0].reset();
    $('label.error').remove()
    $('.input-group').removeClass('input-group-error')

})


// role edit function

$(document).on('click', '#configure', function () {
    var id = $(this).attr('data-id');
    $('#rolid').val(id);
    Editrole(id, languagedata)
    setTimeout(function () {
        $('.permissionsection')[0].scrollIntoView({ behavior: 'smooth', block: 'start' });
    }, 500);
});


// role edit function

$(document).on('click', '#manage', function () {
    var id = $(this).attr('data-id');
    $('#rolid').val(id);

    Editrole(id, languagedata)

    // setTimeout(function () {
    //     $('.permissionsection')[0].scrollIntoView({ behavior: 'smooth', block: 'start' });
    // }, 500);
});


// role edit function

function Editrole(id, languagedata) {
    $.ajax({
        url: "/settings/roles/getroledetail",
        type: "POST",
        async: false,
        data: { "id": id, csrf: $("input[name='csrf']").val() },
        datatype: "json",
        caches: false,
        success: function (result) {
            let jsonrole = $.parseJSON(result);
            console.log("ss", result, jsonrole);
            $('#rolename').val(jsonrole.role.Name)
            $('#roledesc').val(jsonrole.role.Description)

            console.log(jsonrole.role.Id, "admin");
            if (jsonrole.role.Id == 1) {

                $('.roleperm').attr('checked', true);
                $('.saverolperm').attr('disabled', true);
                $('.saverolperm').addClass('roldisabled');

            } else {
                $('.roleperm').attr('checked', false);
                $('.saverolperm').attr('disabled', false);
                $('.saverolperm').removeClass('roldisabled');
            }

            if (jsonrole.permissionid != null) {
                for (let x of jsonrole.permissionid) {
                    $("#Check" + x).attr('checked', true)
                }
            }

            $('.modal-header').children('h5').text(languagedata.updaterole + ' / ' + languagedata.Setting.managepermissions)
        }
    })

    $('#url').val('/settings/roles/updaterole');

    $('.saverolperm').text('Update');


    //category group restrict

    if ($('#Check2').prop('checked')) {

        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').addClass("pointer-events-none opacity-25");
    }

    //category restrict


    if ($('#Check6').prop('checked')) {

        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').addClass("pointer-events-none opacity-25");
    }

    // member group restrict


    if ($('#Check11').prop('checked')) {

        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').addClass("pointer-events-none opacity-25");
    }

    // member restrict

    if ($('#Check15').prop('checked')) {

        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').removeClass("pointer-events-none opacity-25");
    } else {

        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').addClass("pointer-events-none opacity-25");
    }

    // category (init) restrict

    if ($('#Check2:checked, #Check3:checked, #Check4:checked, #Check5:checked').length === 4) {
        $('#headingCategories').removeClass("pointer-events-none opacity-25");
        $('label[for="Check6"]').removeClass("pointer-events-none opacity-25");
    } else {
        $('#headingCategories').addClass("pointer-events-none opacity-25");
        $('label[for="Check6"]').addClass("pointer-events-none opacity-25");
    }



}


// checkbox select function

$(document).on('click', '.selectcheckbox', function () {

    $(".multidelete").addClass("responsive-width")


    var roleid = $(this).attr('data-id')

    var status = $(this).parents('td').siblings('td').find('.tgl-light').val();



    if ($(this).prop('checked')) {

        selectedcheckboxarr.push({ "roleid": roleid, "status": status })

    } else {
        const index = selectedcheckboxarr.findIndex(item => item.roleid === roleid);

        if (index !== -1) {

            selectedcheckboxarr.splice(index, 1);
        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {

        $('.selected-numbers').removeClass("hidden")

        if (selectedcheckboxarr.length == 1) {
            $('#deselectid').text(languagedata.deselect)
        } else if (selectedcheckboxarr.length > 1) {
            $('#deselectid').text(languagedata.deselectall)
        }

        var allSame = selectedcheckboxarr.every(function (item) {
            return item.status === selectedcheckboxarr[0].status;
        });

        var setstatus
        var img;

        if (selectedcheckboxarr[0].status === '1') {
            console.log(img, setstatus, "see");

            setstatus = languagedata.Userss.deactive;

            img = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = languagedata.Userss.active;

            img = "/public/img/Active.svg";

        }


        var htmlContent = '';

        if (allSame) {
            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';

        } else {

            htmlContent = '';
            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')


        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

        if (!allSame) {

            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

            $('.unbulishslt').html("")
        } else {
            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }


    } else {

        $('.selected-numbers').addClass("hidden")
    }

    var allChecked = true;

    $('.selectcheckbox').each(function () {

        if (!$(this).prop('checked')) {

            allChecked = false;

            return false;
        }
    });

    $('#Check').prop('checked', allChecked);

    console.log(selectedcheckboxarr, "checkkkk")
})


// select all checkbox

$(document).on('click', '#Check', function () {

    $(".multidelete").addClass("responsive-width")

    selectedcheckboxarr = []

    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            roldid = $(this).attr('data-id')

            var status = $(this).parents('td').siblings('td').find('.tgl-light').val();


            selectedcheckboxarr.push({ "roleid": roldid, "status": status })
        })

        if (selectedcheckboxarr.length != 0) {
            const deselectText = selectedcheckboxarr.length == 1 ? languagedata.deselect : languagedata.deselectall;
            $('#deselectid').text(deselectText);

            $('.selected-numbers').removeClass("hidden")
        } else {
            $('.selected-numbers').addClass("hidden")
        }
        var allSame = selectedcheckboxarr.every(function (item) {

            return item.status === selectedcheckboxarr[0].status;
        });


        var setstatus

        var img
        if (selectedcheckboxarr[0].status === '1') {

            setstatus = languagedata.Userss.deactive;

            img = "/public/img/In-Active.svg";

        } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = languagedata.Userss.active;

            img = "/public/img/Active.svg";
        }

        var htmlContent = '';

        if (allSame) {

            htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >' + '<span class="max-sm:hidden @[550px]:inline-block hidden">' + setstatus + '</span>';


            $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        } else {

            htmlContent = '';
            $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

        }

        $('#unbulishslt').html(htmlContent);

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

    } else {
        $('.selected-numbers').addClass("hidden")

        selectedcheckboxarr = []

        $('.selectcheckbox').prop('checked', isChecked);
    }
})


// delete model content update

$(document).on('click', '#seleccheckboxdelete', function () {

    $.ajax({
        url: '/settings/roles/chkroleshaveuser',
        type: 'POST',
        async: false,
        data: {
            "roleids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
        },
        dataType: 'json',
        success: function (data) {
            if (data.value) {
                $('#delid').addClass("hidden");
                $('#dltCancelBtn').text(languagedata.ok);
                if (selectedcheckboxarr.length > 1) {

                    $('.deltitle').text(languagedata.Roless.deleteroles)

                    $('#content').text(languagedata.Roless.rolesrestrictmsgs)

                } else {

                    $('.deltitle').text(languagedata.Roless.deleterole)

                    $('#content').text(languagedata.Roless.rolesrestrictmsg)
                }
            } else {
                $('#delid').removeClass("hidden")
                $('#dltCancelBtn').text(languagedata.cancel);

                if (selectedcheckboxarr.length > 1) {

                    $('.deltitle').text(languagedata.Roless.deleteroles)

                    $('#content').text(languagedata.Roless.deleterolescontent)

                } else {

                    $('.deltitle').text(languagedata.Roless.deleterole)

                    $('#content').text(languagedata.Roless.deleterolecontent)
                }

            }
        }
    })


    $("#delid").text($(this).text());

    $('#delid').addClass('checkboxdelete')
})


// status model content update

$(document).on('click', '#unbulishslt', function () {

    if (selectedcheckboxarr.length > 1) {

        $('.deltitle').text($(this).text() + " " + languagedata.Roless.Roles)

        $('#content').text(languagedata.Roless.statusupdaterole + " " + $(this).text() + " " + languagedata.Roless.selectedroles)
    } else {

        $('.deltitle').text($(this).text() + " " + languagedata.Roless.Role)

        $('#content').text(languagedata.Roless.statusupdaterole + " " + $(this).text() + " " + languagedata.Roless.selectedrole)
    }
    $("#delid").text($(this).text());

    $('#delid').addClass('selectedunpublish')

})

// cancel btn function

$("#deleteModal").on("hide.bs.modal", function () {
    $('#delid').removeClass('checkboxdelete');
    $('#delid').removeClass('selectedunpublish');
    $('#delid').attr('href', '');
    $('.delname').text('')
})



//MULTI SELECT DELETE FUNCTION//

$(document).on('click', '.checkboxdelete', function () {

    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/settings/roles/multiselectroledelete',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "roleids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {


            if (data.value == true) {

                setCookie("get-toast", "Role Deleted Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })

})


//Deselectall function//

$(document).on('click', '#deselectid', function () {

    $('.selectcheckbox').prop('checked', false)

    $('#Check').prop('checked', false)

    selectedcheckboxarr = []

    $('.selected-numbers').addClass("hidden")

})


//multi select active and deactive function//

$(document).on('click', '.selectedunpublish', function () {

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    pageno = urlpar.get('page')

    $('.selected-numbers').addClass("hidden")
    $.ajax({
        url: '/settings/roles/multiselectrolestatus',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "roleids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "Role Updated Successfully")

                window.location.href = data.url
            } else {

                setCookie("Alert-msg", "Internal Server Error")

            }

        }
    })


})


// Group name limit of 25 char

$(document).on('keyup', '.checklength', function () {
    var inputVal = $(this).val()

    var inputLength = inputVal.length

    if (inputLength == 25) {
        $(this).siblings('.lengthErr').removeClass('hidden')
    } else {
        $(this).siblings('.lengthErr').addClass('hidden')
    }

})




$('#Check2, #Check3, #Check4, #Check5').on('change', function () {
    // Check if exactly four checkboxes are checked
    if ($('#Check2:checked, #Check3:checked, #Check4:checked, #Check5:checked').length === 4) {
        $('#headingCategories').removeClass("pointer-events-none opacity-25");
        $('label[for="Check6"]').removeClass("pointer-events-none opacity-25");
    } else {
        $('#Check6,#Check7, #Check8, #Check9').prop('checked', false);
        $('#headingCategories').addClass("pointer-events-none opacity-25");
        $('label[for="Check6"], label[for="Check7"], label[for="Check8"], label[for="Check9"]').addClass("pointer-events-none opacity-25");
        $('#collapseCategories').removeClass('show');
        $('#headingCategories button').addClass('collapsed');


    }
});


//category group

$('#Check2').on('change', function () {
    if ($(this).prop('checked')) {
        // Remove the class if #Check15 is checked
        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').removeClass("pointer-events-none opacity-25");

    } else {
        // Add the class if #Check15 is unchecked
        $('#Check3, #Check4, #Check5').prop('checked', false);

        $('label[for="Check3"], label[for="Check4"], label[for="Check5"]').addClass("pointer-events-none opacity-25");

    }
});


//category

$('#Check6').on('change', function () {
    if ($(this).prop('checked')) {
        // Remove the class if #Check15 is checked
        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').removeClass("pointer-events-none opacity-25");

    } else {
        // Add the class if #Check15 is unchecked
        $('#Check7, #Check8, #Check9').prop('checked', false);
        $('label[for="Check7"], label[for="Check8"], label[for="Check9"]').addClass("pointer-events-none opacity-25");

    }
});


// member group
$('#Check11').on('change', function () {
    if ($(this).prop('checked')) {
        // Remove the class if #Check15 is checked
        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').removeClass("pointer-events-none opacity-25");
    } else {
        // Add the class if #Check15 is unchecked
        $('#Check12, #Check13, #Check14').prop('checked', false);
        $('label[for="Check12"], label[for="Check13"], label[for="Check14"]').addClass("pointer-events-none opacity-25");
    }
});

// member

$('#Check15').on('change', function () {
    if ($(this).prop('checked')) {
        // Remove the class if #Check15 is checked
        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').removeClass("pointer-events-none opacity-25");
    } else {
        // Add the class if #Check15 is unchecked
        $('#Check16, #Check17, #Check18').prop('checked', false);
        $('label[for="Check16"], label[for="Check17"], label[for="Check18"]').addClass("pointer-events-none opacity-25");
    }
});


