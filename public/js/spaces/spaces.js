var languagedata
var categoryIds = [];
var spaces = [];
var arraydatacheck =[]
var sdata
var cateid

$(document).ready(async function () {

    var languagepath = $('.language-group>button').attr('data-path')
  
    await $.getJSON(languagepath, function (data) {
        
        languagedata = data
    })


   $('.categorydropdown').each(function(){

    var length = $(this).children('.categorysname').length-1

      $(this).children('.categorysname').each(function(index){

        if(length == index ){

          $(this).next().remove();
        }

      })
    })

    $(".category").each(function(i, data){
       
     if (data.length > 100){

       return truncate(data, 58) + "...";
      }
      return data;
  
    })

});

$(document).keydown(function(event) {
    if (event.ctrlKey && event.key === '/') {
        $(".search").focus().select();
    }
  });

function truncate(str, no_words) {
    return str.split(" ").splice(0, no_words).join(" ");
  }


$("#cancel-btn").on('click', function () {

    $("#rightModal").hide()
    $('#spaceform')[0].reset();
    $('#spimagehide').attr("src", "")
    $("#spacedel-img").hide();
    $("#save").show();
    $("#update").hide();
    $("label.error").hide();
    $('#spaceform').attr("action", "/spaces/createspace")
    $('#spcname').removeClass('input-group-error');
    $('#spcdes').removeClass('input-group-error');
    $('#grbcat').removeClass('input-group-error')
    $("#catdropdown").removeClass('input-group-error')
    $("#spacename-error").hide()
    $("#spacedescription-error").hide()
})


// dropdown filter input box search
$("#searchcategory").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".channels-list-row button").each(function (index, element) {
    var title = $(element).find('span').text().toLowerCase()

      if (title.includes(keyword)) {
            $(element).show()
        } else {
            $(element).hide()

        }
    })
  
})

// media image insert

// $(document).on('click',".file>img",function () {
//     var src = $(this).find('img').attr("src");
//     var data = $("#spimage").val(src)
//     var data1 = $("#spimagehide").attr("src", src);
    
//     if (data1 != "") {
//         $(".heading-three").hide();
//         $("#browse").hide();
//         $("#spimagehide").attr("src", src).show()
//         $("#spacedel-img").show();
//         $("#addnewimageModal").modal('hide')

//     }
//     $("#rightModal").modal('show')

// })
 

$(document).on("click","#mediamodalclose",function (){
    $("#rightModal").modal('show')

})

// add space
$("#addspace , #clickadd").click(function () {
    $('form[name="newspace"]').attr("action", "/spaces/createspace");
    $("#title").text(languagedata?.Spaces?.newspace)

    // input values
    $("#spacename").val("")
    $("#spacedescription").val("")
    $("#spimage").val("")
    $("#spimagehide").attr("src", "");
    $("#triggerId").html("")
    $("#triggerId").text(languagedata?.Spaces?.slectcatgory)
    $("#spacename-error").hide()
    // btns
    $("#uplaodimg").show();
    $("#browse").show();
    $("#spacedel-img").hide()
    $("#save").show();
    $("#clone").hide();
    $("#update").hide();
    $("#spacename-error").hide()
    $("#spacedescription-error").hide()

    $(".input-group").removeClass("input-group-error")

})



function refreshdiv(){
    $('.choose-rel-article').load(location.href + ' .choose-rel-article');
}
$(document).on('click', '#newdd-input', function () {
    $(this).siblings('.dd-c').css('display', 'block')
})

$('#addnewspaceModal').on('click', function (event) {
    if ($('.dd-c').css('visibility') == 'visible' && !$(event.target).is('#newdd-input') && !$(event.target).is('.dd-c')) {
      
    }
})

// Edit functionaltiy
$("body").on("click", "#edit", function () {

    $(".input-group").removeClass("input-group-error")
    $("#spacename-error").hide()
    $("#spacedescription-error").hide()

    // $("#catlist").css("color","#112D55")

    $("#uplaodimg").hide()
    $("#browse").hide()

    var url=window.location.search
    const urlpar= new URLSearchParams(url)
    pageno = urlpar.get('page')
    $("#spacepageno").val(pageno)
    $("#categoryspan").hide()
    $("#spaceform").attr("name", "editspace")
    var data = $(this).attr("data-id");
    var id = $(this).attr("data-cid");
    $("#title").text(languagedata?.Spaces?.updatespace)
    $("#update").show()
    $("#save").hide()
    $("#clone").hide();

    $('form[name="editspace"]').attr("action", "/spaces/updatespace");
    $("#id").val(data);
    $("#catiddd").val(id)

    var details = $(this).parents(".spaceCard-btm");
    var imgdetails = $(this).parents(".spaceCard-child")
    var spimg = imgdetails.find("img").attr("src")
    var spcat = details.find("#spacecategory").html()
    var spname = details.find("#spacecontentname").text()
    var spdesc = details.find("#spacecontent").val()

    $("#uplaodimg").show();
        $("#browse").show();
    $('#spimage').val(spimg)
    $("#spimagehide").attr("src", spimg)
    if (spimg != "") {
        $("#uplaodimg").hide();
        $("#browse").hide();
        $("#spacedel-img").show();

    }else{
        $("#spacedel-img").hide()
    }
    // if ($("#spimagehide").attr("src") === "/public/img/space-default-image.png") {
    //     $("#uplaodimg").hide();
        // $("#browse").hide();
    //     $("#spimagehide").attr("src", "");
    // } else {
    //      $("#uplaodimg").hide();
        // $("#browse").hide();
    //     $("#spimage").show();
    //     $("#spacedel-img").show();
    //     // $("#catdel-img").show();
    // }

    $("#triggerId").html(spcat)
    $("#spacename").val(spname.trim())
    $("#spacedescription").val(spdesc.trim())
    // $("#spimage").val(spimg)

    if (spcat != "") {
        $("input[name=spacecategoryvalue]").rules("remove", "required")

    }

    $('#triggerId').children('li').children('a').css('color','#112D55')

});


// $('#rightModal').on('hidden.bs.modal', function (event) {
//     $('#spaceform')[0].reset();
//     $('#spimagehide').attr("src", "")
//     $("#spacedel-img").hide();
//     $("#save").show();
//     $("#update").hide();
//     $("label.error").hide();
//     $('#spaceform').attr("action", "/spaces/createspace")
//     $('#spcname').removeClass('input-group-error');
//     $('#spcdes').removeClass('input-group-error');
//     $('#grbcat').removeClass('input-group-error')
//     $("#dropdown").removeClass('input-group-error')
//     $('.modal-backdrop').remove()
// })



// delete popup 
// $(document).on('click', '#deletespace', function () {
//     var id = $(this).attr("data-id");
//     $.ajax({
//         url: "/spaces/deletepopup",
//         type: "GET",
//         dataType: "json",
//         data: { "id": id, csrf: $("input[name='csrf']").val() },
//         success: function (results) {
//             if (results) {
//                 $('#content').text("Are You Sure You Want To Delete This Space");
//                 $('#delid').show();
//                 $('#delid').parent('#delete').attr('href', '/spaces/deletespace?id=' + id);
//                 $('#btn3').text(languagedata.no);

//             } else {
//                 $('#content').text(languagedata.delroleinvalid);
//                 $('#delid').parent('#delete').attr('href', '');
//                 $('#delid').hide();
//                 $('#btn3').text(languagedata.ok)

//             }
//         }
//     });
// });

// delete popup 
$(document).on('click', '#deletespace', function () {
    var id = $(this).attr("data-id");
    $(".deltitle").text(languagedata?.Spaces?.deltitle)
    var details = $(this).parents(".spaceCard-btm");
    var spname = details.find("#spacecontentname").text()
    $('.delname').text(spname.trim())
    $('#content').text(languagedata?.Spaces?.delspace);
    $('#delid').show();
    $('#delete').attr('href', '/spaces/deletespace?id=' + id);
    $("a").attr("id", "")
    $('#delcancel').text(languagedata.no);
})

// Clone functionaltiy
$("body").on("click", "#clonebtn", function () {

    $("#spacename-error").hide()
    $("#spacedescription-error").hide()
    // $("#uplaodimg").hide()
    // $("#browse").hide()

    $(".space-path-list").css("color","#112D55")

    
    $("#categoryspan").hide()
    $("#spaceform").attr("name", "clonespace")
    var data = $(this).attr("data-id");
    var id = $(this).attr("data-cid")
    $("#title").text(languagedata?.Spaces?.clonespace)
    $("#clone").show()
    $("#update").hide()
    $("#save").hide()
    $('form[name="clonespace"]').attr("action", "/spaces/clonespace");
    $("#id").val(data);
    $("#catiddd").val(id)

    var details = $(this).parents(".spaceCard-btm");
    var imgdetails = $(this).parents(".spaceCard-child")
    var spimg = imgdetails.find("img").attr("src")
    var spcat = details.find("#spacecategory").html()
    
    var spname = details.find("#spacecontentname").text()
    var spdesc = details.find("#spacecontentdesc").text()

    $("#spimagehide").attr("src", spimg)
    $("#spimage").val(spimg)
    if (spimg != "") {
        $("#uplaodimg").hide();
        $("#browse").hide();
        $("#spacedel-img").show();

    }else{
        $("#spacedel-img").hide()
    }
    // if ($("#spimage").attr("src") === "/public/img/space-default-image.png") {
    //    $("#uplaodimg").hide();
        // $("#browse").hide();
    //     $("#spimage").attr("src", "");
    // } else {
    //     $("#uplaodimg").hide();
        // $("#browse").hide();
    //     $("#spimage").show();
    //     $("#spacedel-img").show();
    //     // $("#catdel-img").show();
    // }

    $("#triggerId").html(spcat)
    $("#spacename").val(spname.trim())
    $("#spacedescription").val(spdesc.trim())
    // $("#spacecategoryid").val(id)
    // $("#spimagehide").val(spimg)
    // if (spcat != "") {

    //     $("input[name=spacecategoryid]").rules("remove", "required")
    // }

});
$(document).on('click', '.fcheck', function () {
    var categoryId = $(this).closest('.newck-group').find('.para:last').data('id').toString();

    var categoryId1 = $(this).attr("data-check").toString();
      
    var isChecked = $(this).prop('checked');
    if (isChecked) {
        if (categoryId !== "" && categoryId1 !== "") {
            categoryIds.push(categoryId);
            arraydatacheck.push(categoryId1);

        }
    } else {
        categoryIds = categoryIds.filter(function (id) {
            return id != categoryId;
        }); 
        arraydatacheck = arraydatacheck.filter(function (cid) {
            return cid != categoryId1;
        });

    }


    var keyword = $('#spacessearch').val();
    if (categoryIds.length > 0 && arraydatacheck.length > 0) {
        $("#filterid").val(categoryIds);
        window.location.href = "/spaces?keyword=" + keyword + "&categoryid=" + categoryIds.join(",")+ "&check=" + arraydatacheck.join(",");
    } else if (keyword != "") {
        window.location.href = "/spaces?keyword=" + keyword
    } else {
        $("#filterid").val("");
        window.location.href = "/spaces";
    }
});


$('.delete-flag').click(function () {
    $('input[name=spaceimagepath]').val("")
    $('#spimagehide').attr('src', '')
    $('#spimage').val('')
    $(this).siblings('h3').show()
    $('#spacedel-img').hide()
    
    $('#browse').show()
    $('#spimage').val('')
})

$.validator.addMethod("customLength", function (value, element, params) {
    var minLength = params[0];
    var maxLength = params[1];
    return this.optional(element) || (value.length >= minLength && value.length <= maxLength);
}, $.validator.format("Please enter between {0} and {1} characters."));


$("#save").click(function () {

    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        id = $("#id").val()
       
        $.ajax({
            url:"/spaces/checkspacename",
            type:"POST",
            async:false,
            data:{"spacename":value,"id":id,csrf:$("input[name='csrf']").val()},
            datatype:"json",
            caches:false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim()!="true"
    })

    $('#spaceform').validate({
        ignore:[],
        rules: {
            spacename: {
                required: true,
                space: true,
                duplicatename:true
            },
            spacedescription: {
                required: true,
                space: true,
                customLength: [1, 250]

            },
            catiddd: {
                required: true
            }
        },
        messages: {
            spacename: {
                required: "* " + languagedata.Spaces.spacenamevalid,
                space: "* " + languagedata.Spaces.spacergx,
                duplicatename:"*" + languagedata.Spaces.namevailderr
            },
            spacedescription: {
                required: "* " + languagedata.Spaces.spacedescvalid,
                space: "* " + languagedata.Spaces.spacergx,
                customLength: "*" + languagedata?.Permission?.descriptionchat,
            },
            catiddd: {
                required: "* " + languagedata.Spaces.spacecategoryvalid,

            },

        }
    });

    var formcheck = $("#spaceform").valid();
    if (formcheck == true) {
        $('#spaceform')[0].submit();
    } else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })

        if ($('#catiddd-error').css('display') !== 'none') {
            $('#grbcat').addClass('input-group-error')
            // $("#catdropdown").addClass('input-group-error')

        }
        else {
            $('#grbcat').removeClass('input-group-error')
            // $("#catdropdown").removeClass('input-group-error')

        }
                             
    }

    return false

})

$("#update").click(function () {
    
    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        id = $("#id").val()
       
        $.ajax({
            url:"/spaces/checkspacename",
            type:"POST",
            async:false,
            data:{"spacename":value,"id":id,csrf:$("input[name='csrf']").val()},
            datatype:"json",
            caches:false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim()!="true"
    })
    $('#spaceform').validate({
        ignore:[],

        rules: {
            spacename: {
                required: true,
                space: true,
                duplicatename:true
            },
            spacedescription: {
                required: true,
                space: true,
                customLength: [1, 250]

            },
            spacecategoryvalue: {
                required: true
            }
        },
        messages: {
            spacename: {
                required: "* " + languagedata.Spaces.spacenamevalid,
                space: "* " + languagedata.Spaces.spacergx,
                duplicatename:"*" + languagedata.Spaces.namevailderr
            },
            spacedescription: {
                required: "* " + languagedata.Spaces.spacedescvalid,
                space: "* " + languagedata.Spaces.spacergx,
                customLength: "*" + languagedata?.Permission?.descriptionchat,
            },
            spacecategoryvalue: {
                required: "* " + languagedata.Spaces.spacecategoryvalid,

            },

        }
    });

  
    var formcheck = $("#spaceform").valid();
    if (formcheck == true) {
        $('#spaceform')[0].submit();
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        // if ($('#catiddd').val()!==''){
        //     $('#grbcat').addClass('input-group-error')
        //     $("#catdropdown").addClass('input-group-error')

        // }else {
        //     $('#grbcat').removeClass('input-group-error')
        //     $("#catdropdown").removeClass('input-group-error')
        // }
    }

    return false
})

$("#clone").click(function () {
    jQuery.validator.addMethod("duplicatename", function (value) {

        var result;
        // id = $("#id").val()
        id=0
       
        $.ajax({
            url:"/spaces/checkspacename",
            type:"POST",
            async:false,
            data:{"spacename":value,"id":id,csrf:$("input[name='csrf']").val()},
            datatype:"json",
            caches:false,
            success: function (data) {
                result = data.trim();
            }
        })
        return result.trim()!="true"
    })

    $('#spaceform').validate({
        ignore:[],

        rules: {
            spacename: {
                required: true,
                space: true,
                duplicatename:true
            },
            spacedescription: {
                required: true,
                space: true,
                customLength: [1, 250]

            },
            spacecategoryvalue: {
                required: true
            }
        },
        messages: {
            spacename: {
                required: "* " + languagedata.Spaces.spacenamevalid,
                space: "* " + languagedata.Spaces.spacergx,
                duplicatename:"*" + languagedata.Spaces.namevailderr
            },
            spacedescription: {
                required: "* " + languagedata.Spaces.spacedescvalid,
                space: "* " + languagedata.Spaces.spacergx,
                customLength: "*" + languagedata?.Permission?.descriptionchat,
            },
            spacecategoryvalue: {
                required: "* " + languagedata.Spaces.spacecategoryvalid,

            },

        }
    });

  

    var formcheck = $("#spaceform").valid();
    if (formcheck == true) {
        $('#spaceform')[0].submit();
    }
    else {
        Validationcheck()
        $(document).on('keyup', ".field", function () {
            Validationcheck()
        })
        // if ($('#catiddd').hasClass('error')) {
        //     $("#catdropdown").addClass('input-group-error')
    
        // }
        // else {
        //     $("#catdropdown").removeClass('input-group-error')
    
        // }
       
    }

    return false
})
function checkImage() {
    const spimage = document.getElementById("spimage");
    const spacedel = document.getElementById("spacedel-img");

    // Check if the src attribute is empty
    if (spimage.src === "") {
        spacedel.style.display = "none";
    } else {
        spacedel.style.display = "block";
    }
}

// avoid the last index image

// const image = document.querySelectorAll('.scrollimg ');
// image.forEach(group => {
//     const image = group.querySelectorAll('img');
//     image[image.length - 1].style.display = 'none';
// });

const spacedes = document.getElementById('spacedescription');
const inputGroup = document.querySelectorAll('.input-group');
const opt = document.querySelectorAll('.select-options')


spacedes.addEventListener('click', () => {

    spacedes.closest('.input-group').classList.add('input-group-focused');

});

spacedes.addEventListener('blur', function () {
    spacedes.closest('.input-group').classList.remove('input-group-focused');
});


function Validationcheck() {


    if ($('#spacename').hasClass('error')) {
        $('#spcname').addClass('input-group-error');
    } else {
        $('#spcname').removeClass('input-group-error');
    }

    if ($('#spacedescription').hasClass('error')) {
        $('#spcdes').addClass('input-group-error');
    } else {
        $('#spcdes').removeClass('input-group-error');
    }
 
}

$(document).on('click','.close',function(){
    $("#spacedel-img").hide();
    $(".input-group").removeClass("input-group-error")

    $('#spcname').removeClass('input-group-error');
    $('#spcdes').removeClass('input-group-error');
    $('#grbcat').removeClass('input-group-error')
    $('#catdropdown').removeClass('input-group-error')

    $('#catiddd-error').hide()
})

// form category search in keyup

$("#searchcatlist").keyup(function () {
    var keyword = $(this).val().trim().toLowerCase()
    $(".choose-category button").each(function (index, element) {
        var title = $(element).text().toLowerCase()
        if (title.includes(keyword)) {
            $(element).show()
            $("#spacenodatafound").hide()

        } else {
            $(element).hide()
            if($('.choose-category button:visible').length==0){
                $("#spacenodatafound").show()
            }
            
        }
    })
})

// search key value empty in dropdown search

$("#triggerId").on("click",function(){

    var keyword = $("#searchcatlist").val()
    $(".choose-category button").each(function (index, element) {
        if(keyword != ""){
        $("#searchcatlist").val("")
        $("#spacenodatafound").hide()
        $(element).show()
      }else{
        $("#searchcatlist").val()
        $("#spacenodatafound").hide()
        $(element).show()
    
      }
    })
   
})


$(document).on('click','#medcancel',function () {
  $("#addnewimageModal").hide()
})

$(document).ready(function(){

    $(".desclist").each(function (index,element){

        var result= $(this).text()

        if (result?.length <= 129 ) {

             $(this).removeAttr("data-bs-original-title")
        
        }
      

    })

       
    $(".searchcategorylist").each(function (index,element){
   

        var id = $(this).attr('data-id')

    
      var newiw=  $('.catdataid').val()

      console.log(newiw,"ffff")
    //     categorylen= $('#newcat'+id).text()
     
    //     $('#newcat'+id).attr('data-bs-original-title',categorylen)
     
    //   var hidden=$(this).children('.categoryactiveid').val()
    //   console.log(hidden,"checkkk")
         
         if (id==$('.catdataid').val()){
     
             console.log("sddddddd")
     
             $(this).addClass('active')
     
         }
     
     
         })
     
     
     
     })
     
     $(".txtclr4").click(function(){
     
         $('.searchcategorylist').removeClass('active')
     
         $(".searchcategorylist").each(function (index,element){
        
     
             var id = $(this).attr('data-id')
          
            var hidden=$(this).children('.categoryactiveid').val()
     
             
              if (id==$('.categoryactiveid').val()){
        
        
                  $(this).addClass('active')
          
              }
          
          
              })


   
})

$(document).on('click','.filterdrop',function(){

    if (cateid !=  $('.catdataid').val()){

        $('.searchcategorylist').removeClass('active')
    }

})

var filtercategoryid

$(".searchcategorylist").on("click", function () {

    $(this).removeClass('active')

     cateid = $(this).attr('data-id')

    $('.catdataid').val(cateid)

    // $('.list-items-select').empty();

    var categoryname =$(this).children('.newcat').text().slice(0,-4)
           
    var fcategoryname =$(this).children('.newcat').attr('data-bs-original-title')

   let fcat =fcategoryname.slice(0,-4)

   let  categoryname1 = categoryname.replaceAll('/', '<img src="/public/img/breadcrumb-arrow.svg" style="padding:5px;" alt="">');

    var Category = $(this).find('.id')
    $.each(Category, function (index, value) {

        if (Category.length - 1) {
            filtercategoryid = $(this).attr("data-id")

        }
    })

    var data = {"categoryid":filtercategoryid }

    var getdata = LoadMoreData(data, space_count, space_array)

     space_array = getdata.array

     space_count = getdata.count
  
        $('.selected-list').show()

        $(".searchcategorylist").each(function (index,element){
   

            var id = $(this).attr('data-id')
    
        
          var newiw=  $('.catdataid').val()
    
          console.log(newiw,"ffff")
             
             if (id==$('.catdataid').val()){
         
                 console.log("sddddddd")
         
                 $(this).addClass('active')
         
             }else{
                $(this).removeClass('active')

             }
         
         
             })
         
         

        // $('.list-items-select').html('<p class="pcategory" style="display: flex" data-bs-custom-class="lms-tooltip" data-bs-toggle="tooltip" data-bs-html="true" data-bs-placement="top" data-bs-original-title="" >' + categoryname1 + '</p>')

        // $('.pcategory').attr('data-bs-original-title', fcat)
})



var keywordname

$("#searchspacename").on("keypress", function () {

    keywordname= $(this).val().trim().toLowerCase()

    var data = {"keyword":keywordname}

    if (event.key === "Enter") {

        event.preventDefault();
        
        var getdata = LoadMoreData(data, space_count, space_array)

        space_array = getdata.array
   
        space_count = getdata.count
      }

   
  
})

/*search redirect home page */

$(document).on('keyup', '#searchspacename', function () {

    if (event.key === 'Backspace') {

    if ($('.search').val() === "") {

         window.location.href = "/spaces/"
    }

}
})
    


// Scroll paganation

    // var spacelimit = 10

    // var modifier = 1
    
    // var space_offset = 1
    
    var space_array = []
    
    var space_count = 0
  
    // var scrollablediv = document.querySelector('.spaceCards')
    
    // var space_divheight = scrollablediv.scrollHeight
    
    // scrollablediv.addEventListener('scroll', () => {
    
    //   var readOnlyHeight = scrollablediv.offsetHeight
   
    //   var scrollposition = scrollablediv.scrollTop
   
    //   var currentscroll = readOnlyHeight + scrollposition
   
    //   if (currentscroll + modifier > space_divheight) {
    
    //     if (scrollposition >= 140 && scrollposition <= 180) {
    
    //       space_offset = space_offset + 1
    
    //     } 

  
    //     if (space_array.length == 0 || space_array.length < space_count) {

   
    //      var data = { "space_offset": space_offset,"space_limit": spacelimit, "space_scrolldata": true }
    
    //      var getdata = LoadMoreData(data, space_count, space_array)
    
    //       space_array = getdata.array
    
    //       space_count = getdata.count
    
    //     }
    
    //   }

            
    // })
 


    function LoadMoreData(data, count, array) {

        $.ajax({
      
          url: "/spaces/list",
      
          data: data,
      
          async: false,
      
          dataType: "json",
      
          success: (result) => {

           console.log("pagnation",result);

            if (keywordname  !="" || filtercategoryid != "") {

                $(".spaceCards").html("")

                $("#searchcount").text(result.FilterCount)

            }

            if (result.FilterCount == 0){

                var nodatafound = `<div class="noData-foundWrapper nodatafound">

                <div class="empty-folder">
                    <img src="/public/img/nodatafilter.svg" alt="">
                </div>
                <h1 class="heading"> `+languagedata.oopsnodata +`</h1>
                <p class="para">              
                </p>
               
            </div>`

            $('.spaceCards').append(nodatafound)

            }
   
            // if (result.hasOwnProperty('TotalSpaceCount')) {
      
            //   count = result.TotalSpaceCount - 10
      
            // } 


            if (result.hasOwnProperty('SpaceDetails')) {
      
              for (let space of result.SpaceDetails) {
      
                array.push(space)
      
                var childs = ""
      
                if (space.CategoryNames != null) {
      
                  for (let child of  space.CategoryNames) {

                var index=space.CategoryNames.findIndex(Obj => Obj.Id == child.Id)
                       
                   if(index == space.CategoryNames.length-1){

                    childs = childs + `<span>` + child.CategoryName+` </span>`
                   }else{

                    childs = childs + `<span>` + child.CategoryName + ` </span> / `

                   }
                 }
                 
                }


            
                var single_list = `<div class="spaceCard-child">
                <a href="/spaces/`+space.Id +`" class="spacecard-link"></a>

                <div class="spaceCard-top">
                
                    <img src="`+space.ImagePath + `"  style="object-fit: fill;width: 100%;height: 100%;border: none;">
                    
                </div>

                <div class="spaceCard-btm">

                    <h3 class="heading-three lms-card-title" href="/spaces/`+space.Id +`" id="spacecontentname">
                     `+space.SpacesName+`
                    </h3>

                    <div class="card-breadCrumb">
                        <ul class="breadcrumb-container" id="spacecategory">
                            <li> <a href="/spaces/`+space.Id +`" class="para-light space-path-list" style="display: flex;gap: 4px;"> `+childs +` </a></li>
                        </ul>
                    </div>

                    <a href="/spaces/`+space.Id +`" class="para-extralight desclist" id="spacecontentdesc"
                            data-bs-custom-class="lms-tooltip" data-bs-toggle="tooltip" data-bs-html="true"
                            data-bs-placement="top"  title="`+ space.SpaceFullDescription+`">
                       `+space.SpacesDescription+`
                       <input type="hidden" value="`+space.SpaceFullDescription +`" name="spacecontent" id="spacecontent">

                    </a>

                    <div class="read-time">
                           
                    <h6 class="para-extralight">Updated: `+space.ModifiedDate +`</h6>
                     <p class="xmintag${space.Id}">Reading Time: `+space.ReadTime + `</p> 
                      </div>



                    <div class="card-buttonOption">
                    <a href="/spaces/`+space.Id+`"> <button class="btn-reg btn-lg primary " id="page${space.Id}">   
                            <img src="/public/img/add.svg" alt="" /> Add Page
                            </button></a>

                        <div class="card-option">
                            <div class="btn-group language-group">
                                <button type="button" class="dropdown-tog" data-bs-toggle="dropdown"
                                    aria-expanded="true">
                                    <img src="/public/img/card-option.svg" alt="">
                                    <img src="/public/img/card-option-bg.svg" alt="">
                                </button>
                                <ul class="dropdown-menu dropdown-menu-end " data-popper-placement="bottom-end"
                                    style="position: absolute; inset: 0px 0px auto auto; margin: 0px; transform: translate(0px, 22px);">
                                    <li><button class="dropdown-item" id="clonebtn" type="button"
                                            data-cid="`+space.PageCategoryId+`" data-id=" `+space.Id +` " data-bs-toggle="modal"
                                            data-bs-target="#rightModal"> <span><img src="/public/img/copy.svg"
                                                    alt=""></span> Clone </button></li>
                                    <li><button class="dropdown-item" id="edit" type="button"
                                    data-cid="`+space.PageCategoryId+`" data-id=" `+space.Id +`" data-bs-toggle="modal"
                                    data-bs-target="#rightModal"> <span><img src="/public/img/edit.svg"
                                                    alt=""></span> Edit </button></li>
                                    <li><button class="dropdown-item" id="deletespace" data-id="`+space.Id +`"
                                            type="button" data-bs-toggle="modal" data-bs-target="#centerModal">
                                            <span><img src="/public/img/delete.svg" alt=""></span> Delete </button>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>

                </div>

              </div> `
      
                $('.spaceCards').append(single_list)

                if(space.FullSpaceAccess==true){
                   
                    $(`#page${space.Id}`).text('Manage Page')

                    $(`.xmintag${space.Id}`).show()
 
                     }else{
 
                         $(`#page${space.Id}`).text('Add Page')

                         $(`.xmintag${space.Id}`).hide()
 
                     }
             }
      
            }
          }
        })
      
        return { "array": array, "count": count }
}

// dropdown category get id

$(".categorydropdown").click(function(){
  
    var id 

    var Category = $(this).find('.id')

    $.each(Category, function (index, value) {

        if (Category.length - 1) {
          id = $(this).attr("id")
          $("#catiddd").val(id)
        }
    })

     if (id == 0) {
            $("#grbcat").addClass('input-group-error')

        }else {
            $("#grbcat").removeClass('input-group-error')
            $("#catiddd-error").css('display', 'none')
    }
    $(this).parents('.dropdown-menu').siblings('a').html($(this).html())
})

$(document).on('click','#removecategory',function(){

    window.location.href="/spaces/"

  }) 


  $(document).on('click','#removecategory',function(){

    window.location.href="/spaces/"

  })