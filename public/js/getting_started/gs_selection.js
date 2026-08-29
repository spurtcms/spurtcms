



$(document).on('click','.channelnextbtn',function(){

    if ($('#channelid').val()==0){



           notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li> <div class="toast-msg flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify" > <img src="/public/img/close-toast.svg" alt="close"> </a> <div> <img src="/public/img/danger-group-12.svg" alt="toast error"> </div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Warning</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] ">Please select Channel.</p></div></div> </li></ul>`;
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function () {
                    $('.toast-msg').fadeOut('slow', function () {
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds
    }else{

          $('#channelform')[0].submit();
    }
})


$(document).on('click','.radiobtn',function(){

    channelid =$(this).attr('data-id')

    choosetype =$(this).attr('data-type')


    $('#channelid').val(channelid)

    $('#choose_type').val(choosetype)
})