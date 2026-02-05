//jquery ready function
$(document).ready(async function () {
    var fieldDiv, languageData;

    //retrieving langugae json data
    var languagepath = $('.language-group>button').attr('data-path')
    await $.getJSON(languagepath, function (data) {
        languageData = data
    })

    //selecting the payload list item
    $(document).on('click', '.payload-list-item', function (e) {
        var payloadType = $(this).text().trim()
        $(this).parents('.dropdown').find('a>p#slct-payload-item').removeClass('text-[#B2B2B2]').addClass('text-[#152027]').text(payloadType)
        $(this).parents('.dropdown').find('input[name=wh_payload]').val(payloadType.toLowerCase())
        if ($('#payload-del-img').length == 0) {
            var dropdownBtn = $(this).parent().siblings('#slct-payload-btn')
            $(dropdownBtn).removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
            var delImg = '<img src="/public/img/removeCover.svg" id="payload-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
            $(delImg).insertAfter(dropdownBtn)
        }
    })

    //removing payoad type
    $(document).on('click', '#payload-del-img', function () {
        $('input[name=wh_payload]').val("")
        $('#slct-payload-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctMethod)
        // var dropdownBtn = $(this).parent().siblings('#slct-payload-btn')
        $('#slct-payload-btn').addClass("bg-[url('/public/img/property-arrow.svg')]").attr('data-bs-toggle', 'dropdown')
        $(this).remove()
    })

    //selecting the method list item
    $(document).on('click', '.method-list-item', function () {
        var method = $(this).text().trim()
        $(this).parents('.dropdown').find('a>p#slct-method-item').removeClass('text-[#B2B2B2]').addClass('text-[#152027]').text(method)
        $(this).parents('.dropdown').find('input[name=wh_method]').val(method.toLowerCase())
        if ($('#method-del-img').length == 0) {
            var dropdownBtn = $(this).parent().siblings('#slct-method-btn')
            $(dropdownBtn).removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
            var delImg = '<img src="/public/img/removeCover.svg" id="method-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
            $(delImg).insertAfter(dropdownBtn)
        }
        $('#wh_method-error').hide()
    })

    // removing request method
    $(document).on('click', '#method-del-img', function () {
        $('input[name=wh_method]').val("")
        $('#slct-method-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctMethod)
        // var dropdownBtn = $(this).parent().siblings('#slct-method-btn')
        $('#slct-method-btn').addClass("bg-[url('/public/img/property-arrow.svg')]").attr('data-bs-toggle', 'dropdown')
        $(this).remove()
    })

    //selecting the event list item
    $(document).on('click', '.event-list-item', function () {
        var eventType = $(this).text().trim()
        $(this).parents('.dropdown').find('a>p#slct-event-item').removeClass('text-[#B2B2B2]').addClass('text-[#152027]').text(eventType)
        $(this).parents('.dropdown').find('input[name=wh_event]').val(eventType.toLowerCase())
        if ($(this).parent().siblings('#event-del-img').length == 0) {
            var dropdownBtn = $(this).parent().siblings('#slct-event-btn')
            $(dropdownBtn).removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
            var delImg = '<img src="/public/img/removeCover.svg" id="event-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
            $(delImg).insertAfter(dropdownBtn)
        }
        $('#wh_event-error').hide()
    })

    // removing event type
    $(document).on('click', '#event-del-img', function () {
        $('input[name=wh_event]').val("")
        $('#slct-event-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctEventType)
        // var dropdownBtn = $(this).parent().siblings('#slct-event-btn')
        $('#slct-event-btn').addClass("bg-[url('/public/img/property-arrow.svg')]").attr('data-bs-toggle', 'dropdown')
        $(this).remove()
    })

    //adding new headers
    $(document).on('click', '#add-header[data-status=0]', function () {
        var parentDiv = $(this).parent('.head-containers').html()
        if ($(this).siblings('input').eq(0).val() != "" && $(this).siblings('input').eq(1).val() != "") {
            $(this).html(`<img src="/public/img/delete.svg">`).attr('data-status', '1')
            $(this).parent().attr('data-status', '1')
            var head = `<div class="grid grid-cols-[1fr_1fr_auto] gap-[16px] head-containers">${parentDiv}</div>`
            $(head).insertBefore($('.head-containers').eq(0))
        }
    })

    // $(document).on('focusout', '.head-containers[data-status=1] input', function (e) {
    //     if ($(this).val() == "" && $(this).siblings('input').val() == "") {
    //         $(this).parent().remove()
    //     }
    // })

    //automatic removing of headers and fields if they are empty valued
    $(document).on('click', function (e) {
        if ($('#modalId').hasClass('show') && ($('.field-containers[data-status=1]').length > 0 || $('.head-containers[data-status=1]').length > 0)) {
            $('.field-containers[data-status=1]').each(function () {
                var fieldNameIp = $(this).find('input[name=wh_field_name]')
                var fieldValueDiv = $(this).find('.dropdown input[name=wh_field_value]')
                var fieldValHost = $(this).find('.dropdown #slct-fieldval-btn #slct-fieldval-item')
                if (!$(e.target).is(fieldNameIp) && !$(e.target).is(fieldValHost)) {
                    var fieldVal = $(fieldNameIp).val()
                    var fieldName = $(fieldValueDiv).val()
                    if (fieldName == "" && fieldVal == "") {
                        $(this).remove()
                    }
                }
            })
            $('.head-containers[data-status=1]').each(function () {
                var keyIp = $(this).find('input[name=wh_header_key]')
                var valueIp = $(this).find('input[name=wh_header_value]')
                if (!$(e.target).is(keyIp) && !$(e.target).is(valueIp)) {
                    var headerKey = $(keyIp).val()
                    var headerValue = $(valueIp).val()
                    if (headerKey == "" && headerValue == "") {
                        $(this).remove()
                    }
                }
            })
        }
    })

    //removing header
    $(document).on('click', '#add-header[data-status=1]', function () {
        $(this).parent().remove()
    })

    //setting webhook status value
    $(document).on('click', '#toggleTwo', function () {
        if ($(this).prop('checked')) {
            $(this).val('1')
        } else {
            $(this).val('0')
        }
    })

    // submitting the form by validating to create and update a webhook
    $(document).on('click', '#webhook-trgrBtn', function (e) {
        var formData = new FormData($('#webhook-form')[0])
        data = {}
        formData.forEach((value, key) => {
            if (key != "wh_header_key" && key != "wh_header_value" && key != "wh_field_name" && key != "wh_field_value") {
                console.log("key:value", key, value);
                data[key] = value
            }
        })

        if (data["wh-status"] == null) {
            var status = $('#toggleTwo').val()
            var param = $('#toggleTwo').attr('name')
            data[param] = status
            console.log("null assign", status);
        }

        if ($('.head-containers[data-status=1]').length > 0) {
            console.log("entr", $('.head-containers[data-status=1]').length > 0);
            data["wh_headers"] = []
            $('.head-containers[data-status=1]').each(function (index, element) {
                var headerKey = $(element).find('input').eq(0).val().trim()
                var headerValue = $(element).find('input').eq(1).val().trim()
                var key = $(element).find('input').eq(0).attr('name')
                var value = $(element).find('input').eq(1).attr('name')
                header = {}
                header[key] = headerKey
                header[value] = headerValue
                data["wh_headers"].push(header)
            })
        }

        if ($('.field-containers[data-status=1]').length > 0) {
            console.log("entr", $('.field-containers[data-status=1]').length > 0);
            data["wh_fields"] = []
            $('.field-containers[data-status=1]').each(function (index, element) {
                var fieldName = $(element).find('input').eq(0).val().trim()
                var fieldValue = $(element).find('.dropdown').find('input[name=wh_field_value]').val().trim()
                var key = $(element).find('input').eq(0).attr('name')
                var value = $(element).find('.dropdown>input[name=wh_field_value]').attr('name')
                field = {}
                field[key] = fieldName
                field[value] = fieldValue
                data["wh_fields"].push(field)
            })
        }

        if ($('#webhook-form').valid()) {
            if ($(this).attr('data-status') == '0') {
                console.log(" webhook create triggered");
                UpdateWebhook(data, '/admin/webhooks/create')
            } else if ($(this).attr('data-status') == '1') {
                console.log(" webhook update triggered");
                UpdateWebhook(data, '/admin/webhooks/update')
            }
        }
    })

    // api call
    function UpdateWebhook(data, url) {
        console.log("data", data);
        $.ajax({
            url: url,
            method: 'POST',
            data: { "webhook": JSON.stringify(data), "csrf": $('#csrf-value').val() },
            success: function (result) {
                console.log("result", result);
                if (result.status == 1) {
                    var currentPage = $('#wh-current-page').val()
                    if (currentPage > 1) {
                        url = "/admin/webhooks?page=" + currentPage
                    } else {
                        url = "/admin/webhooks"
                    }
                    window.location.href = url
                }
            },
            error: function (jqXHR, textStatus, errorThrown) {
                var errorRespoonse = JSON.parse(jqXHR.responseText)
                console.log('Error:', errorRespoonse.status, errorRespoonse, textStatus, errorThrown);
            }
        })
    }

    // setting up webhook form fields in form by calling fetch api
    $(document).on('click', '#wh-edit', async function () {
        var webhookId = $(this).attr('data-id')
        var webhookDetails = await FetchWebhookDetails(webhookId);
        console.log("webhookdetails", webhookDetails);
        if (webhookDetails != null) {
            console.log("webhookdetails", webhookDetails);
            $('input[name=wh_id]').val(webhookId)
            if (webhookDetails.PayloadType != "") {
                $('input[name=wh_payload]').val(webhookDetails.PayloadType)
                var payload = CapitalizeFirstLetter(webhookDetails.PayloadType, "")
                $('#slct-payload-item').removeClass('text-[#B2B2B2]').addClass('text-[#152027]').text(payload)
                $('#payload-container').show()
                $('#slct-payload-btn').removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
                var delImg = '<img src="/public/img/removeCover.svg" id="payload-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
                $(delImg).insertAfter('#slct-payload-btn')
            }
            $('input[name=wh_name]').val(webhookDetails.WebhookName)
            if (webhookDetails.RequestMethod != "") {
                $('input[name=wh_method]').val(webhookDetails.RequestMethod)
                $('#slct-method-btn').removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
                var delImg = '<img src="/public/img/removeCover.svg" id="method-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
                $(delImg).insertAfter('#slct-method-btn')
                var method = CapitalizeFirstLetter(webhookDetails.RequestMethod, "")
                $('#slct-method-item').text(method).removeClass('text-[#B2B2B2]').addClass('text-[#152027]')
            }
            $('input[name=wh_url]').val(webhookDetails.RequestUrl)
            if (webhookDetails.EventType != "") {
                $('input[name=wh_event]').val(webhookDetails.EventType)
                var event = CapitalizeFirstLetter(webhookDetails.EventType, ".")
                $('#slct-event-item').text(event).removeClass('text-[#B2B2B2]').addClass('text-[#152027]')
                $('#slct-event-btn').removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
                var delImg = '<img src="/public/img/removeCover.svg" id="event-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
                $(delImg).insertAfter('#slct-event-btn')
            }
            for (let head in webhookDetails.Headers) {
                var headContainer = `<div id="headers-container" class="grid grid-cols-[1fr_1fr_auto] gap-[16px] head-containers" data-status="1">
                                         <input type="text" placeholder="Header" name="wh_header_key" value="${webhookDetails.Headers[head]["wh_header_key"]}"
                                         class="rounded-[4px] border border-[#EDEDED] bg-[#F7F7F5] p-[12px] h-[36px]  text-[#152027] text-sm font-normal w-full placeholder:text-[#B2B2B2] ">
                                         <input type="text" placeholder="Value" name="wh_header_value" value="${webhookDetails.Headers[head]["wh_header_value"]}"
                                         class="rounded-[4px] border border-[#EDEDED] bg-[#F7F7F5] p-[12px] h-[36px] text-[#152027] text-sm font-normal w-full placeholder:text-[#B2B2B2] ">
                                         <a id="add-header" href="javascript:void(0)" class="bg-[#F7F7F5]  border border-[#EDEDED] text-xl w-[36px] h-[36px] grid place-items-center" data-status="1"><img src="/public/img/delete.svg"></a>
                                         </div>`
                $('#header-section').append(headContainer)
            }
            for (let field in webhookDetails.Fields) {
                var fieldVal = CapitalizeFirstLetter(webhookDetails.Fields[field]["wh_field_value"], " ")
                var delImg = '<img src="/public/img/removeCover.svg" id="fieldval-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
                var fieldContainer = `<div id="field-container" class="grid grid-cols-[1fr_1fr_auto] gap-[16px] field-containers" data-status="1">
                                    <input type="text" placeholder="Field Name" name="wh_field_name" value="${webhookDetails.Fields[field]["wh_field_name"]}"
                                        class="rounded-[4px] border border-[#EDEDED] bg-[#F7F7F5] p-[12px] h-[36px]  text-[#152027] text-sm font-normal w-full placeholder:text-[#B2B2B2] ">
                                    <div class="dropdown">
                                        <input type="hidden" name="wh_field_value" value="${webhookDetails.Fields[field]["wh_field_value"]}">
                                        <a href="javascript:void(0);" id="slct-fieldval-btn"
                                            class="border border-[#EDEDED] pd-3 h-[36px] rounded-[4px]  flex items-center bg-no-repeat bg-[right_12px_center]"
                                            type="button" aria-expanded="false">
                                            <p id="slct-fieldval-item" class="font-normal w-full text-sm text-[#152027]">
                                                ${fieldVal}</p>
                                        </a>
                                        ${delImg}
                                        <ul class="dropdown-menu w-full border-0 rounded-[4px] 
                                            shadow-[0px_8px_24px_-4px_#0000001F] p-[4px_0] !mt-[4px] ">
                                            <li class="fv-list-item"> <a href="javascript:void(0);" data-value="entry title"
                                                    class=" block p-[8px_16px] text-xs font-normal leading-4 text-[#262626] hover:bg-[#F5F5F5] ">
                                                    `+ languageData.webhooks.dropdowns.entryTitle + `</a>
                                            </li>
                                            <li class="fv-list-item"> <a href="javascript:void(0);"
                                                    data-value="channel name"
                                                    class=" block p-[8px_16px] text-xs font-normal leading-4 text-[#262626] hover:bg-[#F5F5F5] ">
                                                    `+ languageData.webhooks.dropdowns.channelName + `</a>
                                            </li>
    
                                            <li class="fv-list-item"> <a href="javascript:void(0);"
                                                    data-value="initiator name"
                                                    class="block p-[8px_16px] text-xs font-normal leading-4 text-[#262626] hover:bg-[#F5F5F5] ">
                                                    `+ languageData.webhooks.dropdowns.eventInitatorName + `</a>
                                            </li>
    
                                            <li class="fv-list-item"> <a href="javascript:void(0);"
                                                    data-value="initiator role"
                                                    class="block p-[8px_16px] text-xs font-normal leading-4 text-[#262626] hover:bg-[#F5F5F5] ">
                                                    `+ languageData.webhooks.dropdowns.eventInitiatorRole + `</a>
                                            </li>
                                        </ul>
                                    </div>
                                    <a id="add-field" href="javascript:void(0)"
                                        class="bg-[#F7F7F5]  border border-[#EDEDED] text-xl w-[36px] h-[36px] grid place-items-center"
                                        data-status="1"><img src="/public/img/delete.svg"></a>
                                </div>`

                $('#field-section').append(fieldContainer)
            }
            if (webhookDetails.IsActive == 1) {
                $('#toggleTwo').prop('checked', true).val(webhookDetails.IsActive)
            }
            $('#webhook-trgrBtn').text(languageData.webhooks.update).attr('data-status', '1')
            $('#modalTitleId').text(languageData.webhooks.updateWebhook)
            $('#modalId').modal('show')
        }
    })

    //fetching webhook details by calling fetch api
    async function FetchWebhookDetails(webhookId) {
        var webhookDetails;
        await $.ajax({
            url: '/admin/webhooks/fetch/' + webhookId,
            dataType: 'json'
        }).then((result) => {
            if (result.status == 1) {
                webhookDetails = result.data
            }
        }).fail((jqXHR, textStatus, errorThrown) => {
            var errorRespoonse = JSON.parse(jqXHR.responseText)
            console.log('Error:', errorRespoonse.status, errorRespoonse, textStatus, errorThrown);
            webhookDetails = null
        })
        return webhookDetails
    }

    //capitalizing text function
    function CapitalizeFirstLetter(content, separator) {
        if (content.length == 0) {
            return ""
        }
        var modContent = ""
        if (separator != "") {
            var splitData = content.split(separator)
            for (let x in splitData) {
                if (x == splitData.length - 1) {
                    modContent = modContent + splitData[x].charAt(0).toUpperCase() + splitData[x].slice(1)
                } else {
                    modContent = modContent + splitData[x].charAt(0).toUpperCase() + splitData[x].slice(1) + separator
                }
            }
        } else {
            modContent = modContent + content.toUpperCase()
        }
        return modContent
    }

    //handling the ui screen while the form modal start hiding
    $('#modalId').on('hide.bs.modal', function () {
        $('input[name=wh_id]').val('')
        $('#toggleTwo').prop('checked', false).val('0')
        $('input[name=wh_payload]').val("")
        // $('#payload-container').hide()
        $('#slct-payload-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctPaylodType)
        $('input[name=wh_name]').val("")
        $('#slct-payload-btn').addClass("bg-[url('/public/img/property-arrow.svg')]")
        $('#payload-del-img').remove()
        $('input[name=wh_method]').val("")
        $('#slct-method-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctMethod)
        $('#slct-method-btn').addClass("bg-[url('/public/img/property-arrow.svg')]")
        $('#method-del-img').remove()
        $('input[name=wh_url]').val("")
        $('input[name=wh_event]').val("")
        $('#event-del-img').remove()
        $('#slct-event-btn').addClass("bg-[url('/public/img/property-arrow.svg')]")
        $('#slct-event-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctEventType)
        $('#toggleTwo').prop('checked', false).val("0")
        $('#webhook-trgrBtn').text('Create').attr('data-status', '0')
        $('#modalTitleId').text(languageData.webhooks.createWebhook)
        $('.head-containers[data-status=1]').remove()
        $('.field-containers[data-status=1]').remove()
        $('.head-containers').each(function (index, element) {
            $(this).find('input').val("")
        })
        $('label.error').hide()
    })

    //triggering webhook delete api in ui modal
    $(document).on('click', '#wh-delete', function () {
        var currentPage = $('#wh-current-page').val()
        var webhookId = $(this).attr("data-id")
        $('#deleteModal .deltitle').text(languageData.webhooks.deleteWebhook)
        $('#deleteModal #content').text(languageData.webhooks.sureDelete)
        var webhookName = $(this).parents('#wh-container' + webhookId).find('.items-center>#wh-flxcol1>#wh-flxcol2>h5').text()
        $('#deleteModal .delname').text(webhookName)
        $('#deleteModal #delid').attr('href', "/admin/webhooks/delete/" + webhookId + "?page=" + currentPage)
        $('#deleteModal').modal('show')
    })

    //setting up the value of selected  payload field value to the its appropriate form field 
    $(document).on('click', '.fv-list-item', function () {
        fieldDiv = $(this).parents('.field-containers').html()
        var fieldVal = $(this).text().trim()
        console.log("data", fieldVal);
        $(this).parents('.dropdown').find('a>p#slct-fieldval-item').removeClass('text-[#B2B2B2]').addClass('text-[#152027]').text(fieldVal)
        $(this).parents('.dropdown').find('input[name=wh_field_value]').val(fieldVal.toLowerCase())
        if ($(this).parent().siblings('#fieldval-del-img').length == 0) {
            var dropdownBtn = $(this).parent().siblings('#slct-fieldval-btn')
            $(dropdownBtn).removeClass("bg-[url('/public/img/property-arrow.svg')]").removeAttr('data-bs-toggle')
            var delImg = '<img src="/public/img/removeCover.svg" id="fieldval-del-img" style="position: absolute;top: 10px;right: 10px;cursor: pointer;">'
            $(delImg).insertAfter(dropdownBtn)
        }
        $('#wh_field_value-error').hide()
    })

    //removing of payload field value
    $(document).on('click', '#fieldval-del-img', function () {
        $(this).siblings('input[name=wh_field_value]').val("")
        $(this).siblings('#slct-fieldval-btn').find('#slct-fieldval-item').removeClass('text-[#152027]').addClass('text-[#B2B2B2]').text(languageData.webhooks.placeholders.slctFieldValue)
        // var dropdownBtn = $(this).parent().siblings('#slct-fieldval-btn')
        $(this).siblings('#slct-fieldval-btn').addClass("bg-[url('/public/img/property-arrow.svg')]").attr('data-bs-toggle', 'dropdown')
        $(this).remove()
    })

    //Adding payload fields to the form
    $(document).on('click', '#add-field[data-status=0]', function () {
        if ($(this).siblings('div').eq(0).find('input[name=wh_field_name]').val() != "" && $(this).siblings('.dropdown').find('input[name=wh_field_value]').val() != "" && $(this).parent().attr('data-status') != "1") {
            $(this).html(`<img src="/public/img/delete.svg">`).attr('data-status', '1')
            $(this).parent().attr('data-status', '1')
            var head = `<div class="grid grid-cols-[1fr_1fr_auto] gap-[16px] field-containers">${fieldDiv}</div>`
            $(head).insertBefore($('.field-containers').eq(0))
            $('.field-containers[data-status=0]').find('.dropdown').find('a>p#slct-fieldval-item').text(languageData.webhooks.placeholders.slctFieldValue)
            $('#wh_field_name-error').hide()
            $('#wh_field_value-error').hide()
        }
    })

    //adding payload fields to the form field
    $(document).on('click', '#add-field[data-status=1]', function () {
        $(this).parent().remove()
    })

    //jquery form validation
    $('#webhook-form').validate({
        ignore: [],
        rules: {
            wh_name: {
                required: true,
                minlength: 3, // Minimum length
                maxlength: 20, // Maximum length
                validate_whname: true
            },
            wh_url: {
                required: true,
                validate_whurl: true,
            },
            wh_method: {
                required: true
            },
            wh_event: {
                required: true
            },
            wh_field_name: {
                required: function () {
                    return $('input[name=wh_payload]').val() != "" && $('.field-containers[data-status=1]').length == 0
                }
            },
            wh_field_value: {
                required: function () {
                    return $('input[name=wh_payload]').val() != "" && $('.field-containers[data-status=1]').length == 0
                },
            }
        },
        messages: {
            wh_name: {
                required: "* " + languageData.webhooks.validationErrors.reqWebhookName,
                minlength: "* " + languageData.webhooks.validationErrors.whnameMinLength,
                maxlength: "* " + languageData.webhooks.validationErrors.whnameMaxLength,
            },
            wh_url: {
                required: "* " + languageData.webhooks.validationErrors.reqUrl
            },
            wh_method: {
                required: "* " + languageData.webhooks.validationErrors.reqMethod
            },
            wh_event: {
                required: "* " + languageData.webhooks.validationErrors.reqEventType
            },
            wh_field_name: {
                required: "* " + languageData.webhooks.validationErrors.reqFieldName
            },
            wh_field_value: {
                required: "* " + languageData.webhooks.validationErrors.reqFieldValue
            },
        },
    });

    //jquery custom validation
    jQuery.validator.addMethod(
        "validate_whname",
        function (value, element) {
            if (/^(?! )[a-zA-Z0-9_-]+(?: [a-zA-Z0-9_-]+)*[^ ]$/.test(value)) {
                return true;
            } else {
                return false;
            }
        },
        "* " + languageData.webhooks.validationErrors.validateWebhookName
    );

    jQuery.validator.addMethod(
        "validate_whurl",
        function (value, element) {
            if (/^(https?):\/\/([a-zA-Z0-9-._~%]+(?:\.[a-zA-Z0-9-._~%]+)*|\[[0-9a-fA-F:.]+\])(:\d{1,5})?(\/[a-zA-Z0-9-._~%!$&'()*+,;=:@/]*)?(\?[a-zA-Z0-9-._~%!$&'()*+,;=:@/?]*)?(#[a-zA-Z0-9-._~%!$&'()*+,;=:@/?]*)?$/.test(value)) {
                return true;
            } else {
                return false;
            }
        },
        "* " + languageData.webhooks.validationErrors.validateWebhookUrl
    );
})

$(document).on('keyup', '#wh-srch-input', function (e) {
    var searchVal = $('#wh-srch-input').attr('value')
    if (e.which == 8 && searchVal != "") {
        var searchKey = $('#wh-srch-input').val()
        if (searchKey == "") {
            window.location.href = "/admin/webhooks/"
        }
    }
})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
  })

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')

    window.location.href = "/admin/webhooks/"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            var value=$(".search").val();
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
  