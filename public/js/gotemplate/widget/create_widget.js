
var languagedata
var selectedLabel

/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
}
)
$(document).ready(function () {

    $('input[name="lang"]').change(function () {
        selectedLabel = $(this).next('label').text().trim();
        $('.navpath').val("")
        $('.product-error').addClass('hidden')
        if (selectedLabel === "Custom URL") {
            $('.customUrlInput').removeClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden');
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.menutype').val("custom_url")
            $('.channelDropdown').addClass('hidden')
        } else if (selectedLabel === "Entries") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').removeClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.widget_type').val('Entries')
        }
        else if (selectedLabel === "Pages") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').removeClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.widget_type').val('Pages')
        } else if (selectedLabel === "Listings") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.categoryDropdown').addClass('hidden')
            $('.listingsDropdown').removeClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.widget_type').val('Listings')
        } else if (selectedLabel === "Categories") {
            $('.customUrlInput').addClass('hidden')
            $('.entriesDropdown').addClass('hidden')
            $('.pagesDropdown').addClass('hidden')
            $('.listingsDropdown').addClass('hidden')
            $('.categoryDropdown').removeClass('hidden')
            $('.channelDropdown').addClass('hidden')
            $('.widget_type').val('Categories')
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
            $('.widget_type').val('Channels')
        }
    });
});

$(document).ready(function () {
    $('.listingsDropdownMenu').on('click', function (e) {
        e.stopPropagation();
    });
    $('.categoryDropdownMenu').on('click', function (e) {
        e.stopPropagation();
    });
    $('.entriesDropdown').on('click', function (e) {
        e.stopPropagation();
    });
    $('.pagesDropdown').on('click', function (e) {
        e.stopPropagation();
    });
    $('.channelDropdown').on('click', function (e) {
        e.stopPropagation();
    });

})


$(document).on('click', '#widgetsave', function () {


    GetProductIds()
    // var productids
    // if ($('.widget_type').val()=="Listings"){

    // let selectedIds = $('.selected-listings:checked').map(function () {
    //     return $(this).data('id');
    // }).get();
    //  productids = selectedIds.join(',');
    // }


    // $('.productids').val(productids);
    // form validation

    $("#widgetform").validate({
         ignore: [],
        rules: {
            title: {
                required: true,
                space: true,
                // duplicategrb: true

            },
            position: {
                required: true,

            },
            // widget_slug: {
            //     required: true
            // },
            status: {
                required: true
            },
            Product_type: {
                required: true
            }
        },
        messages: {
            title: {
                required: "* Please Enter Widget Title",
                space: "* " + languagedata.spacergx,
                // duplicategrb: "*" + languagedata.Members_Group.memgrpnamechk
            },
            position: {
                required: "* Please Enter Widget Position",

            },
            // widget_slug: {
            //     required: "* Please Enter Widget Slug",
            // },
            status: {
                required: "* Please Enter Widget Status",
            },
            Product_type: {
                required: "* Please Enter Widget Product"
            }

        },

    });

    if($('.widget_type').val()==""){

        prodcuttype=false

        $('.product-error').removeClass('hidden')
    }else{
         prodcuttype=true
        $('.product-error').addClass('hidden')
    }  

    var formcheck = $("#widgetform").valid();

    if ((formcheck == true) &&(prodcuttype==true)) {
        $('#widgetform')[0].submit();
    }

    return false
})

function GetProductIds() {
    var productids = '';

    var widgetType = $('.widget_type').val();

    if (widgetType === "Listings") {
        let selectedIds = $('.selected-listings:checked').map(function () {
            return $(this).data('id');
        }).get();
        productids = selectedIds.join(',');

    } else if (widgetType === "Entries") {
        let selectedIds = $('.entriesDropdown .selected-listings:checked').map(function () {
            return $(this).data('id');
        }).get();
        productids = selectedIds.join(',');

    } else if (widgetType === "Categories") {
        let selectedIds = $('.categoryDropdown .selectcheckbox:checked').map(function () {

            let lastSpanId = $(this).closest('li').find('label span').last().attr('data-id');
            return lastSpanId;
        }).get();
        productids = selectedIds.join(',');
    }
    else if (widgetType === "Channels") {
        let selectedIds = $('.channelDropdown .selected-chennals:checked').map(function () {
            return $(this).data('id');
        }).get();
        productids = selectedIds.join(',');

    } else if (widgetType === "Pages") {

        let selectedIds = $('.pagesDropdown .selected-pages:checked').map(function () {
            return $(this).data('id');
        }).get();
        productids = selectedIds.join(',');
    }

    // Now assign the collected IDs to your hidden input
    $('.productids').val(productids);

}
$(document).on('click', '.statusdropdown', function () {

    $(this).closest('ul').removeClass('show');

    statusval = $(this).text().trim()

    console.log(statusval, "checkstatus")

    $('.status_type_span').text(statusval)

    if (statusval == "Active") {

        $('#status').val(1)
    } else {
        $('#status').val(0)

    }
})


$(document).on('click', '.postiondropdown', function () {

    $(this).closest('ul').removeClass('show');

    postionval = $(this).text().trim()

    $('.position_span').text(postionval)

    $('#position').val(postionval)

    $('#position-error').addClass('hidden')
})


$(document).on('click', '.sortdropdown', function () {

    $(this).closest('ul').removeClass('show');

    sortorder = $(this).text().trim()

    $('.sortorder_span').text(sortorder)

    $('#sort_order').val(sortorder)
})

$(document).ready(function () {

    producttype = $('.widget_type').val().trim()

    productids = $('.productids').val()

    console.log(productids, "idssss")

    if (producttype != "") {

        if (producttype == "Entries") {

            console.log("checkdfdfdf")
            $('.entriesDropdown').removeClass('hidden')
            $('#radioEntries').prop('checked', true)

            if (productids) {

                var idsArray = productids.split(',').map(function (item) {
                    return item.trim();
                });


                $('.entriesDropdown ul li').each(function () {
                    var liId = $(this).data('id')?.toString();
                    if (idsArray.includes(liId)) {
                        $(this).find('input[type="checkbox"]').prop('checked', true);
                    }
                });
            }

        }else if  (producttype == "Listings") {

            
            $('.listingsDropdown').removeClass('hidden')
            $('#radioListings').prop('checked', true)

           if (productids) {
       
        var idsArray = productids.split(',').map(function(item) {
            return item.trim();
        });

      
        $('.listingsDropdown ul li').each(function () {
            var liId = $(this).data('id')?.toString();
            if (idsArray.includes(liId)) {
                $(this).find('input[type="checkbox"]').prop('checked', true);
            }
        });
    }

        }else if  (producttype == "Categories") {

            
            $('.categoryDropdown').removeClass('hidden')
            $('#radioCategories').prop('checked', true)

           if (productids) {
       
        var idsArray = productids.split(',').map(function(item) {
            return item.trim();
        });

      
        $('.categoryDropdown ul li').each(function () {
            var liId = $(this).find('label span').last().attr('data-id')?.toString();
            if (idsArray.includes(liId)) {
                $(this).find('input[type="checkbox"]').prop('checked', true);
            }
        });
    }

        } else if (producttype == "Pages") {
 
            $('.pagesDropdown').removeClass('hidden')
            $('#radioPages').prop('checked', true)
 
            if (productids) {
 
                var idsArray = productids.split(',').map(function (item) {
                    return item.trim();
                });
 
                $('.pagesDropdown ul li.page-dropdownlist').each(function () {
 
                    var liId = $(this).data('id')?.toString();
 
                    if (idsArray.includes(liId)) {
                        $(this).find('input[type="checkbox"]').prop('checked', true);
                    }
                });
            }
            
        } else if (producttype == "Channels") {
 
            $('.channelDropdown').removeClass('hidden')
            $('#radioChannels').prop('checked', true)
 
            if (productids) {
 
                var idsArray = productids.split(',').map(function (item) {
                    return item.trim();
                });
 
                $('.channelDropdown ul li').each(function () {
 
                    var liId = $(this).data('id')?.toString();
 
                    if (idsArray.includes(liId)) {
                        $(this).find('input[type="checkbox"]').prop('checked', true);
                    }
                });
            }
        }




    }
})

$(".searchcatlists").keyup(function () {
    var found = false;
    var searchTerm = $(this).val().trim().toLowerCase();

    $(".entry-dropdownlist").each(function () {
        var $categoryItem = $(this);

        var categoryText = $categoryItem.find("label").map(function () {
            return $(this).text().toLowerCase();
        }).get().join(" ");

        var isVisible = categoryText.indexOf(searchTerm) > -1;
        $categoryItem.toggle(isVisible);

        if (isVisible) {
            found = true;
        }
    });

    if (found) {
        $(".noData-foundentry").hide();
    } else {
        $(".noData-foundentry").show();
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



