// var quill;
var ckeditor1;
var userids =[]
var templatestatus=[]

$(document).ready(function(){
    CKEDITORS()
})


$(document).on('click','.open-emailtemp',function(){

console.log("email tem");
    $('#mailtemplate_form').attr('action','/member/settings/updatetemplate')
    console.log("checkclick")
    tempid = $(this).attr('data-id')
    $('#userid').val(tempid)
    var url = window.location.search;
    const urlpar = new URLSearchParams(url);
    pageno = urlpar.get('page');
    $("#pageno").val(pageno)
    $.ajax({
      url: "/member/settings/edittemplate",
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
          $("#emailname").val(result.TemplateName)
          $("#emailsub").val(result.TemplateSubject);
          $("#emaildesc").val(result.TemplateDescription)
          ckeditor1.setData(result.TemplateMessage)
        // quill.clipboard.dangerouslyPasteHTML(result.TemplateMessage);
       }
    }
    })

})


function CKEDITORS() {

    var url = $('#urlid').val();

    CKEDITOR.ClassicEditor.create(document.getElementById("text"), {
        toolbar: {
            // items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'blockQuote', 'imageUpload', 'numberedList', 'bulletedList', 'horizontalLine', 'link', 'code'],
            items: ['heading', 'bold', 'italic', 'alignment', 'underline', 'numberedList', 'bulletedList', 'link'],
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
        // placeholder: languagedata.Channell.plckeditor,
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



// function Quill() {

//     // const toolbarOptions = [
//     //     ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
//     //     ['blockquote', 'code-block'],
//     //     ['link', 'image', 'video', 'formula'],

//     //     [{ 'header': 1 }, { 'header': 2 }],               // custom button values
//     //     [{ 'list': 'ordered'}, { 'list': 'bullet' }, { 'list': 'check' }],
//     //     [{ 'script': 'sub'}, { 'script': 'super' }],      // superscript/subscript
//     //     [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent
//     //     [{ 'direction': 'rtl' }],                         // text direction

//     //     [{ 'size': ['small', false, 'large', 'huge'] }],  // custom dropdown
//     //     [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

//     //     [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
//     //     [{ 'font': [] }],
//     //     [{ 'align': [] }],

//     //     ['clean']                                         // remove formatting button
//     //   ];

//       const toolbarOptions = [
//         [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

//         [{ 'size': ['small', false, 'large'] }],  // custom dropdown

//         ['bold', 'italic', 'underline'],        // toggled buttons

//         [{ 'list': 'ordered'}, { 'list': 'bullet' }],

//         [{ 'align': [] }],

//         ['link']
//       ];

//        quill = new Quill('#text', {
//         modules: {
//           toolbar: toolbarOptions
//         },
//         theme: 'snow'
//       });

// }


//  using quill update

// $('#update-templatebtn').click(function(event){
//     var htmlContent = quill.root.innerHTML;

//     var emailtemplate = htmlContent;

//     // $('#input-tempname').val($("#emailtemp-name").text())

//     $('#ckeditor-data').val(emailtemplate)

//     $('#mailtemplate_form').submit()
// })


$(document).on('click','#update',function(){
    console.log("log update");

     $('#multiselectuser').val(userids)

     var templateJSON = JSON.stringify(templatestatus)

     $("#templatestatus").val(templateJSON)

    $('#membersettingform')[0].submit();


})

$(document).on('click','.allowregis',function(){



    if ($(this).prop('checked')){


        $('.allowregistration').val('1')
    }else{

        $('.allowregistration').val('0')
    }
})

$(document).on('click','#check-1',function(){


   option= $(this).siblings('label').text()

   console.log(option,"optionval")

    if ($(this).prop('checked')){

        $('.memberlogin').val(option)
    }
})

$(document).on('click','#check-2',function(){


    option= $(this).siblings('label').text()

    console.log(option,"optionval")

     if ($(this).prop('checked')){

         $('.memberlogin').val(option)
     }
 })



 $(document).on('click','.checkboxid',function(){

    $('.memberdropdown').addClass('show').css({
        position: 'absolute',
        inset: '0px auto auto 0px',
        margin: '0px',
        transform: 'translate(-1px, 65px)'
    });

   var  userid= $(this).attr('data-id')

   if ($(this).prop('checked')){

    userids.push(userid)

   }else{

    const index = userids.indexOf(userid);

    if (index !== -1) {

             userids.splice(index, 1);
    }

   }

console.log(userids,"useridsss")
 })

 $(document).ready(function(){


    if (!userids.includes('1')) {

        userids.push('1');
    }

   var userid= $('#multiselectuser').val()

   if (userid != "") {

    var values = userid.split(',');

    for (var i = 0; i < values.length; i++) {

        if (values[i] !== "1") {

            userids.push(values[i]);
        }


    }

    $('.checkboxdiv').each(function () {

        if (userid.includes($(this).find('input').attr('data-id'))) {

            $(this).find('input').prop('checked', true);
        }
    })
}
 })

 $(document).on('click','#cancel',function(){

    window.location.href="/member/settings"
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

$(document).on('click', '.closemember', function () {

    $('.memberdropdown').addClass('show').css({
        position: 'absolute',
        inset: '0px auto auto 0px',
        margin: '0px',
        transform: 'translate(-1px, 65px)'
    });

    $('#searchdropdownrole').val('')

    $(this).hide()

   $('.checkboxdiv').show()

   $("#nodatafounddesign").addClass("hidden")


})



// dropdown filter input box search
$("#searchdropdownrole").keyup(function () {


    var keyword = $(this).val().trim().toLowerCase()


    if (keyword !== "") {

        $('.closemember').show()

    }else{

        $('.closemember').hide()
    }


    $(".checkboxdiv").each(function (index,element) {

        var title = $(this).find('p').text().toLowerCase();

        if (title.includes(keyword)) {

            $(element).show()
            $("#nodatafounddesign").addClass("hidden")


        } else {

            $(element).hide()

            if ($('.checkboxdiv:visible').length == 0) {
                $("#nodatafounddesign").removeClass("hidden")

            }
        }
    })
})



$(document).on('click','#update-templatebtn',function(){

    console.log("checking")

    $("form[name='emailtemplateform']").validate({
      rules: {
        tempname:{
            required: true,
            space: true,
        },

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
        tempname:{
            required: "* Please enter Template Name" ,
            space: "* " + languagedata.spacergx,

        },
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

        if ($('#emailname').hasClass('error')){

            $('#namegrb').addClass('input-group-error');
        }else{
            $('#namegrb').removeClass('input-group-error');
        }

        if ($('#emailsub').hasClass('error')){

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

      $('.input-group').removeClass('input-group-error')

      $('.field').removeClass('error')

      $('.error').hide()

    })