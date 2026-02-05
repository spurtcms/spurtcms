var languagedata
var selectedLabel
var categoryarr = [];
/** */
$(document).ready(async function () {

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

        if (uniqueCategoryIds.includes(dataIdNum)) {
            $(this).find('input[type="checkbox"]').prop('checked', true);
        } else {
            $(this).find('input[type="checkbox"]').prop('checked', false);
        }
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

        if (uniquelistingIds.includes(dataIdNum)) {
            $(this).find('input[type="checkbox"]').prop('checked', true);
        } else {
            $(this).find('input[type="checkbox"]').prop('checked', false);
        }
    });


})
//savebutton function
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


    let selectedIds = $('.selected-listings:checked').map(function () {
        return $(this).data('id');
    }).get();
    let joinedIds = selectedIds.join(',');
    $('#listingsids').val(joinedIds);

    if (categoryarr.length > 0) {
        $('#categoryids').val(categoryarr.join(','));
    }
    if (selectedLabel == "Listings") {
        $(".navpath").val("/listings/");
    }
    if (selectedLabel == "Categories") {
        $('.navpath').val("/categories/");
    }


    $("#menuitemform").validate({
        onkeyup: false,
        onfocusout: false,
        onclick: false,
        rules: {
            menu_name: {
                required: true,
                space: true,
                duplicatename: true,
                maxlength: 300,
            },
            urlpath: {
                required: function () {
                    return selectedLabel !== "None";
                },
                space: function () {
                    return selectedLabel !== "None";
                }
            }
        },
        messages: {
            menu_name: {
                required: "*" + languagedata.Menu.labelnameerr,
                space: "* " + languagedata.spacergx,
                duplicatename: "*" + languagedata.Menu.nameduplicateerr,
                maxlength: "*" + languagedata.Menu.namelimiterr
            },
            urlpath: {
                required: "*" + languagedata.Menu.linkpatherr,
                space: "* " + languagedata.spacergx,
            },
        }
    });

    // Trigger validation manually (only runs now, not on keyup)
    var formcheck = $("#menuitemform").valid();
    if (formcheck === true) {
        $('#menuitemform')[0].submit();
    }

    return false;
});

var collapsecount = 0;

var menuitemsarr = []

$(document).on('click', '.addchannelmenu', function (e) {
    e.preventDefault();


    var $accordionBody = $(this).closest('.accordion-body');

    $accordionBody.find('.channelnameinput:checked').each(function () {
        var menuname = $(this).siblings('label').find('p').text().trim();
        var id = $(this).attr('data-id');
        var type = $(this).attr('data-type');
        var parentmenuid = $('#menu_id').val();
        var slug =$(this).attr('data-slug')



        var exists = false;
        $('#accordionExample1 .accordion-item').each(function () {
            var $updateBtn = $(this).find('#updatebtn');
            console.log(menuname, $updateBtn.data('name'), "namesss")
            if (

                $updateBtn.data('type') == type &&
                $updateBtn.data('typeid') == id ||
                $updateBtn.data('name') == menuname
            ) {
                exists = true;

                return false;
            }
        });



        if (!exists) {
            $.ajax({
                url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').attr('data-id'),
                type: "POST",
                async: false,
                data: {
                    "menu_name": menuname,
                    "menu_id": parentmenuid,
                    "urlpath": "",
                    csrf: $("input[name='csrf']").val(),
                    "menu_typeid": id,
                    "type": type,
                     "slug_name":slug,
                    "webid": $('.webid').attr('data-id'),
                },
                dataType: "json",
                cache: false,
                success: function (result) {

                    location.reload();
                }
            });
        }
    });
});



// New separate function for category add
function addCategoryFunction(type, classname) {
    var $accordionBody = $('.' + classname).closest('.accordion-body');

    $accordionBody.find('.channelnameinput:checked').each(function () {
        var $checkbox = $(this);
        var $label = $checkbox.siblings('label');


        var $childSpan = $label.find('span[data-id]').last();

        var subcategoryid = $childSpan.data('id');
        var subcategoryname = $.trim($childSpan.text());
        var parentmenuid = $('#menu_id').val();
        var exists = false;
        $('#accordionExample1 .accordion-item').each(function () {
            var $updateBtn = $(this).find('#updatebtn');
 
            if (

                $updateBtn.data('type') == type &&
                $updateBtn.data('typeid') == subcategoryid &&
                $updateBtn.data('name') == subcategoryname

               
            ) {
                exists = true;

                return false;
            }
        });



        if (!exists) {
            $.ajax({
                url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').data('id'),
                method: "POST",
                dataType: "json",
                async: false,
                data: {
                    menu_name: subcategoryname,
                    menu_id: parentmenuid,
                    menu_typeid: subcategoryid,
                    csrf: $("input[name='csrf']").val(),
                    type: type,
                    webid: $('.webid').data('id')
                },
                success: function () {
                    location.reload();
                }
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
$(document).on('click', '#updatebtn', function () {
    var menuId = $(this).data('id');
    var type = $(this).attr('data-type');
    var typeid = $(this).attr('data-typeid');
    parentmenuid = $(this).attr('data-parent');
    var $accordionItem = $(this).closest('.accordion-body');
    var label = $accordionItem.find('input[placeholder="Navigation Lable"]').val().trim();
    var path = $accordionItem.find('input[placeholder="Path link"]').val();
    var svgimage = $accordionItem.find('.svgval').val()
    var svgDelete = $accordionItem.find('.svgDelete').val()
    var separatewindow = $accordionItem.find('.separatewindow').val()

    var labelValid = label.length <= 300;
    var isDuplicate = false;
     var $editor = $accordionItem.find('#editor-' + menuId);
        
        var quill = $editor.data('quill');
        var editorContent = quill ? quill.root.innerHTML : '';
        
        // Clean content
        editorContent = editorContent.replace(/<p><br><\/p>$/, '').trim();
        
    console.log(editorContent,"editocondtent")
    
    $.ajax({
        url: "/admin/website/menu/checkmenuname",
        type: "POST",
        async: false,
        data: {
            "menu_name": label,
            "webid": $(".templateid").val(),
            "menu_id": menuId,

            csrf: $("input[name='csrf']").val()
        },
        datatype: "json",
        cache: false,
        success: function (data) {
            var result = data.trim();

            console.log(result, "resulttt")
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
                "editorContent":editorContent,

            },
            datatype: "json",
            cache: false,
            success: function (data) {
                result = data
                console.log(result, result.Id, result.ParentId, "resultttt")

                if (result) {

                    console.log("checkconditionsd")
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



$(document).on('click', '#menuitemdeletebtn', function () {

    var menuid = $(this).attr("data-id")
    $("#content").text(languagedata.Menu.removeconfirmation)
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')


    if (pageno == null) {
        $('#delid').attr('href', "/admin/website/menu/deletemenuitem/" + menuid + "?webid=" + $('.templateid').attr('data-id'));

    } else {
        $('#delid').attr('href', "/admin/website/menu/deletemenuitem/" + menuid + "?webid=" + $('.templateid').attr('data-id') + "?page=" + pageno);

    }
    $(".deltitle").text(languagedata.Menu.deletemenu + "?")

    var label = $(this).attr("data-name")
    $('.delname').text(label)

})




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
    setTimeout(function () {
        $('.toast-msg').fadeOut('slow', function () {
            $(this).remove();
        });

        // window.location.reload()
    }, 5000);
}




$('#accordionExample1').sortable({
    handle: '.drag-handle',
    items: '.menuitemsdiv',
    connectWith: '.child-container,.menuitemsdiv',
    placeholder: 'sortable-placeholder',
    start: function (e, ui) {
        ui.item.addClass("dragging");
    },

    receive: function (event, ui) {
        const dragged = ui.item;


        const hasChildren = dragged
            .find('.child-container .childdiv,.group')
            .length > 0;


        const dropTarget = $(this);
        const isDroppingAsChild = dropTarget.hasClass("child-container");


        if (hasChildren && isDroppingAsChild) {
          
            $(ui.sender).sortable('cancel');
        }
    },

    stop: function (e, ui) {
        ui.item.removeClass('dragging');
        handleRootDrop(e, ui.item);

    }
}).disableSelection();


$('.child-container').sortable({
    handle: '.drag-handle',
    items: '> .childdiv',
    connectWith: '#accordionExample1, .child-container,.menuitemsdiv',
    placeholder: 'sortable-placeholder',
    start: function (e, ui) { ui.item.addClass('dragging'); },
    receive: function (e, ui) {
        handleChildDrop(e, ui.item);
    },
    stop: function (e, ui) {
        ui.item.removeClass('dragging');
        handleChildDrop(e, ui.item);
    }
}).disableSelection();
// function handleRootDrop(event, $wrapper) {
//     const $group = $wrapper.find('.group').first();
//     const parentId = getRootParentId();
//     const draggedHasChildren = $wrapper.find('.child-container').length > 0;


//     let $hit = $(document.elementFromPoint(event.clientX, event.clientY));


//     let $target =
//         $hit.closest('.menuitemsdiv').length ? $hit.closest('.menuitemsdiv') :
//             $hit.closest('.childdiv').length ? $hit.closest('.childdiv') :
//                 $hit.closest('.child-container').length ? $hit.closest('.child-container').closest('.menuitemsdiv') :
//                     $hit.closest('.group').length ? $hit.closest('.menuitemsdiv') :
//                         $();


//     if ($target.is($wrapper)) {
//         $target = $();
//     }


//     if (draggedHasChildren) {
//         moveToRoot($wrapper, $group, parentId);
//         return;
//     }

//     if ($target.length) {
//         makeItemChildOf($wrapper, $target);
//         return;
//     }

//     moveToRoot($wrapper, $group, parentId);
// }
function handleRootDrop(event, $wrapper) {
    const $group = $wrapper.find('.group').first();
    const parentId = getRootParentId();
    const draggedHasChildren = $wrapper.find('.child-container').length > 0;

    let $hit = $(document.elementFromPoint(event.clientX, event.clientY));

    let $target =
        $hit.closest('.menuitemsdiv').length ? $hit.closest('.menuitemsdiv') :
        $hit.closest('.childdiv').length ? $hit.closest('.childdiv') :
        $hit.closest('.child-container').length ? $hit.closest('.child-container').closest('.menuitemsdiv') :
        $hit.closest('.group').length ? $hit.closest('.menuitemsdiv') :
        $();

    if ($target.is($wrapper)) {
        $target = $();
    }

    if ($hit.closest('.child-container').length) {
        if (draggedHasChildren) {
            // If dragged item has children, do not allow dropping inside a child-container
            // Instead, you can move to root or just return to prevent drop
            moveToRoot($wrapper, $group, parentId);
            return;
        }
        let $childContainerTarget = $hit.closest('.child-container');
        makeItemChildOf($wrapper, $childContainerTarget.closest('.menuitemsdiv').length ? $childContainerTarget.closest('.menuitemsdiv') : $childContainerTarget);
        return;
    }

    if (draggedHasChildren) {
        moveToRoot($wrapper, $group, parentId);
        return;
    }

    if ($target.length) {
        makeItemChildOf($wrapper, $target);
        return;
    }

    moveToRoot($wrapper, $group, parentId);
}




function moveToRoot($wrapper, $group, parentId) {
    $group.removeClass('ms-[20px] childdiv');
    $group.attr('data-parent', parentId);


    if (!$wrapper.parent().is('#accordionExample1')) {
        $('#accordionExample1').append($wrapper);
    }

    updateOrderForContainer($('#accordionExample1'), parentId);
}


function makeItemChildOf($sourceWrapper, $targetWrapper) {
    let $sourceRow = $sourceWrapper.find('.group').first();
    let targetId = parseInt($targetWrapper.find('.group').first().data('id'), 10);


    let $childContainer = $targetWrapper.find('.child-container').first();


    if (!$childContainer.length) {
        let $accordionItem = $targetWrapper.find('.accordion-item').first();
        let $deleteBtnWrapper = $targetWrapper.find('div').last();

        $childContainer = $('<div class="child-container ui-sortable" data-parent-id="' + targetId + '"></div>');

        console.log("checkconditiond")
        if ($deleteBtnWrapper.length) {
            console.log("checkconditiond1")
            $('.div' + targetId).append($childContainer);
            $childContainer.addClass('ms-[20px] childdiv');
            $childContainer.attr('data-parent', targetId);
            // $childContainer.insertBefore($deleteBtnWrapper);

        } else {
            console.log("checkconditiond2")
            $accordionItem.after($childContainer);
        }

        initializeChildSortable($childContainer);
    }


    $sourceRow.addClass('ms-[20px] childdiv');
    $sourceRow.attr('data-parent', targetId);
    $sourceRow.appendTo($childContainer);
    $sourceWrapper.remove();
    updateOrderForContainer($childContainer, targetId);
}


function initializeChildSortable($childContainer) {
    $childContainer.sortable({
        handle: '.drag-handle',
        items: '> .childdiv',
        connectWith: '#accordionExample1, .menuitemsdiv .child-container, .child-container',
        placeholder: 'sortable-placeholder ms-[20px]',
        start: function (e, ui) {
            ui.item.addClass('dragging');
            ui.placeholder.height(ui.item.height());
        },
        stop: function (e, ui) {
            handleChildDrop(e, ui.item);
        },
        receive: function (e, ui) {

            let $item = ui.item.find('.group').first();
            if ($item.length && !$item.hasClass('childdiv')) {
                $item.addClass('ms-[20px] childdiv');
                $item.attr('data-parent', $(this).data('parent-id'));
            }
        }
    }).disableSelection();
}


function initializeAllChildContainers() {
    $('.child-container').each(function () {
        if (!$(this).hasClass('ui-sortable')) {
            initializeChildSortable($(this));
        }
    });
}


$(document).ready(function () {
    initializeAllChildContainers();
});

function handleChildDrop(event, $child) {
    let $container = $child.closest('.child-container, #accordionExample1, .menuitemsdiv');
    let $sourceContainer = $(event.target)


    if (($container.is('.child-container'))) {

        let parentId = parseInt($container.data('parent-id'), 10);


        $child.addClass('ms-[20px] childdiv');
        $child.attr('data-parent', parentId);
        updateOrderForContainer($container, parentId);
    } else if ($container.is('#accordionExample1')) {

        makeChildIndependent($child);
        updateOrderForContainer($('#accordionExample1'), getRootParentId());
    }

}




function makeChildIndependent($child) {
    $child.removeClass('ms-[20px] childdiv');
    $child.attr('data-parent', getRootParentId());
    let $wrapper = $('<div class="menuitemsdiv"></div>');
    $wrapper.append($child);
    $('#accordionExample1').append($wrapper);
}
function getRootParentId() {
    return parseInt($('.parentmenu_id').val(), 10);
}

function getMenuId($el) {
    return parseInt($el.data('id'), 10);
}


function updateOrderForContainer($container, parentId) {
    let orderData = [];


    if ($container.attr("id") === "accordionExample1") {
        
        $container.children(".menuitemsdiv").each(function (index) {
            let $row = $(this).find("> .group");
            orderData.push({
                menuitem_id: getMenuId($row),
                orderindex: index + 1,
                parentmenu_id: parentId,
                is_child: false
            });
        });
    }



    let $parentItem = $(`.dragitem[data-id="${parentId}"]`);
    let $allChildDivs = $parentItem.find('.childdiv');


    $allChildDivs.each(function (index) {
        let $row = $(this);
        orderData.push({
            menuitem_id: getMenuId($row),
            orderindex: index + 1,
            parentmenu_id: parentId,
            is_child: true
        });
    });




    console.log("ORDER DATA →", orderData);
    if (orderData.length === 0) return;


    let $anyRow;

    if ($container.attr("id") === "accordionExample1") {
        $anyRow = $container.find(".menuitemsdiv .group").first();
    } else {
        $anyRow = $container.find(".childdiv").first();
    }

    if (!$anyRow.length) return;

    let label = $anyRow.find('input[placeholder="Navigation Lable"]').val();
    let path = $anyRow.find('input[placeholder="Path link"]').val();
    let updateBtn = $anyRow.find('#updatebtn');

    $.ajax({
        url: "/admin/website/menu/updatemenuitemsorder",
        type: "POST",
        data: {
            menuitem_id: getMenuId($anyRow),
            menu_name: label,
            parentmenu_id: parentId,
            urlpath: path,
            csrf: $("input[name='csrf']").val(),
            menu_typeid: updateBtn.data('typeid'),
            type: updateBtn.data('type'),
            webid: $('.templateid').attr('data-id'),
            orderData: JSON.stringify(orderData)
        },
        success: function () {
            console.log("Update successful");
             window.location.reload()
        }
    });
}



// // // Initialize sortable on root container - allows moving items amongst root groups or into child containers
// $('#accordionExample1').sortable({
//     connectWith: '.child-container',
//     handle: '.drag-handle',
//     items: '> .group',
//     placeholder: 'sortable-placeholder',
//     start: function (event, ui) {
//         ui.item.addClass('dragging');
//     },
//     stop: function (event, ui) {
//         ui.item.removeClass('dragging');
//         handleDrop(ui.item);
//     }
// }).disableSelection();

// // Initialize sortable on child containers - allows reordering within children only, no cross moves
// $('.child-container').sortable({
//     connectWith: '#accordionExample1',
//     handle: '.drag-handle',
//     items: '> .childdiv',
//     placeholder: 'sortable-placeholder',
//     start: function (event, ui) {
//         ui.item.addClass('dragging');
//     },
//     stop: function (event, ui) {
//         ui.item.removeClass('dragging');
//         handleDrop(ui.item);
//     }
// }).disableSelection();


// function handleDrop($item) {
//     let $droppedOn = $(document.elementFromPoint(event.clientX, event.clientY)).closest('.group, .childdiv');
//     if ($droppedOn.length && !$droppedOn.is($item)) {
//         let newParentId = $droppedOn.data('id');
//         let oldParentId = $item.parent().closest('.group, #accordionExample1').data('id');


//         let hasChildren = $item.find('.childdiv').length > 0;


//         if (hasChildren && newParentId !== oldParentId) {
//             $('#accordionExample1, .child-container').sortable('cancel');
//             showToast("Warning", "Items with children cannot be moved to another parent");
//             return;
//         }


//         if ($droppedOn.hasClass('childdiv')) {
//             $('#accordionExample1, .child-container').sortable('cancel');
//             showToast("Warning", "Child items cannot be dropped onto another item");
//             return;
//         }


//         if (typeof newParentId === "undefined" || newParentId === null) {
//             if ($item.hasClass('ms-[20px]')) {
//                 $item.removeClass('ms-[20px] childdiv');
//                 $('#accordionExample1').append($item);
//             }
//             updateMenuOrder($item, 0);
//         } else {

//             $('.div' + newParentId).append($item);
//             $item.addClass('ms-[20px] childdiv');
//             updateMenuOrder($item, newParentId);
//         }
//     } else {

//         if ($item.hasClass('ms-[20px]')) {
//             $item.removeClass('ms-[20px] childdiv');
//             $('#accordionExample1').append($item);
//         }
//         updateMenuOrder($item, 0);
//     }
// }



// function updateMenuOrder($item, parentId) {
//     if (!parentId) {
//         parentId = parseInt($('.parentmenu_id').val(), 10);

//     }

//     let $parentContainer = $item.parent();
//     let orderData = [];
//     $parentContainer.children('.group, .childdiv').each(function (index) {
//         let $el = $(this);
//         let id = $el.data('id');
//         let isChild = $el.hasClass('ms-[20px]');
//         orderData.push({
//             menuitem_id: id,
//             orderindex: index + 1,
//             parentmenu_id: parentId,
//             is_child: isChild
//         });
//     });

//     let label = $item.find('input[placeholder="Navigation Lable"]').val();
//     let path = $item.find('input[placeholder="Path link"]').val();
//     let updateBtn = $item.find('#updatebtn');

//     $.ajax({
//         url: "/admin/website/menu/updatemenuitemsorder",
//         type: "POST",
//         data: {
//             menuitem_id: $item.data('id'),
//             menu_name: label,
//             parentmenu_id: parentId,
//             urlpath: path,
//             csrf: $("input[name='csrf']").val(),
//             menu_typeid: updateBtn.data('typeid'),
//             type: updateBtn.data('type'),
//             webid: $('.templateid').attr('data-id'),
//             orderData: JSON.stringify(orderData)
//         },
//         success: function (data) {
//             console.log("Update successful");
//         }
//     });
// }



//search function//

$(document).on('keyup', '.searchbtn', function () {
    var keyword = $(this).val().trim().toLowerCase();
    var chkgrb = $(this).parents('.relative').siblings('.tab-content').find('.chk-group');
    var hasMatch = false;

    if (keyword === '') {
        chkgrb.removeClass('hidden');
        $(this).parents('.relative').siblings('#nodatafounddesign').addClass('hidden');
    } else {
        chkgrb.each(function () {
            var pname = $(this).find('p').text().trim().toLowerCase();
            if (pname.includes(keyword)) {
                $(this).removeClass('hidden');
                hasMatch = true;
            } else {
                $(this).addClass('hidden');
            }
        });


        if (hasMatch) {
            $(this).parents('.relative').siblings('#nodatafounddesign').addClass('hidden');
            $(this).parents('.relative').siblings('.tab-content').find('#nodatafounddesignn').removeClass('hidden')
        } else {
            $(this).parents('.relative').siblings('#nodatafounddesign').removeClass('hidden');
            $(this).parents('.relative').siblings('.tab-content').find('#nodatafounddesignn').addClass('hidden')
        }

    }
});



$(document).on('click', '#editbtnn', function () {

    menuname = $(this).attr('data-name')
    menutitle = $(this).attr('data-title')
    menudesc = $(this).attr('data-desc')
    menustatus = $(this).attr('data-status')
    var data = $(this).attr("data-id");
    $("#menuform").attr("name", "editmenu").attr("action", "/admin/website/menu/updatemenu/?webid=" + $('.templateid').attr('data-id'))
    $("input[name=menu_name]").val(menuname.trim());
    $("input[name=menu_title]").val(menutitle.trim());
    $("textarea[name=menu_desc]").val(menudesc.trim());



    if ((menuname.trim() == "Aside") || (menuname.trim() == "Headers") || (menuname.trim() == "Footers")||(menuname.trim()=="SERVICES")||(menuname.trim()=="SOLUTIONS")||(menuname.trim()=="EXPERTISE")) {

        $('#menu_name').addClass('pointer-events-none opacity-50')
    } else {
        $('#menu_name').removeClass('pointer-events-none opacity-50')
    }
    if (menustatus == 1) {
        $('.menustatus').prop('checked', true).val('1');
    } else {
        $('.menustatus').prop('checked', false).val('0');
    }

    $("#savemenubtn").text(languagedata.update)
    $('#modalTitleIdd').text(languagedata.update + languagedata.Menu.menu)
    $("#menu_name-error").hide()
    $("#menu_desc-error").hide()
    $("input[name=menu_id]").val(data);
})


$(document).on('keyup', '.labelname', function () {

    console.log("checkdfffdf",$(this).val())

    if ($(this).val().trim().length >= 300) {

        $(this).removeClass('mb-[24px]')

        $(this).siblings('#lablename-error').removeClass('hidden').addClass('mb-[24px]')

    } else {
        $(this).addClass('mb-[24px]')
        $(this).siblings('#lablename-error').addClass('hidden').removeClass('mb-[24px]')
    }
})

$(document).on('click', '.opentab', function () {


    $('#customumenurl-error').addClass('hidden')
    $('#custommenuname-error').addClass('hidden')

    var activetab = $(this).siblings('.accordion-collapse').children('.accordion-body').children('.tab-content').find('.tab-pane.active');
    var checkboxes = activetab.find('.channelnameinput');
    var $btn = $(this).siblings('.accordion-collapse').children('.accordion-body').find('.courseallselect');



    if (checkboxes.length && checkboxes.filter(':checked').length === checkboxes.length) {
        $btn.text(languagedata.Menu.deselectall);

    } else {


        $btn.text(languagedata.Menu.selectall);
    }
});


$(document).on('click', '.courseallselect', function () {

    var tabcontent = $(this).parent().siblings('.tab-content');
    var activetab = tabcontent.find('.tab-pane.active');
    var checkboxes = activetab.find('.channelnameinput');


    var allChecked = checkboxes.length && checkboxes.filter(':checked').length === checkboxes.length;

    if (allChecked) {
        checkboxes.prop('checked', false);
        $(this).text(languagedata.Menu.selectall);
    } else {
        checkboxes.prop('checked', true);
        $(this).text(languagedata.Menu.deselectall);
    }
});


$(document).on('click', '.channelnameinput', function () {

    var $tabPane = $(this).closest('.tab-pane');
    var $checkboxes = $tabPane.find('.channelnameinput');
    var $courseAllSelectBtn = $tabPane.closest('.tab-content').siblings('.nav').find('.courseallselect');
    var allChecked = $checkboxes.length && $checkboxes.filter(':checked').length === $checkboxes.length;

    if ($courseAllSelectBtn.length) {
        if (allChecked) {
            $courseAllSelectBtn.text(languagedata.Menu.deselectall);
        } else {
            $courseAllSelectBtn.text(languagedata.Menu.selectall);
        }
    }
});


$(document).on('click', '.recentchannel', function () {

    $('#allchannel').removeClass('show active')
})

$(document).on('click', '.recentcourse', function () {

    $('#allcourse').removeClass('show active')
})
$(document).on('click', '.recentform', function () {

    $('#allforms').removeClass('show active')
})
$(document).on('click', '.recentcategory', function () {

    console.log("checkconditon")

    $('#allcategory').removeClass('show active')
})
$(document).on('click', '.recentlistingCategory', function () {

    $('#alllistingCategory').removeClass('show active')
})
$(document).on('click', '.recentpages', function () {

    $('#allpages').removeClass('show active')
})

$(document).on('click', '.cancelmenuitem', function () {

    $('#menu_name-error').hide()
    $('#urlpath-error').hide()
    $('input[name="lang"]:checked').prop('checked', false);
    $('.customUrlInput').addClass('hidden')
    $('.entriesDropdown').addClass('hidden')
    $('.pagesDropdown').addClass('hidden');
    $('.listingsDropdown').addClass('hidden');
    $('.categoryDropdown').addClass('hidden')
    $('.channelDropdown').addClass('hidden')
    $('.navpath').val("")
    $('input[name="menu_name"]').val("")
    $('input[name="meta_title"').val("")
    $('input[name="meta_description"').val("")
    $('input[name="meta_keywords"').val("")
    $('.listingsDropdownMenu li').each(function () {

        $(this).find('input[type="checkbox"]').prop('checked', false);
    })
    $('.categoryDropdownMenu li').each(function () {


        $(this).find('input[type="checkbox"]').prop('checked', false);

    })
})

$(document).on('click', '.addcoursemenu', function () {
    var parentmenuid = $('#menu_id').val();
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

    if (!exists) {
        $.ajax({
            url: "/admin/website/menu/createmenuitems",
            type: "POST",
            async: false,
            data: {
                "menu_name": "All Courses",
                "menu_id": parentmenuid,
                "urlpath": "/courses",
                csrf: $("input[name='csrf']").val(),
                "menu_typeid": "",
                "type": "courses",
                "webid": $('.templateid').attr('data-id'),
            },
            dataType: "json",
            cache: false,
            success: function (result) {

                location.reload();
            }
        })

    }
})

$(document).ready(function () {

    $('input[name="lang"]').change(function () {
        selectedLabel = $(this).next('label').text().trim();
        $('.navpath').val("")
        if (selectedLabel === "Custom URL") {
            $('.customUrlInput').removeClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden');
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.menutype').val("custom_url")
            $('.channelDropdown').addClass('hidden')
        } //  else if (selectedLabel === "Entries") {
        //     $('.customUrlInput').addClass('hidden')
        //     $('.entriesDropdown').removeClass('hidden')
        //     $('.pagesDropdown').addClass('hidden')
        //     $('.listingsDropdown').addClass('hidden')
        //     $('.categoryDropdown').addClass('hidden')
        // }
        else if (selectedLabel === "Pages") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').removeClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.menutype').val("pages")
        } else if (selectedLabel === "Listings") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.listingsDropdown').removeClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.menutype').val("listings")
        } else if (selectedLabel === "Categories") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').removeClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.menutype').val("categories")
            $('.navpath').val("/categories")
        } else if (selectedLabel == "None") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.navpath').val("")
            $('.menutype').val("none")

        } else if (selectedLabel == "Channels") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.channelDropdown').removeClass('hidden')
            $('.menutype').val("channels")
        }
    });
});

$(document).on('keyup', '.customurl', function () {

    url = $(this).val().trim()

    if (url != "") {

        $('#urlpath-error').addClass('hidden')
    } else {

        $('#urlpath-error').removeClass('hidden')
    }

    $('.navpath').val(url)

})

$(document).on('click', '.entryslug', function () {

    url = $(this).attr('data-slug').trim()

    $('.navpath').val("/categories/" + url)
})

$(document).on('click', '.pageslug', function () {

    url = $(this).attr('data-slug').trim()

    $('.navpath').val("/pages/" + url)

    $('.selectpage').text(url)

    $('.menu_typeid').val($(this).attr('data-id'))

})
$(document).on('click', '.channelname', function () {

    url = $(this).attr('data-slug').trim()

    $('.navpath').val("/channel/" + url)

    $('.selectchannel').text(url)
})
$(".searchcatlists").keyup(function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase()

    console.log("checksearch")


    $(".catrgory-dropdownlist a").filter(function () {
        var isVisible = $(this).text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible);
        if (isVisible) found = true;
    })
    if (found) {
        $('.noData-foundWrapper').hide();
    } else {
        $('.noData-foundWrapper').show();
    }
})

$(".searchpagelist").keyup(function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase()


    $(".catrgory-dropdownlistt a").filter(function () {
        var isVisible = $(this).text().toLowerCase().indexOf(searchTerm) > -1;
        $(this).toggle(isVisible);
        if (isVisible) found = true;
    })
    if (found) {
        $('.noData-foundWrapperr').hide();
    } else {
        $('.noData-foundWrapperr').show();
    }
})
$(".searchlistinglist").keyup(function () {
    var searchTerm = $(this).val().trim().toLowerCase();
    var found = false;
    var dropdownContainer = $(this).closest('.dropdown');

    dropdownContainer.find("ul li").each(function () {
        var itemText = $(this).find('label').text().toLowerCase();
        if (itemText.indexOf(searchTerm) > -1) {
            $(this).show();
            found = true;
        } else {
            $(this).hide();
        }
    });
    if (found) {
        dropdownContainer.find('.noData-foundWrapperr').hide();
    } else {
        dropdownContainer.find('.noData-foundWrapperr').show();
    }
});
$(".searchchannellist").keyup(function () {
    var searchTerm = $(this).val().trim().toLowerCase();
    var found = false;
    var dropdownContainer = $(this).closest('.dropdown');

    dropdownContainer.find("ul li").each(function () {
        var itemText = $(this).find('span').text().toLowerCase();
        if (itemText.indexOf(searchTerm) > -1) {
            $(this).show();
            found = true;
        } else {
            $(this).hide();
        }
    });
    if (found) {
        dropdownContainer.find('.noData-foundWrapperr').hide();
    } else {
        dropdownContainer.find('.noData-foundWrapperr').show();
    }
});

$(document).ready(function () {
    $('.listingsDropdownMenu').on('click', function (e) {
        e.stopPropagation();
    });
    $('.categoryDropdownMenu').on('click', function (e) {
        e.stopPropagation();
    });
});

$(".searchcategorylists").on("keyup", function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase();

    $(".catrgory-list").each(function () {
        var $categoryItem = $(this);

        var categoryText = $categoryItem.find("span").map(function () {
            return $(this).text().toLowerCase();
        }).get().join(" ");

        var isVisible = categoryText.indexOf(searchTerm) > -1;
        $categoryItem.toggle(isVisible);

        if (isVisible) {
            found = true;
        }
    });

    if (found) {
        $(".noData-foundcategory").hide();
    } else {
        $(".noData-foundcategory").show();
    }
});

$(".searchlslists").on("keyup", function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase();

    $(".listingcat").each(function () {
        var $categoryItem = $(this);

        var categoryText = $categoryItem.find("span").map(function () {
            return $(this).text().toLowerCase();
        }).get().join(" ");

        var isVisible = categoryText.indexOf(searchTerm) > -1;
        $categoryItem.toggle(isVisible);

        if (isVisible) {
            found = true;
        }
    });

    if (found) {
        $(".noData-foundcategorylist").hide();
    } else {
        $(".noData-foundcategorylist").show();
    }
});

$('.selectcheckbox').on('click', function () {
    var parentLi = $(this).closest('li.catrgory-list');
    var lastSpan = parentLi.find('span').last();
    var dataId = String(lastSpan.attr('data-id'));


    categoryarr = categoryarr.flat();

    if ($(this).prop('checked')) {
        if (!categoryarr.includes(dataId)) {
            categoryarr.push(dataId);
        }
    } else {

        categoryarr = categoryarr.filter(function (id) {

            return id !== dataId;
        });

    }

});





$('.imgupload').on('change', function (e) {
    var file = e.target.files[0];
    if (!file) return;


    if (file.type !== 'image/svg+xml' && !file.name.endsWith('.svg')) {

        $(this).val('');
        $('#svgHidden').val('');
        return;
    }

    var reader = new FileReader();
    reader.onload = function (event) {
        $('#svgHidden').val(event.target.result);
        $(e.target).closest('.imgupldiv').addClass('hidden');
        $('#ImageName').text(file.name);
        $('#imageRemoveDiv').removeClass('hidden');
        $('.svgins').addClass('hidden')
        $('.uploadbtn').addClass('hidden')
        $('.ImageNamediv').removeClass('hidden')
    };
    reader.readAsDataURL(file);
});

// Handle remove image button
$('#imageRemoveDiv ').on('click', function () {
    $('#svgHidden').val('');
    $('.imgupload').val('');
    $('#ImageName').text('');
    $('#imageRemoveDiv').addClass('hidden');
    $('.imgupldiv').removeClass('hidden');
    $('.svgins').removeClass('hidden')
    $('.uploadbtn').removeClass('hidden')
    $('.ImageNamediv').addClass('hidden')
});

$('.deleteimg').on('click', function () {

    $(this).closest('.imageRemoveDiv').addClass('hidden');
    $(this).closest('.imageRemoveDiv').siblings('.imgupldiv').removeClass('hidden');
    $(this).closest('.imageRemoveDiv').siblings('.imgupldiv').find('.svgval').val('');
    $(this).siblings('.svgDelete').val('1')
});


$(document).on('change', '.imguploadd', function (e) {
    var file = e.target.files[0];
    if (!file) return;


    if (file.type !== 'image/svg+xml' && !file.name.endsWith('.svg')) {
        $(this).val('');
        $(this).closest('.imgupldiv').find('.svgval').val('');
        return;
    }

    var reader = new FileReader();
    var uploaddiv = $(this);
    reader.onload = function (event) {

        uploaddiv.closest('.imgupldiv').find('.svgval').val(event.target.result);
        uploaddiv.closest('.imgupldiv').addClass('hidden');
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').removeClass('hidden');
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').find('.ImageName').text(file.name);
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').find('.deleteimg').removeClass('hidden')
        uploaddiv.closest('.imgupldiv').find('.svgins').addClass('hidden');
        uploaddiv.closest('.imgupldiv').siblings('.imageRemoveDiv').find('.svgDelete').val('')
    };

    reader.readAsDataURL(file);
});


//Menu edit button functionality//

$(document).on('click', '.menuitemeditbtn', function () {


    menuid = $(this).attr('data-id')
    webid = $('.webid').val()

    $.ajax({
        url: "/admin/website/menu/editmenuitem/" + menuid,
        type: "GET",
        dataType: "json",
        success: function (result) {
            console.log("user", result);
            $('#modalTitleId').text('Update Menu Item')
            $('#menuitemssavebtn').text('Update item')
            $('#menuitemform').attr('action', '/admin/website/menu/updatemenuitems?webid=' + webid)
            $('input[name="menu_name"]').val(result.Name)
            $('.menuitem_id').val(menuid)
            $('#ImageName').text(result.ImageName)
            $('#separatewindow').val(result.SeparateWindow)
            if (result.SeparateWindow == 1) {
                $('#separatewindow').prop('checked', true)
            } else {
                $('#separatewindow').prop('checked', false)
            }
            $('.navpath').val(result.UrlPath)
            $('input[name="meta_title"').val(result.MetaTitle)
            $('input[name="meta_description"').val(result.MetaDescription)
            $('input[name="meta_keywords"').val(result.MetaKeywords)
            $('.parentmenu_id').val(result.ParentId)
            $('.menu_typeid').val(result.TypeId)
            $('.menu_typeid').val(result.TypeId)

            if (result.ImageName != "") {

                $('.svgins').addClass('hidden')
                $('.uploadbtn').addClass('hidden')
                $('#imageRemoveDiv').removeClass('hidden')
                $('.ImageNamediv').removeClass('hidden')
            }
            $('.menutype').val(result.Type)

            if (result.Type == "none") {

                $('#radionone').prop('checked', true)
                selectedLabel = "None"

            } else if (result.Type == "pages") {

                $('#radioPages').trigger('click')
                $('.navpath').val(result.UrlPath)
                pagename = result.UrlPath.substring(result.UrlPath.indexOf("/pages/") + 7);
                $('.selectpage').text(pagename)
            } else if (result.Type == "categories") {
                $('#radioCategories').trigger('click')
                $('.navpath').val(result.UrlPath)
                var categoryIdsArray = result.CategoryIds.split(',');
                categoryarr.push(categoryIdsArray)
                $('.categoryDropdownMenu li').each(function () {
                    var dataId = $(this).find('span').last().attr('data-id');


                    if (categoryIdsArray.includes(dataId)) {

                        $(this).find('input[type="checkbox"]').prop('checked', true);
                    } else {

                        $(this).find('input[type="checkbox"]').prop('checked', false);
                    }
                })
            } else if (result.Type == "listings") {
                $('#radioListings').trigger('click')
                $('.navpath').val(result.UrlPath)
                var listingIdsArray = result.ListingsIds.split(',');

                $('.listingsDropdownMenu li').each(function () {
                    var dataId = $(this).attr('data-id');


                    if (listingIdsArray.includes(dataId)) {

                        $(this).find('input[type="checkbox"]').prop('checked', true);
                    } else {

                        $(this).find('input[type="checkbox"]').prop('checked', false);
                    }
                });



            } else if (result.Type == "custom_url") {
                $('#radioCustomUrl').trigger('click')
                $('.navpath').val(result.UrlPath)
            } else if (result.Type == "channels") {
                $('#radioChannels').trigger('click')
                $('.navpath').val(result.UrlPath)
                channelname = result.UrlPath.substring(result.UrlPath.indexOf("/channel/") + 9);
                console.log(channelname, "channelname")
                $('.selectchannel').text(channelname)
            }

        }
    })

})


$(document).on('click', '.addmenuitembtn', function () {
    webid = $('.webid').val()
    $('#modalTitleId').text('Add Menu Item')
    $('#menuitemssavebtn').text('Add Item')
    $('#menuitemform').attr('action', '/admin/website/menu/createmenuitems?webid=' + webid)
    selectedLabel = ""
    $('.ImageNamediv').addClass('hidden')
    $('.uploadbtn').removeClass('hidden')
    $('.selectpage').text('Select Page')
    $('.selectchannel').text('Select Channel')

})


$(".menuitemname").on('keyup', function () {
    // Find the error label for this field and hide it
    var errorLabel = $("label[for='menu_name']");
    errorLabel.hide();
});


$(document).on('click', '.addcustommenu', function () {
    var parentmenuid = $('#menu_id').val();
    var url = $('.customumenurl').val()
    menuname = $('.custommenuname').val()

    if (url == "") {

        $('#customumenurl-error').removeClass('hidden')

    }
    if (menuname == "") {
        $('#custommenuname-error').removeClass('hidden')
    }
    if (url != "" && menuname != "") {

        $('#customumenurl-error').addClass('hidden')
        $('#custommenuname-error').addClass('hidden')

        $.ajax({
            url: "/admin/website/menu/createmenuitems?webid=" + $('.webid').data('id'),
            method: "POST",
            dataType: "json",
            async: false,
            data: {
                menu_name: menuname,
                menu_id: parentmenuid,
                urlpath: url,
                menu_typeid: "",
                csrf: $("input[name='csrf']").val(),
                type: "custom_url",
                webid: $('.webid').data('id')
            },
            success: function () {
                location.reload();
            }
        });
    }

})


$(document).on('keyup', '.custommenuname', function () {

    Name = $(this).val().trim()

    if (Name != "") {

        $('#custommenuname-error').addClass('hidden')
    } else {

        $('#custommenuname-error').removeClass('hidden')
    }

})

$(document).on('keyup', '.customumenurl', function () {

    Name = $(this).val().trim()

    if (Name != "") {

        $('#customumenurl-error').addClass('hidden')
    } else {

        $('#customumenurl-error').removeClass('hidden')
    }



})


$(document).on('mousedown', '.quill-editor-container', function (e) {

    e.stopPropagation();

    const quill = $(this).data('quill');
    if (!quill) return;

    // Let browser place cursor first
    setTimeout(() => {
        if (!quill.hasFocus()) {
            quill.focus();
        }
    }, 0);
});


