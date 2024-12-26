$(document).on('click', '#infoCloseBtn', function () {

    $('#hd-crd').addClass('hidden')
    $('#infoShowBtn').removeClass('hidden')
    $('#infoHideBtn').addClass('hidden')

})


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
})

$(document).on("click", ".searchClosebtn", function () {
    $(".search").val('')

    window.location.href = "/templates/"
})

$(document).ready(function () {

    $('.search').on('input', function () {
        if ($(this).val().length >= 1) {
            $(".Closebtn").removeClass("hidden")
        } else {
            $(".Closebtn").addClass("hidden")
        }
    });
})

$(document).on("click", ".tempHoverIcon", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
})