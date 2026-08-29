let button
var channelid = []
var channelname = []

$(document).ready(function () {
    let formimages = $("#formimages").val()

    if (formimages == "") {
        $("#ctimagehide").hide()
        $("#catdel-img").hide()
    } else {
        $("#ctimagehide").show()
        $("#catdel-img").show()
        $("#browse").hide()
        $("#uploadLine").hide()
        $("#uploadFormat").hide()
    }


    $('#catdel-img').click(function () {
        $('#categoryimages').val("")
        $('#ctimagehide').attr('src', '')
        $('#ctimagehide').hide()
        $('#browse').show()
        $("#mediadesc").css("margin-top", "1%")

        $(this).siblings('p,button').show()
        $(this).hide()

    })

    $('.tab-togg').click(function () {
        $('.editor-tabs').toggleClass('translate-x-[100%]');
        $(".imgtoogle").toggleClass('rotate-180')

    });

    $("#formTitle").on("input", function () {
        let formTitle = $(this).val().trim();

        if (!formTitle) {
            if (!$("#formTitleError").length) {
                $(this).after(
                    '<label id="formTitleError" class="text-sm mt-1 error" style="color:#F26674;font-weight: 400;font-size: 0.75rem;line-height: 1rem;">* This field is required</label>'
                );
            }
        } else {
            $("#formTitleError").remove();
        }
    });


});


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

    $(document).on('click', '#formimage', function () {
        $("#prof-crop").val("12")
    })

    // publish btn
    $(document).on("click", "#publish-form", function () {
        let formTitle = $("#formTitle").val().trim();
        if (formTitle != "") {
            button = "publish-form"
            const event = new CustomEvent("getHTML", {
            });
            document.dispatchEvent(event);
            return
        } else {
            $("#formTitleError").remove();
            $("#formTitle").after(
                '<label id="formTitleError" class="text-sm mt-1 error" style="color:#F26674;font-weight: 400;font-size: 0.75rem;line-height: 1rem;">* This field is required</label>'
            );
            $('.editor-tabs').removeClass('translate-x-[100%]');
            $(".imgtoogle").addClass('rotate-180')
        }
    })


    // You get innerHTML here    
    document.addEventListener('saveChange', function (event) {

        spurtdata = event.detail
        if (button == "save-form") {
            homeurl = "/admin/cta/draft"
        } else {
            homeurl = "/admin/cta"
        }

        let first = spurtdata.data[0]

        var image = spurtdata.formImage
        let newimage = image.split('?name=')[1];
        var imagename = spurtdata.formImageName
        let value = first.value
        var id = $("#formid").val()
        let formTitle = $("#formTitle").val().trim();
        let description = $("#description").val().trim()
        let formimages = $("#formimages").val()

        if (channelid.length == 0 && channelname.length == 0) {

            channelid.push($('#defaultchnid').val())

            channelname.push('Default Channel')
        }
        if (id == "") {
            $.ajax({
                url: "/admin/cta/createforms",
                method: "POST",
                dataType: 'json',
                data: { "button": button, "form": JSON.stringify(spurtdata), "image": newimage, "imagename": imagename, csrf: $("input[name='csrf']").val(), "title": formTitle, "description": description, "formimages": formimages, "channelid": channelid, "channelname": channelname },
                success: function (response) {
                    window.location.href = response.data
                }
            })
        } else {
            $.ajax({
                url: "/admin/cta/updateforms",
                method: "POST",
                dataType: 'json',
                data: { "button": button, "form": JSON.stringify(spurtdata), "image": newimage, "imagename": imagename, csrf: $("input[name='csrf']").val(), "title": formTitle, "description": description, "formimages": formimages, "id": id, "channelid": channelid, "channelname": channelname },
                success: function (response) {
                    window.location.href = response.data
                }
            })
        }


    })
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
})