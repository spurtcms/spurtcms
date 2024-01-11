var languagedata
/** */
$(document).ready(async function () {

    var languagecode = $('.language-group>button').attr('data-code')
  
    await $.getJSON("/locales/"+languagecode+".json", function (data) {
        
        languagedata = data
    })
  
   $('.Content').addClass('checked');

});

$("#save").click(function () {

  jQuery.validator.addMethod("duplicatename", function (value) {

    var result;
    category_id = $("#category_id").val()

    $.ajax({
      url: "/categories/checkcategoryname",
      type: "POST",
      async: false,
      data: { "category_name": value, "category_id": category_id, csrf: $("input[name='csrf']").val() },
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
        space: true
        // category_desc: true
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
        space: "* " + languagedata.spacergx

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
      console.log("sd", result);
      if (result.Id != 0) {
        var name = $("#cname").val(result.CategoryName);
        var desc = $("#cdesc").val(result.Description);
        var image = $("#ctimagehide").attr("src", result.ImagePath)

        $("#cname").text(name);
        $("#cdesc").text(desc);
        $("#categoryimage").val(result.ImagePath)



        if (image != "") {
          $("#browse").hide();
          $("#mediadesc").hide();
          $("#catdel-img").show();

        }
        //   if ($("#categoryimage").attr("src") === "") {
        //     $("#uploadimage").show();
        //     $("#choosebtn").show();
        //     // $("#categoryimage").hide();
        //     $("#catdel-img").hide();
        // } else {
        //     $("#uploadimage").hide();
        //     $("#choosebtn").hide();
        //     $("#categoryimage").show();
        //     $("#catdel-img").show();
        // }

        // if ("form[name='edit']") {
        //   $(".categorypdiv").each(function (index, element) {

        //     var id = $(element).attr("data-id")
        //     var categoryname = $(element).find("p").text().trim()
        //     var parentid = $(element).attr("data-parentid")

        //     if (result.Id == id) {
        //       $(element).hide()
        //     } else if (result.Id == parentid) {
        //       $(element).hide()
        //     } else if (categoryname.indexOf(result.CategoryName) != -1) {
        //       $(element).hide()
        //     } else {
        //       $(element).show()
        //     }

        //   })
        // }

        // $(".radbtn").each(function (index, element) {
        //   var Id = $(element).attr("data-id");
        //   if (result.ParentId == Id) {
        //     $(element).prop("checked", true)
        //   }
        // })
      } else {
        // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>'+languagedata.internalserverr+'</span></div>';
        // $(notify_content).insertBefore(".breadcrumbs");
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

// $("#addnewcategoriesModal").on('hidden.bs.modal', function () {
//   $("#submit").show()
//   $("#update").hide()
//   $("#cname").val("")
//   $("#cdesc").val("")
//   $("#cname-error").hide()
//   $("#cdesc-error").hide()

// })

// delete popup
$(document).on('click', '#delete-btn', function () {
  var data = $(this).attr("data-id");
  var parent_id = $("#categoryid").val();
  var del = $(this).closest("tr");

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
        $('#delete').attr('href', '/categories/removecategory?id=' + data + "&&categoryid=" + parent_id);
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
});

// media image insert

// $(document).on('click', ".upload-folders", function () {

//   var src = $(this).find('img').attr("src");

//   $("#categoryimage").val(src)
//   var data = $("#ctimagehide").attr("src", src);

//   if (data != "") {
//     $("#mediadesc").hide();
//     $("#browse").hide();
//     $("#ctimagehide").attr("src", src).show()
//     $("#catdel-img").show()

//   }
//   $("#addnewimageModal").hide()
//   $("#categoryModal").modal('show')

// })

$('#categoryModal').on('hide.bs.modal', function (event) {
  $('.modal-backdrop').remove()
})

$(document).on("click", "#mediamodalclose", function () {
  $("#categoryModal").modal('show')

})

$('#catdel-img').click(function () {
  $('categoryimage').val("")
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
    category_id = $("#category_id").val()

    $.ajax({
      url: "/categories/checkcategoryname",
      type: "POST",
      async: false,
      data: { "category_name": value, "category_id": category_id, csrf: $("input[name='csrf']").val() },
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
        space: true
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
        space: "* " + languagedata.spacergx

      },
      subcategoryid: {
        // required: "* Please select the category"
      }

    }
  });

  var formcheck = $("#form").valid();
  if (formcheck == true) {
    $('#form')[0].submit();

    var id = $(this).val()

    $("#form").attr("action", "")
    var cname = $("#cname").val();
    var cdesc = $("#cdesc").val();
    var categoryid = $("#cid").val();
    var parentid = $("#categoryid").val();
    var image = $("#categoryimage").val();
    var pageno = $("#catpageno").val();
    if ($("#form").valid()) {
      $.ajax({
        url: "/categories/updatesubcategory",
        type: "POST",
        datatype: "json",
        data: { "id": id, "cname": cname, "cdesc": cdesc, "categoryid": categoryid, "image": image, csrf: $("input[name='csrf']").val() },
        success: function (result) {
          if (result) {
            window.location.href = "/categories/addcategory/" + parentid;
  
          } else {
            // notify_content = '<div class="toast-msg sucs-green"> <a id="cancel-notify"> <img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a> <img src="/public/img/group-12.svg" alt="" class="left-img" /> <span>'+languagedata.internalserverr+'</span></div>';
            // $(notify_content).insertBefore(".breadcrumbs");
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

$(document).on('keyup', '.search', function () {

  var categoryid = $("#categoryid").val();
  if ($('.search').val() == "") {
    window.location.href = "/categories/addcategory/" + categoryid

  }
})

$(document).on('click', '#medcancel', function () {
  $("#addnewimageModal").hide()
})

// Get dropdown values and id

$(".drp").on("click", function () {

  var text = $(this).html()

  var id = $(this).attr("data-id")
  $(this).parent('.dropdown-menu').siblings('a').html($(this).html(text))
  $(this).parent().siblings("input").val(id)

  if (id == 0) {
    $("#availablecat").addClass('input-group-error')

  }else {
    $("#availablecat").removeClass('input-group-error')
    $("#cid-error").css('display', 'none')
  }
})

const imagee = document.querySelectorAll('.category-path-list ');
imagee.forEach(group => {
  const image = group.querySelectorAll('img');
  image[image.length - 1].style.display = 'none';
});

