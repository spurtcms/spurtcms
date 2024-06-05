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

// $(document).ready(async function(){

var languagepath = $('.language-group>button').attr('data-path')

$.getJSON(languagepath, function (data) {

  languagedata = data
})

// })

/**STEPS changes body content */
$(document).ready(function () {

  var channelname = $("#channelname").val()

  var channeldesc = $('#channeldesc').val()

  if (channelname == "" && channeldesc == "") {

    $('.channelstep2').hide();

    $('.channelstep3').hide();

  }

  function CheckPrevioushideandshow() {

    console.log(fiedlvalue, "checkkkkkk")

    if (fiedlvalue == null) {

      console.log("condition1")

      fiedlvalue = []
    }

    if ($('.channelstep2').is(':visible')) {

      if (fiedlvalue == '') {

        console.log("conditon2")
        $('.fieldapp').each(function () {

          var obj = CreatePropertiesObjec()

          obj.MasterFieldId = parseInt($(this).attr('data-masterfieldid'))

          obj.NewFieldId = parseInt($(this).attr('data-newfieldid'))

          obj.FieldName = $(this).children('.heading-three').text();

          obj.IconPath = $(this).children('.img-div').children('img').attr('src')

          obj.SectionId = parseInt($(this).parent('.section-fields-content').attr('data-id'));

          obj.SectionNewId = parseInt($(this).parent('.section-fields-content').attr('data-newid'));

          obj.SectionName = $(this).parent('.section-fields-content').children('.section-class').children('input').val();

          fiedlvalue.push(obj)

          console.log(fiedlvalue, "arrayfield")
        })
      }

      $('.fieldapp').each(function () {

        const index = fiedlvalue.findIndex(obj => {

          return obj.FieldId == $(this).attr('data-fieldid') && obj.NewFieldId == $(this).attr('data-newfieldid');

        });

        console.log(index, index < 0);

        if (index == -1) {

          var obj = CreatePropertiesObjec()

          obj.MasterFieldId = parseInt($(this).attr('data-masterfieldid'))

          obj.NewFieldId = parseInt($(this).attr('data-newfieldid'))

          obj.FieldName = $(this).children('.heading-three').text();

          obj.IconPath = $(this).children('.img-div').children('img').attr('src')

          obj.SectionNewId = parseInt($(this).parent('.section-fields-content').attr('data-newid'));

          obj.SectionId = parseInt($(this).parent('.section-fields-content').attr('data-id'));

          obj.SectionName = $(this).parent('.section-fields-content').children('.section-class').children('input').val();

          fiedlvalue.push(obj)
        }

      })

      flg = FieldValidation()

      console.log("flg==", flg);

      if (!flg) {

        return false
      }

      $('#cr-ch-title').hide();

      $('#backto-prev').show();

      $('#cr-ch-title').hide();

      $('#backto-prev').show();

      $('.channelstep2').hide();

      $('.channelstep3').show();

      $('#step2').addClass('prev');

      $('#nextstep').hide();

      if (window.location.href.indexOf("editchannel") > -1) {

        $('.channelupt').show();

      } else {

        $('.channelsave').show();
      }


      $(".channelIdentifier").removeClass('active');

      $('#step3').addClass('active');


    } else {

      $('#cr-ch-title').hide();

      $('#backto-prev').show();

      $(".channelIdentifier").removeClass('active');

      $('#step2').addClass('active');

      $('.channelstep2').show();

      $('.channelstep3').hide();

      $('#nextstep').show();

      if (window.location.href.indexOf("editchannel") > -1) {

        $('.channelupt').hide();

      } else {

        $('.channelsave').hide();
      }



    }
  }

  function Backtoprevious() {

    if ($('.channelstep2').is(':visible')) {

      $('#cr-ch-title').show();

      $('#backto-prev').hide();

      $('.channelstep2').hide();

      $('.channelstep1').show();

      $('#step1').removeClass('prev');

      $('#nextstep').show();

      if (window.location.href.indexOf("editchannel") > -1) {

        $('.channelupt').hide();

      } else {

        $('.channelsave').hide();
      }

      $(".channelIdentifier").removeClass('active');

      $('#step1').addClass('active');

    } else {

      $('#cr-ch-title').hide();

      $('#backto-prev').show();

      $('.channelstep3').hide();

      $('.channelstep2').show();

      $('#step2').removeClass('prev');

      $('#nextstep').show();

      if (window.location.href.indexOf("editchannel") > -1) {

        $('.channelupt').hide();

      } else {

        $('.channelsave').hide();
      }

      $(".channelIdentifier").removeClass('active');

      $('#step2').addClass('active');

    }

  }

  $(document).on('click', '#nextstep', function () {


    $("form[name='channelform']").validate({
      rules: {
        channelname: {
          required: true,
          space: true,

        },
        description: {
          required: true,

        }
      },
      messages: {
        channelname: {
          required: "*" + languagedata.Channell.channamevalid,
          space: "* " + languagedata.spacergx

        },
        description: {
          required: "*" + languagedata.Channell.chandescvalid
        },
      }
    });

    $("form[name='channelform']").valid();

    if ($('#channelname').val() == "" || $('#channeldesc').val() == "") {

      Validationcheck();

      return false

    }

    $('#step1').addClass('prev');

    $('.channelstep1').hide();

    $('.channelstep3').hide();

    CheckPrevioushideandshow()

  })

  $(document).on('click', '#back', function () {

    if ($('#cr-ch-title').is(":visible")) {

      window.location.href = "/channels/"

    } else {

      Backtoprevious()
    }

  })

})

/** */
$(document).ready(function () {
  //main menu
  Sort()

})



/** Category select */
$(document).ready(function () {

  var count = 0

  $('.categorypdiv').each(function () {

    if ($(this).css('display') != "none") {

      count = count + 1

    }

    var id

    $(this).children('.choose-cat-list-col').children('.categoryname').each(function () {

      id = $(this).attr('data-id');

    })

    $(this).children('.category-select-btn').attr('data-id', id).attr('id', 'lastcatid-' + id)
  })

  $('#avacatcount').text(count);

})


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

        if (data.Section != null) {

          sections = data.Section
        }

        if (data.FieldValue != null) {

          fiedlvalue = data.FieldValue
        }

        if (data.SelectedCategory != null) {

          SelectedCategoryValue = data.SelectedCategory
        }

        var sectionstr = ``

        for (let x of data.Section) {

          sectionstr += `<div class="section-fields-content drop-able ui-droppable" style="margin-bottom: 10px;" id="section-` + x['SectionId'] + `` + x['SectionNewId'] + `" data-id ="` + x['SectionId'] + `"  data-newid="` + x['SectionNewId'] + `" data-orderindex="` + x['OrderIndex'] + `">
          <div class="section-class">
            <input type="text" class="sectionname" placeholder="" style="font-size:0.825rem" value="`+ x['SectionName'] + `">
            <button class="sectiondel" style="margin-right: 11px;"><img src="/public/img/bin.svg"></button>   
          </div>
        </div>`

        }

        $('#Sortsection').append(sectionstr);

        for (let x of data.Section) {

          SortableSection(x['SectionId'], x['SectionNewId'])

        }

        for (let [index, x] of data.FieldValue.entries()) {

          $('.section-fields-content').each(function () {

            if ($(this).attr('data-id') == x['SectionId'] && $(this).attr('data-newid') == x['SectionNewId']) {

              var fieldstr = `<div class="section-fields-content-child fieldapp" data-masterfieldid=` + x.MasterFieldId + ` data-newfieldid=` + x.NewFieldId + ` data-fieldid=` + x.FieldId + ` data-orderindex=` + x.OrderIndex + ` id=fieldapp` + x.FieldId + `` + x.NewFieldId + `>
              <input type="hidden" class="orderval" value="">
                  <div class="img-div">
                      <img src="`+ x.IconPath + `" alt="">
                  </div>
                  <h3 class="heading-three" id=fieldtitle`+ x.FieldId + `` + x.NewFieldId + `>` + x.FieldName + `</h3>
                  <a href="javascript:void(0)" class="field-delete delete" data-id="`+ id + `">
                      <img src="/public/img/bin.svg" alt="">
                  </a>
              </div> `

              $(this).append(fieldstr);

            }

          })

          if (fiedlvalue[index].OptionValue == null) {

            fiedlvalue[index].OptionValue = []

          }

        }

        const index1 = fiedlvalue.findIndex(obj => {

          return obj.MasterFieldId == RelationalMember

        });

        if (index1 >= 0) {

          $("#relational-member").draggable("disable");

          $("#relational-member").css('opacity', '0.8');
        }

      }
    })

  }
})

/** Update Channel */
$(document).on('click', '.channelupt', function () {


  $("form[name='channelform']").validate({
    rules: {
      channelname: {
        required: true,
        space: true,

      },
      description: {
        required: true,

      }
    },
    messages: {
      channelname: {
        required: "*" + languagedata.Channell.channamevalid,
        space: "*" + languagedata.spacergx,

      },
      description: {
        required: "*" + languagedata.Channell.chandescvalid
      },
    }
  });

  var formcheck = $("form[name='channelform']").valid();

  if (formcheck == true) {

    // $("form[name='channelform']")[0].submit();
  }

  else {

    Validationcheck()

    $(document).on('keyup', ".field", function () {

      Validationcheck()

    })

    return false

  }


  $('.fieldapp').each(function () {

    const index = fiedlvalue.findIndex(obj => {

      return obj.FieldId == $(this).attr('data-fieldid') && obj.NewFieldId == $(this).attr('data-newfieldid');

    });

    if (index < 0) {

      var obj = CreatePropertiesObjec()

      obj.MasterFieldId = parseInt($(this).attr('data-masterfieldid'))

      obj.NewFieldId = parseInt($(this).attr('data-newfieldid'))

      obj.FieldName = $(this).children('.heading-three').text();

      obj.IconPath = $(this).children('.img-div').children('img').attr('src')

      obj.SectionId = parseInt($(this).parent('.section-fields-content').attr('data-id'));

      obj.SectionNewId = parseInt($(this).parent('.section-fields-content').attr('data-newid'));

      obj.SectionName = $(this).parent('.section-fields-content').children('.section-class').children('input').val();

      fiedlvalue.push(obj)
    }

  })

  SortFieldsandSection()

  flg = FieldValidation()
  flg1 = CateogryValidation()

  if (!flg || !flg1) {

    console.log(flg, flg1);

    return false

  }

  var name = $('#channelname').val();

  var desc = $('#channeldesc').val();

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

      if (data == true) {

        setCookie("get-toast", "Channel Updated Successfully")

        setCookie("Alert-msg", "success", 1)
      }
      else {

        setCookie("Alert-msg", "Internal Server Error")

      }
      window.location.href = "/channels/"

    }
  })

})

/** if section arry empty create default section */
$(document).ready(function () {


  if (sections.length == 0) {

    var sectionobj = SectionObjectCreate()

    sectionobj.SectionName = "New Section"
    sections.push(sectionobj)

    //   $(`<div class="section-fields-content drop-able active ui-droppable" style="margin-bottom: 10px;" id="section-0` + sectionid + `" data-newid="` + sectionid + `" data-orderindex="` + orderindex + `">
    //       <div class="section-class" >
    //         <input type="text" class="sectionname" placeholder="Enter the Name" style="font-size:0.825rem" value="Section 1"/>
    //         <button class="sectiondel" data-bs-toggle="modal" data-bs-target="#centerModal"><img src="/public/img/bin.svg"></button>   
    //       </div>
    //     </div>
    // `).insertAfter('#Sortsection')

    $("#Sortsection").append(`<div class="section-fields-content drop-able active ui-droppable" style="margin-bottom: 10px;" id="section-0` + sectionid + `" data-newid=${sectionid} data-id=0 data-orderindex=${orderindex}>
  <div class="section-class" >
    <input type="text" class="sectionname" placeholder="" style="font-size:0.825rem" value="New Section"/>
    <button class="sectiondel" style="margin-right: 11px;" ><img src="/public/img/bin.svg"></button>   
  </div>
</div>`)

    sectionid++;

    orderindex++;

  }


  // drop
  $(".drop-able").droppable({
    tolerance: 'pointer',

    drop: function (ev, ui) {

      console.log("draggggggg")
      var id = ui.draggable[0]['dataset'].id;
      var name = ui.draggable.text().trim().split("\n");
      var img = ui.draggable.find("img").attr("src");
      var newsection = ui.draggable[0]['id']
      console.log("newsectiondrop", id, name, img, newsection)
      if (newsection != "newsection") {
        $(this).append(`
            <div class="section-fields-content-child fieldapp" data-masterfieldid=`+ id + ` data-newfieldid=` + newFieldid + ` data-fieldid=` + fielid + ` data-orderindex=` + orderindex + ` id=fieldapp` + fielid + `` + newFieldid + `>
            <input type="hidden" class="orderval" value="">
                <div class="img-div">
                    <img src="`+ img + `" alt="">
                </div>
                <h3 class="heading-three" id=fieldtitle`+ fielid + `` + newFieldid + `>` + name[0] + `</h3>
                <a href="javascript:void(0)" class="field-delete delete" data-id="`+ id + `" >
                    <img src="/public/img/bin.svg" alt="">
                </a>
            </div>    
        `)


        newFieldid = newFieldid + 1;
        orderindex = orderindex + 1;
      }

      if (newsection == 'relational-member') {

        $(".fieldapp[data-masterfieldid|='14']").trigger('click');

        var obj = CreatePropertiesObjec()

        fiedlvalue.push(obj)

        console.log(fiedlvalue);

        $("#relational-member").draggable("disable")

        $("#relational-member").css('opacity', '0.6');

      }

    },
  });

  var sortableElement = document.querySelector('.drop-able')
  // Check if the element exists before creating the Sortable instance
  if (sortableElement) {
    var sortable = new Sortable(sortableElement, {
      animation: 200,
      filter: '.non-sortable-item',
      onStart: function (evt) {
        if (evt.from.classList.contains('non-sortable-item')) {
          sortable.option("disabled", true);
        }
      },
      onEnd: function (evt) {
        sortable.option("disabled", false);
      }
    });
  } else {
    console.error("Sortable container element not found.");
  }

})

// drag
$(".drag-able").draggable({
  appendTo: "body",
  cursor: "move",
  helper: "clone",
  revert: true,
  revertDuration: 0
}).disableSelection();

// drag
$(".section-drag-able").draggable({
  appendTo: "body",
  cursor: "move",
  helper: "clone",
  revert: true,
  revertDuration: 0
}).disableSelection();


var sortableElement = document.querySelector('.drop-able')
// Check if the element exists before creating the Sortable instance
if (sortableElement) {
  var sortable = new Sortable(sortableElement, {
    animation: 200,
    filter: '.non-sortable-item',
    onStart: function (evt) {
      if (evt.from.classList.contains('non-sortable-item')) {
        sortable.option("disabled", true);
      }
    },
    onEnd: function (evt) {
      sortable.option("disabled", false);
    }
  });
} else {
  console.error("Sortable container element not found.");
}

// drop
$(".drop-able").droppable({
  tolerance: 'pointer',
  drop: function (ev, ui) {

    var id = ui.draggable[0]['dataset'].id;
    var name = ui.draggable.text().trim().split("\n");
    var img = ui.draggable.find("img").attr("src");
    var newsection = ui.draggable[0]['id']
    console.log("check")

    if (newsection != "newsection") {
      $(this).append(`
            <div class="section-fields-content-child fieldapp" data-masterfieldid=`+ id + ` data-newfieldid=` + newFieldid + ` data-fieldid=` + fielid + ` data-orderindex=` + orderindex + ` id=fieldapp` + fielid + `` + newFieldid + `>
            <input type="hidden" class="orderval" value="">
                <div class="img-div">
                    <img src="`+ img + `" alt="">
                </div>
                <h3 class="heading-three" id=fieldtitle`+ fielid + `` + newFieldid + `>` + name[0] + `</h3>
                <a href="javascript:void(0)" class="field-delete delete" data-id="`+ id + `" >
                    <img src="/public/img/bin.svg" alt="">
                </a>
            </div>    
        `)


      newFieldid = newFieldid + 1;
      orderindex = orderindex + 1;
    }

  },
});

$(".section-drop-able").droppable({
  tolerance: 'pointer',
  drop: function (ev, ui) {
    var draggable = ui.draggable;
    var id = draggable.data('id');
    var name = draggable.text().trim().split("\n");
    var img = draggable.find("img").attr("src");
    var this_id = draggable.attr("id");
    var newsection = this_id;
    if (newsection === "newsection") {
      createDroppable();
    }
  },
});

// Function to create a new droppable element
function createDroppable() {

  var droppable = $("<div>", {
    class: "section-fields-content drop-able newdrop-able",
    style: "margin-bottom: 10px;",
    id: "section-0" + sectionid,
    "data-id": 0,
    "data-newid": sectionid,
    "data-orderindex": orderindex,
    html: `
      <div class="section-class">
      <input type="text" class="sectionname" placeholder="" style="font-size:0.825rem" value="New Section"/>
      <button class="sectiondel" style="margin-right: 11px;" ><img src="/public/img/bin.svg"></button>                              
      </div>
    `
  }).appendTo("#Sortsection");

  orderindex++;


  droppable.droppable({
    drop: function (event, ui) {
      console.log("sectioncreate")
      var id = ui.draggable.data('id');
      var name = ui.draggable.text().trim().split("\n");
      var img = ui.draggable.find("img").attr("src");
      if (ui.draggable[0]['id'] !== "newsection") {
        $(this).append(`
          <div class="section-fields-content-child fieldapp" data-masterfieldid="${id}" data-newfieldid="${newFieldid}" data-fieldid="${fielid}" data-orderindex="${orderindex}" id="fieldapp${fielid}${newFieldid}">
            <input type="hidden" class="orderval" value="">
            <div class="img-div">
              <img src="${img}" alt="">
            </div>
            <h3 class="heading-three" id="fieldtitle${fielid}${newFieldid}">${name[0]}</h3>
            <a href="javascript:void(0)" class="field-delete delete" data-id="${id}" >
              <img src="/public/img/bin.svg" alt="">
            </a>
          </div>
        `);

        newFieldid++;
        orderindex++;

      }
    }
  });


  var sortableElement = document.getElementById('section-0' + sectionid)
  console.log(sortableElement, "elemtn")
  // Check if the element exists before creating the Sortable instance
  if (sortableElement) {
    var sortable = new Sortable(sortableElement, {
      animation: 200,
      filter: '.non-sortable-item',
      onStart: function (evt) {
        if (evt.from.classList.contains('non-sortable-item')) {
          sortable.option("disabled", true);
        }
      },
      onEnd: function (evt) {
        sortable.option("disabled", false);
      }
    });
  } else {
    console.error("Sortable container element not found.");
  }

  var sectionobj = SectionObjectCreate()

  sectionobj.SectionName = "New Section "

  sections.push(sectionobj)

  sectionid++;

}

function Sort() {

  var sortableElement = document.getElementById('Sortsection')
  // Check if the element exists before creating the Sortable instance
  if (sortableElement) {
    var sortable = new Sortable(sortableElement, {
      animation: 200,
      filter: '.non-sortable-item',
      onStart: function (evt) {
        if (evt.from.classList.contains('non-sortable-item')) {
          sortable.option("disabled", true);
        }
      },
      onEnd: function (evt) {
        sortable.option("disabled", false);
      }
    });
  } else {
    console.error("Sortable container element not found.");
  }
}


function SortableSection(id, newid) {

  var sortableElement = document.getElementById('section-' + id.toString() + newid.toString())
  // Check if the element exists before creating the Sortable instance
  if (sortableElement) {
    var sortable = new Sortable(sortableElement, {
      animation: 200,
      filter: '.non-sortable-item',
      onStart: function (evt) {
        if (evt.from.classList.contains('non-sortable-item')) {
          sortable.option("disabled", true);
        }
      },
      onEnd: function (evt) {
        sortable.option("disabled", false);
      }
    });
  } else {
    console.error("Sortable container element not found.");
  }
}


function SortFieldsandSection() {

  var count = 1

  $('.section-fields-content').each(function () {

    $(this).attr("data-orderindex", count);


    if (sections != null){

      for (let [index, x] of sections.entries()) {
  
        if (x.SectionId == $(this).attr('data-id') && x.SectionNewId == $(this).attr('data-newid')) {
  
          sections[index].OrderIndex = count
  
        }
  
      }
    }


    count++

    $(this).children('.fieldapp').each(function () {

      $(this).attr("data-orderindex", count);


      if (fiedlvalue != null) {

        for (let x of fiedlvalue) {

          if (x.FieldId == $(this).attr('data-fieldid') && x.NewFieldId == $(this).attr('data-newfieldid')) {

            x.OrderIndex = count

          }

        }
      }
      count++
    });

  })


}


/**properties empty function */
function PropertiesEmpty() {

  $('#finitialv').val("");

  $('#fplacehold').val("");

  $('#fdatefor').val("");

  $('#ftimefor').val("");

  $('#date-form').html(languagedata?.Channell?.chosedate);

  $('#time-for').html(languagedata?.Channell?.chosetime);

  $('#furl').val("");

  $('#foptname').val("");

  $("#ftext").val("");

  $("#fdesc").val("");

  $('.options-added').html("");

  $('#triggerId').text(languagedata?.Channell?.dateformat);

  $('#time-format').text(languagedata?.Channell?.timeformat);

  $('#fcharvalue').val(0)

  $('#Check2').prop('checked', false)
}

/** */
function FieldBasedProperties(id) {

  $('.prop').hide();

  $('#multi').parent().hide();

  $("#title").css('visibility', 'visible');

  $('#desc').css('display', 'flex');

  $("#distext").css('display', 'flex');

  $('#Check2').parent().show();

  if (id == "2") {

    $("#checkbox").css('display', 'flex');

    $('#fcharallow').css('display', 'flex')

  } else if (id == "3") {

    $("#url").css('display', 'flex');

  } else if (id == "4") {

    $(this).parents('.dd-c').siblings('.newdd-input').prop('checked', false);

    $("#dateform").css('display', 'flex');

    $('#timefor').css('display', 'flex');

  } else if (id == "5") {

    $('#option').css('display', 'grid');

  } else if (id == "6") {

    $("#dateform").css('display', 'flex');

  } else if (id == "7") {

    $('#fcharallow').css('display', 'flex')

  } else if (id == "8") {

    $('#fcharallow').css('display', 'flex')

  } else if (id == "9") {

    $('#option').css('display', 'grid');

  } else if (id == "10") {

    $('#option').css('display', 'grid');

  } else if (id == "11") {

    $('#fcharallow').css('display', 'flex')

  } else if (id == "12") {

    $('#desc').hide();

    $("#distext").css('display', 'flex');

    $("#checkbox").hide();

  } else if (id == "13") {

    $('#desc').hide();

    $("#distext").css('display', 'flex');

    $("#checkbox").hide();

  }else if (id =="14"){
    
    $('#desc').hide();

    $("#distext").css('display', 'none');

    $("#checkbox").hide();

    $('#Check2').parent().hide();

    // $('#multi').parent().show();

  }
}

/*field select on*/
$(document).on('click', '.fieldapp', function () {

  PropertiesEmpty()

  console.log(fiedlvalue);

  $('#field-properties').show();

  $('.fieldapp').removeClass('active');

  $(this).addClass('active');

  var index = $(this).attr('data-newid');

  var img = $(this).children('.img-div').children('img').attr('src')

  var name = $(this).children('.heading-three').text()

  $("#ftext").val(name);

  $('#title').html(`<div class="img-div"><img src="` + img + `" alt=""></div><h3 class="heading-three">` + name + `</h3>`);

  $('.fcheck').prop("checked", false);

  var id = $(this).attr('data-masterfieldid');

  $('#fieldid').attr('data-fieldid', $(this).attr('data-fieldid')).attr('data-newfieldid', $(this).attr('data-newfieldid'))

  $('#fieldid').attr('data-masterfieldid', $(this).attr('data-masterfieldid'))

  $('#fieldid').attr('data-sectionid', $(this).parents('.section-fields-content').attr('data-id')).attr('data-newsectionid', $(this).parents('.section-fields-content').attr('data-newid')).attr('data-sectionname', $(this).parents('.section-fields-content').children('.section-class').children('input').val())

  FieldBasedProperties(id)

  var fieldid = $(this).attr('data-fieldid');

  var newfldid = $(this).attr('data-newfieldid');

  /**This is for select option empty */
  if ($('#fieldid').attr('data-fieldid') == fieldid && $('#fieldid').attr('data-newfieldid') == newfldid) {

    optionvalues = [];

  }

  if (fiedlvalue.length > 0) {

    for (let i = 0; i < fiedlvalue.length; i++) {

      if (fiedlvalue[i]['FieldId'] == fieldid && fiedlvalue[i]['NewFieldId'] == newfldid) {

        $("#ftext").val(fiedlvalue[i]['FieldName'].trim());

        $("#fdesc").val(fiedlvalue[i]['Description']);

        fiedlvalue[i]['Mandatory'] == 1 ? $('#Check2').prop("checked", true) : $('#Check2').prop("checked", false);

        $('#finitialv').val(fiedlvalue[i]['InitialValue']);

        $('#fplacehold').val(fiedlvalue[i]['Placeholder']);

        $('#fdatefor').val(fiedlvalue[i]['DateFormat']);

        $('#ftimefor').val(fiedlvalue[i]['TimeFormat']);

        $('#date-form').html(fiedlvalue[i]['DateFormat']);

        $('#time-for').html(fiedlvalue[i]['TimeFormat']);

        $('#furl').val(fiedlvalue[i]['Url']);

        $('#triggerId').text(fiedlvalue[i]['DateFormat'])

        $('#time-format').text(fiedlvalue[i]['TimeFormat'])

        $('#fcharvalue').val(fiedlvalue[i]['CharacterAllowed'])

        optionvalues = fiedlvalue[i]['OptionValue'];

        if (optionvalues !== null && optionvalues.length > 0) {

          for (let j = 0; j < optionvalues.length; j++) {


            $('.options-added').append(`<div class="addedOption"><img src="/public/img/addOption-drag.svg" alt="" /> <p class="para">${optionvalues[j].Value}</p><button href="javascript:void(0)" class="option-delete delval" data-id="${optionvalues[j].Id}" data-newid=${optionvalues[j].NewId}> <img src="/public/img/addOption-close.svg" alt="" /></button></div>`)
          }

        }

        if (fiedlvalue[i]['FieldName'] != "") {
          $("#fieldname-error").hide();
          $('#fieldname-error').parents('.input-group').removeClass('input-group-error')
        }

        if (fiedlvalue[i]['OptionValue'].length != 0) {
          $('#select-error').hide();
          $('#select-error').parents('.input-group').removeClass('input-group-error');
        }

        if ((fiedlvalue[i]['DateFormat']) != "") {
          $('#date-error').hide();
          $('#date-error').parents('.input-group').removeClass('input-group-error');
        }

        if (fiedlvalue[i]['TimeFormat'] != "") {
          $('#time-error').hide();
          $('#time-error').parents('.input-group').removeClass('input-group-error');
        }

      }
    }
  }

})

/** Section Object create */
function SectionObjectCreate() {

  var obj = {}

  obj.SectionId = 0

  obj.SectionNewId = sectionid

  obj.SectionName = ""

  obj.MasterFieldId = parseInt($('#newsection').attr('data-id'))

  return obj

}

/** Properties Object create */
function CreatePropertiesObjec() {

  var obj = {}

  obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

  obj.FieldId = 0

  obj.NewFieldId = parseInt($('#fieldid').attr('data-newfieldid'));

  obj.SectionId = parseInt($('#fieldid').attr('data-sectionid'))

  obj.SectionNewId = parseInt($('#fieldid').attr('data-newsectionid'));

  obj.FieldName = $('#ftext').val();

  obj.Url = $('#furl').val();

  obj.Mandatory = $('#Check2').is(":checked") == true ? 1 : 0

  obj.IconPath = $('#title').children('.img-div').children('img').attr('src')

  obj.CharacterAllowed = parseInt($('#fcharvalue').val()) == "" ? 0 : parseInt($('#fcharvalue').val());

  obj.DateFormat = ""

  obj.TimeFormat = ""

  obj.OptionValue = []

  return obj
}

/** properties keyup obj creation */
$(document).on('keyup', '.propinput', function () {

  var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

  var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

  const index = fiedlvalue.findIndex(obj => {

    return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

  });

  if (fiedlvalue.length == 0 || index < 0) {

    const obj = CreatePropertiesObjec()

    obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

    fiedlvalue.push(obj)

  } else {

    fiedlvalue[index].FieldName = $('#ftext').val();

    fiedlvalue[index].Url = $('#furl').val();

    fiedlvalue[index].CharacterAllowed = parseInt($('#fcharvalue').val()) == NaN ? 0 : parseInt($('#fcharvalue').val());

  }

  /**center field name change */
  $('#fieldtitle' + fieldid.toString() + newfieldid.toString()).text($('#ftext').val())

})

/** Properties Dropdown TimeFormat*/
$(document).on('click', '.dropdown-item', function () {

  var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

  var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

  $(this).parents('a').siblings('.fidvalinput').val($(this).attr('data-id'))

  $(this).parents('a').siblings('.error').hide();

  const index = fiedlvalue.findIndex(obj => {

    return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

  });

  if (fiedlvalue.length == 0 || index < 0) {

    const obj = CreatePropertiesObjec()

    obj.MasterFieldId = fieldid;

    obj.NewFieldId = newfieldid

    obj.SectionNewId = parseInt($('#fieldid').attr('data-newsectionid'));

    obj.DateFormat = $('#fdatefor').val()

    obj.TimeFormat = $('#ftimefor').val()

    fiedlvalue.push(obj)

  } else {

    fiedlvalue[index].DateFormat = $('#fdatefor').val()

    fiedlvalue[index].TimeFormat = $('#ftimefor').val()

  }

})

/** Properties Option value */
$(document).on('click', '#addvalue', function () {

  var optionvalue = $('#foptname').val();

  if (optionvalue == "") {

    $('#select-error1').show();

    $('#select-error1').parents('.input-group').addClass('input-group-error');

    return

  }

  $('#foptname').val("")

  // $('.options-added').append(`<div class="options-added-row flexx">` + optionvalue + `<button class="delval" data-id="0" data-newid=` + optionid + `> <img src="/public/img/trash-style2.svg" alt=""> </button> </div>`)

  $('.options-added').append(`<div class="addedOption"><img src="/public/img/addOption-drag.svg" alt="" /> <p class="para">${optionvalue}</p><button href="javascript:void(0)" class="option-delete delval" data-id="0" data-newid=${optionid}> <img src="/public/img/addOption-close.svg" alt="" /></button></div>`)

  var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

  var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

  const index = fiedlvalue.findIndex(obj => {

    return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

  });

  var optionobj = {}

  optionobj.Id = 0

  optionobj.NewId = optionid

  optionobj.FieldId = fieldid

  optionobj.NewFieldId = newfieldid

  optionobj.Value = optionvalue

  optionvalues.push(optionobj)

  optionid = optionid + 1

  if (fiedlvalue.length == 0 || index < 0) {

    const obj = CreatePropertiesObjec()

    obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

    obj.NewFieldId = newfieldid

    obj.OptionValue = optionvalues

    obj.SectionNewId = parseInt($('#fieldid').attr('data-newsectionid'));

    fiedlvalue.push(obj)

  } else {

    fiedlvalue[index].OptionValue = optionvalues
  }

  $('#select-error').hide();

  $('#select-error').parents('.input-group').removeClass('input-group-error')

})

/** Delete option value */
$(document).on('click', '.delval', function () {

  var optionid = $(this).attr('data-id');

  var optionnewid = $(this).attr('data-newid');

  var fieldid = $('#fieldid').attr('data-fieldid');

  var newfieldid = $('#fieldid').attr('data-newfieldid');

  const index = fiedlvalue.findIndex(obj => {

    return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

  });

  if (index >= 0) {

    const optionindex = fiedlvalue[index].OptionValue.findIndex(obj => {

      return obj.Id == optionid && obj.NewId == optionnewid;

    });

    if (fiedlvalue[index].OptionValue[optionindex].Id != 0) {

      deleteoption.push(fiedlvalue[index].OptionValue[optionindex])
    }

    fiedlvalue[index].OptionValue.splice(optionindex, 1)

  }

  $(this).parents('.addedOption').remove();

})

/** Delete FieldId Confirmation Popup */
$(document).on('click', '.field-delete', function () {

  $('#centerModal1').modal('show')

  $('.deltitle').text(languagedata?.Channell?.deltitle)

  $('.deldesc').text(languagedata?.Channell?.delcontent)

  $('.delname').text($(this).siblings('.heading-three').text())

  $('#delid1').attr('data-id', $(this).parents('.fieldapp').attr('data-fieldid')).attr('data-newid', $(this).parents('.fieldapp').attr('data-newfieldid'))

  $('#delid1').removeAttr('data-sectionnewid').removeAttr('data-sectionid');

})

/** Delete field */
$(document).on('click', '#delid1', function () {

  var fieldid = $(this).attr('data-id');

  var newfieldid = $(this).attr('data-newid');

  var sectionid = $(this).attr('data-sectionid') !== undefined ? $(this).attr('data-sectionid') : 0;

  var newsectionid = $(this).attr('data-sectionnewid') !== undefined ? $(this).attr('data-sectionnewid') : 0

  if (sectionid != 0 || newsectionid != 0) {

    const sectionindex = sections.findIndex(obj => {

      return obj.SectionId == sectionid && obj.SectionNewId == newsectionid;

    })
    console.log(sectionindex, "sectionindex")
    if (sectionindex >= 0) {

      if (sections[sectionindex].SectionId != 0) {

        deletesecion.push(sections[sectionindex])

      }

      sections.splice(sectionindex, 1)

    }

    for (let x in fiedlvalue) {

      if (fiedlvalue[x].SectionId == sectionid && fiedlvalue[x].SectionNewId == newsectionid) {

        if (fiedlvalue[x].FieldId != 0) {

          console.log(fiedlvalue[x], "sdfdfdfdfd")

          deletefields.push(fiedlvalue[x])

        }

        fiedlvalue.splice(x, 1)

      }

    };

    $('#section-' + sectionid + newsectionid).remove();

  } else {

    const index = fiedlvalue.findIndex(obj => {

      return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

    });

    if (index >= 0) {

      deletefields.push(fiedlvalue[index])

      fiedlvalue.splice(index, 1)

      const index1 = fiedlvalue.findIndex(obj => {

        return obj.MasterFieldId == RelationalMember

      });

      if (index1 == -1) {

        $("#relational-member").draggable("enable");

        $("#relational-member").css('opacity', '1');
      }

    }

    $('#fieldapp' + fieldid.toString() + newfieldid.toString()).remove();

  }

  $('#delcancel1').trigger('click');

  $('#field-properties').hide();

  console.log(deletefields, "fieldvaluex")
})

/** Channel Save Ajax */
$(document).on('click', '.channelsave', function () {

  flg = FieldValidation()

  flg1 = CateogryValidation()

  if (!flg || !flg1) {

    console.log(flg, flg1);

    return false

  }

  var name = $('#channelname').val();

  var desc = $('#channeldesc').val();

  SortFieldsandSection()

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

})

/**section select */
$(document).on('click', '.section-fields-content', function () {

  $('.section-fields-content').removeClass('active');

  $(this).addClass('active');
})

/** input */
$(document).on('click', '.section-title', function () {

  $(this).parent().hide();

  $(this).parent().siblings('.sectioninput').show();

})

/**section delete */
$(document).on('click', '.sectiondel', function () {

  $('#centerModal1').modal('show')

  $('.deltitle').text(languagedata?.Channell?.delsectitle)

  $('.deldesc').text(languagedata?.Channell?.delseccontent)

  $('.delname').text($(this).siblings('input').val());

  $('#delid1').attr('data-sectionnewid', $(this).parents('.section-fields-content').attr('data-newid'))

  $('#delid1').attr('data-sectionid', $(this).parents('.section-fields-content').attr('data-id'))

})

/**section name keyup object save */
$(document).on('keyup', '.sectionname', function () {

  var sectionnewid = $(this).parents('.section-fields-content').attr('data-newid');

  var sectionid = $(this).parents('.section-fields-content').attr('data-id') == undefined ? 0 : $(this).parents('.section-fields-content').attr('data-id');

  const sectionindex = sections.findIndex(obj => {

    return obj.SectionId == sectionid && obj.SectionNewId == sectionnewid;

  })

  if (sectionindex >= 0) {

    sections[sectionindex]['SectionName'] = $(this).val();
  }


})

$(document).on('keyup', '.chanfield', function () {

  Validationcheck()

})

function Validationcheck() {

  if ($('#channelname').hasClass('error')) {

    $('#cname').addClass('input-group-error');

  } else {

    $('#cname').removeClass('input-group-error');
  }

  if ($('#channeldesc').hasClass('error')) {

    $('#cdesc').addClass('input-group-error');

  } else {

    $('#cdesc').removeClass('input-group-error');
  }
}

//**description focus function */
var Desc = document.getElementById('channeldesc');
var inputGroup = document.querySelectorAll('.input-group');

Desc.addEventListener('focus', () => {

  Desc.closest('.input-group').classList.add('input-group-focused');
});
Desc.addEventListener('blur', () => {
  Desc.closest('.input-group').classList.remove('input-group-focused');
});


$('#Sortsection').bind('DOMSubtreeModified', function () {

  SortFieldsandSection()

});


function FieldValidation() {

  console.log(fiedlvalue, "validation")

  // if (fiedlvalue.length == 0) {

  //   notify_content = `<div class="toast-msg dang-red" ><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span> ` + languagedata?.Channell?.fieldselcterr + `</span></div>`;
  //   $(notify_content).insertBefore(".header-rht");
  //   setTimeout(function () {
  //     $('.toast-msg').fadeOut('slow', function () {
  //       $(this).remove();
  //     });
  //   }, 5000); // 5000 milliseconds = 5 seconds

  //   return false

  // }
  if (fiedlvalue != null) {
    for (let x of fiedlvalue) {

      if (x['FieldName'] == "") {

        $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

        $('#fieldname-error').show();

        $('#fieldname-error').parents('.input-group').addClass('input-group-error')

        console.log("fieldname empty");

        return false

      }

      if (x["MasterFieldId"] == 3) {

        if (x['Url'] == "") {

          $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

          $('#url-error').show();

          $('#url-error').parents('.input-group').addClass('input-group-error')

          console.log("url empty");

          return false

        }



      } else if (x["MasterFieldId"] == 4) {

        var flg = true
        var flg2 = true

        if (x['DateFormat'] == "") {

          $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

          $('#date-error').show();

          $('#date-error').parent('.input-group').addClass('input-group-error')

          console.log("dateformat empty");

          flg = false

        }

        if (x['TimeFormat'] == "") {

          $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

          console.log("timeformat empty");

          $('#time-error').show();

          $('#time-error').parents('.input-group').addClass('input-group-error')

          flg2 = false
        }

        if (!flg || !flg2) {

          console.log("11");

          return false

        }


      } else if (x["MasterFieldId"] == 5 || x["MasterFieldId"] == 9 || x["MasterFieldId"] == 10) {

        if (x['OptionValue'].length == 0) {

          $('#select-error').show();

          $('#select-error').parents('.input-group').addClass('input-group-error')

          $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

          console.log("options empty");

          return false

        }

      } else if (x["MasterFieldId"] == 6) {

        if (x['DateFormat'] == "") {

          $('#fieldapp' + x["FieldId"].toString() + x["NewFieldId"].toString()).trigger('click');

          $('#date-error').parents('.input-group').addClass('input-group-error')

          $('#date-error').show();

          console.log("dateformat empty");

          return false

        }

      }

    }
  }


  return true
}

function CateogryValidation() {

  if (SelectedCategoryValue.length == 0) {

    notify_content = `<div class="toast-msg dang-red"><a id="cancel-notify" ><img src="/public/img/x-black.svg" alt="" class="rgt-img" /></a><img src="/public/img/danger-group-12.svg" alt="" class="left-img" /><span> ` + languagedata?.Channell?.selectcatvalid + ` </span></div>`;
    $(notify_content).insertBefore(".header-rht");
    setTimeout(function () {
      $('.toast-msg').fadeOut('slow', function () {
        $(this).remove();
      });
    }, 5000); // 5000 milliseconds = 5 seconds

    return false
  }

  return true
}

$(document).on('keyup', '.fieldinput', function () {

  var value = $(this).val();

  $('#title').children('h3').text(value);

  if (value != "") {

    $(this).parents('.input-group').removeClass('input-group-error')

    $(this).siblings('label').hide();

  } else {

    $(this).parents('.input-group').addClass('input-group-error')

    $(this).siblings('label').show();
  }

})

$(document).on('click', ".propdrop", function () {

  var value = $(this).parents('.dropdown-menu').siblings('.fidvalinput').val()

  if (value != "") {

    $(this).parents('.dropdown-menu').siblings('label.error').hide();

    $(this).parents('.input-group').removeClass('input-group-error');


  }


})

$(document).ready(function () {

  $('.categorypdiv').each(function () {

    var length = $(this).children('.choose-cat-list-col').children('.categoryname').length - 1

    $(this).children('.choose-cat-list-col').children('.categoryname').each(function (index) {

      if (length == index) {

        $(this).next().remove();
      }

    })

  })

  $('.selectedcategorydiv').each(function () {

    var length = $(this).children('.choose-cat-list-col').children('.categoryname').length - 1

    $(this).children('.choose-cat-list-col').children('.categoryname').each(function (index) {

      if (length == index) {

        $(this).next().remove();
      }

    })

  })

})


/**Checked */
$(document).on('click', '#Check2', function () {

  if ($(this).is(':checked')) {

    $(this).prop('checked', true);

  } else {

    $(this).prop('checked', false);

  }

  var fieldid = parseInt($('#fieldid').attr('data-fieldid'));

  var newfieldid = parseInt($('#fieldid').attr('data-newfieldid'));

  const index = fiedlvalue.findIndex(obj => {

    return obj.FieldId == fieldid && obj.NewFieldId == newfieldid;

  });

  console.log(index, $(this).is(':checked'));

  if (fiedlvalue.length == 0 || index < 0) {

    const obj = CreatePropertiesObjec()

    obj.MasterFieldId = parseInt($('#fieldid').attr('data-masterfieldid'));

    fiedlvalue.push(obj)

  } else {

    fiedlvalue[index].Mandatory = $(this).is(':checked') == true ? 1 : 0;

  }

})

$(document).on('keyup', '#foptname', function () {

  $("#select-error").hide();

  if ($(this).val() != "") {

    $("#select-error1").hide();

    $('#select-error1').parents('.input-group').removeClass('input-group-error');

  } else {
    $("#select-error1").show();

    $('#select-error1').parents('.input-group').addClass('input-group-error');
  }
})

$('#channelcreate').on('keyup keypress', function (e) {
  var keyCode = e.keyCode || e.which;
  if (keyCode === 13) {
    e.preventDefault();
    return false;
  }
});


// $("#relational-member").draggable("disable");