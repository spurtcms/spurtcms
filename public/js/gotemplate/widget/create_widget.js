var languagedata;
var selectedLabel;

// ==================
// LOAD LANGUAGE DATA
// ==================
$(document).ready(async function () {
    var languagepath = $('.language-group>button').attr('data-path');

    if (languagepath) {
        try {
            languagedata = await $.getJSON(languagepath);
        } catch (e) {
            console.log("Language load failed", e);
        }
    }
});

// ==================
// TAB SWITCH LOGIC
// ==================
$(document).ready(function () {

    $('input[name="lang"]').change(function () {

        selectedLabel = $(this).next('label').text().trim();

        // reset errors and clear previous mapping validation state
        $('.navpath').val("");
        $('.product-error, .entries-error, .pages-error, .listings-error, .category-error, .channel-error').addClass('hidden');
        $('.title-data-error').addClass('hidden');

        // hide all first
        $('.customUrlInput, .entriesDropdown, .pagesDropdown, .listingsDropdown, .categoryDropdown, .channelDropdown').addClass('hidden');

        switch (selectedLabel) {
            case "Custom URL":
                $('.customUrlInput').removeClass('hidden');
                $('.menutype').val("custom_url");
                break;

            case "Entries":
                $('.entriesDropdown').removeClass('hidden');
                $('.widget_type').val('Entries');
                break;

            case "Pages":
                $('.pagesDropdown').removeClass('hidden');
                $('.widget_type').val('Pages');
                break;

            case "Listings":
                $('.listingsDropdown').removeClass('hidden');
                $('.widget_type').val('Listings');
                break;

            case "Categories":
                $('.categoryDropdown').removeClass('hidden');
                $('.widget_type').val('Categories');
                break;

            case "Channels":
                $('.channelDropdown').removeClass('hidden');
                $('.widget_type').val('Channels');
                break;

            case "None":
                $('.widget_type').val('');
                $('.menutype').val("none");
                break;
        }
    });
});

// ==================
// DROPDOWN STOP PROPAGATION
// ==================
$(document).ready(function () {
    $('.listingsDropdownMenu, .categoryDropdownMenu, .entriesDropdown, .pagesDropdown, .channelDropdownMenu')
        .on('click', function (e) {
            e.stopPropagation();
        });
});

// ==================
// FORM VALIDATION (INIT ONCE)
// ==================
$(document).ready(function () {

    $("#widgetform").validate({
        ignore: [],
        rules: {
            title: {
                required: true
            },
            position: {
                required: true
            }
        },
        messages: {
            title: {
                required: "* Please Enter Widget Title"
            },
            position: {
                required: "* Please Enter Widget Position"
            }
        },

errorPlacement: function (error, element) {

    if (element.attr("name") === "title") {
        $('.title-data-error')
            .removeClass('hidden')
            .text(error.text());

    } else if (element.attr("name") === "position") {
        $('.position-data-error')
            .removeClass('hidden')
            .text(error.text());

    } else {
        error.insertAfter(element); // ✅ only once
    }
},

success: function (label, element) {

    if ($(element).attr("name") === "title") {
        $('.title-data-error').addClass('hidden');
    }

    if ($(element).attr("name") === "position") {
        $('.position-data-error').addClass('hidden');
    }
}

// position

    });
});

// ==================
// SAVE BUTTON
// ==================
$(document).on('click', '#widgetsave', function () {

    GetProductIds();

    var widgetTypeVal = $('.widget_type').val();
    var productIdsVal = $('.productids').val();

    // clear all mapping errors
    $('.product-error, .entries-error, .pages-error, .listings-error, .category-error, .channel-error').addClass('hidden');

    var mappingValid = true;

    if (widgetTypeVal == "") {
        mappingValid = false;
        $('.product-error').removeClass('hidden');
    } else if (!productIdsVal) {
        switch (widgetTypeVal) {
            case "Entries":
                $('.entries-error').removeClass('hidden');
                mappingValid = false;
                break;
            case "Pages":
                $('.pages-error').removeClass('hidden');
                mappingValid = false;
                break;
            case "Listings":
                $('.listings-error').removeClass('hidden');
                mappingValid = false;
                break;
            case "Categories":
                $('.category-error').removeClass('hidden');
                mappingValid = false;
                break;
            case "Channels":
                $('.channel-error').removeClass('hidden');
                mappingValid = false;
                break;
        }
    }

    var formcheck = $("#widgetform").valid();

    if (formcheck && mappingValid) {
        $('#widgetform')[0].submit();
    }

    return false;
});

// ==================
// DYNAMIC ERROR CLEARING
// ==================
$(document).ready(function () {

    // Clear title error on input
    $('input[name="title"]').on('keyup', function () {
        $('.title-data-error').addClass('hidden');
    });

    // Clear mapping errors when selecting a radio button
    $('input[name="lang"]').on('change', function () {
        $('.product-error, .entries-error, .pages-error, .listings-error, .category-error, .channel-error').addClass('hidden');
    });

    // Clear mapping errors when checkboxes are checked
    $(document).on('change', '.selected-listings, .selected-pages, .selected-chennals, .selectcheckbox', function () {
        var widgetType = $('.widget_type').val();
        if (widgetType === 'Entries') {
            $('.entries-error').addClass('hidden');
        } else if (widgetType === 'Listings') {
            $('.listings-error').addClass('hidden');
        } else if (widgetType === 'Pages') {
            $('.pages-error').addClass('hidden');
        } else if (widgetType === 'Categories') {
            $('.category-error').addClass('hidden');
        } else if (widgetType === 'Channels') {
            $('.channel-error').addClass('hidden');
        }
    });
});

// ==================
// GET PRODUCT IDS
// ==================
function GetProductIds() {

    var productids = '';
    var widgetType = $('.widget_type').val();

    if (widgetType === "Listings") {
        productids = $('.selected-listings:checked').map(function () {
            return $(this).data('id');
        }).get().join(',');
    }

    else if (widgetType === "Entries") {
        productids = $('.entriesDropdown .selected-listings:checked').map(function () {
            return $(this).data('id');
        }).get().join(',');
    }

    else if (widgetType === "Categories") {
        productids = $('.categoryDropdown .selectcheckbox:checked').map(function () {
            return $(this).closest('li').find('label span').last().attr('data-id');
        }).get().join(',');
    }

    else if (widgetType === "Channels") {
        productids = $('.channelDropdown .selected-chennals:checked').map(function () {
            return $(this).data('id');
        }).get().join(',');
    }

    else if (widgetType === "Pages") {
        productids = $('.pagesDropdown .selected-pages:checked').map(function () {
            return $(this).data('id');
        }).get().join(',');
    }

    $('.productids').val(productids);
}

// ==================
// STATUS DROPDOWN
// ==================
$(document).on('click', '.statusdropdown', function () {

    $(this).closest('ul').removeClass('show');

    var statusval = $(this).text().trim();
    $('.status_type_span').text(statusval);

    $('#status').val(statusval === "Active" ? 1 : 0);
});

// ==================
// POSITION DROPDOWN
// ==================
$(document).on('click', '.postiondropdown', function () {

    $(this).closest('ul').removeClass('show');

    var postionval = $(this).text().trim();

    $('.position_span').text(postionval);
    $('#position').val(postionval);

    $('.position-data-error').addClass('hidden');
});

// ==================
// SORT DROPDOWN
// ==================
$(document).on('click', '.sortdropdown', function () {

    $(this).closest('ul').removeClass('show');

    var sortorder = $(this).text().trim();

    $('.sortorder_span').text(sortorder);
    $('#sort_order').val(sortorder);
});


// ============================================================================================================


// ── Show/hide widget detail panel ────────────────────────────────────────────



function showWidgetDetail(clickedTd, widgetId) {
    document.getElementById('widget-table-wrapper').style.display = 'none';
    document.getElementById('widget-detail-section').style.display = 'block';
    var pagination = document.getElementById('widget-pagination');
    if (pagination) pagination.style.display = 'none';

   function hideAllButtons() {
        const buttons = document.querySelectorAll('.action-btn');
        buttons.forEach(function(btn) {
            btn.style.display = 'none';
        });
    }

   
    hideAllButtons();

    if (clickedTd) {
        document.querySelectorAll('#pills-widgets tbody td:first-child').forEach(function(td) {
            td.style.fontWeight = '';
            td.style.color = '';

        });

        clickedTd.style.fontWeight = '600';
        clickedTd.style.display = 'none';

    }

    if (widgetId) {
        loadWidgetForEdit(widgetId);
    } else {
        resetWidgetForm();
    }

    window.scrollTo({ top: 0, behavior: 'smooth' });

}

function showWidgetTable() {
    document.getElementById('widget-detail-section').style.display = 'none';
    document.getElementById('widget-table-wrapper').style.display = 'block';
    var pagination = document.getElementById('widget-pagination');
    if (pagination) pagination.style.display = '';
    document.querySelectorAll('.action-btn').forEach(function(btn) {
        btn.style.display = '';
    });
    document.querySelectorAll('#pills-widgets tbody td:first-child').forEach(function(td) {
        td.style.fontWeight = '';
        td.style.color = '';
    });
}

// ── Helper: hide all mapping dropdowns and uncheck all mapping checkboxes ────
function resetMappingDropdowns() {
    document.querySelectorAll('.entriesDropdown, .pagesDropdown, .listingsDropdown, .categoryDropdown, .channelDropdown').forEach(function(el) {
        el.classList.add('hidden');
        el.style.display = '';
    });
    document.querySelectorAll('.selected-listings, .selected-pages, .selected-chennals, .selectcheckbox').forEach(function(chk) {
        chk.checked = false;
    });
    
    // Uncheck all radio buttons
    document.querySelectorAll('input[name="lang"]').forEach(function(r) {
        r.checked = false;
    });
}

// ── Helper: restore mapping selection from saved WidgetType + productIds ─────
function restoreMappingSelection(widgetType, productIdsStr) {
    resetMappingDropdowns();

    if (!widgetType) return;

    var savedIds = (productIdsStr || '').split(',').map(function(s) { return s.trim(); }).filter(Boolean);
    var type = widgetType.toLowerCase();

    // Map widgetType value to the radio input value used in the template
    var radioValueMap = {
        'entries':    'entries',
        'pages':      'pages',
        'listings':   'listings',
        'categories': 'categories',
        'channels':   'Channels'
    };

    // Map widgetType to the dropdown container class
    var dropdownClassMap = {
        'entries':    '.entriesDropdown',
        'pages':      '.pagesDropdown',
        'listings':   '.listingsDropdown',
        'categories': '.categoryDropdown',
        'channels':   '.channelDropdown'
    };

    var radioVal = radioValueMap[type];
    var dropClass = dropdownClassMap[type];

    // Check the correct radio button
    if (radioVal) {
        var radio = document.querySelector('input[name="lang"][value="' + radioVal + '"]');
        if (radio) radio.checked = true;
    }

    // Show the correct dropdown
    if (dropClass) {
        var dropdown = document.querySelector(dropClass);
        if (dropdown) {
            dropdown.classList.remove('hidden');
        }
    }

    // Re-check saved item checkboxes based on type
    if (savedIds.length === 0) return;

    if (type === 'entries' || type === 'listings') {
        // entries: id="Check{id}" class="selected-listings"
        // listings: id="Checklist{id}" class="selected-listings"
        savedIds.forEach(function(id) {
            var prefix = (type === 'entries') ? 'Check' : 'Checklist';
            var chk = document.getElementById(prefix + id);
            if (chk) chk.checked = true;
        });
    } else if (type === 'pages') {
        // pages: id="PageCheck{id}" class="selected-pages"
        savedIds.forEach(function(id) {
            var chk = document.getElementById('PageCheck' + id);
            if (chk) chk.checked = true;
        });
    } else if (type === 'categories') {
        // categories: id="Checkcat{index}" class="selectcheckbox"
        savedIds.forEach(function(id) {
            var chk = document.getElementById('Checkcat' + id);
            if (chk) chk.checked = true;
        });
    } else if (type === 'channels') {
        // channels: id="PageCheck{index}" class="selected-chennals"
        savedIds.forEach(function(id) {
            var chk = document.getElementById('PageCheck' + id);
            if (chk) chk.checked = true;
        });
    }
}

// ── Load widget data into form for editing ───────────────────────────────────
function loadWidgetForEdit(widgetId) {
    var webid  = document.getElementById('webid')      ? document.getElementById('webid').value      : '';
    var tempid = document.getElementById('templateId') ? document.getElementById('templateId').value : '';

    $.ajax({
        url: '/admin/website/widgets/editwidget/' + widgetId + '?webid=' + webid + '&tempid=' + tempid,
        method: 'GET',
        headers: { 'X-Requested-With': 'XMLHttpRequest' },
        success: function(res) {
            var detail = res.widgetdetail;
            if (!detail) return;

            var form = document.getElementById('widgetform');

            // ── Form action & widget ID ──────────────────────────────────
            form.action = '/admin/website/widgets/updatewidget';
            form.querySelector('[name="widget_id"]').value = detail.Id || '';

            // ── Text fields ──────────────────────────────────────────────
            form.querySelector('[name="title"]').value            = detail.Title           || '';
            form.querySelector('[name="long_title"]').value       = detail.LongTitle       || '';
            form.querySelector('[name="widget_slug"]').value      = detail.Slug            || '';
            form.querySelector('[name="widget_limit"]').value     = detail.WidgetLimit     || '0';
            form.querySelector('[name="meta_title"]').value       = detail.MetaTitle       || '';
            form.querySelector('[name="meta_description"]').value = detail.MetaDescription || '';
            form.querySelector('[name="meta_keyword"]').value     = detail.MetaKeywords    || '';

            // ── Position: update hidden input AND visible dropdown label ─
            var positionVal = detail.Position || '';
            document.querySelector('.position_span').textContent = positionVal || 'Position';
            document.getElementById('position').value = positionVal;

            // ── Sort Order: update hidden input AND visible dropdown label ─
            // FIX: sortorder_span was never updated — dropdown label stayed stale
            var sortVal = detail.SortOrder ? String(detail.SortOrder) : '';
            document.querySelector('.sortorder_span').textContent = sortVal || 'Sort Order';
            document.getElementById('sort_order').value = sortVal;

            // ── Status: update hidden input AND visible dropdown label ───
            // FIX: status_type_span was never updated — dropdown label stayed blank
            var statusVal   = detail.Status;
            var statusLabel = (statusVal == 1 || statusVal === 'Active') ? 'Active' : 'InActive';
            var statusNum   = (statusVal == 1 || statusVal === 'Active') ? '1' : '0';
            form.querySelector('[name="status"]').value = statusNum;
            document.querySelector('.status_type_span').textContent = statusLabel;

            // ── ProductIds ───────────────────────────────────────────────
            // FIX: res.widgetdetail.ProductIds does not exist —
            // the Go handler returns product IDs separately as res.productids
            form.querySelector('.productids').value  = res.productids || '';
            form.querySelector('.widget_type').value = detail.WidgetType || '';

            // ── TemplateId ───────────────────────────────────────────────
            // FIX: keep templaId in sync so save/update submits correct template
            var templaIdField = form.querySelector('[name="templaId"]');
            if (templaIdField) {
                templaIdField.value = res.templateid || tempid || '';
            }

            // ── Mapping Product/Category ─────────────────────────────────
            // FIX: restore radio selection + show correct dropdown + re-check saved IDs
            restoreMappingSelection(detail.WidgetType, res.productids);
        },
        error: function() {
            console.error('Failed to load widget details for id:', widgetId);
        }
    });
}

// ── Reset form for create mode ───────────────────────────────────────────────
function resetWidgetForm() {
    var form = document.getElementById('widgetform');
    form.action = '/admin/website/widgets/savewidget';

    form.querySelector('[name="widget_id"]').value        = '';
    form.querySelector('[name="title"]').value            = '';
    form.querySelector('[name="long_title"]').value       = '';
    form.querySelector('[name="widget_slug"]').value      = '';
    form.querySelector('[name="widget_limit"]').value     = '0';
    form.querySelector('[name="status"]').value           = '';
    form.querySelector('[name="meta_title"]').value       = '';
    form.querySelector('[name="meta_description"]').value = '';
    form.querySelector('[name="meta_keyword"]').value     = '';
    form.querySelector('.productids').value               = '';
    form.querySelector('.widget_type').value              = '';

    document.querySelector('.position_span').textContent    = 'Position';
    document.querySelector('.sortorder_span').textContent   = 'Sort Order';
    document.querySelector('.status_type_span').textContent = 'status';
    document.getElementById('position').value               = '';
    document.getElementById('sort_order').value             = '';

    // Reset all mapping dropdowns and checkboxes
    resetMappingDropdowns();

    // Clear all error messages
    $('.title-data-error, .position-data-error, .product-error, .entries-error, .pages-error, .listings-error, .category-error, .channel-error').addClass('hidden');
}

$(document).ready(function () {

    // ── Widgets tab: return to table view ───────────────────────────────
    $("#pills-widgets-tab").on("click", function () {
        showWidgetTable();
    });

    // ── Social media save ────────────────────────────────────────────────
    $('#socialmediaupdate').on('click', function (e) {
        e.preventDefault();
        var socialData = {
            linkedin:  $('#linkedinlink').val(),
            x:         $('#xlink').val(),
            youtube:   $('#youtubelink').val(),
            instagram: $('#instagramlink').val(),
            facebook:  $('#facebooklink').val()
        };
        $('#social_media_data').val(JSON.stringify(socialData));
        $.ajax({
            url: '/admin/website/settingpage?webid=1',
            method: 'POST',
            headers: { 'X-Requested-With': 'XMLHttpRequest' },
            data: $('#SocialmediaForm').serialize(),
            success: function (res) {
                var btn = $('#socialmediaupdate');
                var original = btn.text();
                btn.text('Saved ✓').css('background-color', '#148569');
                setTimeout(function () {
                    btn.text(original).css('background-color', '');
                }, 2000);
            },
            error: function () {
                alert('Failed to save social links. Please try again.');
            }
        });
    });

});