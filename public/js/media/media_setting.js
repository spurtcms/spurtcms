var languagedata

$(document).ready(async function () {

    $('.Content').addClass('checked');

    var languagePath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagePath, function (data) {

        languagedata = data
    })

})


$(document).on('click', '#radio-1', function () {

    if ($(this).is(':checked')) {

        $('.localstoragesetting').removeClass("hidden");

        $('.amazonsetting').addClass("hidden");

        $('.azuresetting').addClass("hidden");
    }
})

$(document).on('click', '#radio-2', function () {

    if ($(this).is(':checked')) {

        $('.amazonsetting').removeClass("hidden");

        $('.localstoragesetting').addClass("hidden");

        $('.azuresetting').addClass("hidden");
    }
})


$(document).on('click', '#radio-3', function () {

    if ($(this).is(':checked')) {

        $('.azuresetting').show();

        $('.localstoragesetting').hide();

        $('.amazonsetting').hide();
    }
})


$(document).on("click", "#savetype", function () {

    var selectedtype = $("input[name='radio']:checked").attr('data-select')

    if (selectedtype == 'local') {

        LocalFormValidation()

    } else if (selectedtype == 'aws') {

        AwsS3FormValidation()

    } else if (selectedtype == 'azure') {

        AzureBlobFormValidation()

    } else if (selectedtype == 'drive') {


    }

    AddErrorLabel()
})

function LocalFormValidation() {

    $("#localsettingform").validate({

        ignore: [],

        rules: {
            local: {
                required: true,
                space: true,
            },
        },
        messages: {
            cname: {
                required: "* Please enter the storage path",
                space: "* " + languagedata.spacergx,
            },
        }
    });

    var formcheck = $("#localsettingform").valid();

    if (formcheck == true) {

        $('#localsettingform')[0].submit();

    }

    return formcheck

}

function AwsS3FormValidation() {

    $("#s3settingform").validate({

        ignore: [],

        rules: {
            awsid: {
                required: true,
                space: true,
            },
            awskey: {
                required: true,
                space: true,
            },
            awsregion: {
                required: true,
                space: true,
            },
            awsbucketname: {
                required: true,
                space: true,
            },
        },
        messages: {
            awsid: {
                required: "* Please enter the awsid",
                space: "* " + languagedata.spacergx,
            },
            awskey: {
                required: "* Please enter the awskey",
                space: "* " + languagedata.spacergx,
            },
            awsregion: {
                required: "* Please enter the awsregion",
                space: "* " + languagedata.spacergx,
            },
            awsbucketname: {
                required: "* Please enter the awsbucketname",
                space: "* " + languagedata.spacergx,
            },
        }
    });

    var formcheck = $("#s3settingform").valid();

    if (formcheck == true) {

        $('#s3settingform')[0].submit();

    }


    return formcheck
}

function AzureBlobFormValidation() {

    $("#azuresettingform").validate({

        ignore: [],

        rules: {
            azureacc: {
                required: true,
                space: true,
            },
            azurekey: {
                required: true,
                space: true,
            },
            azurecontainer: {
                required: true,
                space: true,
            }
        },
        messages: {
            azureacc: {
                required: "* Please enter the azure account name",
                space: "* " + languagedata.spacergx,
            },
            azurekey: {
                required: "* Please enter the azure key",
                space: "* " + languagedata.spacergx,
            },
            azurecontainer: {
                required: "* Please enter the azure container",
                space: "* " + languagedata.spacergx,
            }
        }
    });

    var formcheck = $("#azuresettingform").valid();

    if (formcheck == true) {

        $('#azuresettingform')[0].submit();

    }

    return formcheck

}

function AddErrorLabel() {

    $(".error").each(function () {

        $(this).parents('.input-group').addClass("input-group-error");
    })

}

$(document).on('keyup', '.awsvalid', function () {

    $(this).parents('.input-group').removeClass("input-group-error");
    
    if ($(this).hasClass('.error')) {

        // if ($(this).val() == "") {

        $(this).parents('.input-group').addClass("input-group-error");
        // }
    }

})