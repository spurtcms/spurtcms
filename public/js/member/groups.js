var languagedata

var selectedcheckboxarr = []


$(document).ready(async function () {

     $('.Members').addClass('checked');

     var languagepath = $('.language-group>button').attr('data-path')

     await $.getJSON(languagepath, function (data) {

          languagedata = data
          console.log(languagedata);
     })

     if ($('.selectcheckbox').length === 0) {

          $('.headingcheck').hide();
     }
})


$(document).ready(function(){
     var $textarea = $('#membergroup_desc');
     var $errorMessage = $('#membergroup_desc-error');
     var maxLength = 250;

     $textarea.on('input', function () {
          if ($(this).val().length >= maxLength) {
               // Show error message
               $errorMessage.text(languagedata.Permission.descriptionchat);
          } else {
               // Clear error message if under the limit
               $errorMessage.text('');
          }
     });


     $('#membergroup_desc').on('input', function () {

          let lines = $(this).val().split('\n').length;

          if (lines > 5) {

               let value = $(this).val();

               let linesArray = value.split('\n').slice(0, 5);

               $(this).val(linesArray.join('\n'));
          }
     });
})


// save btn function
$('#save').click(function () {

     jQuery.validator.addMethod("duplicategrb", function (value) {

          var result;
          id = $("#membergroup_id").val()
          console.log("check", id)
          $.ajax({
               url: "/membersgroup/checknameinmembergrp",
               type: "POST",
               async: false,
               data: { "membergroup_name": value, "membergroup_id": id, csrf: $("input[name='csrf']").val() },
               datatype: "json",
               caches: false,
               success: function (data) {
                    result = data.trim();
               }
          })
          return result.trim() != "true"
     })


     // form validation

     $("#membergroup_form").validate({
          rules: {
               membergroup_name: {
                    required: true,
                    space: true,
                    duplicategrb: true

               },
               membergroup_desc: {
                    required: true,
                    space: true,
                    maxlength: 250,
               }
          },
          messages: {
               membergroup_name: {
                    required: "* " + languagedata.Members_Group.memgrpnamevalid,
                    space: "* " + languagedata.spacergx,
                    duplicategrb: "*" + languagedata.Members_Group.memgrpnamechk
               },
               membergroup_desc: {
                    required: "* " + languagedata.Members_Group.memgrpdescvalid,
                    space: "* " + languagedata.spacergx,
                    maxlength: "* " + languagedata.Permission.descriptionchat
               }

          },

     });


     var formcheck = $("#membergroup_form").valid();

     if (formcheck == true) {
          $('#membergroup_form')[0].submit();
     }
     else {
          Validationcheck()
          $(document).on('keyup', ".field", function () {
               Validationcheck()
          })

     }

     return false
})





// member group edit
var edit;

$(".editmembergroup").click(function () {
     console.log("edit member");

     $(".input-group").removeClass("input-group-error")

     var url = window.location.search;
     const urlpar = new URLSearchParams(url);
     pageno = urlpar.get('page');

     $("#title").text(languagedata.Members_Group.updmemgrp)
     var data = $(this).attr("data-id");
     edit = $(this).closest("tr");     
     $("#membergroup_form").attr("action", "/membersgroup/updategroup")
     var name = edit.find("td:eq(1)").text().trim()
     var desc = edit.find("td:eq(2)").text().trim()

     $("#membergroup_name").val(name);
     $("#membergroup_desc").val(desc);
     $("#membergroup_id").val(data);
     $("#update").show();
     $("#save").hide();

     $("#membergroup_name-error").hide();
     $("#membergroup_desc-error").empty();
     $("#memgrbpageno").val(pageno);
});


// cancel btn function

$('#groupcalcelbtn').on('click', function () {
     $("#membergroup_form").attr("action", "")
     $("#membergroup_name").val("");
     $("#membergroup_desc").val("");
})


// new membergroup model open function

$(document).on("click", "#add-btn , #clickadd", function () {
     $("#title").text(languagedata.Members_Group.addmembergrp)
     $("#membergroup_form").attr("action", "/membersgroup/newgroup")
     $(".input-group").removeClass("input-group-error")

     $("#membergroup_name-error").hide();
     $("#membergroup_desc-error").empty();
     $("#save").show();
     $("#update").hide();
     $("#membergroup_name").val("");
     $("#membergroup_desc").val("");
     $(".member-update").hide();
     $(".member-save").show();
     $(".lengthErr").addClass("hidden");
})


// delete model popup function

$(document).on('click', '#delete-btn', function () {

     var MemberGroupId = $(this).attr("data-id");

     var del = $(this).closest("tr");
     $.ajax({
          url: '/membersgroup/chkmemgrphavemember',
          type: 'POST',
          async: false,
          data: { "id": MemberGroupId,csrf: $("input[name='csrf']").val()},
          dataType: 'json',
          success: function (data){
               if(data.value){
                    $('#delid').addClass("hidden");
                    $('#dltCancelBtn').text(languagedata.ok);
                    $("#content").text(languagedata.Members_Group.membergrprestrictmsg)
               }else{
                    $('#delid').removeClass("hidden")
                    $('#content').text(languagedata.Members_Group.delmemgrp);
                    $('#dltCancelBtn').text(languagedata.cancel);

               }
          }
     })
     var url = window.location.search;
     const urlpar = new URLSearchParams(url);
     pageno = urlpar.get('page');

     $(".deltitle").text(languagedata.Members_Group.delmembergrp)
     $('.delname').text(del.find('td:eq(1)').text())
     $('#delcancel').text(languagedata.no);

     if (pageno == null) {

          $('#delid').attr('href', '/membersgroup/deletegroup?id=' + MemberGroupId);

     } else {
          $('#delid').attr('href', '/membersgroup/deletegroup?id=' + MemberGroupId + "&page=" + pageno);

     }

});


// delete model cancel function

$("#deleteModal").on("hide.bs.modal", function () {
     $('#delid').removeClass('checkboxdelete');
     $('#delid').removeClass('selectedunpublish');
     $('.delname').text("");
     $("#delid").attr('href', '');

})

// member group status function

function MemberStatus(id) {
     $('#cb' + id).on('change', function () {
          this.value = this.checked ? 1 : 0;
     }).change();
     var isactive = $('#cb' + id).val();

     $.ajax({
          url: '/membersgroup/groupisactive',
          type: 'POST',
          async: false,
          data: { "id": id, "isactive": isactive, csrf: $("input[name='csrf']").val() },
          dataType: 'json',
          cache: false,
          success: function (result) {
               if (result) {

                    notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#278E2B] bg-[#E2F7E3]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-success.svg" alt = "toast success"></div> <div> <h3 class="text-[#278E2B] text-normal leading-[17px] font-normal mb-[5px] ">Success</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " > ${languagedata.Toast.MemberGroupStatusUpdatedSuccessfully} </p ></div ></div ></li></ul> `;
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                         $('.toast-msg').fadeOut('slow', function () {
                              $(this).remove();
                         });
                    }, 5000); // 5000 milliseconds = 5 seconds

               } else {

                    notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
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



// const memberGroupName = document.getElementById('membergroup_name');
// const memberGroupDesc = document.getElementById('membergroup_desc');
// const inputGroup = document.querySelectorAll('.input-group');


// memberGroupName.addEventListener('focus', () => {

//      memberGroupName.closest('.input-group').classList.add('focus');
// });

// memberGroupDesc.addEventListener('focus', () => {

//      memberGroupDesc.closest('.input-group').classList.add('focus');
// });
// memberGroupName.addEventListener('blur', () => {
//      memberGroupName.closest('.input-group').classList.remove('focus');
// });

// memberGroupDesc.addEventListener('blur', () => {
//      memberGroupDesc.closest('.input-group').classList.remove('focus');
// });

// update function

$('#update').click(function () {


     jQuery.validator.addMethod("duplicategrb", function (value) {

          var result;
          id = $("#membergroup_id").val()
          console.log("check", id)
          $.ajax({
               url: "/membersgroup/checknameinmembergrp",
               type: "POST",
               async: false,
               data: { "membergroup_name": value, "membergroup_id": id, csrf: $("input[name='csrf']").val() },
               datatype: "json",
               caches: false,
               success: function (data) {
                    result = data.trim();
               }
          })
          return result.trim() != "true"
     })


     $("#membergroup_form").validate({
          rules: {
               membergroup_name: {
                    required: true,
                    space: true,
                    duplicategrb: true

               },
               membergroup_desc: {
                    required: true,
                    space: true,
                    maxlength: 250,
               }
          },
          messages: {
               membergroup_name: {
                    required: "* " + languagedata.Members_Group.memgrpnamevalid,
                    space: "* " + languagedata.spacergx,
                    duplicategrb: "*" + languagedata.Members_Group.memgrpnamechk
               },
               membergroup_desc: {
                    required: "* " + languagedata.Members_Group.memgrpdescvalid,
                    space: "* " + languagedata.spacergx,
                    maxlength: "* " + languagedata.Permission.descriptionchat

               }

          },

     });

     var formcheck = $("#membergroup_form").valid();
     if (formcheck == true) {
          $('#membergroup_form')[0].submit();
     }
     else {
          Validationcheck()
          $(document).on('keyup', ".field", function () {
               Validationcheck()
          })
     }

     return false

})


function Validationcheck() {

     if ($('#membergroup_name').hasClass('error')) {
          $('#grpname').addClass('input-group-error');
     } else {
          $('#grpname').removeClass('input-group-error');
     }

     if ($('#membergroup_desc').hasClass('error')) {
          $('#grpdesc').addClass('input-group-error');
     } else {
          $('#grpdesc').removeClass('input-group-error');
     }
}



// search functions strat //

/*search redirect home page */

$(document).on('keyup', '#searchmemgroup', function (event) {
    const searchInput = $(this).val();

    if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {

        if ($(this).val() == "") {

            window.location.href = "/membersgroup/";
        }
    }

    $('.searchClosebtn').toggleClass('hidden', searchInput === "");

})

$(document).on("click", ".Closebtn", function () {
     $(".search").val('')
     $(".Closebtn").addClass("hidden")
     $(".srchBtn-togg").removeClass("pointer-events-none")

   })

   $(document).on("click", ".searchClosebtn", function () {
     $(".search").val('')
     window.location.href = "/membersgroup/"
   })

   $(document).ready(function () {

     $('.search').on('input', function () {

         if ($(this).val().length >= 1) {
             $(".Closebtn").removeClass("hidden")
             $(".srchBtn-togg").addClass("pointer-events-none")

         }else{
           $(".Closebtn").addClass("hidden")
           $(".srchBtn-togg").removeClass("pointer-events-none")

         }
     });
   })

   $(document).on("click", ".hovericon", function () {
     $(".search").val('')
     $(".Closebtn").addClass("hidden")
   })

// search functions end //



$(document).keydown(function (event) {

     if (event.ctrlKey && event.key === '/') {

          $("#searchmemgroup").focus().select();
     }
});


$('form[class=filterform]>img').click(function () {

     if ($(this).siblings('input[name=keyword]').val() != "" && $(this).parents('.transitionSearch').hasClass('active')) {

          var keyword = $(this).siblings('input[name=keyword]').val()

          window.location.href = "/member/?keyword=" + keyword
     }
})



// checkbox select function

$(document).on('click', '.selectcheckbox', function () {

     memberid = $(this).attr('data-id')

     var status = $(this).parents('td').siblings('td').find('.tgl-light').val();

     var hasDelId = $(this).parents('td').siblings("td").find('#delete-btn').length > 0;
     var statusid = $(this).parents('td').siblings('td').find('.tgl-light').length > 0

     console.log(statusid,"statuiddd")


     if (hasDelId) {

         $('#seleccheckboxdelete').show()
     } else {
         $('#seleccheckboxdelete').hide()
     }

     if (statusid) {
         console.log("chekckk")
         $('#unbulishslt').show()

     } else {
         $('#unbulishslt').hide()

         $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')
     }


     if ($(this).prop('checked')) {

          selectedcheckboxarr.push({ "memberid": memberid, "status": status })

     } else {

          const index = selectedcheckboxarr.findIndex(item => item.memberid === memberid);

          if (index !== -1) {

               selectedcheckboxarr.splice(index, 1);
          }

          $('#Check').prop('checked', false)

     }

     if (selectedcheckboxarr.length != 0) {
          $('.selected-numbers').removeClass("hidden")

          if (selectedcheckboxarr.length == 1) {
               $('#deselectid').text(languagedata.deselect)
          } else if (selectedcheckboxarr.length > 1) {
               $('#deselectid').text(languagedata.deselectall)
          }

          var allSame = selectedcheckboxarr.every(function (item) {
               return item.status === selectedcheckboxarr[0].status;
          });

          var setstatus
          var img;

          if (selectedcheckboxarr[0].status === '1') {

               setstatus = languagedata.Memberss.deactive;

               img = "/public/img/In-Active.svg";

          } else if (selectedcheckboxarr[0].status === '0') {

               setstatus = languagedata.Memberss.active;

               img = "/public/img/Active.svg";

          }

          var htmlContent = '';

          if (allSame) {

               htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >'  + '<span class="max-sm:hidden @[550px]:inline-block hidden">'+setstatus+'</span>';

          } else {

               htmlContent = '';

          }

          $('#unbulishslt').html(htmlContent);

          var items

          if (selectedcheckboxarr.length==1){

              items ="Item Selected"
          }else{

              items = languagedata.itemselected
          }
          $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)

          if (!allSame || !statusid) {

               $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')

               $('.unbulishslt').html("")
          } else {

               $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')
          }


     } else {
          $('.selected-numbers').addClass("hidden")
     }

     var allChecked = true;

     $('.selectcheckbox').each(function () {

          if (!$(this).prop('checked')) {

               allChecked = false;

               return false;
          }
     });

     $('#Check').prop('checked', allChecked);

     console.log(selectedcheckboxarr, "checkkkk")
})



//  //ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click', '#Check', function () {

     selectedcheckboxarr = []

     var isChecked = $(this).prop('checked');

     var statusid =$(this).parents('th').siblings('.statushead').length > 0;

     var dperid =$("#dperid").val()

     if (statusid){

          $('#unbulishslt').show();
     }else{
          $('#unbulishslt').hide();
          $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')
     }

     if (dperid=="true"){

          $("#seleccheckboxdelete").show()
     }else{
          $('#seleccheckboxdelete').hide();
     }

     if (isChecked) {

          $('.selectcheckbox').prop('checked', isChecked);

          $('.selectcheckbox').each(function () {

               memberid = $(this).attr('data-id')

               var status = $(this).parents('td').siblings('td').find('.tgl-light').val();

               console.log(status, "state");

               selectedcheckboxarr.push({ "memberid": memberid, "status": status })
          })

          if (selectedcheckboxarr.length != 0) {
               const deselectText = selectedcheckboxarr.length == 1 ? languagedata.deselect : languagedata.deselectall;
               $('#deselectid').text(deselectText);
               $('.selected-numbers').removeClass("hidden")
          } else {
               $('.selected-numbers').addClass("hidden")
          }
          var allSame = selectedcheckboxarr.every(function (item) {

               return item.status === selectedcheckboxarr[0].status;
          });

          console.log(allSame, "allsome");

          var setstatus

          var img
          if (selectedcheckboxarr[0].status === '1') {

               setstatus = languagedata.Memberss.deactive;

               img = "/public/img/In-Active.svg";

          } else if (selectedcheckboxarr[0].status === '0') {

               setstatus = languagedata.Memberss.active;

               img = "/public/img/Active.svg";
          }

          var htmlContent = '';

          if (!allSame || !statusid) {

               htmlContent = '';

               $('#seleccheckboxdelete').removeClass('border-r border-[#717171] mr-[8px] pr-[8px]')


          } else {


               htmlContent = '<img style="width: 14px; height: 14px;" src="' + img + '" >'  + '<span class="max-sm:hidden @[550px]:inline-block hidden">'+setstatus+'</span>';

               $('#seleccheckboxdelete').addClass('border-r border-[#717171] mr-[8px] pr-[8px]')

          }

          $('#unbulishslt').html(htmlContent);

          $('.checkboxlength').text(selectedcheckboxarr.length + " " + languagedata.Memberss.itemsselected)

     } else {
          $('.selected-numbers').addClass("hidden")

          selectedcheckboxarr = []

          $('.selectcheckbox').prop('checked', isChecked);
     }

})


// delete model content upadte

$(document).on('click','#seleccheckboxdelete', function () {

     $.ajax({
          url: '/membersgroup/chkmemgrphavemember',
          type: 'POST',
          async: false,
          data: { "membergrpids": JSON.stringify(selectedcheckboxarr),
                   csrf: $("input[name='csrf']").val(),},
          dataType: 'json',
          success: function (data){
               if(data.value){

                    if (selectedcheckboxarr.length > 1) {

                         $('.deltitle').text(languagedata.Members_Group.deletemembergroups)

                         $('#content').text(languagedata.Members_Group.membergrprestrictmsgs)

                    } else {

                         $('.deltitle').text(languagedata.Members_Group.deletemembergroup)

                         $('#content').text(languagedata.Members_Group.membergrprestrictmsg)
                    }

                     $('#delid').addClass("hidden");
                    $('#dltCancelBtn').text(languagedata.ok);

               }else{

                    if (selectedcheckboxarr.length > 1) {

                         $('.deltitle').text(languagedata.Members_Group.deletemembergroups)

                         $('#content').text(languagedata.Members_Group.deletemembergroupscontents)

                    } else {

                         $('.deltitle').text(languagedata.Members_Group.deletemembergroup)

                         $('#content').text(languagedata.Members_Group.deletemembergroupscontent)
                    }
                    $('#delid').removeClass("hidden")
                    $('#dltCancelBtn').text(languagedata.cancel);

               }
          }
     })



     $("#delid").text($(this).text());
     $('#delid').addClass('checkboxdelete')
})


// status model content update

$(document).on('click', '#unbulishslt', function () {

     if (selectedcheckboxarr.length > 1) {

          $('.deltitle').text($(this).text() + " " + languagedata.Members_Group.membergroups)

          $('#content').text(languagedata.Members_Group.areyousurewantto + $(this).text() + " " + languagedata.Members_Group.selectedmembergroups)
     } else {

          $('.deltitle').text($(this).text() + " " + languagedata.Members_Group.membergroup)

          $('#content').text(languagedata.Members_Group.areyousurewantto + $(this).text() + " " + languagedata.Members_Group.selectedmembergroup)
     }
     $("#delid").text($(this).text());
     $('#delid').addClass('selectedunpublish')

})


//MULTI SELECT DELETE FUNCTION//

$(document).on('click', '.checkboxdelete', function () {

     var url = window.location.href;

     console.log("url", url)

     var pageurl = window.location.search

     const urlpar = new URLSearchParams(pageurl)

     pageno = urlpar.get('page')

     $('.selected-numbers').hide()
     $.ajax({
          url: '/membersgroup/deleteselectedmembergroup',
          type: 'post',
          dataType: 'json',
          async: false,
          data: {
               "memberids": JSON.stringify(selectedcheckboxarr),
               csrf: $("input[name='csrf']").val(),
               "page": pageno


          },
          success: function (data) {

               console.log(data, "result")

               if (data.value == true) {

                    setCookie("get-toast", "Member Group Deleted Successfully")

                    window.location.href = data.url
               } else {

                    window.location.href = data.url

               }

          }
     })

})


//Deselectall function//

$(document).on('click', '#deselectid', function () {

     $('.selectcheckbox').prop('checked', false)

     $('#Check').prop('checked', false)

     selectedcheckboxarr = []

     $('.selected-numbers').addClass("hidden")

})


//multi select active and deactive function//

$(document).on('click', '.selectedunpublish', function () {

     var url = window.location.href;

     console.log("url", url)

     var pageurl = window.location.search

     const urlpar = new URLSearchParams(pageurl)

     pageno = urlpar.get('page')

     $('.selected-numbers').hide()
     $.ajax({
          url: '/membersgroup/multiselectmembergroup',
          type: 'post',
          dataType: 'json',
          async: false,
          data: {
               "memberids": JSON.stringify(selectedcheckboxarr),
               csrf: $("input[name='csrf']").val(),
               "page": pageno


          },
          success: function (data) {

               console.log(data, "result")

               if (data.value == true) {

                    setCookie("get-toast", "memgrpstatusnotify")

                    window.location.href = data.url
               } else {

                    setCookie("Alert-msg", "Internal Server Error")

               }

          }
     })


})


// Group name limit of 25 char

$(document).on('keyup', '.checklength', function () {

     var inputVal = $(this).val()

     var inputLength = inputVal.length

     if (inputLength == 25) {
          $(this).siblings('.lengthErr').removeClass('hidden')
     } else {
          $(this).siblings('.lengthErr').addClass('hidden')
     }

})
