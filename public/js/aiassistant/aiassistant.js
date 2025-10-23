$(document).ready(function () {
    var pText = "Topic";

    if (pText == "Topic") {
        $('#topicinput').show()
        $('#keywordinput').hide()

    } else {
        $('#topicinput').hide()
        $('#keywordinput').show()
    }
    // Next btn initially not-working
    $("#nextbtn").addClass('pointer-events-none opacity-50');
    $('#topic-listdev').hide();

})


// select content type dropdown
$('.selecttypecontent').click(function () { // select topic
    pText = $(this).find('p').text();
    if (pText == "Topic") {
        $('#topicinput').show();
        $('#keywordinput').hide();
        $('#contentdropdown').text(pText);

        $("#closeimgbtn").addClass('hidden');
        $('#nextbtn').show();
        $('#close-textbtn').show();
        $('#topic-listdev').hide();

    } else {  // select keyword
        $('#nextbtn').hide();
        $('#close-textbtn').hide();
        $('#closeimgbtn').removeClass('hidden');

        $('#topicinput').hide();
        $('#keywordinput').show();
        $('#contentdropdown').text(pText);
        $("#nextbtn").addClass('pointer-events-none opacity-50');
        if ($('#topic-listdev').find('ul>li').length > 0) {
            $('#topic-listdev').show();

        } else {
            $('#topic-listdev').hide();
        }
        //    $('#topic-listdev').hide();
    }
})

// Clear all input value when close btn click
$(".nextclosebtn,#generatearticle-closebtn").on("click", function () {
    pText = "Topic";
    if (pText == "Topic") {
        $('#topicinput').show()
        $('#keywordinput').hide()

    } else {
        $('#topicinput').hide()
        $('#keywordinput').show()
    }

    $("#inputtopic").val('');
    $("#inputkeyword").val('');
    $('#selectedtone').text("Friendly")
    $("#tonetype").val("Friendly")
    $('#ecselectedtone').text("Friendly")
    $("#ectonetype").val("Friendly")
    $("#generatetopics-error").addClass('hidden');
    $("#nextbtn").addClass('pointer-events-none opacity-50');
    $("#closeimgbtn").addClass('hidden');
    $('#nextbtn').show();
    $('#close-textbtn').show();
    $("#contentdropdown").text("Topic")
    $('#Id3').removeClass('show');
    $('.offcanvas-backdrop').remove();
    $('#selectedtone').addClass('text-bold-gray')
    $('#dynamoictopic').empty();
    $('#topic-listdev').hide();
})


// Next btn working
$("#inputtopic").on('keydown', function () {
    var topiclength = $(this).val().length;
    if (topiclength > 0) {
        $("#nextbtn").removeClass('pointer-events-none opacity-50');
    } else {
        $("#nextbtn").addClass('pointer-events-none opacity-50');
    }
})

// genarate topics based on keywords
$("#generatetopics").on('click', function () {
    var keywordvalue = $('#inputkeyword').val()

    if (keywordvalue.length > 0) {
        $('#generatetopics-error').addClass('hidden')

        $.ajax({
            url: '/admin/aiassistance/generatetopics',
            type: 'post',
            data: { 'keyword': keywordvalue, csrf: $("input[name='csrf']").val() },
            success: function (result) {

                if (result.data==""){
                    notify_content = `<ul class="fixed top-[56px] right-[16px] z-[9999] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">Please Check Your AI Credentials</p></div></div> </li></ul>`;
    
    
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                        $('.toast-msg').fadeOut('slow', function () {
                            $(this).remove();
                        });
                    }, 5000); 
                }
                const topics = JSON.parse(result.data).topics;
                $('#dynamoictopic').empty();
                for (var i = 0; i < topics.length; i++) {
                    var list = `<li>
                        <input type="radio" class="hidden peer" id="radio${i}" name="radio">
                        <label
                            class="flex genkey items-center  before:w-[16px] before:h-[16px] before:bg-[url('/public/img/keyRadio-uncheck.svg')] mr-[9px] before:bg-contain peer-checked:before:bg-[url('/public/img/keyRadio-check.svg')] before:mr-[9px] cursor-pointer"
                            for="radio${i}">
                            <span
                                class=" w-full rounded-[4px] border border-[#EDEDED] p-[8px_12px] text-[14px] font-normal text-[#262626] leading-[17.5px]">
                                ${topics[i]}
                            </span>
                        </label>
                    </li>`
                    $('#dynamoictopic').append(list);
                }
                $('#topic-listdev').show();

            }
        })
    } else {
        $('#generatetopics-error').removeClass('hidden')
    }

})

// Select Tone
$('.tonetype').on('click', function () {
    var Tonetype = $(this).find('p').text()
    $('#selectedtone').text(Tonetype)
    $("#tonetype").val(Tonetype)
    $('#selectedtone').removeClass('text-bold-gray')

})

$('#nextbtn').on('click', function () {
    var topicval = $('#inputtopic').val();
    $('.dynamickey').empty()
    $.ajax({
        url: '/admin/aiassistance/receivetopicdata',
        type: 'post',
        data: { 'topic': topicval, csrf: $("input[name='csrf']").val() },
        success: function (result) {
            if (result.data==""){
                notify_content = `<ul class="fixed top-[56px] right-[16px] z-[9999] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">Please Check Your AI Credentials</p></div></div> </li></ul>`;


                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); 
            }
            const data = JSON.parse(result.data);
            $('#titletopic').val(topicval);
            $("#articlelength").val(data.articleLength)

            for (let x of data.keywords) {
                div = `<div class="flex space-x-[4px] mt-[8px] items-center px-[12px] py-[6px] rounded bg-[#F0FFFB]">
                            <p class="text-sm font-normal text-[#10A37F] keywords">${x}</p>
                            <a class="cursor-pointer w-[10px] removekw"><img src="/public/img/block-close.svg" alt=""></a>
                        </div>`
                $('.dynamickey').append(div);
            }
            $('#Id3').addClass('show');
            $('#Id2').removeClass('show');
        }
    })
})

$(document).on('click', "#generatearticle", function (event) {
    event.preventDefault();
    let keyword = "";
    $('.keywords').each(function () {
        keyword += $(this).text() + ",";
    });
    keyword = keyword.slice(0, -1);
    $('#keywords').val(keyword);
    $("#generateArticleForm").validate({
        rules: {
            titletopic: {
                required: true,
                maxlength: 30,
            },
            articlelength: {
                required: true,
                range: [300, 5000],
            }
        },
        messages: {
            titletopic: {
                required: "* Please enter your title",
                maxlength: "* Maximum 30 character allowed"
            },
            articlelength: {
                required: "* Please enter your article length",
                range: "* Article length must be between 300 to 5000",
            },
        }
    });

    if ($("#generateArticleForm").valid() == true) {
        $("#articleloader").removeClass("hidden")
        $('#generateArticleForm').submit();
    }

});

$(document).on('click', ".removekw", function () {

    $(this).parent("div").remove()
})

// create user prompt inputs based on topic
$(document).on('click', '.genkey', function () {
    var topicval = $(this).find('span').text();
    $('.dynamickey').empty();
    $.ajax({
        url: '/admin/aiassistance/receivetopicdata',
        type: 'post',
        data: { 'topic': topicval, csrf: $("input[name='csrf']").val() },
        success: function (result) {
            const data = JSON.parse(result.data);
            console.log(data.keywords);
            $('#titletopic').val(data.topic);
            $("#articlelength").val(data.articleLength)

            for (let x of data.keywords) {
                div = `<div class="flex space-x-[4px]  mt-[8px] items-center px-[12px] py-[6px] rounded bg-[#F0FFFB]">
                            <p class="text-sm font-normal text-[#10A37F] keywords">${x}</p>
                            <a class="cursor-pointer w-[10px] removekw"><img src="/public/img/block-close.svg" alt=""></a>
                        </div>`
                $('.dynamickey').append(div);
            }
            $('#Id3').addClass('show');
            $('#Id2').removeClass('show');
        }
    })
})

// create product details based on product name 
$('#ecnextbtn').on('click', function () {
    var productname = $('#inputproduct').val();
    $('.dynamickey').empty()
    $.ajax({
        url: '/admin/aiassistance/receiveproductdata',
        type: 'post',
        data: { 'productname': productname, csrf: $("input[name='csrf']").val() },
        success: function (result) {
            const data = JSON.parse(result.data);
            console.log(data.keywords);
            $('#titleproduct').val(productname);
            $("#descriptionlength").val(data.articleLength)

            for (let x of data.keywords) {
                div = `<div class="flex space-x-[4px]  mt-[8px] items-center px-[12px] py-[6px] rounded bg-[#F0FFFB]">
                            <p class="text-sm font-normal text-[#10A37F] keywords">${x}</p>
                            <a class="cursor-pointer w-[10px] removekw"><img src="/public/img/block-close.svg" alt=""></a>
                        </div>`
                $('.dynamickeyfeature').append(div);
            }
            $('#Id5').addClass('show');
            $('#Id4').removeClass('show');
        }
    })
})
$(document).keydown(function (event) {
    if (event.ctrlKey && event.key === '/') {
        $("#searchkey").focus().select();
    }
});
// Search functionality get back to list
$(document).on('keyup', '#searchkey', function () {

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($('#searchkey').val() === "") {

            window.location.href = "/admin/aiassistance/"

        }
    }
})

$(document).on('keypress', function (event) {
    var topicval = $('#inputtopic').val();
    if (event.key === 'Enter' && topicval != "") {
        $('#nextbtn').click();
    }
});

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
})
$(document).on("click", ".searchClosebtn", function () {
    $("#searchkey").val('')
    window.location.href = "/admin/aiassistance/"
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


$(document).on('keyup', '#articlelength', function () {

    var articleLength = $(this).val();


    var length = parseInt(articleLength, 10);


    var $errorMessage = $('#articlelength-error');


    if (length >= 300 && length <= 5000) {

        $errorMessage.html("");
    } else {

        $errorMessage.html('* Article length must be between 300 to 5000.');
    }
});


$(document).on('click', '#aiarticle', function () {

    count = $("#articlecount").val()
    console.log('Clicked!', count);

    var OwnApiKey = $("#OwnApiKey").val()
    console.log("OwnApiKey:",OwnApiKey);
    
    if (OwnApiKey=="openai") {
        console.log("trueeee");
        

        $('#Id2').css('display', 'block');
        $('#Id2').addClass("show");


        $('body').append('<div class="offcanvas-backdrop fade show"></div>');

    } else {
        if (count == 5) {
            console.log("falseeee");

            notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">Article Limit Reached! You can only create up to 5 articles.</p></div></div> </li></ul>`;


            $(notify_content).insertBefore(".header-rht");
            setTimeout(function () {
                $('.toast-msg').fadeOut('slow', function () {
                    $(this).remove();
                });
            }, 5000); // 5000 milliseconds = 5 seconds
        } else {

            console.log("falseeeeok");
            $('#Id2').css('display', 'block');
            $('#Id2').addClass("show");


            $('body').append('<div class="offcanvas-backdrop fade show"></div>');
        }
    }
})
$(document).on('click', '.nextclosebtn, #cancel-notify', function () {
    $('#Id2').css('display', 'none');
    $('#Id2').removeClass("show");


    $('.offcanvas-backdrop').remove();
});