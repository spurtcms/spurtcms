var languagedata
/** */
$(document).ready(async function () {

     $('.Members').addClass('checked');

     var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })


})

//**description focus function */
const Desc = document.getElementById('membergroup_desc');
const inputgroup = document.querySelectorAll('.input-group');

Desc.addEventListener('focus', () => {

     Desc.closest('.input-group').classList.add('input-group-focused');
});
Desc.addEventListener('blur', () => {
     Desc.closest('.input-group').classList.remove('input-group-focused');
});


// member group edit
var edit;
$("body").on("click", "#editmembergroup", function () {

     $(".input-group").removeClass("input-group-error")

     var url = window.location.search;
     const urlpar = new URLSearchParams(url);
     pageno = urlpar.get('page');

     $("#title").text(languagedata.Members_Group.updmemgrp)
     var data = $(this).attr("data-id");
     edit = $(this).closest("tr");
     $("#membergroup_form").attr("action", "/membersgroup/updategroup")
     var name = edit.find("td:eq(0)").text();
     var desc = edit.find("td:eq(1)").text();
     $("#membergroup_name").val(name);
     $("#membergroup_desc").val(desc);
     $("#membergroup_id").val(data);
     $("#update").show();
     $("#save").hide();

     $("#membergroup_name-error").hide();
     $("#membergroup_desc-error").hide()
     $("#memgrbpageno").val(pageno)
});

// $(document).on('click', "#groupcancel", function () {
//      // $("#groupcancel").hide();
//      $('#grbname').removeClass('validate')
//      $('#grbdes').removeClass('validate')
//      $(".new").show();
//      $(".update").hide();
//      $("#membergroup_name-error").hide();
//      $("#membergroup_name").val("");

// })

/*search */
$(document).on("click", "#filterformsubmit", function () {
     var key = $(this).siblings().children(".search").val();
     if (key == "") {
          window.location.href = "/membersgroup/"
     } else {
          $('.filterform').submit();
     }
})

// new Btn

$(document).on("click", "#add-btn , #clickadd", function () {
     $("#title").text(languagedata.Members_Group.addmembergrp)
     $(".input-group").removeClass("input-group-error")

     $("#membergroup_name-error").hide();
     $("#membergroup_desc-error").hide();
     $("#save").show();
     $("#update").hide();
     $("#membergroup_name").val("");
     $("#membergroup_desc").val("");
     $(".member-update").hide();
     $(".member-save").show();


})
// popup
$(document).on('click', '#delete-btn', function () {

     console.log("text", languagedata.Members_Group.delmemgrp);

     var MemberGroupId = $(this).attr("data-id");
     var del = $(this).closest("tr");

     var url = window.location.search;
     const urlpar = new URLSearchParams(url);
     pageno = urlpar.get('page');

     $.ajax({
          url: "/membersgroup/memberpopup",
          type: "post",
          dataType: "json",
          data: { "id": MemberGroupId, csrf: $("input[name='csrf']").val() },
          success: function (results) {
               $('#content').text(languagedata.Members_Group.delmemgrp);
               $(".deltitle").text(languagedata.Members_Group.delmembergrp)
               $('.delname').text(del.find('td:eq(0)').text())
               $('#delid').show();
               $('#delcancel').text(languagedata.no);

               if (pageno == null) {
                    $('#delid').parent('#delete').attr('href', '/membersgroup/deletegroup?id=' + MemberGroupId);

               } else {
                    $('#delid').parent('#delete').attr('href', '/membersgroup/deletegroup?id=' + MemberGroupId + "&page=" + pageno);

               }


          }
     });
});

function MemberStatus(id) {
     $('#cb' + id).on('change', function () {
          this.value = this.checked ? 1 : 0;
     }).change();
     var isactive = $('#cb' + id).val();

       $.ajax({
          url: '/membersgroup/groupisactive',
          type: 'POST',
          async: false,
          data: {"id": id,"isactive": isactive,csrf: $("input[name='csrf']").val()},
          dataType: 'json',
          cache: false,
          success: function (result) { 
              if (result) {

                    notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>' + languagedata?.Toast?.memgrpstatusnotify + '</span></div>';
                    $(notify_content).insertBefore(".header-rht");
                    setTimeout(function () {
                         $('.toast-msg').fadeOut('slow', function () {
                              $(this).remove();
                         });
                    }, 5000); // 5000 milliseconds = 5 seconds

               } else {
                    
                    notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.internalserverr + '</span></div>';
                    $(notify_content).insertBefore(".breadcrumbs");
                    setTimeout(function () {
                         $('.toast-msg').fadeOut('slow', function () {
                              $(this).remove();
                         });
                    }, 5000); // 5000 milliseconds = 5 seconds

               }
          }
     });
}

const memberGroupName = document.getElementById('membergroup_name');
const memberGroupDesc = document.getElementById('membergroup_desc');
const inputGroup = document.querySelectorAll('.input-group');

memberGroupName.addEventListener('focus', () => {

     memberGroupName.closest('.input-group').classList.add('focus');
});

memberGroupDesc.addEventListener('focus', () => {

     memberGroupDesc.closest('.input-group').classList.add('focus');
});
memberGroupName.addEventListener('blur', () => {
     memberGroupName.closest('.input-group').classList.remove('focus');
});

memberGroupDesc.addEventListener('blur', () => {
     memberGroupDesc.closest('.input-group').classList.remove('focus');
});


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
                    maxlength: "* "+languagedata?.Permission?.descriptionchat
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
                    maxlength: "* "+languagedata.Permission.descriptionchat

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

$(document).on('click', '#back', function () {

     $('#grbname').removeClass('input-group-error')

     $('#grbdes').removeClass('input-group-error')

})

/*search redirect home page */

$(document).on('keyup', '#searchmemgroup', function () {
     if ($('.search').val() === "") {
          window.location.href = "/membersgroup/"
     }
})


$('form[class=filterform]>img').click(function () {

     if ($(this).siblings('input[name=keyword]').val() != "" && $(this).parents('.transitionSearch').hasClass('active')) {

          var keyword = $(this).siblings('input[name=keyword]').val()

          window.location.href = "/member/?keyword=" + keyword
     }
})