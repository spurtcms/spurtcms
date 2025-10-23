
$(document).on('click', '.myprofilebtn', function () {
    $('.myprofiledrop').toggleClass('hidden')
})

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

$(document).on('keyup', '.searchbtn', function() {
    var keyword = $(this).val().toLowerCase().trim();

    
    $('.entrylistdiv ul li').each(function(index) {
       
        if(index === 0) return;

        var text = $(this).text().toLowerCase();
        if(text.indexOf(keyword) !== -1) {
            $(this).removeClass('hidden');
        } else {
            $(this).addClass('hidden');
        }
    });
});


function getCookie(cname) {
    let name = cname + "=";
    let ca = document.cookie.split(';');

    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function delete_cookie(name) {
    document.cookie = name + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}