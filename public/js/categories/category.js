var languagedata

selectedcheckboxarr = []
/** */
$(document).ready(async function () {

  var languagepath = $('.language-group>button').attr('data-path')

  await $.getJSON(languagepath, function (data) {

    languagedata = data
  })

  $('.Content').addClass('checked');

  $('.category-path-list').each(function () {

    var length = $(this).children('.categorylistname').length - 1

    $(this).children('.categorylistname').each(function (index) {

      if (length == index) {

        $(this).next().remove();
      }

    })
  })

});

$(document).ready(function () {
  var $textarea = $('#cdesc');
  var $errorMessage = $('#error-messagedesc');
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
});

$(document).ready(function () {
  var $inputfield = $('#cname');
  var $errorMessagename = $('#error-messagename');
  var maxLength = 50;

  $inputfield.on('input', function () {
    if ($(this).val().length >= maxLength) {
      // Show error message
      $errorMessagename.text("Maximum 50 character allowed");
    } else {
      // Clear error message if under the limit
      $errorMessagename.text('');
    }
  });
});


$(document).ready(function (){
  $('#form').on('keydown', function(event) {
    if (event.key === 'Enter') {      
      event.preventDefault();  
    }
  });
})


$(document).ready(function () {
  $('#cdesc').on('input', function () {

    let lines = $(this).val().split('\n').length;

    if (lines > 5) {

      let value = $(this).val();

      let linesArray = value.split('\n').slice(0, 5);

      $(this).val(linesArray.join('\n'));
    }
  });
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

      },
      subcategoryid: {
        required: true

      }

    },
    messages: {
      cname: {
        required: "* " + languagedata.Categoryy.catnamevalid,
        space: "* " + languagedata.spacergx,
        duplicatename: "*" + languagedata.Categoryy.catnamevailderr

      },
      cdesc: {
        required: "* " + languagedata.Categoryy.catdescvalid,
        space: "* " + languagedata.spacergx,

      },
      subcategoryid: {
        required: "* " + languagedata.Categoryy.selectvaild
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
  var desc = edit.find("td:eq(2)").html();
  $("#triggerId").html(desc);
  $("#triggerId").css("gap", "5px")

  var url = window.location.search
  const urlpar = new URLSearchParams(url)
  pageno = urlpar.get('page')

  var data = $(this).attr("data-id");
  $("#exampleModalLabel").text(languagedata.Categoryy.updatecategory)
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
        var image = $("#ctimagehide").attr("src", "/image-resize?name=" + result.ImagePath)

        $("#cname").text(name);
        $("#cdesc").text(desc);
        $("#categoryimages").val(result.ImagePath)
        $("#cid").val(result.ParentId)
        $("#ctimagehide").attr("src", "/image-resize?name=" + result.ImagePath)
        $("#id").val(result.Id)
        $("#category_id").val(result.Id)


        var keyword = $("#searchcatlists").val()

        if (keyword != "") {
          $("#searchcatlists").val("")
          $(".noData-foundWrapper").hide()

        }

        if (result.ImagePath != "") {
          $("#browse").hide();
          $("#ctimagehide").show();
          $("#catdel-img").show()
          $('#uploadLine').hide()
          $('#uploadFormat').hide()

        } else {
          $("#browse").show();
          $("#ctimagehide").hide();
          $("#catdel-img").hide()
          $('#uploadLine').show()
          $('#uploadFormat').show()

        }


        if ("form[name='edit']") {
          $(".drp").each(function (index, element) {

            var id = $(element).attr("data-id")
            var categoryname = $(element).find("p").text().trim()
            var parentid = $(element).attr("data-parentid")

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
        notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"> <li><div class="flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"><a href="javascript:void(0)" class="absolute right-[8px] top-[8px] "><img src="/public/img/close-toast.svg" alt="close"></a><div><img src="/public/img/toast-error.svg" alt="toast error"></div><div><h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Error</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] "> ` + languagedata.Toast.internalserverr + ` </p></div></div></li><ul>`;
        $(notify_content).insertBefore(".header-rht");
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
  var url = window.location.search
  const urlpar = new URLSearchParams(url)
  pageno = urlpar.get('page')
  // $.ajax({
  //   url: "/categories/deletepopup",
  //   type: "GET",
  //   dataType: "json",
  //   data: { "id": data },
  //   success: function (results) {
  // if (results.Id == 0) {
  $('#content').text(languagedata.Categoryy.delcategory);
  $(".deltitle").text(languagedata.Categoryy.deltitle)
  $('.delname').text(del.find('td:eq(1)').text())
  $('#delid').show();
  if (pageno == null) {
    $('#delid').attr('href', '/categories/removecategory?id=' + data + "&&categoryid=" + parent_id);

  } else {
    $('#delid').attr('href', '/categories/removecategory?id=' + data + "&&categoryid=" + parent_id + "?page=" + pageno);

  }
  $('#btn3').text(languagedata.no);

  // } else {
  //   $('#content').text(languagedata.Categoryy.delcatinvalid);
  //   $('#delete').attr('href', '');
  //   $('#delid').hide();
  //   $('#btn3').text(languagedata.ok)

  // }
  // }
  // });
});


$("#caddbtn , #clickadd").on("click", function () {
  $(".input-group").removeClass("input-group-error")

  var CategoryId = $("#categoryid").val()
  $("#exampleModalLabel").text(languagedata.Categoryy.addcategry)

  $("#form").attr("action", "/categories/addsubcategory/" + CategoryId).attr("method", "POST")

  $("#cname").val("");
  $("#cdesc").val("");
  $("#cname-error").hide()
  $("#cdesc-error").hide()
  $("#cid-error").hide()
  $("#triggerId").html("")
  $("#triggerId").text(languagedata.Categoryy.selctcategry)


  $("#ctimagehide").attr("src", "");
  $("#categoryimages").val("")

  $("#save").show()
  $("#browse").show();
  $("#mediadesc").show();
  $("#update").hide()
  $("#ctimagehide").hide()
  $("#catdel-img").hide()

  $("#form").attr("name", "")


  var keyword = $("#searchcatlists").val()



  if ("form[name='']") {
    $(".drp").each(function (index, element) {
      if (keyword != "") {
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


$('#catdel-img').click(function () {
  $('#categoryimages').val("")
  $('#ctimagehide').attr('src', '')
  $('#ctimagehide').hide()
  $('#browse').show()
  $("#mediadesc").css("margin-top", "1%")

  $(this).siblings('p,button').show()
  $(this).hide()

})


$(document).on("click", "#cancel", function () {
  $("#exampleModalLabel").text("")
  $("#catdel-img").hide();
  $("#cname-error").hide()
  $("#cdesc-error").hide()
  $("#cid-error").hide()
  $('#catname').removeClass('input-group-error');
  $('#catdes').removeClass('input-group-error');
  $('#availablecat').removeClass('input-group-error');
  $('.modal-backdrop').remove()
  $('#ctimagehide').attr('src', '')
  $('#categoryimages').val("")
  $('#cid').val("")
  $('#uploadLine').show()
  $('#uploadFormat').show()
  $("#error-messagename").html("")
  $("#error-messagedesc").html("")

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
      data: { "cname": value, "parentid": parent_id, "categoryid": categoryid, csrf: $("input[name='csrf']").val() },
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
        required: true

      }

    },
    messages: {
      cname: {
        required: "* " + languagedata.Categoryy.catnamevalid,
        space: "* " + languagedata.spacergx,
        duplicatename: "*" + languagedata.Categoryy.catnamevailderr,

      },
      cdesc: {
        required: "* " + languagedata.Categoryy.catdescvalid,
        space: "* " + languagedata.spacergx,
        maxlength: "* " + languagedata.Permission.descriptionchat

      },
      subcategoryid: {
        required: "*" + languagedata.Categoryy.selectvaild
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
    var parentid = $("#categoryid").val();
    var image = $("#categoryimages").val();
    var pageno = $("#catpageno").val();


    if ($("#form").valid()) {
      $.ajax({
        url: "/categories/editsubcategory",
        type: "POST",
        datatype: "json",
        data: { "id": id, "cname": cname, "cdesc": cdesc, "pcategoryid": pcategoryid, "image": image, csrf: $("input[name='csrf']").val() },
        success: function (result) {


          if (result.value == true) {

            if (pageno == 0) {
              window.location.href = "/categories/addcategory/" + parentid;

            } else {
              window.location.href = "/categories/addcategory/" + parentid + "?page=" + pageno;

            }

          } else {
            // notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"> <li><div class="flex  max-sm:max-w-[300px] relative items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"><a href="javascript:void(0)" class="absolute right-[8px] top-[8px] "><img src="/public/img/close-toast.svg" alt="close"></a><div><img src="/public/img/toast-error.svg" alt="toast error"></div><div><h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Error</h3><p class="text-[#262626] text-[12px] font-normal leading-[15px] "> ` + languagedata.Toast.internalserverr + ` </p></div></div></li><ul>`;
            // $(notify_content).insertBefore(".header-rht");
            // setTimeout(function () {
            //   $('.toast-msg').fadeOut('slow', function () {
            //     $(this).remove();
            //   });
            // }, 5000); // 5000 milliseconds = 5 seconds
            window.location.href = result.url

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

  if (event.key === 'Backspace' && window.location.href.indexOf("keyword") > -1) {
    if ($('.search').val() == "") {
      window.location.href = "/categories/addcategory/" + categoryid

    }
  }
})

$(document).on('click', '#medcancel', function () {
  $("#addnewimageModal").hide()
})

$("#addnewimageModal").on('show.bs.modal', function (event) {
  $(this).css("background-color", "rgba(30,41,44,0.46)")
})

$("#categoryModal").on('show.bs.modal', function (event) {
  $(this).css("background-color", "rgba(30,41,44,0.46)")
})

// Get dropdown values and id

$(".drp").on("click", function () {

  var id = $(this).attr("data-id")
  $(this).parents('.dropdown-menu').siblings('a').html($(this).html())
  $("#cid").val(id)

  if (id == 0) {
    $("#availablecat").addClass('input-group-error')

  } else {
    $("#availablecat").removeClass('input-group-error')
    $("#cid-error").css('display', 'none')
  }
})


// form category search in keyup

$("#searchcatlists").keyup(function () {
  var keyword = $(this).val().trim().toLowerCase()
  $(".catrgory-dropdownlist a").each(function (index, element) {
    var title = $(element).text().toLowerCase()
    var categoryid = $("#category_id").val()
    var currentid = $(element).attr("data-id")
    if (title.includes(keyword) && categoryid != currentid) {
      $(element).show()
      $(".noData-foundWrapper").hide()

    } else {
      $(element).hide()
      if ($('.catrgory-dropdownlist button:visible').length == 0) {
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

$(document).on('click', '.selectcheckbox', function () {

  $('#unbulishslt').hide()
  $('#seleccheckboxdelete').removeClass('border-r');

  categorygrbid = $(this).attr('data-id')

  if ($(this).prop('checked')) {

    selectedcheckboxarr.push(categorygrbid)

  } else {

    const index = selectedcheckboxarr.indexOf(categorygrbid);

    if (index !== -1) {

      selectedcheckboxarr.splice(index, 1);
    }

    $('#Check').prop('checked', false)

  }


  if (selectedcheckboxarr.length != 0) {

    $('.selected-numbers').removeClass('hidden')

    var items

    if (selectedcheckboxarr.length == 1) {

      items = "Item Selected"
    } else {

      items = languagedata.itemselected
    }
    $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)


    $('#seleccheckboxdelete').removeClass('border-end')

    $('.unbulishslt').html("")

    if (selectedcheckboxarr.length > 1) {

      $('#deselectid').text("Deselect All")

    } else {
      $('#deselectid').text("Deselect")
    }


  } else {

    $('.selected-numbers').addClass('hidden')
  }

  var allChecked = true;

  $('.selectcheckbox').each(function () {

    if (!$(this).prop('checked')) {

      allChecked = false;

      return false;
    }
  });

  $('#Check').prop('checked', allChecked);

})

//ALL CHECKBOX CHECKED FUNCTION//

$(document).on('click', '#Check', function () {

  $('.selected-numbers').removeClass('hidden')
  $('#unbulishslt').hide()

  selectedcheckboxarr = []

  var isChecked = $(this).prop('checked');

  if (isChecked) {

    $('.selectcheckbox').prop('checked', isChecked);

    $('.selectcheckbox').each(function () {

      categorygrbid = $(this).attr('data-id')

      selectedcheckboxarr.push(categorygrbid)
    })

    $('#seleccheckboxdelete').removeClass('border-r');

    var items

    if (selectedcheckboxarr.length == 1) {

      items = "Item Selected"
    } else {

      items = languagedata.itemselected
    }
    $('.checkboxlength').text(selectedcheckboxarr.length + " " + items)


  } else {


    selectedcheckboxarr = []

    $('.selectcheckbox').prop('checked', isChecked);

    $('.selected-numbers').addClass('hidden')
  }
  if (selectedcheckboxarr.length == 0) {

    $('.selected-numbers').addClass('hidden')
  }

})
$(document).on('click', '#seleccheckboxdelete', function () {

  if (selectedcheckboxarr.length > 1) {

    $('.deltitle').text("Delete Categories?")

    $('#content').text('Are you sure want to delete selected Categories?')


  } else {

    $('.deltitle').text("Delete Category?")

    $('#content').text('Are you sure want to delete selected Category?')
  }


  $('#delid').addClass('checkboxdelete')
})
//MULTI SELECT DELETE FUNCTION//
$(document).on('click', '.checkboxdelete', function () {

  var url = window.location.href;

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
      "page": pageno,
      "parentid": parent_id


    },
    success: function (data) {

      if (data.value == true) {

        setCookie("get-toast", "Category Deleted Successfully")

        window.location.href = data.url
      } else {

        // setCookie("Alert-msg", "Internal Server Error")
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

  $('.selected-numbers').addClass('hidden')

})

$(document).on('click', '#categoryimage', function () {
  $("#prof-crop").val("9")
})

$(document).on("click", ".Closebtn", function () {
  $(".search").val('')
  $(".Closebtn").addClass("hidden")
  $(".srchBtn-togg").removeClass("pointer-events-none")
})

$(document).on("click", ".searchClosebtn", function () {
  $(".search").val('')
  var categoryid = $("#categoryid").val();
  window.location.href = "/categories/addcategory/" + categoryid
})

$(document).ready(function () {

  $('.search').on('input', function () {
      if ($(this).val().length >= 1) {
          $(".Closebtn").removeClass("hidden")
          $(".srchBtn-togg").addClass("pointer-events-none")
      } else {
          $(".Closebtn").addClass("hidden")
          $(".srchBtn-togg").removeClass("pointer-events-none")
      }
  });
})
