var languagedata

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

})
//savebutton function
$(document).on('click', '#menuitemssavebtn', function (e) {
    e.preventDefault(); //


    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        menu_id = $("#menu_id").val()
        menu_name = $("#menu_name").val()

        $.ajax({
            url: "/admin/website/menu/checkmenuname",
            type: "POST",
            async: false,
            data: { "menu_name": value, "webid": $(".templateid").val(), "menu_id": menu_id, csrf: $("input[name='csrf']").val() },
            datatype: "json",
            caches: false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim() != "true"
    })


    let selectedIds = $('.selected-listings:checked').map(function () {
        return $(this).data('id');
    }).get();

    let joinedIds = selectedIds.join(',');

    if (joinedIds!=""){

        $(".navpath").val("/listings/")
    }

    $('#listingsids').val(joinedIds);


    $("#menuitemform").validate({
        rules: {
            menu_name: {
                required: true,
                space: true,
                duplicatename: true,
                maxlength: 15,

            },
            urlpath: {
                required: true,
                // maxlength: 30,
                space: true,
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
                // maxlength: "*" + languagedata.Menu.pathlimiterr
            },
        }
    });


    var formcheck = $("#menuitemform").valid();
    if (formcheck == true) {
        $('#menuitemform')[0].submit();
    }

    return false

})

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


        var exists = false;
        $('#accordionExample1 .accordion-item').each(function () {
            var $updateBtn = $(this).find('#updatebtn');
            if (
                $updateBtn.data('type') == type &&
                $updateBtn.data('typeid') == id
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
                    "menu_name": menuname,
                    "menu_id": parentmenuid,
                    "urlpath": "",
                    csrf: $("input[name='csrf']").val(),
                    "cid": id,
                    "type": type,
                    "webid": $('.templateid').attr('data-id'),
                },
                dataType: "json",
                cache: false,
                success: function (result) {

                    location.reload();
                    // $('.noitemsdiv').addClass('hidden');
                    // var collapseId = "collapse" + result.Id;
                    // var str = `
                    //     <div class="flex items-start gap-[10px] mt-[8px] first-of-type:mt-0 w-full group ui-draggable ui-droppable" data-id="${result.Id}">
                    //         <div class="drag-handle p-[16px_0] min-w-[14px] opacity-[0.5] group-hover:opacity-100 cursor-pointer">
                    //             <img src="/public/img/drag.svg" alt="">
                    //         </div>
                    //         <div class="accordion-item grow max-w-[586px]">
                    //             <button class="bg-[#F0F0F0] p-[12px_20px] rounded-[8px] w-full border-[#EBEBEB] brder border-solid justify-between flex items-center mb-[8px] h-[48px] collapsed"
                    //                 type="button" data-bs-toggle="collapse"
                    //                 data-bs-target="#${collapseId}" aria-expanded="false"
                    //                 aria-controls="${collapseId}">
                    //                 <p class="text-base font-normal leading-[20px] text-[#252525] text-start itemname">
                    //                     ${menuname}</p>
                    //                 <div class="flex items-center gap-[10px]">
                    //                     <p class="text-[13px] font-light text-[#757575] leading-[16px]">
                    //                         ${menuname}</p>
                    //                     <img src="/public/img/accord-arrow.svg" alt="">
                    //                 </div>
                    //             </button>
                    //             <div id="${collapseId}"
                    //                 class="accordion-collapse [&amp;.collapse]:!visible collapse"
                    //                 data-bs-parent="#accordionExample1" style="">
                    //                 <div class="accordion-body p-[24px_14px_20px] border border-solid border-[#E0E0E0] rounded-[8px]">
                    //                     <input type="text" placeholder="Navigation Lable" value="${menuname}"
                    //                         class="p-[15px_14px] border border-solid border-[#EDEDED] rounded-[4px] h-[48px] text-[14px] leading-light placeholder:text-[#8A8A8A] w-full mb-[24px]">
                    //                     <input type="text" placeholder="Path link"
                    //                         class="p-[15px_14px] border border-solid border-[#EDEDED] rounded-[4px] h-[48px] text-[14px] leading-light placeholder:text-[#8A8A8A] w-full mb-[24px]">
                    //                     <div class="flex items-center justify-end gap-[10px]">
                    //                         <a href="#" id="deletebtn" data-id="${result.Id}" data-bs-target="#deleteModal" data-bs-toggle="modal"
                    //                          class="flex justify-center place-items-center bg-[#ffffff]  hover:bg-[#F5F5F5] px-[14px] py-[7px] rounded-[4px] w-fit  h-[32px] font-light text-[14px] text-center leading-tight tracking-[0.7px] whitespace-nowrap border border-solid border-[#E6E6E6] text-[#989898]">
                    //                             Remove
                    //                         </a>
                    //                         <a href="#" id="updatebtn" data-id="${result.Id}" data-type="${result.Type}" data-typeid="${result.TypeId}" data-parent="${result.ParentId}"
                    //                             class="flex justify-center place-items-center bg-[#10A37F] hover:bg-[#148569]  px-[14px] py-[7px] rounded-[4px] w-fit  h-[32px] font-normal text-[14px] text-white text-center leading-tight tracking-[0.7px] whitespace-nowrap">
                    //                             Save
                    //                         </a>
                    //                     </div>
                    //                 </div>
                    //             </div>
                    //         </div>
                    //     </div>
                    // `;
                    // $('#accordionExample1').append(str);
                }
            });
        }
    });
});


$(document).on('click', '#updatebtn', function () {
    var menuId = $(this).data('id');
    var type = $(this).attr('data-type');
    var typeid = $(this).attr('data-typeid');
    parentmenuid = $(this).attr('data-parent');
    var $accordionItem = $(this).closest('.accordion-body');
    var label = $accordionItem.find('input[placeholder="Navigation Lable"]').val().trim();
    var path = $accordionItem.find('input[placeholder="Path link"]').val();
 
    var labelValid = label.length <= 20;
    var isDuplicate = false;
 
   
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
 
            console.log(result,"resulttt")
            if (result == 'true') {
                isDuplicate = true;
            }
        }
    });
 
   
    if (!labelValid) {
        console.log("jjsdsdsd")
        $accordionItem.find('.lablename-error').removeClass('hidden').addClass('mb-[24px]').text('Please Enter 20 Characters Only');
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
                "id": menuId,
                "menu_name": label,
                "webid": $('.templateid').attr('data-id'),
                "parentmenu_id": parentmenuid,
                "urlpath": path,
                csrf: $("input[name='csrf']").val(),
                "cid": typeid,
                "type": type
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



//drag and drop functionlity//
// $('.group').draggable({
//     handle: '.drag-handle',
//     helper: 'clone',
//     revert: 'invalid',
//     zIndex: 100,
//     start: function (event, ui) {

//         if ($(this).has('.group').length) {
//             return false;
//         }
//         $(this).addClass('dragging');
//     },
//     stop: function (event, ui) {
//         $(this).removeClass('dragging');
//     }
// });

// Make all groups draggable by their drag handle
$('.group').draggable({
    handle: '.drag-handle',
    helper: 'clone',
    revert: 'invalid',
    cursor: 'move',
    zIndex: 10000,
    start: function (event, ui) {
        $(this).addClass('dragging');
    },
    stop: function (event, ui) {
        $(this).removeClass('dragging');
    }
});


$('.group.dragitem, #drop-zone').droppable({
    accept: function ($dragged) {
        // Only accept if drop target is NOT a child AND dragged element is a group
        return !$(this).hasClass('ms-[20px]') && $dragged.hasClass('group');
    },
    hoverClass: 'drop-hover',
    tolerance: 'pointer',
    drop: function (event, ui) {
        const $dropTarget = $(this);
        const $draggedItem = ui.draggable;

        // Prevent dropping on restricted children or self
        if ($dropTarget.hasClass('menu-itemdev') || $dropTarget.closest('.ms-[20px]').length || $draggedItem.data('id') === $dropTarget.data('id')) {
            return;
        }

        let parentId = $dropTarget.data('id');
        let flag = typeof parentId === 'undefined';

        if (flag) {
            parentId = $('#menu_id').val();
        }

        // Trigger drop handler with cleaned variable naming
        handleDrop($draggedItem, $dropTarget, parentId, flag);
    }
});

function handleDrop($dragged, $droppedOn, parentId, flag) {
    // Prevent dropping on children or duplicate drops
    if ($droppedOn.hasClass('ms-[20px]')) return;
    if ($dragged.data('id') === $droppedOn.data('id')) return;

    const label = $dragged.find('input[placeholder="Navigation Lable"]').val();
    const path = $dragged.find('input[placeholder="Path link"]').val();
    const updateBtn = $dragged.find('#updatebtn');
    const draggedId = $dragged.data('id');

    $.ajax({
        url: "/admin/website/menu/updatemenuitems",
        type: "POST",
        data: {
            id: draggedId,
            menu_name: label,
            parentmenu_id: parentId,
            urlpath: path,
            csrf: $("input[name='csrf']").val(),
            cid: updateBtn.data('typeid'),
            type: updateBtn.data('type'),
            "webid": $('.templateid').attr('data-id'),
        },
        success: function (data) {
            // Update UI based on whether it is root or child drop
            if (flag) {
                $dragged.removeClass('ms-[20px]');
                $('#accordionExample1').append($dragged);
            } else {
                $dragged.addClass('ms-[20px]');
                $('.div' + parentId).append($dragged);
            }

            // Optional: Toast notification (uncomment if needed)
            // notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >Menu Updated Successfully </p ></div ></div ></li></ul> `;
            // $(notify_content).insertBefore(".header-rht");
            // setTimeout(function () {
            //      $('.toast-msg').fadeOut('slow', function () {
            //          $(this).remove();
            //      });
            // }, 5000);
        }
    });
}
// $(function() {
//   // Make root container sortable for parent items
//   $('#accordionExample1').sortable({
//     connectWith: '.child-container',
//     handle: '.drag-handle',
//     items: '> .group',
//     placeholder: 'sortable-placeholder',
//     tolerance: 'pointer',

//     start: function(e, ui) {
//       ui.placeholder.height(ui.item.outerHeight());
//     },

//     update: function(e, ui) {
//       const $movedItem = ui.item;
//       const movedId = $movedItem.data('id');

//       // If parent container is root or child container, detect new parent accordingly
//       const $newParentContainer = $movedItem.parent();
//       let newParentId = null;

//       if ($newParentContainer.hasClass('child-container')) {
//         newParentId = $newParentContainer.closest('.group').data('id');
//       } else if ($newParentContainer.attr('id') === 'accordionExample1') {
//         newParentId = 0; // Root level
//       }

//       // AJAX to update database
//       $.ajax({
//         url: "/admin/website/menu/updatemenuitems",
//         method: 'POST',
//         data: {
//           id: movedId,
//           parentmenu_id: newParentId,
//           csrf: $("input[name='csrf']").val()
//         },
//         success: function() { window.location.reload()
//              console.log('Menu updated successfully'); },
//         error: function() { console.error('Error updating menu'); }
//       });
//     }
//   }).disableSelection();

//   // Make child containers sortable for child items
//   $('.child-container').sortable({
//     connectWith: '#accordionExample1',
//     handle: '.drag-handle',
//     items: '> .group',
//     placeholder: 'sortable-placeholder',
//     tolerance: 'pointer',

//     start: function(e, ui) {
//       ui.placeholder.height(ui.item.outerHeight());
//     },

//     update: function(e, ui) {
//       const $movedItem = ui.item;
//       const movedId = $movedItem.data('id');

//       const $newParentContainer = $movedItem.parent();
//       let newParentId = null;

//       if ($newParentContainer.hasClass('child-container')) {
//         newParentId = $newParentContainer.closest('.group').data('id');
//       } else if ($newParentContainer.attr('id') === 'accordionExample1') {
//         newParentId = 0; // Root level
//       }

//       $.ajax({
//         url: "/admin/website/menu/updatemenuitems",
//         method: 'POST',
//         data: {
//           id: movedId,
//           parentmenu_id: newParentId,
//           csrf: $("input[name='csrf']").val()
//         },
//         success: function() { console.log('Menu updated successfully'); },
//         error: function() { console.error('Error updating menu'); }
//       });
//     }
//   }).disableSelection();
// });






//search function//

$(document).on('keyup', '.searchbtn', function () {
    var keyword = $(this).val().trim().toLowerCase();
    console.log("jjj")
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
    menudesc = $(this).attr('data-desc')
    menustatus = $(this).attr('data-status')
    var data = $(this).attr("data-id");
    $("#menuform").attr("name", "editmenu").attr("action", "/admin/website/menu/updatemenu/?webid=" + $('.templateid').attr('data-id'))
    $("input[name=menu_name]").val(menuname.trim());
    $("textarea[name=menu_desc]").val(menudesc.trim());
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


    if ($(this).val().trim().length >= 15) {

        $(this).removeClass('mb-[24px]')
        console.log("pppp")
        $(this).siblings('#lablename-error').removeClass('hidden').addClass('mb-[24px]')

    } else {
        $(this).addClass('mb-[24px]')
        $(this).siblings('#lablename-error').addClass('hidden').removeClass('mb-[24px]')
    }
})

$(document).on('click', '.opentab', function () {



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

$(document).on('click', '.cancelmenuitem', function () {

    $('#menu_name-error').hide()
    $('#urlpath-error').hide()
    $('input[name="lang"]:checked').prop('checked', false);
    $('.customUrlInput').addClass('hidden')
    $('.entriesDropdown').addClass('hidden')
    $('.pagesDropdown').addClass('hidden');
    $('.listingsDropdown').addClass('hidden');
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
                "cid": "",
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
        var selectedLabel = $(this).next('label').text().trim();

        if (selectedLabel === "Custom URL") {
            $('.customUrlInput').removeClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden');
            $('.listingsDropdown').addClass('hidden')
        } else if (selectedLabel === "Entries") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').removeClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
        } else if (selectedLabel === "Pages") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').removeClass('hidden')
            $('.listingsDropdown').addClass('hidden')
        } else if (selectedLabel === "Listings") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').removeClass('hidden')
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

    $('.navpath').val("/entries/" + url)
})

$(document).on('click', '.pageslug', function () {

    url = $(this).attr('data-slug').trim()

    $('.navpath').val("/pages/" + url)
})

$(".searchcatlists").keyup(function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase()


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

$(document).ready(function () {
    $('.listingsDropdownMenu').on('click', function (e) {
        e.stopPropagation();
    });
});