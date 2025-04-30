var selectedcheckboxarr = []


$("#Save").on("click", function (event) {    

    $("#aisettingscreate").validate({

        ignore: [],

        rules: {
            apikey: {
                required: true
            },
            apimodel: {
                required: true

            }

        },
        messages: {
            apikey: {
                required: "Please Enter Api Key"

            },
            apimodel: {
                required: "Please Select AI Model"

            }

        }
    });

    var formcheck = $("#aisettingscreate").valid();

    if (formcheck) {

        $("#aisettingscreate").submit();

    }

});

$("#Cancel").on("click", function () {

    $("#modalTitleId").text("Create New AI Module")

    $("#aisettingscreate").attr("action", "/settings/aisettings/create");

    $("#Save").text("Save")

    $("#moduleid").val('')

    $("#aimodule").val("openai")

    $("#apikey").val('');

    $("textarea[name=desc]").text('');

    $("#apimodel").val('')

    $("#triggerspan").text('select');

    $('.aimodule').removeClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#10A37F]");

    $("#openai").addClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#10A37F]")

    $('.dropdown-item1').text('gpt-4o-mini');

    $('.dropdown-item2').text('gpt-3.5');

})


$(document).on('click', '.aimodule', function () {

    id = $(this).data("id")

    value = $(this).data("value")

    $("#aimodule").val(value)

    $('.aimodule').removeClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#ECECEC] border-[#10A37F]")

    $(this).addClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#10A37F]")


    if (id == 1) {

        $('.dropdown-item1').text('gpt-4o-mini')

        $('.dropdown-item2').text('gpt-3.5')

    } else if (id == 2) {

        $('#openai').addClass("hover:border-[#10A37F]")

        $('.dropdown-item1').text('Gemini 2.0 Flash-Lite')

        $('.dropdown-item2').text('Gemini 1.5 Flash')

    } else if (id == 3) {

        $('#openai').addClass("hover:border-[#10A37F]")

        $('.dropdown-item1').text('DeepSeek R1')

        $('.dropdown-item2').text('DeepSeek V3')

    }

})

$(document).on("click", ".dropdown-item", function () {

    var value = $(this).text()

    $("#apimodel").val(value);

    $("#triggerId span").text(value)

    $("#apimodel-error").hide()

})

$(document).ready(function () {

    var $textarea = $('#desc');

    var $errorMessage = $('#error-messagedesc');

    var maxLength = 250;
    
    $textarea.on('input', function () {

        if ($(this).val().length >= maxLength) {

            $errorMessage.text("Maximum 250 character allowed");

        } else {

            $errorMessage.text('');
            
        }
    });

});


$(document).on("click", "#editmodule", function () {

    $("#modalTitleId").text("Update AI Module")

    $("#Save").text("Update")

    $("#aisettingscreate").attr("action", "/settings/aisettings/update");

    var id = $(this).data("id")

    $("#moduleid").val(id)

    edit = $(this).closest("tr");

    var module = edit.find("td:eq(1)").text();

    $("#aimodule").val(module.trim())

    var apiKey = edit.find("input[type='hidden']").val();

    $("#apikey").val(apiKey.trim())

    var desc = edit.find("td:eq(2)").text();

    $("textarea[name=desc]").text(desc.trim());

    var model = edit.find("td:eq(5)").text();

    $("#triggerspan").text(model.trim())

    $("#apimodel").val(model.trim())

    $('.aimodule').removeClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#ECECEC] border-[#10A37F]")

    $("#" + module.trim()).addClass("bg-[#FAFFFE] shadow-[0px_0px_6px_0px_#00000014] border-[#10A37F]")

    if (module.trim() == "openai") {

        $('.dropdown-item1').text('gpt-4o-mini');

        $('.dropdown-item2').text('gpt-3.5');

    } else if (module.trim() == "geminiai") {

        $('#openai').addClass("hover:border-[#10A37F]")

        $('.dropdown-item1').text('Gemini 2.0 Flash-Lite');

        $('.dropdown-item2').text('Gemini 1.5 Flash');

    } else if (module.trim() == "deepseekai") {

        $('#openai').addClass("hover:border-[#10A37F]")

        $('.dropdown-item1').text('DeepSeek R1');

        $('.dropdown-item2').text('DeepSeek V3');

    }


})



function SubscriptionStatus(id) {


    $('#check' + id).on('change', function () {

        this.value = this.checked ? 1 : 0;

    }).change();

    var isactive = $('#check' + id).val()

    console.log("isactive:", isactive, id);

    $.ajax({
        url: '/settings/aisettings/status',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {

            if (result.value) {

                setCookie("get-toast", "Status Changed Successfully")

                window.location.href = result.url

            } else {

                setCookie("Alert-msg", "Internal Server Error")

                window.location.href = result.url

            }
        }
    });


}

//single delete

$(document).on('click', '#delete-btn', function () {

    var Id = $(this).attr("data-id")

    $("#content").text("Are you sure you want to delete this Module?")

    var url = window.location.search

    const urlpar = new URLSearchParams(url)

    pageno = urlpar.get('page')

    if (pageno == null) {

        $('#delid').attr('href', "/settings/aisettings/delete/" + Id);

    } else {

        $('#delid').attr('href', "/settings/aisettings/delete/" + Id + "?page=" + pageno);

    }
    $(".deltitle").text("Delete AI Settings Module?")

    $('.delname').text($(this).parents('tr').find('td:eq(1)').text())

})

//multi delete

$(document).on('click', '.selectcheckbox', function () {

    $(".multidelete").addClass("responsive-width")

    $('#unbulishslt').hide()

    $('#seleccheckboxdelete').removeClass('border-r');

    moduleid = $(this).attr('data-id')


    if ($(this).prop('checked')) {

        selectedcheckboxarr.push(moduleid)

    } else {

        const index = selectedcheckboxarr.indexOf(moduleid);

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

    $(".multidelete").addClass("responsive-width")

    $('#unbulishslt').hide()

    $('#seleccheckboxdelete').removeClass('border-r');

    selectedcheckboxarr = []

    var isChecked = $(this).prop('checked');

    if (isChecked) {

        $('.selectcheckbox').prop('checked', isChecked);

        $('.selectcheckbox').each(function () {

            moduleid = $(this).attr('data-id')

            selectedcheckboxarr.push(moduleid)
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

        $('.deltitle').text("Delete AI Module?")

        $('#content').text('Are you sure want to delete selected AI Modules?')

    } else {

        $('.deltitle').text("Delete AI Module?")

        $('#content').text('Are you sure want to delete selected AI Modules?')
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
        url: '/settings/aisettings/multiselectdelete',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "moduleids": selectedcheckboxarr,
            csrf: $("input[name='csrf']").val(),
            "page": pageno


        },
        success: function (data) {

            if (data.value == true) {

                setCookie("get-toast", "AI Module Deleted Successfully")

                window.location.href = data.url

            } else {

                setCookie("Alert-msg", "Internal Server Error")

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

//Search functions

$(document).on("click", ".Closebtn", function () {

    $(".search").val('')

    $(".Closebtn").addClass("hidden")

    $(".SearchClosebtn").removeClass("hidden")

    $(".srchBtn-togg").removeClass("pointer-events-none")

})

$(document).on("click", ".searchClosebtn", function () {

    $(".search").val('')

    window.location.href = "/settings/aisettings/"

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



