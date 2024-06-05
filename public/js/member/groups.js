var languagedata

var selectedcheckboxarr=[]
/** */
$(document).ready(async function () {

     $('.Members').addClass('checked');

     var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })


    if ($('.selectcheckbox').length === 0) {

        $('.headingcheck').hide();
   }
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
     var name = edit.find("td:eq(1)").text();
     var desc = edit.find("td:eq(2)").text();
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


$(document).on('keyup', '#searchmemgroup', function (event) {
    
     if (event.key === 'Backspace') {
       
         if ($('.search').val() === "") {
         
             window.location.href = "/membersgroup/";
         }
     }
 });
 
 $(document).keydown(function(event) {

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

$(document).on('click','.selectcheckbox',function(){

     memberid =$(this).attr('data-id')

     var status = $(this).parents('td').siblings('td').find('.tgl-light').val();

     console.log(status,"status")

 
    if ($(this).prop('checked')){

        selectedcheckboxarr.push({"memberid":memberid,"status":status})
    
    }else{

        const index = selectedcheckboxarr.findIndex(item => item.memberid === memberid);
    
        if (index !== -1) {

            console.log(index,"sssss")
            selectedcheckboxarr.splice(index, 1);
        }
       
        $('#Check').prop('checked',false)

    }

   
    if (selectedcheckboxarr.length !=0){

        $('.selected-numbers').show()

        var allSame = selectedcheckboxarr.every(function(item) {
            return item.status === selectedcheckboxarr[0].status;
        });
        
        var setstatus
        var img;

           if (selectedcheckboxarr[0].status === '1') {
  
             setstatus = "Deactive";
 
              img = "/public/img/In-Active (1).svg";

        } else if (selectedcheckboxarr[0].status === '0') {

               setstatus = "Active";

               img = "/public/img/Active (1).svg";

          } 

           var htmlContent = '';

             if (allSame) {

              htmlContent = '<img src="' + img + '">' + setstatus;

              } else {

                htmlContent = '';

              }

            $('#unbulishslt').html(htmlContent);
       
        $('.checkboxlength').text(selectedcheckboxarr.length+" " +'items selected')

        if(!allSame){

            $('#seleccheckboxdelete').removeClass('border-end')

            $('.unbulishslt').html("")
        }else{

            $('#seleccheckboxdelete').addClass('border-end')
        }

       
    }else{

        $('.selected-numbers').hide()
    }

        var allChecked = true;

         $('.selectcheckbox').each(function() {

             if (!$(this).prop('checked')) {

             allChecked = false;

            return false; 
         }
    });

          $('#Check').prop('checked', allChecked);

     console.log(selectedcheckboxarr,"checkkkk")
 })

 //ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click','#Check',function(){

     selectedcheckboxarr=[]
 
     var isChecked = $(this).prop('checked');
 
     if (isChecked){
 
         $('.selectcheckbox').prop('checked', isChecked);
 
         $('.selectcheckbox').each(function(){
     
            memberid= $(this).attr('data-id')
 
            var status = $(this).parents('td').siblings('td').find('.tgl-light').val();
     
            selectedcheckboxarr.push({"memberid":memberid,"status":status})
         })
 
         $('.selected-numbers').show()
 
         var allSame = selectedcheckboxarr.every(function(item) {

             return item.status === selectedcheckboxarr[0].status;
         });
         
         var setstatus
 
         var img

         if (selectedcheckboxarr.length !=0){
         if (selectedcheckboxarr[0].status === '1') {
 
          setstatus = "Deactive";

           img = "/public/img/In-Active (1).svg";

     } else if (selectedcheckboxarr[0].status === '0') {

            setstatus = "Active";

            img = "/public/img/Active (1).svg";

       } 

     }
         var htmlContent = '';
 
         if (allSame) {
 
          htmlContent = '<img src="' + img + '">' + setstatus;
 
          $('#seleccheckboxdelete').addClass('border-end')
 
          } else {
 
            htmlContent = '';
 
            $('#seleccheckboxdelete').removeClass('border-end')
 
          }
 
        $('#unbulishslt').html(htmlContent);
 
         $('.checkboxlength').text(selectedcheckboxarr.length+" " +'items selected')
     
     }else{
 
 
         selectedcheckboxarr=[]
 
         $('.selectcheckbox').prop('checked', isChecked);
 
         $('.selected-numbers').hide()
     }
     if (selectedcheckboxarr.length ==0){

          $('.selected-numbers').hide()
      }
 
 })
 
 $(document).on('click','#seleccheckboxdelete',function(){

     if (selectedcheckboxarr.length>1){
          
     $('.deltitle').text("Delete Membergroups?")
 
     $('#content').text('Are you sure want to delete selected Membergroups?')

     }else {

          $('.deltitle').text("Delete Membergroup?")
 
          $('#content').text('Are you sure want to delete selected Membergroup?')
     }
 
 
     $('#delete').addClass('checkboxdelete')
 })
 
 $(document).on('click','#unbulishslt',function(){

     if (selectedcheckboxarr.length>1){

          $('.deltitle').text( $(this).text()+" "+"Membergroups?")
 
          $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Membergroups?")
     }else{

          $('.deltitle').text( $(this).text()+" "+"Membergroup?")
 
          $('#content').text("Are you sure want to " +$(this).text()+" "+"selected Membergroup?")
     }
 
     $('#delete').addClass('selectedunpublish')
 
 })

 //MULTI SELECT DELETE FUNCTION//
$(document).on('click','.checkboxdelete',function(){

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
             "page":pageno
 
             
         },
         success: function (data) {
 
             console.log(data,"result")
 
             if (data.value==true){
 
                 setCookie("get-toast", "Member Group Deleted Successfully")
 
                 window.location.href=data.url
             }else{
 
                 setCookie("Alert-msg", "Internal Server Error")
 
             }
 
         }
     })
 
 })
 //Deselectall function//
 
 $(document).on('click','#deselectid',function(){
 
     $('.selectcheckbox').prop('checked',false)
 
     $('#Check').prop('checked',false)
 
     selectedcheckboxarr=[]
 
     $('.selected-numbers').hide()
     
 })
 
 //multi select active and deactive function//
 
 $(document).on('click','.selectedunpublish',function(){
 
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
             "memberids":JSON.stringify(selectedcheckboxarr),
             csrf: $("input[name='csrf']").val(),
             "page":pageno
 
             
         },
         success: function (data) {
 
             console.log(data,"result")

             if (data.value==true){
 
                 setCookie("get-toast", "memgrpstatusnotify")
 
                 window.location.href=data.url
             }else{
 
                 setCookie("Alert-msg", "Internal Server Error")
 
             }
 
         }
     })
 
 
 })