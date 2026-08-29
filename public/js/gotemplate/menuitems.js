var languagedata
var selectedLabel
var categoryarr = [];
/** */
$(document).ready(async function () {
    initializeAllChildContainers();

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    $('.Content').addClass('checked');

    var exists = false;
    $('#accordionExample1 .accordion-item').each(function () {
        var $updateBtn = $(this).find('#updatebtn');
        if (
            $updateBtn.data('type') == "courses"

        ) {
            exists = true;
            return false;
        }
    });

    if (exists) {


        $('#collapseTwo').find('.channelnameinput').prop('checked', true)
    }
    var selectedcategory = $('.selectedcategoryids').val();
    var selectedlisting = $('.selectedlistingids').val();




    var categoryIdsArray = selectedcategory
        .replace(/[\[\]]/g, '')
        .trim()
        .split(/\s+/)
        .map(id => parseInt(id.trim(), 10))
        .filter(id => !isNaN(id));


    var uniqueCategoryIds = [...new Set(categoryIdsArray)];



    $('.categorybody li').each(function () {
        var dataId = $(this).find('span').last().attr('data-id');
        var dataIdNum = parseInt(dataId, 10);
        $(this).find('input[type="checkbox"]').prop('checked', uniqueCategoryIds.includes(dataIdNum));
    });

    var listingIdsArray = selectedlisting
        .replace(/[\[\]]/g, '')
        .trim()
        .split(/\s+/)
        .map(id => parseInt(id.trim(), 10))
        .filter(id => !isNaN(id));


    var uniquelistingIds = [...new Set(listingIdsArray)];



    $('.listingbody li').each(function () {
        var dataId = $(this).find('span').last().attr('data-id');
        var dataIdNum = parseInt(dataId, 10);
        $(this).find('input[type="checkbox"]').prop('checked', uniquelistingIds.includes(dataIdNum));
    });
});

// ─────────────────────────────────────────────────────────────
// SAVE MENU ITEM
// ─────────────────────────────────────────────────────────────
$(document).on('click', '#menuitemssavebtn', function (e) {
    e.preventDefault();

    if (!$.validator.methods.duplicatename) {
        $.validator.addMethod("duplicatename", function (value) {
            var result;
            $.ajax({
                url: "/admin/website/menu/checkmenuname",
                type: "POST",
                async: false,
                data: {
                    "menu_name": value,
                    "webid": $(".templateid").val(),
                    "menu_id": $(".menuitem_id").val(),
                    "parentmenu_id": $('.parentmenu_id').val(),
                    csrf: $("input[name='csrf']").val()
                },
                datatype: "json",
                caches: false,
                success: function (data) {
                    result = data.trim();
                }
            });
            return result.trim() !== "true";
        });
    }

    let selectedIds = $('.selected-listings:checked').map(function () { return $(this).data('id'); }).get();
    $('#listingsids').val(selectedIds.join(','));
    if (categoryarr.length > 0) { $('#categoryids').val(categoryarr.join(',')); }
    if (selectedLabel == "Listings")   { $(".navpath").val("/listings/"); }
    if (selectedLabel == "Categories") { $('.navpath').val("/categories/"); }

    $("#menuitemform").validate({
        onkeyup: false, onfocusout: false, onclick: false,
        rules: {
            menu_name: { required: true, space: true, duplicatename: true, maxlength: 300 },
            urlpath: {
                required: function () { return selectedLabel !== "None"; },
                space:    function () { return selectedLabel !== "None"; }
            }
        },
        messages: {
            menu_name: {
                required:      "*" + languagedata.Menu.labelnameerr,
                space:         "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Menu.nameduplicateerr,
                maxlength:     "*" + languagedata.Menu.namelimiterr
            },
            urlpath: {
                required: "*" + languagedata.Menu.linkpatherr,
                space:    "* " + languagedata.spacergx
            }
        }
    });

    if ($("#menuitemform").valid() === true) { $('#menuitemform')[0].submit(); }
    return false;
});

// ─────────────────────────────────────────────────────────────
// ADD CHANNEL / PAGE / CATEGORY MENU ITEMS
// ─────────────────────────────────────────────────────────────
$(document).on('click', '.addchannelmenu', function (e) {
    e.preventDefault();
    var $accordionBody = $(this).closest('.accordion-body');
    $accordionBody.find('.channelnameinput:checked').each(function () {
        var menuname      = $(this).siblings('label').find('p').text().trim();
        var id            = $(this).attr('data-id');
        var type          = $(this).attr('data-type');
        var parentmenuid  = $('#menu_id').val();
        var slug          = $(this).attr('data-slug');
        var menu_grouptype = $('#parentmenu_grouptype').val();

        var exists = false;
        $('#accordionExample1 .accordion-item').each(function () {
            var $updateBtn = $(this).find('#updatebtn');
            if (($updateBtn.data('type') == type && $updateBtn.data('typeid') == id) ||
                $updateBtn.data('name') == menuname) {
                exists = true; return false;
            }
        });

        if (!exists) {
            $.ajax({
                url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').attr('data-id'),
                type: "POST", async: false,
                data: {
                    "menu_name": menuname, "menu_id": parentmenuid, "urlpath": "",
                    csrf: $("input[name='csrf']").val(), "menu_typeid": id, "type": type,
                    "slug_name": slug, "webid": $('.webid').attr('data-id'),
                    parentmenu_grouptype: menu_grouptype
                },
                dataType: "json", cache: false,
                success: function () { location.reload(); }
            });
        }
    });
});

function addCategoryFunction(type, classname) {
    var $accordionBody = $('.' + classname).closest('.accordion-body');
    var menu_grouptype = $('#parentmenu_grouptype').val();

    $accordionBody.find('.channelnameinput:checked').each(function () {
        var $childSpan    = $(this).siblings('label').find('span[data-id]').last();
        var subcategoryid = $childSpan.data('id');
        var subcategoryname = $.trim($childSpan.text());
        var parentmenuid  = $('#menu_id').val();

        var exists = false;
        $('#accordionExample1 .accordion-item').each(function () {
            var $updateBtn = $(this).find('#updatebtn');
            if ($updateBtn.data('type') == type && $updateBtn.data('typeid') == subcategoryid &&
                $updateBtn.data('name') == subcategoryname) {
                exists = true; return false;
            }
        });

        if (!exists) {
            $.ajax({
                url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').data('id'),
                method: "POST", dataType: "json", async: false,
                data: {
                    menu_name: subcategoryname, menu_id: parentmenuid, menu_typeid: subcategoryid,
                    csrf: $("input[name='csrf']").val(), type: type,
                    webid: $('.webid').data('id'), parentmenu_grouptype: menu_grouptype
                },
                success: function () { location.reload(); }
            });
        }
    });
}

$(document).off('click.addcategorymenu').on('click.addcategorymenu', '.addcategorymenu', function (e) {
    e.preventDefault();
    addCategoryFunction("categories", "addcategorymenu");
});
$(document).off('click.addlistingcatmenu').on('click.addlistingcatmenu', '.addlistingcatmenu', function (e) {
    e.preventDefault();
    addCategoryFunction("listings", "addlistingcatmenu");
});

// ─────────────────────────────────────────────────────────────
// UPDATE MENU ITEM (inline save)
// ─────────────────────────────────────────────────────────────
$(document).on('click', '#updatebtn', function () {
    var menuId          = $(this).data('id');
    var type            = $(this).attr('data-type');
    var typeid          = $(this).attr('data-typeid');
    var parentmenuid    = $(this).attr('data-parent');
    var $accordionItem  = $(this).closest('.accordion-body');
    var label           = $accordionItem.find('input[placeholder="Navigation Lable"]').val().trim();
    var path            = $accordionItem.find('input[placeholder="Path link"]').val();
    var svgimage        = $accordionItem.find('.svgval').val();
    var svgDelete       = $accordionItem.find('.svgDelete').val();
    var separatewindow  = $accordionItem.find('.separatewindow').val();
    var menu_Group      = $accordionItem.find('input[placeholder="Navigation Lable"]').attr('data-group');
    var labelValid      = label.length <= 300;
    var isDuplicate     = false;

    var $editor      = $accordionItem.find('#editor-' + menuId);
    var quill        = $editor.data('quill');
    var editorContent = quill ? quill.root.innerHTML.replace(/<p><br><\/p>$/, '').trim() : '';

    $.ajax({
        url: "/admin/website/menu/checkmenuname",
        type: "POST", async: false,
        data: {
            "menu_name": label, "webid": $(".templateid").val(),
            "menu_id": menuId, csrf: $("input[name='csrf']").val()
        },
        datatype: "json", cache: false,
        success: function (data) {
            var result = data.trim();

            if (result == 'true') {
                isDuplicate = true;
            }
        }
    });


    if (!labelValid) {



        $accordionItem.find('.lablename-error').removeClass('hidden').addClass('mb-[24px]').text('Please Enter 300 Characters Only');
    } else {
        $accordionItem.find('.lablename-error').addClass('hidden').text('');
    }
    if (isDuplicate) {

        $accordionItem.find('.labelname').removeClass('mb-[24px]')
        $accordionItem.find('.lablename-error').removeClass('hidden').addClass('mb-[24px]').text('Menu name already exists!');
    }


    if (labelValid && !isDuplicate) {
        $.ajax({
            url: "/admin/website/menu/updatemenuitems",
            type: "POST",
            async: false,
            data: {
                "menuitem_id": menuId,
                "menu_name": label,
                "webid": $('.templateid').attr('data-id'),
                "parentmenu_id": parentmenuid,
                "urlpath": path,
                "svgHidden": svgimage,
                "svgDelete": svgDelete,
                csrf: $("input[name='csrf']").val(),
                "menu_typeid": typeid,
                "type": type,
                "metainfo": "false",
                "separatewindow": separatewindow,
                "editorContent": editorContent,
                "menu_group": menu_Group,

            },
            datatype: "json",
            cache: false,
            success: function (data) {
                result = data

                if (result) {

                    window.location.reload();
                    $accordionItem.closest('.accordion-item').find('.itemname').text(result.Name);
                    notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Menu Updated Successfully </p ></div ></div ></li></ul> `;
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();

                        });
                    }, 5000);
                }
            }
        });
    }
});

// ─────────────────────────────────────────────────────────────
// DELETE MENU ITEM
// ─────────────────────────────────────────────────────────────
$(document).on('click', '#menuitemdeletebtn', function () {
    var menuid = $(this).attr("data-id");
    var url    = window.location.search;
    var pageno = new URLSearchParams(url).get('page');
    
    // FIX: Get current view ID from URL path (e.g., /menuitems/394 → 394)
    var pathParts = window.location.pathname.split('/');
    var currentViewId = pathParts[pathParts.length - 1];

    // Build delete URL with redirectid
    var deleteUrl = "/admin/website/menu/deletemenuitem/" + menuid 
        + "?webid=" + $('.templateid').attr('data-id')
        + "&redirectid=" + currentViewId;

    if (pageno != null) {
        deleteUrl += "&page=" + pageno;
    }

    $('#delid').attr('href', deleteUrl);
    $(".deltitle").text(languagedata.Menu.deletemenu + "?");
    $('.delname').text($(this).attr("data-name"));
});

// ─────────────────────────────────────────────────────────────
// TOAST
// ─────────────────────────────────────────────────────────────
function showToast(title, message) {
    let notify_content =
        `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]">
            <li>
                <div class="toast-msg flex max-sm:max-w-[300px] flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]">
                    <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify">
                        <img src="/public/img/close-toast.svg" alt="close">
                    </a>
                    <div><img src="/public/img/toast-error.svg" alt="toast error"></div>
                    <div>
                        <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px]">${title}</h3>
                        <p class="text-[#262626] text-[12px] font-normal leading-[15px]">${message}</p>
                    </div>
                </div>
            </li>
        </ul>`;
    $(notify_content).insertBefore(".header-rht");
    setTimeout(function () { $('.toast-msg').fadeOut('slow', function () { $(this).remove(); }); }, 5000);
}

// =============================================================
// DRAG & DROP — 3-LEVEL SUPPORT
// =============================================================

// ── LEVEL 1 : Root ───────────────────────────────────────────
$('#accordionExample1').sortable({
    handle: '.drag-handle',
    items: '.menuitemsdiv',
    connectWith: '.child-container, .menuitemsdiv',
    placeholder: 'sortable-placeholder',
    start: function (e, ui) { ui.item.addClass('dragging'); },
    receive: function (event, ui) {
        if (ui.item.find('.child-container .childdiv, .group').length > 0 &&
            $(this).hasClass('child-container')) {
            $(ui.sender).sortable('cancel');
        }
    },
    stop: function (e, ui) {
        ui.item.removeClass('dragging');
        handleRootDrop(e, ui.item);
    }
});

// ── LEVEL 2 : Child containers (static) ──────────────────────
$('.child-container').sortable({
    handle: '.drag-handle',
    items: '> .childdiv',
    connectWith: '#accordionExample1, .child-container, .menuitemsdiv',
    placeholder: 'sortable-placeholder',
    start:   function (e, ui) { ui.item.addClass('dragging'); },
    receive: function (e, ui) { handleChildDrop(e, ui.item); },
    stop:    function (e, ui) { ui.item.removeClass('dragging'); handleChildDrop(e, ui.item); }
});

// ── LEVEL 3 : Subchild containers (static) ───────────────────
$('.subchild-container').sortable({
    handle: '.drag-handle',
    items: '> .childdiv',
    connectWith: '#accordionExample1, .child-container, .subchild-container',
    placeholder: 'sortable-placeholder ms-[40px]',
    start: function (e, ui) { ui.item.addClass('dragging'); ui.placeholder.height(ui.item.height()); },
    receive: function (e, ui) {
        if (ui.item.hasClass('menuitemsdiv') || ui.item.find('.child-container, .subchild-container').length > 0) {
            $(ui.sender).sortable('cancel'); return;
        }
        ui.item.addClass('ms-[40px] childdiv').attr('data-parent', parseInt($(this).attr('data-parent-id'), 10));
    },
    stop: function (e, ui) { ui.item.removeClass('dragging'); handleSubChildDrop(e, ui.item); }
});

// ── ROOT DROP ─────────────────────────────────────────────────
function handleRootDrop(event, $wrapper) {
    var $group      = $wrapper.find('.group').first();
    var parentId    = getRootParentId();
    var hasChildren = $wrapper.find('.child-container').length > 0;
    var $hit        = $(document.elementFromPoint(event.clientX, event.clientY));

    var $target =
        $hit.closest('.menuitemsdiv').length   ? $hit.closest('.menuitemsdiv') :
        $hit.closest('.childdiv').length        ? $hit.closest('.childdiv') :
        $hit.closest('.child-container').length ? $hit.closest('.child-container').closest('.menuitemsdiv') :
        $hit.closest('.group').length           ? $hit.closest('.menuitemsdiv') : $();

    if ($target.is($wrapper)) $target = $();

    if ($hit.closest('.child-container').length) {
        if (hasChildren) { moveToRoot($wrapper, $group, parentId); return; }
        var $cc = $hit.closest('.child-container');
        makeItemChildOf($wrapper, $cc.closest('.menuitemsdiv').length ? $cc.closest('.menuitemsdiv') : $cc);
        return;
    }

    if (hasChildren)    { moveToRoot($wrapper, $group, parentId); return; }
    if ($target.length) { makeItemChildOf($wrapper, $target);     return; }
    moveToRoot($wrapper, $group, parentId);
}

// ── MOVE TO ROOT ──────────────────────────────────────────────
function moveToRoot($wrapper, $group, parentId) {
    $group.removeClass('ms-[20px] ms-[40px] childdiv').attr('data-parent', parentId);
    if (!$wrapper.parent().is('#accordionExample1')) { $('#accordionExample1').append($wrapper); }
    updateOrderForContainer($('#accordionExample1'), parentId);
}

// ── MAKE ITEM A CHILD ─────────────────────────────────────────
function makeItemChildOf($sourceWrapper, $targetWrapper) {
    var $sourceRow = $sourceWrapper.find('.group').first();
    var targetId   = parseInt($targetWrapper.find('.group').first().data('id'), 10);

    // FIX: plain .find() — child-container is nested deep inside the target wrapper
    var $childContainer = $targetWrapper.find('.child-container').first();

    if (!$childContainer.length) {
        $childContainer = $('<div class="child-container ms-[20px]" data-parent-id="' + targetId + '"></div>');
        // Append to the accordion-item div that carries the .div{id} class
        $('.div' + targetId).append($childContainer);
        initializeChildSortable($childContainer);
    }

    $sourceRow.addClass('ms-[20px] childdiv').attr('data-parent', targetId).appendTo($childContainer);
    $sourceWrapper.remove();
    updateOrderForContainer($childContainer, targetId);
}

// ── MAKE ITEM A SUBCHILD (LEVEL 3) ───────────────────────────
function makeItemSubChildOf($sourceChild, $targetChild) {
    $targetChild = $targetChild.closest('.childdiv');
    if (!$targetChild.length) return;
    if ($targetChild.parent().hasClass('subchild-container')) return; // block 4th level

    var targetId = parseInt($targetChild.attr('data-id') || $targetChild.find('.group').first().data('id'), 10);
    if (isNaN(targetId)) return;

    var $accordionItem = $targetChild.find('.accordion-item').first();

    var $subContainer = $accordionItem.children('.subchild-container');
    if (!$subContainer.length) {
        $subContainer = $('<div class="subchild-container mt-[8px]"></div>')
            .attr('data-parent-id', targetId).appendTo($accordionItem);
        initializeSubChildSortable($subContainer);
    }

    $sourceChild.removeClass('ms-[20px]').addClass('ms-[40px] childdiv')
        .attr('data-parent', targetId).appendTo($subContainer);
    updateOrderForContainer($subContainer, targetId);
}

// ── CHILD DROP ────────────────────────────────────────────────
function handleChildDrop(event, $child) {
    // Add a third level if dropped on another child, but only if intentionally dragged to the right
    if (event && event.originalEvent) {
        var $hit         = $(document.elementFromPoint(event.originalEvent.clientX, event.originalEvent.clientY));
        var $targetChild = $hit.closest('.childdiv').not($child);

        if ($targetChild.length) {
            var targetRect = $targetChild[0].getBoundingClientRect();
            // Only nest if the mouse is at least 40px to the right of the target's left edge
            if (event.originalEvent.clientX > targetRect.left + 40) {
                makeItemSubChildOf($child, $targetChild);
                return;
            }
        }
    }

    if ($child.closest('#accordionExample1').length && !$child.closest('.child-container').length) {
        makeChildIndependent($child);
        updateOrderForContainer($('#accordionExample1'), getRootParentId());
        return;
    }

    var $immediateParent = $child.parent();
    if ($immediateParent.is('.child-container')) {
        var parentId = parseInt($immediateParent.attr('data-parent-id'), 10);
        if (isNaN(parentId)) return;
        $child.addClass('ms-[20px] childdiv').attr('data-parent', parentId);
        updateOrderForContainer($immediateParent, parentId);
    }
}

// ── SUBCHILD DROP ─────────────────────────────────────────────
function handleSubChildDrop(event, $subChild) {
    if ($subChild.closest('#accordionExample1').length &&
        !$subChild.closest('.child-container, .subchild-container').length) {
        makeSubChildIndependent($subChild); return;
    }

    var $hit    = $(document.elementFromPoint(event.originalEvent.clientX, event.originalEvent.clientY));
    var $target = $hit.closest('.childdiv').not($subChild);
    if ($target.length) { makeSubChildIndependent($subChild); return; }

    var $parent = $subChild.parent();
    if ($parent.is('.subchild-container')) {
        var parentId = parseInt($parent.attr('data-parent-id'), 10);
        $subChild.addClass('ms-[40px] childdiv').attr('data-parent', parentId);
        updateOrderForContainer($parent, parentId);
    }
}

// ── MAKE INDEPENDENT ─────────────────────────────────────────
function makeChildIndependent($child) {
    $child.removeClass('ms-[20px] ms-[40px] childdiv').attr('data-parent', getRootParentId());
    var $wrapper = $('<div class="menuitemsdiv"></div>').append($child);
    $('#accordionExample1').append($wrapper);
}

function makeSubChildIndependent($subChild) {
    $subChild.removeClass('ms-[20px] ms-[40px] childdiv').attr('data-parent', getRootParentId());
    var $wrapper = $('<div class="menuitemsdiv"></div>').append($subChild);
    $('#accordionExample1').append($wrapper);
    updateOrderForContainer($('#accordionExample1'), getRootParentId());
}

// ── INIT: CHILD SORTABLE ──────────────────────────────────────
function initializeChildSortable($childContainer) {
    if ($childContainer.hasClass('ui-sortable')) return;
    $childContainer.sortable({
        handle: '.drag-handle',
        items: '> .childdiv',
        connectWith: '#accordionExample1, .menuitemsdiv .child-container, .child-container',
        placeholder: 'sortable-placeholder ms-[20px]',
        start: function (e, ui) { ui.item.addClass('dragging'); ui.placeholder.height(ui.item.height()); },
        stop: function (e, ui) { ui.item.removeClass('dragging'); handleChildDrop(e, ui.item); },
        receive: function (e, ui) {
            var $item = ui.item.find('.group').first();
            if ($item.length && !$item.hasClass('childdiv')) {
                $item.addClass('ms-[20px] childdiv').attr('data-parent', $(this).attr('data-parent-id'));
            }
        }
    }).disableSelection();
}

// ── INIT: SUBCHILD SORTABLE ───────────────────────────────────
function initializeSubChildSortable($subContainer) {
    if ($subContainer.hasClass('ui-sortable')) return;
    $subContainer.sortable({
        handle: '.drag-handle',
        items: '> .childdiv',
        connectWith: '#accordionExample1, .child-container', // no subchild→subchild
        placeholder: 'sortable-placeholder ms-[40px]',
        start: function (e, ui) { ui.item.addClass('dragging'); ui.placeholder.height(ui.item.height()); },
        receive: function (e, ui) {
            if (ui.item.hasClass('menuitemsdiv') || ui.item.find('.child-container, .subchild-container').length > 0) {
                $(ui.sender).sortable('cancel'); return;
            }
            ui.item.addClass('ms-[40px] childdiv').attr('data-parent', $subContainer.attr('data-parent-id'));
        },
        stop: function (e, ui) { ui.item.removeClass('dragging'); handleSubChildDrop(e, ui.item); }
    }).disableSelection();
}

// ── INIT ALL ON PAGE LOAD ─────────────────────────────────────
function initializeAllChildContainers() {
    $('.child-container').each(function () {
        if (!$(this).hasClass('ui-sortable')) { initializeChildSortable($(this)); }
    });
    $('.subchild-container').each(function () {
        if (!$(this).hasClass('ui-sortable')) { initializeSubChildSortable($(this)); }
    });
}

// ── HELPERS ───────────────────────────────────────────────────
function getRootParentId() { return parseInt($('.parentmenu_id').val(), 10); }
function getMenuId($el)    { return parseInt($el.data('id'), 10); }

// ─────────────────────────────────────────────────────────────
// UPDATE ORDER & SAVE  ← THE FIXED FUNCTION
// ─────────────────────────────────────────────────────────────
function updateOrderForContainer($container, parentId) {
    var orderData = [];

    if ($container.attr('id') === 'accordionExample1') {
        // ── LEVEL 1 ──────────────────────────────────────────
        // FIX: .children('.group')  not  .find('> .group')
        // jQuery's .find() does NOT support the > direct-child combinator.
        $container.children('.menuitemsdiv').each(function (index) {
            var $row = $(this).children('.group'); // ← FIXED
            if (!$row.length) return;
            orderData.push({
                menuitem_id:   getMenuId($row),
                orderindex:    index + 1,
                parentmenu_id: parentId,
                is_child:      false
            });
        });

    } else {
        // ── LEVEL 2 & 3 ──────────────────────────────────────
        // FIX: iterate $container.children('.childdiv') directly.
        // The old code did $(`.group[data-id="${parentId}"]`).closest(...)
        // .find('> .child-container > .childdiv') which:
        //   a) used '>' inside .find() — doesn't work
        //   b) .child-container is nested inside .accordion-item inside
        //      .group, not a direct child — so the traversal found nothing.
        // $container IS the child-container / subchild-container, so just
        // iterate its direct .childdiv children.
        $container.children('.childdiv').each(function (index) {
            var $row = $(this).hasClass('group') ? $(this) : $(this).find('.group').first();
            if (!$row.length) return;
            orderData.push({
                menuitem_id:   getMenuId($row),
                orderindex:    index + 1,
                parentmenu_id: parentId,
                is_child:      true
            });
        });
    }


    // If orderData is still empty the AJAX call is pointless — and this
    // was the silent failure that caused "disappears on refresh".
    if (!orderData.length) {
        return;
    }

    // Representative row for the extra fields the endpoint requires
    var $anyRow;
    if ($container.attr('id') === 'accordionExample1') {
        $anyRow = $container.children('.menuitemsdiv').first().children('.group');
    } else {
        var $first = $container.children('.childdiv').first();
        $anyRow = $first.hasClass('group') ? $first : $first.find('.group').first();
    }
    if (!$anyRow.length) return;

    var label     = $anyRow.find('input[placeholder="Navigation Lable"]').val();
    var path      = $anyRow.find('input[placeholder="Path link"]').val();
    var updateBtn = $anyRow.find('#updatebtn');

    $.ajax({
        url: '/admin/website/menu/updatemenuitemsorder',
        type: 'POST',
        data: {
            menuitem_id:   getMenuId($anyRow),
            menu_name:     label,
            parentmenu_id: parentId,
            urlpath:       path,
            csrf:          $('input[name="csrf"]').val(),
            menu_typeid:   updateBtn.data('typeid'),
            type:          updateBtn.data('type'),
            webid:         $('.templateid').attr('data-id'),
            orderData:     JSON.stringify(orderData)
        },
        success: function () {
            console.log('Order saved successfully');
            window.location.reload();
        },
        error: function (xhr) {
            // Surface failures — previously silent, masking the root cause
            console.error('Order save FAILED:', xhr.status, xhr.responseText);
        }
    });
}

// =============================================================
// SEARCH / FILTER / UI HELPERS (unchanged)
// =============================================================
$(document).on('keyup', '.searchbtn', function () {
    var keyword  = $(this).val().trim().toLowerCase();
    var chkgrb   = $(this).parents('.relative').siblings('.tab-content').find('.chk-group');
    var hasMatch = false;

    if (keyword === '') {
        chkgrb.removeClass('hidden');
        $(this).parents('.relative').siblings('#nodatafounddesign').addClass('hidden');
    } else {
        chkgrb.each(function () {
            var pname = $(this).find('p').text().trim().toLowerCase();
            if (pname.includes(keyword)) { $(this).removeClass('hidden'); hasMatch = true; }
            else { $(this).addClass('hidden'); }
        });
        if (hasMatch) {
            $(this).parents('.relative').siblings('#nodatafounddesign').addClass('hidden');
            $(this).parents('.relative').siblings('.tab-content').find('#nodatafounddesignn').removeClass('hidden');
        } else {
            $(this).parents('.relative').siblings('#nodatafounddesign').removeClass('hidden');
            $(this).parents('.relative').siblings('.tab-content').find('#nodatafounddesignn').addClass('hidden');
        }
    }
});

$(document).on('click', '#editbtnn', function () {
    var menuname  = $(this).attr('data-name');
    var menutitle = $(this).attr('data-title');
    var menudesc  = $(this).attr('data-desc');
    var menu_Group = $(this).attr('data-group');
    var menu_Order = $(this).attr('data-order');
    var formatted_menu_Group = '';
    if (menu_Group) {
        formatted_menu_Group = menu_Group.replace(/-/g, " ").replace(/\b\w/g, c => c.toUpperCase());
    }
    var menustatus = $(this).attr('data-status');
    var data = $(this).attr("data-id");
    $("#menuform").attr("name", "editmenu").attr("action", "/admin/website/menu/updatemenu");
    $("input[name=menu_name]").val(menuname.trim());
    $("input[name=menu_title]").val(menutitle.trim());
    $("textarea[name=menu_desc]").val(menudesc.trim());
    $("input[name=menu_group]").val(menu_Group.trim());
    $("#menu-group-head").text(formatted_menu_Group);
    $("input[name=menu_order]").val(menu_Order.trim());
    $("#menu-order-head").text(menu_Order.trim());
    $('#menu_name').addClass('pointer-events-none opacity-50');
    if (menustatus == 1) { $('.menustatus').prop('checked', true).val('1'); }
    else { $('.menustatus').prop('checked', false).val('0'); }
    $("#savemenubtn").text(languagedata.update);
    $('#modalTitleIdd').text(languagedata.update + languagedata.Menu.menu);
    $("#menu_name-error").hide();
    $("#menu_desc-error").hide();
    $("input[name=menu_id]").val(data);
});

$(document).on('keyup', '.labelname', function () {
    if ($(this).val().trim().length >= 300) {
        $(this).removeClass('mb-[24px]');
        $(this).siblings('#lablename-error').removeClass('hidden').addClass('mb-[24px]');
    } else {
        $(this).addClass('mb-[24px]');
        $(this).siblings('#lablename-error').addClass('hidden').removeClass('mb-[24px]');
    }
});

$(document).on('click', '.opentab', function () {
    $('#customumenurl-error').addClass('hidden');
    $('#custommenuname-error').addClass('hidden');
    var activetab  = $(this).siblings('.accordion-collapse').children('.accordion-body').children('.tab-content').find('.tab-pane.active');
    var checkboxes = activetab.find('.channelnameinput');
    var $btn       = $(this).siblings('.accordion-collapse').children('.accordion-body').find('.courseallselect');
    if (checkboxes.length && checkboxes.filter(':checked').length === checkboxes.length) {
        $btn.text(languagedata.Menu.deselectall);
    } else {
        $btn.text(languagedata.Menu.selectall);
    }
});

$(document).on('click', '.courseallselect', function () {
    var activetab  = $(this).parent().siblings('.tab-content').find('.tab-pane.active');
    var checkboxes = activetab.find('.channelnameinput');
    var allChecked = checkboxes.length && checkboxes.filter(':checked').length === checkboxes.length;
    if (allChecked) { checkboxes.prop('checked', false); $(this).text(languagedata.Menu.selectall); }
    else            { checkboxes.prop('checked', true);  $(this).text(languagedata.Menu.deselectall); }
});

$(document).on('click', '.channelnameinput', function () {
    var $tabPane   = $(this).closest('.tab-pane');
    var $checkboxes = $tabPane.find('.channelnameinput');
    var $btn        = $tabPane.closest('.tab-content').siblings('.nav').find('.courseallselect');
    var allChecked  = $checkboxes.length && $checkboxes.filter(':checked').length === $checkboxes.length;
    if ($btn.length) { $btn.text(allChecked ? languagedata.Menu.deselectall : languagedata.Menu.selectall); }
});

$(document).on('click', '.recentchannel',        function () { $('#allchannel').removeClass('show active'); });
$(document).on('click', '.recentcourse',          function () { $('#allcourse').removeClass('show active'); });
$(document).on('click', '.recentform',            function () { $('#allforms').removeClass('show active'); });
$(document).on('click', '.recentcategory',        function () { $('#allcategory').removeClass('show active'); });
$(document).on('click', '.recentlistingCategory', function () { $('#alllistingCategory').removeClass('show active'); });
$(document).on('click', '.recentpages',           function () { $('#allpages').removeClass('show active'); });

$(document).on('click', '.cancelmenuitem', function () {
    $('#menu_name-error').hide(); $('#urlpath-error').hide();
    $('input[name="lang"]:checked').prop('checked', false);
    $('.customUrlInput, .entriesDropdown, .pagesDropdown, .listingsDropdown, .categoryDropdown, .channelDropdown').addClass('hidden');
    $('.navpath').val(""); $('input[name="menu_name"]').val("");
    $('input[name="meta_title"]').val(""); $('input[name="meta_description"]').val(""); $('input[name="meta_keywords"]').val("");
    $('.listingsDropdownMenu li').find('input[type="checkbox"]').prop('checked', false);
    $('.categoryDropdownMenu li').find('input[type="checkbox"]').prop('checked', false);
});

$(document).on('click', '.addcoursemenu', function () {
    var parentmenuid = $('#menu_id').val();
    var exists = false;
    $('#accordionExample1 .accordion-item').each(function () {
        if ($(this).find('#updatebtn').data('type') == "courses") { exists = true; return false; }
    });
    if (!exists) {
        $.ajax({
            url: "/admin/website/menu/createmenuitems", type: "POST", async: false,
            data: { "menu_name": "All Courses", "menu_id": parentmenuid, "urlpath": "/courses",
                    csrf: $("input[name='csrf']").val(), "menu_typeid": "", "type": "courses",
                    "webid": $('.templateid').attr('data-id') },
            dataType: "json", cache: false, success: function () { location.reload(); }
        });
    }
});

$(document).ready(function () {
    $('input[name="lang"]').change(function () {
        selectedLabel = $(this).next('label').text().trim();
        $('.navpath').val("");
        $('.customUrlInput, .entriesDropdown, .pagesDropdown, .listingsDropdown, .categoryDropdown, .channelDropdown').addClass('hidden');
        if      (selectedLabel === "Custom URL")  { $('.customUrlInput').removeClass('hidden');  $('.menutype').val("custom_url"); }
        else if (selectedLabel === "Entries")     { $('.entriesDropdown').removeClass('hidden');  $('.menutype').val("entries"); }
        else if (selectedLabel === "Pages")       { $('.pagesDropdown').removeClass('hidden');   $('.menutype').val("pages"); }
        else if (selectedLabel === "Listings")    { $('.listingsDropdown').removeClass('hidden'); $('.menutype').val("listings"); }
        else if (selectedLabel === "Categories")  { $('.categoryDropdown').removeClass('hidden'); $('.menutype').val("categories"); $('.navpath').val("/categories"); }
        else if (selectedLabel === "None")        { $('.menutype').val("none"); }
        else if (selectedLabel === "Channels")    { $('.channelDropdown').removeClass('hidden'); $('.menutype').val("channels"); }
    });
});

$(document).on('keyup', '.customurl', function () {
    var url = $(this).val().trim();
    if (url != "") { $('#urlpath-error').addClass('hidden'); } else { $('#urlpath-error').removeClass('hidden'); }
    $('.navpath').val(url);
});

$(document).on('click', '.entryslug',  function () { $('.navpath').val("/categories/" + $(this).attr('data-slug').trim()); });
$(document).on('click', '.pageslug',   function () {
    var url = $(this).attr('data-slug').trim();
    $('.navpath').val("/pages/" + url); $('.selectpage').text(url); $('.menu_typeid').val($(this).attr('data-id'));
});
$(document).on('click', '.channelname', function () {
    var url = $(this).attr('data-slug').trim();
    $('.navpath').val("/channel/" + url); $('.selectchannel').text(url);
});

$(".searchcatlists").keyup(function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    var dropdownContainer = $(this).closest('.dropdown');
    dropdownContainer.find("li.entry-dropdownlist").each(function () {
        var isVisible = $(this).find('label').text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    dropdownContainer.find('.noData-foundentry').toggle(!found);
});

$(".searchpagelist").keyup(function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    var dropdownContainer = $(this).closest('.dropdown');
    dropdownContainer.find("li.page-dropdownlist").each(function () {
        var isVisible = $(this).find('label').text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    dropdownContainer.find('.noData-foundWrapperr').toggle(!found);
});

$(".searchlistinglist").keyup(function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    var dropdownContainer = $(this).closest('.dropdown');
    dropdownContainer.find("ul li").each(function () {
        var isVisible = $(this).find('label').text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    dropdownContainer.find('.noData-foundWrapperr').toggle(!found);
});

$(".searchchannellist").keyup(function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    var dropdownContainer = $(this).closest('.dropdown');
    dropdownContainer.find("ul li").each(function () {
        var isVisible = $(this).find('span').text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    dropdownContainer.find('.noData-foundWrapperr').toggle(!found);
});

$(document).ready(function () {
    $('.listingsDropdownMenu').on('click', function (e) { e.stopPropagation(); });
    $('.categoryDropdownMenu').on('click', function (e) { e.stopPropagation(); });
});

$(".searchcategorylists").on("keyup", function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    $(".catrgory-list").each(function () {
        var categoryText = $(this).find("span").map(function () { return $(this).text().toLowerCase(); }).get().join(" ");
        var isVisible = categoryText.indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    $(".noData-foundcategory").toggle(!found);
});

$(".searchlslists").on("keyup", function () {
    var found = false; var searchTerm = $(this).val().trim().toLowerCase();
    $(".listingcat").each(function () {
        var categoryText = $(this).find("span").map(function () { return $(this).text().toLowerCase(); }).get().join(" ");
        var isVisible = categoryText.indexOf(searchTerm) > -1;
        $(this).toggle(isVisible); if (isVisible) found = true;
    });
    $(".noData-foundcategorylist").toggle(!found);
});

$('.selectcheckbox').on('click', function () {
    var dataId = String($(this).closest('li.catrgory-list').find('span').last().attr('data-id'));
    categoryarr = categoryarr.flat();
    if ($(this).prop('checked')) { if (!categoryarr.includes(dataId)) { categoryarr.push(dataId); } }
    else { categoryarr = categoryarr.filter(function (id) { return id !== dataId; }); }
});

$('.imgupload').on('change', function (e) {
    var file = e.target.files[0];
    if (!file) return;
    if (file.type !== 'image/svg+xml' && !file.name.endsWith('.svg')) { $(this).val(''); $('#svgHidden').val(''); return; }
    var reader = new FileReader();
    reader.onload = function (event) {
        $('#svgHidden').val(event.target.result);
        $(e.target).closest('.imgupldiv').addClass('hidden');
        $('#ImageName').text(file.name); $('#imageRemoveDiv').removeClass('hidden');
        $('.svgins').addClass('hidden'); $('.uploadbtn').addClass('hidden'); $('.ImageNamediv').removeClass('hidden');
    };
    reader.readAsDataURL(file);
});

$('#imageRemoveDiv').on('click', function () {
    $('#svgHidden').val(''); $('.imgupload').val(''); $('#ImageName').text('');
    $('#imageRemoveDiv').addClass('hidden'); $('.imgupldiv').removeClass('hidden');
    $('.svgins').removeClass('hidden'); $('.uploadbtn').removeClass('hidden'); $('.ImageNamediv').addClass('hidden');
});

$('.deleteimg').on('click', function () {
    $(this).closest('.imageRemoveDiv').addClass('hidden');
    $(this).closest('.imageRemoveDiv').siblings('.imgupldiv').removeClass('hidden').find('.svgval').val('');
    $(this).siblings('.svgDelete').val('1');
});

$(document).on('change', '.imguploadd', function (e) {
    var file = e.target.files[0];
    if (!file) return;
    if (file.type !== 'image/svg+xml' && !file.name.endsWith('.svg')) {
        $(this).val(''); $(this).closest('.imgupldiv').find('.svgval').val(''); return;
    }
    var reader = new FileReader(); var uploaddiv = $(this);
    reader.onload = function (event) {
        uploaddiv.closest('.imgupldiv').find('.svgval').val(event.target.result);
        uploaddiv.closest('.imgupldiv').addClass('hidden');
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').removeClass('hidden')
            .find('.ImageName').text(file.name);
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').find('.deleteimg').removeClass('hidden');
        uploaddiv.closest('.imgupldiv').find('.svgins').addClass('hidden');
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').find('.svgDelete').val('');
    };
    reader.readAsDataURL(file);
});

$(document).on('click', '.menuitemeditbtn', function () {
    var menuid = $(this).attr('data-id');
    var webid  = $('.webid').val();
    $.ajax({
        url: "/admin/website/menu/editmenuitem/" + menuid, type: "GET", dataType: "json",
        success: function (result) {
            $('#modalTitleId').text('Update Menu Item'); $('#menuitemssavebtn').text('Update item');
            $('#menuitemform').attr('action', '/admin/website/menu/updatemenuitems?webid=' + webid);
            $('input[name="menu_name"]').val(result.Name); $('.menuitem_id').val(menuid);
            $('#ImageName').text(result.ImageName); $('#separatewindow').val(result.SeparateWindow);
            $('#separatewindow').prop('checked', result.SeparateWindow == 1);
            $('.navpath').val(result.UrlPath);
            $('input[name="meta_title"]').val(result.MetaTitle);
            $('input[name="meta_description"]').val(result.MetaDescription);
            $('input[name="meta_keywords"]').val(result.MetaKeywords);
            $('.parentmenu_id').val(result.ParentId); $('.menu_typeid').val(result.TypeId);
            if (result.ImageName != "") {
                $('.svgins').addClass('hidden'); $('.uploadbtn').addClass('hidden');
                $('#imageRemoveDiv').removeClass('hidden'); $('.ImageNamediv').removeClass('hidden');
            }
            $('.menutype').val(result.Type);
            if      (result.Type == "none")       { $('#radionone').prop('checked', true); selectedLabel = "None"; }
            else if (result.Type == "pages")       { $('#radioPages').trigger('click'); $('.navpath').val(result.UrlPath); $('.selectpage').text(result.UrlPath.substring(result.UrlPath.indexOf("/pages/") + 7)); }
            else if (result.Type == "categories")  {
                $('#radioCategories').trigger('click'); $('.navpath').val(result.UrlPath);
                var categoryIdsArray = result.CategoryIds.split(','); categoryarr.push(categoryIdsArray);
                $('.categoryDropdownMenu li').each(function () {
                    var dataId = $(this).find('span').last().attr('data-id');
                    $(this).find('input[type="checkbox"]').prop('checked', categoryIdsArray.includes(dataId));
                });
            }
            else if (result.Type == "listings")    {
                $('#radioListings').trigger('click'); $('.navpath').val(result.UrlPath);
                var listingIdsArray = result.ListingsIds.split(',');
                $('.listingsDropdownMenu li').each(function () {
                    $(this).find('input[type="checkbox"]').prop('checked', listingIdsArray.includes($(this).attr('data-id')));
                });
            }
            else if (result.Type == "custom_url")  { $('#radioCustomUrl').trigger('click'); $('.navpath').val(result.UrlPath); }
            else if (result.Type == "channels")    {
                $('#radioChannels').trigger('click'); $('.navpath').val(result.UrlPath);
                $('.selectchannel').text(result.UrlPath.substring(result.UrlPath.indexOf("/channel/") + 9));
            }
        }
    });
});

$(document).on('click', '.addmenuitembtn', function () {
    var webid = $('.webid').val();
    $('#modalTitleId').text('Add Menu Item'); $('#menuitemssavebtn').text('Add Item');
    $('#menuitemform').attr('action', '/admin/website/menu/createmenuitems?webid=' + webid);
    selectedLabel = ""; $('.ImageNamediv').addClass('hidden'); $('.uploadbtn').removeClass('hidden');
    $('.selectpage').text('Select Page'); $('.selectchannel').text('Select Channel');
});

$(".menuitemname").on('keyup', function () { $("label[for='menu_name']").hide(); });

$(document).on('click', '.addcustommenu', function () {
    var parentmenuid   = $('#menu_id').val();
    var url            = $('.customumenurl').val();
    var menuname       = $('.custommenuname').val();
    var menu_grouptype = $('#parentmenu_grouptype').val();
    if (url == "")     { $('#customumenurl-error').removeClass('hidden'); }
    if (menuname == "") { $('#custommenuname-error').removeClass('hidden'); }
    if (url != "" && menuname != "") {
        $('#customumenurl-error').addClass('hidden'); $('#custommenuname-error').addClass('hidden');
        $.ajax({
            url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').data('id'),
            method: "POST", dataType: "json", async: false,
            data: { menu_name: menuname, menu_id: parentmenuid, urlpath: url, menu_typeid: "",
                    csrf: $("input[name='csrf']").val(), type: "custom_url",
                    webid: $('.webid').data('id'), parentmenu_grouptype: menu_grouptype },
            success: function () { location.reload(); }
        });
    }
});

$(document).on('keyup', '.custommenuname', function () {
    if ($(this).val().trim() != "") { $('#custommenuname-error').addClass('hidden'); }
    else { $('#custommenuname-error').removeClass('hidden'); }
});
$(document).on('keyup', '.customumenurl', function () {
    if ($(this).val().trim() != "") { $('#customumenurl-error').addClass('hidden'); }
    else { $('#customumenurl-error').removeClass('hidden'); }
});

$(document).on('mousedown', '.quill-editor-container', function (e) {
    e.stopPropagation();
    var quill = $(this).data('quill');
    if (!quill) return;
    setTimeout(function () { if (!quill.hasFocus()) { quill.focus({ preventScroll: true }); } }, 0);
});