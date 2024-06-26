var languagedata

selectedcheckboxarr=[]
/** */
$(document).ready(async function () {

  var languagepath = $('.language-group>button').attr('data-path')
  
  await $.getJSON(languagepath, function (data) {
      
      languagedata = data
  })
  
   $('.Content').addClass('checked');

   $('.category-path-list').each(function(){

    var length = $(this).children('.categorylistname').length-1

      $(this).children('.categorylistname').each(function(index){

        if(length == index ){
          
          $(this).next().remove();
        }

      })
  })

});

$(document).keydown(function(event) {
  if (event.ctrlKey && event.key === '/') {
      $(".search").focus().select();
  }
});

$("#save").click(function () {

  jQuery.validator.addMethod("duplicatename", function (value) {

    var result;
    category_id = $("#categoryid").val()

    $.ajax({
      url: "/categories/checksubcategoryname",
      type: "POST",
      async: false,
      data: { "cname": value, "categoryid": category_id, csrf: $("input[name='csrf']").val() },
      datatype: "json",
      caches: false,
      success: function (data) {
        result = data.trim();
      }
    })
    return result.trim() != "true"
  })


  $("#form").validate({

    ignore: [],

    rules: {
      cname: {
        required: true,
        space: true,
        duplicatename: true

      },
      cdesc: {
        required: true,
        space: true,
        // category_desc: true
        maxlength: 250,
       
      },
      subcategoryid: {
        required: true

      }

    },
    messages: {
      cname: {
        required: "* " + languagedata.Categoryy.catnamevalid,
        space: "* " + languagedata.spacergx,
        duplicatename : "*" + languagedata.Categoryy.catnamevailderr

      },
      cdesc: {
        required: "* " + languagedata.Categoryy.catdescvalid,
        space: "* " + languagedata.spacergx,
        maxlength: "* "+languagedata?.Permission?.descriptionchat

      },
      subcategoryid: {
        required: "*" +  languagedata.Categoryy.selectvaild
      }

    }
  });

  var formcheck = $("#form").valid();
  if (formcheck == true) {
    $('#form')[0].submit();

  }
  else {
    Validationcheck()
    $(document).on('keyup', ".field", function () {
      Validationcheck()
    })

    if ($("#cid").hasClass('error')) {
      $('#availablecat').addClass('input-group-error');
  
    } else {
      $('#availablecat').removeClass('input-group-error');
  
    }
  }

})


// Edit functionaltiy
var edit;
$("body").on("click", "#edit", function () {

  $(".input-group").removeClass("input-group-error")
  $("#cname-error").hide()
  $("#cdesc-error").hide()
  $("#cid-error").hide()

  $("#update").show()
  $("#save").hide()


  edit = $(this).closest("tr");
  var desc = edit.find("td:eq(1)").html();
  $("#triggerId").html(desc);
  $("#triggerId").css("gap","5px")

  var url = window.location.search
  const urlpar = new URLSearchParams(url)
  pageno = urlpar.get('page')

  var data = $(this).attr("data-id");
  $("#title").text(languagedata.Categoryy.updatecategory)
  $("#form").attr("action", "");
  $("#form").attr("name", "edit");
  $("#catpageno").val(pageno)
  var id = $("#category_id").val(data)
  $.ajax({
    url: "/categories/editsubcategory",
    type: "GET",
    dataType: "json",
    data: { "id": data },
    success: function (result) {
      if (result.Id != 0) {
        var name = $("#cname").val(result.CategoryName);
        var desc = $("#cdesc").val(result.Description);
        var image = $("#ctimagehide").attr("src", result.ImagePath)

        $("#cname").text(name);
        $("#cdesc").text(desc);
        $("#categoryimage").val(result.ImagePath)
        $("#cid").val(result.ParentId)

        $("#id").val(result.Id)
        $("#category_id").val(result.Id)

        
        var keyword = $("#searchcatlists").val()

        if(keyword != ""){
          $("#searchcatlists").val("")
          $(".noData-foundWrapper").hide()

        }

        if (result.ImagePath != "") {
          $("#browse").hide();
          $("#mediadesc").hide();
          $("#catdel-img").show()

        }else{
          $("#browse").show();
          $("#mediadesc").show();
          $("#catdel-img").hide()

        }
    

        if ("form[name='edit']") {
          $(".drp").each(function (index, element) {

            var id = $(element).attr("data-id")
            var categoryname = $(element).find("p").text().trim()
            var parentid = $(element).attr("data-parentid")


            console.log("id",id,categoryname,parentid,result.Id);

            if (result.Id == id) {
              $(element).hide()
            } else if (result.Id == parentid) {
              $(element).hide()
            } else if (categoryname.indexOf(result.CategoryName) != -1) {
              $(element).hide()
            } else {
              $(element).show()
            }

          })
        }
      } else {
        
        notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Toast.internalserverr + '</span></div>';
        $(notify_content).insertBefore(".breadcrumbs");
        setTimeout(function () {
          $('.toast-msg').fadeOut('slow', function () {
            $(this).remove();
          });
        }, 5000); // 5000 milliseconds = 5 seconds

      }
    }
  })
});


// delete popup
$(document).on('click', '#delete-btn', function () {
  var data = $(this).attr("data-id");
  var parent_id = $("#categoryid").val();
  var del = $(this).closest("tr");
  var url=window.location.search
    const urlpar= new URLSearchParams(url)
    pageno = urlpar.get('page')
  $.ajax({
    url: "/categories/deletepopup",
    type: "GET",
    dataType: "json",
    data: { "id": data },
    success: function (results) {
      if (results.Id == 0) {
        $('#content').text(languagedata.Categoryy.delcategory);
        $(".deltitle").text(languagedata.Categoryy.deltitle)
        $('.delname').text(del.find('td:eq(0)').text())
        $('#delid').show();
        if (pageno == null){
          $('#delete').attr('href', '/categories/removecategory?id=' + data + "&&categoryid=" + parent_id);

        }else{
          $('#delete').attr('href', '/categories/removecategory?id=' + data + "&&categoryid=" + parent_id + "?page=" +pageno);

        }
        $('#btn3').text(languagedata.no);

      } else {
        $('#content').text(languagedata.Categoryy.delcatinvalid);
        $('#delete').attr('href', '');
        $('#delid').hide();
        $('#btn3').text(languagedata.ok)

      }
    }
  });
});


$("#caddbtn , #clickadd").on("click", function () {
  $(".input-group").removeClass("input-group-error")

  var CategoryId = $("#categoryid").val()
  $("#title").text(languagedata.Categoryy.addcategry)

  $("#form").attr("action", "/categories/addsubcategory/" + CategoryId).attr("method", "POST")

  $("#cname").val("");
  $("#cdesc").val("");
  $("#cname-error").hide()
  $("#cdesc-error").hide()
  $("#cid-error").hide()
  $("#triggerId").html("")
  $("#triggerId").text(languagedata.Categoryy.selctcategry)


  $("#ctimagehide").attr("src", "");
  $("#categoryimage").val("")

  $("#save").show()
  $("#browse").show();
  $("#mediadesc").show();
  $("#update").hide()
  $("#catdel-img").hide()

  $("#form").attr("name", "")


  var keyword = $("#searchcatlists").val()

       

  if ("form[name='']") {
    $(".drp").each(function (index, element) {
      if(keyword != ""){
        $("#searchcatlists").val("")
        $(".noData-foundWrapper").hide()

      }
          $(element).show()
    })
  }
});


$('#categoryModal').on('hide.bs.modal', function (event) {
  $('.modal-backdrop').remove()
})

$(document).on("click", "#mediamodalclose", function () {
  $("#categoryModal").modal('show')

})

$('#catdel-img').click(function () {
  $('#categoryimage').val("")
  $('#ctimagehide').attr('src', '')
  $("#mediadesc").css("margin-top","1%")

  $(this).siblings('h3,button').show()
  $(this).hide()

})


$(document).on("click", "#cancel", function () {
  $("#catdel-img").hide();
  $("#cname-error").hide()
  $("#cdesc-error").hide()
  $("#cid-error").hide()
  $('#catname').removeClass('input-group-error');
  $('#catdes').removeClass('input-group-error');
  $('#availablecat').removeClass('input-group-error');
  $('.modal-backdrop').remove()

})


$('#update').click(function () {

  jQuery.validator.addMethod("duplicatename", function (value) {

    var result;
    parent_id = $("#categoryid").val()
    categoryid = $("#category_id").val()

    $.ajax({
      url: "/categories/checksubcategoryname",
      type: "POST",
      async: false,
      data: { "cname": value, "parentid": parent_id,"categoryid":categoryid, csrf: $("input[name='csrf']").val() },
      datatype: "json",
      caches: false,
      success: function (data) {
        result = data.trim();
      }
    })
    return result.trim() != "true"
  })


  $("#form").validate({

    ignore: [],

    rules: {
      cname: {
        required: true,
        duplicatename: true,
        space: true
      },
      cdesc: {
        required: true,
        space: true,
        maxlength: 250,
        // category_desc: true
      },
      subcategoryid: {
        // required: true

      }

    },
    messages: {
      cname: {
        required: "* " + languagedata.Categoryy.catnamevalid,
        space: "* " + languagedata.spacergx,
        duplicatename : "*" + languagedata.Categoryy.catnamevailderr,

      },
      cdesc: {
        required: "* " + languagedata.Categoryy.catdescvalid,
        space: "* " + languagedata.spacergx,
        maxlength: "* "+languagedata?.Permission?.descriptionchat

      },
      subcategoryid: {
        // required: "* Please select the category"
      }

    }
  });

  var formcheck = $("#form").valid();
  if (formcheck == true) {
    $("#form").attr("action", "")

    var id = $("#category_id").val()

    var cname = $("#cname").val();
    var cdesc = $("#cdesc").val();
    var pcategoryid = $("#cid").val();
    console.log("pcategoryid",pcategoryid);
    var parentid = $("#categoryid").val();
    var image = $("#categoryimage").val();
    var pageno = $("#catpageno").val();

  
    if ($("#form").valid()) {
      $.ajax({
        url: "/categories/editsubcategory",
        type: "POST",
        datatype: "json",
        data: { "id": id, "cname": cname, "cdesc": cdesc, "pcategoryid": pcategoryid, "image": image, csrf: $("input[name='csrf']").val() },
        success: function (result) {
          if (result) {

            if (pageno == 0){
              window.location.href = "/categories/addcategory/" + parentid;

            }else{
              window.location.href = "/categories/addcategory/" + parentid + "?page="+pageno;

            }
  
          } else {
           
            notify_content = '<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span>' + languagedata.Toast.internalserverr + '</span></div>';
            $(notify_content).insertBefore(".breadcrumbs");
            setTimeout(function () {
              $('.toast-msg').fadeOut('slow', function () {
                $(this).remove();
              });
            }, 5000); // 5000 milliseconds = 5 seconds
  
          }
  
        }
      })
      $("#submit").show()
      $("#update").hide()
    }

  }
  else {
    Validationcheck()
    $(document).on('keyup', ".field", function () {
      Validationcheck()
    })

  }

})

function Validationcheck() {

  if ($('#cname').hasClass('error')) {
    $('#catname').addClass('input-group-error');
  } else {
    $('#catname').removeClass('input-group-error');
  }

  if ($('#cdesc').hasClass('error')) {
    $('#catdesc').addClass('input-group-error');
  } else {
    $('#catdesc').removeClass('input-group-error');
  }

}

$(document).on('keyup', '#searchsubcategory', function () {

  var categoryid = $("#categoryid").val();

  if (event.key === 'Backspace') {
  if ($('.search').val() == "") {
    window.location.href = "/categories/addcategory/" + categoryid

  }
}
})

$(document).on('click', '#medcancel', function () {
  $("#addnewimageModal").hide()
})

$("#addnewimageModal").on('show.bs.modal', function (event) {
  $(this).css("background-color","rgba(30,41,44,0.46)")
})

$("#categoryModal").on('show.bs.modal', function (event) {
  $(this).css("background-color","rgba(30,41,44,0.46)")
})

// Get dropdown values and id

$(".drp").on("click", function () {

  var id = $(this).attr("data-id")
  $(this).parents('.dropdown-menu').siblings('a').html($(this).html())
  $(this).parent().siblings("input[name='subcategoryid']").val(id)

  if (id == 0) {
    $("#availablecat").addClass('input-group-error')

  }else {
    $("#availablecat").removeClass('input-group-error')
    $("#cid-error").css('display', 'none')
  }
})


// form category search in keyup

$("#searchcatlists").keyup(function () {
  var keyword = $(this).val().trim().toLowerCase()
  $(".catrgory-dropdownlist button").each(function (index, element) {
      var title = $(element).text().toLowerCase()
      var categoryid = $("#category_id").val()
      var currentid = $(element).attr("data-id")
      if (title.includes(keyword) && categoryid != currentid) {
          $(element).show()
          $(".noData-foundWrapper").hide()

      } else {
          $(element).hide()
          if($('.catrgory-dropdownlist button:visible').length==0){
            $(".noData-foundWrapper").show()
        }
      }
  })
})

// description focus 
const GroupDesc = document.getElementById('cdesc');
const inputGroup = document.querySelectorAll('.input-group');

GroupDesc.addEventListener('focus', () => {

  GroupDesc.closest('.input-group').classList.add('input-group-focused');
});
GroupDesc.addEventListener('blur', () => {
  GroupDesc.closest('.input-group').classList.remove('input-group-focused');
});

$(document).on('click','.selectcheckbox',function(){

  categorygrbid =$(this).attr('data-id')


 if ($(this).prop('checked')){

     selectedcheckboxarr.push(categorygrbid)
 
 }else{

     const index = selectedcheckboxarr.indexOf(categorygrbid);
 
     if (index !== -1) {

         console.log(index,"sssss")
         selectedcheckboxarr.splice(index, 1);
     }
    
     $('#Check').prop('checked',false)

 }


 if (selectedcheckboxarr.length !=0){

     $('.selected-numbers').show()

    
     $('.checkboxlength').text(selectedcheckboxarr.length+" " +'items selected')

         $('#seleccheckboxdelete').removeClass('border-end')

         $('.unbulishslt').html("")
     
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
  
         categorygrbid= $(this).attr('data-id')

         selectedcheckboxarr.push(categorygrbid)
      })

      $('.selected-numbers').show()

     
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
       
  $('.deltitle').text("Delete Categories?")

  $('#content').text('Are you sure want to delete selected Categories?')

  }else {

       $('.deltitle').text("Delete Category?")

       $('#content').text('Are you sure want to delete selected Category?')
  }


  $('#delete').addClass('checkboxdelete')
})
//MULTI SELECT DELETE FUNCTION//
$(document).on('click','.checkboxdelete',function(){

  var url = window.location.href;

  console.log("url", url)

  var pageurl = window.location.search

  const urlpar = new URLSearchParams(pageurl)

  pageno = urlpar.get('page')

  var parent_id = $("#categoryid").val();


  $('.selected-numbers').hide()
  $.ajax({
      url: '/categories/multiselectcategorydelete',
      type: 'post',
      dataType: 'json',
      async: false,
      data: {
          "categoryids": selectedcheckboxarr,
          csrf: $("input[name='csrf']").val(),
          "page":pageno,
          "parentid":parent_id

          
      },
      success: function (data) {

          console.log(data,"result")

          if (data.value==true){

              setCookie("get-toast", "Category Deleted Successfully")

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