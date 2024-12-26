var ckeditor1;
var templatestatus=[]


$(document).ready(function(){
    CKEDITORS()
})

$(document).on('click','.open-emailtemp',function(){
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

    $('#mailtemplate_form').submit()
})

$(document).on('click','#update',function(){


    var templateJSON = JSON.stringify(templatestatus)

    $("#templatestatus").val(templateJSON)
    console.log("templateJSON",templateJSON);

   $('#channelsettingform')[0].submit();


})

$(document).on('click','#cancel',function(){

    window.location.href="/channel/settings"
 })


 $(document).on('click','.templatecb',function(){


   templateid= $(this).attr('data-id')

   var  status

   if($(this).prop('checked')){

    status = '1'
   }else{

    status ='0'
   }

   var existingIndex = templatestatus.findIndex(function(item) {

    return item.templateid === templateid;
  });

  if (existingIndex === -1) {

    templatestatus.push({ "templateid": templateid, "status": status });

  } else {

    templatestatus[existingIndex].status = status;
  }

   console.log(templatestatus,"arraycheckk")

 })