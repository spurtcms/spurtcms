$(document).ready(function () {

    $('.hd-crd-btn').click(function () {

        if ($('#hd-crd').is(':visible')) {
            $('#hd-crd').addClass('hidden').removeClass("show");
            document.cookie = `ctabanner=false; path=/;`;
        } else {
            $('#hd-crd').addClass("show").removeClass('hidden');
            document.cookie = `ctabanner=true; path=/;`;
        }
    });



    $(document).on("click", ".Closebtn", function () {
        $(".search").val('')
        $(".Closebtn").addClass("hidden")
        $(".SearchClosebtn").removeClass("hidden")
        $(".srchBtn-togg").removeClass("pointer-events-none")
    })

    $(document).on("click", ".searchClosebtn", function () {
        $(".search").val('')
        window.location.href = "/admin/cta/form-responses"
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



})
