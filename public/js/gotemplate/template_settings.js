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


// 

$(document).ready(function () {

     // stores template temporarily when channel not selected

    // CLICK ON CHANNEL
    $(document).on("click", ".dropdown-channel .dropdown-menu a", function () {        

        let chosenName = $(this).text().trim();
        let chosenId = $(this).data("id");
        let index = $(this).closest(".dropdown-channel").data("index");

        // Prevent duplicate channel selection
        if (selectedChannels.some(ch => ch.id === chosenId)) {
            $("#Channeltemplate-error").text("This channel is already selected. Please choose another one.").removeClass("hidden");
            return;
        }
        $("#Channeltemplate-error").addClass("hidden");

        // Remove old entry for same index
        selectedChannels = selectedChannels.filter(ch => ch.index !== index);

        // Add Channel
        let newObj = {
            index: index,
            id: chosenId,
            name: chosenName,
            templatetype: ""
        };

        // If template was selected earlier, apply it
        if (pendingTemplates[index]) {
            newObj.templatetype = pendingTemplates[index];
            delete pendingTemplates[index]; // clear once applied
        }

        selectedChannels.push(newObj);

        $(this).closest(".dropdown").find("a:first").text(chosenName);

    });

    // CLICK ON TEMPLATE
    $(document).on("click", ".dropdown-template .dropdown-menu a", function () {

        let chosenTemplate = $(this).text().trim();
        let index = $(this).closest(".dropdown-template").data("index");

        // Find existing channel for this index
        let obj = selectedChannels.find(ch => ch.index === index);

        if (obj) {
            // Channel selected → update template
            obj.templatetype = chosenTemplate;
        } else {
            // Channel NOT selected → store temporarily
            pendingTemplates[index] = chosenTemplate;
        }

        $(this).closest(".dropdown").find("a:first").text(chosenTemplate);

     
    });


    // pair of template type and channel validation


    // UPDATE BUTTON CLICK
$(document).on("click", "#channeltemplateupdate", function (e) {
    
    e.preventDefault();  // stop form submit first

    // If any template exists without channel
    if (Object.keys(pendingTemplates).length > 0) {
        $("#Channeltemplate-error")
            .text("Please select a channel for all chosen templates.")
            .removeClass("hidden");
        return; // stop submit
    }

    // If no channels selected at all (optional, but useful)
    if (selectedChannels.length === 0) {
        $("#Channeltemplate-error")
            .text("Please select at least one channel.")
            .removeClass("hidden");
        return;
    }

    $("#channel_template_data").val(JSON.stringify(selectedChannels));
    $("#Channeltemplate-error").addClass("hidden");
    

    // Everything valid → submit form
    $("#channeltemplatform")[0].submit();
});


});

$(document).ready(function () {

    $("#socialmediaupdate").on("click", function (e) {

        e.preventDefault(); 

        let FaceBookLink = $("#facebooklink").val();
        let LinkedinLink = $("#linkedinlink").val();
        let YoutubeLink = $("#youtubelink").val();
        let XSocialLink = $("#xlink").val();
        let InstaLink = $("#instagramlink").val();

        let obj ={
            "Type" :"facebook",
            "SocialUrl":FaceBookLink,
            "IsActive": 0
        }

        let newObject = {
            FaceBooklink: FaceBookLink,
            Linkedinlink: LinkedinLink,
            Youtubelink: YoutubeLink,
            XSociallink: XSocialLink,
            Instalink: InstaLink
        };

       
        const socialConfig = [
            { type: "FaceBooklink", checkbox: "ck4", input: "#facebooklink" },
            { type: "Linkedinlink", checkbox: "ck5", input: "#linkedinlink" },
            { type: "Youtubelink", checkbox: "ck2", input: "#youtubelink" },
            { type: "XSociallink", checkbox: "ck1", input: "#xlink" },
            { type: "Instalink", checkbox: "ck3", input: "#instagramlink" }
        ];

        let socialMediaData = [];

        socialConfig.forEach(item => {
            const isChecked = document.getElementById(item.checkbox).checked;
            const link = $(item.input).val();

            socialMediaData.push({
                Type: item.type,
                SocialUrl: link,
                IsActive: isChecked ? 1 : 0
            });
        });

        console.log("checked or not ", socialMediaData)

        // let selectLinks = [socialMediaData]

  
        $("#social_media_data").val(JSON.stringify(socialMediaData));

        $("#SocialmediaForm")[0].submit(); 
    });

});

// update social link


let link_json = $("#social_link_json").val()

let savedlinks = JSON.parse(link_json);

savedlinks.forEach((item, index) => {


    if (item.Type == "Instalink") {

        if (item.SocialUrl !="") {

            $("#instagramlink").val(item.SocialUrl)

        }
        if (item.IsActive == 1) {
            $("#ck3").prop("checked", true); // ✅ check the box
        } else {
            $("#ck3").prop("checked", false);
        }

    }
    if (item.Type == "XSociallink") {

        if (item.SocialUrl != "") {

            $("#xlink").val(item.SocialUrl)

        }
        if (item.IsActive == 1) {
            $("#ck1").prop("checked", true); // ✅ check the box
        } else {
            $("#ck1").prop("checked", false);
        }

    }

    if (item.Type == "Youtubelink") {

        if (item.SocialUrl != "") {

            $("#youtubelink").val(item.SocialUrl)

        }
        if (item.IsActive == 1) {
            $("#ck2").prop("checked", true); // ✅ check the box
        } else {
            $("#ck2").prop("checked", false);
        }

    }
    if (item.Type == "FaceBooklink") {

        if (item.SocialUrl != "") {

            $("#facebooklink").val(item.SocialUrl)

        }
        if (item.IsActive == 1) {
            $("#ck4").prop("checked", true); // ✅ check the box
        } else {
            $("#ck4").prop("checked", false);
        }

    }
  
    

    if (item.Type == "Linkedinlink") {

        if (item.SocialUrl != "") {

            $("#linkedinlink").val(item.SocialUrl)

        }
        if (item.IsActive == 1) {
            $("#ck5").prop("checked", true); // ✅ check the box
        } else {
            $("#ck5").prop("checked", false);
        }

    }


  
});




let raw = $("#template_json").val();
    
        let savedChannelTemplate = [];
        try {
            savedChannelTemplate = JSON.parse(raw)
        } catch (e) {
            console.error("Invalid JSON:", raw);
        }
    
    
        $(document).ready(function () {
    
            if (Array.isArray(savedChannelTemplate)) {
    
                savedChannelTemplate.forEach(item => {
    
                    // Fill dropdown visually
                    let channelDropdown = $(`.dropdown-channel[data-index='${item.index}'] a:first`);
                    channelDropdown.text(item.name);
                    channelDropdown.data("selected", item.name);
    
                    let templateDropdown = $(`.dropdown-template[data-index='${item.index}'] a:first`);
                    templateDropdown.text(item.templatetype || "Choose Template");
                    templateDropdown.data("selected", item.templatetype);
    
                    // Add to selectedChannels
                    selectedChannels.push({
                        index: item.index,
                        id: item.id,
                        name: item.name,
                        templatetype: item.templatetype
                    });
                });
    
            }
        });


// set header thame what selected

let headerthame = $("#headerthame").val()

console.log("dfjhdsj", headerthame)

$(document).ready(function () {
    let headerthame = $("#headerthame").val();
    console.log("headerthame:", headerthame);

    if (headerthame) {
        $('input[type="radio"][name="radio"][value="' + headerthame + '"]').prop('checked', true);
    }
});