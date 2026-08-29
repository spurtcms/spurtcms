
var languagedata


/** */
$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')

    await $.getJSON(languagepath, function (data) {

        languagedata = data
    })
}
)

$(document).on('click', '#deletebtn', function () {

   

    var widgetid = $(this).attr("delete-data-id")
 console.log("kkk",widgetid)
    $("#content").text("Are you sure you want to remove this widget")
    var url = window.location.search
    const urlpar = new URLSearchParams(url)
    pageno = urlpar.get('page')
  



    if (pageno == null) {
        $('#delid').attr('href', "/admin/website/widgets/deletewidget/" + widgetid );

    } else {
        $('#delid').attr('href', "/admin/website/widgets/deletewidget/" + widgetid );

    }

   
    $(".deltitle").text( "Delete Page?")
    $('.delname').text($(this).parents('tr').find('td:eq(0) a').text())

})



function showToast(type, message) {
    $('.toast-ul').remove();

    var isSuccess = type === 'success';
    var notify_content = `
        <ul class="toast-ul fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]">
          <li>
            <div class="toast-msg flex max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] ${isSuccess ? 'border-[#278E2B] bg-[#E2F7E3]' : 'border-[#A32D2D] bg-[#FCEBEB]'}">
              <a href="javascript:void(0)" class="cancel-notify absolute right-[8px] top-[8px]">
                <img src="/public/img/close-toast.svg" alt="close">
              </a>
              <div>
                <img src="${isSuccess ? '/public/img/toast-success.svg' : '/public/img/danger-group-12.svg'}" alt="">
              </div>
              <div>
                <h3 class="${isSuccess ? 'text-[#278E2B]' : 'text-[#A32D2D]'} text-normal leading-[17px] font-normal mb-[5px]">
                  ${isSuccess ? 'Success' : 'Error'}
                </h3>
                <p class="text-[#262626] text-[12px] font-normal leading-[15px]">${message}</p>
              </div>
            </div>
          </li>
        </ul>`;

    $('body').append(notify_content);

    setTimeout(function () {
        $('.toast-msg').fadeOut('slow', function () {
            $(this).closest('.toast-ul').remove();
        });
    }, 5000);
}

function WidgetStatus(id) {
    $('#cb' + id).off('change').on('change', function () {
        this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb' + id).val();

    $.ajax({
        url: '/admin/website/widgets/widgetstatuschange',
        type: 'POST',
        async: false,
        data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
        dataType: 'json',
        cache: false,
        success: function (result) {
            if (result) {
                showToast('success', 'Widget Updated Successfully');
                console.log("dddd");
                
            } else {
                showToast('error', languagedata.internalserverr);
            }
        },
        error: function () {
            showToast('error', languagedata.internalserverr);
        }
    });
}


$(document).on('click', '.cancel-notify', function () {
    $(this).closest('.toast-ul').find('.toast-msg').fadeOut('slow', function () {
        $(this).closest('.toast-ul').remove();
    });
});

$(document).on("click", ".Closebtn", function () {
    $(".search").val('')
    $(".Closebtn").addClass("hidden")
    $(".SearchClosebtn").removeClass("hidden")
    $(".srchBtn-togg").removeClass("pointer-events-none")
   
})

$(document).on("click", ".searchClosebtn", function () {

    $(".search").val('')
    
  window.location.href = "/admin/browsetheme/configure/"+$('#templateId').val()


})

$(document).ready(function () {

    $('.search').on('input', function () {

        console.log("chdfdffdfdf")
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

//Filter dropdown function//

$(document).on('click','.statuscls',function(){

    statusval =$(this).text().trim()
    $(".filterleveldropdown").removeClass("show")
    $('.slctstatus').text(statusval)
    $('#statusid').val(statusval)
})