$("#filterform").validate({

    rules: {
        fromdate: {
            required: true,

        },
        todate: {
            required: true,

        }

    },
    messages: {
        fromdate: {
            required: "* Please Enter Form Date",

        },
        todate: {
            required: "* Please Enter To Date",
        }

    }
})

$(document).ready(function () {
    
    $(".hiddenimage").each(function () {
        var value = $(this).val();
        $(this).closest('td').find('.imagesrc').attr("src", value);
    });

    $(".Date").each(function () {
        var value = $(this).val();
        hello = new Date(value)
        $(this).closest('td').find(".DateSet").html(hello.toLocaleDateString())
    });

    $(".Time").each(function () {
        var value = $(this).val();
        hello = new Date(value)
        $(this).closest('td').find(".TimeSet").html(hello.toLocaleTimeString())
    });

})
