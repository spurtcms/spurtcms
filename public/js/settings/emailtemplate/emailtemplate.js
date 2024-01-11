var languagedata
/**This if for Language json file */
$(document).ready(async function () {

  var languagecode = $('.language-group>button').attr('data-code')

  await $.getJSON("/locales/"+languagecode+".json", function (data) {
      
      languagedata = data
  })

});

$(document).on('click','#edit-btn',function(){
      tempid = $(this).attr('data-id')
      $('#userid').val(tempid)
      $.ajax({
        url: "/settings/emails/edit-template",
        type: "GET",
        dataType: "json",
        data: { "id": tempid },
        success: function (result) {
            console.log("result",result)

          if (result.Id != 0) {
             
            var name = $("#tempname").val(result.TemplateName);
            var sub = $("#tempsub").val(result.TemplateSubject);
            var content = $("#temcont").val(result.TemplateMessage);
            $("#tempname").text(name);
            $("#tempsub").text(sub);
            $("#temcont").text(content)
  
          
         }
      }
      })

})
//**Template isactive status *//
function TempStatus(id) {
    $('#cb1' + id).on('change', function () {
      this.value = this.checked ? 1 : 0;
    }).change();
    var isactive = $('#cb1' + id).val();
    $.ajax({
      url: '/settings/emails/templateisactive',
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
             notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>Template Updated Successfully</span></div>';
                $(notify_content).insertBefore(".header-rht");
                setTimeout(function(){
                    $('.toast-msg').fadeOut('slow', function(){
                        $(this).remove();
                    });
                }, 5000); // 5000 milliseconds = 5 seconds
      }
    });
}
//**description focus function */
const Desc = document.getElementById('temcont');
const inputGroup = document.querySelectorAll('.input-group');

Desc.addEventListener('focus', () => {

  Desc.closest('.input-group').classList.add('input-group-focused');
});
Desc.addEventListener('blur', () => {
  Desc.closest('.input-group').classList.remove('input-group-focused');
});

/*search */
$(document).on("click", "#filterformsubmit", function () {
  var key = $(this).siblings().children(".search").val();
  if (key == "") {
      window.location.href = "/settings/emails/"
  } else {
      $('.filterform').submit();
  }
})

$(document).on('keyup','#searchemails',function(){

  if($('.search').val()===""){
      console.log("check")
      window.location.href ="/settings/emails"
      
  }

})

$(document).on('click','#uptemplate',function(){

console.log("checking")

$("#updtemp").validate({
  rules: {
    tempname: {
          required: true,
          space: true,
      },
      tempsub: {
          required: true,
          space: true,
          
      },
      temcont:{
        required: true,
        space: true,
        maxlength: 250,
      },
  },
  messages: {
    tempname: {
          required: "* Please enter templatename" ,
          space: "* " + languagedata.spacergx,
          
      },
      tempsub: {
          required: "* Please enter templatesubject" ,
          space: "* " + languagedata.spacergx,
         
      },
      temcont: {
        required: "* Please enter templatecontent" ,
          space: "* " + languagedata.spacergx,
          maxlength: "* Maximum 250 character allowed"
      }
  }
});

var formcheck = $("#updtemp").valid();
if (formcheck == true) {
  $('#updtemp')[0].submit();

}
else{
  Validationcheck()
  $(document).on('keyup', ".field", function () {
      Validationcheck()
  })
}
})

function Validationcheck(){

  if ($('#tempname').hasClass('error')) {
    $('#namegrb').addClass('input-group-error');
}else{
  $('#namegrb').removeClass('input-group-error');
}

if ($('#tempsub').hasClass('error')){
    $('#subgrb').addClass('input-group-error');
}else{
    $('#subgrb').removeClass('input-group-error');
}

if ($('#temcont').hasClass('error')){
  $('#textgrb').addClass('input-group-error');
}else{
  $('#textgrb').removeClass('input-group-error');
}
}

$(document).on('click','.close',function(){

  $('label.error').remove()
  $('.input-group').removeClass('input-group-error')

})