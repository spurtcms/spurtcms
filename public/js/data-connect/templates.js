$(document).on('click', '#infoCloseBtn', function () {

    $('#hd-crd').addClass('hidden')
    $('#infoShowBtn').removeClass('hidden')
    $('#infoHideBtn').addClass('hidden')

})

$('.hd-crd-btn').click(function () {
  
    if ($('#hd-crd').is(':visible')) {
        $('#hd-crd').addClass('hidden').removeClass("show"); 
        document.cookie = `tempbanner=false; path=/;`;
    } else {
        $('#hd-crd').addClass("show").removeClass('hidden');
        document.cookie = `tempbanner=true; path=/;`; 
    }
});


$(document).on('click', '#infoShowBtn', function () {
    $('#hd-crd').removeClass('hidden')
    $('#infoShowBtn').addClass('hidden')
    $('#infoHideBtn').removeClass('hidden')
})

$(document).on('click', '#infoHideBtn', function () {
    $('#hd-crd').addClass('hidden')
    $('#infoShowBtn').removeClass('hidden')
    $('#infoHideBtn').addClass('hidden')
})


//search bar code below

$(document).on('keypress', '#templateSearchBar', function (e) {

    if (e.which == 13) {
        var searchText = $('#templateSearchBar').val()
        console.log(searchText, "searchText");


        if (searchText != "") {
            $('#templateSearchLink').attr('href', "/templates/?keyword=" + searchText)
            $('#templateSearchBar').attr('value', searchText)
            $('#templateSearchLink').get(0).click()
        } else {
            window.location.href = "/templates/"
        }

    }
})

// on pressing backspace if searcbar is empty show list page
$(document).on('keyup', '#templateSearchBar', function (e) {
    var searchVal = $('#templateSearchBar').attr('value')
    if (e.which == 8 && searchVal != "") {
        var searchKey = $('#templateSearchBar').val()
        if (searchKey == "") {
            window.location.href = "/templates/"
        }
    }
})

$(document).on('click', '#removeTemplateFilter', function () {

    window.location.href = "/templates/"
})

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
  })

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')

    window.location.href = "/templates/"
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

$(document).on("click", ".tempHoverIcon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})