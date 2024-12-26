var languagedata


/*block hide and show*/
$(document).ready(function () {
  $('.hd-crd-btn').click(function () {
    $('#hd-crd, .hd-crd-btn').toggleClass('hide');
  });
});

/** */
$(document).ready(async function () {

  var languagepath = $('.language-group>button').attr('data-path')

  await $.getJSON(languagepath, function (data) {

    languagedata = data
  })

})



// Search functionality get back to list
$(document).on('keyup', '#search', function () {

  if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

    if ($('#search').val() === "") {

      window.location.href = "/channels/"

    }
  }
})
/**Channellist delete */
$(document).on('click', '.channeldelete', function () {

  var url = window.location.search
  const urlpar = new URLSearchParams(url)
  pageno = urlpar.get('page')

  $('.deltitle').text(languagedata.Channell.delchltitle)
  $('.deldesc').text(languagedata.Channell.delchlsection)
  $('.delname').text($(this).parents('.f-chn').find('.chname').text())
  $('#delid').show()
  $('#delid').attr('data-id', $(this).attr('data-id'))
  $('#dltCancelBtn').text("cancel")
  if (pageno == null) {
    $('#delid').attr('href', "/channels/deletechannel?id=" + $(this).attr('data-id'))

  } else {
    $('#delid').attr('href', "/channels/deletechannel?id=" + $(this).attr('data-id') + "&page=" + pageno)

  }



})

/*Channel Status */
function ChannelStatus(id) {
  $('#cb' + id).on('change', function () {
    this.value = this.checked ? 1 : 0;
  }).change();
  var isactive = $('#cb' + id).val();
  $.ajax({
    url: '/channels/status',
    type: 'POST',
    async: false,
    data: {
      id: id,
      isactive: isactive,
      csrf: $("input[name='csrf']").val()
    },
    dataType: 'json',
    cache: false,
    success: function (result) {

      // Removetoast();
      if (result == true) {

        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast.chanstatupdnotify + `</p ></div ></div ></li></ul>`
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
          $('.toast-msg').fadeOut('slow', function () {
            $(this).remove();
          });
        }, 5000); // 5000 milliseconds = 5 seconds

        $(this).val(result.IsActive);
      } else {
        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Toast.internalserverr + `</p ></div ></div ></li></ul>`
        $(notify_content).insertBefore(".header-rht");
        setTimeout(function () {
          $('.toast-msg').fadeOut('slow', function () {
            $(this).remove();
          });
        }, 5000); // 5000 milliseconds = 5 seconds

      }
    }
  });
}

// var modifier = 1

// var scrollablediv = document.querySelector('.channels-container')

// var offset = 0;

// var scrollablediv = document.querySelector('.channels-container');

// var space_divheight = scrollablediv.scrollHeight

// scrollablediv.addEventListener('scroll', function () {
//   // if (scrollablediv.scrollTop + scrollablediv.clientHeight >= scrollablediv.scrollHeight) {

//   var readOnlyHeight = scrollablediv.offsetHeight

//   var scrollposition = scrollablediv.scrollTop

//   var currentscroll = readOnlyHeight + scrollposition

//   if (currentscroll + modifier > space_divheight) {

//     if (scrollposition >= 140 && scrollposition <= 180) {

//       offset = offset + 1;

//     }

//     offset++;

//     $.ajax({
//       url: '/channels/pagination',
//       type: 'POST',
//       async: false,
//       data: {
//         offset: offset,
//         csrf: $("input[name='csrf']").val()
//       },
//       dataType: 'json',
//       cache: false,
//       success: function (result) {
//         var channelstr = ``
//         for (let x of result) {

//           channelstr += `


//                   <div class="channels-card">
//                       <div class="channels-card-top">
//                           <div class="card-top-head">
//                               <div style="display: flex; justify-content: space-between;">
//                                   <h6 class="para-extralight"> `+ languagedata.updated + `` + x.DateString + ` </h6>
//                                   <h6 class="para-extralight">(`+ x.SlugName + ` ` + languagedata.Channell.entriesavailable + `)</h6>
//                               </div>
//                               <h3 class="heading-five">`+ x.ChannelName + `</h3>
//                           </div>
//                           <p class="para-extralight">
//                               `+ x.ChannelDescription + `
//                           </p>
//                       </div>

//                       <div class="channels-card-btm">
//                           <div class="btn-group action-btn-grp">
//                               <button type="button" class="dropdown-toggle primary " data-bs-toggle="dropdown"
//                                   aria-expanded="false">
//                                   Action
//                                   <img src="/public/img/arrow-white-down.svg" alt="" />
//                               </button>
//                               <ul class="dropdown-menu dropdown-menu-end">
//                                   <li>
//                                       <a href="/channels/editchannel/`+ x.Id + `">
//                                           <button class="dropdown-item" type="button">
//                                               <span><img src="/public/img/channel-configuration.svg" alt="" /></span>
//                                               Configuration
//                                           </button></a>
//                                   </li>
//                                   <li>
//                                       <button class="dropdown-item channeldelete" type="button" data-bs-toggle="modal"
//                                           data-bs-target="#centerModal" data-id="`+ x.Id + `">
//                                           <span><img src="/public/img/delete.svg" alt="" /></span> Delete
//                                       </button>
//                                   </li>
//                               </ul>
//                           </div>

//                           <div class="toggle">
//                               <input class="tgl tgl-light" id="cb`+ x.Id + `" type="checkbox"
//                                   onclick="ChannelStatus('`+ x.Id + `')"  checked  />
//                               <label class="tgl-btn" for="cb`+ x.Id + `"></label>
//                           </div>
//                       </div>
//                   </div>

//           `;
//         }
//         $('.channels-cards-container').append(channelstr);

//       }
//     })

//   }

// });

$(document).on('click', "#work-tab", function () {
  $('#create-channel-btn').addClass('hidden')
  $('#records').addClass('hidden')
  $('.pagination').addClass('hidden')
  
  
})

$(document).on('click', "#collection-tab", function () {
  $('#create-channel-btn').removeClass('hidden')
  $('#records').removeClass('hidden')
  $('.pagination').removeClass('hidden')
})

$(document).on('click', '.delete-toast-btn', function () {
  $('.deltitle').text(languagedata.Channell.delchltitle)
  $('.deldesc').text(languagedata.Channell.thischannelcontainsentries)
  $('#delid').hide()
  $('.delname').text($(this).parents('.f-chn').find('.chname').text())
  $('#dltCancelBtn').text("Ok")
})

$(document).on("click", ".Closebtn", function () {
  $(".search").val('')
  $(".Closebtn").addClass("hidden")
  $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
  $(".search").val('')
  window.location.href = "/channels/"
})

$(document).ready(function () {

  $('.search').on('input', function () {
      if ($(this).val().length >= 1) {
          $(".Closebtn").removeClass("hidden")
          $(".srchBtn-togg").addClass("pointer-events-none")
      } else {
          $(".Closebtn").addClass("hidden")
          $(".srchBtn-togg").removeClass("pointer-events-none")
      }
  });
})




