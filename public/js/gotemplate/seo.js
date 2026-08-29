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

  /* ── StoreData save ─────────────────────────────────────────────────────── */
    $('#StoreSave').click(function () {
        $('#StoreData').validate({
            ignore: [],
            rules:    { storetitle: { required: true }, storedescription: { required: true }, storekeyword: { required: true } },
            messages: {
                storetitle:       { required: '* ' + languagedata.Seo.pleaseentertitle },
                storedescription: { required: '* ' + languagedata.Seo.pleaseenterdescripion },
                storekeyword:     { required: '* ' + languagedata.Seo.pleaseenterkeyword }
            }
        });
        if ($('#StoreData').valid()) $('#StoreData')[0].submit();
    });
 
/* ── Sitemap XML upload ───────────────────────────────────────────────────── */
 
/**
* Three display states:
*  'upload'  → show the Upload XML button  (no file selected / file deleted)
*  'server'  → show the saved-file chip    (file already saved on server)
*  'js'      → show the newly-picked chip  (file chosen, not yet saved)
*/
function showSitemapState(state) {
    $('#upload-placeholder').toggleClass('hidden', state !== 'upload');
    $('#server-file-wrap').toggleClass('hidden',   state !== 'server');
    $('#js-file-preview').toggleClass('hidden',    state !== 'js');
}
 
/* File picked */
$('#sitemap').on('change', function () {
    var files = $(this).prop('files');
    if (!files || files.length === 0) return;
 
    var file  = files[0];
    var isXml = ['text/xml', 'application/xml'].includes(file.type)
                || file.name.toLowerCase().endsWith('.xml');
 
    if (!isXml) {
        $('#error-sitemap').text('Please upload a valid XML file only.');
        $(this).val('');
        return;
    }
 
    $('#error-sitemap').text('');
    $('#js-file-name').text(file.name);
    $('#js-file-icon').html('<img src="/public/img/xml.jpg" class="w-[60px] h-[60px] object-contain rounded" alt="xml file">');
    showSitemapState('js');
});
 
/* Delete newly-picked file  →  back to upload */
$(document).on('click', '#deleteFileJs', function () {
    $('#sitemap').val('');
    $('#js-file-name').text('');
    showSitemapState('upload');
});
 
/* Delete existing server file  →  back to upload + mark for removal */
$(document).on('click', '#deleteFile', function () {
    $('#sitemap').val('');
    $('#deleteSitemap').val('1');   // tells the backend to clear the record
    showSitemapState('upload');
});
 
/* SiteMap save */
$('#SiteMapSave').off('click').on('click', function () {
    $('#SiteMap').validate({
        ignore: [],
        rules:    { sitemap: { required: true } },
        messages: { sitemap: { required: '* ' + languagedata.Seo.pleasechooseimage } }
    });
    if ($('#SiteMap').valid()) $('#SiteMap')[0].submit();
});
 
/* Web-banner toggle */
$('.hd-crd-btn').click(function () {
    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass('show');
        document.cookie = 'webbanner=false; path=/;';
    } else {
        $('#hd-crd').addClass('show').removeClass('hidden');
        document.cookie = 'webbanner=true; path=/;';
    }
});



// File type detection for server files
document.addEventListener('DOMContentLoaded', function() {
    const filename = document.getElementById('server-filename')?.textContent?.trim();
    if (!filename) return;
    if (!filename.toLowerCase().endsWith('.xml')) {
        document.getElementById('server-xml-icon').style.display = 'none';
        document.getElementById('server-image-preview').classList.remove('hidden');
    } else {
        document.getElementById('server-xml-icon').src = '/public/img/xml.jpg';
    }
});

// File input handling (existing JS should handle this)
// document.getElementById('sitemap')?.addEventListener('change', function(e) {
//     // Your existing file preview logic
// });

// Add to your seo.js
const sitemapInput = document.getElementById('sitemap');
const sitemapImageInput = document.getElementById('sitemapimage');
const uploadPlaceholder = document.getElementById('upload-placeholder');
const serverFileWrap = document.getElementById('server-file-wrap');
const jsFilePreview = document.getElementById('js-file-preview');
const jsFileIcon = document.getElementById('js-file-icon');
const jsFileName = document.getElementById('js-file-name');
const deleteFile = document.getElementById('deleteFile');
const deleteFileJs = document.getElementById('deleteFileJs');

sitemapInput.addEventListener('change', function(e) {
    const file = e.target.files[0];
    if (file) {
        jsFileName.textContent = file.name;
        
        // Show correct icon
       if (file.name.toLowerCase().endsWith('.xml')) {
    jsFileIcon.innerHTML = `<img src="/public/img/xml.jpg" class="w-[60px] h-[60px] object-contain rounded" alt="xml file">`;
}
     else {
            // Image preview
            const reader = new FileReader();
            reader.onload = (e) => {
                jsFileIcon.innerHTML = `<img src="${e.target.result}" class="w-[20px] h-[20px] object-cover rounded" />`;
            };
            reader.readAsDataURL(file);
        }
        
        // Convert to base64 for backend
        const reader = new FileReader();
        reader.onload = function(e) {
            let mimeType = file.type || 'application/octet-stream';
            if (file.name.toLowerCase().endsWith('.xml')) {
                mimeType = 'application/xml';
            }
            sitemapImageInput.value = `data:${mimeType};base64,${btoa(String.fromCharCode(...new Uint8Array(e.target.result)))}`;
            
            // Show preview
            uploadPlaceholder.classList.add('hidden');
            jsFilePreview.classList.remove('hidden');
            serverFileWrap.classList.add('hidden');
        };
        reader.readAsArrayBuffer(file);
    }
});

// Delete buttons
deleteFile.addEventListener('click', () => {
    sitemapImageInput.value = '';
    serverFileWrap.classList.add('hidden');
    uploadPlaceholder.classList.remove('hidden');
});

deleteFileJs.addEventListener('click', () => {
    sitemapInput.value = '';
    sitemapImageInput.value = '';
    jsFilePreview.classList.add('hidden');
    uploadPlaceholder.classList.remove('hidden');
});

/* SiteMap Save - FIXED validation */
$('#SiteMapSave').off('click').on('click', function () {
    // Destroy existing validation first
    $('#SiteMap').validate().destroy();
    
    $('#SiteMap').validate({
        ignore: [],
        errorClass: 'custom-error',
        errorElement: 'div',
        errorPlacement: function(error, element) {
            // Place error DIRECTLY in your #error-sitemap div
            
            $('#error-sitemap').html(error.text()).show();
        },
        rules: { 
            sitemap: { required: true }  // Changed from sitemapimage to sitemap
        },
        messages: { 
            sitemap: { 
                required: '* ' + languagedata.Seo.pleasechooseimage 
            } 
        }
    });
    
    if ($('#SiteMap').valid()) {
        $('#SiteMap')[0].submit();
    }
});