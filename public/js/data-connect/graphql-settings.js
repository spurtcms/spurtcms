var languagedata, editTknName, editTknDrn, editTknDesc

$(document).ready(async function () {
    var languagepath = $('.language-group>button').attr('data-path')
    await $.getJSON(languagepath, function (data) {
        languagedata = data
    })
})

$(document).on('click', '.tkn-durn.dropdown-item', function () {
    var selectedVal = $(this).text()
    $('#slctdurn').text(selectedVal)
    if ($("#createToken #genTokenBtn").attr('data-status') == "1" && editTknDrn != selectedVal) {
        $("#createToken #genTokenBtn").removeClass('bg-[#D1D1D1]')
        $("#createToken #genTokenBtn").addClass('hover:bg-[#148569]').addClass('bg-[#10A37F]')
        $("#createToken #genTokenBtn").prop('disabled', false).removeClass('cursor-not-allowed')
    } else {
        CheckIsEditDataUnchanged()
    }
})

var isCalled = false
$(document).on('click', '#genTokenBtn[data-status=0]', function (e) {
    e.preventDefault();
    if (isCalled) {
        return
    }

    var formData = new FormData()
    formData.append("csrf", $('#csrf-value').val())
    var fieldVal = []


    $('.apitoken-data').each(function (index, element) {

        if (element.tagName == "INPUT" || element.tagName == "TEXTAREA") {
            var fieldValue = $.trim($(element).val())
            var fieldParam = $(element).attr('name')
        } else {
            var fieldValue = $.trim($(element).text())
            var fieldParam = $(element).attr('name')
        }


        if (element.tagName == "INPUT" || element.tagName == "A") {
            fieldVal.push(fieldValue)
            if (fieldValue != "" && fieldValue != "Select Duration" && fieldValue != "Seleccionar duración" && fieldValue != "Sélectionnez la durée") {

                formData.append(fieldParam, fieldValue)
            } else {
                $(element).siblings('.apiTokenErr').removeClass('hidden')
            }
        } else {
            formData.append(fieldParam, fieldValue)
        }
    })

    let validData = fieldVal.every(d => d != "" && d != "Select Duration" && d != "Seleccionar duración" && d != "Sélectionnez la durée")
    if (validData) {
        isCalled = true
    }

    if (validData && isCalled) {

        $('#genTokenBtn').attr('data-bs-toggle', 'modal')
        $('#genTokenBtn').css('display', 'none')
        $('#genTokenBtn').get(0).click()
        $.ajax({
            url: "/graphql/create",
            type: "POST",
            data: formData,
            contentType: false, // Prevent jQuery from setting Content-Type
            processData: false, // Prevent jQuery from processing the data
            success: function (result) {
                $('.apitoken-data').val('')
                if (result.status == 1 && result.token != "") {
                    $('#copyToken #copy-btn').attr('data-use', "save")
                    $("#copyToken #tokenContentDiv #tokenContent").text(result.token)
                    $("#copyToken #tokenContentDiv #tokenContent").attr('data-value', result.token)
                    $('#copyToken').modal('show')
                }
                isCalled = false

            },
            error: function (jqXHR, textStatus, errorThrown) {
                var errorRespoonse = JSON.parse(jqXHR.responseText)
                console.log('Error:', errorRespoonse.status, errorRespoonse, textStatus, errorThrown);
            }
        });

    }
})


var originalContent = $('#createToken').clone();

$(document).on('click', '#tokenCancelBtn', function () {
    $('#createToken').html(originalContent.html())
})

$(document).on('click', '.tokenDelBtn', function () {
    $('#createToken').html(originalContent.html())
})

$(document).on('click', '#copy-btn', function () {
    $('#createToken').html(originalContent.html())
})

$(document).on('click', '#graphqlCopyCancel', function () {
    window.location.href = "/graphql/"
})

function TokenCopy(token) {
    navigator.clipboard.writeText(token)
}

$(document).on('click', '#copyToken #copy-btn', function (e) {
    var dataUse = $(this).attr('data-use')
    $(this).addClass('hidden')
    $('#copied-btn').removeClass('hidden')
    setTimeout(() => {
        if (dataUse == "save") {
            TokenCopy($('#tokenContent').attr('data-value').trim())
            window.location.href = '/graphql/'
            setCookie('get-toast', 'Graphql Settings Created Successfully')
        } else if (dataUse == "copy") {
            TokenCopy($('#tokenContent').attr('data-value').trim())
            $('#copyToken').modal('hide')
        }
    }, 500)
})

$(document).on('click', ".apitoken-data", function () {
    if ($(this).val() != "select" && $(this).siblings('label.apiTokenErr').hasClass('hidden') != true) {
        $(this).siblings('label.apiTokenErr').addClass('hidden')
    }
})

$(document).on('click', '#key-desc', function () {
    var id = $(this).attr('data-id')
    $.ajax({
        url: "/graphql/edit/" + id,
        type: "GET",
        success: function (result) {
            editTknName = result.data.TokenName
            editTknDesc = result.data.Description
            editTknDrn = result.data.Duration
            $("#createToken #modalTitleId").text(languagedata.Graphql.updatetokenheading)
            $("#createToken #genTokenBtn").attr('data-status', '1')
            $("#createToken #genTokenBtn").addClass('bg-[#D1D1D1]')
            $("#createToken #genTokenBtn").prop('disabled', true).addClass('cursor-not-allowed')
            $("#createToken #genTokenBtn").removeClass('hover:bg-[#148569]')
            $("#createToken #genTokenBtn").text(languagedata.update)
            $("#createToken #tokenName").val(result.data.TokenName)
            $("#createToken #tokenName").attr('data-id', id)
            $("#createToken #tokenDesc").val(result.data.Description)
            $('#tokeninsideContent').text(result.data.Token)
            $('#tokeninsideContent').attr('data-value', result.data.Token)
            $("#slctdurn").text(result.data.Duration);
            $('#apikey').removeClass('hidden')
            if (result.data.IsDefault == 1) {
                $('.dropdown.tokenInputGrp').addClass('hidden')
                $("#createToken #genTokenBtn").attr('default-update', '1')
            }
            $('#key-blck').removeClass('hidden')
            $('#createToken').modal('show')
        }
    })
})

$(document).on('input', '#createToken #tokenName', function () {
    if ($("#createToken #genTokenBtn").attr('data-status') == "1") {
        var currentValue = $(this).val()
        if (currentValue != editTknName) {
            $("#createToken #genTokenBtn").removeClass('bg-[#D1D1D1]')
            $("#createToken #genTokenBtn").addClass('hover:bg-[#148569]').addClass('bg-[#10A37F]')
            $("#createToken #genTokenBtn").prop('disabled', false).removeClass('cursor-not-allowed')
        } else {
            CheckIsEditDataUnchanged()

        }
    }
});

$(document).on('input', '#createToken #tokenDesc', function () {
    if ($("#createToken #genTokenBtn").attr('data-status') == "1") {
        var currentValue = $(this).val()
        if (currentValue != editTknDesc) {
            $("#createToken #genTokenBtn").removeClass('bg-[#D1D1D1]')
            $("#createToken #genTokenBtn").addClass('hover:bg-[#148569]').addClass('bg-[#10A37F]')
            $("#createToken #genTokenBtn").prop('disabled', false).removeClass('cursor-not-allowed')
        } else {
            CheckIsEditDataUnchanged()
        }
    }
});

function CheckIsEditDataUnchanged() {
    var desc = $("#createToken #tokenDesc").val()
    var durn = $("#slctdurn").text();
    var name = $("#createToken #tokenName").val()
    if ((desc == editTknDesc && durn == editTknDrn && name == editTknName) || ($("#createToken #genTokenBtn").attr('default-update') == "1" && durn == editTknDrn && name == editTknName)) {
        $("#createToken #genTokenBtn").addClass('bg-[#D1D1D1]')
        $("#createToken #genTokenBtn").removeClass('hover:bg-[#148569]')
        $("#createToken #genTokenBtn").prop('disabled', true).addClass('cursor-not-allowed')
    }
}

$(document).on('click', '#genTokenBtn[data-status=1]', function (e) {

    var formData = new FormData()

    formData.append("csrf", $('#csrf-value').val())


    var id = $('#createToken #tokenName').attr('data-id')
    formData.append("id", id)

    var fieldVal = []

    $('.apitoken-data').each(function (index, element) {
        if (element.tagName == "INPUT" || element.tagName == "TEXTAREA") {
            var fieldValue = $.trim($(element).val())
            var fieldParam = $(element).attr('name')
        } else {
            var fieldValue = $.trim($(element).text())
            var fieldParam = $(element).attr('name')
        }

        if (element.tagName == "INPUT" || element.tagName == "A") {
            fieldVal.push(fieldValue)
            if (fieldValue != "" && fieldValue != "Select Duration" && fieldValue != "Seleccionar duración" && fieldValue != "Sélectionnez la durée") {

                formData.append(fieldParam, fieldValue)
            } else {
                $(element).siblings('.apiTokenErr').removeClass('hidden')
            }
        } else {
            formData.append(fieldParam, fieldValue)
        }
    })

    let validData = fieldVal.every(d => d != "" && d != "Select Duration" && d != "Seleccionar duración" && d != "Sélectionnez la durée")
    if (validData) {
        $.ajax({
            url: "/graphql/update",
            type: "POST",
            data: formData,
            contentType: false,
            processData: false,
            success: function (result) {
                window.location.href = "/graphql/"
            }
        })
    }

})



//--------------------Delete token start-----------------

$(document).on('click', '.tokenDelBtn', function () {
    var id = $(this).attr('data-id');
    $('.deltitle').text(languagedata.Graphql.deletetoken + " ?")
    $("#content").text(languagedata.Graphql.deletesubheading)
    $(".deleteBtn").attr('href', '/graphql/delete/' + id)
    $('.deleteBtn').attr('id', "delete" + id)

})

//--------------------Delete token end------------------


//search bar code below

$(document).on('keypress', '#graphqlSearchBar', function (e) {

    if (e.which == 13) {
        var searchText = $('#graphqlSearchBar').val()
        console.log(searchText, "searchText");


        if (searchText != "") {
            $('#graphqlSearchlink').attr('href', "/graphql/?keyword=" + searchText)
            $('#graphqlSearchBar').attr('value', searchText)
            $('#graphqlSearchlink').get(0).click()
        } else {
            window.location.href = "/graphql/"
        }

    }
})

// on pressing backspace if searcbar is empty show list page
$(document).on('keyup', '#graphqlSearchBar', function (e) {
    var searchVal = $('#graphqlSearchBar').attr('value')
    if (e.which == 8 && searchVal != "") {
        var searchKey = $('#graphqlSearchBar').val()
        if (searchKey == "") {
            window.location.href = "/graphql/"
        }
    }
})

var selectedcheckboxarr = []

//single check box delete function
console.log(selectedcheckboxarr);


$(document).on('click', '.selectcheckbox', function () {

    var tokenid = $(this).attr('data-id')


    if ($(this).prop('checked')) {

        selectedcheckboxarr.push(tokenid)

    } else {


        var uncheckedId = $(this).attr('data-id')

        const indexVal = selectedcheckboxarr.findIndex((element) => element == uncheckedId)


        if (indexVal != -1) {

            selectedcheckboxarr.splice(indexVal, 1);

        }

        $('#Check').prop('checked', false)

    }


    if (selectedcheckboxarr.length != 0) {

        $('.selected-numbers').removeClass("hidden")

        $('#unbulishslt').css('display', 'none')
        $('#seleccheckboxdelete').removeClass('border-r')

        $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)


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

    console.log(selectedcheckboxarr);

})

$(document).on('click', '#seleccheckboxdelete', function () {

    $('.deltitle').text(languagedata.Graphql.deletetoken + " ?")

    $('#content').text(languagedata.Graphql.deletesubheading)

    $('#delid').addClass('checkboxdelete')

    $('#delid').text("Delete")
})


//  multiselect delete function
$(document).on('click', '.checkboxdelete', function () {

    var url = window.location.href;

    var pageurl = window.location.search

    const urlpar = new URLSearchParams(pageurl)

    var pageno = urlpar.get('page')

    $('.selected-numbers').hide()
    $.ajax({
        url: '/graphql/multideletetokens',
        type: 'post',
        dataType: 'json',
        async: false,
        data: {
            "tokenids": JSON.stringify(selectedcheckboxarr),
            csrf: $("input[name='csrf']").val(),
            "page": pageno
        },
        success: function (data) {
            console.log(data, "data");
            if (data.value == true) {

                setCookie("get-toast", "Graphql Settings Deleted Successfully")
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

    $('.selected-numbers').addClass('hidden');

})

// select all by making master checkbox checked

$(document).on('click', '#Check', function () {

    if ($('#Check').prop('checked')) {

        if ($('.selectcheckbox[data-default=0]').length > 0) {

            $('.selectcheckbox[data-default=0]').each(function () {
                $(this).prop('checked', true)
                var id = $(this).attr('data-id')

                if (!selectedcheckboxarr.includes(id)) {
                    selectedcheckboxarr.push(id)
                }
            });

            $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.itemselected)

            $('.selected-numbers').removeClass('hidden')
            $('#unbulishslt').css('display', 'none')
            $('#seleccheckboxdelete').removeClass('border-r')
        }

    } else {

        $('.selectcheckbox').each(function () {

            $(this).prop('checked', false)
            selectedcheckboxarr = []
        });

        $('.selected-numbers').addClass('hidden')
    }

})

$(document).ready(function () {
    var $textarea = $('#tokenDesc');
    var $errorMessage = $('#error-message');
    var maxLength = 250;

    $textarea.on('input', function () {
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorMessage.text(languagedata.Permission.descriptionchat).removeClass("hidden");
        } else {
            // Clear error message if under the limit
            $errorMessage.text('').addClass('hidden');
        }
    });
});

$(document).ready(function () {
    $('#tokenDesc').on('input', function () {

        let lines = $(this).val().split('\n').length;

        if (lines > 5) {

            let value = $(this).val();

            let linesArray = value.split('\n').slice(0, 5);

            $(this).val(linesArray.join('\n'));
        }
    });
});

$(document).on('click', '#copy-insidebtn', function () {
    TokenCopy($('#tokeninsideContent').attr('data-value').trim())
    $(this).text(languagedata.Graphql.copied)
    setTimeout(() => {
        $(this).html('<img src="/public/img/copy-green.svg">' + languagedata.copy)
    }, 1000)
})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')

    window.location.href = "/graphql/"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
        } else {
            $(".Closebtn").addClass("hidden")
        }
    });
})

$(document).on("click", ".graphqlHoverIcon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})