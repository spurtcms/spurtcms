var languagedata
/** */
$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')

    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })
  
  })
  
/**Channellist delete */
$(document).on('click','.channeldelete',function(){

    $('.deltitle').text(languagedata.Channell.delchltitle)
  
    $('.deldesc').text(languagedata.Channell.delchlsection)
  
    $('.delname').text($(this).parents('.channels-card').children('.channels-card-top').children(".card-top-head").children('h3').text())

    $('#delid').attr('data-id',$(this).attr('data-id'))

    $('#delid').parent('a').attr('href',"/channels/deletechannel?id="+$('#delid').attr('data-id'))


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
        Removetoast();
        if (result.length != 0) {
          notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>'+languagedata.Toast.chanstatupdnotify+'</span></div>';
          $(notify_content).insertBefore(".breadcrumbs");
          setTimeout(function(){
            $('.toast-msg').fadeOut('slow', function(){
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    
          
          $(this).val(result.IsActive);
        } else {
          // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>'+languagedata.internalserverr+'</span></div>';
          // $(notify_content).insertBefore(".breadcrumbs");
          notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>'+languagedata.Toast.internalserverr+'</span></div>';
          $(notify_content).insertBefore(".breadcrumbs");
          setTimeout(function(){
            $('.toast-msg').fadeOut('slow', function(){
                $(this).remove();
            });
        }, 5000); // 5000 milliseconds = 5 seconds
    
        }
      }
    });
  }

  var modifier = 1

  var scrollablediv = document.querySelector('.channels-container')

  var offset = 0;

  var scrollablediv = document.querySelector('.channels-container');

  var space_divheight = scrollablediv.scrollHeight

scrollablediv.addEventListener('scroll', function() {
  // if (scrollablediv.scrollTop + scrollablediv.clientHeight >= scrollablediv.scrollHeight) {

    var readOnlyHeight = scrollablediv.offsetHeight
   
    var scrollposition = scrollablediv.scrollTop
 
    var currentscroll = readOnlyHeight + scrollposition
 
    if (currentscroll + modifier > space_divheight) {
  
      if (scrollposition >= 140 && scrollposition <= 180) {
  
        offset = offset+1;
  
      } 

      offset ++;
     
      $.ajax({
        url: '/channels/pagination',
        type: 'POST',
        async: false,
        data: {
          offset:offset,
          csrf: $("input[name='csrf']").val()
        },
        dataType: 'json',
        cache: false,
        success: function (result) {
          var channelstr = ``
          for (let x of result) {
        
          channelstr +=`
         
            
                  <div class="channels-card">
                      <div class="channels-card-top">
                          <div class="card-top-head">
                              <div style="display: flex; justify-content: space-between;">
                                  <h6 class="para-extralight"> `+languagedata.updated+``+x.DateString+` </h6>
                                  <h6 class="para-extralight">(`+x.SlugName+` `+languagedata.Channell.entriesavailable + `)</h6>
                              </div>
                              <h3 class="heading-five">`+x.ChannelName+`</h3>
                          </div>
                          <p class="para-extralight">
                              `+x.ChannelDescription+`
                          </p>
                      </div>
  
                      <div class="channels-card-btm">
                          <div class="btn-group action-btn-grp">
                              <button type="button" class="dropdown-toggle primary {{$THEMECOLOR}}" data-bs-toggle="dropdown"
                                  aria-expanded="false">
                                  Action
                                  <img src="/public/img/arrow-white-down.svg" alt="" />
                              </button>
                              <ul class="dropdown-menu dropdown-menu-end">
                                  <li>
                                      <a href="/channels/editchannel/`+x.Id+`">
                                          <button class="dropdown-item" type="button">
                                              <span><img src="/public/img/channel-configuration.svg" alt="" /></span>
                                              Configuration
                                          </button></a>
                                  </li>
                                  <li>
                                      <button class="dropdown-item channeldelete" type="button" data-bs-toggle="modal"
                                          data-bs-target="#centerModal" data-id="`+x.Id+`">
                                          <span><img src="/public/img/delete.svg" alt="" /></span> Delete
                                      </button>
                                  </li>
                              </ul>
                          </div>
  
                          <div class="toggle">
                              <input class="tgl tgl-light" id="cb`+x.Id+`" type="checkbox"
                                  onclick="ChannelStatus('`+x.Id+`')"  checked  />
                              <label class="tgl-btn" for="cb`+x.Id+`"></label>
                          </div>
                      </div>
                  </div>
                 
          `;
}
$('.channels-cards-container').append(channelstr);

}
      })
    
    }
  
  });

  // Search functionality get back to list
  $(document).on('keyup','.search',function(){

    if($('.search').val()===""){
        console.log("check")
        window.location.href ="/channels/"
        
    }
  
  })

