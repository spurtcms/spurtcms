var languagedata
/**This if for Language json file */
$(document).ready(async function () {

  var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })

});

$(document).keydown(function(event) {
  if (event.ctrlKey && event.key === '/') {
      $(".search").focus().select();
  }
});

$(document).on('click','#edit-btn',function(){
      tempid = $(this).attr('data-id')
      $('#userid').val(tempid)
      var url = window.location.search;
      const urlpar = new URLSearchParams(url);
      pageno = urlpar.get('page');
      $("#pageno").val(pageno)
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


    if (event.key === 'Backspace') {

  if($('.search').val()===""){
    
      window.location.href ="/settings/emails"
      
  }
    }
})

$(document).on('click','#update-templatebtn',function(){

console.log("checking")

$("form[name='emailtemplateform']").validate({
  rules: {
    
      tempsub: {
          required: true,
          space: true,
          
      },
      tempdesc: {
        required: true,
        space: true,
    },
      temcont:{
        required: true,
        space: true,
        // maxlength: 250,
      },
  },
  messages: {
  
      tempsub: {
          required: "* Please enter Template Subject" ,
          space: "* " + languagedata.spacergx,
         
      },
      tempdesc: {
        required: "* Please enter Template Description" ,
        space: "* " + languagedata.spacergx,
        
    },
      temcont: {
        required: "* Please enter templatecontent" ,
          space: "* " + languagedata.spacergx,
          // maxlength: "* Maximum 250 character allowed"
      }
  }
});

var formcheck = $("#mailtemplate_form").valid();


if (formcheck == true) {
  $('#mailtemplate_form')[0].submit();

}
else{
  Validationcheck()
  $(document).on('keyup', ".field", function () {
      Validationcheck()
  })
}

})

function Validationcheck(){

    if ($('#emailsub').hasClass('error')){

        console.log("checkkkuuuuuuuuuu")
        $('#subgrb').addClass('input-group-error');
    }else{
        $('#subgrb').removeClass('input-group-error');
    }

  if ($('#emaildesc').hasClass('error')) {
    $('#desgrb').addClass('input-group-error');
}else{
  $('#desgrb').removeClass('input-group-error');
}


}

$(document).on('click','.close',function(){

  $('label.error').remove()
  $('.input-group').removeClass('input-group-error')

})

var ckeditor1;
var userids =[]
var templatestatus=[]

$(document).ready(function(){
    CKEDITORS()
})

$(document).on('click','.open-emailtemp',function(){

    console.log("checkclick")
    tempid = $(this).attr('data-id')
    $('#userid').val(tempid)
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');
    $("#pageno").val(pageno)
    $.ajax({
      url: "/settings/emails/edit-template",
      type: "GET",
      dataType: "json",
      data: { "id": tempid },
      success: function (result) {
          console.log("result",result)
        if (result.Id != 0) {
          var pathname = window.location.pathname
          $('#input-pagefind').val(pathname)
          $('#input-tempid').val(tempid)
          $("#emailtemp-name").text(result.TemplateName);
          $("#emailsub").val(result.TemplateSubject);
          $("#emaildesc").val(result.TemplateDescription)
          ckeditor1.setData(result.TemplateMessage)
       }
    }
    })

})


function CKEDITORS() {

    var url = $('#urlid').val();

    CKEDITOR.ClassicEditor.create(document.getElementById("text"), {
        toolbar: {
            items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'horizontalLine', 'link', 'code'],
            shouldNotGroupWhenFull: true
        },
        list: {
            properties: {
                styles: true,
                startIndex: true,
                reversed: true
            }
        },
        heading: {
            options: [
                { model: 'paragraph', title: 'Paragraph', class: 'ck-heading_paragraph' },
                { model: 'heading1', view: 'h1', title: 'Heading 1', class: 'ck-heading_heading1' },
                { model: 'heading2', view: 'h2', title: 'Heading 2', class: 'ck-heading_heading2' },
                { model: 'heading3', view: 'h3', title: 'Heading 3', class: 'ck-heading_heading3' },
                { model: 'heading4', view: 'h4', title: 'Heading 4', class: 'ck-heading_heading4' },
                { model: 'heading5', view: 'h5', title: 'Heading 5', class: 'ck-heading_heading5' },
                { model: 'heading6', view: 'h6', title: 'Heading 6', class: 'ck-heading_heading6' }
            ]
        },
        placeholder: languagedata?.Channell?.plckeditor,
        fontFamily: {
            options: [
                'default',
                'Arial, Helvetica, sans-serif',
                'Courier New, Courier, monospace',
                'Georgia, serif',
                'Lucida Sans Unicode, Lucida Grande, sans-serif',
                'Tahoma, Geneva, sans-serif',
                'Times New Roman, Times, serif',
                'Trebuchet MS, Helvetica, sans-serif',
                'Verdana, Geneva, sans-serif'
            ],
            supportAllValues: true
        },
        fontSize: {
            options: [10, 12, 14, 'default', 18, 20, 22],
            supportAllValues: true
        },
        htmlSupport: {
            allow: [
                {
                    name: /.*/,
                    attributes: true,
                    classes: true,
                    styles: true
                }
            ]
        },
        htmlEmbed: {
            showPreviews: true
        },
        link: {
            decorators: {
                addTargetToExternalLinks: true,
                defaultProtocol: 'https://',
                toggleDownloadable: {
                    mode: 'manual',
                    label: 'Downloadable',
                    attributes: {
                        download: 'file'
                    }
                }
            }
        },
        mention: {
            feeds: [
                {
                    marker: '@',
                    feed: [
                        '@apple', '@bears', '@brownie', '@cake', '@cake', '@candy', '@canes', '@chocolate', '@cookie', '@cotton', '@cream',
                        '@cupcake', '@danish', '@donut', '@dragée', '@fruitcake', '@gingerbread', '@gummi', '@ice', '@jelly-o',
                        '@liquorice', '@macaroon', '@marzipan', '@oat', '@pie', '@plum', '@pudding', '@sesame', '@snaps', '@soufflé',
                        '@sugar', '@sweet', '@topping', '@wafer'
                    ],
                    minimumCharacters: 1
                }
            ]
        },
        removePlugins: [
            'ExportPdf',
            'ExportWord',
            'CKBox',
            'CKFinder',
            'EasyImage',
            'RealTimeCollaborativeComments',
            'RealTimeCollaborativeTrackChanges',
            'RealTimeCollaborativeRevisionHistory',
            'PresenceList',
            'Comments',
            'TrackChanges',
            'TrackChangesData',
            'RevisionHistory',
            'Pagination',
            'WProofreader',
            'MathType',
            'SlashCommand',
            'Template',
            'DocumentOutline',
            'FormatPainter',
            'TableOfContents',
            'PasteFromOfficeEnhanced'
        ]
    }).then(ckeditor => {
        ckeditor1 = ckeditor
        ckeditor.plugins.get('FileRepository').createUploadAdapter = (loader) => {
            return {
                upload: () => {
                    return loader.file.then(file => {
                        return new Promise((resolve, reject) => {
                            const formData = new FormData();
                            formData.append('file', file);
                            formData.append('csrf', $("input[name='csrf']").val())
                            fetch(url + '/channel/imageupload', {
                                method: 'POST',
                                body: formData
                            })
                                .then(response => response.json())
                                .then(data => {
                                    if (data.success) {
                                        resolve({
                                            default: data.url // URL to the uploaded image

                                        });
                                    } else {
                                        reject(data.error);
                                    }
                                })
                                .catch(error => {
                                    reject(error);
                                });
                        });
                    });
                }
            };
        };
    })
        .catch(error => {
            console.error(error);
        });

}

$('#update-templatebtn').click(function(event){

    var emailtemplate = ckeditor1.getData()

    $('#input-tempname').val($("#emailtemp-name").text())

    $('#ckeditor-data').val(emailtemplate)

    // $('#mailtemplate_form').submit()
})

function setroute(){

    window.location.href="/settings/emails/"
}

function setroute1(){

    window.location.href="/settings/emails/emailconfig/"
}