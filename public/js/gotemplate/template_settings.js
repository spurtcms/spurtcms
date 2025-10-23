var languagedata
var selectedcheckboxarr = []
/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    $('.Content').addClass('checked');
})

$(document).ready(function () {

    //Site Name

    $("#siteNameSave").click(function () {

        var value = $("#siteName").val();
        console.log("value::", value);


        if (value.length != "") {
            console.log("value1::", value);
            $('#siteNameForm')[0].submit();

        } else {
            console.log("value2::", value);
            $("#siteName-error").show().text("* " + languagedata.WebsiteSettings.pleaseenteryoursitename)

        }

    })

    $("#siteName").on('input', function () {

        var value = $("#siteName").val();

        if (value.length >= 20) {

            $("#siteName-error").show().text(languagedata.WebsiteSettings.maximumlength20)


        } else if (value.length == "") {

            $("#siteName-error").show().text("* " + languagedata.WebsiteSettings.pleaseenteryoursitename)


        } else {

            $("#siteName-error").hide().text("")

        }
    })

    //Site Logo

    $('.fileInput').on('change', function () {
        var files = $(this).prop('files');
        if (files.length > 0) {
            var file = files[0];

            // Show file name
            $('#filenameDisplay').text(file.name);
            $('#deleteFile').removeClass('hidden');

            var reader = new FileReader();
            reader.onload = function (e) {
                var base64Data = e.target.result;

                $("#sitelogoimage").val(base64Data)
                $("#sitelogo").val(base64Data)
            };
            reader.readAsDataURL(file);
        }

    });

    // Optional: Delete/reset file input
    $('#deleteFile').on('click', function () {
        $('.fileInput').val('');
        $('#filenameDisplay').text(languagedata.Seo.nofilechosen);
        $(this).addClass('hidden');
        $("#sitelogoimage").val("")
        $("#sitelogo").val("")
    });


    $("#siteLogoSave").click(function () {

        $("#siteLogoForm").validate({

            ignore: [],

            rules: {
                sitelogoimage: {
                    required: true
                }
            },
            messages: {
                sitelogoimage: {
                    required: "* " + languagedata.Seo.pleasechooseimage
                }
            }
        })

        var sitemap = $("#siteLogoForm").valid();
        if (sitemap == true) {
            $('#siteLogoForm')[0].submit();

        }

    })

    //Site FavIcon


    $('.fileInputFavIcon').on('change', function () {
        var files = $(this).prop('files');
        if (files.length > 0) {
            var file = files[0];

            // Show file name
            $('#filenameDisplayFavIcon').text(file.name);
            $('#deleteFileFavIcon').removeClass('hidden');

            var reader = new FileReader();
            reader.onload = function (e) {
                var base64Data = e.target.result;

                $("#sitefaviconimage").val(base64Data)
                $("#sitefavicon").val(base64Data)
            };
            reader.readAsDataURL(file);
        }

    });

    // Optional: Delete/reset file input
    $('#deleteFileFavIcon').on('click', function () {
        $('.fileInputFavIcon').val('');
        $('#filenameDisplayFavIcon').text(languagedata.Seo.nofilechosen);
        $(this).addClass('hidden');
        $("#sitefaviconimage").val("")
        $("#sitefavicon").val("")
    });


    $("#siteFavIconSave").click(function () {

        $("#siteFavIconForm").validate({

            ignore: [],

            rules: {
                sitefaviconimage: {
                    required: true
                }
            },
            messages: {
                sitefaviconimage: {
                    required: "* " + languagedata.Seo.pleasechooseimage
                }
            }
        })

        var sitemap = $("#siteFavIconForm").valid();
        if (sitemap == true) {
            $('#siteFavIconForm')[0].submit();

        }

    })


    //Website URL

    $("#websiteUrlSave").click(function (e) {
        e.preventDefault();

        let websiteInput = $("#websiteInput").val();
        const isValid = /^[a-z0-9]*$/.test(websiteInput);

        if (websiteInput === "") {
            $("#websiteInput-error").show().text("* " + languagedata.WebsiteSettings.pleaseenteryourwebsiteurl);
            return;
        }

        if (!isValid) {
            $("#websiteInput-error").show().text(languagedata.WebsiteSettings.onlylowercaselettersandnumbersareallowed);
            return;
        }


        $.ajax({
             url: "/admin/website/checksitename ",
            type: "POST",
            data: {
                "sitename": websiteInput,
                csrf: $("input[name='csrf']").val(),
            },
            dataType: "json",
            cache: false,
            success: function (result) {
                if (result) {
                    $("#websiteInput-error").show().text("Name Already Exists");
                } else {
                    $("#websiteInput-error").hide();
                    $('#websiteUrlForm')[0].submit();
                }
            },

        });
    });


    $("#websiteInput").on('input', function () {

        let value = $(this).val();

        const isValid = /^[a-z0-9]*$/.test(value);

        if (value === '') {

            $("#websiteInput-error").show().text("* " + languagedata.WebsiteSettings.pleaseenteryourwebsiteurl)

        } else if (value.length >= 20) {

            $("#websiteInput-error").show().text(languagedata.WebsiteSettings.maximumlength20)

        } else if (!isValid) {

            $("#websiteInput-error").show().text(languagedata.WebsiteSettings.onlylowercaselettersandnumbersareallowed);

        } else {

            $("#websiteInput-error").hide().text("");

        }
    })
})

$('.hd-crd-btn').click(function () {

    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show");
        document.cookie = `webbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `webbanner=true; path=/;`;
    }
});