$(document).ready(function () {

    $(document).on('click', '.searchbtn', function (e) {
        e.stopPropagation();
        $('.entrylistdiv').removeClass('hidden');
    });


    $('body').on('click', function () {
        $('.entrylistdiv').addClass('hidden');
    });


    $(document).on('click', '.entrylistdiv', function (e) {
        e.stopPropagation();
    });
});

$(document).on('keyup', '.searchbtn', function () {
    var keyword = $(this).val().toLowerCase().trim();

    // Loop through each section (li)
    $('ul > li').each(function () {
        var hasVisible = false;

        // Loop through all .searchtag elements inside this section
        $(this).find('.searchtag').each(function () {
            var text = $(this).data("value").toLowerCase();

            if (text.indexOf(keyword) !== -1) {
                $(this).removeClass('hidden');
                hasVisible = true;
            } else {
                $(this).addClass('hidden');
            }
        });

        // If no visible .searchtag in this section, hide the whole li
        if (!hasVisible) {
            $(this).addClass('hidden');
        } else {
            $(this).removeClass('hidden');
        }
    });
});
