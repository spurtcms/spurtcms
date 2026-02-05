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


    var $pagetitle = $('#pagetitle');
    var $errorpagetitle = $('#error-pagetitle');

    $pagetitle.on('input', function () {
        var maxLength = 95;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorpagetitle.text("* "+languagedata.Seo.titlemustnotexceed55characters);
        } else {
            // Clear error message if under the limit
            $errorpagetitle.text('');
        }
    });

    var $pagedescription = $('#pagedescription');
    var $errorpagedescription = $('#error-pagedescription');

    $pagedescription.on('input', function () {
        var maxLength = 250;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorpagedescription.text("* "+languagedata.Seo.descriptionmustnotexceed155characters);
        } else {
            // Clear error message if under the limit
            $errorpagedescription.text('');
        }
    });

    var $pagekeyword = $('#pagekeyword');
    var $errorpagekeyword = $('#error-pagekeyword');

    $pagekeyword.on('input', function () {
        var maxLength = 250;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorpagekeyword.text("* "+languagedata.Seo.keywordmustnotexceed155characters);
        } else {
            // Clear error message if under the limit
            $errorpagekeyword.text('');
        }
    });

    var $storetitle = $('#storetitle');
    var $errorstoretitle = $('#error-storetitle');

    $storetitle.on('input', function () {
        var maxLength = 55;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorstoretitle.text("* "+languagedata.Seo.titlemustnotexceed55characters);
        } else {
            // Clear error message if under the limit
            $errorstoretitle.text('');
        }
    });

    var $storedescription = $('#storedescription');
    var $errorstoredescription = $('#error-storedescription');

    $storedescription.on('input', function () {
        var maxLength = 155;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorstoredescription.text("* "+languagedata.Seo.descriptionmustnotexceed155characters);
        } else {
            // Clear error message if under the limit
            $errorstoredescription.text('');
        }
    });

    var $storekeyword = $('#storekeyword');
    var $errorstorekeyword = $('#error-storekeyword');

    $storekeyword.on('input', function () {
        var maxLength = 155;
        if ($(this).val().length >= maxLength) {
            // Show error message
            $errorstorekeyword.text("* "+languagedata.Seo.keywordmustnotexceed155characters);
        } else {
            // Clear error message if under the limit
            $errorstorekeyword.text('');
        }
    });



    $("#PageSave").click(function () {

        $("#HomePage").validate({

            ignore: [],

            rules: {
                pagetitle: {
                    required: true
                },
                pagedescription: {
                    required: true
                },
                pagekeyword: {
                    required: true
                }
            },
            messages: {
                pagetitle: {
                    required: "* "+languagedata.Seo.pleaseentertitle
                },
                pagedescription: {
                    required: "* "+languagedata.Seo.pleaseenterdescripion
                },
                pagekeyword: {
                    required: "* "+languagedata.Seo.pleaseenterkeyword
                }
            }
        })

        var pagecheck = $("#HomePage").valid();
        if (pagecheck == true) {
            $('#HomePage')[0].submit();

        }

    })

    $("#StoreSave").click(function () {

        $("#StoreData").validate({

            ignore: [],

            rules: {
                storetitle: {
                    required: true
                },
                storedescription: {
                    required: true
                },
                storekeyword: {
                    required: true
                }
            },
            messages: {
                storetitle: {
                    required: "* "+languagedata.Seo.pleaseentertitle
                },
                storedescription: {
                    required: "* "+languagedata.Seo.pleaseenterdescripion
                },
                storekeyword: {
                    required: "* "+languagedata.Seo.pleaseenterkeyword
                }
            }
        })

        var storecheck = $("#StoreData").valid();
        if (storecheck == true) {
            $('#StoreData')[0].submit();

        }

    })

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

                console.log("base64Data:", base64Data);


                $("#sitemapimage").val(base64Data)
                $("#sitemap").val(base64Data)
            };
            reader.readAsDataURL(file);
        }

    });

    // Optional: Delete/reset file input
    $('#deleteFile').on('click', function () {
        $('.fileInput').val('');
        $('#filenameDisplay').text(languagedata.Seo.nofilechosen);
        $(this).addClass('hidden');
        $("#sitemapimage").val("")
        $("#sitemap").val("")
    });


    $("#SiteMapSave").click(function () {

        $("#SiteMap").validate({

            ignore: [],

            rules: {
                sitemapimage: {
                    required: true
                }
            },
            messages: {
                sitemapimage: {
                    required: "* "+languagedata.Seo.pleasechooseimage
                }
            }
        })

        var sitemap = $("#SiteMap").valid();
        if (sitemap == true) {
            $('#SiteMap')[0].submit();

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