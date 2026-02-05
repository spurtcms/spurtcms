var languagedata

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })

    $('.Content').addClass('checked');

})


//delete func//

$(document).on('click', '#deletebtn', function () {



    var pageid = $(this).attr("data-id")
    console.log("kkk", pageid)
    $("#content").text("Are you sure you want to remove this page")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
    templateid = $('.templateid').val()



    if (pageno == null) {
        $('#delid').attr('href', "/admin/website/pages/deletepage/" + pageid + "?webid=" + templateid);

    } else {
        $('#delid').attr('href', "/admin/website/pages/deletepage/" + pageid + "?page=" + pageno & "?webid=" + templateid);

    }


    $(".deltitle").text("Delete Page?")
      $('.deldesc').text("Are you sure you want to remove this page")
    $('.delname').text($(this).parents('tr').find('td:eq(0) a').text())
      $('#delid').show()

  $('#dltCancelBtn').text("cancel")

})


function PageStatus(id) {
    $('#cb' + id).on('change', function () {
        console.log("printf");
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb' + id).val();

    console.log("check", isactive, id)

    $.ajax({
        url: '/admin/website/pages/pagestatuschange',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {

                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Page Updated Successfully</p ></div ></div ></li></ul> `;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            } else {

                notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds

            }
        }
    });
}

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')
    window.location.href = "/admin/website/pages/" + $('.templateid').val()
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            var value = $(".search").val();
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

//Filter dropdown function//

$(document).on('click', '.statuscls', function () {

    statusval = $(this).text().trim()
    $(".filterleveldropdown").removeClass("show")
    $('.slctstatus').text(statusval)
    $('#statusid').val(statusval)
})

//drag and drop functionality//




// dec 25


const ROOT_PARENT_ID = 0;
const CHILD_MARGIN = 'ms-[20px]';
const SUBCHILD_MARGIN = 'ms-[40px]';
const MAX_LEVEL = 3;


$('#accordionPanelsStayOpenExample').sortable({
    handle: '.drag-handle',
    items: '.sortable-parent',
    connectWith: '.child-container,.sortable-parent', // ✅ YOUR WORKING CONNECTWITH
    placeholder: 'sortable-placeholder',
    start: function (e, ui) { ui.item.addClass("dragging"); },
    receive: function (event, ui) {
        const dragged = ui.item;
        const hasChildren = dragged.find('.child-container .childdiv,.sortable-parent').length > 0;
        const dropTarget = $(this);
        const isDroppingAsChild = dropTarget.hasClass("child-container");

        if (hasChildren && isDroppingAsChild) {
            $(ui.sender).sortable('cancel');
        }
    },
    stop: function (e, ui) {
        ui.item.removeClass('dragging');
        handleRootDrop(e, ui.item); // ✅ YOUR WORKING HANDLER
    }
}).disableSelection();

// 🔹 YOUR ORIGINAL CHILD SORTABLE (KEEP WORKING)
$('.child-container').sortable({
    handle: '.drag-handle',
    items: '> .childdiv',
    connectWith: '#accordionPanelsStayOpenExample, .sortable-parent .child-container, .child-container',
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
// 🔹 SUBCHILD CONTAINER SORTABLE ONLY (Level 3)
$('.subchild-container').sortable({
    handle: '.drag-handle',
    items: '> .childdiv',
    connectWith: '#accordionPanelsStayOpenExample, .child-container ,.subchild-container',
    placeholder: 'sortable-placeholder ms-[40px]',
    start: function(e, ui) {
        ui.item.addClass('dragging');
        ui.placeholder.height(ui.item.height());
    },
    receive: function(e, ui) {
        // BLOCK 4th level & invalid drops
        const $item = ui.item;
        if ($item.hasClass('sortable-parent') || 
            $item.find('.child-container, .subchild-container').length > 0) {
            $(ui.sender).sortable('cancel');
            return;
        }
        const parentId = parseInt($(this).attr('data-parent'), 10);
        $item.addClass('ms-[40px] childdiv').attr('data-parent', parentId);
    },
    stop: function(e, ui) {
        ui.item.removeClass('dragging');
        handleSubChildDrop(e, ui.item);
    }
}).disableSelection();

// 🔹 NEW: SUBCHILD SORTABLE (LEVEL 3 ONLY - NO DEEPER)
function initializeSubChildSortable($subContainer) {
    if ($subContainer.hasClass('ui-sortable')) return;
    
    $subContainer.sortable({
        handle: '.drag-handle',
        items: '> .childdiv',
        connectWith: '#accordionPanelsStayOpenExample, .child-container', // ✅ NO subchild-to-subchild
        placeholder: 'sortable-placeholder ms-[40px]',
        start: function(e, ui) {
            ui.item.addClass('dragging');
            ui.placeholder.height(ui.item.height());
        },
        receive: function(e, ui) {
            // ✅ BLOCK 4th LEVEL & INVALID DROPS
            const $item = ui.item;
            if ($item.hasClass('sortable-parent') || $item.find('.child-container, .subchild-container').length > 0) {
                $(ui.sender).sortable('cancel');
                return;
            }
            $item.addClass('ms-[40px] childdiv').attr('data-parent', $subContainer.attr('data-parent'));
        },
        stop: function(e, ui) {
            ui.item.removeClass('dragging');
            handleSubChildDrop(e, ui.item);
        }
    }).disableSelection();
}

// 🔹 YOUR ORIGINAL FUNCTIONS (UNCHANGED - WORKING)
function handleRootDrop(event, $wrapper) {
    const $group = $wrapper.find('.sortable-parent').first();
    const parentId = getRootParentId();
    const draggedHasChildren = $wrapper.find('.child-container').length > 0;

    let $hit = $(document.elementFromPoint(event.clientX, event.clientY));
    let $target = $hit.closest('.sortable-parent').length ? $hit.closest('.sortable-parent') :
        $hit.closest('.childdiv').length ? $hit.closest('.childdiv') :
            $hit.closest('.child-container').length ? $hit.closest('.child-container').closest('.sortable-parent') :
                $hit.closest('.sortable-parent').length ? $hit.closest('.sortable-parent') : $();

    if ($target.is($wrapper)) $target = $();

    if ($hit.closest('.child-container').length) {
        if (draggedHasChildren) {
            moveToRoot($wrapper, $group, parentId);
            return;
        }
        let $childContainerTarget = $hit.closest('.child-container');
        makeItemChildOf($wrapper, $childContainerTarget.closest('.sortable-parent').length ? $childContainerTarget.closest('.sortable-parent') : $childContainerTarget);
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
    $group.removeClass('ms-[20px] childdiv ms-[40px]');
    $group.attr('data-parent', parentId);

    if (!$wrapper.parent().is('#accordionPanelsStayOpenExample')) {
        $('#accordionPanelsStayOpenExample').append($wrapper);
    }

    updateOrderForContainer($('#accordionPanelsStayOpenExample'), parentId);
}

function makeItemChildOf($sourceWrapper, $targetWrapper) {
    let itemId = parseInt($sourceWrapper.attr('data-id'), 10);
    let targetId = parseInt($targetWrapper.data('id') || $targetWrapper.attr('data-id'), 10);

    let $childContainer = $targetWrapper.find('.child-container').first();

    if (!$childContainer.length) {
        $childContainer = $('<div class="child-container"></div>');
        $childContainer.addClass('ms-[20px]');
        $childContainer.attr('data-parent', targetId);

        let $accordionItem = $targetWrapper.find('.accordion-item').first();
        let $deleteBtnWrapper = $targetWrapper.find('div').last();

        if ($deleteBtnWrapper.length) {
            $targetWrapper.append($childContainer);
        } else {
            $accordionItem.after($childContainer);
        }

        initializeChildSortable($childContainer);
    }

    $sourceWrapper.addClass('ms-[20px] childdiv');
    $sourceWrapper.attr('data-parent', targetId);
    $sourceWrapper.appendTo($childContainer);

    updateOrderForContainer($childContainer, targetId);
}

// 🔹 FIXED: YOUR CHILD SORTABLE INIT (ADDED SUBCHILD SUPPORT)
function initializeChildSortable($childContainer) {
    if ($childContainer.hasClass('ui-sortable')) return;

    $childContainer.sortable({
        handle: '.drag-handle',
        items: '> .childdiv',
        connectWith: '#accordionPanelsStayOpenExample, .sortable-parent .child-container, .child-container',
        placeholder: 'sortable-placeholder ms-[20px]',
        start: function (e, ui) {
            ui.item.addClass('dragging');
            ui.placeholder.height(ui.item.height());
        },
        stop: function (e, ui) {
            handleChildDrop(e, ui.item);
        },
        receive: function (e, ui) {
            let $item = ui.item.find('.sortable-parent').first();
            if ($item.length && !$item.hasClass('childdiv')) {
                $item.addClass('ms-[20px] childdiv');
                $item.attr('data-parent', $(this).attr('data-parent'));
            }
        }
    }).disableSelection();
}

// 🔹 YOUR ORIGINAL CHILD DROP (ENHANCED FOR SUBCHILD)
function handleChildDrop(event, $child) {
    let $hit = $(document.elementFromPoint(event.originalEvent.clientX, event.originalEvent.clientY));
    let $targetChild = $hit.closest('.childdiv').not($child);

   
    if ($targetChild.length) {
        makeItemSubChildOf($child, $targetChild);
        return;
    }

   
    if ($child.closest('#accordionPanelsStayOpenExample').length && !$child.closest('.child-container').length) {
        makeChildIndependent($child);
        updateOrderForContainer($('#accordionPanelsStayOpenExample'), getRootParentId());
        return;
    }

  
    let $immediateParent = $child.parent();
    if ($immediateParent.is('.child-container')) {
        let $actualParent = $immediateParent.closest('.sortable-parent');
        let parentId = parseInt($actualParent.attr('data-id') || $actualParent.data('id'), 10);

        if (isNaN(parentId)) return;

        $child.addClass('ms-[20px] childdiv').attr('data-parent', parentId);
        $immediateParent.attr('data-parent', parentId);
        updateOrderForContainer($immediateParent, parentId);
        return;
    }
}

// 🔹 NEW: SUBCHILD DROP HANDLER (STRICTLY LEVEL 3)
function handleSubChildDrop(event, $subChild) {
    // ROOT check
    if ($subChild.closest('#accordionPanelsStayOpenExample').length && 
        !$subChild.closest('.child-container, .subchild-container').length) {

            console.log("checkindepee",$subChild)
        makeSubChildIndependent($subChild);
        return;
    }
    
    console.log("checksubci")
    // Block subchild → subchild (no 4th level)
    let $hit = $(document.elementFromPoint(event.originalEvent.clientX, event.originalEvent.clientY));
    let $target = $hit.closest('.childdiv').not($subChild);
    if ($target.length) {
        makeSubChildIndependent($subChild);
        return;
    }
    
    // Normal reorder
    let $parent = $subChild.parent();
    if ($parent.is('.subchild-container')) {
       let parentId = parseInt(
  $parent.closest('.menuitemsdiv, .childdiv').attr('data-id'),
  10
);
        $subChild.addClass('ms-[40px] childdiv').attr('data-parent', parentId);
         console.log("checksubci2")
        updateOrderForContainer($parent, parentId);
    }
}

function makeSubChildIndependent($subChild) {

    console.log("makee")
    let pageId = $subChild.attr('data-id');
    $subChild.removeClass('ms-[40px] childdiv').attr('data-parent', 0);
    let $wrapper = $(`<div class="sortable-parent" data-id="${pageId}"></div>`);
    $wrapper.append($subChild);
    $('#accordionPanelsStayOpenExample').append($wrapper);
   updateOrderForContainer($('#accordionPanelsStayOpenExample'), 0);

}


function makeItemSubChildOf($sourceChild, $targetChild) {

  
     $targetChild = $targetChild.closest('.childdiv');
    if (!$targetChild.length) return;

  
    if ($targetChild.parent().hasClass('subchild-container')) {
       
        return;
    }
    console.log("check2level");

    let targetId = parseInt($targetChild.attr('data-id'), 10);
    if (isNaN(targetId)) return;

   
    let $subContainer = $targetChild.children('.subchild-container');

    if (!$subContainer.length) {
        $subContainer = $('<div class="subchild-container ms-[40px]"></div>')
            .attr('data-parent', targetId)
            .appendTo($targetChild);

        initializeSubChildSortable($subContainer);
    }

 
    $sourceChild
        .removeClass('ms-[20px]')
        .addClass('ms-[40px] childdiv')
        .attr('data-parent', targetId)
        .appendTo($subContainer);

  
    updateOrderForContainer($subContainer, targetId);
}



// 🔹 YOUR ORIGINAL FUNCTIONS (UNCHANGED)
function initializeAllChildContainers() {
    $('.child-container').each(function () {
        if (!$(this).hasClass('ui-sortable')) {
            initializeChildSortable($(this));
        }
    });
}

function makeChildIndependent($child) {
    let pageid = $child.attr('data-id');
    $child.removeClass('ms-[20px] childdiv ms-[40px]');
    $child.attr('data-parent', getRootParentId());
    let $wrapper = $('<div class="sortable-parent" data-id="' + pageid + '"></div>');
    $wrapper.append($child);
    $('#accordionPanelsStayOpenExample').append($wrapper);
 
}

// function makeSubChildIndependent($subChild) {

//     console.log("child1")
//     makeChildIndependent($subChild);
// }

function getRootParentId() {
    return parseInt(0);
}

function updateOrderForContainer($container, parentId) {
    let orderData = [];

    if ($container.attr('id') === 'accordionPanelsStayOpenExample') {
        $container.children('.sortable-parent').each(function (index) {
            orderData.push({
                menuitem_id: parseInt($(this).attr('data-id'), 10),
                orderindex: index + 1,
                parentmenu_id: 0,
                is_child: false
            });
        });
    } else {
        $container.children('.childdiv').each(function (index) {
            orderData.push({
                menuitem_id: parseInt($(this).attr('data-id'), 10),
                orderindex: index + 1,
                parentmenu_id: parentId,
                is_child: true
            });
        });
    }

    console.log('ORDER DATA →', orderData);
    if (!orderData.length) return;

    $.ajax({
        url: '/admin/website/pages/updatepagesorder',
        type: 'POST',
        data: {
            parentmenu_id: parentId,
            webid: $('.templateid').data('id'),
            csrf: $('input[name="csrf"]').val(),
            orderData: JSON.stringify(orderData)
        },
        success: function () {
            console.log('Order saved');
             window.location.reload()
        }
    });
}

// 🔹 MASTER INIT (INITIALIZES SUBCHILD TOO)
function initializeAllContainers() {
    initializeAllChildContainers();
    $('.subchild-container').each(function() {
        if (!$(this).hasClass('ui-sortable')) {
            initializeSubChildSortable($(this));
        }
    });
}

// 🔹 FINAL READY (YOUR STRUCTURE)
$(document).ready(function () {
    initializeAllContainers();
});


$(document).on('click', '.delete-toast-btn', function () {
  $('.deltitle').text("Delete Page")
  $('.deldesc').text("This Page contains Sub Page and cannot be deleted.")
  $('#delid').hide()
  $('.delname').text($(this).parents('.f-chn').find('.chname').text())
  $('#dltCancelBtn').text("Ok")
})