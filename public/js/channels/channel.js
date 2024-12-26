/** Gobal Variables */
var languagedata
let sections = []
var fiedlvalue = []
let optionvalues = []
let deleteoption = []
let deletesecion = []
let deletefields = []
let fielid = 0             //db field id
let newFieldid = 1        //temporary field id
let orderindex = 1        //order index field
let optionid = 1
var sectionid = 1
let RelationalMember = 14
let RelationalMedia = 15
let RelationalVideo = 16

$(document).ready(async function () {
  var languagepath = $('.language-group>button').attr('data-path')

  await $.getJSON(languagepath, function (data) {

    languagedata = data

  })

  $('.drag-btn').attr({
    'data-bs-toggle': 'tooltip',
    'data-bs-placement': 'bottom',
    'data-bs-custom-class': 'custom-tooltip',
    'data-bs-title': 'Drag & Drop'
});

$('.edit-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Edit'
});

$('.duplicate-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Clone'
});

$('.remove-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Remove'
});

$('[data-bs-toggle="tooltip"]').tooltip();
})

// Create channel Next btn
$(document).on('click', '#nextstep', function () {
  var currentindex = $("#chn-step").val()
  if (currentindex == 1) {

    $("form[name='channelform']").validate({   // channel name and deiscription validation
      rules: {
        channelname: {
          required: true,
          space: true,

        },
        description: {
          required: true,
          maxlength: 250,
        }
      },
      messages: {
        channelname: {
          required: "*" + languagedata.Channell.channamevalid,
          space: "* " + languagedata.spacergx

        },
        description: {
          required: "*" + languagedata.Channell.chandescvalid,
          maxlength: "*" + languagedata.Permission.descriptionchat
        },
      }
    });

    $("form[name='channelform']").valid();

    if ($("form[name='channelform']").valid() == true) {
      $("#chn-step").val(2)
      $('.channelstep1').hide()
      $('#tab-s1').removeClass("text-[#10A37F]").addClass("text-bold-gray");
      $("#stp-lev").text(languagedata.Channell.steps + " 2 " + languagedata.of + " 3")
      $('.channelstep2').show()
      $('#tab-s2').removeClass("text-bold-gray").addClass("text-[#10A37F]");
    }
  }

  if (currentindex == 2) {
    FieldBasedPropertiesValidation()
    if (isvalied == false) {
      $("#chn-step").val(3)
      $('.channelstep2').hide()
      $('#tab-s2').removeClass("text-[#10A37F]").addClass("text-bold-gray");
      $('#tb1-tb2').removeClass("bg-[#10A37F]").addClass("bg-bold-gray");
      $("#stp-lev").text(languagedata.Channell.steps + " 3 " + languagedata.of + " 3")
      $('.channelstep3').show()
      if ($("#channelid").val() == "") {
        $("#nextstep").text(languagedata.save).addClass("channelsave")
      } else {
        $("#nextstep").text(languagedata.update).addClass("channelsave")
      }
      $('#tab-s3').removeClass("text-bold-gray").addClass("text-[#10A37F]");
      $('#tb2-tb3').removeClass("bg-bold-gray").addClass("bg-[#10A37F]");
    }
  }
})

// Previous btn
$(document).on('click', '#previous-btn', function () {
  var currentindex = $("#chn-step").val()
  if (currentindex == 3) {
    $("#chn-step").val(2)
    $('.channelstep3').hide()
    $('#tab-s3').removeClass("text-[#10A37F]").addClass("text-bold-gray");
    $('#tb2-tb3').removeClass("bg-[#10A37F]").addClass("bg-bold-gray");
    $("#stp-lev").text(languagedata.Channell.steps + " 2 " + languagedata.of + " 3")
    $("#nextstep").text(languagedata.next).removeClass("channelsave")
    $('.channelstep2').show()
    $('#tab-s2').removeClass("text-bold-gray").addClass("text-[#10A37F]");
    $('#tb1-tb2').removeClass("bg-bold-gray").addClass("bg-[#10A37F]");
    fiedlvalue = []
  }
  if (currentindex == 2) {
    $("#chn-step").val(1)
    $('.channelstep2').hide()
    $('#tab-s2').removeClass("text-[#10A37F]").addClass("text-bold-gray");
    $("#stp-lev").text(languagedata.Channell.steps + " 1 " + languagedata.of + " 3")
    $('.channelstep1').show()
    $('#tab-s1').removeClass("text-bold-gray").addClass("text-[#10A37F]");
  }
  if (currentindex == 1) {
    window.location.href = "/channels/"
  }
})

// add channel additional fields
$(document).on('click', '.field-types', function () {

  var msfieldid = $(this).attr("data-id")

  var fieldname = $(this).attr("type-name")

  var imgicon = $(this).find("img").attr("src")

  AddFieldString(msfieldid, "", fieldname, imgicon, "", "", "", "")

})

// additional fields sorting
document.addEventListener('DOMContentLoaded', function () {
  var el = document.getElementById('Sortsection');
  var sortable = Sortable.create(el, {
    handle: '.drag-btn',                          // drag handel
    animation: 150,
    fallbackOnBody: true,
    swapThreshold: 0.65,
    filter: ".constant-item",
    onStart: function (evt) {
      if (evt.item.classList.contains('constant-item')) {
        evt.preventDefault();
      }
    },
    onEnd: function (evt) {
      var items = Array.from(el.children);
      var constantItem = el.querySelector('.constant-item');
      el.appendChild(constantItem);
    }
  });

  sortable.option('filter', '.constant-item');   // remove the non sortable item
});

// additional fields remove
$(document).on('click', '.remove-field', function () {
  $('.deltitle').text(languagedata.Channell.delsectitle)
  $('.deldesc').text(languagedata.Channell.delseccontent)
  $('.delname').text($(this).parents(".sl-fields").find("p").text())
  $("#delid").attr("data-id", $(this).parents(".sl-fields").attr("data-orderindex"))
  $("#delid").attr("data-fieldid", $(this).parents(".sl-fields").find('.field-name').attr('field-id'))
  $(".modal-backdrop").css("display", "block")
})
$(document).on('click', '.deleteBtn', function () {
  delid = $(this).attr("data-id")
  delfieldid = $(this).attr("data-fieldid")
  var obj = {}
  obj.FieldId = parseInt(delfieldid)
  deletefields.push(obj)
  $(".new-field" + delid).remove()
  $("#deleteModal").modal("hide")
  $(".modal-backdrop").css("display", "none")
})

// create duplicate fields
$(document).on('click', '.duplicate-field', function () {

  var msfieldid = $(this).parents(".sl-fields").attr("data-id")
  var fieldname = $(this).parents(".sl-fields").find(".field-name").text()
  var imgicon = $(this).parents(".sl-fields").find(".field-icon").attr("src")

  AddFieldString(msfieldid, "", fieldname, imgicon, "", "", "", "")
})

// field string
function AddFieldString(masterfieldid, fieldid, fieldname, imgicon, dformat, timeformat, mandatoryfd, optval, id) {


  var div = `<div class="border border-[#ECECEC] group-hover bg-white rounded-[4px] p-[16px] relative flex items-center sl-fields h-[56px] new-field` + orderindex + `" data-id="` + masterfieldid + `" data-orderindex="` + orderindex + `">
                            <div class="flex items-center space-x-[8px]">
                                <img src="`+ imgicon + `" class="field-icon" alt="text">
                                <p class="text-bold-black text-sm font-normal mb-0 field-name" opt-val="`+ optval + `" dt-format="` + dformat + `" fl-mandatory="` + mandatoryfd + `" tm-format="` + timeformat + `" master-fieldid="` + masterfieldid + `" field-id="` + fieldid + `">` + fieldname + `</p>
                            </div>
                            <div class="absolute right-3 flex edit">
                                <button class="rounded-[4px] h-6 px-1.5 bg-white drag-btn">
                                    <img src="/public/img/fields-drag.svg">
                                </button>
                                <button class="rounded-[4px] h-6 px-1.5 bg-white edit-field">
                                    <img src="/public/img/fields-edit.svg">
                                </button>
                                <button class="rounded-[4px] h-6 px-1.5 bg-white duplicate-field">
                                    <img src="/public/img/fields-copy.svg">
                                </button>
                                <button class="rounded-[4px] h-6 px-1.5 bg-white remove-field">
                                    <img src="/public/img/fields-delete.svg">
                                </button>
                            </div>
                        </div>`

  $(div).insertBefore("#add-field")
  // if (id == 1) {
  //   $('.sl-fields').find('button').attr('disabled', 'true');

  // } else {
  //   $('.sl-fields').find('button').attr('disabled');
  // }

  orderindex = orderindex + 1;


  $('.drag-btn').attr({
    'data-bs-toggle': 'tooltip',
    'data-bs-placement': 'bottom',
    'data-bs-custom-class': 'custom-tooltip',
    'data-bs-title': 'Drag & Drop'
});

$('.edit-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Edit'
});

$('.duplicate-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Clone'
});

$('.remove-field').attr({
  'data-bs-toggle': 'tooltip',
  'data-bs-placement': 'bottom',
  'data-bs-custom-class': 'custom-tooltip',
  'data-bs-title': 'Remove'
});

$('[data-bs-toggle="tooltip"]').tooltip();
}


$(document).on('click','.remove-field', function() {
  $('#deleteModal').modal('show')
  });



// edit field options
$(document).on('click', '.edit-field', function () {
  $(".opt-delete").parents(".fields-opt").remove()
  dataid = $(this).parents(".sl-fields").attr("data-id")
  $("#fl-input").val($(this).parents(".sl-fields").find(".field-name").text())
  $("#fl-input").attr("msf-id", dataid)
  $("#fl-input").attr("data-id", $(this).parents(".sl-fields").attr("data-orderindex"))
  $("#dt-f").text($(this).parents(".sl-fields").find(".field-name").attr("dt-format"))
  $("#tm-f").text($(this).parents(".sl-fields").find(".field-name").attr("tm-format"))
  flmandatory = $(this).parents(".sl-fields").find(".field-name").attr("fl-mandatory")
  if (flmandatory == 1) {
    $("#Check").prop("checked", true)
  }
  $(this).parents(".sl-fields").find(".opt-input").each(function () {
    var div = ` <div class="flex items-center gap-[12px] fields-opt">
    <div class=" break-words border border-[#F7F7F5] pd-3 min-h-[36px] rounded-[4px] grow text-xs font-normal leading-4 text-[#262626] fl-opt flex items-center justify-between" data-id="${$(this).attr("data-id")}" field-id="${$(this).attr("field-id")}">${$(this).val()} <a href="javascript:void(0);" class="drag-btn min-w-[14px]"><img src="/public/img/drag.svg" alt="drag"></a></div>
    <a href="javascript:void(0);" class="border bg-[#F7F7F5] border-[#EDEDED] rounded-[4px] min-w-[36px] h-[36px] grid place-items-center hover:bg-[#e0e0e0] opt-delete"><img src="/public/img/delete-options.svg" alt="delete"></a></div>`
    $(".add-fp").append(div)
  })
  FieldBasedProperties(dataid)
  $("#Id2").modal("show")
})

// set property field based on field type
function FieldBasedProperties(id) {

  if (id == "2") {

    $(".fl-name").text("Properties - Text")

    $('#fl-input').attr('maxlength', '30');
    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "4") {
    $(".fl-name").text("Properties - Date & Time")

    $(".dt-field").show()
    $(".ti-field").show()
    $(".option-field").hide()

  } else if (id == "5") {

    $(".fl-name").text("Properties - Select")

    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").show()

  } else if (id == "6") {

    $(".fl-name").text("Properties - Date")

    $(".dt-field").show()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "7") {

    $(".fl-name").text("Properties - TextBox")
    $('#fl-input').attr('maxlength', '30');

    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "8") {

    $(".fl-name").text("Properties - TextArea")
    $('#fl-input').attr('maxlength', '30');

    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "9") {


    $(".fl-name").text("Properties - Radio Button")

    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").show()

  } else if (id == "10") {

    $(".fl-name").text("Properties - CheckBox")
    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").show()

  } else if (id == "14") {

    $(".fl-name").text("Properties - Members")
    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "15") {
    $(".fl-name").text("Properties - Media Gallery")
    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()

  } else if (id == "16") {

    $(".fl-name").text("Properties - Video URL")
    $(".dt-field").hide()
    $(".ti-field").hide()
    $(".option-field").hide()
  }
}

$(document).on("click", ".dt-format", function () {    // set dateformat in property

  $("#dt-f").text($(this).text())
  $("#dt-f-error").hide()

})
$(document).on("click", ".tm-format", function () {      // set timeformat in property

  $("#tm-f").text($(this).text())
  $("#tm-f-error").hide()

})

$(document).on("click", "#add-options", function () {       // add options in property
  optval = $("#opt-val").val()
  let exists = false;
  $('.fl-opt').each(function () {
    if ($(this).text().trim() == optval) {
      exists = true
    }
  })
  if ($("#opt-val").val() != "") {
    if (exists == true) {
      $("#opt-val-error").text("Option value already exists").show()
    } else {
      var div = `<div class="flex items-center gap-[12px] fields-opt">
  <div class="break-words border border-[#F7F7F5] pd-3 min-h-[36px] rounded-[4px] grow text-xs font-normal leading-4 text-[#262626] fl-opt flex items-center justify-between ">`+ optval + ` <a href="javascript:void(0);" class="drag-btn min-w-[14px]" ><img src="/public/img/drag.svg" alt="drag"></a></div>
  <a href="javascript:void(0);" class="border bg-[#F7F7F5] border-[#EDEDED] rounded-[4px] min-w-[36px] h-[36px] grid place-items-center hover:bg-[#e0e0e0] opt-delete ">
  <img src="/public/img/delete-options.svg" alt="delete"></a></div>`

      $(".add-fp").append(div)
      $("#opt-val").val("")
    }
  } else {
    $("#opt-val-error").text("Please enter the option").show()
  }
})

// delete property options
$(document).on("click", ".opt-delete", function () {
  id = $(this).parents(".fields-opt").find(".fl-opt").attr("data-id")
  field = $(this).parents(".fields-opt").find(".fl-opt").attr("field-id")
  if (id != "") {
    var delopt = {}
    delopt.Id = parseInt(id)
    delopt.FieldId = field
    deleteoption.push(delopt)

  }
  $(this).parents(".fields-opt").remove()

})

$(document).on('keyup', "#fl-input", function () {   //field name validation
  if ($("#fl-input") == "") {
    $("#fl-input-error").show()
  } else {
    $("#fl-input-error").hide()
  }
})
$(document).on('keyup', "#opt-val", function () {   //field option validation
  if ($("#opt-val") == "") {
    $("#opt-val-error").show()

  } else {
    $("#opt-val-error").hide()

  }
})
// set fields property value
var isvalied = false
$(document).on('click', "#flp-save", function () {
  $('.lengthErr').addClass('hidden')
  $("#opt-val").val("")
  if ($("#fl-input").val() == "") {
    $("#fl-input-error").show()
    isvalied = true
  } else {
    $("#fl-input-error").hide()
    isvalied = false
  }
  if ($("#fl-input").attr("msf-id") == 4 || $("#fl-input").attr("msf-id") == 6) {
    if ($("#dt-f").text() == "") {
      $("#dt-f-error").show()
      isvalied = true
    } else {
      $("#dt-f-error").hide()
      isvalied = false
    }
  }
  if ($("#fl-input").attr("msf-id") == 4) {
    if ($("#tm-f").text() == "") {
      $("#tm-f-error").show()
      isvalied = true
    } else {
      $("#tm-f-error").hide()
      isvalied = false
    }
  }
  if ($("#fl-input").attr("msf-id") == 5 || $("#fl-input").attr("msf-id") == 9 || $("#fl-input").attr("msf-id") == 10) {
    if ($(".fl-opt").text() == "") {
      $("#opt-val-error").text("Please enter the option").show()
      isvalied = true
    } else {
      $("#opt-val-error").hide()
      isvalied = false
    }
  }
  if (isvalied == false) {
    fieldindex = $("#fl-input").attr("data-id")
    $(".new-field" + fieldindex).find(".field-name").text($("#fl-input").val())
    $(".new-field" + fieldindex).find(".field-name").attr("dt-format", $("#dt-f").text())
    $(".new-field" + fieldindex).find(".field-name").attr("tm-format", $("#tm-f").text())
    $(".new-field" + fieldindex).find(".field-name").attr("fl-mandatory", $('#Check').is(":checked") == true ? 1 : 0)
    $('.fl-opt').each(function () {
      var intext = `<input type="hidden" class="opt-input" data-id="${$(this).attr("data-id")}" field-id="${$(this).attr("field-id")}" value="${$(this).text()}">`
      $(".new-field" + fieldindex).find(".field-name").attr("opt-val", 1)
      $(".new-field" + fieldindex).find(".field-name").append(intext)
    })
    $("#Check").prop("checked",false)
    $("#Id2").modal("hide")

  }
})
// modal cancel
$(document).on("click", "#cln-modal", function () {
  $("#Id2").modal("hide")
  $("#tm-f").text("")
  $("#dt-f").text("")
  $("#opt-val").val("")
  $(".opt-delete").parents(".fields-opt").remove()
  $("#Check").prop("checked", false)
  $("#tm-f-error").hide()
  $("#dt-f-error").hide()
  $("#fl-input-error").hide()
  $("#opt-val-error").hide()
  $(".lengthErr").addClass('hidden')

})

// chn save and update
$(document).on('click', '.channelsave', function () {
  fiedlvalue=[]
  GetfieldsValue()  //get field value
  console.log("myfields", fiedlvalue);

  var name = $('#channelname').val();
  var desc = $('#channeldesc').val();
  var url = window.location.search
  const urlpar = new URLSearchParams(url)
  pageno = urlpar.get('page')



  if ($("#channelid").val() == "") {

    if (SelectedCategoryValue.length == 0) {

      notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-error.svg" alt = "toast error"></div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Error</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Channell.selectcatvalid + `</p ></div ></div ></li></ul> `;
      $(notify_content).insertBefore(".header-rht");
      setTimeout(function () {
        $('.toast-msg').fadeOut('slow', function () {
          $(this).remove();
        });
      }, 5000); // 5000 milliseconds = 5 seconds

      return false
    }

    $.ajax({

      url: '/channels/newchannel',
      type: 'post',
      dataType: 'json',
      async: false,
      data: {
        "channelname": name,
        "channeldesc": desc,
        "sections": JSON.stringify({ sections }),
        "fiedlvalue": JSON.stringify({ fiedlvalue }),
        "categoryvalue": SelectedCategoryValue,
        csrf: $("input[name='csrf']").val()
      },
      success: function (data) {

        window.location.href = "/channels/"

      }
    })
  } else {
    var CurrentPage = $('#CurrentPage').val();

    if (SelectedCategoryValue.length == 0) {

      notify_content = `<ul class="fixed top-[56px] right-[16px] z-[1000] grid gap-[8px]"><li><div class="toast-msg flex max-sm:max-w-[300px]  flex relative max-sm:max-w-[300px] items-start gap-[8px] rounded-[2px] p-[12px_20px] border-l-[4px] border-[#FF8964] bg-[#FFF1ED]"> <a href="javascript:void(0)" class="absolute right-[8px] top-[8px]" id="cancel-notify"> <img src="/public/img/close-toast.svg" alt="close"> </a>` + `<div> <img src = "/public/img/toast-error.svg" alt = "toast error"></div> <div> <h3 class="text-[#FF8964] text-normal leading-[17px] font-normal mb-[5px] ">Error</h3> <p class="text-[#262626] text-[12px] font-normal leading-[15px] " >` + languagedata.Channell.selectcatvalid + `</p ></div ></div ></li></ul> `;
      $(notify_content).insertBefore(".header-rht");
      setTimeout(function () {
        $('.toast-msg').fadeOut('slow', function () {
          $(this).remove();
        });
      }, 5000); // 5000 milliseconds = 5 seconds

      return false
    }



    $.ajax({

      url: '/channels/updatechannel',
      type: 'post',
      dataType: 'json',
      async: false,
      data: {
        "id": $("#channelid").val(),
        "channelname": name,
        "channeldesc": desc,
        "sections": JSON.stringify({ sections }),
        "deletesections": JSON.stringify({ deletesecion }),
        "deletefields": JSON.stringify({ deletefields }),
        "deleteoptions": JSON.stringify({ deleteoption }),
        "fiedlvalue": JSON.stringify({ fiedlvalue }),
        "categoryvalue": SelectedCategoryValue,
        csrf: $("input[name='csrf']").val()
      },
      success: function (data) {

        if (CurrentPage == 1) {
          window.location.href = "/channels/"
        } else {
          window.location.href = "/channels/?page=" + CurrentPage
        }

      }
    })
  }

})

// Get fields value
// Get fields value
function GetfieldsValue() {

  var order = 1
  $('.sl-fields').each(function () {

    var obj = {}

    obj.MasterFieldId = parseInt($(this).find(".field-name").attr("master-fieldid"))

    obj.FieldId = parseInt($(this).find(".field-name").attr("field-id"))

    obj.NewFieldId = 0

    obj.SectionId = 0

    obj.SectionNewId = 0

    obj.FieldName = $(this).find(".field-name").text();

    obj.OrderIndex = order

    obj.Url = ""

    obj.Mandatory = parseInt($(this).find(".field-name").attr("fl-mandatory"))

    obj.IconPath = $(this).find(".field-icon").attr("src")

    obj.CharacterAllowed = ""

    obj.DateFormat = $(this).find(".field-name").attr("dt-format")

    obj.TimeFormat = $(this).find(".field-name").attr("tm-format")

    var opt = []

    var optorder = 1

    $(this).find(".opt-input").each(function () {

      var optionobj = {}

      optionobj.Id = parseInt($(this).attr("data-id"))

      optionobj.NewId = 0

      optionobj.FieldId = obj.FieldId

      optionobj.NewFieldId = 0

      optionobj.Value = $(this).val()

      optionobj.OrderIndex = optorder

      opt.push(optionobj)

      optorder++;

    })

    obj.OptionValue = opt

    fiedlvalue.push(obj)

    order++;
  })
}

//field based property validation
function FieldBasedPropertiesValidation() {

  $('.sl-fields').each(function () {
    if (($(this).find(".field-name").attr("master-fieldid") == "4") && ($(this).find(".field-name").attr("dt-format") == "" || $(this).find(".field-name").attr("tm-format") == "")) {

      $(this).find(".edit-field").click()
      if ($(this).find(".field-name").attr("dt-format") == "") {
        $("#dt-f-error").show()
        isvalied = true
      }
      if ($(this).find(".field-name").attr("tm-format") == "") {
        $("#tm-f-error").show()
        isvalied = true
      }
    } else if (($(this).find(".field-name").attr("master-fieldid") == "6") && ($(this).find(".field-name").attr("dt-format") == "")) {

      $(this).find(".edit-field").click()
      if ($(this).find(".field-name").attr("dt-format") == "") {
        $("#dt-f-error").show()
        isvalied = true
      }
    } else if (($(this).find(".field-name").attr("master-fieldid") == "5") && ($(this).find(".field-name").attr("opt-val") == "")) {

      $(this).find(".edit-field").click()
      $("#opt-val-error").show()
      isvalied = true
    } else if (($(this).find(".field-name").attr("master-fieldid") == "9") && ($(this).find(".field-name").attr("opt-val") == "")) {

      $(this).find(".edit-field").click()

      $("#opt-val-error").show()
      isvalied = true
    } else if (($(this).find(".field-name").attr("master-fieldid") == "10") && ($(this).find(".field-name").attr("opt-val") == "")) {

      $(this).find(".edit-field").click()
      $("#opt-val-error").show()
      isvalied = true
    }
  })
}

/**Edit URL */
$(document).ready(function () {

  if (window.location.href.indexOf("editchannel") > -1) {

    var id = $('#channelid').val();

    $.ajax({
      url: '/channels/getfields',
      type: 'post',
      dataType: 'json',
      async: false,
      data: {
        "id": id,
        csrf: $("input[name='csrf']").val()
      },
      success: function (data) {

        if (data.SelectedCategory != null) {

          SelectedCategoryValue = data.SelectedCategory
        }

        for (let [index, x] of data.FieldValue.entries()) {

          AddFieldString(x.MasterFieldId, x.FieldId, x.FieldName, x.IconPath, x.DateFormat, x.TimeFormat, x.Mandatory, 1, id)

          if (x.OptionValue != null) {
            for (let y of x.OptionValue) {

              var intext = `<input type="hidden" class="opt-input" value="${y.Value}" data-id="${y.Id}" field-id="${x.FieldId}">`

              $(".new-field" + (orderindex - 1)).find(".field-name").append(intext)
            }
          }
        }



      }
    })

  }
})

$(document).ready(function () {
  var $textarea = $('#channeldesc');
  var $errorMessage = $('#error-message');
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
  $('#channeldesc').on('input', function () {

    let lines = $(this).val().split('\n').length;

    if (lines > 5) {

      let value = $(this).val();

      let linesArray = value.split('\n').slice(0, 5);

      $(this).val(linesArray.join('\n'));
    }
  });
});

$(document).on("click", ".categorypdiv", function () {
  $(this).find('.category-select-btn').click();
})

$(document).on("click", ".selectedcategorydiv", function () {
  $(this).find('.category-unselect-btn').click();
})

document.addEventListener('DOMContentLoaded', function () {
  var el = document.getElementById('Sortsection2');

  var sortable = Sortable.create(el, {
    handle: '.drag-btn',             // drag handle
    animation: 150,
    fallbackOnBody: true,
    swapThreshold: 0.65,
    // Exclude the first element from sorting
    onEnd: function (evt) {
      var constantItem = el.querySelector('.constant-item');
      el.insertBefore(constantItem, el.firstChild); // Keep the first item in place after sorting
    }
  });
});



// channel fields text limit of 30 char

$(document).on('keyup', '.checklength', function () {

  var inputVal = $(this).val()

  var inputLength = inputVal.length

  if (inputLength == 30) {
       $(this).siblings('.lengthErr').removeClass('hidden')
  } else {
       $(this).siblings('.lengthErr').addClass('hidden')
  }

})