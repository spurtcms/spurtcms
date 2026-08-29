var languagedata
var selectedcheckboxarr = []
let selectedChannels = [];
let pendingTemplates = {};

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {
        languagedata = data
    })

    $('.Content').addClass('checked');
})

$(document).ready(function () {

    // Site Name
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

    // Site Logo
    $('.fileInput').on('change', function () {
        var files = $(this).prop('files');
        if (files.length > 0) {
            var file = files[0];
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
                sitelogoimage: { required: true }
            },
            messages: {
                sitelogoimage: { required: "* " + languagedata.Seo.pleasechooseimage }
            }
        })

        var sitemap = $("#siteLogoForm").valid();
        if (sitemap == true) {
            $('#siteLogoForm')[0].submit();
        }
    })

    // Site FavIcon
    $('.fileInputFavIcon').on('change', function () {
        var files = $(this).prop('files');
        if (files.length > 0) {
            var file = files[0];
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
                sitefaviconimage: { required: true }
            },
            messages: {
                sitefaviconimage: { required: "* " + languagedata.Seo.pleasechooseimage }
            }
        })

        var sitemap = $("#siteFavIconForm").valid();
        if (sitemap == true) {
            $('#siteFavIconForm')[0].submit();
        }
    })

    // Website URL
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
            url: "/admin/website/checksitename",
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

    // ── Social Media Update ──────────────────────────────────────────
    $("#socialmediaupdate").on("click", function (e) {
        e.preventDefault();

        let FaceBookLink  = $("#facebooklink").val();
        let LinkedinLink  = $("#linkedinlink").val();
        let YoutubeLink   = $("#youtubelink").val();
        let XSocialLink   = $("#xlink").val();
        let InstaLink     = $("#instagramlink").val();
console.log(FaceBookLink, LinkedinLink, YoutubeLink, XSocialLink, InstaLink);

        // Remove previous error messages
        $(".error-msg").remove();

        let formvalidate = true;
        const urlPattern = /^(https?:\/\/)[^\s$.?#].[^\s]*$/i;

        if (FaceBookLink !== "" && !urlPattern.test(FaceBookLink)) {
            $("#facebooklink").closest("div").after("<div class='error-msg text-red-500 text-xs mt-1 ml-[125px]'>* Enter Valid Link</div>");
            formvalidate = false;
        }

        if (LinkedinLink !== "" && !urlPattern.test(LinkedinLink)) {
            $("#linkedinlink").closest("div").after("<p class='error-msg text-red-500 text-xs mt-1 ml-[125px]'>* Enter Valid Link</p>");
            formvalidate = false;
        }

        if (YoutubeLink !== "" && !urlPattern.test(YoutubeLink)) {
            $("#youtubelink").closest("div").after("<p class='error-msg text-red-500 text-xs mt-1 ml-[125px]'>* Enter Valid Link</p>");
            formvalidate = false;
        }

        if (XSocialLink !== "" && !urlPattern.test(XSocialLink)) {
            $("#xlink").closest("div").after("<p class='error-msg text-red-500 text-xs mt-1 ml-[125px]'>* Enter Valid Link</p>");
            formvalidate = false;
        }

        if (InstaLink !== "" && !urlPattern.test(InstaLink)) {
            $("#instagramlink").closest("div").after("<p class='error-msg text-red-500 text-xs mt-1 ml-[125px]'>* Enter Valid Link</p>");
            formvalidate = false;
        }

        if (!formvalidate) return;

        const socialConfig = [
            { type: "FaceBooklink", checkbox: "ck4", input: "#facebooklink" },
            { type: "Linkedinlink", checkbox: "ck5", input: "#linkedinlink" },
            { type: "Youtubelink", checkbox: "ck2", input: "#youtubelink" },
            { type: "XSociallink", checkbox: "ck1", input: "#xlink" },
            { type: "Instalink",   checkbox: "ck3", input: "#instagramlink" }
        ];

        let socialMediaData = [];

        socialConfig.forEach(function(item) {
            const isChecked = document.getElementById(item.checkbox) ? document.getElementById(item.checkbox).checked : false;
            const link = $(item.input).val();

            socialMediaData.push({
                Type:      item.type,
                SocialUrl: link,
                IsActive:  isChecked ? 1 : 0
            });
        });

        console.log("social media data:", socialMediaData);

        $("#social_media_data").val(JSON.stringify(socialMediaData));

        // FIX: was "#SocialmediaForm" — now correctly targets the social links form
        $("#socialLinksForm")[0].submit();
    });

})

// ── Header theme card toggle ─────────────────────────────────────────
$('.hd-crd-btn').click(function () {
    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show");
        document.cookie = `webbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `webbanner=true; path=/;`;
    }
});

// ── Channel Template ─────────────────────────────────────────────────
$(document).ready(function () {

    // CLICK ON CHANNEL
    $(document).on("click", ".dropdown-channel .dropdown-menu a", function () {
        let chosenName = $(this).text().trim();
        let chosenId   = $(this).data("id");
        let index      = $(this).closest(".dropdown-channel").data("index");

        if (selectedChannels.some(ch => ch.id === chosenId)) {
            $("#Channeltemplate-error").text("This channel is already selected. Please choose another one.").removeClass("hidden");
            return;
        }
        $("#Channeltemplate-error").addClass("hidden");

        selectedChannels = selectedChannels.filter(ch => ch.index !== index);

        let newObj = {
            index: index,
            id:    chosenId,
            name:  chosenName,
            templatetype: ""
        };

        if (pendingTemplates[index]) {
            newObj.templatetype = pendingTemplates[index];
            delete pendingTemplates[index];
        }

        selectedChannels.push(newObj);
        $(this).closest(".dropdown").find("a:first").text(chosenName);
    });

    // CLICK ON TEMPLATE
    $(document).on("click", ".dropdown-template .dropdown-menu a", function () {
        let chosenTemplate = $(this).text().trim();
        let index          = $(this).closest(".dropdown-template").data("index");

        let obj = selectedChannels.find(ch => ch.index === index);

        if (obj) {
            obj.templatetype = chosenTemplate;
        } else {
            pendingTemplates[index] = chosenTemplate;
        }

        $(this).closest(".dropdown").find("a:first").text(chosenTemplate);
    });

    // UPDATE BUTTON CLICK
    $(document).on("click", "#channeltemplateupdate", function (e) {
        e.preventDefault();

        let websiteInput = $("#websiteInput").val();


        if (websiteInput === "") {
            $("#websiteInput-error").show().text("* " + languagedata.WebsiteSettings.pleaseenteryourwebsiteurl);
            return;
        }


        if (Object.keys(pendingTemplates).length > 0) {
            $("#Channeltemplate-error")
                .text("* Please select a channel for all chosen templates.")
                .removeClass("hidden");
            return;
        }

        if (selectedChannels.length === 0) {
            $("#Channeltemplate-error")
                .text("* Please select at least one channel.")
                .removeClass("hidden");
            return;
        }

        $("#channel_template_data").val(JSON.stringify(selectedChannels));
        $("#Channeltemplate-error").addClass("hidden");
        $("#SocialmediaForm")[0].submit();
    });

});

// ── Restore saved social links ────────────────────────────────────────
$(document).ready(function () {

    let link_json  = $("#social_link_json").val() || "[]";
    let savedlinks = [];

    try {
        savedlinks = JSON.parse(link_json) || [];
    } catch (e) {
        console.warn("social_link_json parse error:", e);
        savedlinks = [];
    }

    savedlinks.forEach(function(item) {

        if (item.Type === "Instalink") {
            if (item.SocialUrl !== "") $("#instagramlink").val(item.SocialUrl);
            $("#ck3").prop("checked", item.IsActive === 1);
        }

        if (item.Type === "XSociallink") {
            if (item.SocialUrl !== "") $("#xlink").val(item.SocialUrl);
            $("#ck1").prop("checked", item.IsActive === 1);
        }

        if (item.Type === "Youtubelink") {
            if (item.SocialUrl !== "") $("#youtubelink").val(item.SocialUrl);
            $("#ck2").prop("checked", item.IsActive === 1);
        }

        if (item.Type === "FaceBooklink") {
            if (item.SocialUrl !== "") $("#facebooklink").val(item.SocialUrl);
            $("#ck4").prop("checked", item.IsActive === 1);
        }

        if (item.Type === "Linkedinlink") {
            if (item.SocialUrl !== "") $("#linkedinlink").val(item.SocialUrl);
            $("#ck5").prop("checked", item.IsActive === 1);
        }
    });

});

// ── Restore saved channel templates ──────────────────────────────────
$(document).ready(function () {

    let raw = $("#template_json").val() || "[]";
    let savedChannelTemplate = [];

    try {
        savedChannelTemplate = JSON.parse(raw) || [];
    } catch (e) {
        console.error("template_json parse error:", e);
        savedChannelTemplate = [];
    }

    if (Array.isArray(savedChannelTemplate)) {
        savedChannelTemplate.forEach(function(item) {
            let channelDropdown = $(`.dropdown-channel[data-index='${item.index}'] a:first`);
            channelDropdown.text(item.name);
            channelDropdown.data("selected", item.name);

            let templateDropdown = $(`.dropdown-template[data-index='${item.index}'] a:first`);
            templateDropdown.text(item.templatetype || "Choose Template");
            templateDropdown.data("selected", item.templatetype);

            selectedChannels.push({
                index:        item.index,
                id:           item.id,
                name:         item.name,
                templatetype: item.templatetype
            });
        });
    }

});

// ── Restore header theme radio selection ─────────────────────────────
$(document).ready(function () {

    let headerthame = $("#headerthame").val() || "";
    console.log("headerthame:", headerthame);

    if (headerthame) {
        $('input[type="radio"][name="headertheme"][value="' + headerthame + '"]').prop('checked', true);
    }

});

// ===================================================================


// ═══════════════════════════════════════════════════════════════════════════
// SITE NAME FORM — inject current social media data before submit
// FIX: prevents brand form from overwriting saved social links with empty
// ═══════════════════════════════════════════════════════════════════════════
document.getElementById('siteNameSave').addEventListener('submit', function () {
    const socialConfig = [
        { type: "FaceBooklink", checkbox: "ck4", input: "#facebooklink"  },
        { type: "Linkedinlink", checkbox: "ck5", input: "#linkedinlink"  },
        { type: "Youtubelink",  checkbox: "ck2", input: "#youtubelink"   },
        { type: "XSociallink",  checkbox: "ck1", input: "#xlink"         },
        { type: "Instalink",    checkbox: "ck3", input: "#instagramlink" }
    ];
    const socialMediaData = socialConfig.map(function(item) {
        const cb = document.getElementById(item.checkbox);
        return {
            Type:      item.type,
            SocialUrl: $(item.input).val(),
            IsActive:  cb && cb.checked ? 1 : 0
        };
    });
    document.getElementById('siteNameSave_social_media_data').value = JSON.stringify(socialMediaData);
});

// ✅ Site Logo — file change

$('#sitelogo-file-input').on('change', function () {

    var file = this.files[0];

    if (!file) return;
 
    var reader = new FileReader();

    reader.onload = function (e) {

        $('#sitelogoimage_data').val(e.target.result);

        $('#sitelogo-preview').attr('src', e.target.result);

        $('#sitelogo-preview-wrap').removeClass('hidden').addClass('flex');

        $('#sitelogo-upload-label').addClass('hidden');

        $('#sitelogo-delete-btn').removeClass('hidden');

        $("#sitelogoDlt").val("0")

    };

    reader.readAsDataURL(file);

});
 
function clearSiteLogo() {

    $('#sitelogoimage_data').val('DELETE');

    $('#sitelogo-file-input').val('');

    $('#sitelogo-preview').attr('src', '');

    $('#sitelogo-preview-wrap').addClass('hidden').removeClass('flex');

    $('#sitelogo-upload-label').removeClass('hidden');

    $('#sitelogo-delete-btn').addClass('hidden');

    $("#sitelogoDlt").val("1")

}
 
// ✅ Site Favicon — file change

$('#sitefavicon-file-input').on('change', function () {

    var file = this.files[0];

    if (!file) return;
 
    var reader = new FileReader();

    reader.onload = function (e) {

        $('#sitefaviconimage_data').val(e.target.result);

        $('#sitefavicon-preview').attr('src', e.target.result);

        $('#sitefavicon-preview-wrap').removeClass('hidden').addClass('flex');

        $('#sitefavicon-upload-label').addClass('hidden');

        $('#sitefavicon-delete-btn').removeClass('hidden');

        $("#sitefaviconDlt").val("0")

    };

    reader.readAsDataURL(file);

});
 
function clearSiteFavicon() {

    $('#sitefaviconimage_data').val('DELETE');

    $('#sitefavicon-file-input').val('');

    $('#sitefavicon-preview').attr('src', '');

    $('#sitefavicon-preview-wrap').addClass('hidden').removeClass('flex');

    $('#sitefavicon-upload-label').removeClass('hidden');

    $('#sitefavicon-delete-btn').addClass('hidden');

    $("#sitefaviconDlt").val("1")

}
 



// ═══════════════════════════════════════════════════════════════════════════
// SITE FAVICON — file → base64 → hidden input inside siteNameSave + preview
// ═══════════════════════════════════════════════════════════════════════════
document.getElementById('sitefavicon-file-input').addEventListener('change', function () {
    var file = this.files[0];
    if (!file) return;
    var reader = new FileReader();
    reader.onload = function (e) {
        document.getElementById('sitefaviconimage_data').value = e.target.result;
        document.getElementById('sitefavicon-preview').src = e.target.result;
        document.getElementById('sitefavicon-preview-wrap').classList.remove('hidden');
        document.getElementById('sitefavicon-preview-wrap').classList.add('flex');
        document.getElementById('sitefavicon-upload-label').classList.add('hidden');
        document.getElementById('sitefavicon-delete-btn').classList.remove('hidden');
    };
    reader.readAsDataURL(file);
});

function clearSiteFavicon() {
    document.getElementById('sitefaviconimage_data').value = 'DELETE';
    document.getElementById('sitefavicon-file-input').value = '';
    document.getElementById('sitefavicon-preview').src = '';
    document.getElementById('sitefavicon-preview-wrap').classList.add('hidden');
    document.getElementById('sitefavicon-preview-wrap').classList.remove('flex');
    document.getElementById('sitefavicon-upload-label').classList.remove('hidden');
    document.getElementById('sitefavicon-delete-btn').classList.add('hidden');
    $("#sitefaviconDlt").val("1")
}

// ═══════════════════════════════════════════════════════════════════════════
// SOCIAL MEDIA FORM — pack links into JSON before submit
// FIX: updated from 'SocialmediaForm' to 'socialLinksForm'
// ═══════════════════════════════════════════════════════════════════════════
document.getElementById('socialLinksForm').addEventListener('submit', function (e) {
    var socialData = {
        linkedin:  document.getElementById('linkedinlink').value,
        x:         document.getElementById('xlink').value,
        youtube:   document.getElementById('youtubelink').value,
        instagram: document.getElementById('instagramlink').value,
        facebook:  document.getElementById('facebooklink').value
    };
    document.getElementById('social_media_data').value = JSON.stringify(socialData);
});

// ═══════════════════════════════════════════════════════════════════════════
// LOAD TEMPLATE SETTING TAB VIA AJAX
// ═══════════════════════════════════════════════════════════════════════════
function loadTemplateSettingAPI() {
    fetch("/admin/website/settings/templatesetting?webid=1")
        .then(function (res) { return res.text(); })
        .then(function (html) {
            var parser = new DOMParser();
            var doc = parser.parseFromString(html, "text/html");
            var newContent = doc.querySelector("#pills-seo");

            if (newContent) {
                document.getElementById("pills-seo").innerHTML = newContent.innerHTML;
                bindDynamicEvents();
            }
        })
        .catch(function (err) { console.error("API Error:", err); });
}

// ═══════════════════════════════════════════════════════════════════════════
// RE-BIND EVENTS after loadTemplateSettingAPI replaces #pills-seo innerHTML
// FIX: updated socialForm lookup from 'SocialmediaForm' to 'socialLinksForm'
// ═══════════════════════════════════════════════════════════════════════════
function bindDynamicEvents() {

    var logoInput = document.getElementById('sitelogo-file-input');
    if (logoInput) {
        logoInput.addEventListener('change', function () {
            var file = this.files[0];
            if (!file) return;
            var reader = new FileReader();
            reader.onload = function (e) {
                document.getElementById('sitelogo-preview').src = e.target.result;
                document.getElementById('sitelogoimage_data').value = e.target.result;
            };
            reader.readAsDataURL(file);
        });
    }

    var faviconInput = document.getElementById('sitefavicon-file-input');
    if (faviconInput) {
        faviconInput.addEventListener('change', function () {
            var file = this.files[0];
            if (!file) return;
            var reader = new FileReader();
            reader.onload = function (e) {
                document.getElementById('sitefavicon-preview').src = e.target.result;
                document.getElementById('sitefaviconimage_data').value = e.target.result;
            };
            reader.readAsDataURL(file);
        });
    }

    // FIX: was 'SocialmediaForm' — now 'socialLinksForm'
    var socialForm = document.getElementById('socialLinksForm');
    if (socialForm) {
        socialForm.addEventListener('submit', function () {
            var socialData = {
                linkedin:  document.getElementById('linkedinlink').value,
                x:         document.getElementById('xlink').value,
                youtube:   document.getElementById('youtubelink').value,
                instagram: document.getElementById('instagramlink').value,
                facebook:  document.getElementById('facebooklink').value
            };
            document.getElementById('social_media_data').value = JSON.stringify(socialData);
        });
    }
}

// ═══════════════════════════════════════════════════════════════════════════
// PRE-POPULATE SOCIAL LINKS from server-rendered JSON
// ═══════════════════════════════════════════════════════════════════════════
document.addEventListener('DOMContentLoaded', function () {
    try {
        var raw = document.getElementById('social_link_json').value;
        if (!raw) return;
        var data = JSON.parse(raw);
        if (data.linkedin)  document.getElementById('linkedinlink').value  = data.linkedin;
        if (data.x)         document.getElementById('xlink').value         = data.x;
        if (data.youtube)   document.getElementById('youtubelink').value   = data.youtube;
        if (data.instagram) document.getElementById('instagramlink').value = data.instagram;
        if (data.facebook)  document.getElementById('facebooklink').value  = data.facebook;
    } catch (e) {
        console.warn('Could not parse social_link_json:', e);
    }
});
