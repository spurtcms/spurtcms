let button
var channelid = []
var channelname = []


$(document).ready(async function () {

    var currentURL = window.location.href;


    if (currentURL.includes('edit')) {

        if ($('#slchannel').attr('data-id') != "") {
            channelid = $('#slchannel').attr('data-id').split(",")
        }
        if ($('#slchannel').val() !== "") {

            channelname = $('#slchannel').val().split(",")
        }

        $(".select-chn").each(function () {
            var currentChannelId = $(this).attr('data-id');


            if (channelid.includes(currentChannelId)) {

                $(this).prop('checked', true);
            } else {

                $(this).prop('checked', false);
            }
        })
    } else {
        // $(".select-chn:first").trigger("click");
    }


    // You get innerHTML here    
    document.addEventListener('saveChange', function (event) {
        let homeurl
        spurtdata = event.detail
        if (button == "save-form") {
            homeurl = "/cta/draft"
        } else {
            homeurl = "/cta"
        }
        let first = spurtdata.data[0]

        var image = spurtdata.formImage
        let newimage = image.split('?name=')[1];
        var imagename =spurtdata.formImageName
        let value = first.value
        var id = $("#formid").val()

        if (channelid.length ==0 && channelname.length ==0){

           channelid.push( $('#defaultchnid').val())

           channelname.push('Default Channel')
        }
        if (id == "") {
            $.ajax({
                url: "/cta/createforms",
                type: "POST",
                async: false,
                data: { "button": button, "form": JSON.stringify(spurtdata), "image":newimage, "imagename":imagename, csrf: $("input[name='csrf']").val(), "title": value, "channelid": channelid, "channelname": channelname },
                datatype: "json",
                caches: false,
                success: function (data) {
                    window.location.href = homeurl
                }
            })
        } else {
            $.ajax({
                url: "/cta/updateforms",
                type: "POST",
                async: false,
                data: { "button": button, "form": JSON.stringify(spurtdata),"image":newimage, "imagename":imagename,csrf: $("input[name='csrf']").val(), "title": value, "id": id, "channelid": channelid, "channelname": channelname },
                datatype: "json",
                caches: false,
                success: function (data) {
                    window.location.href = homeurl
                }
            })
        }


    })
})


// publish btn
$(document).on("click", "#publish-form", function () {
    button = "publish-form"
    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);


})

// save btn
$(document).on("click", "#save-form", function () {
    button = "save-form"
    const event = new CustomEvent("getHTML", {
    });
    document.dispatchEvent(event);

})

// New code 
$(document).on("click", "#sl-chn", function () {
    $("#chn-list").toggleClass("show")
})

// Get channel fields
$(document).on('click', '.select-chn', function () {


    $(".selected-cat").empty()


    channellid = $(this).attr('data-id')

    $("#sl-chn-error").hide()
    $("#chn-name").text($(this).text()).addClass("text-bold-black")
    $("#chn-name").attr("data-slug", $(this).text())
    $("#chn-name").attr("data-id", channelid)

    if ($(this).is(':checked')) {

        channelname.push($(this).attr("data-slug"));
        channelid.push(channellid);
    } else {

        var index1 = channelname.indexOf($(this).attr("data-slug"));
        if (index1 !== -1) {
            channelname.splice(index1, 1);

        }
        var index2 = channelid.indexOf($(this).attr('data-id'));
        if (index2 !== -1) {

            channelid.splice(index2, 1);
        }
    }
    console.log(channelid, channelname, "arraylll")
})